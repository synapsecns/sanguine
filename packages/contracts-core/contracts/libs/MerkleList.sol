// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import { MerkleLib } from "./Merkle.sol";

library MerkleList {
    /**
     * @notice Calculates merkle root for a list of given leafs.
     * Merkle Tree is constructed by padding the list with ZERO values for leafs
     * until list length is a power of two.
     * Merkle Root is calculated for the constructed tree, and recorded in leafs[0].
     * Note: `leafs` values are overwritten in the process to avoid excessive memory allocations.
     * Caller is expected not to allocate memory for the leafs list, and only use leafs[0] value,
     * which is guaranteed to contain the calculated merkle root.
     * @param hashes    List of leafs for the merkle tree (to be overwritten)
     */
    function calculateRoot(bytes32[] memory hashes) internal pure {
        uint256 levelLength = hashes.length;
        // We will be iterating from the "leafs level" up to the "root level" of the Merkle Tree.
        // For every level we will only record "significant values", i.e. not equal to ZERO
        // Repeat until we only have a single hash: this would be the root of the tree
        while (levelLength > 1) {
            // Let H be the height of the "current level". H = 0 for the "root level".
            // Invariant: hashes[0 .. length) are "current level" tree nodes
            // Invariant: bytes32(0) is the value for nodes with indexes [length .. 2**H)

            // Iterate over every pair of (leftChild, rightChild) on the current level
            for (uint256 leftIndex = 0; leftIndex < levelLength; leftIndex += 2) {
                uint256 rightIndex = leftIndex + 1;
                bytes32 leftChild = hashes[leftIndex];
                // Note: rightChild might be ZERO
                bytes32 rightChild = rightIndex < levelLength ? hashes[rightIndex] : bytes32(0);
                // Record the parent hash in the same array. This will not affect
                // further calculations for the same level: (leftIndex >> 1) <= leftIndex.
                hashes[leftIndex >> 1] = MerkleLib.getParent(leftChild, rightChild);
            }
            // Set length for the "parent level"
            levelLength = (levelLength + 1) >> 1;
        }
    }

    function calculateProof(
        bytes32[] memory hashes,
        uint256 index,
        bytes32[] memory proof
    ) internal pure {
        // proof[0] is already set up
        uint256 height = 0;

        uint256 levelLength = hashes.length;

        // We will be iterating from the "leafs level" up to the "root level" of the Merkle Tree.
        // For every level we will only record "significant values", i.e. not equal to ZERO
        // Repeat until we only have a single hash: this would be the root of the tree
        while (levelLength > 1) {
            // Use sibling for the merkle proof
            proof[++height] = (index ^ 1 < levelLength) ? hashes[index ^ 1] : bytes32(0);

            // Let H be the height of the "current level". H = 0 for the "root level".
            // Invariant: hashes[0 .. length) are "current level" tree nodes
            // Invariant: bytes32(0) is the value for nodes with indexes [length .. 2**H)

            // Iterate over every pair of (leftChild, rightChild) on the current level
            for (uint256 leftIndex = 0; leftIndex < levelLength; leftIndex += 2) {
                uint256 rightIndex = leftIndex + 1;
                bytes32 leftChild = hashes[leftIndex];
                // Note: rightChild might be ZERO
                bytes32 rightChild = rightIndex < levelLength ? hashes[rightIndex] : bytes32(0);
                // Record the parent hash in the same array. This will not affect
                // further calculations for the same level: (leftIndex >> 1) <= leftIndex.
                hashes[leftIndex >> 1] = MerkleLib.getParent(leftChild, rightChild);
            }
            // Set length for the "parent level"
            levelLength = (levelLength + 1) >> 1;
            // Traverse to parent node
            index >>= 1;
        }
    }
}
