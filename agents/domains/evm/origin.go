package evm

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/synapsecns/sanguine/agents/contracts/origin"
	"github.com/synapsecns/sanguine/agents/domains"
	"github.com/synapsecns/sanguine/agents/types"
	"github.com/synapsecns/sanguine/ethergo/chain"
	"github.com/synapsecns/sanguine/ethergo/signer/nonce"
)

// NewOriginContract returns a new bound origin contract.
// nolint: staticcheck
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
	// TODO: After origin.go inherits GlobalNotaryRegistry, we can implement suggestAttestations
	// and change this to ProduceAttestations returning a slice []types.Attestation
	return nil, domains.ErrNoUpdate
	/*suggestedUpdate, err := h.contract.SuggestAttestation(&bind.CallOpts{Context: ctx}, 0)
	if err != nil {
		return nil, fmt.Errorf("could not suggest update: %w", err)
	}

	if suggestedUpdate.LatestRoot == [32]byte{} {
		return nil, domains.ErrNoUpdate
	}

	// TODO (joe), this can be cached
	localDomain, err := h.contract.LocalDomain(&bind.CallOpts{Context: ctx})
	if err != nil {
		return nil, fmt.Errorf("could not get local domain: %w", err)
	}

	update := types.NewAttestation(localDomain, suggestedUpdate.LatestNonce, suggestedUpdate.LatestRoot)

	return update, nil*/
}

var _ domains.OriginContract = &originContract{}
