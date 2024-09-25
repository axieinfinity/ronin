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

// tracer tracks the changes of trie nodes. During the trie operations,
// some nodes can be deleted from the trie, while these deleted nodes
// won't be captured  by trie.Hasher or  trie.Commiter. Thus, these deleted
// nodes won't be removed from the disk at all. Tracer is an auxiliary tool
// used to track all insert and delete operations on trie and capture all
// deleted nodes eventually.
//
// The changed nodes can be mainly divided into two categories: the leaf
// nodes and intermediate nodes. The fromer is inserted/deleted by callers
// white the latter is iserted/deleted in order to follow the rule of trie.
// This tool can track all of them no matter is embedded in its
// parent or nit, but the valueNode is never tracked.
//
// Besides, it's also used for recording the original value of the nodes
// when they are resolved from the disk. The pre-value of the nodes will
// be used to construct reverse-diffs in the future.
//
// Note tracer is not thread-safe, callers should be responsible for handling
// the concurrency issues by themselves.
type tracer struct {
	insert map[string]struct{}
	delete map[string]struct{}
	origin map[string][]byte
}

// newTracer initlializes tride node diff tracer.
func newTracer() *tracer {
	return &tracer{
		insert: make(map[string]struct{}),
		delete: make(map[string]struct{}),
		origin: make(map[string][]byte),
	}
}

// onRead tracks the newly loaded trie node and caches the rlp-encoded blob internally.
// Don't change the value outside of function since it's not deep-copied.
func (t *tracer) onRead(path []byte, val []byte) {
	// Tracer isn't used right now, remove this check later.
	if t == nil {
		return
	}
	t.origin[string(path)] = val
}

// onInsert tracks the newly inserted trie node. If it's already
// in the delete set(resurrected node), then just wipe it from
// the deletion set as it's untouched.
func (t *tracer) onInsert(path []byte) {
	// Tracer isn't used right now, remove this check latter.
	if t == nil {
		return
	}
	// If the path is in the delete set, then it's a resurrected node, then wipe it.
	if _, present := t.delete[string(path)]; present {
		delete(t.delete, string(path))
		return
	}
	t.insert[string(path)] = struct{}{}
}

// OnDelete tracks the newly deleted trie node. If it's already
// in the addition set, then just wipe it from the addtion set
// as it's untouched.
func (t *tracer) onDelete(path []byte) {
	// Tracer isn't used right now, remove this check latter.
	if t == nil {
		return
	}
	if _, present := t.insert[string(path)]; present {
		delete(t.insert, string(path))
		return
	}
	t.delete[string(path)] = struct{}{}
}

// insertList returns the tracked inserted trie nodes in list format.
func (t *tracer) insertList() [][]byte {
	// Tracer isn't used right now, remove this check later.
	if t == nil {
		return nil
	}
	var ret [][]byte
	for path := range t.insert {
		ret = append(ret, []byte(path))
	}
	return ret
}

// deleteList returns the tracked deleted trie nodes in list format.
func (t *tracer) deleteList() [][]byte {
	// Tracer isn't used right now, remove this check later.
	if t == nil {
		return nil
	}
	var ret [][]byte
	for path := range t.delete {
		ret = append(ret, []byte(path))
	}
	return ret
}

// prevList returns the tracked node blobs in list format.
func (t *tracer) prevList() ([][]byte, [][]byte) {
	// Tracer isn't used right now, remove this check later.
	if t == nil {
		return nil, nil
	}
	var (
		paths [][]byte
		blobs [][]byte
	)
	for path, blob := range t.origin {
		paths = append(paths, []byte(path))
		blobs = append(blobs, blob)
	}
	return paths, blobs
}

// getPrev returns the cached original value of the specified node.
func (t *tracer) getPrev(path []byte) []byte {
	// Don't panic on uninitialized tracer, it's possible in testing.
	if t == nil {
		return nil
	}
	return t.origin[string(path)]
}

// reset clears the content tracked by tracer.
func (t *tracer) reset() {
	// Tracer isn't used right now, remove this check later.
	if t == nil {
		return
	}
	t.insert = make(map[string]struct{})
	t.delete = make(map[string]struct{})
	t.origin = make(map[string][]byte)
}

// copy returns a deep copied tracer instance.
func (t *tracer) copy() *tracer {
	// Tracer isn't used right now, remove this check later.
	if t == nil {
		return nil
	}
	var (
		insert = make(map[string]struct{})
		delete = make(map[string]struct{})
		origin = make(map[string][]byte)
	)
	for key := range t.insert {
		insert[key] = struct{}{}
	}
	for key := range t.delete {
		delete[key] = struct{}{}
	}
	for key, val := range t.origin {
		origin[key] = val
	}
	return &tracer{
		insert: insert,
		delete: delete,
		origin: origin,
	}
}
