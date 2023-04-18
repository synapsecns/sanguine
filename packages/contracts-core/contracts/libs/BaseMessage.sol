// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {REQUEST_LENGTH, TIPS_LENGTH} from "./Constants.sol";
import {Request, RequestLib} from "./Request.sol";
import {Tips, TipsLib} from "./Tips.sol";
import {MemView, MemViewLib} from "./MemView.sol";

/// @dev BaseMessage is a memory view over the base message supported by Origin-Destination
type BaseMessage is uint256;

/// @dev Attach library functions to BaseMessage
using BaseMessageLib for BaseMessage global;

library BaseMessageLib {
    using MemViewLib for bytes;
    using RequestLib for MemView;
    using TipsLib for MemView;

    /**
     * @dev Memory layout of BaseMessage fields
     * [000 .. 032): sender         bytes32 32 bytes    Sender address on origin chain
     * [032 .. 064): recipient      bytes32 32 bytes    Recipient address on destination chain
     * [064 .. 112): tips           bytes   48 bytes    Tips paid on origin chain
     * [112 .. 120): request        bytes    8 bytes    Request for message execution on destination chain
     * [120 .. AAA): content        bytes   ?? bytes    Content to be passed to recipient
     *
     * The variables below are not supposed to be used outside of the library directly.
     */

    uint256 private constant OFFSET_SENDER = 0;
    uint256 private constant OFFSET_RECIPIENT = 32;
    uint256 private constant OFFSET_TIPS = 64;
    uint256 private constant OFFSET_REQUEST = OFFSET_TIPS + TIPS_LENGTH;
    uint256 private constant OFFSET_CONTENT = OFFSET_REQUEST + REQUEST_LENGTH;

    // ═══════════════════════════════════════════════ BASE MESSAGE ════════════════════════════════════════════════════

    /**
     * @notice Returns a formatted BaseMessage payload with provided fields.
     * @param sender_       Sender address on origin chain
     * @param recipient_    Recipient address on destination chain
     * @param tipsPayload   Formatted payload with tips information
     * @param content_      Raw content to be passed to recipient on destination chain
     * @return Formatted base message
     */
    function formatBaseMessage(
        bytes32 sender_,
        bytes32 recipient_,
        bytes memory tipsPayload,
        bytes memory requestPayload,
        bytes memory content_
    ) internal pure returns (bytes memory) {
        return abi.encodePacked(sender_, recipient_, tipsPayload, requestPayload, content_);
    }

    /**
     * @notice Returns a BaseMessage view over the given payload.
     * @dev Will revert if the payload is not a base message.
     */
    function castToBaseMessage(bytes memory payload) internal pure returns (BaseMessage) {
        return castToBaseMessage(payload.ref());
    }

    /**
     * @notice Casts a memory view to a BaseMessage view.
     * @dev Will revert if the memory view is not over a base message payload.
     */
    function castToBaseMessage(MemView memView) internal pure returns (BaseMessage) {
        require(isBaseMessage(memView), "Not a base message");
        return BaseMessage.wrap(MemView.unwrap(memView));
    }

    /// @notice Checks that a payload is a formatted BaseMessage.
    function isBaseMessage(MemView memView) internal pure returns (bool) {
        // Check if sender, recipient, tips fields exist
        if (memView.len() < OFFSET_CONTENT) return false;
        // Check if tips payload is formatted
        if (!_tips(memView).isTips()) return false;
        // Check if tips payload is formatted
        return _request(memView).isRequest();
        // Content could be empty, so we don't check that
    }

    /// @notice Convenience shortcut for unwrapping a view.
    function unwrap(BaseMessage baseMessage) internal pure returns (MemView) {
        return MemView.wrap(BaseMessage.unwrap(baseMessage));
    }

    // ═══════════════════════════════════════════ BASE MESSAGE SLICING ════════════════════════════════════════════════

    /// @notice Returns sender address on origin chain.
    function sender(BaseMessage baseMessage) internal pure returns (bytes32) {
        // Get the underlying memory view
        MemView memView = baseMessage.unwrap();
        return memView.index({index_: OFFSET_SENDER, bytes_: 32});
    }

    /// @notice Returns recipient address on destination chain.
    function recipient(BaseMessage baseMessage) internal pure returns (bytes32) {
        MemView memView = baseMessage.unwrap();
        return memView.index({index_: OFFSET_RECIPIENT, bytes_: 32});
    }

    /// @notice Returns a typed memory view over the payload with tips paid on origin chain.
    function tips(BaseMessage baseMessage) internal pure returns (Tips) {
        MemView memView = baseMessage.unwrap();
        // We check that tips payload is properly formatted, when the whole payload is wrapped
        // into BaseMessage, so this never reverts.
        return _tips(memView).castToTips();
    }

    /// @notice Returns a typed memory view over the payload with request for message execution on destination chain.
    function request(BaseMessage baseMessage) internal pure returns (Request) {
        MemView memView = baseMessage.unwrap();
        // We check that request payload is properly formatted, when the whole payload is wrapped
        // into BaseMessage, so this never reverts.
        return _request(memView).castToRequest();
    }

    /// @notice Returns an untyped memory view over the content to be passed to recipient.
    function content(BaseMessage baseMessage) internal pure returns (MemView) {
        MemView memView = baseMessage.unwrap();
        return memView.sliceFrom({index_: OFFSET_CONTENT});
    }

    // ══════════════════════════════════════════════ PRIVATE HELPERS ══════════════════════════════════════════════════

    /// @dev Returns an untyped memory view over the tips field without checking
    /// if the whole payload or the tips are properly formatted.
    function _tips(MemView memView) private pure returns (MemView) {
        return memView.slice({index_: OFFSET_TIPS, len_: TIPS_LENGTH});
    }

    /// @dev Returns an untyped memory view over the request field without checking
    /// if the whole payload or the request are properly formatted.
    function _request(MemView memView) private pure returns (MemView) {
        return memView.slice({index_: OFFSET_REQUEST, len_: REQUEST_LENGTH});
    }
}
