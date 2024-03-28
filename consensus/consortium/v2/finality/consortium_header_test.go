package finality

import "testing"

func TestFinalityVoteBitSet(t *testing.T) {
	var bitSet FinalityVoteBitSet

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
