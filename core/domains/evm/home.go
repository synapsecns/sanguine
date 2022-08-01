package evm

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/synapsecns/sanguine/core/contracts/home"
	"github.com/synapsecns/sanguine/core/domains"
	"github.com/synapsecns/sanguine/core/types"
	"github.com/synapsecns/sanguine/ethergo/signer/nonce"
	"github.com/synapsecns/synapse-node/pkg/evm"
	"math/big"
)

// NewHomeContract returns a new bound home contract.
func NewHomeContract(ctx context.Context, client evm.Chain, homeAddress common.Address) (domains.HomeContract, error) {
	boundContract, err := home.NewHomeRef(homeAddress, client)
	if err != nil {
		return nil, fmt.Errorf("could not create %T: %w", &home.HomeRef{}, err)
	}

	nonceManager := nonce.NewNonceManager(ctx, client, client.GetBigChainID())

	return homeContract{
		contract:     boundContract,
		client:       client,
		nonceManager: nonceManager,
	}, nil
}

// homeContract contains an interface for interacting with the home chain that implements
// domains.HomeContract.
type homeContract struct {
	// contract contains the contract handle
	contract home.IHome
	// client is the client
	client evm.Chain
	// nonceManager is the nonce manager used for transacting
	nonceManager nonce.Manager
}

func (h homeContract) FetchSortedMessages(ctx context.Context, from uint32, to uint32) (messages []types.CommittedMessage, err error) {
	rangeFilter := NewRangeFilter(h.contract.Address(), h.client, big.NewInt(int64(from)), big.NewInt(int64(to)), 100, false)

	// blocks until done `
	err = rangeFilter.Start(ctx)
	if err != nil {
		return []types.CommittedMessage{}, fmt.Errorf("could not filter: %w", err)
	}

	filteredLogs, err := rangeFilter.Drain(ctx)
	if err != nil {
		return []types.CommittedMessage{}, fmt.Errorf("could not drain queue: %w", err)
	}

	for _, log := range filteredLogs {
		logType, ok := h.contract.Parser().EventType(log)
		if !ok {
			continue
		}

		if logType == home.DispatchEvent {
			dispatchEvents, ok := h.contract.Parser().ParseDispatch(log)
			// TODO: this should never happen. Maybe we should return an error here?
			if !ok {
				continue
			}

			messages = append(messages, dispatchEvents)
		}
	}

	return messages, nil
}

func (h homeContract) ProduceUpdate(ctx context.Context) (types.Attestation, error) {
	suggestedUpdate, err := h.contract.SuggestUpdate(&bind.CallOpts{Context: ctx})
	if err != nil {
		return nil, fmt.Errorf("could not suggest update: %w", err)
	}

	if suggestedUpdate.Root == [32]byte{} {
		return nil, domains.ErrNoUpdate
	}

	// TODO, this can be cached
	localDomain, err := h.contract.LocalDomain(&bind.CallOpts{Context: ctx})
	if err != nil {
		return nil, fmt.Errorf("could not get local domain: %w", err)
	}

	update := types.NewAttestation(localDomain, suggestedUpdate.Nonce, suggestedUpdate.Root)

	return update, nil
}

var _ domains.HomeContract = &homeContract{}
