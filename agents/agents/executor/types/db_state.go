package types

import "github.com/ethereum/go-ethereum/common"

// DBState is the executor type for interacting with the database representation of a state.
type DBState struct {
	// SnapshotRoot is the snapshot root of the state.
	SnapshotRoot *common.Hash
	// Root is the root of the state.
	Root *common.Hash
	// ChainID is the chain ID of the chain that the state is for.
	ChainID *uint32
	// Nonce is the nonce in the state.
	Nonce *uint32
	// OriginBlockNumber is the block number of the state.
	OriginBlockNumber *uint64
	// OriginTimestamp is the block time of the state.
	OriginTimestamp *uint64
}
