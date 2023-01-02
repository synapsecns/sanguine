// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import "./ByteString.sol";
import { TypeCasts } from "./TypeCasts.sol";

/// @dev Header is a memory over over a formatted message header payload.
type Header is bytes29;

/**
 * @notice Library for versioned formatting [the header part]
 * of [the messages used by Origin and Destination].
 */
library HeaderLib {
    using ByteString for bytes;
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
    ▏*║                              FORMATTERS                              ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

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
     * @notice Returns a Header view over for the given payload.
     * @dev Will revert if the payload is not a header payload.
     */
    function castToHeader(bytes memory _payload) internal pure returns (Header) {
        return castToHeader(_payload.castToRawBytes());
    }

    /**
     * @notice Casts a memory view to a Header view.
     * @dev Will revert if the memory view is not over a header payload.
     */
    function castToHeader(bytes29 _view) internal pure returns (Header) {
        require(isHeader(_view), "Not a header payload");
        return Header.wrap(_view);
    }

    /**
     * @notice Checks that a payload is a formatted Header.
     */
    function isHeader(bytes29 _view) internal pure returns (bool) {
        uint256 length = _view.len();
        // Check if version exists in the payload
        if (length < 2) return false;
        // Check that header version and its length matches
        return _getVersion(_view) == HEADER_VERSION && length == HEADER_LENGTH;
    }

    /// @notice Convenience shortcut for unwrapping a view.
    function unwrap(Header _header) internal pure returns (bytes29) {
        return Header.unwrap(_header);
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                            HEADER SLICING                            ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /// @notice Returns header's version field.
    function version(Header _header) internal pure returns (uint16) {
        // Get the underlying memory view
        bytes29 _view = Header.unwrap(_header);
        return _getVersion(_view);
    }

    /// @notice Returns header's origin field
    function origin(Header _header) internal pure returns (uint32) {
        bytes29 _view = Header.unwrap(_header);
        return uint32(_view.indexUint(OFFSET_ORIGIN, 4));
    }

    /// @notice Returns header's sender field
    function sender(Header _header) internal pure returns (bytes32) {
        bytes29 _view = Header.unwrap(_header);
        return _view.index(OFFSET_SENDER, 32);
    }

    /// @notice Returns header's nonce field
    function nonce(Header _header) internal pure returns (uint32) {
        bytes29 _view = Header.unwrap(_header);
        return uint32(_view.indexUint(OFFSET_NONCE, 4));
    }

    /// @notice Returns header's destination field
    function destination(Header _header) internal pure returns (uint32) {
        bytes29 _view = Header.unwrap(_header);
        return uint32(_view.indexUint(OFFSET_DESTINATION, 4));
    }

    /// @notice Returns header's recipient field as bytes32
    function recipient(Header _header) internal pure returns (bytes32) {
        bytes29 _view = Header.unwrap(_header);
        return _view.index(OFFSET_RECIPIENT, 32);
    }

    /// @notice Returns header's optimistic seconds field
    function optimisticSeconds(Header _header) internal pure returns (uint32) {
        bytes29 _view = Header.unwrap(_header);
        return uint32(_view.indexUint(OFFSET_OPTIMISTIC_SECONDS, 4));
    }

    /// @notice Returns header's recipient field as an address
    function recipientAddress(Header _header) internal pure returns (address) {
        return TypeCasts.bytes32ToAddress(recipient(_header));
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                           PRIVATE HELPERS                            ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /// @notice Returns a version field without checking if payload is properly formatted.
    function _getVersion(bytes29 _view) internal pure returns (uint16) {
        return uint16(_view.indexUint(OFFSET_VERSION, 2));
    }
}
