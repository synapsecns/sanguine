// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {InterchainBatch, ModuleBatchLibHarness} from "../harnesses/ModuleBatchLibHarness.sol";

import {Test} from "forge-std/Test.sol";

// solhint-disable func-name-mixedcase
// solhint-disable ordering
contract ModuleBatchLibTest is Test {
    ModuleBatchLibHarness public libHarness;

    function setUp() public {
        libHarness = new ModuleBatchLibHarness();
    }

    function assertEq(InterchainBatch memory actual, InterchainBatch memory expected) public pure {
        assertEq(actual.srcChainId, expected.srcChainId, "!srcChainId");
        assertEq(actual.dbNonce, expected.dbNonce, "!dbNonce");
        assertEq(actual.batchRoot, expected.batchRoot, "!batchRoot");
    }

    function test_encodeVersionedModuleBatch_roundTrip(
        bytes memory versionedBatch,
        bytes memory moduleData
    )
        public
        view
    {
        bytes memory encoded = libHarness.encodeVersionedModuleBatch(versionedBatch, moduleData);
        (bytes memory decodedVersionedBatch, bytes memory decodedModuleData) =
            libHarness.decodeVersionedModuleBatch(encoded);
        assertEq(decodedVersionedBatch, versionedBatch);
        assertEq(decodedModuleData, moduleData);
    }
}
