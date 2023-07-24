package evm

import (
	"context"
	"fmt"
	"math/big"
	"strings"

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

func (a inboxContract) SubmitStateReportWithSnapshot(ctx context.Context, signer signer.Signer, stateIndex int64, signature signer.Signature, snapPayload []byte, snapSignature []byte) (tx *ethTypes.Transaction, err error) {
	transactor, err := signer.GetTransactor(ctx, a.client.GetBigChainID())
	if err != nil {
		return nil, fmt.Errorf("could not sign tx: %w", err)
	}

	transactOpts, err := a.nonceManager.NewKeyedTransactor(transactor)
	if err != nil {
		return nil, fmt.Errorf("could not create tx: %w", err)
	}

	transactOpts.Context = ctx

	transactOpts.GasLimit = 5000000

	rawSig, err := types.EncodeSignature(signature)
	if err != nil {
		return nil, fmt.Errorf("could not encode signature: %w", err)
	}

	// TODO: Is there a way to get a return value from a contractTransactor call?
	tx, err = a.contract.SubmitStateReportWithSnapshot(transactOpts, big.NewInt(stateIndex), rawSig, snapPayload, snapSignature)
	if err != nil {
		// TODO: Why is this done? And if it is necessary, we should functionalize it.
		if strings.Contains(err.Error(), "nonce too low") {
			a.nonceManager.ClearNonce(signer.Address())
		}
		return nil, fmt.Errorf("could not submit state report: %w", err)
	}

	return tx, nil
}

func (a inboxContract) SubmitSnapshot(transactor *bind.TransactOpts, signer signer.Signer, encodedSnapshot []byte, signature signer.Signature) (tx *ethTypes.Transaction, err error) {
	rawSig, err := types.EncodeSignature(signature)
	if err != nil {
		return nil, fmt.Errorf("could not encode signature: %w", err)
	}

	tx, err = a.contract.SubmitSnapshot(transactor, encodedSnapshot, rawSig)
	if err != nil {
		if strings.Contains(err.Error(), "nonce too low") {
			a.nonceManager.ClearNonce(signer.Address())
		}
		return nil, fmt.Errorf("could not submit sanpshot: %w", err)
	}

	return tx, nil
}

func (a inboxContract) SubmitSnapshotCtx(ctx context.Context, signer signer.Signer, encodedSnapshot []byte, signature signer.Signature) (tx *ethTypes.Transaction, err error) {
	transactor, err := signer.GetTransactor(ctx, a.client.GetBigChainID())
	if err != nil {
		return nil, fmt.Errorf("could not sign tx: %w", err)
	}

	transactOpts, err := a.nonceManager.NewKeyedTransactor(transactor)
	if err != nil {
		return nil, fmt.Errorf("could not create tx: %w", err)
	}

	transactOpts.Context = ctx

	transactOpts.GasLimit = 5000000

	rawSig, err := types.EncodeSignature(signature)
	if err != nil {
		return nil, fmt.Errorf("could not encode signature: %w", err)
	}

	tx, err = a.contract.SubmitSnapshot(transactOpts, encodedSnapshot, rawSig)
	if err != nil {
		if strings.Contains(err.Error(), "nonce too low") {
			a.nonceManager.ClearNonce(signer.Address())
		}
		return nil, fmt.Errorf("could not submit sanpshot: %w", err)
	}

	return tx, nil
}

func (a inboxContract) VerifyAttestation(ctx context.Context, signer signer.Signer, attestation []byte, attSignature []byte) (tx *ethTypes.Transaction, err error) {
	transactor, err := signer.GetTransactor(ctx, a.client.GetBigChainID())
	if err != nil {
		return nil, fmt.Errorf("could not sign tx: %w", err)
	}

	transactOpts, err := a.nonceManager.NewKeyedTransactor(transactor)
	if err != nil {
		return nil, fmt.Errorf("could not create tx: %w", err)
	}

	transactOpts.Context = ctx
	return a.contract.VerifyAttestation(transactOpts, attestation, attSignature)
}

func (a inboxContract) VerifyStateWithAttestation(ctx context.Context, signer signer.Signer, stateIndex int64, snapPayload []byte, attPayload []byte, attSignature []byte) (tx *ethTypes.Transaction, err error) {
	transactor, err := signer.GetTransactor(ctx, a.client.GetBigChainID())
	if err != nil {
		return nil, fmt.Errorf("could not sign tx: %w", err)
	}

	transactOpts, err := a.nonceManager.NewKeyedTransactor(transactor)
	if err != nil {
		return nil, fmt.Errorf("could not create tx: %w", err)
	}

	transactOpts.Context = ctx
	transactOpts.GasLimit = 5000000
	return a.contract.VerifyStateWithAttestation(transactOpts, big.NewInt(stateIndex), snapPayload, attPayload, attSignature)
}

func (a inboxContract) SubmitStateReportWithAttestation(ctx context.Context, signer signer.Signer, stateIndex int64, signature signer.Signature, snapPayload, attPayload, attSignature []byte) (tx *ethTypes.Transaction, err error) {
	transactor, err := signer.GetTransactor(ctx, a.client.GetBigChainID())
	if err != nil {
		return nil, fmt.Errorf("could not sign tx: %w", err)
	}

	transactOpts, err := a.nonceManager.NewKeyedTransactor(transactor)
	if err != nil {
		return nil, fmt.Errorf("could not create tx: %w", err)
	}

	transactOpts.Context = ctx

	transactOpts.GasLimit = 5000000

	rawSig, err := types.EncodeSignature(signature)
	if err != nil {
		return nil, fmt.Errorf("could not encode signature: %w", err)
	}

	// TODO: Is there a way to get a return value from a contractTransactor call?
	tx, err = a.contract.SubmitStateReportWithAttestation(transactOpts, big.NewInt(stateIndex), rawSig, snapPayload, attPayload, attSignature)
	if err != nil {
		// TODO: Why is this done? And if it is necessary, we should functionalize it.
		if strings.Contains(err.Error(), "nonce too low") {
			a.nonceManager.ClearNonce(signer.Address())
		}
		return nil, fmt.Errorf("could not submit state report: %w", err)
	}

	return tx, nil
}

func (a inboxContract) SubmitReceipt(ctx context.Context, signer signer.Signer, rcptPayload []byte, rcptSignature signer.Signature, paddedTips *big.Int, headerHash [32]byte, bodyHash [32]byte) (tx *ethTypes.Transaction, err error) {
	transactor, err := signer.GetTransactor(ctx, a.client.GetBigChainID())
	if err != nil {
		return nil, fmt.Errorf("could not sign tx: %w", err)
	}

	transactOpts, err := a.nonceManager.NewKeyedTransactor(transactor)
	if err != nil {
		return nil, fmt.Errorf("could not create tx: %w", err)
	}

	transactOpts.Context = ctx

	transactOpts.GasLimit = 5000000

	rawSig, err := types.EncodeSignature(rcptSignature)
	if err != nil {
		return nil, fmt.Errorf("could not encode signature: %w", err)
	}

	// TODO: Is there a way to get a return value from a contractTransactor call?
	tx, err = a.contract.SubmitReceipt(transactOpts, rcptPayload, rawSig, paddedTips, headerHash, bodyHash)
	if err != nil {
		// TODO: Why is this done? And if it is necessary, we should functionalize it.
		if strings.Contains(err.Error(), "nonce too low") {
			a.nonceManager.ClearNonce(signer.Address())
		}
		return nil, fmt.Errorf("could not submit state report: %w", err)
	}

	return tx, nil
}

func (a inboxContract) VerifyReceipt(ctx context.Context, signer signer.Signer, rcptPayload []byte, rcptSignature []byte) (tx *ethTypes.Transaction, err error) {
	transactor, err := signer.GetTransactor(ctx, a.client.GetBigChainID())
	if err != nil {
		return nil, fmt.Errorf("could not sign tx: %w", err)
	}

	transactOpts, err := a.nonceManager.NewKeyedTransactor(transactor)
	if err != nil {
		return nil, fmt.Errorf("could not create tx: %w", err)
	}

	transactOpts.Context = ctx
	transactOpts.GasLimit = 5000000
	return a.contract.VerifyReceipt(transactOpts, rcptPayload, rcptSignature)
}

func (a inboxContract) SubmitReceiptReport(ctx context.Context, signer signer.Signer, rcptPayload []byte, rcptSignature []byte, rrSignature []byte) (tx *ethTypes.Transaction, err error) {
	transactor, err := signer.GetTransactor(ctx, a.client.GetBigChainID())
	if err != nil {
		return nil, fmt.Errorf("could not sign tx: %w", err)
	}

	transactOpts, err := a.nonceManager.NewKeyedTransactor(transactor)
	if err != nil {
		return nil, fmt.Errorf("could not create tx: %w", err)
	}

	transactOpts.Context = ctx

	transactOpts.GasLimit = 5000000

	// TODO: Is there a way to get a return value from a contractTransactor call?
	tx, err = a.contract.SubmitReceiptReport(transactOpts, rcptPayload, rcptSignature, rrSignature)
	if err != nil {
		// TODO: Why is this done? And if it is necessary, we should functionalize it.
		if strings.Contains(err.Error(), "nonce too low") {
			a.nonceManager.ClearNonce(signer.Address())
		}
		return nil, fmt.Errorf("could not submit state report: %w", err)
	}

	return tx, nil
}
