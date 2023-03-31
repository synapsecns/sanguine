// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import { AgentStatus } from "../libs/Structures.sol";

interface ILightManager {
    /**
     * @notice Updates agent status, using a proof against the latest known Agent Merkle Root.
     * @dev Will revert if the provided proof doesn't match the latest merkle root.
     * @param agent     Agent address
     * @param status    Structure specifying agent status: (flag, domain, index)
     * @param proof     Merkle proof of Active status for the agent
     */
    function updateAgentStatus(
        address agent,
        AgentStatus memory status,
        bytes32[] memory proof
    ) external;

    /**
     * @notice Updates the root of Agent Merkle Tree that the Light Manager is tracking.
     * Could be only called by a local Destination contract, which is supposed to
     * verify the attested Agent Merkle Roots.
     * @param agentRoot     New Agent Merkle Root
     */
    function setAgentRoot(bytes32 agentRoot) external;
}
