package relconfig

// RebalanceMethod is the method to rebalance.
//
//go:generate go run golang.org/x/tools/cmd/stringer -type=RebalanceMethod
type RebalanceMethod uint8

const (
	// CCTPRebalance is the rebalance method for CCTP.
	CCTPRebalance RebalanceMethod = iota + 1
	// NativeBridgeRebalance is the rebalance method for native bridge.
	NativeBridgeRebalance
)

var stringToRebalanceMethod = map[string]RebalanceMethod{
	"cctp":   CCTPRebalance,
	"native": NativeBridgeRebalance,
}
