package consumer

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/services/explorer/consumer/client"
	"github.com/synapsecns/sanguine/services/explorer/contracts/bridgeconfig"
	"github.com/synapsecns/sanguine/services/scribe/graphql"
	"golang.org/x/sync/errgroup"
	"math/big"
)

// Fetcher is the fetcher for the events. It uses GQL.
type Fetcher struct {
	FetchClient *client.Client
}

// NewFetcher creates a new fetcher.
func NewFetcher(fetchClient *client.Client) *Fetcher {
	return &Fetcher{
		FetchClient: fetchClient,
	}
}

// FetchLogsInRange fetches logs in a range with the GQL client.
func (f Fetcher) FetchLogsInRange(ctx context.Context, chainID uint32, startBlock, endBlock uint64) ([]ethTypes.Log, error) {
	logs := &client.GetLogsRange{}
	page := 1
	for {
		paginatedLogs, err := f.FetchClient.GetLogsRange(ctx, int(chainID), int(startBlock), int(endBlock), page)
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

// TimeToBlockNumber returns the first block number after the given time.
func (f Fetcher) TimeToBlockNumber(ctx context.Context, chainID uint32, startHeight, targetTime uint64) (uint64, error) {
	// get the start and end block
	searchRange, err := f.getSearchRange(ctx, startHeight, chainID)
	if err != nil {
		return 0, fmt.Errorf("could not get search range: %w", err)
	}

	// handle corner cases
	if targetTime <= searchRange.StartTime() {
		return searchRange.startBlock.number, nil
	}
	if targetTime >= searchRange.EndTime() {
		return searchRange.endBlock.number, nil
	}

	// run the binary search
	var i, mid, j uint64
	i = searchRange.startBlock.number
	j = searchRange.endBlock.number

	for i < j {
		mid = (i + j) / 2

		midBlock, err := f.FetchClient.GetBlockTime(ctx, int(chainID), int(mid))
		if err != nil || midBlock == nil || midBlock.Response == nil {
			return 0, fmt.Errorf("could not get mid time: %w", err)
		}

		midTime := uint64(*midBlock.Response)

		if midTime == targetTime {
			return mid, nil
		}

		// If target is less than midBlock element
		// then search in left
		//nolint: nestif // we want to keep the logic for the binary search together
		if targetTime < midTime {
			// If target is greater than previous to mid, return the closest of the two
			midSubBlock, err := f.FetchClient.GetBlockTime(ctx, int(chainID), int(mid-1))
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
			// If target is greater than mid
			midPlusBlock, err := f.FetchClient.GetBlockTime(ctx, int(chainID), int(mid+1))
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

	// only a single element is left after the search. Fetch the block and return it.
	resultingBlock, err := f.FetchClient.GetBlockTime(ctx, int(chainID), int(mid))
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

func (f Fetcher) getSearchRange(ctx context.Context, startHeight uint64, chainID uint32) (*blockRange, error) {
	getEndHeight, err := f.FetchClient.GetLastStoredBlockNumber(ctx, int(chainID))
	if err != nil {
		return nil, fmt.Errorf("could not get end height: %w", err)
	}
	endHeight := uint64(*getEndHeight.Response)
	var output blockRange
	if startHeight == 0 {
		getStartHeight, err := f.FetchClient.GetFirstStoredBlockNumber(ctx, int(chainID))
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
		startTime, err := f.FetchClient.GetBlockTime(ctx, int(chainID), int(startHeight))
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
		endTime, err := f.FetchClient.GetBlockTime(ctx, int(chainID), int(endHeight))
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

// BridgeConfigFetcher is the fetcher for the bridge config contract.
type BridgeConfigFetcher struct {
	bridgeConfig        *bridgeconfig.BridgeConfigRef
	bridgeConfigAddress common.Address
} // TODO switch bridge config based on block number

// NewBridgeConfigFetcher creates a new config fetcher.
// Backend must be an archive backend.
func NewBridgeConfigFetcher(bridgeConfigAddress common.Address, backend bind.ContractBackend) (*BridgeConfigFetcher, error) {
	bridgeConfig, err := bridgeconfig.NewBridgeConfigRef(bridgeConfigAddress, backend)
	if err != nil {
		return nil, fmt.Errorf("could not bind bridge config contract: %w", err)
	}
	return &BridgeConfigFetcher{bridgeConfig, bridgeConfigAddress}, nil
}

// GetTokenID gets the token id from the bridge config contract.
func (b *BridgeConfigFetcher) GetTokenID(ctx context.Context, chainID uint32, tokenAddress common.Address) (tokenID *string, err error) {
	tokenIDStr, err := b.bridgeConfig.GetTokenID(&bind.CallOpts{
		Context: ctx,
	}, tokenAddress, big.NewInt(int64(chainID)))
	if err != nil {
		return nil, fmt.Errorf("could not get token id: %w", err)
	}

	if tokenIDStr == "" {
		return nil, fmt.Errorf("couldn't find token id for address %s and chain id %d: %w", tokenAddress, chainID, ErrTokenDoesNotExist)
	}

	return &tokenIDStr, nil
}

// GetToken gets the token from the bridge config contract.
func (b *BridgeConfigFetcher) GetToken(ctx context.Context, chainID, block uint32, tokenID string) (token *bridgeconfig.BridgeConfigV3Token, err error) {
	tok, err := b.bridgeConfig.GetToken(&bind.CallOpts{
		BlockNumber: big.NewInt(int64(block)),
		Context:     ctx,
	}, tokenID, big.NewInt(int64(chainID)))
	if err != nil {
		// var none bridgeconfig.BridgeConfigV3Token
		return nil, fmt.Errorf("could not get token id: %w", err)
	}
	return &tok, nil
}
