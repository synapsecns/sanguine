// Package evm TODO description
package evm

import (
	"context"
	"fmt"
	"github.com/synapsecns/sanguine/core/metrics"
	client2 "github.com/synapsecns/sanguine/ethergo/client"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/synapsecns/sanguine/agents/config"
	"github.com/synapsecns/sanguine/agents/domains"
)

type evmClient struct {
	// name is the name of the evm client
	name string
	// config is the config of the evm client
	config config.DomainConfig
	// client uses the old synapse client for now
	//nolint: staticcheck
	client client2.EVMChainID
	// origin contains the origin contract
	origin domains.OriginContract
	// summit contains the summit contract
	summit domains.SummitContract
	// destination contains the destination contract
	destination domains.DestinationContract
}

var _ domains.DomainClient = &evmClient{}

// NewEVM creates a new evm client.
func NewEVM(ctx context.Context, name string, domain config.DomainConfig, handler metrics.Handler) (domains.DomainClient, error) {
	underlyingClient, err := client2.DialBackendChainID(ctx, big.NewInt(int64(domain.DomainID)), domain.RPCUrl, handler, client2.Capture(true))
	if err != nil {
		return nil, fmt.Errorf("could not get evm: %w", err)
	}

	boundOrigin, err := NewOriginContract(ctx, underlyingClient, common.HexToAddress(domain.OriginAddress))
	if err != nil {
		return nil, fmt.Errorf("could not bind origin contract: %w", err)
	}

	boundSummit, err := NewSummitContract(ctx, underlyingClient, common.HexToAddress(domain.SummitAddress))
	if err != nil {
		return nil, fmt.Errorf("could not bind attestation contract: %w", err)
	}

	boundDestination, err := NewDestinationContract(ctx, underlyingClient, common.HexToAddress(domain.DestinationAddress))
	if err != nil {
		return nil, fmt.Errorf("could not bind destination contract: %w", err)
	}

	return evmClient{
		name:        name,
		config:      domain,
		client:      underlyingClient,
		summit:      boundSummit,
		origin:      boundOrigin,
		destination: boundDestination,
	}, nil
}

// Name gets the name of the evm client.
func (e evmClient) Name() string {
	return e.name
}

// Config gets the config the evm client was initiated with.
func (e evmClient) Config() config.DomainConfig {
	return e.config
}

// BlockNumber gets the latest block number.
func (e evmClient) BlockNumber(ctx context.Context) (uint32, error) {
	blockNumber, err := e.client.BlockNumber(ctx)
	if err != nil {
		return 0, fmt.Errorf("could not get block number: %w", err)
	}

	return uint32(blockNumber), nil
}

// Origin returns the bound origin contract.
func (e evmClient) Origin() domains.OriginContract {
	return e.origin
}

// Summit gets the summit.
func (e evmClient) Summit() domains.SummitContract {
	return e.summit
}

// Destination gets the destination.
func (e evmClient) Destination() domains.DestinationContract {
	return e.destination
}
