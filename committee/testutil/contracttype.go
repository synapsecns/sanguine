package testutil

import (
	"github.com/ethereum/go-ethereum/common/compiler"
	"github.com/synapsecns/sanguine/committee/contracts/interchaindb"
	"github.com/synapsecns/sanguine/committee/contracts/mocks/gasoraclemock"
	"github.com/synapsecns/sanguine/committee/contracts/synapsemodule"
	"github.com/synapsecns/sanguine/ethergo/backends/base"
	"github.com/synapsecns/sanguine/ethergo/contracts"
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
	// SynapseModuleType is the type of the contract being fetched.
	SynapseModuleType contractTypeImpl = iota + 1 // SynapseModule
	// InterchainDB is the type of the contract being fetched.
	InterchainDB // InterchainDB
	// GasOracleMockType is the type of the contract being fetched.
	GasOracleMockType // GasOracleMock
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
	case SynapseModuleType:
		return synapsemodule.Contracts["solidity/SynapseModule.sol:SynapseModule"]
	case InterchainDB:
		return interchaindb.Contracts["solidity/InterchainDB.sol:InterchainDB"]
	case GasOracleMockType:
		return gasoraclemock.Contracts["solidity/SynapseGasOracleMock.sol:SynapseGasOracleMock"]
	}
	return nil
}
