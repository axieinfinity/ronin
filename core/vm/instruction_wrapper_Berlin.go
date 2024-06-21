
package vm

import "github.com/ethereum/go-ethereum/common/math"


func opStopWrapperBerlin(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
    stack := callContext.Stack
    contract := callContext.Contract
    var err error
    var res []byte	

    in.evm.Context.Counter++

    if sLen := stack.len(); sLen < 0 {
        return nil, &ErrStackUnderflow{stackLen: sLen, required: 0}
    } else if sLen > 1024 {
        return nil, &ErrStackOverflow{stackLen: sLen, limit: 1024}
    }
    if !contract.UseGas(0) {
        return nil, ErrOutOfGas
    }
    res, err = operation.execute(pc, in, callContext)
    if err != nil {
        return nil, err
    }
    *pc++
    return res, nil
}

func opAddWrapperBerlin(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
    stack := callContext.Stack
    contract := callContext.Contract
    var err error
    var res []byte	

    in.evm.Context.Counter++

    if sLen := stack.len(); sLen < 2 {
        return nil, &ErrStackUnderflow{stackLen: sLen, required: 2}
    } else if sLen > 1025 {
        return nil, &ErrStackOverflow{stackLen: sLen, limit: 1025}
    }
    if !contract.UseGas(3) {
        return nil, ErrOutOfGas
    }
    res, err = operation.execute(pc, in, callContext)
    if err != nil {
        return nil, err
    }
    *pc++
    return res, nil
}

func opMulWrapperBerlin(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
    stack := callContext.Stack
    contract := callContext.Contract
    var err error
    var res []byte	

    in.evm.Context.Counter++

    if sLen := stack.len(); sLen < 2 {
        return nil, &ErrStackUnderflow{stackLen: sLen, required: 2}
    } else if sLen > 1025 {
        return nil, &ErrStackOverflow{stackLen: sLen, limit: 1025}
    }
    if !contract.UseGas(5) {
        return nil, ErrOutOfGas
    }
    res, err = operation.execute(pc, in, callContext)
    if err != nil {
        return nil, err
    }
    *pc++
    return res, nil
}

func opSubWrapperBerlin(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
    stack := callContext.Stack
    contract := callContext.Contract
    var err error
    var res []byte	

    in.evm.Context.Counter++

    if sLen := stack.len(); sLen < 2 {
        return nil, &ErrStackUnderflow{stackLen: sLen, required: 2}
    } else if sLen > 1025 {
        return nil, &ErrStackOverflow{stackLen: sLen, limit: 1025}
    }
    if !contract.UseGas(3) {
        return nil, ErrOutOfGas
    }
    res, err = operation.execute(pc, in, callContext)
    if err != nil {
        return nil, err
    }
    *pc++
    return res, nil
}

func opDivWrapperBerlin(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
    stack := callContext.Stack
    contract := callContext.Contract
    var err error
    var res []byte	

    in.evm.Context.Counter++

    if sLen := stack.len(); sLen < 2 {
        return nil, &ErrStackUnderflow{stackLen: sLen, required: 2}
    } else if sLen > 1025 {
        return nil, &ErrStackOverflow{stackLen: sLen, limit: 1025}
    }
    if !contract.UseGas(5) {
        return nil, ErrOutOfGas
    }
    res, err = operation.execute(pc, in, callContext)
    if err != nil {
        return nil, err
    }
    *pc++
    return res, nil
}

func opSdivWrapperBerlin(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
    stack := callContext.Stack
    contract := callContext.Contract
    var err error
    var res []byte	

    in.evm.Context.Counter++

    if sLen := stack.len(); sLen < 2 {
        return nil, &ErrStackUnderflow{stackLen: sLen, required: 2}
    } else if sLen > 1025 {
        return nil, &ErrStackOverflow{stackLen: sLen, limit: 1025}
    }
    if !contract.UseGas(5) {
        return nil, ErrOutOfGas
    }
    res, err = operation.execute(pc, in, callContext)
    if err != nil {
        return nil, err
    }
    *pc++
    return res, nil
}

func opModWrapperBerlin(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
    stack := callContext.Stack
    contract := callContext.Contract
    var err error
    var res []byte	

    in.evm.Context.Counter++

    if sLen := stack.len(); sLen < 2 {
        return nil, &ErrStackUnderflow{stackLen: sLen, required: 2}
    } else if sLen > 1025 {
        return nil, &ErrStackOverflow{stackLen: sLen, limit: 1025}
    }
    if !contract.UseGas(5) {
        return nil, ErrOutOfGas
    }
    res, err = operation.execute(pc, in, callContext)
    if err != nil {
        return nil, err
    }
    *pc++
    return res, nil
}

func opSmodWrapperBerlin(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
    stack := callContext.Stack
    contract := callContext.Contract
    var err error
    var res []byte	

    in.evm.Context.Counter++

    if sLen := stack.len(); sLen < 2 {
        return nil, &ErrStackUnderflow{stackLen: sLen, required: 2}
    } else if sLen > 1025 {
        return nil, &ErrStackOverflow{stackLen: sLen, limit: 1025}
    }
    if !contract.UseGas(5) {
        return nil, ErrOutOfGas
    }
    res, err = operation.execute(pc, in, callContext)
    if err != nil {
        return nil, err
    }
    *pc++
    return res, nil
}

func opAddmodWrapperBerlin(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
    stack := callContext.Stack
    contract := callContext.Contract
    var err error
    var res []byte	

    in.evm.Context.Counter++

    if sLen := stack.len(); sLen < 3 {
        return nil, &ErrStackUnderflow{stackLen: sLen, required: 3}
    } else if sLen > 1026 {
        return nil, &ErrStackOverflow{stackLen: sLen, limit: 1026}
    }
    if !contract.UseGas(8) {
        return nil, ErrOutOfGas
    }
    res, err = operation.execute(pc, in, callContext)
    if err != nil {
        return nil, err
    }
    *pc++
    return res, nil
}

func opMulmodWrapperBerlin(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
    stack := callContext.Stack
    contract := callContext.Contract
    var err error
    var res []byte	

    in.evm.Context.Counter++

    if sLen := stack.len(); sLen < 3 {
        return nil, &ErrStackUnderflow{stackLen: sLen, required: 3}
    } else if sLen > 1026 {
        return nil, &ErrStackOverflow{stackLen: sLen, limit: 1026}
    }
    if !contract.UseGas(8) {
        return nil, ErrOutOfGas
    }
    res, err = operation.execute(pc, in, callContext)
    if err != nil {
        return nil, err
    }
    *pc++
    return res, nil
}

func opExpWrapperBerlin(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
    mem := callContext.Memory
    stack := callContext.Stack
    contract := callContext.Contract
    var err error
    var res []byte	

    in.evm.Context.Counter++

    if sLen := stack.len(); sLen < 2 {
        return nil, &ErrStackUnderflow{stackLen: sLen, required: 2}
    } else if sLen > 1025 {
        return nil, &ErrStackOverflow{stackLen: sLen, limit: 1025}
    }
    var memorySize uint64
    var dynamicCost uint64
    dynamicCost, err = operation.dynamicGas(in.evm, contract, stack, mem, memorySize)
    if err != nil || !contract.UseGas(0 + dynamicCost) {
        return nil, ErrOutOfGas
    }
    if memorySize > 0 {
        mem.Resize(memorySize)
    }
    res, err = operation.execute(pc, in, callContext)
    if err != nil {
        return nil, err
    }
    *pc++
    return res, nil
}

func opSignExtendWrapperBerlin(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
    stack := callContext.Stack
    contract := callContext.Contract
    var err error
    var res []byte	

    in.evm.Context.Counter++

    if sLen := stack.len(); sLen < 2 {
        return nil, &ErrStackUnderflow{stackLen: sLen, required: 2}
    } else if sLen > 1025 {
        return nil, &ErrStackOverflow{stackLen: sLen, limit: 1025}
    }
    if !contract.UseGas(5) {
        return nil, ErrOutOfGas
    }
    res, err = operation.execute(pc, in, callContext)
    if err != nil {
        return nil, err
    }
    *pc++
    return res, nil
}

func opLtWrapperBerlin(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
    stack := callContext.Stack
    contract := callContext.Contract
    var err error
    var res []byte	

    in.evm.Context.Counter++

    if sLen := stack.len(); sLen < 2 {
        return nil, &ErrStackUnderflow{stackLen: sLen, required: 2}
    } else if sLen > 1025 {
        return nil, &ErrStackOverflow{stackLen: sLen, limit: 1025}
    }
    if !contract.UseGas(3) {
        return nil, ErrOutOfGas
    }
    res, err = operation.execute(pc, in, callContext)
    if err != nil {
        return nil, err
    }
    *pc++
    return res, nil
}

func opGtWrapperBerlin(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
    stack := callContext.Stack
    contract := callContext.Contract
    var err error
    var res []byte	

    in.evm.Context.Counter++

    if sLen := stack.len(); sLen < 2 {
        return nil, &ErrStackUnderflow{stackLen: sLen, required: 2}
    } else if sLen > 1025 {
        return nil, &ErrStackOverflow{stackLen: sLen, limit: 1025}
    }
    if !contract.UseGas(3) {
        return nil, ErrOutOfGas
    }
    res, err = operation.execute(pc, in, callContext)
    if err != nil {
        return nil, err
    }
    *pc++
    return res, nil
}

func opSltWrapperBerlin(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
    stack := callContext.Stack
    contract := callContext.Contract
    var err error
    var res []byte	

    in.evm.Context.Counter++

    if sLen := stack.len(); sLen < 2 {
        return nil, &ErrStackUnderflow{stackLen: sLen, required: 2}
    } else if sLen > 1025 {
        return nil, &ErrStackOverflow{stackLen: sLen, limit: 1025}
    }
    if !contract.UseGas(3) {
        return nil, ErrOutOfGas
    }
    res, err = operation.execute(pc, in, callContext)
    if err != nil {
        return nil, err
    }
    *pc++
    return res, nil
}

func opSgtWrapperBerlin(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
    stack := callContext.Stack
    contract := callContext.Contract
    var err error
    var res []byte	

    in.evm.Context.Counter++

    if sLen := stack.len(); sLen < 2 {
        return nil, &ErrStackUnderflow{stackLen: sLen, required: 2}
    } else if sLen > 1025 {
        return nil, &ErrStackOverflow{stackLen: sLen, limit: 1025}
    }
    if !contract.UseGas(3) {
        return nil, ErrOutOfGas
    }
    res, err = operation.execute(pc, in, callContext)
    if err != nil {
        return nil, err
    }
    *pc++
    return res, nil
}

func opEqWrapperBerlin(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
    stack := callContext.Stack
    contract := callContext.Contract
    var err error
    var res []byte	

    in.evm.Context.Counter++

    if sLen := stack.len(); sLen < 2 {
        return nil, &ErrStackUnderflow{stackLen: sLen, required: 2}
    } else if sLen > 1025 {
        return nil, &ErrStackOverflow{stackLen: sLen, limit: 1025}
    }
    if !contract.UseGas(3) {
        return nil, ErrOutOfGas
    }
    res, err = operation.execute(pc, in, callContext)
    if err != nil {
        return nil, err
    }
    *pc++
    return res, nil
}

func opIszeroWrapperBerlin(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
    stack := callContext.Stack
    contract := callContext.Contract
    var err error
    var res []byte	

    in.evm.Context.Counter++

    if sLen := stack.len(); sLen < 1 {
        return nil, &ErrStackUnderflow{stackLen: sLen, required: 1}
    } else if sLen > 1024 {
        return nil, &ErrStackOverflow{stackLen: sLen, limit: 1024}
    }
    if !contract.UseGas(3) {
        return nil, ErrOutOfGas
    }
    res, err = operation.execute(pc, in, callContext)
    if err != nil {
        return nil, err
    }
    *pc++
    return res, nil
}

func opAndWrapperBerlin(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
    stack := callContext.Stack
    contract := callContext.Contract
    var err error
    var res []byte	

    in.evm.Context.Counter++

    if sLen := stack.len(); sLen < 2 {
        return nil, &ErrStackUnderflow{stackLen: sLen, required: 2}
    } else if sLen > 1025 {
        return nil, &ErrStackOverflow{stackLen: sLen, limit: 1025}
    }
    if !contract.UseGas(3) {
        return nil, ErrOutOfGas
    }
    res, err = operation.execute(pc, in, callContext)
    if err != nil {
        return nil, err
    }
    *pc++
    return res, nil
}

func opOrWrapperBerlin(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
    stack := callContext.Stack
    contract := callContext.Contract
    var err error
    var res []byte	

    in.evm.Context.Counter++

    if sLen := stack.len(); sLen < 2 {
        return nil, &ErrStackUnderflow{stackLen: sLen, required: 2}
    } else if sLen > 1025 {
        return nil, &ErrStackOverflow{stackLen: sLen, limit: 1025}
    }
    if !contract.UseGas(3) {
        return nil, ErrOutOfGas
    }
    res, err = operation.execute(pc, in, callContext)
    if err != nil {
        return nil, err
    }
    *pc++
    return res, nil
}

func opXorWrapperBerlin(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
    stack := callContext.Stack
    contract := callContext.Contract
    var err error
    var res []byte	

    in.evm.Context.Counter++

    if sLen := stack.len(); sLen < 2 {
        return nil, &ErrStackUnderflow{stackLen: sLen, required: 2}
    } else if sLen > 1025 {
        return nil, &ErrStackOverflow{stackLen: sLen, limit: 1025}
    }
    if !contract.UseGas(3) {
        return nil, ErrOutOfGas
    }
    res, err = operation.execute(pc, in, callContext)
    if err != nil {
        return nil, err
    }
    *pc++
    return res, nil
}

func opNotWrapperBerlin(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
    stack := callContext.Stack
    contract := callContext.Contract
    var err error
    var res []byte	

    in.evm.Context.Counter++

    if sLen := stack.len(); sLen < 1 {
        return nil, &ErrStackUnderflow{stackLen: sLen, required: 1}
    } else if sLen > 1024 {
        return nil, &ErrStackOverflow{stackLen: sLen, limit: 1024}
    }
    if !contract.UseGas(3) {
        return nil, ErrOutOfGas
    }
    res, err = operation.execute(pc, in, callContext)
    if err != nil {
        return nil, err
    }
    *pc++
    return res, nil
}

func opByteWrapperBerlin(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
    stack := callContext.Stack
    contract := callContext.Contract
    var err error
    var res []byte	

    in.evm.Context.Counter++

    if sLen := stack.len(); sLen < 2 {
        return nil, &ErrStackUnderflow{stackLen: sLen, required: 2}
    } else if sLen > 1025 {
        return nil, &ErrStackOverflow{stackLen: sLen, limit: 1025}
    }
    if !contract.UseGas(3) {
        return nil, ErrOutOfGas
    }
    res, err = operation.execute(pc, in, callContext)
    if err != nil {
        return nil, err
    }
    *pc++
    return res, nil
}

func opSHLWrapperBerlin(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
    stack := callContext.Stack
    contract := callContext.Contract
    var err error
    var res []byte	

    in.evm.Context.Counter++

    if sLen := stack.len(); sLen < 2 {
        return nil, &ErrStackUnderflow{stackLen: sLen, required: 2}
    } else if sLen > 1025 {
        return nil, &ErrStackOverflow{stackLen: sLen, limit: 1025}
    }
    if !contract.UseGas(3) {
        return nil, ErrOutOfGas
    }
    res, err = operation.execute(pc, in, callContext)
    if err != nil {
        return nil, err
    }
    *pc++
    return res, nil
}

func opSHRWrapperBerlin(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
    stack := callContext.Stack
    contract := callContext.Contract
    var err error
    var res []byte	

    in.evm.Context.Counter++

    if sLen := stack.len(); sLen < 2 {
        return nil, &ErrStackUnderflow{stackLen: sLen, required: 2}
    } else if sLen > 1025 {
        return nil, &ErrStackOverflow{stackLen: sLen, limit: 1025}
    }
    if !contract.UseGas(3) {
        return nil, ErrOutOfGas
    }
    res, err = operation.execute(pc, in, callContext)
    if err != nil {
        return nil, err
    }
    *pc++
    return res, nil
}

func opSARWrapperBerlin(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
    stack := callContext.Stack
    contract := callContext.Contract
    var err error
    var res []byte	

    in.evm.Context.Counter++

    if sLen := stack.len(); sLen < 2 {
        return nil, &ErrStackUnderflow{stackLen: sLen, required: 2}
    } else if sLen > 1025 {
        return nil, &ErrStackOverflow{stackLen: sLen, limit: 1025}
    }
    if !contract.UseGas(3) {
        return nil, ErrOutOfGas
    }
    res, err = operation.execute(pc, in, callContext)
    if err != nil {
        return nil, err
    }
    *pc++
    return res, nil
}

func opSha3WrapperBerlin(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
    mem := callContext.Memory
    stack := callContext.Stack
    contract := callContext.Contract
    var err error
    var res []byte	

    in.evm.Context.Counter++

    if sLen := stack.len(); sLen < 2 {
        return nil, &ErrStackUnderflow{stackLen: sLen, required: 2}
    } else if sLen > 1025 {
        return nil, &ErrStackOverflow{stackLen: sLen, limit: 1025}
    }
    var memorySize uint64
    memSize, overflow := operation.memorySize(stack)
    if overflow {
        return nil, ErrGasUintOverflow
    }
    if memorySize, overflow = math.SafeMul(toWordSize(memSize), 32); overflow {
        return nil, ErrGasUintOverflow
    }
    var dynamicCost uint64
    dynamicCost, err = operation.dynamicGas(in.evm, contract, stack, mem, memorySize)
    if err != nil || !contract.UseGas(30 + dynamicCost) {
        return nil, ErrOutOfGas
    }
    if memorySize > 0 {
        mem.Resize(memorySize)
    }
    res, err = operation.execute(pc, in, callContext)
    if err != nil {
        return nil, err
    }
    *pc++
    return res, nil
}

func opAddressWrapperBerlin(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
    stack := callContext.Stack
    contract := callContext.Contract
    var err error
    var res []byte	

    in.evm.Context.Counter++

    if sLen := stack.len(); sLen < 0 {
        return nil, &ErrStackUnderflow{stackLen: sLen, required: 0}
    } else if sLen > 1023 {
        return nil, &ErrStackOverflow{stackLen: sLen, limit: 1023}
    }
    if !contract.UseGas(2) {
        return nil, ErrOutOfGas
    }
    res, err = operation.execute(pc, in, callContext)
    if err != nil {
        return nil, err
    }
    *pc++
    return res, nil
}

func opBalanceWrapperBerlin(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
    mem := callContext.Memory
    stack := callContext.Stack
    contract := callContext.Contract
    var err error
    var res []byte	

    in.evm.Context.Counter++

    if sLen := stack.len(); sLen < 1 {
        return nil, &ErrStackUnderflow{stackLen: sLen, required: 1}
    } else if sLen > 1024 {
        return nil, &ErrStackOverflow{stackLen: sLen, limit: 1024}
    }
    var memorySize uint64
    var dynamicCost uint64
    dynamicCost, err = operation.dynamicGas(in.evm, contract, stack, mem, memorySize)
    if err != nil || !contract.UseGas(100 + dynamicCost) {
        return nil, ErrOutOfGas
    }
    if memorySize > 0 {
        mem.Resize(memorySize)
    }
    res, err = operation.execute(pc, in, callContext)
    if err != nil {
        return nil, err
    }
    *pc++
    return res, nil
}

func opOriginWrapperBerlin(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
    stack := callContext.Stack
    contract := callContext.Contract
    var err error
    var res []byte	

    in.evm.Context.Counter++

    if sLen := stack.len(); sLen < 0 {
        return nil, &ErrStackUnderflow{stackLen: sLen, required: 0}
    } else if sLen > 1023 {
        return nil, &ErrStackOverflow{stackLen: sLen, limit: 1023}
    }
    if !contract.UseGas(2) {
        return nil, ErrOutOfGas
    }
    res, err = operation.execute(pc, in, callContext)
    if err != nil {
        return nil, err
    }
    *pc++
    return res, nil
}

func opCallerWrapperBerlin(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
    stack := callContext.Stack
    contract := callContext.Contract
    var err error
    var res []byte	

    in.evm.Context.Counter++

    if sLen := stack.len(); sLen < 0 {
        return nil, &ErrStackUnderflow{stackLen: sLen, required: 0}
    } else if sLen > 1023 {
        return nil, &ErrStackOverflow{stackLen: sLen, limit: 1023}
    }
    if !contract.UseGas(2) {
        return nil, ErrOutOfGas
    }
    res, err = operation.execute(pc, in, callContext)
    if err != nil {
        return nil, err
    }
    *pc++
    return res, nil
}

func opCallValueWrapperBerlin(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
    stack := callContext.Stack
    contract := callContext.Contract
    var err error
    var res []byte	

    in.evm.Context.Counter++

    if sLen := stack.len(); sLen < 0 {
        return nil, &ErrStackUnderflow{stackLen: sLen, required: 0}
    } else if sLen > 1023 {
        return nil, &ErrStackOverflow{stackLen: sLen, limit: 1023}
    }
    if !contract.UseGas(2) {
        return nil, ErrOutOfGas
    }
    res, err = operation.execute(pc, in, callContext)
    if err != nil {
        return nil, err
    }
    *pc++
    return res, nil
}

func opCallDataLoadWrapperBerlin(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
    stack := callContext.Stack
    contract := callContext.Contract
    var err error
    var res []byte	

    in.evm.Context.Counter++

    if sLen := stack.len(); sLen < 1 {
        return nil, &ErrStackUnderflow{stackLen: sLen, required: 1}
    } else if sLen > 1024 {
        return nil, &ErrStackOverflow{stackLen: sLen, limit: 1024}
    }
    if !contract.UseGas(3) {
        return nil, ErrOutOfGas
    }
    res, err = operation.execute(pc, in, callContext)
    if err != nil {
        return nil, err
    }
    *pc++
    return res, nil
}

func opCallDataSizeWrapperBerlin(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
    stack := callContext.Stack
    contract := callContext.Contract
    var err error
    var res []byte	

    in.evm.Context.Counter++

    if sLen := stack.len(); sLen < 0 {
        return nil, &ErrStackUnderflow{stackLen: sLen, required: 0}
    } else if sLen > 1023 {
        return nil, &ErrStackOverflow{stackLen: sLen, limit: 1023}
    }
    if !contract.UseGas(2) {
        return nil, ErrOutOfGas
    }
    res, err = operation.execute(pc, in, callContext)
    if err != nil {
        return nil, err
    }
    *pc++
    return res, nil
}

func opCallDataCopyWrapperBerlin(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
    mem := callContext.Memory
    stack := callContext.Stack
    contract := callContext.Contract
    var err error
    var res []byte	

    in.evm.Context.Counter++

    if sLen := stack.len(); sLen < 3 {
        return nil, &ErrStackUnderflow{stackLen: sLen, required: 3}
    } else if sLen > 1027 {
        return nil, &ErrStackOverflow{stackLen: sLen, limit: 1027}
    }
    var memorySize uint64
    memSize, overflow := operation.memorySize(stack)
    if overflow {
        return nil, ErrGasUintOverflow
    }
    if memorySize, overflow = math.SafeMul(toWordSize(memSize), 32); overflow {
        return nil, ErrGasUintOverflow
    }
    var dynamicCost uint64
    dynamicCost, err = operation.dynamicGas(in.evm, contract, stack, mem, memorySize)
    if err != nil || !contract.UseGas(3 + dynamicCost) {
        return nil, ErrOutOfGas
    }
    if memorySize > 0 {
        mem.Resize(memorySize)
    }
    res, err = operation.execute(pc, in, callContext)
    if err != nil {
        return nil, err
    }
    *pc++
    return res, nil
}

func opCodeSizeWrapperBerlin(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
    stack := callContext.Stack
    contract := callContext.Contract
    var err error
    var res []byte	

    in.evm.Context.Counter++

    if sLen := stack.len(); sLen < 0 {
        return nil, &ErrStackUnderflow{stackLen: sLen, required: 0}
    } else if sLen > 1023 {
        return nil, &ErrStackOverflow{stackLen: sLen, limit: 1023}
    }
    if !contract.UseGas(2) {
        return nil, ErrOutOfGas
    }
    res, err = operation.execute(pc, in, callContext)
    if err != nil {
        return nil, err
    }
    *pc++
    return res, nil
}

func opCodeCopyWrapperBerlin(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
    mem := callContext.Memory
    stack := callContext.Stack
    contract := callContext.Contract
    var err error
    var res []byte	

    in.evm.Context.Counter++

    if sLen := stack.len(); sLen < 3 {
        return nil, &ErrStackUnderflow{stackLen: sLen, required: 3}
    } else if sLen > 1027 {
        return nil, &ErrStackOverflow{stackLen: sLen, limit: 1027}
    }
    var memorySize uint64
    memSize, overflow := operation.memorySize(stack)
    if overflow {
        return nil, ErrGasUintOverflow
    }
    if memorySize, overflow = math.SafeMul(toWordSize(memSize), 32); overflow {
        return nil, ErrGasUintOverflow
    }
    var dynamicCost uint64
    dynamicCost, err = operation.dynamicGas(in.evm, contract, stack, mem, memorySize)
    if err != nil || !contract.UseGas(3 + dynamicCost) {
        return nil, ErrOutOfGas
    }
    if memorySize > 0 {
        mem.Resize(memorySize)
    }
    res, err = operation.execute(pc, in, callContext)
    if err != nil {
        return nil, err
    }
    *pc++
    return res, nil
}

func opGaspriceWrapperBerlin(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
    stack := callContext.Stack
    contract := callContext.Contract
    var err error
    var res []byte	

    in.evm.Context.Counter++

    if sLen := stack.len(); sLen < 0 {
        return nil, &ErrStackUnderflow{stackLen: sLen, required: 0}
    } else if sLen > 1023 {
        return nil, &ErrStackOverflow{stackLen: sLen, limit: 1023}
    }
    if !contract.UseGas(2) {
        return nil, ErrOutOfGas
    }
    res, err = operation.execute(pc, in, callContext)
    if err != nil {
        return nil, err
    }
    *pc++
    return res, nil
}

func opExtCodeSizeWrapperBerlin(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
    mem := callContext.Memory
    stack := callContext.Stack
    contract := callContext.Contract
    var err error
    var res []byte	

    in.evm.Context.Counter++

    if sLen := stack.len(); sLen < 1 {
        return nil, &ErrStackUnderflow{stackLen: sLen, required: 1}
    } else if sLen > 1024 {
        return nil, &ErrStackOverflow{stackLen: sLen, limit: 1024}
    }
    var memorySize uint64
    var dynamicCost uint64
    dynamicCost, err = operation.dynamicGas(in.evm, contract, stack, mem, memorySize)
    if err != nil || !contract.UseGas(100 + dynamicCost) {
        return nil, ErrOutOfGas
    }
    if memorySize > 0 {
        mem.Resize(memorySize)
    }
    res, err = operation.execute(pc, in, callContext)
    if err != nil {
        return nil, err
    }
    *pc++
    return res, nil
}

func opExtCodeCopyWrapperBerlin(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
    mem := callContext.Memory
    stack := callContext.Stack
    contract := callContext.Contract
    var err error
    var res []byte	

    in.evm.Context.Counter++

    if sLen := stack.len(); sLen < 4 {
        return nil, &ErrStackUnderflow{stackLen: sLen, required: 4}
    } else if sLen > 1028 {
        return nil, &ErrStackOverflow{stackLen: sLen, limit: 1028}
    }
    var memorySize uint64
    memSize, overflow := operation.memorySize(stack)
    if overflow {
        return nil, ErrGasUintOverflow
    }
    if memorySize, overflow = math.SafeMul(toWordSize(memSize), 32); overflow {
        return nil, ErrGasUintOverflow
    }
    var dynamicCost uint64
    dynamicCost, err = operation.dynamicGas(in.evm, contract, stack, mem, memorySize)
    if err != nil || !contract.UseGas(100 + dynamicCost) {
        return nil, ErrOutOfGas
    }
    if memorySize > 0 {
        mem.Resize(memorySize)
    }
    res, err = operation.execute(pc, in, callContext)
    if err != nil {
        return nil, err
    }
    *pc++
    return res, nil
}

func opReturnDataSizeWrapperBerlin(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
    stack := callContext.Stack
    contract := callContext.Contract
    var err error
    var res []byte	

    in.evm.Context.Counter++

    if sLen := stack.len(); sLen < 0 {
        return nil, &ErrStackUnderflow{stackLen: sLen, required: 0}
    } else if sLen > 1023 {
        return nil, &ErrStackOverflow{stackLen: sLen, limit: 1023}
    }
    if !contract.UseGas(2) {
        return nil, ErrOutOfGas
    }
    res, err = operation.execute(pc, in, callContext)
    if err != nil {
        return nil, err
    }
    *pc++
    return res, nil
}

func opReturnDataCopyWrapperBerlin(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
    mem := callContext.Memory
    stack := callContext.Stack
    contract := callContext.Contract
    var err error
    var res []byte	

    in.evm.Context.Counter++

    if sLen := stack.len(); sLen < 3 {
        return nil, &ErrStackUnderflow{stackLen: sLen, required: 3}
    } else if sLen > 1027 {
        return nil, &ErrStackOverflow{stackLen: sLen, limit: 1027}
    }
    var memorySize uint64
    memSize, overflow := operation.memorySize(stack)
    if overflow {
        return nil, ErrGasUintOverflow
    }
    if memorySize, overflow = math.SafeMul(toWordSize(memSize), 32); overflow {
        return nil, ErrGasUintOverflow
    }
    var dynamicCost uint64
    dynamicCost, err = operation.dynamicGas(in.evm, contract, stack, mem, memorySize)
    if err != nil || !contract.UseGas(3 + dynamicCost) {
        return nil, ErrOutOfGas
    }
    if memorySize > 0 {
        mem.Resize(memorySize)
    }
    res, err = operation.execute(pc, in, callContext)
    if err != nil {
        return nil, err
    }
    *pc++
    return res, nil
}

func opExtCodeHashWrapperBerlin(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
    mem := callContext.Memory
    stack := callContext.Stack
    contract := callContext.Contract
    var err error
    var res []byte	

    in.evm.Context.Counter++

    if sLen := stack.len(); sLen < 1 {
        return nil, &ErrStackUnderflow{stackLen: sLen, required: 1}
    } else if sLen > 1024 {
        return nil, &ErrStackOverflow{stackLen: sLen, limit: 1024}
    }
    var memorySize uint64
    var dynamicCost uint64
    dynamicCost, err = operation.dynamicGas(in.evm, contract, stack, mem, memorySize)
    if err != nil || !contract.UseGas(100 + dynamicCost) {
        return nil, ErrOutOfGas
    }
    if memorySize > 0 {
        mem.Resize(memorySize)
    }
    res, err = operation.execute(pc, in, callContext)
    if err != nil {
        return nil, err
    }
    *pc++
    return res, nil
}

func opBlockhashWrapperBerlin(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
    stack := callContext.Stack
    contract := callContext.Contract
    var err error
    var res []byte	

    in.evm.Context.Counter++

    if sLen := stack.len(); sLen < 1 {
        return nil, &ErrStackUnderflow{stackLen: sLen, required: 1}
    } else if sLen > 1024 {
        return nil, &ErrStackOverflow{stackLen: sLen, limit: 1024}
    }
    if !contract.UseGas(20) {
        return nil, ErrOutOfGas
    }
    res, err = operation.execute(pc, in, callContext)
    if err != nil {
        return nil, err
    }
    *pc++
    return res, nil
}

func opCoinbaseWrapperBerlin(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
    stack := callContext.Stack
    contract := callContext.Contract
    var err error
    var res []byte	

    in.evm.Context.Counter++

    if sLen := stack.len(); sLen < 0 {
        return nil, &ErrStackUnderflow{stackLen: sLen, required: 0}
    } else if sLen > 1023 {
        return nil, &ErrStackOverflow{stackLen: sLen, limit: 1023}
    }
    if !contract.UseGas(2) {
        return nil, ErrOutOfGas
    }
    res, err = operation.execute(pc, in, callContext)
    if err != nil {
        return nil, err
    }
    *pc++
    return res, nil
}

func opTimestampWrapperBerlin(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
    stack := callContext.Stack
    contract := callContext.Contract
    var err error
    var res []byte	

    in.evm.Context.Counter++

    if sLen := stack.len(); sLen < 0 {
        return nil, &ErrStackUnderflow{stackLen: sLen, required: 0}
    } else if sLen > 1023 {
        return nil, &ErrStackOverflow{stackLen: sLen, limit: 1023}
    }
    if !contract.UseGas(2) {
        return nil, ErrOutOfGas
    }
    res, err = operation.execute(pc, in, callContext)
    if err != nil {
        return nil, err
    }
    *pc++
    return res, nil
}

func opNumberWrapperBerlin(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
    stack := callContext.Stack
    contract := callContext.Contract
    var err error
    var res []byte	

    in.evm.Context.Counter++

    if sLen := stack.len(); sLen < 0 {
        return nil, &ErrStackUnderflow{stackLen: sLen, required: 0}
    } else if sLen > 1023 {
        return nil, &ErrStackOverflow{stackLen: sLen, limit: 1023}
    }
    if !contract.UseGas(2) {
        return nil, ErrOutOfGas
    }
    res, err = operation.execute(pc, in, callContext)
    if err != nil {
        return nil, err
    }
    *pc++
    return res, nil
}

func opDifficultyWrapperBerlin(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
    stack := callContext.Stack
    contract := callContext.Contract
    var err error
    var res []byte	

    in.evm.Context.Counter++

    if sLen := stack.len(); sLen < 0 {
        return nil, &ErrStackUnderflow{stackLen: sLen, required: 0}
    } else if sLen > 1023 {
        return nil, &ErrStackOverflow{stackLen: sLen, limit: 1023}
    }
    if !contract.UseGas(2) {
        return nil, ErrOutOfGas
    }
    res, err = operation.execute(pc, in, callContext)
    if err != nil {
        return nil, err
    }
    *pc++
    return res, nil
}

func opGasLimitWrapperBerlin(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
    stack := callContext.Stack
    contract := callContext.Contract
    var err error
    var res []byte	

    in.evm.Context.Counter++

    if sLen := stack.len(); sLen < 0 {
        return nil, &ErrStackUnderflow{stackLen: sLen, required: 0}
    } else if sLen > 1023 {
        return nil, &ErrStackOverflow{stackLen: sLen, limit: 1023}
    }
    if !contract.UseGas(2) {
        return nil, ErrOutOfGas
    }
    res, err = operation.execute(pc, in, callContext)
    if err != nil {
        return nil, err
    }
    *pc++
    return res, nil
}

func opChainIDWrapperBerlin(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
    stack := callContext.Stack
    contract := callContext.Contract
    var err error
    var res []byte	

    in.evm.Context.Counter++

    if sLen := stack.len(); sLen < 0 {
        return nil, &ErrStackUnderflow{stackLen: sLen, required: 0}
    } else if sLen > 1023 {
        return nil, &ErrStackOverflow{stackLen: sLen, limit: 1023}
    }
    if !contract.UseGas(2) {
        return nil, ErrOutOfGas
    }
    res, err = operation.execute(pc, in, callContext)
    if err != nil {
        return nil, err
    }
    *pc++
    return res, nil
}

func opSelfBalanceWrapperBerlin(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
    stack := callContext.Stack
    contract := callContext.Contract
    var err error
    var res []byte	

    in.evm.Context.Counter++

    if sLen := stack.len(); sLen < 0 {
        return nil, &ErrStackUnderflow{stackLen: sLen, required: 0}
    } else if sLen > 1023 {
        return nil, &ErrStackOverflow{stackLen: sLen, limit: 1023}
    }
    if !contract.UseGas(5) {
        return nil, ErrOutOfGas
    }
    res, err = operation.execute(pc, in, callContext)
    if err != nil {
        return nil, err
    }
    *pc++
    return res, nil
}

func opPopWrapperBerlin(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
    stack := callContext.Stack
    contract := callContext.Contract
    var err error
    var res []byte	

    in.evm.Context.Counter++

    if sLen := stack.len(); sLen < 1 {
        return nil, &ErrStackUnderflow{stackLen: sLen, required: 1}
    } else if sLen > 1025 {
        return nil, &ErrStackOverflow{stackLen: sLen, limit: 1025}
    }
    if !contract.UseGas(2) {
        return nil, ErrOutOfGas
    }
    res, err = operation.execute(pc, in, callContext)
    if err != nil {
        return nil, err
    }
    *pc++
    return res, nil
}

func opMloadWrapperBerlin(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
    mem := callContext.Memory
    stack := callContext.Stack
    contract := callContext.Contract
    var err error
    var res []byte	

    in.evm.Context.Counter++

    if sLen := stack.len(); sLen < 1 {
        return nil, &ErrStackUnderflow{stackLen: sLen, required: 1}
    } else if sLen > 1024 {
        return nil, &ErrStackOverflow{stackLen: sLen, limit: 1024}
    }
    var memorySize uint64
    memSize, overflow := operation.memorySize(stack)
    if overflow {
        return nil, ErrGasUintOverflow
    }
    if memorySize, overflow = math.SafeMul(toWordSize(memSize), 32); overflow {
        return nil, ErrGasUintOverflow
    }
    var dynamicCost uint64
    dynamicCost, err = operation.dynamicGas(in.evm, contract, stack, mem, memorySize)
    if err != nil || !contract.UseGas(3 + dynamicCost) {
        return nil, ErrOutOfGas
    }
    if memorySize > 0 {
        mem.Resize(memorySize)
    }
    res, err = operation.execute(pc, in, callContext)
    if err != nil {
        return nil, err
    }
    *pc++
    return res, nil
}

func opMstoreWrapperBerlin(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
    mem := callContext.Memory
    stack := callContext.Stack
    contract := callContext.Contract
    var err error
    var res []byte	

    in.evm.Context.Counter++

    if sLen := stack.len(); sLen < 2 {
        return nil, &ErrStackUnderflow{stackLen: sLen, required: 2}
    } else if sLen > 1026 {
        return nil, &ErrStackOverflow{stackLen: sLen, limit: 1026}
    }
    var memorySize uint64
    memSize, overflow := operation.memorySize(stack)
    if overflow {
        return nil, ErrGasUintOverflow
    }
    if memorySize, overflow = math.SafeMul(toWordSize(memSize), 32); overflow {
        return nil, ErrGasUintOverflow
    }
    var dynamicCost uint64
    dynamicCost, err = operation.dynamicGas(in.evm, contract, stack, mem, memorySize)
    if err != nil || !contract.UseGas(3 + dynamicCost) {
        return nil, ErrOutOfGas
    }
    if memorySize > 0 {
        mem.Resize(memorySize)
    }
    res, err = operation.execute(pc, in, callContext)
    if err != nil {
        return nil, err
    }
    *pc++
    return res, nil
}

func opMstore8WrapperBerlin(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
    mem := callContext.Memory
    stack := callContext.Stack
    contract := callContext.Contract
    var err error
    var res []byte	

    in.evm.Context.Counter++

    if sLen := stack.len(); sLen < 2 {
        return nil, &ErrStackUnderflow{stackLen: sLen, required: 2}
    } else if sLen > 1026 {
        return nil, &ErrStackOverflow{stackLen: sLen, limit: 1026}
    }
    var memorySize uint64
    memSize, overflow := operation.memorySize(stack)
    if overflow {
        return nil, ErrGasUintOverflow
    }
    if memorySize, overflow = math.SafeMul(toWordSize(memSize), 32); overflow {
        return nil, ErrGasUintOverflow
    }
    var dynamicCost uint64
    dynamicCost, err = operation.dynamicGas(in.evm, contract, stack, mem, memorySize)
    if err != nil || !contract.UseGas(3 + dynamicCost) {
        return nil, ErrOutOfGas
    }
    if memorySize > 0 {
        mem.Resize(memorySize)
    }
    res, err = operation.execute(pc, in, callContext)
    if err != nil {
        return nil, err
    }
    *pc++
    return res, nil
}

func opSloadWrapperBerlin(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
    mem := callContext.Memory
    stack := callContext.Stack
    contract := callContext.Contract
    var err error
    var res []byte	

    in.evm.Context.Counter++

    if sLen := stack.len(); sLen < 1 {
        return nil, &ErrStackUnderflow{stackLen: sLen, required: 1}
    } else if sLen > 1024 {
        return nil, &ErrStackOverflow{stackLen: sLen, limit: 1024}
    }
    var memorySize uint64
    var dynamicCost uint64
    dynamicCost, err = operation.dynamicGas(in.evm, contract, stack, mem, memorySize)
    if err != nil || !contract.UseGas(0 + dynamicCost) {
        return nil, ErrOutOfGas
    }
    if memorySize > 0 {
        mem.Resize(memorySize)
    }
    res, err = operation.execute(pc, in, callContext)
    if err != nil {
        return nil, err
    }
    *pc++
    return res, nil
}

func opSstoreWrapperBerlin(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
    mem := callContext.Memory
    stack := callContext.Stack
    contract := callContext.Contract
    var err error
    var res []byte	

    in.evm.Context.Counter++

    if sLen := stack.len(); sLen < 2 {
        return nil, &ErrStackUnderflow{stackLen: sLen, required: 2}
    } else if sLen > 1026 {
        return nil, &ErrStackOverflow{stackLen: sLen, limit: 1026}
    }
    if in.readOnly && in.evm.chainRules.IsByzantium {
        return nil, ErrWriteProtection
    }
    var memorySize uint64
    var dynamicCost uint64
    dynamicCost, err = operation.dynamicGas(in.evm, contract, stack, mem, memorySize)
    if err != nil || !contract.UseGas(0 + dynamicCost) {
        return nil, ErrOutOfGas
    }
    if memorySize > 0 {
        mem.Resize(memorySize)
    }
    res, err = operation.execute(pc, in, callContext)
    if err != nil {
        return nil, err
    }
    *pc++
    return res, nil
}

func opJumpWrapperBerlin(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
    stack := callContext.Stack
    contract := callContext.Contract
    var err error
    var res []byte	

    in.evm.Context.Counter++

    if sLen := stack.len(); sLen < 1 {
        return nil, &ErrStackUnderflow{stackLen: sLen, required: 1}
    } else if sLen > 1025 {
        return nil, &ErrStackOverflow{stackLen: sLen, limit: 1025}
    }
    if !contract.UseGas(8) {
        return nil, ErrOutOfGas
    }
    res, err = operation.execute(pc, in, callContext)
    if err != nil {
        return nil, err
    }
    return res, nil
}

func opJumpiWrapperBerlin(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
    stack := callContext.Stack
    contract := callContext.Contract
    var err error
    var res []byte	

    in.evm.Context.Counter++

    if sLen := stack.len(); sLen < 2 {
        return nil, &ErrStackUnderflow{stackLen: sLen, required: 2}
    } else if sLen > 1026 {
        return nil, &ErrStackOverflow{stackLen: sLen, limit: 1026}
    }
    if !contract.UseGas(10) {
        return nil, ErrOutOfGas
    }
    res, err = operation.execute(pc, in, callContext)
    if err != nil {
        return nil, err
    }
    return res, nil
}

func opPcWrapperBerlin(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
    stack := callContext.Stack
    contract := callContext.Contract
    var err error
    var res []byte	

    in.evm.Context.Counter++

    if sLen := stack.len(); sLen < 0 {
        return nil, &ErrStackUnderflow{stackLen: sLen, required: 0}
    } else if sLen > 1023 {
        return nil, &ErrStackOverflow{stackLen: sLen, limit: 1023}
    }
    if !contract.UseGas(2) {
        return nil, ErrOutOfGas
    }
    res, err = operation.execute(pc, in, callContext)
    if err != nil {
        return nil, err
    }
    *pc++
    return res, nil
}

func opMsizeWrapperBerlin(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
    stack := callContext.Stack
    contract := callContext.Contract
    var err error
    var res []byte	

    in.evm.Context.Counter++

    if sLen := stack.len(); sLen < 0 {
        return nil, &ErrStackUnderflow{stackLen: sLen, required: 0}
    } else if sLen > 1023 {
        return nil, &ErrStackOverflow{stackLen: sLen, limit: 1023}
    }
    if !contract.UseGas(2) {
        return nil, ErrOutOfGas
    }
    res, err = operation.execute(pc, in, callContext)
    if err != nil {
        return nil, err
    }
    *pc++
    return res, nil
}

func opGasWrapperBerlin(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
    stack := callContext.Stack
    contract := callContext.Contract
    var err error
    var res []byte	

    in.evm.Context.Counter++

    if sLen := stack.len(); sLen < 0 {
        return nil, &ErrStackUnderflow{stackLen: sLen, required: 0}
    } else if sLen > 1023 {
        return nil, &ErrStackOverflow{stackLen: sLen, limit: 1023}
    }
    if !contract.UseGas(2) {
        return nil, ErrOutOfGas
    }
    res, err = operation.execute(pc, in, callContext)
    if err != nil {
        return nil, err
    }
    *pc++
    return res, nil
}

func opJumpdestWrapperBerlin(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
    stack := callContext.Stack
    contract := callContext.Contract
    var err error
    var res []byte	

    in.evm.Context.Counter++

    if sLen := stack.len(); sLen < 0 {
        return nil, &ErrStackUnderflow{stackLen: sLen, required: 0}
    } else if sLen > 1024 {
        return nil, &ErrStackOverflow{stackLen: sLen, limit: 1024}
    }
    if !contract.UseGas(1) {
        return nil, ErrOutOfGas
    }
    res, err = operation.execute(pc, in, callContext)
    if err != nil {
        return nil, err
    }
    *pc++
    return res, nil
}

func opPush1WrapperBerlin(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
    stack := callContext.Stack
    contract := callContext.Contract
    var err error
    var res []byte	

    in.evm.Context.Counter++

    if sLen := stack.len(); sLen < 0 {
        return nil, &ErrStackUnderflow{stackLen: sLen, required: 0}
    } else if sLen > 1023 {
        return nil, &ErrStackOverflow{stackLen: sLen, limit: 1023}
    }
    if !contract.UseGas(3) {
        return nil, ErrOutOfGas
    }
    res, err = operation.execute(pc, in, callContext)
    if err != nil {
        return nil, err
    }
    *pc++
    return res, nil
}

func opPushWrapperBerlin(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
    stack := callContext.Stack
    contract := callContext.Contract
    var err error
    var res []byte	

    in.evm.Context.Counter++

    if sLen := stack.len(); sLen < 0 {
        return nil, &ErrStackUnderflow{stackLen: sLen, required: 0}
    } else if sLen > 1023 {
        return nil, &ErrStackOverflow{stackLen: sLen, limit: 1023}
    }
    if !contract.UseGas(3) {
        return nil, ErrOutOfGas
    }
    res, err = operation.execute(pc, in, callContext)
    if err != nil {
        return nil, err
    }
    *pc++
    return res, nil
}

func opDupWrapperBerlin(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
    stack := callContext.Stack
    contract := callContext.Contract
    var err error
    var res []byte	

    in.evm.Context.Counter++

    if sLen := stack.len(); sLen < 1 {
        return nil, &ErrStackUnderflow{stackLen: sLen, required: 1}
    } else if sLen > 1023 {
        return nil, &ErrStackOverflow{stackLen: sLen, limit: 1023}
    }
    if !contract.UseGas(3) {
        return nil, ErrOutOfGas
    }
    res, err = operation.execute(pc, in, callContext)
    if err != nil {
        return nil, err
    }
    *pc++
    return res, nil
}

func opSwapWrapperBerlin(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
    stack := callContext.Stack
    contract := callContext.Contract
    var err error
    var res []byte	

    in.evm.Context.Counter++

    if sLen := stack.len(); sLen < 2 {
        return nil, &ErrStackUnderflow{stackLen: sLen, required: 2}
    } else if sLen > 1024 {
        return nil, &ErrStackOverflow{stackLen: sLen, limit: 1024}
    }
    if !contract.UseGas(3) {
        return nil, ErrOutOfGas
    }
    res, err = operation.execute(pc, in, callContext)
    if err != nil {
        return nil, err
    }
    *pc++
    return res, nil
}

func opLogWrapperBerlin(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
    mem := callContext.Memory
    stack := callContext.Stack
    contract := callContext.Contract
    var err error
    var res []byte	

    in.evm.Context.Counter++

    if sLen := stack.len(); sLen < 3 {
        return nil, &ErrStackUnderflow{stackLen: sLen, required: 3}
    } else if sLen > 1027 {
        return nil, &ErrStackOverflow{stackLen: sLen, limit: 1027}
    }
    if in.readOnly && in.evm.chainRules.IsByzantium {
        return nil, ErrWriteProtection
    }
    var memorySize uint64
    memSize, overflow := operation.memorySize(stack)
    if overflow {
        return nil, ErrGasUintOverflow
    }
    if memorySize, overflow = math.SafeMul(toWordSize(memSize), 32); overflow {
        return nil, ErrGasUintOverflow
    }
    var dynamicCost uint64
    dynamicCost, err = operation.dynamicGas(in.evm, contract, stack, mem, memorySize)
    if err != nil || !contract.UseGas(0 + dynamicCost) {
        return nil, ErrOutOfGas
    }
    if memorySize > 0 {
        mem.Resize(memorySize)
    }
    res, err = operation.execute(pc, in, callContext)
    if err != nil {
        return nil, err
    }
    *pc++
    return res, nil
}

func opCreateWrapperBerlin(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
    mem := callContext.Memory
    stack := callContext.Stack
    contract := callContext.Contract
    var err error
    var res []byte	

    in.evm.Context.Counter++

    if sLen := stack.len(); sLen < 3 {
        return nil, &ErrStackUnderflow{stackLen: sLen, required: 3}
    } else if sLen > 1026 {
        return nil, &ErrStackOverflow{stackLen: sLen, limit: 1026}
    }
    if in.readOnly && in.evm.chainRules.IsByzantium {
        return nil, ErrWriteProtection
    }
    var memorySize uint64
    memSize, overflow := operation.memorySize(stack)
    if overflow {
        return nil, ErrGasUintOverflow
    }
    if memorySize, overflow = math.SafeMul(toWordSize(memSize), 32); overflow {
        return nil, ErrGasUintOverflow
    }
    var dynamicCost uint64
    dynamicCost, err = operation.dynamicGas(in.evm, contract, stack, mem, memorySize)
    if err != nil || !contract.UseGas(32000 + dynamicCost) {
        return nil, ErrOutOfGas
    }
    if memorySize > 0 {
        mem.Resize(memorySize)
    }
    res, err = operation.execute(pc, in, callContext)
    in.returnData = res
    if err != nil {
        return nil, err
    }
    *pc++
    return res, nil
}

func opCallWrapperBerlin(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
    mem := callContext.Memory
    stack := callContext.Stack
    contract := callContext.Contract
    var err error
    var res []byte	

    in.evm.Context.Counter++

    if sLen := stack.len(); sLen < 7 {
        return nil, &ErrStackUnderflow{stackLen: sLen, required: 7}
    } else if sLen > 1030 {
        return nil, &ErrStackOverflow{stackLen: sLen, limit: 1030}
    }
    if in.readOnly && in.evm.chainRules.IsByzantium {
        if stack.Back(2).Sign() != 0 {
            return nil, ErrWriteProtection
        }
    }
    var memorySize uint64
    memSize, overflow := operation.memorySize(stack)
    if overflow {
        return nil, ErrGasUintOverflow
    }
    if memorySize, overflow = math.SafeMul(toWordSize(memSize), 32); overflow {
        return nil, ErrGasUintOverflow
    }
    var dynamicCost uint64
    dynamicCost, err = operation.dynamicGas(in.evm, contract, stack, mem, memorySize)
    if err != nil || !contract.UseGas(100 + dynamicCost) {
        return nil, ErrOutOfGas
    }
    if memorySize > 0 {
        mem.Resize(memorySize)
    }
    res, err = operation.execute(pc, in, callContext)
    in.returnData = res
    if err != nil {
        return nil, err
    }
    *pc++
    return res, nil
}

func opCallCodeWrapperBerlin(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
    mem := callContext.Memory
    stack := callContext.Stack
    contract := callContext.Contract
    var err error
    var res []byte	

    in.evm.Context.Counter++

    if sLen := stack.len(); sLen < 7 {
        return nil, &ErrStackUnderflow{stackLen: sLen, required: 7}
    } else if sLen > 1030 {
        return nil, &ErrStackOverflow{stackLen: sLen, limit: 1030}
    }
    var memorySize uint64
    memSize, overflow := operation.memorySize(stack)
    if overflow {
        return nil, ErrGasUintOverflow
    }
    if memorySize, overflow = math.SafeMul(toWordSize(memSize), 32); overflow {
        return nil, ErrGasUintOverflow
    }
    var dynamicCost uint64
    dynamicCost, err = operation.dynamicGas(in.evm, contract, stack, mem, memorySize)
    if err != nil || !contract.UseGas(100 + dynamicCost) {
        return nil, ErrOutOfGas
    }
    if memorySize > 0 {
        mem.Resize(memorySize)
    }
    res, err = operation.execute(pc, in, callContext)
    in.returnData = res
    if err != nil {
        return nil, err
    }
    *pc++
    return res, nil
}

func opReturnWrapperBerlin(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
    mem := callContext.Memory
    stack := callContext.Stack
    contract := callContext.Contract
    var err error
    var res []byte	

    in.evm.Context.Counter++

    if sLen := stack.len(); sLen < 2 {
        return nil, &ErrStackUnderflow{stackLen: sLen, required: 2}
    } else if sLen > 1026 {
        return nil, &ErrStackOverflow{stackLen: sLen, limit: 1026}
    }
    var memorySize uint64
    memSize, overflow := operation.memorySize(stack)
    if overflow {
        return nil, ErrGasUintOverflow
    }
    if memorySize, overflow = math.SafeMul(toWordSize(memSize), 32); overflow {
        return nil, ErrGasUintOverflow
    }
    var dynamicCost uint64
    dynamicCost, err = operation.dynamicGas(in.evm, contract, stack, mem, memorySize)
    if err != nil || !contract.UseGas(0 + dynamicCost) {
        return nil, ErrOutOfGas
    }
    if memorySize > 0 {
        mem.Resize(memorySize)
    }
    res, err = operation.execute(pc, in, callContext)
    if err != nil {
        return nil, err
    }
    *pc++
    return res, nil
}

func opDelegateCallWrapperBerlin(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
    mem := callContext.Memory
    stack := callContext.Stack
    contract := callContext.Contract
    var err error
    var res []byte	

    in.evm.Context.Counter++

    if sLen := stack.len(); sLen < 6 {
        return nil, &ErrStackUnderflow{stackLen: sLen, required: 6}
    } else if sLen > 1029 {
        return nil, &ErrStackOverflow{stackLen: sLen, limit: 1029}
    }
    var memorySize uint64
    memSize, overflow := operation.memorySize(stack)
    if overflow {
        return nil, ErrGasUintOverflow
    }
    if memorySize, overflow = math.SafeMul(toWordSize(memSize), 32); overflow {
        return nil, ErrGasUintOverflow
    }
    var dynamicCost uint64
    dynamicCost, err = operation.dynamicGas(in.evm, contract, stack, mem, memorySize)
    if err != nil || !contract.UseGas(100 + dynamicCost) {
        return nil, ErrOutOfGas
    }
    if memorySize > 0 {
        mem.Resize(memorySize)
    }
    res, err = operation.execute(pc, in, callContext)
    in.returnData = res
    if err != nil {
        return nil, err
    }
    *pc++
    return res, nil
}

func opCreate2WrapperBerlin(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
    mem := callContext.Memory
    stack := callContext.Stack
    contract := callContext.Contract
    var err error
    var res []byte	

    in.evm.Context.Counter++

    if sLen := stack.len(); sLen < 4 {
        return nil, &ErrStackUnderflow{stackLen: sLen, required: 4}
    } else if sLen > 1027 {
        return nil, &ErrStackOverflow{stackLen: sLen, limit: 1027}
    }
    if in.readOnly && in.evm.chainRules.IsByzantium {
        return nil, ErrWriteProtection
    }
    var memorySize uint64
    memSize, overflow := operation.memorySize(stack)
    if overflow {
        return nil, ErrGasUintOverflow
    }
    if memorySize, overflow = math.SafeMul(toWordSize(memSize), 32); overflow {
        return nil, ErrGasUintOverflow
    }
    var dynamicCost uint64
    dynamicCost, err = operation.dynamicGas(in.evm, contract, stack, mem, memorySize)
    if err != nil || !contract.UseGas(32000 + dynamicCost) {
        return nil, ErrOutOfGas
    }
    if memorySize > 0 {
        mem.Resize(memorySize)
    }
    res, err = operation.execute(pc, in, callContext)
    in.returnData = res
    if err != nil {
        return nil, err
    }
    *pc++
    return res, nil
}

func opStaticCallWrapperBerlin(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
    mem := callContext.Memory
    stack := callContext.Stack
    contract := callContext.Contract
    var err error
    var res []byte	

    in.evm.Context.Counter++

    if sLen := stack.len(); sLen < 6 {
        return nil, &ErrStackUnderflow{stackLen: sLen, required: 6}
    } else if sLen > 1029 {
        return nil, &ErrStackOverflow{stackLen: sLen, limit: 1029}
    }
    var memorySize uint64
    memSize, overflow := operation.memorySize(stack)
    if overflow {
        return nil, ErrGasUintOverflow
    }
    if memorySize, overflow = math.SafeMul(toWordSize(memSize), 32); overflow {
        return nil, ErrGasUintOverflow
    }
    var dynamicCost uint64
    dynamicCost, err = operation.dynamicGas(in.evm, contract, stack, mem, memorySize)
    if err != nil || !contract.UseGas(100 + dynamicCost) {
        return nil, ErrOutOfGas
    }
    if memorySize > 0 {
        mem.Resize(memorySize)
    }
    res, err = operation.execute(pc, in, callContext)
    in.returnData = res
    if err != nil {
        return nil, err
    }
    *pc++
    return res, nil
}

func opRevertWrapperBerlin(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
    mem := callContext.Memory
    stack := callContext.Stack
    contract := callContext.Contract
    var err error
    var res []byte	

    in.evm.Context.Counter++

    if sLen := stack.len(); sLen < 2 {
        return nil, &ErrStackUnderflow{stackLen: sLen, required: 2}
    } else if sLen > 1026 {
        return nil, &ErrStackOverflow{stackLen: sLen, limit: 1026}
    }
    var memorySize uint64
    memSize, overflow := operation.memorySize(stack)
    if overflow {
        return nil, ErrGasUintOverflow
    }
    if memorySize, overflow = math.SafeMul(toWordSize(memSize), 32); overflow {
        return nil, ErrGasUintOverflow
    }
    var dynamicCost uint64
    dynamicCost, err = operation.dynamicGas(in.evm, contract, stack, mem, memorySize)
    if err != nil || !contract.UseGas(0 + dynamicCost) {
        return nil, ErrOutOfGas
    }
    if memorySize > 0 {
        mem.Resize(memorySize)
    }
    res, err = operation.execute(pc, in, callContext)
    in.returnData = res
    if err != nil {
        return nil, err
    }
    return res, ErrExecutionReverted
}

func opSuicideWrapperBerlin(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
    mem := callContext.Memory
    stack := callContext.Stack
    contract := callContext.Contract
    var err error
    var res []byte	

    in.evm.Context.Counter++

    if sLen := stack.len(); sLen < 1 {
        return nil, &ErrStackUnderflow{stackLen: sLen, required: 1}
    } else if sLen > 1025 {
        return nil, &ErrStackOverflow{stackLen: sLen, limit: 1025}
    }
    if in.readOnly && in.evm.chainRules.IsByzantium {
        return nil, ErrWriteProtection
    }
    var memorySize uint64
    var dynamicCost uint64
    dynamicCost, err = operation.dynamicGas(in.evm, contract, stack, mem, memorySize)
    if err != nil || !contract.UseGas(5000 + dynamicCost) {
        return nil, ErrOutOfGas
    }
    if memorySize > 0 {
        mem.Resize(memorySize)
    }
    res, err = operation.execute(pc, in, callContext)
    if err != nil {
        return nil, err
    }
    *pc++
    return res, nil
}

func opBaseFeeWrapperBerlin(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
	return nil, nil
}

var stringToWrapperBerlin = map[string]executionWrapperFunc{
    "opStop": opStopWrapperBerlin,
    "opAdd": opAddWrapperBerlin,
    "opMul": opMulWrapperBerlin,
    "opSub": opSubWrapperBerlin,
    "opDiv": opDivWrapperBerlin,
    "opSdiv": opSdivWrapperBerlin,
    "opMod": opModWrapperBerlin,
    "opSmod": opSmodWrapperBerlin,
    "opAddmod": opAddmodWrapperBerlin,
    "opMulmod": opMulmodWrapperBerlin,
    "opExp": opExpWrapperBerlin,
    "opSignExtend": opSignExtendWrapperBerlin,
    "opLt": opLtWrapperBerlin,
    "opGt": opGtWrapperBerlin,
    "opSlt": opSltWrapperBerlin,
    "opSgt": opSgtWrapperBerlin,
    "opEq": opEqWrapperBerlin,
    "opIszero": opIszeroWrapperBerlin,
    "opAnd": opAndWrapperBerlin,
    "opOr": opOrWrapperBerlin,
    "opXor": opXorWrapperBerlin,
    "opNot": opNotWrapperBerlin,
    "opByte": opByteWrapperBerlin,
    "opSHL": opSHLWrapperBerlin,
    "opSHR": opSHRWrapperBerlin,
    "opSAR": opSARWrapperBerlin,
    "opSha3": opSha3WrapperBerlin,
    "opAddress": opAddressWrapperBerlin,
    "opBalance": opBalanceWrapperBerlin,
    "opOrigin": opOriginWrapperBerlin,
    "opCaller": opCallerWrapperBerlin,
    "opCallValue": opCallValueWrapperBerlin,
    "opCallDataLoad": opCallDataLoadWrapperBerlin,
    "opCallDataSize": opCallDataSizeWrapperBerlin,
    "opCallDataCopy": opCallDataCopyWrapperBerlin,
    "opCodeSize": opCodeSizeWrapperBerlin,
    "opCodeCopy": opCodeCopyWrapperBerlin,
    "opGasprice": opGaspriceWrapperBerlin,
    "opExtCodeSize": opExtCodeSizeWrapperBerlin,
    "opExtCodeCopy": opExtCodeCopyWrapperBerlin,
    "opReturnDataSize": opReturnDataSizeWrapperBerlin,
    "opReturnDataCopy": opReturnDataCopyWrapperBerlin,
    "opExtCodeHash": opExtCodeHashWrapperBerlin,
    "opBlockhash": opBlockhashWrapperBerlin,
    "opCoinbase": opCoinbaseWrapperBerlin,
    "opTimestamp": opTimestampWrapperBerlin,
    "opNumber": opNumberWrapperBerlin,
    "opDifficulty": opDifficultyWrapperBerlin,
    "opGasLimit": opGasLimitWrapperBerlin,
    "opChainID": opChainIDWrapperBerlin,
    "opSelfBalance": opSelfBalanceWrapperBerlin,
    "opPop": opPopWrapperBerlin,
    "opMload": opMloadWrapperBerlin,
    "opMstore": opMstoreWrapperBerlin,
    "opMstore8": opMstore8WrapperBerlin,
    "opSload": opSloadWrapperBerlin,
    "opSstore": opSstoreWrapperBerlin,
    "opJump": opJumpWrapperBerlin,
    "opJumpi": opJumpiWrapperBerlin,
    "opPc": opPcWrapperBerlin,
    "opMsize": opMsizeWrapperBerlin,
    "opGas": opGasWrapperBerlin,
    "opJumpdest": opJumpdestWrapperBerlin,
    "opPush1": opPush1WrapperBerlin,
    "opPush": opPushWrapperBerlin,
    "opDup": opDupWrapperBerlin,
    "opSwap": opSwapWrapperBerlin,
    "opLog": opLogWrapperBerlin,
    "opCreate": opCreateWrapperBerlin,
    "opCall": opCallWrapperBerlin,
    "opCallCode": opCallCodeWrapperBerlin,
    "opReturn": opReturnWrapperBerlin,
    "opDelegateCall": opDelegateCallWrapperBerlin,
    "opCreate2": opCreate2WrapperBerlin,
    "opStaticCall": opStaticCallWrapperBerlin,
    "opRevert": opRevertWrapperBerlin,
    "opSuicide": opSuicideWrapperBerlin,
    "opBaseFee": opBaseFeeWrapperBerlin,
}
