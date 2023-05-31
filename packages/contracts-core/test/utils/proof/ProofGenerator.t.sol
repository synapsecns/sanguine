// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {ORIGIN_TREE_HEIGHT} from "../../../contracts/libs/Constants.sol";
import {MerkleTree} from "../../../contracts/libs/merkle/MerkleTree.sol";

import {ProofCutter} from "./ProofCutter.t.sol";

// TODO: move from test directory
contract ProofGenerator is ProofCutter {
    /**
     * @notice Store only non-"zero" values of the merkle tree
     * merkleTree[0] are the leafs
     * merkleTree[1] are keccak256(A, B) where A and B are leafs
     * ...
     * merkleTree[ORIGIN_TREE_HEIGHT][0] is the merkle root
     */
    bytes32[][] internal merkleTree;

    constructor() {
        merkleTree = new bytes32[][](ORIGIN_TREE_HEIGHT + 1);
    }

    /**
     * @notice Creates a Merkle Tree from an array of leafs.
     */
    function createTree(bytes32[] memory leafs) external {
        _clearTree();
        // Copy the leafs into the tree
        uint256 size = _copyLeafs(leafs);
        // Go upwards the tree and construct parent layers one by one
        for (uint256 d = 1; d <= ORIGIN_TREE_HEIGHT; ++d) {
            size = (size + 1) / 2;
            _createLayer(d, size);
        }
    }

    /**
     * @notice Returns a merkle proof for leaf with a given index.
     */
    function getProof(uint256 index) external view returns (bytes32[] memory) {
        bytes32[] memory proof = new bytes32[](ORIGIN_TREE_HEIGHT);
        for (uint256 d = 0; d < ORIGIN_TREE_HEIGHT; ++d) {
            // Get node's neighbor
            if (index % 2 == 0) {
                ++index;
            } else {
                --index;
            }
            proof[d] = getNode(d, index);
            // We need to go deeper
            index = index / 2;
        }
        return cutProof(proof);
    }

    /**
     * @notice Returns merkle root of the tree.
     */
    function getRoot() external view returns (bytes32) {
        return merkleTree[ORIGIN_TREE_HEIGHT][0];
    }

    /**
     * @notice Returns node of the Merkle Tree given its depth and index.
     */
    function getNode(uint256 depth, uint256 index) public view returns (bytes32 node) {
        if (index < merkleTree[depth].length) {
            node = merkleTree[depth][index];
        } else {
            node = bytes32(0);
        }
    }

    /**
     * @notice Clears the merkle tree.
     */
    function _clearTree() internal {
        if (merkleTree[0].length != 0) {
            for (uint256 d = 0; d <= ORIGIN_TREE_HEIGHT; ++d) {
                delete merkleTree[d];
            }
        }
    }

    /**
     * @notice Copies the leafs into the leaf layer of the tree.
     */
    function _copyLeafs(bytes32[] memory leafs) internal returns (uint256 size) {
        size = leafs.length;
        merkleTree[0] = new bytes32[](size);
        for (uint256 i = 0; i < size; ++i) {
            merkleTree[0][i] = leafs[i];
        }
    }

    /**
     * @notice Creates a layer of the tree using its child layer.
     */
    function _createLayer(uint256 depth, uint256 size) internal {
        merkleTree[depth] = new bytes32[](size);
        for (uint256 i = 0; i < size; ++i) {
            merkleTree[depth][i] = _hash(getNode(depth - 1, 2 * i), getNode(depth - 1, 2 * i + 1));
        }
    }

    function _hash(bytes32 left, bytes32 right) internal pure returns (bytes32) {
        if (left != 0 || right != 0) {
            return keccak256(abi.encodePacked(left, right));
        } else {
            return 0;
        }
    }
}
