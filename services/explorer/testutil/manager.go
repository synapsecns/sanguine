package testutil

import (
	"testing"

	"github.com/synapsecns/sanguine/ethergo/manager"
)

// NewDeployManager creates a deploy manager.
func NewDeployManager(t *testing.T) *DeployManager {
	t.Helper()
	parentManager := manager.NewDeployerManager(t,
		NewBridgeConfigV3Deployer,
		NewSynapseBridgeDeployer,
		NewSwapFlashLoanDeployer,
		NewSynapseBridgeV1Deployer,
		NewMessageBusDeployer,
		NewMetaSwapDeployer,
		NewCCTPDeployer,
		NewFastBridgeDeployer,
	)

	return &DeployManager{parentManager}
}

// DeployManager wraps DeployManager and allows typed contract handles to be returned.
type DeployManager struct {
	*manager.DeployerManager
}
