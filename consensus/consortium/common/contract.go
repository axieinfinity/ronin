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
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/consensus"
	roninValidatorSet "github.com/ethereum/go-ethereum/consensus/consortium/generated_contracts/ronin_validator_set"
	slashIndicator "github.com/ethereum/go-ethereum/consensus/consortium/generated_contracts/slash_indicator"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/internal/ethapi"
	"github.com/ethereum/go-ethereum/log"
	chainParams "github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/rpc"
)

var errMethodUnimplemented = errors.New("method is unimplemented")

// getTransactionOpts is a helper function that creates TransactOpts with GasPrice equals 0
func getTransactionOpts(from common.Address, nonce uint64, chainId *big.Int, signTxFn SignerTxFn) *bind.TransactOpts {
	return &bind.TransactOpts{
		From: from,
		// FIXME(linh): Decrease gasLimit later, math.MaxUint64 / 2 is too large
		GasLimit: uint64(math.MaxUint64 / 2),
		GasPrice: big.NewInt(0),
		// Set dummy value always equal 0 since it will be overridden when creating a new message
		Value:  new(big.Int).SetUint64(0),
		Nonce:  new(big.Int).SetUint64(nonce),
		NoSend: true,
		Signer: func(address common.Address, tx *types.Transaction) (*types.Transaction, error) {
			// signTxFn is nil when mining is not enabled, then we just return the transaction directly
			if signTxFn == nil {
				return tx, nil
			}

			return signTxFn(accounts.Account{Address: from}, tx, chainId)
		},
	}
}

// ContractIntegrator is a contract facing to interact with smart contract that supports DPoS
type ContractIntegrator struct {
	chainId             *big.Int
	signer              types.Signer
	roninValidatorSetSC *roninValidatorSet.RoninValidatorSet
	slashIndicatorSC    *slashIndicator.SlashIndicator
	signTxFn            SignerTxFn
	coinbase            common.Address
}

// NewContractIntegrator creates new ContractIntegrator with custom backend and signTxFn
func NewContractIntegrator(config *chainParams.ChainConfig, backend bind.ContractBackend, signTxFn SignerTxFn, coinbase common.Address) (*ContractIntegrator, error) {
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

	return &ContractIntegrator{
		chainId:             config.ChainID,
		roninValidatorSetSC: roninValidatorSetSC,
		slashIndicatorSC:    slashIndicatorSC,
		signTxFn:            signTxFn,
		signer:              types.LatestSignerForChainID(config.ChainID),
		coinbase:            coinbase,
	}, nil
}

// GetValidators retrieves top validators addresses
func (c *ContractIntegrator) GetValidators(blockNumber *big.Int) ([]common.Address, error) {
	callOpts := bind.CallOpts{
		BlockNumber: blockNumber,
	}
	addresses, err := c.roninValidatorSetSC.GetBlockProducers(&callOpts)
	if err != nil {
		return nil, err
	}
	return addresses, nil
}

// WrapUpEpoch distributes rewards to validators and updates validators set
func (c *ContractIntegrator) WrapUpEpoch(opts *ApplyTransactOpts) error {
	nonce := opts.State.GetNonce(c.coinbase)
	tx, err := c.roninValidatorSetSC.WrapUpEpoch(getTransactionOpts(c.coinbase, nonce, c.chainId, c.signTxFn))
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
	)

	if err = ApplyTransaction(msg, opts); err != nil {
		return err
	}

	log.Info("Wrapped up epoch", "block hash", opts.Header.Hash(), "tx hash", tx.Hash().Hex())
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
	tx, err := c.roninValidatorSetSC.SubmitBlockReward(getTransactionOpts(c.coinbase, nonce, c.chainId, c.signTxFn))
	if err != nil {
		return err
	}
	log.Info("Submitted block reward", "block hash", opts.Header.Hash(), "tx hash", tx.Hash().Hex(), "amount", balance.Uint64())

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
	tx, err := c.slashIndicatorSC.SlashUnavailability(getTransactionOpts(c.coinbase, nonce, c.chainId, c.signTxFn), spoiledValidator)
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
	)

	if err = ApplyTransaction(msg, opts); err != nil {
		return err
	}

	return nil
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
	miner := opts.Header.Coinbase
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

	if msg.From() == miner && mining {
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
			return fmt.Errorf("expected tx hash %v, get %v, nonce %d, to %s, value %s, gas %d, gasPrice %s, data %s", expectedHash.String(), actualTx.Hash().String(),
				expectedTx.Nonce(),
				expectedTx.To().String(),
				expectedTx.Value().String(),
				expectedTx.Gas(),
				expectedTx.GasPrice().String(),
				hex.EncodeToString(expectedTx.Data()),
			)
		}
		expectedTx = actualTx
		// move to next
		*receivedTxs = (*receivedTxs)[1:]
	}
	opts.State.Prepare(expectedTx.Hash(), len(*txs))
	gasUsed, err := applyMessage(msg, opts.ApplyMessageOpts)
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
	opts.State.SetNonce(msg.From(), nonce+1)
	return nil
}

// applyMessage creates new evm and applies a transaction to the current state
func applyMessage(
	msg types.Message,
	opts *ApplyMessageOpts,
) (uint64, error) {
	// Create a new context to be used in the EVM environment
	opts.EVMContext.CurrentTransaction = types.NewTransaction(msg.Nonce(), *msg.To(), msg.Value(), msg.Gas(), msg.GasPrice(), msg.Data())
	// Create a new environment which holds all relevant information
	// about the transaction and calling mechanisms.
	vmenv := vm.NewEVM(*opts.EVMContext, vm.TxContext{Origin: msg.From(), GasPrice: big.NewInt(0)}, opts.State, opts.ChainConfig, vm.Config{})
	// Apply the transaction to the current State (included in the env)
	ret, returnGas, err := vmenv.Call(
		vm.AccountRef(msg.From()),
		*msg.To(),
		msg.Data(),
		msg.Gas(),
		msg.Value(),
	)
	if err != nil {
		log.Error("Apply message failed", "message", string(ret), "error", err, "to", msg.To(), "value", msg.Value())
	}
	return msg.Gas() - returnGas, err
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
	gas := (hexutil.Uint64)(uint64(math.MaxUint64 / 2))
	data := (hexutil.Bytes)(call.Data)

	result, err := b.Call(ctx, ethapi.TransactionArgs{
		Gas:  &gas,
		To:   call.To,
		Data: &data,
	}, block, nil)
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
// NOTE(linh): We allow math.MaxUint64 / 2 because it is only called by the validator
func (b *ConsortiumBackend) EstimateGas(ctx context.Context, call ethereum.CallMsg) (gas uint64, err error) {
	return math.MaxUint64 / 2, nil
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
