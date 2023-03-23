package abi

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"strings"
)

// GetSelectorByName takes a function name and a pointer to a `bind.MetaData` object,
// searches for the first function signature in the `Sigs` map of the metadata object
// that matches the function name, and returns the first four bytes of the keccak256 hash
// of the function signature as a `[4]byte` array.
func GetSelectorByName(name string, metadata *bind.MetaData) ([4]byte, error) {
	var matchingSigs []string

	// search for function signatures that match the function name
	for sig, desc := range metadata.Sigs {
		if strings.HasPrefix(desc, name+"(") && strings.HasSuffix(desc, ")") {
			matchingSigs = append(matchingSigs, sig)
		}
	}

	if len(matchingSigs) == 0 {
		// if there are no matching signatures, return an error
		return [4]byte{}, fmt.Errorf("no function with name %s", name)
	} else if len(matchingSigs) > 1 {
		// if there are multiple matching signatures, return an error
		return [4]byte{}, fmt.Errorf("multiple functions with name %s", name)
	}

	// extract the function selector bytes from the only matching signature
	selectorBytes := common.Hex2Bytes(matchingSigs[0])
	var selector [4]byte
	copy(selector[:], selectorBytes)

	return selector, nil
}
