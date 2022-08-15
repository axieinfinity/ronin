package v2

import (
	"errors"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/consensus"
	consortiumCommon "github.com/ethereum/go-ethereum/consensus/consortium/common"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/systemcontracts"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/internal/ethapi"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/rpc"
	lru "github.com/hashicorp/golang-lru"
	"math/big"
	"strings"
	"sync"
	"time"
)

const (
	inmemorySnapshots  = 128  // Number of recent vote snapshots to keep in memory
	inmemorySignatures = 4096 // Number of recent block signatures to keep in memory

	wiggleTime = 1000 * time.Millisecond // Random delay (per signer) to allow concurrent signers
)

// Consortium proof-of-authority protocol constants.
var (
	epochLength = uint64(30000) // Default number of blocks after which to checkpoint

	extraVanity = 32                     // Fixed number of extra-data prefix bytes reserved for signer vanity
	extraSeal   = crypto.SignatureLength // Fixed number of extra-data suffix bytes reserved for signer seal

	emptyNonce = hexutil.MustDecode("0x0000000000000000") // Nonce number should be empty

	uncleHash = types.CalcUncleHash(nil) // Always Keccak256(RLP([])) as uncles are meaningless outside of PoW

	diffInTurn = big.NewInt(7) // Block difficulty for in-turn signatures
	diffNoTurn = big.NewInt(3) // Block difficulty for out-of-turn signatures
)

var (
	// 100 native token
	maxSystemBalance = new(big.Int).Mul(big.NewInt(100), big.NewInt(params.Ether))

	systemContracts = map[common.Address]bool{
		common.HexToAddress(systemcontracts.ValidatorContract): true,
		common.HexToAddress(systemcontracts.SlashContract):     true,
	}
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
)

type SignerTxFn func(accounts.Account, *types.Transaction, *big.Int) (*types.Transaction, error)

func isToSystemContract(to common.Address) bool {
	return systemContracts[to]
}

type Consortium struct {
	chainConfig *params.ChainConfig
	config      *params.ConsortiumConfig // Consensus engine configuration parameters
	genesisHash common.Hash
	db          ethdb.Database // Database to store and retrieve snapshot checkpoints

	recents    *lru.ARCCache // Snapshots for recent block to speed up reorgs
	signatures *lru.ARCCache // Signatures of recent blocks to speed up mining

	val      common.Address // Ethereum address of the signing key
	signer   types.Signer
	signFn   consortiumCommon.SignerFn // Signer function to authorize hashes with
	signTxFn SignerTxFn

	lock sync.RWMutex // Protects the signer fields

	ethAPI          *ethapi.PublicBlockChainAPI
	validatorSetABI abi.ABI
	slashABI        abi.ABI
}

func New(
	chainConfig *params.ChainConfig,
	db ethdb.Database,
	ethAPI *ethapi.PublicBlockChainAPI,
	genesisHash common.Hash,
) *Consortium {
	consortiumConfig := chainConfig.Consortium

	if consortiumConfig != nil && consortiumConfig.Epoch == 0 {
		consortiumConfig.Epoch = epochLength
	}

	// Allocate the snapshot caches and create the engine
	recents, _ := lru.NewARC(inmemorySnapshots)
	signatures, _ := lru.NewARC(inmemorySignatures)
	vABI, _ := abi.JSON(strings.NewReader(validatorSetABI))
	sABI, _ := abi.JSON(strings.NewReader(slashABI))

	return &Consortium{
		chainConfig:     chainConfig,
		config:          consortiumConfig,
		genesisHash:     genesisHash,
		db:              db,
		ethAPI:          ethAPI,
		recents:         recents,
		signatures:      signatures,
		validatorSetABI: vABI,
		slashABI:        sABI,
		signer:          types.NewEIP155Signer(chainConfig.ChainID),
	}
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
	if sender == header.Coinbase && isToSystemContract(*tx.To()) && tx.GasPrice().Cmp(big.NewInt(0)) == 0 {
		return true, nil
	}
	return false, nil
}

func (c *Consortium) IsSystemContract(to *common.Address) bool {
	if to == nil {
		return false
	}
	return isToSystemContract(*to)
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
	return common.Address{}, nil
}

func (c *Consortium) VerifyHeader(chain consensus.ChainHeaderReader, header *types.Header, seal bool) error {
	return nil
}

func (c *Consortium) VerifyHeaders(chain consensus.ChainHeaderReader, headers []*types.Header, seals []bool) (chan<- struct{}, <-chan error) {
	abort := make(chan struct{})
	results := make(chan error, len(headers))

	return abort, results
}

func (c *Consortium) VerifyHeaderAndParents(chain consensus.ChainHeaderReader, header *types.Header, parents []*types.Header) error {
	return nil
}

func (c *Consortium) snapshot(chain consensus.ChainHeaderReader, number uint64, hash common.Hash, parents []*types.Header) (*Snapshot, error) {
	return nil, nil
}

func (c *Consortium) VerifyUncles(chain consensus.ChainReader, block *types.Block) error {
	return nil
}

func (c *Consortium) Prepare(chain consensus.ChainHeaderReader, header *types.Header) error {
	return nil
}

func (c *Consortium) Finalize(chain consensus.ChainHeaderReader, header *types.Header, state *state.StateDB, txs *[]*types.Transaction,
	uncles []*types.Header, receipts *[]*types.Receipt, systemTxs *[]*types.Transaction, usedGas *uint64) error {
	return nil
}

func (c *Consortium) FinalizeAndAssemble(chain consensus.ChainHeaderReader, header *types.Header, state *state.StateDB,
	txs []*types.Transaction, uncles []*types.Header, receipts []*types.Receipt) (*types.Block, error) {
	return nil, nil
}

func (c *Consortium) Delay(chain consensus.ChainReader, header *types.Header) *time.Duration {
	return nil
}

func (c *Consortium) Seal(chain consensus.ChainHeaderReader, block *types.Block, results chan<- *types.Block, stop <-chan struct{}) error {
	return nil
}

func (c *Consortium) SealHash(header *types.Header) common.Hash {
	return common.Hash{}
}

func (c *Consortium) Close() error {
	return nil
}

func (c *Consortium) APIs(chain consensus.ChainHeaderReader) []rpc.API {
	return []rpc.API{}
}

func (c *Consortium) CalcDifficulty(chain consensus.ChainHeaderReader, time uint64, parent *types.Header) *big.Int {
	return nil
}
