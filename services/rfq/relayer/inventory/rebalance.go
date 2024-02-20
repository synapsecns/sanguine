package inventory

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/ethergo/submitter"
	"github.com/synapsecns/sanguine/services/cctp-relayer/contracts/cctp"
)

// Rebalance is the interface for rebalancing inventory.
type Rebalance interface {
	Execute(ctx context.Context, txSubmitter submitter.TransactionSubmitter, handler metrics.Handler) error
}

type cctpRebalance struct {
	origin   int
	dest     int
	token    common.Address
	amount   *big.Int
	contract *cctp.SynapseCCTP
}

func (c *cctpRebalance) Execute(ctx context.Context, txSubmitter submitter.TransactionSubmitter, handler metrics.Handler) error {
	// TODO: implement
	return nil
}
