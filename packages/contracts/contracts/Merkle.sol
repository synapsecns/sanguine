// SPDX-License-Identifier: MIT
pragma solidity 0.8.13;

// ============ Internal Imports ============
import { MerkleLib } from "./libs/Merkle.sol";

/**
 * @title MerkleTreeManager
 * @author Illusory Systems Inc.
 * @notice Contains a Merkle tree instance and
 * exposes view functions for the tree.
 */
contract MerkleTreeManager {
    // ============ Libraries ============

    using MerkleLib for MerkleLib.Tree;
    MerkleLib.Tree public tree;
    bytes32[] public historicalRoots;

    // ============ Upgrade Gap ============

    // gap for upgrade safety
    uint256[48] private __GAP;

    // ============ Public Functions ============

    /**
     * @notice Calculates and returns tree's current root
     */
    function root() public view returns (bytes32) {
        return tree.root();
    }

    /**
     * @notice Returns the number of inserted leaves in the tree (current index)
     */
    function count() public view returns (uint256) {
        return tree.count;
    }

    // ============ Internal Functions ============

    /**
     * @notice Inserts _hash into the Merkle tree and stores the new merkle root.
     */
    function _insertHash(bytes32 _hash) internal {
        tree.insert(_hash);
        historicalRoots.push(tree.root());
    }
}
