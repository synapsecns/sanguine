package types

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
)

// Proof is a merkle proof object. The leaf it's path to the root and its index in the tree.
type Proof interface {
	// Leaf is the leaf in the proof
	Leaf() common.Hash
	// Index is the index in the tree
	Index() uint32
	// Path is the merkle branch
	Path() common.Hash
	// Encode encodes a message
	Encode() ([]byte, error)
}

// proofImpl implements the merkle proof.
type proofImpl struct {
	leaf  common.Hash
	index uint32
	path  common.Hash
}

// ProofEncoder is exported to allow proofs to be encoded/deoded from binary.
type ProofEncoder struct {
	Leaf  common.Hash
	Index uint32
	Path  common.Hash
}

// NewProof creates a new merkle proof.
func NewProof(leaf common.Hash, index uint32, path common.Hash) Proof {
	return proofImpl{
		leaf:  leaf,
		index: index,
		path:  path,
	}
}

// DecodeProof decodes a proof.
func DecodeProof(rawProof []byte) (Proof, error) {
	dec := gob.NewDecoder(bytes.NewReader(rawProof))

	var proofEncoder ProofEncoder
	err := dec.Decode(&proofEncoder)
	if err != nil {
		return nil, fmt.Errorf("could not decode proof: %w", err)
	}

	return proofImpl{
		leaf:  proofEncoder.Leaf,
		index: proofEncoder.Index,
		path:  proofEncoder.Path,
	}, nil
}

func (p proofImpl) Leaf() common.Hash {
	return p.leaf
}

func (p proofImpl) Index() uint32 {
	return p.index
}

func (p proofImpl) Path() common.Hash {
	return p.path
}

func (p proofImpl) Encode() ([]byte, error) {
	var res bytes.Buffer
	enc := gob.NewEncoder(&res)

	proofEncoder := ProofEncoder{
		Leaf:  p.leaf,
		Index: p.index,
		Path:  p.path,
	}

	err := enc.Encode(proofEncoder)
	if err != nil {
		return nil, fmt.Errorf("could not encode %T: %w", p, err)
	}
	return res.Bytes(), nil
}
