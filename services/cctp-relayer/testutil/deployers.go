package testutil

// NewDeployManager creates a deploy manager.
func NewDeployManager(t *testing.T) *DeployManager {
	t.Helper()

	parentManager := manager.NewDeployerManager(t,
		NewInboxDeployer, NewBondingManagerDeployer, NewMockMessageTransmitterDeployer, NewSynapseCCTPDeployer
	)
	return &DeployManager{parentManager}
}

// DeployManager wraps DeployManager and allows typed contract handles to be returned.
type DeployManager struct {
	*manager.DeployerManager
}

// MockMessageTransmitterDeployer deploys the mockmessagetransmitter.
type MockMessageTransmitterDeployer struct {
	*deployer.BaseDeployer
}

// NewMockMessageTransmitterDeployer deploys the light inbox contract.
func NewMockMessageTransmitterDeployer(registry deployer.GetOnlyContractRegistry, backend backends.SimulatedTestBackend) deployer.ContractDeployer {
	return MockMessageTransmitterDeployer{deployer.NewSimpleDeployer(registry, backend, MockMessageTransmitterType)}
}

// Deploy deploys the light manager contract.
func (d MockMessageTransmitterDeployer) Deploy(ctx context.Context) (contracts.DeployedContract, error) {
	return d.DeploySimpleContract(ctx, func(transactOps *bind.TransactOpts, backend bind.ContractBackend) (address common.Address, tx *types.Transaction, data interface{}, err error) {
		// deploy the light inbox contract
		var rawHandle *lightinbox.LightInbox
                // define the domain as the chain id!
		return mockmessagetransmitter.DeployMockMessageTransmitter(transactOps, backend, uint32(d.Backend().GetChainID()))
	}, func(address common.Address, backend bind.ContractBackend) (interface{}, error) {
                 // remember what I said about vm.ContractRef!
		return mockmessagetransmitter.NewMockMessageTransmitterRef(address, backend)
	})
}

// MockMessageTransmitterDeployer deploys the mockmessagetransmitter.
type SynapseCCTPDeployer struct {
	*deployer.BaseDeployer
}

// NewMockMessageTransmitterDeployer deploys the light inbox contract.
func NewSynapseCCTPDeployer(registry deployer.GetOnlyContractRegistry, backend backends.SimulatedTestBackend) deployer.ContractDeployer {
	return SynapseCCTPDeployer{deployer.NewSimpleDeployer(registry, backend, MockMessageTransmitterType)}
}

// Deploy deploys the light manager contract.
func (d SynapseCCTPDeployer) Deploy(ctx context.Context) (contracts.DeployedContract, error) {
	return d.DeploySimpleContract(ctx, func(transactOps *bind.TransactOpts, backend bind.ContractBackend) (address common.Address, tx *types.Transaction, data interface{}, err error) {
                 // ah hah! the dependency
                 messageTransmitter := d.Get(ctx, MockMessageTransmitterType)
                // define the domain as the chain id!
		return cctp.DeploySynapseCCTP(transactOps, backend, messageTransmitter.Address())
	}, func(address common.Address, backend bind.ContractBackend) (interface{}, error) {
                 // remember what I said about vm.ContractRef!
		return cctp.NewSynaseCCTPRef(address, backend)
	})
}

func (d SynapseCCTPDeployer) Dependencies(){
   return []contracts.ContractType{MockMessageTransmitter}
}
