package evm

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/synapsecns/sanguine/agents/contracts/summit"
	"github.com/synapsecns/sanguine/agents/domains"
	"github.com/synapsecns/sanguine/agents/types"
	"github.com/synapsecns/sanguine/core"
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

func (a summitContract) SubmitSnapshot(ctx context.Context, signer signer.Signer, snapshot types.Snapshot) error {
	transactor, err := signer.GetTransactor(ctx, a.client.GetBigChainID())
	if err != nil {
		return fmt.Errorf("could not sign tx: %w", err)
	}

	transactOpts, err := a.nonceManager.NewKeyedTransactor(transactor)
	if err != nil {
		return fmt.Errorf("could not create tx: %w", err)
	}

	transactOpts.Context = ctx

	encodedSnapshot, err := types.EncodeSnapshot(snapshot)
	if err != nil {
		return fmt.Errorf("could not get signed attestations: %w", err)
	}

	hashedSnapshot, err := types.HashRawBytes(encodedSnapshot)
	if err != nil {
		return fmt.Errorf("could not hash snapshot: %w", err)
	}
	signature, err := signer.SignMessage(ctx, core.BytesToSlice(hashedSnapshot), false)
	if err != nil {
		return fmt.Errorf("could not sign message: %w", err)
	}

	_, err = a.contract.SubmitSnapshot(transactOpts, encodedSnapshot, signature.R().Bytes())
	if err != nil {
		return fmt.Errorf("could not submit attestation: %w", err)
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

func (a summitContract) PrimeNonce(ctx context.Context, signer signer.Signer) error {
	_, err := a.nonceManager.GetNextNonce(signer.Address())
	if err != nil {
		return fmt.Errorf("could not prime nonce for signer on collector: %w", err)
	}
	return nil
}
