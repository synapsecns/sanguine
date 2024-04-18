// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

abstract contract ExecutionFeesEvents {
    /// @notice Emitted when an execution fee is added for a transaction.
    /// Sum of all fees for a transaction will be awarded to the executor that completes the transaction.
    /// @param dstChainId       The chain ID of the transaction destination chain.
    /// @param transactionId    The unique identifier of the interchain transaction.
    /// @param totalFee         The total fee added to the transaction so far.
    event ExecutionFeeAdded(uint64 dstChainId, bytes32 indexed transactionId, uint256 totalFee);

    /// @notice Emitted when an executor is recorded for a transaction.
    /// The executor will be awarded the sum of all fees for the transaction.
    /// @param dstChainId       The chain ID of the transaction destination chain.
    /// @param transactionId    The unique identifier of the interchain transaction.
    /// @param executor         The address of the executor that completed the transaction.
    event ExecutorRecorded(uint64 dstChainId, bytes32 indexed transactionId, address indexed executor);

    /// @notice Emitted when execution fees are awarded to an executor.
    /// @param executor         The address of the executor that was awarded the fees.
    /// @param amount           The amount of fees awarded.
    event ExecutionFeesAwarded(address indexed executor, uint256 amount);

    /// @notice Emitted when execution fees are claimed by an executor.
    /// @param executor         The address of the executor that claimed the fees.
    /// @param amount           The amount of fees claimed.
    event ExecutionFeesClaimed(address indexed executor, uint256 amount);
}
