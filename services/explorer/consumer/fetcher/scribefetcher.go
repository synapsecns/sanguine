package fetcher

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/services/explorer/consumer/client"
	"github.com/synapsecns/sanguine/services/scribe/graphql"
	"golang.org/x/sync/errgroup"
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

type block struct {
	number uint64
	time   uint64
}

type blockRange struct {
	startBlock block
	endBlock   block
}

func (b *blockRange) Length() uint64 {
	return b.endBlock.number - b.startBlock.number + 1
}

func (b *blockRange) StartTime() uint64 {
	return b.startBlock.time
}

func (b *blockRange) EndTime() uint64 {
	return b.endBlock.time
}

// TODO Nuke most of this stuff.

// TimeToBlockNumber returns the first block number after the given time.
//
//nolint:gocognit,cyclop
func (s ScribeFetcher) TimeToBlockNumber(ctx context.Context, chainID int, startHeight, targetTime uint64) (uint64, error) {
	searchRange, err := s.getSearchRange(ctx, startHeight, chainID)
	if err != nil {
		return 0, fmt.Errorf("could not get search range: %w", err)
	}

	if targetTime <= searchRange.StartTime() {
		return searchRange.startBlock.number, nil
	}

	if targetTime >= searchRange.EndTime() {
		return searchRange.endBlock.number, nil
	}

	// Run the binary search.
	var i, mid, j uint64

	i = searchRange.startBlock.number
	j = searchRange.endBlock.number

	for i < j {
		mid = (i + j) / 2
		midBlock, err := s.FetchClient.GetBlockTime(ctx, chainID, int(mid))
		if err != nil || midBlock == nil || midBlock.Response == nil {
			return 0, fmt.Errorf("could not get mid time: %w", err)
		}

		midTime := uint64(*midBlock.Response)

		if midTime == targetTime {
			return mid, nil
		}

		// If target is less than midBlock element, then search in left.
		//nolint: nestif // we want to keep the logic for the binary search together
		if targetTime < midTime {
			// If target is greater than previous to mid, return the closest of the two.
			midSubBlock, err := s.FetchClient.GetBlockTime(ctx, chainID, int(mid-1))
			if err != nil || midSubBlock == nil || midSubBlock.Response == nil {
				return 0, fmt.Errorf("could not get mid time: %w", err)
			}

			midSubTime := uint64(*midSubBlock.Response)

			if mid > 0 && targetTime >= midSubTime {
				closestBlock := getClosest(
					block{mid - 1, midSubTime},
					block{mid, midTime},
					targetTime,
				)

				return closestBlock.number, nil
			}

			j = mid
		} else {
			// If target is greater than mid.
			midPlusBlock, err := s.FetchClient.GetBlockTime(ctx, chainID, int(mid+1))
			if err != nil || midPlusBlock == nil || midPlusBlock.Response == nil {
				return 0, fmt.Errorf("could not get mid time: %w", err)
			}

			midPlusTime := uint64(*midPlusBlock.Response)

			if mid <= searchRange.Length()-1 && targetTime <= midPlusTime {
				closestBlock := getClosest(
					block{mid, midTime},
					block{mid + 1, midPlusTime},
					targetTime,
				)

				return closestBlock.number, nil
			}

			i = mid + 1
		}
	}

	// Only a single element is left after the search. Fetch the block and return it.
	resultingBlock, err := s.FetchClient.GetBlockTime(ctx, chainID, int(mid))
	if err != nil || resultingBlock == nil || resultingBlock.Response == nil {
		return 0, fmt.Errorf("could not get mid time: %w", err)
	}

	return mid, nil
}

func getClosest(lesser block, greater block, target uint64) block {
	if (target - lesser.time) >= (greater.time - target) {
		return greater
	}

	return lesser
}

//nolint:cyclop
func (s ScribeFetcher) getSearchRange(ctx context.Context, startHeight uint64, chainID int) (*blockRange, error) {
	getEndHeight, err := s.FetchClient.GetLastStoredBlockNumber(ctx, chainID)
	if err != nil {
		return nil, fmt.Errorf("could not get end height: %w", err)
	}

	endHeight := uint64(*getEndHeight.Response)

	var output blockRange

	if startHeight == 0 {
		getStartHeight, err := s.FetchClient.GetFirstStoredBlockNumber(ctx, chainID)
		if err != nil {
			return nil, fmt.Errorf("could not get start height: %w", err)
		}

		startHeight = uint64(*getStartHeight.Response)
	}

	if endHeight < startHeight {
		return nil, fmt.Errorf("end height must be greater than start height")
	}

	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		startTime, err := s.FetchClient.GetBlockTime(ctx, chainID, int(startHeight))
		if err != nil || startTime == nil || startTime.Response == nil {
			return fmt.Errorf("could not get start time: %w", err)
		}

		output.startBlock = block{
			number: startHeight,
			time:   uint64(*startTime.Response),
		}

		return nil
	})

	g.Go(func() error {
		endTime, err := s.FetchClient.GetBlockTime(ctx, chainID, int(endHeight))
		if err != nil || endTime == nil || endTime.Response == nil {
			return fmt.Errorf("could not get end time: %w", err)
		}

		output.endBlock = block{
			number: endHeight,
			time:   uint64(*endTime.Response),
		}

		return nil
	})

	if err := g.Wait(); err != nil {
		return nil, fmt.Errorf("could not get search range: %w", err)
	}

	return &output, nil
}
