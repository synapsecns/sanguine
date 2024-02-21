package testutil

import (
	"github.com/ethereum/go-ethereum/common/compiler"
	"github.com/synapsecns/sanguine/ethergo/backends/base"
	"github.com/synapsecns/sanguine/ethergo/contracts"
	"github.com/synapsecns/sanguine/sin-executor/contracts/interchainclient"
	"github.com/synapsecns/sanguine/sin-executor/contracts/interchaindb"
	"github.com/synapsecns/sanguine/sin-executor/contracts/mocks/executionfeesmock"
	"github.com/synapsecns/sanguine/sin-executor/contracts/mocks/executionservicemock"
	"github.com/synapsecns/sanguine/sin-executor/contracts/mocks/interchainappmock"
	"github.com/synapsecns/sanguine/sin-executor/contracts/mocks/interchainmodulemock"
	"github.com/synapsecns/sanguine/sin-executor/contracts/mocks/optionslibexport"
)

// set all contact types.
func init() {
	for i := 0; i < len(_contractTypeImpl_index)-1; i++ {
		contractType := contractTypeImpl(i + 1)
		AllContractTypes = append(AllContractTypes, contractType)
		// assert type is correct
		var _ contracts.ContractType = contractType
		// boot time assertion
		if !base.IsVerificationBlacklisted(contractType) {
			if contractType.ContractInfo() == nil {
				panic("contract info is nil")
			}
		}
	}
}

// verifyStringerUpdated verifies stringer is up to date (this index is included in stringer).
func verifyStringerUpdated(contractType contractTypeImpl) {
	if int(contractType) > len(_contractTypeImpl_index) {
		panic("please update stringer before running test again")
	}
}

// AllContractTypes is a list of all contract types. Since we use stringer and this is a testing library, instead
// of manually copying all these out we pull the names out of stringer. In order to make sure stringer is updated, we panic on
// any method called where the index is higher than the stringer array length.
var AllContractTypes []contractTypeImpl

// contractTypeImpl is the type of the contract being saved/fetched.
// we use an interface here so the deploy helper here can be abstracted away from the synapse contracts
//
//go:generate go run golang.org/x/tools/cmd/stringer -type=contractTypeImpl -linecomment
type contractTypeImpl int

const (
	// InterchainClient is the interchain execution client.
	InterchainClient contractTypeImpl = iota + 1 // SynapseModule
	// InterchainDB is the interchain database.
	InterchainDB // SynapseModule
	// InterchainModuleMock is the interchain module mock.
	InterchainModuleMock // InterchainModuleMock
	// InterchainAppMock is the interchain app mock.
	InterchainAppMock // InterchainAppMock
	// OptionsLib is the options library.
	OptionsLib // OptionsLib
	// ExecutionServiceMock is the execution service mock.
	ExecutionServiceMock // ExecutionServiceMock
	// ExecutionFeesMock is the execution fees mock.
	ExecutionFeesMock // ExecutionFeesMock
)

// ID gets the contract type as an id.
func (c contractTypeImpl) ID() int {
	verifyStringerUpdated(c)

	return int(c)
}

// Name gets the name of the contract.
func (c contractTypeImpl) Name() string {
	verifyStringerUpdated(c)

	return c.String()
}

func (c contractTypeImpl) ContractName() string {
	return c.Name()
}

// ContractInfo gets the source code of every contract. See TODO above.
// TODO these should use contract name and maybe come out of the generator.
//
//nolint:cyclop
func (c contractTypeImpl) ContractInfo() *compiler.Contract {
	switch c {
	case InterchainClient:
		return interchainclient.Contracts["solidity/InterchainClientV1.sol:InterchainClientV1"]
	case InterchainDB:
		return interchaindb.Contracts["solidity/InterchainDB.sol:InterchainDB"]
	case InterchainModuleMock:
		return interchainmodulemock.Contracts["solidity/InterchainModuleMock.sol:InterchainModuleMock"]
	case InterchainAppMock:
		return interchainappmock.Contracts["solidity/InterchainAppMock.sol:InterchainAppMock"]
	case OptionsLib:
		return optionslibexport.Contracts["solidity/OptionsLibExport.sol:OptionsLibMocks"]
	case ExecutionServiceMock:
		return executionservicemock.Contracts["solidity/ExecutionServiceMock.sol:ExecutionServiceMock"]
	case ExecutionFeesMock:
		return executionfeesmock.Contracts["solidity/ExecutionfeesMock.sol:ExecutionFeesMock"]
	}
	return nil
}
