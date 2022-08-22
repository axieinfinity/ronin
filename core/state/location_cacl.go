package state

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func GetLocSimpleVariable(slot *big.Int) common.Hash {
	slotHash := common.BigToHash(slot)
	return slotHash
}

func GetLocMappingAtKey(key common.Hash, slot *big.Int) common.Hash {
	slotHash := common.BigToHash(slot)
	retByte := crypto.Keccak256(key.Bytes(), slotHash.Bytes())
	ret := new(big.Int)
	ret.SetBytes(retByte)
	return common.BigToHash(ret)
}

func GetLocDynamicArrAtElement(slotHash common.Hash, index uint64, elementSize uint64) common.Hash {
	slotKecBig := crypto.Keccak256Hash(slotHash.Bytes()).Big()
	//arrBig = slotKecBig + index * elementSize
	arrBig := slotKecBig.Add(slotKecBig, new(big.Int).SetUint64(index*elementSize))
	return common.BigToHash(arrBig)
}

func GetLocFixedArrAtElement(slot *big.Int, index uint64, elementSize uint64) common.Hash {
	arrBig := slot.Add(slot, new(big.Int).SetUint64(index*elementSize))
	return common.BigToHash(arrBig)
}
