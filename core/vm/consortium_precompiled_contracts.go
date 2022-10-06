package vm

import (
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/rlp"
	"golang.org/x/crypto/sha3"
	"io"
	"math/big"
	"math/rand"
	"os"
	"strings"
	"time"
)

var (
	consortiumLogAbi           = `[{"inputs":[{"internalType":"string","name":"message","type":"string"}],"name":"log","outputs":[],"stateMutability":"nonpayable","type":"function"}]`
	consortiumSortValidatorAbi = `[{"inputs":[],"name":"getValidatorCandidates","outputs":[{"internalType":"address[]","name":"","type":"address[]"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"uint256","name":"limit","type":"uint256"}],"name":"sortValidators","outputs":[{"internalType":"address[]","name":"validators","type":"address[]"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"address[]","name":"_poolList","type":"address[]"}],"name":"totalBalances","outputs":[{"internalType":"uint256[]","name":"_balances","type":"uint256[]"}],"stateMutability":"view","type":"function"}]`
	consortiumVerifyHeadersAbi = `[{"outputs":[],"name":"getHeader","inputs":[{"internalType":"uint256","name":"chainId","type":"uint256"},{"internalType":"bytes32","name":"parentHash","type":"bytes32"},{"internalType":"bytes32","name":"ommersHash","type":"bytes32"},{"internalType":"address","name":"coinbase","type":"address"},{"internalType":"bytes32","name":"stateRoot","type":"bytes32"},{"internalType":"bytes32","name":"transactionsRoot","type":"bytes32"},{"internalType":"bytes32","name":"receiptsRoot","type":"bytes32"},{"internalType":"uint8[256]","name":"logsBloom","type":"uint8[256]"},{"internalType":"uint256","name":"difficulty","type":"uint256"},{"internalType":"uint256","name":"number","type":"uint256"},{"internalType":"uint64","name":"gasLimit","type":"uint64"},{"internalType":"uint64","name":"gasUsed","type":"uint64"},{"internalType":"uint64","name":"timestamp","type":"uint64"},{"internalType":"bytes","name":"extraData","type":"bytes"},{"internalType":"bytes32","name":"mixHash","type":"bytes32"},{"internalType":"uint64","name":"nonce","type":"uint64"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"bytes","name":"header1","type":"bytes"},{"internalType":"bytes","name":"header2","type":"bytes"}],"name":"validatingDoubleSignProof","outputs":[{"internalType":"bool","name":"","type":"bool"}],"stateMutability":"view","type":"function"}]`
)

const (
	sortValidatorsMethod = "sortValidators"
	logMethod            = "log"
	getValidatorsMethod  = "getValidatorCandidates"
	totalBalancesMethod  = "totalBalances"
	verifyHeaders        = "validatingDoubleSignProof"
	getHeader            = "getHeader"
	extraVanity          = 32
)

func PrecompiledContractsConsortium(caller ContractRef, evm *EVM) map[common.Address]PrecompiledContract {
	return map[common.Address]PrecompiledContract{
		common.BytesToAddress([]byte{101}): &consortiumLog{},
		common.BytesToAddress([]byte{102}): &consortiumValidatorSorting{caller: caller, evm: evm},
		common.BytesToAddress([]byte{103}): &consortiumVerifyHeaders{caller: caller, evm: evm},
	}
}

type consortiumLog struct{}

func (c *consortiumLog) RequiredGas(input []byte) uint64 {
	return 0
}

func (c *consortiumLog) Run(input []byte) ([]byte, error) {
	if os.Getenv("DEBUG") != "true" {
		return input, nil
	}
	_, method, args, err := loadMethodAndArgs(consortiumLogAbi, input)
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

type consortiumValidatorSorting struct {
	caller ContractRef
	evm    *EVM
}

func (c *consortiumValidatorSorting) RequiredGas(input []byte) uint64 {
	return 0
}

func (c *consortiumValidatorSorting) Run(input []byte) ([]byte, error) {
	if c.evm.ChainConfig().ConsortiumV2Contracts == nil {
		return nil, errors.New("cannot find consortium v2 contracts")
	}
	if !c.evm.ChainConfig().ConsortiumV2Contracts.IsSystemContract(c.caller.Address()) {
		return nil, errors.New("unauthorized sender")
	}
	// get method, args from abi
	smcAbi, method, args, err := loadMethodAndArgs(consortiumSortValidatorAbi, input)
	if err != nil {
		return nil, err
	}
	if method.Name != sortValidatorsMethod {
		return nil, errors.New("invalid method")
	}
	if len(args) != 1 {
		return nil, errors.New(fmt.Sprintf("invalid arguments, expected 1 got %d", len(args)))
	}
	// cast args[0] to number
	limit, ok := args[0].(*big.Int)
	if !ok {
		return nil, errors.New("invalid argument type")
	}
	validators, err := loadValidators(c.evm, smcAbi, c.caller.Address())
	if err != nil {
		return nil, err
	}
	totalBalances, err := loadTotalBalances(c.evm, smcAbi, c.caller.Address(), validators)
	if err != nil {
		return nil, err
	}
	if len(validators) != len(totalBalances) {
		return nil, errors.New("balances and validators length mismatched")
	}
	sortValidators(validators, totalBalances)

	if limit.Int64() > int64(len(validators)) {
		limit = big.NewInt(int64(len(validators)))
	}

	return method.Outputs.Pack(validators[:limit.Int64()])
}

func sortValidators(validators []common.Address, balances []*big.Int) {
	if len(validators) < 2 {
		return
	}

	left, right := 0, len(validators)-1

	pivot := rand.Int() % len(validators)

	validators[pivot], validators[right] = validators[right], validators[pivot]
	balances[pivot], balances[right] = balances[right], balances[pivot]

	for i, _ := range validators {
		cmp := balances[i].Cmp(balances[right])
		addrsCmp := big.NewInt(0).SetBytes(validators[i].Bytes()).Cmp(big.NewInt(0).SetBytes(validators[right].Bytes())) > 0
		if cmp > 0 || (cmp == 0 && addrsCmp) {
			validators[left], validators[i] = validators[i], validators[left]
			balances[left], balances[i] = balances[i], balances[left]
			left++
		}
	}

	validators[left], validators[right] = validators[right], validators[left]
	balances[left], balances[right] = balances[right], balances[left]

	sortValidators(validators[:left], balances[:left])
	sortValidators(validators[left+1:], balances[left+1:])

	return
}

func loadValidators(evm *EVM, smcAbi abi.ABI, sender common.Address) ([]common.Address, error) {
	res, err := staticCall(evm, smcAbi, getValidatorsMethod, evm.ChainConfig().ConsortiumV2Contracts.RoninValidatorSet, sender)
	if err != nil {
		return nil, err
	}
	return *abi.ConvertType(res[0], new([]common.Address)).(*[]common.Address), nil
}

func loadTotalBalances(evm *EVM, smcAbi abi.ABI, sender common.Address, validators []common.Address) ([]*big.Int, error) {
	res, err := staticCall(evm, smcAbi, totalBalancesMethod, evm.ChainConfig().ConsortiumV2Contracts.StakingContract, sender, validators)
	if err != nil {
		return nil, err
	}
	return *abi.ConvertType(res[0], new([]*big.Int)).(*[]*big.Int), nil
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

func loadMethodAndArgs(smcAbi string, input []byte) (abi.ABI, *abi.Method, []interface{}, error) {
	var (
		pAbi   abi.ABI
		err    error
		method *abi.Method
		args   []interface{}
	)
	if pAbi, err = abi.JSON(strings.NewReader(smcAbi)); err != nil {
		return abi.ABI{}, nil, nil, err
	}
	if method, err = pAbi.MethodById(input); err != nil {
		return abi.ABI{}, nil, nil, err
	}
	if args, err = method.Inputs.Unpack(input[4:]); err != nil {
		return abi.ABI{}, nil, nil, err
	}
	return pAbi, method, args, nil
}

type BlockHeader struct {
	ChainId          *big.Int       `abi:"chainId"`
	ParentHash       [32]uint8      `abi:"parentHash"`
	OmmersHash       [32]uint8      `abi:"ommersHash"`
	Benificiary      common.Address `abi:"coinbase"`
	StateRoot        [32]uint8      `abi:"stateRoot"`
	TransactionsRoot [32]uint8      `abi:"transactionsRoot"`
	ReceiptsRoot     [32]uint8      `abi:"receiptsRoot"`
	LogsBloom        [256]uint8     `abi:"logsBloom"`
	Difficulty       *big.Int       `abi:"difficulty"`
	Number           *big.Int       `abi:"number"`
	GasLimit         uint64         `abi:"gasLimit"`
	GasUsed          uint64         `abi:"gasUsed"`
	Timestamp        uint64         `abi:"timestamp"`
	ExtraData        []byte         `abi:"extraData"`
	MixHash          [32]uint8      `abi:"mixHash"`
	Nonce            uint64         `abi:"nonce"`
}

func fromHeader(header *types.Header, chainId *big.Int) *BlockHeader {
	blockHeader := &BlockHeader{
		ChainId:     chainId,
		Difficulty:  header.Difficulty,
		Number:      header.Number,
		GasLimit:    header.GasLimit,
		GasUsed:     header.GasUsed,
		Timestamp:   header.Time,
		Nonce:       header.Nonce.Uint64(),
		ExtraData:   header.Extra,
		LogsBloom:   header.Bloom,
		Benificiary: header.Coinbase,
	}
	copy(blockHeader.ParentHash[:], header.ParentHash.Bytes())
	copy(blockHeader.StateRoot[:], header.Root.Bytes())
	copy(blockHeader.TransactionsRoot[:], header.TxHash.Bytes())
	copy(blockHeader.ReceiptsRoot[:], header.ReceiptHash.Bytes())
	copy(blockHeader.MixHash[:], header.MixDigest.Bytes())
	return blockHeader
}

func (b *BlockHeader) toHeader() *types.Header {
	return &types.Header{
		ParentHash:  common.BytesToHash(b.ParentHash[:]),
		Root:        common.BytesToHash(b.StateRoot[:]),
		TxHash:      common.BytesToHash(b.TransactionsRoot[:]),
		ReceiptHash: common.BytesToHash(b.ReceiptsRoot[:]),
		Bloom:       types.BytesToBloom(b.LogsBloom[:]),
		Difficulty:  b.Difficulty,
		Number:      b.Number,
		GasLimit:    b.GasLimit,
		GasUsed:     b.GasUsed,
		Time:        b.Timestamp,
		Extra:       b.ExtraData,
		MixDigest:   common.BytesToHash(b.MixHash[:]),
		Nonce:       types.EncodeNonce(b.Nonce),
		Coinbase:    b.Benificiary,
	}
}

func (b *BlockHeader) Bytes() ([]byte, error) {
	pAbi, _ := abi.JSON(strings.NewReader(consortiumVerifyHeadersAbi))
	bloom := types.BytesToBloom(b.LogsBloom[:])
	var uncles [32]uint8
	return pAbi.Methods[getHeader].Inputs.Pack(b.ChainId, b.ParentHash, uncles, b.Benificiary, b.StateRoot, b.TransactionsRoot, b.ReceiptsRoot, bloom.Bytes(), b.Difficulty, b.Number, b.GasLimit, b.GasUsed, b.Timestamp, b.ExtraData, b.MixHash, b.Nonce)
}

type consortiumVerifyHeaders struct {
	caller ContractRef
	evm    *EVM
}

func (c *consortiumVerifyHeaders) RequiredGas(input []byte) uint64 {
	return 0
}

func (c *consortiumVerifyHeaders) Run(input []byte) ([]byte, error) {
	if c.evm.ChainConfig().ConsortiumV2Contracts == nil {
		return nil, errors.New("cannot find consortium v2 contracts")
	}
	if !c.evm.ChainConfig().ConsortiumV2Contracts.IsSystemContract(c.caller.Address()) {
		return nil, errors.New("unauthorized sender")
	}
	// get method, args from abi
	smcAbi, method, args, err := loadMethodAndArgs(consortiumVerifyHeadersAbi, input)
	if err != nil {
		return nil, err
	}
	if method.Name != verifyHeaders {
		return nil, errors.New("invalid method")
	}
	if len(args) != 2 {
		return nil, errors.New(fmt.Sprintf("invalid arguments, expected 2 got %d", len(args)))
	}
	// decode header1, header2
	var blockHeader1, blockHeader2 BlockHeader
	if err := c.unpack(smcAbi, &blockHeader1, args[0].([]byte)); err != nil {
		return nil, err
	}
	if err := c.unpack(smcAbi, &blockHeader2, args[1].([]byte)); err != nil {
		return nil, err
	}
	output := c.verify(blockHeader1, blockHeader2)
	return smcAbi.Methods[verifyHeaders].Outputs.Pack(output)
}

func (c *consortiumVerifyHeaders) unpack(smcAbi abi.ABI, v interface{}, input []byte) error {
	args := smcAbi.Methods[getHeader].Inputs
	output, err := args.Unpack(input)
	if err != nil {
		return err
	}
	return args.Copy(v, output)
}

func (c *consortiumVerifyHeaders) getSigner(header BlockHeader) (common.Address, error) {
	if header.Number == nil || header.Timestamp > uint64(time.Now().Unix()) || len(header.ExtraData) < crypto.SignatureLength {
		return common.Address{}, errors.New("invalid header")
	}
	signature := header.ExtraData[len(header.ExtraData)-crypto.SignatureLength:]

	// Recover the public key and the Ethereum address
	pubkey, err := crypto.Ecrecover(SealHash(header.toHeader(), c.evm.ChainConfig().ChainID).Bytes(), signature)
	if err != nil {
		return common.Address{}, err
	}
	var signer common.Address
	copy(signer[:], crypto.Keccak256(pubkey[1:])[12:])

	return signer, nil
}

func (c *consortiumVerifyHeaders) verify(header1, header2 BlockHeader) bool {
	if header1.toHeader().ParentHash.Hex() != header2.toHeader().ParentHash.Hex() {
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
	return signer1.Hex() == signer2.Hex()
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
