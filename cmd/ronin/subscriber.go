package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"runtime"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/eth"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/ethdb/httpdb"
	"github.com/ethereum/go-ethereum/event"
	"github.com/ethereum/go-ethereum/internal/ethapi"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/sasl/plain"
	"github.com/segmentio/kafka-go/sasl/scram"
	"gopkg.in/urfave/cli.v1"
)

const (
	blockEventTopic                = "subscriber.blockEventTopic"
	reOrgBlockEventTopic           = "subscriber.reOrgBlockEventTopic"
	transactionEventTopic          = "subscriber.txEventTopic"
	reorgTransactionEventTopic     = "subscriber.reorgTxEventTopic"
	logsEventTopic                 = "subscriber.logsEventTopic"
	blockConfirmedEventTopic       = "subscriber.blockConfirmedEventTopic"
	transactionConfirmedEventTopic = "subscriber.transactionConfirmedEventTopic"
	logsConfirmedEventTopic        = "subscriber.logsConfirmedEventTopic"
	internalTransactionEventTopic  = "subscriber.internalTransactionEventTopic"
	transactionResultTopic         = "subscriber.transactionResultTopic"
	kafkaPartition                 = "subscriber.kafka.partition"
	kafkaUrl                       = "subscriber.kafka.url"
	maxRetry                       = "subscriber.maxRetry"
	numberOfWorker                 = "subscriber.workers"
	backoff                        = "subscriber.backoff"
	publisherType                  = "subscriber.publisher"
	fromHeight                     = "subscriber.fromHeight"
	kafkaUsername                  = "subscriber.kafka.username"
	kafkaPassword                  = "subscriber.kafka.password"
	kafkaAuthentication            = "subscriber.kafka.authentication"
	queueSize                      = "subscriber.queueSize"
	safeBlockRange                 = "subscriber.safeBlockRange"
	coolDownDuration               = "subscriber.coolDownDuration"
	rpcUrl                         = "subscriber.rpcUrl"
	archiveUrl                     = "subscriber.archiveUrl"
	resyncLimitation               = "subscriber.resync.limit"
	defaultSafeBlockRange          = 10
	statsDuration                  = 30
	defaultWorkers                 = 2048
	defaultMaxQueueSize            = 4096
	defaultCoolDownDuration        = 1
	defaultReSyncLimitation        = 20
)

const (
	pubJob = iota
	sendConfirmBlock
)

var (
	SubscriberFlag = cli.BoolFlag{
		Name:  "subscriber",
		Usage: "subscribes to blockchain event",
	}
	ChainEventFlag = cli.StringFlag{
		Name:  blockEventTopic,
		Usage: "topic name that new block will be published to",
	}
	ReOrgBlockEventFlag = cli.StringFlag{
		Name:  reOrgBlockEventTopic,
		Usage: "topic name that reorged block will be published to",
	}
	TransactionEventFlag = cli.StringFlag{
		Name:  transactionEventTopic,
		Usage: "topic name that new transactions will be published to",
	}
	ReorgTransactionEventFlag = cli.StringFlag{
		Name:  reorgTransactionEventTopic,
		Usage: "topic name that reorg transactions will be published to",
	}
	LogsEventFlag = cli.StringFlag{
		Name:  logsEventTopic,
		Usage: "topic name that new logs will be published to",
	}
	BlockConfirmedEventFlag = cli.StringFlag{
		Name:  blockConfirmedEventTopic,
		Usage: "topic name that confirmed block will be published to",
	}
	TransactionConfirmedEventFlag = cli.StringFlag{
		Name:  transactionConfirmedEventTopic,
		Usage: "topic name that confirmed transaction will be published to",
	}
	LogsConfirmedEventFlag = cli.StringFlag{
		Name:  logsConfirmedEventTopic,
		Usage: "topic name that confirmed logs will be published to",
	}
	InternalTxEventFlag = cli.StringFlag{
		Name:  internalTransactionEventTopic,
		Usage: "topic name that internal transaction message will be published to",
	}
	KafkaPartitionFlag = cli.IntFlag{
		Name:  kafkaPartition,
		Usage: "partition of kafka topic. Default 0",
		Value: 0,
	}
	KafkaUrlFlag = cli.StringFlag{
		Name:  kafkaUrl,
		Usage: "kafka connection url",
	}
	MaxRetryFlag = cli.IntFlag{
		Name:  maxRetry,
		Usage: "maximum retry time for a failed job",
		Value: 100,
	}
	NumberOfWorkerFlag = cli.IntFlag{
		Name:  numberOfWorker,
		Usage: "number of concurrent workers",
		Value: defaultWorkers,
	}
	BackOffFlag = cli.IntFlag{
		Name:  backoff,
		Usage: "the weighted number which is used for exponential backoff that handles failed job",
		Value: 5,
	}
	PublisherFlag = cli.StringFlag{
		Name:  publisherType,
		Usage: "type of publishing framework: kafka, google pub/sub",
	}
	FromHeightFlag = cli.Uint64Flag{
		Name:  fromHeight,
		Usage: "the height that the program starts publishing events",
	}
	kafkaUsernameFlag = cli.StringFlag{
		Name:  kafkaUsername,
		Usage: "username to access kafka",
	}
	kafkaPasswordFlag = cli.StringFlag{
		Name:  kafkaPassword,
		Usage: "Password to access kafka",
	}
	KafkaAuthenticationFlag = cli.StringFlag{
		Name:  kafkaAuthentication,
		Usage: "authentication type. eg: PLAIN, SCRAM-SHA-256, SCRAM-SHA-512",
	}
	QueueSizeFlag = cli.IntFlag{
		Name:  queueSize,
		Usage: "specify size of workers queue and jobs queue",
		Value: defaultMaxQueueSize,
	}
	SafeBlockRangeFlag = cli.IntFlag{
		Name:  safeBlockRange,
		Usage: "confirm block that behind current block height (is sent to new block topic) `confirmAt` blocks",
		Value: defaultSafeBlockRange,
	}
	CoolDownDurationFlag = cli.IntFlag{
		Name:  coolDownDuration,
		Usage: "coolDownDuration is used to sleep for a while when a channel reaches its size",
		Value: defaultCoolDownDuration,
	}
	RPCUrlFlag = cli.StringFlag{
		Name:  rpcUrl,
		Usage: "rpcUrl is used in httpdb to reprocess block in case local db (leveldb) cannot find root key",
	}
	ArchiveUrlFlag = cli.StringFlag{
		Name:  archiveUrl,
		Usage: "archiveUrl is used in httpdb in case rpc url cannot find data",
	}
	ResyncLimitationFlag = cli.IntFlag{
		Name:  resyncLimitation,
		Usage: "number of blocks allowed to resync at once",
	}
	TransactionResultEventFlag = cli.StringFlag{
		Name:  transactionResultTopic,
		Usage: "topic name that transaction result message will be published to",
	}
)

type NewLog struct {
	Address       common.Address `json:"address" gencodec:"required"`
	Topics        []common.Hash  `json:"topics" gencodec:"required"`
	Data          hexutil.Bytes  `json:"data" gencodec:"required"`
	BlockNumber   uint64         `json:"blockNumber"`
	TxHash        common.Hash    `json:"transactionHash" gencodec:"required"`
	TxIndex       uint           `json:"transactionIndex"`
	BlockHash     common.Hash    `json:"blockHash"`
	Index         uint           `json:"logIndex"`
	Removed       bool           `json:"removed"`
	TimeStamp     uint64         `json:"timestamp"`
	PublishedTime int64          `json:"publishedTime"`
}

// NewTransaction represents a transaction that will be published to message broker when new block has been mined
type NewTransaction struct {
	BlockHash         common.Hash     `json:"blockHash"`
	BlockNumber       uint64          `json:"blockNumber"`
	TimeStamp         uint64          `json:"timestamp"`
	From              common.Address  `json:"from"`
	ContractAddress   common.Address  `json:"contractAddress"`
	Status            uint64          `json:"status"`
	Gas               hexutil.Uint64  `json:"gas"`
	GasPrice          *hexutil.Big    `json:"gasPrice"`
	GasUsed           uint64          `json:"gasUsed"`
	CumulativeGasUsed uint64          `json:"cumulativeGasUsed"`
	Hash              common.Hash     `json:"hash"`
	Input             hexutil.Bytes   `json:"input"`
	Nonce             hexutil.Uint64  `json:"nonce"`
	To                *common.Address `json:"to"`
	TransactionIndex  hexutil.Uint    `json:"transactionIndex"`
	Value             *hexutil.Big    `json:"value"`
	V                 *hexutil.Big    `json:"v"`
	R                 *hexutil.Big    `json:"r"`
	S                 *hexutil.Big    `json:"s"`
	PublishedTime     int64           `json:"publishedTime"`
}

// NewBlock represents a block that will be published to message broker when new block has been mined
type NewBlock struct {
	Number               uint64         `json:"number"`
	Hash                 common.Hash    `json:"hash"`
	ParentHash           common.Hash    `json:"parentHash"`
	NumberOfTransactions int            `json:"numberOfTransactions"`
	Nonce                uint64         `json:"nonce"`
	MixHash              common.Hash    `json:"mixHash"`
	LogsBloom            types.Bloom    `json:"logsBloom"`
	StateRoot            common.Hash    `json:"stateRoot"`
	Miner                common.Address `json:"coinbase"`
	Difficulty           *hexutil.Big   `json:"difficulty"`
	ExtraData            hexutil.Bytes  `json:"extraData"`
	Size                 hexutil.Uint64 `json:"size"`
	GasLimit             hexutil.Uint64 `json:"gasLimit"`
	GasUsed              hexutil.Uint64 `json:"gasUsed"`
	TimeStamp            hexutil.Uint64 `json:"timestamp"`
	TransactionsRoot     common.Hash    `json:"transactionsRoot"`
	ReceiptsRoot         common.Hash    `json:"receiptsRoot"`
	PublishedTime        int64          `json:"publishedTime"`
}

type InternalTransaction struct {
	Opcode          string         `json:"opcode"`
	Order           uint64         `json:"order"`
	TransactionHash common.Hash    `json:"transactionHash"`
	Hash            common.Hash    `json:"hash"`
	Type            string         `json:"type"`
	Value           *hexutil.Big   `json:"value"`
	Input           hexutil.Bytes  `json:"input"`
	From            common.Address `json:"from"`
	To              common.Address `json:"to"`
	Success         bool           `json:"success"`
	Error           string         `json:"reason"`
	Height          uint64         `json:"height"`
	BlockHash       common.Hash    `json:"blockHash"`
	BlockTime       uint64         `json:"blockTime"`
}

type TransactionResult struct {
	TransactionHash common.Hash   `json:"transactionHash"`
	UsedGas         uint64        `json:"usedGas"`
	Err             string        `json:"error"`
	ReturnData      hexutil.Bytes `json:"returnData"`
}

func newTransaction(tx *types.Transaction, blockHash common.Hash, blockNumber, timestamp uint64, index int, receipts types.Receipts) *NewTransaction {
	var signer types.Signer = types.FrontierSigner{}
	if tx.Protected() {
		signer = types.NewEIP155Signer(tx.ChainId())
	}
	from, _ := types.Sender(signer, tx)
	v, r, s := tx.RawSignatureValues()

	result := &NewTransaction{
		From:          from,
		TimeStamp:     timestamp,
		Gas:           hexutil.Uint64(tx.Gas()),
		GasPrice:      (*hexutil.Big)(tx.GasPrice()),
		Hash:          tx.Hash(),
		Input:         hexutil.Bytes(tx.Data()),
		Nonce:         hexutil.Uint64(tx.Nonce()),
		To:            tx.To(),
		Value:         (*hexutil.Big)(tx.Value()),
		V:             (*hexutil.Big)(v),
		R:             (*hexutil.Big)(r),
		S:             (*hexutil.Big)(s),
		PublishedTime: time.Now().UnixNano(),
	}
	if blockHash != (common.Hash{}) {
		result.BlockHash = blockHash
		result.BlockNumber = blockNumber
		result.TransactionIndex = hexutil.Uint(index)
	}
	if receipts != nil && len(receipts) > index {
		receipt := receipts[index]
		result.Status = receipt.Status
		result.GasUsed = receipt.GasUsed
		result.CumulativeGasUsed = receipt.CumulativeGasUsed
		result.ContractAddress = receipt.ContractAddress
	}
	return result
}

func newBlock(b *types.Block) *NewBlock {
	head := b.Header()
	return &NewBlock{
		Number:               head.Number.Uint64(),
		Hash:                 b.Hash(),
		ParentHash:           head.ParentHash,
		Nonce:                head.Nonce.Uint64(),
		MixHash:              head.MixDigest,
		LogsBloom:            head.Bloom,
		StateRoot:            head.Root,
		Miner:                head.Coinbase,
		Difficulty:           (*hexutil.Big)(head.Difficulty),
		ExtraData:            head.Extra,
		Size:                 hexutil.Uint64(b.Size()),
		GasLimit:             hexutil.Uint64(head.GasLimit),
		GasUsed:              hexutil.Uint64(head.GasUsed),
		TimeStamp:            hexutil.Uint64(head.Time),
		TransactionsRoot:     head.TxHash,
		ReceiptsRoot:         head.ReceiptHash,
		NumberOfTransactions: b.Transactions().Len(),
		PublishedTime:        time.Now().UnixNano(),
	}
}

func newLog(log *types.Log, timestamp uint64) *NewLog {
	return &NewLog{
		Address:       log.Address,
		Topics:        log.Topics,
		Data:          log.Data,
		BlockNumber:   log.BlockNumber,
		TxHash:        log.TxHash,
		TxIndex:       log.TxIndex,
		BlockHash:     log.BlockHash,
		Index:         log.Index,
		Removed:       log.Removed,
		TimeStamp:     timestamp,
		PublishedTime: time.Now().UnixNano(),
	}
}

func newInternalTx(tx *types.InternalTransaction) *InternalTransaction {
	return &InternalTransaction{
		Opcode:          tx.Opcode,
		Order:           tx.Order,
		TransactionHash: tx.TransactionHash,
		Hash:            tx.Hash(),
		Type:            tx.Type,
		Value:           (*hexutil.Big)(tx.Value),
		Input:           tx.Input,
		From:            tx.From,
		To:              tx.To,
		Success:         tx.Success,
		Error:           tx.Error,
		Height:          tx.Height,
		BlockHash:       tx.BlockHash,
		BlockTime:       tx.BlockTime,
	}
}

// Publisher is used in subscriber to publish message to target message broker
type Publisher interface {
	publish(Job) error
	newMessage(string, []byte) interface{}
	checkConnection() error
	close()
}

type Job struct {
	ID         int32
	Type       int
	Message    []interface{}
	RetryCount int
	NextTry    int64
	MaxTry     int
	BackOff    int
}

func (job Job) Hash() common.Hash {
	return common.BytesToHash([]byte(fmt.Sprintf("j-%d-%d-%d", job.ID, job.RetryCount, job.NextTry)))
}

type Worker struct {
	ctx context.Context

	id int

	publishFn func(Job) error

	// queue is passed from subscriber is used to add workerChan to queue
	queue chan chan Job

	// mainChain is subscriber's jobChan which is used to push job back to subscriber
	mainChan chan<- Job

	// workerChan is used to receive and publishing job
	workerChan chan Job

	closeChan chan struct{}
}

// Subscriber is used to subscribe blockchain event and organized workers
// to publish message into targeted message broker.
type Subscriber struct {
	once      sync.Once
	ctx       context.Context
	cancelCtx context.CancelFunc
	backend   ethapi.Backend
	ethereum  *eth.Ethereum
	db        ethdb.Database

	eventPublisher   Publisher
	chainEvent       chan core.ChainEvent
	resyncEvent      chan core.ChainEvent
	reorgEvent       chan core.ReorgEvent
	removeLogsEvent  chan core.RemovedLogsEvent
	rebirthLogsEvent chan []*types.Log
	internalTxEvent  chan types.InternalTransaction

	// topics params
	chainEventTopic        string
	chainSideTopic         string
	transactionsTopic      string
	reorgTransactionsTopic string
	logsTopic              string
	internalTxTopic        string
	transactionResultTopic string

	confirmedBlockTopic       string
	confirmedTransactionTopic string
	confirmedLogsTopic        string

	// message backoff
	MaxRetry int
	BackOff  int

	// coolDownDuration is used to sleep for a while when a channel reaches its size
	coolDownDuration int

	// start publishing from specific block's height
	// this field is necessary when we don't want to handle data which is already existed.
	FromHeight uint64

	Workers []*Worker

	// Queue holds a list of worker
	Queue chan chan Job

	// JobChan receives new job
	JobChan chan Job

	MaxQueueSize int

	// safeBlockRange is used to send confirmed block
	// confirmed block is behind the current block `confirmBlockAt` height
	safeBlockRange int

	jobId         int32
	processedJobs sync.Map

	resyncing           *atomic.Value
	limitation, counter int32
	isReachedLimitation *atomic.Value
}

type DefaultEventPublisher struct {
	Partition          int
	URL                string
	Username           string
	Password           string
	AuthenticationType string
}

func NewSubscriber(ethereum *eth.Ethereum, backend ethapi.Backend, ctx *cli.Context) *Subscriber {
	workers := defaultWorkers
	if ctx.GlobalIsSet(NumberOfWorkerFlag.Name) {
		workers = ctx.GlobalInt(NumberOfWorkerFlag.Name)
	}

	subCtx, cancelCtx := context.WithCancel(context.Background())
	subs := &Subscriber{
		ethereum:         ethereum,
		backend:          backend,
		ctx:              subCtx,
		cancelCtx:        cancelCtx,
		MaxRetry:         100,
		BackOff:          5,
		Workers:          make([]*Worker, 0),
		MaxQueueSize:     defaultMaxQueueSize,
		safeBlockRange:   defaultSafeBlockRange,
		coolDownDuration: defaultCoolDownDuration,
	}
	if ctx.GlobalIsSet(QueueSizeFlag.Name) {
		queueSize := ctx.GlobalInt(QueueSizeFlag.Name)
		if queueSize > 0 {
			subs.MaxQueueSize = queueSize
		}
	}
	subs.JobChan = make(chan Job, subs.MaxQueueSize)
	subs.Queue = make(chan chan Job, subs.MaxQueueSize)
	subs.chainEvent = make(chan core.ChainEvent, subs.MaxQueueSize)
	subs.reorgEvent = make(chan core.ReorgEvent, subs.MaxQueueSize)
	subs.removeLogsEvent = make(chan core.RemovedLogsEvent, subs.MaxQueueSize)
	subs.rebirthLogsEvent = make(chan []*types.Log, subs.MaxQueueSize)
	subs.resyncEvent = make(chan core.ChainEvent, subs.MaxQueueSize)
	subs.internalTxEvent = make(chan types.InternalTransaction, subs.MaxQueueSize)

	// set event publisher
	handlerType := ctx.GlobalString(PublisherFlag.Name)
	switch handlerType {
	default:
		subs.eventPublisher = NewDefaultEventPublisher(ctx)
	}
	// check connection
	if err := subs.eventPublisher.checkConnection(); err != nil {
		panic(err)
	}
	if ctx.GlobalIsSet(ChainEventFlag.Name) {
		subs.chainEventTopic = ctx.GlobalString(ChainEventFlag.Name)
		backend.SubscribeChainEvent(subs.chainEvent)
	}
	if ctx.GlobalIsSet(ReOrgBlockEventFlag.Name) {
		subs.chainSideTopic = ctx.GlobalString(ReOrgBlockEventFlag.Name)
		backend.SubscribeReorgEvent(subs.reorgEvent)
	}
	if ctx.GlobalIsSet(TransactionEventFlag.Name) {
		subs.transactionsTopic = ctx.GlobalString(TransactionEventFlag.Name)
	}
	if ctx.GlobalIsSet(ReorgTransactionEventFlag.Name) {
		subs.reorgTransactionsTopic = ctx.GlobalString(ReorgTransactionEventFlag.Name)
	}
	if ctx.GlobalIsSet(BlockConfirmedEventFlag.Name) {
		subs.confirmedBlockTopic = ctx.GlobalString(BlockConfirmedEventFlag.Name)
	}
	if ctx.GlobalIsSet(TransactionConfirmedEventFlag.Name) {
		subs.confirmedTransactionTopic = ctx.GlobalString(TransactionConfirmedEventFlag.Name)
	}
	if ctx.GlobalIsSet(LogsConfirmedEventFlag.Name) {
		subs.confirmedLogsTopic = ctx.GlobalString(LogsConfirmedEventFlag.Name)
	}
	if ctx.GlobalIsSet(LogsEventFlag.Name) {
		subs.logsTopic = ctx.GlobalString(LogsEventFlag.Name)
		backend.SubscribeRemovedLogsEvent(subs.removeLogsEvent)
		backend.SubscribeLogsEvent(subs.rebirthLogsEvent)
	}
	if ctx.GlobalIsSet(TransactionResultEventFlag.Name) {
		subs.transactionResultTopic = ctx.GlobalString(TransactionResultEventFlag.Name)
	}
	if ctx.GlobalIsSet(InternalTxEventFlag.Name) {
		subs.internalTxTopic = ctx.GlobalString(InternalTxEventFlag.Name)
		backend.SubscribeInternalTransactionEvent(subs.internalTxEvent)
	}
	if ctx.GlobalIsSet(MaxRetryFlag.Name) {
		subs.MaxRetry = ctx.GlobalInt(MaxRetryFlag.Name)
	}
	if ctx.GlobalIsSet(BackOffFlag.Name) {
		subs.BackOff = ctx.GlobalInt(BackOffFlag.Name)
	}
	if ctx.GlobalIsSet(NumberOfWorkerFlag.Name) {
		workers = ctx.GlobalInt(NumberOfWorkerFlag.Name)
	}
	if ctx.GlobalIsSet(FromHeightFlag.Name) {
		subs.FromHeight = ctx.GlobalUint64(FromHeightFlag.Name)
	}
	if ctx.GlobalIsSet(SafeBlockRangeFlag.Name) {
		subs.safeBlockRange = ctx.GlobalInt(SafeBlockRangeFlag.Name)
	}
	if ctx.GlobalIsSet(CoolDownDurationFlag.Name) {
		subs.coolDownDuration = ctx.GlobalInt(CoolDownDurationFlag.Name)
	}
	if ctx.GlobalIsSet(RPCUrlFlag.Name) {
		subs.db = httpdb.NewDBWithLRU(ctx.GlobalString(RPCUrlFlag.Name), ctx.GlobalString(ArchiveUrlFlag.Name), 0)
	}
	if ctx.GlobalIsSet(ResyncLimitationFlag.Name) {
		subs.limitation = int32(ctx.GlobalInt(ResyncLimitationFlag.Name))
	}
	// init workers
	for i := 0; i < workers; i++ {
		subs.Workers = append(subs.Workers, NewWorker(subs.ctx, i, subs.JobChan, subs.Queue, subs.eventPublisher.publish, subs.MaxQueueSize))
	}
	return subs
}

func (s *Subscriber) CoolDown() {
	if s.coolDownDuration > 0 {
		<-time.NewTicker(time.Duration(s.coolDownDuration) * time.Second).C
	}
}

func (s *Subscriber) SendJob(jobType int, messages ...interface{}) {
	if len(messages) == 0 {
		return
	}
	for len(s.JobChan) >= s.MaxQueueSize {
		log.Debug("JobChan has reached its limit, Sleeping...")
		s.CoolDown()
	}
	s.JobChan <- NewJob(atomic.AddInt32(&s.jobId, 1), jobType, messages, s.MaxRetry, s.BackOff)
}

func (s *Subscriber) HandleNewBlockWithValidation(evt core.ChainEvent) {
	if evt.Block == nil || evt.Block.NumberU64() < s.FromHeight {
		return
	}
	s.HandleNewBlock(evt)
}

// HandleNewBlock handles mined block's data.
// Block/Transaction/Transaction's logs data will be submitted to job channel and be handled by subscriber's workers
func (s *Subscriber) HandleNewBlock(evt core.ChainEvent) {
	block, logs, receipts := evt.Block, evt.Logs, evt.Receipts
	txs := block.Transactions()

	// init messages slice
	messages := make([]interface{}, 0)

	if s.chainEventTopic != "" {
		blockData, err := json.Marshal(newBlock(block))
		if err != nil {
			log.Error("[HandleNewBlock]Marshal Block Data", "error", err, "blockHeight", block.NumberU64())
			return
		}
		messages = append(messages, s.eventPublisher.newMessage(s.chainEventTopic, blockData))
	}
	// call send confirmed block with block behind with current block `confirmBlockAt` blocks
	if s.safeBlockRange > 0 {
		confirmedBlockHeight := block.NumberU64() - uint64(s.safeBlockRange)
		s.SendJob(sendConfirmBlock, func(jobId int32) error {
			log.Info("Processing confirm block job", "jobId", jobId, "height", confirmedBlockHeight)
			return s.SendConfirmedBlock(confirmedBlockHeight)
		})
	}

	// handle sending new transactions
	messages = append(messages, s.HandleNewTransactions(s.transactionsTopic, s.logsTopic, block.Hash(), block.NumberU64(), block.Time(), txs, receipts)...)

	log.Info("[HandleNewBlock] sending new block messages to jobChan", "messages", len(messages), "height", block.NumberU64(), "txs", len(txs), "logs", len(logs))
	s.SendJob(pubJob, messages...)
}

func (s *Subscriber) SendConfirmedBlock(height uint64) error {
	if height < 0 {
		return nil
	}
	messages := make([]interface{}, 0)
	if s.confirmedBlockTopic != "" {
		// get block by number
		block, err := s.backend.BlockByNumber(context.Background(), rpc.BlockNumber(height))
		if err != nil {
			log.Error("[Subscriber][HandleConfirmedBlock] BlockByNumber", "err", err, "height", height)
			return err
		}
		if block == nil {
			log.Debug("[Subscriber][HandleConfirmedBlock] Could not find block", "height", height)
			return errors.New(fmt.Sprintf("block %d not found", height))
		}
		// get receipts by hash
		receipts, err := s.backend.GetReceipts(context.Background(), block.Hash())
		if err != nil {
			log.Error("[Subscriber][HandleConfirmedBlock] GetReceipts", "err", err, "hash", block.Hash().Hex())
			return err
		}
		// marshal block
		blockData, err := json.Marshal(newBlock(block))
		if err != nil {
			log.Error("[Subscriber][HandleConfirmedBlock] Marshal Block Data", "error", err, "height", height)
			return err
		}
		messages = append(messages, s.eventPublisher.newMessage(s.confirmedBlockTopic, blockData))

		if block.Transactions().Len() != len(receipts) {
			log.Error("[Subscriber][HandleConfirmedBlock] mismatched txs len, receipts len or execution results len",
				"height", height, "txs len", block.Transactions().Len(), "receipts len", len(receipts))
			return errors.New("transactions and receipts are mismatched")
		}
		messages = append(messages, s.HandleNewTransactions(s.confirmedTransactionTopic, s.confirmedLogsTopic,
			block.Hash(), height, block.Time(), block.Transactions(), receipts)...)

		// reprocess block
		reprocessesMsg, err := s.reprocessBlock(block)
		if err != nil {
			return err
		}
		messages = append(messages, reprocessesMsg...)
		log.Info("[HandleConfirmedBlock] sending confirmed block messages to jobChan", "messages", len(messages), "height", height)
		s.SendJob(pubJob, messages...)
	}
	if s.limitation > 0 && s.resyncing.Load().(bool) {
		counter := atomic.AddInt32(&s.counter, -1)
		if counter < s.limitation {
			s.isReachedLimitation.Store(false)
		}
	}
	return nil
}

// HandleReorgBlock handles reOrg block event and push relevant block and transactions to message brokers using eventPublisher
func (s *Subscriber) HandleReorgBlock(evt core.ReorgEvent) {
	block := evt.Block
	if block == nil || block.NumberU64() < s.FromHeight {
		return
	}
	txs := block.Transactions()

	// init messages slice
	messages := make([]interface{}, 0)

	if s.chainSideTopic != "" {
		blockData, err := json.Marshal(newBlock(block))
		if err != nil {
			log.Error("[HandleReorgBlock]Marshal Block Data", "error", err, "blockHeight", block.NumberU64())
			return
		}
		messages = append(messages, s.eventPublisher.newMessage(s.chainSideTopic, blockData))
	}
	messages = append(messages, s.HandleNewTransactions(s.reorgTransactionsTopic, s.logsTopic, block.Hash(), block.NumberU64(), block.Time(), txs, nil)...)

	log.Info("[HandleReorgBlock] sending reOrg block messages to jobChan", "messages", len(messages), "height", block.NumberU64(), "txs", len(txs))
	s.SendJob(pubJob, messages...)
}

// HandleNewTransactions converts transaction to readable transaction (JSON) based on transactions list and receipts list and push them message broker.
// if there is any topic within receipts call HandleLogs also to add all Logs to messages
func (s *Subscriber) HandleNewTransactions(topic, logsTopic string, hash common.Hash, number uint64, timestamp uint64, txs types.Transactions, receipts types.Receipts) []interface{} {
	messages := make([]interface{}, 0)
	if topic != "" {
		for i, tx := range txs {
			transaction := newTransaction(tx, hash, number, timestamp, i, receipts)
			txData, err := json.Marshal(transaction)
			if err != nil {
				log.Error("[HandleNewTransactions]Marshal Transaction Data", "error", err, "blockHeight", number, "index", i)
				continue
			}
			messages = append(messages, s.eventPublisher.newMessage(topic, txData))
			if receipts != nil && len(receipts) == len(txs) {
				messages = append(messages, s.HandleLogs(logsTopic, hash, tx.Hash(), number, timestamp, uint(i), receipts[i].Logs)...)
			}
		}
	}
	return messages
}

// HandleLogs converts list of logs to binary and add to published messaged
// When syncing using snap/fast mode, log does not contain txHash, blockHash, blockNumber and txIndex
// Therefore we add these variables from params and update each log with these params.
func (s *Subscriber) HandleLogs(topic string, hash, txHash common.Hash, number uint64, timestamp uint64, txIndex uint, logs []*types.Log) []interface{} {
	messages := make([]interface{}, 0)
	if topic != "" {
		for _, l := range logs {
			if l.BlockNumber < s.FromHeight {
				return messages
			}
			l.TxHash = txHash
			l.BlockHash = hash
			l.BlockNumber = number
			l.TxIndex = txIndex
			logData, err := json.Marshal(newLog(l, timestamp))
			if err != nil {
				log.Error("[HandleLogs]Marshal log data", "err", err, "blockHeight", l.BlockNumber, "index", l.TxIndex)
				continue
			}
			messages = append(messages, s.eventPublisher.newMessage(topic, logData))
		}
	}
	return messages
}

// HandleRemoveRebirthLogs handles removedLogsEvent from blockchain.
// these logs are called when reorg occurs, and they need to be removed.
func (s *Subscriber) HandleRemoveRebirthLogs(logs []*types.Log) {
	messages := make([]interface{}, 0)
	blockTimes := make(map[uint64]uint64)
	if s.logsTopic != "" {
		for _, l := range logs {
			if l.BlockNumber < s.FromHeight {
				return
			}
			// block time at current number is not find then find it in database
			if blockTimes[l.BlockNumber] == 0 {
				header, err := s.backend.HeaderByHash(context.Background(), l.BlockHash)
				if err != nil {
					log.Error("[HandleRemoveRebirthLogs] Get Header by block height", "number", l.BlockNumber, "err", err)
					continue
				}
				blockTimes[l.BlockNumber] = header.Time
			}
			logData, err := json.Marshal(newLog(l, blockTimes[l.BlockNumber]))
			if err != nil {
				log.Error("[HandleRemoveRebirthLogs] Marshal log data", "err", err, "blockHeight", l.BlockNumber, "index", l.TxIndex)
				continue
			}
			messages = append(messages, s.eventPublisher.newMessage(s.logsTopic, logData))
		}
	}
	log.Info("[HandleRemoveRebirthLogs] sending remove/rebirth logs messages to jobChan", "messages", len(messages))
	s.SendJob(pubJob, messages...)
}

// Start starts a subscriber which do the following:
// - Starts all workers
// - Resyncs if fromHeight is less than current backend's height
// - Run all event handlers: ChainEvent, ChainSideEvent, NewBlock, RemovedLogsEvent, RebirthLogsEvent, ctx Done, etc.
func (s *Subscriber) Start() chan struct{} {
	done := make(chan struct{}, 1)
	for _, worker := range s.Workers {
		go worker.start()
	}
	// run all events listeners
	go func() {
		statsTicker := time.NewTicker(statsDuration * time.Second)
		for {
			select {
			case evt := <-s.reorgEvent:
				go s.HandleReorgBlock(evt)
			case evt := <-s.chainEvent:
				go s.HandleNewBlockWithValidation(evt)
			case evt := <-s.removeLogsEvent:
				go s.HandleRemoveRebirthLogs(evt.Logs)
			case evt := <-s.rebirthLogsEvent:
				go s.HandleRemoveRebirthLogs(evt)
			case job := <-s.JobChan:
				// get 1 workerCh from queue and push job to this channel
				hash := job.Hash()
				if _, ok := s.processedJobs.Load(hash); ok {
					continue
				}
				s.processedJobs.Store(hash, struct{}{})
				log.Info("jobChan received a job", "jobId", job.ID, "nextTry", job.NextTry)
				workerCh := <-s.Queue
				workerCh <- job
			case <-s.ctx.Done():
				s.once.Do(func() {
					close(s.Queue)
					close(s.JobChan)
					close(s.reorgEvent)
					close(s.chainEvent)
					close(s.rebirthLogsEvent)
					close(s.removeLogsEvent)
					statsTicker.Stop()
				})
				break
			case <-statsTicker.C:
				log.Info("subscriber stats",
					"WorkerQueueSize", len(s.Queue),
					"jobChan", len(s.JobChan),
					"chainEvent", len(s.chainEvent),
					"numberOfGoRoutines", runtime.NumGoroutine(),
				)
			}
		}
	}()
	// get past blocks if fromHeight < currentHeight
	s.resync()
	done <- struct{}{}
	return done
}

// resync is used when we want to resync blocks that exist on stateDb for some reason.
// Such as message broker is down and message cannot be published to services.
func (s *Subscriber) resync() {
	s.resyncing.Store(true)
	// if fromHeight is update to date or greater than currentHeight then do nothing
	currentHeader := s.backend.CurrentHeader().Number.Uint64()
	if s.FromHeight == 0 || s.FromHeight >= currentHeader {
		return
	}
	// loop until
	for s.FromHeight < currentHeader {
		height := s.FromHeight
		if s.limitation > 0 {
			rs := s.isReachedLimitation.Load()
			if rs.(bool) {
				s.CoolDown()
				continue
			}
			if atomic.AddInt32(&s.counter, 1) >= s.limitation {
				s.isReachedLimitation.Store(true)
			}
		}
		s.SendJob(sendConfirmBlock, func(jobId int32) error {
			log.Info("Processing confirm block job", "jobId", jobId, "height", height)
			return s.SendConfirmedBlock(height)
		})
		s.FromHeight++
	}
	s.resyncing.Store(false)
}

func (s *Subscriber) reprocessBlock(block *types.Block) ([]interface{}, error) {
	messages := make([]interface{}, 0)
	if s.internalTxTopic != "" || s.transactionResultTopic != "" {
		// get parent from blockchain first, if it can
		parent, err := s.backend.BlockByNumber(context.Background(), rpc.BlockNumber(block.NumberU64()-1))
		if err != nil {
			return nil, err
		}
		errCh := make(chan error, 1)
		startTime := time.Now()
		log.Info("[Subscriber][reprocessBlock] reprocessing block", "height", block.NumberU64(), "numOfTxs", len(block.Transactions()))
		txResults, internalTxs := reprocessBlock(s.ethereum.BlockChain(), parent.Root(), block.Header(), block.Transactions(), s.ethereum.APIBackend.ChainConfig(), state.NewDatabase(s.db), errCh)
		err = <-errCh
		if err != nil {
			return nil, err
		}
		log.Info("[Subscriber][reprocessBlock] finish reprocessing block", "height", block.NumberU64(), "txResults", len(txResults), "internalTxs", len(internalTxs), "elapsed", time.Now().Unix()-startTime.Unix())
		if s.internalTxTopic != "" {
			for _, tx := range internalTxs {
				internalTx := newInternalTx(tx)
				data, err := json.Marshal(internalTx)
				if err != nil {
					log.Error("[Subscriber][HandleInternalTransactionEvent] Marshal data", "err", err)
					return nil, err
				}
				log.Info("[Subscriber][reprocessBlock] adding message to internalTxTopic", "topic", s.internalTxTopic, "tx", tx.Hash().Hex(), "height", tx.Height, "parent", tx.TransactionHash.Hex())
				messages = append(messages, s.eventPublisher.newMessage(s.internalTxTopic, data))
			}
		}
		if s.transactionResultTopic != "" {
			for _, tx := range txResults {
				result, err := json.Marshal(tx)
				if err != nil {
					log.Error("[Subscriber][HandleTransactionResult] marshal error", "err", err)
					return nil, err
				}
				log.Info("[Subscriber][reprocessBlock] adding message to txResultTopic", "topic", s.transactionResultTopic, "tx", tx.TransactionHash.Hex(), "height", block.NumberU64())
				messages = append(messages, s.eventPublisher.newMessage(s.transactionResultTopic, result))
			}
		}
	}
	return messages, nil
}

func (s *Subscriber) Close() {
	s.cancelCtx()
}

func NewJob(id int32, jobType int, message []interface{}, maxTry, backOff int) Job {
	return Job{
		ID:      id,
		Type:    jobType,
		Message: message,
		MaxTry:  maxTry,
		BackOff: backOff,
	}
}

func NewWorker(ctx context.Context, id int, mainChan chan Job, queue chan chan Job, publishFn func(job Job) error, size int) *Worker {
	return &Worker{
		ctx:        ctx,
		id:         id,
		workerChan: make(chan Job, size),
		mainChan:   mainChan,
		queue:      queue,
		publishFn:  publishFn,
	}
}

func (w *Worker) String() string {
	return fmt.Sprintf("{ id: %d, currentSize: %d }", w.id, len(w.workerChan))
}

func (w *Worker) start() {
	for {
		// push worker chan into queue
		w.queue <- w.workerChan
		select {
		case job := <-w.workerChan:
			log.Info("processing job", "id", job.ID, "nextTry", job.NextTry, "retryCount", job.RetryCount, "type", job.Type)
			if job.NextTry == 0 || job.NextTry <= time.Now().Unix() {
				w.processJob(job)
				continue
			}
			// push the job back to mainChan
			w.mainChan <- job
		case <-w.ctx.Done():
			close(w.workerChan)
			return
		}
	}
}

func (w *Worker) processJob(job Job) {
	var err error
	switch job.Type {
	case pubJob:
		if err = w.publishFn(job); err != nil {
			log.Error("[worker][publishing]", "err", err, "id", job.ID)
			break
		}
	case sendConfirmBlock:
		if len(job.Message) == 0 {
			return
		}
		sendConfirmBlockFunc := job.Message[0].(func(int32) error)
		if err = sendConfirmBlockFunc(job.ID); err != nil {
			log.Error("[worker][reprocessBlock]", "err", err, "id", job.ID)
			break
		}
	}

	if err != nil {
		if job.RetryCount+1 > job.MaxTry {
			log.Info("[Worker][processJob] job reaches its maxTry", "message", job.Message)
			return
		}

		job.RetryCount++
		job.NextTry = time.Now().Unix() + int64(job.RetryCount*job.BackOff)
		// push the job back to mainChan
		w.mainChan <- job
		log.Info("[Worker][processJob] job failed, added back to jobChan", "id", job.ID, "retryCount", job.RetryCount, "nextTry", job.NextTry)
	}

}

func NewDefaultEventPublisher(ctx *cli.Context) *DefaultEventPublisher {
	return &DefaultEventPublisher{
		Partition: ctx.GlobalInt(KafkaPartitionFlag.Name),
		URL:       ctx.GlobalString(KafkaUrlFlag.Name),
	}
}

func (s *DefaultEventPublisher) publish(job Job) error {
	var messages []kafka.Message
	for _, message := range job.Message {
		messages = append(messages, message.(kafka.Message))
	}
	w := &kafka.Writer{
		Addr:         kafka.TCP(s.URL),
		Compression:  kafka.Snappy,
		WriteTimeout: 10 * time.Second,
	}
	defer w.Close()
	return w.WriteMessages(context.Background(), messages...)
}

func (s *DefaultEventPublisher) newMessage(topic string, data []byte) interface{} {
	return kafka.Message{Topic: topic, Value: data}
}

func (s *DefaultEventPublisher) checkConnection() error {
	dialer, err := s.getDialer()
	if err != nil {
		return err
	}
	conn, err := dialer.Dial("tcp", s.URL)
	if err != nil {
		return err
	}
	return conn.Close()
}

func (s *DefaultEventPublisher) getDialer() (*kafka.Dialer, error) {
	dialer := &kafka.Dialer{
		Timeout:   10 * time.Second,
		DualStack: true,
	}
	if s.Username != "" {
		var err error
		switch strings.ToUpper(s.AuthenticationType) {
		case scram.SHA512.Name():
			dialer.SASLMechanism, err = scram.Mechanism(scram.SHA512, s.Username, s.Password)
		case scram.SHA256.Name():
			dialer.SASLMechanism, err = scram.Mechanism(scram.SHA256, s.Username, s.Password)
		default:
			dialer.SASLMechanism = plain.Mechanism{Username: s.Username, Password: s.Password}
		}
		if err != nil {
			return nil, err
		}
	}
	return dialer, nil
}

func (s *DefaultEventPublisher) close() {}

type ReprocessBlockRequest struct {
	BC           core.ChainContext
	ParentRoot   common.Hash
	Header       *types.Header
	Transactions []*types.Transaction
	ChainConfig  *params.ChainConfig
	DB           state.Database
	ErrorChannel chan error
}

type ReprocessBlockResponse struct {
	TransactionResults   []*TransactionResult
	InternalTransactions []*types.InternalTransaction
}

func reprocessBlock2(req *ReprocessBlockRequest) *ReprocessBlockResponse {
	var (
		internalTxFeed event.Feed
		usedGas        = new(uint64)
		gp             = new(core.GasPool).AddGas(req.Header.GasLimit)
		internalTxsCh  = make(chan *types.InternalTransaction, 10000)
		res            ReprocessBlockResponse
	)
	internalTxFeed.Subscribe(internalTxsCh)
	opEvents := []*vm.PublishEvent{
		{
			OpCodes: []vm.OpCode{vm.CALL},
			Event:   core.NewInternalTransferOrSmcCallEvent(&internalTxFeed),
		},
		{
			OpCodes: []vm.OpCode{vm.CREATE, vm.CREATE2},
			Event:   core.NewInternalTransactionContractCreation(&internalTxFeed),
		},
	}
	defer func() {
		if err := recover(); err != nil {
			req.ErrorChannel <- errors.New(fmt.Sprintf("panic occurred: %v", err))
		}
	}()
	log.Debug("[reprocessBlock] start reprocessBlock", "Height", req.Header.Number.Uint64())
	blockContext := core.NewEVMBlockContext(req.Header, req.BC, nil, opEvents...)
	statedb, err := state.New(req.ParentRoot, req.DB, nil)
	if err != nil {
		req.ErrorChannel <- err
		return nil
	}
	vmenv := vm.NewEVM(blockContext, vm.TxContext{}, statedb, req.ChainConfig, vm.Config{})
	res.TransactionResults = make([]*TransactionResult, len(req.Transactions))
	// Iterate over and process the individual transactions
	startTime := time.Now()
	for i, tx := range req.Transactions {
		log.Info("[reprocessBlock] processing tx", "index", i, "tx", tx.Hash().Hex(), "txNonce", tx.Nonce())
		// set current transaction in block context to each transaction
		vmenv.Context.CurrentTransaction = tx
		// reset counter to start counting opcodes in new transaction
		vmenv.Context.Counter = 0
		msg, err := tx.AsMessage(types.MakeSigner(req.ChainConfig, req.Header.Number), req.Header.BaseFee)
		if err != nil {
			errMsg := fmt.Errorf("could not apply tx %d [%v]: %w", i, tx.Hash().Hex(), err)
			log.Error("[reprocessBlock] tx to message", "err", err, "tx", tx.Hash().Hex(), "block", req.Header.Number.Uint64())
			req.ErrorChannel <- errMsg
			return nil
		}
		nonce := statedb.GetNonce(msg.From())
		log.Debug("[reprocessBlock] get sender nonce", "sender", msg.From().Hex(), "nonce", nonce, "index", i, "txNonce", msg.Nonce(), "tx", tx.Hash().Hex())
		statedb.Prepare(tx.Hash(), i)
		_, result, err := core.ApplyTransactionWithVMContext(msg, req.ChainConfig, nil, nil, gp, statedb, req.Header.Number, req.Header.Hash(), tx, usedGas, vmenv)
		if err != nil {
			errMsg := fmt.Errorf("could not apply tx %d [%v] [block: %d]: %w", i, tx.Hash().Hex(), req.Header.Number.Uint64(), err)
			log.Error("[reprocessBlock] ApplyMessage", "err", err, "tx", tx.Hash().Hex(), "block", req.Header.Number.Uint64())
			req.ErrorChannel <- errMsg
			return nil
		}
		res.TransactionResults[i] = &TransactionResult{
			TransactionHash: tx.Hash(),
			UsedGas:         result.UsedGas,
		}
		if result.Err != nil {
			res.TransactionResults[i].Err = result.Err.Error()
		}
		if result.ReturnData != nil {
			res.TransactionResults[i].ReturnData = result.ReturnData
		}
	}
	endTime := time.Now()
	log.Info("[reprocessBlock] finish processing tx", "txs", len(res.TransactionResults), "internalTxs", len(internalTxsCh), "block", req.Header.Number.Uint64(), "elapsed", endTime.Unix()-startTime.Unix())
	close(internalTxsCh)
	if len(internalTxsCh) > 0 {
		log.Debug("[reprocessBlock] getting internalTxs", "txs", len(internalTxsCh))
		for internalTx := range internalTxsCh {
			res.InternalTransactions = append(res.InternalTransactions, internalTx)
		}
	}
	req.ErrorChannel <- nil
	return &res
}

func reprocessBlock(bc core.ChainContext, parentRoot common.Hash, header *types.Header, txs types.Transactions, chainConfig *params.ChainConfig, db state.Database, errCh chan error) (txResults []*TransactionResult, internalTxs []*types.InternalTransaction) {
	var (
		internalTxFeed event.Feed
		usedGas        = new(uint64)
		gp             = new(core.GasPool).AddGas(header.GasLimit)
		internalTxsCh  = make(chan *types.InternalTransaction, 10000)
	)
	internalTxFeed.Subscribe(internalTxsCh)
	opEvents := []*vm.PublishEvent{
		{
			OpCodes: []vm.OpCode{vm.CALL},
			Event:   core.NewInternalTransferOrSmcCallEvent(&internalTxFeed),
		},
		{
			OpCodes: []vm.OpCode{vm.CREATE, vm.CREATE2},
			Event:   core.NewInternalTransactionContractCreation(&internalTxFeed),
		},
	}
	defer func() {
		if err := recover(); err != nil {
			errCh <- errors.New(fmt.Sprintf("panic occurred: %v", err))
		}
	}()
	log.Debug("[reprocessBlock] start reprocessBlock", "Height", header.Number.Uint64())
	blockContext := core.NewEVMBlockContext(header, bc, nil, opEvents...)
	statedb, err := state.New(parentRoot, db, nil)
	if err != nil {
		errCh <- err
		return
	}
	vmenv := vm.NewEVM(blockContext, vm.TxContext{}, statedb, chainConfig, vm.Config{})
	txResults = make([]*TransactionResult, len(txs))
	// Iterate over and process the individual transactions
	startTime := time.Now()
	for i, tx := range txs {
		log.Info("[reprocessBlock] processing tx", "index", i, "tx", tx.Hash().Hex(), "txNonce", tx.Nonce())
		// set current transaction in block context to each transaction
		vmenv.Context.CurrentTransaction = tx
		// reset counter to start counting opcodes in new transaction
		vmenv.Context.Counter = 0
		msg, err := tx.AsMessage(types.MakeSigner(chainConfig, header.Number), header.BaseFee)
		if err != nil {
			errMsg := fmt.Errorf("could not apply tx %d [%v]: %w", i, tx.Hash().Hex(), err)
			log.Error("[reprocessBlock] tx to message", "err", err, "tx", tx.Hash().Hex(), "block", header.Number.Uint64())
			errCh <- errMsg
			return
		}
		nonce := statedb.GetNonce(msg.From())
		log.Debug("[reprocessBlock] get sender nonce", "sender", msg.From().Hex(), "nonce", nonce, "index", i, "txNonce", msg.Nonce(), "tx", tx.Hash().Hex())
		statedb.Prepare(tx.Hash(), i)
		_, result, err := core.ApplyTransactionWithVMContext(msg, chainConfig, nil, nil, gp, statedb, header.Number, header.Hash(), tx, usedGas, vmenv)
		if err != nil {
			errMsg := fmt.Errorf("could not apply tx %d [%v] [block: %d]: %w", i, tx.Hash().Hex(), header.Number.Uint64(), err)
			log.Error("[reprocessBlock] ApplyMessage", "err", err, "tx", tx.Hash().Hex(), "block", header.Number.Uint64())
			errCh <- errMsg
			return
		}
		txResults[i] = &TransactionResult{
			TransactionHash: tx.Hash(),
			UsedGas:         result.UsedGas,
		}
		if result.Err != nil {
			txResults[i].Err = result.Err.Error()
		}
		if result.ReturnData != nil {
			txResults[i].ReturnData = result.ReturnData
		}
	}
	endTime := time.Now()
	log.Info("[reprocessBlock] finish processing tx", "txs", len(txResults), "internalTxs", len(internalTxsCh), "block", header.Number.Uint64(), "elapsed", endTime.Unix()-startTime.Unix())
	close(internalTxsCh)
	if len(internalTxsCh) > 0 {
		log.Debug("[reprocessBlock] getting internalTxs", "txs", len(internalTxsCh))
		for internalTx := range internalTxsCh {
			internalTxs = append(internalTxs, internalTx)
		}
	}
	errCh <- nil
	return
}
