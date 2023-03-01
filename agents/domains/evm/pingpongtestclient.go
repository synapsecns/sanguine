package evm

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/event"
	"github.com/synapsecns/sanguine/agents/contracts/test/pingpongclient"
	"github.com/synapsecns/sanguine/agents/domains"
	"github.com/synapsecns/sanguine/ethergo/chain"
	"github.com/synapsecns/sanguine/ethergo/signer/nonce"
	"github.com/synapsecns/sanguine/ethergo/signer/signer"
)

// NewPingPongClientContract returns a bound ping pong test client contract.
//
//nolint:staticcheck
func NewPingPongClientContract(ctx context.Context, client chain.Chain, pingPongClientAddress common.Address) (domains.PingPongClientContract, error) {
	boundCountract, err := pingpongclient.NewPingPongClientRef(pingPongClientAddress, client)
	if err != nil {
		return nil, fmt.Errorf("could not create %T: %w", &pingpongclient.PingPongClientRef{}, err)
	}

	nonceManager := nonce.NewNonceManager(ctx, client, client.GetBigChainID())
	return pingPongClientContract{
		contract:     boundCountract,
		client:       client,
		nonceManager: nonceManager,
	}, nil
}

type pingPongClientContract struct {
	// contract contains the conract handle
	contract *pingpongclient.PingPongClientRef
	// client contains the evm client
	//nolint: staticcheck
	client chain.Chain
	// nonceManager is the nonce manager used for transacting with the chain
	nonceManager nonce.Manager
}

func (a pingPongClientContract) DoPing(ctx context.Context, signer signer.Signer, destination uint32, recipient common.Address, pings uint16) error {
	transactOpts, err := a.transactOptsSetup(ctx, signer)
	if err != nil {
		return fmt.Errorf("could not setup transact opts: %w", err)
	}

	_, err = a.contract.DoPing(transactOpts, destination, recipient, pings)
	if err != nil {
		return fmt.Errorf("could not send ping: %w", err)
	}

	return nil
}

func (a pingPongClientContract) WatchPingSent(ctx context.Context, sink chan<- *pingpongclient.PingPongClientPingSent) (event.Subscription, error) {
	sub, err := a.contract.WatchPingSent(&bind.WatchOpts{Context: ctx}, sink)
	if err != nil {
		return nil, fmt.Errorf("could set up channel to watch ping sent: %w", err)
	}

	return sub, nil
}

func (a pingPongClientContract) WatchPongReceived(ctx context.Context, sink chan<- *pingpongclient.PingPongClientPongReceived) (event.Subscription, error) {
	sub, err := a.contract.WatchPongReceived(&bind.WatchOpts{Context: ctx}, sink)
	if err != nil {
		return nil, fmt.Errorf("could set up channel to watch pong received: %w", err)
	}

	return sub, nil
}

func (a pingPongClientContract) transactOptsSetup(ctx context.Context, signer signer.Signer) (*bind.TransactOpts, error) {
	transactor, err := signer.GetTransactor(ctx, a.client.GetBigChainID())
	if err != nil {
		return nil, fmt.Errorf("could not sign tx: %w", err)
	}

	transactOpts, err := a.nonceManager.NewKeyedTransactor(transactor)
	if err != nil {
		return nil, fmt.Errorf("could not create tx: %w", err)
	}

	transactOpts.Context = ctx

	return transactOpts, nil
}

var _ domains.PingPongClientContract = &pingPongClientContract{}
