// SPDX-License-Identifier: MIT

pragma solidity 0.8.17;

import {Header, HeaderLib, MessageFlag} from "../../../../contracts/libs/stack/Header.sol";

/**
 * @notice Exposes Header methods for testing against golang.
 */
contract HeaderHarness {
    // Note: we don't add an empty test() function here, as it currently leads
    // to zero coverage on the corresponding library.

    // ══════════════════════════════════════════════════ GETTERS ══════════════════════════════════════════════════════

    /// @notice Returns header's flag field
    function flag(uint256 paddedHeader) public pure returns (MessageFlag) {
        return HeaderLib.wrapPadded(paddedHeader).flag();
    }

    /// @notice Returns header's origin field
    function origin(uint256 paddedHeader) public pure returns (uint32) {
        return HeaderLib.wrapPadded(paddedHeader).origin();
    }

    /// @notice Returns header's nonce field
    function nonce(uint256 paddedHeader) public pure returns (uint32) {
        return HeaderLib.wrapPadded(paddedHeader).nonce();
    }

    /// @notice Returns header's destination field
    function destination(uint256 paddedHeader) public pure returns (uint32) {
        return HeaderLib.wrapPadded(paddedHeader).destination();
    }

    /// @notice Returns header's optimistic seconds field
    function optimisticPeriod(uint256 paddedHeader) public pure returns (uint32) {
        return HeaderLib.wrapPadded(paddedHeader).optimisticPeriod();
    }

    // ════════════════════════════════════════════════ FORMATTERS ═════════════════════════════════════════════════════

    function encodeHeader(uint8 flag_, uint32 origin_, uint32 nonce_, uint32 destination_, uint32 optimisticPeriod_)
        public
        pure
        returns (uint136)
    {
        Header header = HeaderLib.encodeHeader(MessageFlag(flag_), origin_, nonce_, destination_, optimisticPeriod_);
        return Header.unwrap(header);
    }

    function isHeader(uint256 paddedHeader) public pure returns (bool) {
        return HeaderLib.isHeader(paddedHeader);
    }

    function wrapPadded(uint256 paddedHeader) public pure returns (uint136) {
        Header header = HeaderLib.wrapPadded(paddedHeader);
        return Header.unwrap(header);
    }

    function leaf(uint256 paddedHeader) public pure returns (bytes32) {
        return HeaderLib.wrapPadded(paddedHeader).leaf();
    }
}
