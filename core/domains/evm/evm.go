// Package evm TODO description
package evm

import (
	"context"
	"errors"
	"fmt"
	"github.com/cockroachdb/pebble"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/synapsecns/sanguine/core/config"
	"github.com/synapsecns/sanguine/core/contracts/xappconfig"
	"github.com/synapsecns/sanguine/core/db"
	"github.com/synapsecns/sanguine/core/domains"
	"github.com/synapsecns/synapse-node/pkg/evm"
	"github.com/synapsecns/synapse-node/pkg/teller/backfiller"
	"math/big"
)

type evmClient struct {
	// name is the name of the evm client
	name string
	// config is the config of the evm client
	config config.DomainConfig
	// client uses the old synapse client for now
	client evm.Chain
	// db stores the db handle
	db db.DB
	// xAppConfig is the xAppConfig handle
	xAppConfig *xappconfig.XAppConfigRef
}

var _ domains.DomainClient = &evmClient{}

// NewEVM creates a new evm client.
func NewEVM(ctx context.Context, name string, domain config.DomainConfig, db db.DB) (domains.DomainClient, error) {
	underlyingClient, err := evm.NewFromURL(ctx, domain.RPCUrl)
	if err != nil {
		return nil, fmt.Errorf("could not get evm: %w", err)
	}

	xAppConfig, err := xappconfig.NewXAppConfigRef(common.HexToAddress(domain.XAppConfigAddress), underlyingClient)
	if err != nil {
		return nil, fmt.Errorf("could not create xappconfig handle: %w", err)
	}

	return evmClient{
		name:       name,
		config:     domain,
		client:     underlyingClient,
		db:         db,
		xAppConfig: xAppConfig,
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

// FetchStoredUpdates fetches stored updates.
func (e evmClient) FetchStoredUpdates(ctx context.Context, from uint32, to uint32) error {
	indexedHeight, err := e.db.GetIndexedHeight(fmt.Sprintf("%d", e.config.DomainID))
	if errors.Is(err, pebble.ErrNotFound) {
		// Get deployed height of contract
		indexedHeight = 0
	} else if err != nil {
		return fmt.Errorf("could not get indexed height: %w", err)
	}
	_ = indexedHeight

	homeAddress, err := e.xAppConfig.Home(&bind.CallOpts{Context: ctx})
	if err != nil {
		return fmt.Errorf("could not add home address: %w", err)
	}

	currentHeight, err := e.client.BlockNumber(ctx)
	if err != nil {
		return fmt.Errorf("could not set current height: %w", err)
	}

	rangeFilter := backfiller.NewRangeFilter(homeAddress, e.client, big.NewInt(int64(from)), new(big.Int).SetUint64(currentHeight), 1200, true)

	err = rangeFilter.Start(ctx)
	if err != nil {
		return fmt.Errorf("could not filter ranges: %w", err)
	}
	return nil
}
