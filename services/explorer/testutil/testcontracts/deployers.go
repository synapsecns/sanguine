package testcontracts

import (
	"context"
	"fmt"
	"math/big"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/ethergo/backends"
	"github.com/synapsecns/sanguine/ethergo/contracts"
	"github.com/synapsecns/sanguine/ethergo/deployer"
	"github.com/synapsecns/sanguine/services/explorer/contracts/bridge/testbridge"
	"github.com/synapsecns/sanguine/services/explorer/contracts/bridge/testbridgev1"
	"github.com/synapsecns/sanguine/services/explorer/contracts/bridgeconfig"
	"github.com/synapsecns/sanguine/services/explorer/contracts/cctp/testcctp"
	"github.com/synapsecns/sanguine/services/explorer/contracts/messagebus/testmessagebus"
	"github.com/synapsecns/sanguine/services/explorer/contracts/metaswap/testmetaswap"
	"github.com/synapsecns/sanguine/services/explorer/contracts/swap/testswap"
	"github.com/synapsecns/sanguine/services/explorer/testutil"
)

// TestSynapseBridgeDeployer is the type of the test bridge deployer.
type TestSynapseBridgeDeployer struct {
	*deployer.BaseDeployer
}

// TestSynapseBridgeV1Deployer is the type of the test bridge deployer.
type TestSynapseBridgeV1Deployer struct {
	*deployer.BaseDeployer
}

// TestSwapFlashLoanDeployer is the type of the test swap deployer.
type TestSwapFlashLoanDeployer struct {
	*deployer.BaseDeployer
}

// BridgeConfigV3Deployer is the type of the bridge config v3 deployer.
type BridgeConfigV3Deployer struct {
	*deployer.BaseDeployer
}

// TestMessageBusUpgradeableDeployer is the type of the test message bus upgradeable deployer.
type TestMessageBusUpgradeableDeployer struct {
	*deployer.BaseDeployer
}

// TestMetaSwapDeployer is the type of the test meta swap deployer.
type TestMetaSwapDeployer struct {
	*deployer.BaseDeployer
}

// TestCCTPDeployer is the type of the test cctp deployer.
type TestCCTPDeployer struct {
	*deployer.BaseDeployer
}

// NewTestSynapseBridgeDeployer creates a new test bridge deployer.
func NewTestSynapseBridgeDeployer(registry deployer.GetOnlyContractRegistry, backend backends.SimulatedTestBackend) deployer.ContractDeployer {
	return TestSynapseBridgeDeployer{deployer.NewSimpleDeployer(registry, backend, TestSynapseBridgeType)}
}

// NewTestSynapseBridgeV1Deployer creates a new test bridge v1 deployer.
func NewTestSynapseBridgeV1Deployer(registry deployer.GetOnlyContractRegistry, backend backends.SimulatedTestBackend) deployer.ContractDeployer {
	return TestSynapseBridgeV1Deployer{deployer.NewSimpleDeployer(registry, backend, TestSynapseBridgeV1Type)}
}

// NewTestSwapFlashLoanDeployer creates a new test swap deployer.
func NewTestSwapFlashLoanDeployer(registry deployer.GetOnlyContractRegistry, backend backends.SimulatedTestBackend) deployer.ContractDeployer {
	return TestSwapFlashLoanDeployer{deployer.NewSimpleDeployer(registry, backend, TestSwapFlashLoanType)}
}

// NewBridgeConfigV3Deployer creates a new bridge config v2 client.
func NewBridgeConfigV3Deployer(registry deployer.GetOnlyContractRegistry, backend backends.SimulatedTestBackend) deployer.ContractDeployer {
	return BridgeConfigV3Deployer{deployer.NewSimpleDeployer(registry, backend, testutil.BridgeConfigTypeV3)}
}

// NewTestMessageBusDeployer creates a new test message bus deployer.
func NewTestMessageBusDeployer(registry deployer.GetOnlyContractRegistry, backend backends.SimulatedTestBackend) deployer.ContractDeployer {
	return TestMessageBusUpgradeableDeployer{deployer.NewSimpleDeployer(registry, backend, TestMessageBusType)}
}

// NewTestMetaSwapDeployer creates a new test meta swap deployer.
func NewTestMetaSwapDeployer(registry deployer.GetOnlyContractRegistry, backend backends.SimulatedTestBackend) deployer.ContractDeployer {
	return TestMetaSwapDeployer{deployer.NewSimpleDeployer(registry, backend, TestMetaSwapType)}
}

// NewTestCCTPDeployer creates a new test cctp client.
func NewTestCCTPDeployer(registry deployer.GetOnlyContractRegistry, backend backends.SimulatedTestBackend) deployer.ContractDeployer {
	return TestCCTPDeployer{deployer.NewSimpleDeployer(registry, backend, TestCCTPType)}
}

// Deploy deploys a test bridge.
func (t TestSynapseBridgeDeployer) Deploy(ctx context.Context) (contracts.DeployedContract, error) {
	return t.DeploySimpleContract(ctx, func(transactOps *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, interface{}, error) {
		return testbridge.DeployTestSynapseBridge(transactOps, backend)
	}, func(address common.Address, backend bind.ContractBackend) (interface{}, error) {
		return testbridge.NewTestBridgeRef(address, backend)
	})
}

// Deploy deploys a test bridge v1.
func (t TestSynapseBridgeV1Deployer) Deploy(ctx context.Context) (contracts.DeployedContract, error) {
	return t.DeploySimpleContract(ctx, func(transactOps *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, interface{}, error) {
		return testbridgev1.DeployTestSynapseBridgeV1(transactOps, backend)
	}, func(address common.Address, backend bind.ContractBackend) (interface{}, error) {
		return testbridgev1.NewTestBridgeV1Ref(address, backend)
	})
}

// Deploy deploys a test swap.
func (t TestSwapFlashLoanDeployer) Deploy(ctx context.Context) (contracts.DeployedContract, error) {
	return t.DeploySimpleContract(ctx, func(transactOps *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, interface{}, error) {
		return testswap.DeployTestSwapFlashLoan(transactOps, backend)
	}, func(address common.Address, backend bind.ContractBackend) (interface{}, error) {
		return testswap.NewTestSwapRef(address, backend)
	})
}

// Deploy deploys bridge config v3
//
//nolint:dupl
func (n BridgeConfigV3Deployer) Deploy(ctx context.Context) (contracts.DeployedContract, error) {
	return n.DeploySimpleContract(ctx, func(transactOps *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, interface{}, error) {
		deployAddress, tx, handler, err := bridgeconfig.DeployBridgeConfigV3(transactOps, backend)
		if err != nil {
			return common.Address{}, nil, nil, fmt.Errorf("could not deploy bridge config: %w", err)
		}

		// wait for confirm, we need this to grant role
		n.Backend().WaitForConfirmation(ctx, tx)

		// https://github.com/synapsecns/synapse-contracts/pull/13 introduces a breaking change where the BRIDGEMANAGER_ROLE is not automatically granted to the
		// deployer. We fix that here by granting the role to the owner
		bridgeManagerRole, err := handler.BRIDGEMANAGERROLE(&bind.CallOpts{Context: ctx})
		if err != nil {
			return common.Address{}, nil, nil, fmt.Errorf("could not get bridge manager role: %w", err)
		}

		auth := n.Backend().GetTxContext(ctx, &transactOps.From)
		// grant the role
		grantTx, err := handler.GrantRole(auth.TransactOpts, bridgeManagerRole, auth.From)
		if err != nil {
			return common.Address{}, nil, nil, fmt.Errorf("could not grant bridge manager role: %w", err)
		}

		n.Backend().WaitForConfirmation(ctx, grantTx)

		return deployAddress, tx, handler, nil
	}, func(address common.Address, backend bind.ContractBackend) (interface{}, error) {
		return bridgeconfig.NewBridgeConfigRef(address, backend)
	})
}

// Deploy deploys a test message bus contract.
func (t TestMessageBusUpgradeableDeployer) Deploy(ctx context.Context) (contracts.DeployedContract, error) {
	return t.DeploySimpleContract(ctx, func(transactOps *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, interface{}, error) {
		return testmessagebus.DeployTestMessageBusUpgradeable(transactOps, backend)
	}, func(address common.Address, backend bind.ContractBackend) (interface{}, error) {
		return testmessagebus.NewTestMessageBusRef(address, backend)
	})
}

// Deploy deploys a test meta swap contract.
func (t TestMetaSwapDeployer) Deploy(ctx context.Context) (contracts.DeployedContract, error) {
	return t.DeploySimpleContract(ctx, func(transactOps *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, interface{}, error) {
		return testmetaswap.DeployTestMetaSwap(transactOps, backend)
	}, func(address common.Address, backend bind.ContractBackend) (interface{}, error) {
		return testmetaswap.NewTestMetaSwapRef(address, backend)
	})
}

// Deploy deploys CCTP contract
//
//nolint:dupl
func (n TestCCTPDeployer) Deploy(ctx context.Context) (contracts.DeployedContract, error) {
	tokenMessengerContract, err := n.DeploySimpleContract(ctx, func(transactOps *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, interface{}, error) {
		return testcctp.DeployMessageTransmitter(transactOps, backend)
	}, func(address common.Address, backend bind.ContractBackend) (interface{}, error) {
		return testcctp.NewMessageTransmitter(address, backend)
	})
	if err != nil {
		return nil, fmt.Errorf("failed to deploy tokenMessengerContract %w", err)
	}
	return n.DeploySimpleContract(ctx, func(transactOps *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, interface{}, error) {
		// Create mock owner
		owner := common.BigToAddress(big.NewInt(gofakeit.Int64()))

		return testcctp.DeployTestSynapseCCTP(transactOps, backend, tokenMessengerContract.Address(), owner)
	}, func(address common.Address, backend bind.ContractBackend) (interface{}, error) {
		return testcctp.NewTestCCTPRef(address, backend)
	})
}

var _ deployer.ContractDeployer = &TestSynapseBridgeDeployer{}
var _ deployer.ContractDeployer = &TestSynapseBridgeV1Deployer{}
var _ deployer.ContractDeployer = &TestSwapFlashLoanDeployer{}
var _ deployer.ContractDeployer = &BridgeConfigV3Deployer{}
var _ deployer.ContractDeployer = &TestMessageBusUpgradeableDeployer{}
var _ deployer.ContractDeployer = &TestMetaSwapDeployer{}
var _ deployer.ContractDeployer = &TestCCTPDeployer{}
