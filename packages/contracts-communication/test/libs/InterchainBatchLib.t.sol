// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {InterchainBatch, InterchainBatchLibHarness, BatchKey} from "../harnesses/InterchainBatchLibHarness.sol";

import {Test} from "forge-std/Test.sol";

// solhint-disable func-name-mixedcase
// solhint-disable ordering
contract InterchainBatchLibTest is Test {
    InterchainBatchLibHarness public libHarness;

    InterchainBatch public mockBatch = InterchainBatch({srcChainId: 1, dbNonce: 2, batchRoot: bytes32(uint256(3))});

    function setUp() public {
        libHarness = new InterchainBatchLibHarness();
    }

    function assertEq(InterchainBatch memory actual, InterchainBatch memory expected) public pure {
        assertEq(actual.srcChainId, expected.srcChainId, "!srcChainId");
        assertEq(actual.dbNonce, expected.dbNonce, "!dbNonce");
        assertEq(actual.batchRoot, expected.batchRoot, "!batchRoot");
    }

    function test_constructLocalBatch() public {
        vm.chainId(1);
        uint64 dbNonce = 2;
        bytes32 batchRoot = bytes32(uint256(3));
        InterchainBatch memory actual = libHarness.constructLocalBatch(dbNonce, batchRoot);
        assertEq(actual, mockBatch);
    }

    function test_constructLocalBatch(uint64 chainId, uint64 dbNonce, bytes32 batchRoot) public {
        vm.chainId(chainId);
        InterchainBatch memory expected = InterchainBatch(chainId, dbNonce, batchRoot);
        InterchainBatch memory actual = libHarness.constructLocalBatch(dbNonce, batchRoot);
        assertEq(actual, expected);
    }

    function test_encodeBatch_roundTrip(InterchainBatch memory batch) public view {
        bytes memory encoded = libHarness.encodeBatch(batch);
        InterchainBatch memory decoded = libHarness.decodeBatch(encoded);
        assertEq(decoded, batch);
    }

    function test_encodeBatchFromMemory_roundTrip(InterchainBatch memory batch) public view {
        bytes memory encoded = libHarness.encodeBatch(batch);
        InterchainBatch memory decoded = libHarness.decodeBatchFromMemory(encoded);
        assertEq(decoded, batch);
    }

    function test_encodeBatchKey_roundTrip(uint64 srcChainId, uint64 dbNonce) public view {
        BatchKey key = libHarness.encodeBatchKey(srcChainId, dbNonce);
        (uint64 decodedSrcChainId, uint64 decodedDbNonce) = libHarness.decodeBatchKey(key);
        assertEq(decodedSrcChainId, srcChainId);
        assertEq(decodedDbNonce, dbNonce);
    }
}
