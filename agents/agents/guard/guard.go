package guard

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

// Guard scans origins for latest state and submits snapshots to the Summit.
type Guard struct {
	bondedSigner       signer.Signer
	unbondedSigner     signer.Signer
	domains            []domains.DomainClient
	summitDomain       domains.DomainClient
	refreshInterval    time.Duration
	summitLatestStates map[uint32]types.State
	originLatestStates map[uint32]types.State
}

// NewGuard creates a new guard.
//
//nolint:cyclop
func NewGuard(ctx context.Context, cfg config.AgentConfig) (_ Guard, err error) {
	guard := Guard{
		refreshInterval: time.Second * time.Duration(cfg.RefreshIntervalSeconds),
	}
	guard.domains = []domains.DomainClient{}

	guard.bondedSigner, err = config.SignerFromConfig(ctx, cfg.BondedSigner)
	if err != nil {
		return Guard{}, fmt.Errorf("error with bondedSigner, could not create guard: %w", err)
	}

	guard.unbondedSigner, err = config.SignerFromConfig(ctx, cfg.UnbondedSigner)
	if err != nil {
		return Guard{}, fmt.Errorf("error with unbondedSigner, could not create guard: %w", err)
	}

	for domainName, domain := range cfg.Domains {
		domainClient, err := evm.NewEVM(ctx, domainName, domain)
		if err != nil {
			return Guard{}, fmt.Errorf("failing to create evm for domain, could not create guard for: %w", err)
		}
		guard.domains = append(guard.domains, domainClient)
		if domain.DomainID == cfg.SummitDomainID {
			guard.summitDomain = domainClient
		}
	}

	guard.summitLatestStates = make(map[uint32]types.State, len(guard.domains))
	guard.originLatestStates = make(map[uint32]types.State, len(guard.domains))

	return guard, nil
}

func (g Guard) loadSummitLatestStates(ctx context.Context) {
	for _, domain := range g.domains {
		originID := domain.Config().DomainID
		latestState, err := domain.Summit().GetLatestAgentState(ctx, originID, g.bondedSigner)
		if err != nil {
			latestState = nil
			logger.Errorf("Failed calling GetLatestAgentState for originID on the Summit contract: %d, err = %v", originID, err)
		}
		if latestState != nil {
			g.summitLatestStates[originID] = latestState
		}
	}
}

func (g Guard) loadOriginLatestStates(ctx context.Context) {
	for _, domain := range g.domains {
		originID := domain.Config().DomainID
		latestState, err := domain.Origin().SuggestLatestState(ctx)
		if err != nil {
			latestState = nil
			logger.Errorf("Failed calling SuggestLatestState for originID: %d on the Origin contract: %v", originID, err)
		} else if latestState == nil {
			logger.Errorf("No latest state found for origin id %d", originID)
		}
		if latestState != nil {
			g.originLatestStates[originID] = latestState
		}
	}
}

func (g Guard) getLatestSnapshot() (types.Snapshot, map[uint32]types.State) {
	statesToSubmit := make(map[uint32]types.State, len(g.domains))
	for _, domain := range g.domains {
		originID := domain.Config().DomainID
		summitLatest, ok := g.summitLatestStates[originID]
		if !ok {
			summitLatest = nil
		}
		originLatest, ok := g.originLatestStates[originID]
		if !ok {
			continue
		}
		if summitLatest.Nonce() >= originLatest.Nonce() {
			// Here this guard already submitted this state
			continue
		}
		statesToSubmit[originID] = originLatest
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

func (g Guard) submitLatestSnapshot(ctx context.Context) {
	snapshot, statesToSubmit := g.getLatestSnapshot()
	if snapshot == nil {
		return
	}

	snapshotSignature, encodedSnapshot, _, err := snapshot.SignSnapshot(ctx, g.bondedSigner)
	if err != nil {
		logger.Errorf("Error signing snapshot: %v", err)
	} else {
		err := g.summitDomain.Summit().SubmitSnapshot(ctx, g.unbondedSigner, encodedSnapshot, snapshotSignature)
		if err != nil {
			logger.Errorf("Failed to submit snapshot to summit: %v", err)
		} else {
			for originID, state := range statesToSubmit {
				g.summitLatestStates[originID] = state
			}
		}
	}
}

// Start starts the guard.
//
//nolint:cyclop
func (g Guard) Start(ctx context.Context) error {
	// First initialize a map to track what was the last state signed by this guard
	g.loadSummitLatestStates(ctx)

	for {
		select {
		// parent loop terminated
		case <-ctx.Done():
			return nil
		case <-time.After(g.refreshInterval):
			g.loadOriginLatestStates(ctx)
			g.submitLatestSnapshot(ctx)
		}
	}
}
