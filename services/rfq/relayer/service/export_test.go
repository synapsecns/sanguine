package service

import "context"

// StartChainParser exports chain parser for testing.
func (r *Relayer) StartChainParser(ctx context.Context) error {
	return r.startChainParser(ctx)
}
