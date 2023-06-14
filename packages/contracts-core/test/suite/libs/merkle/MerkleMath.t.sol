// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {MerkleMath} from "../../../../contracts/libs/merkle/MerkleMath.sol";

import {SynapseLibraryTest} from "../../../utils/SynapseLibraryTest.t.sol";
import {MerkleMathHarness} from "../../../harnesses/libs/merkle/MerkleMathHarness.t.sol";

// solhint-disable func-name-mixedcase
contract MerkleMathLibraryTest is SynapseLibraryTest {
    uint256 public constant HEIGHT = 8;
    uint256 public constant MAX_LENGTH = 1 << HEIGHT;

    MerkleMathHarness public libHarness;

    function setUp() public {
        libHarness = new MerkleMathHarness();
    }

    function test_calculateRoot(uint256 length) public {
        // length should be in [1 .. MAX_LENGTH] range
        length = bound(length, 1, MAX_LENGTH);
        bytes32[] memory hashes = generateHashes(length);
        bytes32[] memory extended = extendHashes(hashes);
        bytes32 expectedRoot = calculateRoot(extended);
        bytes32 root = libHarness.calculateRoot(hashes, HEIGHT);
        assertEq(root, expectedRoot, "Merkle Root incorrect");
    }

    function test_calculateProof(uint256 length, uint256 index) public {
        // length should be in [1 .. MAX_LENGTH] range
        length = bound(length, 1, MAX_LENGTH);
        bytes32[] memory hashes = generateHashes(length);
        // index should be in [0 .. MAX_LENGTH) range
        index = bound(index, 0, MAX_LENGTH - 1);
        // Check proofs for zero leafs outside of the list as well
        bytes32 node = index < length ? leaf(index) : bytes32(0);
        bytes32 expectedRoot = calculateRoot(extendHashes(hashes));
        bytes32[] memory proof = libHarness.calculateProof(hashes, index);
        bytes32 root = MerkleMath.proofRoot(index, node, proof, HEIGHT);
        assertEq(root, expectedRoot, "!calculateProof");
    }

    function test_getParent(bytes32 leftChild, bytes32 rightChild) public {
        assertEq(libHarness.getParent(leftChild, rightChild), getParent(leftChild, rightChild));
    }

    function test_getParent_zero(bytes32 node) public {
        test_getParent(0, node);
        test_getParent(node, 0);
        test_getParent(0, 0);
    }

    function test_getParent(bytes32 node, bytes32 sibling, uint256 leafIndex, uint256 nodeHeight) public {
        nodeHeight = nodeHeight % HEIGHT;
        leafIndex = leafIndex % MAX_LENGTH;
        bytes32 parent = libHarness.getParent(node, sibling, leafIndex, nodeHeight);
        // Get the index of the node on level `nodeHeight` the straightforward way
        for (uint256 h = 0; h < nodeHeight; ++h) {
            leafIndex = leafIndex / 2;
        }
        // `node` is a left child if the index is even
        bytes32 expected = leafIndex % 2 == 0 ? getParent(node, sibling) : getParent(sibling, node);
        assertEq(parent, expected);
    }

    function test_proofRoot(bytes32 node, uint256 index, bytes32[] memory proof, uint256 height) public {
        height = bound(height, 1, 32);
        index = index % (1 << height);
        if (proof.length > height) {
            // Resize proof to be the same length
            assembly {
                mstore(proof, height)
            }
            require(proof.length == height, "Dirty assembly trick failed");
        }
        bytes32 root = libHarness.proofRoot(index, node, proof, height);
        bytes32 expected = node;
        for (uint256 h = 0; h < height; ++h) {
            bytes32 sibling = h < proof.length ? proof[h] : bytes32(0);
            if (index % 2 == 0) {
                expected = getParent(expected, sibling);
            } else {
                expected = getParent(sibling, expected);
            }
            index = index / 2;
        }
        assertEq(root, expected);
    }

    // ══════════════════════════════════════════════════ HELPERS ══════════════════════════════════════════════════════

    /// @dev Calculate merkle root for a list of 2**N leafs in the most straightforward way.
    function calculateRoot(bytes32[] memory hashes) public pure returns (bytes32) {
        if (hashes.length == 1) return hashes[0];
        uint256 length = hashes.length / 2;
        bytes32[] memory parents = new bytes32[](length);
        for (uint256 i = 0; i < length; ++i) {
            parents[i] = getParent(hashes[2 * i], hashes[2 * i + 1]);
        }
        return calculateRoot(parents);
    }

    /// @dev Generate N different hashes for tests.
    function generateHashes(uint256 length) public pure returns (bytes32[] memory hashes) {
        hashes = new bytes32[](length);
        for (uint256 i = 0; i < length; ++i) {
            hashes[i] = leaf(i);
        }
    }

    function leaf(uint256 index) public pure returns (bytes32) {
        return keccak256(abi.encode("Leaf", index));
    }

    function getParent(bytes32 leftLeaf, bytes32 rightLeaf) public pure returns (bytes32) {
        if (leftLeaf == bytes32(0) && rightLeaf == bytes32(0)) return bytes32(0);
        return keccak256(abi.encodePacked(leftLeaf, rightLeaf));
    }

    /// @dev Extend `hashes` with `zeroHash` values until list length is MAX_LENGTH
    function extendHashes(bytes32[] memory hashes) public pure returns (bytes32[] memory extended) {
        extended = new bytes32[](MAX_LENGTH);
        for (uint256 i = 0; i < hashes.length; ++i) {
            extended[i] = hashes[i];
        }
        // The remaining items are bytes32(0)
    }
}
