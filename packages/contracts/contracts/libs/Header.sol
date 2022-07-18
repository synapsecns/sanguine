// SPDX-License-Identifier: MIT
pragma solidity 0.8.13;

import { TypedMemView } from "./TypedMemView.sol";
import { TypeCasts } from "./TypeCasts.sol";
import { Message } from "./Message.sol";

/**
 * @notice Library for versioned formatting [the header part] of [the messages used by Home and Replicas].
 */
library Header {
    using TypedMemView for bytes;
    using TypedMemView for bytes29;

    uint16 internal constant HEADER_VERSION = 1;

    /**
     * @dev Header memory layout
     * [000 .. 002): version            uint16   2 bytes
     * [002 .. 006): originDomain       uint32   4 bytes
     * [006 .. 038): sender             bytes32 32 bytes
     * [038 .. 042): nonce              uint32   4 bytes
     * [042 .. 046): destinationDomain  uint32   4 bytes
     * [046 .. 078): recipient          bytes32 32 bytes
     * [078 .. 082): optimisticSeconds  uint32   4 bytes
     */

    uint256 private constant OFFSET_ORIGIN = 2;
    uint256 private constant OFFSET_SENDER = 6;
    uint256 private constant OFFSET_NONCE = 38;
    uint256 private constant OFFSET_DESTINATION = 42;
    uint256 private constant OFFSET_RECIPIENT = 46;
    uint256 private constant OFFSET_OPTIMISTIC_SECONDS = 78;

    modifier onlyHeader(bytes29 _view) {
        _view.assertType(Message.HEADER_TYPE);
        _;
    }

    function formatHeader(
        uint32 _originDomain,
        bytes32 _sender,
        uint32 _nonce,
        uint32 _destinationDomain,
        bytes32 _recipient,
        uint32 _optimisticSeconds
    ) internal pure returns (bytes memory) {
        return
            abi.encodePacked(
                HEADER_VERSION,
                _originDomain,
                _sender,
                _nonce,
                _destinationDomain,
                _recipient,
                _optimisticSeconds
            );
    }

    function headerView(bytes memory _header) internal pure returns (bytes29) {
        return _header.ref(Message.HEADER_TYPE);
    }

    function headerVersion(bytes29 _header) internal pure onlyHeader(_header) returns (uint16) {
        return uint16(_header.indexUint(0, 2));
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
