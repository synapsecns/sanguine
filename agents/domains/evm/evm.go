// Package evm TODO description
package evm

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/synapsecns/sanguine/agents/config"
	"github.com/synapsecns/sanguine/agents/domains"
	"github.com/synapsecns/sanguine/ethergo/chain"
)

type evmClient struct {
	// name is the name of the evm client
	name string
	// config is the config of the evm client
	config config.DomainConfig
	// client uses the old synapse client for now
	//nolint: staticcheck
	client chain.Chain
	// origin contains the origin contract
	origin domains.OriginContract
	// summit contains the summit contract
	summit domains.SummitContract
	// destination contains the destination contract
	destination domains.DestinationContract
	// lightManager contains the lightManager contract
	lightManager domains.LightManagerContract
	// bondingManager contains the bondingManager contract
	bondingManager domains.BondingManagerContract
	// lightInbox contains the lightInbox contract
	lightInbox domains.LightInboxContract
	// inbox contains the inbox contract
	inbox domains.InboxContract
}

var _ domains.DomainClient = &evmClient{}

// NewEVM creates a new evm client.
//
//nolint:nestif
func NewEVM(ctx context.Context, name string, domain config.DomainConfig, chainRPCURL string) (domains.DomainClient, error) {
	underlyingClient, err := chain.NewFromURL(ctx, chainRPCURL)
	if err != nil {
		return nil, fmt.Errorf("could not get evm: %w", err)
	}

	boundOrigin, err := NewOriginContract(ctx, underlyingClient, common.HexToAddress(domain.OriginAddress))
	if err != nil {
		return nil, fmt.Errorf("could not bind origin contract: %w", err)
	}

	boundDestination, err := NewDestinationContract(ctx, underlyingClient, common.HexToAddress(domain.DestinationAddress))
	if err != nil {
		return nil, fmt.Errorf("could not bind destination contract: %w", err)
	}

	var boundSummit domains.SummitContract
	var boundBondingManager domains.BondingManagerContract
	var boundInbox domains.InboxContract
	var boundLightManager domains.LightManagerContract
	var boundLightInbox domains.LightInboxContract
	if domain.SummitAddress == "" {
		boundLightManager, err = NewLightManagerContract(ctx, underlyingClient, common.HexToAddress(domain.LightManagerAddress))
		if err != nil {
			return nil, fmt.Errorf("could not bind light manager contract: %w", err)
		}
		boundLightInbox, err = NewLightInboxContract(ctx, underlyingClient, common.HexToAddress(domain.LightInboxAddress))
		if err != nil {
			return nil, fmt.Errorf("could not bind light inbox contract: %w", err)
		}
	} else {
		boundSummit, err = NewSummitContract(ctx, underlyingClient, common.HexToAddress(domain.SummitAddress))
		if err != nil {
			return nil, fmt.Errorf("could not bind summit contract: %w", err)
		}

		boundBondingManager, err = NewBondingManagerContract(ctx, underlyingClient, common.HexToAddress(domain.BondingManagerAddress))
		if err != nil {
			return nil, fmt.Errorf("could not bind bonding manager contract: %w", err)
		}
		boundInbox, err = NewInboxContract(ctx, underlyingClient, common.HexToAddress(domain.InboxAddress))
		if err != nil {
			return nil, fmt.Errorf("could not bind inbox contract: %w", err)
		}
	}

	return evmClient{
		name:           name,
		config:         domain,
		client:         underlyingClient,
		summit:         boundSummit,
		origin:         boundOrigin,
		destination:    boundDestination,
		lightManager:   boundLightManager,
		lightInbox:     boundLightInbox,
		bondingManager: boundBondingManager,
		inbox:          boundInbox,
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

// LightManager gets the light manager.
func (e evmClient) LightManager() domains.LightManagerContract {
	return e.lightManager
}

// BondingManager gets the bonding manager.
func (e evmClient) BondingManager() domains.BondingManagerContract {
	return e.bondingManager
}

// LightInbox gets the light inbox.
func (e evmClient) LightInbox() domains.LightInboxContract {
	return e.lightInbox
}

// Inbox gets the inbox.
func (e evmClient) Inbox() domains.InboxContract {
	return e.inbox
}
