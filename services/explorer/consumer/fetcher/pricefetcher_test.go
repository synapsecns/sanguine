package fetcher_test

import (
	"context"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/services/explorer/consumer/fetcher"
	"testing"
)

func TestGetDefiLlamaData(t *testing.T) {
	amount, symbol := fetcher.GetDefiLlamaData(context.Background(), 1648680149, "ethereum")
	NotNil(t, amount)
	NotNil(t, symbol)
	Equal(t, 3386.03, *amount)
}
