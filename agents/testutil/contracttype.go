package testutil

import (
	"github.com/ethereum/go-ethereum/common/compiler"
	"github.com/synapsecns/sanguine/agents/contracts/bondingmanager"
	"github.com/synapsecns/sanguine/agents/contracts/destination"
	"github.com/synapsecns/sanguine/agents/contracts/gasoracle"
	"github.com/synapsecns/sanguine/agents/contracts/inbox"
	"github.com/synapsecns/sanguine/agents/contracts/lightinbox"
	"github.com/synapsecns/sanguine/agents/contracts/lightmanager"
	"github.com/synapsecns/sanguine/agents/contracts/origin"
	"github.com/synapsecns/sanguine/agents/contracts/summit"
	"github.com/synapsecns/sanguine/agents/contracts/test/attestationharness"
	"github.com/synapsecns/sanguine/agents/contracts/test/basemessageharness"
	"github.com/synapsecns/sanguine/agents/contracts/test/bondingmanagerharness"
	"github.com/synapsecns/sanguine/agents/contracts/test/destinationharness"
	gasdataharness "github.com/synapsecns/sanguine/agents/contracts/test/gasdata"
	"github.com/synapsecns/sanguine/agents/contracts/test/headerharness"
	"github.com/synapsecns/sanguine/agents/contracts/test/lightmanagerharness"
	"github.com/synapsecns/sanguine/agents/contracts/test/messageharness"
	"github.com/synapsecns/sanguine/agents/contracts/test/originharness"
	"github.com/synapsecns/sanguine/agents/contracts/test/pingpongclient"
	"github.com/synapsecns/sanguine/agents/contracts/test/receiptharness"
	"github.com/synapsecns/sanguine/agents/contracts/test/requestharness"
	"github.com/synapsecns/sanguine/agents/contracts/test/snapshotharness"
	"github.com/synapsecns/sanguine/agents/contracts/test/stateharness"
	"github.com/synapsecns/sanguine/agents/contracts/test/summitharness"
	"github.com/synapsecns/sanguine/agents/contracts/test/testclient"
	"github.com/synapsecns/sanguine/agents/contracts/test/tipsharness"
	"github.com/synapsecns/sanguine/agents/testutil/agentstestcontract"
	"github.com/synapsecns/sanguine/ethergo/contracts"
)

// set all contact types.
func init() {
	for i := 1; i < len(_contractTypeImpl_index); i++ {
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
	OriginType contractTypeImpl = iota + 1 // Origin
	// MessageHarnessType is the type of the message harness contract.
	MessageHarnessType // MessageHarness
	// BaseMessageHarnessType is the type of the base message harness contract.
	BaseMessageHarnessType // BaseMessageHarness
	// ReceiptHarnessType is the type of the receipt harness contract.
	ReceiptHarnessType // ReceiptHarness
	// RequestHarnessType is the type of the request harness contract.
	RequestHarnessType // RequestHarness
	// OriginHarnessType is the origin harness type.
	OriginHarnessType // OriginHarness
	// StateHarnessType is the state harness type.
	StateHarnessType
	// SnapshotHarnessType is the snapshot harness type.
	SnapshotHarnessType
	// AttestationHarnessType is the type of the attestation harness.
	AttestationHarnessType
	// TipsHarnessType is the type of the tips harness.
	TipsHarnessType
	// HeaderHarnessType is the type of the header harness.
	HeaderHarnessType
	// DestinationHarnessType is the destination harness type.
	DestinationHarnessType // DestinationHarness
	// SummitHarnessType is the summit harness type.
	SummitHarnessType // SummitHarness
	// SummitType is the type of the summit.
	SummitType // Summit
	// DestinationType is the type of the destination.
	DestinationType // Destination
	// AgentsTestContractType is the type of the agents test contract.
	AgentsTestContractType // AgentsTestContract
	// TestClientType is the type of the test client.
	TestClientType // TestClient
	// PingPongClientType is the type of the test client.
	PingPongClientType // PingPongClient
	// LightManagerHarnessType is the light manager harness type.
	LightManagerHarnessType // LightManagerHarness
	// BondingManagerHarnessType is the bonding manager harness type.
	BondingManagerHarnessType // BondingManagerHarness
	// LightManagerType is the light manager type.
	LightManagerType // LightManager
	// BondingManagerType is the bonding manager type.
	BondingManagerType // BondingManager
	// GasDataHarnessType is the gasData harness type.
	GasDataHarnessType
	// GasOracleType is the gas oracle type.
	GasOracleType // GasOracle
	// InboxType is the inbox type.
	InboxType // Inbox
	// LightInboxType is the light inbox type.
	LightInboxType // LightInbox
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
	case BaseMessageHarnessType:
		return basemessageharness.Contracts["solidity/BaseMessageHarness.t.sol:BaseMessageHarness"]
	case ReceiptHarnessType:
		return receiptharness.Contracts["solidity/ReceiptHarness.t.sol:ReceiptHarness"]
	case RequestHarnessType:
		return requestharness.Contracts["solidity/RequestHarness.t.sol:RequestHarness"]
	case OriginHarnessType:
		return originharness.Contracts["solidity/OriginHarness.t.sol:OriginHarness"]
	case StateHarnessType:
		return stateharness.Contracts["solidity/StateHarness.t.sol:StateHarness"]
	case SnapshotHarnessType:
		return snapshotharness.Contracts["solidity/SnapshotHarness.t.sol:SnapshotHarness"]
	case AttestationHarnessType:
		return attestationharness.Contracts["solidity/AttestationHarness.t.sol:AttestationHarness"]
	case DestinationHarnessType:
		return destinationharness.Contracts["solidity/DestinationHarness.t.sol:DestinationHarness"]
	case SummitHarnessType:
		return summitharness.Contracts["solidity/SummitHarness.t.sol:SummitHarness"]
	case TipsHarnessType:
		return tipsharness.Contracts["solidity/TipsHarness.t.sol:TipsHarness"]
	case SummitType:
		return summit.Contracts["solidity/Summit.sol:Summit"]
	case DestinationType:
		return destination.Contracts["solidity/Destination.sol:Destination"]
	case HeaderHarnessType:
		return headerharness.Contracts["solidity/HeaderHarness.t.sol:HeaderHarness"]
	case AgentsTestContractType:
		return agentstestcontract.Contracts["solidity/AgentsTestContract.sol:AgentsTestContract"]
	case TestClientType:
		return testclient.Contracts["solidity/TestClient.sol:TestClient"]
	case PingPongClientType:
		return pingpongclient.Contracts["solidity/PingPongClient.sol:PingPongClient"]
	case LightManagerHarnessType:
		return lightmanagerharness.Contracts["solidity/LightManagerHarness.t.sol:LightManagerHarness"]
	case BondingManagerHarnessType:
		return bondingmanagerharness.Contracts["solidity/BondingManagerHarness.t.sol:BondingManagerHarness"]
	case LightManagerType:
		return lightmanager.Contracts["solidity/LightManager.sol:LightManager"]
	case BondingManagerType:
		return bondingmanager.Contracts["solidity/BondingManager.sol:BondingManager"]
	case GasDataHarnessType:
		return gasdataharness.Contracts["solidity/GasDataHarness.t.sol:GasDataHarness"]
	case GasOracleType:
		return gasoracle.Contracts["solidity/GasOracle.sol:GasOracle"]
	case InboxType:
		return inbox.Contracts["solidity/Inbox.sol:Inbox"]
	case LightInboxType:
		return lightinbox.Contracts["solidity/LightInbox.sol:LightInbox"]
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
