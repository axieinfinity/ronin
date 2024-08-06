package monitor

import (
	"errors"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus"
	"github.com/ethereum/go-ethereum/consensus/consortium/v2/finality"
	"github.com/ethereum/go-ethereum/core/types"
	blsCommon "github.com/ethereum/go-ethereum/crypto/bls/common"
	"github.com/ethereum/go-ethereum/log"
	lru "github.com/hashicorp/golang-lru/v2"
)

const finalityVoteCache = 100

type blockInformation struct {
	blockHash           common.Hash
	voterPublicKey      []blsCommon.PublicKey
	voterAddress        []common.Address
	aggregatedSignature blsCommon.Signature
}

type FinalityVoteMonitor struct {
	chain         consensus.ChainHeaderReader
	engine        consensus.FastFinalityPoSA
	observedVotes *lru.Cache[uint64, []blockInformation]
	alerter       *slackAlerter
}

func NewFinalityVoteMonitor(
	chain consensus.ChainHeaderReader,
	engine consensus.FastFinalityPoSA,
) (*FinalityVoteMonitor, error) {
	observedVotes, err := lru.New[uint64, []blockInformation](finalityVoteCache)
	if err != nil {
		return nil, err
	}

	return &FinalityVoteMonitor{
		chain:         chain,
		engine:        engine,
		observedVotes: observedVotes,
		alerter:       NewSlackAlert(),
	}, nil
}

func prettyPrintPublicKey(publicKey []blsCommon.PublicKey) string {
	result := "[ "
	for _, key := range publicKey {
		result += common.Bytes2Hex(key.Marshal()) + ", "
	}

	return result + " ]"
}

func prettyPrintAddress(addresses []common.Address) string {
	result := "[ "
	for _, address := range addresses {
		result += address.String() + ", "
	}

	return result + " ]"
}

func (monitor *FinalityVoteMonitor) CheckFinalityVote(block *types.Block) error {
	extraData, err := finality.DecodeExtraV2(block.Extra(), monitor.chain.Config(), block.Number())
	// This should not happen because the block has been verified
	if err != nil {
		log.Error("Unexpected error when decode extradata", "err", err)
		return err
	}

	if extraData.HasFinalityVote == 1 {
		blockValidator := monitor.engine.GetFinalityVoterAt(
			monitor.chain,
			block.NumberU64()-1,
			block.ParentHash(),
		)

		var (
			voterPublicKey []blsCommon.PublicKey
			voterAddress   []common.Address
		)

		position := extraData.FinalityVotedValidators.Indices()
		for _, pos := range position {
			voterPublicKey = append(voterPublicKey, blockValidator[pos].BlsPublicKey)
			voterAddress = append(voterAddress, blockValidator[pos].Address)
		}

		return monitor.checkSameHeightVote(
			block.NumberU64(),
			block.Hash(),
			voterPublicKey,
			voterAddress,
			extraData.AggregatedFinalityVotes,
		)
	}

	return nil
}

func (monitor *FinalityVoteMonitor) checkSameHeightVote(
	blockNumber uint64,
	blockHash common.Hash,
	voterPublicKey []blsCommon.PublicKey,
	voterAddress []common.Address,
	aggregatedSignature blsCommon.Signature,
) error {
	blockInfo, ok := monitor.observedVotes.Get(blockNumber)
	if !ok {
		monitor.observedVotes.Add(blockNumber, []blockInformation{
			{
				blockHash:           blockHash,
				voterPublicKey:      voterPublicKey,
				voterAddress:        voterAddress,
				aggregatedSignature: aggregatedSignature,
			},
		})
		return nil
	}

	violated := false
	for _, block := range blockInfo {
		// 2 blocks are the same, it's not likely to happen
		if block.blockHash == blockHash {
			continue
		}

		for _, cachePublicKey := range block.voterPublicKey {
			for _, blockPublicKey := range voterPublicKey {
				if blockPublicKey.Equals(cachePublicKey) {
					alertHeader := "Fast finality rule is violated"
					alertFormat := "- Voter public key: %s\n" +
						"- Block number: %d\n" +
						"- Block 1 hash: %s\n" +
						"- Block 1 voter public key: %s\n" +
						"- Block 1 voter address: %s\n" +
						"- Block 1 aggregated signature: %s\n" +
						"- Block 2 hash: %s\n" +
						"- Block 2 voter public key: %s\n" +
						"- Block 2 voter address: %s\n" +
						"- Block 2 aggregated signature: %s\n"

					alertBody := fmt.Sprintf(
						alertFormat,
						common.Bytes2Hex(blockPublicKey.Marshal()),
						blockNumber,
						block.blockHash,
						prettyPrintPublicKey(block.voterPublicKey),
						prettyPrintAddress(block.voterAddress),
						common.Bytes2Hex(block.aggregatedSignature.Marshal()),
						blockHash,
						prettyPrintPublicKey(voterPublicKey),
						prettyPrintAddress(voterAddress),
						common.Bytes2Hex(aggregatedSignature.Marshal()),
					)

					if monitor.alerter != nil {
						monitor.alerter.Alert(alertHeader, alertBody)
					}
					log.Error(alertHeader, "message", alertBody)

					violated = true
				}
			}
		}
	}

	blockInfo = append(blockInfo, blockInformation{
		blockHash:           blockHash,
		voterPublicKey:      voterPublicKey,
		voterAddress:        voterAddress,
		aggregatedSignature: aggregatedSignature,
	})

	monitor.observedVotes.Add(blockNumber, blockInfo)

	if violated {
		return errors.New("finality rule violated")
	}
	return nil
}
