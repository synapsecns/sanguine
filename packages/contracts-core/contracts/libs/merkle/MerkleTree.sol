// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {MerkleMath} from "./MerkleMath.sol";
import {AGENT_TREE_HEIGHT, ORIGIN_TREE_HEIGHT} from "../Constants.sol";
import {NotEnoughLeafs, MerkleTreeFull, LeafNotProven} from "../Errors.sol";

/// `BaseTree` is a struct representing incremental merkle tree.
/// - Contains only the current branch.
/// - The number of inserted leaves is stored externally, and is later supplied for tree operations.
/// - Has a fixed height of `ORIGIN_TREE_HEIGHT`.
/// - Can store up to `2 ** ORIGIN_TREE_HEIGHT - 1` leaves.
/// > Note: the hash function for the tree `H(x, y)` is defined as:
/// > - `H(0,0) = 0`
/// > - `H(x,y) = keccak256(x, y), when (x != 0 || y != 0)`
/// ## Invariant for leafs
/// - The leftmost empty leaf has `index == count`, where `count` is the amount of the inserted leafs so far.
/// - Value for any empty leaf or node is bytes32(0).
/// ## Invariant for current branch
/// `branch[i]` is always the value of a node on the i-th level.
/// Levels are numbered from leafs to root: `0 .. ORIGIN_TREE_HEIGHT`.
/// `branch[i]` stores the value for the node, such that:
//  - The node is a "left child" (e.g. has an even index).
/// - The node must have two non-empty children.
/// - Out of all level's "left child" nodes with "non-empty children",
/// the one with the biggest index (the rightmost one) is stored as `branch[i]`.
/// > __`branch` could be used to form a proof of inclusion for the first empty leaf (`index == count`).__
/// _Here is how:_
/// - Let's walk along the path from the "first empty leaf" to the root.
/// - i-th bit of the "first empty leaf" index (which is equal to `count`) determines if the path's node
/// for this i-th level is a "left child" or a "right child".
/// - i-th bit in `count` is 0 → we are the left child on this level → sibling is the right child
/// that does not exist yet → `proof[i] = bytes32(0)`.
/// - i-th bit in `count` is 1 → we are the right child on this level → sibling is the left child
/// sibling is the rightmost "left child" node on the level → `proof[i] = branch[i]`.
/// > Therefore `proof[i] = (count & (1 << i)) == 0 ? bytes32(0) : branch[i])`
struct BaseTree {
    bytes32[ORIGIN_TREE_HEIGHT] branch;
}

using MerkleTree for BaseTree global;

/// `HistoricalTree` is an incremental merkle tree keeping track of its historical merkle roots.
/// > - `roots[N]` is the root of the tree after `N` leafs were inserted
/// > - `roots[0] == bytes32(0)`
/// @param tree     Incremental merkle tree
/// @param roots    Historical merkle roots of the tree
struct HistoricalTree {
    BaseTree tree;
    bytes32[] roots;
}

using MerkleTree for HistoricalTree global;

/// `DynamicTree` is a struct representing a Merkle Tree with `2**AGENT_TREE_HEIGHT` leaves.
/// - A single operation is available: update value for leaf with an arbitrary index (which might be a non-empty leaf).
/// - This is done by requesting the proof of inclusion for the old value, which is used to both
/// verify the old value, and calculate the new root.
/// > Based on Original idea from [ER forum post](https://ethresear.ch/t/efficient-on-chain-dynamic-merkle-tree/11054).
struct DynamicTree {
    bytes32 root;
}

using MerkleTree for DynamicTree global;

/// MerkleTree is work based on Nomad's Merkle.sol, which is used under MIT OR Apache-2.0
/// [link](https://github.com/nomad-xyz/monorepo/blob/main/packages/contracts-core/contracts/libs/Merkle.sol).
/// With the following changes:
/// - Adapted for Solidity 0.8.x.
/// - Amount of tree leaves stored externally.
/// - Added thorough documentation.
/// - `H(0,0) = 0` optimization from [ER forum post](https://ethresear.ch/t/optimizing-sparse-merkle-trees/3751/6).
/// > Nomad's Merkle.sol is work based on eth2 deposit contract, which is used under CC0-1.0
///[link](https://github.com/ethereum/deposit_contract/blob/dev/deposit_contract/contracts/validator_registration.v.py).
/// With the following changes:
/// > - Implemented in Solidity 0.7.6 (eth2 deposit contract implemented in Vyper).
/// > - `H() = keccak256()` is used as the hashing function instead of `sha256()`.
library MerkleTree {
    /// @dev For root calculation we need at least one empty leaf, thus the minus one in the formula.
    uint256 internal constant MAX_LEAVES = 2 ** ORIGIN_TREE_HEIGHT - 1;

    // ═════════════════════════════════════════════════ BASE TREE ═════════════════════════════════════════════════════

    /**
     * @notice Inserts `node` into merkle tree
     * @dev Reverts if tree is full
     * @param newCount  Amount of inserted leaves in the tree after the insertion (i.e. current + 1)
     * @param node      Element to insert into tree
     */
    function insertBase(BaseTree storage tree, uint256 newCount, bytes32 node) internal {
        if (newCount > MAX_LEAVES) revert MerkleTreeFull();
        // We go up the tree following the branch from the empty leaf AFTER the just inserted one.
        // We stop when we find the first "right child" node.
        // Its sibling is now the rightmost "left child" node that has two non-empty children.
        // Therefore we need to update `tree.branch` value on this level.
        // One could see that `tree.branch` value on lower and higher levels remain unchanged.

        // Loop invariant: `node` is the current level's value for the branch from JUST INSERTED leaf
        for (uint256 i = 0; i < ORIGIN_TREE_HEIGHT;) {
            if ((newCount & 1) == 1) {
                // Found the first "right child" node on the branch from EMPTY leaf
                // `node` is the value for node on branch from JUST INSERTED leaf
                // Which in this case is the "left child".
                // We update tree.branch and exit
                tree.branch[i] = node;
                return;
            }
            // On the branch from EMPTY leaf this is still a "left child".
            // Meaning on branch from JUST INSERTED leaf, `node` is a right child.
            // We compute value for `node` parent using `tree.branch` invariant:
            // This is the rightmost "left child" node, which would be sibling of `node`.
            node = MerkleMath.getParent(tree.branch[i], node);
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
        // To calculate the root we follow the branch of first EMPTY leaf (index == count)
        for (uint256 i = 0; i < ORIGIN_TREE_HEIGHT;) {
            // Check if we are the left or the right child on the current level
            if ((count & 1) == 1) {
                // We are the right child. Our sibling is the "rightmost" "left-child" node
                // that has two non-empty children → sibling is `tree.branch[i]`
                current = MerkleMath.getParent(tree.branch[i], current);
            } else {
                // We are the left child. Our sibling does not exist yet → sibling is EMPTY
                current = MerkleMath.getParent(current, bytes32(0));
            }
            // Get the parent index, and go to the next tree level
            count >>= 1;
            unchecked {
                ++i;
            }
        }
    }

    // ══════════════════════════════════════════════ HISTORICAL TREE ══════════════════════════════════════════════════

    /// @notice Initializes the historical roots for the tree by inserting an empty root.
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
        if (count >= tree.roots.length) revert NotEnoughLeafs();
        return tree.roots[count];
    }

    // ═══════════════════════════════════════════════ DYNAMIC TREE ════════════════════════════════════════════════════

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
        if (MerkleMath.proofRoot(index, oldValue, branch, AGENT_TREE_HEIGHT) != tree.root) {
            revert LeafNotProven();
        }
        // New root is new value + the same proof (values for sibling nodes are not updated)
        newRoot = MerkleMath.proofRoot(index, newValue, branch, AGENT_TREE_HEIGHT);
        // Write the new root
        tree.root = newRoot;
    }
}
