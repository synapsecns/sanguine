package base

import (
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
)

// GetChainID exports getChainID for testing.
func GetChainID(tx *types.Transaction) (hasType bool, chainID *big.Int) {
	return getChainID(tx)
}

// GetGormFieldName wraps getGormFieldName to export it for testing.
func GetGormFieldName(model interface{}, fieldName string) string {
	return getGormFieldName(model, fieldName)
}
