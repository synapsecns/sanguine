// SPDX-License-Identifier: MIT
pragma solidity 0.8.13;

import { MerkleLib } from "../../contracts/libs/Merkle.sol";

contract ProofGenerator {
    uint256 public constant TREE_DEPTH = 32;

    /**
     * @notice Store only non-"zero" values of the merkle tree
     * merkleTree[0] are the leafs
     * merkleTree[1] are keccak256(A, B) where A and B are leafs
     * ...
     * merkleTree[TREE_DEPTH][0] is the merkle root
     */
    bytes32[][] internal merkleTree;

    bytes32[] internal zeroHashes;

    constructor() {
        zeroHashes = MerkleLib.zeroHashes();
        merkleTree = new bytes32[][](TREE_DEPTH + 1);
    }

    /**
     * @notice Creates a Merkle Tree from an array of leafs.
     */
    function createTree(bytes32[] memory _leafs) external {
        _clearTree();
        // Copy the leafs into the tree
        uint256 size = _copyLeafs(_leafs);
        // Go upwards the tree and construct parent layers one by one
        for (uint256 d = 1; d <= TREE_DEPTH; ++d) {
            size = (size + 1) / 2;
            _createLayer(d, size);
        }
    }

    /**
     * @notice Returns a merkle proof for leaf with a given index.
     */
    function getProof(uint256 _index) external view returns (bytes32[TREE_DEPTH] memory proof) {
        for (uint256 d = 0; d < TREE_DEPTH; ++d) {
            // Get node's neighbor
            if (_index % 2 == 0) {
                ++_index;
            } else {
                --_index;
            }
            proof[d] = getNode(d, _index);
            // We need to go deeper
            _index = _index / 2;
        }
    }

    /**
     * @notice Returns merkle root of the tree.
     */
    function getRoot() external view returns (bytes32) {
        return merkleTree[TREE_DEPTH][0];
    }

    /**
     * @notice Returns node of the Merkle Tree given its depth and index.
     */
    function getNode(uint256 _depth, uint256 _index) public view returns (bytes32 node) {
        if (_index < merkleTree[_depth].length) {
            node = merkleTree[_depth][_index];
        } else {
            node = zeroHashes[_depth];
        }
    }

    /**
     * @notice Clears the merkle tree.
     */
    function _clearTree() internal {
        if (merkleTree[0].length != 0) {
            for (uint256 d = 0; d <= TREE_DEPTH; ++d) {
                delete merkleTree[d];
            }
        }
    }

    /**
     * @notice Copies the leafs into the leaf layer of the tree.
     */
    function _copyLeafs(bytes32[] memory _leafs) internal returns (uint256 size) {
        size = _leafs.length;
        merkleTree[0] = new bytes32[](size);
        for (uint256 i = 0; i < size; ++i) {
            merkleTree[0][i] = _leafs[i];
        }
    }

    /**
     * @notice Creates a layer of the tree using its child layer.
     */
    function _createLayer(uint256 _depth, uint256 _size) internal {
        merkleTree[_depth] = new bytes32[](_size);
        for (uint256 i = 0; i < _size; ++i) {
            merkleTree[_depth][i] = keccak256(
                abi.encodePacked(getNode(_depth - 1, 2 * i), getNode(_depth - 1, 2 * i + 1))
            );
        }
    }
}
