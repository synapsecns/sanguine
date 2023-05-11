// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {MerkleMath, AGENT_TREE_HEIGHT} from "../../../contracts/libs/merkle/MerkleTree.sol";

import {ProofCutter} from "./ProofCutter.t.sol";

// TODO: move from test directory
contract DynamicProofGenerator is ProofCutter {
    /**
     * @notice Store only non-"zero" values of the merkle tree
     * merkleTree[0] are the leafs
     * merkleTree[1] are keccak256(A, B) where A and B are leafs
     * ...
     * merkleTree[AGENT_TREE_HEIGHT][0] is the merkle root
     */
    mapping(uint256 => mapping(uint256 => bytes32)) internal merkleTree;

    /// @notice Updates a leaf in the Merkle Tree.
    function update(uint256 index, bytes32 newValue) external {
        merkleTree[0][index] = newValue;
        for (uint256 h = 1; h <= AGENT_TREE_HEIGHT; ++h) {
            // Traverse to parent
            index >>= 1;
            merkleTree[h][index] = MerkleMath.getParent(merkleTree[h - 1][2 * index], merkleTree[h - 1][2 * index + 1]);
        }
    }

    /// @notice Returns current value for the leaf.
    function getLeaf(uint256 index) external view returns (bytes32) {
        return merkleTree[0][index];
    }

    /// @notice Returns merkle root of the tree.
    function getRoot() external view returns (bytes32) {
        return merkleTree[AGENT_TREE_HEIGHT][0];
    }

    /// @notice Returns a merkle proof for leaf with a given index.
    function getProof(uint256 index) external view returns (bytes32[] memory) {
        bytes32[] memory proof = new bytes32[](AGENT_TREE_HEIGHT);
        for (uint256 h = 0; h < AGENT_TREE_HEIGHT; ++h) {
            // Get node's sibling
            proof[h] = merkleTree[h][index ^ 1];
            // Traverse to parent
            index >>= 1;
        }
        return cutProof(proof);
    }
}
