// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {BaseMessageLib} from "./BaseMessage.sol";
import {ByteString} from "./ByteString.sol";
import {Header, HEADER_LENGTH, HeaderLib} from "./Header.sol";
import {SystemMessageLib} from "./SystemMessage.sol";
import {TypedMemView} from "./TypedMemView.sol";

/// @dev Message is a memory over over a formatted message payload.
type Message is bytes29;

/// @dev Attach library functions to Message
using MessageLib for Message global;

/// @dev Types of messages supported by Origin-Destination
/// - System: message sent between system contracts located on different chains
/// - Base: message sent by protocol user, contains tips
enum MessageFlag {
    System,
    Base
}

/// @dev Attach library functions to MessageFlag
using MessageLib for MessageFlag global;

/**
 * @notice Library for formatting the various messages supported by Origin and Destination.
 */
library MessageLib {
    using BaseMessageLib for bytes29;
    using ByteString for bytes;
    using HeaderLib for bytes29;
    using SystemMessageLib for bytes29;
    using TypedMemView for bytes29;

    /**
     * @dev Message memory layout
     * [000 .. 001): flag       uint8    1 byte     Flag specifying the type of message
     * [001 .. 017): header     bytes   16 bytes    Formatted payload with general routing information
     * [017 .. AAA): body       bytes   ?? bytes    Formatted payload (according to flag) with message body
     *
     * The variables below are not supposed to be used outside of the library directly.
     */

    uint256 private constant OFFSET_FLAG = 0;
    uint256 private constant OFFSET_HEADER = 1;
    uint256 private constant OFFSET_BODY = OFFSET_HEADER + HEADER_LENGTH;

    // ══════════════════════════════════════════════════ MESSAGE ══════════════════════════════════════════════════════

    /**
     * @notice Returns formatted message with provided fields.
     * @param flag_     Flag specifying the type of message
     * @param header_   Formatted payload with general routing information
     * @param body_     Formatted payload (according to flag) with message body
     * @return Formatted message
     */
    function formatMessage(MessageFlag flag_, bytes memory header_, bytes memory body_)
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
        return castToMessage(payload.castToRawBytes());
    }

    /**
     * @notice Casts a memory view to a Message view.
     * @dev Will revert if the memory view is not over a message payload.
     */
    function castToMessage(bytes29 view_) internal pure returns (Message) {
        require(isMessage(view_), "Not a message payload");
        return Message.wrap(view_);
    }

    /**
     * @notice Checks that a payload is a formatted Message.
     */
    function isMessage(bytes29 view_) internal pure returns (bool) {
        uint256 length = view_.len();
        // Check if flag and header exist in the payload
        if (length < OFFSET_BODY) return false;
        uint8 flag_ = _flag(view_);
        // Check that Flag fits into MessageFlag enum
        if (flag_ > uint8(type(MessageFlag).max)) return false;
        // Check that Header is formatted
        if (!_header(view_).isHeader()) return false;
        // Check that body is formatted according to the flag
        // Only System/Base message flags exist
        if (flag_ == uint8(MessageFlag.System)) {
            // Check if body is a formatted system message
            return _body(view_).isSystemMessage();
        } else {
            // Check if body is a formatted base message
            return _body(view_).isBaseMessage();
        }
    }

    /// @notice Convenience shortcut for unwrapping a view.
    function unwrap(Message message) internal pure returns (bytes29) {
        return Message.unwrap(message);
    }

    /// @notice Returns message's hash: a leaf to be inserted in the Merkle tree.
    function leaf(Message message) internal pure returns (bytes32) {
        bytes29 view_ = message.unwrap();
        return view_.keccak();
    }

    // ══════════════════════════════════════════════ MESSAGE SLICING ══════════════════════════════════════════════════

    /// @notice Returns message's flag.
    function flag(Message message) internal pure returns (MessageFlag) {
        bytes29 view_ = message.unwrap();
        // We check that flag fits into enum, when payload is wrapped
        // into Message, so this never reverts
        return MessageFlag(_flag(view_));
    }

    /// @notice Returns message's header field as a Header view.
    function header(Message message) internal pure returns (Header) {
        bytes29 view_ = message.unwrap();
        // We check that header is properly formatted, when payload is wrapped
        // into Message, so this never reverts.
        return _header(view_).castToHeader();
    }

    /// @notice Returns message's body field as an untyped memory view.
    function body(Message message) internal pure returns (bytes29) {
        bytes29 view_ = message.unwrap();
        return _body(view_);
    }

    // ══════════════════════════════════════════════ PRIVATE HELPERS ══════════════════════════════════════════════════

    /// @dev Returns message's flag without checking that it fits into MessageFlag enum.
    function _flag(bytes29 view_) private pure returns (uint8) {
        return uint8(view_.indexUint({index_: OFFSET_FLAG, bytes_: 1}));
    }

    /// @dev Returns an untyped memory view over the header field without checking
    /// if the whole payload or the header are properly formatted.
    function _header(bytes29 view_) private pure returns (bytes29) {
        return view_.slice({index_: OFFSET_HEADER, len_: HEADER_LENGTH, newType: 0});
    }

    /// @dev Returns an untyped memory view over the body field without checking
    /// if the whole payload or the body are properly formatted.
    function _body(bytes29 view_) private pure returns (bytes29) {
        return view_.sliceFrom({index_: OFFSET_BODY, newType: 0});
    }
}
