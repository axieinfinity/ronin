package proxy

import (
	"context"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/bloombits"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/eth/ethconfig"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/ethdb/httpdb"
	"github.com/ethereum/go-ethereum/event"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/ethereum/go-ethereum/trie"
	"math/big"
	"sync/atomic"
	"time"
)

const (
	defaultSafeBlockRange = 10
	// freezerHeaderTable indicates the name of the freezer header table.
	freezerHeaderTable = "headers"

	// freezerHashTable indicates the name of the freezer canonical hash table.
	freezerHashTable = "hashes"

	// freezerBodiesTable indicates the name of the freezer block body table.
	freezerBodiesTable = "bodies"
)

// backend implements interface ethapi.Backend which is used to init new VM
type backend struct {
	db             ethdb.Database
	stateCache     state.Database
	ethConfig      *ethconfig.Config
	currentBlock   *atomic.Value
	rpc            *ethclient.Client
	fgpClient      *ethclient.Client
	chainConfig    *params.ChainConfig
	safeBlockRange uint
}

func (b *backend) Config() *params.ChainConfig {
	return b.ChainConfig()
}

func (b *backend) GetHeaderByNumber(number uint64) *types.Header {
	header, _ := b.HeaderByNumber(context.Background(), rpc.BlockNumber(number))
	return header
}

func (b *backend) GetHeaderByHash(hash common.Hash) *types.Header {
	header, _ := b.HeaderByHash(context.Background(), hash)
	return header
}

func (b *backend) DB() ethdb.Database {
	return b.db
}

func (b *backend) StateCache() state.Database {
	return b.stateCache
}

func (b *backend) PendingBlockAndReceipts() (*types.Block, types.Receipts) {
	return nil, nil
}

func (b *backend) SubscribeInternalTransactionEvent(ch chan<- types.InternalTransaction) event.Subscription {
	return nil
}

func (b *backend) SyncProgress() ethereum.SyncProgress {
	return ethereum.SyncProgress{}
}

func (b *backend) SuggestGasTipCap(ctx context.Context) (*big.Int, error) {
	return nil, nil
}

func (b *backend) FeeHistory(ctx context.Context, blockCount int, lastBlock rpc.BlockNumber, rewardPercentiles []float64) (*big.Int, [][]*big.Int, []*big.Int, []float64, error) {
	return nil, nil, nil, nil, nil
}

func (b *backend) RPCEVMTimeout() time.Duration {
	return b.ethConfig.RPCEVMTimeout
}

func (b *backend) TxPoolContentFrom(addr common.Address) (types.Transactions, types.Transactions) {
	return nil, nil
}

func NewBackend(cfg *Config, ethConfig *ethconfig.Config) (*backend, error) {
	var db *httpdb.DB
	if cfg.Redis != nil {
		db = httpdb.NewDBWithRedis(cfg.RpcUrl, cfg.ArchiveUrl, cfg.Redis.Expiration, cfg.Redis.Options)
	} else {
		db = httpdb.NewDBWithLRU(cfg.RpcUrl, cfg.ArchiveUrl, cfg.DBCachedSize)
	}
	rpcClient, err := ethclient.Dial(cfg.RpcUrl)
	if err != nil {
		return nil, err
	}
	chainConfig, _, err := core.SetupGenesisBlockWithOverride(db, nil, nil)
	if err != nil {
		return nil, err
	}
	b := &backend{
		db:             db,
		ethConfig:      ethConfig,
		rpc:            rpcClient,
		chainConfig:    chainConfig,
		currentBlock:   &atomic.Value{},
		safeBlockRange: defaultSafeBlockRange,
		stateCache:     state.NewDatabaseWithConfig(db, &trie.Config{Cache: 16}),
	}
	if cfg.FreeGasProxy != "" {
		if b.fgpClient, err = ethclient.Dial(cfg.FreeGasProxy); err != nil {
			return nil, err
		}
	}
	if cfg.SafeBlockRange > 0 {
		b.safeBlockRange = cfg.SafeBlockRange
	}
	return b, nil
}

func (b *backend) ChainDb() ethdb.Database           { return b.db }
func (b *backend) AccountManager() *accounts.Manager { return nil }
func (b *backend) ExtRPCEnabled() bool               { return true }
func (b *backend) RPCGasCap() uint64                 { return b.ethConfig.RPCGasCap }   // global gas cap for eth_call over rpc: DoS protection
func (b *backend) RPCTxFeeCap() float64              { return b.ethConfig.RPCTxFeeCap } // global tx fee cap for all transaction related APIs
func (b *backend) UnprotectedAllowed() bool          { return false }                   // allows only for EIP155 transactions.

func (b *backend) SetHead(number uint64) {}

func (b *backend) HeaderByNumber(ctx context.Context, number rpc.BlockNumber) (*types.Header, error) {
	block, err := b.BlockByNumber(ctx, number)
	if err != nil {
		return nil, err
	}
	if block == nil {
		log.Error("[proxy][backend][HeaderByNumber] block not found", "number", number.Int64())
		return nil, errors.New(fmt.Sprintf("block %d not found", number.Int64()))
	}
	return block.Header(), nil
}

func (b *backend) HeaderByHash(ctx context.Context, hash common.Hash) (*types.Header, error) {
	block, err := b.BlockByHash(ctx, hash)
	if err != nil {
		return nil, err
	}
	return block.Header(), nil
}

func (b *backend) HeaderByNumberOrHash(ctx context.Context, blockNrOrHash rpc.BlockNumberOrHash) (*types.Header, error) {
	if blockNrOrHash.BlockNumber != nil {
		return b.HeaderByNumber(ctx, *blockNrOrHash.BlockNumber)
	}
	return b.HeaderByHash(ctx, *blockNrOrHash.BlockHash)
}

func (b *backend) CurrentHeader() *types.Header {
	block := b.CurrentBlock()
	if block == nil {
		return nil
	}
	return block.Header()
}

func (b *backend) writeAncient(block *types.Block) {
	log.Debug("[backend] start writing ancient", "number", block.NumberU64())
	header, err := rlp.EncodeToBytes(block.Header())
	if err != nil {
		log.Debug("[backend] encode header", "err", err, "block", block.NumberU64())
	} else if err = httpdb.PutAncient(b.db, freezerHeaderTable, block.NumberU64(), header); err != nil {
		log.Debug("[backend] error while saving block's header to ancient", "err", err, "block", block.NumberU64())
	}

	body, err := rlp.EncodeToBytes(block.Body())
	if err != nil {
		log.Debug("[backend] encode body", "err", err, "block", block.NumberU64())
	} else if err = httpdb.PutAncient(b.db, freezerBodiesTable, block.NumberU64(), body); err != nil {
		log.Debug("[backend] error while saving block's body to ancient", "err", err, "block", block.NumberU64())
	}

	if err = httpdb.PutAncient(b.db, freezerHashTable, block.NumberU64(), block.Hash().Bytes()); err != nil {
		log.Debug("[backend] error while saving block's hash to ancient", "err", err, "block", block.NumberU64())
	}
}

func (b *backend) writeBlock(block *types.Block) {
	// cache current block and relevant fields
	rawdb.WriteCanonicalHash(b.db, block.Hash(), block.NumberU64())
	rawdb.WriteHeaderNumber(b.db, block.Hash(), block.NumberU64())
	rawdb.WriteHeader(b.db, block.Header())
	rawdb.WriteTxLookupEntriesByBlock(b.db, block)
	b.writeAncient(block)

	// check if previous hashes were reorged or not
	go func() {
		parentHash := block.ParentHash()
		number := block.NumberU64() - 1
		// checkPoint is to make sure the loop won't loop from millions of blocks to 0
		checkPoint := block.NumberU64() - uint64(b.safeBlockRange)
		for number > checkPoint {
			log.Debug("reorg checking", "number", number)
			// loop until mismatch found or number does not exist
			hash := rawdb.ReadCanonicalHash(b.db, number)
			// there are 2 conditions:
			// - hash matches with parentHash => continue with previous block
			// - hash does not match with parentHash => it might be reorged => call get block by parentHash and end the loop to prevent overlapping writeBlock
			if parentHash.Hex() == hash.Hex() {
				prevBlock := rawdb.ReadBlock(b.db, parentHash, number)
				if prevBlock != nil {
					parentHash = prevBlock.ParentHash()
					number--
					continue
				}
			}
			// remove reorg block hash from parent
			rawdb.DeleteCanonicalHash(b.db, number)

			httpdb.RemoveAncient(b.db, freezerBodiesTable, number)
			httpdb.RemoveAncient(b.db, freezerHeaderTable, number)
			httpdb.RemoveAncient(b.db, freezerHashTable, number)

			// start getting block's data from parentHash
			_, err := b.BlockByHash(context.Background(), parentHash)
			if err != nil {
				log.Error("error while getting block in double check", "err", err, "hash", parentHash.Hex())
			}
			return
		}
	}()
}

func (b *backend) CurrentBlock() *types.Block {
	var (
		currentBlock, block *types.Block
		exist        bool
		err          error
	)
	if currentBlock, exist = b.currentBlock.Load().(*types.Block); exist {
		log.Debug("[backend][CurrentBlock] currentBlock is not nil", "height", currentBlock.NumberU64())
		// return currentBlock if block's time + block's period > now
		if currentBlock.Time()+b.ChainConfig().Consortium.Period > uint64(time.Now().Unix()) {
			return currentBlock
		}
	}
	log.Debug("calling rpc client to get current block")
	latest, err := b.rpc.BlockNumber(context.Background())
	if err != nil {
		log.Error("[backend][CurrentBlock] get latest blockNumber", "err", err)
		return currentBlock
	}
	log.Debug("[backend][CurrentBlock] getting block by latest", "number", latest)
	block, err = b.BlockByNumber(context.Background(), rpc.BlockNumber(latest))
	if block != nil {
		log.Debug("[backend][CurrentBlock] block is found, start caching", "number", latest)
		b.currentBlock.Store(block)
		b.writeBlock(block)
		return block
	}
	log.Debug("[backend][CurrentBlock] something went wrong when loading latest block", "number", latest, "err", err)
	return currentBlock
}

func (b *backend) BlockByNumber(ctx context.Context, number rpc.BlockNumber) (*types.Block, error) {
	if number == rpc.LatestBlockNumber || number == rpc.PendingBlockNumber {
		block := b.CurrentBlock()
		if block != nil {
			return block, nil
		}
		return nil, errors.New("cannot get latest block")
	}
	// getting hash from number
	hash := rawdb.ReadCanonicalHash(b.db, uint64(number))
	log.Debug("[proxy][backend][BlockByNumber] getting block from hash and number", "hash", hash.Hex(), "number", number)
	block := rawdb.ReadBlock(b.db, hash, uint64(number))
	if block != nil {
		return block, nil
	}
	return nil, errors.New(fmt.Sprintf("block %d not found", number.Int64()))
}

func (b *backend) BlockByHash(ctx context.Context, hash common.Hash) (*types.Block, error) {
	blockNumber := rawdb.ReadHeaderNumber(b.db, hash)
	if blockNumber == nil {
		return nil, errors.New(fmt.Sprintf("block not found by hash: %s", hash.Hex()))
	}
	block := rawdb.ReadBlock(b.db, hash, *blockNumber)
	if block != nil {
		return block, nil
	}
	return nil, errors.New(fmt.Sprintf("block %d not found", *blockNumber))
}

func (b *backend) BlockByNumberOrHash(ctx context.Context, blockNrOrHash rpc.BlockNumberOrHash) (*types.Block, error) {
	if blockNrOrHash.BlockHash != nil {
		return b.BlockByHash(ctx, *blockNrOrHash.BlockHash)
	}
	if blockNrOrHash.BlockNumber != nil {
		return b.BlockByNumber(ctx, *blockNrOrHash.BlockNumber)
	}
	return nil, errors.New("invalid arguments; neither block nor hash specified")
}

func (b *backend) StateAndHeaderByNumber(ctx context.Context, number rpc.BlockNumber) (*state.StateDB, *types.Header, error) {
	header, err := b.HeaderByNumber(ctx, number)
	if err != nil {
		return nil, nil, err
	}
	if header == nil {
		return nil, nil, errors.New("header not found")
	}
	stateDb, err := state.New(header.Root, b.stateCache, nil)
	return stateDb, header, err
}

func (b *backend) StateAndHeaderByNumberOrHash(ctx context.Context, blockNrOrHash rpc.BlockNumberOrHash) (*state.StateDB, *types.Header, error) {
	if blockNr, ok := blockNrOrHash.Number(); ok {
		return b.StateAndHeaderByNumber(ctx, blockNr)
	}
	if hash, ok := blockNrOrHash.Hash(); ok {
		header, err := b.HeaderByHash(ctx, hash)
		if err != nil {
			return nil, nil, err
		}
		if header == nil {
			return nil, nil, errors.New("header for hash not found")
		}
		if blockNrOrHash.RequireCanonical && rawdb.ReadCanonicalHash(b.db, header.Number.Uint64()) != hash {
			return nil, nil, errors.New("hash is not currently canonical")
		}
		stateDb, err := state.New(header.Root, b.stateCache, nil)
		return stateDb, header, err
	}
	return nil, nil, errors.New("invalid arguments; neither block nor hash specified")
}

func (b *backend) GetReceipts(ctx context.Context, hash common.Hash) (types.Receipts, error) {
	block, err := b.BlockByHash(ctx, hash)
	if err != nil {
		return nil, err
	}
	receipts := rawdb.ReadReceipts(b.db, block.Hash(), block.NumberU64(), b.ChainConfig())
	if receipts != nil {
		return receipts, nil
	}
	log.Warn(fmt.Sprintf("[proxy][backend] receipts not found at hash:%s and number:%d", block.Hash().Hex(), block.NumberU64()))
	return nil, nil
}

func (b *backend) GetTd(ctx context.Context, hash common.Hash) *big.Int {
	block, err := b.BlockByHash(ctx, hash)
	if err != nil {
		log.Error("[GetTd] error while getting block by hash", "err", err, "hash", hash.Hex())
		return nil
	}
	if block == nil {
		return nil
	}
	return rawdb.ReadTd(b.db, hash, block.NumberU64())
}

func (b *backend) GetEVM(ctx context.Context, msg core.Message, state *state.StateDB, header *types.Header, vmConfig *vm.Config) (*vm.EVM, func() error, error) {
	vmError := func() error { return nil }
	if vmConfig == nil {
		vmConfig = &vm.Config{
			EnablePreimageRecording: b.ethConfig.EnablePreimageRecording,
		}
	}
	txContext := core.NewEVMTxContext(msg)
	blockContext := core.NewEVMBlockContext(header, b, nil)
	return vm.NewEVM(blockContext, txContext, state, b.ChainConfig(), *vmConfig), vmError, nil
}

func (b *backend) SendTx(ctx context.Context, signedTx *types.Transaction) error {
	return b.rpc.SendTransaction(ctx, signedTx)
}

func (b *backend) GetTransaction(ctx context.Context, txHash common.Hash) (*types.Transaction, common.Hash, uint64, uint64, error) {
	blockNumber := rawdb.ReadTxLookupEntry(b.db, txHash)
	if blockNumber == nil {
		return nil, common.Hash{}, 0, 0, errors.New(fmt.Sprintf("block not found with txHash %s", txHash.Hex()))
	}
	block, err := b.BlockByNumber(ctx, rpc.BlockNumber(*blockNumber))
	if err != nil {
		return nil, common.Hash{}, 0, 0, err
	}
	for txIndex, tx := range block.Transactions() {
		if tx.Hash() == txHash {
			return tx, block.Hash(), block.NumberU64(), uint64(txIndex), nil
		}
	}
	return nil, common.Hash{}, 0, 0, errors.New(fmt.Sprintf("Transaction not found at block:%d hash:%s tx:%s", block.NumberU64(), block.Hash().Hex(), txHash.Hex()))
}

func (b *backend) GetPoolTransactions() (types.Transactions, error) { return nil, nil }

func (b *backend) GetPoolTransaction(txHash common.Hash) *types.Transaction { return nil }

func (b *backend) GetPoolNonce(ctx context.Context, addr common.Address) (uint64, error) {
	return 0, nil
}

func (b *backend) Stats() (pending int, queued int) { return -1, -1 }

func (b *backend) TxPoolContent() (map[common.Address]types.Transactions, map[common.Address]types.Transactions) {
	return nil, nil
}

func (b *backend) SubscribeNewTxsEvent(chan<- core.NewTxsEvent) event.Subscription { return nil }

func (b *backend) BloomStatus() (uint64, uint64) { return 0, 0 }

func (b *backend) GetLogs(ctx context.Context, blockHash common.Hash) ([][]*types.Log, error) {
	receipts, err := b.GetReceipts(ctx, blockHash)
	if err != nil {
		return nil, err
	}
	if receipts == nil {
		return nil, nil
	}
	logs := make([][]*types.Log, len(receipts))
	for i, receipt := range receipts {
		logs[i] = receipt.Logs
	}
	return logs, nil
}

func (b *backend) ChainConfig() *params.ChainConfig {
	return b.chainConfig
}

func (b *backend) Engine() consensus.Engine {
	return nil
}

func (b *backend) GetHeader(hash common.Hash, number uint64) *types.Header {
	return rawdb.ReadHeader(b.db, hash, number)
}

func (b *backend) ServiceFilter(ctx context.Context, session *bloombits.MatcherSession) {}
func (b *backend) SubscribeChainEvent(ch chan<- core.ChainEvent) event.Subscription     { return nil }
func (b *backend) SubscribeChainHeadEvent(ch chan<- core.ChainHeadEvent) event.Subscription {
	return nil
}
func (b *backend) SubscribeChainSideEvent(ch chan<- core.ChainSideEvent) event.Subscription {
	return nil
}
func (b *backend) SubscribeLogsEvent(ch chan<- []*types.Log) event.Subscription        { return nil }
func (b *backend) SubscribePendingLogsEvent(ch chan<- []*types.Log) event.Subscription { return nil }
func (b *backend) SubscribeRemovedLogsEvent(ch chan<- core.RemovedLogsEvent) event.Subscription {
	return nil
}
func (b *backend) SubscribeReorgEvent(ch chan<- core.ReorgEvent) event.Subscription { return nil }
