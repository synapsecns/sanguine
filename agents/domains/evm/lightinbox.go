package evm

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/agents/contracts/lightinbox"
	"github.com/synapsecns/sanguine/agents/domains"
	"github.com/synapsecns/sanguine/agents/types"
	"github.com/synapsecns/sanguine/ethergo/chain"
	"github.com/synapsecns/sanguine/ethergo/signer/nonce"
	"github.com/synapsecns/sanguine/ethergo/signer/signer"
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

//nolint:dupl
func (a lightInboxContract) SubmitStateReportWithSnapshot(transactor *bind.TransactOpts, stateIndex uint8, signature signer.Signature, snapPayload []byte, snapSignature []byte) (tx *ethTypes.Transaction, err error) {
	rawSig, err := types.EncodeSignature(signature)
	if err != nil {
		return nil, fmt.Errorf("could not encode signature: %w", err)
	}

	// TODO: Is there a way to get a return value from a contractTransactor call?
	tx, err = a.contract.SubmitStateReportWithSnapshot(transactor, stateIndex, rawSig, snapPayload, snapSignature)
	if err != nil {
		return nil, fmt.Errorf("could not submit state report: %w", err)
	}

	return tx, nil
}

func (a lightInboxContract) VerifyStateWithSnapshot(transactor *bind.TransactOpts, stateIndex uint8, snapPayload []byte, snapSignature []byte) (tx *ethTypes.Transaction, err error) {
	tx, err = a.contract.VerifyStateWithSnapshot(transactor, stateIndex, snapPayload, snapSignature)
	if err != nil {
		return nil, fmt.Errorf("could not submit state report: %w", err)
	}

	return tx, nil
}

func (a lightInboxContract) SubmitAttestationReport(transactor *bind.TransactOpts, attestation, arSignature, attSignature []byte) (tx *ethTypes.Transaction, err error) {
	tx, err = a.contract.SubmitAttestationReport(transactor, attestation, arSignature, attSignature)
	if err != nil {
		return nil, fmt.Errorf("could not submit state report: %w", err)
	}

	return tx, nil
}

func (a lightInboxContract) VerifyStateWithAttestation(transactor *bind.TransactOpts, stateIndex uint8, snapPayload []byte, attPayload []byte, attSignature []byte) (tx *ethTypes.Transaction, err error) {
	tx, err = a.contract.VerifyStateWithAttestation(transactor, stateIndex, snapPayload, attPayload, attSignature)
	if err != nil {
		return nil, fmt.Errorf("could not submit state report: %w", err)
	}

	return tx, nil
}

func (a lightInboxContract) VerifyReceipt(transactor *bind.TransactOpts, rcptPayload []byte, rcptSignature []byte) (tx *ethTypes.Transaction, err error) {
	tx, err = a.contract.VerifyReceipt(transactor, rcptPayload, rcptSignature)
	if err != nil {
		return nil, fmt.Errorf("could not submit state report: %w", err)
	}

	return tx, nil
}

//nolint:dupl
func (a lightInboxContract) SubmitStateReportWithAttestation(transactor *bind.TransactOpts, stateIndex uint8, signature signer.Signature, snapPayload, attPayload, attSignature []byte) (tx *ethTypes.Transaction, err error) {
	rawSig, err := types.EncodeSignature(signature)
	if err != nil {
		return nil, fmt.Errorf("could not encode signature: %w", err)
	}

	tx, err = a.contract.SubmitStateReportWithAttestation(transactor, stateIndex, rawSig, snapPayload, attPayload, attSignature)
	if err != nil {
		return nil, fmt.Errorf("could not submit state report: %w", err)
	}

	return tx, nil
}
