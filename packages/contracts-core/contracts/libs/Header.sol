// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import { TypedMemView } from "./TypedMemView.sol";
import { TypeCasts } from "./TypeCasts.sol";
import { SynapseTypes } from "./SynapseTypes.sol";

/**
 * @notice Library for versioned formatting [the header part]
 * of [the messages used by Origin and Destination].
 */
library Header {
    using TypedMemView for bytes;
    using TypedMemView for bytes29;

    uint16 internal constant HEADER_VERSION = 1;

    /**
     * @dev Header memory layout
     * [000 .. 002): version            uint16   2 bytes
     * [002 .. 006): origin             uint32   4 bytes
     * [006 .. 038): sender             bytes32 32 bytes
     * [038 .. 042): nonce              uint32   4 bytes
     * [042 .. 046): destination        uint32   4 bytes
     * [046 .. 078): recipient          bytes32 32 bytes
     * [078 .. 082): optimisticSeconds  uint32   4 bytes
     */

    uint256 internal constant OFFSET_VERSION = 0;
    uint256 internal constant OFFSET_ORIGIN = 2;
    uint256 internal constant OFFSET_SENDER = 6;
    uint256 internal constant OFFSET_NONCE = 38;
    uint256 internal constant OFFSET_DESTINATION = 42;
    uint256 internal constant OFFSET_RECIPIENT = 46;
    uint256 internal constant OFFSET_OPTIMISTIC_SECONDS = 78;

    uint256 internal constant HEADER_LENGTH = 82;

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                              MODIFIERS                               ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    modifier onlyHeader(bytes29 _view) {
        _view.assertType(SynapseTypes.MESSAGE_HEADER);
        _;
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                              FORMATTERS                              ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @notice Returns a properly typed bytes29 pointer for a header payload.
     */
    function castToHeader(bytes memory _payload) internal pure returns (bytes29) {
        return _payload.ref(SynapseTypes.MESSAGE_HEADER);
    }

    /**
     * @notice Returns a formatted Header payload with provided fields
     * @param _origin               Domain of origin chain
     * @param _sender               Address that sent the message
     * @param _nonce                Message nonce on origin chain
     * @param _destination          Domain of destination chain
     * @param _recipient            Address that will receive the message
     * @param _optimisticSeconds    Optimistic period for message execution
     * @return Formatted header
     **/
    function formatHeader(
        uint32 _origin,
        bytes32 _sender,
        uint32 _nonce,
        uint32 _destination,
        bytes32 _recipient,
        uint32 _optimisticSeconds
    ) internal pure returns (bytes memory) {
        return
            abi.encodePacked(
                HEADER_VERSION,
                _origin,
                _sender,
                _nonce,
                _destination,
                _recipient,
                _optimisticSeconds
            );
    }

    /**
     * @notice Checks that a payload is a formatted Header.
     */
    function isHeader(bytes29 _view) internal pure returns (bool) {
        uint256 length = _view.len();
        // Check if version exists in the payload
        if (length < 2) return false;
        // Check that header version and its length matches
        return headerVersion(_view) == HEADER_VERSION && length == HEADER_LENGTH;
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                            HEADER SLICING                            ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /// @notice Returns header's version field.
    function headerVersion(bytes29 _header) internal pure onlyHeader(_header) returns (uint16) {
        return uint16(_header.indexUint(OFFSET_VERSION, 2));
    }

    /// @notice Returns header's origin field
    function origin(bytes29 _header) internal pure onlyHeader(_header) returns (uint32) {
        return uint32(_header.indexUint(OFFSET_ORIGIN, 4));
    }

    /// @notice Returns header's sender field
    function sender(bytes29 _header) internal pure onlyHeader(_header) returns (bytes32) {
        return _header.index(OFFSET_SENDER, 32);
    }

    /// @notice Returns header's nonce field
    function nonce(bytes29 _header) internal pure onlyHeader(_header) returns (uint32) {
        return uint32(_header.indexUint(OFFSET_NONCE, 4));
    }

    /// @notice Returns header's destination field
    function destination(bytes29 _header) internal pure onlyHeader(_header) returns (uint32) {
        return uint32(_header.indexUint(OFFSET_DESTINATION, 4));
    }

    /// @notice Returns header's recipient field as bytes32
    function recipient(bytes29 _header) internal pure onlyHeader(_header) returns (bytes32) {
        return _header.index(OFFSET_RECIPIENT, 32);
    }

    /// @notice Returns header's optimistic seconds field
    function optimisticSeconds(bytes29 _header) internal pure onlyHeader(_header) returns (uint32) {
        return uint32(_header.indexUint(OFFSET_OPTIMISTIC_SECONDS, 4));
    }

    /// @notice Returns header's recipient field as an address
    function recipientAddress(bytes29 _header) internal pure returns (address) {
        return TypeCasts.bytes32ToAddress(recipient(_header));
    }
}
