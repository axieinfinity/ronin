
package vm

import "github.com/ethereum/go-ethereum/common/math"


func opStopWrapperTangerineWhistle(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opAddWrapperTangerineWhistle(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opMulWrapperTangerineWhistle(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opSubWrapperTangerineWhistle(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opDivWrapperTangerineWhistle(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opSdivWrapperTangerineWhistle(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opModWrapperTangerineWhistle(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opSmodWrapperTangerineWhistle(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opAddmodWrapperTangerineWhistle(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opMulmodWrapperTangerineWhistle(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opExpWrapperTangerineWhistle(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opSignExtendWrapperTangerineWhistle(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opLtWrapperTangerineWhistle(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opGtWrapperTangerineWhistle(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opSltWrapperTangerineWhistle(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opSgtWrapperTangerineWhistle(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opEqWrapperTangerineWhistle(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opIszeroWrapperTangerineWhistle(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opAndWrapperTangerineWhistle(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opOrWrapperTangerineWhistle(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opXorWrapperTangerineWhistle(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opNotWrapperTangerineWhistle(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opByteWrapperTangerineWhistle(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opSha3WrapperTangerineWhistle(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opAddressWrapperTangerineWhistle(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opBalanceWrapperTangerineWhistle(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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
    if !contract.UseGas(400) {
        return nil, ErrOutOfGas
    }
    res, err = operation.execute(pc, in, callContext)
    if err != nil {
        return nil, err
    }
    *pc++
    return res, nil
}

func opOriginWrapperTangerineWhistle(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opCallerWrapperTangerineWhistle(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opCallValueWrapperTangerineWhistle(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opCallDataLoadWrapperTangerineWhistle(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opCallDataSizeWrapperTangerineWhistle(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opCallDataCopyWrapperTangerineWhistle(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opCodeSizeWrapperTangerineWhistle(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opCodeCopyWrapperTangerineWhistle(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opGaspriceWrapperTangerineWhistle(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opExtCodeSizeWrapperTangerineWhistle(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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
    if !contract.UseGas(700) {
        return nil, ErrOutOfGas
    }
    res, err = operation.execute(pc, in, callContext)
    if err != nil {
        return nil, err
    }
    *pc++
    return res, nil
}

func opExtCodeCopyWrapperTangerineWhistle(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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
    if err != nil || !contract.UseGas(700 + dynamicCost) {
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

func opBlockhashWrapperTangerineWhistle(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opCoinbaseWrapperTangerineWhistle(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opTimestampWrapperTangerineWhistle(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opNumberWrapperTangerineWhistle(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opDifficultyWrapperTangerineWhistle(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opGasLimitWrapperTangerineWhistle(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opPopWrapperTangerineWhistle(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opMloadWrapperTangerineWhistle(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opMstoreWrapperTangerineWhistle(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opMstore8WrapperTangerineWhistle(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opSloadWrapperTangerineWhistle(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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
    if !contract.UseGas(200) {
        return nil, ErrOutOfGas
    }
    res, err = operation.execute(pc, in, callContext)
    if err != nil {
        return nil, err
    }
    *pc++
    return res, nil
}

func opSstoreWrapperTangerineWhistle(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opJumpWrapperTangerineWhistle(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opJumpiWrapperTangerineWhistle(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opPcWrapperTangerineWhistle(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opMsizeWrapperTangerineWhistle(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opGasWrapperTangerineWhistle(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opJumpdestWrapperTangerineWhistle(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opPush1WrapperTangerineWhistle(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opPushWrapperTangerineWhistle(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opDupWrapperTangerineWhistle(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opSwapWrapperTangerineWhistle(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opLogWrapperTangerineWhistle(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opCreateWrapperTangerineWhistle(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opCallWrapperTangerineWhistle(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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
    if err != nil || !contract.UseGas(700 + dynamicCost) {
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

func opCallCodeWrapperTangerineWhistle(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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
    if err != nil || !contract.UseGas(700 + dynamicCost) {
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

func opReturnWrapperTangerineWhistle(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opDelegateCallWrapperTangerineWhistle(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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
    if err != nil || !contract.UseGas(700 + dynamicCost) {
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

func opSuicideWrapperTangerineWhistle(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opSHLWrapperTangerineWhistle(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
	return nil, nil
}

func opSHRWrapperTangerineWhistle(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
	return nil, nil
}

func opSARWrapperTangerineWhistle(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
	return nil, nil
}

func opReturnDataSizeWrapperTangerineWhistle(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
	return nil, nil
}

func opReturnDataCopyWrapperTangerineWhistle(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
	return nil, nil
}

func opExtCodeHashWrapperTangerineWhistle(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
	return nil, nil
}

func opChainIDWrapperTangerineWhistle(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
	return nil, nil
}

func opSelfBalanceWrapperTangerineWhistle(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
	return nil, nil
}

func opBaseFeeWrapperTangerineWhistle(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
	return nil, nil
}

func opCreate2WrapperTangerineWhistle(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
	return nil, nil
}

func opStaticCallWrapperTangerineWhistle(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
	return nil, nil
}

func opRevertWrapperTangerineWhistle(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
	return nil, nil
}

var stringToWrapperTangerineWhistle = map[string]executionWrapperFunc{
    "opStop": opStopWrapperTangerineWhistle,
    "opAdd": opAddWrapperTangerineWhistle,
    "opMul": opMulWrapperTangerineWhistle,
    "opSub": opSubWrapperTangerineWhistle,
    "opDiv": opDivWrapperTangerineWhistle,
    "opSdiv": opSdivWrapperTangerineWhistle,
    "opMod": opModWrapperTangerineWhistle,
    "opSmod": opSmodWrapperTangerineWhistle,
    "opAddmod": opAddmodWrapperTangerineWhistle,
    "opMulmod": opMulmodWrapperTangerineWhistle,
    "opExp": opExpWrapperTangerineWhistle,
    "opSignExtend": opSignExtendWrapperTangerineWhistle,
    "opLt": opLtWrapperTangerineWhistle,
    "opGt": opGtWrapperTangerineWhistle,
    "opSlt": opSltWrapperTangerineWhistle,
    "opSgt": opSgtWrapperTangerineWhistle,
    "opEq": opEqWrapperTangerineWhistle,
    "opIszero": opIszeroWrapperTangerineWhistle,
    "opAnd": opAndWrapperTangerineWhistle,
    "opOr": opOrWrapperTangerineWhistle,
    "opXor": opXorWrapperTangerineWhistle,
    "opNot": opNotWrapperTangerineWhistle,
    "opByte": opByteWrapperTangerineWhistle,
    "opSha3": opSha3WrapperTangerineWhistle,
    "opAddress": opAddressWrapperTangerineWhistle,
    "opBalance": opBalanceWrapperTangerineWhistle,
    "opOrigin": opOriginWrapperTangerineWhistle,
    "opCaller": opCallerWrapperTangerineWhistle,
    "opCallValue": opCallValueWrapperTangerineWhistle,
    "opCallDataLoad": opCallDataLoadWrapperTangerineWhistle,
    "opCallDataSize": opCallDataSizeWrapperTangerineWhistle,
    "opCallDataCopy": opCallDataCopyWrapperTangerineWhistle,
    "opCodeSize": opCodeSizeWrapperTangerineWhistle,
    "opCodeCopy": opCodeCopyWrapperTangerineWhistle,
    "opGasprice": opGaspriceWrapperTangerineWhistle,
    "opExtCodeSize": opExtCodeSizeWrapperTangerineWhistle,
    "opExtCodeCopy": opExtCodeCopyWrapperTangerineWhistle,
    "opBlockhash": opBlockhashWrapperTangerineWhistle,
    "opCoinbase": opCoinbaseWrapperTangerineWhistle,
    "opTimestamp": opTimestampWrapperTangerineWhistle,
    "opNumber": opNumberWrapperTangerineWhistle,
    "opDifficulty": opDifficultyWrapperTangerineWhistle,
    "opGasLimit": opGasLimitWrapperTangerineWhistle,
    "opPop": opPopWrapperTangerineWhistle,
    "opMload": opMloadWrapperTangerineWhistle,
    "opMstore": opMstoreWrapperTangerineWhistle,
    "opMstore8": opMstore8WrapperTangerineWhistle,
    "opSload": opSloadWrapperTangerineWhistle,
    "opSstore": opSstoreWrapperTangerineWhistle,
    "opJump": opJumpWrapperTangerineWhistle,
    "opJumpi": opJumpiWrapperTangerineWhistle,
    "opPc": opPcWrapperTangerineWhistle,
    "opMsize": opMsizeWrapperTangerineWhistle,
    "opGas": opGasWrapperTangerineWhistle,
    "opJumpdest": opJumpdestWrapperTangerineWhistle,
    "opPush1": opPush1WrapperTangerineWhistle,
    "opPush": opPushWrapperTangerineWhistle,
    "opDup": opDupWrapperTangerineWhistle,
    "opSwap": opSwapWrapperTangerineWhistle,
    "opLog": opLogWrapperTangerineWhistle,
    "opCreate": opCreateWrapperTangerineWhistle,
    "opCall": opCallWrapperTangerineWhistle,
    "opCallCode": opCallCodeWrapperTangerineWhistle,
    "opReturn": opReturnWrapperTangerineWhistle,
    "opDelegateCall": opDelegateCallWrapperTangerineWhistle,
    "opSuicide": opSuicideWrapperTangerineWhistle,
    "opSHL": opSHLWrapperTangerineWhistle,
    "opSHR": opSHRWrapperTangerineWhistle,
    "opSAR": opSARWrapperTangerineWhistle,
    "opReturnDataSize": opReturnDataSizeWrapperTangerineWhistle,
    "opReturnDataCopy": opReturnDataCopyWrapperTangerineWhistle,
    "opExtCodeHash": opExtCodeHashWrapperTangerineWhistle,
    "opChainID": opChainIDWrapperTangerineWhistle,
    "opSelfBalance": opSelfBalanceWrapperTangerineWhistle,
    "opBaseFee": opBaseFeeWrapperTangerineWhistle,
    "opCreate2": opCreate2WrapperTangerineWhistle,
    "opStaticCall": opStaticCallWrapperTangerineWhistle,
    "opRevert": opRevertWrapperTangerineWhistle,
}
