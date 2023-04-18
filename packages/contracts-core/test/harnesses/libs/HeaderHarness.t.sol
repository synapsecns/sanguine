// SPDX-License-Identifier: MIT

pragma solidity 0.8.17;

import {Header, HeaderLib, MemView, MemViewLib} from "../../../contracts/libs/Header.sol";

/**
 * @notice Exposes Header methods for testing against golang.
 */
contract HeaderHarness {
    using HeaderLib for bytes;
    using HeaderLib for MemView;
    using MemViewLib for bytes;

    // Note: we don't add an empty test() function here, as it currently leads
    // to zero coverage on the corresponding library.

    // ══════════════════════════════════════════════════ GETTERS ══════════════════════════════════════════════════════

    function castToHeader(bytes memory payload) public view returns (bytes memory) {
        // Walkaround to get the forge coverage working on libraries, see
        // https://github.com/foundry-rs/foundry/pull/3128#issuecomment-1241245086
        Header header = HeaderLib.castToHeader(payload);
        return header.unwrap().clone();
    }

    /// @notice Returns header's origin field
    function origin(bytes memory payload) public pure returns (uint32) {
        return payload.castToHeader().origin();
    }

    /// @notice Returns header's nonce field
    function nonce(bytes memory payload) public pure returns (uint32) {
        return payload.castToHeader().nonce();
    }

    /// @notice Returns header's destination field
    function destination(bytes memory payload) public pure returns (uint32) {
        return payload.castToHeader().destination();
    }

    /// @notice Returns header's optimistic seconds field
    function optimisticPeriod(bytes memory payload) public pure returns (uint32) {
        return payload.castToHeader().optimisticPeriod();
    }

    function isHeader(bytes memory payload) public pure returns (bool) {
        return payload.ref().isHeader();
    }

    // ════════════════════════════════════════════════ FORMATTERS ═════════════════════════════════════════════════════

    function formatHeader(uint32 origin_, uint32 nonce_, uint32 destination_, uint32 optimisticPeriod_)
        public
        pure
        returns (bytes memory)
    {
        return HeaderLib.formatHeader(origin_, nonce_, destination_, optimisticPeriod_);
    }
}
