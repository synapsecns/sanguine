package types

// DBMessage is the executor type for interacting with the database representation of a message.
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
	// AttestationNonce is the nonce of the attestation that was used to set the minimum time.
	// If this is 0, then the minimum time was not set by an attestation.
	AttestationNonce *uint32
	// MinimumTime is the minimum time that the message can be executed.
	MinimumTime *uint64
}
