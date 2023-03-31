// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {ByteString} from "./ByteString.sol";
import {Header, HeaderLib} from "./Header.sol";
import {Tips, TipsLib} from "./Tips.sol";
import {TypedMemView} from "./TypedMemView.sol";

/// @dev Message is a memory over over a formatted message payload.
type Message is bytes29;
/// @dev Attach library functions to Message

using {
    MessageLib.unwrap,
    MessageLib.version,
    MessageLib.header,
    MessageLib.tips,
    MessageLib.body,
    MessageLib.leaf
} for Message global;

/**
 * @notice  Library for versioned formatting the messages used by Origin and Destination.
 */
library MessageLib {
    using HeaderLib for bytes29;
    using TipsLib for bytes29;
    using ByteString for bytes;
    using TypedMemView for bytes29;

    enum Parts {
        Version,
        Header,
        Tips,
        Body
    }

    /**
     * @dev This is only updated if the whole message structure is changed,
     *      i.e. if a new part is added.
     *      If already existing part is changed, the message version does not get bumped.
     */
    uint16 internal constant MESSAGE_VERSION = 1;

    /**
     * @dev Message memory layout
     * [000 .. 002): version            uint16  2 bytes
     * [002 .. 004): header length      uint16  2 bytes (length == AAA - 6)
     * [004 .. 006): tips length        uint16  2 bytes (length == BBB - AAA)
     * [006 .. AAA): header             bytes   ? bytes
     * [AAA .. BBB): tips               bytes   ? bytes
     * [BBB .. CCC): body               bytes   ? bytes (length could be zero)
     */

    uint256 internal constant OFFSET_VERSION = 0;

    /// @dev How much bytes is used for storing the version, or a single offset value
    uint8 internal constant TWO_BYTES = 2;
    /// @dev This value reflects the header offset in the latest message version
    uint16 internal constant OFFSET_HEADER = TWO_BYTES * uint8(type(Parts).max);

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                              FORMATTERS                              ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @notice Returns formatted message with provided fields
     * @param header_       Formatted header payload
     * @param tips_         Formatted tips payload
     * @param messageBody   Raw bytes of message body
     * @return Formatted message
     */
    function formatMessage(bytes memory header_, bytes memory tips_, bytes memory messageBody)
        internal
        pure
        returns (bytes memory)
    {
        // Header and Tips are supposed to fit within 65535 bytes
        return
            abi.encodePacked(MESSAGE_VERSION, uint16(header_.length), uint16(tips_.length), header_, tips_, messageBody);
    }

    /**
     * @notice Returns formatted message with provided fields
     * @param origin                Domain of origin chain
     * @param sender                Address that sent the message
     * @param nonce                 Message nonce on origin chain
     * @param destination           Domain of destination chain
     * @param recipient             Address that will receive the message
     * @param optimisticSeconds     Optimistic period for message execution
     * @param tips_                 Formatted tips payload
     * @param messageBody           Raw bytes of message body
     * @return Formatted message
     */
    function formatMessage(
        uint32 origin,
        bytes32 sender,
        uint32 nonce,
        uint32 destination,
        bytes32 recipient,
        uint32 optimisticSeconds,
        bytes memory tips_,
        bytes memory messageBody
    ) internal pure returns (bytes memory) {
        return formatMessage(
            HeaderLib.formatHeader(origin, sender, nonce, destination, recipient, optimisticSeconds), tips_, messageBody
        );
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
        // Check if version and lengths exist in the payload
        if (length < OFFSET_HEADER) return false;
        // Check message version
        if (_getVersion(view_) != MESSAGE_VERSION) return false;

        uint256 headerLength = _getLen(view_, Parts.Header);
        uint256 tipsLength = _getLen(view_, Parts.Tips);
        // Header and Tips need to exist
        // Body could be empty, thus >
        if (OFFSET_HEADER + headerLength + tipsLength > length) return false;

        // Check header for being a formatted header payload
        // Check tips for being a formatted tips payload
        if (!_getHeader(view_).isHeader() || !_getTips(view_).isTips()) return false;
        return true;
    }

    /// @notice Convenience shortcut for unwrapping a view.
    function unwrap(Message message) internal pure returns (bytes29) {
        return Message.unwrap(message);
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                           MESSAGE SLICING                            ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /// @notice Returns message's version field.
    function version(Message message) internal pure returns (uint16) {
        // Get the underlying memory view
        bytes29 view_ = message.unwrap();
        return _getVersion(view_);
    }

    /// @notice Returns message's header field as a Header view.
    function header(Message message) internal pure returns (Header) {
        bytes29 view_ = message.unwrap();
        return _getHeader(view_).castToHeader();
    }

    /// @notice Returns message's tips field as a Tips view.
    function tips(Message message) internal pure returns (Tips) {
        bytes29 view_ = message.unwrap();
        return _getTips(view_).castToTips();
    }

    /// @notice Returns message's body field as a generic memory view.
    function body(Message message) internal pure returns (bytes29) {
        bytes29 view_ = message.unwrap();
        // Determine index where message body payload starts
        uint256 index = OFFSET_HEADER + _getLen(view_, Parts.Header) + _getLen(view_, Parts.Tips);
        return view_.sliceFrom({index_: index, newType: 0});
    }

    /// @notice Returns message's hash: a leaf to be inserted in the Merkle tree.
    function leaf(Message message) internal pure returns (bytes32) {
        bytes29 view_ = message.unwrap();
        return view_.keccak();
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                           PRIVATE HELPERS                            ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /// @dev Returns length for a given part of the message
    /// without checking if the payload is properly formatted.
    function _getLen(bytes29 view_, Parts part) private pure returns (uint256) {
        return view_.indexUint(uint256(part) * TWO_BYTES, TWO_BYTES);
    }

    /// @dev Returns a version field without checking if the payload is properly formatted.
    function _getVersion(bytes29 view_) private pure returns (uint16) {
        return uint16(view_.indexUint(OFFSET_VERSION, 2));
    }

    /// @dev Returns a generic memory view over the header field without checking
    /// if the whole payload or the header are properly formatted.
    function _getHeader(bytes29 view_) private pure returns (bytes29) {
        uint256 length = _getLen(view_, Parts.Header);
        return view_.slice({index_: OFFSET_HEADER, len_: length, newType: 0});
    }

    /// @dev Returns a generic memory view over the tips field without checking
    /// if the whole payload or the tips are properly formatted.
    function _getTips(bytes29 view_) private pure returns (bytes29) {
        // Determine index where tips payload starts
        uint256 indexFrom = OFFSET_HEADER + _getLen(view_, Parts.Header);
        uint256 length = _getLen(view_, Parts.Tips);
        return view_.slice({index_: indexFrom, len_: length, newType: 0});
    }
}
