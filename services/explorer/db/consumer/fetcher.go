package consumer

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/services/explorer/contracts/bridgeconfig"
	"github.com/synapsecns/sanguine/services/explorer/db/consumer/client"
	"github.com/synapsecns/sanguine/services/scribe/graphql"
	"math/big"
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

type BridgeConfigFetcher struct {
	bridgeConfig        *bridgeconfig.BridgeConfigRef
	bridgeConfigAddress common.Address
}

// NewBridgeConfigFetcher creates a new config fetcher.
func NewBridgeConfigFetcher(bridgeConfigAddress common.Address, backend bind.ContractBackend) (*BridgeConfigFetcher, error) {
	bridgeConfig, err := bridgeconfig.NewBridgeConfigRef(bridgeConfigAddress, backend)
	if err != nil {
		return nil, fmt.Errorf("could not bind bridge config contract: %w", err)
	}
	return &BridgeConfigFetcher{bridgeConfig, bridgeConfigAddress}, nil
}

func (b *BridgeConfigFetcher) GetTokenID(ctx context.Context, chainID, block uint32, tokenAddress common.Address) (tokenID *string, err error) {
	tokenIDStr, err := b.bridgeConfig.GetTokenID(&bind.CallOpts{
		BlockNumber: big.NewInt(int64(block)),
		Context:     ctx,
	}, tokenAddress, big.NewInt(int64(chainID)))
	if err != nil {
		return nil, fmt.Errorf("could not get token id: %w", err)
	}

	if tokenIDStr == "" {
		return nil, fmt.Errorf("couldn't find token id for address %s and chain id %d: %w", b.bridgeConfigAddress.String(), chainID, ErrTokenDoesNotExist)
	}

	return &tokenIDStr, nil
}

func (b *BridgeConfigFetcher) GetToken(ctx context.Context, chainID, block uint32, tokenId string) (token bridgeconfig.BridgeConfigV3Token, err error) {
	tok, err := b.bridgeConfig.GetToken(&bind.CallOpts{
		BlockNumber: big.NewInt(int64(block)),
		Context:     ctx,
	}, tokenId, big.NewInt(int64(chainID)))
	if err != nil {
		var none bridgeconfig.BridgeConfigV3Token
		return none, fmt.Errorf("could not get token id: %w", err)
	}
	return tok, nil
}
