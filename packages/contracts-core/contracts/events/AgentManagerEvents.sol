// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {AgentFlag} from "../libs/Structures.sol";

abstract contract AgentManagerEvents {
    /**
     * @notice Emitted whenever a Dispute is opened between two agents. This happens when a Guard submits
     * their report for the Notary-signed statement to `StatementInbox`.
     * @param disputeIndex  Index of the dispute in the global list of all opened disputes
     * @param guardIndex    Index of the Guard in the Agent Merkle Tree
     * @param notaryIndex   Index of the Notary in the Agent Merkle Tree
     */
    event DisputeOpened(uint256 disputeIndex, uint32 guardIndex, uint32 notaryIndex);

    /**
     * @notice Emitted whenever a Dispute is resolved. This happens when an Agent who was in Dispute is slashed.
     * Note: this won't be emitted, if an Agent was slashed without being in Dispute.
     * @param disputeIndex  Index of the dispute in the global list of all opened disputes
     * @param slashedIndex  Index of the slashed agent in the Agent Merkle Tree
     * @param rivalIndex    Index of the rival agent in the Agent Merkle Tree
     * @param fraudProver   Address who provided fraud proof to resolve the Dispute
     */
    event DisputeResolved(uint256 disputeIndex, uint32 slashedIndex, uint32 rivalIndex, address fraudProver);

    // ═══════════════════════════════════════════════ DATA UPDATED ════════════════════════════════════════════════════

    /**
     * @notice Emitted whenever the root of the Agent Merkle Tree is updated.
     * @param newRoot   New agent merkle root
     */
    event RootUpdated(bytes32 newRoot);

    /**
     * @notice Emitted after the contract owner proposes a new agent root to resolve the stuck chain.
     * @param newRoot   New agent merkle root that was proposed
     */
    event AgentRootProposed(bytes32 newRoot);

    /**
     * @notice Emitted after the contract owner cancels the previously proposed agent root.
     * @param proposedRoot  Agent merkle root that was proposed
     */
    event ProposedAgentRootCancelled(bytes32 proposedRoot);

    /**
     * @notice Emitted after the contract owner resolves the previously proposed agent root.
     * @param proposedRoot  New agent merkle root that was resolved
     */
    event ProposedAgentRootResolved(bytes32 proposedRoot);

    /**
     * @notice Emitted whenever a status of the agent is updated.
     * @dev Only Active/Unstaking/Resting/Slashed flags could be stored in the Agent Merkle Tree.
     * Unknown flag is the default (zero) value and is used to represent agents that never
     * interacted with the BondingManager contract.
     * Fraudulent flag is the value for the agent who has been proven to commit fraud, but their
     * status hasn't been updated to Slashed in the Agent Merkle Tree. This is due to the fact
     * that the update of the status requires a merkle proof of the old status, and happens
     * in a separate transaction because of that.
     * @param flag      Flag defining agent status:
     * @param domain    Domain assigned to the agent (ZERO for Guards)
     * @param agent     Agent address
     */
    event StatusUpdated(AgentFlag flag, uint32 indexed domain, address indexed agent);
}
