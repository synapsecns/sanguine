package types

import "fmt"

func init() {
	// boot sanity check that we re-generated with stringer
	expectedChainCount := len(_ChainType_index) - 1
	actualChainCount := len(AllChainTypes())

	if actualChainCount != expectedChainCount {
		panic(fmt.Sprintf("Expected exactly %d chains, got %d. You may need to rerun stringer", expectedChainCount, actualChainCount))
	}
}

// ChainType is the type of chain being used (e.g. evm).
//
//go:generate go run golang.org/x/tools/cmd/stringer -type=ChainType -linecomment
type ChainType uint8

const (
	// EVM is a chain that conforms to the evm standard.
	EVM ChainType = 0
)

// AllChainTypes gets all chain types for the chain.
func AllChainTypes() []ChainType {
	return []ChainType{EVM}
}
