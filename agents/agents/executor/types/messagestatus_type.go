package types

// MessageStatusType is the type for message statuses on the Destination.
//
//go:generate go run golang.org/x/tools/cmd/stringer -type=MessageStatusType -linecomment
type MessageStatusType int

const (
	// None is if the message has not been processed.
	None MessageStatusType = iota // None
	// Failed is if the message failed to be processed.
	Failed // Failed
	// Success is if the message was processed successfully.
	Success // Success
)
