package evm

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/synapsecns/sanguine/agents/contracts/test/testclient"
	"github.com/synapsecns/sanguine/agents/domains"
	"github.com/synapsecns/sanguine/ethergo/chain"
	"github.com/synapsecns/sanguine/ethergo/signer/nonce"
	"github.com/synapsecns/sanguine/ethergo/signer/signer"
)

// NewTestClientContract returns a bound test client contract.
//
//nolint:staticcheck
func NewTestClientContract(ctx context.Context, client chain.Chain, testClientAddress common.Address) (domains.TestClientContract, error) {
	boundCountract, err := testclient.NewTestClientRef(testClientAddress, client)
	if err != nil {
		return nil, fmt.Errorf("could not create %T: %w", &testclient.TestClientRef{}, err)
	}

	nonceManager := nonce.NewNonceManager(ctx, client, client.GetBigChainID())
	return testClientContract{
		contract:     boundCountract,
		client:       client,
		nonceManager: nonceManager,
	}, nil
}

type testClientContract struct {
	// contract contains the conract handle
	contract *testclient.TestClientRef
	// client contains the evm client
	//nolint: staticcheck
	client chain.Chain
	// nonceManager is the nonce manager used for transacting with the chain
	nonceManager nonce.Manager
}

func (a testClientContract) SendMessage(ctx context.Context, signer signer.Signer, destination uint32, recipient common.Address, optimisticSeconds uint32, gasLimit uint64, version uint32, message []byte) error {
	transactOpts, err := a.transactOptsSetup(ctx, signer)
	if err != nil {
		return fmt.Errorf("could not setup transact opts: %w", err)
	}

	_, err = a.contract.SendMessage(transactOpts, destination, recipient, optimisticSeconds, gasLimit, version, message)
	if err != nil {
		return fmt.Errorf("could not send message: %w", err)
	}

	return nil
}

func (a testClientContract) transactOptsSetup(ctx context.Context, signer signer.Signer) (*bind.TransactOpts, error) {
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

var _ domains.TestClientContract = &testClientContract{}
