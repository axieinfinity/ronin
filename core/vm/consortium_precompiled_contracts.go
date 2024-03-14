package vm

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"math/big"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/bls/blst"
	blsCommon "github.com/ethereum/go-ethereum/crypto/bls/common"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/rlp"
	"golang.org/x/crypto/sha3"
)

const (
	LogContract = iota
	SortValidator
	VerifyHeaders
	PickValidatorSet
	GetDoubleSignSlashingConfig
	ValidateFinalityVoteProof
	ValidateProofOfPossession
	PickValidatorSetBeacon
	NumOfAbis
)

var (
	rawConsortiumLogAbi                = `[{"inputs":[{"internalType":"string","name":"message","type":"string"}],"name":"log","outputs":[],"stateMutability":"nonpayable","type":"function"}]`
	rawConsortiumSortValidatorAbi      = `[{"inputs":[],"stateMutability":"nonpayable","type":"constructor"},{"inputs":[{"internalType":"address[]","name":"validators","type":"address[]"},{"internalType":"uint256[]","name":"weights","type":"uint256[]"}],"name":"sortValidators","outputs":[{"internalType":"address[]","name":"_validators","type":"address[]"}],"stateMutability":"view","type":"function"}]`
	rawConsortiumVerifyHeadersAbi      = `[{"outputs":[],"name":"getHeader","inputs":[{"internalType":"uint256","name":"chainId","type":"uint256"},{"internalType":"bytes32","name":"parentHash","type":"bytes32"},{"internalType":"bytes32","name":"ommersHash","type":"bytes32"},{"internalType":"address","name":"coinbase","type":"address"},{"internalType":"bytes32","name":"stateRoot","type":"bytes32"},{"internalType":"bytes32","name":"transactionsRoot","type":"bytes32"},{"internalType":"bytes32","name":"receiptsRoot","type":"bytes32"},{"internalType":"uint8[256]","name":"logsBloom","type":"uint8[256]"},{"internalType":"uint256","name":"difficulty","type":"uint256"},{"internalType":"uint256","name":"number","type":"uint256"},{"internalType":"uint64","name":"gasLimit","type":"uint64"},{"internalType":"uint64","name":"gasUsed","type":"uint64"},{"internalType":"uint64","name":"timestamp","type":"uint64"},{"internalType":"bytes","name":"extraData","type":"bytes"},{"internalType":"bytes32","name":"mixHash","type":"bytes32"},{"internalType":"uint64","name":"nonce","type":"uint64"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"address","name":"consensusAddr","type":"address"},{"internalType":"bytes","name":"header1","type":"bytes"},{"internalType":"bytes","name":"header2","type":"bytes"}],"name":"validatingDoubleSignProof","outputs":[{"internalType":"bool","name":"","type":"bool"}],"stateMutability":"view","type":"function"}]`
	rawConsortiumPickValidatorSetAbi   = `[{"inputs":[],"stateMutability":"nonpayable","type":"constructor"},{"inputs":[{"internalType":"address[]","name":"_candidates","type":"address[]"},{"internalType":"uint256[]","name":"_weights","type":"uint256[]"},{"internalType":"uint256[]","name":"_trustedWeights","type":"uint256[]"},{"internalType":"uint256","name":"_maxValidatorNumber","type":"uint256"},{"internalType":"uint256","name":"_maxPrioritizedValidatorNumber","type":"uint256"}],"name":"pickValidatorSet","outputs":[{"internalType":"address[]","name":"_validators","type":"address[]"}],"stateMutability":"view","type":"function"}]`
	rawGetDoubleSignSlashingConfigsAbi = `[{"inputs":[],"name":"getDoubleSignSlashingConfigs","outputs":[{"internalType":"uint256","name":"","type":"uint256"},{"internalType":"uint256","name":"","type":"uint256"},{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"}]`
	rawValidateFinalityVoteProofAbi    = `[{"inputs":[{"internalType":"bytes","name":"voterPublicKey","type":"bytes"},{"internalType":"uint256","name":"targetBlockNumber","type":"uint256"},{"internalType":"bytes32[2]","name":"targetBlockHash","type":"bytes32[2]"},{"internalType":"bytes[][2]","name":"listOfPublicKey","type":"bytes[][2]"},{"internalType":"bytes[2]","name":"aggregatedSignature","type":"bytes[2]"}],"name":"validateFinalityVoteProof","outputs":[{"internalType":"bool","name":"","type":"bool"}],"stateMutability":"view","type":"function"}]`
	rawValidateProofOfPossessionAbi    = `[{"inputs":[{"internalType":"bytes","name":"publicKey","type":"bytes"},{"internalType":"bytes","name":"signature","type":"bytes"}],"name":"validateProofOfPossession","outputs":[{"internalType":"bool","name":"","type":"bool"}],"stateMutability":"view","type":"function"}]`
	rawPickValidatorSetBeaconAbi       = `[{"inputs":[{"internalType":"uint256","name":"period","type":"uint256"},{"internalType":"uint256","name":"epoch","type":"uint256"}],"name":"pickValidatorSet","outputs":[{"internalType":"address[]","name":"pickedValidatorIds","type":"address[]"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"uint256","name":"beacon","type":"uint256"},{"internalType":"uint256","name":"period","type":"uint256"},{"internalType":"uint256","name":"numGovernanceValidator","type":"uint256"},{"internalType":"uint256","name":"numStandardValidator","type":"uint256"},{"internalType":"uint256","name":"numRotatingValidator","type":"uint256"},{"internalType":"address[]","name":"ids","type":"address[]"},{"internalType":"uint256[]","name":"stakedAmounts","type":"uint256[]"},{"internalType":"uint256[]","name":"trustedWeights","type":"uint256[]"}],"name":"requestSortValidatorSet","outputs":[],"stateMutability":"nonpayable","type":"function"}]`

	rawABIs = [NumOfAbis]string{
		LogContract:                 rawConsortiumLogAbi,
		SortValidator:               rawConsortiumSortValidatorAbi,
		VerifyHeaders:               rawConsortiumVerifyHeadersAbi,
		PickValidatorSet:            rawConsortiumPickValidatorSetAbi,
		GetDoubleSignSlashingConfig: rawGetDoubleSignSlashingConfigsAbi,
		ValidateFinalityVoteProof:   rawValidateFinalityVoteProofAbi,
		ValidateProofOfPossession:   rawValidateProofOfPossessionAbi,
		PickValidatorSetBeacon:      rawPickValidatorSetBeaconAbi,
	}

	unmarshalledABIs              = [NumOfAbis]*abi.ABI{}
	PickValidatorSetBeaconAddress = common.BytesToAddress([]byte{107})
)

const (
	sortValidatorsMethod         = "sortValidators"
	pickValidatorSetMethod       = "pickValidatorSet"
	logMethod                    = "log"
	getValidatorsMethod          = "getValidatorCandidates"
	totalBalancesMethod          = "totalBalances"
	verifyHeaders                = "validatingDoubleSignProof"
	getHeader                    = "getHeader"
	getDoubleSignSlashingConfigs = "getDoubleSignSlashingConfigs"
	extraVanity                  = 32

	validateFinalityVoteProof = "validateFinalityVoteProof"
	validateProofOfPossession = "validateProofOfPossession"
	maxBlsPublicKeyListLength = 100

	requestSortValidatorSet   = "requestSortValidatorSet"
	maxNumberOfEpochPerPeriod = 144

	// pickValidatorSetBeacon storage slot
	periodSlot                        = 0
	numberOfEpochsSlot                = 1
	numberOfNonRotatingValidatorsSlot = 2
	numberOfRotatingValidatorsSlot    = 3
	validatorArrayStartSlot           = 100
)

func init() {
	for i, rawABI := range rawABIs {
		unmarshalledABI, err := abi.JSON(strings.NewReader(rawABI))
		if err != nil {
			log.Error("Failed to unmarshalled precompiled ABI", "num", i)
		} else {
			unmarshalledABIs[i] = &unmarshalledABI
		}
	}
}

func PrecompiledContractsConsortium(caller ContractRef, evm *EVM) map[common.Address]PrecompiledContract {
	return map[common.Address]PrecompiledContract{
		common.BytesToAddress([]byte{101}): &consortiumLog{},
		common.BytesToAddress([]byte{102}): &consortiumValidatorSorting{caller: caller, evm: evm},
		common.BytesToAddress([]byte{103}): &consortiumVerifyHeaders{caller: caller, evm: evm},
		common.BytesToAddress([]byte{104}): &consortiumPickValidatorSet{caller: caller, evm: evm},
		common.BytesToAddress([]byte{105}): &consortiumValidateFinalityProof{caller: caller, evm: evm},
	}
}

func PrecompiledContractsConsortiumMiko(caller ContractRef, evm *EVM) map[common.Address]PrecompiledContract {
	contracts := PrecompiledContractsConsortium(caller, evm)
	contracts[common.BytesToAddress([]byte{106})] = &consortiumValidateProofOfPossession{caller: caller, evm: evm}
	return contracts
}

func PrecompiledContractsConsortiumTripp(caller ContractRef, evm *EVM) map[common.Address]PrecompiledContract {
	contracts := PrecompiledContractsConsortiumMiko(caller, evm)
	contracts[common.BytesToAddress([]byte{107})] = &pickValidatorSetBeacon{caller: caller, evm: evm}
	return contracts
}

type consortiumLog struct{}

func (c *consortiumLog) RequiredGas(_ []byte) uint64 {
	return 0
}

func (c *consortiumLog) Run(input []byte) ([]byte, error) {
	if os.Getenv("DEBUG") != "true" {
		return input, nil
	}
	_, method, args, err := loadMethodAndArgs(LogContract, input)
	if err != nil {
		return nil, err
	}
	switch method.Name {
	case logMethod:
		if len(args) == 0 {
			return input, nil
		}
		if _, ok := args[0].(string); ok {
			log.Info("[consortiumLog] log message from smart contract", "message", args[0].(string))
		}
	}
	return input, nil
}

func isSystemContractCaller(caller ContractRef, evm *EVM) error {
	// These 2 fields are nil in benchmark only
	if caller != nil && evm != nil {
		if evm.ChainConfig().ConsortiumV2Contracts == nil {
			return errors.New("cannot find consortium v2 contracts")
		}
		if !evm.ChainConfig().ConsortiumV2Contracts.IsSystemContract(caller.Address()) {
			return errors.New("unauthorized sender")
		}
	}

	return nil
}

type consortiumPickValidatorSet struct {
	caller ContractRef
	evm    *EVM
}

func (c *consortiumPickValidatorSet) RequiredGas(input []byte) uint64 {
	// c.evm is nil in benchmark
	if c.evm == nil || c.evm.chainRules.IsMiko {
		// We approximate the number of validators by dividing the length of input by
		// length of address (20). This is likely an overestimation because there are
		// slices of weight, maxValidatorNumber and maxPrioritizedValidatorNumber in
		// the input too.
		return uint64((len(input) / common.AddressLength)) * params.ValidatorSortingBaseGas
	} else {
		return 0
	}
}

func (c *consortiumPickValidatorSet) Run(input []byte) ([]byte, error) {
	if err := isSystemContractCaller(c.caller, c.evm); err != nil {
		return nil, err
	}
	// get method, args from abi
	_, method, args, err := loadMethodAndArgs(PickValidatorSet, input)
	if err != nil {
		return nil, err
	}
	if method.Name != pickValidatorSetMethod {
		return nil, errors.New("invalid method")
	}

	if len(args) != 5 {
		return nil, fmt.Errorf("invalid arguments, expected 5 got %d", len(args))
	}

	// cast args[0] to list addresses
	candidates, ok := args[0].([]common.Address)
	if !ok {
		return nil, errors.New("invalid candidateList argument type")
	}

	// cast args[1] to list big int
	weights, ok := args[1].([]*big.Int)
	if !ok {
		return nil, errors.New("invalid weights argument type")
	}

	trustedWeights, ok := args[2].([]*big.Int)
	if !ok {
		return nil, errors.New("invalid trustedWeights argument type")
	}

	if len(candidates) != len(weights) || len(weights) != len(trustedWeights) {
		return nil, errors.New("array length is mismatch")
	}

	maxValidatorNumber, ok := args[3].(*big.Int)
	if !ok {
		return nil, errors.New("invalid maxValidatorNumber argument type")
	}

	maxPrioritizedValidatorNumber, ok := args[4].(*big.Int)
	if !ok {
		return nil, errors.New("invalid maxPrioritizedValidatorNumber argument type")
	}
	log.Debug("Precompiled pick validator set", "candidates", candidates, "weights", weights, "trustedWeights", trustedWeights, "maxValidatorNumber", maxValidatorNumber, "maxPrioritizedValidatorNumber", maxPrioritizedValidatorNumber)

	newValidatorCount := maxValidatorNumber.Uint64()
	candidateLen := uint64(len(candidates))
	if newValidatorCount > candidateLen {
		newValidatorCount = candidateLen
	}

	candidateMap := createCandidateMap(candidates, trustedWeights)

	// Sort candidates in place
	sortValidators(candidates, weights)

	updateIsTrustedOrganizations(candidates, trustedWeights, candidateMap)

	// If the length of trusted nodes reach the maxPrioritizedValidatorNumber, then the other trusted nodes
	// will be treated as normal nodes
	arrangeValidatorCandidates(candidates, newValidatorCount, trustedWeights, maxPrioritizedValidatorNumber)

	// Since the arrangeValidatorCandidates updates candidates in place.
	// If the length of candidates is greater than newValidatorCount then
	// cut the items down to newValidatorCount
	if candidateLen > newValidatorCount {
		candidates = candidates[:newValidatorCount]
	}

	log.Debug("Precompiled pick validator set", "candidates", candidates)

	return method.Outputs.Pack(candidates)
}

// createCandidateMap maps isTrustedOrganization with candidate before sorting to prevent indexes are changed
func createCandidateMap(candidates []common.Address, isTrustedOrganizations []*big.Int) map[common.Address]*big.Int {
	candidateMap := map[common.Address]*big.Int{}
	for i, address := range candidates {
		candidateMap[address] = isTrustedOrganizations[i]
	}

	return candidateMap
}

// updateIsTrustedOrganizations updates the data of isTrustedOrganizations
func updateIsTrustedOrganizations(candidates []common.Address, isTrustedOrganizations []*big.Int, candidateMap map[common.Address]*big.Int) {
	for i, address := range candidates {
		isTrustedOrganizations[i] = candidateMap[address]
	}
}

func arrangeValidatorCandidates(candidates []common.Address, newValidatorCount uint64, isTrustedOrganizations []*big.Int, maxPrioritizedValidatorNumber *big.Int) {
	var waitingCandidates []common.Address
	var prioritySlotCounter uint64

	for i := 0; i < len(candidates); i++ {
		if isTrustedOrganizations[i].Cmp(big0) > 0 && prioritySlotCounter < maxPrioritizedValidatorNumber.Uint64() {
			candidates[prioritySlotCounter] = candidates[i]
			prioritySlotCounter++
			continue
		}
		waitingCandidates = append(waitingCandidates, candidates[i])
	}

	var waitingCounter uint64
	for i := prioritySlotCounter; i < newValidatorCount; i++ {
		candidates[i] = waitingCandidates[waitingCounter]
		waitingCounter++
	}
}

type consortiumValidatorSorting struct {
	caller ContractRef
	evm    *EVM
}

func (c *consortiumValidatorSorting) RequiredGas(input []byte) uint64 {
	// c.evm is nil in benchmark
	if c.evm == nil || c.evm.chainRules.IsMiko {
		// We approximate the number of validators by dividing the length of input by
		// length of address (20). This is likely an overestimation because there is
		// a slice of weight in the input too.
		return uint64((len(input) / common.AddressLength)) * params.ValidatorSortingBaseGas
	} else {
		return 0
	}
}

func (c *consortiumValidatorSorting) Run(input []byte) ([]byte, error) {
	if err := isSystemContractCaller(c.caller, c.evm); err != nil {
		return nil, err
	}
	// get method, args from abi
	_, method, args, err := loadMethodAndArgs(SortValidator, input)
	if err != nil {
		return nil, err
	}
	if method.Name != sortValidatorsMethod {
		return nil, errors.New("invalid method")
	}
	if len(args) != 2 {
		return nil, fmt.Errorf("invalid arguments, expected 2 got %d", len(args))
	}
	// cast args[0] to list addresses
	validators, ok := args[0].([]common.Address)
	if !ok {
		return nil, errors.New("invalid first argument type")
	}

	// cast args[1] to list big int
	weights, ok := args[1].([]*big.Int)
	if !ok {
		return nil, errors.New("invalid second argument type")
	}

	if len(validators) != len(weights) {
		return nil, errors.New("balances and validators length mismatched")
	}
	sortValidators(validators, weights)

	log.Debug("Precompiled sorting", "validators", validators, "weights", weights)

	return method.Outputs.Pack(validators)
}

func sortValidators(validators []common.Address, weights []*big.Int) {
	if len(validators) < 2 {
		return
	}
	// start sorting validators
	var validatorWithWeights []validatorWithWeight
	for i := range validators {
		validatorWithWeights = append(validatorWithWeights, validatorWithWeight{
			address: validators[i],
			weight:  weights[i],
		})
	}

	sort.Sort(sortByWeight(validatorWithWeights))

	for i, validator := range validatorWithWeights {
		validators[i] = validator.address
		weights[i] = validator.weight
	}
}

func staticCall(evm *EVM, smcAbi abi.ABI, method string, contract, sender common.Address, args ...interface{}) ([]interface{}, error) {
	inputParams, err := smcAbi.Pack(method, args...)
	if err != nil {
		return nil, err
	}
	ret, _, err := evm.StaticCall(AccountRef(sender), contract, inputParams, math.MaxUint64/2)
	if err != nil {
		return nil, err
	}
	out, err := smcAbi.Unpack(method, ret)
	if err != nil {
		return nil, err
	}
	if len(out) == 0 {
		return nil, errors.New("data not found")
	}
	return out, nil
}

func loadMethodAndArgs(contractIndex int, input []byte) (abi.ABI, *abi.Method, []interface{}, error) {
	var (
		pAbi   abi.ABI
		err    error
		method *abi.Method
		args   []interface{}
	)
	if contractIndex < 0 || contractIndex >= len(unmarshalledABIs) || unmarshalledABIs[contractIndex] == nil {
		return abi.ABI{}, nil, nil, errors.New("invalid contract index")
	}
	pAbi = *unmarshalledABIs[contractIndex]
	if method, err = pAbi.MethodById(input); err != nil {
		return abi.ABI{}, nil, nil, err
	}
	if args, err = method.Inputs.Unpack(input[4:]); err != nil {
		return abi.ABI{}, nil, nil, err
	}
	return pAbi, method, args, nil
}

const doubleSigningOffsetTest = 28800

type consortiumVerifyHeaders struct {
	caller ContractRef
	evm    *EVM

	test bool
}

func (c *consortiumVerifyHeaders) RequiredGas(_ []byte) uint64 {
	// c.evm is nil in benchmark
	if c.evm == nil || c.evm.chainRules.IsMiko {
		return params.VerifyFinalityHeadersProofGas
	} else {
		return 0
	}
}

func (c *consortiumVerifyHeaders) Run(input []byte) ([]byte, error) {
	if err := isSystemContractCaller(c.caller, c.evm); err != nil {
		return nil, err
	}
	// get method, args from abi
	smcAbi, method, args, err := loadMethodAndArgs(VerifyHeaders, input)
	if err != nil {
		return nil, err
	}
	if method.Name != verifyHeaders {
		return nil, errors.New("invalid method")
	}
	if len(args) != 3 {
		return nil, fmt.Errorf("invalid arguments, expected 2 got %d", len(args))
	}
	consensusAddr, ok := args[0].(common.Address)
	if !ok {
		return nil, errors.New("invalid first argument type")
	}
	// decode header1, header2
	var blockHeader1, blockHeader2 types.BlockHeader
	if err := c.unpack(smcAbi, &blockHeader1, args[1].([]byte)); err != nil {
		return nil, err
	}
	if err := c.unpack(smcAbi, &blockHeader2, args[2].([]byte)); err != nil {
		return nil, err
	}
	output := c.verify(consensusAddr, blockHeader1, blockHeader2)
	return smcAbi.Methods[verifyHeaders].Outputs.Pack(output)
}

func (c *consortiumVerifyHeaders) unpack(smcAbi abi.ABI, v interface{}, input []byte) error {
	// use `getHeader` abi which contains params defined in `BlockHeader` to unmarshal `input` data into `BlockHeader`
	args := smcAbi.Methods[getHeader].Inputs
	output, err := args.Unpack(input)
	if err != nil {
		return err
	}
	return args.Copy(v, output)
}

func (c *consortiumVerifyHeaders) getSigner(header types.BlockHeader) (common.Address, error) {
	if header.Number == nil || header.Timestamp > uint64(time.Now().Unix()) || len(header.ExtraData) < crypto.SignatureLength {
		return common.Address{}, errors.New("invalid header")
	}
	signature := header.ExtraData[len(header.ExtraData)-crypto.SignatureLength:]

	// Recover the public key and the Ethereum address
	pubkey, err := crypto.Ecrecover(SealHash(header.ToHeader(), header.ChainId).Bytes(), signature)
	if err != nil {
		return common.Address{}, err
	}
	r := new(big.Int).SetBytes(signature[:32])
	s := new(big.Int).SetBytes(signature[32:64])
	v := signature[64]
	if !crypto.ValidateSignatureValues(v, r, s, true) {
		return common.Address{}, err
	}

	var signer common.Address
	copy(signer[:], crypto.Keccak256(pubkey[1:])[12:])

	return signer, nil
}

func (c *consortiumVerifyHeaders) verify(consensusAddr common.Address, header1, header2 types.BlockHeader) bool {
	var maxOffset *big.Int

	// c.evm s nil in benchmark, so we skip this check in benchmark
	if c.evm != nil && !c.evm.chainConfig.IsConsortiumV2(header1.Number) {
		return false
	}
	if header1.ToHeader().ParentHash.Hex() != header2.ToHeader().ParentHash.Hex() {
		return false
	}
	if len(header1.ExtraData) < crypto.SignatureLength || len(header2.ExtraData) < crypto.SignatureLength {
		return false
	}
	if bytes.Equal(SealHash(header1.ToHeader(), header1.ChainId).Bytes(), SealHash(header2.ToHeader(), header2.ChainId).Bytes()) {
		return false
	}
	signer1, err := c.getSigner(header1)
	if err != nil {
		log.Trace("[consortiumVerifyHeaders][verify] error while getting signer from header1", "err", err)
		return false
	}
	signer2, err := c.getSigner(header2)
	if err != nil {
		log.Trace("[consortiumVerifyHeaders][verify] error while getting signer from header2", "err", err)
		return false
	}
	if unmarshalledABIs[GetDoubleSignSlashingConfig] == nil {
		return false
	}
	methodAbi := *unmarshalledABIs[GetDoubleSignSlashingConfig]

	if c.test {
		maxOffset = big.NewInt(doubleSigningOffsetTest)
	} else {
		if c.evm.chainConfig.ConsortiumV2Contracts == nil {
			return false
		} else {
			rawMaxOffset, err := staticCall(
				c.evm,
				methodAbi,
				getDoubleSignSlashingConfigs,
				c.evm.chainConfig.ConsortiumV2Contracts.SlashIndicator,
				common.Address{},
			)
			if err != nil {
				log.Error("Failed to get double sign config", "err", err)
				return false
			}
			if len(rawMaxOffset) < 3 {
				log.Error("Invalid output length when getting double sign config", "length", len(rawMaxOffset))
				return false
			}
			maxOffset = rawMaxOffset[2].(*big.Int)
		}
	}

	// c.evm is nil in benchmark, so we skip this check in benchmark
	if c.evm != nil {
		currentBlock := c.evm.Context.BlockNumber
		// What if current block < header1.Number?
		if currentBlock.Cmp(header1.Number) > 0 && new(big.Int).Sub(currentBlock, header1.Number).Cmp(maxOffset) > 0 {
			return false
		}
	}

	return signer1.Hex() == signer2.Hex() &&
		signer2.Hex() == header2.Benificiary.Hex() &&
		bytes.Equal(consensusAddr.Bytes(), signer1.Bytes())
}

// SealHash returns the hash of a block prior to it being sealed.
func SealHash(header *types.Header, chainId *big.Int) (hash common.Hash) {
	hasher := sha3.NewLegacyKeccak256()
	encodeSigHeader(hasher, header, chainId)
	hasher.Sum(hash[:0])
	return hash
}

func encodeSigHeader(w io.Writer, header *types.Header, chainId *big.Int) {
	err := rlp.Encode(w, []interface{}{
		chainId,
		header.ParentHash,
		header.UncleHash,
		header.Coinbase,
		header.Root,
		header.TxHash,
		header.ReceiptHash,
		header.Bloom,
		header.Difficulty,
		header.Number,
		header.GasLimit,
		header.GasUsed,
		header.Time,
		header.Extra[:len(header.Extra)-crypto.SignatureLength], // Yes, this will panic if extra is too short
		header.MixDigest,
		header.Nonce,
	})
	if err != nil {
		panic("can't encode: " + err.Error())
	}
}

type consortiumValidateFinalityProof struct {
	caller ContractRef
	evm    *EVM
}

func (contract *consortiumValidateFinalityProof) RequiredGas(input []byte) uint64 {
	return params.ValidateFinalityProofGas
}

func (contract *consortiumValidateFinalityProof) Run(input []byte) ([]byte, error) {
	if err := isSystemContractCaller(contract.caller, contract.evm); err != nil {
		return nil, err
	}
	_, method, args, err := loadMethodAndArgs(ValidateFinalityVoteProof, input)
	if err != nil {
		return nil, err
	}
	if method.Name != validateFinalityVoteProof {
		return nil, errors.New("invalid method")
	}
	if len(args) != 5 {
		return nil, fmt.Errorf("invalid arguments, expect 5 got %d", len(args))
	}

	rawVoterPublicKey, ok := args[0].([]byte)
	if !ok {
		return nil, errors.New("invalid voter public key")
	}

	targetBlockNumber, ok := args[1].(*big.Int)
	if !ok {
		return nil, errors.New("invalid target block number")
	}
	if !targetBlockNumber.IsUint64() {
		return nil, errors.New("malformed target block number")
	}

	targetBlockHashes, ok := args[2].([2][32]byte)
	if !ok {
		return nil, errors.New("invalid target block hashes")
	}

	if targetBlockHashes[0] == targetBlockHashes[1] {
		return nil, errors.New("block hash is the same")
	}

	listOfRawPublicKey, ok := args[3].([2][][]byte)
	if !ok {
		return nil, errors.New("invalid target block number")
	}

	rawAggregatedSignatures, ok := args[4].([2][]byte)
	if !ok {
		return nil, errors.New("invalid aggregated signature")
	}

	voterPublicKey, err := blst.PublicKeyFromBytes(rawVoterPublicKey)
	if err != nil {
		return nil, errors.New("malformed voter public key")
	}

	var listOfPublicKey [2][]blsCommon.PublicKey
	for block := range listOfRawPublicKey {
		voterInPublicKeyList := false
		for _, rawKey := range listOfRawPublicKey[block] {
			publicKey, err := blst.PublicKeyFromBytes(rawKey)
			if err != nil {
				return nil, errors.New("malformed public key in list of public keys")
			}

			if publicKey.Equals(voterPublicKey) {
				voterInPublicKeyList = true
			}

			listOfPublicKey[block] = append(listOfPublicKey[block], publicKey)
		}

		if !voterInPublicKeyList {
			return nil, errors.New("reported voter does not in public key list")
		}
	}

	for _, list := range listOfPublicKey {
		if len(list) > maxBlsPublicKeyListLength {
			return nil, errors.New("public key list is too long")
		}
	}

	var aggregatedSignature [2]blsCommon.Signature
	for block, rawSignature := range rawAggregatedSignatures {
		signature, err := blst.SignatureFromBytes(rawSignature)
		if err != nil {
			return nil, errors.New("malformed signature")
		}

		aggregatedSignature[block] = signature
	}

	for block := 0; block < 2; block++ {
		voteData := types.VoteData{
			TargetNumber: targetBlockNumber.Uint64(),
			TargetHash:   targetBlockHashes[block],
		}
		digest := voteData.Hash()
		if !aggregatedSignature[block].FastAggregateVerify(listOfPublicKey[block], digest) {
			return nil, errors.New("failed to verify signature")
		}
	}

	return method.Outputs.Pack(true)
}

type consortiumValidateProofOfPossession struct {
	caller ContractRef
	evm    *EVM
}

func (contract *consortiumValidateProofOfPossession) RequiredGas(input []byte) uint64 {
	return params.ValidateProofOfPossession
}

func (contract *consortiumValidateProofOfPossession) Run(input []byte) ([]byte, error) {
	if err := isSystemContractCaller(contract.caller, contract.evm); err != nil {
		return nil, err
	}
	_, method, args, err := loadMethodAndArgs(ValidateProofOfPossession, input)
	if err != nil {
		return nil, err
	}
	if method.Name != validateProofOfPossession {
		return nil, errors.New("invalid method")
	}
	if len(args) != 2 {
		return nil, fmt.Errorf("invalid arguments, expect 2 got %d", len(args))
	}

	rawPublicKey, ok := args[0].([]byte)
	if !ok {
		return nil, errors.New("invalid voter public key")
	}

	rawSignature, ok := args[1].([]byte)
	if !ok {
		return nil, errors.New("invalid proof signature")
	}

	blsPublicKey, err := blst.PublicKeyFromBytes(rawPublicKey)
	if err != nil {
		return nil, errors.New("malformed voter public key")
	}

	blsSignature, err := blst.SignatureFromBytes(rawSignature)
	if err != nil {
		return nil, errors.New("malformed proof signature")
	}

	if !blsSignature.VerifyProof(blsPublicKey, rawPublicKey) {
		return nil, errors.New("invalid possession proof")
	}

	return method.Outputs.Pack(true)
}

// This precompiled contract has 2 methods
// - requestSortValidatorSet: is called at the end of old period, with the information about
// beacon, staked amount of validator candidates. Sort and pick validator set for each epoch
// in the next period and store the result into contract's storage
// - pickValidatorSet: get the validator (block producer) list of an epoch
//
// The contract storage layout
// Storage slot 0: period number
// Storage slot 1: number of epochs in a period
// Storage slot 2: number of non-rotating validators
// Storage slot 3: number of rotating validators
//
// Storage slot 100: store N non-rotating validator addresses
// Storage slot 100 + N: 2D array of rotating validators in each epoch
// The storage slot of rotating validator ith in epoch jth is
// 100 + N + j * number of rotating validators in each epoch + i
type pickValidatorSetBeacon struct {
	caller ContractRef
	evm    *EVM
	// This is true only when running benchmark
	skipPeriodCheck bool
}

func (contract *pickValidatorSetBeacon) RequiredGas(input []byte) uint64 {
	_, method, args, err := loadMethodAndArgs(PickValidatorSetBeacon, input)
	if err != nil {
		return math.MaxUint64
	}

	if method.Name == requestSortValidatorSet {
		numRotatingValidator, ok := args[4].(*big.Int)
		if !ok {
			return math.MaxUint64
		}
		return numRotatingValidator.Uint64() * maxNumberOfEpochPerPeriod * params.SstoreClearGas
	} else {
		return maxNumberOfEpochPerPeriod * params.SloadGasEIP2200
	}
}

func (contract *pickValidatorSetBeacon) Run(input []byte) ([]byte, error) {
	if err := isSystemContractCaller(contract.caller, contract.evm); err != nil {
		return nil, err
	}

	_, method, args, err := loadMethodAndArgs(PickValidatorSetBeacon, input)
	if err != nil {
		return nil, err
	}

	if method.Name == requestSortValidatorSet {
		return contract.requestSortValidator(method, args)
	} else {
		return contract.pickValidator(method, args)
	}
}

func (contract *pickValidatorSetBeacon) pickValidator(method *abi.Method, args []interface{}) ([]byte, error) {
	if len(args) != 2 {
		return nil, fmt.Errorf("invalid arguments, expected 2 got %d", len(args))
	}
	period, ok := args[0].(*big.Int)
	if !ok {
		return nil, errors.New("invalid period argument type")
	}
	epoch, ok := args[1].(*big.Int)
	if !ok {
		return nil, errors.New("invalid epoch argument type")
	}
	validators, err := contract.readValidatorAtEpoch(int(period.Int64()), int(epoch.Int64()))
	if err != nil {
		return nil, err
	}
	output, err := method.Outputs.Pack(validators)
	if err != nil {
		return nil, err
	}

	return output, nil
}

func (contract *pickValidatorSetBeacon) readMetadataFromStorage() (
	period int,
	numOfEpochs int,
	numOfNonRotatingValidators int,
	numOfRotatingValidators int,
) {
	stateDB := contract.evm.StateDB
	contractAddress := PickValidatorSetBeaconAddress

	periodValue := stateDB.GetState(contractAddress, common.BigToHash(big.NewInt(periodSlot)))
	numOfEpochsValue := stateDB.GetState(contractAddress, common.BigToHash(big.NewInt(numberOfEpochsSlot)))
	numOfNonRotatingValidatorsValue := stateDB.GetState(
		contractAddress,
		common.BigToHash(big.NewInt(numberOfNonRotatingValidatorsSlot)),
	)
	numOfRotatingValidatorsValue := stateDB.GetState(
		contractAddress,
		common.BigToHash(big.NewInt(numberOfRotatingValidatorsSlot)),
	)

	return int(periodValue.Big().Int64()),
		int(numOfEpochsValue.Big().Int64()),
		int(numOfNonRotatingValidatorsValue.Big().Int64()),
		int(numOfRotatingValidatorsValue.Big().Int64())
}

func (contract *pickValidatorSetBeacon) readValidatorAtEpoch(period int, epochNumber int) ([]common.Address, error) {
	stateDB := contract.evm.StateDB
	contractAddress := PickValidatorSetBeaconAddress
	storedPeriod, numOfEpochs, numOfNonRotatingValidators, numOfRotatingValidators := contract.readMetadataFromStorage()

	if storedPeriod != period {
		return nil, fmt.Errorf("queried period mismatches with stored one, queried: %d, stored: %d", period, storedPeriod)
	}

	// Queried epoch number starts from 1, but the stored index starts from 0.
	// So we need to subtract 1 from queried epoch to get the correct stored index.
	epochNumber = epochNumber - 1
	if epochNumber < 0 || epochNumber > numOfEpochs {
		return nil, errors.New("invalid epoch number")
	}

	consensusAddrs := make([]common.Address, 0, numOfNonRotatingValidators+numOfRotatingValidators)
	// Read non-rotating validators
	for i := 0; i < numOfNonRotatingValidators; i++ {
		value := stateDB.GetState(contractAddress, common.BigToHash(big.NewInt(int64(validatorArrayStartSlot+i))))
		address := common.BytesToAddress(value.Bytes())
		consensusAddrs = append(consensusAddrs, address)
	}

	// Read rotating validators
	startSlot := validatorArrayStartSlot + numOfNonRotatingValidators + epochNumber*numOfRotatingValidators
	for i := 0; i < numOfRotatingValidators; i++ {
		value := stateDB.GetState(contractAddress, common.BigToHash(big.NewInt(int64(startSlot+i))))
		address := common.BytesToAddress(value.Bytes())
		consensusAddrs = append(consensusAddrs, address)
	}

	return consensusAddrs, nil
}

// rotating validators dimension: [numOfEpochs][numOfRotatingValidators]common.Address
func (contract *pickValidatorSetBeacon) writeToStorage(
	period int,
	nonRotatingValidators []common.Address,
	rotatingValidators [][]common.Address,
) {
	stateDB := contract.evm.StateDB
	contractAddress := PickValidatorSetBeaconAddress

	var (
		numOfRotatingValidators int
		numOfEpochs             int = maxNumberOfEpochPerPeriod
	)
	if len(rotatingValidators) != 0 {
		numOfRotatingValidators = len(rotatingValidators[0])
		numOfEpochs = len(rotatingValidators)
	}

	stateDB.SetState(
		contractAddress,
		common.BigToHash(big.NewInt(periodSlot)),
		common.BigToHash(big.NewInt(int64(period))),
	)
	stateDB.SetState(
		contractAddress,
		common.BigToHash(big.NewInt(numberOfEpochsSlot)),
		common.BigToHash(big.NewInt(int64(numOfEpochs))),
	)
	stateDB.SetState(
		contractAddress,
		common.BigToHash(big.NewInt(numberOfNonRotatingValidatorsSlot)),
		common.BigToHash(big.NewInt(int64(len(nonRotatingValidators)))),
	)
	stateDB.SetState(
		contractAddress,
		common.BigToHash(big.NewInt(numberOfRotatingValidatorsSlot)),
		common.BigToHash(big.NewInt(int64(numOfRotatingValidators))),
	)

	// Write non-rotating validators
	for i, validator := range nonRotatingValidators {
		stateDB.SetState(
			contractAddress,
			common.BigToHash(big.NewInt(int64(validatorArrayStartSlot+i))),
			common.BytesToHash(validator.Bytes()),
		)
	}

	// Write rotating validators
	startSlot := validatorArrayStartSlot + len(nonRotatingValidators)
	for epochNumber, validators := range rotatingValidators {
		for valIndex, validator := range validators {
			stateDB.SetState(
				contractAddress,
				common.BigToHash(big.NewInt(int64(startSlot+epochNumber*numOfRotatingValidators+valIndex))),
				common.BytesToHash(validator.Bytes()),
			)
		}
	}
}

func (contract *pickValidatorSetBeacon) cleanupStorage(oldEnd int, newEnd int) {
	stateDB := contract.evm.StateDB
	contractAddress := PickValidatorSetBeaconAddress

	for i := newEnd; i < oldEnd; i++ {
		stateDB.SetState(contractAddress, common.BigToHash(big.NewInt(int64(i))), common.Hash{})
	}
}

// calculateValidatorWeight calculates the weight to choose rotating validators
// uint256 random = hash(beacon || epochNumber || address)
// random = higher128(random) xor lower128(random)
// stakeAmount = stakedAmount / 10**18
// weight =  random * stakedAmount * stakedAmount
// with || is the concatenation operation
func calculateValidatorWeight(
	hasher crypto.KeccakState,
	beacon *big.Int,
	epochNumber int,
	consensusAddr common.Address,
	stakedAmount *big.Int,
) *big.Int {
	var (
		hashbuf common.Hash
		output  [32]byte
	)

	hasher.Reset()
	beacon.FillBytes(output[:])
	hasher.Write(output[:])

	big.NewInt(int64(epochNumber)).FillBytes(output[:])
	hasher.Write(output[:])

	new(big.Int).SetBytes(consensusAddr.Bytes()).FillBytes(output[:])
	hasher.Write(output[:])

	hasher.Read(hashbuf[:])

	ether := new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil)
	amount := new(big.Int).Div(stakedAmount, ether)
	amount = new(big.Int).Mul(amount, amount)

	hashResult := make([]byte, len(hashbuf)/2)
	for i := range hashResult {
		hashResult[i] = hashbuf[i] ^ hashbuf[i+16]
	}

	randomHash := new(big.Int).SetBytes(hashResult)
	return new(big.Int).Mul(randomHash, amount)
}

type validatorWithWeight struct {
	address common.Address
	weight  *big.Int
}

// Sort validator based on weight in descending order
type sortByWeight []validatorWithWeight

func (validators sortByWeight) Len() int { return len(validators) }
func (validators sortByWeight) Swap(i, j int) {
	validators[i], validators[j] = validators[j], validators[i]
}
func (validators sortByWeight) Less(i, j int) bool {
	cmp := validators[i].weight.Cmp(validators[j].weight)
	if cmp != 0 {
		return cmp > 0
	} else {
		return bytes.Compare(validators[i].address[:], validators[j].address[:]) > 0
	}
}

// pickNonRotatingValidator assumes that len(consensusAddress) > numGovernanceValidator + numStandardValidator.
// It always picks numGovernanceValidator + numStandardValidator validators. In case, the available governance
// validators are fewer than the maximum number, the remaining number is transfered to standard maximum number
// (if we cannot pick enough governance validator, we will pick more standard validator).
func pickNonRotatingValidator(
	numGovernanceValidator int,
	numStandardValidator int,
	consensusAddress []common.Address,
	stakedAmounts []*big.Int,
	isGovernanceValidator []*big.Int,
) []common.Address {
	var governanceValidator []validatorWithWeight
	for i := range isGovernanceValidator {
		if isGovernanceValidator[i].Cmp(common.Big1) == 0 {
			governanceValidator = append(governanceValidator, validatorWithWeight{
				address: consensusAddress[i],
				weight:  stakedAmounts[i],
			})
		}
	}

	if len(governanceValidator) > numGovernanceValidator {
		sort.Sort(sortByWeight(governanceValidator))
		governanceValidator = governanceValidator[:numGovernanceValidator]
	} else {
		// The number of governance validators is fewer than the maximum governance
		// validator, the remainning number is transfered to standard validator case
		numStandardValidator += numGovernanceValidator - len(governanceValidator)
	}

	chosen := make(map[common.Address]struct{})
	for _, validator := range governanceValidator {
		chosen[validator.address] = struct{}{}
	}

	var standardValidator []validatorWithWeight
	for i := range consensusAddress {
		if _, ok := chosen[consensusAddress[i]]; !ok {
			standardValidator = append(standardValidator, validatorWithWeight{
				address: consensusAddress[i],
				weight:  stakedAmounts[i],
			})
		}
	}

	if len(standardValidator) > numStandardValidator {
		sort.Sort(sortByWeight(standardValidator))
		standardValidator = standardValidator[:numStandardValidator]
	}

	nonRotatingValidator := make([]common.Address, 0, len(governanceValidator)+len(standardValidator))
	for _, validator := range governanceValidator {
		nonRotatingValidator = append(nonRotatingValidator, validator.address)
	}
	for _, validator := range standardValidator {
		nonRotatingValidator = append(nonRotatingValidator, validator.address)
	}

	return nonRotatingValidator
}

func (contract *pickValidatorSetBeacon) requestSortValidator(method *abi.Method, args []interface{}) ([]byte, error) {
	if len(args) != 8 {
		return nil, fmt.Errorf("invalid arguments, expected 4 got %d", len(args))
	}
	beacon, ok := args[0].(*big.Int)
	if !ok {
		return nil, errors.New("invalid beacon argument type")
	}
	period, ok := args[1].(*big.Int)
	if !ok {
		return nil, errors.New("invalid period argument type")
	}
	numGovernanceValidator, ok := args[2].(*big.Int)
	if !ok {
		return nil, errors.New("invalid number of governance validators argument type")
	}
	numStandardValidator, ok := args[3].(*big.Int)
	if !ok {
		return nil, errors.New("invalid number of standard validators argument type")
	}
	numRotatingValidator, ok := args[4].(*big.Int)
	if !ok {
		return nil, errors.New("invalid number of rotating validators argument type")
	}
	consensusAddrs, ok := args[5].([]common.Address)
	if !ok {
		return nil, errors.New("invalid consensus address argument type")
	}
	stakedAmounts, ok := args[6].([]*big.Int)
	if !ok {
		return nil, errors.New("invalid staked amount argument type")
	}
	isGovernanceValidator, ok := args[7].([]*big.Int)
	if !ok {
		return nil, errors.New("invalid is govnernance validator argument type")
	}
	if len(consensusAddrs) != len(stakedAmounts) {
		return nil, errors.New("consensus addresses and staked amounts length mismatched")
	}
	if len(isGovernanceValidator) != len(stakedAmounts) {
		return nil, errors.New("is governance validator and staked amounts length mismatched")
	}

	oldPeriod, numOfEpochs, numOfPreviousNonRotatingValidators, numOfPreviousRotatingValidators := contract.readMetadataFromStorage()
	if !contract.skipPeriodCheck && oldPeriod >= int(period.Int64()) {
		return nil, errors.New("new period is fewer or equals to stored period")
	}

	// The number of validator candidates is too few, just pick all candidates. In this case, number of non-rotating
	// validators is higher than the sum of number of governance validators and number of standard validators
	if len(consensusAddrs) <=
		int(numGovernanceValidator.Int64())+int(numStandardValidator.Int64())+int(numRotatingValidator.Int64()) {

		contract.writeToStorage(int(period.Int64()), consensusAddrs, nil)

		oldEnd := validatorArrayStartSlot + numOfPreviousNonRotatingValidators + numOfEpochs*numOfPreviousRotatingValidators
		newEnd := validatorArrayStartSlot + len(consensusAddrs)
		contract.cleanupStorage(oldEnd, newEnd)
		return nil, nil
	}

	// Pick governance and standard validators, this set does not change across
	// different epoch
	nonRotatingValidators := pickNonRotatingValidator(
		int(numGovernanceValidator.Int64()),
		int(numStandardValidator.Int64()),
		consensusAddrs,
		stakedAmounts,
		isGovernanceValidator,
	)

	chosen := make(map[common.Address]struct{})
	for _, validator := range nonRotatingValidators {
		chosen[validator] = struct{}{}
	}

	var rotatingValidatorCandidates []validatorWithWeight
	for i := range consensusAddrs {
		if _, ok := chosen[consensusAddrs[i]]; !ok {
			rotatingValidatorCandidates = append(rotatingValidatorCandidates, validatorWithWeight{
				address: consensusAddrs[i],
				weight:  stakedAmounts[i],
			})
		}
	}

	rotatingValidators := make([][]common.Address, maxNumberOfEpochPerPeriod)
	for epochNumber := range rotatingValidators {
		rotatingValidators[epochNumber] = make([]common.Address, numRotatingValidator.Int64())
	}

	rotatingValidatorCandidatesWithWeight := make([]validatorWithWeight, len(rotatingValidatorCandidates))
	hasher := sha3.NewLegacyKeccak256().(crypto.KeccakState)
	// The epoch number starts from 1 ranges from [1, maxNumberOfEpochPerPeriod]
	for epochNumber := 1; epochNumber <= maxNumberOfEpochPerPeriod; epochNumber++ {
		for i, validator := range rotatingValidatorCandidates {
			weight := calculateValidatorWeight(
				hasher,
				beacon,
				epochNumber,
				validator.address,
				validator.weight,
			)

			rotatingValidatorCandidatesWithWeight[i] = validatorWithWeight{
				address: validator.address,
				weight:  weight,
			}
		}

		sort.Sort(sortByWeight(rotatingValidatorCandidatesWithWeight))
		// Pick numRotatingValidator with the highest weight
		for i := 0; i < int(numRotatingValidator.Int64()); i++ {
			rotatingValidators[epochNumber-1][i] = rotatingValidatorCandidatesWithWeight[i].address
		}
	}
	contract.writeToStorage(int(period.Int64()), nonRotatingValidators, rotatingValidators)

	// Clean up the storage if old validator list is larger than the new one
	oldEnd := validatorArrayStartSlot + numOfPreviousNonRotatingValidators + numOfEpochs*numOfPreviousRotatingValidators
	newEnd := validatorArrayStartSlot + len(nonRotatingValidators) + maxNumberOfEpochPerPeriod*int(numRotatingValidator.Int64())
	contract.cleanupStorage(oldEnd, newEnd)

	return nil, nil
}
