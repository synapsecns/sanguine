package guard

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/agents/types"
	"github.com/synapsecns/sanguine/core/retry"
)

// handleSnapshot checks a snapshot for invalid states.
// If an invalid state is found, initiate slashing and submit a state report.
//
//nolint:cyclop,gocognit
func (g Guard) handleSnapshot(ctx context.Context, log ethTypes.Log) error {
	fraudSnapshot, err := g.inboxParser.ParseSnapshotAccepted(log)
	if err != nil {
		return fmt.Errorf("could not parse snapshot accepted: %w", err)
	}

	// Verify each state in the snapshot.
	for si, s := range fraudSnapshot.Snapshot.States() {
		stateIndex, state := si, s
		isSlashable, err := g.isStateSlashable(ctx, state)
		if err != nil {
			return fmt.Errorf("could not handle state: %w", err)
		}
		if !isSlashable {
			continue
		}

		// Initiate slashing on origin.
		_, err = g.txSubmitter.SubmitTransaction(ctx, big.NewInt(int64(state.Origin())), func(transactor *bind.TransactOpts) (tx *ethTypes.Transaction, err error) {
			tx, err = g.domains[state.Origin()].LightInbox().VerifyStateWithSnapshot(
				transactor,
				int64(stateIndex),
				fraudSnapshot.Payload,
				fraudSnapshot.Signature,
			)
			if err != nil {
				return nil, fmt.Errorf("could not verify state with snapshot: %w", err)
			}

			return
		})
		if err != nil {
			return fmt.Errorf("could not submit VerifyStateWithSnapshot tx: %w", err)
		}

		// Check if we should submit the state report.
		shouldSubmit, err := g.shouldSubmitStateReport(ctx, fraudSnapshot)
		if err != nil {
			return fmt.Errorf("could not check if should submit state report: %w", err)
		}
		if !shouldSubmit {
			return nil
		}

		// Submit the state report to summit.
		srSignature, _, _, err := state.SignState(ctx, g.bondedSigner)
		if err != nil {
			return fmt.Errorf("could not sign state: %w", err)
		}
		ok, err := g.prepareStateReport(ctx, fraudSnapshot.Agent, g.summitDomainID)
		if err != nil {
			return fmt.Errorf("could not prepare state report on summit: %w", err)
		}
		if !ok {
			continue
		}
		_, err = g.txSubmitter.SubmitTransaction(ctx, big.NewInt(int64(g.summitDomainID)), func(transactor *bind.TransactOpts) (tx *ethTypes.Transaction, err error) {
			tx, err = g.domains[g.summitDomainID].Inbox().SubmitStateReportWithSnapshot(
				transactor,
				int64(stateIndex),
				srSignature,
				fraudSnapshot.Payload,
				fraudSnapshot.Signature,
			)
			if err != nil {
				return nil, fmt.Errorf("could not submit state report with snapshot to summit: %w", err)
			}

			return
		})
		if err != nil {
			return fmt.Errorf("could not submit SubmitStateReportWithSnapshot tx: %w", err)
		}

		// Submit the state report to the remote chain.
		ok, err = g.prepareStateReport(ctx, fraudSnapshot.Agent, fraudSnapshot.AgentDomain)
		if err != nil {
			return fmt.Errorf("could not prepare state report on summit: %w", err)
		}
		if !ok {
			continue
		}
		_, err = g.txSubmitter.SubmitTransaction(ctx, big.NewInt(int64(fraudSnapshot.AgentDomain)), func(transactor *bind.TransactOpts) (tx *ethTypes.Transaction, err error) {
			tx, err = g.domains[fraudSnapshot.AgentDomain].LightInbox().SubmitStateReportWithSnapshot(
				transactor,
				int64(stateIndex),
				srSignature,
				fraudSnapshot.Payload,
				fraudSnapshot.Signature,
			)
			if err != nil {
				return nil, fmt.Errorf("could not submit state report with snapshot to agent domain %d: %w", fraudSnapshot.AgentDomain, err)
			}

			return
		})
		if err != nil {
			return fmt.Errorf("could not submit SubmitStateReportWithSnapshot tx: %w", err)
		}
	}

	return nil
}

// Only submit a state report if we are not on Summit, and the snapshot
// agent is not currently in dispute.
func (g Guard) shouldSubmitStateReport(ctx context.Context, snapshot *types.FraudSnapshot) (bool, error) {
	var disputeStatus types.DisputeStatus
	var err error
	contractCall := func(ctx context.Context) error {
		disputeStatus, err = g.domains[g.summitDomainID].BondingManager().GetDisputeStatus(ctx, snapshot.Agent)
		if err != nil {
			return fmt.Errorf("could not get dispute status: %w", err)
		}

		return nil
	}
	err = retry.WithBackoff(ctx, contractCall, g.retryConfig...)
	if err != nil {
		return false, fmt.Errorf("could not get dispute status: %w", err)
	}

	isNotary := snapshot.AgentDomain != 0
	isNotInDispute := disputeStatus.Flag() == types.DisputeFlagNone
	shouldSubmit := isNotary && isNotInDispute
	return shouldSubmit, nil
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

// handleAttestation checks whether an attestation is valid.
// If invalid, initiate slashing and/or submit a fraud report.
func (g Guard) handleAttestation(ctx context.Context, log ethTypes.Log) error {
	fraudAttestation, err := g.lightInboxParser.ParseAttestationAccepted(log)
	if err != nil {
		return fmt.Errorf("could not parse attestation accepted: %w", err)
	}

	var isValid bool
	contractCall := func(ctx context.Context) error {
		isValid, err = g.domains[g.summitDomainID].Summit().IsValidAttestation(ctx, fraudAttestation.Payload)
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
		return g.handleValidAttestation(ctx, fraudAttestation)
	}

	return g.handleInvalidAttestation(ctx, fraudAttestation)
}

// handleValidAttestation handles an attestation that is valid, but may
// attest to a snapshot that contains an invalid state.
//
//nolint:cyclop,gocognit
func (g Guard) handleValidAttestation(ctx context.Context, fraudAttestation *types.FraudAttestation) error {
	// Fetch the attested snapshot.
	var snapshot types.Snapshot
	var err error
	contractCall := func(ctx context.Context) error {
		snapshot, err = g.domains[g.summitDomainID].Summit().GetNotarySnapshot(ctx, fraudAttestation.Payload)
		if err != nil {
			return fmt.Errorf("could not get snapshot: %w", err)
		}

		return nil
	}
	err = retry.WithBackoff(ctx, contractCall, g.retryConfig...)
	if err != nil {
		return fmt.Errorf("could not get snapshot: %w", err)
	}

	snapPayload, err := snapshot.Encode()
	if err != nil {
		return fmt.Errorf("could not encode snapshot: %w", err)
	}

	// Verify each state in the snapshot.
	for si, s := range snapshot.States() {
		stateIndex, state := si, s
		isSlashable, err := g.isStateSlashable(ctx, state)
		if err != nil {
			return fmt.Errorf("could not check if state is slashable: %w", err)
		}
		if !isSlashable {
			continue
		}

		// Initiate slashing on origin.
		_, err = g.txSubmitter.SubmitTransaction(ctx, big.NewInt(int64(state.Origin())), func(transactor *bind.TransactOpts) (tx *ethTypes.Transaction, err error) {
			tx, err = g.domains[state.Origin()].LightInbox().VerifyStateWithAttestation(
				transactor,
				int64(stateIndex),
				snapPayload,
				fraudAttestation.Payload,
				fraudAttestation.Signature,
			)
			if err != nil {
				return nil, fmt.Errorf("could not verify state with attestation: %w", err)
			}

			return
		})

		if err != nil {
			return fmt.Errorf("could not submit VerifyStateWithAttestation tx: %w", err)
		}

		// Submit the state report on summit.
		srSignature, _, _, err := state.SignState(ctx, g.bondedSigner)
		if err != nil {
			return fmt.Errorf("could not sign state: %w", err)
		}
		ok, err := g.prepareStateReport(ctx, fraudAttestation.Notary, g.summitDomainID)
		if err != nil {
			return fmt.Errorf("could not prepare state report on summit: %w", err)
		}
		if !ok {
			continue
		}
		_, err = g.txSubmitter.SubmitTransaction(ctx, big.NewInt(int64(g.summitDomainID)), func(transactor *bind.TransactOpts) (tx *ethTypes.Transaction, err error) {
			tx, err = g.domains[g.summitDomainID].Inbox().SubmitStateReportWithAttestation(
				transactor,
				int64(stateIndex),
				srSignature,
				snapPayload,
				fraudAttestation.Payload,
				fraudAttestation.Signature,
			)
			if err != nil {
				return nil, fmt.Errorf("could not submit state report with attestation on summit: %w", err)
			}

			return
		})
		if err != nil {
			return fmt.Errorf("could not submit SubmitStateReportWithAttestation tx: %w", err)
		}

		// Submit the state report on the remote chain.
		ok, err = g.prepareStateReport(ctx, fraudAttestation.Notary, fraudAttestation.AgentDomain)
		if err != nil {
			return fmt.Errorf("could not prepare state report on remote: %w", err)
		}
		if !ok {
			continue
		}
		_, err = g.txSubmitter.SubmitTransaction(ctx, big.NewInt(int64(fraudAttestation.AgentDomain)), func(transactor *bind.TransactOpts) (tx *ethTypes.Transaction, err error) {
			tx, err = g.domains[fraudAttestation.AgentDomain].LightInbox().SubmitStateReportWithAttestation(
				transactor,
				int64(stateIndex),
				srSignature,
				snapPayload,
				fraudAttestation.Payload,
				fraudAttestation.Signature,
			)
			if err != nil {
				return nil, fmt.Errorf("could not submit state report with attestation on agent domain %d: %w", fraudAttestation.AgentDomain, err)
			}

			return
		})
		if err != nil {
			return fmt.Errorf("could not submit SubmitStateReportWithAttestation tx: %w", err)
		}
	}

	return nil
}

// prepareStateReport checks if the given agent is in a slashable status (Active or Unstaking),
// and relays the agent status from Summit to the given chain if necessary.
//
//nolint:cyclop
func (g Guard) prepareStateReport(ctx context.Context, agent common.Address, chainID uint32) (ok bool, err error) {
	var agentStatus types.AgentStatus
	//nolint:nestif
	if chainID == g.summitDomainID {
		contractCall := func(ctx context.Context) error {
			agentStatus, err = g.domains[chainID].BondingManager().GetAgentStatus(ctx, agent)
			if err != nil {
				return fmt.Errorf("could not get agent status: %w", err)
			}

			return nil
		}
		err = retry.WithBackoff(ctx, contractCall, g.retryConfig...)
		if err != nil {
			return false, fmt.Errorf("could not get agent status: %w", err)
		}
	} else {
		contractCall := func(ctx context.Context) error {
			agentStatus, err = g.domains[chainID].LightManager().GetAgentStatus(ctx, agent)
			if err != nil {
				return fmt.Errorf("could not get agent status: %w", err)
			}

			return nil
		}
		err = retry.WithBackoff(ctx, contractCall, g.retryConfig...)
		if err != nil {
			return false, fmt.Errorf("could not get agent status: %w", err)
		}
	}
	if err != nil {
		return false, fmt.Errorf("could not get agent status: %w", err)
	}

	//nolint:exhaustive
	switch agentStatus.Flag() {
	case types.AgentFlagUnknown:
		if chainID == g.summitDomainID {
			return false, fmt.Errorf("cannot submit state report for Unknown agent on summit")
		}
		// Update the agent status to active using the last known root on remote chain.
		err = g.guardDB.StoreRelayableAgentStatus(
			ctx,
			agent,
			types.AgentFlagUnknown,
			types.AgentFlagActive,
			chainID,
		)
		if err != nil {
			return false, fmt.Errorf("could not store relayable agent status: %w", err)
		}
		err = g.updateAgentStatus(ctx, chainID)
		if err != nil {
			return false, err
		}
		return true, nil
	case types.AgentFlagActive, types.AgentFlagUnstaking:
		return true, nil
	default:
		return false, nil
	}
}

// handleInvalidAttestation handles an invalid attestation by initiating slashing on summit,
// then submitting an attestation fraud report on the accused agent's Domain.
func (g Guard) handleInvalidAttestation(ctx context.Context, fraudAttestation *types.FraudAttestation) error {
	// Initiate slashing for invalid attestation.
	_, err := g.txSubmitter.SubmitTransaction(ctx, big.NewInt(int64(g.summitDomainID)), func(transactor *bind.TransactOpts) (tx *ethTypes.Transaction, err error) {
		tx, err = g.domains[g.summitDomainID].Inbox().VerifyAttestation(
			transactor,
			fraudAttestation.Payload,
			fraudAttestation.Signature,
		)
		if err != nil {
			return nil, fmt.Errorf("could not verify attestation: %w", err)
		}

		return
	})
	if err != nil {
		return fmt.Errorf("could not submit VerifyAttestation tx: %w", err)
	}

	// Submit a fraud report by calling `submitAttestationReport()` on the remote chain.
	arSignature, _, _, err := fraudAttestation.Attestation.SignAttestation(ctx, g.bondedSigner, false)
	if err != nil {
		return fmt.Errorf("could not sign attestation: %w", err)
	}
	arSignatureEncoded, err := types.EncodeSignature(arSignature)
	if err != nil {
		return fmt.Errorf("could not encode signature: %w", err)
	}
	_, err = g.txSubmitter.SubmitTransaction(ctx, big.NewInt(int64(fraudAttestation.AgentDomain)), func(transactor *bind.TransactOpts) (tx *ethTypes.Transaction, err error) {
		tx, err = g.domains[fraudAttestation.AgentDomain].LightInbox().SubmitAttestationReport(
			transactor,
			fraudAttestation.Payload,
			arSignatureEncoded,
			fraudAttestation.Signature,
		)
		if err != nil {
			return nil, fmt.Errorf("could not submit attestation report: %w", err)
		}

		return
	})
	if err != nil {
		return fmt.Errorf("could not submit SubmitAttestationReport tx: %w", err)
	}

	return nil
}

//nolint:cyclop
func (g Guard) handleReceipt(ctx context.Context, log ethTypes.Log) error {
	fraudReceipt, err := g.inboxParser.ParseReceiptAccepted(log)
	if err != nil {
		return fmt.Errorf("could not parse receipt accepted: %w", err)
	}

	// Validate the receipt.
	receipt, err := types.DecodeReceipt(fraudReceipt.RcptPayload)
	if err != nil {
		return fmt.Errorf("could not decode receipt: %w", err)
	}

	var isValid bool
	contractCall := func(ctx context.Context) error {
		isValid, err = g.domains[receipt.Destination()].Destination().IsValidReceipt(ctx, fraudReceipt.RcptPayload)
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

	// Initiate slashing for an invalid receipt, and optionally submit a fraud report.
	//nolint:nestif
	if receipt.Destination() == g.summitDomainID {
		_, err = g.txSubmitter.SubmitTransaction(ctx, big.NewInt(int64(receipt.Destination())), func(transactor *bind.TransactOpts) (tx *ethTypes.Transaction, err error) {
			tx, err = g.domains[receipt.Destination()].Inbox().VerifyReceipt(
				transactor,
				fraudReceipt.RcptPayload,
				fraudReceipt.RcptSignature,
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
				fraudReceipt.RcptPayload,
				fraudReceipt.RcptSignature,
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
				fraudReceipt.RcptPayload,
				fraudReceipt.RcptSignature,
				rrReceiptBytes,
			)
			if err != nil {
				return nil, fmt.Errorf("could not submit receipt report: %w", err)
			}

			return
		})
		if err != nil {
			return fmt.Errorf("could not submit SubmitReceiptReport tx: %w", err)
		}
	}

	return nil
}

// handleStatusUpdated stores models related to a StatusUpdated event.
//
//nolint:cyclop,gocognit
func (g Guard) handleStatusUpdated(ctx context.Context, log ethTypes.Log, chainID uint32) error {
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

		var remoteStatus types.AgentStatus
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
			// Fetch the current remote status and check whether the status is synced.
			contractCall := func(ctx context.Context) error {
				remoteStatus, err = g.domains[statusUpdated.Domain].LightManager().GetAgentStatus(ctx, statusUpdated.Agent)
				if err != nil {
					return fmt.Errorf("could not get agent status: %w", err)
				}

				return nil
			}
			err = retry.WithBackoff(ctx, contractCall, g.retryConfig...)
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
func (g Guard) handleRootUpdated(ctx context.Context, log ethTypes.Log, chainID uint32) error {
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
func (g Guard) updateAgentStatus(ctx context.Context, chainID uint32) error {
	eligibleAgentTrees, err := g.guardDB.GetRelayableAgentStatuses(ctx, chainID)
	if err != nil {
		return fmt.Errorf("could not get update agent status parameters: %w", err)
	}

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

	localRootBlockNumber, err := g.guardDB.GetSummitBlockNumberForRoot(ctx, common.BytesToHash(localRoot[:]).String())
	if err != nil {
		return fmt.Errorf("could not get block number for local root: %w", err)
	}

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
			logger.Infof("Relaying agent status for agent %s on chain %d", tree.AgentAddress.String(), chainID)
			// Fetch the agent status to be relayed from Summit.
			var agentStatus types.AgentStatus
			contractCall := func(ctx context.Context) error {
				agentStatus, err = g.domains[g.summitDomainID].BondingManager().GetAgentStatus(ctx, tree.AgentAddress)
				if err != nil {
					return fmt.Errorf("could not get agent status: %w", err)
				}

				return nil
			}
			err = retry.WithBackoff(ctx, contractCall, g.retryConfig...)
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
