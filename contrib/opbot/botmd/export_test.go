package botmd

import (
	"context"

	"github.com/synapsecns/sanguine/ethergo/client"
)

func StripLinks(input string) string {
	return stripLinks(input)
}

func GetTxAge(ctx context.Context, client client.EVM, res Status) string {
	return getTxAge(ctx, client, res)
}
