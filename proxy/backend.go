package proxy

import (
	"context"
	"errors"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/consensus"
	"github.com/ethereum/go-ethereum/consensus/consortium"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/bloombits"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/eth/downloader"
	"github.com/ethereum/go-ethereum/eth/ethconfig"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/ethdb/httpdb"
	"github.com/ethereum/go-ethereum/event"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/ethereum/go-ethereum/trie"
	lru "github.com/hashicorp/golang-lru"
	"math/big"
)

const receiptsCacheLimit = 32

// RPCBlock represents a block that will serialize to the RPC representation of a block
type RPCBlock struct {
	Number           string            `json:"number"`
	Hash             string            `json:"hash"`
	ParentHash       string            `json:"parentHash"`
	Nonce            string            `json:"nonce"`
	MixHash          string            `json:"mixHash"`
	LogsBloom        string            `json:"logsBloom"`
	StateRoot        string            `json:"stateRoot"`
	Miner            string            `json:"coinbase"`
	Difficulty       string            `json:"difficulty"`
	ExtraData        string            `json:"extraData"`
	Size             string            `json:"size"`
	GasLimit         string            `json:"gasLimit"`
	GasUsed          string            `json:"gasUsed"`
	TimeStamp        string            `json:"timestamp"`
	TransactionsRoot string            `json:"transactionsRoot"`
	ReceiptsRoot     string            `json:"receiptsRoot"`
	Uncles           []string          `json:"uncles"`
	Sha3Uncles       string            `json:"sha3Uncles"`
	Transactions     []*RPCTransaction `json:"transactions"`
}

// RPCTransaction represents a transaction that will serialize to the RPC representation of a transaction
type RPCTransaction struct {
	BlockHash        string `json:"blockHash"`
	BlockNumber      string `json:"blockNumber"`
	From             string `json:"from"`
	Gas              string `json:"gas"`
	GasPrice         string `json:"gasPrice"`
	Hash             string `json:"hash"`
	Input            string `json:"input"`
	Nonce            string `json:"nonce"`
	To               string `json:"to"`
	TransactionIndex string `json:"transactionIndex"`
	Value            string `json:"value"`
	V                string `json:"v"`
	R                string `json:"r"`
	S                string `json:"s"`
}

// backend implements interface ethapi.Backend which is used to init new VM
type backend struct {
	db            ethdb.Database
	ethConfig     *ethconfig.Config
	hc            *core.HeaderChain
	receiptsCache *lru.Cache // Cache for the most recent receipts per block
	client        *rpc.Client
}

func newBackend(db *httpdb.DB, ethConfig *ethconfig.Config, rpcUrl string) (*backend, error) {
	receiptsCache, _ := lru.New(receiptsCacheLimit)
	client, err := rpc.DialHTTP(rpcUrl)
	if err != nil {
		return nil, err
	}
	return &backend{db: db, ethConfig: ethConfig, receiptsCache: receiptsCache, client: client}, nil
}

func (b *backend) getLatestBlockNumber() (uint64, error) {
	block, err := b.getBlockByNumber(-1)
	if err != nil {
		return 0, err
	}
	return block.NumberU64(), nil
}

func (b *backend) getBlockByNumber(num rpc.BlockNumber) (*types.Block, error) {
	var (
		block    RPCBlock
		blockNum string
	)
	if num == rpc.LatestBlockNumber {
		blockNum = "latest"
	} else {
		blockNum = hexutil.EncodeUint64(uint64(num.Int64()))
	}
	if err := b.client.Call(&block, "eth_getBlockByNumber", blockNum, true); err != nil {
		return nil, err
	}
	return b.toBlock(block)
}

func (b *backend) toBlock(rpcBlock RPCBlock) (*types.Block, error) {
	blockNumber, err := hexutil.DecodeUint64(rpcBlock.Number)
	if err != nil {
		return nil, err
	}
	hash := common.HexToHash(rpcBlock.Hash)
	header := b.GetHeader(hash, blockNumber)
	receipts, err := b.GetReceipts(context.Background(), hash)
	if err != nil {
		return nil, err
	}
	txs := make(types.Transactions, 0)
	for _, tx := range rpcBlock.Transactions {
		nonce, err := hexutil.DecodeUint64(tx.Nonce)
		if err != nil {
			return nil, err
		}
		gas, err := hexutil.DecodeUint64(tx.Gas)
		if err != nil {
			return nil, err
		}
		to := common.HexToAddress(tx.To)
		txs = append(txs, types.NewTx(&types.LegacyTx{
			Nonce:    nonce,
			To:       &to,
			Value:    hexutil.MustDecodeBig(tx.Value),
			Gas:      gas,
			GasPrice: hexutil.MustDecodeBig(tx.GasPrice),
			Data:     hexutil.MustDecode(tx.Input),
		}))
	}
	return types.NewBlock(header, txs, nil, receipts, new(trie.Trie)), nil
}

func (b *backend) getBlockByHash(hash common.Hash) (*types.Block, error) {
	var block RPCBlock
	if err := b.client.Call(&block, "eth_getBlockByHash", hash.Hex(), true); err != nil {
		return nil, err
	}
	return b.toBlock(block)
}

func (b *backend) Downloader() *downloader.Downloader                 { return nil }
func (b *backend) SuggestPrice(ctx context.Context) (*big.Int, error) { return nil, nil }
func (b *backend) ChainDb() ethdb.Database                            { return b.db }
func (b *backend) AccountManager() *accounts.Manager                  { return nil }
func (b *backend) ExtRPCEnabled() bool                                { return true }
func (b *backend) RPCGasCap() uint64                                  { return b.ethConfig.RPCGasCap }   // global gas cap for eth_call over rpc: DoS protection
func (b *backend) RPCTxFeeCap() float64                               { return b.ethConfig.RPCTxFeeCap } // global tx fee cap for all transaction related APIs
func (b *backend) UnprotectedAllowed() bool                           { return false }                   // allows only for EIP155 transactions.

// Blockchain API
func (b *backend) SetHead(number uint64) {}
func (b *backend) HeaderByNumber(ctx context.Context, number rpc.BlockNumber) (*types.Header, error) {
	// Pending block is only known by the miner
	if number == rpc.PendingBlockNumber {
		return nil, nil
	}
	// Otherwise resolve and return the block
	if number == rpc.LatestBlockNumber {
		num, err := b.getLatestBlockNumber()
		if err != nil {
			return nil, err
		}
		number = rpc.BlockNumber(num)
	}
	return b.hc.GetHeaderByNumber(uint64(number)), nil
}
func (b *backend) HeaderByHash(ctx context.Context, hash common.Hash) (*types.Header, error) {
	return b.hc.GetHeaderByHash(hash), nil
}
func (b *backend) HeaderByNumberOrHash(ctx context.Context, blockNrOrHash rpc.BlockNumberOrHash) (*types.Header, error) {
	if blockNrOrHash.BlockNumber != nil {
		return b.HeaderByNumber(ctx, *blockNrOrHash.BlockNumber)
	}
	return b.HeaderByHash(ctx, *blockNrOrHash.BlockHash)
}
func (b *backend) CurrentHeader() *types.Header {
	num, err := b.getLatestBlockNumber()
	if err != nil {
		return nil
	}
	return b.hc.GetHeaderByNumber(num)
}
func (b *backend) CurrentBlock() *types.Block {
	block, err := b.getBlockByNumber(-1)
	if err != nil {
		return nil
	}
	return block
}
func (b *backend) BlockByNumber(ctx context.Context, number rpc.BlockNumber) (*types.Block, error) {
	return b.getBlockByNumber(number)
}
func (b *backend) BlockByHash(ctx context.Context, hash common.Hash) (*types.Block, error) {
	return b.getBlockByHash(hash)
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
	receipts := rawdb.ReadReceipts(b.db, hash, *number, &params.ChainConfig{})
	if receipts == nil {
		return nil, nil
	}
	b.receiptsCache.Add(hash, receipts)
	return receipts, nil
}
func (b *backend) GetTd(ctx context.Context, hash common.Hash) *big.Int {
	number := b.hc.GetBlockNumber(hash)
	if number == nil {
		return nil
	}
	return b.hc.GetTd(hash, *number)
}
func (b *backend) GetEVM(ctx context.Context, msg core.Message, state *state.StateDB, header *types.Header, vmConfig *vm.Config) (*vm.EVM, func() error, error) {
	vmError := func() error { return nil }
	if vmConfig == nil {
		vmConfig = &vm.Config{
			EnablePreimageRecording: b.ethConfig.EnablePreimageRecording,
			EWASMInterpreter:        b.ethConfig.EWASMInterpreter,
			EVMInterpreter:          b.ethConfig.EVMInterpreter,
		}
	}
	txContext := core.NewEVMTxContext(msg)
	blockContext := core.NewEVMBlockContext(header, b, nil)
	return vm.NewEVM(blockContext, txContext, state, &params.ChainConfig{}, *vmConfig), vmError, nil
}
func (b *backend) SubscribeChainEvent(ch chan<- core.ChainEvent) event.Subscription { return nil }
func (b *backend) SubscribeChainHeadEvent(ch chan<- core.ChainHeadEvent) event.Subscription {
	return nil
}
func (b *backend) SubscribeChainSideEvent(ch chan<- core.ChainSideEvent) event.Subscription {
	return nil
}

// Transaction pool API
func (b *backend) SendTx(ctx context.Context, signedTx *types.Transaction) error { return nil }
func (b *backend) GetTransaction(ctx context.Context, txHash common.Hash) (*types.Transaction, common.Hash, uint64, uint64, error) {
	tx, blockHash, blockNumber, index := rawdb.ReadTransaction(b.db, txHash)
	return tx, blockHash, blockNumber, index, nil
}
func (b *backend) GetPoolTransactions() (types.Transactions, error)         { return nil, nil }
func (b *backend) GetPoolTransaction(txHash common.Hash) *types.Transaction { return nil }
func (b *backend) GetPoolNonce(ctx context.Context, addr common.Address) (uint64, error) {
	return 0, nil
}
func (b *backend) Stats() (pending int, queued int) { return -1, -1 }
func (b *backend) TxPoolContent() (map[common.Address]types.Transactions, map[common.Address]types.Transactions) {
	return nil, nil
}
func (b *backend) SubscribeNewTxsEvent(chan<- core.NewTxsEvent) event.Subscription { return nil }

// Filter API
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
func (b *backend) ServiceFilter(ctx context.Context, session *bloombits.MatcherSession) {}
func (b *backend) SubscribeLogsEvent(ch chan<- []*types.Log) event.Subscription         { return nil }
func (b *backend) SubscribePendingLogsEvent(ch chan<- []*types.Log) event.Subscription  { return nil }
func (b *backend) SubscribeRemovedLogsEvent(ch chan<- core.RemovedLogsEvent) event.Subscription {
	return nil
}
func (b *backend) SubscribeReorgEvent(ch chan<- core.ReorgEvent) event.Subscription { return nil }

func (b *backend) ChainConfig() *params.ChainConfig { return &params.ChainConfig{} }
func (b *backend) Engine() consensus.Engine {
	return consortium.New(&params.ConsortiumConfig{}, b.db)
}
func (b *backend) GetHeader(hash common.Hash, number uint64) *types.Header {
	return b.hc.GetHeader(hash, number)
}
