package abiutil

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"strings"
)

// MustGetSelectorByName is a wrapper around `GetSelectorByName` that panics if an error is returned.
func MustGetSelectorByName(name string, metadata *bind.MetaData) [4]byte {
	selector, err := GetSelectorByName(name, metadata)
	if err != nil {
		panic(err)
	}

	return selector
}

// GetSelectorByName takes a function name and a pointer to a `bind.MetaData` object,
// searches for the first function testContract in the `Sigs` map of the metadata object
// that matches the function name, and returns the first four bytes of the keccak256 hash
// of the function testContract as a `[4]byte` array.
func GetSelectorByName(name string, metadata *bind.MetaData) ([4]byte, error) {
	matchingSig, err := GetStringSelectorByName(name, metadata)
	if err != nil {
		return [4]byte{}, err
	}
	// extract the function selector bytes from the only matching testContract
	selectorBytes := common.Hex2Bytes(matchingSig)
	var selector [4]byte
	copy(selector[:], selectorBytes)

	return selector, nil
}

// GetStringSelectorByName takes a function name and a pointer to a `bind.MetaData` object,
// searches for the first function testContract in the `Sigs` map of the metadata object
// that matches the function name, and returns the first four bytes of the keccak256 hash
// of the function testContract as a string.
func GetStringSelectorByName(name string, metadata *bind.MetaData) (string, error) {
	var matchingSigs []string

	// search for function signatures that match the function name
	for sig, desc := range metadata.Sigs {
		if strings.HasPrefix(desc, name+"(") && strings.HasSuffix(desc, ")") {
			matchingSigs = append(matchingSigs, sig)
		}
	}

	if len(matchingSigs) == 0 {
		// if there are no matching signatures, return an error
		return "", fmt.Errorf("no function with name %s", name)
	} else if len(matchingSigs) > 1 {
		// if there are multiple matching signatures, return an error
		return "", fmt.Errorf("multiple functions with name %s", name)
	}

	// extract the function selector bytes from the only matching testContract
	return matchingSigs[0], nil
}
