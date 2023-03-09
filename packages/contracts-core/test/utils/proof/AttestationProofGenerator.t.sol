// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import { StateLib } from "../../../contracts/libs/State.sol";

contract AttestationProofGenerator {
    using StateLib for bytes;

    uint256 internal leafs;
    bytes[] internal states;

    uint256 public height;
    bytes32 public root;

    // (height => (index => node))
    mapping(uint256 => mapping(uint256 => bytes32)) internal nodes;

    function acceptSnapshot(bytes[] memory snapshotStates) external {
        delete states;
        uint256 amount = 1;
        while (amount < snapshotStates.length) {
            amount <<= 1;
        }
        uint256 h = 0;
        // Copy leafs
        for (uint256 i = 0; i < snapshotStates.length; ++i) {
            states.push(snapshotStates[i]);
            nodes[h][i] = snapshotStates[i].castToState().hash();
        }
        // Remaining leafs are ZERO
        for (uint256 i = snapshotStates.length; i < amount; ++i) {
            nodes[h][i] = bytes32(0);
        }
        leafs = amount;
        // Save all parents all the way to the root
        while (amount > 1) {
            ++h;
            amount >>= 1;
            for (uint256 i = 0; i < amount; ++i) {
                nodes[h][i] = keccak256(bytes.concat(nodes[h - 1][2 * i], nodes[h - 1][2 * i + 1]));
            }
        }
        root = nodes[h][0];
        height = h;
    }

    function generateProof(uint256 index) external view returns (bytes32[] memory proof) {
        require(index < states.length, "Out of range");
        proof = new bytes32[](height + 1);
        // First element is "right sub-leaf"
        (, proof[0]) = states[index].castToState().subLeafs();
        for (uint256 h = 0; h < height; ++h) {
            // Get sibling on the current level
            proof[h + 1] = nodes[h][index ^ 1];
            // Traverse to parent
            index >>= 1;
        }
    }
}
