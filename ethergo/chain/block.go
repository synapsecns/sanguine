package chain

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/core/types"
	lru "github.com/hashicorp/golang-lru"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
)

// HeaderByTime gets the closest block to a given time.
//
// It works by running a binary search against block headers from genesis until now.
// the user can optionally pass in a nullable startBlock *big.Int. If the block is left empty the search is performed from genesis
// the worst case time complexity makes this a cheap operation on any chain, but if done frequently a startBlock should be passed in as a heuristic
//
//nolint:gocognit,cyclop // we ignore complexity warnings here: breaking this up makes it much harder to read.
func (b *baseChain) HeaderByTime(ctx context.Context, startHeight *big.Int, targetTime time.Time) (*types.Header, error) {
	// grab the MeteredEVMClient explicitly to avoid ambiguous references
	client, err := newHeaderCacheClient(b, 4)
	if err != nil {
		return nil, fmt.Errorf("could not create header cache client: %w", err)
	}
	target := uint64(targetTime.Unix())

	// get the start block and end block
	searchRange, err := getSearchRange(ctx, client, startHeight)
	if err != nil {
		return nil, err
	}

	// handle corner cases
	if target <= searchRange.StartTime() {
		return searchRange.startBlock, nil
	}
	if target >= searchRange.EndTime() {
		return searchRange.endBlock, nil
	}

	// run the binary search
	var i, mid, j uint64
	i = searchRange.startBlock.Number.Uint64()
	j = searchRange.endBlock.Number.Uint64()

	for i < j {
		mid = (i + j) / 2

		midBlock, err := client.HeaderByNumber(ctx, big.NewInt(0).SetUint64(mid))
		if err != nil {
			return nil, errors.Wrapf(err, "could not get block %d", mid)
		}

		if midBlock.Time == target {
			return midBlock, nil
		}

		// If target is less than midBlock element
		// then search in left
		//nolint: nestif // we want to keep the logic for the binary search together
		if target < midBlock.Time {
			// If target is greater than previous to mid, return closest of two
			midSubBlock, err := client.HeaderByNumber(ctx, big.NewInt(0).SetUint64(mid-1))
			if err != nil {
				return nil, errors.Wrapf(err, "could not get block %d", mid-1)
			}

			if mid > 0 && target >= midSubBlock.Time {
				return getClosest(midSubBlock, midBlock, target), nil
			}

			j = mid
		} else {
			// If target is greater than mid
			midPlusBlock, err := client.HeaderByNumber(ctx, big.NewInt(0).SetUint64(mid+1))
			if err != nil {
				return nil, errors.Wrapf(err, "could not get block %d", mid+1)
			}

			if mid <= searchRange.Length()-1 && target <= midPlusBlock.Time {
				return getClosest(midBlock, midPlusBlock, target), nil
			}
			i = mid + 1
		}
	}

	// only a single element is left after the search. Fetch the block (from cache) and return it.
	resultingBlock, err := client.HeaderByNumber(ctx, big.NewInt(int64(mid)))
	if err != nil {
		return nil, fmt.Errorf("could not get resultng block")
	}

	return resultingBlock, nil
}

// getClosest gets the closer value to the target of the two blocks
//
// We find the closest by taking the difference between the target and both values
// this method assumes that blockB is greater than blockA and target lies between the two.
func getClosest(lesser *types.Header, greater *types.Header, target uint64) *types.Header {
	if (target - lesser.Time) >= (greater.Time - target) {
		return greater
	}
	return lesser
}

// getSearchRange gets the block range used for performing a binary search.
// an error is returned if either startBlock or endBlock is null. Uses 1 if startBlock is nil.
func getSearchRange(ctx context.Context, client ethereum.ChainReader, startHeight *big.Int) (_ *blockRange, err error) {
	var output blockRange
	if startHeight == nil || startHeight.Cmp(big.NewInt(0)) == 0 {
		startHeight = big.NewInt(1)
	}

	g, ctx := errgroup.WithContext(ctx)
	g.Go(func() error {
		output.startBlock, err = client.HeaderByNumber(ctx, startHeight)
		if err != nil {
			return fmt.Errorf("could not get start block: %w", err)
		}
		return nil
	})

	g.Go(func() error {
		output.endBlock, err = client.HeaderByNumber(ctx, nil)
		if err != nil {
			return fmt.Errorf("could not get start block: %w", err)
		}
		return nil
	})

	err = g.Wait()
	if err != nil {
		return nil, fmt.Errorf("unable to get search range: %w", err)
	}
	return &output, nil
}

// blockRange contains data about a range of block
// accessor methods are exposed for performing a binary search.
type blockRange struct {
	// startBlock of the range. This should not be nil
	startBlock *types.Header

	// endBlock of the range. Will always be greater than start block
	endBlock *types.Header
}

// Length gets the length of the block range as a uint64.
func (b *blockRange) Length() uint64 {
	return big.NewInt(0).Sub(b.endBlock.Number, b.startBlock.Number).Uint64() + 1
}

// StartTime gets the start time of the range as a unix timestamp.
func (b *blockRange) StartTime() uint64 {
	return b.startBlock.Time
}

// EndTime gets the end time of the range as a unix timestamp.
func (b *blockRange) EndTime() uint64 {
	return b.endBlock.Time
}

// headerCacheClient is a client that caches the last count headers
// we use this here so we don't have to store midBlock/handle instances of cached vs uncached blocks differently.
type headerCacheClient struct {
	ethereum.ChainReader
	// headerCache holds the headers
	headerCache *lru.Cache
}

// newHeaderCacheClient creates a client that contains a header cache
// the cache checks if the header exists and if not fetches it. If the cache size
// exceeds maxSize the oldest item is removed.
//
// this cache is intended for use with HeaderByTime to prevent refetching blocks when iterating
// but can be used generally. This will also not handle reorgs since doing so would require a refetch.
func newHeaderCacheClient(ethClient ethereum.ChainReader, count int) (*headerCacheClient, error) {
	cache, err := lru.New(count)
	if err != nil {
		return nil, fmt.Errorf("could not create header cache: %w", err)
	}
	return &headerCacheClient{
		ChainReader: ethClient,
		headerCache: cache,
	}, nil
}

// HeaderByNumber wraps header cache client and introduces caching.
func (c *headerCacheClient) HeaderByNumber(ctx context.Context, number *big.Int) (header *types.Header, err error) {
	// handle latest block query. We cache this by number after the fetch
	if number == nil {
		return c.fetchAndStoreHeader(ctx, number)
	}

	rawHeader, hasKey := c.headerCache.Get(number.Uint64())
	if hasKey {
		var ok bool
		header, ok = rawHeader.(*types.Header)
		// we should always be able to cast here, but if we can't fetch and store the header
		if !ok {
			logger.Warnf("couldn't cast header of type %T to %T. cache may be corrupted", rawHeader, &types.Header{})
			return c.fetchAndStoreHeader(ctx, number)
		}
		return header, nil
	}

	return c.fetchAndStoreHeader(ctx, number)
}

// fetchAndStoreHeader fetches and stores the header bypassing the cache.
func (c *headerCacheClient) fetchAndStoreHeader(ctx context.Context, number *big.Int) (header *types.Header, err error) {
	fetchedBlock, err := c.ChainReader.HeaderByNumber(ctx, number)
	if err != nil {
		return nil, fmt.Errorf("could not fetch latest block: %w", err)
	}
	// add the header to the cache
	c.headerCache.Add(fetchedBlock.Number.Uint64(), fetchedBlock)
	return fetchedBlock, nil
}
