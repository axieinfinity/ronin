package main

import (
	"context"
	"encoding/json"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/internal/ethapi"
	"github.com/ethereum/go-ethereum/log"
	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/sasl/plain"
	"github.com/segmentio/kafka-go/sasl/scram"
	"gopkg.in/urfave/cli.v1"
	"time"
)

const (
	chainHeadEventTopic   = "subscriber.blockEventTopic"
	transactionEventTopic = "subscriber.txEventTopic"
	kafkaPartition        = "subscriber.kafka.partition"
	kafkaUrl              = "subscriber.kafka.url"
	maxRetry              = "subscriber.maxRetry"
	numberOfWorker        = "subscriber.workers"
	backoff               = "subscriber.backoff"
	publisherType         = "subscriber.publisher"
	fromHeight            = "subscriber.fromHeight"
	kafkaUsername         = "subscriber.kafka.username"
	kafkaPassword         = "subscriber.kafka.password"
	kafkaAuthentication   = "subscriber.kafka.authentication"
	queueSize             = "subscriber.queueSize"
)

var (
	SubscriberFlag = cli.BoolFlag{
		Name:  "subscriber",
		Usage: "subscribes to blockchain event",
	}
	ChainHeadEventFlag = cli.StringFlag{
		Name:  chainHeadEventTopic,
		Usage: "topic name that new block will be published to",
	}
	TransactionEventFlag = cli.StringFlag{
		Name:  transactionEventTopic,
		Usage: "topic name that new transactions will be published to",
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
	}
	NumberOfWorkerFlag = cli.IntFlag{
		Name:  numberOfWorker,
		Usage: "number of concurrent workers",
	}
	BackOffFlag = cli.IntFlag{
		Name:  backoff,
		Usage: "the weighted number which is used for exponential backoff that handles failed job",
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
		Usage: "authentication type. eg: plain, sasl",
	}
	QueueSizeFlag = cli.IntFlag{
		Name:  queueSize,
		Usage: "specify size of workers queue and jobs queue",
	}
	SubscriberFlags = []cli.Flag{
		SubscriberFlag,
		ChainHeadEventFlag,
		TransactionEventFlag,
		KafkaPartitionFlag,
		KafkaUrlFlag,
		KafkaAuthenticationFlag,
		kafkaUsernameFlag,
		kafkaPasswordFlag,
		MaxRetryFlag,
		NumberOfWorkerFlag,
		BackOffFlag,
		PublisherFlag,
		FromHeightFlag,
	}
)

// RPCTransaction represents a transaction that will serialize to the RPC representation of a transaction
type NewTransaction struct {
	BlockHash        common.Hash     `json:"blockHash"`
	BlockNumber      uint64          `json:"blockNumber"`
	From             common.Address  `json:"from"`
	Gas              hexutil.Uint64  `json:"gas"`
	GasPrice         *hexutil.Big    `json:"gasPrice"`
	Hash             common.Hash     `json:"hash"`
	Input            hexutil.Bytes   `json:"input"`
	Nonce            hexutil.Uint64  `json:"nonce"`
	To               *common.Address `json:"to"`
	TransactionIndex hexutil.Uint    `json:"transactionIndex"`
	Value            *hexutil.Big    `json:"value"`
	V                *hexutil.Big    `json:"v"`
	R                *hexutil.Big    `json:"r"`
	S                *hexutil.Big    `json:"s"`
}

type NewBlock struct {
	Number           uint64         `json:"number"`
	Hash             common.Hash    `json:"hash"`
	ParentHash       common.Hash    `json:"parentHash"`
	Nonce            uint64         `json:"nonce"`
	MixHash          common.Hash    `json:"mixHash"`
	LogsBloom        types.Bloom    `json:"logsBloom"`
	StateRoot        common.Hash    `json:"stateRoot"`
	Miner            common.Address `json:"coinbase"`
	Difficulty       *hexutil.Big   `json:"difficulty"`
	ExtraData        hexutil.Bytes  `json:"extraData"`
	Size             hexutil.Uint64 `json:"size"`
	GasLimit         hexutil.Uint64 `json:"gasLimit"`
	GasUsed          hexutil.Uint64 `json:"gasUsed"`
	TimeStamp        hexutil.Uint64 `json:"timestamp"`
	TransactionsRoot common.Hash    `json:"transactionsRoot"`
	ReceiptsRoot     common.Hash    `json:"receiptsRoot"`
}

// newRPCTransaction returns a transaction that will serialize to the RPC
// representation, with the given location metadata set (if available).
func newTransaction(tx *types.Transaction, blockHash common.Hash, blockNumber uint64, index uint64) *NewTransaction {
	var signer types.Signer = types.FrontierSigner{}
	if tx.Protected() {
		signer = types.NewEIP155Signer(tx.ChainId())
	}
	from, _ := types.Sender(signer, tx)
	v, r, s := tx.RawSignatureValues()

	result := &NewTransaction{
		From:     from,
		Gas:      hexutil.Uint64(tx.Gas()),
		GasPrice: (*hexutil.Big)(tx.GasPrice()),
		Hash:     tx.Hash(),
		Input:    hexutil.Bytes(tx.Data()),
		Nonce:    hexutil.Uint64(tx.Nonce()),
		To:       tx.To(),
		Value:    (*hexutil.Big)(tx.Value()),
		V:        (*hexutil.Big)(v),
		R:        (*hexutil.Big)(r),
		S:        (*hexutil.Big)(s),
	}
	if blockHash != (common.Hash{}) {
		result.BlockHash = blockHash
		result.BlockNumber = blockNumber
		result.TransactionIndex = hexutil.Uint(index)
	}
	return result
}

func newBlock(b *types.Block) *NewBlock {
	head := b.Header()
	return &NewBlock{
		Number:           head.Number.Uint64(),
		Hash:             b.Hash(),
		ParentHash:       head.ParentHash,
		Nonce:            head.Nonce.Uint64(),
		MixHash:          head.MixDigest,
		LogsBloom:        head.Bloom,
		StateRoot:        head.Root,
		Miner:            head.Coinbase,
		Difficulty:       (*hexutil.Big)(head.Difficulty),
		ExtraData:        head.Extra,
		Size:             hexutil.Uint64(b.Size()),
		GasLimit:         hexutil.Uint64(head.GasLimit),
		GasUsed:          hexutil.Uint64(head.GasUsed),
		TimeStamp:        hexutil.Uint64(head.Time),
		TransactionsRoot: head.TxHash,
		ReceiptsRoot:     head.ReceiptHash,
	}
}

// Publisher is used in subscriber to publish message to target message broker
type Publisher interface {
	publish(topic string, data []byte) error
	setConn(topic string) error
	close()
}

type Job struct {
	Topic      string
	Message    []byte
	RetryCount int
	NextTry    int
	MaxTry     int
	BackOff    int
}

type Worker struct {
	publishFn func(topic string, message []byte) error

	// queue is passed from subscriber is used to add workerChan to queue
	queue chan chan Job

	// mainChain is subscriber's jobChan which is used to push job back to subscriber
	mainChan chan Job

	// workerChan is used to receive and publishing job
	workerChan chan Job

	closeChan chan struct{}
}

// Subscriber is used to subscribe blockchain event and organized workers
// to publish message into targeted message broker.
type Subscriber struct {
	eventPublisher      Publisher
	chainHeadEvent      chan core.ChainHeadEvent
	chainHeadEventTopic string
	transactionsTopic   string
	closeCh             chan struct{}
	MaxRetry            int
	BackOff             int

	// start publishing from specific block's height
	// this field is necessary when we don't want to handle data which is already existed.
	FromHeight uint64

	Workers []*Worker

	// Queue holds a list of worker
	Queue chan chan Job

	// JobChan receives new job
	JobChan chan Job
}

type DefaultEventPublisher struct {
	Partition          int
	URL                string
	Username           string
	Password           string
	AuthenticationType string
	Connections        map[string]*kafka.Conn
}

func NewSubscriber(eth ethapi.Backend, ctx *cli.Context) *Subscriber {
	workers := 1
	subs := &Subscriber{
		chainHeadEvent: make(chan core.ChainHeadEvent, 1),
		closeCh:        make(chan struct{}, 1),
		MaxRetry:       100,
		BackOff:        5,
		Workers:        make([]*Worker, 0),
		Queue:          make(chan chan Job, 1000),
		JobChan:        make(chan Job, 1000),
	}

	queueSize := ctx.GlobalInt(QueueSizeFlag.Name)
	if queueSize > 0 {
		subs.JobChan = make(chan Job, queueSize)
		subs.Queue = make(chan chan Job, queueSize)
	}

	// set event publisher
	handlerType := ctx.GlobalString(PublisherFlag.Name)
	switch handlerType {
	default:
		subs.eventPublisher = NewDefaultEventPublisher(ctx)
	}

	if ctx.GlobalIsSet(ChainHeadEventFlag.Name) {
		subs.chainHeadEventTopic = ctx.GlobalString(ChainHeadEventFlag.Name)
		eth.SubscribeChainHeadEvent(subs.chainHeadEvent)
		if err := subs.eventPublisher.setConn(subs.chainHeadEventTopic); err != nil {
			panic(err)
		}
	}
	if ctx.GlobalIsSet(TransactionEventFlag.Name) {
		subs.transactionsTopic = ctx.GlobalString(TransactionEventFlag.Name)
		if err := subs.eventPublisher.setConn(subs.transactionsTopic); err != nil {
			panic(err)
		}
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

	// init workers
	for i := 0; i < workers; i++ {
		subs.Workers = append(subs.Workers, NewWorker(subs.JobChan, subs.Queue, subs.eventPublisher.publish))
	}
	return subs
}

// HandleNewBlock handles mined block's data.
// Block/Transaction data will be submitted to job channel and be handled by subscriber's workers
func (s *Subscriber) HandleNewBlock(block *types.Block) {
	// if block's height is less than subscribed height then do nothing.
	if block == nil || block.NumberU64() < s.FromHeight {
		return
	}
	if s.chainHeadEventTopic != "" {
		blockData, err := json.Marshal(newBlock(block))
		if err != nil {
			log.Error("[HandleNewBlock]Marshal Block Data", "error", err)
			return
		}
		s.JobChan <- NewJob(s.chainHeadEventTopic, blockData, s.MaxRetry, s.BackOff)
	}
	if s.transactionsTopic != "" && block != nil {
		for i, tx := range block.Transactions() {
			txData, err := json.Marshal(newRPCTransaction(tx, block.Hash(), block.NumberU64(), uint64(i)))
			if err != nil {
				log.Error("[HandleNewBlock]Marshal Transaction Data", "error", err, "blockHeight", block.NumberU64(), "index", i)
				continue
			}
			s.JobChan <- NewJob(s.transactionsTopic, txData, s.MaxRetry, s.BackOff)
		}
	}
}

func (s *Subscriber) Start() {
	for _, worker := range s.Workers {
		go worker.start()
	}
	for {
		select {
		case evt := <-s.chainHeadEvent:
			s.HandleNewBlock(evt.Block)
		case job := <-s.JobChan:
			// get a worker from queue
			jobChan := <-s.Queue
			jobChan <- job
		case <-s.closeCh:
			close(s.closeCh)
			return
		}
	}
}

func (s *Subscriber) Close() {
	// close all workers
	for _, worker := range s.Workers {
		worker.close()
	}
	// close event publisher
	s.eventPublisher.close()
	// close listener
	s.closeCh <- struct{}{}
}

func NewJob(topic string, message []byte, maxTry, backOff int) Job {
	return Job{
		Topic:   topic,
		Message: message,
		MaxTry:  maxTry,
		BackOff: backOff,
	}
}

func NewWorker(mainChan chan Job, queue chan chan Job, publishFn func(string, []byte) error) *Worker {
	return &Worker{
		workerChan: make(chan Job),
		mainChan:   mainChan,
		queue:      queue,
		publishFn:  publishFn,
	}
}

func (w *Worker) start() {
	for {
		// push worker chan into queue
		w.queue <- w.workerChan
		select {
		case job := <-w.workerChan:
			if job.NextTry == 0 || job.NextTry <= time.Now().Second() {
				if err := w.publishFn(job.Topic, job.Message); err != nil {
					// check if this job reaches maxTry or not
					// if it is not send it back to mainChan
					if job.RetryCount+1 >= job.MaxTry {
						return
					}
					job.RetryCount += 1
					job.NextTry = time.Now().Second() + (job.RetryCount * job.BackOff)
				}
			}
			// push the job back to mainChan
			w.mainChan <- job
		case <-w.closeChan:
			close(w.workerChan)
			close(w.closeChan)
			return
		}
	}
}

func (w *Worker) close() {
	w.closeChan <- struct{}{}
}

func NewDefaultEventPublisher(ctx *cli.Context) *DefaultEventPublisher {
	return &DefaultEventPublisher{
		Partition:   ctx.GlobalInt(KafkaPartitionFlag.Name),
		URL:         ctx.GlobalString(KafkaUrlFlag.Name),
		Connections: make(map[string]*kafka.Conn),
	}
}

func (s *DefaultEventPublisher) publish(topic string, data []byte) error {
	conn := s.Connections[topic]
	if conn != nil {
		if _, err := conn.WriteMessages(kafka.Message{
			Value: data,
		}); err != nil {
			log.Error("[DefaultEventPublisher][publish]Write Message", "err", err, "data", string(data))
			return err
		}
	}
	return nil
}

// setConn inits connection for each topic
func (s *DefaultEventPublisher) setConn(topic string) error {
	// if connection exists then do nothing
	if s.Connections[topic] == nil {
		conn, err := s.conn(topic)
		if err != nil {
			log.Error("[DefaultEventPublisher][publish]Get kafka connection", "err", err, "topic", topic)
			return err
		}
		s.Connections[topic] = conn
	}
	return nil
}

func (s *DefaultEventPublisher) close() {
	// close all connections
	for _, conn := range s.Connections {
		conn.Close()
	}
}

// conn establishes a connection to kafka based on a specific topic
func (s *DefaultEventPublisher) conn(topic string) (*kafka.Conn, error) {
	dialer := &kafka.Dialer{
		Timeout:   10 * time.Second,
		DualStack: true,
	}
	if s.Username != "" {
		var err error
		switch s.AuthenticationType {
		case scram.SHA512.Name():
			dialer.SASLMechanism, err = scram.Mechanism(scram.SHA512, s.Username, s.Password)
			if err != nil {
				return nil, err
			}
		default:
			dialer.SASLMechanism = plain.Mechanism{Username: s.Username, Password: s.Password}
		}
	}
	return dialer.DialLeader(context.Background(), "tcp", s.URL, topic, s.Partition)
}
