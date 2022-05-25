package prover

import (
	"github.com/prysmaticlabs/prysm/container/trie"
	"github.com/synapsecns/sanguine/core/db"
)

// ProverSync syncs the db with sparse merkle trie
type ProverSync struct {
	db   db.DB
	tree trie.SparseMerkleTrie
}
