package confirmedtofinalized

import (
	"errors"
	"github.com/ethereum/go-ethereum/rpc"
)

var (
	latestBlock    []byte
	finalizedBlock []byte
)

func init() {
	var err error
	latestBlock, err = rpc.LatestBlockNumber.MarshalText()
	if err != nil {
		panic(errors.New("could not marshall test from latest block number"))
	}

	finalizedBlock, err = rpc.FinalizedBlockNumber.MarshalText()
	if err != nil {
		panic(errors.New("could not marshall test from finalized block number"))
	}
}
