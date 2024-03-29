// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

library MathLib {
    /// @notice Rounds up to the nearest multiple of 32.
    /// Note: Returns zero on overflows instead of reverting. This is fine for practical
    /// use cases, as this is used for determining the size of the payload in memory.
    function roundUpToWord(uint256 x) internal pure returns (uint256) {
        unchecked {
            return (x + 31) & ~uint256(31);
        }
    }
}
