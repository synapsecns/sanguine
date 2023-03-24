// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import { MerkleLib } from "../../../contracts/libs/Merkle.sol";

// TODO: move from test directory
contract ProofGenerator {
    uint256 public constant ORIGIN_TREE_HEIGHT = 32;

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
    function createTree(bytes32[] memory _leafs) external {
        _clearTree();
        // Copy the leafs into the tree
        uint256 size = _copyLeafs(_leafs);
        // Go upwards the tree and construct parent layers one by one
        for (uint256 d = 1; d <= ORIGIN_TREE_HEIGHT; ++d) {
            size = (size + 1) / 2;
            _createLayer(d, size);
        }
    }

    /**
     * @notice Returns a merkle proof for leaf with a given index.
     */
    function getProof(uint256 _index)
        external
        view
        returns (bytes32[ORIGIN_TREE_HEIGHT] memory proof)
    {
        for (uint256 d = 0; d < ORIGIN_TREE_HEIGHT; ++d) {
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
        return merkleTree[ORIGIN_TREE_HEIGHT][0];
    }

    /**
     * @notice Returns node of the Merkle Tree given its depth and index.
     */
    function getNode(uint256 _depth, uint256 _index) public view returns (bytes32 node) {
        if (_index < merkleTree[_depth].length) {
            node = merkleTree[_depth][_index];
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
            merkleTree[_depth][i] = _hash(
                getNode(_depth - 1, 2 * i),
                getNode(_depth - 1, 2 * i + 1)
            );
        }
    }

    function _hash(bytes32 _left, bytes32 _right) internal pure returns (bytes32) {
        if (_left != 0 || _right != 0) {
            return keccak256(abi.encodePacked(_left, _right));
        } else {
            return 0;
        }
    }
}
