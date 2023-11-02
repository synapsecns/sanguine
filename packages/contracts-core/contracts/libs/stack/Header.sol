// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {FlagOutOfRange} from "../Errors.sol";

/// Header is encoded data with "general routing information".
type Header is uint136;

using HeaderLib for Header global;

/// Types of messages supported by Origin-Destination
/// - Base: message sent by protocol user, contains tips
/// - Manager: message sent between AgentManager contracts located on different chains, no tips
enum MessageFlag {
    Base,
    Manager
}

using HeaderLib for MessageFlag global;

/// Library for formatting _the header part_ of _the messages used by Origin and Destination_.
/// - Header represents general information for routing a Message for Origin and Destination.
/// - Header occupies a single storage word, and thus is stored on stack instead of being stored in memory.
///
/// # Header stack layout (from highest bits to lowest)
///
/// | Position   | Field            | Type   | Bytes | Description                             |
/// | ---------- | ---------------- | ------ | ----- | --------------------------------------- |
/// | (017..016] | flag             | uint8  | 1     | Flag specifying the type of message     |
/// | (016..012] | origin           | uint32 | 4     | Domain where message originated         |
/// | (012..008] | nonce            | uint32 | 4     | Message nonce on the origin domain      |
/// | (008..004] | destination      | uint32 | 4     | Domain where message will be executed   |
/// | (004..000] | optimisticPeriod | uint32 | 4     | Optimistic period that will be enforced |
library HeaderLib {
    /// @dev Amount of bits to shift to flag field
    uint136 private constant SHIFT_FLAG = 16 * 8;
    /// @dev Amount of bits to shift to origin field
    uint136 private constant SHIFT_ORIGIN = 12 * 8;
    /// @dev Amount of bits to shift to nonce field
    uint136 private constant SHIFT_NONCE = 8 * 8;
    /// @dev Amount of bits to shift to destination field
    uint136 private constant SHIFT_DESTINATION = 4 * 8;

    /// @notice Returns an encoded header with provided fields
    /// @param origin_              Domain of origin chain
    /// @param nonce_               Message nonce on origin chain
    /// @param destination_         Domain of destination chain
    /// @param optimisticPeriod_    Optimistic period for message execution
    function encodeHeader(
        MessageFlag flag_,
        uint32 origin_,
        uint32 nonce_,
        uint32 destination_,
        uint32 optimisticPeriod_
    ) internal pure returns (Header) {
        // All casts below are upcasts, so they are safe
        // forgefmt: disable-next-item
        return Header.wrap(
            uint136(uint8(flag_)) << SHIFT_FLAG |
            uint136(origin_) << SHIFT_ORIGIN |
            uint136(nonce_) << SHIFT_NONCE |
            uint136(destination_) << SHIFT_DESTINATION |
            uint136(optimisticPeriod_)
        );
    }

    /// @notice Checks that the header is a valid encoded header.
    function isHeader(uint256 paddedHeader) internal pure returns (bool) {
        // Check that flag is within range
        return _flag(paddedHeader) <= uint8(type(MessageFlag).max);
    }

    /// @notice Wraps the padded encoded request into a Header-typed value.
    /// @dev The "padded" header is simply an encoded header casted to uint256 (highest bits are set to zero).
    /// Casting to uint256 is done automatically in Solidity, so no extra actions from consumers are needed.
    /// The highest bits are discarded, so that the contracts dealing with encoded headers
    /// don't need to be updated, if a new field is added.
    function wrapPadded(uint256 paddedHeader) internal pure returns (Header) {
        // Check that flag is within range
        if (!isHeader(paddedHeader)) revert FlagOutOfRange();
        // Casting to uint136 will truncate the highest bits, which is the behavior we want
        return Header.wrap(uint136(paddedHeader));
    }

    /// @notice Returns header's hash: a leaf to be inserted in the "Message mini-Merkle tree".
    function leaf(Header header) internal pure returns (bytes32 hashedHeader) {
        // solhint-disable-next-line no-inline-assembly
        assembly {
            // Store header in scratch space
            mstore(0, header)
            // Compute hash of header padded to 32 bytes
            hashedHeader := keccak256(0, 32)
        }
    }

    // ══════════════════════════════════════════════ HEADER SLICING ═══════════════════════════════════════════════════

    /// @notice Returns header's flag field
    function flag(Header header) internal pure returns (MessageFlag) {
        // We check that flag is within range when wrapping the header, so this cast is safe
        return MessageFlag(_flag(Header.unwrap(header)));
    }

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

    // ══════════════════════════════════════════════ PRIVATE HELPERS ══════════════════════════════════════════════════

    /// @dev Returns header's flag field without casting to MessageFlag
    function _flag(uint256 paddedHeader) private pure returns (uint8) {
        // Casting to uint8 will truncate the highest bits, which is the behavior we want
        return uint8(paddedHeader >> SHIFT_FLAG);
    }
}
