package fetcher

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/services/explorer/consumer/client"
	"github.com/synapsecns/sanguine/services/scribe/graphql"
)

// ScribeFetcher is the fetcher for the events. It uses GQL.
type ScribeFetcher struct {
	FetchClient *client.Client
}

// NewFetcher creates a new fetcher.
func NewFetcher(fetchClient *client.Client) *ScribeFetcher {
	return &ScribeFetcher{
		FetchClient: fetchClient,
	}
}

// FetchLastBlock fetches the last block that Scribe has stored.
func (s ScribeFetcher) FetchLastBlock(ctx context.Context, chainID uint32) (uint64, error) {
	getEndHeight, err := s.FetchClient.GetLastStoredBlockNumber(ctx, int(chainID))

	if err != nil && getEndHeight == nil {
		return 0, fmt.Errorf("could not get end height: %w", err)
	}

	return uint64(*getEndHeight.Response), nil
}

// FetchTxSender fetches the sender of a transaction.
func (s ScribeFetcher) FetchTxSender(ctx context.Context, chainID uint32, txHash string) (string, error) {
	sender, err := s.FetchClient.GetTxSender(ctx, int(chainID), txHash)

	if err != nil || sender == nil || sender.Response == nil {
		return "", fmt.Errorf("could not get sender: %w", err)
	}

	return *sender.Response, nil
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
