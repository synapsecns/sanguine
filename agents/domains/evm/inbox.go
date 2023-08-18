package evm

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
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

func (a inboxContract) SubmitSnapshot(transactor *bind.TransactOpts, signer signer.Signer, encodedSnapshot []byte, signature signer.Signature) (tx *ethTypes.Transaction, err error) {
	rawSig, err := types.EncodeSignature(signature)
	if err != nil {
		return nil, fmt.Errorf("could not encode signature: %w", err)
	}

	tx, err = a.contract.SubmitSnapshot(transactor, encodedSnapshot, rawSig)
	if err != nil {
		if strings.Contains(err.Error(), "nonce too low") {
			a.nonceManager.ClearNonce(signer.Address())
		}
		return nil, fmt.Errorf("could not submit sanpshot: %w", err)
	}

	return tx, nil
}
