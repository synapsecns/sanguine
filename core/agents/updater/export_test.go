package updater

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
)

// Update wraps update and produces one if the update has not been produced.
func (u UpdateProducer) Update(ctx context.Context) error {
	return u.update(ctx)
}

// Update wraps update and submits an update if one exists in the db
func (u UpdateSubmitter) Update(ctx context.Context, committedRoot common.Hash) (common.Hash, error) {
	return u.update(ctx, committedRoot)
}
