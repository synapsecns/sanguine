package types

// MessageState represents the state transitions of a CCTP transfer.
type MessageState int

const (
	// Pending indicates the USDC transfer has been initiated on the origin chain and is pending attestation.
	Pending MessageState = iota + 1
	// Attested indicates the USDC transfer is waiting for submission on the destination chain.
	Attested
	// Submitted indicates the USDC transfer has been confirmed on the destination chain.
	Submitted
	// Complete indicates the USDC transfer has been completed on the destination chain.
	Complete
)
