//  SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {Test} from "forge-std/Test.sol";

import {BaseTree, MerkleMath, AGENT_TREE_HEIGHT} from "../../../contracts/libs/merkle/MerkleTree.sol";
import {DynamicProofGenerator} from "./DynamicProofGenerator.t.sol";

import {Random} from "../libs/Random.t.sol";

// solhint-disable func-name-mixedcase
contract DynamicProofGeneratorTest is Test {
    BaseTree internal tree;
    DynamicProofGenerator internal proofGen;

    function setUp() public {
        proofGen = new DynamicProofGenerator();
    }

    function test_rootParity(Random memory random) public {
        _checkRootParity(random, 10, 69);
    }

    function test_rootParity_withEmptyLeafs(Random memory random) public {
        _checkRootParity(random, 20, 15);
    }

    function test_getProof() public {
        uint256 amount = 100;
        uint256 updates = 4200;
        Random memory random = Random("Very random much wow");
        bytes32[] memory leafs = new bytes32[](amount);
        for (uint256 i = 0; i < updates; ++i) {
            uint256 index = random.nextUint256() % amount;
            bytes32 value = random.next();
            leafs[index] = value;
            proofGen.update(index, value);
        }
        bytes32 root = proofGen.getRoot();
        // Check the produced proofs for the latest values
        for (uint256 i = 0; i < amount; ++i) {
            bytes32[] memory proof = proofGen.getProof(i);
            assertEq(MerkleMath.proofRoot(i, leafs[i], proof, AGENT_TREE_HEIGHT), root);
        }
    }

    function _checkRootParity(Random memory random, uint256 amount, uint256 updates) internal {
        bytes32[] memory leafs = new bytes32[](amount);
        for (uint256 i = 0; i < updates; ++i) {
            uint256 index = random.nextUint256() % amount;
            bytes32 value = random.next();
            leafs[index] = value;
            proofGen.update(index, value);
        }
        // Construct the Tree to test against
        for (uint256 i = 0; i < amount; ++i) {
            tree.insertBase(i + 1, leafs[i]);
        }
        assertEq(proofGen.getRoot(), tree.rootBase(amount));
    }
}
