// Copyright 2014 The go-ethereum Authors
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

// bitvec is a bit vector which maps bytes in a program.
// An unset bit means the byte is an opcode, a set bit means
// it's data (i.e. argument of PUSHxx).
type bitvec []byte

func (bits bitvec) setN(flag uint16, pos uint64) {
	a := flag >> (pos % 8)
	bits[pos/8] |= byte(a >> 8)
	bits[pos/8+1] = byte(a)
}

func (bits bitvec) set8(pos uint64) {
	a := byte(0xFF >> (pos % 8))
	bits[pos/8] |= a
	bits[pos/8+1] = ^a
}

func (bits bitvec) set16(pos uint64) {
	a := byte(0xFF >> (pos % 8))
	bits[pos/8] |= a
	bits[pos/8+1] = 0xFF
	bits[pos/8+2] = ^a
}

// codeSegment checks if the position is in a code segment.
func (bits *bitvec) codeSegment(pos uint64) bool {
	return ((*bits)[pos/8] & (0x80 >> (pos % 8))) == 0
}

// codeBitmap collects data locations in code.
func codeBitmap(code []byte) bitvec {
	// The bitmap is 4 bytes longer than necessary, in case the code
	// ends with a PUSH32, the algorithm will push zeroes onto the
	// bitvector outside the bounds of the actual code.
	bits := make(bitvec, len(code)/8+1+5)
	return codeBitmapInternal(code, bits)
}

// codeBitmapInternal is the internal implementation of codeBitmap.
// It exists for the purpose of being able to run benchmark tests
// without dynamic allocations affecting the results.
func codeBitmapInternal(code, bits bitvec) bitvec {
	for pc := uint64(0); pc < uint64(len(code)); {
		op := OpCode(code[pc])
		pc++
		if op < PUSH1 || op > PUSH32 {
			continue
		}
		numbits := op - PUSH1 + 1
		if numbits >= 8 {
			for ; numbits >= 16; numbits -= 16 {
				bits.set16(pc)
				pc += 16
			}
			for ; numbits >= 8; numbits -= 8 {
				bits.set8(pc)
				pc += 8
			}
		}

		// create the mask based on numbits
		// set2BitsMask = uint16(0b1100_0000_0000_0000)
		// set3BitsMask = uint16(0b1110_0000_0000_0000)
		// set4BitsMask = uint16(0b1111_0000_0000_0000)
		// set5BitsMask = uint16(0b1111_1000_0000_0000)
		// set6BitsMask = uint16(0b1111_1100_0000_0000)
		// set7BitsMask = uint16(0b1111_1110_0000_0000)
		var mask uint16 = (1 << numbits) - 1
		mask = mask << (16 - numbits)
		bits.setN(mask, pc)
		pc += uint64(numbits)
	}
	return bits
}
