// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

abstract contract InterchainClientV1Events {
    event ExecutionFeesSet(address executionFees);
    event LinkedClientSet(uint256 chainId, bytes32 client);

    // TODO: figure out indexing

    event InterchainTransactionSent(
        bytes32 indexed transactionId,
        uint256 indexed dbNonce,
        uint64 indexed entryIndex,
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
        uint64 indexed entryIndex,
        uint256 srcChainId,
        bytes32 srcSender,
        bytes32 dstReceiver
    );

    event ExecutionProofWritten(
        bytes32 indexed transactionId, uint256 indexed dbNonce, uint64 indexed entryIndex, address executor
    );
}
