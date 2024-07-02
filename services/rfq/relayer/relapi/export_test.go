package relapi

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/synapsecns/sanguine/services/rfq/relayer/relconfig"
)

// TokenIDExists checks if a token ID exists in the config.
func TokenIDExists(cfg relconfig.Config, tokenAddress common.Address, chainID int) bool {
	return tokenIDExists(cfg, tokenAddress, chainID)
}
