// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

interface IExecutionService {
    /// @notice Request the execution of an Interchain Transaction on a remote chain.
    /// Note: the off-chain actor needs to fetch the transaction payload from the InterchainClient
    /// event with the same transactionId, then execute the transaction on the remote chain:
    /// `dstInterchainClient.executeTransaction(transactionPayload)`
    /// Once the execution is confirmed on the source chain, the off-chain actor will be able
    /// to claim `executionFee` in the ExecutionFees contract.
    /// @dev Could only be called by `InterchainClient` contracts.
    /// Will revert if the execution fee is not big enough.
    /// @param dstChainId           The chain id of the destination chain.
    /// @param txPayloadSize        The size of the transaction payload to use for the execution.
    /// @param transactionId        The id of the transaction to execute.
    /// @param executionFee         The fee paid for the execution.
    /// @param options              The options to use for the execution.
    function requestExecution(
        uint256 dstChainId,
        uint256 txPayloadSize,
        bytes32 transactionId,
        uint256 executionFee,
        bytes memory options
    )
        external;

    /// @notice Get the execution fee for executing an Interchain Transaction on a remote chain.
    /// @param dstChainId           The chain id of the destination chain.
    /// @param txPayloadSize        The size of the transaction payload to use for the execution.
    /// @param options              The options to use for the execution.
    function getExecutionFee(
        uint256 dstChainId,
        uint256 txPayloadSize,
        bytes memory options
    )
        external
        view
        returns (uint256);
}
