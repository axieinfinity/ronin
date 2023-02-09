package v2

import (
	"bytes"
	"encoding/json"
	"errors"
	"math/big"
	"sort"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus"
	v1 "github.com/ethereum/go-ethereum/consensus/consortium/v1"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/internal/ethapi"
	"github.com/ethereum/go-ethereum/params"
	lru "github.com/hashicorp/golang-lru"
)

// Snapshot is the state of the authorization validators at a given point in time.
type Snapshot struct {
	chainConfig *params.ChainConfig
	config      *params.ConsortiumConfig // Consensus engine parameters to fine tune behavior
	ethAPI      *ethapi.PublicBlockChainAPI
	sigCache    *lru.ARCCache // Cache of recent block signatures to speed up ecrecover

	Number     uint64                      `json:"number"`     // Block number where the snapshot was created
	Hash       common.Hash                 `json:"hash"`       // Block hash where the snapshot was created
	Validators map[common.Address]struct{} `json:"validators"` // Set of authorized validators at this moment
	Recents    map[uint64]common.Address   `json:"recents"`    // Set of recent validators for spam protections
}

// validatorsAscending implements the sort interface to allow sorting a list of addresses
type validatorsAscending []common.Address

func (s validatorsAscending) Len() int           { return len(s) }
func (s validatorsAscending) Less(i, j int) bool { return bytes.Compare(s[i][:], s[j][:]) < 0 }
func (s validatorsAscending) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

// newSnapshot creates a new snapshot with the specified startup parameters. This
// method does not initialize the set of recent validators, so only ever use if for
// the genesis block
func newSnapshot(chainConfig *params.ChainConfig, config *params.ConsortiumConfig, sigcache *lru.ARCCache, number uint64, hash common.Hash, validators []common.Address, ethAPI *ethapi.PublicBlockChainAPI) *Snapshot {
	snap := &Snapshot{
		chainConfig: chainConfig,
		config:      config,
		ethAPI:      ethAPI,
		sigCache:    sigcache,
		Number:      number,
		Hash:        hash,
		Recents:     make(map[uint64]common.Address),
		Validators:  make(map[common.Address]struct{}),
	}
	for _, v := range validators {
		snap.Validators[v] = struct{}{}
	}
	return snap
}

// loadSnapshot loads an existing snapshot from the database.
func loadSnapshot(
	config *params.ConsortiumConfig,
	sigcache *lru.ARCCache,
	db ethdb.Database,
	hash common.Hash,
	ethAPI *ethapi.PublicBlockChainAPI,
	chainConfig *params.ChainConfig,
) (*Snapshot, error) {
	blob, err := db.Get(append([]byte("consortium-"), hash[:]...))
	if err != nil {
		return nil, err
	}
	snap := new(Snapshot)
	if err := json.Unmarshal(blob, snap); err != nil {
		return nil, err
	}
	snap.config = config
	snap.sigCache = sigcache
	snap.ethAPI = ethAPI
	snap.chainConfig = chainConfig

	return snap, nil
}

// store inserts the snapshot into the database.
func (s *Snapshot) store(db ethdb.Database) error {
	blob, err := json.Marshal(s)
	if err != nil {
		return err
	}
	return db.Put(append([]byte("consortium-"), s.Hash[:]...), blob)
}

// copy creates a deep copy of the snapshot.
func (s *Snapshot) copy() *Snapshot {
	cpy := &Snapshot{
		chainConfig: s.chainConfig,
		config:      s.config,
		ethAPI:      s.ethAPI,
		sigCache:    s.sigCache,
		Number:      s.Number,
		Hash:        s.Hash,
		Validators:  make(map[common.Address]struct{}),
		Recents:     make(map[uint64]common.Address),
	}

	for v := range s.Validators {
		cpy.Validators[v] = struct{}{}
	}
	for block, v := range s.Recents {
		cpy.Recents[block] = v
	}
	return cpy
}

// apply creates a new authorization snapshot by applying the given headers to
// the original one.
func (s *Snapshot) apply(headers []*types.Header, chain consensus.ChainHeaderReader, parents []*types.Header, chainId *big.Int) (*Snapshot, error) {
	// Allow passing in no headers for cleaner code
	if len(headers) == 0 {
		return s, nil
	}
	// Sanity check that the headers can be applied
	for i := 0; i < len(headers)-1; i++ {
		if headers[i+1].Number.Uint64() != headers[i].Number.Uint64()+1 {
			return nil, errOutOfRangeChain
		}
		if !bytes.Equal(headers[i+1].ParentHash.Bytes(), headers[i].Hash().Bytes()) {
			return nil, errBlockHashInconsistent
		}
	}
	if headers[0].Number.Uint64() != s.Number+1 {
		return nil, errOutOfRangeChain
	}
	if !bytes.Equal(headers[0].ParentHash.Bytes(), s.Hash.Bytes()) {
		return nil, errBlockHashInconsistent
	}
	// Iterate through the headers and create a new snapshot
	snap := s.copy()

	// Number of consecutive blocks out of which a validator may only sign one.
	// Must be len(snap.Validators)/2 + 1 to enforce majority consensus on a chain
	for _, header := range headers {
		number := header.Number.Uint64()
		// Delete the oldest validators from the recent list to allow it signing again
		if limit := uint64(len(snap.Validators)/2 + 1); number >= limit {
			delete(snap.Recents, number-limit)
		}
		// Resolve the authorization key and check against signers
		var (
			validator common.Address
			err       error
		)
		// If the headers come from v1 the block hash function does not include chainId,
		// we need to use the correct ecrecover function the get the correct signer
		if !snap.chainConfig.IsConsortiumV2(header.Number) {
			validator, err = v1.Ecrecover(header, s.sigCache)
		} else {
			validator, err = ecrecover(header, s.sigCache, chainId)
		}
		if err != nil {
			return nil, err
		}
		if _, ok := snap.Validators[validator]; !ok {
			return nil, errUnauthorizedValidator
		}
		for _, recent := range snap.Recents {
			if recent == validator {
				return nil, errRecentlySigned
			}
		}
		snap.Recents[number] = validator
		// Change the validator set base on the size of the validators set
		if number > 0 && number%s.config.EpochV2 == uint64(len(snap.Validators)/2) {
			// Get the most recent checkpoint header
			checkpointHeader := FindAncientHeader(header, uint64(len(snap.Validators)/2), chain, parents)
			if checkpointHeader == nil {
				return nil, consensus.ErrUnknownAncestor
			}

			validatorBytes := checkpointHeader.Extra[extraVanity : len(checkpointHeader.Extra)-extraSeal]
			// Get validator set from headers and use that for new validator set
			newValArr, err := ParseValidators(validatorBytes)
			if err != nil {
				return nil, err
			}
			newVals := make(map[common.Address]struct{}, len(newValArr))
			for _, val := range newValArr {
				newVals[val] = struct{}{}
			}
			oldLimit := len(snap.Validators)/2 + 1
			newLimit := len(newVals)/2 + 1
			if newLimit < oldLimit {
				for i := 0; i < oldLimit-newLimit; i++ {
					delete(snap.Recents, number-uint64(newLimit)-uint64(i))
				}
			}
			snap.Validators = newVals
		}
	}
	snap.Number += uint64(len(headers))
	snap.Hash = headers[len(headers)-1].Hash()
	return snap, nil
}

// validators retrieves the list of validators in ascending order.
func (s *Snapshot) validators() []common.Address {
	validators := make([]common.Address, 0, len(s.Validators))
	for v := range s.Validators {
		validators = append(validators, v)
	}
	sort.Sort(validatorsAscending(validators))
	return validators
}

// inturn returns if a validator at a given block height is in-turn or not.
func (s *Snapshot) inturn(validator common.Address) bool {
	validators := s.validators()
	offset := (s.Number + 1) % uint64(len(validators))
	return validators[offset] == validator
}

// sealableValidators finds the validators that are not in recent sign list, which mean they can seal
// a new block.
// This function returns the position of input validator in the sealable validators list and the length
// of that list. In case the input validator is not in sealable validators list, position is -1
func (s *Snapshot) sealableValidators(validator common.Address) (position, numOfSealableValidators int) {
	validators := s.validators()
	sealable := make([]common.Address, 0, len(validators))
	for _, val := range validators {
		allowToSeal := true
		for seen, recent := range s.Recents {
			if recent == val {
				if limit := uint64(len(validators)/2 + 1); seen > s.Number+1-limit {
					allowToSeal = false
					break
				}
			}
		}
		if allowToSeal {
			sealable = append(sealable, val)
		}
	}

	numOfSealableValidators = len(sealable)
	for i, sealableVal := range sealable {
		if validator == sealableVal {
			return i, numOfSealableValidators
		}
	}

	return -1, numOfSealableValidators
}

// supposeValidator returns the in-turn validator at a given block height
func (s *Snapshot) supposeValidator() common.Address {
	validators := s.validators()
	index := (s.Number + 1) % uint64(len(validators))
	return validators[index]
}

// ParseValidators retrieves the list of validators
func ParseValidators(validatorsBytes []byte) ([]common.Address, error) {
	if len(validatorsBytes)%validatorBytesLength != 0 {
		return nil, errors.New("invalid validators bytes")
	}
	n := len(validatorsBytes) / validatorBytesLength
	result := make([]common.Address, n)
	for i := 0; i < n; i++ {
		address := make([]byte, validatorBytesLength)
		copy(address, validatorsBytes[i*validatorBytesLength:(i+1)*validatorBytesLength])
		result[i] = common.BytesToAddress(address)
	}
	return result, nil
}

// FindAncientHeader finds the most recent checkpoint header
// Travel through the candidateParents to find the ancient header.
// If all headers in candidateParents have the number is larger than the header number,
// the search function will return the index, but it is not valid if we check with the
// header since the number and hash is not equals. The candidateParents is
// only available when it downloads blocks from the network.
// Otherwise, the candidateParents is nil, and it will be found by header hash and number.
func FindAncientHeader(header *types.Header, ite uint64, chain consensus.ChainHeaderReader, candidateParents []*types.Header) *types.Header {
	ancient := header
	for i := uint64(1); i <= ite; i++ {
		parentHash := ancient.ParentHash
		parentHeight := ancient.Number.Uint64() - 1
		found := false
		if len(candidateParents) > 0 {
			index := sort.Search(len(candidateParents), func(i int) bool {
				return candidateParents[i].Number.Uint64() >= parentHeight
			})
			if index < len(candidateParents) && candidateParents[index].Number.Uint64() == parentHeight &&
				candidateParents[index].Hash() == parentHash {
				ancient = candidateParents[index]
				found = true
			}
		}
		if !found {
			ancient = chain.GetHeader(parentHash, parentHeight)
			found = true
		}
		if ancient == nil || !found {
			return nil
		}
	}
	return ancient
}
