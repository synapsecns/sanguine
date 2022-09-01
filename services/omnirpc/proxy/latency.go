package proxy

import (
	"context"
	"github.com/synapsecns/sanguine/services/omnirpc/latency"
	"sort"
)

// reorderRPCs used for the round robin based on latency.
func (r *RPCProxy) reorderRPCs(ctx context.Context, chainID int) {
	rpcList := r.rpcMap.ChainID(chainID)

	if len(rpcList) == 0 {
		return
	}

	latencyList := latency.GetRPCLatency(ctx, rpcTimeout, rpcList)

	// sort loweset->highest latency
	sort.Slice(latencyList, func(i, j int) bool {
		// ignore latencies with an error
		if latencyList[i].HasError {
			return false
		}

		ageDifference := latencyList[i].BlockAge - latencyList[j].BlockAge
		if ageDifference == 0 {
			return latencyList[i].Latency < latencyList[j].Latency
		} else if ageDifference > 0 {
			return false
		}

		return true
	})

	var newOrder []string
	for _, rpcItem := range latencyList {
		newOrder = append(newOrder, rpcItem.URL)
	}

	r.rpcMap.PutChainID(chainID, newOrder)
}
