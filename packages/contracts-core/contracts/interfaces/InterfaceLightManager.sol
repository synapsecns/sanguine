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
     * by setting the agent root.
     * This could only be called if no fresh data has been submitted by the Notaries to the Inbox,
     * indicating that the chain is stuck for one of the reasons:
     * - All active Notaries are in Dispute.
     * - No active Notaries exist under the current agent root.
     * @param agentRoot_    New Agent Merkle Root
     */
    function setAgentRootWhenStuck(bytes32 agentRoot_) external;

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
}
