package relconfig

import "fmt"

// RebalanceMethod is the method to rebalance.
type RebalanceMethod uint8

const (
	// RebalanceMethodNone is the default rebalance method.
	RebalanceMethodNone RebalanceMethod = iota
	// RebalanceMethodSynapseCCTP is the rebalance method for CCTP.
	RebalanceMethodSynapseCCTP
	// RebalanceMethodCircleCCTP is the rebalance method for Circle CCTP.
	RebalanceMethodCircleCCTP
	// RebalanceMethodScroll is the rebalance method for Scroll.
	RebalanceMethodScroll
)

// RebalanceMethodFromString converts a string to a RebalanceMethod.
func RebalanceMethodFromString(str string) (RebalanceMethod, error) {
	switch str {
	case "synapsecctp":
		return RebalanceMethodSynapseCCTP, nil
	case "circlecctp":
		return RebalanceMethodCircleCCTP, nil
	case "scroll":
		return RebalanceMethodScroll, nil
	case "":
		return RebalanceMethodNone, nil
	default:
		return RebalanceMethodNone, fmt.Errorf("invalid rebalance method: %s", str)
	}
}

func (i RebalanceMethod) String() string {
	switch i {
	case RebalanceMethodNone:
		return ""
	case RebalanceMethodSynapseCCTP:
		return "synapsecctp"
	case RebalanceMethodCircleCCTP:
		return "circlecctp"
	case RebalanceMethodScroll:
		return "scroll"
	default:
		return ""
	}
}
