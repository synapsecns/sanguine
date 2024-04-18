// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

interface IExecutionFees {
    error ExecutionFees__AlreadyRecorded(uint64 dstChainId, bytes32 transactionId, address executor);
    error ExecutionFees__ZeroAddress();
    error ExecutionFees__ZeroAmount();

    /// @notice Add the execution fee for a transaction. The attached value will be added to the
    /// rewards for the executor completing the transaction.
    /// Note: this could be used to store the execution fee for a new transaction, or to add more
    /// funds to the execution fee of an existing transaction. Therefore this function is payable,
    /// and does not implement any caller restrictions.
    /// @dev Will revert if the executor is already recorded for the transaction.
    /// @param dstChainId           The chain id of the destination chain.
    /// @param transactionId        The id of the transaction to add the execution fee to.
    function addExecutionFee(uint64 dstChainId, bytes32 transactionId) external payable;

    /// @notice Record the executor (who completed the transaction) for a transaction,
    /// and update the accumulated rewards for the executor.
    /// @dev Could only be called by the Recorder.
    /// @param dstChainId           The chain id of the destination chain.
    /// @param transactionId        The id of the transaction to record the executor for.
    /// @param executor             The address of the executor who completed the transaction.
    function recordExecutor(uint64 dstChainId, bytes32 transactionId, address executor) external;

    /// @notice Allows the executor to claim their unclaimed rewards.
    /// @dev Will revert if the executor has no unclaimed rewards.
    function claimExecutionFees(address executor) external;

    // ═══════════════════════════════════════════════════ VIEWS ═══════════════════════════════════════════════════════

    /// @notice Get the accumulated rewards for an executor.
    /// @param executor             The address of the executor to get the rewards for.
    function accumulatedRewards(address executor) external view returns (uint256 accumulated);

    /// @notice Get the unclaimed rewards for an executor.
    /// @param executor             The address of the executor to get the rewards for.
    function unclaimedRewards(address executor) external view returns (uint256 unclaimed);

    /// @notice Get the total execution fee for a transaction.
    /// @param dstChainId           The chain id of the destination chain.
    /// @param transactionId        The id of the transaction to get the execution fee for.
    function executionFee(uint64 dstChainId, bytes32 transactionId) external view returns (uint256 fee);

    /// @notice Get the address of the recorded executor for a transaction.
    /// @dev Will return address(0) if the executor is not recorded.
    /// @param dstChainId           The chain id of the destination chain.
    /// @param transactionId        The id of the transaction to get the recorded executor for.
    function recordedExecutor(uint64 dstChainId, bytes32 transactionId) external view returns (address executor);
}
