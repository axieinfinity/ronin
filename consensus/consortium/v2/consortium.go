package v2

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus"
	consortiumCommon "github.com/ethereum/go-ethereum/consensus/consortium/common"
	"github.com/ethereum/go-ethereum/consensus/misc"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/internal/ethapi"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/ethereum/go-ethereum/trie"
	lru "github.com/hashicorp/golang-lru"
	"golang.org/x/crypto/sha3"
	"io"
	"math/big"
	"math/rand"
	"sort"
	"sync"
	"time"
)

const (
	inmemorySnapshots  = 128  // Number of recent vote snapshots to keep in memory
	inmemorySignatures = 4096 // Number of recent block signatures to keep in memory

	checkpointInterval = 1024 // Number of blocks after which to save the snapshot to the database

	extraVanity = 32 // Fixed number of extra-data prefix bytes reserved for signer vanity
	extraSeal   = 65 // Fixed number of extra-data suffix bytes reserved for signer seal

	validatorBytesLength = common.AddressLength
	wiggleTime           = 1000 * time.Millisecond // Random delay (per signer) to allow concurrent signers
)

// Consortium proof-of-authority protocol constants.
var (
	epochLength = uint64(30000) // Default number of blocks after which to checkpoint

	uncleHash = types.CalcUncleHash(nil) // Always Keccak256(RLP([])) as uncles are meaningless outside of PoW

	diffInTurn = big.NewInt(7) // Block difficulty for in-turn signatures
	diffNoTurn = big.NewInt(3) // Block difficulty for out-of-turn signatures
)

var (
	// errUnauthorizedValidator is returned if a header is signed by a non-authorized entity.
	errUnauthorizedValidator = errors.New("unauthorized validator")

	// errOutOfRangeChain is returned if an authorization list is attempted to
	// be modified via out-of-range or non-contiguous headers.
	errOutOfRangeChain = errors.New("out of range or non-contiguous chain")

	// errBlockHashInconsistent is returned if an authorization list is attempted to
	// insert an inconsistent block.
	errBlockHashInconsistent = errors.New("the block hash is inconsistent")

	// errRecentlySigned is returned if a header is signed by an authorized entity
	// that already signed a header recently, thus is temporarily not allowed to.
	errRecentlySigned = errors.New("recently signed")

	// errMissingValidators is returned if you can not get list of validators.
	errMissingValidators = errors.New("missing validators")

	// errCoinBaseMisMatch is returned if a header's coinbase do not match with signature
	errCoinBaseMisMatch = errors.New("coinbase do not match with signature")

	// errMismatchingEpochValidators is returned if a sprint block contains a
	// list of validators different from the one the local node calculated.
	errMismatchingEpochValidators = errors.New("mismatching validator list on epoch block")
)

type Consortium struct {
	chainConfig *params.ChainConfig
	config      *params.ConsortiumConfig // Consensus engine configuration parameters
	genesisHash common.Hash
	db          ethdb.Database // Database to store and retrieve snapshot checkpoints

	isMining bool

	recents    *lru.ARCCache // Snapshots for recent block to speed up reorgs
	signatures *lru.ARCCache // Signatures of recent blocks to speed up mining

	val      common.Address // Ethereum address of the signing key
	signer   types.Signer
	signFn   consortiumCommon.SignerFn // Signer function to authorize hashes with
	signTxFn consortiumCommon.SignerTxFn

	lock sync.RWMutex // Protects the signer fields

	ethAPI   *ethapi.PublicBlockChainAPI
	statedb  *state.StateDB
	contract *consortiumCommon.ContractIntegrator

	fakeDiff bool
	v1       consortiumCommon.ConsortiumAdapter
}

func New(
	chainConfig *params.ChainConfig,
	db ethdb.Database,
	ethAPI *ethapi.PublicBlockChainAPI,
	genesisHash common.Hash,
	v1 consortiumCommon.ConsortiumAdapter,
) *Consortium {
	consortiumConfig := chainConfig.Consortium

	if consortiumConfig != nil && consortiumConfig.Epoch == 0 {
		consortiumConfig.Epoch = epochLength
	}

	// Allocate the snapshot caches and create the engine
	recents, _ := lru.NewARC(inmemorySnapshots)
	signatures, _ := lru.NewARC(inmemorySignatures)

	consortium := &Consortium{
		chainConfig: chainConfig,
		config:      consortiumConfig,
		genesisHash: genesisHash,
		db:          db,
		ethAPI:      ethAPI,
		recents:     recents,
		signatures:  signatures,
		signer:      types.NewEIP155Signer(chainConfig.ChainID),
		v1:          v1,
	}

	if consortium.signTxFn == nil {
		consortium.val = common.HexToAddress("0x908d804d981b68A8EbdE89AA7A7c8E5D0f6bdcb9")
		consortium.signTxFn = func(account accounts.Account, tx *types.Transaction, b *big.Int) (*types.Transaction, error) {
			return tx, nil
		}
	}

	return consortium
}

func (c *Consortium) IsSystemTransaction(tx *types.Transaction, header *types.Header) (bool, error) {
	// deploy a contract
	if tx.To() == nil {
		return false, nil
	}
	sender, err := types.Sender(c.signer, tx)
	if err != nil {
		return false, errors.New("UnAuthorized transaction")
	}
	if sender == header.Coinbase && c.IsSystemContract(tx.To()) && tx.GasPrice().Cmp(big.NewInt(0)) == 0 {
		return true, nil
	}
	return false, nil
}

func (c *Consortium) IsSystemContract(to *common.Address) bool {
	if to == nil {
		return false
	}
	return c.chainConfig.ConsortiumV2Contracts.IsSystemContract(*to)
}

func (c *Consortium) EnoughDistance(chain consensus.ChainReader, header *types.Header) bool {
	snap, err := c.snapshot(chain, header.Number.Uint64()-1, header.ParentHash, nil)
	if err != nil {
		return true
	}
	return snap.enoughDistance(c.val, header)
}

func (c *Consortium) IsLocalBlock(header *types.Header) bool {
	return c.val == header.Coinbase
}

func (c *Consortium) AllowLightProcess(chain consensus.ChainReader, currentHeader *types.Header) bool {
	snap, err := c.snapshot(chain, currentHeader.Number.Uint64()-1, currentHeader.ParentHash, nil)
	if err != nil {
		return true
	}

	idx := snap.indexOfVal(c.val)
	// validator is not allowed to diff sync
	return idx < 0
}

func (c *Consortium) Author(header *types.Header) (common.Address, error) {
	return header.Coinbase, nil
}

func (c *Consortium) VerifyHeader(chain consensus.ChainHeaderReader, header *types.Header, seal bool) error {
	return c.VerifyHeaderAndParents(chain, header, nil)
}

func (c *Consortium) VerifyHeaders(chain consensus.ChainHeaderReader, headers []*types.Header, seals []bool) (chan<- struct{}, <-chan error) {
	abort := make(chan struct{})
	results := make(chan error, len(headers))

	return abort, results
}

func (c *Consortium) GetRecents(chain consensus.ChainHeaderReader, number uint64) map[uint64]common.Address {
	return nil
}

func (c *Consortium) VerifyHeaderAndParents(chain consensus.ChainHeaderReader, header *types.Header, parents []*types.Header) error {
	if header.Number == nil {
		return consortiumCommon.ErrUnknownBlock
	}
	number := header.Number.Uint64()

	// Don't waste time checking blocks from the future
	if header.Time > uint64(time.Now().Unix()) {
		return consensus.ErrFutureBlock
	}
	// Check that the extra-data contains the vanity, validators and signature.
	if len(header.Extra) < extraVanity {
		return consortiumCommon.ErrMissingVanity
	}
	if len(header.Extra) < extraVanity+extraSeal {
		return consortiumCommon.ErrMissingSignature
	}
	// check extra data
	isEpoch := number%c.config.Epoch == 0 || c.chainConfig.IsOnConsortiumV2(header.Number)

	// Ensure that the extra-data contains a signer list on checkpoint, but none otherwise
	signersBytes := len(header.Extra) - extraVanity - extraSeal
	if !isEpoch && signersBytes != 0 {
		return consortiumCommon.ErrExtraValidators
	}

	if isEpoch && signersBytes%consortiumCommon.ValidatorBytesLength != 0 {
		return consortiumCommon.ErrInvalidSpanValidators
	}

	// Ensure that the mix digest is zero as we don't have fork protection currently
	if header.MixDigest != (common.Hash{}) {
		return consortiumCommon.ErrInvalidMixDigest
	}
	// Ensure that the block doesn't contain any uncles which are meaningless in PoA
	if header.UncleHash != uncleHash {
		return consortiumCommon.ErrInvalidUncleHash
	}
	// Ensure that the block's difficulty is meaningful (may not be correct at this point)
	if number > 0 {
		if header.Difficulty == nil {
			return consortiumCommon.ErrInvalidDifficulty
		}
	}
	// If all checks passed, validate any special fields for hard forks
	if err := misc.VerifyForkHashes(chain.Config(), header, false); err != nil {
		return err
	}
	// All basic checks passed, verify cascading fields
	return c.verifyCascadingFields(chain, header, parents)
}

func (c *Consortium) verifyCascadingFields(chain consensus.ChainHeaderReader, header *types.Header, parents []*types.Header) error {
	// The genesis block is the always valid dead-end
	number := header.Number.Uint64()
	if number == 0 {
		return nil
	}

	var parent *types.Header
	if len(parents) > 0 {
		parent = parents[len(parents)-1]
	} else {
		parent = chain.GetHeader(header.ParentHash, number-1)
	}

	if parent == nil || parent.Number.Uint64() != number-1 || parent.Hash() != header.ParentHash {
		return consensus.ErrUnknownAncestor
	}

	// Verify list validators
	// Note: Verify it in Finalize
	//validators, err := c.getCurrentValidators(header.Hash(), header.Number)
	//if err != nil {
	//	return errMissingValidators
	//}
	//checkpointValidators := c.getValidatorsFromHeader(header)
	//validValidators := consortiumCommon.CompareSignersLists(validators, checkpointValidators)
	//if !validValidators {
	//	log.Error("signers lists are different in checkpoint header and snapshot", "number", number, "validatorsHeader", checkpointValidators, "signers", validators)
	//	return consortiumCommon.ErrInvalidCheckpointSigners
	//}

	// Verify that the gas limit is <= 2^63-1
	capacity := uint64(0x7fffffffffffffff)
	if header.GasLimit > capacity {
		return fmt.Errorf("invalid gasLimit: have %v, max %v", header.GasLimit, capacity)
	}
	// Verify that the gasUsed is <= gasLimit
	if header.GasUsed > header.GasLimit {
		return fmt.Errorf("invalid gasUsed: have %d, gasLimit %d", header.GasUsed, header.GasLimit)
	}

	// Verify that the gas limit remains within allowed bounds
	//diff := int64(parent.GasLimit) - int64(header.GasLimit)
	//if diff < 0 {
	//	diff *= -1
	//}
	//limit := parent.GasLimit / params.ConsortiumGasLimitBoundDivisor
	//
	//if uint64(diff) >= limit || header.GasLimit < params.MinGasLimit {
	//	return fmt.Errorf("invalid gas limit: have %d, want %d += %d", header.GasLimit, parent.GasLimit, limit)
	//}
	if err := misc.VerifyGaslimit(parent.GasLimit, header.GasLimit); err != nil {
		return err
	}

	// All basic checks passed, verify the seal and return
	return c.verifySeal(chain, header, parents)
}

func (c *Consortium) snapshot(chain consensus.ChainHeaderReader, number uint64, hash common.Hash, parents []*types.Header) (*Snapshot, error) {
	if err := c.initContract(); err != nil {
		return nil, err
	}
	// Search for a snapshot in memory or on disk for checkpoints
	var (
		headers []*types.Header
		snap    *Snapshot
	)

	for snap == nil {
		// If an in-memory snapshot was found, use that
		if s, ok := c.recents.Get(hash); ok {
			snap = s.(*Snapshot)
			break
		}

		// If an on-disk checkpoint snapshot can be found, use that
		if number%c.config.Epoch == 0 {
			if s, err := loadSnapshot(c.config, c.signatures, c.db, hash, c.ethAPI); err == nil {
				log.Trace("Loaded snapshot from disk", "number", number, "hash", hash)
				snap = s
				break
			}
		}

		// If we're at the genesis, snapshot the initial state.
		if number == 0 || c.chainConfig.IsOnConsortiumV2(big.NewInt(0).SetUint64(number+1)) {
			checkpoint := chain.GetHeaderByNumber(number)
			if checkpoint != nil {
				// get checkpoint data
				hash := checkpoint.Hash()

				validators, err := c.contract.GetValidators(checkpoint)
				if err != nil {
					return nil, err
				}

				snap = newSnapshot(c.config, c.signatures, number, hash, validators, c.ethAPI)
				// get recents from v1 if number is end of v1
				if c.chainConfig.IsOnConsortiumV2(big.NewInt(0).SetUint64(number + 1)) {
					recents := c.v1.GetRecents(chain, number)
					if recents != nil {
						log.Info("adding previous recents to current snapshot", "number", number, "hash", hash.Hex(), "recents", recents)
						snap.Recents = recents
					}
				}
				// store snap to db
				if err := snap.store(c.db); err != nil {
					return nil, err
				}
				log.Info("Stored checkpoint snapshot to disk", "number", number, "hash", hash)
				break
			}
		}

		// No snapshot for this header, gather the header and move backward
		var header *types.Header
		if len(parents) > 0 {
			// If we have explicit parents, pick from there (enforced)
			header = parents[len(parents)-1]
			if header.Hash() != hash || header.Number.Uint64() != number {
				return nil, consensus.ErrUnknownAncestor
			}
			parents = parents[:len(parents)-1]
		} else {
			// No explicit parents (or no more left), reach out to the database
			header = chain.GetHeader(hash, number)
			if header == nil {
				return nil, consensus.ErrUnknownAncestor
			}
		}
		headers = append(headers, header)
		number, hash = number-1, header.ParentHash
	}

	// check if snapshot is nil
	if snap == nil {
		return nil, fmt.Errorf("unknown error while retrieving snapshot at block number %v", number)
	}

	// Previous snapshot found, apply any pending headers on top of it
	for i := 0; i < len(headers)/2; i++ {
		headers[i], headers[len(headers)-1-i] = headers[len(headers)-1-i], headers[i]
	}

	snap, err := snap.apply(headers, chain, parents, c.chainConfig.ChainID)
	if err != nil {
		return nil, err
	}
	c.recents.Add(snap.Hash, snap)

	// If we've generated a new checkpoint snapshot, save to disk
	if snap.Number%c.config.Epoch == 0 && len(headers) > 0 {
		if err = snap.store(c.db); err != nil {
			return nil, err
		}
		log.Trace("Stored snapshot to disk", "number", snap.Number, "hash", snap.Hash)
	}
	return snap, err
}

func (c *Consortium) VerifyUncles(chain consensus.ChainReader, block *types.Block) error {
	if len(block.Uncles()) > 0 {
		return errors.New("uncles not allowed")
	}
	return nil
}

// verifySeal checks whether the signature contained in the header satisfies the
// consensus protocol requirements. The method accepts an optional list of parent
// headers that aren't yet part of the local blockchain to generate the snapshots
// from.
func (c *Consortium) verifySeal(chain consensus.ChainHeaderReader, header *types.Header, parents []*types.Header) error {
	// Verifying the genesis block is not supported
	number := header.Number.Uint64()
	if number == 0 {
		return consortiumCommon.ErrUnknownBlock
	}
	// Retrieve the snapshot needed to verify this header and cache it
	snap, err := c.snapshot(chain, number-1, header.ParentHash, parents)
	if err != nil {
		return err
	}

	// Resolve the authorization key and check against validators
	signer, err := ecrecover(header, c.signatures, c.chainConfig.ChainID)
	if err != nil {
		return err
	}

	if signer != header.Coinbase {
		return errCoinBaseMisMatch
	}

	if _, ok := snap.Validators[signer]; !ok {
		return errUnauthorizedValidator
	}

	for seen, recent := range snap.Recents {
		if recent == signer {
			// Signer is among recents, only fail if the current block doesn't shift it out
			if limit := uint64(len(snap.Validators)/2 + 1); seen > number-limit {
				return consortiumCommon.ErrRecentlySigned
			}
		}
	}

	// Ensure that the difficulty corresponds to the turn-ness of the signer
	if !c.fakeDiff {
		inturn := snap.inturn(signer)
		if inturn && header.Difficulty.Cmp(diffInTurn) != 0 {
			return consortiumCommon.ErrWrongDifficulty
		}
		if !inturn && header.Difficulty.Cmp(diffNoTurn) != 0 {
			return consortiumCommon.ErrWrongDifficulty
		}
	}

	return nil
}

func (c *Consortium) Prepare(chain consensus.ChainHeaderReader, header *types.Header) error {
	if err := c.initContract(); err != nil {
		return err
	}

	header.Coinbase = c.val
	header.Nonce = types.BlockNonce{}

	number := header.Number.Uint64()
	snap, err := c.snapshot(chain, number-1, header.ParentHash, nil)
	if err != nil {
		return err
	}

	// Set the correct difficulty
	header.Difficulty = CalcDifficulty(snap, c.val)

	// Ensure the extra data has all it's components
	if len(header.Extra) < extraVanity {
		header.Extra = append(header.Extra, bytes.Repeat([]byte{0x00}, extraVanity-len(header.Extra))...)
	}
	header.Extra = header.Extra[:extraVanity]

	if number%c.config.Epoch == 0 || c.chainConfig.IsOnConsortiumV2(big.NewInt(int64(number))) {
		newValidators, err := c.contract.GetValidators(header)
		if err != nil {
			return err
		}
		// sort validator by address
		sort.Sort(validatorsAscending(newValidators))
		for _, validator := range newValidators {
			header.Extra = append(header.Extra, validator.Bytes()...)
		}
	}

	// add extra seal space
	header.Extra = append(header.Extra, make([]byte, extraSeal)...)

	// Mix digest is reserved for now, set to empty
	header.MixDigest = common.Hash{}

	// Ensure the timestamp has the correct delay
	parent := chain.GetHeader(header.ParentHash, number-1)
	if parent == nil {
		return consensus.ErrUnknownAncestor
	}
	header.Time = c.blockTimeForConsortiumV2Fork(snap, header, parent)
	if header.Time < uint64(time.Now().Unix()) {
		header.Time = uint64(time.Now().Unix())
	}
	return nil
}

func (c *Consortium) Finalize(chain consensus.ChainHeaderReader, header *types.Header, state *state.StateDB, txs *[]*types.Transaction,
	uncles []*types.Header, receipts *[]*types.Receipt, systemTxs *[]*types.Transaction, usedGas *uint64) error {
	if err := c.initContract(); err != nil {
		return err
	}
	// warn if not in majority fork
	number := header.Number.Uint64()
	snap, err := c.snapshot(chain, number-1, header.ParentHash, nil)
	if err != nil {
		return err
	}
	transactOpts := &consortiumCommon.ApplyTransactOpts{
		ApplyMessageOpts: &consortiumCommon.ApplyMessageOpts{
			State:        state,
			Header:       header,
			ChainConfig:  c.chainConfig,
			ChainContext: consortiumCommon.ChainContext{Chain: chain, Consortium: c},
		},
		Txs:         txs,
		Receipts:    receipts,
		ReceivedTxs: systemTxs,
		UsedGas:     usedGas,
		Mining:      false,
		Signer:      c.signer,
		SignTxFn:    c.signTxFn,
		EthAPI:      c.ethAPI,
	}

	// If the block is a epoch end block, verify the validator list
	// The verification can only be done when the state is ready, it can't be done in VerifyHeader.
	if header.Number.Uint64()%c.config.Epoch == 0 {
		newValidators, err := c.contract.GetValidators(header)
		if err != nil {
			return err
		}
		// sort validator by address
		sort.Sort(validatorsAscending(newValidators))
		validatorsBytes := make([]byte, len(newValidators)*validatorBytesLength)
		for i, validator := range newValidators {
			copy(validatorsBytes[i*validatorBytesLength:], validator.Bytes())
		}

		extraSuffix := len(header.Extra) - extraSeal
		if !bytes.Equal(header.Extra[extraVanity:extraSuffix], validatorsBytes) {
			return errMismatchingEpochValidators
		}
	}
	// No block rewards in PoA, so the state remains as is and uncles are dropped
	if header.Difficulty.Cmp(diffInTurn) != 0 {
		spoiledVal := snap.supposeValidator()
		signedRecently := false
		for _, recent := range snap.Recents {
			if recent == spoiledVal {
				signedRecently = true
				break
			}
		}
		if !signedRecently {
			err = c.contract.Slash(transactOpts, spoiledVal)
			if err != nil {
				// it is possible that slash validator failed because of the slash channel is disabled.
				log.Error("slash validator failed", "block hash", header.Hash(), "address", spoiledVal)
			}
		}
	}

	if header.Number.Uint64()%c.config.Epoch == c.config.Epoch-1 {
		if err := c.contract.WrapUpEpoch(transactOpts); err != nil {
			log.Error("Failed to update validators", "err", err)
		}
	}

	err = c.contract.SubmitBlockReward(transactOpts)
	if err != nil {
		return err
	}
	if len(*systemTxs) > 0 {
		return errors.New("the length of systemTxs do not match")
	}
	return nil
}

func (c *Consortium) FinalizeAndAssemble(chain consensus.ChainHeaderReader, header *types.Header, state *state.StateDB,
	txs []*types.Transaction, uncles []*types.Header, receipts []*types.Receipt) (*types.Block, []*types.Receipt, error) {
	log.Info("FinalizeAndAssemble", "Difficulty", header.Difficulty.Uint64(), "diffInTurn", diffInTurn.Uint64())
	if err := c.initContract(); err != nil {
		return nil, nil, err
	}

	// No block rewards in PoA, so the state remains as is and uncles are dropped
	if txs == nil {
		txs = make([]*types.Transaction, 0)
	}
	if receipts == nil {
		receipts = make([]*types.Receipt, 0)
	}

	transactOpts := &consortiumCommon.ApplyTransactOpts{
		ApplyMessageOpts: &consortiumCommon.ApplyMessageOpts{
			State:        state,
			Header:       header,
			ChainConfig:  c.chainConfig,
			ChainContext: consortiumCommon.ChainContext{Chain: chain, Consortium: c},
		},
		Txs:         &txs,
		Receipts:    &receipts,
		ReceivedTxs: nil,
		UsedGas:     &header.GasUsed,
		Mining:      true,
		Signer:      c.signer,
		SignTxFn:    c.signTxFn,
	}
	if header.Difficulty.Cmp(diffInTurn) != 0 {
		number := header.Number.Uint64()
		snap, err := c.snapshot(chain, number-1, header.ParentHash, nil)
		if err != nil {
			return nil, nil, err
		}
		spoiledVal := snap.supposeValidator()
		signedRecently := false
		for _, recent := range snap.Recents {
			if recent == spoiledVal {
				signedRecently = true
				break
			}
		}
		if !signedRecently {
			err = c.contract.Slash(transactOpts, spoiledVal)
			if err != nil {
				// it is possible that slash validator failed because of the slash channel is disabled.
				log.Error("Slash validator failed", "block", header.Number.Uint64(), "hash", header.Hash().Hex(), "address", spoiledVal, "error", err)
			}
		}
	}

	if header.Number.Uint64()%c.config.Epoch == c.config.Epoch-1 {
		if err := c.contract.WrapUpEpoch(transactOpts); err != nil {
			log.Error("Wrap up epoch failed", "block", header.Number.Uint64(), "hash", header.Hash().Hex(), "error", err)
		}
	}

	err := c.contract.SubmitBlockReward(transactOpts)
	if err != nil {
		return nil, nil, err
	}
	// should not happen. Once happen, stop the node is better than broadcast the block
	if header.GasLimit < header.GasUsed {
		return nil, nil, errors.New("gas consumption of system txs exceed the gas limit")
	}
	header.UncleHash = types.CalcUncleHash(nil)
	var blk *types.Block
	var rootHash common.Hash
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		rootHash = state.IntermediateRoot(chain.Config().IsEIP158(header.Number))
		wg.Done()
	}()
	go func() {
		blk = types.NewBlock(header, *transactOpts.Txs, nil, *transactOpts.Receipts, trie.NewStackTrie(nil))
		wg.Done()
	}()
	wg.Wait()
	blk.SetRoot(rootHash)
	// Assemble and return the final block for sealing
	return blk, *transactOpts.Receipts, nil
}

func (c *Consortium) Authorize(signer common.Address, signFn consortiumCommon.SignerFn, signTxFn consortiumCommon.SignerTxFn) {
	log.Info("Loaded authorize", "signFn", signFn, "signTxFn", signTxFn)
	c.lock.Lock()
	defer c.lock.Unlock()

	c.isMining = true
	c.val = signer
	c.signFn = signFn
	c.signTxFn = signTxFn
}

func (c *Consortium) Delay(chain consensus.ChainReader, header *types.Header) *time.Duration {
	number := header.Number.Uint64()
	snap, err := c.snapshot(chain, number-1, header.ParentHash, nil)
	if err != nil {
		return nil
	}
	delay := c.delayForConsortiumV2Fork(snap, header)
	// The blocking time should be no more than half of period
	half := time.Duration(c.config.Period) * time.Second / 2
	if delay > half {
		delay = half
	}
	return &delay
}

func (c *Consortium) Seal(chain consensus.ChainHeaderReader, block *types.Block, results chan<- *types.Block, stop <-chan struct{}) error {
	header := block.Header()

	// Sealing the genesis block is not supported
	number := header.Number.Uint64()
	if number == 0 {
		return consortiumCommon.ErrUnknownBlock
	}
	// For 0-period chains, refuse to seal empty blocks (no reward but would spin sealing)
	if c.config.Period == 0 && len(block.Transactions()) == 0 {
		log.Info("Sealing paused, waiting for transactions")
		return nil
	}
	// Don't hold the val fields for the entire sealing procedure
	c.lock.RLock()
	val, signFn := c.val, c.signFn
	c.lock.RUnlock()

	snap, err := c.snapshot(chain, number-1, header.ParentHash, nil)
	if err != nil {
		return err
	}

	// Bail out if we're unauthorized to sign a block
	if _, authorized := snap.Validators[val]; !authorized {
		return errUnauthorizedValidator
	}

	// If we're amongst the recent signers, wait for the next block
	for seen, recent := range snap.Recents {
		if recent == val {
			// Signer is among recents, only wait if the current block doesn't shift it out
			if limit := uint64(len(snap.Validators)/2 + 1); number < limit || seen > number-limit {
				log.Info("Signed recently, must wait for others")
				return nil
			}
		}
	}

	// Sweet, the protocol permits us to sign the block, wait for our time
	delay := time.Unix(int64(header.Time), 0).Sub(time.Now()) // nolint: gosimple
	if header.Difficulty.Cmp(diffInTurn) != 0 {
		// It's not our turn explicitly to sign, delay it a bit
		wiggle := time.Duration(len(snap.Validators)/2+1) * wiggleTime
		delay += time.Duration(rand.Int63n(int64(wiggle))) + wiggleTime // delay for 0.5s more

		log.Trace("Out-of-turn signing requested", "wiggle", common.PrettyDuration(wiggle))
	}

	log.Info("Sealing block with", "number", number, "delay", delay, "headerDifficulty", header.Difficulty, "val", val.Hex(), "txs", len(block.Transactions()))

	// Sign all the things!
	sig, err := signFn(accounts.Account{Address: val}, accounts.MimetypeConsortium, consortiumRLP(header, c.chainConfig.ChainID))
	if err != nil {
		return err
	}
	copy(header.Extra[len(header.Extra)-extraSeal:], sig)

	// Wait until sealing is terminated or delay timeout.
	log.Trace("Waiting for slot to sign and propagate", "delay", common.PrettyDuration(delay))
	go func() {
		select {
		case <-stop:
			return
		case <-time.After(delay):
		}

		select {
		case results <- block.WithSeal(header):
		default:
			log.Warn("Sealing result is not read by miner", "sealhash", SealHash(header, c.chainConfig.ChainID))
		}
	}()

	return nil
}

func (c *Consortium) SealHash(header *types.Header) common.Hash {
	return SealHash(header, c.chainConfig.ChainID)
}

func (c *Consortium) Close() error {
	return nil
}

func (c *Consortium) APIs(chain consensus.ChainHeaderReader) []rpc.API {
	return []rpc.API{}
}

func (c *Consortium) CalcDifficulty(chain consensus.ChainHeaderReader, time uint64, parent *types.Header) *big.Int {
	snap, err := c.snapshot(chain, parent.Number.Uint64(), parent.Hash(), nil)
	if err != nil {
		return nil
	}
	return CalcDifficulty(snap, c.val)
}

// CalcDifficulty is the difficulty adjustment algorithm. It returns the difficulty
// that a new block should have based on the previous blocks in the chain and the
// current signer.
func CalcDifficulty(snap *Snapshot, signer common.Address) *big.Int {
	if snap.inturn(signer) {
		return new(big.Int).Set(diffInTurn)
	}
	return new(big.Int).Set(diffNoTurn)
}

func (c *Consortium) getValidatorsFromHeader(header *types.Header) []common.Address {
	extraSuffix := len(header.Extra) - consortiumCommon.ExtraSeal
	return consortiumCommon.ExtractAddressFromBytes(header.Extra[extraVanity:extraSuffix])
}

func (c *Consortium) initContract() error {
	contract, err := consortiumCommon.NewContractIntegrator(c.chainConfig, consortiumCommon.NewConsortiumBackend(c.ethAPI), c.signTxFn, c.val)
	if err != nil {
		return err
	}
	c.contract = contract

	return nil
}

// Check if it is the turn of the signer from the last checkpoint
func (c *Consortium) signerInTurn(signer common.Address, number uint64, validators []common.Address) bool {
	lastCheckpoint := number / c.config.Epoch * c.config.Epoch
	index := (number - lastCheckpoint) % uint64(len(validators))
	return validators[index] == signer
}

// ecrecover extracts the Ethereum account address from a signed header.
func ecrecover(header *types.Header, sigcache *lru.ARCCache, chainId *big.Int) (common.Address, error) {
	// If the signature's already cached, return that
	hash := header.Hash()
	if address, known := sigcache.Get(hash); known {
		return address.(common.Address), nil
	}
	// Retrieve the signature from the header extra-data
	if len(header.Extra) < consortiumCommon.ExtraSeal {
		return common.Address{}, consortiumCommon.ErrMissingSignature
	}
	signature := header.Extra[len(header.Extra)-consortiumCommon.ExtraSeal:]

	// Recover the public key and the Ethereum address
	pubkey, err := crypto.Ecrecover(SealHash(header, chainId).Bytes(), signature)
	if err != nil {
		return common.Address{}, err
	}
	var signer common.Address
	copy(signer[:], crypto.Keccak256(pubkey[1:])[12:])

	sigcache.Add(hash, signer)
	return signer, nil
}

// SealHash returns the hash of a block prior to it being sealed.
func SealHash(header *types.Header, chainId *big.Int) (hash common.Hash) {
	hasher := sha3.NewLegacyKeccak256()
	encodeSigHeader(hasher, header, chainId)
	hasher.Sum(hash[:0])
	return hash
}

// consortiumRLP returns the rlp bytes which needs to be signed for the proof-of-authority
// sealing. The RLP to sign consists of the entire header apart from the 65 byte signature
// contained at the end of the extra data.
//
// Note, the method requires the extra data to be at least 65 bytes, otherwise it
// panics. This is done to avoid accidentally using both forms (signature present
// or not), which could be abused to produce different hashes for the same header.
func consortiumRLP(header *types.Header, chainId *big.Int) []byte {
	b := new(bytes.Buffer)
	encodeSigHeader(b, header, chainId)
	return b.Bytes()
}

func encodeSigHeader(w io.Writer, header *types.Header, chainId *big.Int) {
	err := rlp.Encode(w, []interface{}{
		chainId,
		header.ParentHash,
		header.UncleHash,
		header.Coinbase,
		header.Root,
		header.TxHash,
		header.ReceiptHash,
		header.Bloom,
		header.Difficulty,
		header.Number,
		header.GasLimit,
		header.GasUsed,
		header.Time,
		header.Extra[:len(header.Extra)-crypto.SignatureLength], // Yes, this will panic if extra is too short
		header.MixDigest,
		header.Nonce,
	})
	if err != nil {
		panic("can't encode: " + err.Error())
	}
}
