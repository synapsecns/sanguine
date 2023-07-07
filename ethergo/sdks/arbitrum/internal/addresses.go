package internal

import "github.com/ethereum/go-ethereum/common"

// GetGasInfoAddress returns the default node interface address.
func GetGasInfoAddress() common.Address {
	return common.HexToAddress("0x6e")
}

// GetNodeInterfaceAddress returns the default node interface address.
func GetNodeInterfaceAddress() common.Address {
	return common.HexToAddress("0xc8")
}
