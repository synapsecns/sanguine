package guard

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/agents/contracts/inbox"
	"github.com/synapsecns/sanguine/agents/types"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/core/retry"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

// handleSnapshotAccepted checks a snapshot for invalid states.
// If an invalid state is found, initiate slashing and submit a state report.
//
//nolint:cyclop,gocognit
func (g Guard) handleSnapshotAccepted(parentCtx context.Context, log ethTypes.Log) error {
	ctx, _ := g.handler.Tracer().Start(parentCtx, "handleSnapshotAccepted")

	snapshotData, err := g.inboxParser.ParseSnapshotAccepted(log)
	if err != nil {
		return fmt.Errorf("could not parse snapshot accepted: %w", err)
	}

	err = g.handleSnapshot(ctx, snapshotData.Snapshot, snapshotData)
	if err != nil {
		return fmt.Errorf("could not handle snapshot: %w", err)
	}
	return nil
}

// Determine which chains need to receive the state report.
// A chain should receive a state report if the fraudulent Notary has not yet been reported on that chain-
// That is, a dispute has not yet been opened with this notary on that chain.
func (g Guard) getStateReportChains(ctx context.Context, chainID uint32, agent common.Address) ([]uint32, error) {
	stateReportChains := []uint32{}
	for _, reportChainID := range []uint32{g.summitDomainID, chainID} {
		status, err := g.getDisputeStatus(ctx, agent)
		if err != nil {
			return []uint32{}, err
		}
		isNotary := reportChainID != 0
		isNotInDispute := status.Flag() == types.DisputeFlagNone
		if isNotary && isNotInDispute {
			stateReportChains = append(stateReportChains, reportChainID)
		}
	}
	return stateReportChains, nil
}

// isStateSlashable checks if a state is slashable, i.e. if the state is valid on the
// Origin, and if the agent is in a slashable status.
func (g Guard) isStateSlashable(ctx context.Context, state types.State) (bool, error) {
	statePayload, err := state.Encode()
	if err != nil {
		return false, fmt.Errorf("could not encode state: %w", err)
	}

	// Verify that the state is valid w.r.t. Origin.
	var isValid bool
	contractCall := func(ctx context.Context) error {
		isValid, err = g.domains[state.Origin()].Origin().IsValidState(ctx, statePayload)
		if err != nil {
			return fmt.Errorf("could not check validity of state: %w", err)
		}

		return nil
	}
	err = retry.WithBackoff(ctx, contractCall, g.retryConfig...)
	if err != nil {
		return false, fmt.Errorf("could not check validity of state: %w", err)
	}
	return !isValid, nil
}

// handleAttestationAccepted checks whether an attestation is valid.
// If invalid, initiate slashing and/or submit a fraud report.
func (g Guard) handleAttestationAccepted(parentCtx context.Context, log ethTypes.Log) error {
	ctx, _ := g.handler.Tracer().Start(parentCtx, "handleAttestationAccepted")

	attestationData, err := g.lightInboxParser.ParseAttestationAccepted(log)
	if err != nil {
		return fmt.Errorf("could not parse attestation accepted: %w", err)
	}

	var isValid bool
	contractCall := func(ctx context.Context) error {
		isValid, err = g.domains[g.summitDomainID].Summit().IsValidAttestation(ctx, attestationData.AttestationPayload())
		if err != nil {
			return fmt.Errorf("could not check validity of attestation: %w", err)
		}

		return nil
	}
	err = retry.WithBackoff(ctx, contractCall, g.retryConfig...)
	if err != nil {
		return fmt.Errorf("could not check validity of attestation: %w", err)
	}

	if isValid {
		return g.handleValidAttestation(ctx, attestationData)
	}

	return g.handleInvalidAttestation(ctx, attestationData)
}

// handleValidAttestation handles an attestation that is valid, but may
// attest to a snapshot that contains an invalid state.
//
//nolint:cyclop,gocognit
func (g Guard) handleValidAttestation(parentCtx context.Context, attestationData *types.AttestationWithMetadata) (err error) {
	ctx, span := g.handler.Tracer().Start(parentCtx, "handleValidAttestation", trace.WithAttributes(
		attribute.String(metrics.Agent, attestationData.Agent().String()),
		attribute.Int(metrics.AgentDomain, int(attestationData.AgentDomain())),
	))
	defer metrics.EndSpanWithErr(span, err)

	// Fetch the attested snapshot.
	var snapshot types.Snapshot
	contractCall := func(ctx context.Context) error {
		snapshot, err = g.domains[g.summitDomainID].Summit().GetNotarySnapshot(ctx, attestationData.AttestationPayload())
		if err != nil {
			return fmt.Errorf("could not get snapshot: %w", err)
		}

		return nil
	}
	err = retry.WithBackoff(ctx, contractCall, g.retryConfig...)
	if err != nil {
		return fmt.Errorf("could not get snapshot: %w", err)
	}

	// Set the SnapshotPayload so that it can be passed to contract calls.
	snapshotPayload, err := snapshot.Encode()
	if err != nil {
		return fmt.Errorf("could not encode snapshot: %w", err)
	}
	attestationData.SetSnapshotPayload(snapshotPayload)

	err = g.handleSnapshot(ctx, snapshot, attestationData)
	if err != nil {
		return fmt.Errorf("could not handle snapshot: %w", err)
	}

	return nil
}

// handleSnapshot handles a snapshot by validating each state in the snapshot.
// If an invalid state is found, initiate slashing and submit state reports on eligible chains.
func (g Guard) handleSnapshot(parentCtx context.Context, snapshot types.Snapshot, data types.StateValidationData) (err error) {
	snapRoot, _, _ := snapshot.SnapshotRootAndProofs()
	ctx, span := g.handler.Tracer().Start(parentCtx, "handleSnapshot", trace.WithAttributes(
		attribute.String(metrics.SnapRoot, common.BytesToHash(snapRoot[:]).String()),
	))
	defer metrics.EndSpanWithErr(span, err)

	// Process each state in the snapshot.
	for si, s := range snapshot.States() {
		stateIndex, state := si, s
		isSlashable, err := g.isStateSlashable(ctx, state)
		if err != nil {
			err = fmt.Errorf("could not handle state: %w", err)
			return err
		}
		if !isSlashable {
			continue
		}

		// Initiate slashing on origin.
		err = g.verifyState(ctx, state, stateIndex, data)
		if err != nil {
			err = fmt.Errorf("could not verify state: %w", err)
			return err
		}

		// Evaluate which chains need a state report.
		stateReportChains, err := g.getStateReportChains(ctx, data.AgentDomain(), data.Agent())
		if err != nil {
			err = fmt.Errorf("could not get state report chains: %w", err)
			return err
		}
		stateReportChainsInt := make([]int, len(stateReportChains))
		for i, chainID := range stateReportChains {
			stateReportChainsInt[i] = int(chainID)
		}
		span.AddEvent("got state report chains", trace.WithAttributes(
			attribute.IntSlice("state_report_chains", stateReportChainsInt),
		))

		// Submit the state report on each eligible chain.
		// If a notary has not been reported anywhere,
		// report should be submitted on both summit and remote.
		for _, chainID := range stateReportChains {
			err = g.submitStateReport(ctx, chainID, state, stateIndex, data)
			if err != nil {
				err = fmt.Errorf("could not submit state report: %w", err)
				return err
			}
		}
	}
	return nil
}

// handleInvalidAttestation handles an invalid attestation by initiating slashing on summit,
// then submitting an attestation fraud report on the accused agent's Domain.
func (g Guard) handleInvalidAttestation(parentCtx context.Context, attestationData *types.AttestationWithMetadata) (err error) {
	ctx, span := g.handler.Tracer().Start(parentCtx, "handleInvalidAttestation", trace.WithAttributes(
		attribute.String(metrics.Agent, attestationData.Agent().String()),
		attribute.Int(metrics.AgentDomain, int(attestationData.AgentDomain())),
	))
	defer metrics.EndSpanWithErr(span, err)

	// Initiate slashing for invalid attestation.
	_, err = g.txSubmitter.SubmitTransaction(ctx, big.NewInt(int64(g.summitDomainID)), func(transactor *bind.TransactOpts) (tx *ethTypes.Transaction, err error) {
		tx, err = g.domains[g.summitDomainID].Inbox().VerifyAttestation(
			transactor,
			attestationData.AttestationPayload(),
			attestationData.AttestationSignature(),
		)
		if err != nil {
			return nil, fmt.Errorf("could not verify attestation: %w", err)
		}

		return
	})
	if err != nil {
		err = fmt.Errorf("could not submit VerifyAttestation tx: %w", err)
		return err
	}

	// Submit a fraud report by calling `submitAttestationReport()` on the remote chain.
	arSignature, _, _, err := attestationData.Attestation.SignAttestation(ctx, g.bondedSigner, false)
	if err != nil {
		return fmt.Errorf("could not sign attestation: %w", err)
	}
	arSignatureEncoded, err := types.EncodeSignature(arSignature)
	if err != nil {
		return fmt.Errorf("could not encode signature: %w", err)
	}
	_, err = g.txSubmitter.SubmitTransaction(ctx, big.NewInt(int64(attestationData.AgentDomain())), func(transactor *bind.TransactOpts) (tx *ethTypes.Transaction, err error) {
		tx, err = g.domains[attestationData.AgentDomain()].LightInbox().SubmitAttestationReport(
			transactor,
			attestationData.AttestationPayload(),
			arSignatureEncoded,
			attestationData.AttestationSignature(),
		)
		if err != nil {
			return nil, fmt.Errorf("could not submit attestation report: %w", err)
		}

		return
	})
	if err != nil {
		err = fmt.Errorf("could not submit SubmitAttestationReport tx: %w", err)
		return err
	}

	return nil
}

// handleReceiptAccepted checks whether a receipt is valid and submits a receipt report if not.
//
//nolint:cyclop
func (g Guard) handleReceiptAccepted(parentCtx context.Context, log ethTypes.Log) error {
	ctx, _ := g.handler.Tracer().Start(parentCtx, "handleReceiptAccepted")

	event, err := g.inboxParser.ParseReceiptAccepted(log)
	if err != nil {
		return fmt.Errorf("could not parse receipt accepted: %w", err)
	}

	// Validate the receipt.
	receipt, err := types.DecodeReceipt(event.RcptPayload)
	if err != nil {
		return fmt.Errorf("could not decode receipt: %w", err)
	}

	var isValid bool
	contractCall := func(ctx context.Context) error {
		isValid, err = g.domains[receipt.Destination()].Destination().IsValidReceipt(ctx, event.RcptPayload)
		if err != nil {
			return fmt.Errorf("could not check validity of attestation: %w", err)
		}

		return nil
	}
	err = retry.WithBackoff(ctx, contractCall, g.retryConfig...)
	if err != nil {
		return fmt.Errorf("could not check validity of attestation: %w", err)
	}
	if isValid {
		return nil
	}

	err = g.handleInvalidReceipt(ctx, receipt, event)
	if err != nil {
		return fmt.Errorf("could not handle invalid receipt: %w", err)
	}

	return nil
}

func (g Guard) handleInvalidReceipt(parentCtx context.Context, receipt types.Receipt, event *inbox.InboxReceiptAccepted) (err error) {
	leaf := receipt.MessageHash()
	ctx, span := g.handler.Tracer().Start(parentCtx, "handleInvalidReceipt", trace.WithAttributes(
		attribute.String("attestation_notary", receipt.AttestationNotary().String()),
		attribute.Int(metrics.Destination, int(receipt.Destination())),
		attribute.String(metrics.MessageLeaf, common.BytesToHash(leaf[:]).String()),
	))
	defer metrics.EndSpanWithErr(span, err)

	// Initiate slashing for an invalid receipt, and optionally submit a fraud report.
	//nolint:nestif
	if receipt.Destination() == g.summitDomainID {
		_, err = g.txSubmitter.SubmitTransaction(ctx, big.NewInt(int64(receipt.Destination())), func(transactor *bind.TransactOpts) (tx *ethTypes.Transaction, err error) {
			tx, err = g.domains[receipt.Destination()].Inbox().VerifyReceipt(
				transactor,
				event.RcptPayload,
				event.RcptSignature,
			)
			if err != nil {
				return nil, fmt.Errorf("could not verify receipt: %w", err)
			}

			return
		})
		if err != nil {
			return fmt.Errorf("could not submit VerifyReceipt tx: %w", err)
		}
	} else {
		_, err = g.txSubmitter.SubmitTransaction(ctx, big.NewInt(int64(receipt.Destination())), func(transactor *bind.TransactOpts) (tx *ethTypes.Transaction, err error) {
			tx, err = g.domains[receipt.Destination()].LightInbox().VerifyReceipt(
				transactor,
				event.RcptPayload,
				event.RcptSignature,
			)
			if err != nil {
				return nil, fmt.Errorf("could not verify receipt: %w", err)
			}

			return
		})
		if err != nil {
			return fmt.Errorf("could not submit VerifyReceipt tx: %w", err)
		}

		rrReceipt, _, _, err := receipt.SignReceipt(ctx, g.bondedSigner, false)
		if err != nil {
			return fmt.Errorf("could not sign receipt: %w", err)
		}
		rrReceiptBytes, err := types.EncodeSignature(rrReceipt)
		if err != nil {
			return fmt.Errorf("could not encode receipt: %w", err)
		}
		_, err = g.txSubmitter.SubmitTransaction(ctx, big.NewInt(int64(g.summitDomainID)), func(transactor *bind.TransactOpts) (tx *ethTypes.Transaction, err error) {
			tx, err = g.domains[g.summitDomainID].Inbox().SubmitReceiptReport(
				transactor,
				event.RcptPayload,
				event.RcptSignature,
				rrReceiptBytes,
			)
			if err != nil {
				return nil, fmt.Errorf("could not submit receipt report: %w", err)
			}

			return
		})
		if err != nil {
			err = fmt.Errorf("could not submit SubmitReceiptReport tx: %w", err)
			return err
		}
	}
	return nil
}

// handleStatusUpdated stores models related to a StatusUpdated event.
//
//nolint:cyclop,gocognit
func (g Guard) handleStatusUpdated(parentCtx context.Context, log ethTypes.Log, chainID uint32) error {
	ctx, _ := g.handler.Tracer().Start(parentCtx, "handleStatusAccepted")

	statusUpdated, err := g.bondingManagerParser.ParseStatusUpdated(log)
	if err != nil {
		return fmt.Errorf("could not parse status updated: %w", err)
	}

	//nolint:exhaustive
	switch types.AgentFlagType(statusUpdated.Flag) {
	case types.AgentFlagFraudulent:
		var agentProof [][32]byte
		contractCall := func(ctx context.Context) error {
			agentProof, err = g.domains[g.summitDomainID].BondingManager().GetProof(ctx, statusUpdated.Agent)
			if err != nil {
				return fmt.Errorf("could not get proof: %w", err)
			}

			return nil
		}
		err = retry.WithBackoff(ctx, contractCall, g.retryConfig...)
		if err != nil {
			return fmt.Errorf("could not get proof: %w", err)
		}

		_, err = g.txSubmitter.SubmitTransaction(ctx, big.NewInt(int64(g.summitDomainID)), func(transactor *bind.TransactOpts) (tx *ethTypes.Transaction, err error) {
			tx, err = g.domains[g.summitDomainID].BondingManager().CompleteSlashing(
				transactor,
				statusUpdated.Domain,
				statusUpdated.Agent,
				agentProof,
			)
			if err != nil {
				return nil, fmt.Errorf("could not complete slashing: %w", err)
			}

			return
		})
		if err != nil {
			return fmt.Errorf("could not submit CompleteSlashing tx: %w", err)
		}
	case types.AgentFlagSlashed:
		var agentRoot [32]byte
		contractCall := func(ctx context.Context) error {
			agentRoot, err = g.domains[g.summitDomainID].BondingManager().GetAgentRoot(ctx)
			if err != nil {
				return fmt.Errorf("could not get agent root: %w", err)
			}

			return nil
		}
		err = retry.WithBackoff(ctx, contractCall, g.retryConfig...)
		if err != nil {
			return fmt.Errorf("could not get agent root: %w", err)
		}

		var agentProof [][32]byte
		contractCall = func(ctx context.Context) error {
			agentProof, err = g.domains[g.summitDomainID].BondingManager().GetProof(ctx, statusUpdated.Agent)
			if err != nil {
				return fmt.Errorf("could not get proof: %w", err)
			}

			return nil
		}
		err = retry.WithBackoff(ctx, contractCall, g.retryConfig...)
		if err != nil {
			return fmt.Errorf("could not get proof: %w", err)
		}

		if chainID == g.summitDomainID {
			err = g.guardDB.StoreAgentTree(
				ctx,
				agentRoot,
				statusUpdated.Agent,
				log.BlockNumber,
				agentProof,
			)
			if err != nil {
				return fmt.Errorf("could not store agent tree: %w", err)
			}

			err = g.guardDB.StoreAgentRoot(
				ctx,
				agentRoot,
				log.BlockNumber,
			)
			if err != nil {
				return fmt.Errorf("could not store agent root: %w", err)
			}
		}

		if statusUpdated.Domain != 0 {
			remoteStatus, err := g.getAgentStatus(ctx, statusUpdated.Domain, statusUpdated.Agent)
			if err != nil {
				return fmt.Errorf("could not get agent status: %w", err)
			}

			if remoteStatus.Flag() == types.AgentFlagType(statusUpdated.Flag) {
				return nil
			}

			// If not synced, store a relayable agent status.
			err = g.guardDB.StoreRelayableAgentStatus(
				ctx,
				statusUpdated.Agent,
				remoteStatus.Flag(),
				types.AgentFlagType(statusUpdated.Flag),
				statusUpdated.Domain,
			)
			if err != nil {
				return fmt.Errorf("could not store relayable agent status: %w", err)
			}
		}
	default:
		logger.Errorf("Witnessed agent status updated, but not handling [status=%d, agent=%s]", statusUpdated.Flag, statusUpdated.Agent)
	}

	return nil
}

// handleRootUpdated stores models related to a RootUpdated event.
func (g Guard) handleRootUpdated(parentCtx context.Context, log ethTypes.Log, chainID uint32) error {
	ctx, _ := g.handler.Tracer().Start(parentCtx, "handleRootUpdated")

	if chainID == g.summitDomainID {
		newRoot, err := g.bondingManagerParser.ParseRootUpdated(log)
		if err != nil || newRoot == nil {
			return fmt.Errorf("could not parse root updated: %w", err)
		}
		err = g.guardDB.StoreAgentRoot(
			ctx,
			*newRoot,
			log.BlockNumber,
		)
		if err != nil {
			return fmt.Errorf("could not store agent root: %w", err)
		}
	}

	return nil
}

// updateAgentStatuses updates the status of all agents on all chains, except for summit.
func (g Guard) updateAgentStatuses(ctx context.Context) error {
	for _, domain := range g.domains {
		chainID := domain.Config().DomainID
		if chainID == g.summitDomainID {
			continue
		}

		err := g.updateAgentStatus(ctx, chainID)
		if err != nil {
			return err
		}
	}
	return nil
}

// updateAgentStatus updates the status for each agent with a pending agent tree model,
// and open dispute on remote chain.
//
//nolint:cyclop,gocognit
func (g Guard) updateAgentStatus(parentCtx context.Context, chainID uint32) (err error) {
	ctx, span := g.handler.Tracer().Start(parentCtx, "updateAgentStatus", trace.WithAttributes(
		attribute.Int(metrics.ChainID, int(chainID)),
	))
	defer metrics.EndSpanWithErr(span, err)

	eligibleAgentTrees, err := g.guardDB.GetRelayableAgentStatuses(ctx, chainID)
	if err != nil {
		return fmt.Errorf("could not get update agent status parameters: %w", err)
	}

	span.AddEvent("got eligible agent trees", trace.WithAttributes(
		attribute.Int("numAgentTrees", len(eligibleAgentTrees)),
	))
	if len(eligibleAgentTrees) == 0 {
		return nil
	}

	var localRoot [32]byte
	contractCall := func(ctx context.Context) error {
		localRoot, err = g.domains[chainID].LightManager().GetAgentRoot(ctx)
		if err != nil {
			return fmt.Errorf("could not get agent root: %w", err)
		}

		return nil
	}
	err = retry.WithBackoff(ctx, contractCall, g.retryConfig...)
	if err != nil {
		return fmt.Errorf("could not get agent root: %w", err)
	}

	span.AddEvent("got local root", trace.WithAttributes(
		attribute.String("localRoot", common.BytesToHash(localRoot[:]).String()),
	))
	localRootBlockNumber, err := g.guardDB.GetSummitBlockNumberForRoot(ctx, common.BytesToHash(localRoot[:]).String())
	if err != nil {
		return fmt.Errorf("could not get block number for local root: %w", err)
	}
	span.AddEvent("got local root block number", trace.WithAttributes(
		attribute.Int("localRootBlockNumber", int(localRootBlockNumber)),
	))

	// Filter the eligible agent roots by the given block number and call updateAgentStatus().
	for _, t := range eligibleAgentTrees {
		tree := t
		// Get the first recorded summit block number for the tree agent root.
		treeBlockNumber, err := g.guardDB.GetSummitBlockNumberForRoot(ctx, tree.AgentRoot)
		if err != nil {
			return fmt.Errorf("could not get block number for local root: %w", err)
		}
		//nolint:nestif
		if localRootBlockNumber >= treeBlockNumber {
			span.AddEvent("updating agent status", trace.WithAttributes(
				attribute.Int("chainID", int(chainID)),
				attribute.String("agentAddress", tree.AgentAddress.String()),
			))
			logger.Infof("Relaying agent status for agent %s on chain %d", tree.AgentAddress.String(), chainID)

			// Fetch the agent status to be relayed from Summit.
			agentStatus, err := g.getAgentStatus(ctx, g.summitDomainID, tree.AgentAddress)
			if err != nil {
				return fmt.Errorf("could not get agent status: %w", err)
			}

			if agentStatus.Domain() != chainID {
				continue
			}

			// Update agent status on remote.
			_, err = g.txSubmitter.SubmitTransaction(ctx, big.NewInt(int64(chainID)), func(transactor *bind.TransactOpts) (tx *ethTypes.Transaction, err error) {
				tx, err = g.domains[chainID].LightManager().UpdateAgentStatus(
					transactor,
					tree.AgentAddress,
					agentStatus,
					tree.Proof,
				)
				if err != nil {
					return nil, fmt.Errorf("could not submit UpdateAgentStatus tx: %w", err)
				}
				logger.Infof("Updated agent status on chain %d for agent %s: %s [hash: %s]", chainID, tree.AgentAddress.String(), agentStatus.Flag().String(), tx.Hash())
				return
			})
			if err != nil {
				return fmt.Errorf("could not submit UpdateAgentStatus tx: %w", err)
			}

			// Mark the relayable status as Relayed.
			err = g.guardDB.UpdateAgentStatusRelayedState(
				ctx,
				tree.AgentAddress,
				types.Relayed,
			)
			if err != nil {
				return fmt.Errorf("could not update agent status relayed state: %w", err)
			}
		}
	}

	return nil
}
