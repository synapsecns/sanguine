package testutil

import (
	"github.com/ethereum/go-ethereum/common/compiler"
	"github.com/synapsecns/sanguine/ethergo/backends/base"
	"github.com/synapsecns/sanguine/ethergo/contracts"
	"github.com/synapsecns/sanguine/services/rfq/contracts/fastbridge"
	"github.com/synapsecns/sanguine/services/rfq/contracts/testcontracts/dai"
	"github.com/synapsecns/sanguine/services/rfq/contracts/testcontracts/fastbridgemock"
	"github.com/synapsecns/sanguine/services/rfq/contracts/testcontracts/mockerc20"
	"github.com/synapsecns/sanguine/services/rfq/contracts/testcontracts/usdc"
	"github.com/synapsecns/sanguine/services/rfq/contracts/testcontracts/weth9"
)

// set all contact types.
func init() {
	base.AddToVerificationBlacklist(USDTType)
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
	// FastBridgeType is the type of the fast bridge contract.
	FastBridgeType contractTypeImpl = iota + 1 // FastBridge
	// MockERC20Type is a mock erc20 contract.
	MockERC20Type // MockERC20
	// FastBridgeMockType is a mock contract for testing fast bridge interactions
	// TODO: rename  contract to MockFastBridge.
	FastBridgeMockType // FastBridgeMock
	// WETH9Type  is the weth 9 contract.
	WETH9Type // WETH9
	// USDTType is the tether type.
	USDTType // USDT
	// USDCType is the type of the usdc contract.
	USDCType // USDC
	// DAIType is the dai contract.
	DAIType // DAI is the dai contract type
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
	case FastBridgeType:
		return fastbridge.Contracts["solidity/FastBridge.sol:FastBridge"]
	case MockERC20Type:
		return mockerc20.Contracts["solidity/MockERC20.sol:MockERC20"]
	case FastBridgeMockType:
		return fastbridgemock.Contracts["solidity/FastBridgeMock.sol:FastBridgeMock"]
	case WETH9Type:
		return weth9.Contracts["/solidity/WETH9.sol:WETH9"]
	case USDTType:
		panic("method not supported, token is verification blacklisted")
	case USDCType:
		return usdc.Contracts["/solidity/FiatToken.sol:FiatTokenV2"]
	case DAIType:
		return dai.Contracts["/solidity/dai.sol:Dai"]
	}
	return nil
}

var _ contracts.ContractType = contractTypeImpl(1)
