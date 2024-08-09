package types

import "github.com/ethereum/go-ethereum/common"

type BlobSidecars []*BlobSidecar

type BlobSidecar struct {
	BlobTxSidecar
	TxHash common.Hash
}

func NewBlobSidecarFromTx(tx *Transaction) *BlobSidecar {
	if tx.BlobTxSidecar() == nil {
		return nil
	}
	return &BlobSidecar{
		BlobTxSidecar: *tx.BlobTxSidecar(),
		TxHash:        tx.Hash(),
	}
}
