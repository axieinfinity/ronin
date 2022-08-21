package common

import (
	"bytes"
	"context"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/consensus"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/systemcontracts/generated_contracts/validators"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/internal/ethapi"
	"github.com/ethereum/go-ethereum/log"
	chainParams "github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/rpc"
	"math/big"
)

var errMethodUnimplemented = errors.New("method is unimplemented")

type ContractIntegrator struct {
	signer      types.Signer
	validatorSC *validators.Validators
}

func NewContractIntegrator(config *chainParams.ChainConfig, backend bind.ContractBackend) (*ContractIntegrator, error) {
	validatorSC, err := validators.NewValidators(config.ConsortiumV2Contracts.ValidatorSC, backend)
	if err != nil {
		return nil, err
	}

	return &ContractIntegrator{
		validatorSC: validatorSC,
	}, nil
}

func (c *ContractIntegrator) GetValidators(header *types.Header) ([]common.Address, error) {
	addresses, err := c.validatorSC.GetValidators(&bind.CallOpts{
		BlockNumber: new(big.Int).Sub(header.Number, common.Big1),
	})
	if err != nil {
		return nil, err
	}

	return addresses, nil
}

func (c *ContractIntegrator) UpdateValidators(header *types.Header, opts *ApplyTransactOpts) error {
	coinbase := opts.Header.Coinbase
	nonce := opts.State.GetNonce(coinbase)

	tx, err := c.validatorSC.UpdateValidators(&bind.TransactOpts{
		From:     coinbase,
		GasLimit: math.MaxUint64 / 2,
		GasPrice: big.NewInt(0),
		Value:    new(big.Int).SetUint64(0),
		Nonce:    new(big.Int).SetUint64(nonce),
		NoSend:   true,
	})
	if err != nil {
		return err
	}

	msg, err := tx.AsMessage(c.signer, big.NewInt(0))
	if err != nil {
		return err
	}

	err = applyTransaction(msg, opts)
	if err != nil {
		return err
	}

	return err
}

func (c *ContractIntegrator) DistributeRewards(to common.Address, opts *ApplyTransactOpts) error {
	coinbase := opts.Header.Coinbase
	balance := opts.State.GetBalance(consensus.SystemAddress)
	if balance.Cmp(common.Big0) <= 0 {
		return nil
	}
	opts.State.SetBalance(consensus.SystemAddress, big.NewInt(0))
	opts.State.AddBalance(coinbase, balance)

	log.Trace("distribute to validator contract", "block hash", opts.Header.Hash(), "amount", balance)
	nonce := opts.State.GetNonce(coinbase)
	tx, err := c.validatorSC.DepositReward(&bind.TransactOpts{
		From:     coinbase,
		GasLimit: math.MaxUint64 / 2,
		GasPrice: big.NewInt(0),
		Value:    balance,
		Nonce:    new(big.Int).SetUint64(nonce),
		NoSend:   true,
	}, to)
	if err != nil {
		return err
	}

	msg, err := tx.AsMessage(c.signer, big.NewInt(0))
	if err != nil {
		return err
	}

	err = applyTransaction(msg, opts)
	if err != nil {
		return err
	}

	return nil
}

func (c *ContractIntegrator) Slash(to common.Address, opts *ApplyTransactOpts) error {
	return nil
}

type ApplyMessageOpts struct {
	State        *state.StateDB
	Header       *types.Header
	ChainConfig  *chainParams.ChainConfig
	ChainContext core.ChainContext
}

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

func applyTransaction(msg types.Message, opts *ApplyTransactOpts) (err error) {
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
		return err
	}
	*txs = append(*txs, expectedTx)
	var root []byte
	if chainConfig.IsByzantium(opts.Header.Number) {
		opts.State.Finalise(true)
	} else {
		root = opts.State.IntermediateRoot(chainConfig.IsEIP158(opts.Header.Number)).Bytes()
	}
	*usedGas += gasUsed
	receipt := types.NewReceipt(root, false, *usedGas)
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

func applyMessage(
	msg types.Message,
	opts *ApplyMessageOpts,
) (uint64, error) {
	// Create a new context to be used in the EVM environment
	context := core.NewEVMBlockContext(opts.Header, opts.ChainContext, nil)
	// Create a new environment which holds all relevant information
	// about the transaction and calling mechanisms.
	vmenv := vm.NewEVM(context, vm.TxContext{Origin: msg.From(), GasPrice: big.NewInt(0)}, opts.State, opts.ChainConfig, vm.Config{})
	// Apply the transaction to the current State (included in the env)
	ret, returnGas, err := vmenv.Call(
		vm.AccountRef(msg.From()),
		*msg.To(),
		msg.Data(),
		msg.Gas(),
		msg.Value(),
	)
	if err != nil {
		log.Error("apply message failed", "msg", string(ret), "err", err)
	}
	return msg.Gas() - returnGas, err
}

type ConsortiumBackend struct {
	ee *ethapi.PublicBlockChainAPI
}

func NewConsortiumBackend(ee *ethapi.PublicBlockChainAPI) *ConsortiumBackend {
	return &ConsortiumBackend{
		ee,
	}
}

func (b *ConsortiumBackend) CodeAt(ctx context.Context, contract common.Address, blockNumber *big.Int) ([]byte, error) {
	block := rpc.BlockNumberOrHashWithNumber(rpc.BlockNumber(blockNumber.Int64()))
	result, err := b.ee.GetCode(ctx, contract, block)
	if err != nil {
		return nil, err
	}

	return result.MarshalText()
}

func (b *ConsortiumBackend) CallContract(ctx context.Context, call ethereum.CallMsg, blockNumber *big.Int) ([]byte, error) {
	block := rpc.BlockNumberOrHashWithNumber(rpc.BlockNumber(blockNumber.Int64()))
	gas := (hexutil.Uint64)(uint64(math.MaxUint64 / 2))
	data := (hexutil.Bytes)(call.Data)

	result, err := b.ee.Call(ctx, ethapi.TransactionArgs{
		Gas:  &gas,
		To:   call.To,
		Data: &data,
	}, block, nil)
	if err != nil {
		return nil, err
	}

	return result.MarshalText()
}

func (b *ConsortiumBackend) HeaderByNumber(ctx context.Context, number *big.Int) (*types.Header, error) {
	return b.ee.GetHeader(ctx, rpc.BlockNumber(number.Int64()))
}

func (b *ConsortiumBackend) PendingCodeAt(ctx context.Context, account common.Address) ([]byte, error) {
	return nil, errMethodUnimplemented
}

func (b *ConsortiumBackend) PendingNonceAt(ctx context.Context, account common.Address) (uint64, error) {
	return 0, errMethodUnimplemented
}

func (b *ConsortiumBackend) SuggestGasPrice(ctx context.Context) (*big.Int, error) {
	return big.NewInt(0), nil
}

func (b *ConsortiumBackend) SuggestGasTipCap(ctx context.Context) (*big.Int, error) {
	return big.NewInt(0), nil
}

func (b *ConsortiumBackend) EstimateGas(ctx context.Context, call ethereum.CallMsg) (gas uint64, err error) {
	return math.MaxUint64 / 2, nil
}

func (b *ConsortiumBackend) SendTransaction(ctx context.Context, tx *types.Transaction) error {
	// No need to send transaction
	return errMethodUnimplemented
}

func (b *ConsortiumBackend) FilterLogs(ctx context.Context, query ethereum.FilterQuery) ([]types.Log, error) {
	return nil, errMethodUnimplemented
}

func (b *ConsortiumBackend) SubscribeFilterLogs(ctx context.Context, query ethereum.FilterQuery, ch chan<- types.Log) (ethereum.Subscription, error) {
	return nil, errMethodUnimplemented
}
