package origin

import (
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
)

// OriginRef is a bound origin contract that returns the address of the contract.
//
//nolint:golint
type OriginRef struct {
	*Origin
	address common.Address
	parser  Parser
}

// Address is the contract address.
func (s OriginRef) Address() common.Address {
	return s.address
}

// Parser returns the origin parser.
func (s OriginRef) Parser() Parser {
	return s.parser
}

// NewOriginRef creates a new origin contract with a contract ref.
func NewOriginRef(address common.Address, backend bind.ContractBackend) (*OriginRef, error) {
	originContract, err := NewOrigin(address, backend)
	if err != nil {
		return nil, err
	}

	parser, err := NewParser(address)
	if err != nil {
		return nil, fmt.Errorf("could not create parser: %w", err)
	}

	return &OriginRef{
		Origin:  originContract,
		address: address,
		parser:  parser,
	}, nil
}

var _ vm.ContractRef = OriginRef{}

// IOrigin wraps the generated origin interface code.
type IOrigin interface {
	IOriginCaller
	IOriginFilterer
	IOriginTransactor
	vm.ContractRef
	Parser() Parser
}
