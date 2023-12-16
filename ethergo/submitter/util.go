package submitter

import (
	"github.com/synapsecns/sanguine/ethergo/chain/gas"
	"github.com/synapsecns/sanguine/ethergo/submitter/db"
	"sort"
)

// sortTxesByChainID sorts a slice of transactions by nonce.
func sortTxesByChainID(txs []db.TX) map[uint64][]db.TX {
	txesByChainID := make(map[uint64][]db.TX)
	// put the transactions in a map by chain id
	for _, t := range txs {
		txesByChainID[t.ChainId().Uint64()] = append(txesByChainID[t.ChainId().Uint64()], t)
	}

	for key := range txesByChainID {
		key := key // capture range variable
		// sort the transactions by nonce
		sort.Slice(txesByChainID[key], func(i, j int) bool {
			iNonce := txesByChainID[key][i].Nonce()
			jNonce := txesByChainID[key][j].Nonce()

			if iNonce == jNonce {
				gasCmp := gas.CompareGas(txesByChainID[key][i].Transaction, txesByChainID[key][j].Transaction, nil)
				if gasCmp == 0 || gasCmp == 1 {
					return false
				}
				return true
			}
			return iNonce < jNonce
		})
	}

	return txesByChainID
}

// groupTxesByNonce groups a slice of transactions by nonce.
// this will not differentiate between transactions with different chain ids.
func groupTxesByNonce(txs []db.TX) map[uint64][]db.TX {
	txesByNonce := make(map[uint64][]db.TX)
	for _, t := range txs {
		txesByNonce[t.Nonce()] = append(txesByNonce[t.Nonce()], t)
	}

	return txesByNonce
}
