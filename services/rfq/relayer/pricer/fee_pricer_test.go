package pricer_test

import (
	"github.com/synapsecns/sanguine/ethergo/submitter/mocks"
	"github.com/synapsecns/sanguine/services/rfq/relayer/pricer"
)

func (s *PricerSuite) TestGetOriginFee() {
	clientFetcher := &mocks.ClientFetcher{}
	feePricer := pricer.NewFeePricer(s.feePricerConfig, s.chainConfigs, clientFetcher)
}
