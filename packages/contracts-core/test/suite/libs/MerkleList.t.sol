// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {MerkleLib} from "../../../contracts/libs/Merkle.sol";

import {SynapseLibraryTest} from "../../utils/SynapseLibraryTest.t.sol";
import {MerkleListHarness} from "../../harnesses/libs/MerkleListHarness.t.sol";

// solhint-disable func-name-mixedcase
contract MerkleListLibraryTest is SynapseLibraryTest {
    uint256 public constant HEIGHT = 8;
    uint256 public constant MAX_LENGTH = 1 << HEIGHT;

    MerkleListHarness internal libHarness;

    function setUp() public {
        libHarness = new MerkleListHarness();
    }

    function test_calculateRoot(uint256 length) public {
        // length should be in [1 .. MAX_LENGTH] range
        length = bound(length, 1, MAX_LENGTH);
        bytes32[] memory hashes = _generateHashes(length);
        bytes32[] memory extended = _extendHashes(hashes);
        bytes32 expectedRoot = _calculateRoot(extended);
        bytes32 root = libHarness.calculateRoot(hashes, HEIGHT);
        assertEq(root, expectedRoot, "Merkle Root incorrect");
    }

    function test_calculateProof(uint256 length, uint256 index) public {
        // length should be in [1 .. MAX_LENGTH] range
        length = bound(length, 1, MAX_LENGTH);
        bytes32[] memory hashes = _generateHashes(length);
        // index should be in [0 .. MAX_LENGTH) range
        index = bound(index, 0, MAX_LENGTH - 1);
        // Check proofs for zero leafs outside of the list as well
        bytes32 node = index < length ? leaf(index) : bytes32(0);
        bytes32 expectedRoot = _calculateRoot(_extendHashes(hashes));
        bytes32[] memory proof = libHarness.calculateProof(hashes, index);
        bytes32 root = MerkleLib.proofRoot(index, node, proof, HEIGHT);
        assertEq(root, expectedRoot, "!calculateProof");
    }

    /// @dev Calculate merkle root for a list of 2**N leafs in the most straightforward way.
    function _calculateRoot(bytes32[] memory hashes) internal pure returns (bytes32) {
        if (hashes.length == 1) return hashes[0];
        uint256 length = hashes.length / 2;
        bytes32[] memory parents = new bytes32[](length);
        for (uint256 i = 0; i < length; ++i) {
            parents[i] = _getParent(hashes[2 * i], hashes[2 * i + 1]);
        }
        return _calculateRoot(parents);
    }

    /// @dev Generate N different hashes for tests.
    function _generateHashes(uint256 length) internal pure returns (bytes32[] memory hashes) {
        hashes = new bytes32[](length);
        for (uint256 i = 0; i < length; ++i) {
            hashes[i] = leaf(i);
        }
    }

    function leaf(uint256 index) internal pure returns (bytes32) {
        return keccak256(abi.encode("Leaf", index));
    }

    function _getParent(bytes32 leftLeaf, bytes32 rightLeaf) internal pure returns (bytes32) {
        if (leftLeaf == bytes32(0) && rightLeaf == bytes32(0)) return bytes32(0);
        return keccak256(bytes.concat(leftLeaf, rightLeaf));
    }

    /// @dev Extend `hashes` with `zeroHash` values until list length is MAX_LENGTH
    function _extendHashes(bytes32[] memory hashes) internal pure returns (bytes32[] memory extended) {
        extended = new bytes32[](MAX_LENGTH);
        for (uint256 i = 0; i < hashes.length; ++i) {
            extended[i] = hashes[i];
        }
        // The remaining items are bytes32(0)
    }
}
