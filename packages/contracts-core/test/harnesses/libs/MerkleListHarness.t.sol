// SPDX-License-Identifier: MIT

pragma solidity 0.8.17;

import "../../../contracts/libs/MerkleList.sol";

contract MerkleListHarness {
    function calculateRoot(bytes32[] memory hashes) public pure returns (bytes32) {
        MerkleList.calculateRoot(hashes);
        return hashes[0];
    }

    function calculateRoot(bytes32[] memory hashes, bytes32 zeroHash)
        public
        pure
        returns (bytes32)
    {
        MerkleList.calculateRoot(hashes, zeroHash);
        return hashes[0];
    }
}
