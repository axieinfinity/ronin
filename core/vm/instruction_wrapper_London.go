
package vm

import "github.com/ethereum/go-ethereum/common/math"


func opStopWrapperLondon(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opAddWrapperLondon(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opMulWrapperLondon(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opSubWrapperLondon(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opDivWrapperLondon(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opSdivWrapperLondon(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opModWrapperLondon(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opSmodWrapperLondon(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opAddmodWrapperLondon(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opMulmodWrapperLondon(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opExpWrapperLondon(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opSignExtendWrapperLondon(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opLtWrapperLondon(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opGtWrapperLondon(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opSltWrapperLondon(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opSgtWrapperLondon(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opEqWrapperLondon(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opIszeroWrapperLondon(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opAndWrapperLondon(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opOrWrapperLondon(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opXorWrapperLondon(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opNotWrapperLondon(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opByteWrapperLondon(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opSHLWrapperLondon(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opSHRWrapperLondon(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opSARWrapperLondon(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opSha3WrapperLondon(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opAddressWrapperLondon(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opBalanceWrapperLondon(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opOriginWrapperLondon(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opCallerWrapperLondon(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opCallValueWrapperLondon(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opCallDataLoadWrapperLondon(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opCallDataSizeWrapperLondon(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opCallDataCopyWrapperLondon(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opCodeSizeWrapperLondon(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opCodeCopyWrapperLondon(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opGaspriceWrapperLondon(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opExtCodeSizeWrapperLondon(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opExtCodeCopyWrapperLondon(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opReturnDataSizeWrapperLondon(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opReturnDataCopyWrapperLondon(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opExtCodeHashWrapperLondon(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opBlockhashWrapperLondon(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opCoinbaseWrapperLondon(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opTimestampWrapperLondon(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opNumberWrapperLondon(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opDifficultyWrapperLondon(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opGasLimitWrapperLondon(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opChainIDWrapperLondon(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opSelfBalanceWrapperLondon(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opBaseFeeWrapperLondon(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opPopWrapperLondon(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opMloadWrapperLondon(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opMstoreWrapperLondon(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opMstore8WrapperLondon(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opSloadWrapperLondon(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opSstoreWrapperLondon(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opJumpWrapperLondon(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opJumpiWrapperLondon(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opPcWrapperLondon(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opMsizeWrapperLondon(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opGasWrapperLondon(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opJumpdestWrapperLondon(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opPush1WrapperLondon(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opPushWrapperLondon(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opDupWrapperLondon(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opSwapWrapperLondon(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opLogWrapperLondon(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opCreateWrapperLondon(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opCallWrapperLondon(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opCallCodeWrapperLondon(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opReturnWrapperLondon(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opDelegateCallWrapperLondon(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opCreate2WrapperLondon(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opStaticCallWrapperLondon(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opRevertWrapperLondon(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opSuicideWrapperLondon(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

var stringToWrapperLondon = map[string]executionWrapperFunc{
    "opStop": opStopWrapperLondon,
    "opAdd": opAddWrapperLondon,
    "opMul": opMulWrapperLondon,
    "opSub": opSubWrapperLondon,
    "opDiv": opDivWrapperLondon,
    "opSdiv": opSdivWrapperLondon,
    "opMod": opModWrapperLondon,
    "opSmod": opSmodWrapperLondon,
    "opAddmod": opAddmodWrapperLondon,
    "opMulmod": opMulmodWrapperLondon,
    "opExp": opExpWrapperLondon,
    "opSignExtend": opSignExtendWrapperLondon,
    "opLt": opLtWrapperLondon,
    "opGt": opGtWrapperLondon,
    "opSlt": opSltWrapperLondon,
    "opSgt": opSgtWrapperLondon,
    "opEq": opEqWrapperLondon,
    "opIszero": opIszeroWrapperLondon,
    "opAnd": opAndWrapperLondon,
    "opOr": opOrWrapperLondon,
    "opXor": opXorWrapperLondon,
    "opNot": opNotWrapperLondon,
    "opByte": opByteWrapperLondon,
    "opSHL": opSHLWrapperLondon,
    "opSHR": opSHRWrapperLondon,
    "opSAR": opSARWrapperLondon,
    "opSha3": opSha3WrapperLondon,
    "opAddress": opAddressWrapperLondon,
    "opBalance": opBalanceWrapperLondon,
    "opOrigin": opOriginWrapperLondon,
    "opCaller": opCallerWrapperLondon,
    "opCallValue": opCallValueWrapperLondon,
    "opCallDataLoad": opCallDataLoadWrapperLondon,
    "opCallDataSize": opCallDataSizeWrapperLondon,
    "opCallDataCopy": opCallDataCopyWrapperLondon,
    "opCodeSize": opCodeSizeWrapperLondon,
    "opCodeCopy": opCodeCopyWrapperLondon,
    "opGasprice": opGaspriceWrapperLondon,
    "opExtCodeSize": opExtCodeSizeWrapperLondon,
    "opExtCodeCopy": opExtCodeCopyWrapperLondon,
    "opReturnDataSize": opReturnDataSizeWrapperLondon,
    "opReturnDataCopy": opReturnDataCopyWrapperLondon,
    "opExtCodeHash": opExtCodeHashWrapperLondon,
    "opBlockhash": opBlockhashWrapperLondon,
    "opCoinbase": opCoinbaseWrapperLondon,
    "opTimestamp": opTimestampWrapperLondon,
    "opNumber": opNumberWrapperLondon,
    "opDifficulty": opDifficultyWrapperLondon,
    "opGasLimit": opGasLimitWrapperLondon,
    "opChainID": opChainIDWrapperLondon,
    "opSelfBalance": opSelfBalanceWrapperLondon,
    "opBaseFee": opBaseFeeWrapperLondon,
    "opPop": opPopWrapperLondon,
    "opMload": opMloadWrapperLondon,
    "opMstore": opMstoreWrapperLondon,
    "opMstore8": opMstore8WrapperLondon,
    "opSload": opSloadWrapperLondon,
    "opSstore": opSstoreWrapperLondon,
    "opJump": opJumpWrapperLondon,
    "opJumpi": opJumpiWrapperLondon,
    "opPc": opPcWrapperLondon,
    "opMsize": opMsizeWrapperLondon,
    "opGas": opGasWrapperLondon,
    "opJumpdest": opJumpdestWrapperLondon,
    "opPush1": opPush1WrapperLondon,
    "opPush": opPushWrapperLondon,
    "opDup": opDupWrapperLondon,
    "opSwap": opSwapWrapperLondon,
    "opLog": opLogWrapperLondon,
    "opCreate": opCreateWrapperLondon,
    "opCall": opCallWrapperLondon,
    "opCallCode": opCallCodeWrapperLondon,
    "opReturn": opReturnWrapperLondon,
    "opDelegateCall": opDelegateCallWrapperLondon,
    "opCreate2": opCreate2WrapperLondon,
    "opStaticCall": opStaticCallWrapperLondon,
    "opRevert": opRevertWrapperLondon,
    "opSuicide": opSuicideWrapperLondon,
}
