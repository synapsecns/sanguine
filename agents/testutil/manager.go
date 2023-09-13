package testutil

import (
	"testing"

	"github.com/synapsecns/sanguine/ethergo/manager"
)

// NewDeployManager creates a deploy manager.
func NewDeployManager(t *testing.T) *DeployManager {
	t.Helper()

	parentManager := manager.NewDeployerManager(t,
		NewInboxDeployer, NewBondingManagerDeployer, NewBondingManagerHarnessDeployer,
		NewLightInboxDeployer, NewLightManagerDeployer, NewLightManagerHarnessDeployer,
		NewOriginDeployer, NewMessageHarnessDeployer, NewOriginHarnessDeployer, NewSummitDeployer,
		NewTipsHarnessDeployer, NewDestinationDeployer, NewDestinationHarnessDeployer, NewSummitHarnessDeployer,
		NewHeaderHarnessDeployer, NewAgentsTestContractDeployer, NewTestClientDeployer, NewPingPongClientDeployer,
		NewStateHarnessDeployer, NewSnapshotHarnessDeployer, NewAttestationHarnessDeployer, NewGasDataHarnessDeployer,
		NewGasOracleDeployer, NewBaseMessageHarnessDeployer, NewRequestHarnessDeployer, NewReceiptHarnessDeployer,
	)
	return &DeployManager{parentManager}
}

// DeployManager wraps DeployManager and allows typed contract handles to be returned.
type DeployManager struct {
	*manager.DeployerManager
}
