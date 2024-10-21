package vote

import (
	"container/heap"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math/big"
	"os"
	"path/filepath"
	"testing"
	"time"

	wallet "github.com/ethereum/go-ethereum/accounts/bls"
	"github.com/ethereum/go-ethereum/crypto/bls"
	"github.com/google/uuid"
	keystorev4 "github.com/wealdtech/go-eth2-wallet-encryptor-keystorev4"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus"
	"github.com/ethereum/go-ethereum/consensus/ethash"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/crypto"
	blsCommon "github.com/ethereum/go-ethereum/crypto/bls/common"
	"github.com/ethereum/go-ethereum/eth/downloader"
	"github.com/ethereum/go-ethereum/event"
	"github.com/ethereum/go-ethereum/params"
)

var (
	// testKey is a private key to use for funding a tester account.
	testKey, _ = crypto.HexToECDSA("b71c71a67e1177ad4e901695e1b4b9ee17ae16c6668d313eac2f96dbcda3f291")

	// testAddr is the Ethereum address of the tester account.
	testAddr = crypto.PubkeyToAddress(testKey.PublicKey)

	password = "secretPassword"

	timeThreshold = 30
)

type mockPOSA struct {
	consensus.FastFinalityPoSA
}

// testBackend is a mock implementation of the live Ethereum message handler.
type testBackend struct {
	eventMux *event.TypeMux
}

func newTestBackend() *testBackend {
	return &testBackend{eventMux: new(event.TypeMux)}
}
func (b *testBackend) IsMining() bool           { return true }
func (b *testBackend) EventMux() *event.TypeMux { return b.eventMux }

func (p *mockPOSA) GetJustifiedBlock(chain consensus.ChainHeaderReader, blockNumber uint64, blockHash common.Hash) (uint64, common.Hash) {
	return 0, common.Hash{}
}

func (m *mockPOSA) VerifyVote(chain consensus.ChainHeaderReader, vote *types.VoteEnvelope) error {
	return nil
}

func (m *mockPOSA) IsFinalityVoterAt(chain consensus.ChainHeaderReader, header *types.Header) bool {
	return true
}

func (pool *VotePool) verifyStructureSizeOfVotePool(curVotes, futureVotes, curVotesPq, futureVotesPq int) bool {
	for i := 0; i < timeThreshold; i++ {
		time.Sleep(1 * time.Second)
		poolCurVotes, poolCurVotesPq, poolFutureVotes, poolFutureVotesPq := pool.stats()
		if poolCurVotes == curVotes && poolFutureVotes == futureVotes && poolCurVotesPq == curVotesPq && poolFutureVotesPq == futureVotesPq {
			return true
		}
	}
	return false
}

func TestValidVotePool(t *testing.T) {
	testVotePool(t, true)
}

func TestInvalidVotePool(t *testing.T) {
	testVotePool(t, false)
}

func testVotePool(t *testing.T, isValidRules bool) {
	walletPasswordDir, walletDir := setUpKeyManager(t)

	// Create a database pre-initialize with a genesis block
	db := rawdb.NewMemoryDatabase()
	genesis := (&core.Genesis{
		Config:  params.TestChainConfig,
		Alloc:   core.GenesisAlloc{testAddr: {Balance: big.NewInt(1000000)}},
		BaseFee: big.NewInt(params.InitialBaseFee),
	}).MustCommit(db)
	chain, _ := core.NewBlockChain(db, nil, params.TestChainConfig, ethash.NewFullFaker(), vm.Config{}, nil, nil)

	mux := new(event.TypeMux)
	mockEngine := &mockPOSA{}

	// Create vote pool
	votePool := NewVotePool(chain, mockEngine, 22)

	// Create vote manager
	// Create a temporary file for the votes journal
	file, err := ioutil.TempFile("", "")
	if err != nil {
		t.Fatalf("failed to create temporary file path: %v", err)
	}
	journal := file.Name()
	defer os.Remove(journal)

	// Clean up the temporary file, we only need the path for now
	file.Close()
	os.Remove(journal)

	var (
		voteManager *VoteManager
	)
	if isValidRules {
		voteManager, err = NewVoteManager(newTestBackend(), db, params.TestChainConfig, chain, votePool, true, walletPasswordDir, walletDir, mockEngine, nil)
	} else {
		voteManager, err = NewVoteManager(newTestBackend(), db, params.TestChainConfig, chain, votePool, true, walletPasswordDir, walletDir, mockEngine, &Debug{ValidateRule: func(header *types.Header) error {
			return errors.New("mock error")
		}})
	}

	if err != nil {
		t.Fatalf("failed to create vote managers")
	}

	// Send the done event of downloader
	time.Sleep(10 * time.Millisecond)
	mux.Post(downloader.DoneEvent{})

	bs, _ := core.GenerateChain(params.TestChainConfig, genesis, ethash.NewFaker(), db, 1, nil, true)
	if _, err := chain.InsertChain(bs, nil); err != nil {
		panic(err)
	}
	for i := 0; i < 10; i++ {
		bs, _ = core.GenerateChain(params.TestChainConfig, bs[len(bs)-1], ethash.NewFaker(), db, 1, nil, true)
		if _, err := chain.InsertChain(bs, nil); err != nil {
			panic(err)
		}
	}

	if !isValidRules {
		if votePool.verifyStructureSizeOfVotePool(11, 0, 11, 0) {
			t.Fatalf("put vote failed")
		}
		return
	}

	if !votePool.verifyStructureSizeOfVotePool(11, 0, 11, 0) {
		t.Fatalf("put vote failed")
	}

	// Verify if votesPq is min heap
	votesPq := votePool.curVotesPq
	pqBuffer := make([]*types.VoteData, 0)
	lastVotedBlockNumber := uint64(0)
	for votesPq.Len() > 0 {
		voteData := heap.Pop(votesPq).(*types.VoteData)
		if voteData.TargetNumber < lastVotedBlockNumber {
			t.Fatalf("votesPq verification failed")
		}
		lastVotedBlockNumber = voteData.TargetNumber
		pqBuffer = append(pqBuffer, voteData)
	}
	for _, voteData := range pqBuffer {
		heap.Push(votesPq, voteData)
	}

	bs, _ = core.GenerateChain(params.TestChainConfig, bs[len(bs)-1], ethash.NewFaker(), db, 1, nil, true)
	if _, err := chain.InsertChain(bs, nil); err != nil {
		panic(err)
	}

	if !votePool.verifyStructureSizeOfVotePool(12, 0, 12, 0) {
		t.Fatalf("put vote failed")
	}

	for i := 0; i < 256; i++ {
		bs, _ = core.GenerateChain(params.TestChainConfig, bs[len(bs)-1], ethash.NewFaker(), db, 1, nil, true)
		if _, err := chain.InsertChain(bs, nil); err != nil {
			panic(err)
		}
	}

	// currently chain size is 268, and votePool should be pruned, so vote pool size should be 256!
	if !votePool.verifyStructureSizeOfVotePool(256, 0, 256, 0) {
		t.Fatalf("put vote failed")
	}

	// Test invalid vote whose number larger than latestHeader + 13
	invalidVote := &types.VoteEnvelope{
		RawVoteEnvelope: types.RawVoteEnvelope{
			Data: &types.VoteData{
				TargetNumber: 1000,
			},
		},
	}
	voteManager.pool.PutVote("", invalidVote)

	if !votePool.verifyStructureSizeOfVotePool(256, 0, 256, 0) {
		t.Fatalf("put vote failed")
	}

	votes := votePool.GetVotes()
	if len(votes) != 256 {
		t.Fatalf("get votes failed")
	}

	// Test future votes scenario: votes number within latestBlockHeader ~ latestBlockHeader + 13
	futureVote := &types.VoteEnvelope{
		RawVoteEnvelope: types.RawVoteEnvelope{
			Data: &types.VoteData{
				TargetNumber: 279,
			},
		},
	}
	if err := voteManager.signer.SignVote(futureVote); err != nil {
		t.Fatalf("sign vote failed")
	}
	voteManager.pool.PutVote("", futureVote)

	if !votePool.verifyStructureSizeOfVotePool(256, 1, 256, 1) {
		t.Fatalf("put vote failed")
	}

	// Test duplicate vote case, shouldn'd be put into vote pool
	duplicateVote := &types.VoteEnvelope{
		RawVoteEnvelope: types.RawVoteEnvelope{
			Data: &types.VoteData{
				TargetNumber: 279,
			},
		},
	}
	if err := voteManager.signer.SignVote(duplicateVote); err != nil {
		t.Fatalf("sign vote failed")
	}
	voteManager.pool.PutVote("", duplicateVote)

	if !votePool.verifyStructureSizeOfVotePool(256, 1, 256, 1) {
		t.Fatalf("put vote failed")
	}

	// Test future votes larger than latestBlockNumber + 13 should be rejected
	futureVote = &types.VoteEnvelope{
		RawVoteEnvelope: types.RawVoteEnvelope{
			Data: &types.VoteData{
				TargetNumber: 282,
				TargetHash:   common.Hash{},
			},
		},
	}
	voteManager.pool.PutVote("", futureVote)
	if !votePool.verifyStructureSizeOfVotePool(256, 1, 256, 1) {
		t.Fatalf("put vote failed")
	}

	// Test transfer votes from future to cur, latest block header is #288 after the following generation
	// For the above BlockNumber 279, it did not have blockHash, should be assigned as well below.
	curNumber := 268
	var futureBlockHash common.Hash
	for i := 0; i < 20; i++ {
		bs, _ = core.GenerateChain(params.TestChainConfig, bs[len(bs)-1], ethash.NewFaker(), db, 1, nil, true)
		curNumber += 1
		if curNumber == 279 {
			futureBlockHash = bs[0].Hash()
			futureVotesMap := votePool.futureVotes
			voteBox := futureVotesMap[common.Hash{}]
			futureVotesMap[futureBlockHash] = voteBox
			delete(futureVotesMap, common.Hash{})
			futureVotesPq := votePool.futureVotesPq
			futureVotesPq.Peek().TargetHash = futureBlockHash
		}
		if _, err := chain.InsertChain(bs, nil); err != nil {
			panic(err)
		}
	}

	done := false
	for i := 0; i < timeThreshold; i++ {
		time.Sleep(1 * time.Second)
		if len(votePool.FetchVoteByBlockHash(futureBlockHash)) == 2 {
			done = true
			break
		}
	}
	if !done {
		t.Fatalf("transfer vote failed")
	}

	// Pruner will keep the size of votePool as latestBlockHeader-255~latestBlockHeader, then final result should be 256!
	if !votePool.verifyStructureSizeOfVotePool(256, 0, 256, 0) {
		t.Fatalf("put vote failed")
	}

	for i := 0; i < 224; i++ {
		bs, _ = core.GenerateChain(params.TestChainConfig, bs[len(bs)-1], ethash.NewFaker(), db, 1, nil, true)
		if _, err := chain.InsertChain(bs, nil); err != nil {
			panic(err)
		}
	}

	bs, _ = core.GenerateChain(params.TestChainConfig, bs[len(bs)-1], ethash.NewFaker(), db, 1, nil, true)
	if _, err := chain.InsertChain(bs, nil); err != nil {
		panic(err)
	}
}

func setUpKeyManager(t *testing.T) (string, string) {
	walletDir := filepath.Join(t.TempDir(), "wallet")
	walletPasswordDir := filepath.Join(t.TempDir(), "password")
	if err := os.MkdirAll(filepath.Dir(walletPasswordDir), 0700); err != nil {
		t.Fatalf("failed to create walletPassword dir: %v", err)
	}
	if err := ioutil.WriteFile(walletPasswordDir, []byte(password), 0600); err != nil {
		t.Fatalf("failed to write wallet password dir: %v", err)
	}
	if err := os.MkdirAll(walletDir, 0700); err != nil {
		t.Fatalf("failed to create wallet dir: %v", err)
	}
	w, err := wallet.New(walletDir, walletPasswordDir)
	if err != nil {
		t.Fatalf("failed to create wallet: %v", err)
	}
	km, _ := wallet.NewKeyManager(context.Background(), w)
	secretKey, _ := bls.RandKey()
	encryptor := keystorev4.New()
	pubKeyBytes := secretKey.PublicKey().Marshal()
	cryptoFields, err := encryptor.Encrypt(secretKey.Marshal(), password)
	if err != nil {
		t.Fatalf("failed: %v", err)
	}

	id, _ := uuid.NewRandom()
	keystore := &wallet.Keystore{
		Crypto:  cryptoFields,
		ID:      id.String(),
		Pubkey:  fmt.Sprintf("%x", pubKeyBytes),
		Version: encryptor.Version(),
		Name:    encryptor.Name(),
	}

	encodedFile, _ := json.MarshalIndent(keystore, "", "\t")
	keyStoreDir := filepath.Join(t.TempDir(), "keystore")
	keystoreFile, _ := os.Create(fmt.Sprintf("%s/keystore-%s.json", keyStoreDir, "publichh"))
	keystoreFile.Write(encodedFile)
	km.ImportKeystores(context.Background(), []*wallet.Keystore{keystore}, []string{password})
	return walletPasswordDir, walletDir
}

func generateVote(
	blockNumber int,
	blockHash common.Hash,
	secretKey blsCommon.SecretKey,
) *types.VoteEnvelope {
	voteData := types.VoteData{
		TargetNumber: uint64(blockNumber),
		TargetHash:   blockHash,
	}
	digest := voteData.Hash()
	signature := secretKey.Sign(digest[:])

	vote := &types.VoteEnvelope{
		RawVoteEnvelope: types.RawVoteEnvelope{
			PublicKey: types.BLSPublicKey(secretKey.PublicKey().Marshal()),
			Signature: types.BLSSignature(signature.Marshal()),
			Data:      &voteData,
		},
	}

	return vote
}

func TestVotePoolDosProtection(t *testing.T) {
	secretKey, err := bls.RandKey()
	if err != nil {
		t.Fatalf("Failed to create secret key, err %s", err)
	}

	// Create a database pre-initialize with a genesis block
	db := rawdb.NewMemoryDatabase()
	genesis := (&core.Genesis{
		Config:  params.TestChainConfig,
		Alloc:   core.GenesisAlloc{testAddr: {Balance: big.NewInt(1000000)}},
		BaseFee: big.NewInt(params.InitialBaseFee),
	}).MustCommit(db)
	chain, _ := core.NewBlockChain(db, nil, params.TestChainConfig, ethash.NewFullFaker(), vm.Config{}, nil, nil)

	bs, _ := core.GenerateChain(params.TestChainConfig, genesis, ethash.NewFaker(), db, 25, nil, true)
	if _, err := chain.InsertChain(bs[:1], nil); err != nil {
		panic(err)
	}
	mockEngine := &mockPOSA{}

	// Create vote pool
	votePool := NewVotePool(chain, mockEngine, 22)

	for i := 0; i < maxFutureVotePerPeer; i++ {
		vote := generateVote(1, common.BigToHash(big.NewInt(int64(i+1))), secretKey)
		votePool.PutVote("AAAA", vote)
		time.Sleep(100 * time.Millisecond)
	}

	_, _, _, futureVoteQueueLength := votePool.stats()
	if futureVoteQueueLength != maxFutureVotePerPeer {
		t.Fatalf("Future vote pool length, expect %d have %d", maxFutureVotePerPeer, futureVoteQueueLength)
	}
	numFutureVotePerPeer := votePool.getNumberOfFutureVoteByPeer("AAAA")
	if numFutureVotePerPeer != maxFutureVotePerPeer {
		t.Fatalf("Number of future vote per peer, expect %d have %d", maxFutureVotePerPeer, numFutureVotePerPeer)
	}

	// This vote is dropped due to DOS protection
	vote := generateVote(1, common.BigToHash(big.NewInt(int64(maxFutureVoteAmountPerBlock+1))), secretKey)
	votePool.PutVote("AAAA", vote)
	time.Sleep(100 * time.Millisecond)
	_, _, _, futureVoteQueueLength = votePool.stats()
	if futureVoteQueueLength != maxFutureVotePerPeer {
		t.Fatalf("Future vote pool length, expect %d have %d", maxFutureVotePerPeer, futureVoteQueueLength)
	}
	numFutureVotePerPeer = votePool.getNumberOfFutureVoteByPeer("AAAA")
	if numFutureVotePerPeer != maxFutureVotePerPeer {
		t.Fatalf("Number of future vote per peer, expect %d have %d", maxFutureVotePerPeer, numFutureVotePerPeer)
	}

	// Vote from different peer must be accepted
	vote = generateVote(1, common.BigToHash(big.NewInt(int64(maxFutureVoteAmountPerBlock+2))), secretKey)
	votePool.PutVote("BBBB", vote)
	time.Sleep(100 * time.Millisecond)
	_, _, _, futureVoteQueueLength = votePool.stats()
	if futureVoteQueueLength != maxFutureVotePerPeer+1 {
		t.Fatalf("Future vote pool length, expect %d have %d", maxFutureVotePerPeer, futureVoteQueueLength)
	}
	numFutureVotePerPeer = votePool.getNumberOfFutureVoteByPeer("AAAA")
	if numFutureVotePerPeer != maxFutureVotePerPeer {
		t.Fatalf("Number of future vote per peer, expect %d have %d", maxFutureVotePerPeer, numFutureVotePerPeer)
	}
	numFutureVotePerPeer = votePool.getNumberOfFutureVoteByPeer("BBBB")
	if numFutureVotePerPeer != 1 {
		t.Fatalf("Number of future vote per peer, expect %d have %d", 1, numFutureVotePerPeer)
	}

	// One vote is not queued twice
	votePool.PutVote("CCCC", vote)
	time.Sleep(100 * time.Millisecond)
	_, _, _, futureVoteQueueLength = votePool.stats()
	if futureVoteQueueLength != maxFutureVotePerPeer+1 {
		t.Fatalf("Future vote pool length, expect %d have %d", maxFutureVotePerPeer, futureVoteQueueLength)
	}
	numFutureVotePerPeer = votePool.getNumberOfFutureVoteByPeer("CCCC")
	if numFutureVotePerPeer != 0 {
		t.Fatalf("Number of future vote per peer, expect %d have %d", 0, numFutureVotePerPeer)
	}

	if _, err := chain.InsertChain(bs[1:], nil); err != nil {
		panic(err)
	}
	time.Sleep(100 * time.Millisecond)
	// Future vote must be transferred to current and failed the verification,
	// numFutureVotePerPeer decreases
	_, _, _, futureVoteQueueLength = votePool.stats()
	if futureVoteQueueLength != 0 {
		t.Fatalf("Future vote pool length, expect %d have %d", 0, futureVoteQueueLength)
	}
	numFutureVotePerPeer = votePool.getNumberOfFutureVoteByPeer("AAAA")
	if numFutureVotePerPeer != 0 {
		t.Fatalf("Number of future vote per peer, expect %d have %d", 0, numFutureVotePerPeer)
	}
}

type mockPOSAv2 struct {
	consensus.FastFinalityPoSA
}

func (p *mockPOSAv2) GetJustifiedNumberAndHash(chain consensus.ChainHeaderReader, header *types.Header) (uint64, common.Hash, error) {
	parentHeader := chain.GetHeaderByHash(header.ParentHash)
	if parentHeader == nil {
		return 0, common.Hash{}, fmt.Errorf("unexpected error")
	}
	return parentHeader.Number.Uint64(), parentHeader.Hash(), nil
}

func (m *mockPOSAv2) VerifyVote(chain consensus.ChainHeaderReader, vote *types.VoteEnvelope) error {
	header := chain.GetHeaderByHash(vote.Data.TargetHash)
	if header == nil {
		return errors.New("header not found")
	}

	if header.Number.Uint64() != vote.Data.TargetNumber {
		return errors.New("wrong target number in vote")
	}

	return nil
}

func (m *mockPOSAv2) IsFinalityVoterAt(chain consensus.ChainHeaderReader, header *types.Header) bool {
	return true
}

func TestVotePoolWrongTargetNumber(t *testing.T) {
	secretKey, err := bls.RandKey()
	if err != nil {
		t.Fatalf("Failed to create secret key, err %s", err)
	}

	// Create a database pre-initialize with a genesis block
	db := rawdb.NewMemoryDatabase()
	genesis := (&core.Genesis{
		Config:  params.TestChainConfig,
		Alloc:   core.GenesisAlloc{testAddr: {Balance: big.NewInt(1000000)}},
		BaseFee: big.NewInt(params.InitialBaseFee),
	}).MustCommit(db)
	chain, _ := core.NewBlockChain(db, nil, params.TestChainConfig, ethash.NewFullFaker(), vm.Config{}, nil, nil)

	bs, _ := core.GenerateChain(params.TestChainConfig, genesis, ethash.NewFaker(), db, 1, nil, true)
	if _, err := chain.InsertChain(bs[:1], nil); err != nil {
		panic(err)
	}
	mockEngine := &mockPOSAv2{}

	// Create vote pool
	votePool := NewVotePool(chain, mockEngine, 22)

	// bs[0] is the block 1 so the target block number must be 1.
	// Here we provide wrong target number 0
	vote := generateVote(0, bs[0].Hash(), secretKey)
	votePool.PutVote("AAAA", vote)
	time.Sleep(100 * time.Millisecond)

	if len(votePool.curVotes) != 0 {
		t.Fatalf("Current vote length, expect %d have %d", 0, len(votePool.curVotes))
	}
}
