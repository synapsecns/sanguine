package withdraw

import (
	"context"

	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/services/rfq/relayer/relapi"
)

// Withdrawer is a wrapper around the RelayerClient's Withdraw method which allows
// a relayer to withdraw at once multiple ERC20s or the native token.
type Withdrawer interface {
	Withdraw(ctx context.Context, withdrawRequests relapi.WithdrawRequest) (*relapi.WithdrawResponse, error)
}

type withdrawerImpl struct {
	client relapi.RelayerClient
}

func NewWithdrawer(handler metrics.Handler, url string) Withdrawer {
	return &withdrawerImpl{
		client: relapi.NewRelayerClient(handler, url),
	}
}

// TODO: support multiple withdraw requests in one cli command (via config?)
func (w *withdrawerImpl) Withdraw(ctx context.Context, withdrawRequest relapi.WithdrawRequest) (*relapi.WithdrawResponse, error) {
	return w.client.Withdraw(ctx, &withdrawRequest)
}
