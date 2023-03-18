// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

// Here we define common enums and structures to enable their easier reusing later.

/// @dev Potential senders/recipients of a system message
enum SystemEntity {
    Origin,
    Destination,
    AgentManager
}

/**
 * @notice Unified struct for off-chain agent storing
 * @dev Both Guards and Notaries are stored this way.
 * `domain == 0` refers to Guards, who are active on every domain
 * `domain != 0` refers to Notaries, who are active on a single domain
 * @param bonded    Whether agent bonded or unbonded
 * @param domain    Domain, where agent is active
 * @param account   Off-chain agent address
 */
struct AgentInfo {
    // TODO: This won't be needed when Agents Merkle Tree is implemented
    uint32 domain;
    address account;
    bool bonded;
}
