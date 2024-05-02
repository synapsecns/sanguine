package relconfig

import "fmt"

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

// RebalanceMethodFromString converts a string to a RebalanceMethod.
func RebalanceMethodFromString(str string) (RebalanceMethod, error) {
	switch str {
	case "synapsecctp":
		return RebalanceMethodSynapseCCTP, nil
	case "circlecctp":
		return RebalanceMethodCircleCCTP, nil
	case "native":
		return RebalanceMethodNative, nil
	case "":
		return RebalanceMethodNone, nil
	default:
		return RebalanceMethodNone, fmt.Errorf("invalid rebalance method: %s", str)
	}
}

// CoalesceRebalanceMethods coalesces two rebalance methods.
func CoalesceRebalanceMethods(a, b RebalanceMethod) RebalanceMethod {
	if a == b {
		return a
	}
	return RebalanceMethodNone
}
