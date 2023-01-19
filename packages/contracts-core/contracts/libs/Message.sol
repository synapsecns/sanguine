// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import "./Header.sol";
import "./Tips.sol";

/// @dev Message is a memory over over a formatted message payload.
type Message is bytes29;

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
     * @param _header       Formatted header payload
     * @param _tips         Formatted tips payload
     * @param _messageBody  Raw bytes of message body
     * @return Formatted message
     **/
    function formatMessage(
        bytes memory _header,
        bytes memory _tips,
        bytes memory _messageBody
    ) internal pure returns (bytes memory) {
        // Header and Tips are supposed to fit within 65535 bytes
        return
            abi.encodePacked(
                MESSAGE_VERSION,
                uint16(_header.length),
                uint16(_tips.length),
                _header,
                _tips,
                _messageBody
            );
    }

    /**
     * @notice Returns formatted message with provided fields
     * @param _origin               Domain of origin chain
     * @param _sender               Address that sent the message
     * @param _nonce                Message nonce on origin chain
     * @param _destination          Domain of destination chain
     * @param _recipient            Address that will receive the message
     * @param _optimisticSeconds    Optimistic period for message execution
     * @param _tips                 Formatted tips payload
     * @param _messageBody          Raw bytes of message body
     * @return Formatted message
     **/
    function formatMessage(
        uint32 _origin,
        bytes32 _sender,
        uint32 _nonce,
        uint32 _destination,
        bytes32 _recipient,
        uint32 _optimisticSeconds,
        bytes memory _tips,
        bytes memory _messageBody
    ) internal pure returns (bytes memory) {
        return
            formatMessage(
                HeaderLib.formatHeader(
                    _origin,
                    _sender,
                    _nonce,
                    _destination,
                    _recipient,
                    _optimisticSeconds
                ),
                _tips,
                _messageBody
            );
    }

    /**
     * @notice Returns a Message view over for the given payload.
     * @dev Will revert if the payload is not a message payload.
     */
    function castToMessage(bytes memory _payload) internal pure returns (Message) {
        return castToMessage(_payload.castToRawBytes());
    }

    /**
     * @notice Casts a memory view to a Message view.
     * @dev Will revert if the memory view is not over a message payload.
     */
    function castToMessage(bytes29 _view) internal pure returns (Message) {
        require(isMessage(_view), "Not a message payload");
        return Message.wrap(_view);
    }

    /**
     * @notice Checks that a payload is a formatted Message.
     */
    function isMessage(bytes29 _view) internal pure returns (bool) {
        uint256 length = _view.len();
        // Check if version and lengths exist in the payload
        if (length < OFFSET_HEADER) return false;
        // Check message version
        if (_getVersion(_view) != MESSAGE_VERSION) return false;

        uint256 headerLength = _getLen(_view, Parts.Header);
        uint256 tipsLength = _getLen(_view, Parts.Tips);
        // Header and Tips need to exist
        // Body could be empty, thus >
        if (OFFSET_HEADER + headerLength + tipsLength > length) return false;

        // Check header for being a formatted header payload
        // Check tips for being a formatted tips payload
        if (!_getHeader(_view).isHeader() || !_getTips(_view).isTips()) return false;
        return true;
    }

    /// @notice Convenience shortcut for unwrapping a view.
    function unwrap(Message _msg) internal pure returns (bytes29) {
        return Message.unwrap(_msg);
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                           MESSAGE SLICING                            ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /// @notice Returns message's version field.
    function version(Message _msg) internal pure returns (uint16) {
        // Get the underlying memory view
        bytes29 _view = unwrap(_msg);
        return _getVersion(_view);
    }

    /// @notice Returns message's header field as a Header view.
    function header(Message _msg) internal pure returns (Header) {
        bytes29 _view = unwrap(_msg);
        return _getHeader(_view).castToHeader();
    }

    /// @notice Returns message's tips field as a Tips view.
    function tips(Message _msg) internal pure returns (Tips) {
        bytes29 _view = unwrap(_msg);
        return _getTips(_view).castToTips();
    }

    /// @notice Returns message's body field as a generic memory view.
    function body(Message _msg) internal pure returns (bytes29) {
        bytes29 _view = unwrap(_msg);
        // Determine index where message body payload starts
        uint256 index = OFFSET_HEADER + _getLen(_view, Parts.Header) + _getLen(_view, Parts.Tips);
        return _view.sliceFrom({ _index: index, newType: 0 });
    }

    /// @notice Returns message's hash: a leaf to be inserted in the Merkle tree.
    function leaf(Message _msg) internal pure returns (bytes32) {
        bytes29 _view = unwrap(_msg);
        return _view.keccak();
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                           PRIVATE HELPERS                            ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /// @dev Returns length for a given part of the message
    /// without checking if the payload is properly formatted.
    function _getLen(bytes29 _view, Parts _part) private pure returns (uint256) {
        return _view.indexUint(uint256(_part) * TWO_BYTES, TWO_BYTES);
    }

    /// @dev Returns a version field without checking if the payload is properly formatted.
    function _getVersion(bytes29 _view) private pure returns (uint16) {
        return uint16(_view.indexUint(OFFSET_VERSION, 2));
    }

    /// @dev Returns a generic memory view over the header field without checking
    /// if the whole payload or the header are properly formatted.
    function _getHeader(bytes29 _view) private pure returns (bytes29) {
        uint256 length = _getLen(_view, Parts.Header);
        return _view.slice({ _index: OFFSET_HEADER, _len: length, newType: 0 });
    }

    /// @dev Returns a generic memory view over the tips field without checking
    /// if the whole payload or the tips are properly formatted.
    function _getTips(bytes29 _view) private pure returns (bytes29) {
        // Determine index where tips payload starts
        uint256 indexFrom = OFFSET_HEADER + _getLen(_view, Parts.Header);
        uint256 length = _getLen(_view, Parts.Tips);
        return _view.slice({ _index: indexFrom, _len: length, newType: 0 });
    }
}
