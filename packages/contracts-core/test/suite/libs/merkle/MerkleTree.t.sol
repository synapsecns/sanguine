// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {AGENT_TREE_HEIGHT, ORIGIN_TREE_HEIGHT} from "../../../../contracts/libs/merkle/MerkleTree.sol";

import {MerkleMathHarness} from "../../../harnesses/libs/merkle/MerkleMathHarness.t.sol";
import {MerkleTreeHarness} from "../../../harnesses/libs/merkle/MerkleTreeHarness.t.sol";

import {Random} from "../../../utils/libs/Random.t.sol";
import {SynapseLibraryTest} from "../../../utils/SynapseLibraryTest.t.sol";

// solhint-disable func-name-mixedcase
contract MerkleTreeTest is SynapseLibraryTest {
    uint256 public constant RUNS = 1000;
    uint256 public constant DYNAMIC_AMOUNT = 200;

    MerkleTreeHarness public libHarness;
    MerkleMathHarness public mathHarness;
    bytes32[] public leafs;

    function setUp() public {
        libHarness = new MerkleTreeHarness();
        mathHarness = new MerkleMathHarness();
    }

    function test_baseTree() public {
        assertEq(libHarness.rootBase(0), 0);
        for (uint256 i = 0; i < RUNS; ++i) {
            bytes32 leaf = getLeaf(i);
            leafs.push(leaf);
            libHarness.insertBase(i + 1, leaf);
            assertEq(libHarness.rootBase(i + 1), mathHarness.calculateRoot(leafs, ORIGIN_TREE_HEIGHT));
        }
    }

    function test_historicalTree() public {
        bytes32[] memory roots = new bytes32[](RUNS + 1);
        roots[0] = 0;
        libHarness.initializeRoots();
        for (uint256 i = 0; i < RUNS; ++i) {
            bytes32 leaf = getLeaf(i);
            leafs.push(leaf);
            libHarness.insert(leaf);
            roots[i + 1] = mathHarness.calculateRoot(leafs, ORIGIN_TREE_HEIGHT);
        }
        for (uint256 i = 0; i <= RUNS; ++i) {
            assertEq(libHarness.root(i), roots[i]);
        }
    }

    function test_historicalTree_revert_initializeTwice() public {
        libHarness.initializeRoots();
        vm.expectRevert();
        libHarness.initializeRoots();
    }

    function test_dynamicTree() public {
        for (uint256 i = 0; i < DYNAMIC_AMOUNT; ++i) {
            leafs.push(0);
        }
        Random memory random = Random("very random seed");
        for (uint256 i = 0; i < RUNS; ++i) {
            uint256 index = random.nextUint256() % DYNAMIC_AMOUNT;
            bytes32 oldValue = leafs[index];
            bytes32 newValue = random.next();
            bytes32[] memory proof = mathHarness.calculateProof(leafs, index);
            bytes32 newRoot = libHarness.update(index, oldValue, proof, newValue);
            leafs[index] = newValue;
            assertEq(newRoot, mathHarness.calculateRoot(leafs, AGENT_TREE_HEIGHT));
        }
    }

    function getLeaf(uint256 index) public pure returns (bytes32) {
        return keccak256(abi.encode("Leaf", index));
    }
}
