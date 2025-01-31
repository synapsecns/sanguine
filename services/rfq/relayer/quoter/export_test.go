package quoter

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/synapsecns/sanguine/services/rfq/api/model"
	"github.com/synapsecns/sanguine/services/rfq/relayer/relconfig"
)

func (m *Manager) GenerateQuotes(ctx context.Context, chainID int, address common.Address, balance *big.Int, inv map[int]map[common.Address]*big.Int) ([]model.PutRelayerQuoteRequest, error) {
	// nolint: errcheck
	return m.generateQuotes(ctx, chainID, address, balance, inv)
}

func (m *Manager) GetOriginAmount(ctx context.Context, input QuoteInput) (*big.Int, error) {
	return m.getOriginAmount(ctx, input)
}

func (m *Manager) GetDestAmount(ctx context.Context, quoteAmount *big.Int, tokenName string, input QuoteInput) (*big.Int, error) {
	return m.getDestAmount(ctx, quoteAmount, tokenName, input)
}

func (m *Manager) GenerateActiveRFQ(ctx context.Context, msg *model.ActiveRFQMessage) (resp *model.ActiveRFQMessage, err error) {
	return m.generateActiveRFQ(ctx, msg)
}

func (m *Manager) SetConfig(cfg relconfig.Config) {
	m.config = cfg
}

func (m *Manager) SetRelayPaused(relayPaused bool) {
	m.relayPaused.Store(relayPaused)
}
