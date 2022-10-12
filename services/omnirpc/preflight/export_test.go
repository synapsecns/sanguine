package preflight

import "github.com/synapsecns/sanguine/services/omnirpc/types"

// CheckTransaction exports the checkTransaction method for testing
func CheckTransaction(rpcRequest types.IRPCRequest, chainID uint64) (ok bool, _ error) {
	//nolint: wrapcheck
	return checkTransaction(rpcRequest, chainID)
}
