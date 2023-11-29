package evm

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/ethergo/chain"

	"github.com/ethereum/go-ethereum/common"
	"github.com/synapsecns/sanguine/agents/contracts/destination"
	"github.com/synapsecns/sanguine/agents/domains"
	"github.com/synapsecns/sanguine/agents/types"
	"github.com/synapsecns/sanguine/ethergo/signer/nonce"
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

func (a destinationContract) Execute(transactor *bind.TransactOpts, message types.Message, originProof [32][32]byte, snapshotProof [][32]byte, index uint8, gasLimit uint64) (tx *ethTypes.Transaction, err error) {
	encodedMessage, err := types.EncodeMessage(message)
	if err != nil {
		return nil, fmt.Errorf("could not encode message: %w", err)
	}

	tx, err = a.contract.Execute(transactor, encodedMessage, originProof[:], snapshotProof, index, gasLimit)
	if err != nil {
		return nil, fmt.Errorf("could not execute message: %w", err)
	}

	return tx, nil
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

func (a destinationContract) GetAttestationNonce(ctx context.Context, snapRoot [32]byte) (uint32, error) {
	attNonce, err := a.contract.GetAttestationNonce(&bind.CallOpts{Context: ctx}, snapRoot)
	if err != nil {
		return uint32(0), fmt.Errorf("could not get attNonce for snapRoot: %w", err)
	}

	return attNonce, nil
}

func (a destinationContract) LastAttestationNonce(ctx context.Context, index uint32) (uint32, error) {
	attNonce, err := a.contract.LastAttestationNonce(&bind.CallOpts{Context: ctx}, index)
	if err != nil {
		return uint32(0), fmt.Errorf("could not get last attNonce: %w", err)
	}

	return attNonce, nil
}

func (a destinationContract) MessageStatus(ctx context.Context, message types.Message) (uint8, error) {
	messageLeaf, err := message.ToLeaf()
	if err != nil {
		return 0, fmt.Errorf("could not get message leaf: %w", err)
	}

	status, err := a.contract.MessageStatus(&bind.CallOpts{Context: ctx}, messageLeaf)
	if err != nil {
		return 0, fmt.Errorf("could not get message status: %w", err)
	}

	return status, nil
}

//nolint:wrapcheck
func (a destinationContract) IsValidReceipt(ctx context.Context, rcptPayload []byte) (bool, error) {
	return a.contract.IsValidReceipt(&bind.CallOpts{Context: ctx}, rcptPayload)
}

func (a destinationContract) PassAgentRoot(transactor *bind.TransactOpts) (*ethTypes.Transaction, error) {
	tx, err := a.contract.PassAgentRoot(transactor)
	if err != nil {
		return nil, fmt.Errorf("could not pass agent root: %w", err)
	}
	return tx, nil
}
