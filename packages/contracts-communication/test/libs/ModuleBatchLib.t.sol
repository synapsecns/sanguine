// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {InterchainBatch, ModuleBatchLibHarness} from "../harnesses/ModuleBatchLibHarness.sol";

import {Test} from "forge-std/Test.sol";

// solhint-disable func-name-mixedcase
contract ModuleBatchLibTest is Test {
    ModuleBatchLibHarness public libHarness;

    function setUp() public {
        libHarness = new ModuleBatchLibHarness();
    }

    function assertEq(InterchainBatch memory actual, InterchainBatch memory expected) public {
        assertEq(actual.srcChainId, expected.srcChainId, "!srcChainId");
        assertEq(actual.dbNonce, expected.dbNonce, "!dbNonce");
        assertEq(actual.batchRoot, expected.batchRoot, "!batchRoot");
    }

    function test_roundTrip(InterchainBatch memory batch, bytes memory moduleData) public {
        bytes memory encoded = libHarness.encodeModuleBatch(batch, moduleData);
        (InterchainBatch memory decodedBatch, bytes memory decodedModuleData) = libHarness.decodeModuleBatch(encoded);
        assertEq(decodedBatch, batch);
        assertEq(decodedModuleData, moduleData);
    }
}
