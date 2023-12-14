package inventory_test

import (
	"github.com/synapsecns/sanguine/core/metrics"
	omnirpcClient "github.com/synapsecns/sanguine/services/omnirpc/client"
	"github.com/synapsecns/sanguine/services/rfq/relayer/inventory"
	"github.com/synapsecns/sanguine/services/rfq/relayer/relconfig"
)

func (i *InventoryTestSuite) TestInventoryBoot() {
	cfg := relconfig.Config{
		Tokens: map[int][]string{},
	}

	for _, backend := range i.backends {
		handle, _ := i.manager.GetMockERC20(i.GetTestContext(), backend)
		cfg.Tokens[int(backend.GetChainID())] = []string{handle.Address().String()}
	}

	im, err := inventory.NewInventoryManager(i.GetTestContext(), omnirpcClient.NewOmnirpcClient(i.omnirpcURL, metrics.Get()), metrics.Get(), cfg, i.relayer.Address(), i.db)
	i.Require().NoError(err)

	_ = im

}
