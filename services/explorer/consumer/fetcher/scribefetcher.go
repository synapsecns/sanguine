package fetcher

import (
	"context"
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/jpillora/backoff"
	"github.com/synapsecns/sanguine/services/explorer/consumer/client"
	"github.com/synapsecns/sanguine/services/scribe/graphql"
)

// ScribeFetcher is the fetcher for the events. It uses GQL.
type ScribeFetcher struct {
	FetchClient *client.Client
}

const retryThreshold = 5

// NewFetcher creates a new fetcher.
func NewFetcher(fetchClient *client.Client) *ScribeFetcher {
	return &ScribeFetcher{
		FetchClient: fetchClient,
	}
}

// FetchTxSender fetches the sender of a transaction.
func (s ScribeFetcher) FetchTxSender(ctx context.Context, chainID uint32, txHash string) (string, error) {
	b := &backoff.Backoff{
		Factor: 2,
		Jitter: true,
		Min:    10 * time.Millisecond,
		Max:    3 * time.Second,
	}
	timeout := time.Duration(0)
RETRY:
	select {
	case <-ctx.Done():

		return "", nil
	case <-time.After(timeout):
		sender, err := s.FetchClient.GetTxSender(ctx, int(chainID), txHash)

		if err != nil {
			logger.Warnf("could not get sender for tx, trying again %s: %v", txHash, err)
			timeout = b.Duration()
			goto RETRY
		}

		if sender == nil || sender.Response == nil {
			logger.Warnf("could not get sender for tx, invalid tx likely (arb legacy, v,r,x, etc.) %s: %v", txHash)
			*sender.Response = ""
		}

		return *sender.Response, nil
	}
}

// FetchLastIndexed fetches the last indexed block per contract.
func (s ScribeFetcher) FetchLastIndexed(ctx context.Context, chainID uint32, contractAddress string) (uint64, error) {
	lastIndexed, err := s.FetchClient.GetLastIndexed(ctx, int(chainID), contractAddress)
	if err != nil || lastIndexed == nil || lastIndexed.Response == nil {
		return 0, fmt.Errorf("could not get last indexed for contract %s: %w", contractAddress, err)
	}
	return uint64(*lastIndexed.Response), nil
}

// FetchLogsInRange fetches logs in a range with the GQL client.
func (s ScribeFetcher) FetchLogsInRange(ctx context.Context, chainID uint32, startBlock, endBlock uint64, contractAddress common.Address) ([]ethTypes.Log, error) {
	logs := &client.GetLogsRange{}
	page := 1
	contractAddressString := contractAddress.String()

	for {
		paginatedLogs, err := s.FetchClient.GetLogsRange(ctx, int(chainID), int(startBlock), int(endBlock), page, &contractAddressString)
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

// FetchBlockTime fetches the timestamp of a block.
func (s ScribeFetcher) FetchBlockTime(ctx context.Context, chainID int, blockNumber int) (*int, error) {
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
		logger.Errorf("could not get block time for block %d on chainID %d after %d attempts", blockNumber, chainID, retryThreshold)
		return nil, fmt.Errorf("could not get block time for block %d on chainID %d after %d attempts", blockNumber, chainID, retryThreshold)
	}
	select {
	case <-ctx.Done():

		return nil, fmt.Errorf("could not get timestamp for block, context canceled %d: %d", chainID, blockNumber)
	case <-time.After(timeout):

		timeStamp, err := s.FetchClient.GetBlockTime(ctx, chainID, blockNumber)

		if err != nil {
			logger.Warnf("could not get timestamp for block, trying again %d: %v", blockNumber, err)
			timeout = b.Duration()
			goto RETRY
		}

		if timeStamp == nil || timeStamp.Response == nil {
			logger.Warnf("could not get timestamp for block, invalid blocktime %d: %d", chainID, blockNumber)
			return nil, fmt.Errorf("could not get timestamp for block, invalid blocktime %d: %d", chainID, blockNumber)
		}

		return timeStamp.Response, nil
	}
}

// FetchTx fetches the transaction of a log.
func (s ScribeFetcher) FetchTx(ctx context.Context, tx string, chainID int, blockNumber int) (*uint64, *string, error) {
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
		logger.Errorf("could not get tx after %d attempts for hash %s on chain %d trying blocktime", retryThreshold, tx, chainID)
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

		res, err := s.FetchClient.GetTransactions(ctx, chainID, 1, &tx)

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
