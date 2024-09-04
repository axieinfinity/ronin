
package vm

import "github.com/ethereum/go-ethereum/common/math"


func opStopWrapperIstanbul(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opAddWrapperIstanbul(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opMulWrapperIstanbul(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opSubWrapperIstanbul(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opDivWrapperIstanbul(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opSdivWrapperIstanbul(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opModWrapperIstanbul(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opSmodWrapperIstanbul(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opAddmodWrapperIstanbul(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opMulmodWrapperIstanbul(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opExpWrapperIstanbul(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opSignExtendWrapperIstanbul(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opLtWrapperIstanbul(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opGtWrapperIstanbul(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opSltWrapperIstanbul(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opSgtWrapperIstanbul(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opEqWrapperIstanbul(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opIszeroWrapperIstanbul(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opAndWrapperIstanbul(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opOrWrapperIstanbul(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opXorWrapperIstanbul(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opNotWrapperIstanbul(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opByteWrapperIstanbul(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opSHLWrapperIstanbul(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opSHRWrapperIstanbul(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opSARWrapperIstanbul(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opSha3WrapperIstanbul(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opAddressWrapperIstanbul(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opBalanceWrapperIstanbul(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opOriginWrapperIstanbul(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opCallerWrapperIstanbul(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opCallValueWrapperIstanbul(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opCallDataLoadWrapperIstanbul(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opCallDataSizeWrapperIstanbul(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opCallDataCopyWrapperIstanbul(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opCodeSizeWrapperIstanbul(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opCodeCopyWrapperIstanbul(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opGaspriceWrapperIstanbul(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opExtCodeSizeWrapperIstanbul(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opExtCodeCopyWrapperIstanbul(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opReturnDataSizeWrapperIstanbul(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opReturnDataCopyWrapperIstanbul(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opExtCodeHashWrapperIstanbul(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opBlockhashWrapperIstanbul(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opCoinbaseWrapperIstanbul(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opTimestampWrapperIstanbul(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opNumberWrapperIstanbul(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opDifficultyWrapperIstanbul(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opGasLimitWrapperIstanbul(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opChainIDWrapperIstanbul(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opSelfBalanceWrapperIstanbul(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opPopWrapperIstanbul(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opMloadWrapperIstanbul(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opMstoreWrapperIstanbul(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opMstore8WrapperIstanbul(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opSloadWrapperIstanbul(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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
    if !contract.UseGas(800) {
        return nil, ErrOutOfGas
    }
    res, err = operation.execute(pc, in, callContext)
    if err != nil {
        return nil, err
    }
    *pc++
    return res, nil
}

func opSstoreWrapperIstanbul(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opJumpWrapperIstanbul(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opJumpiWrapperIstanbul(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opPcWrapperIstanbul(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opMsizeWrapperIstanbul(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opGasWrapperIstanbul(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opJumpdestWrapperIstanbul(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opPush1WrapperIstanbul(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opPushWrapperIstanbul(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opDupWrapperIstanbul(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opSwapWrapperIstanbul(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opLogWrapperIstanbul(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opCreateWrapperIstanbul(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opCallWrapperIstanbul(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opCallCodeWrapperIstanbul(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opReturnWrapperIstanbul(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opDelegateCallWrapperIstanbul(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opCreate2WrapperIstanbul(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opStaticCallWrapperIstanbul(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opRevertWrapperIstanbul(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opSuicideWrapperIstanbul(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
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

func opBaseFeeWrapperIstanbul(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
	return nil, nil
}

var stringToWrapperIstanbul = map[string]executionWrapperFunc{
    "opStop": opStopWrapperIstanbul,
    "opAdd": opAddWrapperIstanbul,
    "opMul": opMulWrapperIstanbul,
    "opSub": opSubWrapperIstanbul,
    "opDiv": opDivWrapperIstanbul,
    "opSdiv": opSdivWrapperIstanbul,
    "opMod": opModWrapperIstanbul,
    "opSmod": opSmodWrapperIstanbul,
    "opAddmod": opAddmodWrapperIstanbul,
    "opMulmod": opMulmodWrapperIstanbul,
    "opExp": opExpWrapperIstanbul,
    "opSignExtend": opSignExtendWrapperIstanbul,
    "opLt": opLtWrapperIstanbul,
    "opGt": opGtWrapperIstanbul,
    "opSlt": opSltWrapperIstanbul,
    "opSgt": opSgtWrapperIstanbul,
    "opEq": opEqWrapperIstanbul,
    "opIszero": opIszeroWrapperIstanbul,
    "opAnd": opAndWrapperIstanbul,
    "opOr": opOrWrapperIstanbul,
    "opXor": opXorWrapperIstanbul,
    "opNot": opNotWrapperIstanbul,
    "opByte": opByteWrapperIstanbul,
    "opSHL": opSHLWrapperIstanbul,
    "opSHR": opSHRWrapperIstanbul,
    "opSAR": opSARWrapperIstanbul,
    "opSha3": opSha3WrapperIstanbul,
    "opAddress": opAddressWrapperIstanbul,
    "opBalance": opBalanceWrapperIstanbul,
    "opOrigin": opOriginWrapperIstanbul,
    "opCaller": opCallerWrapperIstanbul,
    "opCallValue": opCallValueWrapperIstanbul,
    "opCallDataLoad": opCallDataLoadWrapperIstanbul,
    "opCallDataSize": opCallDataSizeWrapperIstanbul,
    "opCallDataCopy": opCallDataCopyWrapperIstanbul,
    "opCodeSize": opCodeSizeWrapperIstanbul,
    "opCodeCopy": opCodeCopyWrapperIstanbul,
    "opGasprice": opGaspriceWrapperIstanbul,
    "opExtCodeSize": opExtCodeSizeWrapperIstanbul,
    "opExtCodeCopy": opExtCodeCopyWrapperIstanbul,
    "opReturnDataSize": opReturnDataSizeWrapperIstanbul,
    "opReturnDataCopy": opReturnDataCopyWrapperIstanbul,
    "opExtCodeHash": opExtCodeHashWrapperIstanbul,
    "opBlockhash": opBlockhashWrapperIstanbul,
    "opCoinbase": opCoinbaseWrapperIstanbul,
    "opTimestamp": opTimestampWrapperIstanbul,
    "opNumber": opNumberWrapperIstanbul,
    "opDifficulty": opDifficultyWrapperIstanbul,
    "opGasLimit": opGasLimitWrapperIstanbul,
    "opChainID": opChainIDWrapperIstanbul,
    "opSelfBalance": opSelfBalanceWrapperIstanbul,
    "opPop": opPopWrapperIstanbul,
    "opMload": opMloadWrapperIstanbul,
    "opMstore": opMstoreWrapperIstanbul,
    "opMstore8": opMstore8WrapperIstanbul,
    "opSload": opSloadWrapperIstanbul,
    "opSstore": opSstoreWrapperIstanbul,
    "opJump": opJumpWrapperIstanbul,
    "opJumpi": opJumpiWrapperIstanbul,
    "opPc": opPcWrapperIstanbul,
    "opMsize": opMsizeWrapperIstanbul,
    "opGas": opGasWrapperIstanbul,
    "opJumpdest": opJumpdestWrapperIstanbul,
    "opPush1": opPush1WrapperIstanbul,
    "opPush": opPushWrapperIstanbul,
    "opDup": opDupWrapperIstanbul,
    "opSwap": opSwapWrapperIstanbul,
    "opLog": opLogWrapperIstanbul,
    "opCreate": opCreateWrapperIstanbul,
    "opCall": opCallWrapperIstanbul,
    "opCallCode": opCallCodeWrapperIstanbul,
    "opReturn": opReturnWrapperIstanbul,
    "opDelegateCall": opDelegateCallWrapperIstanbul,
    "opCreate2": opCreate2WrapperIstanbul,
    "opStaticCall": opStaticCallWrapperIstanbul,
    "opRevert": opRevertWrapperIstanbul,
    "opSuicide": opSuicideWrapperIstanbul,
    "opBaseFee": opBaseFeeWrapperIstanbul,
}
