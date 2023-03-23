package contracts

import "github.com/ethereum/go-ethereum/common/compiler"

// ContractType is a contract type interface that contracts need to comply with.
//
//go:generate go run github.com/vektra/mockery/v2 --name ContractType --output ./mocks --case=underscore
type ContractType interface {
	// ID gets the unique identifier for the contracts
	ID() int
	// Name gets a the contracts name
	Name() string
	// ContractInfo gets the contract info from the compiler contract.
	ContractInfo() *compiler.Contract
	// ContractName gets the name fo the deployed contract
	ContractName() string
}
