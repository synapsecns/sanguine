package relapi

import (
	"context"
	"fmt"
)

// Withdrawer is a wrapper around the RelayerClient's Withdraw method which allows
// a relayer to withdraw at once multiple ERC20s or the native token.
type Withdrawer interface {
	Withdraw(ctx context.Context, withdrawRequests WithdrawRequest) (*WithdrawResponse, error)
}

type withdrawerImpl struct {
	client RelayerClient
}

// NewWithdrawer creates a new Withdrawer.
func NewWithdrawer(c RelayerClient) Withdrawer {
	return &withdrawerImpl{
		client: c,
	}
}

// TODO: support multiple withdraw requests in one cli command (via config?)
// Withdraw withdraws the given amount of tokens to the given address.
func (w *withdrawerImpl) Withdraw(ctx context.Context, withdrawRequest WithdrawRequest) (*WithdrawResponse, error) {
	res, err := w.client.Withdraw(ctx, &withdrawRequest)
	if err != nil {
		return nil, fmt.Errorf("could not withdraw: %w", err)
	}
	return res, nil
}
