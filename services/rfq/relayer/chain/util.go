package chain

import "github.com/ethereum/go-ethereum/common"

var ethAddress = common.HexToAddress("0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE")

// IsGasToken returns true if the given token is the gas token.
func IsGasToken(token common.Address) bool {
	return token == ethAddress
}
