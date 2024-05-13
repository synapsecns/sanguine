// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

interface IMessageBus {
    error MessageBus__ReceiverNotEVM(bytes32 receiver);

    function sendMessage(
        bytes32 receiver,
        uint256 dstChainId,
        bytes memory message,
        bytes memory options
    )
        external
        payable;
    function setGasBuffer(uint64 gasBuffer_) external;
    function setMessageLengthEstimate(uint256 length) external;

    function messageLengthEstimate() external view returns (uint256);
    function estimateFee(uint256 dstChainId, bytes memory options) external view returns (uint256);
    function estimateFeeExact(
        uint256 dstChainId,
        bytes memory options,
        uint256 messageLen
    )
        external
        view
        returns (uint256);
    function nonce() external view returns (uint64);
}
