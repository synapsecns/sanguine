package types

import "github.com/ethereum/go-ethereum/common"

type IndexerConfig struct {
	Contracts            []common.Address
	GetLogsRange         uint64
	GetLogsBatchAmount   uint64
	StoreConcurrency     int
	ChainID              uint32
	StartHeight          uint64
	EndHeight            uint64
	ConcurrencyThreshold uint64
}
