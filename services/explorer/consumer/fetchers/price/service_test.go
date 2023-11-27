package price_test

import (
	"context"
	"github.com/synapsecns/sanguine/services/explorer/consumer/fetchers/price"
	"testing"

	. "github.com/stretchr/testify/assert"
)

func TestGetDefiLlamaData(t *testing.T) {
	priceFetcher, err := price.NewPriceFetcher()
	Nil(t, err)
	amount := priceFetcher.GetPriceData(context.Background(), 1648680149, "ethereum")
	NotNil(t, amount)
	Equal(t, 3386, int(*amount))
}
