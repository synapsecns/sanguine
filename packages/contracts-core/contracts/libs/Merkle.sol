// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import { TREE_DEPTH } from "./Constants.sol";

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
///     - Levels are numbered from 0 (leafs) to TREE_DEPTH (root)
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
    bytes32[TREE_DEPTH] branch;
}
using { MerkleLib.insertBase, MerkleLib.rootBase } for BaseTree global;

/// @notice Incremental merkle tree keeping track of its historical merkle roots.
/// @dev roots[N] is the root of the tree after N leafs were inserted
/// @param tree     Incremental merkle tree
/// @param roots    Historical merkle roots of the tree
struct HistoricalTree {
    BaseTree tree;
    bytes32[] roots;
}
using { MerkleLib.initializeRoots, MerkleLib.insert, MerkleLib.root } for HistoricalTree global;

/// @notice Struct representing a Dynamic Merkle Tree with 2**TREE_DEPTH leaves
/// A single operation is available: update value for existing leaf (which might be ZERO).
/// This is done by requesting the proof of inclusion for the old value, which is used to
/// verify the old value, and calculate the new root.
/// Based on Original idea from https://ethresear.ch/t/efficient-on-chain-dynamic-merkle-tree/11054
struct DynamicTree {
    bytes32 root;
}
using { MerkleLib.update } for DynamicTree global;

library MerkleLib {
    uint256 internal constant MAX_LEAVES = 2**TREE_DEPTH - 1;

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                              BASE TREE                               ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @notice Inserts `_node` into merkle tree
     * @dev Reverts if tree is full
     * @param _newCount Amount of inserted leaves in the tree after the insertion (i.e. current + 1)
     * @param _node     Element to insert into tree
     **/
    function insertBase(
        BaseTree storage _tree,
        uint256 _newCount,
        bytes32 _node
    ) internal {
        require(_newCount <= MAX_LEAVES, "merkle tree full");
        // We go up the tree following the branch from the zero leaf AFTER the just inserted one.
        // We stop when we find the first "right child" node.
        // Its sibling is now the rightmost "left child" node that has both children as non-zero.
        // Therefore we need to update `tree.branch` value on this level.
        // One could see that `tree.branch` value on lower and higher levels remain unchanged.

        // Loop invariant: `node` is the current level's value for the branch from JUST INSERTED leaf
        for (uint256 i = 0; i < TREE_DEPTH; ) {
            if ((_newCount & 1) == 1) {
                // Found the first "right child" node on the branch from ZERO leaf
                // `node` is the value for node on branch from JUST INSERTED leaf
                // Which in this case is the "left child".
                // We update tree.branch and exit
                _tree.branch[i] = _node;
                return;
            }
            // On the branch from ZERO leaf this is still "left child".
            // Meaning on branch from JUST INSERTED leaf, `node` is right child
            // We compute value for `node` parent using `tree.branch` invariant:
            // This is the rightmost "left child" node, which would be sibling of `node`
            _node = getParent(_tree.branch[i], _node);
            // Get the parent index, and go to the next tree level
            _newCount >>= 1;
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
     * @param _count    Current amount of inserted leaves in the tree
     * @return _current Calculated root of `_tree`
     **/
    function rootBase(BaseTree storage _tree, uint256 _count)
        internal
        view
        returns (bytes32 _current)
    {
        // To calculate the root we follow the branch of first ZERO leaf (index == count)
        for (uint256 i = 0; i < TREE_DEPTH; ) {
            // Check if we are the left or the right child on the current level
            if ((_count & 1) == 1) {
                // We are the right child. Our sibling is the "rightmost" "left-child" node
                // that has two non-zero children => sibling is tree.branch[i]
                _current = getParent(_tree.branch[i], _current);
            } else {
                // We are the left child. Our sibling does not exist yet => sibling is ZERO
                _current = getParent(_current, bytes32(0));
            }
            // Get the parent index, and go to the next tree level
            _count >>= 1;
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
    function initializeRoots(HistoricalTree storage _tree) internal returns (bytes32 savedRoot) {
        // This should only be called once, when the contract is initialized
        assert(_tree.roots.length == 0);
        // Save root for empty merkle tree: bytes32(0)
        _tree.roots.push(savedRoot);
    }

    /// @notice Inserts a new leaf into the merkle tree.
    /// @dev Reverts if tree is full.
    /// @param _node        Element to insert into tree
    /// @return newRoot     Merkle root after the leaf was inserted
    function insert(HistoricalTree storage _tree, bytes32 _node)
        internal
        returns (bytes32 newRoot)
    {
        // Tree count after the new leaf will be inserted (we store roots[0] as root of empty tree)
        uint256 newCount = _tree.roots.length;
        _tree.tree.insertBase(newCount, _node);
        // Save the new root
        newRoot = _tree.tree.rootBase(newCount);
        _tree.roots.push(newRoot);
    }

    /// @notice Returns the historical root of the merkle tree.
    /// @dev Reverts if not enough leafs have been inserted.
    /// @param _count           Amount of leafs in the tree at some point of time
    /// @return historicalRoot  Merkle root after `_count` leafs were inserted
    function root(HistoricalTree storage _tree, uint256 _count)
        internal
        view
        returns (bytes32 historicalRoot)
    {
        require(_count < _tree.roots.length, "Not enough leafs inserted");
        return _tree.roots[_count];
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                             DYNAMIC TREE                             ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @notice Updates the value for the leaf with the given index in the Dynamic Merkle Tree.
     * @dev Will revert if incorrect proof of inclusion for old value is supplied.
     * @param _tree         Dynamic merkle tree
     * @param _index        Index of the leaf to update
     * @param _oldValue     Previous value of the leaf
     * @param _branch       Proof of inclusion of previous value into the tree
     * @param _newValue     New leaf value to assign
     * @return newRoot      New value for the Merkle Root after the leaf is updated
     */
    function update(
        DynamicTree storage _tree,
        uint256 _index,
        bytes32 _oldValue,
        bytes32[TREE_DEPTH] memory _branch,
        bytes32 _newValue
    ) internal returns (bytes32 newRoot) {
        // Check that the old value + proof result in a correct root
        require(branchRoot(_oldValue, _branch, _index) == _tree.root, "Incorrect proof");
        // New root is new value + the same proof (values for sibling nodes are not updated)
        newRoot = branchRoot(_newValue, _branch, _index);
        // Write the new root
        _tree.root = newRoot;
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                               HELPERS                                ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @notice Calculates and returns the merkle root for the given leaf
     * `_item`, a merkle branch, and the index of `_item` in the tree.
     * @param _item Merkle leaf
     * @param _branch Merkle proof
     * @param _index Index of `_item` in tree
     * @return _current Calculated merkle root
     **/
    function branchRoot(
        bytes32 _item,
        bytes32[TREE_DEPTH] memory _branch,
        uint256 _index
    ) internal pure returns (bytes32 _current) {
        _current = _item;
        // Go up the tree levels from leaf to root
        for (uint256 i = 0; i < TREE_DEPTH; ) {
            _current = getParent(_current, _branch[i], _index, i);
            unchecked {
                ++i;
            }
        }
    }

    /**
     * @notice Calculates and returns the merkle root for the given leaf
     * `_item`, a merkle branch, and the index of `_item` in the tree.
     * @dev This extra function is required for Merkle Trees with flexible height.
     * @param _item Merkle leaf
     * @param _branch Merkle proof
     * @param _index Index of `_item` in tree
     * @return _current Calculated merkle root
     **/
    function branchRoot(
        bytes32 _item,
        bytes32[] memory _branch,
        uint256 _index
    ) internal pure returns (bytes32 _current) {
        _current = _item;
        // Go up the tree levels from leaf to root
        for (uint256 i = 0; i < _branch.length; ) {
            _current = getParent(_current, _branch[i], _index, i);
            unchecked {
                ++i;
            }
        }
    }

    /**
     * @notice Calculates the parent of a node on the path from one of the leafs to root.
     * @param _node         Node on a path from tree leaf to root
     * @param _sibling      Sibling for a given node
     * @param _leafIndex    Index of the tree leaf
     * @param _nodeHeight   "Level height" for `_node` (ZERO for leafs, TREE_DEPTH for root)
     */
    function getParent(
        bytes32 _node,
        bytes32 _sibling,
        uint256 _leafIndex,
        uint256 _nodeHeight
    ) internal pure returns (bytes32 parent) {
        // Index for `node` on its "tree level" is (leafIndex / 2**height)
        // "Left child" has even index, "right child" has odd index
        if ((_leafIndex >> _nodeHeight) & 1 == 0) {
            // Left child
            return getParent(_node, _sibling);
        } else {
            // Right child
            return getParent(_sibling, _node);
        }
    }

    /// @notice Calculates the parent of tow nodes in the merkle tree.
    /// @dev We use implementation with H(0,0) = 0
    /// This makes EVERY empty node in the tree equal to ZERO,
    /// saving us from storing H(0,0), H(H(0,0), H(0, 0)), and so on
    /// @param _leftChild   Left child of the calculated node
    /// @param _rightChild  Right child of the calculated node
    /// @return parent      Value for the node having above mentioned children
    function getParent(bytes32 _leftChild, bytes32 _rightChild)
        internal
        pure
        returns (bytes32 parent)
    {
        if (_leftChild == bytes32(0) && _rightChild == bytes32(0)) {
            return 0;
        } else {
            return keccak256(bytes.concat(_leftChild, _rightChild));
        }
    }
}
