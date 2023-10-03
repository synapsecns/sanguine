package evm

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/synapsecns/sanguine/agents/contracts/origin"
	"github.com/synapsecns/sanguine/agents/domains"
	"github.com/synapsecns/sanguine/agents/types"
	"github.com/synapsecns/sanguine/ethergo/chain"
	"github.com/synapsecns/sanguine/ethergo/signer/nonce"
)

// NewOriginContract returns a new bound origin contract.
//
//nolint:staticcheck
func NewOriginContract(ctx context.Context, client chain.Chain, originAddress common.Address) (domains.OriginContract, error) {
	boundContract, err := origin.NewOriginRef(originAddress, client)
	if err != nil {
		return nil, fmt.Errorf("could not create %T: %w", &origin.OriginRef{}, err)
	}

	nonceManager := nonce.NewNonceManager(ctx, client, client.GetBigChainID())

	return originContract{
		contract:     boundContract,
		client:       client,
		nonceManager: nonceManager,
	}, nil
}

// originContract contains an interface for interacting with the origin chain that implements
// domains.OriginContract.
type originContract struct {
	// contract contains the contract handle
	contract *origin.OriginRef
	// client is the client
	//nolint: staticcheck
	client chain.Chain
	// nonceManager is the nonce manager used for transacting
	nonceManager nonce.Manager
}

func (o originContract) GetContractRef() *origin.OriginRef {
	return o.contract
}

func (o originContract) IsValidState(ctx context.Context, statePayload []byte) (isValid bool, err error) {
	isValid, err = o.contract.IsValidState(&bind.CallOpts{Context: ctx}, statePayload)
	if err != nil {
		return false, fmt.Errorf("could not check if state is valid: %w", err)
	}

	return isValid, nil
}

func (o originContract) SuggestLatestState(ctx context.Context) (types.State, error) {
	suggestedStateRaw, err := o.contract.SuggestLatestState(&bind.CallOpts{Context: ctx})
	if err != nil {
		return nil, fmt.Errorf("could not get suggested latest state: %w", err)
	}

	if len(suggestedStateRaw) == 0 {
		//nolint:nilnil
		return nil, nil
	}

	suggestedState, err := types.DecodeState(suggestedStateRaw)
	if err != nil {
		return nil, fmt.Errorf("could not decode suggested state: %w", err)
	}

	return suggestedState, nil
}

func (o originContract) SuggestState(ctx context.Context, nonce uint32) (types.State, error) {
	suggestedStateRaw, err := o.contract.SuggestState(&bind.CallOpts{Context: ctx}, nonce)
	if err != nil {
		return nil, fmt.Errorf("could not get suggested state: %w", err)
	}

	if len(suggestedStateRaw) == 0 {
		//nolint:nilnil
		return nil, nil
	}

	suggestedState, err := types.DecodeState(suggestedStateRaw)
	if err != nil {
		return nil, fmt.Errorf("could not decode suggested state: %w", err)
	}

	return suggestedState, nil
}

var _ domains.OriginContract = &originContract{}
