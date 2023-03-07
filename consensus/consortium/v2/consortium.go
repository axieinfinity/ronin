package v2

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"math/big"
	"math/rand"
	"sort"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/core"

	"github.com/common-nighthawk/go-figure"

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
)

const (
	inmemorySnapshots  = 128  // Number of recent vote snapshots to keep in memory
	inmemorySignatures = 4096 // Number of recent block signatures to keep in memory

	extraVanity = 32 // Fixed number of extra-data prefix bytes reserved for signer vanity
	extraSeal   = 65 // Fixed number of extra-data suffix bytes reserved for signer seal

	validatorBytesLength = common.AddressLength
	wiggleTime           = 1000 * time.Millisecond // Random delay (per signer) to allow concurrent signers
)

// Consortium delegated proof-of-stake protocol constants.
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

	// errCoinBaseMisMatch is returned if a header's coinbase do not match with signature
	errCoinBaseMisMatch = errors.New("coinbase do not match with signature")

	// errMismatchingEpochValidators is returned if a sprint block contains a
	// list of validators different from the one the local node calculated.
	errMismatchingEpochValidators = errors.New("mismatching validator list on epoch block")
)

// Consortium is the delegated proof-of-stake consensus engine proposed to support the
// Ronin to become more decentralized
type Consortium struct {
	chainConfig *params.ChainConfig
	config      *params.ConsortiumConfig // Consensus engine configuration parameters
	forkedBlock uint64
	genesisHash common.Hash
	db          ethdb.Database // Database to store and retrieve snapshot checkpoints

	recents    *lru.ARCCache // Snapshots for recent block to speed up reorgs
	signatures *lru.ARCCache // Signatures of recent blocks to speed up mining

	val      common.Address // Ethereum address of the signing key
	signer   types.Signer
	signFn   consortiumCommon.SignerFn // Signer function to authorize hashes with
	signTxFn consortiumCommon.SignerTxFn

	lock sync.RWMutex // Protects the signer fields

	ethAPI   *ethapi.PublicBlockChainAPI
	contract *consortiumCommon.ContractIntegrator

	fakeDiff bool
	v1       consortiumCommon.ConsortiumAdapter
}

// New creates a Consortium delegated proof-of-stake consensus engine
func New(
	chainConfig *params.ChainConfig,
	db ethdb.Database,
	ethAPI *ethapi.PublicBlockChainAPI,
	genesisHash common.Hash,
	v1 consortiumCommon.ConsortiumAdapter,
) *Consortium {
	consortiumConfig := chainConfig.Consortium

	if consortiumConfig != nil && consortiumConfig.EpochV2 == 0 {
		consortiumConfig.EpochV2 = epochLength
	}

	// Allocate the snapshot caches and create the engine
	recents, _ := lru.NewARC(inmemorySnapshots)
	signatures, _ := lru.NewARC(inmemorySignatures)

	return &Consortium{
		chainConfig: chainConfig,
		config:      consortiumConfig,
		genesisHash: genesisHash,
		db:          db,
		ethAPI:      ethAPI,
		recents:     recents,
		signatures:  signatures,
		signer:      types.NewEIP155Signer(chainConfig.ChainID),
		v1:          v1,
		forkedBlock: chainConfig.ConsortiumV2Block.Uint64(),
	}
}

// IsSystemTransaction implements consensus.PoSA, checking whether a transaction is a system
// transaction or not.
// A system transaction is a transaction that has the recipient of the contract address
// is defined in params.ConsortiumV2Contracts
func (c *Consortium) IsSystemTransaction(tx *types.Transaction, header *types.Header) (bool, error) {
	// deploy a contract
	if tx.To() == nil {
		return false, nil
	}
	sender, err := types.Sender(c.signer, tx)
	if err != nil {
		return false, errors.New("UnAuthorized transaction")
	}
	if sender == header.Coinbase && c.IsSystemContract(tx.To()) {
		return true, nil
	}
	return false, nil
}

// IsSystemContract implements consensus.PoSA, checking whether a contract is a system
// contract or not
// A system contract is a contract is defined in params.ConsortiumV2Contracts
func (c *Consortium) IsSystemContract(to *common.Address) bool {
	if to == nil {
		return false
	}
	return c.chainConfig.ConsortiumV2Contracts.IsSystemContract(*to)
}

// Author implements consensus.Engine, returning the coinbase directly
func (c *Consortium) Author(header *types.Header) (common.Address, error) {
	return header.Coinbase, nil
}

// VerifyHeader checks whether a header conforms to the consensus rules.
func (c *Consortium) VerifyHeader(chain consensus.ChainHeaderReader, header *types.Header, seal bool) error {
	return c.VerifyHeaderAndParents(chain, header, nil)
}

// VerifyHeaders implements consensus.Engine, always returning an empty abort and results channels.
// This method will be handled consortium/main.go instead
func (c *Consortium) VerifyHeaders(chain consensus.ChainHeaderReader, headers []*types.Header, seals []bool) (chan<- struct{}, <-chan error) {
	abort := make(chan struct{})
	results := make(chan error, len(headers))

	return abort, results
}

// GetRecents implements common.ConsortiumAdapter, always returning nil
// as this consensus mechanism doesn't need to get recents
func (c *Consortium) GetRecents(chain consensus.ChainHeaderReader, number uint64) map[uint64]common.Address {
	return nil
}

// VerifyHeaderAndParents checks whether a header conforms to the consensus rules.The
// caller may optionally pass in a batch of parents (ascending order) to avoid
// looking those up from the database. This is useful for concurrently verifying
// a batch of new headers.
func (c *Consortium) VerifyHeaderAndParents(chain consensus.ChainHeaderReader, header *types.Header, parents []*types.Header) error {
	if header.Number == nil {
		return consortiumCommon.ErrUnknownBlock
	}
	number := header.Number.Uint64()

	// Check that the extra-data contains the vanity, validators and signature.
	if len(header.Extra) < extraVanity {
		return consortiumCommon.ErrMissingVanity
	}
	if len(header.Extra) < extraVanity+extraSeal {
		return consortiumCommon.ErrMissingSignature
	}
	// Check extra data
	isEpoch := number%c.config.EpochV2 == 0 || c.chainConfig.IsOnConsortiumV2(header.Number)

	// Ensure that the extra-data contains a signer list on checkpoint, but none otherwise
	signersBytes := len(header.Extra) - extraVanity - extraSeal
	if !isEpoch && signersBytes != 0 {
		return consortiumCommon.ErrExtraValidators
	}

	if isEpoch && signersBytes%common.AddressLength != 0 {
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

// verifyCascadingFields verifies all the header fields that are not standalone,
// rather depend on a batch of previous headers. The caller may optionally pass
// in a batch of parents (ascending order) to avoid looking those up from the
// database. This is useful for concurrently verifying a batch of new headers.
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

	// Verify that the gas limit is <= 2^63-1
	capacity := uint64(0x7fffffffffffffff)
	if header.GasLimit > capacity {
		return fmt.Errorf("invalid gasLimit: have %v, max %v", header.GasLimit, capacity)
	}
	// Verify that the gasUsed is <= gasLimit
	if header.GasUsed > header.GasLimit {
		return fmt.Errorf("invalid gasUsed: have %d, gasLimit %d", header.GasUsed, header.GasLimit)
	}

	if err := misc.VerifyGaslimit(parent.GasLimit, header.GasLimit); err != nil {
		return err
	}

	// Retrieve the snapshot needed to verify this header and cache it
	snap, err := c.snapshot(chain, number-1, header.ParentHash, parents)
	if err != nil {
		return err
	}
	if err = c.verifyHeaderTime(header, parent, snap); err != nil {
		return err
	}

	// All basic checks passed, verify the seal and return
	return c.verifySeal(chain, header, parents, snap)
}

// snapshot retrieves the authorization snapshot at a given point in time.
func (c *Consortium) snapshot(chain consensus.ChainHeaderReader, number uint64, hash common.Hash, parents []*types.Header) (*Snapshot, error) {
	if err := c.initContract(); err != nil {
		return nil, err
	}
	// Search for a snapshot in memory or on disk for checkpoints
	var (
		headers    []*types.Header
		snap       *Snapshot
		cpyParents = make([]*types.Header, len(parents))
	)
	// NOTE(linh): We must copy parents before going to the loop because parents are modified.
	// 	If not, the FindAncientHeader function can not find its block ancestor
	copy(cpyParents, parents)

	for snap == nil {
		// If an in-memory snapshot was found, use that
		if s, ok := c.recents.Get(hash); ok {
			snap = s.(*Snapshot)
			break
		}

		// init snapshot if it is at forkedBlock
		if number == c.forkedBlock-1 {
			var (
				err        error
				validators []common.Address
			)
			snap, err = loadSnapshot(c.config, c.signatures, c.db, hash, c.ethAPI, c.chainConfig)
			if err == nil {
				log.Trace("Loaded snapshot from disk", "number", number, "hash", hash.Hex())
				break
			}

			// get validators set from number
			validators, err = c.contract.GetValidators(big.NewInt(0).SetUint64(number))
			if err != nil {
				log.Error("Load validators at the beginning failed", "err", err)
				return nil, err
			}
			snap = newSnapshot(c.chainConfig, c.config, c.signatures, number, hash, validators, c.ethAPI)

			// load v1 recent list to prevent recent producing-block-validators produce block again
			snapV1 := c.v1.GetSnapshot(chain, number, parents)

			// NOTE(linh): In version 1, the snapshot is not used correctly, so we must clean up
			// 	incorrect data in the recent list before going to version 2
			// 	Example: The current block is 1000, and the recents list is
			// 	[2: address1, 3: address2, ...,998: addressN - 1,999: addressN]
			// 	we need to remove elements which are not continuously
			// 	So the final result must be [998: addressN - 1,999: addressN]
			snap.Recents = consortiumCommon.RemoveOutdatedRecents(snapV1.Recents, number)

			if err := snap.store(c.db); err != nil {
				return nil, err
			}
			log.Info("Stored checkpoint snapshot to disk", "number", number, "hash", hash)
			figure.NewColorFigure("Welcome to DPOS", "", "green", true).Print()
			break
		}

		// If an on-disk checkpoint snapshot can be found, use that
		if number%c.config.EpochV2 == 0 {
			var err error
			snap, err = loadSnapshot(c.config, c.signatures, c.db, hash, c.ethAPI, c.chainConfig)
			if err != nil {
				log.Debug("Load snapshot failed", "number", number, "hash", hash.Hex())
			} else {
				log.Trace("Loaded snapshot from disk", "number", number, "hash", hash.Hex())
				break
			}
		}

		// No snapshot for this header, gather the header and move backward
		// NOTE: We are modifying parents in here
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

	// Checking if snapshot is nil
	if snap == nil {
		return nil, fmt.Errorf("unknown error while retrieving snapshot at block number %v", number)
	}

	// Previous snapshot found, apply any pending headers on top of it
	for i := 0; i < len(headers)/2; i++ {
		headers[i], headers[len(headers)-1-i] = headers[len(headers)-1-i], headers[i]
	}

	snap, err := snap.apply(headers, chain, cpyParents, c.chainConfig.ChainID)
	if err != nil {
		return nil, err
	}
	c.recents.Add(snap.Hash, snap)

	// If we've generated a new checkpoint snapshot, save to disk
	if snap.Number%c.config.EpochV2 == 0 && len(headers) > 0 {
		if err = snap.store(c.db); err != nil {
			return nil, err
		}
		log.Trace("Stored snapshot to disk", "number", snap.Number, "hash", snap.Hash)
	}
	return snap, err
}

// VerifyUncles implements consensus.Engine, always returning an error for any
// uncles as this consensus mechanism doesn't permit uncles.
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
func (c *Consortium) verifySeal(chain consensus.ChainHeaderReader, header *types.Header, parents []*types.Header, snap *Snapshot) error {
	// Verifying the genesis block is not supported
	number := header.Number.Uint64()
	if number == 0 {
		return consortiumCommon.ErrUnknownBlock
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

func backOffTime(header *types.Header, snapshot *Snapshot) uint64 {
	coinbase := header.Coinbase
	if snapshot.inturn(coinbase) {
		return 0
	}

	position, numOfSealableValidators := snapshot.sealableValidators(coinbase)
	// This block doesn't pass the recent check and will fail later, returns
	// dummy value for delay here
	if position == -1 {
		return 0
	}

	source := rand.NewSource(header.Number.Int64())
	rand := rand.New(source)

	// Every validator that can seal a block may have different delay
	// The delay equals to their random delayMultiplier * wiggleTime.
	// The delayMultiplier is random in range [1, numOfSealableValidators]
	delayMultiplier := make([]int, numOfSealableValidators)
	for i := range delayMultiplier {
		delayMultiplier[i] = i + 1
	}
	rand.Shuffle(len(delayMultiplier), func(i, j int) {
		delayMultiplier[i], delayMultiplier[j] = delayMultiplier[j], delayMultiplier[i]
	})

	return uint64((int(wiggleTime) + delayMultiplier[position]*int(wiggleTime)/2) / int(time.Second))
}

func (c *Consortium) computeHeaderTime(header, parent *types.Header, snapshot *Snapshot) uint64 {
	headerTime := parent.Time + c.config.Period

	if c.chainConfig.IsBuba(header.Number) {
		headerTime += backOffTime(header, snapshot)
	}

	if headerTime < uint64(time.Now().Unix()) {
		headerTime = uint64(time.Now().Unix())
	}
	return headerTime
}

func (c *Consortium) verifyHeaderTime(header, parent *types.Header, snapshot *Snapshot) error {
	if header.Time > uint64(time.Now().Unix()) {
		return consensus.ErrFutureBlock
	}

	if c.chainConfig.IsBuba(header.Number) {
		expectedHeaderTime := parent.Time + c.config.Period + backOffTime(header, snapshot)
		if header.Time < expectedHeaderTime {
			return consensus.ErrFutureBlock
		}
	}

	return nil
}

// Prepare implements consensus.Engine, preparing all the consensus fields of the
// header for running the transactions on top.
func (c *Consortium) Prepare(chain consensus.ChainHeaderReader, header *types.Header) error {
	if err := c.initContract(); err != nil {
		return err
	}

	coinbase, _, _ := c.readSigner()
	header.Coinbase = coinbase
	header.Nonce = types.BlockNonce{}

	number := header.Number.Uint64()
	snap, err := c.snapshot(chain, number-1, header.ParentHash, nil)
	if err != nil {
		return err
	}

	// Set the correct difficulty
	header.Difficulty = CalcDifficulty(snap, coinbase)

	// Ensure the extra data has all it's components
	if len(header.Extra) < extraVanity {
		header.Extra = append(header.Extra, bytes.Repeat([]byte{0x00}, extraVanity-len(header.Extra))...)
	}
	header.Extra = header.Extra[:extraVanity]

	if number%c.config.EpochV2 == 0 || c.chainConfig.IsOnConsortiumV2(big.NewInt(int64(number))) {
		// This block is not inserted, the transactions in this block are not applied, so we need
		// the call GetValidators at the context of previous block
		newValidators, err := c.contract.GetValidators(new(big.Int).Sub(header.Number, common.Big1))
		if err != nil {
			return err
		}
		// Sort validators by address
		sort.Sort(validatorsAscending(newValidators))
		for _, validator := range newValidators {
			header.Extra = append(header.Extra, validator.Bytes()...)
		}
	}

	// Add extra seal space
	header.Extra = append(header.Extra, make([]byte, extraSeal)...)

	// Mix digest is reserved for now, set to empty
	header.MixDigest = common.Hash{}

	// Ensure the timestamp has the correct delay
	parent := chain.GetHeader(header.ParentHash, number-1)
	if parent == nil {
		return consensus.ErrUnknownAncestor
	}

	header.Time = c.computeHeaderTime(header, parent, snap)
	return nil
}

func (c *Consortium) submitBlockReward(transactOpts *consortiumCommon.ApplyTransactOpts) error {
	if err := c.contract.SubmitBlockReward(transactOpts); err != nil {
		log.Error("Failed to submit block reward", "err", err)
		return err
	}
	return nil
}

func (c *Consortium) processSystemTransactions(chain consensus.ChainHeaderReader, header *types.Header,
	transactOpts *consortiumCommon.ApplyTransactOpts, isFinalizeAndAssemble bool) error {

	if header.Difficulty.Cmp(diffInTurn) != 0 {
		number := header.Number.Uint64()
		snap, err := c.snapshot(chain, number-1, header.ParentHash, nil)
		if err != nil {
			return err
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
			if !isFinalizeAndAssemble {
				log.Info("Slash validator", "number", header.Number, "spoiled", spoiledVal)
			}
			if err := c.contract.Slash(transactOpts, spoiledVal); err != nil {
				// it is possible that slash validator failed because of the slash channel is disabled.
				log.Error("Failed to slash validator", "block hash", header.Hash(), "address", spoiledVal)
				return err
			}
		}
	}

	// Previously, we call WrapUpEpoch before SubmitBlockReward which is the wrong order.
	// We create a hardfork here to fix the contract call order.
	if c.chainConfig.IsPuffy(header.Number) {
		if err := c.submitBlockReward(transactOpts); err != nil {
			return err
		}
	}

	if header.Number.Uint64()%c.config.EpochV2 == c.config.EpochV2-1 {
		if err := c.contract.WrapUpEpoch(transactOpts); err != nil {
			log.Error("Failed to wrap up epoch", "err", err)
			return err
		}
	}

	if !c.chainConfig.IsPuffy(header.Number) {
		return c.submitBlockReward(transactOpts)
	}

	return nil
}

// Finalize implements consensus.Engine that calls three methods from smart contracts:
// - WrapUpEpoch at epoch to distribute rewards and sort the validators set
// - Slash the validator who does not sign if it is in-turn
// - SubmitBlockRewards of the current block
func (c *Consortium) Finalize(chain consensus.ChainHeaderReader, header *types.Header, state *state.StateDB, txs *[]*types.Transaction,
	uncles []*types.Header, receipts *[]*types.Receipt, systemTxs *[]*types.Transaction, internalTxs *[]*types.InternalTransaction, usedGas *uint64) error {
	if err := c.initContract(); err != nil {
		return err
	}
	_, _, signTxFn := c.readSigner()
	evmContext := core.NewEVMBlockContext(header, consortiumCommon.ChainContext{Chain: chain, Consortium: c}, &header.Coinbase, chain.OpEvents()...)
	transactOpts := &consortiumCommon.ApplyTransactOpts{
		ApplyMessageOpts: &consortiumCommon.ApplyMessageOpts{
			State:       state,
			Header:      header,
			ChainConfig: c.chainConfig,
			EVMContext:  &evmContext,
		},
		Txs:      txs,
		Receipts: receipts,
		// a.k.a. System Txs
		// systemTxs is received from other nodes
		ReceivedTxs: systemTxs,
		UsedGas:     usedGas,
		Mining:      false,
		Signer:      c.signer,
		SignTxFn:    signTxFn,
		EthAPI:      c.ethAPI,
	}

	// If the block is an epoch end block, verify the validator list
	// The verification can only be done when the state is ready, it can't be done in VerifyHeader.
	if header.Number.Uint64()%c.config.EpochV2 == 0 {
		// The GetValidators in Prepare is called on the context of previous block so here it must
		// be called on context of previous block too
		newValidators, err := c.contract.GetValidators(new(big.Int).Sub(header.Number, common.Big1))
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

	if err := c.processSystemTransactions(chain, header, transactOpts, false); err != nil {
		return err
	}
	if len(*transactOpts.EVMContext.InternalTransactions) > 0 {
		*internalTxs = append(*internalTxs, *transactOpts.EVMContext.InternalTransactions...)
	}
	if len(*systemTxs) > 0 {
		return errors.New("the length of systemTxs do not match")
	}
	return nil
}

// FinalizeAndAssemble implements consensus.Engine that calls three methods from smart contracts:
// - WrapUpEpoch at epoch to distribute rewards and sort the validators set
// - Slash the validator who does not sign if it is in-turn
// - SubmitBlockRewards of the current block
func (c *Consortium) FinalizeAndAssemble(chain consensus.ChainHeaderReader, header *types.Header, state *state.StateDB,
	txs []*types.Transaction, uncles []*types.Header, receipts []*types.Receipt) (*types.Block, []*types.Receipt, error) {
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
	_, _, signTxFn := c.readSigner()
	evmContext := core.NewEVMBlockContext(header, consortiumCommon.ChainContext{Chain: chain, Consortium: c}, &header.Coinbase, chain.OpEvents()...)
	transactOpts := &consortiumCommon.ApplyTransactOpts{
		ApplyMessageOpts: &consortiumCommon.ApplyMessageOpts{
			State:       state,
			Header:      header,
			ChainConfig: c.chainConfig,
			EVMContext:  &evmContext,
		},
		Txs:      &txs,
		Receipts: &receipts,
		// a.k.a. System Txs
		// It always equals nil since FinalizeAndAssemble doesn't receive any transactions
		ReceivedTxs: nil,
		UsedGas:     &header.GasUsed,
		Mining:      true,
		Signer:      c.signer,
		SignTxFn:    signTxFn,
	}

	if err := c.processSystemTransactions(chain, header, transactOpts, true); err != nil {
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

// Authorize injects a private key into the consensus engine to mint new blocks with
func (c *Consortium) Authorize(signer common.Address, signFn consortiumCommon.SignerFn, signTxFn consortiumCommon.SignerTxFn) {
	c.lock.Lock()
	defer c.lock.Unlock()

	c.val = signer
	c.signFn = signFn
	c.signTxFn = signTxFn
}

// Seal implements consensus.Engine, attempting to create a sealed block using
// the local signing credentials.
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
	val, signFn, _ := c.readSigner()

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
				return consortiumCommon.ErrRecentlySigned
			}
		}
	}

	// Sweet, the protocol permits us to sign the block, wait for our time
	// After the Buba hardfork, the delay is included in header time already
	delay := time.Until(time.Unix(int64(header.Time), 0))
	if !c.chainConfig.IsBuba(block.Number()) {
		if header.Difficulty.Cmp(diffInTurn) != 0 {
			// It's not our turn explicitly to sign, delay it a bit
			wiggle := time.Duration(len(snap.Validators)/2+1) * wiggleTime
			delay += time.Duration(rand.Int63n(int64(wiggle))) + wiggleTime // delay for 0.5s more

			log.Trace("Out-of-turn signing requested", "wiggle", common.PrettyDuration(wiggle))
		}
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

// SealHash returns the hash of a block prior to it being sealed.
func (c *Consortium) SealHash(header *types.Header) common.Hash {
	return SealHash(header, c.chainConfig.ChainID)
}

// Close implements consensus.Engine. It's a noop for Consortium as there are no background threads.
func (c *Consortium) Close() error {
	return nil
}

// APIs are backward compatible with the v1, so we do not to implement it again
func (c *Consortium) APIs(chain consensus.ChainHeaderReader) []rpc.API {
	return []rpc.API{}
}

// CalcDifficulty is the difficulty adjustment algorithm. It returns the difficulty
// that a new block should have:
// * DIFF_NOTURN(3) if (BLOCK_NUMBER + 1) / VALIDATOR_COUNT != VALIDATOR_INDEX
// * DIFF_INTURN(7) if (BLOCK_NUMBER + 1) / VALIDATOR_COUNT == VALIDATOR_INDEX
func (c *Consortium) CalcDifficulty(chain consensus.ChainHeaderReader, time uint64, parent *types.Header) *big.Int {
	snap, err := c.snapshot(chain, parent.Number.Uint64(), parent.Hash(), nil)
	if err != nil {
		return nil
	}
	coinbase, _, _ := c.readSigner()
	return CalcDifficulty(snap, coinbase)
}

// CalcDifficulty is the difficulty adjustment algorithm. It returns the difficulty
// that a new block should have based on the previous blocks in the chain and the
// current validator.
func CalcDifficulty(snap *Snapshot, signer common.Address) *big.Int {
	if snap.inturn(signer) {
		return new(big.Int).Set(diffInTurn)
	}
	return new(big.Int).Set(diffNoTurn)
}

// initContract creates NewContractIntegrator instance
func (c *Consortium) initContract() error {
	coinbase, _, signTxFn := c.readSigner()
	contract, err := consortiumCommon.NewContractIntegrator(c.chainConfig, consortiumCommon.NewConsortiumBackend(c.ethAPI), signTxFn, coinbase)
	if err != nil {
		return err
	}
	c.contract = contract

	return nil
}

func (c *Consortium) readSigner() (common.Address, consortiumCommon.SignerFn, consortiumCommon.SignerTxFn) {
	c.lock.RLock()
	defer c.lock.RUnlock()

	return c.val, c.signFn, c.signTxFn
}

// ecrecover extracts the Ronin account address from a signed header.
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

// encodeSigHeader encodes the whole header with chainId.
// chainID was introduced in EIP-155 to prevent replay attacks between the main ETH and ETC chains,
// which both have a networkID of 1
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
		header.Extra[:len(header.Extra)-consortiumCommon.ExtraSeal], // Yes, this will panic if extra is too short
		header.MixDigest,
		header.Nonce,
	})
	if err != nil {
		panic("can't encode: " + err.Error())
	}
}
