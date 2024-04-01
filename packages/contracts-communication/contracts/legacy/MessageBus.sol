// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {MessageBusEvents} from "./events/MessageBusEvents.sol";
import {IMessageBus} from "./interfaces/IMessageBus.sol";
import {ILegacyReceiver} from "./interfaces/ILegacyReceiver.sol";
import {LegacyMessageLib} from "./libs/LegacyMessage.sol";
import {LegacyOptionsLib} from "./libs/LegacyOptions.sol";

import {ICAppV1, OptionsV1} from "../apps/ICAppV1.sol";
import {TypeCasts} from "../libs/TypeCasts.sol";

import {SafeCast} from "@openzeppelin/contracts/utils/math/SafeCast.sol";

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
        address dstReceiver = TypeCasts.bytes32ToAddress(receiver);
        if (TypeCasts.addressToBytes32(dstReceiver) != receiver) {
            revert MessageBus__NotEVMReceiver(receiver);
        }
        uint64 cachedNonce = nonce++;
        // Note: we are using the internal nonce here to generate the unique message ID.
        // This is used for tracking purposes and is not used for replay protection.
        // Instead, we rely on the Interchain Client to provide replay protection.
        bytes memory encodedLegacyMsg = LegacyMessageLib.encodeLegacyMessage({
            srcSender: msg.sender,
            dstReceiver: dstReceiver,
            srcNonce: cachedNonce,
            message: message
        });
        _sendToLinkedApp({
            dstChainId: SafeCast.toUint64(dstChainId),
            messageFee: msg.value,
            options: _icOptionsV1(options),
            message: encodedLegacyMsg
        });
        emit MessageSent({
            sender: msg.sender,
            srcChainID: block.chainid,
            receiver: receiver,
            dstChainId: dstChainId,
            message: message,
            nonce: cachedNonce,
            options: options,
            fee: msg.value,
            messageId: keccak256(encodedLegacyMsg)
        });
    }

    /// @inheritdoc IMessageBus
    function setMessageLengthEstimate(uint256 length) external onlyRole(IC_GOVERNOR_ROLE) {
        messageLengthEstimate = length;
        emit MessageLengthEstimateSet(length);
    }

    /// @inheritdoc IMessageBus
    function estimateFee(uint256 dstChainId, bytes calldata options) external view returns (uint256) {
        return estimateFeeExact(dstChainId, options, messageLengthEstimate);
    }

    /// @inheritdoc IMessageBus
    function estimateFeeExact(
        uint256 dstChainId,
        bytes calldata options,
        uint256 messageLen
    )
        public
        view
        returns (uint256)
    {
        return _getMessageFee({
            dstChainId: SafeCast.toUint64(dstChainId),
            options: _icOptionsV1(options),
            messageLen: messageLen
        });
    }

    /// @dev Internal logic for receiving messages. At this point the validity of the message is already checked.
    function _receiveMessage(
        uint64 srcChainId,
        bytes32, // sender
        uint64, // dbNonce
        uint64, // entryIndex
        bytes calldata encodedLegacyMsg
    )
        internal
        override
    {
        (address srcSender, address dstReceiver, uint64 srcNonce, bytes memory message) =
            LegacyMessageLib.decodeLegacyMessage(encodedLegacyMsg);
        // Note: we rely on the Interchain Client to provide replay protection.
        ILegacyReceiver(dstReceiver).executeMessage({
            srcAddress: TypeCasts.addressToBytes32(srcSender),
            srcChainId: srcChainId,
            message: message,
            executor: msg.sender
        });
        emit Executed({
            messageId: keccak256(encodedLegacyMsg),
            status: TxStatus.Success,
            dstAddress: dstReceiver,
            srcChainId: uint64(srcChainId),
            srcNonce: srcNonce
        });
    }

    /// @dev Convert legacy MessageBus options to Interchain OptionsV1.
    function _icOptionsV1(bytes calldata options) internal pure returns (OptionsV1 memory) {
        return OptionsV1({gasLimit: LegacyOptionsLib.decodeLegacyOptions(options), gasAirdrop: 0});
    }
}
