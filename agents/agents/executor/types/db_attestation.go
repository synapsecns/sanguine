package types

import "github.com/ethereum/go-ethereum/common"

// DBAttestation is the executor type for interacting with the database representation of an attestation.
type DBAttestation struct {
	// ChainID is the chain ID of the chain that the attestation is for.
	ChainID *uint32
	// Destination is the destination chain id of the attestation.
	Destination *uint32
	// Nonce is the nonce of the attestation.
	Nonce *uint32
	// Root is the root of the attestation.
	Root *common.Hash
	// DestinationBlockNumber is the block number of the attestation as it was submitted on the destination.
	DestinationBlockNumber *uint64
	// DestinationBlockTime is the block time of the attestation as it was submitted on the destination.
	DestinationBlockTime *uint64
}
