package vkmswallet

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
	"testing"
)

var defaultWalletConfig = &WalletConfig{
	KeyUsageTokenPath: "./key.token",
	VKMSAddress:       "127.0.0.1:51978",
	SourceAddress:     ":6969",
}

var wallet *Wallet

func initWallet() (err error) {
	if wallet == nil {
		wallet, err = NewWallet(defaultWalletConfig)
	}
	return
}

func TestWallet_SignTx_LegacyTx(t *testing.T) {
	// init the wallet
	err := initWallet()
	if err != nil {
		t.Fatal(err)
	}
	acc := wallet.Accounts()[0]

	// create a dummy legacy transaction
	var chainId *big.Int = nil
	txData := &types.LegacyTx{
		Nonce:    0,
		GasPrice: big.NewInt(0),
		Gas:      0,
		To:       &common.Address{},
		Value:    big.NewInt(2),
		Data:     nil,
	}
	tx := types.NewTx(txData)

	// sign the transaction
	tx, err = wallet.SignTx(acc, tx, chainId)
	if err != nil {
		t.Fatal(err)
	}

	// verify the signature (by recover the sender and compare with the wallet's account address)
	signer := types.LatestSignerForChainID(chainId)
	accAddr, err := signer.Sender(tx)
	if err != nil {
		t.Fatal(err)
	}
	if acc.Address != accAddr {
		t.Fatal("address mismatch")
	}
}

func TestWallet_SignTx_DynamicFeeTx(t *testing.T) {
	// init the wallet
	err := initWallet()
	if err != nil {
		t.Fatal(err)
	}
	acc := wallet.Accounts()[0]

	// create a dummy dynamic-free transaction
	var chainId = big.NewInt(1337)
	txData := &types.DynamicFeeTx{
		ChainID:    chainId,
		Nonce:      0,
		GasTipCap:  big.NewInt(0),
		GasFeeCap:  big.NewInt(0),
		Gas:        0,
		To:         &common.Address{},
		Value:      big.NewInt(2),
		Data:       nil,
		AccessList: nil,
	}
	tx := types.NewTx(txData)

	// sign the transaction
	tx, err = wallet.SignTx(acc, tx, chainId)
	if err != nil {
		t.Fatal(err)
	}

	// verify the signature (by recover the sender and compare with the wallet's account address)
	signer := types.LatestSignerForChainID(chainId)
	accAddr, err := signer.Sender(tx)
	if err != nil {
		t.Fatal(err)
	}
	if acc.Address != accAddr {
		t.Fatal("address mismatch")
	}
}

func TestWallet_SignTx_AccessListTx(t *testing.T) {
	// init the wallet
	err := initWallet()
	if err != nil {
		t.Fatal(err)
	}
	acc := wallet.Accounts()[0]

	// create a dummy dynamic-free transaction
	var chainId = big.NewInt(1337)
	txData := &types.AccessListTx{
		ChainID:    chainId,
		Nonce:      0,
		GasPrice:   big.NewInt(0),
		Gas:        0,
		To:         &common.Address{},
		Value:      big.NewInt(2),
		Data:       nil,
		AccessList: nil,
	}
	tx := types.NewTx(txData)

	// sign the transaction
	tx, err = wallet.SignTx(acc, tx, chainId)
	if err != nil {
		t.Fatal(err)
	}

	// verify the signature (by recover the sender and compare with the wallet's account address)
	signer := types.LatestSignerForChainID(chainId)
	accAddr, err := signer.Sender(tx)
	if err != nil {
		t.Fatal(err)
	}
	if acc.Address != accAddr {
		t.Fatal("address mismatch")
	}
}
