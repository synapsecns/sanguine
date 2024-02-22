// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

abstract contract ExecutionFeesEvents {
    event ExecutionFeeAdded(uint256 dstChainId, bytes32 indexed transactionId, uint256 fee);
    event ExecutorRecorded(uint256 dstChainId, bytes32 indexed transactionId, address indexed executor);
    event ExecutionFeesClaimed(address indexed executor, uint256 amount);
}
