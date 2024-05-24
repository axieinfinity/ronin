package state

import (
	"encoding/binary"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func GetLocSimpleVariable(slot uint64) common.Hash {
	slotHash := common.BigToHash(new(big.Int).SetUint64(slot))
	return slotHash
}

func GetLocMappingAtKey(key common.Hash, slot uint64) common.Hash {
	var buffer []byte

	buffer = key.Bytes()
	// Write 8-byte slot to 32-byte space in big endian order.
	// First write 24 0-bytes then write 8-slot in big endian.
	buffer = common.PadTo(buffer, len(buffer)+24)
	buffer = binary.BigEndian.AppendUint64(buffer, slot)
	return crypto.Keccak256Hash(buffer)
}

func GetLocDynamicArrAtElement(slotHash common.Hash, index uint64, elementSize uint64) common.Hash {
	slotKecBig := crypto.Keccak256Hash(slotHash.Bytes()).Big()
	//arrBig = slotKecBig + index * elementSize
	arrBig := slotKecBig.Add(slotKecBig, new(big.Int).SetUint64(index*elementSize))
	return common.BigToHash(arrBig)
}

func GetLocFixedArrAtElement(slot uint64, index uint64, elementSize uint64) common.Hash {
	slotBig := new(big.Int).SetUint64(slot)
	arrBig := slotBig.Add(slotBig, new(big.Int).SetUint64(index*elementSize))
	return common.BigToHash(arrBig)
}
