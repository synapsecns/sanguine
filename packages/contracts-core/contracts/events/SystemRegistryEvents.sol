// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

abstract contract SystemRegistryEvents {
    /**
     * @notice Emitted when an Agent is slashed.
     * @param domain    Domain where a slashed Agent was active
     * @param agent     Address of the slashed agent
     * @param prover    Account that supplied proof leading to agent slashing
     */
    event AgentSlashed(uint32 indexed domain, address indexed agent, address prover);
}
