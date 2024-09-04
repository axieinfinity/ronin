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
	"github.com/ethereum/go-ethereum/params"
)

type (
	executionWrapperFunc func(operation *operation, pc *uint64, interpreter *EVMInterpreter, callContext *ScopeContext) ([]byte, error)
	executionFunc        func(pc *uint64, interpreter *EVMInterpreter, callContext *ScopeContext) ([]byte, error)
	gasFunc              func(*EVM, *Contract, *Stack, *Memory, uint64) (uint64, error) // last parameter is the requested memory size as a uint64
	// memorySizeFunc returns the required size, and whether the operation overflowed a uint64
	memorySizeFunc func(*Stack) (size uint64, overflow bool)
)

type operation struct {
	// opcode name to get wrapper function
	opcodeName string
	// execute is the operation function
	executeWrapper executionWrapperFunc
	execute        executionFunc
	constantGas    uint64
	dynamicGas     gasFunc
	// minStack tells how many stack items are required
	minStack int
	// maxStack specifies the max length the stack can have for this operation
	// to not overflow the stack.
	maxStack int

	// memorySize returns the memory size required for the operation
	memorySize memorySizeFunc

	halts   bool // indicates whether the operation should halt further execution
	jumps   bool // indicates whether the program counter should not increment
	writes  bool // determines whether this a state modifying operation
	reverts bool // determines whether the operation reverts state (implicitly halts)
	returns bool // determines whether the operations sets the return data content
}

var (
	frontierInstructionSet         = newFrontierInstructionSet()
	homesteadInstructionSet        = newHomesteadInstructionSet()
	tangerineWhistleInstructionSet = newTangerineWhistleInstructionSet()
	spuriousDragonInstructionSet   = newSpuriousDragonInstructionSet()
	byzantiumInstructionSet        = newByzantiumInstructionSet()
	constantinopleInstructionSet   = newConstantinopleInstructionSet()
	istanbulInstructionSet         = newIstanbulInstructionSet()
	berlinInstructionSet           = newBerlinInstructionSet()
	londonInstructionSet           = newLondonInstructionSet()
	shanghaiInstructionSet         = newShanghaiInstructionSet()
)

// JumpTable contains the EVM opcodes supported at a given fork.
type JumpTable [256]*operation

func newShanghaiInstructionSet() JumpTable {
	instructionSet := newLondonInstructionSet()
	enable3855(&instructionSet) // PUSH0 instruction
	enable3860(&instructionSet) // Limit and meter initcode
	return instructionSet
}

// newLondonInstructionSet returns the frontier, homestead, byzantium,
// contantinople, istanbul, petersburg, berlin and london instructions.
func newLondonInstructionSet() JumpTable {
	instructionSet := newBerlinInstructionSet()
	enable3529(&instructionSet) // EIP-3529: Reduction in refunds https://eips.ethereum.org/EIPS/eip-3529
	enable3198(&instructionSet) // Base fee opcode https://eips.ethereum.org/EIPS/eip-3198
	for _, instruction := range instructionSet {
		if instruction != nil {
			instruction.executeWrapper = stringToWrapperLondon[instruction.opcodeName]
		}
	}
	return instructionSet
}

// newBerlinInstructionSet returns the frontier, homestead, byzantium,
// contantinople, istanbul, petersburg and berlin instructions.
func newBerlinInstructionSet() JumpTable {
	instructionSet := newIstanbulInstructionSet()
	enable2929(&instructionSet) // Access lists for trie accesses https://eips.ethereum.org/EIPS/eip-2929
	for _, instruction := range instructionSet {
		if instruction != nil {
			instruction.executeWrapper = stringToWrapperBerlin[instruction.opcodeName]
		}
	}
	return instructionSet
}

// newIstanbulInstructionSet returns the frontier, homestead, byzantium,
// contantinople, istanbul and petersburg instructions.
func newIstanbulInstructionSet() JumpTable {
	instructionSet := newConstantinopleInstructionSet()

	enable1344(&instructionSet) // ChainID opcode - https://eips.ethereum.org/EIPS/eip-1344
	enable1884(&instructionSet) // Reprice reader opcodes - https://eips.ethereum.org/EIPS/eip-1884
	enable2200(&instructionSet) // Net metered SSTORE - https://eips.ethereum.org/EIPS/eip-2200
	for _, instruction := range instructionSet {
		if instruction != nil {
			instruction.executeWrapper = stringToWrapperIstanbul[instruction.opcodeName]
		}
	}
	return instructionSet
}

// newConstantinopleInstructionSet returns the frontier, homestead,
// byzantium and contantinople instructions.
func newConstantinopleInstructionSet() JumpTable {
	instructionSet := newByzantiumInstructionSet()
	instructionSet[SHL] = &operation{
		opcodeName:    "opSHL",
		execute:        opSHL,
		constantGas:    GasFastestStep,
		minStack:       minStack(2, 1),
		maxStack:       maxStack(2, 1),
	}
	instructionSet[SHR] = &operation{
		opcodeName:    "opSHR",
		execute:        opSHR,
		constantGas:    GasFastestStep,
		minStack:       minStack(2, 1),
		maxStack:       maxStack(2, 1),
	}
	instructionSet[SAR] = &operation{
		opcodeName:    "opSAR",
		execute:        opSAR,
		constantGas:    GasFastestStep,
		minStack:       minStack(2, 1),
		maxStack:       maxStack(2, 1),
	}
	instructionSet[EXTCODEHASH] = &operation{
		opcodeName:    "opExtCodeHash",
		execute:        opExtCodeHash,
		constantGas:    params.ExtcodeHashGasConstantinople,
		minStack:       minStack(1, 1),
		maxStack:       maxStack(1, 1),
	}
	instructionSet[CREATE2] = &operation{
		opcodeName:    "opCreate2",
		execute:        opCreate2,
		constantGas:    params.Create2Gas,
		dynamicGas:     gasCreate2,
		minStack:       minStack(4, 1),
		maxStack:       maxStack(4, 1),
		memorySize:     memoryCreate2,
		writes:         true,
		returns:        true,
	}
	for _, instruction := range instructionSet {
		if instruction != nil {
			instruction.executeWrapper = stringToWrapperConstantinople[instruction.opcodeName]
		}
	}
	return instructionSet
}

// newByzantiumInstructionSet returns the frontier, homestead and
// byzantium instructions.
func newByzantiumInstructionSet() JumpTable {
	instructionSet := newSpuriousDragonInstructionSet()
	instructionSet[STATICCALL] = &operation{
		opcodeName:    "opStaticCall",
		execute:        opStaticCall,
		constantGas:    params.CallGasEIP150,
		dynamicGas:     gasStaticCall,
		minStack:       minStack(6, 1),
		maxStack:       maxStack(6, 1),
		memorySize:     memoryStaticCall,
		returns:        true,
	}
	instructionSet[RETURNDATASIZE] = &operation{
		opcodeName:    "opReturnDataSize",
		execute:        opReturnDataSize,
		constantGas:    GasQuickStep,
		minStack:       minStack(0, 1),
		maxStack:       maxStack(0, 1),
	}
	instructionSet[RETURNDATACOPY] = &operation{
		opcodeName:    "opReturnDataCopy",
		execute:        opReturnDataCopy,
		constantGas:    GasFastestStep,
		dynamicGas:     gasReturnDataCopy,
		minStack:       minStack(3, 0),
		maxStack:       maxStack(3, 0),
		memorySize:     memoryReturnDataCopy,
	}
	instructionSet[REVERT] = &operation{
		opcodeName:    "opRevert",
		execute:        opRevert,
		dynamicGas:     gasRevert,
		minStack:       minStack(2, 0),
		maxStack:       maxStack(2, 0),
		memorySize:     memoryRevert,
		reverts:        true,
		returns:        true,
	}
	for _, instruction := range instructionSet {
		if instruction != nil {
			instruction.executeWrapper = stringToWrapperByzantium[instruction.opcodeName]
		}
	}
	return instructionSet
}

// EIP 158 a.k.a Spurious Dragon
func newSpuriousDragonInstructionSet() JumpTable {
	instructionSet := newTangerineWhistleInstructionSet()
	instructionSet[EXP].dynamicGas = gasExpEIP158
	for _, instruction := range instructionSet {
		if instruction != nil {
			instruction.executeWrapper = stringToWrapperSpuriousDragon[instruction.opcodeName]
		}
	}
	return instructionSet

}

// EIP 150 a.k.a Tangerine Whistle
func newTangerineWhistleInstructionSet() JumpTable {
	instructionSet := newHomesteadInstructionSet()
	instructionSet[BALANCE].constantGas = params.BalanceGasEIP150
	instructionSet[EXTCODESIZE].constantGas = params.ExtcodeSizeGasEIP150
	instructionSet[SLOAD].constantGas = params.SloadGasEIP150
	instructionSet[EXTCODECOPY].constantGas = params.ExtcodeCopyBaseEIP150
	instructionSet[CALL].constantGas = params.CallGasEIP150
	instructionSet[CALLCODE].constantGas = params.CallGasEIP150
	instructionSet[DELEGATECALL].constantGas = params.CallGasEIP150
	for _, instruction := range instructionSet {
		if instruction != nil {
			instruction.executeWrapper = stringToWrapperTangerineWhistle[instruction.opcodeName]
		}
	}
	return instructionSet
}

// newHomesteadInstructionSet returns the frontier and homestead
// instructions that can be executed during the homestead phase.
func newHomesteadInstructionSet() JumpTable {
	instructionSet := newFrontierInstructionSet()
	instructionSet[DELEGATECALL] = &operation{
		opcodeName:    "opDelegateCall",
		execute:        opDelegateCall,
		dynamicGas:     gasDelegateCall,
		constantGas:    params.CallGasFrontier,
		minStack:       minStack(6, 1),
		maxStack:       maxStack(6, 1),
		memorySize:     memoryDelegateCall,
		returns:        true,
	}
	for _, instruction := range instructionSet {
		if instruction != nil {
			instruction.executeWrapper = stringToWrapperHomestead[instruction.opcodeName]
		}
	}
	return instructionSet
}

// newFrontierInstructionSet returns the frontier instructions
// that can be executed during the frontier phase.
func newFrontierInstructionSet() JumpTable {
	instructionSet := JumpTable{
		STOP: {
			opcodeName:    "opStop",
			execute:        opStop,
			constantGas:    0,
			minStack:       minStack(0, 0),
			maxStack:       maxStack(0, 0),
			halts:          true,
		},
		ADD: {
			opcodeName:    "opAdd",
			execute:        opAdd,
			constantGas:    GasFastestStep,
			minStack:       minStack(2, 1),
			maxStack:       maxStack(2, 1),
		},
		MUL: {
			opcodeName:    "opMul",
			execute:        opMul,
			constantGas:    GasFastStep,
			minStack:       minStack(2, 1),
			maxStack:       maxStack(2, 1),
		},
		SUB: {
			opcodeName:    "opSub",
			execute:        opSub,
			constantGas:    GasFastestStep,
			minStack:       minStack(2, 1),
			maxStack:       maxStack(2, 1),
		},
		DIV: {
			opcodeName:    "opDiv",
			execute:        opDiv,
			constantGas:    GasFastStep,
			minStack:       minStack(2, 1),
			maxStack:       maxStack(2, 1),
		},
		SDIV: {
			opcodeName:    "opSdiv",
			execute:        opSdiv,
			constantGas:    GasFastStep,
			minStack:       minStack(2, 1),
			maxStack:       maxStack(2, 1),
		},
		MOD: {
			opcodeName:    "opMod",
			execute:        opMod,
			constantGas:    GasFastStep,
			minStack:       minStack(2, 1),
			maxStack:       maxStack(2, 1),
		},
		SMOD: {
			opcodeName:    "opSmod",
			execute:        opSmod,
			constantGas:    GasFastStep,
			minStack:       minStack(2, 1),
			maxStack:       maxStack(2, 1),
		},
		ADDMOD: {
			opcodeName:    "opAddmod",
			execute:        opAddmod,
			constantGas:    GasMidStep,
			minStack:       minStack(3, 1),
			maxStack:       maxStack(3, 1),
		},
		MULMOD: {
			opcodeName:    "opMulmod",
			execute:        opMulmod,
			constantGas:    GasMidStep,
			minStack:       minStack(3, 1),
			maxStack:       maxStack(3, 1),
		},
		EXP: {
			opcodeName:    "opExp",
			execute:        opExp,
			dynamicGas:     gasExpFrontier,
			minStack:       minStack(2, 1),
			maxStack:       maxStack(2, 1),
		},
		SIGNEXTEND: {
			opcodeName:    "opSignExtend",
			execute:        opSignExtend,
			constantGas:    GasFastStep,
			minStack:       minStack(2, 1),
			maxStack:       maxStack(2, 1),
		},
		LT: {
			opcodeName:    "opLt",
			execute:        opLt,
			constantGas:    GasFastestStep,
			minStack:       minStack(2, 1),
			maxStack:       maxStack(2, 1),
		},
		GT: {
			opcodeName:    "opGt",
			execute:        opGt,
			constantGas:    GasFastestStep,
			minStack:       minStack(2, 1),
			maxStack:       maxStack(2, 1),
		},
		SLT: {
			opcodeName:    "opSlt",
			execute:        opSlt,
			constantGas:    GasFastestStep,
			minStack:       minStack(2, 1),
			maxStack:       maxStack(2, 1),
		},
		SGT: {
			opcodeName:    "opSgt",
			execute:        opSgt,
			constantGas:    GasFastestStep,
			minStack:       minStack(2, 1),
			maxStack:       maxStack(2, 1),
		},
		EQ: {
			opcodeName:    "opEq",
			execute:        opEq,
			constantGas:    GasFastestStep,
			minStack:       minStack(2, 1),
			maxStack:       maxStack(2, 1),
		},
		ISZERO: {
			opcodeName:    "opIszero",
			execute:        opIszero,
			constantGas:    GasFastestStep,
			minStack:       minStack(1, 1),
			maxStack:       maxStack(1, 1),
		},
		AND: {
			opcodeName:    "opAnd",
			execute:        opAnd,
			constantGas:    GasFastestStep,
			minStack:       minStack(2, 1),
			maxStack:       maxStack(2, 1),
		},
		XOR: {
			opcodeName:    "opXor",
			execute:        opXor,
			constantGas:    GasFastestStep,
			minStack:       minStack(2, 1),
			maxStack:       maxStack(2, 1),
		},
		OR: {
			opcodeName:    "opOr",
			execute:        opOr,
			constantGas:    GasFastestStep,
			minStack:       minStack(2, 1),
			maxStack:       maxStack(2, 1),
		},
		NOT: {
			opcodeName:    "opNot",
			execute:        opNot,
			constantGas:    GasFastestStep,
			minStack:       minStack(1, 1),
			maxStack:       maxStack(1, 1),
		},
		BYTE: {
			opcodeName:    "opByte",
			execute:        opByte,
			constantGas:    GasFastestStep,
			minStack:       minStack(2, 1),
			maxStack:       maxStack(2, 1),
		},
		SHA3: {
			opcodeName:    "opSha3",
			execute:        opSha3,
			constantGas:    params.Sha3Gas,
			dynamicGas:     gasSha3,
			minStack:       minStack(2, 1),
			maxStack:       maxStack(2, 1),
			memorySize:     memorySha3,
		},
		ADDRESS: {
			opcodeName:    "opAddress",
			execute:        opAddress,
			constantGas:    GasQuickStep,
			minStack:       minStack(0, 1),
			maxStack:       maxStack(0, 1),
		},
		BALANCE: {
			opcodeName:    "opBalance",
			execute:        opBalance,
			constantGas:    params.BalanceGasFrontier,
			minStack:       minStack(1, 1),
			maxStack:       maxStack(1, 1),
		},
		ORIGIN: {
			opcodeName:    "opOrigin",
			execute:        opOrigin,
			constantGas:    GasQuickStep,
			minStack:       minStack(0, 1),
			maxStack:       maxStack(0, 1),
		},
		CALLER: {
			opcodeName:    "opCaller",
			execute:        opCaller,
			constantGas:    GasQuickStep,
			minStack:       minStack(0, 1),
			maxStack:       maxStack(0, 1),
		},
		CALLVALUE: {
			opcodeName:    "opCallValue",
			execute:        opCallValue,
			constantGas:    GasQuickStep,
			minStack:       minStack(0, 1),
			maxStack:       maxStack(0, 1),
		},
		CALLDATALOAD: {
			opcodeName:    "opCallDataLoad",
			execute:        opCallDataLoad,
			constantGas:    GasFastestStep,
			minStack:       minStack(1, 1),
			maxStack:       maxStack(1, 1),
		},
		CALLDATASIZE: {
			opcodeName:    "opCallDataSize",
			execute:        opCallDataSize,
			constantGas:    GasQuickStep,
			minStack:       minStack(0, 1),
			maxStack:       maxStack(0, 1),
		},
		CALLDATACOPY: {
			opcodeName:    "opCallDataCopy",
			execute:        opCallDataCopy,
			constantGas:    GasFastestStep,
			dynamicGas:     gasCallDataCopy,
			minStack:       minStack(3, 0),
			maxStack:       maxStack(3, 0),
			memorySize:     memoryCallDataCopy,
		},
		CODESIZE: {
			opcodeName:    "opCodeSize",
			execute:        opCodeSize,
			constantGas:    GasQuickStep,
			minStack:       minStack(0, 1),
			maxStack:       maxStack(0, 1),
		},
		CODECOPY: {
			opcodeName:    "opCodeCopy",
			execute:        opCodeCopy,
			constantGas:    GasFastestStep,
			dynamicGas:     gasCodeCopy,
			minStack:       minStack(3, 0),
			maxStack:       maxStack(3, 0),
			memorySize:     memoryCodeCopy,
		},
		GASPRICE: {
			opcodeName:    "opGasprice",
			execute:        opGasprice,
			constantGas:    GasQuickStep,
			minStack:       minStack(0, 1),
			maxStack:       maxStack(0, 1),
		},
		EXTCODESIZE: {
			opcodeName:    "opExtCodeSize",
			execute:        opExtCodeSize,
			constantGas:    params.ExtcodeSizeGasFrontier,
			minStack:       minStack(1, 1),
			maxStack:       maxStack(1, 1),
		},
		EXTCODECOPY: {
			opcodeName:    "opExtCodeCopy",
			execute:        opExtCodeCopy,
			constantGas:    params.ExtcodeCopyBaseFrontier,
			dynamicGas:     gasExtCodeCopy,
			minStack:       minStack(4, 0),
			maxStack:       maxStack(4, 0),
			memorySize:     memoryExtCodeCopy,
		},
		BLOCKHASH: {
			opcodeName:    "opBlockhash",
			execute:        opBlockhash,
			constantGas:    GasExtStep,
			minStack:       minStack(1, 1),
			maxStack:       maxStack(1, 1),
		},
		COINBASE: {
			opcodeName:    "opCoinbase",
			execute:        opCoinbase,
			constantGas:    GasQuickStep,
			minStack:       minStack(0, 1),
			maxStack:       maxStack(0, 1),
		},
		TIMESTAMP: {
			opcodeName:    "opTimestamp",
			execute:        opTimestamp,
			constantGas:    GasQuickStep,
			minStack:       minStack(0, 1),
			maxStack:       maxStack(0, 1),
		},
		NUMBER: {
			opcodeName:    "opNumber",
			execute:        opNumber,
			constantGas:    GasQuickStep,
			minStack:       minStack(0, 1),
			maxStack:       maxStack(0, 1),
		},
		DIFFICULTY: {
			opcodeName:    "opDifficulty",
			execute:        opDifficulty,
			constantGas:    GasQuickStep,
			minStack:       minStack(0, 1),
			maxStack:       maxStack(0, 1),
		},
		GASLIMIT: {
			opcodeName:    "opGasLimit",
			execute:        opGasLimit,
			constantGas:    GasQuickStep,
			minStack:       minStack(0, 1),
			maxStack:       maxStack(0, 1),
		},
		POP: {
			opcodeName:    "opPop",
			execute:        opPop,
			constantGas:    GasQuickStep,
			minStack:       minStack(1, 0),
			maxStack:       maxStack(1, 0),
		},
		MLOAD: {
			opcodeName:    "opMload",
			execute:        opMload,
			constantGas:    GasFastestStep,
			dynamicGas:     gasMLoad,
			minStack:       minStack(1, 1),
			maxStack:       maxStack(1, 1),
			memorySize:     memoryMLoad,
		},
		MSTORE: {
			opcodeName:    "opMstore",
			execute:        opMstore,
			constantGas:    GasFastestStep,
			dynamicGas:     gasMStore,
			minStack:       minStack(2, 0),
			maxStack:       maxStack(2, 0),
			memorySize:     memoryMStore,
		},
		MSTORE8: {
			opcodeName:    "opMstore8",
			execute:        opMstore8,
			constantGas:    GasFastestStep,
			dynamicGas:     gasMStore8,
			memorySize:     memoryMStore8,
			minStack:       minStack(2, 0),
			maxStack:       maxStack(2, 0),
		},
		SLOAD: {
			opcodeName:    "opSload",
			execute:        opSload,
			constantGas:    params.SloadGasFrontier,
			minStack:       minStack(1, 1),
			maxStack:       maxStack(1, 1),
		},
		SSTORE: {
			opcodeName:    "opSstore",
			execute:        opSstore,
			dynamicGas:     gasSStore,
			minStack:       minStack(2, 0),
			maxStack:       maxStack(2, 0),
			writes:         true,
		},
		JUMP: {
			opcodeName:    "opJump",
			execute:        opJump,
			constantGas:    GasMidStep,
			minStack:       minStack(1, 0),
			maxStack:       maxStack(1, 0),
			jumps:          true,
		},
		JUMPI: {
			opcodeName:    "opJumpi",
			execute:        opJumpi,
			constantGas:    GasSlowStep,
			minStack:       minStack(2, 0),
			maxStack:       maxStack(2, 0),
			jumps:          true,
		},
		PC: {
			opcodeName:    "opPc",
			execute:        opPc,
			constantGas:    GasQuickStep,
			minStack:       minStack(0, 1),
			maxStack:       maxStack(0, 1),
		},
		MSIZE: {
			opcodeName:    "opMsize",
			execute:        opMsize,
			constantGas:    GasQuickStep,
			minStack:       minStack(0, 1),
			maxStack:       maxStack(0, 1),
		},
		GAS: {
			opcodeName:    "opGas",
			execute:        opGas,
			constantGas:    GasQuickStep,
			minStack:       minStack(0, 1),
			maxStack:       maxStack(0, 1),
		},
		JUMPDEST: {
			opcodeName:    "opJumpdest",
			execute:        opJumpdest,
			constantGas:    params.JumpdestGas,
			minStack:       minStack(0, 0),
			maxStack:       maxStack(0, 0),
		},
		PUSH1: {
			opcodeName:    "opPush1",
			execute:        opPush1,
			constantGas:    GasFastestStep,
			minStack:       minStack(0, 1),
			maxStack:       maxStack(0, 1),
		},
		PUSH2: {
			opcodeName:    "opPush",
			execute:        makePush(2, 2),
			constantGas:    GasFastestStep,
			minStack:       minStack(0, 1),
			maxStack:       maxStack(0, 1),
		},
		PUSH3: {
			opcodeName:    "opPush",
			execute:        makePush(3, 3),
			constantGas:    GasFastestStep,
			minStack:       minStack(0, 1),
			maxStack:       maxStack(0, 1),
		},
		PUSH4: {
			opcodeName:    "opPush",
			execute:        makePush(4, 4),
			constantGas:    GasFastestStep,
			minStack:       minStack(0, 1),
			maxStack:       maxStack(0, 1),
		},
		PUSH5: {
			opcodeName:    "opPush",
			execute:        makePush(5, 5),
			constantGas:    GasFastestStep,
			minStack:       minStack(0, 1),
			maxStack:       maxStack(0, 1),
		},
		PUSH6: {
			opcodeName:    "opPush",
			execute:        makePush(6, 6),
			constantGas:    GasFastestStep,
			minStack:       minStack(0, 1),
			maxStack:       maxStack(0, 1),
		},
		PUSH7: {
			opcodeName:    "opPush",
			execute:        makePush(7, 7),
			constantGas:    GasFastestStep,
			minStack:       minStack(0, 1),
			maxStack:       maxStack(0, 1),
		},
		PUSH8: {
			opcodeName:    "opPush",
			execute:        makePush(8, 8),
			constantGas:    GasFastestStep,
			minStack:       minStack(0, 1),
			maxStack:       maxStack(0, 1),
		},
		PUSH9: {
			opcodeName:    "opPush",
			execute:        makePush(9, 9),
			constantGas:    GasFastestStep,
			minStack:       minStack(0, 1),
			maxStack:       maxStack(0, 1),
		},
		PUSH10: {
			opcodeName:    "opPush",
			execute:        makePush(10, 10),
			constantGas:    GasFastestStep,
			minStack:       minStack(0, 1),
			maxStack:       maxStack(0, 1),
		},
		PUSH11: {
			opcodeName:    "opPush",
			execute:        makePush(11, 11),
			constantGas:    GasFastestStep,
			minStack:       minStack(0, 1),
			maxStack:       maxStack(0, 1),
		},
		PUSH12: {
			opcodeName:    "opPush",
			execute:        makePush(12, 12),
			constantGas:    GasFastestStep,
			minStack:       minStack(0, 1),
			maxStack:       maxStack(0, 1),
		},
		PUSH13: {
			opcodeName:    "opPush",
			execute:        makePush(13, 13),
			constantGas:    GasFastestStep,
			minStack:       minStack(0, 1),
			maxStack:       maxStack(0, 1),
		},
		PUSH14: {
			opcodeName:    "opPush",
			execute:        makePush(14, 14),
			constantGas:    GasFastestStep,
			minStack:       minStack(0, 1),
			maxStack:       maxStack(0, 1),
		},
		PUSH15: {
			opcodeName:    "opPush",
			execute:        makePush(15, 15),
			constantGas:    GasFastestStep,
			minStack:       minStack(0, 1),
			maxStack:       maxStack(0, 1),
		},
		PUSH16: {
			opcodeName:    "opPush",
			execute:        makePush(16, 16),
			constantGas:    GasFastestStep,
			minStack:       minStack(0, 1),
			maxStack:       maxStack(0, 1),
		},
		PUSH17: {
			opcodeName:    "opPush",
			execute:        makePush(17, 17),
			constantGas:    GasFastestStep,
			minStack:       minStack(0, 1),
			maxStack:       maxStack(0, 1),
		},
		PUSH18: {
			opcodeName:    "opPush",
			execute:        makePush(18, 18),
			constantGas:    GasFastestStep,
			minStack:       minStack(0, 1),
			maxStack:       maxStack(0, 1),
		},
		PUSH19: {
			opcodeName:    "opPush",
			execute:        makePush(19, 19),
			constantGas:    GasFastestStep,
			minStack:       minStack(0, 1),
			maxStack:       maxStack(0, 1),
		},
		PUSH20: {
			opcodeName:    "opPush",
			execute:        makePush(20, 20),
			constantGas:    GasFastestStep,
			minStack:       minStack(0, 1),
			maxStack:       maxStack(0, 1),
		},
		PUSH21: {
			opcodeName:    "opPush",
			execute:        makePush(21, 21),
			constantGas:    GasFastestStep,
			minStack:       minStack(0, 1),
			maxStack:       maxStack(0, 1),
		},
		PUSH22: {
			opcodeName:    "opPush",
			execute:        makePush(22, 22),
			constantGas:    GasFastestStep,
			minStack:       minStack(0, 1),
			maxStack:       maxStack(0, 1),
		},
		PUSH23: {
			opcodeName:    "opPush",
			execute:        makePush(23, 23),
			constantGas:    GasFastestStep,
			minStack:       minStack(0, 1),
			maxStack:       maxStack(0, 1),
		},
		PUSH24: {
			opcodeName:    "opPush",
			execute:        makePush(24, 24),
			constantGas:    GasFastestStep,
			minStack:       minStack(0, 1),
			maxStack:       maxStack(0, 1),
		},
		PUSH25: {
			opcodeName:    "opPush",
			execute:        makePush(25, 25),
			constantGas:    GasFastestStep,
			minStack:       minStack(0, 1),
			maxStack:       maxStack(0, 1),
		},
		PUSH26: {
			opcodeName:    "opPush",
			execute:        makePush(26, 26),
			constantGas:    GasFastestStep,
			minStack:       minStack(0, 1),
			maxStack:       maxStack(0, 1),
		},
		PUSH27: {
			opcodeName:    "opPush",
			execute:        makePush(27, 27),
			constantGas:    GasFastestStep,
			minStack:       minStack(0, 1),
			maxStack:       maxStack(0, 1),
		},
		PUSH28: {
			opcodeName:    "opPush",
			execute:        makePush(28, 28),
			constantGas:    GasFastestStep,
			minStack:       minStack(0, 1),
			maxStack:       maxStack(0, 1),
		},
		PUSH29: {
			opcodeName:    "opPush",
			execute:        makePush(29, 29),
			constantGas:    GasFastestStep,
			minStack:       minStack(0, 1),
			maxStack:       maxStack(0, 1),
		},
		PUSH30: {
			opcodeName:    "opPush",
			execute:        makePush(30, 30),
			constantGas:    GasFastestStep,
			minStack:       minStack(0, 1),
			maxStack:       maxStack(0, 1),
		},
		PUSH31: {
			opcodeName:    "opPush",
			execute:        makePush(31, 31),
			constantGas:    GasFastestStep,
			minStack:       minStack(0, 1),
			maxStack:       maxStack(0, 1),
		},
		PUSH32: {
			opcodeName:    "opPush",
			execute:        makePush(32, 32),
			constantGas:    GasFastestStep,
			minStack:       minStack(0, 1),
			maxStack:       maxStack(0, 1),
		},
		DUP1: {
			opcodeName:    "opDup",
			execute:        makeDup(1),
			constantGas:    GasFastestStep,
			minStack:       minDupStack(1),
			maxStack:       maxDupStack(1),
		},
		DUP2: {
			opcodeName:    "opDup",
			execute:        makeDup(2),
			constantGas:    GasFastestStep,
			minStack:       minDupStack(2),
			maxStack:       maxDupStack(2),
		},
		DUP3: {
			opcodeName:    "opDup",
			execute:        makeDup(3),
			constantGas:    GasFastestStep,
			minStack:       minDupStack(3),
			maxStack:       maxDupStack(3),
		},
		DUP4: {
			opcodeName:    "opDup",
			execute:        makeDup(4),
			constantGas:    GasFastestStep,
			minStack:       minDupStack(4),
			maxStack:       maxDupStack(4),
		},
		DUP5: {
			opcodeName:    "opDup",
			execute:        makeDup(5),
			constantGas:    GasFastestStep,
			minStack:       minDupStack(5),
			maxStack:       maxDupStack(5),
		},
		DUP6: {
			opcodeName:    "opDup",
			execute:        makeDup(6),
			constantGas:    GasFastestStep,
			minStack:       minDupStack(6),
			maxStack:       maxDupStack(6),
		},
		DUP7: {
			opcodeName:    "opDup",
			execute:        makeDup(7),
			constantGas:    GasFastestStep,
			minStack:       minDupStack(7),
			maxStack:       maxDupStack(7),
		},
		DUP8: {
			opcodeName:    "opDup",
			execute:        makeDup(8),
			constantGas:    GasFastestStep,
			minStack:       minDupStack(8),
			maxStack:       maxDupStack(8),
		},
		DUP9: {
			opcodeName:    "opDup",
			execute:        makeDup(9),
			constantGas:    GasFastestStep,
			minStack:       minDupStack(9),
			maxStack:       maxDupStack(9),
		},
		DUP10: {
			opcodeName:    "opDup",
			execute:        makeDup(10),
			constantGas:    GasFastestStep,
			minStack:       minDupStack(10),
			maxStack:       maxDupStack(10),
		},
		DUP11: {
			opcodeName:    "opDup",
			execute:        makeDup(11),
			constantGas:    GasFastestStep,
			minStack:       minDupStack(11),
			maxStack:       maxDupStack(11),
		},
		DUP12: {
			opcodeName:    "opDup",
			execute:        makeDup(12),
			constantGas:    GasFastestStep,
			minStack:       minDupStack(12),
			maxStack:       maxDupStack(12),
		},
		DUP13: {
			opcodeName:    "opDup",
			execute:        makeDup(13),
			constantGas:    GasFastestStep,
			minStack:       minDupStack(13),
			maxStack:       maxDupStack(13),
		},
		DUP14: {
			opcodeName:    "opDup",
			execute:        makeDup(14),
			constantGas:    GasFastestStep,
			minStack:       minDupStack(14),
			maxStack:       maxDupStack(14),
		},
		DUP15: {
			opcodeName:    "opDup",
			execute:        makeDup(15),
			constantGas:    GasFastestStep,
			minStack:       minDupStack(15),
			maxStack:       maxDupStack(15),
		},
		DUP16: {
			opcodeName:    "opDup",
			execute:        makeDup(16),
			constantGas:    GasFastestStep,
			minStack:       minDupStack(16),
			maxStack:       maxDupStack(16),
		},
		SWAP1: {
			opcodeName:    "opSwap",
			execute:        makeSwap(1),
			constantGas:    GasFastestStep,
			minStack:       minSwapStack(2),
			maxStack:       maxSwapStack(2),
		},
		SWAP2: {
			opcodeName:    "opSwap",
			execute:        makeSwap(2),
			constantGas:    GasFastestStep,
			minStack:       minSwapStack(3),
			maxStack:       maxSwapStack(3),
		},
		SWAP3: {
			opcodeName:    "opSwap",
			execute:        makeSwap(3),
			constantGas:    GasFastestStep,
			minStack:       minSwapStack(4),
			maxStack:       maxSwapStack(4),
		},
		SWAP4: {
			opcodeName:    "opSwap",
			execute:        makeSwap(4),
			constantGas:    GasFastestStep,
			minStack:       minSwapStack(5),
			maxStack:       maxSwapStack(5),
		},
		SWAP5: {
			opcodeName:    "opSwap",
			execute:        makeSwap(5),
			constantGas:    GasFastestStep,
			minStack:       minSwapStack(6),
			maxStack:       maxSwapStack(6),
		},
		SWAP6: {
			opcodeName:    "opSwap",
			execute:        makeSwap(6),
			constantGas:    GasFastestStep,
			minStack:       minSwapStack(7),
			maxStack:       maxSwapStack(7),
		},
		SWAP7: {
			opcodeName:    "opSwap",
			execute:        makeSwap(7),
			constantGas:    GasFastestStep,
			minStack:       minSwapStack(8),
			maxStack:       maxSwapStack(8),
		},
		SWAP8: {
			opcodeName:    "opSwap",
			execute:        makeSwap(8),
			constantGas:    GasFastestStep,
			minStack:       minSwapStack(9),
			maxStack:       maxSwapStack(9),
		},
		SWAP9: {
			opcodeName:    "opSwap",
			execute:        makeSwap(9),
			constantGas:    GasFastestStep,
			minStack:       minSwapStack(10),
			maxStack:       maxSwapStack(10),
		},
		SWAP10: {
			opcodeName:    "opSwap",
			execute:        makeSwap(10),
			constantGas:    GasFastestStep,
			minStack:       minSwapStack(11),
			maxStack:       maxSwapStack(11),
		},
		SWAP11: {
			opcodeName:    "opSwap",
			execute:        makeSwap(11),
			constantGas:    GasFastestStep,
			minStack:       minSwapStack(12),
			maxStack:       maxSwapStack(12),
		},
		SWAP12: {
			opcodeName:    "opSwap",
			execute:        makeSwap(12),
			constantGas:    GasFastestStep,
			minStack:       minSwapStack(13),
			maxStack:       maxSwapStack(13),
		},
		SWAP13: {
			opcodeName:    "opSwap",
			execute:        makeSwap(13),
			constantGas:    GasFastestStep,
			minStack:       minSwapStack(14),
			maxStack:       maxSwapStack(14),
		},
		SWAP14: {
			opcodeName:    "opSwap",
			execute:        makeSwap(14),
			constantGas:    GasFastestStep,
			minStack:       minSwapStack(15),
			maxStack:       maxSwapStack(15),
		},
		SWAP15: {
			opcodeName:    "opSwap",
			execute:        makeSwap(15),
			constantGas:    GasFastestStep,
			minStack:       minSwapStack(16),
			maxStack:       maxSwapStack(16),
		},
		SWAP16: {
			opcodeName:    "opSwap",
			execute:        makeSwap(16),
			constantGas:    GasFastestStep,
			minStack:       minSwapStack(17),
			maxStack:       maxSwapStack(17),
		},
		LOG0: {
			opcodeName:    "opSwap",
			execute:        makeLog(0),
			dynamicGas:     makeGasLog(0),
			minStack:       minStack(2, 0),
			maxStack:       maxStack(2, 0),
			memorySize:     memoryLog,
			writes:         true,
		},
		LOG1: {
			opcodeName:    "opLog",
			execute:        makeLog(1),
			dynamicGas:     makeGasLog(1),
			minStack:       minStack(3, 0),
			maxStack:       maxStack(3, 0),
			memorySize:     memoryLog,
			writes:         true,
		},
		LOG2: {
			opcodeName:    "opLog",
			execute:        makeLog(2),
			dynamicGas:     makeGasLog(2),
			minStack:       minStack(4, 0),
			maxStack:       maxStack(4, 0),
			memorySize:     memoryLog,
			writes:         true,
		},
		LOG3: {
			opcodeName:    "opLog",
			execute:        makeLog(3),
			dynamicGas:     makeGasLog(3),
			minStack:       minStack(5, 0),
			maxStack:       maxStack(5, 0),
			memorySize:     memoryLog,
			writes:         true,
		},
		LOG4: {
			opcodeName:    "opLog",
			execute:        makeLog(4),
			dynamicGas:     makeGasLog(4),
			minStack:       minStack(6, 0),
			maxStack:       maxStack(6, 0),
			memorySize:     memoryLog,
			writes:         true,
		},
		CREATE: {
			opcodeName:    "opCreate",
			execute:        opCreate,
			constantGas:    params.CreateGas,
			dynamicGas:     gasCreate,
			minStack:       minStack(3, 1),
			maxStack:       maxStack(3, 1),
			memorySize:     memoryCreate,
			writes:         true,
			returns:        true,
		},
		CALL: {
			opcodeName:    "opCall",
			execute:        opCall,
			constantGas:    params.CallGasFrontier,
			dynamicGas:     gasCall,
			minStack:       minStack(7, 1),
			maxStack:       maxStack(7, 1),
			memorySize:     memoryCall,
			returns:        true,
		},
		CALLCODE: {
			opcodeName:    "opCallCode",
			execute:        opCallCode,
			constantGas:    params.CallGasFrontier,
			dynamicGas:     gasCallCode,
			minStack:       minStack(7, 1),
			maxStack:       maxStack(7, 1),
			memorySize:     memoryCall,
			returns:        true,
		},
		RETURN: {
			opcodeName:    "opReturn",
			execute:        opReturn,
			dynamicGas:     gasReturn,
			minStack:       minStack(2, 0),
			maxStack:       maxStack(2, 0),
			memorySize:     memoryReturn,
			halts:          true,
		},
		SELFDESTRUCT: {
			opcodeName:    "opSuicide",
			execute:        opSuicide,
			dynamicGas:     gasSelfdestruct,
			minStack:       minStack(1, 0),
			maxStack:       maxStack(1, 0),
			halts:          true,
			writes:         true,
		},
	}
	for _, instruction := range instructionSet {
		if instruction != nil {
			instruction.executeWrapper = stringToWrapperFrontier[instruction.opcodeName]
		}
	}
	return instructionSet
}

// Export instruction set attributes

type OperationAttribute struct {
	Supported bool
	Name string
	MinStack int
	MaxStack int
	ConstantGas uint64
	DynamicCost bool
	MemorySize bool
	Halts bool
	Jumps bool
	Writes bool
	Reverts bool
	Returns bool
}

func ExportInstructionSet(rules params.Rules) []OperationAttribute{
	var instructionSet JumpTable
	switch {
	case rules.IsLondon:
		instructionSet = londonInstructionSet
	case rules.IsBerlin:
		instructionSet = berlinInstructionSet
	case rules.IsIstanbul:
		instructionSet = istanbulInstructionSet
	case rules.IsConstantinople:
		instructionSet = constantinopleInstructionSet
	case rules.IsByzantium:
		instructionSet = byzantiumInstructionSet
	case rules.IsEIP158:
		instructionSet = spuriousDragonInstructionSet
	case rules.IsEIP150:
		instructionSet = tangerineWhistleInstructionSet
	case rules.IsHomestead:
		instructionSet = homesteadInstructionSet
	default:
		instructionSet = frontierInstructionSet
	}
	// Eliminate duplicate operations (like PUSH, DUP,...), and to add unsupported operations
	existOperations := make(map[string]bool)
	var operations []OperationAttribute
	for _, op := range instructionSet {
		if op != nil {
			functionName := op.opcodeName
			if !existOperations[functionName] {
				existOperations[functionName] = true
				attribute := OperationAttribute{
					Supported: true,
					Name: functionName,
					MinStack: op.minStack,
					MaxStack: op.maxStack,
					ConstantGas: op.constantGas,
					DynamicCost: op.dynamicGas != nil,
					MemorySize: op.memorySize != nil,
					Halts: op.halts,
					Jumps: op.jumps,
					Writes: op.writes,
					Reverts: op.reverts,
					Returns: op.returns,
				}
				operations = append(operations, attribute)
			}
		}
	}
	// Also add unsupported operations for generating dummy wrappers
	for _, op := range londonInstructionSet {
		if op != nil {
			functionName := op.opcodeName
			if !existOperations[functionName] {
				existOperations[functionName] = true
				attribute := OperationAttribute{
					Supported: false,
					Name: functionName,
					MinStack: op.minStack,
					MaxStack: op.maxStack,
					ConstantGas: op.constantGas,
					DynamicCost: op.dynamicGas != nil,
					MemorySize: op.memorySize != nil,
					Halts: op.halts,
					Jumps: op.jumps,
					Writes: op.writes,
					Reverts: op.reverts,
					Returns: op.returns,
				}
				operations = append(operations, attribute)
			}
		}
	}
	return operations
}

func copyJumpTable(source *JumpTable) *JumpTable {
	dest := *source
	for i, op := range source {
		if op != nil {
			opCopy := *op
			dest[i] = &opCopy
		}
	}
	return &dest
}
