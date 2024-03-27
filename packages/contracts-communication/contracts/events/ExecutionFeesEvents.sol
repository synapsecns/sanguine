// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

abstract contract ExecutionFeesEvents {
    event ExecutionFeeAdded(uint256 dstChainId, bytes32 indexed transactionId, uint256 totalFee);
    event ExecutorRecorded(uint256 dstChainId, bytes32 indexed transactionId, address indexed executor);

    event ExecutionFeesAwarded(address indexed executor, uint256 amount);
    event ExecutionFeesClaimed(address indexed executor, uint256 amount);
}
