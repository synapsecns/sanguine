// SPDX-License-Identifier: MIT

pragma solidity 0.8.17;

import {MerkleMath} from "../../../../contracts/libs/merkle/MerkleMath.sol";

contract MerkleMathHarness {
    // Note: we don't add an empty test() function here, as it currently leads
    // to zero coverage on the corresponding library.

    // ═════════════════════════════════════════ BASIC MERKLE CALCULATIONS ═════════════════════════════════════════════

    function proofRoot(uint256 index, bytes32 leaf, bytes32[] memory proof, uint256 height)
        public
        pure
        returns (bytes32)
    {
        return MerkleMath.proofRoot(index, leaf, proof, height);
    }

    function getParent(bytes32 node, bytes32 sibling, uint256 leafIndex, uint256 nodeHeight)
        public
        pure
        returns (bytes32)
    {
        return MerkleMath.getParent(node, sibling, leafIndex, nodeHeight);
    }

    function getParent(bytes32 leftChild, bytes32 rightChild) public pure returns (bytes32) {
        return MerkleMath.getParent(leftChild, rightChild);
    }

    // ════════════════════════════════ ROOT/PROOF CALCULATION FOR A LIST OF LEAFS ═════════════════════════════════════

    function calculateRoot(bytes32[] memory hashes, uint256 height) public pure returns (bytes32) {
        MerkleMath.calculateRoot(hashes, height);
        return hashes[0];
    }

    function calculateProof(bytes32[] memory hashes, uint256 index) public pure returns (bytes32[] memory proof) {
        return MerkleMath.calculateProof(hashes, index);
    }
}
