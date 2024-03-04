// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {InterchainEntry, ModuleEntryLibHarness} from "../harnesses/ModuleEntryLibHarness.sol";

import {Test} from "forge-std/Test.sol";

// solhint-disable func-name-mixedcase
contract ModuleEntryLibTest is Test {
    ModuleEntryLibHarness public libHarness;

    function setUp() public {
        libHarness = new ModuleEntryLibHarness();
    }

    function assertEq(InterchainEntry memory actual, InterchainEntry memory expected) public {
        assertEq(actual.srcChainId, expected.srcChainId, "!srcChainId");
        assertEq(actual.dbNonce, expected.dbNonce, "!dbNonce");
        assertEq(actual.srcWriter, expected.srcWriter, "!srcWriter");
        assertEq(actual.dataHash, expected.dataHash, "!dataHash");
    }

    function test_roundTrip(InterchainEntry memory entry, bytes memory moduleData) public {
        bytes memory encoded = libHarness.encodeModuleEntry(entry, moduleData);
        (InterchainEntry memory decodedEntry, bytes memory decodedModuleData) = libHarness.decodeModuleEntry(encoded);
        assertEq(decodedEntry, entry);
        assertEq(decodedModuleData, moduleData);
    }
}
