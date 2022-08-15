package testutil

import (
	"github.com/ethereum/go-ethereum/common/compiler"
	"github.com/synapsecns/sanguine/core/contracts/attestationcollector"
	"github.com/synapsecns/sanguine/core/contracts/notarymanager"
	"github.com/synapsecns/sanguine/core/contracts/origin"
	"github.com/synapsecns/sanguine/core/contracts/replicamanager"
	"github.com/synapsecns/sanguine/core/contracts/test/attestationharness"
	"github.com/synapsecns/sanguine/core/contracts/test/headerharness"
	"github.com/synapsecns/sanguine/core/contracts/test/messageharness"
	"github.com/synapsecns/sanguine/core/contracts/test/originharness"
	"github.com/synapsecns/sanguine/core/contracts/test/replicamanagerharness"
	"github.com/synapsecns/sanguine/core/contracts/test/tipsharness"
	"github.com/synapsecns/sanguine/ethergo/deployer"
)

// set all contact types.
func init() {
	for i := 0; i < len(_contractTypeImpl_index); i++ {
		contractType := contractTypeImpl(i)
		AllContractTypes = append(AllContractTypes, contractType)
		// assert type is correct
		var _ deployer.ContractType = contractType
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
//go:generate go run golang.org/x/tools/cmd/stringer -type=contractTypeImpl -linecomment
type contractTypeImpl int

const (
	// OriginType is the type of the origin.
	OriginType contractTypeImpl = 0 // Origin
	// MessageHarnessType is the type of the message harness contract.
	MessageHarnessType contractTypeImpl = iota // MessageHarness
	// OriginHarnessType is the origin harness type.
	OriginHarnessType contractTypeImpl = iota // OriginHarness
	// AttestationHarnessType is the attestation harness type.
	AttestationHarnessType contractTypeImpl = iota
	// TipsHarnessType is the type of the tips harness.
	TipsHarnessType contractTypeImpl = iota
	// HeaderHarnessType is the tyoe of the header harness.
	HeaderHarnessType contractTypeImpl = iota
	// ReplicaManagerHarnessType is the replica manager harness type.
	ReplicaManagerHarnessType contractTypeImpl = iota
	// NotaryManagerType is the type of the update manager.
	NotaryManagerType contractTypeImpl = iota // UpdaterManager
	// AttestationCollectorType is the type of the attestation collector.
	AttestationCollectorType contractTypeImpl = iota // AttestationCollector
	// ReplicaManagerType is the type of the replica manager.
	ReplicaManagerType contractTypeImpl = iota // ReplicaManager
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
//nolint: cyclop
func (c contractTypeImpl) ContractInfo() *compiler.Contract {
	switch c {
	case OriginType:
		return origin.Contracts["solidity/Origin.sol:Origin"]
	case MessageHarnessType:
		return messageharness.Contracts["solidity/MessageHarness.sol:MessageHarness"]
	case OriginHarnessType:
		return originharness.Contracts["solidity/OriginHarness.sol:OriginHarness"]
	case AttestationHarnessType:
		return attestationharness.Contracts["solidity/AttestationHarness.sol:AttestationHarness"]
	case ReplicaManagerHarnessType:
		return replicamanagerharness.Contracts["solidity/ReplicaManagerHarness.sol:ReplicaManagerHarness"]
	case NotaryManagerType:
		return notarymanager.Contracts["solidity/UpdaterManager.sol:UpdaterManager"]
	case TipsHarnessType:
		return tipsharness.Contracts["solidity/TipsHarness.sol:TipsHarness"]
	case AttestationCollectorType:
		return attestationcollector.Contracts["solidity/AttestationCollector.sol:AttestationCollector"]
	case ReplicaManagerType:
		return replicamanager.Contracts["solidity/ReplicaManager.sol:ReplicaManager"]
	case HeaderHarnessType:
		return headerharness.Contracts["solidity/HeaderHarness.sol:HeaderHarness"]
	default:
		panic("not yet implemented")
	}
}

// ContractName gets the name of the deployed contract.
func (c contractTypeImpl) ContractName() string {
	return c.Name()
}

// make sure contractTypeImpl conforms to deployer.ContractType.
var _ deployer.ContractType = contractTypeImpl(1)
