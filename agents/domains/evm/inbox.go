package evm

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/synapsecns/sanguine/agents/contracts/inbox"
	"github.com/synapsecns/sanguine/agents/domains"
	"github.com/synapsecns/sanguine/agents/types"
	"github.com/synapsecns/sanguine/ethergo/chain"
	"github.com/synapsecns/sanguine/ethergo/signer/nonce"
	"github.com/synapsecns/sanguine/ethergo/signer/signer"
	"strings"
)

// NewInboxContract returns a bound inbox contract.
//
//nolint:staticcheck
func NewInboxContract(ctx context.Context, client chain.Chain, inboxAddress common.Address) (domains.InboxContract, error) {
	boundCountract, err := inbox.NewInboxRef(inboxAddress, client)
	if err != nil {
		return nil, fmt.Errorf("could not create %T: %w", &inbox.InboxRef{}, err)
	}

	nonceManager := nonce.NewNonceManager(ctx, client, client.GetBigChainID())
	return inboxContract{
		contract:     boundCountract,
		client:       client,
		nonceManager: nonceManager,
	}, nil
}

type inboxContract struct {
	// contract contains the conract handle
	contract *inbox.InboxRef
	// client contains the evm client
	//nolint: staticcheck
	client chain.Chain
	// nonceManager is the nonce manager used for transacting with the chain
	nonceManager nonce.Manager
}

func (a inboxContract) SubmitSnapshot(ctx context.Context, signer signer.Signer, encodedSnapshot []byte, signature signer.Signature) error {
	transactor, err := signer.GetTransactor(ctx, a.client.GetBigChainID())
	if err != nil {
		return fmt.Errorf("could not sign tx: %w", err)
	}

	transactOpts, err := a.nonceManager.NewKeyedTransactor(transactor)
	if err != nil {
		return fmt.Errorf("could not create tx: %w", err)
	}

	transactOpts.Context = ctx

	rawSig, err := types.EncodeSignature(signature)
	if err != nil {
		return fmt.Errorf("could not encode signature: %w", err)
	}
	_, err = a.contract.SubmitSnapshot(transactOpts, encodedSnapshot, rawSig)
	if err != nil {
		if strings.Contains(err.Error(), "nonce too low") {
			a.nonceManager.ClearNonce(signer.Address())
		}
		return fmt.Errorf("could not submit sanpshot: %w", err)
	}

	return nil
}
