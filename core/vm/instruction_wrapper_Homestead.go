
package vm

import "github.com/ethereum/go-ethereum/common/math"


func opStopWrapperHomestead(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opAddWrapperHomestead(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opMulWrapperHomestead(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opSubWrapperHomestead(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opDivWrapperHomestead(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opSdivWrapperHomestead(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opModWrapperHomestead(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opSmodWrapperHomestead(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opAddmodWrapperHomestead(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opMulmodWrapperHomestead(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opExpWrapperHomestead(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opSignExtendWrapperHomestead(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opLtWrapperHomestead(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opGtWrapperHomestead(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opSltWrapperHomestead(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opSgtWrapperHomestead(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opEqWrapperHomestead(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opIszeroWrapperHomestead(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opAndWrapperHomestead(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opOrWrapperHomestead(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opXorWrapperHomestead(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opNotWrapperHomestead(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opByteWrapperHomestead(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opSha3WrapperHomestead(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opAddressWrapperHomestead(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opBalanceWrapperHomestead(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opOriginWrapperHomestead(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opCallerWrapperHomestead(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opCallValueWrapperHomestead(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opCallDataLoadWrapperHomestead(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opCallDataSizeWrapperHomestead(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opCallDataCopyWrapperHomestead(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opCodeSizeWrapperHomestead(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opCodeCopyWrapperHomestead(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opGaspriceWrapperHomestead(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opExtCodeSizeWrapperHomestead(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opExtCodeCopyWrapperHomestead(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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
    if err != nil || !contract.UseGas(20 + dynamicCost) {
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

func opBlockhashWrapperHomestead(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opCoinbaseWrapperHomestead(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opTimestampWrapperHomestead(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opNumberWrapperHomestead(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opDifficultyWrapperHomestead(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opGasLimitWrapperHomestead(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opPopWrapperHomestead(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opMloadWrapperHomestead(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opMstoreWrapperHomestead(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opMstore8WrapperHomestead(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opSloadWrapperHomestead(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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
    if !contract.UseGas(50) {
        return nil, ErrOutOfGas
    }
    res, err = operation.execute(pc, in, callContext)
    if err != nil {
        return nil, err
    }
    *pc++
    return res, nil
}

func opSstoreWrapperHomestead(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opJumpWrapperHomestead(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opJumpiWrapperHomestead(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opPcWrapperHomestead(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opMsizeWrapperHomestead(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opGasWrapperHomestead(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opJumpdestWrapperHomestead(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opPush1WrapperHomestead(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opPushWrapperHomestead(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opDupWrapperHomestead(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opSwapWrapperHomestead(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opLogWrapperHomestead(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opCreateWrapperHomestead(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opCallWrapperHomestead(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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
    if err != nil || !contract.UseGas(40 + dynamicCost) {
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

func opCallCodeWrapperHomestead(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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
    if err != nil || !contract.UseGas(40 + dynamicCost) {
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

func opReturnWrapperHomestead(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opDelegateCallWrapperHomestead(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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
    if err != nil || !contract.UseGas(40 + dynamicCost) {
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

func opSuicideWrapperHomestead(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opSHLWrapperHomestead(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
	return nil, nil
}

func opSHRWrapperHomestead(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
	return nil, nil
}

func opSARWrapperHomestead(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
	return nil, nil
}

func opReturnDataSizeWrapperHomestead(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
	return nil, nil
}

func opReturnDataCopyWrapperHomestead(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
	return nil, nil
}

func opExtCodeHashWrapperHomestead(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
	return nil, nil
}

func opChainIDWrapperHomestead(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
	return nil, nil
}

func opSelfBalanceWrapperHomestead(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
	return nil, nil
}

func opBaseFeeWrapperHomestead(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
	return nil, nil
}

func opCreate2WrapperHomestead(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
	return nil, nil
}

func opStaticCallWrapperHomestead(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
	return nil, nil
}

func opRevertWrapperHomestead(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
	return nil, nil
}

var stringToWrapperHomestead = map[string]executionWrapperFunc{
    "opStop": opStopWrapperHomestead,
    "opAdd": opAddWrapperHomestead,
    "opMul": opMulWrapperHomestead,
    "opSub": opSubWrapperHomestead,
    "opDiv": opDivWrapperHomestead,
    "opSdiv": opSdivWrapperHomestead,
    "opMod": opModWrapperHomestead,
    "opSmod": opSmodWrapperHomestead,
    "opAddmod": opAddmodWrapperHomestead,
    "opMulmod": opMulmodWrapperHomestead,
    "opExp": opExpWrapperHomestead,
    "opSignExtend": opSignExtendWrapperHomestead,
    "opLt": opLtWrapperHomestead,
    "opGt": opGtWrapperHomestead,
    "opSlt": opSltWrapperHomestead,
    "opSgt": opSgtWrapperHomestead,
    "opEq": opEqWrapperHomestead,
    "opIszero": opIszeroWrapperHomestead,
    "opAnd": opAndWrapperHomestead,
    "opOr": opOrWrapperHomestead,
    "opXor": opXorWrapperHomestead,
    "opNot": opNotWrapperHomestead,
    "opByte": opByteWrapperHomestead,
    "opSha3": opSha3WrapperHomestead,
    "opAddress": opAddressWrapperHomestead,
    "opBalance": opBalanceWrapperHomestead,
    "opOrigin": opOriginWrapperHomestead,
    "opCaller": opCallerWrapperHomestead,
    "opCallValue": opCallValueWrapperHomestead,
    "opCallDataLoad": opCallDataLoadWrapperHomestead,
    "opCallDataSize": opCallDataSizeWrapperHomestead,
    "opCallDataCopy": opCallDataCopyWrapperHomestead,
    "opCodeSize": opCodeSizeWrapperHomestead,
    "opCodeCopy": opCodeCopyWrapperHomestead,
    "opGasprice": opGaspriceWrapperHomestead,
    "opExtCodeSize": opExtCodeSizeWrapperHomestead,
    "opExtCodeCopy": opExtCodeCopyWrapperHomestead,
    "opBlockhash": opBlockhashWrapperHomestead,
    "opCoinbase": opCoinbaseWrapperHomestead,
    "opTimestamp": opTimestampWrapperHomestead,
    "opNumber": opNumberWrapperHomestead,
    "opDifficulty": opDifficultyWrapperHomestead,
    "opGasLimit": opGasLimitWrapperHomestead,
    "opPop": opPopWrapperHomestead,
    "opMload": opMloadWrapperHomestead,
    "opMstore": opMstoreWrapperHomestead,
    "opMstore8": opMstore8WrapperHomestead,
    "opSload": opSloadWrapperHomestead,
    "opSstore": opSstoreWrapperHomestead,
    "opJump": opJumpWrapperHomestead,
    "opJumpi": opJumpiWrapperHomestead,
    "opPc": opPcWrapperHomestead,
    "opMsize": opMsizeWrapperHomestead,
    "opGas": opGasWrapperHomestead,
    "opJumpdest": opJumpdestWrapperHomestead,
    "opPush1": opPush1WrapperHomestead,
    "opPush": opPushWrapperHomestead,
    "opDup": opDupWrapperHomestead,
    "opSwap": opSwapWrapperHomestead,
    "opLog": opLogWrapperHomestead,
    "opCreate": opCreateWrapperHomestead,
    "opCall": opCallWrapperHomestead,
    "opCallCode": opCallCodeWrapperHomestead,
    "opReturn": opReturnWrapperHomestead,
    "opDelegateCall": opDelegateCallWrapperHomestead,
    "opSuicide": opSuicideWrapperHomestead,
    "opSHL": opSHLWrapperHomestead,
    "opSHR": opSHRWrapperHomestead,
    "opSAR": opSARWrapperHomestead,
    "opReturnDataSize": opReturnDataSizeWrapperHomestead,
    "opReturnDataCopy": opReturnDataCopyWrapperHomestead,
    "opExtCodeHash": opExtCodeHashWrapperHomestead,
    "opChainID": opChainIDWrapperHomestead,
    "opSelfBalance": opSelfBalanceWrapperHomestead,
    "opBaseFee": opBaseFeeWrapperHomestead,
    "opCreate2": opCreate2WrapperHomestead,
    "opStaticCall": opStaticCallWrapperHomestead,
    "opRevert": opRevertWrapperHomestead,
}
