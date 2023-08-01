package guard

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/agents/types"
)

func (g Guard) handleSnapshot(ctx context.Context, log ethTypes.Log) error {
	fraudSnapshot, err := g.inboxParser.ParseSnapshotAccepted(log)
	if err != nil {
		return fmt.Errorf("could not parse snapshot accepted: %w", err)
	}

	// Handle each state in the snapshot
	for stateIndex, state := range fraudSnapshot.Snapshot.States() {
		isSlashable, err := g.isStateSlashable(ctx, state, fraudSnapshot.Agent)
		if err != nil {
			return fmt.Errorf("could not handle state: %w", err)
		}
		if !isSlashable {
			continue
		}

		// Initiate slashing on origin.
		_, err = g.domains[state.Origin()].LightInbox().VerifyStateWithSnapshot(
			ctx,
			g.unbondedSigner,
			int64(stateIndex),
			fraudSnapshot.Payload,
			fraudSnapshot.Signature,
		)
		if err != nil {
			return fmt.Errorf("could not verify state with snapshot: %w", err)
		}

		// Check if we should submit the state report
		shouldSubmit, err := g.shouldSubmitStateReport(ctx, fraudSnapshot)
		if err != nil {
			return fmt.Errorf("could not check if should submit state report: %w", err)
		}
		if !shouldSubmit {
			continue
		}

		// Submit the state report
		srSignature, _, _, err := state.SignState(ctx, g.bondedSigner)
		if err != nil {
			return fmt.Errorf("could not sign state: %w", err)
		}
		_, err = g.domains[g.summitDomainID].Inbox().SubmitStateReportWithSnapshot(
			ctx,
			g.unbondedSigner,
			int64(stateIndex),
			srSignature,
			fraudSnapshot.Payload,
			fraudSnapshot.Signature,
		)
		if err != nil {
			return fmt.Errorf("could not submit state report with snapshot: %w", err)
		}
	}

	return nil
}

// Only submit a state report if we are not on Summit, and the snapshot
// agent is not currently in dispute.
func (g Guard) shouldSubmitStateReport(ctx context.Context, snapshot *types.FraudSnapshot) (bool, error) {
	disputeStatus, err := g.domains[g.summitDomainID].BondingManager().GetDisputeStatus(ctx, snapshot.Agent)
	if err != nil {
		return false, fmt.Errorf("could not get dispute status: %w", err)
	}

	isNotSummit := snapshot.AgentDomain != 0
	isNotInDispute := disputeStatus.Flag() == types.DisputeFlagNone
	shouldSubmit := isNotSummit && isNotInDispute
	return shouldSubmit, nil
}

// Verify a state against a snapshot.
// If invalid, submit a state report on summit.
func (g Guard) isStateSlashable(ctx context.Context, state types.State, agent common.Address) (bool, error) {
	statePayload, err := types.EncodeState(state)
	if err != nil {
		return false, fmt.Errorf("could not encode state: %w", err)
	}

	// Verify that the state is valid w.r.t. Origin
	isValid, err := g.domains[state.Origin()].Origin().IsValidState(
		ctx,
		statePayload,
	)
	if err != nil {
		return false, fmt.Errorf("could not check validity of state: %w", err)
	}
	if isValid {
		return true, nil
	}

	// Verify that the agent is slashable
	agentStatus, err := g.domains[state.Origin()].LightManager().GetAgentStatus(ctx, agent)
	if err != nil {
		return false, fmt.Errorf("could not get agent status: %w", err)
	}
	return isSlashable(agentStatus.Flag()), nil
}

//nolint:gocognit,cyclop
func (g Guard) handleAttestation(ctx context.Context, log ethTypes.Log) error {
	fraudAttestation, err := g.lightInboxParser.ParseAttestationAccepted(log)
	if err != nil {
		return fmt.Errorf("could not parse attestation accepted: %w", err)
	}

	isValid, err := g.domains[g.summitDomainID].Summit().IsValidAttestation(ctx, fraudAttestation.Payload)
	if err != nil {
		return fmt.Errorf("could not check validity of attestation: %w", err)
	}

	//nolint:nestif
	if isValid {
		// The attestation is valid, but may have a state not matching that of Origin.
		// Fetch the snapshot, then verify each individual state with the attestation.
		snapshot, err := g.domains[g.summitDomainID].Summit().GetNotarySnapshot(ctx, fraudAttestation.Payload)
		if err != nil {
			return fmt.Errorf("could not get snapshot: %w", err)
		}

		for stateIndex, state := range snapshot.States() {
			snapPayload, err := types.EncodeSnapshot(snapshot)
			if err != nil {
				return fmt.Errorf("could not encode snapshot: %w", err)
			}

			isSlashable, err := g.isStateSlashable(ctx, state, fraudAttestation.Notary)
			if err != nil {
				return fmt.Errorf("could not check if state is slashable: %w", err)
			}
			if !isSlashable {
				continue
			}

			// Initiate slashing on origin.
			_, err = g.domains[state.Origin()].LightInbox().VerifyStateWithAttestation(
				ctx,
				g.unbondedSigner,
				int64(stateIndex),
				snapPayload,
				fraudAttestation.Payload,
				fraudAttestation.Signature,
			)
			if err != nil {
				return fmt.Errorf("could not verify state with attestation: %w", err)
			}

			// Submit the state report.
			srSignature, _, _, err := state.SignState(ctx, g.bondedSigner)
			if err != nil {
				return fmt.Errorf("could not sign state: %w", err)
			}
			_, err = g.domains[g.summitDomainID].Inbox().SubmitStateReportWithAttestation(
				ctx,
				g.unbondedSigner,
				int64(stateIndex),
				srSignature,
				snapPayload,
				fraudAttestation.Payload,
				fraudAttestation.Signature,
			)
			if err != nil {
				return fmt.Errorf("could not submit state report with attestation: %w", err)
			}
		}
		return nil
	}

	// If attestation is invalid, we need to slash the agent
	// by calling `verifyAttestation()` on the summit domain.
	_, err = g.domains[g.summitDomainID].Inbox().VerifyAttestation(
		ctx,
		g.unbondedSigner,
		fraudAttestation.Payload,
		fraudAttestation.Signature,
	)
	if err != nil {
		return fmt.Errorf("could not verify attestation: %w", err)
	}

	// Finally, we submit a fraud report by calling `submitAttestationReport()` on the remote chain.
	arSignature, _, _, err := fraudAttestation.Attestation.SignAttestation(ctx, g.bondedSigner, false)
	if err != nil {
		return fmt.Errorf("could not sign attestation: %w", err)
	}
	arSignatureEncoded, err := types.EncodeSignature(arSignature)
	if err != nil {
		return fmt.Errorf("could not encode signature: %w", err)
	}

	// Call `submitAttestationReport` on the notary's associated remote domain.
	_, err = g.domains[fraudAttestation.AgentDomain].LightInbox().SubmitAttestationReport(
		ctx,
		g.unbondedSigner,
		fraudAttestation.Payload,
		arSignatureEncoded,
		fraudAttestation.Signature,
	)
	if err != nil {
		return fmt.Errorf("could not submit attestation report: %w", err)
	}

	return nil
}

//nolint:cyclop
func (g Guard) handleReceipt(ctx context.Context, log ethTypes.Log) error {
	fraudReceipt, err := g.inboxParser.ParseReceiptAccepted(log)
	if err != nil {
		return fmt.Errorf("could not parse receipt accepted: %w", err)
	}

	receipt, err := types.DecodeReceipt(fraudReceipt.RcptPayload)
	if err != nil {
		return fmt.Errorf("could not decode receipt: %w", err)
	}
	isValid, err := g.domains[receipt.Destination()].Destination().IsValidReceipt(ctx, fraudReceipt.RcptPayload)
	if err != nil {
		return fmt.Errorf("could not check validity of attestation: %w", err)
	}

	//nolint:nestif
	if !isValid {
		// TODO: merge this logic once solidity interfaces are de-duped
		if receipt.Destination() == g.summitDomainID {
			_, err = g.domains[receipt.Destination()].Inbox().VerifyReceipt(ctx, g.unbondedSigner, fraudReceipt.RcptPayload, fraudReceipt.RcptSignature)
			if err != nil {
				return fmt.Errorf("could not verify receipt: %w", err)
			}
		} else {
			_, err = g.domains[receipt.Destination()].LightInbox().VerifyReceipt(ctx, g.unbondedSigner, fraudReceipt.RcptPayload, fraudReceipt.RcptSignature)
			if err != nil {
				return fmt.Errorf("could not verify receipt: %w", err)
			}
			rrReceipt, _, _, err := receipt.SignReceipt(ctx, g.bondedSigner, false)
			if err != nil {
				return fmt.Errorf("could not sign receipt: %w", err)
			}
			rrReceiptBytes, err := types.EncodeSignature(rrReceipt)
			if err != nil {
				return fmt.Errorf("could not encode receipt: %w", err)
			}
			_, err = g.domains[g.summitDomainID].Inbox().SubmitReceiptReport(ctx, g.unbondedSigner, fraudReceipt.RcptPayload, fraudReceipt.RcptSignature, rrReceiptBytes)
			if err != nil {
				return fmt.Errorf("could not submit receipt report: %w", err)
			}
		}
	}

	return nil
}

//nolint:cyclop
func (g Guard) handleStatusUpdated(ctx context.Context, log ethTypes.Log, chainID uint32) error {
	statusUpdated, err := g.bondingManagerParser.ParseStatusUpdated(log)
	if err != nil {
		return fmt.Errorf("could not parse status updated: %w", err)
	}

	//nolint:exhaustive
	switch types.AgentFlagType(statusUpdated.Flag) {
	case types.AgentFlagFraudulent:
		agentProof, err := g.domains[g.summitDomainID].BondingManager().GetProof(ctx, statusUpdated.Agent)
		if err != nil {
			return fmt.Errorf("could not get proof: %w", err)
		}

		_, err = g.domains[g.summitDomainID].BondingManager().CompleteSlashing(
			ctx,
			g.unbondedSigner,
			statusUpdated.Domain,
			statusUpdated.Agent,
			agentProof,
		)
		if err != nil {
			return fmt.Errorf("could not complete slashing: %w", err)
		}
	case types.AgentFlagSlashed:
		agentRoot, err := g.domains[g.summitDomainID].BondingManager().GetAgentRoot(ctx)
		if err != nil {
			return fmt.Errorf("could not get agent root: %w", err)
		}

		agentProof, err := g.domains[g.summitDomainID].BondingManager().GetProof(ctx, statusUpdated.Agent)
		if err != nil {
			return fmt.Errorf("could not get proof: %w", err)
		}

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
			chainID,
			log.BlockNumber,
		)
		if err != nil {
			return fmt.Errorf("could not store agent root: %w", err)
		}
	default:
		logger.Infof("Witnessed agent status updated, but not handling [status=%d, agent=%s]", statusUpdated.Flag, statusUpdated.Agent)
	}

	return nil
}

func (g Guard) handleDisputeOpened(ctx context.Context, log ethTypes.Log) error {
	disputeOpened, err := g.parseDisputeOpened(log)
	if err != nil {
		return fmt.Errorf("could not parse dispute opened: %w", err)
	}

	_, guardAddress, err := g.domains[g.summitDomainID].BondingManager().GetAgent(ctx, big.NewInt(int64(disputeOpened.guardIndex)))
	if err != nil {
		return fmt.Errorf("could not get agent: %w", err)
	}

	_, notaryAddress, err := g.domains[g.summitDomainID].BondingManager().GetAgent(ctx, big.NewInt(int64(disputeOpened.notaryIndex)))
	if err != nil {
		return fmt.Errorf("could not get agent: %w", err)
	}

	// Store the dispute in the database.
	err = g.guardDB.StoreDispute(
		ctx,
		disputeOpened.disputeIndex,
		types.Opened,
		guardAddress,
		disputeOpened.notaryIndex,
		notaryAddress,
	)
	if err != nil {
		return fmt.Errorf("could not store dispute: %w", err)
	}

	return nil
}

// disputeOpened is a wrapper struct used to merge the
// lightmanager.DisputeOpened and bondingmangaer.DisputeOpened structs.
type disputeOpened struct {
	disputeIndex *big.Int
	guardIndex   uint32
	notaryIndex  uint32
}

func (g Guard) parseDisputeOpened(log ethTypes.Log) (*disputeOpened, error) {
	disputeOpenedLight, err := g.lightManagerParser.ParseDisputeOpened(log)
	if err == nil {
		return &disputeOpened{
			disputeIndex: disputeOpenedLight.DisputeIndex,
			guardIndex:   disputeOpenedLight.GuardIndex,
			notaryIndex:  disputeOpenedLight.NotaryIndex,
		}, nil
	}
	disputeOpenedBonding, err := g.bondingManagerParser.ParseDisputeOpened(log)
	if err == nil {
		return &disputeOpened{
			disputeIndex: disputeOpenedBonding.DisputeIndex,
			guardIndex:   disputeOpenedBonding.GuardIndex,
			notaryIndex:  disputeOpenedBonding.NotaryIndex,
		}, nil
	}
	return nil, fmt.Errorf("could not parse dispute opened: %w", err)
}

func (g Guard) handleRootUpdated(ctx context.Context, log ethTypes.Log, chainID uint32) error {
	newRoot, err := g.bondingManagerParser.ParseRootUpdated(log)
	if err != nil || newRoot == nil {
		return fmt.Errorf("could not parse root updated: %w", err)
	}

	err = g.guardDB.StoreAgentRoot(
		ctx,
		*newRoot,
		chainID,
		log.BlockNumber,
	)
	if err != nil {
		return fmt.Errorf("could not store agent root: %w", err)
	}

	return nil
}

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

func (g Guard) updateAgentStatus(ctx context.Context, chainID uint32) error {
	eligibleAgentTrees, err := g.guardDB.GetUpdateAgentStatusParameters(ctx)
	if err != nil {
		return fmt.Errorf("could not get update agent status parameters: %w", err)
	}

	if len(eligibleAgentTrees) == 0 {
		return nil
	}

	blockNumber, err := g.guardDB.GetLatestConfirmedSummitBlockNumber(ctx, chainID)
	if err != nil {
		return fmt.Errorf("could not get latest confirmed summit block number: %w", err)
	}

	// Filter the eligible agent roots by the given block number and call updateAgentStatus()
	for _, tree := range eligibleAgentTrees {
		if tree.BlockNumber >= blockNumber {
			agentStatus, err := g.domains[g.summitDomainID].BondingManager().GetAgentStatus(ctx, tree.AgentAddress)
			if err != nil {
				return fmt.Errorf("could not get agent status: %w", err)
			}
			_, err = g.domains[chainID].LightManager().UpdateAgentStatus(
				ctx,
				g.unbondedSigner,
				tree.AgentAddress,
				agentStatus,
				tree.Proof,
			)
			if err != nil {
				return fmt.Errorf("could not update agent status: %w", err)
			}
		}
	}

	return nil
}
