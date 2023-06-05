//  SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {Test} from "forge-std/Test.sol";

import {BaseTree, MerkleMath} from "../../../contracts/libs/merkle/MerkleTree.sol";
import {HistoricalProofGenerator} from "./HistoricalProofGenerator.t.sol";

// solhint-disable func-name-mixedcase
contract HistoricalProofGeneratorTest is Test {
    BaseTree internal tree;
    HistoricalProofGenerator internal proofGen;

    uint256 internal constant ORIGIN_TREE_HEIGHT = 32;
    uint256 internal constant AMOUNT = 10;
    bytes32[] internal leafs;

    function setUp() public {
        proofGen = new HistoricalProofGenerator();
        leafs = new bytes32[](AMOUNT);
        for (uint256 i = 0; i < AMOUNT; ++i) {
            leafs[i] = keccak256(abi.encode("test", i));
        }
    }

    function test_insert() public {
        // Just to get an idea of how much gas this could've wasted if implemented on-chain
        for (uint256 index = 0; index < AMOUNT; ++index) {
            proofGen.insert(leafs[index]);
        }
    }

    function test_getRoot() public {
        bytes32[] memory roots = new bytes32[](AMOUNT + 1);
        roots[0] = tree.rootBase(0);
        // Sanity check against precomputed root for an empty Merkle tree
        assertEq(roots[0], bytes32(0), "!root(empty tree)");
        for (uint256 count = 1; count <= AMOUNT; ++count) {
            proofGen.insert(leafs[count - 1]);
            tree.insertBase(count, leafs[count - 1]);
            // Save merkle root after inserting `count` leafs
            roots[count] = tree.rootBase(count);
        }
        for (uint256 count = 0; count <= AMOUNT; ++count) {
            assertEq(proofGen.getRoot(count), roots[count]);
        }
    }

    function test_generateProof() public {
        test_getRoot();
        for (uint256 count = 1; count <= AMOUNT; ++count) {
            // Use a "historical root" for proving
            // root is already tested against the MerkleTree implementation
            bytes32 root = proofGen.getRoot(count);
            // Check generated proofs for every leaf that precedes the "historical root"
            for (uint256 index = 0; index < count; ++index) {
                bytes32[] memory proof = proofGen.getProof(index, count);
                assertEq(MerkleMath.proofRoot(index, leafs[index], proof, ORIGIN_TREE_HEIGHT), root, "Invalid proof");
            }
        }
    }

    function test_allElementsEqual() public {
        // Weird scenario where all leafs are equal.
        // Everything should be working though.
        for (uint256 i = 0; i < AMOUNT; ++i) {
            leafs[i] = keccak256(abi.encode("test"));
        }
        test_generateProof();
    }
}
