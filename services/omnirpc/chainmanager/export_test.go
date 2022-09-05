package chainmanager

import "github.com/synapsecns/sanguine/services/omnirpc/rpcinfo"

// SortInfoList exports sortInfoList for testing.
func SortInfoList(rpcInfoList []rpcinfo.Result) []rpcinfo.Result {
	return sortInfoList(rpcInfoList)
}
