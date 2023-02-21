package testutil

import (
	"github.com/ethereum/go-ethereum/common/compiler"
	"github.com/synapsecns/sanguine/agents/contracts/attestationcollector"
	"github.com/synapsecns/sanguine/agents/contracts/destination"
	"github.com/synapsecns/sanguine/agents/contracts/origin"
	"github.com/synapsecns/sanguine/agents/contracts/test/attestationharness"
	"github.com/synapsecns/sanguine/agents/contracts/test/destinationharness"
	"github.com/synapsecns/sanguine/agents/contracts/test/headerharness"
	"github.com/synapsecns/sanguine/agents/contracts/test/messageharness"
	"github.com/synapsecns/sanguine/agents/contracts/test/originharness"
	"github.com/synapsecns/sanguine/agents/contracts/test/testclient"
	"github.com/synapsecns/sanguine/agents/contracts/test/tipsharness"
	"github.com/synapsecns/sanguine/agents/testutil/agentstestcontract"
	"github.com/synapsecns/sanguine/ethergo/contracts"
)

// set all contact types.
func init() {
	for i := 0; i < len(_contractTypeImpl_index)-1; i++ {
		contractType := contractTypeImpl(i)
		AllContractTypes = append(AllContractTypes, contractType)
		// assert type is correct
		var _ contracts.ContractType = contractType
		// boot time assertion
		if contractType.ContractInfo() == nil {
			panic("contract info is nil")
		}
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
	// OriginType is the type of the origin.
	OriginType contractTypeImpl = iota // Origin
	// MessageHarnessType is the type of the message harness contract.
	MessageHarnessType // MessageHarness
	// OriginHarnessType is the origin harness type.
	OriginHarnessType // OriginHarness
	// AttestationHarnessType is the attestation harness type.
	AttestationHarnessType
	// TipsHarnessType is the type of the tips harness.
	TipsHarnessType
	// HeaderHarnessType is the type of the header harness.
	HeaderHarnessType
	// DestinationHarnessType is the destination harness type.
	DestinationHarnessType // DestinationHarness
	// AttestationCollectorType is the type of the attestation collector.
	AttestationCollectorType // AttestationCollector
	// DestinationType is the type of the destination.
	DestinationType // Destination
	// AgentsTestContractType is the type of the agents test contract.
	AgentsTestContractType // AgentsTestContract
	// TestClientType is the type of the test client.
	TestClientType // TestClient
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
//
//nolint:cyclop
func (c contractTypeImpl) ContractInfo() *compiler.Contract {
	switch c {
	case OriginType:
		return origin.Contracts["solidity/Origin.sol:Origin"]
	case MessageHarnessType:
		return messageharness.Contracts["solidity/MessageHarness.t.sol:MessageHarness"]
	case OriginHarnessType:
		return originharness.Contracts["solidity/OriginHarness.t.sol:OriginHarness"]
	case AttestationHarnessType:
		return attestationharness.Contracts["solidity/AttestationHarness.t.sol:AttestationHarness"]
	case DestinationHarnessType:
		return destinationharness.Contracts["solidity/DestinationHarness.t.sol:DestinationHarness"]
	case TipsHarnessType:
		return tipsharness.Contracts["solidity/TipsHarness.t.sol:TipsHarness"]
	case AttestationCollectorType:
		return attestationcollector.Contracts["solidity/AttestationCollector.sol:AttestationCollector"]
	case DestinationType:
		return destination.Contracts["solidity/Destination.sol:Destination"]
	case HeaderHarnessType:
		return headerharness.Contracts["solidity/HeaderHarness.t.sol:HeaderHarness"]
	case AgentsTestContractType:
		return agentstestcontract.Contracts["solidity/AgentsTestContract.sol:AgentsTestContract"]
	case TestClientType:
		return testclient.Contracts["solidity/TestClient.sol:TestClient"]
	default:
		panic("not yet implemented")
	}
}

// ContractName gets the name of the deployed contract.
func (c contractTypeImpl) ContractName() string {
	return c.Name()
}

// make sure contractTypeImpl conforms to contracts.ContractType.
var _ contracts.ContractType = contractTypeImpl(1)
