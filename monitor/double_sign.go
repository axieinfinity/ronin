package monitor

import (
	"bytes"
	"encoding/hex"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
	lru "github.com/hashicorp/golang-lru"
)

const monitorBlockRange = 20

type DoubleSignMonitor struct {
	observerdBlocks *lru.Cache
}

func NewDoubleSignMonitor() (*DoubleSignMonitor, error) {
	observerdBlocks, err := lru.New(monitorBlockRange)
	if err != nil {
		return nil, err
	}
	monitor := DoubleSignMonitor{
		observerdBlocks: observerdBlocks,
	}

	return &monitor, nil
}

func getSignature(blockHeader *types.Header) string {
	signature := blockHeader.Extra[len(blockHeader.Extra)-crypto.SignatureLength:]
	return "0x" + hex.EncodeToString(signature)
}

func (monitor *DoubleSignMonitor) CheckDoubleSign(blockHeader *types.Header) {
	if rawBlockHeader, ok := monitor.observerdBlocks.Get(blockHeader.ParentHash); ok {
		blockHeaders, _ := rawBlockHeader.([]*types.Header)
		for _, header := range blockHeaders {
			// Simple check for monitoring only
			if !bytes.Equal(header.Hash().Bytes(), blockHeader.Hash().Bytes()) &&
				bytes.Equal(header.Coinbase[:], blockHeader.Coinbase[:]) {
				log.Error("Double sign detected", "block number", header.Number, "signer", header.Coinbase,
					"block 1 hash", header.Hash().Hex(), "block 1 signature", getSignature(header),
					"block 2 hash", blockHeader.Hash().Hex(), "block 2 signature", getSignature(blockHeader),
				)
				break
			}
		}
	} else {
		blockHeaders := []*types.Header{blockHeader}
		monitor.observerdBlocks.Add(blockHeader.ParentHash, blockHeaders)
	}
}
