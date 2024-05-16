package manager

import (
	"context"

	"github.com/synapsecns/sanguine/committee/devnet/config"
	"github.com/synapsecns/sanguine/committee/devnet/provisioner"
	"github.com/synapsecns/sanguine/committee/devnet/sender"
	"golang.org/x/sync/errgroup"
)

// Manager provisions the contracts and submits verifications to the network.
type Manager struct {
	*provisioner.Provisioner
	*sender.Sender
}

// NewManager creates a new Manager.
func NewManager(provisioner *provisioner.Provisioner, sender *sender.Sender) *Manager {
	return &Manager{
		Provisioner: provisioner,
		Sender:      sender,
	}
}

func (m *Manager) Start(ctx context.Context, pc config.ProvisionerConfig, sc config.SenderConfig) error {
	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		return m.Provisioner.Run(ctx, pc)
	})

	if m.Sender != nil {
		g.Go(func() error {
			return m.Sender.Start(ctx, sc)
		})
	}

	if err := g.Wait(); err != nil {
		return err
	}

	return nil
}
