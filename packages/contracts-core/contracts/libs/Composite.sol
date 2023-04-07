// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

library Composite {
    /// @notice Merges two uint32 values into a combined uint64 value.
    function mergeUint32(uint32 first, uint32 second) internal pure returns (uint64 combined) {
        return uint64(first) << 32 | second;
    }

    /// @notice Splits Merges a combined uint64 value into two uint32 values.
    function splitUint32(uint64 combined) internal pure returns (uint32 first, uint32 second) {
        first = uint32(combined >> 32);
        second = uint32(combined & type(uint32).max);
    }
}
