// SPDX-License-Identifier: MIT

pragma solidity 0.8.17;

import "../../../contracts/libs/MerkleList.sol";

contract MerkleListHarness {
    function calculateRoot(bytes32[] memory hashes) public pure returns (bytes32) {
        MerkleList.calculateRoot(hashes);
        return hashes[0];
    }
}
