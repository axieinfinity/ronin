package proxy

import (
	"context"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/internal/ethapi"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/params"
	"math/big"
)

type API struct {
	*ethapi.PublicTransactionPoolAPI
	*ethapi.PublicBlockChainAPI
	b *backend
}

func newAPI(b *backend) *API {
	return &API{
		ethapi.NewPublicTransactionPoolAPI(b, &ethapi.AddrLocker{}),
		ethapi.NewPublicBlockChainAPI(b),
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
	} else {
		// If the transaction fee cap is already specified, ensure the
		// fee of the given transaction is _reasonable_.
		if err := checkTxFee(tx.GasPrice(), tx.Gas(), b.RPCTxFeeCap()); err != nil {
			return common.Hash{}, err
		}
		if !b.UnprotectedAllowed() && !tx.Protected() {
			// Ensure only eip155 signed transactions are submitted if EIP155Required is set.
			return common.Hash{}, errors.New("only replay-protected (EIP-155) transactions allowed over RPC")
		}
		if err := b.SendTx(ctx, tx); err != nil {
			return common.Hash{}, err
		}
	}
	// Print a log with full tx details for manual investigations and interventions
	signer := types.MakeSigner(b.ChainConfig(), b.CurrentBlock().Number())
	from, err := types.Sender(signer, tx)
	if err != nil {
		return common.Hash{}, err
	}

	if tx.To() == nil {
		addr := crypto.CreateAddress(from, tx.Nonce())
		log.Info("Submitted contract creation", "hash", tx.Hash().Hex(), "from", from, "nonce", tx.Nonce(), "contract", addr.Hex(), "value", tx.Value())
	} else {
		log.Info("Submitted transaction", "hash", tx.Hash().Hex(), "from", from, "nonce", tx.Nonce(), "recipient", tx.To(), "value", tx.Value())
	}
	return tx.Hash(), nil
}

// checkTxFee is an internal function used to check whether the fee of
// the given transaction is _reasonable_(under the cap).
func checkTxFee(gasPrice *big.Int, gas uint64, cap float64) error {
	// Short circuit if there is no cap for transaction fee at all.
	if cap == 0 {
		return nil
	}
	feeEth := new(big.Float).Quo(new(big.Float).SetInt(new(big.Int).Mul(gasPrice, new(big.Int).SetUint64(gas))), new(big.Float).SetInt(big.NewInt(params.Ether)))
	feeFloat, _ := feeEth.Float64()
	if feeFloat > cap {
		return fmt.Errorf("tx fee (%.2f ether) exceeds the configured cap (%.2f ether)", feeFloat, cap)
	}
	return nil
}
