// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {ByteString} from "./ByteString.sol";
import {TIPS_LENGTH} from "./Constants.sol";
import {Tips, TipsLib} from "./Tips.sol";
import {TypedMemView} from "./TypedMemView.sol";

/// @dev BaseMessage is a memory view over the base message supported by Origin-Destination
type BaseMessage is bytes29;

/// @dev Attach library functions to BaseMessage
using {
    BaseMessageLib.unwrap,
    BaseMessageLib.sender,
    BaseMessageLib.recipient,
    BaseMessageLib.tips,
    BaseMessageLib.content
} for BaseMessage global;

library BaseMessageLib {
    using ByteString for bytes;
    using TipsLib for bytes29;
    using TypedMemView for bytes29;

    /**
     * @dev Memory layout of BaseMessage fields
     * [000 .. 032): sender         bytes32 32 bytes    Sender address on origin chain
     * [032 .. 064): recipient      bytes32 32 bytes    Recipient address on destination chain
     * [064 .. 112): tips           bytes   48 bytes    Tips paid on origin chain
     * [112 .. AAA): content        bytes   ?? bytes    Content to be passed to recipient
     *
     * The variables below are not supposed to be used outside of the library directly.
     */

    uint256 private constant OFFSET_SENDER = 0;
    uint256 private constant OFFSET_RECIPIENT = 32;
    uint256 private constant OFFSET_TIPS = 64;
    uint256 private constant OFFSET_CONTENT = OFFSET_TIPS + TIPS_LENGTH;

    // ═══════════════════════════════════════════════ BASE MESSAGE ════════════════════════════════════════════════════

    /**
     * @notice Returns a formatted BaseMessage payload with provided fields.
     * @param sender_       Sender address on origin chain
     * @param recipient_    Recipient address on destination chain
     * @param tipsPayload   Formatted payload with tips information
     * @param content_      Raw content to be passed to recipient on destination chain
     * @return Formatted base message
     */
    function formatBaseMessage(bytes32 sender_, bytes32 recipient_, bytes memory tipsPayload, bytes memory content_)
        internal
        pure
        returns (bytes memory)
    {
        return abi.encodePacked(sender_, recipient_, tipsPayload, content_);
    }

    /**
     * @notice Returns a BaseMessage view over the given payload.
     * @dev Will revert if the payload is not a base message.
     */
    function castToBaseMessage(bytes memory payload) internal pure returns (BaseMessage) {
        return castToBaseMessage(payload.castToRawBytes());
    }

    /**
     * @notice Casts a memory view to a BaseMessage view.
     * @dev Will revert if the memory view is not over a base message payload.
     */
    function castToBaseMessage(bytes29 view_) internal pure returns (BaseMessage) {
        require(isBaseMessage(view_), "Not a base message");
        return BaseMessage.wrap(view_);
    }

    /// @notice Checks that a payload is a formatted BaseMessage.
    function isBaseMessage(bytes29 view_) internal pure returns (bool) {
        // Check if sender, recipient, tips fields exist
        if (view_.len() < OFFSET_CONTENT) return false;
        // Check if tips payload is formatted
        return _tips(view_).isTips();
        // Content could be empty, so we don't check that
    }

    /// @notice Convenience shortcut for unwrapping a view.
    function unwrap(BaseMessage baseMessage) internal pure returns (bytes29) {
        return BaseMessage.unwrap(baseMessage);
    }

    // ═══════════════════════════════════════════ BASE MESSAGE SLICING ════════════════════════════════════════════════

    /// @notice Returns sender address on origin chain.
    function sender(BaseMessage baseMessage) internal pure returns (bytes32) {
        // Get the underlying memory view
        bytes29 view_ = baseMessage.unwrap();
        return view_.index({index_: OFFSET_SENDER, bytes_: 32});
    }

    /// @notice Returns recipient address on destination chain.
    function recipient(BaseMessage baseMessage) internal pure returns (bytes32) {
        bytes29 view_ = baseMessage.unwrap();
        return view_.index({index_: OFFSET_RECIPIENT, bytes_: 32});
    }

    /// @notice Returns a typed memory view over the payload with tips paid on origin chain.
    function tips(BaseMessage baseMessage) internal pure returns (Tips) {
        bytes29 view_ = baseMessage.unwrap();
        // We check that tips payload is properly formatted, when the whole payload is wrapped
        // into BaseMessage, so this never reverts.
        return _tips(view_).castToTips();
    }

    /// @notice Returns an untyped memory view over the content to be passed to recipient.
    function content(BaseMessage baseMessage) internal pure returns (bytes29) {
        bytes29 view_ = baseMessage.unwrap();
        return view_.sliceFrom({index_: OFFSET_CONTENT, newType: 0});
    }

    // ══════════════════════════════════════════════ PRIVATE HELPERS ══════════════════════════════════════════════════

    /// @dev Returns an untyped memory view over the tips field without checking
    /// if the whole payload or the tips are properly formatted.
    function _tips(bytes29 view_) private pure returns (bytes29) {
        return view_.slice({index_: OFFSET_TIPS, len_: TIPS_LENGTH, newType: 0});
    }
}
