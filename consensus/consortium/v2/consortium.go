package v2

import (
	"bytes"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"math"
	"math/big"
	"math/rand"
	"sort"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/core"

	"github.com/common-nighthawk/go-figure"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus"
	consortiumCommon "github.com/ethereum/go-ethereum/consensus/consortium/common"
	"github.com/ethereum/go-ethereum/consensus/consortium/v2/finality"
	"github.com/ethereum/go-ethereum/consensus/misc"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/bls/blst"
	blsCommon "github.com/ethereum/go-ethereum/crypto/bls/common"
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

	wiggleTime          = 1000 * time.Millisecond // Random delay (per signer) to allow concurrent signers
	unSealableValidator = -1

	finalityRatio                  float64 = 2.0 / 3
	assemblingFinalityVoteDuration         = 1 * time.Second
	MaxValidatorCandidates                 = 64 // Maximum number of validator candidates (aka voters for a block).
	dayInSeconds                           = uint64(86400)
)

// Consortium delegated proof-of-stake protocol constants.
var (
	epochLength = uint64(30000) // Default number of blocks after which to checkpoint

	uncleHash = types.CalcUncleHash(nil) // Always Keccak256(RLP([])) as uncles are meaningless outside of PoW

	diffInTurn = big.NewInt(7) // Block difficulty for in-turn signatures
	diffNoTurn = big.NewInt(3) // Block difficulty for out-of-turn signatures

	// The proxy contract's implementation slot
	// https://github.com/OpenZeppelin/openzeppelin-contracts-upgradeable/blob/v4.7.3/contracts/proxy/ERC1967/ERC1967UpgradeUpgradeable.sol#L34
	implementationSlot = common.HexToHash("360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc")
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

	// errMismatchingValidators is returned if a sprint block contains a
	// list of validators different from the one the local node calculated.
	errMismatchingValidators = errors.New("mismatching validator list")
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

	lock     sync.RWMutex              // Protects the below 4 fields
	val      common.Address            // Ethereum address of the signing key
	signFn   consortiumCommon.SignerFn // Signer function to authorize hashes with
	signTxFn consortiumCommon.SignerTxFn
	contract consortiumCommon.ContractInteraction

	signer types.Signer
	ethAPI *ethapi.PublicBlockChainAPI

	fakeDiff bool
	v1       consortiumCommon.ConsortiumAdapter

	votePool consensus.VotePool

	// This is used in unit test only
	testTrippEffective bool
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

	consortium := Consortium{
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
	err := consortium.initContract(common.Address{}, nil)
	if err != nil {
		log.Error("Failed to init system contract caller", "err", err)
	}

	return &consortium
}

// IsSystemMessage implements consensus.PoSA, checking whether a transaction is a system
// transaction or not.
// A system transaction is a transaction that has the recipient of the contract address
// is defined in params.ConsortiumV2Contracts
func (c *Consortium) IsSystemMessage(msg core.Message, header *types.Header) bool {
	// deploy a contract
	if msg.To() == nil {
		return false
	}
	if c.chainConfig.IsBuba(header.Number) {
		if msg.From() == header.Coinbase && c.IsSystemContract(msg.To()) {
			return true
		}
	} else {
		if msg.From() == header.Coinbase && c.IsSystemContract(msg.To()) && msg.GasPrice().Cmp(big.NewInt(0)) == 0 {
			return true
		}
	}
	return false
}

// In normal case, IsSystemTransaction in consortium/main.go is used instead of this function. This function
// is only used in testing when we create standalone consortium v2 engine without the v1
func (c *Consortium) IsSystemTransaction(tx *types.Transaction, header *types.Header) (bool, error) {
	msg, err := tx.AsMessage(types.MakeSigner(c.chainConfig, header.Number), header.BaseFee)
	if err != nil {
		return false, err
	}
	return c.IsSystemMessage(msg, header), nil
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
// In normal case, VerifyHeaders in consortium/main.go is used instead of this function. This function
// is only used in testing when we create standalone consortium v2 engine without the v1
func (c *Consortium) VerifyHeaders(chain consensus.ChainHeaderReader, headers []*types.Header, seals []bool) (chan<- struct{}, <-chan error) {
	abort := make(chan struct{})
	results := make(chan error, len(headers))

	go func() {
		for i, header := range headers {
			err := c.VerifyHeaderAndParents(chain, header, headers[:i])
			select {
			case <-abort:
				return
			case results <- err:
			}
		}
	}()

	return abort, results
}

// GetRecents implements common.ConsortiumAdapter, always returning nil
// as this consensus mechanism doesn't need to get recents
func (c *Consortium) GetRecents(chain consensus.ChainHeaderReader, number uint64) map[uint64]common.Address {
	return nil
}

// VerifyVote check if the finality voter is in the validator set, it assumes the signature is
// already verified
func (c *Consortium) VerifyVote(chain consensus.ChainHeaderReader, vote *types.VoteEnvelope) error {
	header := chain.GetHeaderByHash(vote.Data.TargetHash)
	if header == nil {
		return errors.New("header not found")
	}

	if header.Number.Uint64() != vote.Data.TargetNumber {
		return finality.ErrInvalidTargetNumber
	}

	// Look at the comment assembleFinalityVote in function for the
	// detailed explanation on the snapshot we need to get to verify the
	// finality vote.
	// Here we want to verify vote for TargetNumber, so we get snapshot
	// at TargetNumber.
	snap, err := c.snapshot(chain, vote.Data.TargetNumber, vote.Data.TargetHash, nil)
	if err != nil {
		return err
	}

	publicKey, err := blst.PublicKeyFromBytes(vote.PublicKey[:])
	if err != nil {
		return err
	}
	if !snap.inBlsPublicKeySet(publicKey) {
		return finality.ErrUnauthorizedFinalityVoter
	}

	return nil
}

// verifyFinalitySignatures verifies the finality signatures in the block header
func (c *Consortium) verifyFinalitySignatures(
	chain consensus.ChainHeaderReader,
	finalityVotedValidators finality.FinalityVoteBitSet,
	finalitySignatures blsCommon.Signature,
	header *types.Header,
	parents []*types.Header,
) error {
	isTrippEffective := c.IsTrippEffective(chain, header)
	snap, err := c.snapshot(chain, header.Number.Uint64()-1, header.ParentHash, parents)
	if err != nil {
		return err
	}

	voteData := types.VoteData{
		TargetNumber: header.Number.Uint64() - 1,
		TargetHash:   header.ParentHash,
	}
	digest := voteData.Hash()

	// verify aggregated signature
	var (
		publicKeys            []blsCommon.PublicKey
		accumulatedVoteWeight int
		finalityThreshold     int
	)
	votedValidatorPositions := finalityVotedValidators.Indices()
	for _, position := range votedValidatorPositions {
		if position >= len(snap.ValidatorsWithBlsPub) {
			return finality.ErrInvalidFinalityVotedBitSet
		}
		publicKeys = append(publicKeys, snap.ValidatorsWithBlsPub[position].BlsPublicKey)
		if isTrippEffective {
			accumulatedVoteWeight += int(snap.ValidatorsWithBlsPub[position].Weight)
		} else {
			accumulatedVoteWeight += 1
		}
	}

	if isTrippEffective {
		finalityThreshold = int(math.Floor(finalityRatio*float64(consortiumCommon.MaxFinalityVotePercentage))) + 1
	} else {
		finalityThreshold = int(math.Floor(finalityRatio*float64(len(snap.ValidatorsWithBlsPub)))) + 1
	}

	if accumulatedVoteWeight < finalityThreshold {
		return finality.ErrNotEnoughFinalityVote
	}

	if !finalitySignatures.FastAggregateVerify(publicKeys, digest) {
		return finality.ErrFinalitySignatureVerificationFailed
	}

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

	// Ensure that the mix digest is zero as we don't have fork protection currently
	if header.MixDigest != (common.Hash{}) {
		return consortiumCommon.ErrInvalidMixDigest
	}
	// Ensure that the block doesn't contain any uncles which are meaningless in PoA
	if header.UncleHash != uncleHash {
		return consortiumCommon.ErrInvalidUncleHash
	}
	// Ensure that the block's difficulty is meaningful (may not be correct at this point)
	if header.Number.Uint64() > 0 {
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

	// Check extra data
	isShillin := c.chainConfig.IsShillin(header.Number)
	isTrippEffective := c.IsTrippEffective(chain, header)
	extraData, err := finality.DecodeExtraV2(header.Extra, c.chainConfig, header.Number)
	if err != nil {
		return err
	}

	// Check extra data
	isEpoch := number%c.config.EpochV2 == 0 || c.chainConfig.IsOnConsortiumV2(header.Number)

	if !isEpoch && len(extraData.CheckpointValidators) != 0 {
		return consortiumCommon.ErrExtraValidators
	}

	if isTrippEffective && isEpoch && len(extraData.BlockProducers) == 0 {
		return consortiumCommon.ErrExtraValidators
	}

	if isTrippEffective && c.IsPeriodBlock(chain, header) && len(extraData.CheckpointValidators) == 0 {
		return consortiumCommon.ErrExtraValidators
	}

	if !isTrippEffective && len(extraData.BlockProducers) != 0 {
		return consortiumCommon.ErrExtraValidators
	}

	if isShillin && extraData.HasFinalityVote == 1 {
		if err := c.verifyFinalitySignatures(
			chain,
			extraData.FinalityVotedValidators,
			extraData.AggregatedFinalityVotes,
			header,
			parents,
		); err != nil {
			return err
		}
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

		// this case is only happened in mock mode
		if number == 0 {
			validators, err := c.contract.GetBlockProducers(common.Big0)
			if err != nil {
				return nil, err
			}
			snap = newSnapshot(c.chainConfig, c.config, c.signatures, number, hash, validators, nil, c.ethAPI)
			break
		}

		// init snapshot if it is at forkedBlock
		if number == c.forkedBlock-1 {
			var (
				err error
			)
			snap, err = loadSnapshot(c.config, c.signatures, c.db, hash, c.ethAPI, c.chainConfig)
			if err == nil {
				log.Trace("Loaded snapshot from disk", "number", number, "hash", hash.Hex())
				break
			}

			// get validators set from number
			_, _, _, contract := c.readSignerAndContract()
			validators, err := contract.GetBlockProducers(big.NewInt(0).SetUint64(number))
			if err != nil {
				log.Error("Load validators at the beginning failed", "err", err)
				return nil, err
			}

			snap = newSnapshot(c.chainConfig, c.config, c.signatures, number, hash, validators, nil, c.ethAPI)

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
			if header == nil || header.Hash() != hash || header.Number.Uint64() != number {
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
	log.Trace("Checking snapshot data", "number", snap.Number, "validators", snap.validators())
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

	if !snap.inInValidatorSet(signer) {
		return errUnauthorizedValidator
	}

	if snap.IsRecentlySigned(signer) {
		return consortiumCommon.ErrRecentlySigned
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

func backOffTime(header *types.Header, snapshot *Snapshot, chainConfig *params.ChainConfig) uint64 {
	coinbase := header.Coinbase
	if snapshot.inturn(coinbase) {
		return 0
	}

	position, numOfSealableValidators := snapshot.sealableValidators(coinbase)
	// This block doesn't pass the recent check and will fail later, returns
	// dummy value for delay here
	if position == unSealableValidator {
		return 0
	}

	initialDelay := time.Second
	if chainConfig.IsOlek(new(big.Int).SetUint64(snapshot.Number + 1)) {
		inturnValidator := snapshot.supposeValidator()
		pos, _ := snapshot.sealableValidators(inturnValidator)
		if pos == unSealableValidator {
			initialDelay = 0
		}
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

	if chainConfig.IsOlek(new(big.Int).SetUint64(snapshot.Number + 1)) {
		return uint64((int(initialDelay) + (delayMultiplier[position]-1)*int(wiggleTime)) / int(time.Second))
	} else {
		return uint64((int(initialDelay) + delayMultiplier[position]*int(wiggleTime)/2) / int(time.Second))
	}
}

func (c *Consortium) computeHeaderTime(header, parent *types.Header, snapshot *Snapshot) uint64 {
	headerTime := parent.Time + c.config.Period

	if c.chainConfig.IsBuba(header.Number) {
		headerTime += backOffTime(header, snapshot, c.chainConfig)
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
		expectedHeaderTime := parent.Time + c.config.Period + backOffTime(header, snapshot, c.chainConfig)
		if header.Time < expectedHeaderTime {
			return consensus.ErrFutureBlock
		}
	}

	return nil
}

func (c *Consortium) getCheckpointValidatorsFromContract(
	chain consensus.ChainHeaderReader,
	isPeriodBlock bool,
	header *types.Header,
) ([]finality.ValidatorWithBlsPub, []common.Address, error) {
	parentBlockNumber := new(big.Int).Sub(header.Number, common.Big1)
	_, _, _, contract := c.readSignerAndContract()

	blockProducers, err := contract.GetBlockProducers(parentBlockNumber)
	if err != nil {
		return nil, nil, err
	}

	var checkpointValidators []finality.ValidatorWithBlsPub

	// After Tripp, bls key and staked amount are read only once at
	// the start of new period, whereas block producer address is read
	// at the start of every epoch.
	if c.IsTrippEffective(chain, header) {
		sort.Sort(validatorsAscending(blockProducers))
		if !isPeriodBlock {
			return nil, blockProducers, nil
		}
		validatorCandidates, err := contract.GetValidatorCandidates(parentBlockNumber)
		if err != nil {
			return nil, nil, err
		}
		// After Tripp, there is a upper bound for the number of validator candidates
		// (aka voters for blocks) during a period, which is ensured by the contract.
		// However, we add additional check here to ensure that the field FinalityVoteBitSet
		// (of type uint64) works in proper manner.
		if len(validatorCandidates) > MaxValidatorCandidates {
			validatorCandidates = validatorCandidates[:MaxValidatorCandidates]
		}
		stakedAmounts, err := c.contract.GetStakedAmount(header.Number, validatorCandidates)
		if err != nil {
			return nil, nil, err
		}
		weights := consortiumCommon.NormalizeFinalityVoteWeight(stakedAmounts)
		for i, candidate := range validatorCandidates {
			blsPublicKey, err := contract.GetBlsPublicKey(parentBlockNumber, candidate)
			if err == nil {
				checkpointValidators = append(checkpointValidators, finality.ValidatorWithBlsPub{
					Address:      candidate,
					BlsPublicKey: blsPublicKey,
					Weight:       weights[i],
				})
			}
		}
		return checkpointValidators, blockProducers, nil
	}

	var (
		blsPublicKeys      []blsCommon.PublicKey
		filteredValidators []common.Address = blockProducers
	)
	isShillin := c.chainConfig.IsShillin(header.Number)
	if isShillin {
		// The filteredValidators shares the same underlying array with newValidators
		// See more: https://github.com/golang/go/wiki/SliceTricks#filtering-without-allocating
		filteredValidators = filteredValidators[:0]
		for _, validator := range blockProducers {
			blsPublicKey, err := contract.GetBlsPublicKey(parentBlockNumber, validator)
			if err == nil {
				filteredValidators = append(filteredValidators, validator)
				blsPublicKeys = append(blsPublicKeys, blsPublicKey)
			}
		}
	}

	for i := range filteredValidators {
		validatorWithBlsPub := finality.ValidatorWithBlsPub{
			Address: filteredValidators[i],
		}
		if isShillin {
			validatorWithBlsPub.BlsPublicKey = blsPublicKeys[i]
		}
		checkpointValidators = append(checkpointValidators, validatorWithBlsPub)
	}

	// sort validator by address
	sort.Sort(finality.CheckpointValidatorAscending(checkpointValidators))
	return checkpointValidators, nil, nil
}

// Prepare implements consensus.Engine, preparing all the consensus fields of the
// header for running the transactions on top.
func (c *Consortium) Prepare(chain consensus.ChainHeaderReader, header *types.Header) error {
	coinbase, _, _, _ := c.readSignerAndContract()
	header.Coinbase = coinbase
	header.Nonce = types.BlockNonce{}

	number := header.Number.Uint64()
	snap, err := c.snapshot(chain, number-1, header.ParentHash, nil)
	if err != nil {
		return err
	}

	// Set the correct difficulty
	header.Difficulty = CalcDifficulty(snap, coinbase)

	var extraData finality.HeaderExtraData

	if number%c.config.EpochV2 == 0 || c.chainConfig.IsOnConsortiumV2(big.NewInt(int64(number))) {
		isPeriodBlock := c.IsPeriodBlock(chain, header)
		checkpointValidator, blockProducers, err := c.getCheckpointValidatorsFromContract(chain, isPeriodBlock, header)
		if err != nil {
			return err
		}
		// After Tripp, validator candidates are read only once at
		// the start of new period, whereas block producer address is read
		// at the start of every epoch.
		if c.IsTrippEffective(chain, header) {
			if isPeriodBlock {
				extraData.CheckpointValidators = checkpointValidator
			}
			extraData.BlockProducers = blockProducers
		} else {
			extraData.CheckpointValidators = checkpointValidator
		}
	}

	// After Shillin, extraData.hasFinalityVote = 0 here as we does
	// not assemble finality vote yet. Let's wait some time for the
	// finality votes to be broadcasted around the network. The
	// finality votes are assembled later in Seal function.
	header.Extra, err = extraData.EncodeV2(c.chainConfig, header.Number)
	if err != nil {
		return err
	}

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

func (c *Consortium) processSystemTransactions(chain consensus.ChainHeaderReader, header *types.Header,
	transactOpts *consortiumCommon.ApplyTransactOpts, isFinalizeAndAssemble bool) error {

	snap, err := c.snapshot(chain, header.Number.Uint64()-1, header.ParentHash, nil)
	if err != nil {
		return err
	}

	_, _, _, contract := c.readSignerAndContract()

	// If the parent's block includes the finality votes, distribute reward for the voters
	parentNumber := new(big.Int).Sub(header.Number, common.Big1)
	if c.chainConfig.IsShillin(parentNumber) {
		parentHeader := chain.GetHeader(header.ParentHash, header.Number.Uint64()-1)
		extraData, err := finality.DecodeExtraV2(parentHeader.Extra, c.chainConfig, parentNumber)
		if err != nil {
			return err
		}
		if extraData.HasFinalityVote == 1 {
			parentSnap, err := c.snapshot(chain, parentHeader.Number.Uint64()-1, parentHeader.ParentHash, nil)
			if err != nil {
				return err
			}

			votedValidatorPositions := extraData.FinalityVotedValidators.Indices()
			var votedValidators []common.Address
			for _, position := range votedValidatorPositions {
				// The header has been verified so there must be no out of bound here
				votedValidators = append(votedValidators, parentSnap.ValidatorsWithBlsPub[position].Address)
			}

			if err := contract.FinalityReward(transactOpts, votedValidators); err != nil {
				log.Error("Failed to finality reward validator", "err", err)
				return err
			}
		}
	}

	if header.Difficulty.Cmp(diffInTurn) != 0 {
		spoiledVal := snap.supposeValidator()
		signedRecently := false
		if c.chainConfig.IsOlek(header.Number) {
			signedRecently = snap.IsRecentlySigned(spoiledVal)
		} else {
			for _, recent := range snap.Recents {
				if recent == spoiledVal {
					signedRecently = true
					break
				}
			}
		}
		if !signedRecently {
			if !isFinalizeAndAssemble {
				log.Info("Slash validator", "number", header.Number, "spoiled", spoiledVal)
			}
			if err := contract.Slash(transactOpts, spoiledVal); err != nil {
				// it is possible that slash validator failed because of the slash channel is disabled.
				log.Error("Failed to slash validator", "block hash", header.Hash(), "address", spoiledVal)
				return err
			}
		}
	}

	// Previously, we call WrapUpEpoch before SubmitBlockReward which is the wrong order.
	// We create a hardfork here to fix the contract call order.
	if c.chainConfig.IsPuffy(header.Number) {
		if err := contract.SubmitBlockReward(transactOpts); err != nil {
			log.Error("Failed to submit block reward", "err", err)
			return err
		}
	}

	if header.Number.Uint64()%c.config.EpochV2 == c.config.EpochV2-1 {
		if err := contract.WrapUpEpoch(transactOpts); err != nil {
			log.Error("Failed to wrap up epoch", "err", err)
			return err
		}
	}

	if !c.chainConfig.IsPuffy(header.Number) {
		if err := contract.SubmitBlockReward(transactOpts); err != nil {
			log.Error("Failed to submit block reward", "err", err)
			return err
		}
	}

	return nil
}

func (c *Consortium) upgradeRoninTrustedOrg(blockNumber *big.Int, state *state.StateDB) {
	// The upgrade only happens in 1 block: Miko hardfork block
	if c.chainConfig.MikoBlock != nil && c.chainConfig.MikoBlock.Cmp(blockNumber) == 0 {
		state.SetState(
			c.chainConfig.RoninTrustedOrgUpgrade.ProxyAddress,
			implementationSlot,
			c.chainConfig.RoninTrustedOrgUpgrade.ImplementationAddress.Hash(),
		)
	}
}

// Finalize implements consensus.Engine that calls three methods from smart contracts:
// - WrapUpEpoch at epoch to distribute rewards and sort the validators set
// - Slash the validator who does not sign if it is in-turn
// - SubmitBlockRewards of the current block
func (c *Consortium) Finalize(chain consensus.ChainHeaderReader, header *types.Header, state *state.StateDB, txs *[]*types.Transaction,
	uncles []*types.Header, receipts *[]*types.Receipt, systemTxs *[]*types.Transaction, internalTxs *[]*types.InternalTransaction, usedGas *uint64) error {
	_, _, signTxFn, _ := c.readSignerAndContract()
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
		isPeriodBlock := c.IsPeriodBlock(chain, header)
		checkpointValidators, blockProducers, err := c.getCheckpointValidatorsFromContract(chain, isPeriodBlock, header)
		if err != nil {
			return err
		}
		extraData, err := finality.DecodeExtraV2(header.Extra, c.chainConfig, header.Number)
		if err != nil {
			return err
		}

		// If isTripp and new period, read all validator candidates and
		// their amounts, check with stored data in header
		if c.IsTrippEffective(chain, header) {
			if len(blockProducers) != len(extraData.BlockProducers) {
				return errMismatchingValidators
			}
			for i := range blockProducers {
				if blockProducers[i] != extraData.BlockProducers[i] {
					return errMismatchingValidators
				}
			}
			if isPeriodBlock {
				if len(checkpointValidators) != len(extraData.CheckpointValidators) {
					return errMismatchingValidators
				}
				for i := range checkpointValidators {
					if checkpointValidators[i].Address != extraData.CheckpointValidators[i].Address {
						return errMismatchingValidators
					}
					if !checkpointValidators[i].BlsPublicKey.Equals(extraData.CheckpointValidators[i].BlsPublicKey) {
						return errMismatchingValidators
					}
					if checkpointValidators[i].Weight != extraData.CheckpointValidators[i].Weight {
						return errMismatchingValidators
					}
				}
			}
		} else {
			if len(checkpointValidators) != len(extraData.CheckpointValidators) {
				return errMismatchingValidators
			}

			isShillin := c.chainConfig.IsShillin(header.Number)
			for i := range checkpointValidators {
				if checkpointValidators[i].Address != extraData.CheckpointValidators[i].Address {
					return errMismatchingValidators
				}

				if isShillin {
					if !checkpointValidators[i].BlsPublicKey.Equals(extraData.CheckpointValidators[i].BlsPublicKey) {
						return errMismatchingValidators
					}
				}
			}
		}
	}

	if err := c.processSystemTransactions(chain, header, transactOpts, false); err != nil {
		return err
	}
	c.upgradeRoninTrustedOrg(header.Number, state)
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
	// No block rewards in PoA, so the state remains as is and uncles are dropped
	if txs == nil {
		txs = make([]*types.Transaction, 0)
	}
	if receipts == nil {
		receipts = make([]*types.Receipt, 0)
	}
	_, _, signTxFn, _ := c.readSignerAndContract()
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
	c.upgradeRoninTrustedOrg(header.Number, state)

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

	err := c.initContract(signer, signTxFn)
	if err != nil {
		log.Error("Failed to init system contract caller", "err", err)
	}
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
	val, signFn, _, _ := c.readSignerAndContract()

	snap, err := c.snapshot(chain, number-1, header.ParentHash, nil)
	if err != nil {
		return err
	}

	// Bail out if we're unauthorized to sign a block
	if !snap.inInValidatorSet(val) {
		return errUnauthorizedValidator
	}

	// If we're amongst the recent signers, wait for the next block
	if snap.IsRecentlySigned(val) {
		return consortiumCommon.ErrRecentlySigned
	}

	// Sweet, the protocol permits us to sign the block, wait for our time
	// After the Buba hardfork, the delay is included in header time already
	delay := time.Until(time.Unix(int64(header.Time), 0))
	if !c.chainConfig.IsBuba(block.Number()) {
		if header.Difficulty.Cmp(diffInTurn) != 0 {
			// It's not our turn explicitly to sign, delay it a bit
			wiggle := time.Duration(len(snap.validators())/2+1) * wiggleTime
			delay += time.Duration(rand.Int63n(int64(wiggle))) + wiggleTime // delay for 0.5s more

			log.Trace("Out-of-turn signing requested", "wiggle", common.PrettyDuration(wiggle))
		}
	}
	log.Info("Sealing block with", "number", number, "delay", delay, "headerDifficulty", header.Difficulty, "val", val.Hex(), "txs", len(block.Transactions()))

	// Wait until sealing is terminated or delay timeout.
	log.Trace("Waiting for slot to sign and propagate", "delay", common.PrettyDuration(delay))
	go func() {
		select {
		case <-stop:
			return
		case <-time.After(delay - assemblingFinalityVoteDuration):
			c.assembleFinalityVote(chain, header, snap)

			// Sign all the things!
			sig, err := signFn(accounts.Account{Address: val}, accounts.MimetypeConsortium, consortiumRLP(header, c.chainConfig.ChainID))
			if err != nil {
				log.Error("Failed to seal block", "err", err)
				return
			}
			copy(header.Extra[len(header.Extra)-consortiumCommon.ExtraSeal:], sig)
		}

		delay = time.Until(time.Unix(int64(header.Time), 0))
		select {
		case <-stop:
			return
		case <-time.After(delay):
		}

		select {
		case results <- block.WithSeal(header):
		default:
			log.Warn("Sealing result is not read by miner", "sealhash", calculateSealHash(header, c.chainConfig.ChainID))
		}
	}()

	return nil
}

// SealHash returns the hash of a block prior to it being sealed.
func (c *Consortium) SealHash(header *types.Header) common.Hash {
	isShillin := c.chainConfig.IsShillin(header.Number)
	if isShillin {
		// After Shillin, this consensus.SealHash function does not
		// return the real hash used for sealing because the real
		// hash changes after the FinalizeAndAssemble call. As this
		// function is used by worker only to store and look up the
		// sealing tasks, we just return the hash of header without
		// the finality vote, so this hash remains unchanged after
		// FinalizeAndAssemble call.
		copyHeader := types.CopyHeader(header)

		extraData, err := finality.DecodeExtraV2(copyHeader.Extra, c.chainConfig, header.Number)
		if err != nil {
			log.Error("Failed to decode header extra data", "err", err)
		}
		extraData.HasFinalityVote = 0
		copyHeader.Extra, err = extraData.EncodeV2(c.chainConfig, header.Number)
		if err != nil {
			log.Error("Failed to encode header extra data", "err", err)
		}
		return calculateSealHash(copyHeader, c.chainConfig.ChainID)
	} else {
		return calculateSealHash(header, c.chainConfig.ChainID)
	}
}

// Close implements consensus.Engine. It's a noop for Consortium as there are no background threads.
func (c *Consortium) Close() error {
	return nil
}

// APIs are backward compatible with the v1, so we do not to implement it again
func (c *Consortium) APIs(chain consensus.ChainHeaderReader) []rpc.API {
	return []rpc.API{
		{
			Namespace: "consortiumv2",
			Version:   "1.0",
			Service:   &consortiumV2Api{chain: chain, consortium: c},
			Public:    false,
		},
	}
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
	coinbase, _, _, _ := c.readSignerAndContract()
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
func (c *Consortium) initContract(coinbase common.Address, signTxFn consortiumCommon.SignerTxFn) error {
	if consortiumCommon.Validators != nil {
		c.contract = &consortiumCommon.MockContract{}
		return nil
	}
	var err error
	c.contract, err = consortiumCommon.NewContractIntegrator(c.chainConfig, consortiumCommon.NewConsortiumBackend(c.ethAPI), signTxFn, coinbase)
	return err
}

func (c *Consortium) readSignerAndContract() (
	common.Address,
	consortiumCommon.SignerFn,
	consortiumCommon.SignerTxFn,
	consortiumCommon.ContractInteraction,
) {
	c.lock.RLock()
	defer c.lock.RUnlock()

	return c.val, c.signFn, c.signTxFn, c.contract
}

// GetBestParentBlock goes backward in the canonical chain to find if the miner can
// create a chain which has more difficulty than current chain. In case the miner
// cannot create a better chain, this function returns the head block of current
// canonical chain.
func (c *Consortium) GetBestParentBlock(chain *core.BlockChain) (*types.Block, bool) {
	signer, _, _, _ := c.readSignerAndContract()

	currentBlock := chain.CurrentBlock()
	block := currentBlock
	prevBlock := chain.GetBlockByHash(block.ParentHash())
	diffculty := block.Difficulty().Int64()
	for diffculty < diffInTurn.Int64() {
		snap, err := c.snapshot(chain, block.NumberU64()-1, block.ParentHash(), nil)
		if err != nil {
			return currentBlock, false
		}
		// Miner can create an inturn block which helps the chain to have
		// greater diffculty
		if snap.supposeValidator() == signer {
			if !snap.IsRecentlySigned(signer) {
				return prevBlock, true
			}
		}

		block = prevBlock
		prevBlock = chain.GetBlockByHash(block.ParentHash())
		diffculty += block.Difficulty().Int64()
	}

	return currentBlock, false
}

// GetJustifiedBlock gets the fast finality justified block
func (c *Consortium) GetJustifiedBlock(chain consensus.ChainHeaderReader, blockNumber uint64, blockHash common.Hash) (uint64, common.Hash) {
	snap, err := c.snapshot(chain, blockNumber, blockHash, nil)
	if err != nil {
		log.Error("Failed to get snapshot", "err", err)
		return 0, common.Hash{}
	}

	return snap.JustifiedBlockNumber, snap.JustifiedBlockHash
}

// assembleFinalityVote collects finality votes from vote pool and assembles
// them into block header
//
// block (N) <- block (N + 1)
// Block (N) is justified means there are enough finality votes for block (N) in
// block (N + 1)
// The finality vote in block (N + 1) is verified by validator set that are able
// to produce block (N + 1) (ignoring the recently signed rule) which is in
// snapshot (N)
// So here when including the vote for header.Number - 1 into header.Number, the
// snapshot provided must be at header.Number - 1
func (c *Consortium) assembleFinalityVote(chain consensus.ChainHeaderReader, header *types.Header, snap *Snapshot) {
	if c.chainConfig.IsShillin(header.Number) {
		var (
			signatures              []blsCommon.Signature
			finalityVotedValidators finality.FinalityVoteBitSet
			finalityThreshold       int
			accumulatedVoteWeight   int
		)

		isTrippEffective := c.IsTrippEffective(chain, header)
		if isTrippEffective {
			finalityThreshold = int(math.Floor(finalityRatio*float64(consortiumCommon.MaxFinalityVotePercentage))) + 1
		} else {
			finalityThreshold = int(math.Floor(finalityRatio*float64(len(snap.ValidatorsWithBlsPub)))) + 1
		}

		// We assume the signature has been verified in vote pool
		// so we do not verify signature here
		if c.votePool != nil {
			votes := c.votePool.FetchVoteByBlockHash(header.ParentHash)
			// Before Tripp (!isTripp), every vote has the same weight so if the number of votes
			// is lower than threshold, we can short-circuit and skip all the checks
			if isTrippEffective || len(votes) >= finalityThreshold {
				for _, vote := range votes {
					publicKey, err := blst.PublicKeyFromBytes(vote.PublicKey[:])
					if err != nil {
						log.Warn("Malformed public key from vote pool", "err", err)
						continue
					}
					authorized := false
					for valPosition, validator := range snap.ValidatorsWithBlsPub {
						if publicKey.Equals(validator.BlsPublicKey) {
							authorized = true
							signature, err := blst.SignatureFromBytes(vote.Signature[:])
							if err != nil {
								log.Warn("Malformed signature from vote pool", "err", err)
								break
							}
							if finalityVotedValidators.GetBit(valPosition) != 0 {
								log.Warn("More than 1 vote from validator", "address", validator.Address.Hex(),
									"blockHash", header.Hash(), "blockNumber", header.Number)
								break
							}
							signatures = append(signatures, signature)
							finalityVotedValidators.SetBit(valPosition)
							if isTrippEffective {
								accumulatedVoteWeight += int(snap.ValidatorsWithBlsPub[valPosition].Weight)
							}
							break
						}
					}
					if !authorized {
						log.Warn("Unauthorized voter's signature from vote pool", "publicKey", hex.EncodeToString(publicKey.Marshal()))
					}
				}

				if !isTrippEffective {
					accumulatedVoteWeight = len(finalityVotedValidators.Indices())
				}
				if accumulatedVoteWeight >= finalityThreshold {
					extraData, err := finality.DecodeExtraV2(header.Extra, c.chainConfig, header.Number)
					if err != nil {
						// This should not happen
						log.Error("Failed to decode header extra data", "err", err)
						return
					}
					extraData.HasFinalityVote = 1
					extraData.FinalityVotedValidators = finalityVotedValidators
					extraData.AggregatedFinalityVotes = blst.AggregateSignatures(signatures)
					header.Extra, err = extraData.EncodeV2(c.chainConfig, header.Number)
					if err != nil {
						log.Error("Failed to encode header extra data", "err", err)
						return
					}
				}
			}
		}
	}
}

// GetFinalizedBlock gets the fast finality finalized block
func (c *Consortium) GetFinalizedBlock(
	chain consensus.ChainHeaderReader,
	headNumber uint64,
	headHash common.Hash,
) (uint64, common.Hash) {
	var (
		justifiedNumber, descendantJustifiedNumber uint64
		justifiedHash, descendantJustifiedHash     common.Hash
	)

	justifiedNumber = headNumber
	justifiedHash = headHash

	for {
		// When getting the snapshot at block N, the maximum justified number is N - 1.
		// Here, we want to check if the block at justifiedNumber - 1 is justified too.
		// So, the snapshot we need to look up is at justifiedNumber.
		justifiedNumber, justifiedHash = c.GetJustifiedBlock(chain, justifiedNumber, justifiedHash)
		if justifiedNumber == 0 {
			return 0, common.Hash{}
		}

		// Check if the block is justified and its direct descendant is also justified
		if descendantJustifiedNumber != 0 && descendantJustifiedNumber-1 == justifiedNumber {
			// Check if the justified block and its justified direct descendant are voted by the
			// same set of validators.
			// The validator set verifies finality vote for block (N) is in the snapshot (N)
			descendantSnap, err := c.snapshot(chain, descendantJustifiedNumber, descendantJustifiedHash, nil)
			if err != nil {
				return 0, common.Hash{}
			}

			snap, err := c.snapshot(chain, justifiedNumber, justifiedHash, nil)
			if err != nil {
				return 0, common.Hash{}
			}

			descendantValidator := descendantSnap.validators()
			snapValidator := snap.validators()

			if len(descendantValidator) == len(snapValidator) {
				var i int
				for i = 0; i < len(descendantValidator); i++ {
					if descendantValidator[i] != snapValidator[i] {
						break
					}
				}

				if i == len(descendantValidator) {
					return justifiedNumber, justifiedHash
				}
			}
		}

		descendantJustifiedNumber = justifiedNumber
		descendantJustifiedHash = justifiedHash
	}
}

// SetVotePool sets the finality vote pool to be used by consensus
// engine
func (c *Consortium) SetVotePool(votePool consensus.VotePool) {
	c.votePool = votePool
}

// IsFinalityVoterAt is used to check if we can vote for header.Number (the vote
// is included at header.Number + 1). As explained in assembleFinalityVote, the vote
// for header.Number is verified by the validator set at snapshot at block.Number.
// So here we get the snapshot at block.Number not at block.Number - 1
func (c *Consortium) IsFinalityVoterAt(chain consensus.ChainHeaderReader, header *types.Header) bool {
	snap, err := c.snapshot(chain, header.Number.Uint64(), header.Hash(), nil)
	if err != nil {
		return false
	}

	nodeValidator, _, _, _ := c.readSignerAndContract()
	// After Tripp, voting process is openned for a wider set of validator candidates
	// (at most 64 validators), which are stored in ValidatorsWithBlsPub of HeaderExtraData
	if c.IsTrippEffective(chain, header) {
		return snap.inVoterSet(nodeValidator)
	}
	return snap.inInValidatorSet(nodeValidator)
}

// GetFinalityVoterAt gets the validator that can vote for block number
// (the vote is included in block number + 1), so get the snapshot at
// block number
func (c *Consortium) GetFinalityVoterAt(
	chain consensus.ChainHeaderReader,
	blockNumber uint64,
	blockHash common.Hash,
) []finality.ValidatorWithBlsPub {
	snap, err := c.snapshot(chain, blockNumber, blockHash, nil)
	if err != nil {
		return nil
	}

	return snap.ValidatorsWithBlsPub
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
	pubkey, err := crypto.Ecrecover(calculateSealHash(header, chainId).Bytes(), signature)
	if err != nil {
		return common.Address{}, err
	}
	var signer common.Address
	copy(signer[:], crypto.Keccak256(pubkey[1:])[12:])

	sigcache.Add(hash, signer)
	return signer, nil
}

// calculateSealHash returns the hash of a block prior to it being sealed.
func calculateSealHash(header *types.Header, chainId *big.Int) (hash common.Hash) {
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

// getLastCheckpointHeader returns the last checkpoint header, this function is used as a fallback when we cannot
// get the snapshot to query the period number
func (c *Consortium) getLastCheckpointHeader(chain consensus.ChainHeaderReader, currentHeader *types.Header) *types.Header {
	current := currentHeader
	for {
		parentNumber := current.Number.Uint64() - 1
		parentHash := current.ParentHash
		current = chain.GetHeader(parentHash, parentNumber)
		if current == nil {
			log.Error("Failed to get block", "number", parentNumber, "hash", parentHash.Hex())
			return nil
		}

		if current.Number.Uint64()%c.config.EpochV2 == 0 {
			break
		}
	}

	return current
}

// IsPeriodBlock returns indicator whether a block is a period checkpoint block or not,
// which is the first checkpoint block (block % EpochV2 == 0) after 00:00 UTC everyday.
func (c *Consortium) IsPeriodBlock(chain consensus.ChainHeaderReader, header *types.Header) bool {
	number := header.Number.Uint64()
	if number%c.config.EpochV2 != 0 || !chain.Config().IsTripp(header.Number) {
		return false
	}
	snap, err := c.snapshot(chain, number-1, header.ParentHash, nil)

	// If error happens when derive snapshot or current period is absent, we recursively find
	// the nearest epoch block; and determine whether the header is one day ahead of that neighbor.
	if err != nil {
		log.Warn("Fail to get snapshot at", "block", number-1, "err", err)
		parent := c.getLastCheckpointHeader(chain, header)
		if parent == nil {
			return false
		}
		return uint64(header.Time/dayInSeconds) > uint64(parent.Time/dayInSeconds)
	}
	if snap.CurrentPeriod == 0 {
		return false
	}
	return uint64(header.Time/dayInSeconds) > snap.CurrentPeriod
}

// IsTrippEffective returns if the finality vote rule change is effective. This change is effective
// after the period of Tripp block not right after the Tripp block.
func (c *Consortium) IsTrippEffective(chain consensus.ChainHeaderReader, header *types.Header) bool {
	if c.chainConfig.IsTripp(header.Number) {
		if c.testTrippEffective {
			return true
		}

		// This is a small optimization. When we pass the Tripp block with so
		// many blocks, we don't need to go through checks below
		if header.Number.Uint64() > c.chainConfig.TrippBlock.Uint64()+28800 {
			return true
		}

		// If it is the checkpoint block, check its period number with the configured one
		if header.Number.Uint64()%c.chainConfig.Consortium.EpochV2 == 0 {
			return header.Time/dayInSeconds > c.chainConfig.TrippPeriod.Uint64()
		}

		// else check the period number of the last checkpoint header with the configured one
		snap, err := c.snapshot(chain, header.Number.Uint64()-1, header.ParentHash, nil)
		if err != nil {
			log.Error("Failed to get snapshot", "err", err)
			parent := c.getLastCheckpointHeader(chain, header)
			if parent == nil {
				return false
			}

			return parent.Time/dayInSeconds > c.chainConfig.TrippBlock.Uint64()
		}

		if snap.CurrentPeriod > c.chainConfig.TrippPeriod.Uint64() {
			return true
		}
	}

	return false
}
