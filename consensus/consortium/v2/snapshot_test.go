package v2

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/params"
)

type mockChainReader struct {
	headerMapping map[common.Hash]*types.Header
}

func (chainReader *mockChainReader) Config() *params.ChainConfig  { return nil }
func (chainReader *mockChainReader) CurrentHeader() *types.Header { return nil }
func (chainReader *mockChainReader) GetHeader(hash common.Hash, number uint64) *types.Header {
	return chainReader.headerMapping[hash]
}
func (chainReader *mockChainReader) GetHeaderByNumber(number uint64) *types.Header  { return nil }
func (chainReader *mockChainReader) GetHeaderByHash(hash common.Hash) *types.Header { return nil }
func (chainReader *mockChainReader) DB() ethdb.Database                             { return nil }
func (chainReader *mockChainReader) StateCache() state.Database                     { return nil }
func (chainReader *mockChainReader) OpEvents() []*vm.PublishEvent                   { return nil }

func TestFindCheckpointHeader(t *testing.T) {
	// Case 1: checkpoint header is at block 5 (in parent list)
	// parent list ranges from [0, 10)
	parents := make([]*types.Header, 10)
	for i := range parents {
		parents[i] = &types.Header{Number: big.NewInt(int64(i)), Coinbase: common.BigToAddress(big.NewInt(int64(i)))}
	}

	currentHeader := &types.Header{Number: big.NewInt(10)}
	checkpointHeader := findAncestorHeader(currentHeader, 5, nil, parents)
	if checkpointHeader.Number.Cmp(big.NewInt(5)) != 0 && checkpointHeader.Coinbase != common.BigToAddress(big.NewInt(5)) {
		t.Fatalf("Expect checkpoint header number: %d, got: %d", 5, checkpointHeader.Number.Int64())
	}

	// Case 2: checkpoint header is at 5 (lower than parent list)
	// parent list ranges from [10, 20)
	for i := range parents {
		parents[i] = &types.Header{Number: big.NewInt(int64(i + 10)), ParentHash: common.BigToHash(big.NewInt(int64(i + 10 - 1)))}
	}
	mockChain := mockChainReader{
		headerMapping: make(map[common.Hash]*types.Header),
	}
	// create mock chain 1
	for i := 5; i < 10; i++ {
		mockChain.headerMapping[common.BigToHash(big.NewInt(int64(100+i)))] = &types.Header{
			Number:     big.NewInt(int64(i)),
			ParentHash: common.BigToHash(big.NewInt(int64(100 + i - 1))),
		}
	}

	// create mock chain 2
	for i := 5; i < 10; i++ {
		mockChain.headerMapping[common.BigToHash(big.NewInt(int64(i)))] = &types.Header{
			Number:     big.NewInt(int64(i)),
			ParentHash: common.BigToHash(big.NewInt(int64(i - 1))),
		}
	}

	currentHeader = &types.Header{ParentHash: common.BigToHash(big.NewInt(19)), Number: big.NewInt(20)}
	// Must traverse and get the correct header in chain 2
	checkpointHeader = findAncestorHeader(currentHeader, 5, &mockChain, parents)
	if checkpointHeader == nil {
		t.Fatal("Failed to find checkpoint header")
	}
	if checkpointHeader.Number.Cmp(big.NewInt(5)) != 0 && checkpointHeader.ParentHash != common.BigToHash(big.NewInt(int64(4))) {
		t.Fatalf("Expect checkpoint header number %d, parent hash: %s, got number: %d, parent hash: %s",
			5, common.BigToHash(big.NewInt(int64(4))),
			checkpointHeader.Number.Int64(), checkpointHeader.ParentHash,
		)
	}

	// Case 3: find checkpoint header with nil parent list
	currentHeader = &types.Header{Number: big.NewInt(10), ParentHash: common.BigToHash(big.NewInt(109))}
	checkpointHeader = findAncestorHeader(currentHeader, 5, &mockChain, nil)
	// Must traverse and get the correct header in chain 1
	if checkpointHeader == nil {
		t.Fatal("Failed to find checkpoint header")
	}
	if checkpointHeader.Number.Cmp(big.NewInt(5)) != 0 && checkpointHeader.ParentHash != common.BigToHash(big.NewInt(int64(104))) {
		t.Fatalf("Expect checkpoint header number %d, parent hash: %s, got number: %d, parent hash: %s",
			5, common.BigToHash(big.NewInt(int64(104))),
			checkpointHeader.Number.Int64(), checkpointHeader.ParentHash,
		)
	}

	// Case 4: checkpoint header is higher than parent list, this must not happen
	// but the function must not crash in this case
	// parent list ranges from [0, 10)
	parents = make([]*types.Header, 10)
	for i := range parents {
		parents[i] = &types.Header{Number: big.NewInt(int64(i)), Coinbase: common.BigToAddress(big.NewInt(int64(i)))}
	}
	checkpointHeader = findAncestorHeader(nil, 10, nil, parents)
	if checkpointHeader != nil {
		t.Fatalf("Expect %v checkpoint header, got %v", nil, checkpointHeader)
	}
}
