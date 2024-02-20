// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

abstract contract InterchainClientV1Events {
    // TODO: figure out indexing

    event InterchainTransactionSent(
        bytes32 indexed transactionId,
        uint256 indexed dbNonce,
        uint256 dstChainId,
        bytes32 srcSender,
        bytes32 dstReceiver
    );

    event InterchainExecutionRequested(
        bytes32 indexed transactionId, uint256 dstChainId, uint256 executionFee, bytes encodedTransaction
    );

    event InterchainOptionsV1(bytes32 indexed transactionId, uint256 gasLimit, uint256 gasAirdrop);

    event InterchainTransactionReceived(
        bytes32 indexed transactionId,
        uint256 indexed dbNonce,
        uint256 srcChainId,
        bytes32 srcSender,
        bytes32 dstReceiver
    );
}
