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
/// @dev See Attestation.sol: (bytes32,bytes32, uint32,uint40,uint40): 32+4+5+5
uint256 constant ATTESTATION_LENGTH = 78;
/// @dev See Receipt.sol: (uint32,uint32,bytes32,bytes32,uint8,address,address,address,tips): 4+4+32+32+1+20+20+20+tips
uint256 constant RECEIPT_LENGTH = 133 + TIPS_LENGTH;
/// @dev See State.sol: (bytes32,uint32,uint32,uint40,uint40): 32+4+4+5+5
uint256 constant STATE_LENGTH = 50;
/// @dev Maximum amount of states in a single snapshot. Each state produces two leafs in the tree
uint256 constant SNAPSHOT_MAX_STATES = 1 << (SNAPSHOT_TREE_HEIGHT - 1);
// ══════════════════════════════════ MESSAGE ══════════════════════════════════
/// @dev See Header.sol: (uint32,uint32,uint32,uint32): 4+4+4+4
uint256 constant HEADER_LENGTH = 16;
/// @dev See Request.sol: (uint64): 8
uint256 constant REQUEST_LENGTH = 8;
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
bytes32 constant SNAPSHOT_SALT = keccak256("SNAPSHOT_SALT");
bytes32 constant STATE_REPORT_SALT = keccak256("STATE_REPORT_SALT");
// ════════════════════════════════ DESTINATION ════════════════════════════════
uint256 constant AGENT_ROOT_OPTIMISTIC_PERIOD = 1 days;
// ══════════════════════════════════ ORIGIN ═══════════════════════════════════
/// @dev Maximum bytes per message = 2 KiB (somewhat arbitrarily set to begin)
uint256 constant MAX_CONTENT_BYTES = 2 * 2 ** 10;
// ═══════════════════════════════ SYSTEM ROUTER ═══════════════════════════════
/// @dev Custom address used for sending and receiving system messages.
/// - Origin will dispatch messages from SystemRouter as if they were "sent by this sender".
/// - Destination will reroute messages "sent to this recipient" to SystemRouter.
/// - As a result: only SystemRouter messages will have this value as both sender and recipient.
/// Note: all bits except for lower 20 bytes are set to 1.
/// Note: TypeCasts.bytes32ToAddress(SYSTEM_ROUTER) == address(0)
bytes32 constant SYSTEM_ROUTER = bytes32(type(uint256).max << 160);
