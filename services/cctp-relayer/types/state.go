package types

type MessageState int

const (
	// Pending indicates the USDC transfer has been initiated on the origin chain and is pending attestation.
	Pending MessageState = iota + 1
	// Attested indicates the USDC transfer is waiting for submission on the destination chain.
	Attested
	// Complete indicates the USDC transfer has been confirmed on the destination chain.
	Complete
)
