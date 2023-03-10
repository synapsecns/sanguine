package evm

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	"github.com/ethereum/go-ethereum/common"
	"github.com/synapsecns/sanguine/agents/contracts/destination"
	"github.com/synapsecns/sanguine/agents/domains"
	"github.com/synapsecns/sanguine/agents/types"
	"github.com/synapsecns/sanguine/ethergo/chain"
	"github.com/synapsecns/sanguine/ethergo/signer/nonce"
	"github.com/synapsecns/sanguine/ethergo/signer/signer"
)

// NewDestinationContract returns a bound destination contract.
//
//nolint:staticcheck
func NewDestinationContract(ctx context.Context, client chain.Chain, destinationAddress common.Address) (domains.DestinationContract, error) {
	boundCountract, err := destination.NewDestinationRef(destinationAddress, client)
	if err != nil {
		return nil, fmt.Errorf("could not create %T: %w", &destination.DestinationRef{}, err)
	}

	nonceManager := nonce.NewNonceManager(ctx, client, client.GetBigChainID())
	return destinationContract{
		contract:     boundCountract,
		client:       client,
		nonceManager: nonceManager,
	}, nil
}

type destinationContract struct {
	// contract contains the conract handle
	contract *destination.DestinationRef
	// client contains the evm client
	//nolint: staticcheck
	client chain.Chain
	// nonceManager is the nonce manager used for transacting with the chain
	nonceManager nonce.Manager
}

// func (a destinationContract) SubmitAttestation(ctx context.Context, signer signer.Signer, attestation types.SignedAttestation) error {
//	// TODO (joeallen): FIX ME
//	/*transactOpts, err := a.transactOptsSetup(ctx, signer)
//	if err != nil {
//		return fmt.Errorf("could not setup transact opts: %w", err)
//	}
//
//	encodedAttestation, err := types.EncodeSignedAttestation(attestation)
//	if err != nil {
//		return fmt.Errorf("could not get signed attestations: %w", err)
//	}
//
//	_, err = a.contract.SubmitAttestation(transactOpts, encodedAttestation)
//	if err != nil {
//		return fmt.Errorf("could not submit attestation: %w", err)
//	}*/
//
//	return nil
//}

func (a destinationContract) Execute(ctx context.Context, signer signer.Signer, message types.Message, proof [32][32]byte, index *big.Int) error {
	transactOpts, err := a.transactOptsSetup(ctx, signer)
	if err != nil {
		return fmt.Errorf("could not setup transact opts: %w", err)
	}

	encodedMessage, err := types.EncodeMessage(message)
	if err != nil {
		return fmt.Errorf("could not encode message: %w", err)
	}

	// TODO (joeallen): FIX ME
	_, err = a.contract.Execute(transactOpts, encodedMessage, proof, [][32]byte{}, index)
	if err != nil {
		return fmt.Errorf("could not execute message: %w", err)
	}

	return nil
}

func (a destinationContract) transactOptsSetup(ctx context.Context, signer signer.Signer) (*bind.TransactOpts, error) {
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

func (a destinationContract) AttestationsAmount(ctx context.Context) (uint64, error) {
	attestationsAmountBigInt, err := a.contract.AttestationsAmount(&bind.CallOpts{Context: ctx})
	if err != nil {
		return uint64(0), fmt.Errorf("could not get submitted at for origin and root: %w", err)
	}

	if attestationsAmountBigInt == nil {
		return uint64(0), nil
	}

	return attestationsAmountBigInt.Uint64(), nil
}

func (a destinationContract) SubmitAttestation(ctx context.Context, signer signer.Signer, attPayload []byte, signature signer.Signature) error {
	transactOpts, err := a.transactOptsSetup(ctx, signer)
	if err != nil {
		return fmt.Errorf("could not setup transact opts: %w", err)
	}

	rawSig, err := types.EncodeSignature(signature)
	if err != nil {
		return fmt.Errorf("could not encode signature: %w", err)
	}

	_, err = a.contract.SubmitAttestation(transactOpts, attPayload, rawSig)
	if err != nil {
		return fmt.Errorf("could not submit attestation: %w", err)
	}

	return nil
}
