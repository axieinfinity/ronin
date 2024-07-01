package state

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/rawdb"
)

func TestIsAddressBlacklisted(t *testing.T) {
	blackListContract := common.Address{0x11}

	statedb, _ := New(common.Hash{}, NewDatabase(rawdb.NewMemoryDatabase()), nil)

	blacklistedAddress := common.BigToAddress(common.Big3)
	// Blacklist address 0x000..0003
	statedb.SetState(
		blackListContract,
		common.HexToHash("0x7dfe757ecd65cbd7922a9c0161e935dd7fdbcc0e999689c7d31633896b1fc60b"),
		common.BigToHash(common.Big1),
	)
	if !IsAddressBlacklisted(statedb, &blackListContract, &blacklistedAddress) {
		t.Fatalf("Expect address %s to be blacklisted", blacklistedAddress.String())
	}

	notBlacklistedAddress := common.BigToAddress(big.NewInt(10))
	if IsAddressBlacklisted(statedb, &blackListContract, &notBlacklistedAddress) {
		t.Fatalf("Expect address %s to be not blacklisted", notBlacklistedAddress.String())
	}

	statedb.SetState(blackListContract, common.BigToHash(common.Big2), common.BigToHash(common.Big1))
	if IsAddressBlacklisted(statedb, &blackListContract, &blacklistedAddress) {
		t.Fatalf("Expect address %s to be not blacklisted", blacklistedAddress.String())
	}
}

func BenchmarkIsAddressBlacklisted(b *testing.B) {
	statedb, _ := New(common.Hash{}, NewDatabase(rawdb.NewMemoryDatabase()), nil)
	blackListContract := common.Address{0x11}

	queriedAddress := common.BigToAddress(common.Big3)

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		IsAddressBlacklisted(statedb, &blackListContract, &queriedAddress)
	}
}
