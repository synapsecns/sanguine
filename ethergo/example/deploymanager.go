package example

import (
	"context"
	"github.com/synapsecns/sanguine/ethergo/backends"
	"github.com/synapsecns/sanguine/ethergo/contracts"
	"github.com/synapsecns/sanguine/ethergo/example/counter"
	"github.com/synapsecns/sanguine/ethergo/manager"
	"testing"
)

// DeployManager wraps DeployManager and allows typed contract handles to be returned.
type DeployManager struct {
	*manager.DeployerManager
}

// NewDeployManager creates a new DeployManager.
func NewDeployManager(t *testing.T) *DeployManager {
	t.Helper()

	parentManager := manager.NewDeployerManager(t, NewCounterDeployer)
	return &DeployManager{parentManager}
}

// GetCounter gets the pre-created counter.
func (d *DeployManager) GetCounter(ctx context.Context, backend backends.SimulatedTestBackend) (contract contracts.DeployedContract, handle *counter.CounterRef) {
	d.T().Helper()

	return manager.GetContract[*counter.CounterRef](ctx, d.T(), d, backend, CounterType)
}
