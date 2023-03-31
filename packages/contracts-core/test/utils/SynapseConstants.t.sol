// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

contract SynapseConstants {
    // ============ Domains ============
    uint256 internal constant DOMAINS = 3;
    uint32 internal constant DOMAIN_LOCAL = 1000;
    uint32 internal constant DOMAIN_REMOTE = 1500;
    // TODO: replace placeholder value
    uint32 internal constant DOMAIN_SYNAPSE = 10;
    // ============ Actors ============
    uint256 internal constant NOTARIES_PER_CHAIN = 4;
    uint256 internal constant GUARDS = 4;
    // ============ App ============
    uint32 internal constant APP_OPTIMISTIC_SECONDS = 60;
    // ============ Merkle ============
    uint256 internal constant ORIGIN_TREE_HEIGHT = 32;
    // ============ Message ============
    // Maximum bytes per message = 2 KiB
    // (somewhat arbitrarily set to begin)
    uint256 public constant MAX_CONTENT_BYTES = 2 * 2**10;
    // ============ Bonding ============
    uint256 public constant BONDING_OPTIMISTIC_PERIOD = 1 days;
}
