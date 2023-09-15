package types

import "github.com/ethereum/go-ethereum/common"

// IndexerConfig holds metadata for the indexer. It is used to pass data uniformly and used in logging.
type IndexerConfig struct {
	Addresses            []common.Address
	GetLogsRange         uint64
	GetLogsBatchAmount   uint64
	StoreConcurrency     int
	ChainID              uint32
	StartHeight          uint64
	EndHeight            uint64
	ConcurrencyThreshold uint64
	Topics               [][]common.Hash
}
