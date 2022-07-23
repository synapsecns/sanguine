package home

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
)

// HomeRef is a bound home contract that returns the address of the contract.
//nolint: golint
type HomeRef struct {
	*Home
	address common.Address
	parser  Parser
}

// Address is the contract address.
func (s HomeRef) Address() common.Address {
	return s.address
}

// Parser returns the home parser.
func (s HomeRef) Parser() Parser {
	return s.parser
}

// NewHomeRef creates a new home contract with a contract ref and an interface for parsing events into custom types.
func NewHomeRef(address common.Address, backend bind.ContractBackend) (*HomeRef, error) {
	homeContract, err := NewHome(address, backend)
	if err != nil {
		return nil, err
	}

	parser, err := NewParser(address)
	if err != nil {
		return nil, fmt.Errorf("could not create parser: %w", err)
	}

	return &HomeRef{
		Home:    homeContract,
		address: address,
		parser:  parser,
	}, nil
}

var _ vm.ContractRef = HomeRef{}

// IHome wraps the generated home interface code.
type IHome interface {
	IHomeCaller
	IHomeFilterer
	IHomeTransactor
	vm.ContractRef
	Parser() Parser
}
