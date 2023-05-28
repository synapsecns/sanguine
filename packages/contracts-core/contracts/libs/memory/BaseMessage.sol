// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {MemView, MemViewLib} from "./MemView.sol";
import {REQUEST_LENGTH, TIPS_LENGTH} from "../Constants.sol";
import {UnformattedBaseMessage} from "../Errors.sol";
import {MerkleMath} from "../merkle/MerkleMath.sol";
import {Request, RequestLib} from "../stack/Request.sol";
import {Tips, TipsLib} from "../stack/Tips.sol";

/// BaseMessage is a memory view over the base message supported by Origin-Destination
type BaseMessage is uint256;

using BaseMessageLib for BaseMessage global;

/// BaseMessage structure represents a base message sent via the Origin-Destination contracts.
/// - It only contains data relevant to the base message, the rest of data is encoded in the message header.
/// - `sender` and `recipient` for EVM chains are EVM addresses casted to bytes32, while preserving left-alignment.
/// - `tips` and `request` parameters are specified by a message sender
/// > Origin will calculate minimum tips for given request and content length, and will reject messages with tips
/// lower than that.
///
/// # Memory layout of BaseMessage fields
///
/// | Position   | Field     | Type    | Bytes | Description                            |
/// | ---------- | --------- | ------- | ----- | -------------------------------------- |
/// | [000..032) | tips      | uint256 | 32    | Encoded tips paid on origin chain      |
/// | [032..064) | sender    | bytes32 | 32    | Sender address on origin chain         |
/// | [064..096) | recipient | bytes32 | 32    | Recipient address on destination chain |
/// | [096..116) | request   | uint160 | 20    | Encoded request for message execution  |
/// | [104..AAA) | content   | bytes   | ??    | Content to be passed to recipient      |
library BaseMessageLib {
    using MemViewLib for bytes;

    /// @dev The variables below are not supposed to be used outside of the library directly.
    uint256 private constant OFFSET_TIPS = 0;
    uint256 private constant OFFSET_SENDER = 32;
    uint256 private constant OFFSET_RECIPIENT = 64;
    uint256 private constant OFFSET_REQUEST = OFFSET_RECIPIENT + TIPS_LENGTH;
    uint256 private constant OFFSET_CONTENT = OFFSET_REQUEST + REQUEST_LENGTH;

    // ═══════════════════════════════════════════════ BASE MESSAGE ════════════════════════════════════════════════════

    /**
     * @notice Returns a formatted BaseMessage payload with provided fields.
     * @param tips_         Encoded tips information
     * @param sender_       Sender address on origin chain
     * @param recipient_    Recipient address on destination chain
     * @param request_      Encoded request for message execution
     * @param content_      Raw content to be passed to recipient on destination chain
     * @return Formatted base message
     */
    function formatBaseMessage(Tips tips_, bytes32 sender_, bytes32 recipient_, Request request_, bytes memory content_)
        internal
        pure
        returns (bytes memory)
    {
        return abi.encodePacked(tips_, sender_, recipient_, request_, content_);
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
        if (!isBaseMessage(memView)) revert UnformattedBaseMessage();
        return BaseMessage.wrap(MemView.unwrap(memView));
    }

    /// @notice Checks that a payload is a formatted BaseMessage.
    function isBaseMessage(MemView memView) internal pure returns (bool) {
        // Check if sender, recipient, tips fields exist
        return (memView.len() >= OFFSET_CONTENT);
        // Content could be empty, so we don't check that
    }

    /// @notice Convenience shortcut for unwrapping a view.
    function unwrap(BaseMessage baseMessage) internal pure returns (MemView) {
        return MemView.wrap(BaseMessage.unwrap(baseMessage));
    }

    /// @notice Returns baseMessage's hash: a leaf to be inserted in the "Message mini-Merkle tree".
    function leaf(BaseMessage baseMessage) internal pure returns (bytes32) {
        // We hash "tips" and "everything but tips" to make tips proofs easier to verify
        return MerkleMath.getParent(baseMessage.tips().leaf(), baseMessage.bodyLeaf());
    }

    /// @notice Returns hash for the "everything but tips" part of the base message.
    function bodyLeaf(BaseMessage baseMessage) internal pure returns (bytes32) {
        return baseMessage.unwrap().sliceFrom({index_: OFFSET_SENDER}).keccak();
    }

    // ═══════════════════════════════════════════ BASE MESSAGE SLICING ════════════════════════════════════════════════

    /// @notice Returns encoded tips paid on origin chain.
    function tips(BaseMessage baseMessage) internal pure returns (Tips) {
        return TipsLib.wrapPadded((baseMessage.unwrap().indexUint({index_: OFFSET_TIPS, bytes_: TIPS_LENGTH})));
    }

    /// @notice Returns sender address on origin chain.
    function sender(BaseMessage baseMessage) internal pure returns (bytes32) {
        return baseMessage.unwrap().index({index_: OFFSET_SENDER, bytes_: 32});
    }

    /// @notice Returns recipient address on destination chain.
    function recipient(BaseMessage baseMessage) internal pure returns (bytes32) {
        return baseMessage.unwrap().index({index_: OFFSET_RECIPIENT, bytes_: 32});
    }

    /// @notice Returns an encoded request for message execution on destination chain.
    function request(BaseMessage baseMessage) internal pure returns (Request) {
        return RequestLib.wrapPadded((baseMessage.unwrap().indexUint({index_: OFFSET_REQUEST, bytes_: REQUEST_LENGTH})));
    }

    /// @notice Returns an untyped memory view over the content to be passed to recipient.
    function content(BaseMessage baseMessage) internal pure returns (MemView) {
        return baseMessage.unwrap().sliceFrom({index_: OFFSET_CONTENT});
    }
}
