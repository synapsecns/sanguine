package client

import (
	"context"
	"github.com/synapsecns/sanguine/core"
	"github.com/synapsecns/sanguine/core/metrics"
	"math/big"
)

type chainIDClientImpl struct {
	EVM
	chainID *big.Int
}

// GetBigChainID returns the chain ID as a big.Int.
func (c chainIDClientImpl) GetBigChainID() *big.Int {
	return core.CopyBigInt(c.chainID)
}

// EVMChainID is an EVM client that also exposes the chain ID.
type EVMChainID interface {
	EVM
	GetBigChainID() *big.Int
}

// DialBackendChainID dials an EVM backend and returns an EVM client.
func DialBackendChainID(ctx context.Context, chainID *big.Int, url string, handler metrics.Handler, opts ...Options) (_ EVMChainID, err error) {
	underlying, err := DialBackend(ctx, url, handler, opts...)
	if err != nil {
		return nil, err
	}
	return chainIDClientImpl{
		EVM:     underlying,
		chainID: core.CopyBigInt(chainID),
	}, nil
}
