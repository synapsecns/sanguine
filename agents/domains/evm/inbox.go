package evm

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/agents/contracts/inbox"
	"github.com/synapsecns/sanguine/agents/domains"
	"github.com/synapsecns/sanguine/agents/types"
	"github.com/synapsecns/sanguine/ethergo/chain"
	"github.com/synapsecns/sanguine/ethergo/signer/nonce"
	"github.com/synapsecns/sanguine/ethergo/signer/signer"
)

// NewInboxContract returns a bound inbox contract.
//
//nolint:staticcheck
func NewInboxContract(ctx context.Context, client chain.Chain, inboxAddress common.Address) (domains.InboxContract, error) {
	boundCountract, err := inbox.NewInboxRef(inboxAddress, client)
	if err != nil {
		return nil, fmt.Errorf("could not create %T: %w", &inbox.InboxRef{}, err)
	}

	nonceManager := nonce.NewNonceManager(ctx, client, client.GetBigChainID())
	return inboxContract{
		contract:     boundCountract,
		client:       client,
		nonceManager: nonceManager,
	}, nil
}

type inboxContract struct {
	lightInboxContract
	// contract contains the conract handle
	contract *inbox.InboxRef
	// client contains the evm client
	//nolint: staticcheck
	client chain.Chain
	// nonceManager is the nonce manager used for transacting with the chain
	nonceManager nonce.Manager
}

//nolint:dupl
func (a inboxContract) SubmitStateReportWithSnapshot(transactor *bind.TransactOpts, stateIndex uint8, signature signer.Signature, snapPayload []byte, snapSignature []byte) (tx *ethTypes.Transaction, err error) {
	rawSig, err := types.EncodeSignature(signature)
	if err != nil {
		return nil, fmt.Errorf("could not encode signature: %w", err)
	}

	tx, err = a.contract.SubmitStateReportWithSnapshot(transactor, stateIndex, rawSig, snapPayload, snapSignature)
	if err != nil {
		return nil, fmt.Errorf("could not submit state report: %w", err)
	}

	return tx, nil
}

func (a inboxContract) SubmitSnapshot(transactor *bind.TransactOpts, encodedSnapshot []byte, signature signer.Signature) (tx *ethTypes.Transaction, err error) {
	rawSig, err := types.EncodeSignature(signature)
	if err != nil {
		return nil, fmt.Errorf("could not encode signature: %w", err)
	}

	tx, err = a.contract.SubmitSnapshot(transactor, encodedSnapshot, rawSig)
	if err != nil {
		return nil, fmt.Errorf("could not submit sanpshot: %w", err)
	}

	return tx, nil
}

func (a inboxContract) VerifyAttestation(transactor *bind.TransactOpts, attestation []byte, attSignature []byte) (tx *ethTypes.Transaction, err error) {
	tx, err = a.contract.VerifyAttestation(transactor, attestation, attSignature)
	if err != nil {
		return nil, fmt.Errorf("could not submit attestation: %w", err)
	}

	return tx, nil
}

func (a inboxContract) SubmitStateReportWithAttestation(transactor *bind.TransactOpts, stateIndex uint8, signature signer.Signature, snapPayload, attPayload, attSignature []byte) (tx *ethTypes.Transaction, err error) {
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

func (a inboxContract) SubmitReceipt(transactor *bind.TransactOpts, rcptPayload []byte, rcptSignature signer.Signature, paddedTips *big.Int, headerHash [32]byte, bodyHash [32]byte) (tx *ethTypes.Transaction, err error) {
	rawSig, err := types.EncodeSignature(rcptSignature)
	if err != nil {
		return nil, fmt.Errorf("could not encode signature: %w", err)
	}

	tx, err = a.contract.SubmitReceipt(transactor, rcptPayload, rawSig, paddedTips, headerHash, bodyHash)
	if err != nil {
		return nil, fmt.Errorf("could not submit state report: %w", err)
	}

	return tx, nil
}

func (a inboxContract) VerifyReceipt(transactor *bind.TransactOpts, rcptPayload []byte, rcptSignature []byte) (tx *ethTypes.Transaction, err error) {
	tx, err = a.contract.VerifyReceipt(transactor, rcptPayload, rcptSignature)
	if err != nil {
		return nil, fmt.Errorf("could not submit state report: %w", err)
	}

	return tx, nil
}

func (a inboxContract) SubmitReceiptReport(transactor *bind.TransactOpts, rcptPayload []byte, rcptSignature []byte, rrSignature []byte) (tx *ethTypes.Transaction, err error) {
	tx, err = a.contract.SubmitReceiptReport(transactor, rcptPayload, rcptSignature, rrSignature)
	if err != nil {
		return nil, fmt.Errorf("could not submit receipt report: %w", err)
	}

	return tx, nil
}
