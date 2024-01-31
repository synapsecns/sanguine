package pricer

import (
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/ethergo/submitter"
	"github.com/synapsecns/sanguine/services/rfq/relayer/relconfig"
)

func NewFeePricerWithFetchers(config relconfig.Config, clientFetcher submitter.ClientFetcher, handler metrics.Handler, priceFetcher CoingeckoPriceFetcher) FeePricer {
	pricer := NewFeePricer(config, clientFetcher, handler)
	feePricerImpl, _ := pricer.(*feePricer)
	feePricerImpl.priceFetcher = priceFetcher
	return feePricerImpl
}
