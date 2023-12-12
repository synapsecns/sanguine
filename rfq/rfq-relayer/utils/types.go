package utils

import "github.com/ethereum/go-ethereum/core/types"

type WrappedLog struct {
	OriginChainID uint32
	Log           types.Log
}
