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

package tests

import (
	"bufio"
	"bytes"
	"fmt"
	"reflect"
	"testing"

	"github.com/ethereum/go-ethereum/eth/tracers/logger"

	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/state/snapshot"
	"github.com/ethereum/go-ethereum/core/vm"
)

func TestState(t *testing.T) {
	t.Parallel()

	st := new(testMatcher)
	// Long tests:
	st.slow(`^stAttackTest/ContractCreationSpam`)
	st.slow(`^stBadOpcode/badOpcodes`)
	st.slow(`^stPreCompiledContracts/modexp`)
	st.slow(`^stQuadraticComplexityTest/`)
	st.slow(`^stStaticCall/static_Call50000`)
	st.slow(`^stStaticCall/static_Return50000`)
	st.slow(`^stSystemOperationsTest/CallRecursiveBomb`)
	st.slow(`^stTransactionTest/Opcodes_TransactionInit`)

	// Very time consuming
	st.skipLoad(`^stTimeConsuming/`)
	st.skipLoad(`.*vmPerformance/loop.*`)

	// Uses 1GB RAM per tested fork
	st.skipLoad(`^stStaticCall/static_Call1MB`)

	// Broken tests:
	// Expected failures:
	st.skipLoad(`^stPreCompiledContracts`) // Ronin contract are not match with Ethereum

	//st.fails(`^stRevertTest/RevertPrecompiledTouch(_storage)?\.json/Byzantium/0`, "bug in test")
	//st.fails(`^stRevertTest/RevertPrecompiledTouch(_storage)?\.json/Byzantium/3`, "bug in test")
	//st.fails(`^stRevertTest/RevertPrecompiledTouch(_storage)?\.json/Constantinople/0`, "bug in test")
	//st.fails(`^stRevertTest/RevertPrecompiledTouch(_storage)?\.json/Constantinople/3`, "bug in test")
	//st.fails(`^stRevertTest/RevertPrecompiledTouch(_storage)?\.json/ConstantinopleFix/0`, "bug in test")
	//st.fails(`^stRevertTest/RevertPrecompiledTouch(_storage)?\.json/ConstantinopleFix/3`, "bug in test")

	// For Istanbul, older tests were moved into LegacyTests
	for _, dir := range []string{
		stateTestDir,
		legacyStateTestDir,
	} {
		st.walk(t, dir, func(t *testing.T, name string, test *StateTest) {
			for _, subtest := range test.Subtests() {
				subtest := subtest
				key := fmt.Sprintf("%s/%d", subtest.Fork, subtest.Index)

				t.Run(key+"/hash/trie", func(t *testing.T) {
					withTrace(t, test.gasLimit(subtest), func(vmconfig vm.Config) error {
						var result error
						test.Run(subtest, vmconfig, false, rawdb.HashScheme, func(err error, snaps *snapshot.Tree, state *state.StateDB) {
							result = st.checkFailure(t, err)
						})
						return result
					})
				})
				t.Run(key+"/hash/snap", func(t *testing.T) {
					withTrace(t, test.gasLimit(subtest), func(vmconfig vm.Config) error {
						var result error
						test.Run(subtest, vmconfig, true, rawdb.HashScheme, func(err error, snaps *snapshot.Tree, state *state.StateDB) {
							if snaps != nil && state != nil {
								if _, err := snaps.Journal(state.IntermediateRoot(false)); err != nil {
									result = err
									return
								}
							}
							result = st.checkFailure(t, err)
						})
						return result
					})
				})
				t.Run(key+"/path/trie", func(t *testing.T) {
					withTrace(t, test.gasLimit(subtest), func(vmconfig vm.Config) error {
						var result error
						test.Run(subtest, vmconfig, false, rawdb.PathScheme, func(err error, snaps *snapshot.Tree, state *state.StateDB) {
							result = st.checkFailure(t, err)
						})
						return result
					})
				})
				t.Run(key+"/path/snap", func(t *testing.T) {
					withTrace(t, test.gasLimit(subtest), func(vmconfig vm.Config) error {
						var result error
						test.Run(subtest, vmconfig, true, rawdb.PathScheme, func(err error, snaps *snapshot.Tree, state *state.StateDB) {
							if snaps != nil && state != nil {
								if _, err := snaps.Journal(state.IntermediateRoot(false)); err != nil {
									result = err
									return
								}
							}
							result = st.checkFailure(t, err)
						})
						return result
					})
				})
			}
		})
	}
}

// Transactions with gasLimit above this value will not get a VM trace on failure.
const traceErrorLimit = 400000

func withTrace(t *testing.T, gasLimit uint64, test func(vm.Config) error) {
	// Use config from command line arguments.
	config := vm.Config{}
	err := test(config)
	if err == nil {
		return
	}

	// Test failed, re-run with tracing enabled.
	t.Error(err)
	if gasLimit > traceErrorLimit {
		t.Log("gas limit too high for EVM trace")
		return
	}
	buf := new(bytes.Buffer)
	w := bufio.NewWriter(buf)
	tracer := logger.NewJSONLogger(&logger.Config{}, w)
	config.Debug, config.Tracer = true, tracer
	err2 := test(config)
	if !reflect.DeepEqual(err, err2) {
		t.Errorf("different error for second run: %v", err2)
	}
	w.Flush()
	if buf.Len() == 0 {
		t.Log("no EVM operation logs generated")
	} else {
		t.Log("EVM operation log:\n" + buf.String())
	}
	// t.Logf("EVM output: 0x%x", tracer.Output())
	// t.Logf("EVM error: %v", tracer.Error())
}
