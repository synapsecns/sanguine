// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {AgentStatus, DisputeFlag} from "../libs/Structures.sol";

interface IAgentManager {
    /**
     * @notice Allows Inbox to open a Dispute between a Guard and a Notary, if they are both not in Dispute already.
     * > Will revert if any of these is true:
     * > - Caller is not Inbox.
     * > - Guard or Notary is already in Dispute.
     * @param guardIndex    Index of the Guard in the Agent Merkle Tree
     * @param notaryIndex   Index of the Notary in the Agent Merkle Tree
     */
    function openDispute(uint32 guardIndex, uint32 notaryIndex) external;

    /**
     * @notice Allows Inbox to slash an agent, if their fraud was proven.
     * > Will revert if any of these is true:
     * > - Caller is not Inbox.
     * > - Domain doesn't match the saved agent domain.
     * @param domain    Domain where the Agent is active
     * @param agent     Address of the Agent
     * @param prover    Address that initially provided fraud proof
     */
    function slashAgent(uint32 domain, address agent, address prover) external;

    // ═══════════════════════════════════════════════════ VIEWS ═══════════════════════════════════════════════════════

    /**
     * @notice Returns the latest known root of the Agent Merkle Tree.
     */
    function agentRoot() external view returns (bytes32);

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
     * @notice Returns the number of opened Disputes.
     * @dev This includes the Disputes that have been resolved already.
     */
    function getDisputesAmount() external view returns (uint256);

    /**
     * @notice Returns information about the dispute with the given index.
     * @dev Will revert if dispute with given index hasn't been opened yet.
     * @param index             Dispute index
     * @return guard            Address of the Guard in the Dispute
     * @return notary           Address of the Notary in the Dispute
     * @return slashedAgent     Address of the Agent who was slashed when Dispute was resolved
     * @return fraudProver      Address who provided fraud proof to resolve the Dispute
     * @return reportPayload    Raw payload with report data that led to the Dispute
     * @return reportSignature  Guard signature for the report payload
     */
    function getDispute(uint256 index)
        external
        view
        returns (
            address guard,
            address notary,
            address slashedAgent,
            address fraudProver,
            bytes memory reportPayload,
            bytes memory reportSignature
        );

    /**
     * @notice Returns the current Dispute status of a given agent. See Structures.sol for details.
     * @dev Every returned value will be set to zero if agent was not slashed and is not in Dispute.
     * `rival` and `disputePtr` will be set to zero if the agent was slashed without being in Dispute.
     * @param agent         Agent address
     * @return flag         Flag describing the current Dispute status for the agent: None/Pending/Slashed
     * @return rival        Address of the rival agent in the Dispute
     * @return fraudProver  Address who provided fraud proof to resolve the Dispute
     * @return disputePtr   Index of the opened Dispute PLUS ONE. Zero if agent is not in Dispute.
     */
    function disputeStatus(address agent)
        external
        view
        returns (DisputeFlag flag, address rival, address fraudProver, uint256 disputePtr);
}
