// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

abstract contract MessageBusEvents {
    enum TxStatus {
        Null,
        Success,
        Fail
    }

    event Executed(
        bytes32 indexed messageId, TxStatus status, address indexed dstAddress, uint64 srcChainId, uint64 srcNonce
    );

    event MessageSent(
        address indexed sender,
        uint256 srcChainID,
        bytes32 receiver,
        uint256 indexed dstChainId,
        bytes message,
        uint64 nonce,
        bytes options,
        uint256 fee,
        bytes32 indexed messageId
    );

    event GasBufferSet(uint64 gasBuffer);
    event MessageLengthEstimateSet(uint256 length);
}
