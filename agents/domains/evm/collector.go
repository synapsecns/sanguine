package evm

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/synapsecns/sanguine/agents/contracts/attestationcollector"
	"github.com/synapsecns/sanguine/agents/domains"
	"github.com/synapsecns/sanguine/agents/types"
	"github.com/synapsecns/sanguine/ethergo/chain"
	"github.com/synapsecns/sanguine/ethergo/signer/nonce"
	"github.com/synapsecns/sanguine/ethergo/signer/signer"
)

// NewAttestationCollectorContract returns a bound attestation collector contract.
// nolint: staticcheck
func NewAttestationCollectorContract(ctx context.Context, client chain.Chain, attestationAddress common.Address) (domains.AttestationCollectorContract, error) {
	boundCountract, err := attestationcollector.NewAttestationCollectorRef(attestationAddress, client)
	if err != nil {
		return nil, fmt.Errorf("could not create %T: %w", &attestationcollector.AttestationCollectorRef{}, err)
	}

	nonceManager := nonce.NewNonceManager(ctx, client, client.GetBigChainID())
	return attestationCollectorContract{
		contract:     boundCountract,
		client:       client,
		nonceManager: nonceManager,
	}, nil
}

type attestationCollectorContract struct {
	// contract contains the conract handle
	contract *attestationcollector.AttestationCollectorRef
	// client contains the evm client
	//nolint: staticcheck
	client chain.Chain
	// nonceManager is the nonce manager used for transacting with the chain
	nonceManager nonce.Manager
}

func (a attestationCollectorContract) AddNotary(transactOpts *bind.TransactOpts, destinationID uint32, signer signer.Signer) error {
	_, err := a.contract.AddNotary(transactOpts, destinationID, signer.Address())
	if err != nil {
		return fmt.Errorf("could not add notary: %w", err)
	}

	return nil
}

func (a attestationCollectorContract) SubmitAttestation(ctx context.Context, signer signer.Signer, attestation types.SignedAttestation) error {
	transactor, err := signer.GetTransactor(a.client.GetBigChainID())
	if err != nil {
		return fmt.Errorf("could not sign tx: %w", err)
	}

	transactOpts, err := a.nonceManager.NewKeyedTransactor(transactor)
	if err != nil {
		return fmt.Errorf("could not create tx: %w", err)
	}

	transactOpts.Context = ctx

	encodedAttestation, err := types.EncodeSignedAttestation(attestation)
	if err != nil {
		return fmt.Errorf("could not get signed attestations: %w", err)
	}

	_, err = a.contract.SubmitAttestation(transactOpts, encodedAttestation)
	if err != nil {
		return fmt.Errorf("could not submit attestation: %w", err)
	}

	return nil
}

func (a attestationCollectorContract) GetLatestNonce(ctx context.Context, origin uint32, destination uint32, signer signer.Signer) (nonce uint32, currBlockNumber uint64, err error) {
	latestNonce, currBlock, err := a.contract.AttestationCollector.GetLatestNonce(&bind.CallOpts{Context: ctx}, origin, destination, signer.Address())
	if err != nil {
		return 0, uint64(0), fmt.Errorf("could not retrieve latest nonce: %w", err)
	}

	return latestNonce, currBlock.Uint64(), nil
}
