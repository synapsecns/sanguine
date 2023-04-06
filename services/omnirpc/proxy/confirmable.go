package proxy

import (
	"encoding/json"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/ethereum/go-ethereum/eth/filters"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/hashicorp/go-multierror"
	"github.com/hedzr/cmdr/tool"
	"github.com/synapsecns/sanguine/ethergo/client"
	rpc2 "github.com/synapsecns/sanguine/ethergo/parser/rpc"
	"golang.org/x/exp/slices"
	"math/big"
)

func isConfirmable(r rpc2.Request) (bool, error) {
	// TODO: should we error on default?
	// TODO: look at RPCMethod.Comparable for lower, necessary?
	//nolint: exhaustive
	switch client.RPCMethod(r.Method) {
	case client.BlockByNumberMethod, client.PendingTransactionCountMethod:
		return isBlockNumConfirmable(r.Params[0]), nil
	case client.BlockNumberMethod, client.SyncProgressMethod, client.GasPriceMethod, client.MaxPriorityMethod, client.EstimateGasMethod:
		return false, nil
	case client.GetBalanceMethod, client.GetCodeMethod, client.TransactionCountMethod, client.CallMethod:
		return isBlockNumConfirmable(r.Params[1]), nil
	case client.StorageAtMethod:
		return isBlockNumConfirmable(r.Params[2]), nil
	case client.GetLogsMethod:
		return isFilterArgConfirmable(r.Params[0])
	// not confirmable because tx could be pending. We might want to handle w/ omnicast though
	// left separate for comment
	case client.SendRawTransactionMethod:
		return false, nil
	}
	return true, nil
}

func areConfirmable(r rpc2.Requests) (_ bool, errs error) {
	unconfirmable := false

	for i, request := range r {
		canConfirm, err := isConfirmable(request)
		if err != nil {
			errs = multierror.Append(errs, fmt.Errorf("request at index %d: %s is not parsable", i, spew.Sprint(request)))
		}

		if !canConfirm {
			unconfirmable = true
		}
	}

	if errs != nil {
		//nolint:wrapcheck
		return false, errs
	}

	return !unconfirmable, nil
}

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
