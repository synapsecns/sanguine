package quoter

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/synapsecns/sanguine/services/rfq/api/model"
	"github.com/synapsecns/sanguine/services/rfq/relayer/relconfig"
)

func (m *Manager) GenerateQuotes(ctx context.Context, chainID int, address common.Address, balance *big.Int) ([]model.PutQuoteRequest, error) {
	// nolint: errcheck
	return m.generateQuotes(ctx, chainID, address, balance)
}

func (m *Manager) GetOriginAmount(ctx context.Context, origin, dest int, address common.Address, balance *big.Int) (*big.Int, error) {
	return m.getOriginAmount(ctx, origin, dest, address, balance)
}

func (m *Manager) GetDestAmount(ctx context.Context, quoteAmount *big.Int, chainID int, tokenName string) (*big.Int, error) {
	return m.getDestAmount(ctx, quoteAmount, chainID, tokenName)
}

func (m *Manager) SetConfig(cfg relconfig.Config) {
	m.config = cfg
}

func (m *Manager) SetRelayPaused(relayPaused bool) {
	m.relayPaused.Store(relayPaused)
}
