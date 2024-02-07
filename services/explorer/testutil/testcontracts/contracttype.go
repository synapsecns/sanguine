package testcontracts

import (
	"github.com/ethereum/go-ethereum/common/compiler"
	"github.com/synapsecns/sanguine/ethergo/contracts"
	"github.com/synapsecns/sanguine/services/explorer/contracts/bridge/testbridge"
	"github.com/synapsecns/sanguine/services/explorer/contracts/bridge/testbridgev1"
	"github.com/synapsecns/sanguine/services/explorer/contracts/bridgeconfig"
	"github.com/synapsecns/sanguine/services/explorer/contracts/cctp/testcctp"
	"github.com/synapsecns/sanguine/services/explorer/contracts/messagebus/testmessagebus"
	"github.com/synapsecns/sanguine/services/explorer/contracts/metaswap/testmetaswap"
	"github.com/synapsecns/sanguine/services/explorer/contracts/swap/testswap"
)

func init() {
	for i := 0; i < len(_contractTypeImpl_index); i++ {
		contractType := contractTypeImpl(i)
		AllContractTypes = append(AllContractTypes, contractType)
		// assert type is correct
		var _ contracts.ContractType = contractType

		// Checks for discrepancies in contract reference. Will panic if contract is not properly configured.
		contractType.ContractName()
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
	// TestBridgeConfigTypeV3 is the bridge config contract type.
	TestBridgeConfigTypeV3 contractTypeImpl = iota
	// TestSynapseBridgeType is the bridge contract type.
	TestSynapseBridgeType
	// TestSwapFlashLoanType is the swap contract type.
	TestSwapFlashLoanType
	// TestSynapseBridgeV1Type is the bridge contract type for V1.
	TestSynapseBridgeV1Type
	// TestMessageBusType is the message bus type.
	TestMessageBusType
	// TestMetaSwapType is the meta swap type.
	TestMetaSwapType
	// TestCCTPType is cctp contract type.
	TestCCTPType
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
	case TestBridgeConfigTypeV3:
		return bridgeconfig.Contracts["/solidity/BridgeConfigV3_flat.sol:BridgeConfigV3"]
	case TestSynapseBridgeType:
		return testbridge.Contracts["/solidity/TestSynapseBridge.sol:TestSynapseBridge"]
	case TestSwapFlashLoanType:
		return testswap.Contracts["/solidity/TestSwapFlashLoan.sol:TestSwapFlashLoan"]
	case TestSynapseBridgeV1Type:
		return testbridgev1.Contracts["/solidity/TestSynapseBridgeV1.sol:TestSynapseBridgeV1"]
	case TestMessageBusType:
		return testmessagebus.Contracts["/solidity/TestMessageBusUpgradeable.sol:TestMessageBusUpgradeable"]
	case TestMetaSwapType:
		return testmetaswap.Contracts["/solidity/TestMetaSwapV1.sol:TestMetaSwap"]
	case TestCCTPType:
		return testcctp.Contracts["solidity/TestSynapseCCTPV1.sol:TestSynapseCCTP"]

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
