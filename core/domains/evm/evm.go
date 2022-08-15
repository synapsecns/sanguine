// Package evm TODO description
package evm

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
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
	// origin contains the origin contract
	origin domains.OriginContract
	// attestationCollecotr contains the attestation collector contract
	attestationCollector domains.AttestationCollectorContract
}

var _ domains.DomainClient = &evmClient{}

// NewEVM creates a new evm client.
func NewEVM(ctx context.Context, name string, domain config.DomainConfig) (domains.DomainClient, error) {
	underlyingClient, err := evm.NewFromURL(ctx, domain.RPCUrl)
	if err != nil {
		return nil, fmt.Errorf("could not get evm: %w", err)
	}

	boundOrigin, err := NewOriginContract(ctx, underlyingClient, common.HexToAddress(domain.OriginAddress))
	if err != nil {
		return nil, fmt.Errorf("could not bind origin contract: %w", err)
	}

	boundCollector, err := NewAttestationCollectorContract(ctx, underlyingClient, common.HexToAddress(domain.AttesationCollectorAddress))
	if err != nil {
		return nil, fmt.Errorf("could not bind attestation contract: %w", err)
	}

	return evmClient{
		name:                 name,
		config:               domain,
		client:               underlyingClient,
		attestationCollector: boundCollector,
		origin:               boundOrigin,
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

// AttestationCollector gets the attestation collector.
func (e evmClient) AttestationCollector() domains.AttestationCollectorContract {
	return e.attestationCollector
}
