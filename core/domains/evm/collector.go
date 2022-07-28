package evm

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/synapsecns/sanguine/core/contracts/attestationcollector"
	"github.com/synapsecns/sanguine/core/domains"
	"github.com/synapsecns/sanguine/ethergo/signer/nonce"
	"github.com/synapsecns/sanguine/ethergo/signer/signer"
	"github.com/synapsecns/synapse-node/pkg/evm"
)

// NewAttestationCollectorContract returns a bound attestation collector contract
func NewAttestationCollectorContract(ctx context.Context, client evm.Chain, attestationAddress common.Address) (domains.AttestationCollectorContract, error) {
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
	client evm.Chain
	// nonceManager is the nonce manager used for transacting with the chain
	nonceManager nonce.Manager
}

// SubmitAttestation submits an attestation to the attestation collector
func (a attestationCollectorContract) SubmitAttestation(signer signer.Signer, updater common.Hash, attestation []byte) error {
	transactor, err := signer.GetTransactor(a.client.GetBigChainID())
	if err != nil {
		return fmt.Errorf("could not sign tx: %w", err)
	}

	transactOpts, err := a.nonceManager.NewKeyedTransactor(transactor)
	if err != nil {
		return fmt.Errorf("could not create tx: %w", err)
	}

	_, err = a.contract.SubmitAttestation(transactOpts, common.HexToAddress(updater.Hex()), attestation)
	if err != nil {
		return fmt.Errorf("could not submit attestation: %w", err)
	}

	return nil
}
