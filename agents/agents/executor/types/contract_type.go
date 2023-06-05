package types

// ContractType is the type of contract.
//
//go:generate go run golang.org/x/tools/cmd/stringer -type=ContractType -linecomment
type ContractType int

const (
	// OriginContract is a contract that dispatches messages on any chain.
	OriginContract ContractType = iota + 1 // OriginContract
	// DestinationContract is a contract that receives messages on a specific chain.
	DestinationContract // DestinationContract
	// SummitContract is a contract that stores the combined data from all the chains.
	SummitContract // SummitContract
	// Other is any other contract.
	Other // Other
)
