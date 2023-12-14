package notary

import (
	"context"

	"github.com/synapsecns/sanguine/agents/types"
)

func (n *Notary) EnsureNotaryActive(ctx context.Context) error {
	return n.ensureNotaryActive(ctx)
}

func (n *Notary) LoadMyLatestStates(ctx context.Context) {
	n.loadMyLatestStates(ctx)
}

func (n *Notary) MyLatestStates(ctx context.Context) map[uint32]types.State {
	return n.myLatestStates
}
