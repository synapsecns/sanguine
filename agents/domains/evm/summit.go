package evm

import (
	"context"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/event"
	"github.com/synapsecns/sanguine/agents/contracts/summit"
	"github.com/synapsecns/sanguine/agents/domains"
	"github.com/synapsecns/sanguine/agents/types"
	"github.com/synapsecns/sanguine/ethergo/chain"
	"github.com/synapsecns/sanguine/ethergo/signer/nonce"
	"github.com/synapsecns/sanguine/ethergo/signer/signer"
)

// NewSummitContract returns a bound summit contract.
//
//nolint:staticcheck
func NewSummitContract(ctx context.Context, client chain.Chain, summitAddress common.Address) (domains.SummitContract, error) {
	boundCountract, err := summit.NewSummitRef(summitAddress, client)
	if err != nil {
		return nil, fmt.Errorf("could not create %T: %w", &summit.SummitRef{}, err)
	}

	nonceManager := nonce.NewNonceManager(ctx, client, client.GetBigChainID())
	return summitContract{
		contract:     boundCountract,
		client:       client,
		nonceManager: nonceManager,
	}, nil
}

type summitContract struct {
	// contract contains the conract handle
	contract *summit.SummitRef
	// client contains the evm client
	//nolint: staticcheck
	client chain.Chain
	// nonceManager is the nonce manager used for transacting with the chain
	nonceManager nonce.Manager
}

func (a summitContract) AddAgent(transactOpts *bind.TransactOpts, domainID uint32, signer signer.Signer) error {
	_, err := a.contract.AddAgent(transactOpts, domainID, signer.Address())
	if err != nil {
		return fmt.Errorf("could not add notary: %w", err)
	}

	return nil
}

func (a summitContract) SubmitSnapshot(ctx context.Context, signer signer.Signer, encodedSnapshot []byte, signature signer.Signature) error {
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

func (a summitContract) GetLatestState(ctx context.Context, origin uint32) (types.State, error) {
	rawState, err := a.contract.GetLatestState(&bind.CallOpts{Context: ctx}, origin)
	if err != nil {
		return nil, fmt.Errorf("could not retrieve latest state: %w", err)
	}

	state, err := types.DecodeState(rawState)
	if err != nil {
		return nil, fmt.Errorf("could not decode state: %w", err)
	}

	return state, nil
}

func (a summitContract) GetLatestAgentState(ctx context.Context, origin uint32, bondedAgentSigner signer.Signer) (types.State, error) {
	rawState, err := a.contract.GetLatestAgentState(&bind.CallOpts{Context: ctx}, origin, bondedAgentSigner.Address())
	if err != nil {
		return nil, fmt.Errorf("could not retrieve latest agent state: %w", err)
	}

	state, err := types.DecodeState(rawState)
	if err != nil {
		return nil, fmt.Errorf("could not decode state: %w", err)
	}

	return state, nil
}

func (a summitContract) WatchAttestationSaved(ctx context.Context, sink chan<- *summit.SummitAttestationSaved) (event.Subscription, error) {
	sub, err := a.contract.WatchAttestationSaved(&bind.WatchOpts{Context: ctx}, sink)
	if err != nil {
		return nil, fmt.Errorf("could set up channel to watch attestation saved: %w", err)
	}

	return sub, nil
}
