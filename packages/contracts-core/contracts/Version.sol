// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

/**
 * @title Versioned
 * @notice Version getter for contracts. Doesn't use any storage slots, meaning
 * it will never cause any troubles with the upgradeable contracts. For instance, this contract
 * can be added or removed from the inheritance chain without shifting the storage layout.
 **/
abstract contract Versioned {
    /**
     * @notice Struct that is mimicking the storage layout of a string with 32 bytes or less.
     * Length is limited by 32, so the whole string payload takes two memory words:
     * @param length    String length
     * @param data      String characters
     */
    struct _ShortString {
        uint256 length;
        bytes32 data;
    }

    /// @dev Length of the "version string"
    uint256 private immutable _length;
    /// @dev Bytes representation of the "version string".
    /// Strings with length over 32 are not supported!
    bytes32 private immutable _data;

    constructor(string memory _version) {
        _length = bytes(_version).length;
        require(_length <= 32, "String length over 32");
        // bytes32 is left-aligned => this will store the byte representation of the string
        // with the trailing zeroes to complete the 32-byte word
        _data = bytes32(bytes(_version));
    }

    function version() external view returns (string memory versionString) {
        // Load the immutable values to form the version string
        _ShortString memory str = _ShortString(_length, _data);
        // The only way to do this cast is doing some dirty assembly
        assembly {
            // solhint-disable-previous-line no-inline-assembly
            versionString := str
        }
    }
}

abstract contract Version0_0_1 is Versioned {
    // solhint-disable-next-line no-empty-blocks
    constructor() Versioned("0.0.1") {}
}
