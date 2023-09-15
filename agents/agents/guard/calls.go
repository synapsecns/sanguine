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
	"github.com/synapsecns/sanguine/ethergo/signer/signer"
)

type agentStatusContract interface {
	// GetAgentStatus returns the current agent status for the given agent.
	GetAgentStatus(ctx context.Context, address common.Address) (types.AgentStatus, error)
}

// getAgentStatus fetches the agent status of an agent from the given chain.
func (g Guard) getAgentStatus(ctx context.Context, chainID uint32, agent common.Address) (agentStatus types.AgentStatus, err error) {
	var contract agentStatusContract
	if chainID == g.summitDomainID {
		contract = g.domains[chainID].BondingManager()
	} else {
		contract = g.domains[chainID].LightManager()
	}
	contractCall := func(ctx context.Context) error {
		agentStatus, err = contract.GetAgentStatus(ctx, agent)
		if err != nil {
			return fmt.Errorf("could not get agent status from contract: %w", err)
		}
		return nil
	}
	err = retry.WithBackoff(ctx, contractCall, g.retryConfig...)
	if err != nil {
		return nil, fmt.Errorf("could not get agent status after retry: %w", err)
	}
	return agentStatus, nil
}

// verifyState verifies a state on a given chain.
func (g Guard) verifyState(ctx context.Context, state types.State, stateIndex int, data types.StateValidationData) (err error) {
	var submitFunc func(transactor *bind.TransactOpts) (tx *ethTypes.Transaction, err error)
	if types.HasAttestation(data) {
		submitFunc = func(transactor *bind.TransactOpts) (tx *ethTypes.Transaction, err error) {
			tx, err = g.domains[state.Origin()].LightInbox().VerifyStateWithAttestation(
				transactor,
				int64(stateIndex),
				data.SnapshotPayload(),
				data.AttestationPayload(),
				data.AttestationSignature(),
			)
			return
		}
	} else {
		submitFunc = func(transactor *bind.TransactOpts) (tx *ethTypes.Transaction, err error) {
			tx, err = g.domains[state.Origin()].LightInbox().VerifyStateWithSnapshot(
				transactor,
				int64(stateIndex),
				data.SnapshotPayload(),
				data.SnapshotSignature(),
			)
			return
		}
	}

	// Ensure the agent that provided the snapshot is active on origin.
	ok, err := g.ensureAgentActive(ctx, data.Agent(), state.Origin())
	if err != nil {
		return fmt.Errorf("could not ensure agent is active: %w", err)
	}
	if !ok {
		logger.Infof("Agent %s is not active on chain %d; not verifying snapshot state", data.Agent().Hex(), state.Origin())
		return nil
	}

	_, err = g.txSubmitter.SubmitTransaction(ctx, big.NewInt(int64(state.Origin())), submitFunc)
	if err != nil {
		return fmt.Errorf("could not verify state on chain %d: %w", state.Origin(), err)
	}
	return nil
}

type stateReportContract interface {
	// SubmitStateReportWithSnapshot reports to the inbox that a state within a snapshot is invalid.
	SubmitStateReportWithSnapshot(transactor *bind.TransactOpts, stateIndex int64, signature signer.Signature, snapPayload []byte, snapSignature []byte) (tx *ethTypes.Transaction, err error)
	// SubmitStateReportWithAttestation submits a state report corresponding to an attesation for an invalid state.
	SubmitStateReportWithAttestation(transactor *bind.TransactOpts, stateIndex int64, signature signer.Signature, snapPayload, attPayload, attSignature []byte) (tx *ethTypes.Transaction, err error)
}

// submitStateReport submits a state report to the given chain, provided a snapshot or attestation.
func (g Guard) submitStateReport(ctx context.Context, chainID uint32, state types.State, stateIndex int, data types.StateValidationData) (err error) {
	var contract stateReportContract
	if chainID == g.summitDomainID {
		contract = g.domains[chainID].Inbox()
	} else {
		contract = g.domains[chainID].LightInbox()
	}

	var submitFunc func(transactor *bind.TransactOpts) (tx *ethTypes.Transaction, err error)
	srSignature, _, _, err := state.SignState(ctx, g.bondedSigner)
	if err != nil {
		return fmt.Errorf("could not sign state: %w", err)
	}
	if types.HasAttestation(data) {
		submitFunc = func(transactor *bind.TransactOpts) (tx *ethTypes.Transaction, err error) {
			tx, err = contract.SubmitStateReportWithAttestation(
				transactor,
				int64(stateIndex),
				srSignature,
				data.SnapshotPayload(),
				data.AttestationPayload(),
				data.AttestationSignature(),
			)
			return
		}
	} else {
		submitFunc = func(transactor *bind.TransactOpts) (tx *ethTypes.Transaction, err error) {
			tx, err = contract.SubmitStateReportWithSnapshot(
				transactor,
				int64(stateIndex),
				srSignature,
				data.SnapshotPayload(),
				data.SnapshotSignature(),
			)
			return
		}
	}

	// Ensure the agent that provided the snapshot is active on the agent's respective domain.
	ok, err := g.ensureAgentActive(ctx, data.Agent(), chainID)
	if err != nil {
		return fmt.Errorf("could not ensure agent is active: %w", err)
	}
	if !ok {
		logger.Infof("Agent %s is not active on chain %d; not verifying snapshot state", data.Agent().Hex(), chainID)
		return nil
	}

	_, err = g.txSubmitter.SubmitTransaction(ctx, big.NewInt(int64(chainID)), submitFunc)
	if err != nil {
		return fmt.Errorf("could not submit state report to chain %d: %w", chainID, err)
	}
	return nil
}

// getDisputeStatus fetches the dispute status of an agent from Summit.
func (g Guard) getDisputeStatus(ctx context.Context, agent common.Address) (status types.DisputeStatus, err error) {
	contractCall := func(ctx context.Context) error {
		status, err = g.domains[g.summitDomainID].BondingManager().GetDisputeStatus(ctx, agent)
		if err != nil {
			return fmt.Errorf("could not get dispute status: %w", err)
		}
		return nil
	}
	err = retry.WithBackoff(ctx, contractCall, g.retryConfig...)
	if err != nil {
		return nil, fmt.Errorf("could not get dispute status: %w", err)
	}
	return status, nil
}

// ensureAgentActive checks if the given agent is in a slashable status (Active or Unstaking),
// and relays the agent status from Summit to the given chain if necessary.
func (g Guard) ensureAgentActive(ctx context.Context, agent common.Address, chainID uint32) (ok bool, err error) {
	agentStatus, err := g.getAgentStatus(ctx, chainID, agent)
	if err != nil {
		return false, fmt.Errorf("could not get agent status: %w", err)
	}

	//nolint:exhaustive
	switch agentStatus.Flag() {
	case types.AgentFlagUnknown:
		if chainID == g.summitDomainID {
			return false, fmt.Errorf("cannot submit state report for Unknown agent on summit")
		}
		// Fetch the agent status from Summit.
		agentStatusSummit, err := g.getAgentStatus(ctx, g.summitDomainID, agent)
		if err != nil {
			return false, fmt.Errorf("could not get agent status: %w", err)
		}
		if agentStatusSummit.Flag() != types.AgentFlagActive && agentStatusSummit.Flag() != types.AgentFlagUnstaking {
			return false, fmt.Errorf("agent is not active or unstaking on summit: %s [status=%s]", agent.Hex(), agentStatusSummit.Flag().String())
		}
		// Update the agent status using the last known root on remote chain.
		err = g.relayAgentStatus(ctx, agent, chainID, agentStatusSummit.Flag())
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

// relayAgentStatus relays an Active agent status from Summit to a remote
// chain where the agent is unknown.
func (g Guard) relayAgentStatus(ctx context.Context, agent common.Address, chainID uint32, flag types.AgentFlagType) error {
	err := g.guardDB.StoreRelayableAgentStatus(
		ctx,
		agent,
		types.AgentFlagUnknown,
		flag,
		chainID,
	)
	if err != nil {
		return fmt.Errorf("could not store relayable agent status: %w", err)
	}
	err = g.updateAgentStatus(ctx, chainID)
	if err != nil {
		return err
	}
	return nil
}
