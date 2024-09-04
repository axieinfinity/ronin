package main

import (
	"encoding/json"
	"os"
	"text/template"

	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/params"
)

type InstructionSetJSON struct {
    Version string `json:"version"`
    InstructionSet []vm.OperationAttribute `json:"instruction_set"`
}
func ExportInstructionSet(version string) {
	// Get the instruction set
	var rules params.Rules
    switch {
    case version == "Homestead":
        rules = params.Rules{
            IsHomestead: true,
        }
    case version == "TangerineWhistle":
        rules = params.Rules{
            IsEIP150: true,
        }
    case version == "SpuriousDragon":
        rules = params.Rules{
            IsEIP158: true,
        }
    case version == "Byzantium":
        rules = params.Rules{
            IsByzantium: true,
        }
    case version == "Constantinople":
        rules = params.Rules{
            IsConstantinople: true,
        }
    case version == "Istanbul":
        rules = params.Rules{
            IsIstanbul: true,
        }
    case version == "Berlin":
        rules = params.Rules{
            IsBerlin: true,
        }
    case version == "London":
        rules = params.Rules{
            IsLondon: true,
        }
    default:
        // frontier by default
        rules = params.Rules{
            
        }
    }
	instructionSet := vm.ExportInstructionSet(rules)

	// Export the instruction set to json
	file, err := os.Create("instruction_set.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	err = encoder.Encode(
        InstructionSetJSON{
            Version: version,
            InstructionSet: instructionSet,
        },
    )
	if err != nil {
		panic(err)
	}
}

func ImportInstructionSet() InstructionSetJSON {
	// Import the instruction set from json
	file, err := os.Open("instruction_set.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
    var instructionSetJSON InstructionSetJSON
	err = decoder.Decode(&instructionSetJSON)
	if err != nil {
		panic(err)
	}
	return instructionSetJSON
}

func OpCodeWrapperGenerator() {
	// Get instruction set
	instructionSetJSON := ImportInstructionSet()

	// Define the template
	tmpl := template.Must(template.New("opcode").Parse(`
package vm

import "github.com/ethereum/go-ethereum/common/math"
{{$parent := .}}
{{range .InstructionSet}}
func {{.Name}}Wrapper{{$parent.Version}}(operation *operation, pc *uint64, in *EVMInterpreter, callContext *ScopeContext) ([]byte, error) {
{{- if .Supported}}
	{{- if or .MemorySize .DynamicCost}}
    mem := callContext.Memory
	{{- end}}
    stack := callContext.Stack
    contract := callContext.Contract
    var err error
    var res []byte	

    in.evm.Context.Counter++

    if sLen := stack.len(); sLen < {{.MinStack}} {
        return nil, &ErrStackUnderflow{stackLen: sLen, required: {{.MinStack}}}
    } else if sLen > {{.MaxStack}} {
        return nil, &ErrStackOverflow{stackLen: sLen, limit: {{.MaxStack}}}
    }
    {{- if or .Writes (eq .Name "opCall")}}
    if in.readOnly && in.evm.chainRules.IsByzantium {
        {{- if .Writes}}
        return nil, ErrWriteProtection
        {{- else if (eq .Name "opCall")}}
        if stack.Back(2).Sign() != 0 {
            return nil, ErrWriteProtection
        }
        {{- end}}
    }
    {{- end}}
    {{- if not .DynamicCost}}
    if !contract.UseGas({{.ConstantGas}}) {
        return nil, ErrOutOfGas
    }
    {{- end}}
    {{- if or .MemorySize .DynamicCost}}
    var memorySize uint64
    {{- end}}
    {{- if .MemorySize}}
    memSize, overflow := operation.memorySize(stack)
    if overflow {
        return nil, ErrGasUintOverflow
    }
    if memorySize, overflow = math.SafeMul(toWordSize(memSize), 32); overflow {
        return nil, ErrGasUintOverflow
    }
    {{- end}}
    {{- if .DynamicCost}}
    var dynamicCost uint64
    dynamicCost, err = operation.dynamicGas(in.evm, contract, stack, mem, memorySize)
    if err != nil || !contract.UseGas({{.ConstantGas}} + dynamicCost) {
        return nil, ErrOutOfGas
    }
    {{- end}}
    {{- if or .MemorySize .DynamicCost}}
    if memorySize > 0 {
        mem.Resize(memorySize)
    }
    {{- end}}
    res, err = operation.execute(pc, in, callContext)
    {{- if .Returns}}
    in.returnData = res
    {{- end}}
    if err != nil {
        return nil, err
    }
    {{- if .Reverts}}
    return res, ErrExecutionReverted
    {{- end}}
    {{- if and (not .Jumps) (not .Reverts)}}
    *pc++
    {{- end}}
    {{- if not .Reverts}}
    return res, nil
    {{- end}}
{{- else}}
	return nil, nil
{{- end}}
}
{{end}}
var stringToWrapper{{$parent.Version}} = map[string]executionWrapperFunc{
{{- range .InstructionSet}}
    "{{.Name}}": {{.Name}}Wrapper{{$parent.Version}},{{end}}
}
`))

	// Create the source file
    file, err := os.Create("instruction_wrapper_" + instructionSetJSON.Version + ".go")
    if err != nil {
        panic(err)
    }
    defer file.Close()

	// Execute the template and write to the source file
    err = tmpl.Execute(file, instructionSetJSON)
    if err != nil {
        panic(err)
    }
}

func main() {
    if len(os.Args) > 1 {
        version := os.Args[1]
        ExportInstructionSet(version)
        OpCodeWrapperGenerator()
    } else {
        // export all by default
        versions := []string{"Homestead", "TangerineWhistle", "SpuriousDragon", "Byzantium", "Constantinople", "Istanbul", "Berlin", "London", "Frontier"}
        for _, version := range versions {
            ExportInstructionSet(version)
            OpCodeWrapperGenerator()
        }
    }
}
