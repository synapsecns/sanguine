// SPDX-License-Identifier: MIT
pragma solidity 0.8.13;

import { TypedMemView } from "./TypedMemView.sol";

import { Header } from "./Header.sol";

/**
 * @title Message Library
 * @author Illusory Systems Inc.
 * @notice Library for formatted messages used by Home and Replica.
 **/
library Message {
    using TypedMemView for bytes;
    using TypedMemView for bytes29;

    uint40 internal constant MESSAGE_TYPE = 1;
    uint40 internal constant HEADER_TYPE = 2;
    uint40 internal constant BODY_TYPE = 3;

    modifier onlyMessage(bytes29 _view) {
        _view.assertType(MESSAGE_TYPE);
        _;
    }

    /**
     * @dev Message memory layout
     * [000 .. 002): header.length  uint16  2 bytes
     * [002 .. AAA): header         bytes   ? bytes
     * [AAA .. BBB): body           bytes   ? bytes
     */

    uint256 internal constant OFFSET_HEADER = 2;

    /**
     * @notice Returns formatted (packed) message with provided fields
     * @param _header Formatted header
     * @param _messageBody Raw bytes of message body
     * @return Formatted message
     **/
    function formatMessage(bytes memory _header, bytes memory _messageBody)
        internal
        pure
        returns (bytes memory)
    {
        // Header is always supposed to fit in 65535 bytes
        uint16 length = uint16(_header.length);
        return abi.encodePacked(length, _header, _messageBody);
    }

    /**
     * @notice Returns leaf of formatted message with provided fields.
     * @param _header Formatted header
     * @param _messageBody Raw bytes of message body
     * @return Leaf (hash) of formatted message
     **/
    function messageHash(bytes memory _header, bytes memory _messageBody)
        internal
        pure
        returns (bytes32)
    {
        return keccak256(formatMessage(_header, _messageBody));
    }

    function messageView(bytes memory _message) internal pure returns (bytes29) {
        return _message.ref(MESSAGE_TYPE);
    }

    /// @notice Returns message's header field length
    function headerLength(bytes29 _message) internal pure onlyMessage(_message) returns (uint16) {
        return uint16(_message.indexUint(0, 2));
    }

    /// @notice Returns message's header field as bytes29 (refer to TypedMemView library for details on bytes29 type)
    function header(bytes29 _message) internal pure onlyMessage(_message) returns (bytes29) {
        return _message.slice(OFFSET_HEADER, headerLength(_message), HEADER_TYPE);
    }

    /// @notice Returns message's body field as bytes29 (refer to TypedMemView library for details on bytes29 type)
    function body(bytes29 _message) internal pure onlyMessage(_message) returns (bytes29) {
        uint256 bodyLength = _message.len() - (OFFSET_HEADER + headerLength(_message));
        return _message.postfix(bodyLength, BODY_TYPE);
    }

    /// @notice Returns leaf of the formatted message.
    function leaf(bytes29 _message) internal pure onlyMessage(_message) returns (bytes32) {
        // TODO: do we actually need this?
        return _message.keccak();
    }
}
