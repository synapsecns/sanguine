package types

// DisputeProcessedStatus is the status of a dispute on Summit.
// This enum is used for tracking the status of disputes in the Dispute db model.
type DisputeProcessedStatus uint8

const (
	// Opened is when a dispute has been opened but has not been resolved.
	Opened DisputeProcessedStatus = iota
	// Resolved is when a dispute has been resolved on Summit, but agent status has not been updated on the remote chain.
	Resolved
	// Propagated is when a dispute has been resolved on Summit, and agent status has been updated on the remote chain.
	Propagated
)
