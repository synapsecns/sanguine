// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {InterchainEntry, InterchainEntryLibHarness, EntryKey} from "../harnesses/InterchainEntryLibHarness.sol";

import {Test} from "forge-std/Test.sol";

// solhint-disable func-name-mixedcase
// solhint-disable ordering
contract InterchainEntryLibTest is Test {
    InterchainEntryLibHarness public libHarness;

    function setUp() public {
        libHarness = new InterchainEntryLibHarness();
    }

    function assertEq(InterchainEntry memory actual, InterchainEntry memory expected) public pure {
        assertEq(actual.srcChainId, expected.srcChainId, "!srcChainId");
        assertEq(actual.dbNonce, expected.dbNonce, "!dbNonce");
        assertEq(actual.entryValue, expected.entryValue, "!entryValue");
    }

    function test_constructLocalEntry() public {
        vm.chainId(1);
        uint64 dbNonce = 2;
        bytes32 entryValue = bytes32(uint256(3));
        InterchainEntry memory actual = libHarness.constructLocalEntry(dbNonce, entryValue);
        InterchainEntry memory expected = InterchainEntry({srcChainId: 1, dbNonce: dbNonce, entryValue: entryValue});
        assertEq(actual, expected);
    }

    function test_constructLocalEntry(uint64 chainId, uint64 dbNonce, bytes32 entryValue) public {
        vm.chainId(chainId);
        InterchainEntry memory expected =
            InterchainEntry({srcChainId: chainId, dbNonce: dbNonce, entryValue: entryValue});
        InterchainEntry memory actual = libHarness.constructLocalEntry(dbNonce, entryValue);
        assertEq(actual, expected);
    }

    function test_getEntryValue_address() public view {
        address srcWriter = address(3);
        bytes32 digest = bytes32(uint256(4));
        bytes32 expected = keccak256(abi.encode(srcWriter, digest));
        assertEq(libHarness.getEntryValue(srcWriter, digest), expected);
    }

    function test_getEntryValue_address(address srcWriter, bytes32 digest) public view {
        bytes32 expected = keccak256(abi.encode(srcWriter, digest));
        assertEq(libHarness.getEntryValue(srcWriter, digest), expected);
    }

    function test_getEntryValue_bytes32() public view {
        bytes32 srcWriter = bytes32(uint256(3));
        bytes32 digest = bytes32(uint256(4));
        bytes32 expected = keccak256(abi.encode(srcWriter, digest));
        assertEq(libHarness.getEntryValue(srcWriter, digest), expected);
    }

    function test_getEntryValue_bytes32(bytes32 srcWriter, bytes32 digest) public view {
        bytes32 expected = keccak256(abi.encode(srcWriter, digest));
        assertEq(libHarness.getEntryValue(srcWriter, digest), expected);
    }

    function test_encodeEntry_roundTrip(InterchainEntry memory entry) public view {
        bytes memory encoded = libHarness.encodeEntry(entry);
        InterchainEntry memory decoded = libHarness.decodeEntry(encoded);
        assertEq(decoded, entry);
    }

    function test_encodeEntryFromMemory_roundTrip(InterchainEntry memory entry) public view {
        bytes memory encoded = libHarness.encodeEntry(entry);
        InterchainEntry memory decoded = libHarness.decodeEntryFromMemory(encoded);
        assertEq(decoded, entry);
    }

    function test_encodeEntryKey_roundTrip(uint64 srcChainId, uint64 dbNonce) public view {
        EntryKey key = libHarness.encodeEntryKey(srcChainId, dbNonce);
        (uint64 decodedSrcChainId, uint64 decodedDbNonce) = libHarness.decodeEntryKey(key);
        assertEq(decodedSrcChainId, srcChainId);
        assertEq(decodedDbNonce, dbNonce);
    }
}
