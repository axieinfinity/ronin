package types

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/crypto"
)

type FullGiftTicket struct {
	GiftTicket
	V, R, S *big.Int // Signature provided by payer
}

type GiftTicket struct {
	Nonce          *big.Int
	Payer          common.Address   // The address of payer for the gift ticket
	Allowance      *big.Int         // Max amount the payer pays for gas fee
	Recipients     []common.Address // The recipient account that can use gift ticket, empty array means everyone can use the gift ticket
	ExpirationTime *big.Int         // The gift ticket's expiration timestamp
	MaxUse         *big.Int         // The gift ticket's max use time
}

func (g *GiftTicket) Hash() (common.Hash, error) {
	recipients := make([]interface{}, len(g.Recipients))
	for i, recipient := range g.Recipients {
		recipients[i] = recipient.Hex()
	}
	typedData := TypedData{
		Types: Types{
			"EIP712Domain": []Type{
				{Name: "name", Type: "string"},
				{Name: "version", Type: "string"},
				{Name: "chainId", Type: "uint256"},
			},
			"GiftTicket": []Type{
				{Name: "nonce", Type: "uint256"},
				{Name: "payer", Type: "address"},
				{Name: "allowance", Type: "uint256"},
				{Name: "recipients", Type: "address[]"},
				{Name: "expirationTime", Type: "uint256"},
				{Name: "maxUse", Type: "uint256"},
			},
		},
		Domain: TypedDataDomain{
			Name:    "TransactionPass",
			Version: "1",
			ChainId: math.NewHexOrDecimal256(2021),
		},
		PrimaryType: "GiftTicket",
		Message: TypedDataMessage{
			"nonce":          g.Nonce.String(),
			"payer":          g.Payer.String(),
			"allowance":      g.Allowance.String(),
			"recipients":     recipients,
			"expirationTime": g.ExpirationTime.String(),
			"maxUse":         g.MaxUse.String(),
		},
	}
	domainSeparator, err := typedData.HashStruct("EIP712Domain", typedData.Domain.Map())
	if err != nil {
		return common.Hash{}, err
	}
	typedDataHash, err := typedData.HashStruct(typedData.PrimaryType, typedData.Message)
	if err != nil {
		return common.Hash{}, err
	}
	rawData := []byte(fmt.Sprintf("\x19\x01%s%s", string(domainSeparator), string(typedDataHash)))
	return common.BytesToHash(crypto.Keccak256(rawData)), nil
}

func (g *FullGiftTicket) copy() *FullGiftTicket {
	if g == nil {
		return nil
	}
	cpy := &FullGiftTicket{
		GiftTicket: GiftTicket{
			Nonce:          new(big.Int),
			Allowance:      new(big.Int),
			Recipients:     make([]common.Address, len(g.Recipients)),
			ExpirationTime: new(big.Int),
			MaxUse:         new(big.Int),
		},
		V: new(big.Int),
		R: new(big.Int),
		S: new(big.Int),
	}
	copy(cpy.Payer[:], g.Payer[:])
	if g.Nonce != nil {
		cpy.Nonce.Set(g.Nonce)
	}
	if g.Allowance != nil {
		cpy.Allowance.Set(g.Allowance)
	}
	copy(cpy.Recipients, g.Recipients)
	if g.ExpirationTime != nil {
		cpy.ExpirationTime.Set(g.ExpirationTime)
	}
	if g.MaxUse != nil {
		cpy.MaxUse.Set(g.MaxUse)
	}
	if g.V != nil {
		cpy.V.Set(g.V)
	}
	if g.R != nil {
		cpy.R.Set(g.R)
	}
	if g.S != nil {
		cpy.S.Set(g.S)
	}
	return cpy
}

type TransactionPassTx struct {
	*LegacyTx
	FullGiftTicket FullGiftTicket
}

func (tx *TransactionPassTx) copy() TxData {
	cpy := &TransactionPassTx{
		LegacyTx:       tx.LegacyTx.copy().(*LegacyTx),
		FullGiftTicket: *tx.FullGiftTicket.copy(),
	}
	return cpy
}

func (tx *TransactionPassTx) txType() byte { return TransactionPassTxType }
