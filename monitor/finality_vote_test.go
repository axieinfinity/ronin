package monitor

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto/bls/blst"
	blsCommon "github.com/ethereum/go-ethereum/crypto/bls/common"
)

func TestCheckSameHeightVote(t *testing.T) {
	monitor, err := NewFinalityVoteMonitor(nil, nil)
	if err != nil {
		t.Fatalf("Failed to create finality vote monitor, err %s", err)
	}

	key1, err := blst.RandKey()
	if err != nil {
		t.Fatalf("Failed to create bls key, err %s", err)
	}
	address1 := common.Address{0x1}

	key2, err := blst.RandKey()
	if err != nil {
		t.Fatalf("Failed to create bls key, err %s", err)
	}
	address2 := common.Address{0x2}

	voterPublicKey := []blsCommon.PublicKey{key1.PublicKey()}
	voterAddress := []common.Address{address1}
	if monitor.checkSameHeightVote(0, common.Hash{0x1}, voterPublicKey, voterAddress) != nil {
		t.Fatalf("Expect no error when checkSameHeightVote")
	}

	voterPublicKey = []blsCommon.PublicKey{key2.PublicKey()}
	voterAddress = []common.Address{address2}
	if monitor.checkSameHeightVote(0, common.Hash{0x2}, voterPublicKey, voterAddress) != nil {
		t.Fatalf("Expect no error when checkSameHeightVote")
	}

	voterPublicKey = []blsCommon.PublicKey{key2.PublicKey()}
	voterAddress = []common.Address{address2}
	if monitor.checkSameHeightVote(0, common.Hash{0x3}, voterPublicKey, voterAddress) == nil {
		t.Fatalf("Expect error when checkSameHeightVote")
	}
}
