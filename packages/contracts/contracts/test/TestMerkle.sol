// SPDX-License-Identifier: MIT
pragma solidity 0.8.13;

import "../Merkle.sol";

contract TestMerkle is MerkleTreeManager {
    using MerkleLib for MerkleLib.Tree;

    // solhint-disable-next-line no-empty-blocks
    constructor() MerkleTreeManager() {}

    function insert(bytes32 _node) external {
        tree.insert(_node);
    }

    function branchRoot(
        bytes32 _leaf,
        bytes32[32] calldata _proof,
        uint256 _index
    ) external pure returns (bytes32 _node) {
        return MerkleLib.branchRoot(_leaf, _proof, _index);
    }
}
