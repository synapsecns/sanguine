package inventory

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/synapsecns/sanguine/services/rfq/relayer/relconfig"
)

// GetRebalance is a wrapper around the internal getRebalance function.
func GetRebalances(ctx context.Context, cfg relconfig.Config, inv map[int]map[common.Address]*TokenMetadata) (rebalances map[string]*RebalanceData, err error) {
	return getRebalances(ctx, cfg, inv)
}
