// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {SNAPSHOT_TREE_HEIGHT} from "../../../contracts/libs/Constants.sol";
import {MerkleMath} from "../../../contracts/libs/merkle/MerkleMath.sol";
import {StateLib} from "../../../contracts/libs/memory/State.sol";

import {ProofCutter} from "./ProofCutter.t.sol";

// TODO: move from test directory
contract AttestationProofGenerator is ProofCutter {
    using StateLib for bytes;

    bytes32 public root;

    // (height => (index => node))
    mapping(uint256 => mapping(uint256 => bytes32)) internal nodes;

    function acceptSnapshot(bytes[] memory snapshotStates) external {
        uint256 amount = 1 << SNAPSHOT_TREE_HEIGHT;
        uint256 h = 0;
        // Copy leafs
        for (uint256 i = 0; i < snapshotStates.length; ++i) {
            (nodes[h][2 * i], nodes[h][2 * i + 1]) = snapshotStates[i].castToState().subLeafs();
        }
        // Remaining leafs are ZERO
        for (uint256 i = 2 * snapshotStates.length; i < amount; ++i) {
            nodes[h][i] = bytes32(0);
        }
        // Save all parents all the way to the root
        while (amount > 1) {
            ++h;
            amount >>= 1;
            for (uint256 i = 0; i < amount; ++i) {
                nodes[h][i] = MerkleMath.getParent(nodes[h - 1][2 * i], nodes[h - 1][2 * i + 1]);
            }
        }
        root = nodes[h][0];
    }

    function generateProof(uint256 stateIndex) external view returns (bytes32[] memory) {
        // Index of State's left leaf
        uint256 index = stateIndex * 2;
        require(index < (1 << SNAPSHOT_TREE_HEIGHT), "Out of range");
        bytes32[] memory proof = new bytes32[](SNAPSHOT_TREE_HEIGHT);
        for (uint256 h = 0; h < SNAPSHOT_TREE_HEIGHT; ++h) {
            // Get sibling on the current level
            proof[h] = nodes[h][index ^ 1];
            // Traverse to parent
            index >>= 1;
        }
        return cutProof(proof);
    }
}
