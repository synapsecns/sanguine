package base

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	agentTypes "github.com/synapsecns/sanguine/agents/types"
)

// StoreRelayableAgentStatus stores a relayable agent status.
func (s Store) StoreRelayableAgentStatus(
	ctx context.Context,
	agentAddress common.Address,
	staleFlag agentTypes.AgentFlagType,
	updatedFlag agentTypes.AgentFlagType,
	domain uint32,
) error {
	dbTx := s.DB().WithContext(ctx).
		Create(&RelayableAgentStatus{
			AgentAddress:            agentAddress.String(),
			StaleFlag:               staleFlag,
			UpdatedFlag:             updatedFlag,
			Domain:                  domain,
			AgentStatusRelayedState: agentTypes.Queued,
		})

	if dbTx.Error != nil {
		return fmt.Errorf("failed to store relayable agent status: %w", dbTx.Error)
	}

	return nil
}

// UpdateAgentStatusRelayedState updates the state for a RelayableAgentStatus.
func (s Store) UpdateAgentStatusRelayedState(
	ctx context.Context,
	agentAddress common.Address,
	state agentTypes.AgentStatusRelayedState,
) error {
	mask := RelayableAgentStatus{
		AgentAddress: agentAddress.String(),
	}

	dbTx := s.DB().WithContext(ctx).
		Model(&RelayableAgentStatus{}).
		Where(mask).
		Update(AgentStatusRelayedStateFieldName, state)
	if dbTx.Error != nil {
		return fmt.Errorf("failed to update relayed state: %w", dbTx.Error)
	}
	return nil
}
