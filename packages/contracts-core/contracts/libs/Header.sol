// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

/// Header is encoded data with "general routing information".
type Header is uint128;

using HeaderLib for Header global;

/// Library for formatting _the header part_ of _the messages used by Origin and Destination_.
/// - Header represents general information for routing a Message for Origin and Destination.
/// - Header occupies a single storage word, and thus is stored on stack instead of being stored in memory.
///
/// # Header stack layout (from highest bits to lowest)
///
/// | Position   | Field            | Type   | Bytes | Description                             |
/// | ---------- | ---------------- | ------ | ----- | --------------------------------------- |
/// | (016..012] | origin           | uint32 | 4     | Domain where message originated         |
/// | (012..008] | nonce            | uint32 | 4     | Message nonce on the origin domain      |
/// | (008..004] | destination      | uint32 | 4     | Domain where message will be executed   |
/// | (004..000] | optimisticPeriod | uint32 | 4     | Optimistic period that will be enforced |
library HeaderLib {
    /// @dev Amount of bits to shift to origin field
    uint128 private constant SHIFT_ORIGIN = 12 * 8;
    /// @dev Amount of bits to shift to nonce field
    uint128 private constant SHIFT_NONCE = 8 * 8;
    /// @dev Amount of bits to shift to destination field
    uint128 private constant SHIFT_DESTINATION = 4 * 8;

    /// @notice Returns an encoded header with provided fields
    /// @param origin_              Domain of origin chain
    /// @param nonce_               Message nonce on origin chain
    /// @param destination_         Domain of destination chain
    /// @param optimisticPeriod_    Optimistic period for message execution
    function encodeHeader(uint32 origin_, uint32 nonce_, uint32 destination_, uint32 optimisticPeriod_)
        internal
        pure
        returns (Header)
    {
        return Header.wrap(
            uint128(origin_) << SHIFT_ORIGIN | uint128(nonce_) << SHIFT_NONCE
                | uint128(destination_) << SHIFT_DESTINATION | uint128(optimisticPeriod_)
        );
    }

    /// @notice Wraps the padded encoded request into a Header-typed value.
    /// @dev The "padded" header is simply an encoded header casted to uint256 (highest bits are set to zero).
    /// Casting to uint256 is done automatically in Solidity, so no extra actions from consumers are needed.
    /// The highest bits are discarded, so that the contracts dealing with encoded headers
    /// don't need to be updated, if a new field is added.
    function wrapPadded(uint256 paddedHeader) internal pure returns (Header) {
        return Header.wrap(uint128(paddedHeader));
    }

    // ══════════════════════════════════════════════ HEADER SLICING ═══════════════════════════════════════════════════

    /// @notice Returns header's origin field
    function origin(Header header) internal pure returns (uint32) {
        // Casting to uint32 will truncate the highest bits, which is the behavior we want
        return uint32(Header.unwrap(header) >> SHIFT_ORIGIN);
    }

    /// @notice Returns header's nonce field
    function nonce(Header header) internal pure returns (uint32) {
        // Casting to uint32 will truncate the highest bits, which is the behavior we want
        return uint32(Header.unwrap(header) >> SHIFT_NONCE);
    }

    /// @notice Returns header's destination field
    function destination(Header header) internal pure returns (uint32) {
        // Casting to uint32 will truncate the highest bits, which is the behavior we want
        return uint32(Header.unwrap(header) >> SHIFT_DESTINATION);
    }

    /// @notice Returns header's optimistic seconds field
    function optimisticPeriod(Header header) internal pure returns (uint32) {
        // Casting to uint32 will truncate the highest bits, which is the behavior we want
        return uint32(Header.unwrap(header));
    }
}
