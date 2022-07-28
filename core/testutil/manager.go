package testutil

import (
	"github.com/synapsecns/sanguine/ethergo/manager"
	"testing"
)

// NewDeployManager creates a deploy manager.
func NewDeployManager(t *testing.T) *DeployManager {
	t.Helper()

	parentManager := manager.NewDeployerManager(t, NewHomeDeployer, NewXAppConfigDeployer, NewMessageHarnessDeployer, NewHomeHarnessDeployer, NewUpdateManagerDeployer, NewAttestationCollectorDeployer)
	return &DeployManager{parentManager}
}

// DeployManager wraps DeployManager and allows typed contract handles to be returned.
type DeployManager struct {
	*manager.DeployerManager
}
