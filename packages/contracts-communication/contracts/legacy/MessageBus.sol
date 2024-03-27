// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {MessageBusEvents} from "./events/MessageBusEvents.sol";
import {IMessageBus} from "./interfaces/IMessageBus.sol";
import {ILegacyReceiver} from "./interfaces/ILegacyReceiver.sol";
import {LegacyMessageLib} from "./libs/LegacyMessage.sol";
import {LegacyOptionsLib} from "./libs/LegacyOptions.sol";

import {ICAppV1, OptionsV1} from "../apps/ICAppV1.sol";
import {TypeCasts} from "../libs/TypeCasts.sol";

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
        // TODO: do we want to include dbNonce/entryIndex here to enforce non-replayability?
        bytes memory encodedLegacyMsg = LegacyMessageLib.encodeLegacyMessage({
            srcSender: msg.sender,
            dstReceiver: dstReceiver,
            srcNonce: cachedNonce,
            message: message
        });
        _sendToLinkedApp({
            dstChainId: dstChainId,
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
        return estimateFeeExact(dstChainId, messageLengthEstimate, options);
    }

    /// @inheritdoc IMessageBus
    function estimateFeeExact(
        uint256 dstChainId,
        uint256 messageLen,
        bytes calldata options
    )
        public
        view
        returns (uint256)
    {
        return _getMessageFee({dstChainId: dstChainId, options: _icOptionsV1(options), message: new bytes(messageLen)});
    }

    /// @dev Internal logic for receiving messages. At this point the validity of the message is already checked.
    function _receiveMessage(
        uint256 srcChainId,
        bytes32 sender,
        uint256 dbNonce,
        uint64 entryIndex,
        bytes calldata encodedLegacyMsg
    )
        internal
        override
    {
        (address srcSender, address dstReceiver, uint64 srcNonce, bytes memory message) =
            LegacyMessageLib.decodeLegacyMessage(encodedLegacyMsg);
        // TODO: do we want to enforce non-replayability here, or do we rely on InterchainClient?
        ILegacyReceiver(dstReceiver).executeMessage({
            srcAddress: TypeCasts.addressToBytes32(srcSender),
            srcChainId: srcChainId,
            message: message,
            // TODO: this is Interchain Client address. Do we need executor EOA instead (tx.origin)?
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

    function _icOptionsV1(bytes calldata options) internal view returns (OptionsV1 memory) {
        return OptionsV1({gasLimit: LegacyOptionsLib.decodeLegacyOptions(options), gasAirdrop: 0});
    }
}
