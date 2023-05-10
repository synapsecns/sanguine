// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

// Here we define common constants to enable their easier reusing later.

// ══════════════════════════════════ MERKLE ═══════════════════════════════════
/// @dev Height of the Agent Merkle Tree
uint256 constant AGENT_TREE_HEIGHT = 32;
/// @dev Height of the Origin Merkle Tree
uint256 constant ORIGIN_TREE_HEIGHT = 32;
/// @dev Height of the Snapshot Merkle Tree. Allows up to 64 leafs, e.g. up to 32 states
uint256 constant SNAPSHOT_TREE_HEIGHT = 6;
// ══════════════════════════════════ STRUCTS ══════════════════════════════════
/// @dev See Attestation.sol: (bytes32,bytes32,uint32,uint40,uint40): 32+32+4+5+5
uint256 constant ATTESTATION_LENGTH = 78;
/// @dev See Receipt.sol: (uint32,uint32,bytes32,bytes32,uint8,address,address,address,tips): 4+4+32+32+1+20+20+20
uint256 constant GAS_DATA_LENGTH = 12;
/// @dev See Request.sol: (uint64,uint96): 8+12
uint256 constant RECEIPT_BODY_LENGTH = 133;
/// @dev See Receipt.sol: body + tips
uint256 constant RECEIPT_LENGTH = RECEIPT_BODY_LENGTH + TIPS_LENGTH;
/// @dev See State.sol: (bytes32,uint32,uint32,uint40,uint40,GasData): 32+4+4+5+5+len(GasData)
uint256 constant STATE_LENGTH = 50 + GAS_DATA_LENGTH;
/// @dev Maximum amount of states in a single snapshot. Each state produces two leafs in the tree
uint256 constant SNAPSHOT_MAX_STATES = 1 << (SNAPSHOT_TREE_HEIGHT - 1);
// ══════════════════════════════════ MESSAGE ══════════════════════════════════
/// @dev See Header.sol: (uint32,uint32,uint32,uint32): 4+4+4+4
uint256 constant HEADER_LENGTH = 16;
/// @dev See GasData.sol: (uint16,uint16,uint16,uint16,uint16,uint16): 2+2+2+2+2+2
uint256 constant REQUEST_LENGTH = 20;
/// @dev See Tips.sol: (uint64,uint64,uint64,uint64): 8+8+8+8
uint256 constant TIPS_LENGTH = 32;
/// @dev The amount of discarded last bits when encoding tip values
uint256 constant TIPS_GRANULARITY = 32;
/// @dev Tip values could be only the multiples of TIPS_MULTIPLIER
uint256 constant TIPS_MULTIPLIER = 1 << TIPS_GRANULARITY;
// ══════════════════════════════ STATEMENT SALTS ══════════════════════════════
/// @dev Salts for signing various statements
bytes32 constant ATTESTATION_SALT = keccak256("ATTESTATION_SALT");
bytes32 constant ATTESTATION_REPORT_SALT = keccak256("ATTESTATION_REPORT_SALT");
bytes32 constant RECEIPT_SALT = keccak256("RECEIPT_SALT");
bytes32 constant RECEIPT_REPORT_SALT = keccak256("RECEIPT_REPORT_SALT");
bytes32 constant SNAPSHOT_SALT = keccak256("SNAPSHOT_SALT");
bytes32 constant STATE_REPORT_SALT = keccak256("STATE_REPORT_SALT");
// ═════════════════════════════════ PROTOCOL ══════════════════════════════════
/// @dev Optimistic period for new agent roots in LightManager
uint32 constant AGENT_ROOT_OPTIMISTIC_PERIOD = 1 days;
uint32 constant BONDING_OPTIMISTIC_PERIOD = 1 days;
/// @dev Amount of time without fresh data from Notaries before contract owner can resolve stuck disputes manually
uint256 constant FRESH_DATA_TIMEOUT = 4 hours;
/// @dev Maximum bytes per message = 2 KiB (somewhat arbitrarily set to begin)
uint256 constant MAX_CONTENT_BYTES = 2 * 2 ** 10;
/// @dev Domain of the Synapse Chain
// TODO: replace the placeholder with actual value (for MVP this is Optimism chainId)
uint32 constant SYNAPSE_DOMAIN = 10;
