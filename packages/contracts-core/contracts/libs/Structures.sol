// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

// Here we define common enums and structures to enable their easier reusing later.

/// @dev Potential senders/recipients of a system message
enum SystemEntity {
    Origin,
    Destination,
    AgentManager
}

/// @dev Potential statuses for the off-chain bonded agent:
/// - Unknown: never provided a bond => signature not valid
/// - Active: has a bond in BondingManager => signature valid
/// - Unstaking: has a bond in BondingManager, initiated the unstaking => signature not valid
/// - Resting: used to have a bond in BondingManager, successfully unstaked => signature not valid
/// - Fraudulent: proven to commit fraud, value in Merkle Tree not updated => signature not valid
/// - Slashed: proven to commit fraud, value in Merkle Tree was updated => signature not valid
/// Unstaked agent could later be added back to THE SAME domain by staking a bond again.
/// Honest agent: Unknown -> Active -> unstaking -> Resting -> Active ...
/// Malicious agent: Unknown -> Active -> Fraudulent -> Slashed
/// Malicious agent: Unknown -> Active -> Unstaking -> Fraudulent -> Slashed
enum AgentFlag {
    Unknown,
    Active,
    Unstaking,
    Resting,
    Fraudulent,
    Slashed
}

/// @notice Struct for storing an agent in the BondingManager contract.
struct AgentStatus {
    AgentFlag flag;
    uint32 domain;
    uint32 index;
    // 184 bits available for tight packing
}

/// @notice Struct representing information about a slashed agent.
struct SlashStatus {
    bool isSlashed;
    address prover;
    // 88 bits available for tight packing
}
