package common

import (
	"errors"
	"fmt"
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
	// Sender is an empty account, we must not get core.ErrSenderNoEOA
	if errors.Is(err, core.ErrSenderNoEOA) {
		t.Fatalf("Don't expect err: %s, have %s", core.ErrSenderNoEOA, err)
	}

	// Sender is not an empty account but still has no code, we must
	// not get core.ErrSenderNoEOA
	state.SetBalance(sender, common.Big1)
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
	if errors.Is(err, core.ErrSenderNoEOA) {
		t.Fatalf("Don't expect err: %s, have %s", core.ErrSenderNoEOA, err)
	}

	// Sender has code, is not an EOA
	state.SetCode(sender, []byte{0x1})
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

	// Sender is an EOA but the sender of system transaction
	// does not match with coinbase
	state.SetCode(sender, []byte{})
	coinbase := common.Address{0x2}
	err = ApplyTransaction(
		msg,
		&ApplyTransactOpts{
			ApplyMessageOpts: &ApplyMessageOpts{
				State:  state,
				Header: &types.Header{Coinbase: coinbase},
			},
			Signer: signer,
		},
	)

	expectedErr := fmt.Errorf("sender of system transaction is not coinbase, sender: %s, coinbase: %s", sender, coinbase)
	if err == nil || err.Error() != expectedErr.Error() {
		t.Fatalf("Expect err: %s, have %s", expectedErr, err)
	}
}
