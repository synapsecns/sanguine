package fastbridge

// BridgeStatus is an enum for the on-chain status of a request
//
//go:generate go run golang.org/x/tools/cmd/stringer -type=BridgeStatus -linecomment
type BridgeStatus uint8

// DO NOT USE IOTA! These are meant to reflect on chain statuses.
const (
	// NULL is the default value for a bridge status.
	NULL BridgeStatus = 0 // NULL
	// REQUESTED is the status for a request that has been made.
	REQUESTED BridgeStatus = 1 // REQUESTED
	// RelayerProved is the status for a request that has been proved by a relayer.
	RelayerProved BridgeStatus = 2 // RELAYER_PROVED
	// RelayerClaimed is the status for a request that has been claimed by a relayer.
	RelayerClaimed BridgeStatus = 3 // RELAYER_CLAIMED
	// REFUNDED is the status for a request that has been refunded.
	REFUNDED BridgeStatus = 4 // REFUNDED
)

func (b BridgeStatus) Int() uint8 {
	return uint8(b)
}

// set all contact types.
func init() {
	for i := 0; i < len(_BridgeStatus_index)-1; i++ {
		contractType := BridgeStatus(i)
		allBridgeStatuses = append(allBridgeStatuses, contractType)
		// assert type is correct
	}
}

// allBridgeStatuses is a list of all bridge statuses Since we use stringer and this is a testing library, instead
// // of manually copying all these out we pull the names out of stringer. In order to make sure stringer is updated, we panic on
// // any method called where the index is higher than the stringer array length.
var allBridgeStatuses []BridgeStatus
