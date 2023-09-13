package types

import (
	"context"
	"fmt"
	"math"

	"github.com/ethereum/go-ethereum/common"
	"github.com/synapsecns/sanguine/core/merkle"
	"github.com/synapsecns/sanguine/ethergo/signer/signer"
)

// Snapshot is the snapshot interface.
type Snapshot interface {
	Encoder
	// States are the states of the snapshot.
	States() []State

	// SnapshotRootAndProofs returns the snapshot root, calculated from the states, as well as each state's proof.
	SnapshotRootAndProofs() ([32]byte, [][][]byte, error)
	// SignSnapshot returns the signature of the snapshot payload signed by the signer
	SignSnapshot(ctx context.Context, signer signer.Signer) (signer.Signature, []byte, common.Hash, error)
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
	tree := merkle.NewTree(merkle.SnapshotTreeHeight)

	for _, state := range s.states {
		leftLeaf, rightLeaf, err := state.SubLeaves()
		if err != nil {
			return [32]byte{}, nil, fmt.Errorf("failed to hash state: %w", err)
		}

		tree.Insert(leftLeaf[:])
		tree.Insert(rightLeaf[:])
	}

	snapshotRoot, err := tree.Root(uint32(len(s.states) * 2))
	if err != nil {
		return [32]byte{}, nil, fmt.Errorf("failed to get snapshot root: %w", err)
	}

	var snapshotRootB32 [32]byte
	copy(snapshotRootB32[:], snapshotRoot)

	proofs := make([][][]byte, len(s.states)*2)
	for i := 0; i < len(s.states)*2; i += 2 {
		proofs[i/2], err = tree.MerkleProof(uint32(i), uint32(len(s.states)*2))
		if err != nil {
			return [32]byte{}, nil, fmt.Errorf("failed to get merkle proof: %w", err)
		}
	}

	return snapshotRootB32, proofs, nil
}

// TreeHeight returns the height of the merkle tree given `len(states)*2` leaves.
func (s snapshot) TreeHeight() uint32 {
	return uint32(math.Log2(float64(len(s.states) * 2)))
}

func (s snapshot) SignSnapshot(ctx context.Context, signer signer.Signer) (signer.Signature, []byte, common.Hash, error) {
	return signEncoder(ctx, signer, s, SnapshotValidSalt)
}

var _ Snapshot = &snapshot{}
