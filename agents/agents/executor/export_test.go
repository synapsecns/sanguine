package executor

import ethTypes "github.com/ethereum/go-ethereum/core/types"

// GetLogChan gets a log channel.
func (e *Executor) GetLogChan(chainID uint32) chan *ethTypes.Log {
	return e.logChans[chainID]
}
