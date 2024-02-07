package cctp

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
)

// SynapseCCTPRef is a bound cctp contract that conforms to vm.ContractRef.
//
//nolint:golint
type SynapseCCTPRef struct {
	*SynapseCCTP
	address common.Address
}

// Address is the contract address.
func (s SynapseCCTPRef) Address() common.Address {
	return s.address
}

// NewSynapseCCTPRef creates a new SynapseCCTPRef contract with a contract ref.
func NewSynapseCCTPRef(address common.Address, backend bind.ContractBackend) (*SynapseCCTPRef, error) {
	cctpContract, err := NewSynapseCCTP(address, backend)
	if err != nil {
		return nil, err
	}

	return &SynapseCCTPRef{
		SynapseCCTP: cctpContract,
		address:     address,
	}, nil
}

var _ vm.ContractRef = SynapseCCTPRef{}
