// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {Test} from "forge-std/Test.sol";

import {BaseTree, MerkleMath} from "../../../contracts/libs/merkle/MerkleTree.sol";
import {ORIGIN_TREE_HEIGHT, ProofGenerator} from "./ProofGenerator.t.sol";

// solhint-disable func-name-mixedcase
contract ProofGeneratorTest is Test {
    uint256 internal constant MAX_COUNT = 10;

    BaseTree internal tree;
    ProofGenerator internal gen;
    uint256 internal length;
    bytes32[] internal leafs;

    function setUp() public {
        gen = new ProofGenerator();
    }

    function test_createTree() public {
        length = MAX_COUNT;
        _checkCreateTree();
    }

    function test_generateProofs() public {
        test_createTree();
        _checkGenerateProofs();
    }

    function test_recreateTree_biggerSize() public {
        test_createTree();
        length = MAX_COUNT * 2;
        _checkCreateTree();
        _checkGenerateProofs();
    }

    function test_recreateTree_smallerSize() public {
        test_createTree();
        length = MAX_COUNT / 2;
        _checkCreateTree();
        _checkGenerateProofs();
    }

    function _createTestData() internal {
        delete tree;
        leafs = new bytes32[](length);
        for (uint256 i = 0; i < length; ++i) {
            bytes32 node = keccak256(abi.encode(length, i));
            leafs[i] = node;
            tree.insertBase(i + 1, node);
        }
    }

    function _checkCreateTree() internal {
        _createTestData();
        gen.createTree(leafs);
        // Leafs should match the lowest depth level
        for (uint256 i = 0; i < length; ++i) {
            assertEq(gen.getNode(0, i), leafs[i], "!leaf");
        }
        // Non-existing leaf should be zero
        assertEq(gen.getNode(0, length), bytes32(0), "!zero");
        // Merkle root should match
        assertEq(gen.getRoot(), tree.rootBase(length), "!root");
    }

    function _checkGenerateProofs() internal {
        bytes32 root = tree.rootBase(length);
        // Should be able to generate a valid proof for any existing leafs
        for (uint256 i = 0; i < length; ++i) {
            bytes32[] memory proof = gen.getProof(i);
            assertEq(MerkleMath.proofRoot(i, leafs[i], proof, ORIGIN_TREE_HEIGHT), root, "!proof");
        }
        // Cool side effect: could prove message non-inclusion at next index
        {
            uint256 index = length;
            // Should be able to generate a valid proof for a null leaf
            bytes32[] memory proof = gen.getProof(index);
            assertEq(MerkleMath.proofRoot(index, bytes32(0), proof, ORIGIN_TREE_HEIGHT), root, "!proof");
        }
        // Cool side effect: could prove message non-inclusion at index from a distant future
        {
            uint256 index = length + 42_069;
            // Should be able to generate a valid proof for a null leaf
            bytes32[] memory proof = gen.getProof(index);
            assertEq(MerkleMath.proofRoot(index, bytes32(0), proof, ORIGIN_TREE_HEIGHT), root, "!proof");
        }
    }
}
