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
    /// @notice Default message length value for backwards compatible fee estimation.
    uint256 public messageLengthEstimate;
    /// @notice Nonce for generating unique message IDs: incremented for each message sent.
    uint64 public nonce;
    /// @notice Gas buffer to be added to the gas limit of the cross-chain message.
    uint64 public gasBuffer = 20_000;

    constructor(address admin) ICAppV1(admin) {}

    /// @notice Sends a message to a receiving contract address on another chain.
    /// Sender must make sure that the message is unique and not a duplicate message.
    /// @dev Legacy MessageBus only supports V1 of the options format, which specifies only the gas limit.
    /// @param receiver     The bytes32 address of the destination contract to be called
    /// @param dstChainId   The destination chain ID - typically, standard EVM chain ID, but differs on nonEVM chains
    /// @param message      The arbitrary payload to pass to the destination chain receiver
    /// @param options      Versioned struct used to instruct relayer on how to proceed with gas limits
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
            revert MessageBus__ReceiverNotEVM(receiver);
        }
        // Read the nonce and gas buffer from the same storage slot.
        (uint64 cachedNonce, uint64 cachedGasBuffer) = (nonce, gasBuffer);
        ++nonce;
        // Note: we are using the internal nonce here to generate the unique message ID.
        // This is used for tracking purposes and is not used for replay protection.
        // Instead, we rely on the Interchain Client to provide replay protection.
        bytes memory encodedLegacyMsg = LegacyMessageLib.encodeLegacyMessage({
            srcSender: msg.sender,
            dstReceiver: dstReceiver,
            srcNonce: cachedNonce,
            message: message
        });
        OptionsV1 memory icOptions = _icOptionsV1(options, cachedGasBuffer);
        _sendToLinkedApp({
            dstChainId: SafeCast.toUint64(dstChainId),
            messageFee: msg.value,
            options: icOptions,
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

    /// @notice Allows the Interchain Governor to set the gas buffer for sending the interchain messages.
    /// Note: The gas buffer is added to the gas limit requested by the sending app to cover the gas usage
    /// of the MessageBus contract on the destination chain.
    function setGasBuffer(uint64 gasBuffer_) external onlyRole(IC_GOVERNOR_ROLE) {
        gasBuffer = gasBuffer_;
        emit GasBufferSet(gasBuffer_);
    }

    /// @notice Allows the Interchain Governor to set the message length in bytes to be used for fee estimation.
    /// This does not affect the sendMessage function, but only the fee estimation.
    function setMessageLengthEstimate(uint256 length) external onlyRole(IC_GOVERNOR_ROLE) {
        messageLengthEstimate = length;
        emit MessageLengthEstimateSet(length);
    }

    /// @notice Returns srcGasToken fee to charge in wei for the cross-chain message based on the gas limit.
    /// @dev This function is using the preset message length to estimate the gas fee. This should cover most cases,
    /// if the message length is lower than the preset value. For more accurate fee estimation, use estimateFeeExact.
    /// @param dstChainId   The destination chain ID - typically, standard EVM chain ID, but differs on nonEVM chains
    /// @param options      Versioned struct used to instruct relayer on how to proceed with gas limits
    function estimateFee(uint256 dstChainId, bytes calldata options) external view returns (uint256) {
        return estimateFeeExact(dstChainId, options, messageLengthEstimate);
    }

    /// @notice Returns srcGasToken fee to charge in wei for the cross-chain message based on the message length
    /// and the gas limit.
    /// @param dstChainId   The destination chain ID - typically, standard EVM chain ID, but differs on nonEVM chains
    /// @param options      Versioned struct used to instruct relayer on how to proceed with gas limits
    /// @param messageLen   The length of the message to be sent in bytes
    function estimateFeeExact(
        uint256 dstChainId,
        bytes calldata options,
        uint256 messageLen
    )
        public
        view
        returns (uint256)
    {
        uint256 legacyMsgLen = LegacyMessageLib.payloadSize(messageLen);
        return _getMessageFee({
            dstChainId: SafeCast.toUint64(dstChainId),
            options: _icOptionsV1(options, gasBuffer),
            messageLen: legacyMsgLen
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

    /// @dev Convert legacy MessageBus options to Interchain OptionsV1. A gas buffer is added to reserve gas for the
    /// destination MessageBus to forward the received message and emit the event after execution. This is necessary
    /// to ensure that the App's requested gas limit is honored.
    function _icOptionsV1(bytes calldata options, uint64 sendGasBuffer) internal pure returns (OptionsV1 memory) {
        return OptionsV1({gasLimit: LegacyOptionsLib.decodeLegacyOptions(options) + sendGasBuffer, gasAirdrop: 0});
    }
}
