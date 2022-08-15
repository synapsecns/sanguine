package evm

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/synapsecns/sanguine/core/contracts/origin"
	"github.com/synapsecns/sanguine/core/domains"
	"github.com/synapsecns/sanguine/core/types"
	"github.com/synapsecns/sanguine/ethergo/signer/nonce"
	"github.com/synapsecns/synapse-node/pkg/evm"
	"math/big"
)

// NewOriginContract returns a new bound origin contract.
func NewOriginContract(ctx context.Context, client evm.Chain, originAddress common.Address) (domains.OriginContract, error) {
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
	client evm.Chain
	// nonceManager is the nonce manager used for transacting
	nonceManager nonce.Manager
}

func (h originContract) FetchSortedMessages(ctx context.Context, from uint32, to uint32) (messages []types.CommittedMessage, err error) {
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

		if logType == origin.DispatchEvent {
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

func (h originContract) ProduceAttestation(ctx context.Context) (types.Attestation, error) {
	suggestedUpdate, err := h.contract.SuggestAttestation(&bind.CallOpts{Context: ctx})
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

var _ domains.OriginContract = &originContract{}
