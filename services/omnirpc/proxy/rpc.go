package proxy

import (
	"fmt"
	"github.com/ethereum/go-ethereum/eth/filters"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/goccy/go-json"
	"github.com/hedzr/cmdr/tool"
	"github.com/synapsecns/sanguine/services/omnirpc/types"
	"golang.org/x/exp/slices"
	"math/big"
)

func isBlockNumConfirmable(arg json.RawMessage) bool {
	// nonConfirmableBlockNumArgs is a list of non numerical block args
	var nonConfirmableBlockNumArgs = []string{"latest", "pending"}

	return !slices.Contains(nonConfirmableBlockNumArgs, tool.StripQuotes(string(arg)))
}

// isFilterArgConfirmable checks if filter.filterCriteria is confirmable.
func isFilterArgConfirmable(arg json.RawMessage) (bool, error) {
	// cast latest block number to a big int for comparison
	latestBlockNumber := new(big.Int).SetInt64(rpc.LatestBlockNumber.Int64())

	filterCriteria := filters.FilterCriteria{}
	err := filterCriteria.UnmarshalJSON(arg)
	if err != nil {
		return false, fmt.Errorf("could not unmarshall filter: %w", err)
	}

	// Block filter requested, construct a single-shot filter
	if filterCriteria.BlockHash != nil {
		return true, nil
	}

	usesLatest := filterCriteria.FromBlock.Cmp(latestBlockNumber) == 0 || filterCriteria.ToBlock.Cmp(latestBlockNumber) == 0
	return !usesLatest, nil
}

func isConfirmable(r types.IRPCRequest) (bool, error) {
	// TODO: handle batch methods
	// TODO: should we error on default?
	//nolint: exhaustive
	switch r.GetMethod() {
	case types.BlockByNumberMethod, types.PendingTransactionCountMethod:
		return isBlockNumConfirmable(r.GetParams()[0]), nil
	case types.BlockNumberMethod, types.SyncProgressMethod, types.GasPriceMethod, types.MaxPriorityMethod, types.EstimateGasMethod:
		return false, nil
	case types.GetBalanceMethod, types.GetCodeMethod, types.TransactionCountMethod, types.CallMethod:
		return isBlockNumConfirmable(r.GetParams()[1]), nil
	case types.StorageAtMethod:
		return isBlockNumConfirmable(r.GetParams()[2]), nil
	case types.GetLogsMethod:
		return isFilterArgConfirmable(r.GetParams()[0])
	// not confirmable because tx could be pending. We might want to handle w/ omnicast though
	// left separate for comment
	case types.SendRawTransactionMethod:
		return false, nil
	}
	return true, nil
}
