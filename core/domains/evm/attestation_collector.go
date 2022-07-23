package evm

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/synapsecns/sanguine/core/contracts/attestationcollector"
	"github.com/synapsecns/sanguine/core/domains"
	"github.com/synapsecns/sanguine/ethergo/signer/nonce"
	"github.com/synapsecns/synapse-node/pkg/evm"
)

type attestationCollector struct {
	// contract is the
	contract *attestationcollector.AttestationCollectorRef
	// client is the client used for interacting with the chain
	client evm.Chain
	// nonceManager is the nonce manager used for transacting
	// TODO: switch w/ db transactor
	nonceManager nonce.Manager
}

// NewAttestationCollector creates a handle for posting attestations.
func NewAttestationCollector(ctx context.Context, client evm.Chain, attestationCollectorAddress common.Address) (domains.AttestationCollector, error) {
	boundContract, err := attestationcollector.NewAttestationCollectorRef(attestationCollectorAddress, client)
	if err != nil {
		return nil, fmt.Errorf("could not get bound attestation contract to %T: %w", attestationcollector.AttestationCollectorRef{}, err)
	}

	nonceManager := nonce.NewNonceManager(ctx, client, client.GetBigChainID())

	return attestationCollector{
		contract:     boundContract,
		client:       client,
		nonceManager: nonceManager,
	}, nil
}

func (a *attestationCollector) SubmitAttestation(updater common.Hash, attestation []byte) {
	
}

var _ domains.AttestationCollector = &attestationCollector{}
