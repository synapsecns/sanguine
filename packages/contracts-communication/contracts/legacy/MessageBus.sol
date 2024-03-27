// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {MessageBusEvents} from "./events/MessageBusEvents.sol";
import {IMessageBus} from "./interfaces/IMessageBus.sol";

import {ICAppV1} from "../apps/ICAppV1.sol";

contract MessageBus is ICAppV1, MessageBusEvents, IMessageBus {
    uint256 public messageLengthEstimate;
    uint64 public nonce;

    constructor(address admin) ICAppV1(admin) {}

    /// @inheritdoc IMessageBus
    function sendMessage(
        bytes32 receiver,
        uint256 dstChainId,
        bytes calldata message,
        bytes calldata options
    )
        external
        payable
    {
        // TODO: implement
    }

    /// @inheritdoc IMessageBus
    function setMessageLengthEstimate(uint256 length) external {
        // TODO: implement
    }

    /// @inheritdoc IMessageBus
    function estimateFee(uint256 dstChainId, bytes calldata options) external view returns (uint256) {
        // TODO: implement
    }

    /// @inheritdoc IMessageBus
    function estimateFeeExact(
        uint256 dstChainId,
        uint256 messageLen,
        bytes calldata options
    )
        external
        view
        returns (uint256)
    {
        // TODO: implement
    }

    /// @dev Internal logic for receiving messages. At this point the validity of the message is already checked.
    function _receiveMessage(
        uint256 srcChainId,
        bytes32 sender,
        uint256 dbNonce,
        uint64 entryIndex,
        bytes calldata message
    )
        internal
        override
    {
        // TODO: implement
    }
}
