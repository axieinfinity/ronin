package common

import (
	"errors"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

func TestApplyTransactionSender(t *testing.T) {
	signer := types.NewMikoSigner(big.NewInt(2020))
	privateKey, _ := crypto.GenerateKey()
	sender := crypto.PubkeyToAddress(privateKey.PublicKey)

	tx, err := types.SignNewTx(privateKey, signer, &types.LegacyTx{
		Nonce:    0,
		GasPrice: common.Big0,
		Gas:      21000,
		To:       &sender,
	})
	if err != nil {
		t.Fatalf("Failed to sign transaction, err: %s", err)
	}

	state, err := state.New(common.Hash{}, state.NewDatabase(rawdb.NewMemoryDatabase()), nil)
	if err != nil {
		t.Fatalf("Failed to create stateDB, err %s", err)
	}

	msg, err := tx.AsMessage(signer, nil)
	if err != nil {
		t.Fatalf("Failed to create message, err %s", err)
	}
	err = ApplyTransaction(
		msg,
		&ApplyTransactOpts{
			ApplyMessageOpts: &ApplyMessageOpts{
				State:  state,
				Header: &types.Header{},
			},
			Signer: signer,
		},
	)
	if !errors.Is(err, core.ErrSenderNoEOA) {
		t.Fatalf("Expect err: %s, have %s", core.ErrSenderNoEOA, err)
	}
}
