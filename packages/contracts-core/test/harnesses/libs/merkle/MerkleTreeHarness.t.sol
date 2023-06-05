// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {BaseTree, HistoricalTree, DynamicTree, MerkleTree} from "../../../../contracts/libs/merkle/MerkleTree.sol";

// solhint-disable ordering
contract MerkleTreeHarness {
    BaseTree internal baseTree;
    HistoricalTree internal historicalTree;
    DynamicTree internal dynamicTree;

    function insertBase(uint256 newCount, bytes32 node) public {
        baseTree.insertBase(newCount, node);
    }

    function rootBase(uint256 count) public view returns (bytes32) {
        // Walkaround to get the forge coverage working on libraries, see
        // https://github.com/foundry-rs/foundry/pull/3128#issuecomment-1241245086
        bytes32 result = MerkleTree.rootBase(baseTree, count);
        return result;
    }

    function initializeRoots() public returns (bytes32) {
        return historicalTree.initializeRoots();
    }

    function insert(bytes32 node) public returns (bytes32) {
        return historicalTree.insert(node);
    }

    function root(uint256 count) public view returns (bytes32) {
        return historicalTree.root(count);
    }

    function update(uint256 index, bytes32 oldValue, bytes32[] memory branch, bytes32 newValue)
        public
        returns (bytes32)
    {
        return dynamicTree.update(index, oldValue, branch, newValue);
    }
}
