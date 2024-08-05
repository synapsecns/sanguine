package service

import (
	"context"

	"github.com/synapsecns/sanguine/services/rfq/relayer/quoter"
	"github.com/synapsecns/sanguine/services/rfq/relayer/reldb"
)

// StartChainParser exports chain parser for testing.
func (r *Relayer) StartChainParser(ctx context.Context) error {
	return r.startChainIndexers(ctx)
}

func (r *Relayer) SetQuoter(quoter quoter.Quoter) {
	r.quoter = quoter
}

func (r *Relayer) DB() reldb.Service {
	return r.db
}
