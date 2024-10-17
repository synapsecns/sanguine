package mixins

import (
	"bytes"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/ethergo/parser/rpc"
)

// ReqToTX converts a request to a transaction.
func ReqToTX(req rpc.Request) (tx *types.Transaction, err error) {
	tx = new(types.Transaction)

	hex := common.FromHex(string(bytes.ReplaceAll(req.Params[0], []byte{'"'}, []byte{})))
	err = tx.UnmarshalBinary(hex)
	if err != nil {
		return nil, fmt.Errorf("could not unmarshal transaction: %w", err)
	}

	return tx, nil
}
