package manager_test

import (
	"context"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/compiler"
	"github.com/ethereum/go-ethereum/core/types"
	backendMocks "github.com/synapsecns/sanguine/ethergo/backends/mocks"
	contractMocks "github.com/synapsecns/sanguine/ethergo/contracts/mocks"
	deployerMocks "github.com/synapsecns/sanguine/ethergo/deployer/mocks"
	"math/big"
	"reflect"
	"testing"

	"github.com/stretchr/testify/mock"

	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/ethergo/contracts"
	"github.com/synapsecns/sanguine/ethergo/manager"
)

var mockCtx = mock.AnythingOfType(reflect.TypeOf(context.Background()).String())

func TestDeployerManager_Get(t *testing.T) {
	backend := &backendMocks.SimulatedTestBackend{}
	backend.On("GetBigChainID").Return(big.NewInt(1))

	mgr := manager.NewDeployerManager(t)

	contractRegistry := new(deployerMocks.ContractRegistry)

	contractType1 := makeContractType()
	contractType2 := makeContractType()

	deployedContract1 := makeDeployedContract()
	deployedContract2 := makeDeployedContract()

	contractRegistry.On("Get", mockCtx, contractType1).Return(deployedContract1, nil)
	contractRegistry.On("Get", mockCtx, contractType2).Return(deployedContract2, nil)

	mgr.SetContractRegistry(backend, contractRegistry)

	dc2 := mgr.Get(context.Background(), backend, contractType2)
	Equal(t, dc2, deployedContract2)

	dc1 := mgr.Get(context.Background(), backend, contractType1)
	Equal(t, dc1, deployedContract1)
}

func TestDeployManager(t *testing.T) {
	mgr := manager.NewDeployerManager(t)
	Equal(t, mgr.T(), t)

	Panics(t, func() {
		mgr.SetT(nil)
	})
}

func makeContractType() contracts.ContractType {
	contractType := new(contractMocks.ContractType)
	contractType.On("Name").Return(gofakeit.Word())
	contractType.On("ID").Return(int(gofakeit.Int64()))
	contractType.On("ContractInfo").Return(&compiler.Contract{})
	contractType.On("ContractName").Return(gofakeit.Word())
	return contractType
}

func makeDeployedContract() contracts.DeployedContract {
	owner := common.BigToAddress(new(big.Int).SetUint64(gofakeit.Uint64()))

	deployedContract := new(contractMocks.DeployedContract)
	deployedContract.On("Address").Return(common.BigToAddress(new(big.Int).SetUint64(gofakeit.Uint64())))
	deployedContract.On("ContractHandle").Return("")
	deployedContract.On("Owner").Return(owner)
	address := common.BigToAddress(new(big.Int).SetUint64(gofakeit.Uint64()))
	deployedContract.On("DeployTx").Return(types.NewTx(&types.LegacyTx{
		Nonce:    gofakeit.Uint64(),
		GasPrice: big.NewInt(gofakeit.Int64()),
		Gas:      gofakeit.Uint64(),
		To:       &address,
		Value:    big.NewInt(gofakeit.Int64()),
		Data:     []byte(gofakeit.Word()),
		V:        new(big.Int).SetUint64(gofakeit.Uint64()),
		R:        new(big.Int).SetUint64(gofakeit.Uint64()),
		S:        new(big.Int).SetUint64(gofakeit.Uint64()),
	}))
	deployedContract.On("ChainID").Return(big.NewInt(gofakeit.Int64()))
	deployedContract.On("OwnerPtr").Return(owner)
	return deployedContract
}

func TestGetContractRegistry(t *testing.T) {
	t.Skip("TODO")
}

func TestGetDeployedContracts(t *testing.T) {
	t.Skip("TODO")
}

func TestAssertDependniesCorrect(t *testing.T) {
	t.Skip("TODO")
}
