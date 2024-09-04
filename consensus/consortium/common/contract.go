package common

import (
	"bytes"
	"context"
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/consensus"
	finalityTracking "github.com/ethereum/go-ethereum/consensus/consortium/generated_contracts/finality_tracking"
	legacyProfile "github.com/ethereum/go-ethereum/consensus/consortium/generated_contracts/legacy_profile"
	"github.com/ethereum/go-ethereum/consensus/consortium/generated_contracts/profile"
	roninValidatorSet "github.com/ethereum/go-ethereum/consensus/consortium/generated_contracts/ronin_validator_set"
	slashIndicator "github.com/ethereum/go-ethereum/consensus/consortium/generated_contracts/slash_indicator"
	"github.com/ethereum/go-ethereum/consensus/consortium/generated_contracts/staking"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/bls/blst"
	blsCommon "github.com/ethereum/go-ethereum/crypto/bls/common"
	"github.com/ethereum/go-ethereum/internal/ethapi"
	"github.com/ethereum/go-ethereum/log"
	chainParams "github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/rpc"
)

var errMethodUnimplemented = errors.New("method is unimplemented")

// getTransactionOpts is a helper function that creates TransactOpts with GasPrice equals 0
func getTransactionOpts(from common.Address, nonce uint64, signTxFn SignerTxFn, isVenoki bool) *bind.TransactOpts {
	var gasLimit uint64
	if isVenoki {
		gasLimit = systemTransactionGasLimit
	} else {
		gasLimit = math.MaxUint64 / 2
	}

	return &bind.TransactOpts{
		From:     from,
		GasLimit: gasLimit,
		GasPrice: big.NewInt(0),
		// Set dummy value always equal 0 since it will be overridden when creating a new message
		Value:  new(big.Int).SetUint64(0),
		Nonce:  new(big.Int).SetUint64(nonce),
		NoSend: true,
		Signer: func(address common.Address, tx *types.Transaction) (*types.Transaction, error) {
			// The transaction signing will happen later in ApplyTransaction
			return tx, nil
		},
	}
}

type ContractInteraction interface {
	WrapUpEpoch(opts *ApplyTransactOpts) error
	SubmitBlockReward(opts *ApplyTransactOpts) error
	Slash(opts *ApplyTransactOpts, spoiledValidator common.Address) error
	FinalityReward(opts *ApplyTransactOpts, votedValidators []common.Address) error
	GetBlockProducers(blockHash common.Hash, blockNumber *big.Int) ([]common.Address, error)
	GetValidatorCandidates(blockHash common.Hash, blockNumber *big.Int) ([]common.Address, error)
	GetBlsPublicKey(blockHash common.Hash, blockNumber *big.Int, validator common.Address) (blsCommon.PublicKey, error)
	GetStakedAmount(blockHash common.Hash, blockNumber *big.Int, validators []common.Address) ([]*big.Int, error)
	GetMaxValidatorNumber(blockHash common.Hash, blockNumber *big.Int) (*big.Int, error)
}

// ContractIntegrator is a contract facing to interact with smart contract that supports DPoS
type ContractIntegrator struct {
	chainConfig         *chainParams.ChainConfig
	signer              types.Signer
	roninValidatorSetSC *roninValidatorSet.RoninValidatorSet
	slashIndicatorSC    *slashIndicator.SlashIndicator
	finalityTrackingSC  *finalityTracking.FinalityTracking

	roninValidatorSetABI *abi.ABI
	legacyProfileABI     *abi.ABI
	profileABI           *abi.ABI
	stakingABI           *abi.ABI

	ethAPI *ethapi.PublicBlockChainAPI

	signTxFn SignerTxFn
	coinbase common.Address

	// This is used in unit test only
	contractCallHook func(method string) []byte
}

// NewContractIntegrator creates new ContractIntegrator with custom backend and signTxFn
func NewContractIntegrator(
	config *chainParams.ChainConfig,
	backend bind.ContractBackend,
	signTxFn SignerTxFn,
	coinbase common.Address,
	ethAPI *ethapi.PublicBlockChainAPI,
) (*ContractIntegrator, error) {
	// Create Ronin Validator Set smart contract
	roninValidatorSetSC, err := roninValidatorSet.NewRoninValidatorSet(config.ConsortiumV2Contracts.RoninValidatorSet, backend)
	if err != nil {
		return nil, err
	}
	// Create Slash Indicator smart contract
	slashIndicatorSC, err := slashIndicator.NewSlashIndicator(config.ConsortiumV2Contracts.SlashIndicator, backend)
	if err != nil {
		return nil, err
	}
	// Create Finality Tracking contract instance
	finalityTrackingSC, err := finalityTracking.NewFinalityTracking(config.ConsortiumV2Contracts.FinalityTracking, backend)
	if err != nil {
		return nil, err
	}

	roninValidatorSetABI, err := roninValidatorSet.RoninValidatorSetMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	stakingABI, err := staking.StakingMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	profileABI, err := profile.ProfileMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	legacyProfileABI, err := legacyProfile.ProfileMetaData.GetAbi()
	if err != nil {
		return nil, err
	}

	return &ContractIntegrator{
		chainConfig:         config,
		roninValidatorSetSC: roninValidatorSetSC,
		slashIndicatorSC:    slashIndicatorSC,
		finalityTrackingSC:  finalityTrackingSC,

		roninValidatorSetABI: roninValidatorSetABI,
		profileABI:           profileABI,
		legacyProfileABI:     legacyProfileABI,
		stakingABI:           stakingABI,

		ethAPI: ethAPI,

		signTxFn: signTxFn,
		signer:   types.LatestSignerForChainID(config.ChainID),
		coinbase: coinbase,
	}, nil
}

// GetBlockProducers retrieves block producer addresses
func (c *ContractIntegrator) GetBlockProducers(blockHash common.Hash, blockNumber *big.Int) ([]common.Address, error) {
	blockNr := rpc.BlockNumberOrHashWithNumber(rpc.BlockNumber(blockNumber.Int64()))
	if c.chainConfig.IsTripp(blockNumber) {
		blockNr = rpc.BlockNumberOrHashWithHash(blockHash, false)
	}

	var addresses []common.Address
	err := c.contractCall(c.roninValidatorSetABI, c.chainConfig.ConsortiumV2Contracts.RoninValidatorSet,
		"getBlockProducers", blockNr, &addresses)
	if err != nil {
		return nil, err
	}
	return addresses, nil
}

func (c *ContractIntegrator) GetValidatorCandidates(blockHash common.Hash, blockNumber *big.Int) ([]common.Address, error) {
	blockNr := rpc.BlockNumberOrHashWithNumber(rpc.BlockNumber(blockNumber.Int64()))
	if c.chainConfig.IsTripp(blockNumber) {
		blockNr = rpc.BlockNumberOrHashWithHash(blockHash, false)
	}

	var addresses []common.Address
	err := c.contractCall(c.roninValidatorSetABI, c.chainConfig.ConsortiumV2Contracts.RoninValidatorSet,
		"getValidatorCandidates", blockNr, &addresses)
	if err != nil {
		return nil, err
	}
	return addresses, nil
}

// WrapUpEpoch distributes rewards to validators and updates validators set
func (c *ContractIntegrator) WrapUpEpoch(opts *ApplyTransactOpts) error {
	nonce := opts.State.GetNonce(c.coinbase)
	isVenoki := c.chainConfig.IsVenoki(opts.Header.Number)
	tx, err := c.roninValidatorSetSC.WrapUpEpoch(getTransactionOpts(c.coinbase, nonce, c.signTxFn, isVenoki))
	if err != nil {
		return err
	}
	msg := types.NewMessage(
		opts.Header.Coinbase,
		tx.To(),
		opts.State.GetNonce(opts.Header.Coinbase),
		tx.Value(),
		tx.Gas(),
		big.NewInt(0),
		big.NewInt(0),
		big.NewInt(0),
		tx.Data(),
		tx.AccessList(),
		false,
		nil,
		nil,
	)

	if err = ApplyTransaction(msg, opts); err != nil {
		return err
	}

	log.Info("Wrapped up epoch", "block", opts.Header.Number)
	return err
}

// SubmitBlockReward submits a transaction to the ValidatorSetSC that gets balance from sender
// plus bonus from vesting contract in order to updates mining reward and delegating reward
func (c *ContractIntegrator) SubmitBlockReward(opts *ApplyTransactOpts) error {
	coinbase := opts.Header.Coinbase
	balance := opts.State.GetBalance(consensus.SystemAddress)
	opts.State.SetBalance(consensus.SystemAddress, big.NewInt(0))
	opts.State.AddBalance(coinbase, balance)

	nonce := opts.State.GetNonce(c.coinbase)
	isVenoki := c.chainConfig.IsVenoki(opts.Header.Number)
	tx, err := c.roninValidatorSetSC.SubmitBlockReward(getTransactionOpts(c.coinbase, nonce, c.signTxFn, isVenoki))
	if err != nil {
		return err
	}
	log.Debug("Submitted block reward", "block", opts.Header.Number, "amount", balance.Uint64())

	msg := types.NewMessage(
		opts.Header.Coinbase,
		tx.To(),
		opts.State.GetNonce(opts.Header.Coinbase),
		// Reassign value with the current balance. It will be overridden the current one.
		balance,
		tx.Gas(),
		big.NewInt(0),
		big.NewInt(0),
		big.NewInt(0),
		tx.Data(),
		tx.AccessList(),
		false,
		nil,
		nil,
	)

	if err = ApplyTransaction(msg, opts); err != nil {
		return err
	}

	return nil
}

// Slash submits a transaction to the SlashIndicatorSC that checks the unavailability of the coinbase
// and calls the slash method corresponding
func (c *ContractIntegrator) Slash(opts *ApplyTransactOpts, spoiledValidator common.Address) error {
	nonce := opts.State.GetNonce(c.coinbase)
	isVenoki := c.chainConfig.IsVenoki(opts.Header.Number)
	tx, err := c.slashIndicatorSC.SlashUnavailability(getTransactionOpts(c.coinbase, nonce, c.signTxFn, isVenoki), spoiledValidator)
	if err != nil {
		return err
	}

	msg := types.NewMessage(
		opts.Header.Coinbase,
		tx.To(),
		opts.State.GetNonce(opts.Header.Coinbase),
		tx.Value(),
		tx.Gas(),
		big.NewInt(0),
		big.NewInt(0),
		big.NewInt(0),
		tx.Data(),
		tx.AccessList(),
		false,
		nil,
		nil,
	)

	if err = ApplyTransaction(msg, opts); err != nil {
		return err
	}

	return nil
}

func (c *ContractIntegrator) FinalityReward(opts *ApplyTransactOpts, votedValidators []common.Address) error {
	nonce := opts.State.GetNonce(c.coinbase)
	isVenoki := c.chainConfig.IsVenoki(opts.Header.Number)
	tx, err := c.finalityTrackingSC.RecordFinality(getTransactionOpts(c.coinbase, nonce, c.signTxFn, isVenoki), votedValidators)
	if err != nil {
		return err
	}

	msg := types.NewMessage(
		opts.Header.Coinbase,
		tx.To(),
		opts.State.GetNonce(opts.Header.Coinbase),
		tx.Value(),
		tx.Gas(),
		big.NewInt(0),
		big.NewInt(0),
		big.NewInt(0),
		tx.Data(),
		tx.AccessList(),
		false,
		nil,
		nil,
	)

	if err = ApplyTransaction(msg, opts); err != nil {
		return err
	}

	return nil
}

func (c *ContractIntegrator) GetBlsPublicKey(blockHash common.Hash, blockNumber *big.Int, validator common.Address) (blsCommon.PublicKey, error) {
	if c.chainConfig.IsTripp(blockNumber) {
		return c.getBlsPublicKey(blockHash, blockNumber, validator)
	} else {
		return c.getBlsPublicKeyLegacy(blockHash, blockNumber, validator)
	}
}

func (c *ContractIntegrator) getBlsPublicKeyLegacy(blockHash common.Hash, blockNumber *big.Int, validator common.Address) (blsCommon.PublicKey, error) {
	blockNr := rpc.BlockNumberOrHashWithNumber(rpc.BlockNumber(blockNumber.Int64()))
	if c.chainConfig.IsTripp(blockNumber) {
		blockNr = rpc.BlockNumberOrHashWithHash(blockHash, false)
	}

	var validatorProfile legacyProfile.IProfileCandidateProfile
	err := c.contractCall(c.legacyProfileABI, c.chainConfig.ConsortiumV2Contracts.ProfileContract,
		"getId2Profile", blockNr, &validatorProfile, validator)
	if err != nil {
		return nil, err
	}

	blsPublicKey, err := blst.PublicKeyFromBytes(validatorProfile.Pubkey)
	if err != nil {
		return nil, err
	}
	return blsPublicKey, nil
}

func (c *ContractIntegrator) getBlsPublicKey(blockHash common.Hash, blockNumber *big.Int, validator common.Address) (blsCommon.PublicKey, error) {
	blockNr := rpc.BlockNumberOrHashWithNumber(rpc.BlockNumber(blockNumber.Int64()))
	if c.chainConfig.IsTripp(blockNumber) {
		blockNr = rpc.BlockNumberOrHashWithHash(blockHash, false)
	}

	var validatorId common.Address
	err := c.contractCall(c.profileABI, c.chainConfig.ConsortiumV2Contracts.ProfileContract,
		"getConsensus2Id", blockNr, &validatorId, validator)
	if err != nil {
		return nil, err
	}

	var rawPublicKey []byte
	err = c.contractCall(c.profileABI, c.chainConfig.ConsortiumV2Contracts.ProfileContract,
		"getId2Pubkey", blockNr, &rawPublicKey, validatorId)
	if err != nil {
		return nil, err
	}

	blsPublicKey, err := blst.PublicKeyFromBytes(rawPublicKey)
	if err != nil {
		return nil, err
	}
	return blsPublicKey, nil
}

func (c *ContractIntegrator) GetStakedAmount(blockHash common.Hash, blockNumber *big.Int, validators []common.Address) ([]*big.Int, error) {
	blockNr := rpc.BlockNumberOrHashWithNumber(rpc.BlockNumber(blockNumber.Int64()))
	if c.chainConfig.IsTripp(blockNumber) {
		blockNr = rpc.BlockNumberOrHashWithHash(blockHash, false)
	}

	var stakedAmounts []*big.Int
	err := c.contractCall(c.stakingABI, c.chainConfig.ConsortiumV2Contracts.StakingContract,
		"getManyStakingTotals", blockNr, &stakedAmounts, validators)
	if err != nil {
		return nil, err
	}
	return stakedAmounts, nil
}

func (c *ContractIntegrator) GetMaxValidatorNumber(blockHash common.Hash, blockNumber *big.Int) (*big.Int, error) {
	blockNr := rpc.BlockNumberOrHashWithNumber(rpc.BlockNumber(blockNumber.Int64()))
	if c.chainConfig.IsTripp(blockNumber) {
		blockNr = rpc.BlockNumberOrHashWithHash(blockHash, false)
	}

	var maxValidatorNumber *big.Int
	err := c.contractCall(c.roninValidatorSetABI, c.chainConfig.ConsortiumV2Contracts.RoninValidatorSet,
		"maxValidatorNumber", blockNr, &maxValidatorNumber)
	if err != nil {
		return nil, err
	}

	return maxValidatorNumber, nil
}

// Note: this function only supports 1 output
func (c *ContractIntegrator) contractCall(
	abi *abi.ABI,
	address common.Address,
	method string,
	blockNrOrHash rpc.BlockNumberOrHash,
	output interface{},
	input ...interface{},
) error {
	data, err := abi.Pack(method, input...)
	if err != nil {
		log.Error("Failed to pack tx's data", "error", err)
		return err
	}

	var result []byte
	if c.contractCallHook != nil {
		result = c.contractCallHook(method)
	} else {
		// do smart contract call
		msgData := (hexutil.Bytes)(data)
		to := address
		gas := (hexutil.Uint64)(systemTransactionGasLimit)
		result, err = c.ethAPI.Call(context.Background(), ethapi.TransactionArgs{
			Gas:  &gas,
			To:   &to,
			Data: &msgData,
		}, blockNrOrHash, nil, nil)
		if err != nil {
			return err
		}
	}

	return abi.UnpackIntoInterface(&output, method, result)
}

// ApplyMessageOpts is the collection of options to fine tune a contract call request.
type ApplyMessageOpts struct {
	State       *state.StateDB
	Header      *types.Header
	ChainConfig *chainParams.ChainConfig
	EVMContext  *vm.BlockContext
}

// ApplyTransactOpts is the collection of authorization data required to create a
// valid transaction.
type ApplyTransactOpts struct {
	*ApplyMessageOpts
	Txs         *[]*types.Transaction
	Receipts    *[]*types.Receipt
	ReceivedTxs *[]*types.Transaction
	UsedGas     *uint64
	Mining      bool
	Signer      types.Signer
	SignTxFn    SignerTxFn
	EthAPI      *ethapi.PublicBlockChainAPI
}

// ApplyTransaction attempts to apply a transaction to the given state database
// and uses the input parameters for its environment. It returns nil if applied success
// and an error if the transaction failed, indicating the block was invalid.
func ApplyTransaction(msg types.Message, opts *ApplyTransactOpts) (err error) {
	var failed bool

	signer := opts.Signer
	signTxFn := opts.SignTxFn
	coinbase := opts.Header.Coinbase
	mining := opts.Mining
	chainConfig := opts.ChainConfig
	receivedTxs := opts.ReceivedTxs
	txs := opts.Txs
	header := opts.Header
	receipts := opts.Receipts
	usedGas := opts.UsedGas
	nonce := msg.Nonce()

	// TODO(linh): This function is deprecated. Shall we replace it with NewTx?
	expectedTx := types.NewTransaction(nonce, *msg.To(), msg.Value(), msg.Gas(), msg.GasPrice(), msg.Data())
	expectedHash := signer.Hash(expectedTx)

	sender := msg.From()
	// An empty/non-existing account's code hash is 0x000...00, while an existing account with no code has code hash
	// that is equal to crypto.Keccak256Hash(nil)
	if codeHash := opts.State.GetCodeHash(sender); codeHash != crypto.Keccak256Hash(nil) && codeHash != (common.Hash{}) {
		return fmt.Errorf("%w: address %v, codehash: %s", core.ErrSenderNoEOA, sender.Hex(), codeHash)
	}

	if sender != coinbase {
		return fmt.Errorf("sender of system transaction is not coinbase, sender: %s, coinbase: %s", sender, coinbase)
	}

	if mining {
		expectedTx, err = signTxFn(accounts.Account{Address: msg.From()}, expectedTx, chainConfig.ChainID)
		if err != nil {
			return err
		}
	} else {
		if receivedTxs == nil || len(*receivedTxs) == 0 || (*receivedTxs)[0] == nil {
			return errors.New("supposed to get a actual transaction, but get none")
		}
		// receivedTxs (a.k.a systemTxs) is collected by the method Process of state_processor
		// The system transaction is the transaction that have the receiver is ConsortiumV2Contracts
		actualTx := (*receivedTxs)[0]
		if !bytes.Equal(signer.Hash(actualTx).Bytes(), expectedHash.Bytes()) {
			return fmt.Errorf(
				"expected transaction: nonce %d, to %s, value %s, gas %d, gasPrice %s, data %s\n"+
					"got transaction: nonce %d, to %s, value %s, gas %d, gasPrice %s, data %s, hash %s",
				expectedTx.Nonce(), expectedTx.To().String(), expectedTx.Value().String(), expectedTx.Gas(),
				expectedTx.GasPrice().String(), hex.EncodeToString(expectedTx.Data()),
				actualTx.Nonce(), actualTx.To().String(), actualTx.Value().String(), actualTx.Gas(),
				actualTx.GasPrice().String(), hex.EncodeToString(actualTx.Data()), actualTx.Hash(),
			)
		}
		expectedTx = actualTx
		// move to next
		*receivedTxs = (*receivedTxs)[1:]
	}
	opts.State.SetTxContext(expectedTx.Hash(), len(*txs))
	opts.State.SetNonce(msg.From(), nonce+1)
	gasUsed, err := applyMessage(opts.ApplyMessageOpts, expectedTx)
	if err != nil {
		failed = true
	} else {
		failed = false
	}
	log.Debug("Applied transaction", "gasUsed", gasUsed)

	*txs = append(*txs, expectedTx)
	var root []byte
	if chainConfig.IsByzantium(opts.Header.Number) {
		opts.State.Finalise(true)
	} else {
		root = opts.State.IntermediateRoot(chainConfig.IsEIP158(opts.Header.Number)).Bytes()
	}
	*usedGas += gasUsed

	// TODO(linh): This function is deprecated. Shall we replace it with Receipt struct?
	receipt := types.NewReceipt(root, failed, *usedGas)
	receipt.TxHash = expectedTx.Hash()
	receipt.GasUsed = gasUsed

	// Set the receipt logs and create a bloom for filtering
	receipt.Logs = opts.State.GetLogs(expectedTx.Hash(), header.Hash())
	receipt.Bloom = types.CreateBloom(types.Receipts{receipt})
	receipt.BlockHash = header.Hash()
	receipt.BlockNumber = header.Number
	receipt.TransactionIndex = uint(opts.State.TxIndex())
	*receipts = append(*receipts, receipt)
	return nil
}

// applyMessage creates new evm and applies a transaction to the current state
func applyMessage(
	opts *ApplyMessageOpts,
	tx *types.Transaction,
) (uint64, error) {
	// Create a new context to be used in the EVM environment
	opts.EVMContext.CurrentTransaction = tx
	from, _ := types.Sender(types.MakeSigner(opts.ChainConfig, opts.Header.Number), tx)

	chainRules := opts.ChainConfig.Rules(opts.Header.Number)
	if chainRules.IsShanghai {
		opts.State.Prepare(chainRules, from, from, tx.To(), vm.ActivePrecompiles(chainRules), nil)
	} else if chainRules.IsBerlin {
		opts.State.ResetAccessList()
	}

	// Create a new environment which holds all relevant information
	// about the transaction and calling mechanisms.
	vmenv := vm.NewEVM(*opts.EVMContext, vm.TxContext{Origin: from, GasPrice: big.NewInt(0)}, opts.State, opts.ChainConfig, vm.Config{})
	// Apply the transaction to the current State (included in the env)
	ret, returnGas, err := vmenv.Call(
		vm.AccountRef(from),
		*tx.To(),
		tx.Data(),
		tx.Gas(),
		tx.Value(),
	)
	if err != nil {
		log.Error("Apply message failed", "message", string(ret), "error", err, "to", tx.To(), "value", tx.Value())
	}
	return tx.Gas() - returnGas, err
}

// ConsortiumBackend is a custom backend that supports call smart contract by using *ethapi.PublicBlockChainAPI
type ConsortiumBackend struct {
	*ethapi.PublicBlockChainAPI
}

// NewConsortiumBackend creates new ConsortiumBackend from *ethapi.PublicBlockChainAPI
func NewConsortiumBackend(ee *ethapi.PublicBlockChainAPI) *ConsortiumBackend {
	return &ConsortiumBackend{
		ee,
	}
}

// CodeAt returns the code of the given account. This is needed to differentiate
// between contract internal errors and the local chain being out of sync.
func (b *ConsortiumBackend) CodeAt(ctx context.Context, contract common.Address, blockNumber *big.Int) ([]byte, error) {
	blkNumber := rpc.LatestBlockNumber
	if blockNumber != nil {
		blkNumber = rpc.BlockNumber(blockNumber.Int64())
	}
	block := rpc.BlockNumberOrHashWithNumber(blkNumber)
	result, err := b.GetCode(ctx, contract, block)
	if err != nil {
		return nil, err
	}

	return result.MarshalText()
}

// CallContract executes an Ethereum contract call with the specified data as the
// input.
func (b *ConsortiumBackend) CallContract(ctx context.Context, call ethereum.CallMsg, blockNumber *big.Int) ([]byte, error) {
	blkNumber := rpc.LatestBlockNumber
	if blockNumber != nil {
		blkNumber = rpc.BlockNumber(blockNumber.Int64())
	}
	block := rpc.BlockNumberOrHashWithNumber(blkNumber)
	gas := (hexutil.Uint64)(systemTransactionGasLimit)
	data := (hexutil.Bytes)(call.Data)

	result, err := b.Call(ctx, ethapi.TransactionArgs{
		Gas:  &gas,
		To:   call.To,
		Data: &data,
	}, block, nil, nil)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// HeaderByNumber returns a block header from the current canonical chain. If
// number is nil, the latest known header is returned.
func (b *ConsortiumBackend) HeaderByNumber(ctx context.Context, number *big.Int) (*types.Header, error) {
	return b.GetHeader(ctx, rpc.BlockNumber(number.Int64()))
}

// PendingCodeAt returns the code of the given account in the pending state.
// NOTE(linh): This method is never called, implement for interface purposes
func (b *ConsortiumBackend) PendingCodeAt(ctx context.Context, account common.Address) ([]byte, error) {
	return nil, errMethodUnimplemented
}

// PendingNonceAt retrieves the current pending nonce associated with an account.
// NOTE(linh): This method is never called, implement for interface purposes
func (b *ConsortiumBackend) PendingNonceAt(ctx context.Context, account common.Address) (uint64, error) {
	return 0, errMethodUnimplemented
}

// SuggestGasPrice retrieves the currently suggested gas price to allow a timely
// execution of a transaction.
// NOTE(linh): This method is never called, implement for interface purposes
func (b *ConsortiumBackend) SuggestGasPrice(ctx context.Context) (*big.Int, error) {
	return big.NewInt(0), nil
}

// SuggestGasTipCap retrieves the currently suggested 1559 priority fee to allow
// a timely execution of a transaction.
// NOTE(linh): This method is never called, implement for interface purposes
func (b *ConsortiumBackend) SuggestGasTipCap(ctx context.Context) (*big.Int, error) {
	return big.NewInt(0), nil
}

// EstimateGas tries to estimate the gas needed to execute a specific
// transaction based on the current pending state of the backend blockchain.
func (b *ConsortiumBackend) EstimateGas(ctx context.Context, call ethereum.CallMsg) (gas uint64, err error) {
	return systemTransactionGasLimit, nil
}

// SendTransaction injects the transaction into the pending pool for execution.
// NOTE(linh): This method is never called, implement for interface purposes. We call ApplyTransaction directly
func (b *ConsortiumBackend) SendTransaction(ctx context.Context, tx *types.Transaction) error {
	// No need to send transaction
	return errMethodUnimplemented
}

// FilterLogs executes a log filter operation, blocking during execution and
// returning all the results in one batch.
// NOTE(linh): This method is never called, implement for interface purposes
func (b *ConsortiumBackend) FilterLogs(ctx context.Context, query ethereum.FilterQuery) ([]types.Log, error) {
	return nil, errMethodUnimplemented
}

// SubscribeFilterLogs creates a background log filtering operation, returning
// a subscription immediately, which can be used to stream the found events.
// NOTE(linh): This method is never called, implement for interface purposes
func (b *ConsortiumBackend) SubscribeFilterLogs(ctx context.Context, query ethereum.FilterQuery, ch chan<- types.Log) (ethereum.Subscription, error) {
	return nil, errMethodUnimplemented
}
