package chain

import "github.com/ethereum/go-ethereum/common"

// EthAddress is the address of a chain's native gas token.
var EthAddress = common.HexToAddress("0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE")

// IsGasToken returns true if the given token is the gas token.
func IsGasToken(token common.Address) bool {
	return token == EthAddress
}
