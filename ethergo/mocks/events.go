package mocks

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"math/rand"
	"testing"
	"time"
)

// NewMockHash creates a mock hash.
func NewMockHash(tb testing.TB) common.Hash {
	tb.Helper()
	//nolint: gosec
	src := rand.New(rand.NewSource(time.Now().UnixNano()))
	eventTopics := common.Hash{}.Generate(src, common.HashLength)
	hash, ok := eventTopics.Interface().(common.Hash)
	if !ok {
		tb.Errorf("failed to generate hash")
	}
	return hash
}

// GetMockLogs gets eventCount mock events.
func GetMockLogs(tb testing.TB, eventCount int) (logs []types.Log) {
	tb.Helper()
	eventAddress := MockAddress()
	eventTopic := []common.Hash{NewMockHash(tb)}
	blockHash := NewMockHash(tb)
	txHash := NewMockHash(tb)
	for i := 0; i < eventCount; i++ {
		logs = append(logs, types.Log{
			Address: eventAddress,
			Topics:  eventTopic,
			// TODO mock me
			Data:        nil,
			BlockNumber: 0,
			TxHash:      txHash,
			TxIndex:     uint(i),
			BlockHash:   blockHash,
			Index:       uint(i),
			Removed:     false,
		})
	}
	return
}
