package guard

import (
	"context"
	"fmt"
	"github.com/synapsecns/sanguine/core/metrics"
	signerConfig "github.com/synapsecns/sanguine/ethergo/signer/config"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
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
	// TODO: change to metrics type
	originLatestStates map[uint32]types.State
	handler            metrics.Handler
}

// NewGuard creates a new guard.
//
//nolint:cyclop
func NewGuard(ctx context.Context, cfg config.AgentConfig, handler metrics.Handler) (_ Guard, err error) {
	guard := Guard{
		refreshInterval: time.Second * time.Duration(cfg.RefreshIntervalSeconds),
	}
	guard.domains = []domains.DomainClient{}

	guard.bondedSigner, err = signerConfig.SignerFromConfig(ctx, cfg.BondedSigner)
	if err != nil {
		return Guard{}, fmt.Errorf("error with bondedSigner, could not create guard: %w", err)
	}

	guard.unbondedSigner, err = signerConfig.SignerFromConfig(ctx, cfg.UnbondedSigner)
	if err != nil {
		return Guard{}, fmt.Errorf("error with unbondedSigner, could not create guard: %w", err)
	}

	for domainName, domain := range cfg.Domains {
		var domainClient domains.DomainClient

		chainRPCURL := fmt.Sprintf("%s/confirmations/1/rpc/%d", cfg.BaseOmnirpcURL, domain.DomainID)
		domainClient, err = evm.NewEVM(ctx, domainName, domain, chainRPCURL)
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

	guard.handler = handler

	return guard, nil
}

func (g Guard) getDomainByDomainID(domainID uint32) domains.DomainClient {
	for _, domain := range g.domains {
		if domain.Config().DomainID == domainID {
			return domain
		}
	}
	return nil
}

//nolint:cyclop
func (g Guard) loadSummitLatestStates(parentCtx context.Context) {
	for _, domain := range g.domains {
		ctx, span := g.handler.Tracer().Start(parentCtx, "loadSummitLatestStates", trace.WithAttributes(
			attribute.Int("domain", int(domain.Config().DomainID)),
		))

		originID := domain.Config().DomainID
		latestState, err := g.summitDomain.Summit().GetLatestAgentState(ctx, originID, g.bondedSigner)
		if err != nil {
			latestState = nil
			logger.Errorf("Failed calling GetLatestAgentState for originID %d on the Summit contract: err = %v", originID, err)
			span.AddEvent("Failed calling GetLatestAgentState for originID on the Summit contract", trace.WithAttributes(
				attribute.Int("originID", int(originID)),
				attribute.String("err", err.Error()),
			))
		}
		if latestState != nil && latestState.Nonce() > uint32(0) {
			g.summitLatestStates[originID] = latestState
		}

		span.End()
	}
}

//nolint:cyclop
func (g Guard) loadOriginLatestStates(parentCtx context.Context) {
	for _, domain := range g.domains {
		ctx, span := g.handler.Tracer().Start(parentCtx, "loadOriginLatestStates", trace.WithAttributes(
			attribute.Int("domain", int(domain.Config().DomainID)),
		))

		originID := domain.Config().DomainID
		latestState, err := domain.Origin().SuggestLatestState(ctx)
		if err != nil {
			latestState = nil
			logger.Errorf("Failed calling SuggestLatestState for originID %d on the Origin contract: %v", originID, err)
			span.AddEvent("Failed calling SuggestLatestState for originID on the Origin contract", trace.WithAttributes(
				attribute.Int("originID", int(originID)),
				attribute.String("err", err.Error()),
			))
		} else if latestState == nil || latestState.Nonce() == uint32(0) {
			logger.Errorf("No latest state found for origin id %d", originID)
			span.AddEvent("No latest state found for origin id", trace.WithAttributes(
				attribute.Int("originID", int(originID)),
			))
		}
		if latestState != nil {
			// TODO: if overwriting, end span and start a new one
			g.originLatestStates[originID] = latestState
		}

		span.End()
	}
}

//nolint:cyclop
func (g Guard) getLatestSnapshot() (types.Snapshot, map[uint32]types.State) {
	statesToSubmit := make(map[uint32]types.State, len(g.domains))
	for _, domain := range g.domains {
		originID := domain.Config().DomainID
		summitLatest, ok := g.summitLatestStates[originID]
		if !ok || summitLatest == nil || summitLatest.Nonce() == 0 {
			summitLatest = nil
		}
		originLatest, ok := g.originLatestStates[originID]
		if !ok || originLatest == nil || originLatest.Nonce() == 0 {
			continue
		}
		if summitLatest != nil && summitLatest.Nonce() >= originLatest.Nonce() {
			// Here this guard already submitted this state
			continue
		}
		// TODO: add event for submitting that state
		statesToSubmit[originID] = originLatest
	}
	snapshotStates := make([]types.State, 0, len(statesToSubmit))
	for _, state := range statesToSubmit {
		if state.Nonce() == 0 {
			continue
		}
		snapshotStates = append(snapshotStates, state)
	}
	if len(snapshotStates) > 0 {
		return types.NewSnapshot(snapshotStates), statesToSubmit
	}
	//nolint:nilnil
	return nil, nil
}

//nolint:cyclop
func (g Guard) submitLatestSnapshot(parentCtx context.Context) {
	ctx, span := g.handler.Tracer().Start(parentCtx, "submitLatestSnapshot", trace.WithAttributes(
		attribute.Int("domain", int(g.summitDomain.Config().DomainID)),
	))

	defer func() {
		span.End()
	}()

	snapshot, statesToSubmit := g.getLatestSnapshot()
	if snapshot == nil {
		return
	}

	snapshotSignature, encodedSnapshot, _, err := snapshot.SignSnapshot(ctx, g.bondedSigner)
	if err != nil {
		logger.Errorf("Error signing snapshot: %v", err)
		span.AddEvent("Error signing snapshot", trace.WithAttributes(
			attribute.String("err", err.Error()),
		))
	} else {
		err = g.summitDomain.Inbox().SubmitSnapshot(ctx, g.unbondedSigner, encodedSnapshot, snapshotSignature)
		if err != nil {
			logger.Errorf("Failed to submit snapshot to inbox: %v", err)
			span.AddEvent("Failed to submit snapshot to inbox", trace.WithAttributes(
				attribute.String("err", err.Error()),
			))
		} else {
			for originID, state := range statesToSubmit {
				g.summitLatestStates[originID] = state
			}
		}
	}
}

//nolint:cyclop,gocognit
func (g Guard) runAttestationFraudDetection(parentCtx context.Context) {
	for _, domain := range g.domains {
		ctx, span := g.handler.Tracer().Start(parentCtx, "loadNumAttestations", trace.WithAttributes(
			attribute.Int("domain", int(domain.Config().DomainID)),
		))

		if domain.Config().DomainID == g.summitDomain.Config().DomainID {
			continue
		}

		numAttestations, err := domain.Destination().AttestationsAmount(ctx)
		if err != nil {
			logger.Errorf("Failed calling AttestationsAmount for destinationID %d on the Destination contract: err = %v", domain.Config().DomainID, err)
			span.AddEvent("Failed calling AttestationsAmount for destinationID on the Destination contract", trace.WithAttributes(
				attribute.Int("destinationID", int(domain.Config().DomainID)),
				attribute.String("err", err.Error()),
			))
		}

		fmt.Printf("CRONIN destinationID %d domain.Config().DomainID has an Attestation count of %v\n", domain.Config().DomainID, numAttestations)
		for count := numAttestations; count > 0; count-- {
			i := count - 1
			if count == 3 {
				fmt.Printf("CRONIN count == 3\n")
			}
			attestation, attestationSignature, err := domain.Destination().GetAttestation(ctx, i)
			if err != nil {
				logger.Errorf("Failed calling GetAttestation for destinationID %d and index %v on the Destination contract: err = %v", domain.Config().DomainID, i, err)
				span.AddEvent("Failed calling GetAttestation for destinationID and index on the Destination contract", trace.WithAttributes(
					attribute.Int("destinationID", int(domain.Config().DomainID)),
					attribute.Int64("index", int64(i)),
					attribute.String("err", err.Error()),
				))
				continue
			}
			if attestation == nil || attestation.Nonce() == 0 {
				continue
			}

			isValidAttestation, err := g.summitDomain.Summit().IsValidAttestation(ctx, attestation)
			if err != nil {
				// TODO (joe): This could be very serious if we can't reach Summit for an extended period of time.
				// We need to monitor and alert cases when an attestation can't be checked.
				// If need be, we should manually pause the destination from using this attestation.
				logger.Errorf("Failed calling IsValidAttestation for destinationID %d and index %d on the Summit contract: err = %v", domain.Config().DomainID, i, err)
				span.AddEvent("Failed calling IsValidAttestation for destinationID and index on the Summit contract", trace.WithAttributes(
					attribute.Int("destinationID", int(domain.Config().DomainID)),
					attribute.Int64("index", int64(i)),
					attribute.String("err", err.Error()),
				))
				continue
			}

			// TODO (joe): First look up snapshot on Summit to get all states
			// Then iterate through all states and check on the corresponding origin chains if it is ok
			attestationRaw, err := types.EncodeAttestation(attestation)
			if err != nil {
				fmt.Printf("CRONIN error encoding attestation\n")
				logger.Errorf("Failed EncodeAttestation for destinationID %d and index %d: err = %v", domain.Config().DomainID, i, err)
				span.AddEvent("Failed calling EncodeAttestation for destinationID and index", trace.WithAttributes(
					attribute.Int("destinationID", int(domain.Config().DomainID)),
					attribute.Int64("index", int64(i)),
					attribute.String("err", err.Error()),
				))
				continue
			}

			attestationSignatureRaw, err := types.EncodeSignature(attestationSignature)
			if err != nil {
				fmt.Printf("CRONIN error encoding attestation signature\n")
				logger.Errorf("Failed EncodeSignature for attestation on destinationID %d and index %d: err = %v", domain.Config().DomainID, i, err)
				span.AddEvent("Failed calling EncodeSignature for attestation on destinationID and index", trace.WithAttributes(
					attribute.Int("destinationID", int(domain.Config().DomainID)),
					attribute.Int64("index", int64(i)),
					attribute.String("err", err.Error()),
				))
				continue
			}

			recoveredNotaryAddress, err := attestation.RecoverSignerAddress(ctx, attestationSignature)
			if err != nil {
				fmt.Printf("CRONIN error in RecoverSignerAddress %v\n", err)
				logger.Errorf("Could not get pub key address from signature of bad attestation on destinationID %d and index %d: err = %v", domain.Config().DomainID, i, err)
				span.AddEvent("Could not get pub key address from signature of bad attestation on destinationID and index", trace.WithAttributes(
					attribute.Int("destinationID", int(domain.Config().DomainID)),
					attribute.Int64("index", int64(i)),
					attribute.String("err", err.Error()),
				))
				continue
			}

			if isValidAttestation {
				fmt.Printf("CRONIN Attestation is valid!!!\n")

				snapshot /*snapshotSignature*/, _, err := g.summitDomain.Summit().GetNotarySnapshot(ctx, attestationRaw)
				if err != nil {
					logger.Errorf("Failed to GetNotarySnapshot on Summit: %v", err)
					span.AddEvent("Failed to GetNotarySnapshot on Summit", trace.WithAttributes(
						attribute.String("err", err.Error()),
					))
					continue
				}

				fmt.Printf("CRONIN got Notary snapshot with this many states %v\n", len(snapshot.States()))
				for stateIndex, state := range snapshot.States() {
					fmt.Printf("CRONIN notary snapshot state[%v]: origin(%v), nonce(%v), blockNumber(%v), timeStamp(%v)\n",
						i, state.Origin(), state.Nonce(), state.BlockNumber(), state.Timestamp().Uint64())

					originDomain := g.getDomainByDomainID(state.Origin())
					if originDomain == nil {
						logger.Errorf("Do not have the ability to check state from origin domain id: %v", state.Origin())
						span.AddEvent("Do not have the ability to check state from origin domain id", trace.WithAttributes(
							attribute.Int64("originDomainID", int64(state.Origin())),
						))
						continue
					}
					isValidState, err := originDomain.Origin().IsValidState(ctx, state)
					if err != nil {
						// TODO (joe): This could be very serious if we can't reach the Origin for an extended period of time.
						// We need to monitor and alert cases when a state can't be checked.
						// If need be, we should manually pause the Synapse Chain from using this state.
						logger.Errorf("Failed calling IsValidState for originID %v, destinationID %d and index %d on the Origin contract: err = %v", state.Origin(), domain.Config().DomainID, i, err)
						span.AddEvent("Failed calling IsValidState for originID, destinationID and index on the Summit contract", trace.WithAttributes(
							attribute.Int("originID", int(state.Origin())),
							attribute.Int("destinationID", int(domain.Config().DomainID)),
							attribute.Int64("index", int64(i)),
							attribute.String("err", err.Error()),
						))
						continue
					}
					if !isValidState {
						fmt.Printf("CRONIN state is NOT valid!!! We found a FRAUDULENT STATE!!!!\n")
						// First do fraud report to Origin to initiate slashing
						// Then notify summit of pending fraud

						_, snapProofs, err := snapshot.SnapshotRootAndProofs()
						if err != nil {
							logger.Errorf("Failed getting snap root and proofs for bad state on origin %v, destinationID %d and index %d on the Origin contract: err = %v",
								state.Origin(), domain.Config().DomainID, stateIndex, err)
							span.AddEvent("Failed getting snap root and proofs for bad state on originID, destinationID and index on the Summit contract", trace.WithAttributes(
								attribute.Int("originID", int(state.Origin())),
								attribute.Int("destinationID", int(domain.Config().DomainID)),
								attribute.Int64("index", int64(stateIndex)),
								attribute.String("err", err.Error()),
							))
							continue
						}

						snapProof := snapProofs[stateIndex]

						err = originDomain.LightInbox().VerifyStateWithSnapshotProof(
							ctx,
							g.unbondedSigner,
							uint64(stateIndex),
							state,
							snapProof,
							attestationRaw,
							attestationSignature)
						if err != nil {
							logger.Errorf("Failed to call VerifyAttestation on Inbox: %v", err)
							span.AddEvent("Failed to call VerifyStateWithSnapshotProof on inbox", trace.WithAttributes(
								attribute.String("err", err.Error()),
							))
							continue
						}

					} else {
						fmt.Printf("CRONIN state is valid!!!\n")
					}
				}
			} else {
				// TODO (joe): Submit fraud report here
				fmt.Printf("CRONIN Attestation is NOT valid!!! WE FOUND FRAUD!!!!!!\n")

				notaryAgentStatus, err := g.summitDomain.BondingManager().GetAgentStatus(ctx, recoveredNotaryAddress)
				if err != nil {
					fmt.Printf("CRONIN GetAgentStatus with attestationNotaryAddress got error %v\n", err)
					logger.Errorf("Failed to GetAgentStatus for bad notary on destinationID %d and index %d: err = %v", domain.Config().DomainID, i, err)
					span.AddEvent("Failed to GetAgentStatus for bad notary on destinationID and index", trace.WithAttributes(
						attribute.Int("destinationID", int(domain.Config().DomainID)),
						attribute.Int64("index", int64(i)),
						attribute.String("err", err.Error()),
					))
					continue
				}

				if !(notaryAgentStatus.Flag() == uint8(types.AgentFlagActive) || notaryAgentStatus.Flag() == uint8(types.AgentFlagUnstaking)) {
					fmt.Printf("CRONIN Notary has already been slashed on Summit and has a status of %v\n", notaryAgentStatus.Flag())
					continue
				}
				fmt.Printf("CRONIN Notary has NOT been slashed on Summit and has a status of %v", notaryAgentStatus.Flag())

				err = g.summitDomain.Inbox().VerifyAttestation(ctx, g.unbondedSigner, attestationRaw, attestationSignatureRaw)
				if err != nil {
					logger.Errorf("Failed to call VerifyAttestation on Inbox: %v", err)
					span.AddEvent("Failed to call VerifyAttestation on inbox", trace.WithAttributes(
						attribute.String("err", err.Error()),
					))
					continue
				}

				// TODO (joe): Then submit fraud report on destination
			}
		}

		span.End()
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
			logger.Info("Guard exiting without error")
			return nil
		case <-time.After(g.refreshInterval):
			g.loadOriginLatestStates(ctx)
			g.submitLatestSnapshot(ctx)
			g.runAttestationFraudDetection(ctx)
		}
	}
}
