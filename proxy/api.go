package proxy

import (
	"context"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/rawdb"
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

// BlockNumber returns the block number of the chain head.
func (api *API) BlockNumber() (hexutil.Uint64, error) {
	header, err := api.b.HeaderByNumber(context.Background(), rpc.LatestBlockNumber)
	if err != nil {
		log.Error("[api][BlockNumber] cannot get latest number", "err", err)
		return hexutil.Uint64(0), err
	}
	return hexutil.Uint64(header.Number.Uint64()), nil
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
	gp, err := api.b.rpc.SuggestGasPrice(ctx)
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
		nonce, err := api.b.rpc.NonceAt(ctx, address, big.NewInt(int64(blockNr)))
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

func (api *API) GetFreeGasRequests(ctx context.Context, address common.Address) (int64, error) {
	if api.b.fgpClient == nil {
		return -1, errors.New("eth_getFreeGasRequests does not exist")
	}
	var result int64
	if err := api.b.fgpClient.RpcClient().Call(&result, "eth_getFreeGasRequests", address); err != nil {
		return -1, err
	}
	return result, nil
}

func (api *API) GetTransactionReceipt(ctx context.Context, hash common.Hash) (map[string]interface{}, error) {
	var receipt *types.Receipt
	tx, blockHash, blockNumber, index, err := api.b.GetTransaction(ctx, hash)
	if err != nil || tx == nil {
		log.Warn("[proxy][backend] transaction not found or error occurred - calling rpc directly", "err", err)
		// directly call via rpc and archive
		if receipt, err = api.b.rpc.TransactionReceipt(ctx, hash); err != nil {
			log.Warn("[proxy][backend] failed on getting transactionReceipt via rpc - try with archive", "err", err, "hash", hash.Hex())
			if api.b.archive != nil {
				if receipt, err = api.b.archive.TransactionReceipt(ctx, hash); err != nil {
					log.Warn("[proxy][backend] failed on getting transactionReceipt via archive", "err", err, "hash", hash.Hex())
					return nil, nil
				}
			}
		}
		// cache receipt to db
		api.b.writeReceiptAncient(receipt)
	} else {
		receipts := rawdb.ReadReceipts(api.b.db, blockHash, blockNumber, api.b.ChainConfig())
		if receipts == nil || len(receipts) <= int(index) {
			log.Warn(fmt.Sprintf("[proxy][backend] receipts not found at hash:%s and number:%d", blockHash.Hex(), blockNumber))
			return nil, nil
		}
		receipt = receipts[index]
	}

	// Derive the sender.
	bigblock := new(big.Int).SetUint64(blockNumber)
	signer := types.MakeSigner(api.b.ChainConfig(), bigblock)
	from, _ := types.Sender(signer, tx)

	fields := map[string]interface{}{
		"blockHash":         blockHash,
		"blockNumber":       hexutil.Uint64(blockNumber),
		"transactionHash":   hash,
		"transactionIndex":  hexutil.Uint64(index),
		"from":              from,
		"to":                tx.To(),
		"gasUsed":           hexutil.Uint64(receipt.GasUsed),
		"cumulativeGasUsed": hexutil.Uint64(receipt.CumulativeGasUsed),
		"contractAddress":   nil,
		"logs":              receipt.Logs,
		"logsBloom":         receipt.Bloom,
		"type":              hexutil.Uint(tx.Type()),
	}
	// Assign the effective gas price paid
	if !api.b.ChainConfig().IsLondon(bigblock) {
		fields["effectiveGasPrice"] = hexutil.Uint64(tx.GasPrice().Uint64())
	} else {
		header, err := api.b.HeaderByHash(ctx, blockHash)
		if err != nil {
			return nil, err
		}
		gasPrice := new(big.Int).Add(header.BaseFee, tx.EffectiveGasTipValue(header.BaseFee))
		fields["effectiveGasPrice"] = hexutil.Uint64(gasPrice.Uint64())
	}
	// Assign receipt status or post state.
	if len(receipt.PostState) > 0 {
		fields["root"] = hexutil.Bytes(receipt.PostState)
	} else {
		fields["status"] = hexutil.Uint(receipt.Status)
	}
	if receipt.Logs == nil {
		fields["logs"] = [][]*types.Log{}
	}
	// If the ContractAddress is 20 0x0 bytes, assume it is not a contract creation
	if receipt.ContractAddress != (common.Address{}) {
		fields["contractAddress"] = receipt.ContractAddress
	}
	return fields, nil
}
