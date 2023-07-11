package evm

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/agents/contracts/lightinbox"
	"github.com/synapsecns/sanguine/agents/domains"
	"github.com/synapsecns/sanguine/agents/types"
	"github.com/synapsecns/sanguine/ethergo/chain"
	"github.com/synapsecns/sanguine/ethergo/signer/nonce"
	"github.com/synapsecns/sanguine/ethergo/signer/signer"
	"math/big"
)

// NewLightInboxContract returns a bound light inbox contract.
//
//nolint:staticcheck
func NewLightInboxContract(ctx context.Context, client chain.Chain, lightInboxAddress common.Address) (domains.LightInboxContract, error) {
	boundCountract, err := lightinbox.NewLightInboxRef(lightInboxAddress, client)
	if err != nil {
		return nil, fmt.Errorf("could not create %T: %w", &lightinbox.LightInboxRef{}, err)
	}

	nonceManager := nonce.NewNonceManager(ctx, client, client.GetBigChainID())
	return lightInboxContract{
		contract:     boundCountract,
		client:       client,
		nonceManager: nonceManager,
	}, nil
}

type lightInboxContract struct {
	// contract contains the contract handle
	contract *lightinbox.LightInboxRef
	// client contains the evm client
	//nolint: staticcheck
	client chain.Chain
	// nonceManager is the nonce manager used for transacting with the chain
	nonceManager nonce.Manager
}

func (a lightInboxContract) SubmitAttestation(
	transactor *bind.TransactOpts,
	attPayload []byte,
	signature signer.Signature,
	agentRoot [32]byte,
	snapGas []*big.Int,
) (tx *ethTypes.Transaction, err error) {
	rawSig, err := types.EncodeSignature(signature)
	if err != nil {
		return nil, fmt.Errorf("could not encode signature: %w", err)
	}

	tx, err = a.contract.SubmitAttestation(transactor, attPayload, rawSig, agentRoot, snapGas)
	if err != nil {
		return nil, fmt.Errorf("could not submit attestation: %w", err)
	}

	return tx, nil
}
