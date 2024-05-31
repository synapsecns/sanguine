// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

abstract contract InterchainClientV1Events {
    /// @notice Emitted when the default Guard module is set.
    /// @param guard    The address of the Guard module that will be used by default.
    event DefaultGuardSet(address guard);

    /// @notice Emitted when the default Module is set.
    /// @param module   The address of the Module that will be used by default.
    event DefaultModuleSet(address module);

    /// @notice Emitted when the InterchainClientV1 deployment on a remote chain is linked.
    /// @param chainId   The chain ID of the remote chain.
    /// @param client    The address of the InterchainClientV1 deployment on the remote chain.
    event LinkedClientSet(uint64 chainId, bytes32 client);

    /// @notice Emitted when a new interchain transaction is sent through the InterchainClientV1.
    /// The Receiver on the destination chain will receive the specified message once the transaction is executed.
    /// @param transactionId    The unique identifier of the interchain transaction.
    /// @param dbNonce          The nonce of entry containing the transaction.
    /// @param dstChainId       The chain ID of the destination chain.
    /// @param srcSender        The sender of the transaction on the source chain.
    /// @param dstReceiver      The receiver of the transaction on the destination chain.
    /// @param verificationFee  The fee paid to verify the entry on the destination chain.
    /// @param executionFee     The fee paid to execute the transaction on the destination chain.
    /// @param options          The execution options for the transaction.
    /// @param message          The payload of the message being sent.
    event InterchainTransactionSent(
        bytes32 indexed transactionId,
        uint64 dbNonce,
        uint64 dstChainId,
        bytes32 indexed srcSender,
        bytes32 indexed dstReceiver,
        uint256 verificationFee,
        uint256 executionFee,
        bytes options,
        bytes message
    );

    /// @notice Emitted when an interchain transaction is received by the InterchainClientV1.
    /// The Receiver on the destination chain has just received the message sent from the source chain.
    /// @param transactionId    The unique identifier of the interchain transaction.
    /// @param dbNonce          The nonce of entry containing the transaction.
    /// @param srcChainId       The chain ID of the source chain.
    /// @param srcSender        The sender of the transaction on the source chain.
    /// @param dstReceiver      The receiver of the transaction on the destination chain.
    event InterchainTransactionReceived(
        bytes32 indexed transactionId,
        uint64 dbNonce,
        uint64 srcChainId,
        bytes32 indexed srcSender,
        bytes32 indexed dstReceiver
    );

    /// @notice Emitted when the proof of execution is written to InterchainDB. This allows the source chain
    /// to verify that the transaction was executed by a specific executor, if necessary.
    /// @param transactionId    The unique identifier of the interchain transaction.
    /// @param dbNonce          The nonce of entry containing the transaction.
    /// @param executor         The address of the executor that completed the transaction.
    event ExecutionProofWritten(bytes32 indexed transactionId, uint64 dbNonce, address indexed executor);
}
