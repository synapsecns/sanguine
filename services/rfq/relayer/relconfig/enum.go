package relconfig

// RebalanceMethod is the method to rebalance.
//
//go:generate go run golang.org/x/tools/cmd/stringer -type=RebalanceMethod
type RebalanceMethod uint8

const (
	// RebalanceMethodNone is the default rebalance method.
	RebalanceMethodNone RebalanceMethod = iota
	// RebalanceMethodSynapseCCTP is the rebalance method for CCTP.
	RebalanceMethodSynapseCCTP
	// RebalanceMethodCircleCCTP is the rebalance method for Circle CCTP.
	RebalanceMethodCircleCCTP
	// RebalanceMethodNative is the rebalance method for native bridge.
	RebalanceMethodNative
)
