package db

// DBAttestation is the executor type for interacting with the database representation of an attestation.
//
//nolint:golint,revive
type DBAttestation struct {
	// Destination is the destination of the attestation.
	Destination *uint32
	// SnapshotRoot is the snapshot root.
	SnapshotRoot *string
	// DataHash is the agent root and SnapGasHash combined into a single hash.
	DataHash *string
	// AttestationNonce is the nonce of the attestation.
	AttestationNonce *uint32
	// SummitBlockNumber is the block number when the attestation was created in Summit.
	SummitBlockNumber *uint64
	// SummitTimestamp is the timestamp of the block when the attestation was created in Summit.
	SummitTimestamp *uint64
	// DestinationBlockNumber is the block number that the attestation was submitted on the destination.
	DestinationBlockNumber *uint64
	// DestinationTimestamp is the timestamp of the block that the attestation was submitted on the destination.
	DestinationTimestamp *uint64
}
