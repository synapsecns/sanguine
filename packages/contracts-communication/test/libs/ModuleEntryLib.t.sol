// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {InterchainEntry} from "../../contracts/libs/InterchainEntry.sol";

import {ModuleEntryLibHarness} from "../harnesses/ModuleEntryLibHarness.sol";

import {Test} from "forge-std/Test.sol";

// solhint-disable func-name-mixedcase
// solhint-disable ordering
contract ModuleEntryLibTest is Test {
    ModuleEntryLibHarness public libHarness;

    function setUp() public {
        libHarness = new ModuleEntryLibHarness();
    }

    function assertEq(InterchainEntry memory actual, InterchainEntry memory expected) public pure {
        assertEq(actual.srcChainId, expected.srcChainId, "!srcChainId");
        assertEq(actual.dbNonce, expected.dbNonce, "!dbNonce");
        assertEq(actual.entryValue, expected.entryValue, "!entryValue");
    }

    function test_encodeVersionedModuleEntry_roundTrip(
        bytes memory versionedEntry,
        bytes memory moduleData
    )
        public
        view
    {
        bytes memory encoded = libHarness.encodeVersionedModuleEntry(versionedEntry, moduleData);
        (bytes memory decodedVersionedEntry, bytes memory decodedModuleData) =
            libHarness.decodeVersionedModuleEntry(encoded);
        assertEq(decodedVersionedEntry, versionedEntry);
        assertEq(decodedModuleData, moduleData);
    }
}
