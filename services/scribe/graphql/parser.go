package graphql

import (
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/services/scribe/graphql/server/graph/model"
)

// ParseLog converts a log from GraphQL into an ethType log.
func ParseLog(log interface{}) (*ethTypes.Log, error) {
	marshalledLog, err := json.Marshal(log)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal log: %w", err)
	}
	var unmarshalledLog model.Log
	err = json.Unmarshal(marshalledLog, &unmarshalledLog)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal log: %w", err)
	}
	return buildLogFromModelLogs(unmarshalledLog), nil
}

func buildLogFromModelLogs(log model.Log) *ethTypes.Log {
	var topics []common.Hash
	for _, topic := range log.Topics {
		topics = append(topics, common.HexToHash(topic))
	}

	return &ethTypes.Log{
		Address:     common.HexToAddress(log.ContractAddress),
		Topics:      topics,
		Data:        common.FromHex(log.Data),
		BlockNumber: uint64(log.BlockNumber),
		TxHash:      common.HexToHash(log.TxHash),
		TxIndex:     uint(log.TxIndex),
		BlockHash:   common.HexToHash(log.BlockHash),
		Index:       uint(log.Index),
		Removed:     log.Removed,
	}
}
