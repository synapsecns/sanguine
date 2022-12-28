//  SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import "forge-std/Test.sol";

import { MerkleLib } from "../../../contracts/libs/Merkle.sol";
import { HistoricalProofGenerator } from "./HistoricalProofGenerator.t.sol";

// solhint-disable func-name-mixedcase
contract HistoricalProofGeneratorTest is Test {
    using MerkleLib for MerkleLib.Tree;

    MerkleLib.Tree internal tree;
    HistoricalProofGenerator internal proofGen;

    uint256 internal constant TREE_DEPTH = 32;
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
        roots[0] = tree.root(0);
        // Sanity check against precomputed root for an empty Merkle tree
        assertEq(
            roots[0],
            hex"27ae5ba08d7291c96c8cbddcc148bf48a6d68c7974b94356f53754ef6171d757",
            "!root(empty tree)"
        );
        for (uint256 count = 1; count <= AMOUNT; ++count) {
            proofGen.insert(leafs[count - 1]);
            tree.insert(count, leafs[count - 1]);
            // Save merkle root after inserting `count` leafs
            roots[count] = tree.root(count);
        }
        for (uint256 count = 0; count <= AMOUNT; ++count) {
            assertEq(proofGen.getRoot(count), roots[count]);
        }
    }

    function test_generateProof() public {
        test_getRoot();
        for (uint256 count = 1; count <= AMOUNT; ++count) {
            // Use a "historical root" for proving
            // root is already tested against the MerkleLib implementation
            bytes32 root = proofGen.getRoot(count);
            // Check generated proofs for every leaf that precedes the "historical root"
            for (uint256 index = 0; index < count; ++index) {
                bytes32[TREE_DEPTH] memory proof = proofGen.getProof(index, count);
                assertEq(MerkleLib.branchRoot(leafs[index], proof, index), root, "Invalid proof");
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
