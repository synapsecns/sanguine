// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {DynamicTree} from "../../../../contracts/libs/merkle/MerkleTree.sol";

contract DynamicTreeHarness {
    DynamicTree internal tree;

    function update(uint256 index, bytes32 oldValue, bytes32[] memory branch, bytes32 newValue) external {
        tree.update(index, oldValue, branch, newValue);
    }

    function root() external view returns (bytes32) {
        return tree.root;
    }
}
