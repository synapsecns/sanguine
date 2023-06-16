package db

// DBMessage is the executor type for interacting with the database representation of a message.
//
//nolint:golint,revive
type DBMessage struct {
	// ChainID is the chain ID of the chain that the message is for.
	ChainID *uint32
	// Destination is the destination chain id of the message.
	Destination *uint32
	// Nonce is the nonce of the message.
	Nonce *uint32
	// Message is the message.
	Message *[]byte
	// BlockNumber is the block number of the message.
	BlockNumber *uint64
	// Executed is if the message has been executed.
	Executed *bool
	// MinimumTimeSet is if the MinimumTime field has been set from an Attestation.
	MinimumTimeSet *bool
	// MinimumTime is the minimum time that the message can be executed.
	MinimumTime *uint64
}
