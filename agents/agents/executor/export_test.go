package executor

import (
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/core/merkle"
)

// GetLogChan gets a log channel.
func (e Executor) GetLogChan(chainID uint32) chan *ethTypes.Log {
	return e.chainExecutors[chainID].logChan
}

// GetMerkleTree gets a merkle tree.
func (e Executor) GetMerkleTree(chainID uint32, domain uint32) *merkle.HistoricalTree {
	return e.chainExecutors[chainID].merkleTrees[domain]
}
