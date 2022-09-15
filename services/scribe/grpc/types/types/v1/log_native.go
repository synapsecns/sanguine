package pbscribe

import (
	"github.com/ethereum/go-ethereum/core/types"
)

// FromNativeLog converts a native log to a proto log.
func FromNativeLog(log *types.Log) *Log {
	return &Log{
		Address:     FromNativeAddress(log.Address),
		Topics:      FromNativeHashes(log.Topics),
		Data:        log.Data,
		BlockNumber: log.BlockNumber,
		TxHash:      FromNativeHash(log.TxHash),
		TxIndex:     uint64(log.TxIndex),
		BlockHash:   FromNativeHash(log.BlockHash),
		Index:       uint64(log.Index),
		Removed:     log.Removed,
	}
}

// ToLog converts a log type to a native log.
func (x *Log) ToLog() *types.Log {
	return &types.Log{
		Address:     x.GetAddress().ToAddress(),
		Topics:      ToNativeHashes(x.Topics),
		Data:        x.GetData(),
		BlockNumber: x.GetBlockNumber(),
		TxHash:      x.GetTxHash().ToHash(),
		TxIndex:     uint(x.GetTxIndex()),
		BlockHash:   x.BlockHash.ToHash(),
		Index:       uint(x.Index),
		Removed:     x.Removed,
	}
}

// FromNativeLogs is a helper function for converting a batch of logs all at once.
func FromNativeLogs(logs []*types.Log) (res []*Log) {
	for _, log := range logs {
		res = append(res, FromNativeLog(log))
	}
	return res
}

// ToNativeLogs is a helper function for converting logs to native logs.
func ToNativeLogs(logs []*Log) (res []*types.Log) {
	for _, log := range logs {
		res = append(res, log.ToLog())
	}
	return res
}
