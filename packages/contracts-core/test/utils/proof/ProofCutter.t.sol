// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

abstract contract ProofCutter {
    function cutProof(bytes32[] memory proof) public pure returns (bytes32[] memory shortened) {
        uint256 length = proof.length;
        // Figure out when trailing zeroes begin
        while (length > 0 && proof[length - 1] == bytes32(0)) --length;
        // Copy everything but the trailing zeroes
        shortened = new bytes32[](length);
        for (uint256 i = 0; i < length; ++i) {
            shortened[i] = proof[i];
        }
    }
}
