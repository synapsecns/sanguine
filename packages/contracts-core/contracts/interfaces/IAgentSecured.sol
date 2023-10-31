// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {AgentStatus, DisputeStatus} from "../libs/Structures.sol";

interface IAgentSecured {
    /**
     * @notice Local AgentManager should call this function to indicate that a dispute
     * between a Guard and a Notary has been opened.
     * @param guardIndex    Index of the Guard in the Agent Merkle Tree
     * @param notaryIndex   Index of the Notary in the Agent Merkle Tree
     */
    function openDispute(uint32 guardIndex, uint32 notaryIndex) external;

    /**
     * @notice Local AgentManager should call this function to indicate that a dispute
     * has been resolved due to one of the agents being slashed.
     * > `rivalIndex` will be ZERO, if the slashed agent was not in the Dispute.
     * @param slashedIndex  Index of the slashed agent in the Agent Merkle Tree
     * @param rivalIndex    Index of the their Dispute Rival in the Agent Merkle Tree
     */
    function resolveDispute(uint32 slashedIndex, uint32 rivalIndex) external;

    // ═══════════════════════════════════════════════════ VIEWS ═══════════════════════════════════════════════════════

    /**
     * @notice Returns the address of the local AgentManager contract, which is treated as
     * the "source of truth" for agent statuses.
     */
    function agentManager() external view returns (address);

    /**
     * @notice Returns the address of the local Inbox contract, which is treated as
     * the "source of truth" for agent-signed statements.
     * @dev Inbox passes verified agent statements to `IAgentSecured` contract.
     */
    function inbox() external view returns (address);

    /**
     * @notice Returns (flag, domain, index) for a given agent. See Structures.sol for details.
     * @dev Will return AgentFlag.Fraudulent for agents that have been proven to commit fraud,
     * but their status is not updated to Slashed yet.
     * @param agent     Agent address
     * @return          Status for the given agent: (flag, domain, index).
     */
    function agentStatus(address agent) external view returns (AgentStatus memory);

    /**
     * @notice Returns agent address and their current status for a given agent index.
     * @dev Will return empty values if agent with given index doesn't exist.
     * @param index     Agent index in the Agent Merkle Tree
     * @return agent    Agent address
     * @return status   Status for the given agent: (flag, domain, index)
     */
    function getAgent(uint256 index) external view returns (address agent, AgentStatus memory status);

    /**
     * @notice Returns (flag, openedAt, resolvedAt) that describes the latest status of
     * the latest dispute for an agent with a given index.
     * @dev Will return empty values if agent with given index doesn't exist.
     * @param agentIndex    Agent index in the Agent Merkle Tree
     * @return              Latest dispute status for the given agent: (flag, openedAt, resolvedAt)
     */
    function latestDisputeStatus(uint32 agentIndex) external view returns (DisputeStatus memory);
}
