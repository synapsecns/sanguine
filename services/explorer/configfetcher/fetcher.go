package configfetcher

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/synapsecns/sanguine/services/explorer/contracts/bridgeconfig"
	"math/big"
)

type Fetcher interface {
	// GetTokenID gets the token id for the bridge config at a given block number
	GetTokenID(ctx context.Context, chainID, block uint32, tokenAddress common.Address) (tokenID string, err error)
	// GetToken gets the token for the bridge config at a given block number
	GetToken(ctx context.Context, chainID, block uint32, tokenId string) (token bridgeconfig.BridgeConfigV3Token, err error)
}

type fetcher struct {
	bridgeConfig *bridgeconfig.BridgeConfigRef
}

func (f fetcher) GetTokenID(ctx context.Context, chainID, block uint32, tokenAddress common.Address) (tokenID string, err error) {
	tokenID, err = f.bridgeConfig.GetTokenID(&bind.CallOpts{
		BlockNumber: big.NewInt(int64(block)),
		Context:     ctx,
	}, tokenAddress, big.NewInt(int64(chainID)))
	if err != nil {
		return "", fmt.Errorf("could not get token id: %w", err)
	}

	if tokenID == "" {
		return "", fmt.Errorf("couldn't find token id for address %s and chain id %d: %w", tokenAddress.String(), chainID, ErrTokenDoesNotExist)
	}

	return tokenID, nil
}

func (f fetcher) GetToken(ctx context.Context, chainID, block uint32, tokenId string) (token bridgeconfig.BridgeConfigV3Token, err error) {
	tok, err := f.bridgeConfig.GetToken(&bind.CallOpts{
		BlockNumber: big.NewInt(int64(block)),
		Context:     ctx,
	}, tokenId, big.NewInt(int64(chainID)))
	if err != nil {
		var none bridgeconfig.BridgeConfigV3Token
		return none, fmt.Errorf("could not get token id: %w", err)
	}
	return tok, nil
}

var _ Fetcher = &fetcher{}

// NewFetcher creates a new config fetcher.
func NewFetcher(bridgeConfigAddress common.Address, backend bind.ContractBackend) (Fetcher, error) {
	bridgeConfig, err := bridgeconfig.NewBridgeConfigRef(bridgeConfigAddress, backend)
	if err != nil {
		return nil, fmt.Errorf("could not bind bridge config contract: %w", err)
	}

	return &fetcher{bridgeConfig}, nil
}
