// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import { TypedMemView } from "./TypedMemView.sol";

import { Header } from "./Header.sol";
import { Tips } from "./Tips.sol";
import { SynapseTypes } from "./SynapseTypes.sol";

/**
 * @notice  Library for versioned formatting the messages used by Origin and Destination.
 */
library Message {
    using Header for bytes29;
    using Tips for bytes29;
    using TypedMemView for bytes;
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
    ▏*║                              MODIFIERS                               ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    modifier onlyMessage(bytes29 _view) {
        _view.assertType(SynapseTypes.MESSAGE);
        _;
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                              FORMATTERS                              ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @notice Returns a properly typed bytes29 pointer for a message payload.
     */
    function castToMessage(bytes memory _payload) internal pure returns (bytes29) {
        return _payload.ref(SynapseTypes.MESSAGE);
    }

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
                Header.formatHeader(
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
     * @notice Checks that a payload is a formatted Message.
     */
    function isMessage(bytes29 _view) internal pure returns (bool) {
        uint256 length = _view.len();
        // Check if version and lengths exist in the payload
        if (length < OFFSET_HEADER) return false;
        // Check message version
        if (messageVersion(_view) != MESSAGE_VERSION) return false;

        uint256 headerLength = _loadLength(_view, Parts.Header);
        uint256 tipsLength = _loadLength(_view, Parts.Tips);
        // Header and Tips need to exist
        // Body could be empty, thus >
        if (OFFSET_HEADER + headerLength + tipsLength > length) return false;

        // Check header for being a formatted header payload
        // Check tips for being a formatted tips payload
        if (!header(_view).isHeader() || !tips(_view).isTips()) return false;
        return true;
    }

    /**
     * @notice Returns leaf of formatted message with provided fields.
     * @param _header       Formatted header payload
     * @param _tips         Formatted tips payload
     * @param _messageBody  Raw bytes of message body
     * @return Leaf (hash) of formatted message
     **/
    function messageHash(
        bytes memory _header,
        bytes memory _tips,
        bytes memory _messageBody
    ) internal pure returns (bytes32) {
        return keccak256(formatMessage(_header, _tips, _messageBody));
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                           MESSAGE SLICING                            ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /// @notice Returns message's version field.
    function messageVersion(bytes29 _view) internal pure onlyMessage(_view) returns (uint16) {
        return uint16(_view.indexUint(OFFSET_VERSION, 2));
    }

    /// @notice Returns message's header field as bytes29 pointer.
    function header(bytes29 _view) internal pure onlyMessage(_view) returns (bytes29) {
        return
            _view.slice(
                OFFSET_HEADER,
                _loadLength(_view, Parts.Header),
                SynapseTypes.MESSAGE_HEADER
            );
    }

    /// @notice Returns message's tips field as bytes29 pointer.
    function tips(bytes29 _view) internal pure onlyMessage(_view) returns (bytes29) {
        return
            _view.slice(
                OFFSET_HEADER + _loadLength(_view, Parts.Header),
                _loadLength(_view, Parts.Tips),
                SynapseTypes.MESSAGE_TIPS
            );
    }

    /// @notice Returns message's body field as bytes29 pointer.
    function body(bytes29 _view) internal pure onlyMessage(_view) returns (bytes29) {
        return
            _view.sliceFrom(
                OFFSET_HEADER + _loadLength(_view, Parts.Header) + _loadLength(_view, Parts.Tips),
                SynapseTypes.MESSAGE_BODY
            );
    }

    /// @notice Loads length for a given part of the message
    function _loadLength(bytes29 _view, Parts _part) private pure returns (uint256) {
        return _view.indexUint(uint256(_part) * TWO_BYTES, TWO_BYTES);
    }
}
