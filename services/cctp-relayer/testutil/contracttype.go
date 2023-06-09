package testutil

import (
	"github.com/ethereum/go-ethereum/common/compiler"
	"github.com/synapsecns/sanguine/ethergo/contracts"
	"github.com/synapsecns/sanguine/services/cctp-relayer/contracts/cctp"
	"github.com/synapsecns/sanguine/services/cctp-relayer/contracts/mockmessagetransmitter"
	"github.com/synapsecns/sanguine/services/cctp-relayer/contracts/mockmintburntoken"
	"github.com/synapsecns/sanguine/services/cctp-relayer/contracts/mocktokenmessenger"
	"github.com/synapsecns/sanguine/services/cctp-relayer/contracts/mocktokenminter"
)

// set all contact types.
func init() {
	for i := 0; i < len(_contractTypeImpl_index)-1; i++ {
		contractType := contractTypeImpl(i + 1)
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

// contractTypeImpl is the type of the contract being saved/fetched.
// we use an interface here so the deploy helper here can be abstracted away from the synapse contracts
//
//go:generate go run golang.org/x/tools/cmd/stringer -type=contractTypeImpl -linecomment
type contractTypeImpl int

const (
	// SynapseCCTPType is the type of the synapse cctp contract.
	SynapseCCTPType contractTypeImpl = iota + 1 // SynapseCCTP
	// MockMessageTransmitterType is the type of the mock message transmitter contract.
	MockMessageTransmitterType // MockMessageTransmitter
	// MockTokenMessengerType is the type of the mock token messenger contract.
	MockTokenMessengerType // MockTokenMessenger
	// MockMintBurnTokenType is the type of the mock token messenger contract.
	MockMintBurnTokenType // MockMintBurnToken
	// MockTokenMinterType is the type of the mock token minter contract.
	MockTokenMinterType // MockTokenMinter
)

// verifyStringerUpdated verifies stringer is up to date (this index is included in stringer).
func verifyStringerUpdated(contractType contractTypeImpl) {
	if int(contractType) > len(_contractTypeImpl_index) {
		panic("please update stringer before running test again")
	}
}

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
	case SynapseCCTPType:
		return cctp.Contracts["solidity/SynapseCCTP.sol:SynapseCCTP"]
	case MockMessageTransmitterType:
		return mockmessagetransmitter.Contracts["solidity/MockMessageTransmitter.sol:MockMessageTransmitter"]
	case MockTokenMessengerType:
		return mocktokenmessenger.Contracts["solidity/MockTokenMessenger.sol:MockTokenMessenger"]
	case MockMintBurnTokenType:
		return mockmintburntoken.Contracts["solidity/MockMintBurnToken.sol:MockMintBurnToken"]
	case MockTokenMinterType:
		return mocktokenminter.Contracts["solidity/MockTokenMinter.sol:MockTokenMinter"]
	default:
		panic("not yet implemented")
	}
}

// make sure contractTypeImpl conforms to contracts.ContractType.
var _ contracts.ContractType = contractTypeImpl(1)
