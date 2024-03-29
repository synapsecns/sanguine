// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {InterchainBatch, InterchainBatchLibHarness} from "../harnesses/InterchainBatchLibHarness.sol";

import {Test} from "forge-std/Test.sol";

// solhint-disable func-name-mixedcase
contract InterchainBatchLibTest is Test {
    InterchainBatchLibHarness public libHarness;

    InterchainBatch public mockBatch = InterchainBatch({srcChainId: 1, dbNonce: 2, batchRoot: bytes32(uint256(3))});

    function setUp() public {
        libHarness = new InterchainBatchLibHarness();
    }

    function assertEq(InterchainBatch memory actual, InterchainBatch memory expected) public {
        assertEq(actual.srcChainId, expected.srcChainId, "!srcChainId");
        assertEq(actual.dbNonce, expected.dbNonce, "!dbNonce");
        assertEq(actual.batchRoot, expected.batchRoot, "!batchRoot");
    }

    function test_constructLocalBatch() public {
        vm.chainId(1);
        uint256 dbNonce = 2;
        bytes32 batchRoot = bytes32(uint256(3));
        InterchainBatch memory actual = libHarness.constructLocalBatch(dbNonce, batchRoot);
        assertEq(actual, mockBatch);
    }

    function test_constructLocalBatch(uint64 chainId, uint256 dbNonce, bytes32 batchRoot) public {
        vm.chainId(chainId);
        InterchainBatch memory expected = InterchainBatch(chainId, dbNonce, batchRoot);
        InterchainBatch memory actual = libHarness.constructLocalBatch(dbNonce, batchRoot);
        assertEq(actual, expected);
    }

    function test_batchKey() public {
        bytes32 expected = keccak256(abi.encode(1, 2));
        bytes32 actual = libHarness.batchKey(mockBatch);
        assertEq(actual, expected);
    }

    function test_batchKey(InterchainBatch memory batch) public {
        bytes32 expected = keccak256(abi.encode(batch.srcChainId, batch.dbNonce));
        bytes32 actual = libHarness.batchKey(batch);
        assertEq(actual, expected);
    }

    function test_encodeVersionedBatchFromMemory_roundTrip(uint16 dbVersion, InterchainBatch memory batch) public {
        bytes memory versionedBatch = libHarness.encodeVersionedBatch(dbVersion, batch);
        (uint16 dbVersion_, InterchainBatch memory batch_) = libHarness.decodeVersionedBatchFromMemory(versionedBatch);
        assertEq(dbVersion_, dbVersion);
        assertEq(batch_, batch);
    }

    function test_encodeVersionedBatch_roundTrip(uint16 dbVersion, InterchainBatch memory batch) public {
        bytes memory versionedBatch = libHarness.encodeVersionedBatch(dbVersion, batch);
        (uint16 dbVersion_, InterchainBatch memory batch_) = libHarness.decodeVersionedBatch(versionedBatch);
        assertEq(dbVersion_, dbVersion);
        assertEq(batch_, batch);
    }

    function test_encodeBatch_roundTrip(InterchainBatch memory batch) public {
        bytes memory encoded = libHarness.encodeBatch(batch);
        InterchainBatch memory decoded = libHarness.decodeBatch(encoded);
        assertEq(decoded, batch);
    }

    function test_encodeBatchFromMemory_roundTrip(InterchainBatch memory batch) public {
        bytes memory encoded = libHarness.encodeBatch(batch);
        InterchainBatch memory decoded = libHarness.decodeBatchFromMemory(encoded);
        assertEq(decoded, batch);
    }
}
