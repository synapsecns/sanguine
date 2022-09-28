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
	"math/big"
)

// Fetcher is the fetcher for the events. It uses GQL.
type Fetcher struct {
	fetchClient *client.Client
}

// NewFetcher creates a new fetcher.
func NewFetcher(fetchClient *client.Client) *Fetcher {
	return &Fetcher{
		fetchClient: fetchClient,
	}
}

// FetchLogsInRange fetches logs in a range with the GQL client.
func (f Fetcher) FetchLogsInRange(ctx context.Context, chainID uint32, startBlock, endBlock uint64) ([]ethTypes.Log, error) {
	logs := &client.GetLogsRange{}
	page := 1
	for {
		paginatedLogs, err := f.fetchClient.GetLogsRange(ctx, int(chainID), int(startBlock), int(endBlock), page)
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
