// Copyright 2022 The go-ethereum Authors
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

package trie

import (
	"fmt"
	"sort"
	"strings"

	"github.com/ethereum/go-ethereum/common"
)

// memoryNode is all the information we know about a single cached trie node
// in the memory.
type memoryNode struct {
	hash common.Hash // Node hash, computed by hashing rlp value, empty for deleted nodes
	node []byte      // Encoded node blob, nil for deleted nodes
}

// memorySize returns the total memory size used by this node.
// nolint:unused
func (n *memoryNode) memorySize(pathlen int) int {
	return len(n.node) + common.HashLength + pathlen
}

// isDeleted returns the indicator if the node is marked as deleted.
func (n *memoryNode) isDeleted() bool {
	return n.hash == (common.Hash{})
}

// rlp returns the raw rlp encoded blob of the cached trie node, either directly
// from the cache, or by regenerating it from the collapsed node.
// nolint:unused
func (n *memoryNode) rlp() []byte {
	return n.node
}

// obj returns the decoded and expanded trie node, either directly from the cache,
// or by regenerating it from the rlp encoded blob.
// nolint:unused
func (n *memoryNode) obj() node {
	return mustDecodeNode(n.hash[:], n.node)
}

// nodeWithPrev wraps the memoryNode with the previous node value.
type nodeWithPrev struct {
	*memoryNode
	prev []byte // RLP-encoded previous value, nil means it's non-existent
}

// unwrap returns the internal memoryNode object.
// nolint:unused
func (n *nodeWithPrev) unwrap() *memoryNode {
	return n.memoryNode
}

// memorySize returns the total memory size used by this node. It overloads
// the function in memoryNode by counting the size of previous value as well.
// nolint: unused
func (n *nodeWithPrev) memorySize(key int) int {
	return n.memoryNode.memorySize(key) + len(n.prev)
}

// NodeSet contains all dirty nodes collected during the commit operation
// Each node is keyed by path. It's not the thread-safe to use.
type NodeSet struct {
	owner   common.Hash // the identifier of the trie
	leaves  []*leaf     // the list of dirty leaves
	updates int         // the count of updated and inserted nodes
	deletes int         // the count of deleted nodes

	// The set of all dirty nodes. Dirty nodes include newly inserted nodes,
	// deleted nodes and updated nodes. The original value of the newly
	// inserted node must be nil, and the original value of the other two
	// types must be non-nil.
	nodes map[string]*nodeWithPrev
}

// NewNodeSet initializes an empty node set to be used for tracking dirty nodes
// from a specific account or storage trie. The owner is zero for the account
// trie and the owning account address hash for storage tries.
func NewNodeSet(owner common.Hash) *NodeSet {
	return &NodeSet{
		owner: owner,
		nodes: make(map[string]*nodeWithPrev),
	}
}

// forEachWithOrder iterates the dirty nodes with the order from bottom to top,
// right to left, nodes with the longest path will be iterated first.
func (set *NodeSet) forEachWithOrder(callback func(path string, n *memoryNode)) {
	var paths sort.StringSlice
	for path := range set.nodes {
		paths = append(paths, path)
	}
	// Bottom-up, longest path first
	sort.Sort(sort.Reverse(paths))
	for _, path := range paths {
		callback(path, set.nodes[path].unwrap())
	}
}

// addNode adds the provided dirty node into set.
func (set *NodeSet) addNode(path []byte, n *nodeWithPrev) {
	if n.isDeleted() {
		set.deletes += 1
	} else {
		set.updates += 1
	}
	set.nodes[string(path)] = n
}

// addLeaf collects the provided leaf node into set.
func (set *NodeSet) addLeaf(leaf *leaf) {
	set.leaves = append(set.leaves, leaf)
}

// Size returns the number of updated and deleted nodes contained in the set.
func (set *NodeSet) Size() (int, int) {
	return set.updates, set.deletes
}

// Hashes returns the hashes of all updated nodes.
func (set *NodeSet) Hashes() []common.Hash {
	var ret []common.Hash
	for _, node := range set.nodes {
		ret = append(ret, node.hash)
	}
	return ret
}

// Summary returns a string-representation of the NodeSet.
func (set *NodeSet) Summary() string {
	var out = new(strings.Builder)
	fmt.Fprintf(out, "nodeset owner: %v\n", set.owner)
	if set.nodes != nil {
		for path, n := range set.nodes {
			// Deletion
			if n.isDeleted() {
				fmt.Fprintf(out, "  [-]: %x prev: %x\n", path, n.prev)
				continue
			}
			// Insertion
			if len(n.prev) == 0 {
				fmt.Fprintf(out, "  [+]: %x -> %v\n", path, n.hash)
				continue
			}
			// Update
			fmt.Fprintf(out, "  [*]: %x -> %v prev: %x\n", path, n.hash, n.prev)
		}
	}
	for _, n := range set.leaves {
		fmt.Fprintf(out, "[leaf]: %v\n", n)
	}
	return out.String()
}

// MergedNodeSet represents a merged dirty node set for a group of tries.
type MergedNodeSet struct {
	sets map[common.Hash]*NodeSet
}

// NewMergedNodeSet initializes an empty merged set.
func NewMergedNodeSet() *MergedNodeSet {
	return &MergedNodeSet{sets: make(map[common.Hash]*NodeSet)}
}

// NewWithNodeSet constructs a merged nodeset with the provided single set.
func NewWithNodeSet(set *NodeSet) *MergedNodeSet {
	merged := NewMergedNodeSet()
	merged.Merge(set)
	return merged
}

// Merge merges the provided dirty nodes of a trie into the set. The assumption
// is held that no duplicated set belonging to the same trie will be merged twice.
func (set *MergedNodeSet) Merge(other *NodeSet) error {
	_, present := set.sets[other.owner]
	if present {
		return fmt.Errorf("duplicate trie for owner %#x", other.owner)
	}
	set.sets[other.owner] = other
	return nil
}
