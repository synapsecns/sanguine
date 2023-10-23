// Package fetcher gets data from scribe.
package fetcher

import (
	"context"
	"fmt"
	"time"

	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/services/sinner/types"

	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/jpillora/backoff"
	"github.com/synapsecns/sanguine/services/scribe/graphql"
	"github.com/synapsecns/sanguine/services/sinner/fetcher/client"
	"github.com/synapsecns/sanguine/services/sinner/logger"
)

// ScribeFetcher is the interface for fetching events. It uses GQL.
type ScribeFetcher interface {
	// FetchLastIndexed fetches the last indexed block per contract.
	FetchLastIndexed(ctx context.Context, chainID uint32, contractAddress string) (uint64, error)
	// FetchLogsInRange fetches logs in a range with the GQL client.
	FetchLogsInRange(ctx context.Context, chainID uint32, startBlock, endBlock uint64, contractAddress common.Address) ([]ethTypes.Log, error)
	// FetchBlockTime fetches the timestamp of a block.
	FetchBlockTime(ctx context.Context, chainID int, blockNumber int) (*int, error)
	// FetchTx fetches the transaction.
	FetchTx(ctx context.Context, tx string, chainID int, blockNumber int) (*uint64, *string, error)
	// FetchTxsInRange fetches transactions in a range.
	FetchTxsInRange(ctx context.Context, chainID uint32, startBlock uint64, endBlock uint64) ([]types.TxSupplementalInfo, error)
}

type scribeFetcherImpl struct {
	underlyingClient *client.Client
	handler          metrics.Handler
}

const retryThreshold = 5

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

// FetchLogsInRange fetches logs in a range with the GQL client.
func (s scribeFetcherImpl) FetchLogsInRange(ctx context.Context, chainID uint32, startBlock, endBlock uint64, contractAddress common.Address) ([]ethTypes.Log, error) {
	logs := &client.GetLogsRange{}
	page := 1
	contractAddressString := contractAddress.String()

	for {
		paginatedLogs, err := s.underlyingClient.GetLogsRange(ctx, int(chainID), int(startBlock), int(endBlock), page, &contractAddressString)
		if err != nil {
			return nil, fmt.Errorf("could not get logs: %w", err)
		}
		if len(paginatedLogs.Response) == 0 {
			break
		}

		logs.Response = append(logs.Response, paginatedLogs.Response...)
		page++
	}

	var parsedLogs []ethTypes.Log
	for _, log := range logs.Response {
		parsedLog, err := graphql.ParseLog(*log)
		if err != nil {
			return nil, fmt.Errorf("could not parse log: %w", err)
		}

		parsedLogs = append(parsedLogs, *parsedLog)
	}

	return parsedLogs, nil
}

func (s scribeFetcherImpl) FetchBlockTime(ctx context.Context, chainID int, blockNumber int) (*int, error) {
	b := &backoff.Backoff{
		Factor: 2,
		Jitter: true,
		Min:    10 * time.Millisecond,
		Max:    3 * time.Second,
	}
	timeout := time.Duration(0)
	attempts := 0
RETRY:
	attempts++
	if attempts > retryThreshold {
		logger.ReportSinnerError(fmt.Errorf("could not get block time for block %d on chainID %d after %d attempts", blockNumber, chainID, retryThreshold), uint32(chainID), logger.ScribeFetchFailure)
		return nil, fmt.Errorf("could not get block time for block %d on chainID %d after %d attempts", blockNumber, chainID, retryThreshold)
	}
	select {
	case <-ctx.Done():

		return nil, fmt.Errorf("could not get timestamp for block, context canceled %d: %d", chainID, blockNumber)
	case <-time.After(timeout):

		timeStamp, err := s.underlyingClient.GetBlockTime(ctx, chainID, blockNumber)

		if err != nil {
			logger.ReportSinnerError(fmt.Errorf("could not get timestamp for block, trying again %d: %w", blockNumber, err), uint32(chainID), logger.ScribeFetchFailure)
			timeout = b.Duration()
			goto RETRY
		}

		if timeStamp == nil || timeStamp.Response == nil {
			logger.ReportSinnerError(fmt.Errorf("could not get timestamp for block, invalid blocktime %d: %d", chainID, blockNumber), uint32(chainID), logger.ScribeFetchFailure)

			return nil, fmt.Errorf("could not get timestamp for block, invalid blocktime %d: %d", chainID, blockNumber)
		}

		return timeStamp.Response, nil
	}
}

// FetchTx fetches the transaction of a log.
func (s scribeFetcherImpl) FetchTx(ctx context.Context, tx string, chainID int, blockNumber int) (*uint64, *string, error) {
	b := &backoff.Backoff{
		Factor: 2,
		Jitter: true,
		Min:    10 * time.Millisecond,
		Max:    3 * time.Second,
	}
	attempts := 0
	timeout := time.Duration(0)
RETRY:
	attempts++
	if attempts > retryThreshold {
		logger.ReportSinnerError(fmt.Errorf("could not get tx after %d attempts for hash %s", retryThreshold, tx), uint32(chainID), logger.ScribeFetchFailure)
		auxiliaryBlocktime, err := s.FetchBlockTime(ctx, chainID, blockNumber)
		if err != nil {
			return nil, nil, fmt.Errorf("could not get tx for log, after trying to get blocktime, invalid response %d: %s", chainID, tx)
		}
		sender := ""
		blocktime := uint64(*auxiliaryBlocktime)
		return &blocktime, &sender, nil
	}

	select {
	case <-ctx.Done():
		return nil, nil, fmt.Errorf("could not get tx for log, context canceled %d: %s", chainID, tx)
	case <-time.After(timeout):

		res, err := s.underlyingClient.GetTransactions(ctx, chainID, 1, &tx)

		if err != nil || res == nil || res.Response == nil || len(res.Response) == 0 {
			timeout = b.Duration()
			goto RETRY
		}

		resTx := res.Response[0]
		sender := resTx.Sender
		blocktime := uint64(resTx.Timestamp)
		return &blocktime, &sender, nil
	}
}

// FetchTxsInRange fetches tx in a range with the GQL client.
func (s scribeFetcherImpl) FetchTxsInRange(ctx context.Context, chainID uint32, startBlock uint64, endBlock uint64) ([]types.TxSupplementalInfo, error) {
	txs := &client.GetTransactionsRange{}
	page := 1

	for {
		paginatedTxs, err := s.underlyingClient.GetTransactionsRange(ctx, int(chainID), int(startBlock), int(endBlock), page)
		if err != nil {
			return nil, fmt.Errorf("could not get txs: %w", err)
		}
		if len(paginatedTxs.Response) == 0 {
			break
		}

		txs.Response = append(txs.Response, paginatedTxs.Response...)
		page++
	}

	return parseTx(txs), nil
}

// parseTx converts a tx from GraphQL into a model log.
func parseTx(txs *client.GetTransactionsRange) []types.TxSupplementalInfo {
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
