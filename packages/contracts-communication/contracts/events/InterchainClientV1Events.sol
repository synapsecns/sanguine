// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

abstract contract InterchainClientV1Events {
    // TODO: figure out indexing

    event InterchainTransactionSent(
        bytes32 indexed transactionId,
        uint256 indexed dbNonce,
        uint256 dstChainId,
        bytes32 srcSender,
        bytes32 dstReceiver,
        uint256 verificationFee,
        uint256 executionFee,
        bytes options,
        bytes message
    );

    event InterchainTransactionReceived(
        bytes32 indexed transactionId,
        uint256 indexed dbNonce,
        uint256 srcChainId,
        bytes32 srcSender,
        bytes32 dstReceiver
    );
}
