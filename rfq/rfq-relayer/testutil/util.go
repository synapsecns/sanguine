package testutil

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
)

// GenerateTestLog generates a test log.
func GenerateTestLog() *types.Log {
	return &types.Log{
		Address:     common.BigToAddress(big.NewInt(gofakeit.Int64())),
		Topics:      []common.Hash{common.HexToHash(big.NewInt(gofakeit.Int64()).String())},
		Data:        []byte{1, 2, 3},
		BlockNumber: 42,
		TxHash:      common.BigToHash(big.NewInt(gofakeit.Int64())),
		BlockHash:   common.BigToHash(big.NewInt(gofakeit.Int64())),
		Removed:     false,
		TxIndex:     uint(gofakeit.Uint8()),
		Index:       uint(gofakeit.Uint8()),
	}
}
