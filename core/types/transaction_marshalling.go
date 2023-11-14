// Copyright 2021 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package types

import (
	"encoding/json"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// txJSON is the JSON representation of transactions.
type txJSON struct {
	Type hexutil.Uint64 `json:"type"`

	// Common transaction fields:
	Nonce                *hexutil.Uint64 `json:"nonce"`
	GasPrice             *hexutil.Big    `json:"gasPrice"`
	MaxPriorityFeePerGas *hexutil.Big    `json:"maxPriorityFeePerGas"`
	MaxFeePerGas         *hexutil.Big    `json:"maxFeePerGas"`
	Gas                  *hexutil.Uint64 `json:"gas"`
	Value                *hexutil.Big    `json:"value"`
	Data                 *hexutil.Bytes  `json:"input"`
	V                    *hexutil.Big    `json:"v"`
	R                    *hexutil.Big    `json:"r"`
	S                    *hexutil.Big    `json:"s"`
	To                   *common.Address `json:"to"`

	// Access list transaction fields:
	ChainID    *hexutil.Big `json:"chainId,omitempty"`
	AccessList *AccessList  `json:"accessList,omitempty"`

	// Sponsored transaction fields
	ExpiredTime *hexutil.Uint64 `json:"expiredTime,omitempty"`
	PayerV      *hexutil.Big    `json:"payerV,omitempty"`
	PayerR      *hexutil.Big    `json:"payerR,omitempty"`
	PayerS      *hexutil.Big    `json:"payerS,omitempty"`

	// Only used for encoding:
	Hash common.Hash `json:"hash"`
}

// MarshalJSON marshals as JSON with a hash.
func (t *Transaction) MarshalJSON() ([]byte, error) {
	var enc txJSON
	// These are set for all tx types.
	enc.Hash = t.Hash()
	enc.Type = hexutil.Uint64(t.Type())
	nonce := t.Nonce()
	gas := t.Gas()
	data := t.Data()
	v, r, s := t.RawSignatureValues()
	enc.Nonce = (*hexutil.Uint64)(&nonce)
	enc.Gas = (*hexutil.Uint64)(&gas)
	enc.Value = (*hexutil.Big)(t.Value())
	enc.Data = (*hexutil.Bytes)(&data)
	enc.To = t.To()
	enc.V = (*hexutil.Big)(v)
	enc.R = (*hexutil.Big)(r)
	enc.S = (*hexutil.Big)(s)

	// Other fields are set conditionally depending on tx type.
	switch tx := t.inner.(type) {
	case *LegacyTx:
		enc.GasPrice = (*hexutil.Big)(tx.GasPrice)
	case *AccessListTx:
		enc.GasPrice = (*hexutil.Big)(tx.GasPrice)
		enc.ChainID = (*hexutil.Big)(tx.ChainID)
		enc.AccessList = &tx.AccessList
	case *DynamicFeeTx:
		enc.ChainID = (*hexutil.Big)(tx.ChainID)
		enc.AccessList = &tx.AccessList
		enc.MaxFeePerGas = (*hexutil.Big)(tx.GasFeeCap)
		enc.MaxPriorityFeePerGas = (*hexutil.Big)(tx.GasTipCap)
	case *SponsoredTx:
		enc.ChainID = (*hexutil.Big)(tx.ChainID)
		enc.MaxFeePerGas = (*hexutil.Big)(tx.GasFeeCap)
		enc.MaxPriorityFeePerGas = (*hexutil.Big)(tx.GasTipCap)
		enc.ExpiredTime = (*hexutil.Uint64)(&tx.ExpiredTime)
		enc.PayerV = (*hexutil.Big)(tx.PayerV)
		enc.PayerR = (*hexutil.Big)(tx.PayerR)
		enc.PayerS = (*hexutil.Big)(tx.PayerS)
	}
	return json.Marshal(&enc)
}

// UnmarshalJSON unmarshals from JSON.
func (t *Transaction) UnmarshalJSON(input []byte) error {
	var dec txJSON
	if err := json.Unmarshal(input, &dec); err != nil {
		return err
	}

	var to *common.Address
	if dec.To != nil {
		to = dec.To
	}
	if dec.Nonce == nil {
		return errors.New("missing required field 'nonce' in transaction")
	}
	nonce := uint64(*dec.Nonce)
	if dec.Gas == nil {
		return errors.New("missing required field 'gas' in transaction")
	}
	gas := uint64(*dec.Gas)
	if dec.Value == nil {
		return errors.New("missing required field 'value' in transaction")
	}
	value := (*big.Int)(dec.Value)
	if dec.Data == nil {
		return errors.New("missing required field 'input' in transaction")
	}
	data := *dec.Data
	if dec.V == nil {
		return errors.New("missing required field 'v' in transaction")
	}
	v := (*big.Int)(dec.V)
	if dec.R == nil {
		return errors.New("missing required field 'r' in transaction")
	}
	r := (*big.Int)(dec.R)
	if dec.S == nil {
		return errors.New("missing required field 's' in transaction")
	}
	s := (*big.Int)(dec.S)
	withSignature := v.Sign() != 0 || r.Sign() != 0 || s.Sign() != 0
	if withSignature {
		maybeProtected := false
		if dec.Type == LegacyTxType {
			maybeProtected = true
		}

		if err := sanityCheckSignature(v, r, s, maybeProtected); err != nil {
			return err
		}
	}

	// Decode / verify fields according to transaction type.
	var inner TxData
	switch dec.Type {
	case LegacyTxType:
		itx := LegacyTx{
			Nonce: nonce,
			Gas:   gas,
			To:    to,
			Value: value,
			Data:  data,
			V:     v,
			R:     r,
			S:     s,
		}
		inner = &itx
		if dec.GasPrice == nil {
			return errors.New("missing required field 'gasPrice' in transaction")
		}
		itx.GasPrice = (*big.Int)(dec.GasPrice)

	case AccessListTxType:
		itx := AccessListTx{
			Nonce: nonce,
			Gas:   gas,
			To:    to,
			Value: value,
			Data:  data,
			V:     v,
			R:     r,
			S:     s,
		}
		inner = &itx
		// Access list is optional for now.
		if dec.AccessList != nil {
			itx.AccessList = *dec.AccessList
		}
		if dec.ChainID == nil {
			return errors.New("missing required field 'chainId' in transaction")
		}
		itx.ChainID = (*big.Int)(dec.ChainID)
		if dec.GasPrice == nil {
			return errors.New("missing required field 'gasPrice' in transaction")
		}
		itx.GasPrice = (*big.Int)(dec.GasPrice)

	case DynamicFeeTxType:
		itx := DynamicFeeTx{
			Nonce: nonce,
			Gas:   gas,
			To:    to,
			Value: value,
			Data:  data,
			V:     v,
			R:     r,
			S:     s,
		}
		inner = &itx
		// Access list is optional for now.
		if dec.AccessList != nil {
			itx.AccessList = *dec.AccessList
		}
		if dec.ChainID == nil {
			return errors.New("missing required field 'chainId' in transaction")
		}
		itx.ChainID = (*big.Int)(dec.ChainID)
		if dec.MaxPriorityFeePerGas == nil {
			return errors.New("missing required field 'maxPriorityFeePerGas' for txdata")
		}
		itx.GasTipCap = (*big.Int)(dec.MaxPriorityFeePerGas)
		if dec.MaxFeePerGas == nil {
			return errors.New("missing required field 'maxFeePerGas' for txdata")
		}
		itx.GasFeeCap = (*big.Int)(dec.MaxFeePerGas)

	case SponsoredTxType:
		itx := SponsoredTx{
			Nonce: nonce,
			Gas:   gas,
			To:    to,
			Value: value,
			Data:  data,
			V:     v,
			R:     r,
			S:     s,
		}
		inner = &itx
		if dec.ChainID == nil {
			return errors.New("missing required field 'chainId' in transaction")
		}
		itx.ChainID = (*big.Int)(dec.ChainID)
		if dec.MaxPriorityFeePerGas == nil {
			return errors.New("missing required field 'maxPriorityFeePerGas' for txdata")
		}
		itx.GasTipCap = (*big.Int)(dec.MaxPriorityFeePerGas)
		if dec.MaxFeePerGas == nil {
			return errors.New("missing required field 'maxFeePerGas' for txdata")
		}
		itx.GasFeeCap = (*big.Int)(dec.MaxFeePerGas)
		if dec.ExpiredTime == nil {
			return errors.New("missing required field 'expiredTime' in transaction")
		}
		itx.ExpiredTime = uint64(*dec.ExpiredTime)
		if dec.PayerV == nil {
			return errors.New("missing required field 'payerV' in transaction")
		}
		itx.PayerV = (*big.Int)(dec.PayerV)
		if dec.PayerR == nil {
			return errors.New("missing required field 'payerR' in transaction")
		}
		itx.PayerR = (*big.Int)(dec.PayerR)
		if dec.PayerS == nil {
			return errors.New("missing required field 'payerS' in transaction")
		}
		itx.PayerS = (*big.Int)(dec.PayerS)
		if err := sanityCheckSignature(itx.PayerV, itx.PayerR, itx.PayerS, false); err != nil {
			return err
		}

	default:
		return ErrTxTypeNotSupported
	}

	// Now set the inner transaction.
	t.setDecoded(inner, 0)

	// TODO: check hash here?
	return nil
}
