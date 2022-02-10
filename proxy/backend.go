package proxy

import (
	"context"
	"errors"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus"
	"github.com/ethereum/go-ethereum/consensus/consortium"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/bloombits"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/eth/ethconfig"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/event"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/rpc"
	lru "github.com/hashicorp/golang-lru"
	"math/big"
	"sync/atomic"
	"time"
)

const (
	blocksCacheLimit   = 128
	receiptsCacheLimit = 128
)

// backend implements interface ethapi.Backend which is used to init new VM
type backend struct {
	db        ethdb.Database
	ethConfig *ethconfig.Config
	hc        *core.HeaderChain
	// Cache for the most recent receipts per block
	receiptsCache *lru.Cache
	// Cache for the most recent block
	blocksCache *lru.Cache
	// Cache for most recent block number
	numbersCache *lru.Cache
	currentBlock *atomic.Value
	client      *ethclient.Client
	fgpClient   *ethclient.Client
	chainConfig *params.ChainConfig
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

func newBackend(db ethdb.Database, ethConfig *ethconfig.Config, rpcUrl, fgp string) (*backend, error) {
	receiptsCache, _ := lru.New(receiptsCacheLimit)
	blocksCache, _ := lru.New(blocksCacheLimit)
	numbersCache, _ := lru.New(blocksCacheLimit)
	client, err := ethclient.Dial(rpcUrl)
	if err != nil {
		return nil, err
	}
	chainConfig, _, err := core.SetupGenesisBlockWithOverride(db, nil, nil)
	if err != nil {
		return nil, err
	}
	b := &backend{
		db: db,
		ethConfig: ethConfig,
		receiptsCache: receiptsCache,
		blocksCache: blocksCache,
		numbersCache: numbersCache,
		client: client,
		chainConfig: chainConfig,
		currentBlock: &atomic.Value{},
	}
	if fgp != "" {
		if b.fgpClient, err = ethclient.Dial(fgp); err != nil {
			return nil, err
		}
	}
	b.hc, err = core.NewHeaderChain(db, b.ChainConfig(), &consortium.Consortium{}, nil)
	if err != nil {
		return nil, err
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

func (b *backend) cacheBlock(block *types.Block) {
	b.blocksCache.Add(block.Hash(), block)
	b.numbersCache.Add(block.NumberU64(), block.Hash())

	// check if previous hashes were reorged or not
	go func() {
		parentHash := block.ParentHash()
		number := block.NumberU64() - 1
		for number > 0 {
			// loop until mismatch found or number does not exist
			hash, ok := b.numbersCache.Get(number)
			if ok {
				// there are 2 conditions:
				// - hash matches with parentHash => continue with previous block
				// - hash does not match with parentHash => it might be reorged => call get block by parentHash and end the loop to prevent overlapping cacheBlock
				if parentHash.Hex() == hash.(common.Hash).Hex() {
					prevBlock, exist := b.blocksCache.Get(parentHash)
					if exist {
						parentHash = prevBlock.(*types.Block).ParentHash()
						number--
						continue
					}
				}
				// remove canonical cache from db if any
				rawdb.DeleteCanonicalHash(b.db, number)
				_, err := b.BlockByHash(context.Background(), parentHash)
				if err != nil {
					log.Error("error while getting block in double check", "err", err, "hash", parentHash.Hex())
				}
			}
			// block height does not exist in cache
			return
		}
	}()
}

func (b *backend) CurrentBlock() *types.Block {
	currentBlock := b.currentBlock.Load()
	if currentBlock != nil && currentBlock.(*types.Block).Time() + b.ChainConfig().Consortium.Period > uint64(time.Now().Unix()) {
		return currentBlock.(*types.Block)
	}
	log.Trace("calling rpc client to get current block")
	block, err := b.client.BlockByNumber(context.Background(), nil)
	if err != nil {
		return nil
	}
	b.currentBlock.Store(block)
	b.cacheBlock(block)
	return block
}

func (b *backend) BlockByNumber(ctx context.Context, number rpc.BlockNumber) (*types.Block, error) {
	if hash, ok := b.numbersCache.Get(uint64(number)); ok {
		if block, exist := b.blocksCache.Get(hash.(common.Hash)); exist {
			return block.(*types.Block), nil
		}
	}
	log.Trace("block number cannot be found in cache, calling rpc client", "number", number)
	block, err := b.client.BlockByNumber(ctx, big.NewInt(int64(number)))
	if err != nil {
		return nil, err
	}
	b.cacheBlock(block)
	return block, nil
}

func (b *backend) BlockByHash(ctx context.Context, hash common.Hash) (*types.Block, error) {
	if block, ok := b.blocksCache.Get(hash); ok {
		return block.(*types.Block), nil
	}
	log.Trace("block hash cannot be found in cache, calling rpc client", "hash", hash.Hex())
	block, err := b.client.BlockByHash(ctx, hash)
	if err != nil {
		return nil, err
	}
	b.cacheBlock(block)
	return block, nil
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
	stateDb, err := state.New(header.Root, state.NewDatabaseWithConfig(b.db, nil), nil)
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
		if blockNrOrHash.RequireCanonical && b.hc.GetCanonicalHash(header.Number.Uint64()) != hash {
			return nil, nil, errors.New("hash is not currently canonical")
		}
		stateDb, err := state.New(header.Root, state.NewDatabaseWithConfig(b.db, nil), nil)
		return stateDb, header, err
	}
	return nil, nil, errors.New("invalid arguments; neither block nor hash specified")
}

func (b *backend) GetReceipts(ctx context.Context, hash common.Hash) (types.Receipts, error) {
	if receipts, ok := b.receiptsCache.Get(hash); ok {
		return receipts.(types.Receipts), nil
	}
	number := rawdb.ReadHeaderNumber(b.db, hash)
	if number == nil {
		return nil, nil
	}
	receipts := rawdb.ReadReceipts(b.db, hash, *number, b.ChainConfig())
	if receipts == nil {
		return nil, nil
	}
	b.receiptsCache.Add(hash, receipts)
	return receipts, nil
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
	return b.hc.GetTd(hash, block.NumberU64())
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
	return b.client.SendTransaction(ctx, signedTx)
}

func (b *backend) GetTransaction(ctx context.Context, txHash common.Hash) (*types.Transaction, common.Hash, uint64, uint64, error) {
	tx, blockHash, blockNumber, index := rawdb.ReadTransaction(b.db, txHash)
	return tx, blockHash, blockNumber, index, nil
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
	return consortium.New(&params.ConsortiumConfig{}, b.db)
}

func (b *backend) GetHeader(hash common.Hash, number uint64) *types.Header {
	return b.hc.GetHeader(hash, number)
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
