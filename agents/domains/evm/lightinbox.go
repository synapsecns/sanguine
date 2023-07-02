package evm

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
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

//nolint:dupl
func (a lightInboxContract) transactOptsSetup(ctx context.Context, signer signer.Signer) (*bind.TransactOpts, error) {
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

func (a lightInboxContract) SubmitAttestation(
	ctx context.Context,
	signer signer.Signer,
	attPayload []byte,
	signature types.Signature,
	agentRoot [32]byte,
	snapGas []*big.Int,
) error {
	transactOpts, err := a.transactOptsSetup(ctx, signer)
	if err != nil {
		return fmt.Errorf("could not setup transact opts: %w", err)
	}

	transactOpts.GasLimit = uint64(10000000)

	rawSig, err := types.EncodeSignature(signature)
	if err != nil {
		return fmt.Errorf("could not encode signature: %w", err)
	}

	_, err = a.contract.SubmitAttestation(transactOpts, attPayload, rawSig, agentRoot, snapGas)
	if err != nil {
		return fmt.Errorf("could not submit attestation: %w", err)
	}

	return nil
}

func (a lightInboxContract) VerifyStateWithSnapshotProof(
	ctx context.Context,
	signer signer.Signer,
	index uint64,
	state types.State,
	snapProof [][]byte,
	attPayload []byte,
	attSignature types.Signature,
) error {
	transactOpts, err := a.transactOptsSetup(ctx, signer)
	if err != nil {
		return fmt.Errorf("could not setup transact opts: %w", err)
	}

	transactOpts.GasLimit = uint64(10000000)

	rawAttSig, err := types.EncodeSignature(attSignature)
	if err != nil {
		return fmt.Errorf("could not encode signature: %w", err)
	}

	indexBigInt := new(big.Int).SetUint64(index)

	rawState, err := types.EncodeState(state)
	if err != nil {
		return fmt.Errorf("could not econde state: %w", err)
	}

	snapProofEVM := make([][32]byte, len(snapProof))
	for i, snapProofNode := range snapProof {
		copy(snapProofEVM[i][:], snapProofNode[:])
	}

	_, err = a.contract.VerifyStateWithSnapshotProof(transactOpts, indexBigInt, rawState, snapProofEVM, attPayload, rawAttSig)
	if err != nil {
		return fmt.Errorf("could not call VerifyStateWithSnapshotProof: %w", err)
	}

	return nil
}
