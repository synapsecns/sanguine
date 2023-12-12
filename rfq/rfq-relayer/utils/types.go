package utils

import "github.com/ethereum/go-ethereum/core/types"

// WrappedLog is a struct containing the origin chain ID and a log.
type WrappedLog struct {
	OriginChainID uint32
	Log           types.Log
}
