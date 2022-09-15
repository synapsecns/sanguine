package consumer

import (
	"context"
	"fmt"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/services/explorer/db/consumer/client"
	"github.com/synapsecns/sanguine/services/scribe/graphql"
)

type Fetcher struct {
	fetchClient *client.Client
}

func NewFetcher(fetchClient *client.Client) *Fetcher {
	return &Fetcher{
		fetchClient: fetchClient,
	}
}

func (f Fetcher) FetchLogsInRange(ctx context.Context, chainID uint32, startBlock, endBlock uint64) ([]ethTypes.Log, error) {
	logs, err := f.fetchClient.GetLogsRange(ctx, int(chainID), int(startBlock), int(endBlock), 1)
	if err != nil {
		return nil, fmt.Errorf("could not fetch logs: %w", err)
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
