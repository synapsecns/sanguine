// Package fetcher gets data from scribe.
package fetcher

import (
	"context"
	"fmt"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/services/sinner/types"

	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/services/scribe/graphql"
	"github.com/synapsecns/sanguine/services/sinner/fetcher/client"
)

// ScribeFetcher is the interface for fetching events. It uses GQL.
type ScribeFetcher interface {
	// FetchLastIndexed fetches the last indexed block per contract.
	FetchLastIndexed(ctx context.Context, chainID uint32, contractAddress string) (uint64, error)
	// FetchLogsAndTransactionsRange fetches logs and transactions in a range with the GQL client.
	FetchLogsAndTransactionsRange(ctx context.Context, chainID uint32, startBlock, endBlock uint64, contractAddress common.Address, page int) ([]ethTypes.Log, []types.TxSupplementalInfo, error)
}

type scribeFetcherImpl struct {
	underlyingClient *client.Client
	handler          metrics.Handler
}

// NewFetcher creates a new fetcher.
func NewFetcher(fetchClient *client.Client, handler metrics.Handler) ScribeFetcher {
	return &scribeFetcherImpl{
		underlyingClient: fetchClient,
		handler:          handler,
	}
}

// FetchLastIndexed fetches the last indexed block per contract.
func (s scribeFetcherImpl) FetchLastIndexed(ctx context.Context, chainID uint32, contractAddress string) (uint64, error) {
	lastIndexed, err := s.underlyingClient.GetLastIndexed(ctx, int(chainID), contractAddress)
	if err != nil || lastIndexed == nil || lastIndexed.Response == nil {
		return 0, fmt.Errorf("could not get last indexed for contract %s: %w", contractAddress, err)
	}
	return uint64(*lastIndexed.Response), nil
}

func (s scribeFetcherImpl) FetchLogsAndTransactionsRange(ctx context.Context, chainID uint32, startBlock, endBlock uint64, contractAddress common.Address, page int) ([]ethTypes.Log, []types.TxSupplementalInfo, error) {
	// Get data in ascending order.
	res, err := s.underlyingClient.GetLogsAndTransactionsRange(ctx, int(chainID), int(startBlock), int(endBlock), page, contractAddress.String(), true)
	if err != nil {
		return nil, nil, fmt.Errorf("could not get logs and txs: %w", err)
	}
	if len(res.LogsRange) == 0 {
		return []ethTypes.Log{}, []types.TxSupplementalInfo{}, nil
	}
	var parsedLogs []ethTypes.Log
	for _, log := range res.LogsRange {
		parsedLog, err := graphql.ParseLog(*log)
		if err != nil {
			return nil, nil, fmt.Errorf("could not parse log: %w", err)
		}

		parsedLogs = append(parsedLogs, *parsedLog)
	}
	return parsedLogs, parseTx(res), nil
}

// parseTx converts a tx from GraphQL into a model log.
func parseTx(txs *client.GetLogsAndTransactionsRange) []types.TxSupplementalInfo {
	var txSupplementalInfo []types.TxSupplementalInfo

	for _, tx := range txs.Response {
		newTx := types.TxSupplementalInfo{
			TxHash:    tx.TxHash,
			Sender:    tx.Sender,
			Timestamp: tx.Timestamp,
		}
		txSupplementalInfo = append(txSupplementalInfo, newTx)
	}

	return txSupplementalInfo
}
