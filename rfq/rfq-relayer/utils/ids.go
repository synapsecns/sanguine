package utils

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
)

// GenerateTokenID generates a token ID from the token address and chain ID.
func GenerateTokenID(chainID uint32, tokenAddress common.Address) string {
	return fmt.Sprintf("%d-%s", chainID, tokenAddress.String())
}

// GenerateQuoteID generates a quote ID from the origin chain ID, origin token address, destination chain ID, and destination token address.
func GenerateQuoteID(originChainID uint32, originToken common.Address, destChainID uint32, destToken common.Address) string {
	return fmt.Sprintf("%d-%s-%d-%s", originChainID, originToken.String(), destChainID, destToken.String())
}
