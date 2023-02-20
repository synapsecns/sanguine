// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

// Here we define common constants for Structures Libraries
// to enable their easier reusing later.

/// @dev See State.sol: (bytes32,uint32,uint32,uint40,uint40): 32+4+4+5+5
uint256 constant STATE_LENGTH = 50;

/// @dev Maximum amount of states in a single snapshot
uint256 constant SNAPSHOT_MAX_STATES = 32;
