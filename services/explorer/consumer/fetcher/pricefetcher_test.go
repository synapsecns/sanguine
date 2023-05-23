package fetcher_test

import (
	"context"
	"testing"

	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/services/explorer/consumer/fetcher"
)

func TestGetDefiLlamaData(t *testing.T) {
	amount := fetcher.GetDefiLlamaData(context.Background(), 1648680149, "ethereum")
	NotNil(t, amount)
	Equal(t, 3386, int(*amount))
}
