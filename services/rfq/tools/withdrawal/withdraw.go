// Package withdraw provides a wrapper around the RelayerClient's Withdraw method which allows a relayer to withdraw
// at ERC20s or the native token.
package withdrawal

import (
	"context"

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

// NewWithdrawer creates a new Withdrawer.
func NewWithdrawer(c relapi.RelayerClient) Withdrawer {
	return &withdrawerImpl{
		client: c,
	}
}

// TODO: support multiple withdraw requests in one cli command (via config?)
// Withdraw withdraws the given amount of tokens to the given address.
func (w *withdrawerImpl) Withdraw(ctx context.Context, withdrawRequest relapi.WithdrawRequest) (*relapi.WithdrawResponse, error) {
	return w.client.Withdraw(ctx, &withdrawRequest)
}
