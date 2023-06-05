package merkle

import (
	"bytes"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
)

// StateKey implements a key for the historical state map.
type StateKey struct {
	h     uint32
	x     uint32
	count uint32
}

// HistoricalTree implements a merkle tree with the ability to generate historical
// state of the tree. This includes historical roots, as well as historical proofs.
type HistoricalTree struct {
	// state[stateKey] is the value for a tree element:
	//   - With [height = stakeKey.h] (increasing from leafs to root)
	//   - With [x-coord = stateKey.x] (increasing from older leafs to newer)
	//   - When stateKey.count leafs were inserted in the merkle tree
	state map[StateKey][]byte
	// zeroHashes[H] is the value for a tree element:
	//   - With [height = H] (increasing from leafs to root)
	//	 - That doesn't have any non-zero children
	zeroHashes [][]byte
	// treeCount is the total amount of inserted leafs
	treeCount uint32
	// treeHeight is the height of the merkle tree
	treeHeight uint32
}

/**
 * Store historical non-"zero" values of the FULL merkle tree.
 * Full merkle tree consists of 2**TREE_DEPTH "zero" leafs, which are
 * getting populated throughout time. Once a new leaf is added, all elements
 * in the merkle tree on the path from root to the leaf are updated.
 * The goal of this contract is to store only the significant values.
 *
 * merkleTree[H][X][N] is the value for the tree element:
 * - With [height = H] (increasing from leafs to root)
 * - With [x-coord = X] (increasing from older leafs to newer)
 * - When N leafs were inserted in the merkle tree
 *
 * 1. Height (H):
 * merkleTree[0] are the leafs
 * merkleTree[1] are keccak256(A, B) where A and B are leafs
 * ...
 * merkleTree[TREE_DEPTH] is the merkle root level
 *
 * 2. Coordinate (X):
 * A merkle tree can have up to 2**(32-H) elements on a level with height=H
 * Therefore:
 * merkleTree[0][0] is the first leaf
 * merkleTree[0][1] is the second leaf
 * merkleTree[1][0] is their parent
 * merkleTree[1][1] is parent of merkleTree[0][2] and merkleTree[0][3]
 * merkleTree[2][0] is parent of merkleTree[1][0] and merkleTree[1][1]
 * ...
 * merkleTree[TREE_DEPTH][0] is the merkle root
 *
 * 3. Historical state (N).
 * Every element of the full merkle tree has three chronological "stages".
 * a. Element value did not change after the latest leaf insertion. Meaning that
 *    all element's children are "zero" elements, and element itself is "zero".
 *    Requires: 0 <= N <= X*(2**H)
 * b. Element value changed after the latest leaf insertion. Meaning that
 *    at least one of the children is non-zero.
 *    Requires: X*(2**H) < N <= (X+1)*(2**H)
 * c. Element value stopped changing after the latest leaf insertion. Meaning that
 *    all element children are already non-zero.
 *    Requires: (X+1)*(2**H) < N
 *
 * Thus we actually need to store tree element value for N in range (X*(2**H), (X+1)*(2**H)]
 * The amount of "significant" values (stage b) is 2**H.
 *
 * We're using a map to avoid dealing with dynamic arrays.
 */

// MessageTreeHeight is the depth of the merkle tree that is used in the messaging contracts.
const MessageTreeHeight uint32 = 32

// SnapshotTreeHeight is the depth of the merkle tree that is used in the snapshot contracts.
const SnapshotTreeHeight uint32 = 6

// NewTree returns an empty Merkle Tree.
func NewTree(treeHeight uint32) *HistoricalTree {
	return &HistoricalTree{
		state:      make(map[StateKey][]byte),
		zeroHashes: generateZeroHashes(treeHeight),
		treeCount:  0,
		treeHeight: treeHeight,
	}
}

// NewTreeFromItems returns a new Merkle Tree from a slice of byte slices.
func NewTreeFromItems(items [][]byte, treeHeight uint32) *HistoricalTree {
	tree := NewTree(treeHeight)
	for _, item := range items {
		tree.Insert(item)
	}
	return tree
}

// BranchRoot calculates the merkle root given the item and the proof.
func BranchRoot(item []byte, index uint32, proof [][]byte, treeHeight uint32) ([]byte, error) {
	if len(proof) != int(treeHeight) {
		return nil, fmt.Errorf("incorrect proof length: %d; should be: %d", len(proof), treeHeight)
	}
	node := item
	for h := uint32(0); h < treeHeight; h++ {
		if (index>>h)&1 == 0 {
			// We were the left child
			node = getParent(node, proof[h])
		} else {
			// We were the right child
			node = getParent(proof[h], node)
		}
	}
	return node, nil
}

// VerifyMerkleProof verifies a Merkle branch against a root of a tree.
func VerifyMerkleProof(root, item []byte, index uint32, proof [][]byte, treeHeight uint32) bool {
	branchRoot, err := BranchRoot(item, index, proof, treeHeight)
	if err != nil {
		return false
	}
	return bytes.Equal(root, branchRoot)
}

// Insert inserts a new leaf into the merkle tree. This is done using O(1) time.
func (m *HistoricalTree) Insert(item []byte) {
	x := m.treeCount
	newCount := x + 1
	saveElementState(m, 0, x, newCount, item)
	for h := uint32(1); h <= m.treeHeight; h++ {
		// Traverse to parent
		x >>= 1
		// Children have [height = h - 1]
		// And X-coordinates [2 * x] and [2 * x + 1]
		leftChild := fetchTreeElementState(m, h-1, x<<1, newCount)
		rightChild := fetchTreeElementState(m, h-1, (x<<1)+1, newCount)
		parent := getParent(leftChild, rightChild)
		saveElementState(m, h, x, newCount, parent)
	}
	m.treeCount = newCount
}

// Items returns the list of items that were inserted in the Merkle tree.
func (m *HistoricalTree) Items() [][]byte {
	items := make([][]byte, m.treeCount)
	for x := uint32(0); x < m.treeCount; x++ {
		// H=0 is the leaf level.
		items[x] = fetchTreeElementState(m, 0, x, m.treeCount)
	}
	return items
}

// NumOfItems returns the amount of leafs inserted in the merkle tree.
func (m *HistoricalTree) NumOfItems() uint32 {
	return m.treeCount
}

// Item returns the inserted item with the given `index`.
func (m *HistoricalTree) Item(index uint32) ([]byte, error) {
	if index >= m.treeCount {
		return nil, fmt.Errorf("not enough leafs; inserted: %d, requested index: %d", m.treeCount, index)
	}
	// H=0 is the leaf level.
	return fetchTreeElementState(m, 0, index, m.treeCount), nil
}

// Root returns the merkle root of the tree after a certain amount of leafs were inserted.
// This is done using O(1) time.
func (m *HistoricalTree) Root(count uint32) ([]byte, error) {
	if count > m.treeCount {
		return nil, fmt.Errorf("not enough leafs; inserted: %d, requested root for count: %d", m.treeCount, count)
	}
	// H=m.treeHeight is the root level.
	return fetchTreeElementState(m, m.treeHeight, 0, count), nil
}

// MerkleProof returns the proof of inclusion:
//   - For leaf with given `index` MerkleProof
//   - At the time when `count` leafs have been inserted
//
// This is done using O(1) time.
func (m *HistoricalTree) MerkleProof(index uint32, count uint32) ([][]byte, error) {
	if count > m.treeCount {
		return nil, fmt.Errorf("not enough leafs; inserted: %d, requested proof for count: %d", m.treeCount, count)
	}
	if index >= count {
		return nil, fmt.Errorf("merkle index out of range; count: %d, requested proof for index: %d", count, index)
	}
	proof := make([][]byte, m.treeHeight)
	for h := uint32(0); h < m.treeHeight; h++ {
		// First, determine X-axis of the element's sibling
		siblingX := index ^ 1
		// Get sibling state at the time when `count` leafs were added
		proof[h] = fetchTreeElementState(m, h, siblingX, count)
		// Traverse to parent
		index >>= 1
	}
	return proof, nil
}

// generateZeroHashes returns the default "zero" values for elements from bottom to top (leaf to root).
func generateZeroHashes(treeHeight uint32) [][]byte {
	zeroHashes := make([][]byte, treeHeight+1)
	// zeroHashes[0] is bytes32(0).
	zeroHashes[0] = make([]byte, 32)
	// Calculate "zero" element value for other heights.
	// That is the value for an element in the merkle tree, when all their children are "zero".
	for h := uint32(0); h < treeHeight; h++ {
		zeroHashes[h+1] = getParent(zeroHashes[h], zeroHashes[h])
	}
	return zeroHashes
}

// fetchTreeElementState returns a tree element:
//   - With [height = H] (increasing from leafs to root)
//   - With [x-coord = X] (increasing from older leafs to newer)
//   - When `count` leafs were inserted in the merkle tree
func fetchTreeElementState(m *HistoricalTree, h uint32, x uint32, count uint32) []byte {
	// We do cast to uint64 here, as (1 << 32) overflows uint32
	firstChildLeafIndex := uint64(x) << h // x * (2**H)
	childLeafsAmount := uint64(1) << h    // 2**H
	switch {
	case uint64(count) <= firstChildLeafIndex:
		// Stage A: not enough leafs were inserted, element is still zero.
		return m.zeroHashes[h]
	case uint64(count) <= firstChildLeafIndex+childLeafsAmount:
		// Stage B: tree element was updated after last leaf insertion.
		key := StateKey{h, x, count}
		return m.state[key]
	default:
		// Stage C: tree element was not updated after last leaf insertion.
		// Use last saved value.
		key := StateKey{h, x, uint32(firstChildLeafIndex + childLeafsAmount)}
		return m.state[key]
	}
}

// getParent calculates a parent node in the merkle tree given its children.
func getParent(leftChild []byte, rightChild []byte) []byte {
	var leftChildB32, rightChildB32 [32]byte
	copy(leftChildB32[:], leftChild)
	copy(rightChildB32[:], rightChild)

	if leftChildB32 == [32]byte{} && rightChildB32 == [32]byte{} {
		return leftChild
	}

	return crypto.Keccak256(append(leftChild, rightChild...))
}

// saveElementState stores the historical value for a given tree node.
func saveElementState(m *HistoricalTree, h uint32, x uint32, count uint32, item []byte) {
	key := StateKey{h, x, count}
	m.state[key] = item
}
