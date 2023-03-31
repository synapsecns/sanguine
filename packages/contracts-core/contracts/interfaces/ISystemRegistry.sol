// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {AgentStatus} from "../libs/Structures.sol";

interface ISystemRegistry {
    /**
     * @notice Local AgentManager should call this function to indicate that the agent
     * has been slashed, either on local or remote chain.
     * @param domain    Domain where the slashed agent was active
     * @param agent     Address of the slashed Agent
     * @param prover    Account that supplied proof leading to agent slashing
     */
    function managerSlash(uint32 domain, address agent, address prover) external;

    /**
     * @notice Returns (flag, domain, index) for a given agent. See Structures.sol for details.
     * @param agent     Agent address
     * @return          Status for the given agent: (flag, domain, index).
     */
    function agentStatus(address agent) external view returns (AgentStatus memory);
}
