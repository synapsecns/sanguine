// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {Number, NumberLib} from "../../../../contracts/libs/stack/Number.sol";

/**
 * @notice Exposes NumberLib methods for testing against golang.
 */
contract NumberHarness {
    // Note: we don't add an empty test() function here, as it currently leads
    // to zero coverage on the corresponding library.

    function compress(uint256 value) public pure returns (Number) {
        return NumberLib.compress(value);
    }

    function decompress(Number number) public pure returns (uint256) {
        return NumberLib.decompress(number);
    }

    function mostSignificantBit(uint256 x) public pure returns (uint256) {
        return NumberLib.mostSignificantBit(x);
    }
}
