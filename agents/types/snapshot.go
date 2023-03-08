package types

import (
	"github.com/synapsecns/sanguine/core/merkle"
	"math"
)

// Snapshot is the snapshot interface.
type Snapshot interface {
	// States are the states of the snapshot.
	States() []State

	// SnapshotRootAndProofs returns the snapshot root, calculated from the states, as well as each state's proof.
	SnapshotRootAndProofs() ([32]byte, [][][]byte, error)
}

type snapshot struct {
	states []State
}

// NewSnapshot creates a new snapshot.
func NewSnapshot(states []State) Snapshot {
	return &snapshot{
		states: states,
	}
}

func (s snapshot) States() []State {
	return s.states
}

func (s snapshot) SnapshotRootAndProofs() ([32]byte, [][][]byte, error) {
	tree := merkle.NewTree(s.getTreeHeight())

	for _, state := range s.states {
		hash, err := state.Hash()
		if err != nil {
			return [32]byte{}, nil, err
		}

		tree.Insert(hash[:])
	}

	snapshotRoot, err := tree.Root(uint32(len(s.states)))
	if err != nil {
		return [32]byte{}, nil, err
	}

	var snapshotRootB32 [32]byte
	copy(snapshotRootB32[:], snapshotRoot[:])

	proofs := make([][][]byte, len(s.states))
	for i := 0; i < len(s.states); i++ {
		proofs[i], err = tree.MerkleProof(uint32(i), uint32(len(s.states)))
		if err != nil {
			return [32]byte{}, nil, err
		}
	}

	return snapshotRootB32, proofs, nil
}

// getTreeHeight returns the height of the merkle tree given `len(states)` leafs.
func (s snapshot) getTreeHeight() uint32 {
	return uint32(math.Log2(float64(len(s.states))))

}

var _ Snapshot = &snapshot{}
