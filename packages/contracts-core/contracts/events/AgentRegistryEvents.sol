// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

abstract contract AgentRegistryEvents {
    /*
     * @notice Emitted when a new Agent is added.
     * @param domain    Domain where a Agent was added
     * @param account   Address of the added agent
     */
    event AgentAdded(uint32 indexed domain, address indexed account);

    /**
     * @notice Emitted when a new Agent is removed.
     * @param domain    Domain where a Agent was removed
     * @param account   Address of the removed agent
     */
    event AgentRemoved(uint32 indexed domain, address indexed account);

    /**
     * @notice Emitted when the first agent is added for the domain
     * @param domain    Domain where the first Agent was added
     */
    event DomainActivated(uint32 indexed domain);

    /**
     * @notice Emitted when the last agent is removed from the domain
     * @param domain    Domain where the last Agent was removed
     */
    event DomainDeactivated(uint32 indexed domain);
}
