// Copyright 2015 The go-ethereum Authors
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

package vm

import (
	"fmt"

	"github.com/holiman/uint256"
)

// Memory implements a simple memory model for the ethereum virtual machine.
type Memory struct {
	store       []byte
	lastGasCost uint64
}

// NewMemory returns a new memory model.
func NewMemory() *Memory {
	return &Memory{}
}

// Set sets offset + size to value
func (m *Memory) Set(offset, size uint64, value []byte) {
	// It's possible the offset is greater than 0 and size equals 0. This is because
	// the calcMemSize (common.go) could potentially return 0 when size is zero (NO-OP)
	if size > 0 {
		// length of store may never be less than offset + size.
		// The store should be resized PRIOR to setting the memory
		if offset+size > uint64(len(m.store)) {
			panic("invalid memory: store empty")
		}
		copy(m.store[offset:offset+size], value)
	}
}

// Set32 sets the 32 bytes starting at offset to the value of val, left-padded with zeroes to
// 32 bytes.
func (m *Memory) Set32(offset uint64, val *uint256.Int) {
	// length of store may never be less than offset + size.
	// The store should be resized PRIOR to setting the memory
	if offset+32 > uint64(len(m.store)) {
		panic("invalid memory: store empty")
	}
	// Fill in relevant bits
	fastWriteToArray32((*[32]byte)(m.store[offset:offset+32]), val)
}

// fastWriteToArray32 is the same as WriteToArray32 in uint256 package
// but with the loop unrolling manually for reducing branch instructions
func fastWriteToArray32(dest *[32]byte, val *uint256.Int) {
	// Unroll this loop manually
	//
	// for i := 0; i < 32; i++ {
	//     dest[31-i] = byte(val[i/8] >> uint64(8*(i%8)))
	// }
	dest[31] = byte(val[0] >> uint64(0))
	dest[30] = byte(val[0] >> uint64(8))
	dest[29] = byte(val[0] >> uint64(16))
	dest[28] = byte(val[0] >> uint64(24))
	dest[27] = byte(val[0] >> uint64(32))
	dest[26] = byte(val[0] >> uint64(40))
	dest[25] = byte(val[0] >> uint64(48))
	dest[24] = byte(val[0] >> uint64(56))
	dest[23] = byte(val[1] >> uint64(0))
	dest[22] = byte(val[1] >> uint64(8))
	dest[21] = byte(val[1] >> uint64(16))
	dest[20] = byte(val[1] >> uint64(24))
	dest[19] = byte(val[1] >> uint64(32))
	dest[18] = byte(val[1] >> uint64(40))
	dest[17] = byte(val[1] >> uint64(48))
	dest[16] = byte(val[1] >> uint64(56))
	dest[15] = byte(val[2] >> uint64(0))
	dest[14] = byte(val[2] >> uint64(8))
	dest[13] = byte(val[2] >> uint64(16))
	dest[12] = byte(val[2] >> uint64(24))
	dest[11] = byte(val[2] >> uint64(32))
	dest[10] = byte(val[2] >> uint64(40))
	dest[9] = byte(val[2] >> uint64(48))
	dest[8] = byte(val[2] >> uint64(56))
	dest[7] = byte(val[3] >> uint64(0))
	dest[6] = byte(val[3] >> uint64(8))
	dest[5] = byte(val[3] >> uint64(16))
	dest[4] = byte(val[3] >> uint64(24))
	dest[3] = byte(val[3] >> uint64(32))
	dest[2] = byte(val[3] >> uint64(40))
	dest[1] = byte(val[3] >> uint64(48))
	dest[0] = byte(val[3] >> uint64(56))
}

// Resize resizes the memory to size
func (m *Memory) Resize(size uint64) {
	if uint64(m.Len()) < size {
		m.store = append(m.store, make([]byte, size-uint64(m.Len()))...)
	}
}

// Get returns offset + size as a new slice
func (m *Memory) GetCopy(offset, size int64) (cpy []byte) {
	if size == 0 {
		return nil
	}

	if len(m.store) > int(offset) {
		cpy = make([]byte, size)
		copy(cpy, m.store[offset:offset+size])

		return
	}

	return
}

// GetPtr returns the offset + size
func (m *Memory) GetPtr(offset, size int64) []byte {
	if size == 0 {
		return nil
	}

	if len(m.store) > int(offset) {
		return m.store[offset : offset+size]
	}

	return nil
}

// Len returns the length of the backing slice
func (m *Memory) Len() int {
	return len(m.store)
}

// Data returns the backing slice
func (m *Memory) Data() []byte {
	return m.store
}

// Print dumps the content of the memory.
func (m *Memory) Print() {
	fmt.Printf("### mem %d bytes ###\n", len(m.store))
	if len(m.store) > 0 {
		addr := 0
		for i := 0; i+32 <= len(m.store); i += 32 {
			fmt.Printf("%03d: % x\n", addr, m.store[i:i+32])
			addr++
		}
	} else {
		fmt.Println("-- empty --")
	}
	fmt.Println("####################")
}
