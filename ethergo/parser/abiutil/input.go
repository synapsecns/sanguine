package abiutil

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

const selectorLength = 4

// UnpackInputDataToInterface unpacks input data to interface.
func UnpackInputDataToInterface(v interface{}, input []byte, metadata *bind.MetaData) error {
	abiData, err := metadata.GetAbi()
	if err != nil {
		return fmt.Errorf("failed to get abiData: %w", err)
	}

	method, err := getMethod(input, abiData)
	if err != nil {
		return fmt.Errorf("failed to get method by id: %w", err)
	}

	inputs, err := method.Inputs.Unpack(input[selectorLength:])
	if err != nil {
		return fmt.Errorf("failed to unpack inputs: %w", err)
	}

	err = method.Inputs.Copy(v, inputs)
	if err != nil {
		return fmt.Errorf("failed to copy inputs: %w", err)
	}
	return nil
}

// UnpackInputData takes a function name and a pointer to a `bind.MetaData` object,.
func UnpackInputData(input []byte, metadata *bind.MetaData) ([]interface{}, error) {
	abiData, err := metadata.GetAbi()
	if err != nil {
		return nil, fmt.Errorf("failed to get abiData: %w", err)
	}

	method, err := getMethod(input, abiData)
	if err != nil {
		return nil, fmt.Errorf("failed to get method by id: %w", err)
	}

	res, err := method.Inputs.Unpack(input[selectorLength:])
	if err != nil {
		return nil, fmt.Errorf("failed to unpack inputs: %w", err)
	}
	return res, nil
}

// getMethod takes a function name and a pointer to a `bind.MetaData` object,
// and returns the `abi.Method` object for that function.
// If the function is not found, an error is returned.
func getMethod(input []byte, abiData *abi.ABI) (*abi.Method, error) {
	if len(input) < selectorLength {
		return nil, fmt.Errorf("input too short")
	}

	// get the selector from the input, this will be the first 4 bytes
	selector := [selectorLength]byte{}
	copy(selector[:], input[:selectorLength])

	method, err := abiData.MethodById(selector[:])
	if err != nil {
		return nil, fmt.Errorf("failed to get method by id: %w", err)
	}
	return method, nil
}
