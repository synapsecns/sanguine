package client

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/core/metrics"
	"math/big"
)

// HarmonyVM is a strict superset of evm w/ custom harmony methods
type HarmonyVM interface {
	EVM
	// FilterHarmonyLogs returns the logs that satisfy the supplied filter query.
	FilterHarmonyLogs(ctx context.Context, query ethereum.FilterQuery) (logs []types.Log, err error)
	// HarmonyTransactionReceipt returns the receipt of a transaction by transaction hash.
	HarmonyTransactionReceipt(ctx context.Context, txHash common.Hash) (receipt *types.Receipt, err error)
}

type harmonyClient struct {
	*clientImpl
}

// DialHarmonyBackend dials a harmony backend
func DialHarmonyBackend(ctx context.Context, url string, handler metrics.Handler, ops ...Options) (HarmonyVM, error) {
	evm, err := DialBackend(ctx, url, handler, ops...)
	if err != nil {
		return nil, fmt.Errorf("could not dial backend: %w", err)
	}
	client, ok := evm.(*clientImpl)
	if !ok {
		return nil, fmt.Errorf("could not cast to clientImpl")
	}

	return &harmonyClient{client}, nil
}

func (h *harmonyClient) FilterHarmonyLogs(ctx context.Context, query ethereum.FilterQuery) (logs []types.Log, err error) {
	requestCtx, span := h.clientImpl.startSpan(ctx, HarmonyGetLogsMethod)
	defer func() {
		metrics.EndSpanWithErr(span, err)
	}()

	var result []types.Log
	arg, err := toFilterArg(query)
	if err != nil {
		return nil, err
	}
	err = h.clientImpl.CallContext(requestCtx, &result, HarmonyGetLogsMethod.String(), arg)
	return result, err
}

// HarmonyTransactionReceipt calls TransactionReceipt on the underlying client
//
//nolint:wrapcheck
func (h *harmonyClient) HarmonyTransactionReceipt(ctx context.Context, txHash common.Hash) (receipt *types.Receipt, err error) {
	requestCtx, span := h.clientImpl.startSpan(ctx, HarmonyGetReceiptMethod)
	defer func() {
		metrics.EndSpanWithErr(span, err)
	}()

	var r *types.Receipt
	err = h.clientImpl.CallContext(requestCtx, &r, HarmonyGetReceiptMethod.String(), txHash)
	if err == nil {
		if r == nil {
			return nil, ethereum.NotFound
		}
	}
	return r, err
}

func toFilterArg(q ethereum.FilterQuery) (interface{}, error) {
	arg := map[string]interface{}{
		"address": q.Addresses,
		"topics":  q.Topics,
	}
	if q.BlockHash != nil {
		arg["blockHash"] = *q.BlockHash
		if q.FromBlock != nil || q.ToBlock != nil {
			return nil, fmt.Errorf("cannot specify both BlockHash and FromBlock/ToBlock")
		}
	} else {
		if q.FromBlock == nil {
			arg["fromBlock"] = "0x0"
		} else {
			arg["fromBlock"] = toBlockNumArg(q.FromBlock)
		}
		arg["toBlock"] = toBlockNumArg(q.ToBlock)
	}
	return arg, nil
}

func toBlockNumArg(number *big.Int) string {
	if number == nil {
		return "latest"
	}
	pending := big.NewInt(-1)
	if number.Cmp(pending) == 0 {
		return "pending"
	}
	return hexutil.EncodeBig(number)
}
