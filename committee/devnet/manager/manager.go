package manager

import (
	"context"

	"github.com/synapsecns/sanguine/committee/devnet/config"
	"github.com/synapsecns/sanguine/committee/devnet/provisioner"
	"github.com/synapsecns/sanguine/committee/devnet/sender"
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

func (m *Manager) Start(ctx context.Context, pc config.ProvisionerConfig, sc *config.SenderConfig) error {

	if err := m.Provisioner.Run(ctx, pc); err != nil {
		return err
	}

	if sc != nil {
		if err := m.Sender.Start(ctx, sc); err != nil {
			return err
		}
	}

	return nil
}
