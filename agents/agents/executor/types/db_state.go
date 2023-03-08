package types

// DBState is the executor type for interacting with the database representation of a state.
type DBState struct {
	// SnapshotRoot is the snapshot root.
	SnapshotRoot *string
	// Root is the origin Merkle tree's root.
	Root *string
	// ChainID is the origin chain id.
	ChainID *uint32
	// Nonce is the origin Merkle tree's nonce.
	Nonce *uint32
	// OriginBlockNumber is the block number that the state was taken from on the origin.
	OriginBlockNumber *uint64
	// OriginTimestamp is the timestamp of the block that the state was taken from on the origin.
	OriginTimestamp *uint64
	// Proof is the Snapshot Merkle Tree proof for the state.
	Proof *[][]byte
	// TreeHeight is the height of the Snapshot Merkle Tree that the state belongs to.
	TreeHeight *uint32
}
