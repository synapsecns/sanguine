// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import { AgentStatus } from "../libs/Structures.sol";

interface IAgentManager {
    /**
     * @notice Local SystemRegistry should call this function to indicate that the agent
     * is proven to commit fraud in the SystemRegistry.
     * @dev On Synapse Chain this initiates the process of agent slashing. It could be immediately
     * completed by anyone calling completeSlashing() providing a correct merkle proof
     * for the OLD agent status.
     * @param _domain   Domain where the slashed agent was active
     * @param _agent    Address of the slashed Agent
     * @param _prover   Address that initially provided fraud proof in SystemRegistry
     */
    function registrySlash(
        uint32 _domain,
        address _agent,
        address _prover
    ) external;

    // ═════════════════════════════════ VIEWS ═════════════════════════════════

    /**
     * @notice Returns the latest known root of the Agent Merkle Tree.
     */
    function agentRoot() external view returns (bytes32);

    /**
     * @notice Returns (flag, domain, index) for a given agent. See Structures.sol for details.
     * @dev Will return AgentFlag.Fraudulent for agents that have been proven to commit fraud,
     * but their status is not updated to Slashed yet.
     * @param _agent    Agent address
     * @return          Status for the given agent: (flag, domain, index).
     */
    function agentStatus(address _agent) external view returns (AgentStatus memory);

    /**
     * @notice Returns whether the agent has been slashed.
     * @param _agent        Agent address
     * @return isSlashed    Whether the agent has been slashed
     * @return prover       Address that presented the proof of fraud committed by the agent
     */
    function slashStatus(address _agent) external view returns (bool isSlashed, address prover);
}
