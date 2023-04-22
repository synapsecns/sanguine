package fetcher_test

import (
	"context"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/services/explorer/consumer/fetcher"
	"testing"
)

func TestGetDefiLlamaData(t *testing.T) {
	amount := fetcher.GetDefiLlamaData(context.Background(), 1648680149, "ethereum")
	NotNil(t, amount)
	Equal(t, 3386, int(*amount))
}

//
// func TestGetCoinGeckoPriceData(t *testing.T) {
//	amount := fetcher.GetCoinGeckoPriceData(context.Background(), 1648680149, "ethereum", 3)
//	NotNil(t, amount)
//	fmt.Println(*amount)
//}
