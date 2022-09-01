package common

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus/consortium/generated_contracts/validators"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"
	"io/ioutil"
	"math/big"
	"testing"
)

const contractAddress = "0x089f10d52008F962f9E09EFBD2E5275BFf56045b"

func loadKey() (*keystore.Key, error) {
	keyjson, err := ioutil.ReadFile("/Users/mac/coding/ronin/PoS/local1/keystore/UTC--2022-08-21T07-22-03.047965000Z--da0479bed856764502249bec9a3acd1c3da2cf23")
	if err != nil {
		return nil, err
	}
	password := "123456"
	return keystore.DecryptKey(keyjson, password)
}

func TestDeployContract(t *testing.T) {
	key, err := loadKey()
	if err != nil {
		t.Fatal(err)
	}

	client, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		t.Fatal(err)
	}

	address := common.HexToAddress("da0479bed856764502249bec9a3acd1c3da2cf23")
	nonce, err := client.NonceAt(context.Background(), address, nil)
	if err != nil {
		t.Fatal(err)
	}

	contractAddress, tx, _, err := validators.DeployValidators(&bind.TransactOpts{
		From:  address,
		Nonce: big.NewInt(int64(nonce)),
		Signer: func(addr common.Address, tx *types.Transaction) (*types.Transaction, error) {
			signer := types.LatestSignerForChainID(big.NewInt(2022))
			return types.SignTx(tx, signer, key.PrivateKey)
		},
		Context: context.Background(),
	}, client, "test")

	if err != nil {
		t.Fatal(err)
	}
	println(tx.Hash().Hex())
	println(fmt.Sprintf("contractAddres:%s", contractAddress.Hex()))
}

func TestAddNode(t *testing.T) {
	client, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		t.Fatal(err)
	}
	key, err := loadKey()
	if err != nil {
		t.Fatal(err)
	}
	address := common.HexToAddress("da0479bed856764502249bec9a3acd1c3da2cf23")
	contractIntegrator, err := NewContractIntegrator(&params.ChainConfig{
		ConsortiumV2Contracts: &params.ConsortiumV2Contracts{
			ValidatorSC: common.HexToAddress(contractAddress),
		},
	}, client, nil, address)
	if err != nil {
		t.Fatal(err)
	}
	nonce, err := client.NonceAt(context.Background(), address, nil)
	if err != nil {
		t.Fatal(err)
	}
	tx, err := contractIntegrator.validatorSC.AddNode(&bind.TransactOpts{
		From:  address,
		Nonce: big.NewInt(int64(nonce)),
		Signer: func(addr common.Address, tx *types.Transaction) (*types.Transaction, error) {
			signer := types.LatestSignerForChainID(big.NewInt(2022))
			return types.SignTx(tx, signer, key.PrivateKey)
		},
		Context: context.Background(),
	}, address, address)
	if err != nil {
		t.Fatal(err)
	}
	println(fmt.Sprintf("address:%s has been added, tx:%s", address.Hex(), tx.Hash().Hex()))

	address2 := common.HexToAddress("bd1baea7e8a4f6c156039adc536c5bbce68add59")
	tx, err = contractIntegrator.validatorSC.AddNode(&bind.TransactOpts{
		From:  address,
		Nonce: big.NewInt(int64(nonce + 1)),
		Signer: func(addr common.Address, tx *types.Transaction) (*types.Transaction, error) {
			signer := types.LatestSignerForChainID(big.NewInt(2022))
			return types.SignTx(tx, signer, key.PrivateKey)
		},
		Context: context.Background(),
	}, address2, address2)
	if err != nil {
		t.Fatal(err)
	}
	println(fmt.Sprintf("address:%s has been added, tx:%s", address2.Hex(), tx.Hash().Hex()))
}

func TestGetNonce(t *testing.T) {
	client, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		t.Fatal(err)
	}
	key, err := loadKey()
	if err != nil {
		t.Fatal(err)
	}
	nonce, err := client.NonceAt(context.Background(), key.Address, nil)
	if err != nil {
		t.Fatal(err)
	}
	println(nonce)
}

func TestAddMoreNode(t *testing.T) {
	client, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		t.Fatal(err)
	}
	key, err := loadKey()
	if err != nil {
		t.Fatal(err)
	}
	address := common.HexToAddress("089c3107402ae0d06d5953347a4c82ac8ce66f6c")
	contractIntegrator, err := NewContractIntegrator(&params.ChainConfig{
		ConsortiumV2Contracts: &params.ConsortiumV2Contracts{
			ValidatorSC: common.HexToAddress(contractAddress),
		},
	}, client, nil, address)
	if err != nil {
		t.Fatal(err)
	}
	nonce, err := client.NonceAt(context.Background(), key.Address, nil)
	if err != nil {
		t.Fatal(err)
	}
	tx, err := contractIntegrator.validatorSC.AddNode(&bind.TransactOpts{
		From:  key.Address,
		Nonce: big.NewInt(int64(nonce)),
		Signer: func(addr common.Address, tx *types.Transaction) (*types.Transaction, error) {
			signer := types.LatestSignerForChainID(big.NewInt(2022))
			return types.SignTx(tx, signer, key.PrivateKey)
		},
		Context: context.Background(),
	}, address, address)
	if err != nil {
		t.Fatal(err)
	}
	println(fmt.Sprintf("address:%s has been added, tx:%s", address.Hex(), tx.Hash().Hex()))

}

func TestGetLatestValidators(t *testing.T) {
	client, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		t.Fatal(err)
	}
	contractIntegrator, err := NewContractIntegrator(&params.ChainConfig{
		ConsortiumV2Contracts: &params.ConsortiumV2Contracts{
			ValidatorSC: common.HexToAddress(contractAddress),
		},
	}, client, nil, common.Address{})
	if err != nil {
		t.Fatal(err)
	}
	vals, err := contractIntegrator.validatorSC.GetValidators(nil)
	if err != nil {
		t.Fatal(err)
	}
	for _, val := range vals {
		println(val.Hex())
	}
}

func TestTransfer(t *testing.T) {
	client, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		t.Fatal(err)
	}
	key, err := loadKey()
	if err != nil {
		t.Fatal(err)
	}
	nonce, err := client.NonceAt(context.Background(), key.Address, nil)
	if err != nil {
		t.Fatal(err)
	}
	to := common.HexToAddress("0xBD1baEa7e8a4F6C156039Adc536C5BBcE68ADd59")
	signer := types.LatestSignerForChainID(big.NewInt(2022))
	tx, err := types.SignNewTx(key.PrivateKey, signer,
		&types.LegacyTx{
			Nonce:    nonce,
			GasPrice: big.NewInt(500),
			Gas:      21000,
			To:       &to,
			Value:    big.NewInt(0).Mul(big.NewInt(1), big.NewInt(0).Exp(big.NewInt(10), big.NewInt(18), nil)),
		})
	if err != nil {
		t.Fatal(err)
	}
	if err = client.SendTransaction(context.Background(), tx); err != nil {
		t.Fatal(err)
	}
	println(tx.Hash().Hex())
}
