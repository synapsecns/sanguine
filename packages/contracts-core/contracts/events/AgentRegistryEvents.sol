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
     * @notice Emitted when an Agent is removed.
     * @param domain    Domain where a removed Agent was active
     * @param account   Address of the removed agent
     */
    event AgentRemoved(uint32 indexed domain, address indexed account);

    /**
     * @notice Emitted when an Agent is slashed.
     * @param domain    Domain where a slashed Agent was active
     * @param account   Address of the slashed agent
     */
    event AgentSlashed(uint32 indexed domain, address indexed account);

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

    /**
     * @notice Emitted when a Dispute between a Guard and a Notary is initiated
     * by a Guard submitting a Report on invalid statement signed by a Notary.
     * @param guard     Address of the Guard who submitted a Report
     * @param domain    Domain where the Notary is active
     * @param notary    Address of the Notary who signed a reported statement
     */
    event Dispute(address guard, uint32 domain, address notary);
}
