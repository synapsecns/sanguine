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
//
//nolint:staticcheck
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

func (a attestationCollectorContract) AddAgent(transactOpts *bind.TransactOpts, domainID uint32, signer signer.Signer) error {
	// TODO (joeallen): FIX ME
	//_, err := a.contract.AddAgent(transactOpts, domainID, signer.Address())
	//if err != nil {
	//	return fmt.Errorf("could not add notary: %w", err)
	//}

	return nil
}

func (a attestationCollectorContract) SubmitAttestation(ctx context.Context, signer signer.Signer, attestation types.SignedAttestation) error {
	// TODO (joeallen): FIX ME
	/*transactor, err := signer.GetTransactor(ctx, a.client.GetBigChainID())
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
	}*/

	return nil
}

func (a attestationCollectorContract) GetLatestNonce(ctx context.Context, origin uint32, destination uint32, bondedAgentSigner signer.Signer) (nonce uint32, err error) {
	// TODO (joeallen): FIX ME
	/*latestNonce, err := a.contract.GetLatestNonce(&bind.CallOpts{Context: ctx}, origin, destination, bondedAgentSigner.Address())
	if err != nil {
		return 0, fmt.Errorf("could not retrieve latest nonce: %w", err)
	}

	return latestNonce, nil*/
	return uint32(0), nil
}

func (a attestationCollectorContract) GetAttestation(ctx context.Context, origin, destination, nonce uint32) (types.SignedAttestation, error) {
	// TODO (joeallen): FIX ME
	/*rawAttestation, err := a.contract.GetAttestation(&bind.CallOpts{Context: ctx}, origin, destination, nonce)
	if err != nil {
		if err.Error() == "execution reverted: Unknown nonce" {
			return nil, domains.ErrNoUpdate
		}
		return nil, fmt.Errorf("could not retrieve attestation: %w", err)
	}

	signedAttesation, err := types.DecodeSignedAttestation(rawAttestation)
	if err != nil {
		return nil, fmt.Errorf("could not decode attestation: %w", err)
	}

	return signedAttesation, nil*/
	return nil, nil
}

func (a attestationCollectorContract) GetRoot(ctx context.Context, origin, destination, nonce uint32) ([32]byte, error) {
	// TODO (joeallen): FIX ME
	//root, err := a.contract.GetRoot(&bind.CallOpts{Context: ctx}, origin, destination, nonce)
	//if err != nil {
	//	return [32]byte{}, fmt.Errorf("could not retrieve root: %w", err)
	//}

	//return root, nil
	return [32]byte{}, nil
}

func (a attestationCollectorContract) PrimeNonce(ctx context.Context, signer signer.Signer) error {
	_, err := a.nonceManager.GetNextNonce(signer.Address())
	if err != nil {
		return fmt.Errorf("could not prime nonce for signer on collector: %w", err)
	}
	return nil
}
