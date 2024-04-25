// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

interface IExecutionService {
    /// @notice Request the execution of an Interchain Transaction on a remote chain in exchange for
    /// the execution fee, attached to the transaction as `msg.value`.
    /// Note: the off-chain actor needs to fetch the transaction payload from the InterchainClient
    /// event with the same transactionId, then execute the transaction on the remote chain:
    /// `dstInterchainClient.executeTransaction(transactionPayload)`
    /// @dev Could only be called by `InterchainClient` contracts.
    /// Will revert if the execution fee is not big enough.
    /// @param dstChainId           The chain id of the destination chain.
    /// @param txPayloadSize        The size of the transaction payload to use for the execution.
    /// @param transactionId        The id of the transaction to execute.
    /// @param options              The options to use for the execution.
    function requestTxExecution(
        uint64 dstChainId,
        uint256 txPayloadSize,
        bytes32 transactionId,
        bytes memory options
    )
        external
        payable;

    /// @notice Get the address of the EOA account that will be used to execute transactions on the
    /// remote chains.
    function executorEOA() external view returns (address);

    /// @notice Get the execution fee for executing an Interchain Transaction on a remote chain.
    /// @param dstChainId           The chain id of the destination chain.
    /// @param txPayloadSize        The size of the transaction payload to use for the execution.
    /// @param options              The options to use for the execution.
    function getExecutionFee(
        uint64 dstChainId,
        uint256 txPayloadSize,
        bytes memory options
    )
        external
        view
        returns (uint256);
}
