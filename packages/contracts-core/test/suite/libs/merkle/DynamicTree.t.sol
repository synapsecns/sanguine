// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {Test} from "forge-std/Test.sol";

import {DynamicTreeHarness} from "../../../harnesses/libs/merkle/DynamicTreeHarness.t.sol";
import {DynamicProofGenerator} from "../../../utils/proof/DynamicProofGenerator.t.sol";

import {Random} from "../../../utils/libs/Random.t.sol";

// solhint-disable func-name-mixedcase
contract DynamicTreeTest is Test {
    DynamicProofGenerator internal proofGen;
    DynamicTreeHarness internal tree;

    uint256 internal constant UPDATES = 50;

    function setUp() public {
        tree = new DynamicTreeHarness();
        proofGen = new DynamicProofGenerator();
    }

    function test_smallIndexes(Random memory random) public {
        for (uint256 i = 0; i < UPDATES; ++i) {
            _testUpdateValue(random, 16);
        }
    }

    function test_mediumIndexes(Random memory random) public {
        for (uint256 i = 0; i < UPDATES; ++i) {
            _testUpdateValue(random, 420);
        }
    }

    function test_unlimitedIndexes(Random memory random) public {
        for (uint256 i = 0; i < UPDATES; ++i) {
            _testUpdateValue(random, type(uint32).max);
        }
    }

    function _testUpdateValue(Random memory random, uint256 maxIndex) internal {
        uint256 index = random.nextUint32() % maxIndex;
        bytes32 oldValue = proofGen.getLeaf(index);
        bytes32[] memory proof = proofGen.getProof(index);
        bytes32 newValue = random.next();
        tree.update(index, oldValue, proof, newValue);
        proofGen.update(index, newValue);
        assertEq(tree.root(), proofGen.getRoot());
    }
}
