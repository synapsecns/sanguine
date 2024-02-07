package testutil

import (
	"github.com/ethereum/go-ethereum/common/compiler"
	"github.com/synapsecns/sanguine/ethergo/contracts"
	"github.com/synapsecns/sanguine/services/explorer/contracts/bridge"
	"github.com/synapsecns/sanguine/services/explorer/contracts/bridge/bridgev1"
	"github.com/synapsecns/sanguine/services/explorer/contracts/bridgeconfig"
	"github.com/synapsecns/sanguine/services/explorer/contracts/cctp"
	"github.com/synapsecns/sanguine/services/explorer/contracts/messagebus"
	"github.com/synapsecns/sanguine/services/explorer/contracts/metaswap"
	"github.com/synapsecns/sanguine/services/explorer/contracts/swap"
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
	// BridgeConfigTypeV3 is the bridge config contract type.
	BridgeConfigTypeV3 contractTypeImpl = iota
	// SynapseBridgeType is the bridge contract type.
	SynapseBridgeType
	// SwapFlashLoanType is the swap contract type.
	SwapFlashLoanType
	// SynapseBridgeV1Type is the swap contract type.
	SynapseBridgeV1Type
	// MessageBusType is the messagebus contract type.
	MessageBusType
	// MetaSwapType is the metaswap contract type.
	MetaSwapType
	// CCTPType is cctp contract type.
	CCTPType
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
	case BridgeConfigTypeV3:
		return bridgeconfig.Contracts["/solidity/BridgeConfigV3_flat.sol:BridgeConfigV3"]
	case SynapseBridgeType:
		return bridge.Contracts["/solidity/SynapseBridgeV2_flat.sol:SynapseBridge"]
	case SwapFlashLoanType:
		return swap.Contracts["/solidity/SwapFlashLoanV1_flat.sol:SwapFlashLoan"]
	case SynapseBridgeV1Type:
		return bridgev1.Contracts["/solidity/SynapseBridgeV1_flat.sol:SynapseBridge"]
	case MessageBusType:
		return messagebus.Contracts["/solidity/MessageBusUpgradeableV1_flat.sol:MessageBusUpgradeable"]
	case MetaSwapType:
		return metaswap.Contracts["/solidity/MetaSwapV1_flat.sol:MetaSwap"]
	case CCTPType:
		return cctp.Contracts["solidity/SynapseCCTPV1_flat.sol:SynapseCCTP"]
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
