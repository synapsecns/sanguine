package abi

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

// UnpackInputData takes a function name and a pointer to a `bind.MetaData` object,
func UnpackInputData(input []byte, metadata *bind.MetaData) error {
	abi, err := metadata.GetAbi()
	if err != nil {
		return fmt.Errorf("failed to get abi: %w", err)
	}

	if len(input) < 4 {
		return fmt.Errorf("input too short")
	}

	// get the selector from the input, this will be the first 4 bytes
	selector := [4]byte{}
	copy(selector[:], input[:4])

}
