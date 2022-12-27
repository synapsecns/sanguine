package merkle_test

import (
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/core/merkle"
)

const leafsAmount uint32 = 100

func TestMerkleTree(t *testing.T) {
	tree := merkle.NewTree()
	// Generate test leafs
	leafs := make([][]byte, leafsAmount)
	for i := uint32(0); i < leafsAmount; i++ {
		leafs[i] = fakeLeaf()
		fmt.Printf("%d: %v\n", i, leafs[i])
	}
	// Insert test leafs
	for i := uint32(0); i < leafsAmount; i++ {
		tree.Insert(leafs[i])
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
			Equal(t, len(proof), int(merkle.TreeDepth))
			for i := 0; i < len(proof); i++ {
				// Each element should be a bytes32 value.
				Equal(t, len(proof[i]), 32)
			}
			branchRoot, err := merkle.BranchRoot(leafs[index], index, proof)
			Nil(t, err)
			// Should match the correct bytes32 value for the root.
			Equal(t, root, branchRoot)
			True(t, merkle.VerifyMerkleProof(root, leafs[index], index, proof))
		}
	}
}

func fakeLeaf() []byte {
	leaf := make([]byte, 32)
	for i := 0; i < 32; i++ {
		leaf[i] = gofakeit.Uint8()
	}
	return leaf
}
