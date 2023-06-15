// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {ProofCutter} from "./ProofCutter.t.sol";

// TODO: move from test directory
contract HistoricalProofGenerator is ProofCutter {
    uint256 public constant ORIGIN_TREE_HEIGHT = 32;

    /**
     * @notice Store historical non-"zero" values of the FULL merkle tree.
     * Full merkle tree consists of 2**ORIGIN_TREE_HEIGHT "zero" leafs, which are
     * getting populated throughout time. Once a new leaf is added, all elements
     * in the merkle tree on the path from root to the leaf are updated.
     * The goal of this contract is to store only the significant values.
     *
     * merkleTree[H][X][N] is the value for the tree element:
     * - With [height = H] (increasing from leafs to root)
     * - With [x-coord = X] (increasing from older leafs to newer)
     * - When N leafs were inserted in the merkle tree
     *
     * 1. Height (H):
     * merkleTree[0] are the leafs
     * merkleTree[1] are keccak256(A, B) where A and B are leafs
     * ...
     * merkleTree[ORIGIN_TREE_HEIGHT] is the merkle root level
     *
     * 2. Coordinate (X):
     * A merkle tree can have up to 2**(32-H) elements on a level with height=H
     * Therefore:
     * merkleTree[0][0] is the first leaf
     * merkleTree[0][1] is the second leaf
     * merkleTree[1][0] is their parent
     * merkleTree[1][1] is parent of merkleTree[0][2] and merkleTree[0][3]
     * merkleTree[2][0] is parent of merkleTree[1][0] and merkleTree[1][1]
     * ...
     * merkleTree[ORIGIN_TREE_HEIGHT][0] is the merkle root
     *
     * 3. Historical state (N).
     * Every element of the full merkle tree has three chronological "stages".
     * a. Element value did not change after the latest leaf insertion. Meaning that
     *    all element's children are "zero" elements, and element itself is "zero".
     *    Requires: 0 <= N <= X*(2**H)
     * b. Element value changed after the latest leaf insertion. Meaning that
     *    at least one of the children is non-zero.
     *    Requires: X*(2**H) < N <= (X+1)*(2**H)
     * c. Element value stopped changing after the latest leaf insertion. Meaning that
     *    all element children are already non-zero.
     *    Requires: (X+1)*(2**H) < N
     *
     * Thus we actually need to store tree element value for N in range (X*(2**H), (X+1)*(2**H)]
     * The amount of "significant" values (stage b) is 2**H.
     *
     * We're using mapping to avoid dealing with dynamic arrays in Solidity.
     */
    mapping(uint256 => mapping(uint256 => mapping(uint256 => bytes32))) internal merkleTree;
    // Amount of inserted leaves in the tree
    uint256 internal treeCount;

    /**
     * @notice Insert a new leaf into the tree and update the historical
     * merkle tree. O(1)
     */
    function insert(bytes32 leaf) external {
        uint256 x = treeCount;
        uint256 newCount = x + 1;
        merkleTree[0][x][newCount] = leaf;
        for (uint256 h = 1; h <= ORIGIN_TREE_HEIGHT; ++h) {
            // Traverse to parent
            x = x >> 1;
            // Children have [height = h - 1]
            // And X-coordinates [2 * x] and [2 * x + 1]
            merkleTree[h][x][newCount] = _hash(
                _fetchSavedTreeElement(h - 1, (x << 1), newCount), _fetchSavedTreeElement(h - 1, (x << 1) + 1, newCount)
            );
        }
        ++treeCount;
    }

    /// @notice Returns leaf with a given index.
    function getLeaf(uint256 index) external view returns (bytes32) {
        return merkleTree[0][index][index + 1];
    }

    /**
     * @notice Calculate root of merkle tree at the time when
     * `count` leafs have been inserted. O(1)
     */
    function getRoot(uint256 count) external view returns (bytes32) {
        require(count <= treeCount, "Not enough leafs inserted");
        return _fetchSavedTreeElement({h: ORIGIN_TREE_HEIGHT, x: 0, count: count});
    }

    /**
     * @notice Generate proof of inclusion for leaf with given `index`,
     * for the current merkle tree. O(1)
     */
    function getLatestProof(uint256 index) external view returns (bytes32[] memory proof) {
        return this.getProof(index, treeCount);
    }

    /**
     * @notice Generate proof of inclusion for leaf with given `index`,
     * at the time when `count` leafs have been inserted. O(1)
     */
    function getProof(uint256 index, uint256 count) external view returns (bytes32[] memory) {
        require(index < count, "Out of range");
        require(count <= treeCount, "Not enough leafs inserted");
        bytes32[] memory proof = new bytes32[](ORIGIN_TREE_HEIGHT);
        for (uint256 h = 0; h < ORIGIN_TREE_HEIGHT; ++h) {
            // First, determine X-axis of the element's sibling
            uint256 siblingX = (index & 1 == 0) ? index + 1 : index - 1;
            // Get sibling state at the time when `nonce` leafs were added
            proof[h] = _fetchSavedTreeElement(h, siblingX, count);
            // Traverse to parent
            index = index >> 1;
        }
        return cutProof(proof);
    }

    /**
     * @notice Return tree element with height `h`, x-coordinate `x`, after
     * `count` leafs have been inserted. O(1)
     */
    function _fetchSavedTreeElement(uint256 h, uint256 x, uint256 count) internal view returns (bytes32 savedValue) {
        // Should be probably named greatgreat...grandchild, as this
        // references the children in the very bottom (leafs)
        uint256 firstChildLeafIndex = x << h; // x * (2**H)
        uint256 childLeafsAmount = 1 << h; // 2**H
        if (count <= firstChildLeafIndex) {
            // Stage A: not enough leafs were inserted, element is still zero
            savedValue = bytes32(0);
        } else if (count <= firstChildLeafIndex + childLeafsAmount) {
            // Stage B: tree element was updated after last leaf insertion
            savedValue = merkleTree[h][x][count];
            // Sanity check, can't be zero at this point
            require(savedValue != bytes32(0), "Stage B");
        } else {
            // Stage C: tree element was not updated after last leaf insertion
            // Use last saved value
            savedValue = merkleTree[h][x][firstChildLeafIndex + childLeafsAmount];
            // Sanity check, can't be zero at this point
            require(savedValue != bytes32(0), "Stage C");
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
