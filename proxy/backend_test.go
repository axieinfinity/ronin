package proxy

import (
	"github.com/ethereum/go-ethereum/eth/ethconfig"
	"github.com/ethereum/go-ethereum/ethdb/httpdb"
	"testing"
	"time"
)

func TestBackend_getLatestBlockNumber(t *testing.T) {
	url := "https://api.roninchain.com/rpc"
	db := httpdb.NewDB(url, 5*time.Second, 5*time.Second)
	backend, err := newBackend(db, &ethconfig.Config{}, url)
	if err != nil {
		t.Fatal(err)
	}
	num, err := backend.getLatestBlockNumber()
	if err != nil {
		t.Fatal(err)
	}
	if num <= 0 {
		t.Fatal("invalid block number")
	}
}
