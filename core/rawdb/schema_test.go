// Copyright 2020 The go-ethereum Authors
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

package rawdb

import (
	"bytes"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

var (
	bytes4  = []byte{0x00, 0x01, 0x02, 0x03}
	bytes20 = []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f, 0x00, 0x01, 0x02, 0x03}
	bytes32 = []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f, 0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f}
	bytes63 = []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f, 0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f, 0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f, 0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e}
	bytes64 = []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f, 0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f, 0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f, 0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f}
	bytes65 = []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f, 0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f, 0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f, 0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f, 0x00}
)

func TestIsLegacyTrieNode(t *testing.T) {
	tests := []struct {
		name      string
		inputData []byte
		inputKey  []byte
		expected  bool
	}{
		{
			name:     "empty",
			inputKey: []byte{},
			expected: false,
		},
		{
			name:     "non-legacy (too short)",
			inputKey: []byte{0x00, 0x01, 0x02, 0x03},
			expected: false,
		},
		{
			name:      "legacy",
			inputData: []byte{0x00, 0x01, 0x02, 0x03},
			inputKey:  crypto.Keccak256([]byte{0x00, 0x01, 0x02, 0x03}),
			expected:  true,
		},
		{
			name:     "non-legacy (too long)",
			inputKey: []byte{0x00, 0x01, 0x02, 0x03, 0x00, 0x00, 0x00, 0x00, 0x00},
			expected: false,
		},
		{
			name:      "non-legacy (key is not hash of data)",
			inputData: []byte{0x00, 0x01, 0x02, 0x03},
			inputKey:  crypto.Keccak256([]byte{0x00, 0x01, 0x02, 0x04}),
			expected:  false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if actual := IsLegacyTrieNode(test.inputKey, test.inputData); actual != test.expected {
				t.Errorf("expected %v, got %v", test.expected, actual)
			}
		})
	}
}

func TestResolveAccountTrieNodeKey(t *testing.T) {
	tests := []struct {
		name          string
		inputKey      []byte
		expectedCheck bool
		expectedKey   []byte
	}{
		{
			name:          "empty",
			inputKey:      []byte{},
			expectedCheck: false,
			expectedKey:   nil,
		},
		{
			name:          "non account prefixed",
			inputKey:      bytes4,
			expectedCheck: false,
			expectedKey:   nil,
		},
		{
			name:          "storage prefixed",
			inputKey:      append(TrieNodeStoragePrefix, bytes4...),
			expectedCheck: false,
			expectedKey:   nil,
		},
		{
			name:          "account prefixed length 4",
			inputKey:      accountTrieNodeKey(bytes4),
			expectedCheck: true,
			expectedKey:   bytes4,
		},
		{
			name:          "account prefixed length 20",
			inputKey:      accountTrieNodeKey(bytes20),
			expectedCheck: true,
			expectedKey:   bytes20,
		},
		{
			name:          "account prefixed length 63",
			inputKey:      accountTrieNodeKey(bytes63),
			expectedCheck: true,
			expectedKey:   bytes63,
		},
		{
			name:          "account prefixed length 64",
			inputKey:      accountTrieNodeKey(bytes64),
			expectedCheck: false,
			expectedKey:   nil,
		},
		{
			name:          "account prefixed length 65",
			inputKey:      accountTrieNodeKey(bytes65),
			expectedCheck: false,
			expectedKey:   nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if check, key := ResolveAccountTrieNodeKey(test.inputKey); check != test.expectedCheck || !bytes.Equal(key, test.expectedKey) {
				t.Errorf("expected %v, %v, got %v, %v", test.expectedCheck, test.expectedKey, check, key)
			}
		})
	}
}

func TestResolveStorageTrieNode(t *testing.T) {
	tests := []struct {
		name          string
		inputKey      []byte
		expectedCheck bool
		expectedHash  common.Hash
		expectedKey   []byte
	}{
		{
			name:          "empty",
			inputKey:      []byte{},
			expectedCheck: false,
			expectedHash:  common.Hash{},
			expectedKey:   nil,
		},
		{
			name:          "non storage prefixed",
			inputKey:      []byte{0x00, 0x01, 0x02, 0x03},
			expectedCheck: false,
			expectedHash:  common.Hash{},
			expectedKey:   nil,
		},
		{
			name:          "account prefixed",
			inputKey:      accountTrieNodeKey(bytes4),
			expectedCheck: false,
			expectedHash:  common.Hash{},
			expectedKey:   nil,
		},
		{
			name:          "storage prefixed hash 20 length 4",
			inputKey:      append(append(TrieNodeStoragePrefix, bytes20...), bytes4...),
			expectedCheck: false,
			expectedHash:  common.Hash{},
			expectedKey:   nil,
		},
		{
			name:          "storage prefixed hash 32 length 4",
			inputKey:      storageTrieNodeKey(common.BytesToHash(bytes32), bytes4),
			expectedCheck: true,
			expectedHash:  common.BytesToHash(bytes32),
			expectedKey:   bytes4,
		},
		{
			name:          "storage prefixed hash 32 length 20",
			inputKey:      storageTrieNodeKey(common.BytesToHash(bytes20), bytes20),
			expectedCheck: true,
			expectedHash:  common.BytesToHash(bytes20),
			expectedKey:   bytes20,
		},
		{
			name:          "storage prefixed hash 32 length 63",
			inputKey:      storageTrieNodeKey(common.BytesToHash(bytes65), bytes63),
			expectedCheck: true,
			expectedHash:  common.BytesToHash(bytes65),
			expectedKey:   bytes63,
		},
		{
			name:          "storage prefixed hash 32 length 64",
			inputKey:      storageTrieNodeKey(common.BytesToHash(bytes32), bytes64),
			expectedCheck: false,
			expectedHash:  common.Hash{},
			expectedKey:   nil,
		},
		{
			name:          "storage prefixed hash 32 length 65",
			inputKey:      storageTrieNodeKey(common.BytesToHash(bytes32), bytes65),
			expectedCheck: false,
			expectedHash:  common.Hash{},
			expectedKey:   nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if check, hash, key := ResolveStorageTrieNode(test.inputKey); check != test.expectedCheck || !bytes.Equal(key, test.expectedKey) || hash != test.expectedHash {
				t.Errorf("expected %v, %v, %v, got %v, %v, %v", test.expectedCheck, test.expectedHash, test.expectedKey, check, hash, key)
			}
		})
	}
}
