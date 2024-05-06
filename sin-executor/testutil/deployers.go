package testutil

import (
	"context"
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/ethergo/contracts"
	"github.com/synapsecns/sanguine/ethergo/deployer"
	"github.com/synapsecns/sanguine/ethergo/manager"
	"github.com/synapsecns/sanguine/sin-executor/contracts/executionservice"
	"github.com/synapsecns/sanguine/sin-executor/contracts/interchainclient"
	"github.com/synapsecns/sanguine/sin-executor/contracts/interchaindb"
	"github.com/synapsecns/sanguine/sin-executor/contracts/mocks/gasoraclemock"
	"github.com/synapsecns/sanguine/sin-executor/contracts/mocks/interchainapp"
	"github.com/synapsecns/sanguine/sin-executor/contracts/mocks/interchainmodulemock"
	"github.com/synapsecns/sanguine/sin-executor/contracts/mocks/optionslibexport"
)

// DeployManager wraps DeployManager and allows typed contract handles to be returned.
type DeployManager struct {
	*manager.DeployerManager
}

// NewDeployManager creates a new DeployManager.
func NewDeployManager(t *testing.T) *DeployManager {
	t.Helper()

	return &DeployManager{manager.NewDeployerManager(t, interchainClientDeployer, interchainDBDeployer, interchainModuleMockDeployer, interchainAppMockDeployer, optionsLibMock, executionService, gasOracleMock)}
}

var (
	// TODO: this looks horrible, rip out the functional deployer.
	interchainClientDeployer = deployer.NewFunctionalDeployer(InterchainClient,
		func(ctx context.Context, helpers deployer.IFunctionalDeployer,
			transactOps *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, interface{}, error) {
			idb := helpers.Registry().Get(ctx, InterchainDB)

			deployAddress, deployTx, deployIface, err := interchainclient.DeployInterchainClientV1(transactOps, backend, idb.Address(), transactOps.From)
			if err != nil {
				return common.Address{}, nil, nil, fmt.Errorf("could not deploy interchain client: %w", err)
			}

			helpers.Backend().WaitForConfirmation(ctx, deployTx)

			return deployAddress, deployTx, deployIface, nil
		}, func(address common.Address, backend bind.ContractBackend) (interface{}, error) {
			return interchainclient.NewInterchainClientRef(address, backend)
		}, []contracts.ContractType{InterchainDB})

	interchainDBDeployer = deployer.NewFunctionalDeployer(InterchainDB,
		func(ctx context.Context, helpers deployer.IFunctionalDeployer, transactOps *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, interface{}, error) {
			deployAddress, deployTx, deployIface, err := interchaindb.DeployInterchainDB(transactOps, backend)
			if err != nil {
				return common.Address{}, nil, nil, fmt.Errorf("could not deploy interchain db: %w", err)
			}
			helpers.Backend().WaitForConfirmation(ctx, deployTx)

			return deployAddress, deployTx, deployIface, nil
		}, func(address common.Address, backend bind.ContractBackend) (interface{}, error) {
			return interchaindb.NewInterchainDBRef(address, backend)
		}, []contracts.ContractType{})

	interchainModuleMockDeployer = deployer.NewFunctionalDeployer(InterchainModuleMock, func(_ context.Context, _ deployer.IFunctionalDeployer, transactOps *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, interface{}, error) {
		return interchainmodulemock.DeployInterchainModuleMock(transactOps, backend)
	}, func(address common.Address, backend bind.ContractBackend) (interface{}, error) {
		return interchainmodulemock.NewInterchainModuleMockRef(address, backend)
	}, []contracts.ContractType{})

	interchainAppMockDeployer = deployer.NewFunctionalDeployer(InterchainApp, func(ctx context.Context, helpers deployer.IFunctionalDeployer, transactOps *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, interface{}, error) {
		clientContract := helpers.Registry().Get(ctx, InterchainClient)
		sendingModule := helpers.Registry().Get(ctx, InterchainModuleMock)

		appAddress, appTx, appMock, err := interchainapp.DeployExampleAppV1(transactOps, backend, transactOps.From)
		if err != nil {
			return common.Address{}, nil, nil, fmt.Errorf("could not deploy interchain app mock: %w", err)
		}
		helpers.Backend().WaitForConfirmation(ctx, appTx)

		tx, err := appMock.AddInterchainClient(transactOps, clientContract.Address(), true)
		if err != nil {
			return common.Address{}, nil, nil, fmt.Errorf("could not set interchain client: %w", err)
		}
		helpers.Backend().WaitForConfirmation(ctx, tx)

		tx, err = appMock.AddTrustedModule(transactOps, sendingModule.Address())
		if err != nil {
			return common.Address{}, nil, nil, fmt.Errorf("could not set trusted module: %w", err)
		}
		helpers.Backend().WaitForConfirmation(ctx, tx)

		return appAddress, appTx, appMock, nil
	}, func(address common.Address, backend bind.ContractBackend) (interface{}, error) {
		return interchainapp.NewInterchainAppRef(address, backend)
	}, []contracts.ContractType{InterchainClient, InterchainModuleMock})

	optionsLibMock = deployer.NewFunctionalDeployer(OptionsLib, func(_ context.Context, _ deployer.IFunctionalDeployer, transactOps *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, interface{}, error) {
		return optionslibexport.DeployOptionsLibMocks(transactOps, backend)
	}, func(address common.Address, backend bind.ContractBackend) (interface{}, error) {
		return optionslibexport.NewOptionsLibExportRef(address, backend)
	}, []contracts.ContractType{})

	executionService = deployer.NewFunctionalDeployer(ExecutionService, func(ctx context.Context, helpers deployer.IFunctionalDeployer, transactOps *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, interface{}, error) {
		address, deployTx, iface, err := executionservice.DeploySynapseExecutionServiceV1Harness(transactOps, backend)
		if err != nil {
			return common.Address{}, nil, nil, fmt.Errorf("could not deploy execution service: %w", err)
		}
		helpers.Backend().WaitForConfirmation(ctx, deployTx)

		tx, err := iface.SetGasOracle(transactOps, helpers.Registry().Get(ctx, GasOracleMock).Address())
		if err != nil {
			return common.Address{}, nil, nil, fmt.Errorf("could not set gas oracle: %w", err)
		}
		helpers.Backend().WaitForConfirmation(ctx, tx)

		icClientRole, err := iface.ICCLIENTROLE(&bind.CallOpts{Context: ctx})
		if err != nil {
			return common.Address{}, nil, nil, fmt.Errorf("could not get ic client role: %w", err)
		}

		tx, err = iface.GrantRole(transactOps, icClientRole, helpers.Registry().Get(ctx, InterchainClient).Address())
		if err != nil {
			return common.Address{}, nil, nil, fmt.Errorf("could not set interchain client: %w", err)
		}
		helpers.Backend().WaitForConfirmation(ctx, tx)

		return address, tx, iface, nil
	}, func(address common.Address, backend bind.ContractBackend) (interface{}, error) {
		return executionservice.NewSynapseExecutionServiceV1HarnessRef(address, backend)
	}, []contracts.ContractType{GasOracleMock, InterchainClient})

	gasOracleMock = deployer.NewFunctionalDeployer(GasOracleMock, func(_ context.Context, _ deployer.IFunctionalDeployer, transactOps *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, interface{}, error) {
		return gasoraclemock.DeploySynapseGasOracleMock(transactOps, backend)
	}, func(address common.Address, backend bind.ContractBackend) (interface{}, error) {
		return gasoraclemock.NewGasOracleMockRef(address, backend)
	}, []contracts.ContractType{})
)
