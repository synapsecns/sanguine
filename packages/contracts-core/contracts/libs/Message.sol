// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {BaseMessageLib} from "./BaseMessage.sol";
import {ByteString} from "./ByteString.sol";
import {HEADER_LENGTH} from "./Constants.sol";
import {Header, HeaderLib} from "./Header.sol";
import {MemView, MemViewLib} from "./MemView.sol";

/// Message is a memory over over a formatted message payload.
type Message is uint256;

using MessageLib for Message global;

/// Types of messages supported by Origin-Destination
/// - Base: message sent by protocol user, contains tips
/// - Manager: message sent between AgentManager contracts located on different chains, no tips
enum MessageFlag {
    Base,
    Manager
}

using MessageLib for MessageFlag global;

/// Library for formatting the various messages supported by Origin and Destination.
///
/// # Message memory layout
///
/// | Position   | Field  | Type    | Bytes | Description                                             |
/// | ---------- | ------ | ------- | ----- | ------------------------------------------------------- |
/// | [000..001) | flag   | uint8   | 1     | Flag specifying the type of message                     |
/// | [001..017) | header | uint128 | 16    | Encoded general routing information for the message     |
/// | [017..AAA) | body   | bytes   | ??    | Formatted payload (according to flag) with message body |
library MessageLib {
    using BaseMessageLib for MemView;
    using ByteString for MemView;
    using MemViewLib for bytes;
    using HeaderLib for MemView;

    /// @dev The variables below are not supposed to be used outside of the library directly.
    uint256 private constant OFFSET_FLAG = 0;
    uint256 private constant OFFSET_HEADER = 1;
    uint256 private constant OFFSET_BODY = OFFSET_HEADER + HEADER_LENGTH;

    // ══════════════════════════════════════════════════ MESSAGE ══════════════════════════════════════════════════════

    /**
     * @notice Returns formatted message with provided fields.
     * @param flag_     Flag specifying the type of message
     * @param header_   Encoded general routing information for the message
     * @param body_     Formatted payload (according to flag) with message body
     * @return Formatted message
     */
    function formatMessage(MessageFlag flag_, Header header_, bytes memory body_)
        internal
        pure
        returns (bytes memory)
    {
        return abi.encodePacked(flag_, header_, body_);
    }

    /**
     * @notice Returns a Message view over for the given payload.
     * @dev Will revert if the payload is not a message payload.
     */
    function castToMessage(bytes memory payload) internal pure returns (Message) {
        return castToMessage(payload.ref());
    }

    /**
     * @notice Casts a memory view to a Message view.
     * @dev Will revert if the memory view is not over a message payload.
     */
    function castToMessage(MemView memView) internal pure returns (Message) {
        require(isMessage(memView), "Not a message payload");
        return Message.wrap(MemView.unwrap(memView));
    }

    /**
     * @notice Checks that a payload is a formatted Message.
     */
    function isMessage(MemView memView) internal pure returns (bool) {
        uint256 length = memView.len();
        // Check if flag and header exist in the payload
        if (length < OFFSET_BODY) return false;
        uint8 flag_ = _flag(memView);
        // Check that Flag fits into MessageFlag enum
        if (flag_ > uint8(type(MessageFlag).max)) return false;
        // Check that body is formatted according to the flag
        // Only Base/Manager message flags exist
        if (flag_ == uint8(MessageFlag.Base)) {
            // Check if body is a formatted base message
            return _body(memView).isBaseMessage();
        } else {
            // Check if body is a formatted calldata for AgentManager call
            return _body(memView).isCallData();
        }
    }

    /// @notice Convenience shortcut for unwrapping a view.
    function unwrap(Message message) internal pure returns (MemView) {
        return MemView.wrap(Message.unwrap(message));
    }

    /// @notice Returns message's hash: a leaf to be inserted in the Merkle tree.
    function leaf(Message message) internal pure returns (bytes32) {
        MemView memView = message.unwrap();
        return memView.keccak();
    }

    // ══════════════════════════════════════════════ MESSAGE SLICING ══════════════════════════════════════════════════

    /// @notice Returns message's flag.
    function flag(Message message) internal pure returns (MessageFlag) {
        MemView memView = message.unwrap();
        // We check that flag fits into enum, when payload is wrapped
        // into Message, so this never reverts
        return MessageFlag(_flag(memView));
    }

    /// @notice Returns message's encoded header field.
    function header(Message message) internal pure returns (Header) {
        return HeaderLib.wrapPadded((message.unwrap().indexUint({index_: OFFSET_HEADER, bytes_: HEADER_LENGTH})));
    }

    /// @notice Returns message's body field as an untyped memory view.
    function body(Message message) internal pure returns (MemView) {
        MemView memView = message.unwrap();
        return _body(memView);
    }

    // ══════════════════════════════════════════════ PRIVATE HELPERS ══════════════════════════════════════════════════

    /// @dev Returns message's flag without checking that it fits into MessageFlag enum.
    function _flag(MemView memView) private pure returns (uint8) {
        return uint8(memView.indexUint({index_: OFFSET_FLAG, bytes_: 1}));
    }

    /// @dev Returns an untyped memory view over the body field without checking
    /// if the whole payload or the body are properly formatted.
    function _body(MemView memView) private pure returns (MemView) {
        return memView.sliceFrom({index_: OFFSET_BODY});
    }
}
