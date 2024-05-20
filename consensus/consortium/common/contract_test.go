package common

import (
	"bytes"
	"errors"
	"fmt"
	"math/big"
	"reflect"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	legacyProfile "github.com/ethereum/go-ethereum/consensus/consortium/generated_contracts/legacy_profile"
	"github.com/ethereum/go-ethereum/consensus/consortium/generated_contracts/profile"
	roninValidatorSet "github.com/ethereum/go-ethereum/consensus/consortium/generated_contracts/ronin_validator_set"
	"github.com/ethereum/go-ethereum/consensus/consortium/generated_contracts/staking"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/params"
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

func TestContractCall(t *testing.T) {
	roninValidatorSetABI, err := roninValidatorSet.RoninValidatorSetMetaData.GetAbi()
	if err != nil {
		t.Fatal(err)
	}
	stakingABI, err := staking.StakingMetaData.GetAbi()
	if err != nil {
		t.Fatal(err)
	}
	profileABI, err := profile.ProfileMetaData.GetAbi()
	if err != nil {
		t.Fatal(err)
	}
	legacyProfileABI, err := legacyProfile.ProfileMetaData.GetAbi()
	if err != nil {
		t.Fatal(err)
	}

	c := ContractIntegrator{
		roninValidatorSetABI: roninValidatorSetABI,
		profileABI:           profileABI,
		legacyProfileABI:     legacyProfileABI,
		stakingABI:           stakingABI,
		chainConfig: &params.ChainConfig{
			ConsortiumV2Contracts: &params.ConsortiumV2Contracts{},
		},
	}
	c.contractCallHook = func(method string) []byte {
		switch method {
		case "getBlockProducers":
			return common.Hex2Bytes("0000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000001600000000000000000000000032d619dc6188409cebbc52f921ab306f07db085b000000000000000000000000fc3e31519b551bd594235dd0ef014375a87c4e21000000000000000000000000f41af21f0a800dc4d86efb14ad46cfb9884fdf38000000000000000000000000210744c64eea863cf0f972e5aebc683b98fb19840000000000000000000000006e46924371d0e910769aabe0d867590deac20684000000000000000000000000e07d7e56588a6fd860c5073c70a099658c060f3d000000000000000000000000ec702628f44c31acc56c3a59555be47e1f16eb1e000000000000000000000000ee11d2016e9f2fae606b2f12986811f4abbe621500000000000000000000000052349003240770727900b06a3b3a90f5c0219ade0000000000000000000000009b959d27840a31988410ee69991bcf0110d61f020000000000000000000000008eec4f1c0878f73e8e09c1be78ac1465cc16544d000000000000000000000000d11d9842babd5209b9b1155e46f5878c989125b70000000000000000000000006aaabf51c5f6d2d93212cf7dad73d67afa0148d000000000000000000000000005ad3ded6fcc510324af8e2631717af6da5c8b5b000000000000000000000000ae53daac1bf3c4633d4921b8c3f8d579e757f5bc0000000000000000000000004e7ea047ec7e95c7a02cb117128b94ccdd8356bf000000000000000000000000edcafc4ad8097c2012980a2a7087d74b86bddaf900000000000000000000000047cfcb64f8ea44d6ea7fab32f13efa2f8e65eec1000000000000000000000000ca54a1700e0403dcb531f8db4ae3847758b90b010000000000000000000000002bddcaae1c6ccd53e436179b3fc07307ee6f3ef80000000000000000000000004125217ce8868553e1f61bb030426efd330c2d6800000000000000000000000061089875ff9e506ae78c7fe9f7c388416520e386")
		case "getValidatorCandidates":
			return common.Hex2Bytes("0000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000001700000000000000000000000052c0dcd83aa1999ba6c3b0324c8299e30207373c000000000000000000000000f41af21f0a800dc4d86efb14ad46cfb9884fdf38000000000000000000000000e07d7e56588a6fd860c5073c70a099658c060f3d00000000000000000000000052349003240770727900b06a3b3a90f5c0219ade0000000000000000000000002bddcaae1c6ccd53e436179b3fc07307ee6f3ef8000000000000000000000000ec702628f44c31acc56c3a59555be47e1f16eb1e0000000000000000000000004125217ce8868553e1f61bb030426efd330c2d68000000000000000000000000d11d9842babd5209b9b1155e46f5878c989125b700000000000000000000000061089875ff9e506ae78c7fe9f7c388416520e3860000000000000000000000006aaabf51c5f6d2d93212cf7dad73d67afa0148d000000000000000000000000047cfcb64f8ea44d6ea7fab32f13efa2f8e65eec10000000000000000000000008eec4f1c0878f73e8e09c1be78ac1465cc16544d0000000000000000000000009b959d27840a31988410ee69991bcf0110d61f02000000000000000000000000ee11d2016e9f2fae606b2f12986811f4abbe6215000000000000000000000000ca54a1700e0403dcb531f8db4ae3847758b90b010000000000000000000000004e7ea047ec7e95c7a02cb117128b94ccdd8356bf0000000000000000000000006e46924371d0e910769aabe0d867590deac20684000000000000000000000000ae53daac1bf3c4633d4921b8c3f8d579e757f5bc00000000000000000000000005ad3ded6fcc510324af8e2631717af6da5c8b5b00000000000000000000000032d619dc6188409cebbc52f921ab306f07db085b000000000000000000000000210744c64eea863cf0f972e5aebc683b98fb1984000000000000000000000000edcafc4ad8097c2012980a2a7087d74b86bddaf9000000000000000000000000fc3e31519b551bd594235dd0ef014375a87c4e21")
		case "getId2Profile":
			return common.Hex2Bytes("000000000000000000000000000000000000000000000000000000000000002000000000000000000000000032d619dc6188409cebbc52f921ab306f07db085b00000000000000000000000032d619dc6188409cebbc52f921ab306f07db085b0000000000000000000000004e0a599e4dff57965e0dd5bc680f43cc864364c20000000000000000000000004e0a599e4dff57965e0dd5bc680f43cc864364c2000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000001200000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000018000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000030851de18c85f472a7107c584937ac5c6c1caad0c3cb2d4d0977231b91dd669ba4735cf0257a1d52853c3c42f78a287dbe000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000")
		case "getConsensus2Id":
			return common.Hex2Bytes("000000000000000000000000a85ddddceeab43dccaa259dd4936ac104386f9aa")
		case "getId2Pubkey":
			return common.Hex2Bytes("000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000000309256ab3792329b85dc7b633a3f7f99d8f84a8924a27576d89323988f09871deaeb82a18248cd02af3e7837c91d38b62900000000000000000000000000000000")
		case "getManyStakingTotals":
			return common.Hex2Bytes("0000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000008e68da9abcf91c0578116")
		case "maxValidatorNumber":
			return common.Hex2Bytes("0000000000000000000000000000000000000000000000000000000000000064")
		default:
			return nil
		}
	}

	{
		addresses, err := c.GetBlockProducers(common.Hash{}, common.Big0)
		if err != nil {
			t.Fatal(err)
		}

		expectedAddresses := []common.Address{
			common.HexToAddress("0x32D619Dc6188409CebbC52f921Ab306F07DB085b"),
			common.HexToAddress("0xFc3e31519B551bd594235dd0eF014375a87C4e21"),
			common.HexToAddress("0xf41Af21F0A800dc4d86efB14ad46cfb9884FDf38"),
			common.HexToAddress("0x210744C64Eea863Cf0f972e5AEBC683b98fB1984"),
			common.HexToAddress("0x6E46924371d0e910769aaBE0d867590deAC20684"),
			common.HexToAddress("0xE07D7e56588a6FD860c5073c70a099658C060F3D"),
			common.HexToAddress("0xeC702628F44C31aCc56C3A59555be47e1f16eB1e"),
			common.HexToAddress("0xEE11d2016e9f2faE606b2F12986811F4abbe6215"),
			common.HexToAddress("0x52349003240770727900b06a3B3a90f5c0219ADe"),
			common.HexToAddress("0x9B959D27840a31988410Ee69991BCF0110D61F02"),
			common.HexToAddress("0x8Eec4F1c0878F73E8e09C1be78aC1465Cc16544D"),
			common.HexToAddress("0xd11D9842baBd5209b9B1155e46f5878c989125b7"),
			common.HexToAddress("0x6aaABf51C5F6D2D93212Cf7DAD73D67AFa0148d0"),
			common.HexToAddress("0x05ad3Ded6fcc510324Af8e2631717af6dA5C8B5B"),
			common.HexToAddress("0xae53daAC1BF3c4633d4921B8C3F8d579e757F5Bc"),
			common.HexToAddress("0x4E7EA047EC7E95c7a02CB117128B94CCDd8356bf"),
			common.HexToAddress("0xedCafC4Ad8097c2012980A2a7087d74B86bDDAf9"),
			common.HexToAddress("0x47cfcb64f8EA44d6Ea7FAB32f13EFa2f8E65Eec1"),
			common.HexToAddress("0xca54a1700e0403Dcb531f8dB4aE3847758b90B01"),
			common.HexToAddress("0x2bdDcaAE1C6cCd53E436179B3fc07307ee6f3eF8"),
			common.HexToAddress("0x4125217cE8868553e1f61BB030426eFD330c2D68"),
			common.HexToAddress("0x61089875fF9e506ae78C7FE9f7c388416520E386"),
		}

		if !reflect.DeepEqual(addresses, expectedAddresses) {
			t.Fatalf("Block producer mismatches, got %+v\n, expect: %+v", addresses, expectedAddresses)
		}
	}

	{
		addresses, err := c.GetValidatorCandidates(common.Hash{}, common.Big0)
		if err != nil {
			t.Fatal(err)
		}

		expectedAddresses := []common.Address{
			common.HexToAddress("0x52C0dcd83aa1999BA6c3b0324C8299E30207373C"),
			common.HexToAddress("0xf41Af21F0A800dc4d86efB14ad46cfb9884FDf38"),
			common.HexToAddress("0xE07D7e56588a6FD860c5073c70a099658C060F3D"),
			common.HexToAddress("0x52349003240770727900b06a3B3a90f5c0219ADe"),
			common.HexToAddress("0x2bdDcaAE1C6cCd53E436179B3fc07307ee6f3eF8"),
			common.HexToAddress("0xeC702628F44C31aCc56C3A59555be47e1f16eB1e"),
			common.HexToAddress("0x4125217cE8868553e1f61BB030426eFD330c2D68"),
			common.HexToAddress("0xd11D9842baBd5209b9B1155e46f5878c989125b7"),
			common.HexToAddress("0x61089875fF9e506ae78C7FE9f7c388416520E386"),
			common.HexToAddress("0x6aaABf51C5F6D2D93212Cf7DAD73D67AFa0148d0"),
			common.HexToAddress("0x47cfcb64f8EA44d6Ea7FAB32f13EFa2f8E65Eec1"),
			common.HexToAddress("0x8Eec4F1c0878F73E8e09C1be78aC1465Cc16544D"),
			common.HexToAddress("0x9B959D27840a31988410Ee69991BCF0110D61F02"),
			common.HexToAddress("0xEE11d2016e9f2faE606b2F12986811F4abbe6215"),
			common.HexToAddress("0xca54a1700e0403Dcb531f8dB4aE3847758b90B01"),
			common.HexToAddress("0x4E7EA047EC7E95c7a02CB117128B94CCDd8356bf"),
			common.HexToAddress("0x6E46924371d0e910769aaBE0d867590deAC20684"),
			common.HexToAddress("0xae53daAC1BF3c4633d4921B8C3F8d579e757F5Bc"),
			common.HexToAddress("0x05ad3Ded6fcc510324Af8e2631717af6dA5C8B5B"),
			common.HexToAddress("0x32D619Dc6188409CebbC52f921Ab306F07DB085b"),
			common.HexToAddress("0x210744C64Eea863Cf0f972e5AEBC683b98fB1984"),
			common.HexToAddress("0xedCafC4Ad8097c2012980A2a7087d74B86bDDAf9"),
			common.HexToAddress("0xFc3e31519B551bd594235dd0eF014375a87C4e21"),
		}

		if !reflect.DeepEqual(addresses, expectedAddresses) {
			t.Fatalf("Block producer mismatches, got %+v\n, expect: %+v", addresses, expectedAddresses)
		}
	}

	{
		validatorAddress := common.HexToAddress("0x32D619Dc6188409CebbC52f921Ab306F07DB085b")
		blsPublicKey, err := c.getBlsPublicKeyLegacy(common.Hash{}, common.Big0, validatorAddress)
		if err != nil {
			t.Fatal(err)
		}

		expectedKey := common.Hex2Bytes("851de18c85f472a7107c584937ac5c6c1caad0c3cb2d4d0977231b91dd669ba4735cf0257a1d52853c3c42f78a287dbe")
		if !bytes.Equal(blsPublicKey.Marshal(), expectedKey) {
			t.Fatalf("BLS key mismatches, got %v, expect: %v", blsPublicKey.Marshal(), expectedKey)
		}
	}

	{
		validatorAddress := common.HexToAddress("0xA85ddDdCeEaB43DccAa259dd4936aC104386F9aa")
		blsPublicKey, err := c.getBlsPublicKey(common.Hash{}, common.Big0, validatorAddress)
		if err != nil {
			t.Fatal(err)
		}

		expectedKey := common.Hex2Bytes("9256ab3792329b85dc7b633a3f7f99d8f84a8924a27576d89323988f09871deaeb82a18248cd02af3e7837c91d38b629")
		if !bytes.Equal(blsPublicKey.Marshal(), expectedKey) {
			t.Fatalf("BLS key mismatches, got %v, expect: %v", blsPublicKey.Marshal(), expectedKey)
		}
	}

	{
		validatorAddresses := []common.Address{common.HexToAddress("0x32D619Dc6188409CebbC52f921Ab306F07DB085b")}
		stakedAmounts, err := c.GetStakedAmount(common.Hash{}, common.Big0, validatorAddresses)
		if err != nil {
			t.Fatal(err)
		}

		if len(stakedAmounts) != 1 {
			t.Fatalf("Length of staked amounts mismatches, got %d expect %d", len(stakedAmounts), 1)
		}

		expectedAmount := big.NewInt(1076016406498)
		expectedAmount = expectedAmount.Mul(expectedAmount, new(big.Int).Exp(big.NewInt(10), big.NewInt(13), nil))
		expectedAmount = expectedAmount.Add(expectedAmount, big.NewInt(5283175088406))
		if stakedAmounts[0].Cmp(expectedAmount) != 0 {
			t.Fatalf("Staked amounts mismatches, got %+v expect %+v", stakedAmounts, expectedAmount)
		}
	}

	{
		var maxValidatorNumber *big.Int
		maxValidatorNumber, err := c.GetMaxValidatorNumber(common.Hash{}, common.Big0)
		if err != nil {
			t.Fatal(err)
		}
		expectedNumber := big.NewInt(100)
		if maxValidatorNumber.Cmp(expectedNumber) != 0 {
			t.Fatalf("Max validator mismatches, got %+v expect %+v", maxValidatorNumber, expectedNumber)
		}
	}
}
