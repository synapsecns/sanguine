package testutil

import (
	"github.com/ethereum/go-ethereum/common/compiler"
	"github.com/synapsecns/sanguine/ethergo/contracts"
	"github.com/synapsecns/sanguine/services/cctp-relayer/contracts/cctp"
	"github.com/synapsecns/sanguine/services/cctp-relayer/contracts/mockmessagetransmitter"
)

// contractTypeImpl is the type of the contract being saved/fetched.
// we use an interface here so the deploy helper here can be abstracted away from the synapse contracts
//
//go:generate go run golang.org/x/tools/cmd/stringer -type=contractTypeImpl -linecomment
type contractTypeImpl int

const (
	// SynapseCCTP is the type of the synapse cctp contract.
	SynapseCCTPType contractTypeImpl = iota + 1 // SynapseCCTP
	// MockMessageTransmitter is the type of the mock message transmitter contract.
	MockMessageTransmitterType // MockMessageTransmitter
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
		return cctp.Contracts["solidity/SynapseCCTP.sol/SynapseCCTP"]
	case MockMessageTransmitterType:
		return mockmessagetransmitter.Contracts["solidity/MockMessageTransmitter.sol:MockMessageTransmitter"]
	default:
		panic("not yet implemented")
	}
}

// make sure contractTypeImpl conforms to contracts.ContractType.
var _ contracts.ContractType = contractTypeImpl(1)
