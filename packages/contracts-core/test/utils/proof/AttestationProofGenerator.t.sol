// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import { MerkleLib } from "../../../contracts/libs/Merkle.sol";
import { StateLib } from "../../../contracts/libs/State.sol";

// TODO: move from test directory
contract AttestationProofGenerator {
    using StateLib for bytes;

    uint256 public height;
    bytes32 public root;

    // (height => (index => node))
    mapping(uint256 => mapping(uint256 => bytes32)) internal nodes;

    function acceptSnapshot(bytes[] memory snapshotStates) external {
        uint256 amount = 1;
        while (amount < 2 * snapshotStates.length) {
            amount <<= 1;
        }
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
                nodes[h][i] = MerkleLib.getParent(nodes[h - 1][2 * i], nodes[h - 1][2 * i + 1]);
            }
        }
        root = nodes[h][0];
        height = h;
    }

    function generateProof(uint256 stateIndex) external view returns (bytes32[] memory proof) {
        // Index of State's left leaf
        uint256 index = stateIndex * 2;
        require(index < (1 << height), "Out of range");
        proof = new bytes32[](height);
        for (uint256 h = 0; h < height; ++h) {
            // Get sibling on the current level
            proof[h] = nodes[h][index ^ 1];
            // Traverse to parent
            index >>= 1;
        }
    }
}
