package finality

import (
	"encoding/json"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto/bls/blst"
)

func TestFinalityVoteBitSet(t *testing.T) {
	var bitSet BitSet

	bitSet.SetBit(0)
	bitSet.SetBit(40)
	// index >= 64 has no effect
	bitSet.SetBit(64)

	// 2 ** 40 + 2 ** 0
	if uint64(bitSet) != 1099511627777 {
		t.Fatalf("Wrong bitset value, exp %d got %d", 1099511627777, uint64(bitSet))
	}

	indices := bitSet.Indices()
	if len(indices) != 2 {
		t.Fatalf("Wrong indices length, exp %d got %d", 2, len(indices))
	}
	if indices[0] != 0 {
		t.Fatalf("Wrong index, exp %d got %d", 0, indices[0])
	}
	if indices[1] != 40 {
		t.Fatalf("Wrong index, exp %d got %d", 40, indices[1])
	}

	if bitSet.GetBit(40) != 1 {
		t.Fatalf("Wrong bit, exp %d got %d", 1, bitSet.GetBit(40))
	}
	if bitSet.GetBit(50) != 0 {
		t.Fatalf("Wrong bit, exp %d got %d", 1, bitSet.GetBit(50))
	}
	// index >= 64 returns 0
	if bitSet.GetBit(70) != 0 {
		t.Fatalf("Wrong bit, exp %d got %d", 0, bitSet.GetBit(70))
	}
}

func TestMarshallJsonValidatorWithPub(t *testing.T) {
	blsPubkey, err := blst.PublicKeyFromBytes(common.Hex2Bytes("affe116dad1eb59bda4dc6a6442a891c1b19502c07acff7304a070a5d53b7bf89c01d6985a1ad0d86b009b39e5cd7a9e"))
	if err != nil {
		t.Fatal(err)
	}
	validator := ValidatorWithBlsPub{
		Address:      common.BigToAddress(big.NewInt(1)),
		BlsPublicKey: blsPubkey,
		Weight:       100,
	}
	marshalled, err := json.Marshal(validator)
	if err != nil {
		t.Fatal(err)
	}

	var unmarshalledValidator ValidatorWithBlsPub
	err = json.Unmarshal(marshalled, &unmarshalledValidator)
	if err != nil {
		t.Fatal(err)
	}
	if validator.Address != unmarshalledValidator.Address {
		t.Fatalf("Address mismatches, got %v expect %v", unmarshalledValidator.Address, validator.Address)
	}
	if !validator.BlsPublicKey.Equals(unmarshalledValidator.BlsPublicKey) {
		t.Fatalf("BLS public key mismatches, got %v expect %v", unmarshalledValidator.BlsPublicKey, validator.BlsPublicKey)
	}
	if validator.Weight != unmarshalledValidator.Weight {
		t.Fatalf("Weight mismatches, got %d expect %d", unmarshalledValidator.Weight, validator.Weight)
	}
}
