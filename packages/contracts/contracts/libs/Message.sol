// SPDX-License-Identifier: MIT
pragma solidity 0.8.13;

import { TypedMemView } from "./TypedMemView.sol";

import { Header } from "./Header.sol";
import { SynapseTypes } from "./SynapseTypes.sol";

/**
 * @title Message Library
 * @author Illusory Systems Inc.
 * @notice Library for versioned formatted messages used by Origin and Mirror.
 **/
library Message {
    using TypedMemView for bytes;
    using TypedMemView for bytes29;

    /**
     * @dev This is only updated if the whole message structure is changed,
     *      i.e. if a new part is added.
     *      If already existing part is changed, the message version does not get bumped.
     */
    uint16 internal constant MESSAGE_VERSION = 1;

    /// @dev Parts.Last is used only for marking the last element of the enum
    enum Parts {
        Version,
        Header,
        Tips,
        Body,
        Last
    }

    uint40 internal constant MESSAGE_TYPE = 1337;
    uint40 internal constant TIPS_TYPE = uint40(Parts.Tips);
    uint40 internal constant BODY_TYPE = uint40(Parts.Body);

    modifier onlyMessage(bytes29 _view) {
        _view.assertType(MESSAGE_TYPE);
        _;
    }

    /**
     * @dev Message memory layout
     *      All offsets are stored for backwards compatibility
     * [000 .. 002): version            uint16  2 bytes
     * [002 .. 004): header offset = 8  uint16  2 bytes
     * [004 .. 006): tips offset (AAA)  uint16  2 bytes
     * [006 .. 008): body offset (BBB)  uint16  2 bytes
     * [008 .. AAA): header             bytes   ? bytes
     * [AAA .. BBB): tips               bytes   ? bytes
     * [BBB .. CCC): body               bytes   ? bytes
     */

    /// @dev How much bytes is used for storing the version, or a single offset value
    uint8 internal constant TWO_BYTES = 2;
    /// @dev This value reflects the header offset in the latest message version
    uint16 internal constant HEADER_OFFSET = TWO_BYTES * uint8(Parts.Last);

    /**
     * @notice Returns formatted (packed) message with provided fields
     * @param _header Formatted header
     * @param _messageBody Raw bytes of message body
     * @return Formatted message
     **/
    function formatMessage(
        bytes memory _header,
        bytes memory _tips,
        bytes memory _messageBody
    ) internal pure returns (bytes memory) {
        // Version + Offsets + Header + Tips are supposed to fit within 65535 bytes
        uint16 tipsOffset = HEADER_OFFSET + uint16(_header.length);
        uint16 bodyOffset = tipsOffset + uint16(_tips.length);
        return
            abi.encodePacked(
                MESSAGE_VERSION,
                HEADER_OFFSET,
                tipsOffset,
                bodyOffset,
                _header,
                _tips,
                _messageBody
            );
    }

    /**
     * @notice Returns leaf of formatted message with provided fields.
     * @param _header Formatted header
     * @param _messageBody Raw bytes of message body
     * @return Leaf (hash) of formatted message
     **/
    function messageHash(
        bytes memory _header,
        bytes memory _tips,
        bytes memory _messageBody
    ) internal pure returns (bytes32) {
        return keccak256(formatMessage(_header, _tips, _messageBody));
    }

    function messageView(bytes memory _message) internal pure returns (bytes29) {
        return _message.ref(MESSAGE_TYPE);
    }

    /// @notice Returns message's header field as bytes29 (refer to TypedMemView library for details on bytes29 type)
    function header(bytes29 _message) internal pure onlyMessage(_message) returns (bytes29) {
        return
            _between(
                _message,
                _loadOffset(_message, Parts.Header),
                _loadOffset(_message, Parts.Tips),
                SynapseTypes.MESSAGE_HEADER
            );
    }

    /// @notice Returns message's tips field as bytes29 (refer to TypedMemView library for details on bytes29 type)
    function tips(bytes29 _message) internal pure onlyMessage(_message) returns (bytes29) {
        return
            _between(
                _message,
                _loadOffset(_message, Parts.Tips),
                _loadOffset(_message, Parts.Body),
                TIPS_TYPE
            );
    }

    /// @notice Returns message's body field as bytes29 (refer to TypedMemView library for details on bytes29 type)
    function body(bytes29 _message) internal pure onlyMessage(_message) returns (bytes29) {
        return _between(_message, _loadOffset(_message, Parts.Body), _message.len(), BODY_TYPE);
    }

    /// @notice Returns leaf of the formatted message.
    function leaf(bytes29 _message) internal pure onlyMessage(_message) returns (bytes32) {
        // TODO: do we actually need this?
        return _message.keccak();
    }

    function _between(
        bytes29 _message,
        uint256 _from,
        uint256 _to,
        uint40 _newType
    ) private pure returns (bytes29) {
        return _message.slice(_from, _to - _from, _newType);
    }

    /// @notice Loads offset for a given part of the message
    function _loadOffset(bytes29 _message, Parts _part) private pure returns (uint256) {
        return _message.indexUint(uint256(_part) * TWO_BYTES, TWO_BYTES);
    }
}
