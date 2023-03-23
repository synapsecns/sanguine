// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import { MerkleLib } from "../../../contracts/libs/Merkle.sol";

import { SynapseLibraryTest } from "../../utils/SynapseLibraryTest.t.sol";
import { MerkleListHarness } from "../../harnesses/libs/MerkleListHarness.t.sol";

// solhint-disable func-name-mixedcase
contract MerkleListLibraryTest is SynapseLibraryTest {
    uint256 public constant MAX_LENGTH = 32;

    MerkleListHarness internal libHarness;

    function setUp() public override {
        libHarness = new MerkleListHarness();
    }

    function test_calculateRoot(uint256 length) public {
        // length should be in [1 .. MAX_LENGTH] range
        length = bound(length, 1, MAX_LENGTH);
        bytes32[] memory hashes = _generateHashes(length);
        bytes32[] memory extended = _extendHashes(hashes);
        bytes32 expectedRoot = _calculateRoot(extended);
        bytes32 root = libHarness.calculateRoot(hashes);
        assertEq(root, expectedRoot, "Merkle Root incorrect");
    }

    function test_calculateProof(uint256 length, uint256 index) public {
        // length should be in [1 .. MAX_LENGTH] range
        length = bound(length, 1, MAX_LENGTH);
        // index should be in [0 .. length) range
        index = bound(index, 0, length - 1);
        bytes32[] memory hashes = _generateHashes(length);
        bytes32[] memory extended = _extendHashes(hashes);
        bytes32 expectedRoot = _calculateRoot(extended);

        bytes32[] memory proof = libHarness.calculateProof(hashes, index);
        // First element is "right sub-leaf"
        proof[0] = _rightLeaf(index);
        // Index of "left sub-leaf" is twice the index of state
        bytes32 root = MerkleLib.branchRoot(_leftLeaf(index), proof, 2 * index);
        assertEq(root, expectedRoot, "!calculateProof");
    }

    /// @dev Calculate merkle root for a list of 2**N leafs in the most straightforward way.
    function _calculateRoot(bytes32[] memory hashes) internal pure returns (bytes32) {
        if (hashes.length == 1) return hashes[0];
        uint256 length = hashes.length / 2;
        bytes32[] memory parents = new bytes32[](length);
        for (uint256 i = 0; i < length; ++i) {
            parents[i] = keccak256(bytes.concat(hashes[2 * i], hashes[2 * i + 1]));
        }
        return _calculateRoot(parents);
    }

    /// @dev Generate N different hashes for tests.
    function _generateHashes(uint256 length) internal pure returns (bytes32[] memory hashes) {
        hashes = new bytes32[](length);
        for (uint256 i = 0; i < length; ++i) {
            hashes[i] = keccak256(bytes.concat(_leftLeaf(i), _rightLeaf(i)));
        }
    }

    function _leftLeaf(uint256 index) internal pure returns (bytes32) {
        return keccak256(abi.encode("Left", index));
    }

    function _rightLeaf(uint256 index) internal pure returns (bytes32) {
        return keccak256(abi.encode("Right", index));
    }

    /// @dev Extend `hashes` with `zeroHash` values until list length is a power of two.
    function _extendHashes(bytes32[] memory hashes)
        internal
        pure
        returns (bytes32[] memory extended)
    {
        uint256 length = hashes.length;
        // Find the lowest power of two that is greater or equal than length
        uint256 lengthExtended = 1;
        while (lengthExtended < length) {
            lengthExtended *= 2;
        }
        extended = new bytes32[](lengthExtended);
        for (uint256 i = 0; i < length; ++i) {
            extended[i] = hashes[i];
        }
        // The remaining items are bytes32(0)
    }
}
