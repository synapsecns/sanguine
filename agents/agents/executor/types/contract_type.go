package types

// ContractType is the type of contract.
//
//go:generate go run golang.org/x/tools/cmd/stringer -type=ContractType -linecomment
type ContractType int

const (
	// OriginContract is the Origin contract.
	OriginContract ContractType = iota + 1 // OriginContract
	// DestinationContract is the Destination contract.
	DestinationContract // DestinationContract
	// LightInboxContract is the LightInbox contract.
	LightInboxContract // LightInboxContract
	// InboxContract is the Inbox contract.
	InboxContract // InboxContract
	// SummitContract is the Summit contract.
	SummitContract // SummitContract
	// Other is the Other contract.
	Other // Other
)
