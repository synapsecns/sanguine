// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {BaseMessageLib} from "./BaseMessage.sol";
import {ByteString} from "./ByteString.sol";
import {HEADER_LENGTH} from "../Constants.sol";
import {MemView, MemViewLib} from "./MemView.sol";
import {UnformattedMessage} from "../Errors.sol";
import {MerkleMath} from "../merkle/MerkleMath.sol";
import {Header, HeaderLib, MessageFlag} from "../stack/Header.sol";

/// Message is a memory over over a formatted message payload.
type Message is uint256;

using MessageLib for Message global;

/// Library for formatting the various messages supported by Origin and Destination.
///
/// # Message memory layout
///
/// | Position   | Field  | Type    | Bytes | Description                                             |
/// | ---------- | ------ | ------- | ----- | ------------------------------------------------------- |
/// | [000..017) | header | uint136 | 17    | Encoded general routing information for the message     |
/// | [017..AAA) | body   | bytes   | ??    | Formatted payload (according to flag) with message body |
library MessageLib {
    using BaseMessageLib for MemView;
    using ByteString for MemView;
    using MemViewLib for bytes;
    using HeaderLib for MemView;

    /// @dev The variables below are not supposed to be used outside of the library directly.
    uint256 private constant OFFSET_HEADER = 0;
    uint256 private constant OFFSET_BODY = OFFSET_HEADER + HEADER_LENGTH;

    // ══════════════════════════════════════════════════ MESSAGE ══════════════════════════════════════════════════════

    /**
     * @notice Returns formatted message with provided fields.
     * @param header_   Encoded general routing information for the message
     * @param body_     Formatted payload (according to flag) with message body
     * @return Formatted message
     */
    function formatMessage(Header header_, bytes memory body_) internal pure returns (bytes memory) {
        return abi.encodePacked(header_, body_);
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
        if (!isMessage(memView)) revert UnformattedMessage();
        return Message.wrap(MemView.unwrap(memView));
    }

    /**
     * @notice Checks that a payload is a formatted Message.
     */
    function isMessage(MemView memView) internal pure returns (bool) {
        uint256 length = memView.len();
        // Check if headers exist in the payload
        if (length < OFFSET_BODY) return false;
        // Check that Header is valid
        uint256 paddedHeader = _header(memView);
        if (!HeaderLib.isHeader(paddedHeader)) return false;
        // Check that body is formatted according to the flag
        // Only Base/Manager message flags exist
        if (HeaderLib.wrapPadded(paddedHeader).flag() == MessageFlag.Base) {
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
        // We hash header and body separately to make message proofs easier to verify
        Header header_ = message.header();
        // Only Base/Manager message flags exist
        if (header_.flag() == MessageFlag.Base) {
            return MerkleMath.getParent(header_.leaf(), message.body().castToBaseMessage().leaf());
        } else {
            return MerkleMath.getParent(header_.leaf(), message.body().castToCallData().leaf());
        }
    }

    // ══════════════════════════════════════════════ MESSAGE SLICING ══════════════════════════════════════════════════

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

    /// @dev Returns message's padded header without checking that it is a valid header.
    function _header(MemView memView) private pure returns (uint256) {
        return memView.indexUint({index_: OFFSET_HEADER, bytes_: HEADER_LENGTH});
    }

    /// @dev Returns an untyped memory view over the body field without checking
    /// if the whole payload or the body are properly formatted.
    function _body(MemView memView) private pure returns (MemView) {
        return memView.sliceFrom({index_: OFFSET_BODY});
    }
}
