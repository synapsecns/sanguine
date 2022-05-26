// Package evm TODO description
package evm

import (
	"context"
	"fmt"
	"github.com/synapsecns/sanguine/core/config"
	"github.com/synapsecns/sanguine/core/domains"
	"github.com/synapsecns/synapse-node/pkg/evm"
)

type evmClient struct {
	// name is the name of the evm client
	name string
	// config is the config of the evm client
	config config.DomainConfig
	// client uses the old synapse client for now
	client evm.Chain
}

var _ domains.DomainClient = &evmClient{}

// NewEVM creates a new evm client.
func NewEVM(ctx context.Context, name string, domain config.DomainConfig) (domains.DomainClient, error) {
	underlyingClient, err := evm.NewFromURL(ctx, domain.RPCUrl)
	if err != nil {
		return nil, fmt.Errorf("could not get evm: %w", err)
	}

	return evmClient{
		name:   name,
		config: domain,
		client: underlyingClient,
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
