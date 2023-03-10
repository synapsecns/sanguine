package notary

import (
	"context"
	"fmt"
	"time"

	"github.com/synapsecns/sanguine/agents/config"
	"github.com/synapsecns/sanguine/agents/domains"
	"github.com/synapsecns/sanguine/agents/domains/evm"
	"github.com/synapsecns/sanguine/agents/types"
	"github.com/synapsecns/sanguine/ethergo/signer/signer"
)

// Notary checks the Summit for that latest states signed by guards, validates those states on origin,
// then signs and submits the snapshot to Summit.
type Notary struct {
	bondedSigner            signer.Signer
	unbondedSigner          signer.Signer
	domains                 []domains.DomainClient
	summitDomain            domains.DomainClient
	destinationDomain       domains.DomainClient
	refreshInterval         time.Duration
	summitMyLatestStates    map[uint32]types.State
	summitGuardLatestStates map[uint32]types.State
}

// NewNotary creates a new notary.
//
//nolint:cyclop
func NewNotary(ctx context.Context, cfg config.AgentConfig) (_ Notary, err error) {
	notary := Notary{
		refreshInterval: time.Second * time.Duration(cfg.RefreshIntervalSeconds),
	}
	notary.domains = []domains.DomainClient{}

	notary.bondedSigner, err = config.SignerFromConfig(ctx, cfg.BondedSigner)
	if err != nil {
		return Notary{}, fmt.Errorf("error with bondedSigner, could not create notary: %w", err)
	}

	notary.unbondedSigner, err = config.SignerFromConfig(ctx, cfg.UnbondedSigner)
	if err != nil {
		return Notary{}, fmt.Errorf("error with unbondedSigner, could not create notary: %w", err)
	}

	for domainName, domain := range cfg.Domains {
		domainClient, err := evm.NewEVM(ctx, domainName, domain)
		if err != nil {
			return Notary{}, fmt.Errorf("failing to create evm for domain, could not create notary for: %w", err)
		}
		notary.domains = append(notary.domains, domainClient)
		if domain.DomainID == cfg.SummitDomainID {
			notary.summitDomain = domainClient
		}
		if domain.DomainID == cfg.DomainID {
			notary.destinationDomain = domainClient
		}
	}

	notary.summitMyLatestStates = make(map[uint32]types.State, len(notary.domains))
	notary.summitGuardLatestStates = make(map[uint32]types.State, len(notary.domains))

	return notary, nil
}

func (n Notary) loadSummitMyLatestStates(ctx context.Context) {
	for _, domain := range n.domains {
		originID := domain.Config().DomainID
		myLatestState, err := domain.Summit().GetLatestAgentState(ctx, originID, n.bondedSigner)
		if err != nil {
			myLatestState = nil
			logger.Errorf("Failed calling GetLatestAgentState for originID on the Summit contract: %d, err = %v", originID, err)
		}
		if myLatestState != nil {
			n.summitMyLatestStates[originID] = myLatestState
		}
	}
}

func (n Notary) loadSummitGuardLatestStates(ctx context.Context) {
	for _, domain := range n.domains {
		originID := domain.Config().DomainID

		guardLatestState, err := domain.Summit().GetLatestState(ctx, originID)
		if err != nil {
			guardLatestState = nil
			logger.Errorf("Failed calling GetLatestState for originID %d on the Summit contract: err = %v", originID, err)
		}
		if guardLatestState != nil {
			n.summitGuardLatestStates[originID] = guardLatestState
		}
	}
}

func (n Notary) isValidOnOrigin(ctx context.Context, state types.State, domain domains.DomainClient) bool {
	stateOnOrigin, err := domain.Origin().SuggestState(ctx, state.Nonce())
	if err != nil {
		logger.Errorf("Failed calling SuggestState for originID %d on the Origin contract: err = %v", err)
		// return false since we weren't able to validate the state on the origin
		return false
	}

	if stateOnOrigin.Root() != state.Root() {
		logger.Errorf("State roots do not equal")
		return false
	}

	if stateOnOrigin.Origin() != state.Origin() {
		logger.Errorf("State origins do not equal")
		return false
	}

	if stateOnOrigin.Nonce() != state.Nonce() {
		logger.Errorf("State nonces do not equal")
		return false
	}

	if stateOnOrigin.BlockNumber() != state.BlockNumber() {
		logger.Errorf("State block numbers do not equal")
		return false
	}

	if stateOnOrigin.Timestamp() == nil || state.Timestamp() == nil {
		logger.Errorf("One or both of the state timestamps are nil")
		return false
	}

	if stateOnOrigin.Timestamp().Uint64() != state.Timestamp().Uint64() {
		logger.Errorf("State timestamps do not equal")
		return false
	}

	stateOnOriginHash, err := stateOnOrigin.Hash()
	if err != nil {
		logger.Errorf("Error computing state on origin hash")
		return false
	}

	stateHash, err := state.Hash()
	if err != nil {
		logger.Errorf("Error computing state on summit hash")
		return false
	}

	if stateOnOriginHash != stateHash {
		logger.Errorf("State hashes do not equal")
		return false
	}

	return true
}

func (n Notary) getLatestSnapshot(ctx context.Context) (types.Snapshot, map[uint32]types.State) {
	statesToSubmit := make(map[uint32]types.State, len(n.domains))
	for _, domain := range n.domains {
		originID := domain.Config().DomainID
		summitMyLatest, ok := n.summitMyLatestStates[originID]
		if !ok {
			summitMyLatest = nil
		}
		summitGuardLatest, ok := n.summitGuardLatestStates[originID]
		if !ok {
			continue
		}

		if summitMyLatest.Nonce() >= summitGuardLatest.Nonce() {
			// Here this notary already submitted this state
			continue
		}
		if !n.isValidOnOrigin(ctx, summitGuardLatest, domain) {
			logger.Errorf("State not valid on origin %d, nonce %d",
				summitGuardLatest.Origin(),
				summitGuardLatest.Nonce())
			continue
		}
		statesToSubmit[originID] = summitGuardLatest
	}
	snapshotStates := make([]types.State, 0, len(statesToSubmit))
	for _, state := range statesToSubmit {
		snapshotStates = append(snapshotStates, state)
	}
	if len(snapshotStates) > 0 {
		return types.NewSnapshot(snapshotStates), statesToSubmit
	}
	//nolint:nilnil
	return nil, nil
}

func (n Notary) submitLatestSnapshot(ctx context.Context) {
	snapshot, statesToSubmit := n.getLatestSnapshot(ctx)
	if snapshot == nil {
		return
	}

	snapshotSignature, encodedSnapshot, _, err := snapshot.SignSnapshot(ctx, n.bondedSigner)
	if err != nil {
		logger.Errorf("Error signing snapshot: %v", err)
	} else {
		err := n.summitDomain.Summit().SubmitSnapshot(ctx, n.unbondedSigner, encodedSnapshot, snapshotSignature)
		if err != nil {
			logger.Errorf("Failed to submit snapshot to summit: %v", err)
		} else {
			for originID, state := range statesToSubmit {
				n.summitMyLatestStates[originID] = state
			}
		}
	}
}

// Start starts the notary.
func (n Notary) Start(ctx context.Context) error {
	// First initialize a map to track what was the last state signed by this notary
	n.loadSummitMyLatestStates(ctx)

	for {
		select {
		// parent loop terminated
		case <-ctx.Done():
			logger.Info("Notary exiting without error")
			return nil
		case <-time.After(n.refreshInterval):
			n.loadSummitGuardLatestStates(ctx)
			n.submitLatestSnapshot(ctx)
		}
	}
}
