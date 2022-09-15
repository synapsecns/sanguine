package util

import "github.com/ethereum/go-ethereum/core/types"

// LogPointer wraps a log in a pointer.
func LogPointer(log types.Log) *types.Log {
	return &log
}

// LogsPointer wraps logs in a pointer
// TODO: consider abstracting this out into a generic.
func LogsPointer(logs []types.Log) (res []*types.Log) {
	for _, log := range logs {
		res = append(res, LogPointer(log))
	}
	return res
}
