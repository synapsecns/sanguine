// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import { DynamicTree, AGENT_TREE_HEIGHT } from "../../../contracts/libs/Merkle.sol";

contract DynamicTreeHarness {
    DynamicTree internal tree;

    function update(
        uint256 _index,
        bytes32 _oldValue,
        bytes32[AGENT_TREE_HEIGHT] memory _branch,
        bytes32 _newValue
    ) external {
        tree.update(_index, _oldValue, _branch, _newValue);
    }

    function root() external view returns (bytes32) {
        return tree.root;
    }
}
