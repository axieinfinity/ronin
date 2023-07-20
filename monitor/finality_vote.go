package monitor

import (
	"errors"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus"
	"github.com/ethereum/go-ethereum/consensus/consortium/v2/finality"
	"github.com/ethereum/go-ethereum/core/types"
	blsCommon "github.com/ethereum/go-ethereum/crypto/bls/common"
	"github.com/ethereum/go-ethereum/log"
	lru "github.com/hashicorp/golang-lru"
)

const finalityVoteCache = 100

type blockInformation struct {
	blockHash      common.Hash
	voterPublicKey []blsCommon.PublicKey
	voterAddress   []common.Address
}

type FinalityVoteMonitor struct {
	chain         consensus.ChainHeaderReader
	engine        consensus.FastFinalityPoSA
	observedVotes *lru.Cache
}

func NewFinalityVoteMonitor(
	chain consensus.ChainHeaderReader,
	engine consensus.FastFinalityPoSA,
) (*FinalityVoteMonitor, error) {
	observedVotes, err := lru.New(finalityVoteCache)
	if err != nil {
		return nil, err
	}

	return &FinalityVoteMonitor{
		engine:        engine,
		observedVotes: observedVotes,
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
	extraData, err := finality.DecodeExtra(block.Extra(), true)
	// This should not happen because the block has been verified
	if err != nil {
		log.Error("Unexpected error when decode extradata", "err", err)
		return err
	}

	if extraData.HasFinalityVote == 1 {
		blockValidator := monitor.engine.GetActiveValidatorAt(
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
		)
	}

	return nil
}

func (monitor *FinalityVoteMonitor) checkSameHeightVote(
	blockNumber uint64,
	blockHash common.Hash,
	voterPublicKey []blsCommon.PublicKey,
	voterAddress []common.Address,
) error {
	rawBlockInfo, ok := monitor.observedVotes.Get(blockNumber)
	if !ok {
		monitor.observedVotes.Add(blockNumber, []blockInformation{
			{
				blockHash:      blockHash,
				voterPublicKey: voterPublicKey,
				voterAddress:   voterAddress,
			},
		})
		return nil
	}

	violated := false
	blockInfo := rawBlockInfo.([]blockInformation)

	for _, block := range blockInfo {
		// 2 blocks are the same, it's not likely to happen
		if block.blockHash == blockHash {
			continue
		}

		for _, cachePublicKey := range block.voterPublicKey {
			for _, blockPublicKey := range voterPublicKey {
				if blockPublicKey.Equals(cachePublicKey) {
					log.Error(
						"Fast finality rule is violated",
						"voter public key", common.Bytes2Hex(blockPublicKey.Marshal()),
						"block number", blockNumber,
						"block 1 hash", block.blockHash,
						"block 1 voter public key", prettyPrintPublicKey(block.voterPublicKey),
						"block 1 voter address", prettyPrintAddress(block.voterAddress),
						"block 2 hash", blockHash,
						"block 2 voter public key", prettyPrintPublicKey(voterPublicKey),
						"block 2 voter address", prettyPrintAddress(voterAddress),
					)
					violated = true
				}
			}
		}
	}

	blockInfo = append(blockInfo, blockInformation{
		blockHash:      blockHash,
		voterPublicKey: voterPublicKey,
		voterAddress:   voterAddress,
	})

	monitor.observedVotes.Add(blockNumber, blockInfo)

	if violated {
		return errors.New("finality rule violated")
	}
	return nil
}
