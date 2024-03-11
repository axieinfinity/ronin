package common

import (
	"errors"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto/bls/blst"
	blsCommon "github.com/ethereum/go-ethereum/crypto/bls/common"
	"github.com/ethereum/go-ethereum/log"
)

var Validators *MockValidators

type MockValidators struct {
	validators    []common.Address
	blsPublicKeys map[common.Address]blsCommon.PublicKey
}

func SetMockValidators(validators, publicKeys string) error {
	vals := strings.Split(validators, ",")
	pubs := strings.Split(publicKeys, ",")
	if len(vals) != len(pubs) {
		return errors.New("mismatch length between mock validators and mock blsPubKey")
	}
	Validators = &MockValidators{
		validators:    make([]common.Address, len(vals)),
		blsPublicKeys: make(map[common.Address]blsCommon.PublicKey),
	}
	for i, val := range vals {
		Validators.validators[i] = common.HexToAddress(val)
		pubKey, err := blst.PublicKeyFromBytes(common.Hex2Bytes(pubs[i]))
		if err != nil {
			return err
		}
		Validators.blsPublicKeys[Validators.validators[i]] = pubKey
	}
	return nil
}

func (m *MockValidators) GetValidators() []common.Address {
	return m.validators
}

func (m *MockValidators) GetPublicKey(addr common.Address) (blsCommon.PublicKey, error) {
	if key, ok := m.blsPublicKeys[addr]; ok {
		return key, nil
	}
	return nil, errors.New("public key not found")
}

type MockContract struct {
}

func (contract *MockContract) GetValidators(*big.Int) ([]common.Address, error) {
	return Validators.GetValidators(), nil
}

func (contract *MockContract) WrapUpEpoch(*ApplyTransactOpts) error {
	log.Info("WrapUpEpoch")
	return nil
}

func (contract *MockContract) SubmitBlockReward(*ApplyTransactOpts) error {
	log.Info("SubmitBlockReward")
	return nil
}

func (contract *MockContract) Slash(*ApplyTransactOpts, common.Address) error {
	log.Info("Slash")
	return nil
}

func (contract *MockContract) FinalityReward(*ApplyTransactOpts, []common.Address) error {
	log.Info("FinalityReward")
	return nil
}

func (contract *MockContract) GetBlsPublicKey(_ *big.Int, addr common.Address) (blsCommon.PublicKey, error) {
	return Validators.GetPublicKey(addr)
}

func (contract *MockContract) GetStakedAmount(_ *big.Int, _ []common.Address) ([]*big.Int, error) {
	return nil, nil
}
