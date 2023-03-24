// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import { MerkleLib, AGENT_TREE_HEIGHT } from "../../../contracts/libs/Merkle.sol";

// TODO: move from test directory
contract DynamicProofGenerator {
    /**
     * @notice Store only non-"zero" values of the merkle tree
     * merkleTree[0] are the leafs
     * merkleTree[1] are keccak256(A, B) where A and B are leafs
     * ...
     * merkleTree[AGENT_TREE_HEIGHT][0] is the merkle root
     */
    mapping(uint256 => mapping(uint256 => bytes32)) internal merkleTree;

    /// @notice Updates a leaf in the Merkle Tree.
    function update(uint256 _index, bytes32 _newValue) external {
        merkleTree[0][_index] = _newValue;
        for (uint256 h = 1; h <= AGENT_TREE_HEIGHT; ++h) {
            // Traverse to parent
            _index >>= 1;
            merkleTree[h][_index] = MerkleLib.getParent(
                merkleTree[h - 1][2 * _index],
                merkleTree[h - 1][2 * _index + 1]
            );
        }
    }

    /// @notice Returns current value for the leaf.
    function getLeaf(uint256 _index) external view returns (bytes32) {
        return merkleTree[0][_index];
    }

    /// @notice Returns merkle root of the tree.
    function getRoot() external view returns (bytes32) {
        return merkleTree[AGENT_TREE_HEIGHT][0];
    }

    /// @notice Returns a merkle proof for leaf with a given index.
    function getProof(uint256 _index)
        external
        view
        returns (bytes32[AGENT_TREE_HEIGHT] memory proof)
    {
        for (uint256 h = 0; h < AGENT_TREE_HEIGHT; ++h) {
            // Get node's sibling
            proof[h] = merkleTree[h][_index ^ 1];
            // Traverse to parent
            _index >>= 1;
        }
    }
}
