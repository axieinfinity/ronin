package vote

import (
	"github.com/ethereum/go-ethereum/consensus"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/eth/downloader"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/event"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/metrics"
	"github.com/ethereum/go-ethereum/params"
)

var votesManagerCounter = metrics.NewRegisteredCounter("votesManager/local", nil)

// Backend wraps all methods required for voting.
type Backend interface {
	IsMining() bool
	EventMux() *event.TypeMux
}

type Debug struct {
	ValidateRule func(header *types.Header) error
}

// VoteManager will handle the vote produced by self.
type VoteManager struct {
	eth Backend
	db  ethdb.Database

	chain       *core.BlockChain
	chainconfig *params.ChainConfig

	chainHeadCh  chan core.ChainHeadEvent
	chainHeadSub event.Subscription

	pool   *VotePool
	signer *VoteSigner

	engine consensus.FastFinalityPoSA

	// debug is a set of function which are used to debug any function called in VoteManager
	debug *Debug
}

func NewVoteManager(
	eth Backend,
	db ethdb.Database,
	chainconfig *params.ChainConfig,
	chain *core.BlockChain,
	pool *VotePool,
	blsPasswordPath, blsWalletPath string,
	engine consensus.FastFinalityPoSA,
	debug *Debug,
) (*VoteManager, error) {
	voteManager := &VoteManager{
		eth: eth,
		db:  db,

		chain:       chain,
		chainconfig: chainconfig,
		chainHeadCh: make(chan core.ChainHeadEvent, chainHeadChanSize),

		pool:   pool,
		engine: engine,
		debug:  debug,
	}

	// Create voteSigner.
	voteSigner, err := NewVoteSigner(blsPasswordPath, blsWalletPath)
	if err != nil {
		return nil, err
	}
	log.Info("Create voteSigner successfully")
	voteManager.signer = voteSigner

	// Subscribe to chain head event.
	voteManager.chainHeadSub = voteManager.chain.SubscribeChainHeadEvent(voteManager.chainHeadCh)

	go voteManager.loop()

	return voteManager, nil
}

func (voteManager *VoteManager) loop() {
	log.Debug("vote manager routine loop started")
	events := voteManager.eth.EventMux().Subscribe(downloader.StartEvent{}, downloader.DoneEvent{}, downloader.FailedEvent{})
	defer func() {
		log.Debug("vote manager loop defer func occur")
		if !events.Closed() {
			log.Debug("event not closed, unsubscribed by vote manager loop")
			events.Unsubscribe()
		}
	}()

	dlEventCh := events.Chan()

	startVote := true
	for {
		select {
		case ev := <-dlEventCh:
			if ev == nil {
				log.Debug("dlEvent is nil, continue")
				continue
			}
			switch ev.Data.(type) {
			case downloader.StartEvent:
				log.Debug("downloader is in startEvent mode, will not startVote")
				startVote = false
			case downloader.FailedEvent:
				log.Debug("downloader is in FailedEvent mode, set startVote flag as true")
				startVote = true
			case downloader.DoneEvent:
				log.Debug("downloader is in DoneEvent mode, set the startVote flag to true")
				startVote = true
			}
		case cHead := <-voteManager.chainHeadCh:
			if !startVote {
				log.Debug("startVote flag is false, continue")
				continue
			}
			if !voteManager.eth.IsMining() {
				log.Debug("skip voting because mining is disabled, continue")
				continue
			}

			if cHead.Block == nil {
				log.Debug("cHead.Block is nil, continue")
				continue
			}

			curHead := cHead.Block.Header()
			// Check if cur validator is within the validatorSet at curHead
			if !voteManager.engine.IsActiveValidatorAt(voteManager.chain, curHead) {
				log.Debug("cur validator is not within the validatorSet at curHead")
				continue
			}

			// Vote for curBlockHeader block.
			vote := &types.VoteData{
				TargetNumber: curHead.Number.Uint64(),
				TargetHash:   curHead.Hash(),
			}
			voteMessage := &types.VoteEnvelope{
				RawVoteEnvelope: types.RawVoteEnvelope{
					Data: vote,
				},
			}

			// Put Vote into journal and VotesPool if we are active validator and allow to sign it.
			if ok := voteManager.UnderRules(curHead); ok {
				log.Debug("curHead is underRules for voting")
				if err := voteManager.signer.SignVote(voteMessage); err != nil {
					log.Error("Failed to sign vote", "err", err, "votedBlockNumber", voteMessage.Data.TargetNumber, "votedBlockHash", voteMessage.Data.TargetHash, "voteMessageHash", voteMessage.Hash())
					votesSigningErrorCounter.Inc(1)
					continue
				}
				rawdb.WriteHighestFinalityVote(voteManager.db, curHead.Number.Uint64())

				log.Debug("vote manager produced vote", "votedBlockNumber", voteMessage.Data.TargetNumber, "votedBlockHash", voteMessage.Data.TargetHash, "voteMessageHash", voteMessage.Hash())
				voteManager.pool.PutVote(voteMessage)
				votesManagerCounter.Inc(1)
			}
		case <-voteManager.chainHeadSub.Err():
			log.Debug("voteManager subscribed chainHead failed")
			return
		}
	}
}

// UnderRules checks if the produced header under the following rules:
// A validator must not publish two distinct votes for the same height. (Rule 1)
// Validators always vote for their canonical chain’s latest block. (Rule 2)
func (voteManager *VoteManager) UnderRules(header *types.Header) bool {
	// call debug method
	if voteManager.debug != nil && voteManager.debug.ValidateRule != nil {
		if err := voteManager.debug.ValidateRule(header); err != nil {
			log.Debug("error while call debug.ValidateRule", "err", err)
			return false
		}
	}

	highestFinalityVote := rawdb.ReadHighestFinalityVote(voteManager.db)

	// Rule: A validator only votes for the block with a bigger block height than its previous vote
	// This rule implies rule: A validator must not publish two distinct votes for the same height
	targetNumber := header.Number.Uint64()
	if highestFinalityVote != nil && targetNumber <= *highestFinalityVote {
		log.Debug("err: A validator must not publish two distinct votes for the same height.")
		return false
	}

	// Rule: Validators always vote for their canonical chain’s latest block.
	// Since the header subscribed to is the canonical chain, so this rule is satisfied by default.
	log.Debug("All rules check passed")
	return true
}
