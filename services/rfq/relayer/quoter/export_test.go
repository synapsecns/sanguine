package quoter

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/synapsecns/sanguine/services/rfq/api/model"
	"math/big"
)

func (m *Manager) GenerateQuotes(ctx context.Context, chainID int, address common.Address, balance *big.Int) ([]model.PutQuoteRequest, error) {
	// nolint: errcheck
	return m.generateQuotes(ctx, chainID, address, balance)
}
