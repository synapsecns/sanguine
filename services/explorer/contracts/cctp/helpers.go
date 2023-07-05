package cctp

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
)

// CCTPRef
//
//nolint:golint
type CCTPRef struct {
	*SynapseCCTP
	address common.Address
}

// Address is the contract address.
func (s CCTPRef) Address() common.Address {
	return s.address
}

// NewCCTPRef
//
//nolint:golint
func NewCCTPRef(address common.Address, backend bind.ContractBackend) (*CCTPRef, error) {
	synapseCCTP, err := NewSynapseCCTP(address, backend)
	if err != nil {
		return nil, err
	}
	return &CCTPRef{
		SynapseCCTP: synapseCCTP,
		address:     address,
	}, nil
}

var _ vm.ContractRef = &BridgeConfigRef{}
