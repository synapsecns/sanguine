package merkle_test

import (
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/core/merkle"
)

const (
	leafsAmount uint32 = 100
)

func TestMerkleTree(t *testing.T) {
	tree := merkle.NewTree(merkle.MessageTreeHeight)
	// Generate test leafs
	leafs := make([][]byte, leafsAmount)
	for i := uint32(0); i < leafsAmount; i++ {
		leafs[i] = fakeLeaf()
		fmt.Printf("%d: %v\n", i, leafs[i])
	}
	// Insert test leafs
	for i := uint32(0); i < leafsAmount; i++ {
		tree.Insert(leafs[i])
		Equal(t, tree.NumOfItems(), i+1)
	}
	// Check Items()
	items := tree.Items()
	Equal(t, len(items), int(leafsAmount))
	for i := uint32(0); i < leafsAmount; i++ {
		Equal(t, items[i], leafs[i])
	}
	// Check Item()
	for i := uint32(0); i < leafsAmount; i++ {
		item, err := tree.Item(i)
		Nil(t, err)
		Equal(t, item, leafs[i])
	}
	// Check proofs
	for count := uint32(1); count <= leafsAmount; count++ {
		// Get root after `count` leafs were inserted
		root, err := tree.Root(count)
		Nil(t, err)
		// Root should be a bytes32 value.
		Equal(t, len(root), 32)
		for index := uint32(0); index < count; index++ {
			proof, err := tree.MerkleProof(index, count)
			Nil(t, err)
			// Proof should have length equal to `TreeDepth`.
			Equal(t, len(proof), int(merkle.MessageTreeHeight))
			for i := 0; i < len(proof); i++ {
				// Each element should be a bytes32 value.
				Equal(t, len(proof[i]), 32)
			}
			branchRoot, err := merkle.BranchRoot(leafs[index], index, proof, merkle.MessageTreeHeight)
			Nil(t, err)
			// Should match the correct bytes32 value for the root.
			Equal(t, root, branchRoot)
			True(t, merkle.VerifyMerkleProof(root, leafs[index], index, proof, merkle.MessageTreeHeight))
		}
	}
}

func TestIncorrectRequests(t *testing.T) {
	tree := merkle.NewTree(merkle.MessageTreeHeight)
	// Generate test leafs.
	leafs := make([][]byte, leafsAmount)
	for i := uint32(0); i < leafsAmount; i++ {
		leafs[i] = fakeLeaf()
	}
	// Insert test leafs.
	for i := uint32(0); i < leafsAmount; i++ {
		tree.Insert(leafs[i])
	}
	// Check Item() with index out of bound.
	item, err := tree.Item(leafsAmount)
	Nil(t, item)
	NotNil(t, err)
	// Check Root() with count out of bound.
	root, err := tree.Root(leafsAmount + 1)
	Nil(t, root)
	NotNil(t, err)
	// Check MerkleProof() with count out of bound.
	proof, err := tree.MerkleProof(0, leafsAmount+1)
	Nil(t, proof)
	NotNil(t, err)
	// Check MerkleProof() with index >= count.
	proof, err = tree.MerkleProof(10, 10)
	Nil(t, proof)
	NotNil(t, err)
	// Generate a valid proof.
	proof, err = tree.MerkleProof(0, leafsAmount)
	NotNil(t, proof)
	Nil(t, err)
	// Check BranchRoot() with incorrect proof length.
	proof = make([][]byte, merkle.MessageTreeHeight+1)
	branchRoot, err := merkle.BranchRoot(leafs[0], leafsAmount, proof, merkle.MessageTreeHeight)
	Nil(t, branchRoot)
	NotNil(t, err)
	// Get current and historical root.
	root, err = tree.Root(leafsAmount)
	NotNil(t, root)
	Nil(t, err)
	histRoot, err := tree.Root(leafsAmount - 1)
	NotNil(t, histRoot)
	Nil(t, err)
	// Check VerifyMerkleProof() with incorrect proof length.
	False(t, merkle.VerifyMerkleProof(root, leafs[0], 0, proof, merkle.MessageTreeHeight))
	// Generate proofs against current root, they should not work with the historical root.
	for i := uint32(0); i < leafsAmount; i++ {
		proof, err = tree.MerkleProof(i, leafsAmount)
		NotNil(t, proof)
		Nil(t, err)
		// Should not be able to prove against the historical root.
		False(t, merkle.VerifyMerkleProof(histRoot, leafs[i], i, proof, merkle.MessageTreeHeight))
		// Should be able to prove against the current root.
		True(t, merkle.VerifyMerkleProof(root, leafs[i], i, proof, merkle.MessageTreeHeight))
	}
}

func TestNewTreeFromItems(t *testing.T) {
	tree := merkle.NewTree(merkle.MessageTreeHeight)
	// Generate test leafs
	leafs := make([][]byte, leafsAmount)
	for i := uint32(0); i < leafsAmount; i++ {
		leafs[i] = fakeLeaf()
		fmt.Printf("%d: %v\n", i, leafs[i])
	}
	// Insert test leafs
	for i := uint32(0); i < leafsAmount; i++ {
		tree.Insert(leafs[i])
		Equal(t, tree.NumOfItems(), i+1)
	}
	// Get items and generate a new tree from them
	items := tree.Items()
	newTree := merkle.NewTreeFromItems(items, merkle.MessageTreeHeight)
	// Check that the number of items are the same
	Equal(t, tree.NumOfItems(), newTree.NumOfItems())
	// Check that the new tree has the same root
	root, err := tree.Root(leafsAmount)
	Nil(t, err)
	newRoot, err := newTree.Root(leafsAmount)
	Nil(t, err)
	Equal(t, root, newRoot)
}

func fakeLeaf() []byte {
	leaf := make([]byte, 32)
	for i := 0; i < 32; i++ {
		leaf[i] = gofakeit.Uint8()
	}
	return leaf
}
