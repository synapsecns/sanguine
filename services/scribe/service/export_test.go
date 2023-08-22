package service

import (
	"github.com/synapsecns/sanguine/services/scribe/config"
)

// GetLivefillContracts returns the array of livefill contracts for testing.
func (c *ChainIndexer) GetLivefillContracts() []config.ContractConfig {
	return c.livefillContracts
}
