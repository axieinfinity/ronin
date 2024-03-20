package v2

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus"
	consortiumCommon "github.com/ethereum/go-ethereum/consensus/consortium/common"
	"github.com/ethereum/go-ethereum/consensus/consortium/v2/finality"
)

type consortiumV2Api struct {
	chain      consensus.ChainHeaderReader
	consortium *Consortium
}

// GetValidatorAtHash returns the authorized validators that can seal block hash with
// their BLS public key if available
func (api *consortiumV2Api) GetValidatorAtHash(hash common.Hash) ([]finality.ValidatorWithBlsPub, error) {
	header := api.chain.GetHeaderByHash(hash)
	if header == nil {
		return nil, consortiumCommon.ErrUnknownBlock
	}

	snap, err := api.consortium.snapshot(api.chain, header.Number.Uint64()-1, header.ParentHash, nil)
	if err != nil {
		return nil, err
	}

	if snap.ValidatorsWithBlsPub != nil {
		return snap.ValidatorsWithBlsPub, nil
	}

	var validators []finality.ValidatorWithBlsPub
	for validator := range snap.Validators {
		validators = append(validators, finality.ValidatorWithBlsPub{
			Address: validator,
		})
	}

	return validators, nil
}

type finalityVote struct {
	Signature      string   `json:"signature"`
	VoterPublicKey []string `json:"voterPublicKey"`
	VoterAddress   []string `json:"voterAddress"`
}

// GetFinalityVoteAtHash returns the finality vote at block hash
func (api *consortiumV2Api) GetFinalityVoteAtHash(hash common.Hash) (*finalityVote, error) {
	header := api.chain.GetHeaderByHash(hash)
	if header == nil {
		return nil, consortiumCommon.ErrUnknownBlock
	}

	extraData, err := finality.DecodeExtraV2(header.Extra, api.consortium.chainConfig, header.Number)
	if err != nil {
		return nil, err
	}

	if extraData.HasFinalityVote == 0 {
		return nil, nil
	}

	var vote finalityVote
	vote.Signature = common.Bytes2Hex(extraData.AggregatedFinalityVotes.Marshal())

	snap, err := api.consortium.snapshot(api.chain, header.Number.Uint64()-1, header.ParentHash, nil)
	if err != nil {
		return nil, err
	}
	position := extraData.FinalityVotedValidators.Indices()
	for _, pos := range position {
		validator := snap.ValidatorsWithBlsPub[pos]
		vote.VoterAddress = append(vote.VoterAddress, validator.Address.Hex())
		vote.VoterPublicKey = append(vote.VoterPublicKey, common.Bytes2Hex(validator.BlsPublicKey.Marshal()))
	}

	return &vote, nil
}
