// SPDX-License-Identifier: MIT
pragma solidity ^0.8.4;

library ReadableSymbol {
    error ReadableSymbol__StringTooLong();

    /// @notice Converts a string to a bytes31. String must be 31 characters or less.
    function toBytes31(string memory str) internal pure returns (bytes31) {
        if (bytes(str).length > 31) revert ReadableSymbol__StringTooLong();
        return bytes31(bytes(str));
    }

    /// @notice Converts a bytes31 to a string.
    /// @dev This method is not optimized and should be used for off-chain calls only.
    /// Additionally, everything after the first null byte is dropped.
    function toString(bytes31 symbol) internal pure returns (string memory) {
        // Find the first null byte
        uint256 length = 31;
        for (uint256 i = 0; i < 31; i++) {
            if (symbol[i] == 0) {
                length = i;
                break;
            }
        }
        bytes memory result = new bytes(length);
        for (uint256 i = 0; i < length; i++) {
            result[i] = symbol[i];
        }
        return string(result);
    }
}
