package executor

import (
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/prysmaticlabs/prysm/shared/trieutil"
)

// GetLogChan gets a log channel.
func (e *Executor) GetLogChan(chainID uint32) chan *ethTypes.Log {
	return e.logChans[chainID]
}

// GetMerkleTree gets a merkle tree.
func (e *Executor) GetMerkleTree(chainID uint32, domain uint32) *trieutil.SparseMerkleTrie {
	return e.merkleTrees[chainID][domain]
}
