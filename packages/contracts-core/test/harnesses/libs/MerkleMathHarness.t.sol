// SPDX-License-Identifier: MIT

pragma solidity 0.8.17;

import {MerkleMath} from "../../../contracts/libs/MerkleMath.sol";

contract MerkleMathHarness {
    // Note: we don't add an empty test() function here, as it currently leads
    // to zero coverage on the corresponding library.

    function calculateRoot(bytes32[] memory hashes, uint256 height) public pure returns (bytes32) {
        MerkleMath.calculateRoot(hashes, height);
        return hashes[0];
    }

    function calculateProof(bytes32[] memory hashes, uint256 index) public pure returns (bytes32[] memory proof) {
        return MerkleMath.calculateProof(hashes, index);
    }
}
