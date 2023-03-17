// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import { EMPTY_ROOT, TREE_DEPTH } from "./Constants.sol";

// work based on Merkle.sol, which is used under MIT OR Apache-2.0
// Changes:
//  - Adapted for Solidity 0.8.x
//  - Amount of tree leaves stored externally
//  - Added thorough documentation
// https://github.com/nomad-xyz/monorepo/blob/main/packages/contracts-core/contracts/libs/Merkle.sol

// work based on eth2 deposit contract, which is used under CC0-1.0
// Changes:
//  - Implemented in Solidity 0.7.6 (eth2 impl is Vyper)
//  - H() = keccak256() is used as the hashing function instead of sha256()

/// @notice Struct representing incremental merkle tree. Contains the current branch, while
/// the number of inserted leaves are stored externally, and is later supplied for tree operation.
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
            _node = keccak256(abi.encodePacked(_tree.branch[i], _node));
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
     * @notice Calculates and returns`_tree`'s current root given array of zero
     * hashes
     * @param _count    Current amount of inserted leaves in the tree
     * @param _zeroes   Array of zero hashes
     * @return _current Calculated root of `_tree`
     **/
    function rootWithCtx(
        BaseTree storage _tree,
        uint256 _count,
        bytes32[TREE_DEPTH] memory _zeroes
    ) private view returns (bytes32 _current) {
        // To calculate the root we follow the branch of first ZERO leaf (index == count)
        for (uint256 i = 0; i < TREE_DEPTH; ) {
            // i-th bit tells if we are the right of the left child on i-th level
            uint256 _ithBit = (_count >> i) & 0x01;
            if (_ithBit == 1) {
                // We are the right child. Our sibling is the "rightmost" "left-child" node
                // that has two non-zero children: tree.branch[i]
                _current = keccak256(abi.encodePacked(_tree.branch[i], _current));
            } else {
                // We are the left child. Our sibling does not exist yet: zeroes[i]
                _current = keccak256(abi.encodePacked(_current, _zeroes[i]));
            }
            unchecked {
                ++i;
            }
        }
    }

    /**
     * @notice Calculates and returns`_tree`'s current root
     * @param _count    Current amount of inserted leaves in the tree
     * @return Calculated root of `_tree`
     **/
    function rootBase(BaseTree storage _tree, uint256 _count) internal view returns (bytes32) {
        return rootWithCtx(_tree, _count, zeroHashes());
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                           HISTORICAL TREE                            ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /// @notice Initializes the historical roots for the tree by inserting
    /// a precomputed root of an empty Merkle Tree.
    function initializeRoots(HistoricalTree storage _tree) internal returns (bytes32 savedRoot) {
        // This should only be called once, when the contract is initialized
        assert(_tree.roots.length == 0);
        // Save root for empty merkle tree
        savedRoot = EMPTY_ROOT;
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
    ▏*║                               HELPERS                                ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /// @notice Returns array of TREE_DEPTH zero hashes
    /// @return _zeroes Array of TREE_DEPTH zero hashes
    function zeroHashes() internal pure returns (bytes32[TREE_DEPTH] memory _zeroes) {
        _zeroes[0] = Z_0;
        _zeroes[1] = Z_1;
        _zeroes[2] = Z_2;
        _zeroes[3] = Z_3;
        _zeroes[4] = Z_4;
        _zeroes[5] = Z_5;
        _zeroes[6] = Z_6;
        _zeroes[7] = Z_7;
        _zeroes[8] = Z_8;
        _zeroes[9] = Z_9;
        _zeroes[10] = Z_10;
        _zeroes[11] = Z_11;
        _zeroes[12] = Z_12;
        _zeroes[13] = Z_13;
        _zeroes[14] = Z_14;
        _zeroes[15] = Z_15;
        _zeroes[16] = Z_16;
        _zeroes[17] = Z_17;
        _zeroes[18] = Z_18;
        _zeroes[19] = Z_19;
        _zeroes[20] = Z_20;
        _zeroes[21] = Z_21;
        _zeroes[22] = Z_22;
        _zeroes[23] = Z_23;
        _zeroes[24] = Z_24;
        _zeroes[25] = Z_25;
        _zeroes[26] = Z_26;
        _zeroes[27] = Z_27;
        _zeroes[28] = Z_28;
        _zeroes[29] = Z_29;
        _zeroes[30] = Z_30;
        _zeroes[31] = Z_31;
    }

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
            return keccak256(bytes.concat(_node, _sibling));
        } else {
            // Right child
            return keccak256(bytes.concat(_sibling, _node));
        }
    }

    // keccak256 zero hashes
    bytes32 internal constant Z_0 =
        hex"0000000000000000000000000000000000000000000000000000000000000000";
    bytes32 internal constant Z_1 =
        hex"ad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb5";
    bytes32 internal constant Z_2 =
        hex"b4c11951957c6f8f642c4af61cd6b24640fec6dc7fc607ee8206a99e92410d30";
    bytes32 internal constant Z_3 =
        hex"21ddb9a356815c3fac1026b6dec5df3124afbadb485c9ba5a3e3398a04b7ba85";
    bytes32 internal constant Z_4 =
        hex"e58769b32a1beaf1ea27375a44095a0d1fb664ce2dd358e7fcbfb78c26a19344";
    bytes32 internal constant Z_5 =
        hex"0eb01ebfc9ed27500cd4dfc979272d1f0913cc9f66540d7e8005811109e1cf2d";
    bytes32 internal constant Z_6 =
        hex"887c22bd8750d34016ac3c66b5ff102dacdd73f6b014e710b51e8022af9a1968";
    bytes32 internal constant Z_7 =
        hex"ffd70157e48063fc33c97a050f7f640233bf646cc98d9524c6b92bcf3ab56f83";
    bytes32 internal constant Z_8 =
        hex"9867cc5f7f196b93bae1e27e6320742445d290f2263827498b54fec539f756af";
    bytes32 internal constant Z_9 =
        hex"cefad4e508c098b9a7e1d8feb19955fb02ba9675585078710969d3440f5054e0";
    bytes32 internal constant Z_10 =
        hex"f9dc3e7fe016e050eff260334f18a5d4fe391d82092319f5964f2e2eb7c1c3a5";
    bytes32 internal constant Z_11 =
        hex"f8b13a49e282f609c317a833fb8d976d11517c571d1221a265d25af778ecf892";
    bytes32 internal constant Z_12 =
        hex"3490c6ceeb450aecdc82e28293031d10c7d73bf85e57bf041a97360aa2c5d99c";
    bytes32 internal constant Z_13 =
        hex"c1df82d9c4b87413eae2ef048f94b4d3554cea73d92b0f7af96e0271c691e2bb";
    bytes32 internal constant Z_14 =
        hex"5c67add7c6caf302256adedf7ab114da0acfe870d449a3a489f781d659e8becc";
    bytes32 internal constant Z_15 =
        hex"da7bce9f4e8618b6bd2f4132ce798cdc7a60e7e1460a7299e3c6342a579626d2";
    bytes32 internal constant Z_16 =
        hex"2733e50f526ec2fa19a22b31e8ed50f23cd1fdf94c9154ed3a7609a2f1ff981f";
    bytes32 internal constant Z_17 =
        hex"e1d3b5c807b281e4683cc6d6315cf95b9ade8641defcb32372f1c126e398ef7a";
    bytes32 internal constant Z_18 =
        hex"5a2dce0a8a7f68bb74560f8f71837c2c2ebbcbf7fffb42ae1896f13f7c7479a0";
    bytes32 internal constant Z_19 =
        hex"b46a28b6f55540f89444f63de0378e3d121be09e06cc9ded1c20e65876d36aa0";
    bytes32 internal constant Z_20 =
        hex"c65e9645644786b620e2dd2ad648ddfcbf4a7e5b1a3a4ecfe7f64667a3f0b7e2";
    bytes32 internal constant Z_21 =
        hex"f4418588ed35a2458cffeb39b93d26f18d2ab13bdce6aee58e7b99359ec2dfd9";
    bytes32 internal constant Z_22 =
        hex"5a9c16dc00d6ef18b7933a6f8dc65ccb55667138776f7dea101070dc8796e377";
    bytes32 internal constant Z_23 =
        hex"4df84f40ae0c8229d0d6069e5c8f39a7c299677a09d367fc7b05e3bc380ee652";
    bytes32 internal constant Z_24 =
        hex"cdc72595f74c7b1043d0e1ffbab734648c838dfb0527d971b602bc216c9619ef";
    bytes32 internal constant Z_25 =
        hex"0abf5ac974a1ed57f4050aa510dd9c74f508277b39d7973bb2dfccc5eeb0618d";
    bytes32 internal constant Z_26 =
        hex"b8cd74046ff337f0a7bf2c8e03e10f642c1886798d71806ab1e888d9e5ee87d0";
    bytes32 internal constant Z_27 =
        hex"838c5655cb21c6cb83313b5a631175dff4963772cce9108188b34ac87c81c41e";
    bytes32 internal constant Z_28 =
        hex"662ee4dd2dd7b2bc707961b1e646c4047669dcb6584f0d8d770daf5d7e7deb2e";
    bytes32 internal constant Z_29 =
        hex"388ab20e2573d171a88108e79d820e98f26c0b84aa8b2f4aa4968dbb818ea322";
    bytes32 internal constant Z_30 =
        hex"93237c50ba75ee485f4c22adf2f741400bdf8d6a9cc7df7ecae576221665d735";
    bytes32 internal constant Z_31 =
        hex"8448818bb4ae4562849e949e17ac16e0be16688e156b5cf15e098c627c0056a9";
}
