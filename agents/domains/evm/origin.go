package evm

import (
	"context"
	"fmt"
	"github.com/synapsecns/sanguine/ethergo/chain"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/synapsecns/sanguine/agents/contracts/origin"
	"github.com/synapsecns/sanguine/agents/domains"
	"github.com/synapsecns/sanguine/agents/types"
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
	contract origin.IOrigin
	// client is the client
	//nolint: staticcheck
	client chain.Chain
	// nonceManager is the nonce manager used for transacting
	nonceManager nonce.Manager
}

func (o originContract) FetchSortedMessages(ctx context.Context, from uint32, to uint32) (messages []types.Message, err error) {
	rangeFilter := NewRangeFilter(o.contract.Address(), o.client, big.NewInt(int64(from)), big.NewInt(int64(to)), 100, false)

	// blocks until done `
	err = rangeFilter.Start(ctx)
	if err != nil {
		return []types.Message{}, fmt.Errorf("could not filter: %w", err)
	}

	filteredLogs, err := rangeFilter.Drain(ctx)
	if err != nil {
		return []types.Message{}, fmt.Errorf("could not drain queue: %w", err)
	}

	for _, log := range filteredLogs {
		logType, ok := o.contract.Parser().EventType(log)
		if !ok {
			continue
		}

		if logType == origin.SentEvent {
			sentEvents, ok := o.contract.Parser().ParseSent(log)
			// TODO: this should never happen. Maybe we should return an error here?
			if !ok {
				continue
			}

			messages = append(messages, sentEvents)
		}
	}

	return messages, nil
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
