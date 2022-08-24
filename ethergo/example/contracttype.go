package example

import (
	"github.com/ethereum/go-ethereum/common/compiler"
	"github.com/synapsecns/sanguine/ethergo/contracts"
	"github.com/synapsecns/sanguine/ethergo/example/counter"
)

// set all contact types.
func init() {
	for i := 0; i < len(_contractTypeImpl_index); i++ {
		contractType := contractTypeImpl(i)
		AllContractTypes = append(AllContractTypes, contractType)
		// assert type is correct
		var _ contracts.ContractType = contractType
	}
}

// AllContractTypes is a list of all contract types. Since we use stringer and this is a testing library, instead
// of manually copying all these out we pull the names out of stringer. In order to make sure stringer is updated, we panic on
// any method called where the index is higher than the stringer array length.
var AllContractTypes []contractTypeImpl

// verifyStringerUpdated verifies stringer is up to date (this index is included in stringer).
func verifyStringerUpdated(contractType contractTypeImpl) {
	if int(contractType) > len(_contractTypeImpl_index) {
		panic("please update stringer before running test again")
	}
}

// contractTypeImpl is the type of the contract being saved/fetched.
// we use an interface here so the deploy helper here can be abstracted away from the synapse contracts
//
//go:generate go run golang.org/x/tools/cmd/stringer -type=contractTypeImpl -linecomment
type contractTypeImpl int

const (
	// CounterType is the type of the counter contract.
	CounterType contractTypeImpl = 0 // CounterType
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

// ContractInfo gets the source code of every contract. See TODO above.
// TODO these should use contract name and maybe come out of the generator.
func (c contractTypeImpl) ContractInfo() *compiler.Contract {
	switch c {
	case CounterType:
		return counter.Contracts["/solidity/counter.sol:Counter"]
	default:
		panic("not yet implemented")
	}
}

// ContractName gets the name of the deployed contract.
func (c contractTypeImpl) ContractName() string {
	return c.Name()
}

// make sure contractTypeImpl conforms to deployer.ContractType.
var _ contracts.ContractType = contractTypeImpl(1)
