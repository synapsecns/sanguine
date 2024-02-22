// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

interface IExecutionFees {
    /// @notice Add the execution fee for a transaction. The attached value will be added to the
    /// rewards for the executor completing the transaction.
    /// Note: this could be used to store the execution fee for a new transaction, or to add more
    /// funds to the execution fee of an existing transaction. Therefore this function is payable,
    /// and does not implement any caller restrictions.
    /// @dev Will revert if the executor is already recorded for the transaction.
    /// @param dstChainId           The chain id of the destination chain.
    /// @param transactionId        The id of the transaction to add the execution fee to.
    function addExecutionFee(uint256 dstChainId, bytes32 transactionId) external payable;

    /// @notice Record the executor (who completed the transaction) for a transaction,
    /// and update the accumulated rewards for the executor.
    /// @dev Could only be called by the Recorder.
    /// @param dstChainId           The chain id of the destination chain.
    /// @param transactionId        The id of the transaction to record the executor for.
    /// @param executor             The address of the executor who completed the transaction.
    function recordExecutor(uint256 dstChainId, bytes32 transactionId, address executor) external;

    /// @notice Allows the executor to claim their unclaimed rewards.
    /// @dev Will revert if the executor has no unclaimed rewards.
    function claimExecutionFees(address executor) external;

    // ═══════════════════════════════════════════════════ VIEWS ═══════════════════════════════════════════════════════

    /// @notice Get the accumulated rewards for an executor.
    /// @param executor             The address of the executor to get the rewards for.
    function getAccumulatedRewards(address executor) external view returns (uint256 accumulated);

    /// @notice Get the unclaimed rewards for an executor.
    /// @param executor             The address of the executor to get the rewards for.
    function getUnclaimedRewards(address executor) external view returns (uint256 unclaimed);
}
