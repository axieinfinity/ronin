package proxy

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/internal/ethapi"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/rpc"
	"math/big"
)

type API struct {
	*ethapi.PublicTransactionPoolAPI
	*ethapi.PublicBlockChainAPI
	*ethapi.PublicEthereumAPI
	b *backend
}

func newAPI(b *backend) *API {
	return &API{
		ethapi.NewPublicTransactionPoolAPI(b, &ethapi.AddrLocker{}),
		ethapi.NewPublicBlockChainAPI(b),
		ethapi.NewPublicEthereumAPI(b),
		b,
	}
}

func (api *API) SendTransaction(ctx context.Context, args ethapi.TransactionArgs) (common.Hash, error) {
	log.Trace("Sending transaction from proxy")
	return SubmitTransaction(ctx, api.b, args.ToTransaction())
}

func (api *API) SendRawTransaction(ctx context.Context, input hexutil.Bytes) (common.Hash, error) {
	log.Trace("Sending raw transaction from proxy")
	tx := new(types.Transaction)
	if err := tx.UnmarshalBinary(input); err != nil {
		return common.Hash{}, err
	}
	return SubmitTransaction(ctx, api.b, tx)
}

// SubmitTransaction is a helper function that submits tx via rpcUrl or freeGasProxyUrl and logs a message.
func SubmitTransaction(ctx context.Context, b *backend, tx *types.Transaction) (common.Hash, error) {
	if tx.Gas() == 0 && b.fgpClient != nil {
		if err := b.fgpClient.SendTransaction(ctx, tx); err != nil {
			return common.Hash{}, err
		}
	}
	return ethapi.SubmitTransaction(ctx, b, tx)
}

// EstimateGas returns an estimate of the amount of gas needed to execute the
// given transaction against the current pending block.
func (api *API) EstimateGas(ctx context.Context, args ethapi.TransactionArgs, blockNrOrHash *rpc.BlockNumberOrHash) (hexutil.Uint64, error) {
	bNrOrHash := rpc.BlockNumberOrHashWithNumber(rpc.LatestBlockNumber)
	if blockNrOrHash != nil {
		bNrOrHash = *blockNrOrHash
	}
	return ethapi.DoEstimateGas(ctx, api.b, args, bNrOrHash, api.b.RPCGasCap())
}

func (api *API) GasPrice(ctx context.Context) (*hexutil.Big, error) {
	gp, err := api.b.client.SuggestGasPrice(ctx)
	if err != nil {
		return nil, err
	}
	return (*hexutil.Big)(gp), nil
}

func (api *API) ChainId() (*hexutil.Big, error) {
	return (*hexutil.Big)(api.b.ChainConfig().ChainID), nil
}

// GetTransactionCount returns the number of transactions the given address has sent for the given block number
func (api *API) GetTransactionCount(ctx context.Context, address common.Address, blockNrOrHash rpc.BlockNumberOrHash) (*hexutil.Uint64, error) {
	// Ask transaction pool for the nonce which includes pending transactions
	if blockNr, ok := blockNrOrHash.Number(); ok && blockNr == rpc.PendingBlockNumber {
		nonce, err := api.b.client.NonceAt(ctx, address, big.NewInt(int64(blockNr)))
		if err != nil {
			return nil, err
		}
		return (*hexutil.Uint64)(&nonce), nil
	}
	// Resolve block number and use its state to ask for the nonce
	state, _, err := api.b.StateAndHeaderByNumberOrHash(ctx, blockNrOrHash)
	if state == nil || err != nil {
		return nil, err
	}
	nonce := state.GetNonce(address)
	return (*hexutil.Uint64)(&nonce), state.Error()
}
