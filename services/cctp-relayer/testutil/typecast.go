package testutil

import (
	"context"
	"github.com/synapsecns/sanguine/ethergo/backends"
	"github.com/synapsecns/sanguine/ethergo/contracts"
	"github.com/synapsecns/sanguine/ethergo/manager"
	"github.com/synapsecns/sanguine/services/cctp-relayer/contracts/cctp"
	"github.com/synapsecns/sanguine/services/cctp-relayer/contracts/mockmessagetransmitter"
	"github.com/synapsecns/sanguine/services/cctp-relayer/contracts/mockmintburntoken"
	"github.com/synapsecns/sanguine/services/cctp-relayer/contracts/mocktokenmessenger"
	"github.com/synapsecns/sanguine/services/cctp-relayer/contracts/mocktokenminter"
)

// GetSynapseCCTP gets a typecast synapsecctp contract.
func (d *DeployManager) GetSynapseCCTP(ctx context.Context, backend backends.SimulatedTestBackend) (contract contracts.DeployedContract, handle *cctp.SynapseCCTPRef) {
	d.T().Helper()

	return manager.GetContract[*cctp.SynapseCCTPRef](ctx, d.T(), d, backend, SynapseCCTPType)
}

// GetMockMessageTransmitterType gets a mock token message transmitter contract.
func (d *DeployManager) GetMockMessageTransmitterType(ctx context.Context, backend backends.SimulatedTestBackend) (contract contracts.DeployedContract, handle *mockmessagetransmitter.MockMessageTransmitterRef) {
	d.T().Helper()

	return manager.GetContract[*mockmessagetransmitter.MockMessageTransmitterRef](ctx, d.T(), d, backend, MockMessageTransmitterType)
}

// GetMockTokenMessengerType gets a mock token transmitter contract.
func (d *DeployManager) GetMockTokenMessengerType(ctx context.Context, backend backends.SimulatedTestBackend) (contract contracts.DeployedContract, handle *mocktokenmessenger.MockTokenMessengerRef) {
	d.T().Helper()

	return manager.GetContract[*mocktokenmessenger.MockTokenMessengerRef](ctx, d.T(), d, backend, MockTokenMessengerType)
}

// GetMockMintBurnTokenType gets a mock token transmitter contract.
func (d *DeployManager) GetMockMintBurnTokenType(ctx context.Context, backend backends.SimulatedTestBackend) (contract contracts.DeployedContract, handle *mockmintburntoken.MockMintBurnTokenRef) {
	d.T().Helper()

	return manager.GetContract[*mockmintburntoken.MockMintBurnTokenRef](ctx, d.T(), d, backend, MockMintBurnTokenType)
}

// GetMockTokenMinter gets a mock token minter contract.
func (d *DeployManager) GetMockTokenMinter(ctx context.Context, backend backends.SimulatedTestBackend) (contract contracts.DeployedContract, handle *mocktokenminter.MockTokenMinterRef) {
	d.T().Helper()

	return manager.GetContract[*mocktokenminter.MockTokenMinterRef](ctx, d.T(), d, backend, MockTokenMinterType)
}
