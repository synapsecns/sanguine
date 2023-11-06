package erc20

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
)

// ERC20Contract is an interface for an ERC20 contract.
type ERC20Contract interface {
	// Address gets the contract address.
	Address() common.Address
	// ChainID gets the chainID this contract is deployed on.
	ChainID() uint32
	// GetTokenData retrieves token data such as decimal and symbol from an ERC20 token.
	GetTokenData(ctx context.Context) (uint8, string, error)
}

// ERC20Ref is a reference to a deployed erc20 contract.
type ERC20Ref struct {
	*ERC20
	chainID uint32
	address common.Address
}

// NewERC20Ref creates a ERC20Contract interface.
func NewERC20Ref(address common.Address, chainID uint32, backend bind.ContractBackend) (ERC20Contract, error) {
	erc20, err := NewERC20(address, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20Ref{
		ERC20:   erc20,
		chainID: chainID,
		address: address,
	}, nil
}

// Address gets the contract address.
func (s ERC20Ref) Address() common.Address {
	return s.address
}

// ChainID gets the chainID this contract is deployed on.
func (s ERC20Ref) ChainID() uint32 {
	return s.chainID
}

// GetTokenData retrieves token data such as decimal and symbol from an ERC20 token.
func (s ERC20Ref) GetTokenData(ctx context.Context) (uint8, string, error) {
	decimal, err := s.Decimals(&bind.CallOpts{Context: ctx})
	if err != nil {
		return 0, "", fmt.Errorf("could not get token decimal: %w", err)
	}

	symbol, err := s.Symbol(&bind.CallOpts{Context: ctx})
	if err != nil {
		return 0, "", fmt.Errorf("could not get token symbol: %w", err)
	}

	return decimal, symbol, nil
}

var _ vm.ContractRef = &ERC20Ref{}
