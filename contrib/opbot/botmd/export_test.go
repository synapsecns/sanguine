package botmd

import (
	"context"

	"github.com/synapsecns/sanguine/ethergo/client"
	"github.com/synapsecns/sanguine/services/rfq/relayer/relapi"
)

func StripLinks(input string) string {
	return stripLinks(input)
}

func GetTxAge(ctx context.Context, client client.EVM, res *relapi.GetQuoteRequestStatusResponse) string {
	return getTxAge(ctx, client, res)
}
