// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {AGENT_TREE_HEIGHT, ORIGIN_TREE_HEIGHT} from "./Constants.sol";

// work based on Merkle.sol, which is used under MIT OR Apache-2.0:
// https://github.com/nomad-xyz/monorepo/blob/main/packages/contracts-core/contracts/libs/Merkle.sol
// Changes:
//  - Adapted for Solidity 0.8.x
//  - Amount of tree leaves stored externally
//  - Added thorough documentation
//  - H(0,0) = 0 optimization is implemented (https://ethresear.ch/t/optimizing-sparse-merkle-trees/3751/6)

// Nomad's Merkle.sol is work based on eth2 deposit contract, which is used under CC0-1.0:
// https://github.com/ethereum/deposit_contract/blob/dev/deposit_contract/contracts/validator_registration.v.py
// Changes:
//  - Implemented in Solidity 0.7.6 (eth2 impl is Vyper)
//  - H() = keccak256() is used as the hashing function instead of sha256()

/// @notice Struct representing incremental merkle tree. Contains the current branch, while
/// the number of inserted leaves are stored externally, and is later supplied for tree operation.
/// Note: the hash function for the tree H(x, y) is defined as:
/// - H(0,0) = 0
/// - H(x,y) = keccak256(x, y), if x != 0 or y != 0
/// @dev Following invariant is enforced:
/// - First empty leaf has index `count`, where `count` is the amount of the inserted leafs so far
/// - Value for the empty leaf is zeroes[0] = bytes32(0)
/// - Value for node having empty children zeroes[i] = H(zeroes[i-1], zeroes[i-1])
/// - branch[i] is the value of a node on the i-th level:
///     - Levels are numbered from 0 (leafs) to ORIGIN_TREE_HEIGHT (root)
///     - branch[i] stores the value for the node, that is a "left child"
///     - The stored node must have non-zero values for both their children
///     - Out of all level's "left child" nodes with "non-zero children",
///       the one with the biggest index (the rightmost one) is stored.
/// - Therefore, proof of inclusion for the first ZERO leaf (`index == count`) is:
///     - i-th bit in `count` is 0 => we are the left child on this level => sibling is the right child
///       sibling does not exist yet
///         - Therefore proof[i] = zeroes[i]
///     - i-th bit in `count` is 1 => we are the right child on this level => sibling is the left child
///       sibling is the rightmost "left child" node on the level
///         - Therefore proof[i] = branch[i]
struct BaseTree {
    bytes32[ORIGIN_TREE_HEIGHT] branch;
}

using {MerkleLib.insertBase, MerkleLib.rootBase} for BaseTree global;

/// @notice Incremental merkle tree keeping track of its historical merkle roots.
/// @dev roots[N] is the root of the tree after N leafs were inserted
/// @param tree     Incremental merkle tree
/// @param roots    Historical merkle roots of the tree
struct HistoricalTree {
    BaseTree tree;
    bytes32[] roots;
}

using {MerkleLib.initializeRoots, MerkleLib.insert, MerkleLib.root} for HistoricalTree global;

/// @notice Struct representing a Dynamic Merkle Tree with 2**AGENT_TREE_HEIGHT leaves
/// A single operation is available: update value for existing leaf (which might be ZERO).
/// This is done by requesting the proof of inclusion for the old value, which is used to
/// verify the old value, and calculate the new root.
/// Based on Original idea from https://ethresear.ch/t/efficient-on-chain-dynamic-merkle-tree/11054
struct DynamicTree {
    bytes32 root;
}

using {MerkleLib.update} for DynamicTree global;

library MerkleLib {
    uint256 internal constant MAX_LEAVES = 2 ** ORIGIN_TREE_HEIGHT - 1;

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                              BASE TREE                               ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @notice Inserts `node` into merkle tree
     * @dev Reverts if tree is full
     * @param newCount  Amount of inserted leaves in the tree after the insertion (i.e. current + 1)
     * @param node      Element to insert into tree
     */
    function insertBase(BaseTree storage tree, uint256 newCount, bytes32 node) internal {
        require(newCount <= MAX_LEAVES, "merkle tree full");
        // We go up the tree following the branch from the zero leaf AFTER the just inserted one.
        // We stop when we find the first "right child" node.
        // Its sibling is now the rightmost "left child" node that has both children as non-zero.
        // Therefore we need to update `tree.branch` value on this level.
        // One could see that `tree.branch` value on lower and higher levels remain unchanged.

        // Loop invariant: `node` is the current level's value for the branch from JUST INSERTED leaf
        for (uint256 i = 0; i < ORIGIN_TREE_HEIGHT;) {
            if ((newCount & 1) == 1) {
                // Found the first "right child" node on the branch from ZERO leaf
                // `node` is the value for node on branch from JUST INSERTED leaf
                // Which in this case is the "left child".
                // We update tree.branch and exit
                tree.branch[i] = node;
                return;
            }
            // On the branch from ZERO leaf this is still "left child".
            // Meaning on branch from JUST INSERTED leaf, `node` is right child
            // We compute value for `node` parent using `tree.branch` invariant:
            // This is the rightmost "left child" node, which would be sibling of `node`
            node = getParent(tree.branch[i], node);
            // Get the parent index, and go to the next tree level
            newCount >>= 1;
            unchecked {
                ++i;
            }
        }
        // As the loop should always end prematurely with the `return` statement,
        // this code should be unreachable. We assert `false` just to be safe.
        assert(false);
    }

    /**
     * @notice Calculates and returns current root of the merkle tree.
     * @param count     Current amount of inserted leaves in the tree
     * @return current  Calculated root of `tree`
     */
    function rootBase(BaseTree storage tree, uint256 count) internal view returns (bytes32 current) {
        // To calculate the root we follow the branch of first ZERO leaf (index == count)
        for (uint256 i = 0; i < ORIGIN_TREE_HEIGHT;) {
            // Check if we are the left or the right child on the current level
            if ((count & 1) == 1) {
                // We are the right child. Our sibling is the "rightmost" "left-child" node
                // that has two non-zero children => sibling is tree.branch[i]
                current = getParent(tree.branch[i], current);
            } else {
                // We are the left child. Our sibling does not exist yet => sibling is ZERO
                current = getParent(current, bytes32(0));
            }
            // Get the parent index, and go to the next tree level
            count >>= 1;
            unchecked {
                ++i;
            }
        }
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                           HISTORICAL TREE                            ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /// @notice Initializes the historical roots for the tree by inserting
    /// a precomputed root of an empty Merkle Tree.
    // solhint-disable-next-line ordering
    function initializeRoots(HistoricalTree storage tree) internal returns (bytes32 savedRoot) {
        // This should only be called once, when the contract is initialized
        assert(tree.roots.length == 0);
        // Save root for empty merkle tree: bytes32(0)
        tree.roots.push(savedRoot);
    }

    /// @notice Inserts a new leaf into the merkle tree.
    /// @dev Reverts if tree is full.
    /// @param node         Element to insert into tree
    /// @return newRoot     Merkle root after the leaf was inserted
    function insert(HistoricalTree storage tree, bytes32 node) internal returns (bytes32 newRoot) {
        // Tree count after the new leaf will be inserted (we store roots[0] as root of empty tree)
        uint256 newCount = tree.roots.length;
        tree.tree.insertBase(newCount, node);
        // Save the new root
        newRoot = tree.tree.rootBase(newCount);
        tree.roots.push(newRoot);
    }

    /// @notice Returns the historical root of the merkle tree.
    /// @dev Reverts if not enough leafs have been inserted.
    /// @param count            Amount of leafs in the tree at some point of time
    /// @return historicalRoot  Merkle root after `count` leafs were inserted
    function root(HistoricalTree storage tree, uint256 count) internal view returns (bytes32 historicalRoot) {
        require(count < tree.roots.length, "Not enough leafs inserted");
        return tree.roots[count];
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                             DYNAMIC TREE                             ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @notice Updates the value for the leaf with the given index in the Dynamic Merkle Tree.
     * @dev Will revert if incorrect proof of inclusion for old value is supplied.
     * @param tree          Dynamic merkle tree
     * @param index         Index of the leaf to update
     * @param oldValue      Previous value of the leaf
     * @param branch        Proof of inclusion of previous value into the tree
     * @param newValue      New leaf value to assign
     * @return newRoot      New value for the Merkle Root after the leaf is updated
     */
    function update(
        DynamicTree storage tree,
        uint256 index,
        bytes32 oldValue,
        bytes32[] memory branch,
        bytes32 newValue
    ) internal returns (bytes32 newRoot) {
        // Check that the old value + proof result in a correct root
        require(proofRoot(index, oldValue, branch, AGENT_TREE_HEIGHT) == tree.root, "Incorrect proof");
        // New root is new value + the same proof (values for sibling nodes are not updated)
        newRoot = proofRoot(index, newValue, branch, AGENT_TREE_HEIGHT);
        // Write the new root
        tree.root = newRoot;
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                               HELPERS                                ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @notice Calculates the merkle root for the given leaf and merkle proof.
     * @dev Will revert if proof length exceeds the tree height.
     * @param index     Index of `leaf` in tree
     * @param leaf      Leaf of the merkle tree
     * @param proof     Proof of inclusion of `leaf` in the tree
     * @param height    Height of the merkle tree
     * @return root_    Calculated Merkle Root
     */
    function proofRoot(uint256 index, bytes32 leaf, bytes32[] memory proof, uint256 height)
        internal
        pure
        returns (bytes32 root_)
    {
        // Proof length could not exceed the tree height
        uint256 proofLen = proof.length;
        require(proofLen <= height, "Proof too long");
        root_ = leaf;
        // Go up the tree levels from the leaf following the proof
        for (uint256 h = 0; h < proofLen; ++h) {
            // Get a sibling node on current level: this is proof[h]
            root_ = getParent(root_, proof[h], index, h);
        }
        // Go up to the root: the remaining siblings are ZERO
        for (uint256 h = proofLen; h < height; ++h) {
            root_ = getParent(root_, bytes32(0), index, h);
        }
    }

    /**
     * @notice Calculates the parent of a node on the path from one of the leafs to root.
     * @param node          Node on a path from tree leaf to root
     * @param sibling       Sibling for a given node
     * @param leafIndex     Index of the tree leaf
     * @param nodeHeight    "Level height" for `node` (ZERO for leafs, ORIGIN_TREE_HEIGHT for root)
     */
    function getParent(bytes32 node, bytes32 sibling, uint256 leafIndex, uint256 nodeHeight)
        internal
        pure
        returns (bytes32 parent)
    {
        // Index for `node` on its "tree level" is (leafIndex / 2**height)
        // "Left child" has even index, "right child" has odd index
        if ((leafIndex >> nodeHeight) & 1 == 0) {
            // Left child
            return getParent(node, sibling);
        } else {
            // Right child
            return getParent(sibling, node);
        }
    }

    /// @notice Calculates the parent of tow nodes in the merkle tree.
    /// @dev We use implementation with H(0,0) = 0
    /// This makes EVERY empty node in the tree equal to ZERO,
    /// saving us from storing H(0,0), H(H(0,0), H(0, 0)), and so on
    /// @param leftChild    Left child of the calculated node
    /// @param rightChild   Right child of the calculated node
    /// @return parent      Value for the node having above mentioned children
    function getParent(bytes32 leftChild, bytes32 rightChild) internal pure returns (bytes32 parent) {
        if (leftChild == bytes32(0) && rightChild == bytes32(0)) {
            return 0;
        } else {
            return keccak256(bytes.concat(leftChild, rightChild));
        }
    }
}
