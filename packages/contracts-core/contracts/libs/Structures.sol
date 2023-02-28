// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

// Here we define common constants for Structures Libraries
// to enable their easier reusing later.

/// @dev See Attestation.sol: (bytes32,uint8,uint32,uint40,uint40): 32+1+4+5+5
uint256 constant ATTESTATION_LENGTH = 47;

/// @dev See State.sol: (bytes32,uint32,uint32,uint40,uint40): 32+4+4+5+5
uint256 constant STATE_LENGTH = 50;

/// @dev Maximum amount of states in a single snapshot
uint256 constant SNAPSHOT_MAX_STATES = 32;

/// @dev Root for an empty Origin Merkle Tree.
bytes32 constant EMPTY_ROOT = hex"27ae5ba08d7291c96c8cbddcc148bf48a6d68c7974b94356f53754ef6171d757";

/// @dev Depth of the Origin Merkle Tree
uint256 constant ORIGIN_TREE_DEPTH = 32;
