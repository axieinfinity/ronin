package vm

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"math/big"
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
	VerifyHeadersVenoki
	PickValidatorSet
	GetDoubleSignSlashingConfig
	ValidateFinalityVoteProof
	ValidateProofOfPossession
	NumOfAbis
)

var (
	rawConsortiumLogAbi                = `[{"inputs":[{"internalType":"string","name":"message","type":"string"}],"name":"log","outputs":[],"stateMutability":"nonpayable","type":"function"}]`
	rawConsortiumSortValidatorAbi      = `[{"inputs":[],"stateMutability":"nonpayable","type":"constructor"},{"inputs":[{"internalType":"address[]","name":"validators","type":"address[]"},{"internalType":"uint256[]","name":"weights","type":"uint256[]"}],"name":"sortValidators","outputs":[{"internalType":"address[]","name":"_validators","type":"address[]"}],"stateMutability":"view","type":"function"}]`
	rawConsortiumVerifyHeadersAbi      = `[{"outputs":[],"name":"getHeader","inputs":[{"internalType":"uint256","name":"chainId","type":"uint256"},{"internalType":"bytes32","name":"parentHash","type":"bytes32"},{"internalType":"bytes32","name":"ommersHash","type":"bytes32"},{"internalType":"address","name":"coinbase","type":"address"},{"internalType":"bytes32","name":"stateRoot","type":"bytes32"},{"internalType":"bytes32","name":"transactionsRoot","type":"bytes32"},{"internalType":"bytes32","name":"receiptsRoot","type":"bytes32"},{"internalType":"uint8[256]","name":"logsBloom","type":"uint8[256]"},{"internalType":"uint256","name":"difficulty","type":"uint256"},{"internalType":"uint256","name":"number","type":"uint256"},{"internalType":"uint64","name":"gasLimit","type":"uint64"},{"internalType":"uint64","name":"gasUsed","type":"uint64"},{"internalType":"uint64","name":"timestamp","type":"uint64"},{"internalType":"bytes","name":"extraData","type":"bytes"},{"internalType":"bytes32","name":"mixHash","type":"bytes32"},{"internalType":"uint64","name":"nonce","type":"uint64"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"address","name":"consensusAddr","type":"address"},{"internalType":"bytes","name":"header1","type":"bytes"},{"internalType":"bytes","name":"header2","type":"bytes"}],"name":"validatingDoubleSignProof","outputs":[{"internalType":"bool","name":"","type":"bool"}],"stateMutability":"view","type":"function"}]`
	rawConsortiumVerifyHeadersV2Abi    = `[{"outputs":[],"name":"getHeader","inputs":[{"internalType":"uint256","name":"chainId","type":"uint256"},{"internalType":"bytes32","name":"parentHash","type":"bytes32"},{"internalType":"bytes32","name":"ommersHash","type":"bytes32"},{"internalType":"address","name":"coinbase","type":"address"},{"internalType":"bytes32","name":"stateRoot","type":"bytes32"},{"internalType":"bytes32","name":"transactionsRoot","type":"bytes32"},{"internalType":"bytes32","name":"receiptsRoot","type":"bytes32"},{"internalType":"uint8[256]","name":"logsBloom","type":"uint8[256]"},{"internalType":"uint256","name":"difficulty","type":"uint256"},{"internalType":"uint256","name":"number","type":"uint256"},{"internalType":"uint64","name":"gasLimit","type":"uint64"},{"internalType":"uint64","name":"gasUsed","type":"uint64"},{"internalType":"uint64","name":"timestamp","type":"uint64"},{"internalType":"bytes","name":"extraData","type":"bytes"},{"internalType":"bytes32","name":"mixHash","type":"bytes32"},{"internalType":"uint64","name":"nonce","type":"uint64"},{"internalType":"uint256","name":"baseFee","type":"uint256"},{"internalType":"uint64","name":"blobGasUsed","type":"uint64"},{"internalType":"uint64","name":"excessBlobGas","type":"uint64"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"address","name":"consensusAddr","type":"address"},{"internalType":"bytes","name":"header1","type":"bytes"},{"internalType":"bytes","name":"header2","type":"bytes"}],"name":"validatingDoubleSignProof","outputs":[{"internalType":"bool","name":"","type":"bool"}],"stateMutability":"view","type":"function"}]`
	rawConsortiumPickValidatorSetAbi   = `[{"inputs":[],"stateMutability":"nonpayable","type":"constructor"},{"inputs":[{"internalType":"address[]","name":"_candidates","type":"address[]"},{"internalType":"uint256[]","name":"_weights","type":"uint256[]"},{"internalType":"uint256[]","name":"_trustedWeights","type":"uint256[]"},{"internalType":"uint256","name":"_maxValidatorNumber","type":"uint256"},{"internalType":"uint256","name":"_maxPrioritizedValidatorNumber","type":"uint256"}],"name":"pickValidatorSet","outputs":[{"internalType":"address[]","name":"_validators","type":"address[]"}],"stateMutability":"view","type":"function"}]`
	rawGetDoubleSignSlashingConfigsAbi = `[{"inputs":[],"name":"getDoubleSignSlashingConfigs","outputs":[{"internalType":"uint256","name":"","type":"uint256"},{"internalType":"uint256","name":"","type":"uint256"},{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"}]`
	rawValidateFinalityVoteProofAbi    = `[{"inputs":[{"internalType":"bytes","name":"voterPublicKey","type":"bytes"},{"internalType":"uint256","name":"targetBlockNumber","type":"uint256"},{"internalType":"bytes32[2]","name":"targetBlockHash","type":"bytes32[2]"},{"internalType":"bytes[][2]","name":"listOfPublicKey","type":"bytes[][2]"},{"internalType":"bytes[2]","name":"aggregatedSignature","type":"bytes[2]"}],"name":"validateFinalityVoteProof","outputs":[{"internalType":"bool","name":"","type":"bool"}],"stateMutability":"view","type":"function"}]`
	rawValidateProofOfPossessionAbi    = `[{"inputs":[{"internalType":"bytes","name":"publicKey","type":"bytes"},{"internalType":"bytes","name":"signature","type":"bytes"}],"name":"validateProofOfPossession","outputs":[{"internalType":"bool","name":"","type":"bool"}],"stateMutability":"view","type":"function"}]`

	rawABIs = [NumOfAbis]string{
		LogContract:                 rawConsortiumLogAbi,
		SortValidator:               rawConsortiumSortValidatorAbi,
		VerifyHeaders:               rawConsortiumVerifyHeadersAbi,
		VerifyHeadersVenoki:         rawConsortiumVerifyHeadersV2Abi,
		PickValidatorSet:            rawConsortiumPickValidatorSetAbi,
		GetDoubleSignSlashingConfig: rawGetDoubleSignSlashingConfigsAbi,
		ValidateFinalityVoteProof:   rawValidateFinalityVoteProofAbi,
		ValidateProofOfPossession:   rawValidateProofOfPossessionAbi,
	}

	unmarshalledABIs = [NumOfAbis]*abi.ABI{}
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

type PrecompiledContractWithInit interface {
	PrecompiledContract
	Init(caller ContractRef, evm *EVM)
}

type consortiumLog struct{}

func (c *consortiumLog) Init(_ ContractRef, _ *EVM) {
}

func (c *consortiumLog) RequiredGas(_ []byte) uint64 {
	return 0
}

func (c *consortiumLog) Run(input []byte) ([]byte, error) {
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

func (c *consortiumPickValidatorSet) Init(caller ContractRef, evm *EVM) {
	c.caller = caller
	c.evm = evm
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
		return nil, errors.New(fmt.Sprintf("invalid arguments, expected 5 got %d", len(args)))
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

func (c *consortiumValidatorSorting) Init(caller ContractRef, evm *EVM) {
	c.caller = caller
	c.evm = evm
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
		return nil, errors.New(fmt.Sprintf("invalid arguments, expected 2 got %d", len(args)))
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

type SortableValidators struct {
	validators []common.Address
	weights    []*big.Int
}

func (s *SortableValidators) Len() int {
	return len(s.validators)
}

func (s *SortableValidators) Less(i, j int) bool {
	cmp := s.weights[i].Cmp(s.weights[j])

	if cmp == 0 {
		return new(big.Int).SetBytes(s.validators[i].Bytes()).Cmp(new(big.Int).SetBytes(s.validators[j].Bytes())) > 0
	}

	return cmp > 0
}

func (s *SortableValidators) Swap(i, j int) {
	s.validators[i], s.validators[j] = s.validators[j], s.validators[i]
	s.weights[i], s.weights[j] = s.weights[j], s.weights[i]
}

func sortValidators(validators []common.Address, weights []*big.Int) {
	if len(validators) < 2 {
		return
	}
	// start sorting validators
	vals := &SortableValidators{validators: validators, weights: weights}
	sort.Sort(vals)
	return
}

type SmartContractCaller struct {
	evm    *EVM
	smcAbi abi.ABI
	sender common.Address
}

func (c *SmartContractCaller) validators() ([]common.Address, error) {
	res, err := c.staticCall(getValidatorsMethod, c.evm.ChainConfig().ConsortiumV2Contracts.RoninValidatorSet)
	if err != nil {
		return nil, err
	}
	return *abi.ConvertType(res[0], new([]common.Address)).(*[]common.Address), nil
}

func (c *SmartContractCaller) totalBalances(validators []common.Address) ([]*big.Int, error) {
	res, err := c.staticCall(totalBalancesMethod, c.evm.ChainConfig().ConsortiumV2Contracts.RoninValidatorSet, validators)
	if err != nil {
		return nil, err
	}
	return *abi.ConvertType(res[0], new([]*big.Int)).(*[]*big.Int), nil
}

func (c *SmartContractCaller) staticCall(method string, contract common.Address, args ...interface{}) ([]interface{}, error) {
	return staticCall(c.evm, c.smcAbi, method, contract, c.sender, args...)
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

func (c *consortiumVerifyHeaders) Init(caller ContractRef, evm *EVM) {
	c.caller = caller
	c.evm = evm
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
	isVenoki := c.evm.chainRules.IsVenoki
	contractIdx := VerifyHeaders
	if isVenoki {
		contractIdx = VerifyHeadersVenoki
	}

	// get method, args from abi and check if method is valid
	smcAbi, method, args, err := loadMethodAndArgs(contractIdx, input)
	if err != nil {
		return nil, err
	}
	if method.Name != verifyHeaders {
		return nil, errors.New("invalid method")
	}
	if len(args) != 3 {
		return nil, errors.New(fmt.Sprintf("invalid arguments, expected 2 got %d", len(args)))
	}
	consensusAddr, ok := args[0].(common.Address)
	if !ok {
		return nil, errors.New("invalid first argument type")
	}
	blockHeader1, err := c.unpackHeader(smcAbi, args[1].([]byte), isVenoki)
	if err != nil {
		return nil, err
	}
	blockHeader2, err := c.unpackHeader(smcAbi, args[2].([]byte), isVenoki)
	if err != nil {
		return nil, err
	}
	output := c.verify(consensusAddr, blockHeader1, blockHeader2)
	return smcAbi.Methods[verifyHeaders].Outputs.Pack(output)
}

func (c *consortiumVerifyHeaders) unpackHeader(abi abi.ABI, input []byte, isVenoki bool) (types.BlockHeader, error) {
	if isVenoki {
		var blochHeader types.BlockHeaderV2
		if err := c.unpack(abi, &blochHeader, input); err != nil {
			return nil, err
		}

		return &blochHeader, nil
	}

	var blockHeader types.BlockHeaderV1
	if err := c.unpack(abi, &blockHeader, input); err != nil {
		return nil, err
	}
	return &blockHeader, nil
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

func (c *consortiumVerifyHeaders) getSigner(blockHeader types.BlockHeader) (common.Address, error) {
	header := blockHeader.ToHeader()
	if header.Number == nil || header.Time > uint64(time.Now().Unix()) || len(header.Extra) < crypto.SignatureLength {
		return common.Address{}, errors.New("invalid header")
	}
	// Recover the public key and the Ethereum address
	signature := header.Extra[len(header.Extra)-crypto.SignatureLength:]
	signedHash := SealHash(blockHeader).Bytes()
	pubkey, err := crypto.Ecrecover(signedHash, signature)
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

// verify returns true if 2 blocks has the same signer (consensus address), same block number
// but with different seal hash
func (c *consortiumVerifyHeaders) verify(consensusAddr common.Address, header1, header2 types.BlockHeader) bool {
	var maxOffset *big.Int

	// c.evm s nil in benchmark, so we skip this check in benchmark
	if c.evm != nil && !c.evm.chainConfig.IsConsortiumV2(header1.GetNumber()) {
		return false
	}
	if header1.ToHeader().ParentHash.Hex() != header2.ToHeader().ParentHash.Hex() {
		return false
	}
	if len(header1.GetExtraData()) < crypto.SignatureLength || len(header2.GetExtraData()) < crypto.SignatureLength {
		return false
	}
	if bytes.Equal(SealHash(header1).Bytes(), SealHash(header2).Bytes()) {
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

	if c.test {
		maxOffset = big.NewInt(doubleSigningOffsetTest)
	} else {
		methodAbi := *unmarshalledABIs[GetDoubleSignSlashingConfig]
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
		if currentBlock.Cmp(header1.GetNumber()) > 0 && new(big.Int).Sub(currentBlock, header1.GetNumber()).Cmp(maxOffset) > 0 {
			return false
		}
	}

	return signer1.Hex() == signer2.Hex() &&
		signer2.Hex() == header2.GetBenificiary().Hex() &&
		bytes.Equal(consensusAddr.Bytes(), signer1.Bytes())
}

// SealHash returns the hash of a block prior to it being sealed.
func SealHash(header types.BlockHeader) (hash common.Hash) {
	hasher := sha3.NewLegacyKeccak256()
	encodeSigHeader(hasher, header.ToHeader(), header.GetChainId())
	hasher.Sum(hash[:0])
	return hash
}

func encodeSigHeader(w io.Writer, header *types.Header, chainId *big.Int) {
	enc := []interface{}{
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
	}
	if header.BaseFee != nil {
		enc = append(enc, header.BaseFee)
	}
	// blob fields are assumed to had been verified
	if header.BlobGasUsed != nil {
		enc = append(enc, header.BlobGasUsed)
		enc = append(enc, header.ExcessBlobGas)
	}
	if err := rlp.Encode(w, enc); err != nil {
		panic("can't encode: " + err.Error())
	}
}

type consortiumValidateFinalityProof struct {
	caller ContractRef
	evm    *EVM
}

func (c *consortiumValidateFinalityProof) Init(caller ContractRef, evm *EVM) {
	c.caller = caller
	c.evm = evm
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

func (c *consortiumValidateProofOfPossession) Init(caller ContractRef, evm *EVM) {
	c.caller = caller
	c.evm = evm
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
