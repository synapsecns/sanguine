package near

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rpc"
)

func getTransactions(ctx context.Context, c AuroraClient, txHashes []common.Hash, blockHash common.Hash) (txes []*types.Transaction, err error) {
	if len(txHashes) == 0 {
		return txes, nil
	}

	reqs := make([]rpc.BatchElem, len(txHashes))
	res := make([]*rpcTransaction, len(txHashes))

	for i := range reqs {
		reqs[i] = rpc.BatchElem{
			Method: "eth_getTransactionByHash",
			Args:   []interface{}{txHashes[i], hexutil.EncodeUint64(uint64(i))},
			Result: &res[i],
		}
	}

	if err := c.BatchCallContext(ctx, reqs); err != nil {
		//nolint: wrapcheck
		return nil, err
	}

	for i := range reqs {
		if reqs[i].Error != nil {
			return nil, reqs[i].Error
		}

		if res[i] == nil {
			return nil, fmt.Errorf("got null transaction for %s", reqs[i].Args)
		}
	}

	for _, tx := range res {
		if tx.From != nil {
			setSenderFromServer(tx.tx, *tx.From, blockHash)
		}
		txes = append(txes, tx.tx)
	}
	return txes, nil
}

type rpcTransaction struct {
	tx *types.Transaction
	txExtraInfo
}

// see:  https://git.io/Jy1OO
type txExtraInfo struct {
	BlockNumber *string         `json:"blockNumber,omitempty"`
	BlockHash   *common.Hash    `json:"blockHash,omitempty"`
	From        *common.Address `json:"from,omitempty"`
}

//nolint:wrapcheck
func (tx *rpcTransaction) UnmarshalJSON(msg []byte) error {
	if err := json.Unmarshal(msg, &tx.tx); err != nil {
		//nolint:wrapcheck
		return err
	}
	return json.Unmarshal(msg, &tx.txExtraInfo)
}
