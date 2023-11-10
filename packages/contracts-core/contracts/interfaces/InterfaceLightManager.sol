// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {AgentStatus} from "../libs/Structures.sol";

interface InterfaceLightManager {
    /**
     * @notice Updates agent status, using a proof against the latest known Agent Merkle Root.
     * @dev Will revert if the provided proof doesn't match the latest merkle root.
     * @param agent     Agent address
     * @param status    Structure specifying agent status: (flag, domain, index)
     * @param proof     Merkle proof of Active status for the agent
     */
    function updateAgentStatus(address agent, AgentStatus memory status, bytes32[] memory proof) external;

    /**
     * @notice Updates the root of Agent Merkle Tree that the Light Manager is tracking.
     * Could be only called by a local Destination contract, which is supposed to
     * verify the attested Agent Merkle Roots.
     * @param agentRoot_    New Agent Merkle Root
     */
    function setAgentRoot(bytes32 agentRoot_) external;

    /**
     * @notice Allows contract owner to set the agent root to resolve the "stuck" chain
     * by proposing the new agent root. The contract owner will be able to resolve the proposed
     * agent root after a certain period of time.
     * Note: this function could be called multiple times, each time the timer will be reset.
     * This could only be called if no fresh data has been submitted by the Notaries to the Inbox,
     * indicating that the chain is stuck for one of the reasons:
     * - All active Notaries are in Dispute.
     * - No active Notaries exist under the current agent root.
     * @dev Will revert if any of the following conditions is met:
     * - Caller is not the contract owner.
     * - Agent root is empty.
     * - The chain is not in a stuck state (has recently received a fresh data from the Notaries).
     * @param agentRoot_    New Agent Merkle Root that is proposed to be set
     */
    function proposeAgentRootWhenStuck(bytes32 agentRoot_) external;

    /**
     * @notice Allows contract owner to cancel the previously proposed agent root.
     * @dev Will revert if any of the following conditions is met:
     * - Caller is not the contract owner.
     * - No agent root was proposed.
     */
    function cancelProposedAgentRoot() external;

    /**
     * @notice Allows contract owner to resolve the previously proposed agent root.
     * This will update the agent root, allowing the agents to update their status, effectively
     * resolving the "stuck" chain.
     * @dev Will revert if any of the following conditions is met:
     * - Caller is not the contract owner.
     * - No agent root was proposed.
     * - Not enough time has passed since the agent root was proposed.
     */
    function resolveProposedAgentRoot() external;

    /**
     * @notice Withdraws locked base message tips from local Origin to the recipient.
     * @dev Could only be remote-called by BondingManager contract on Synapse Chain.
     * Note: as an extra security check this function returns its own selector, so that
     * Destination could verify that a "remote" function was called when executing a manager message.
     * @param recipient     Address to withdraw tips to
     * @param amount        Tips value to withdraw
     */
    function remoteWithdrawTips(uint32 msgOrigin, uint256 proofMaturity, address recipient, uint256 amount)
        external
        returns (bytes4 magicValue);

    // ═══════════════════════════════════════════════════ VIEWS ═══════════════════════════════════════════════════════

    /**
     * @notice Returns the latest proposed agent root and the timestamp when it was proposed.
     * @dev Will return zero values if no agent root was proposed, or if the proposed agent root
     * was already resolved.
     */
    function proposedAgentRootData() external view returns (bytes32 agentRoot_, uint256 proposedAt_);
}
