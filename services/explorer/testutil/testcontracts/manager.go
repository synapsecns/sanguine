package testcontracts

import (
	"github.com/synapsecns/sanguine/ethergo/manager"
	"testing"
)

// NewDeployManager creates a deploy manager.
func NewDeployManager(t *testing.T) *DeployManager {
	t.Helper()

	parentManager := manager.NewDeployerManager(t,
		NewTestSwapFlashLoanDeployer,
		NewBridgeConfigV3Deployer,
		NewTestSynapseBridgeDeployer,
		NewTestSynapseBridgeV1Deployer,
		NewTestMessageBusDeployer,
		NewTestMetaSwapDeployer,
		NewTestCCTPDeployer,
	)
	return &DeployManager{parentManager}
}

// DeployManager wraps DeployManager and allows typed contract handles to be returned.
type DeployManager struct {
	*manager.DeployerManager
}
