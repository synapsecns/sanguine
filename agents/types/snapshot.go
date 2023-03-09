package types

import (
	"context"
	"fmt"
	"math"

	"github.com/ethereum/go-ethereum/common"
	"github.com/synapsecns/sanguine/core"
	"github.com/synapsecns/sanguine/core/merkle"
	"github.com/synapsecns/sanguine/ethergo/signer/signer"
)

// Snapshot is the snapshot interface.
type Snapshot interface {
	// States are the states of the snapshot.
	States() []State

	// SnapshotRootAndProofs returns the snapshot root, calculated from the states, as well as each state's proof.
	SnapshotRootAndProofs() ([32]byte, [][][]byte, error)
	// TreeHeight returns the height of the merkle tree given `len(states)` leafs.
	TreeHeight() uint32
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
	tree := merkle.NewTree(s.TreeHeight())

	for _, state := range s.states {
		hash, err := state.Hash()
		if err != nil {
			return [32]byte{}, nil, fmt.Errorf("failed to hash state: %w", err)
		}

		tree.Insert(hash[:])
	}

	snapshotRoot, err := tree.Root(uint32(len(s.states)))
	if err != nil {
		return [32]byte{}, nil, fmt.Errorf("failed to get snapshot root: %w", err)
	}

	var snapshotRootB32 [32]byte
	copy(snapshotRootB32[:], snapshotRoot)

	proofs := make([][][]byte, len(s.states))
	for i := 0; i < len(s.states); i++ {
		proofs[i], err = tree.MerkleProof(uint32(i), uint32(len(s.states)))
		if err != nil {
			return [32]byte{}, nil, fmt.Errorf("failed to get merkle proof: %w", err)
		}
	}

	return snapshotRootB32, proofs, nil
}

// TreeHeight returns the height of the merkle tree given `len(states)` leafs.
func (s snapshot) TreeHeight() uint32 {
	return uint32(math.Log2(float64(len(s.states))))
}

func (s snapshot) SignSnapshot(ctx context.Context, signer signer.Signer) (signer.Signature, []byte, common.Hash, error) {
	encodedSnapshot, err := EncodeSnapshot(s)
	if err != nil {
		return nil, nil, common.Hash{}, fmt.Errorf("could not encode snapshot: %w", err)
	}

	hashedSnapshot, err := HashRawBytes(encodedSnapshot)
	if err != nil {
		return nil, nil, common.Hash{}, fmt.Errorf("could not hash snapshot: %w", err)
	}
	signature, err := signer.SignMessage(ctx, core.BytesToSlice(hashedSnapshot), false)
	if err != nil {
		return nil, nil, common.Hash{}, fmt.Errorf("could not sign snapshot: %w", err)
	}
	return signature, encodedSnapshot, hashedSnapshot, nil
}

var _ Snapshot = &snapshot{}
