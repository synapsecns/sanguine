// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {InterchainEntry, InterchainEntryLibHarness} from "../harnesses/InterchainEntryLibHarness.sol";

import {Test} from "forge-std/Test.sol";

contract InterchainEntryLibTest is Test {
    InterchainEntryLibHarness public libHarness;

    InterchainEntry public mockEntry =
        InterchainEntry({srcChainId: 1, dbNonce: 2, srcWriter: bytes32(uint256(3)), dataHash: bytes32(uint256(4))});

    function setUp() public {
        libHarness = new InterchainEntryLibHarness();
    }

    function assertEq(InterchainEntry memory actual, InterchainEntry memory expected) public {
        assertEq(actual.srcChainId, expected.srcChainId, "!srcChainId");
        assertEq(actual.dbNonce, expected.dbNonce, "!dbNonce");
        assertEq(actual.srcWriter, expected.srcWriter, "!srcWriter");
        assertEq(actual.dataHash, expected.dataHash, "!dataHash");
    }

    function test_constructLocalEntry() public {
        vm.chainId(1);
        uint256 dbNonce = 2;
        address srcWriter = address(3);
        bytes32 dataHash = bytes32(uint256(4));
        InterchainEntry memory actual = libHarness.constructLocalEntry(dbNonce, srcWriter, dataHash);
        assertEq(actual, mockEntry);
    }

    function test_constructLocalEntry(uint64 chainId, uint256 dbNonce, address srcWriter, bytes32 dataHash) public {
        vm.chainId(chainId);
        InterchainEntry memory expected = InterchainEntry({
            srcChainId: chainId,
            dbNonce: dbNonce,
            srcWriter: bytes32(uint256(uint160(srcWriter))),
            dataHash: dataHash
        });
        InterchainEntry memory actual = libHarness.constructLocalEntry(dbNonce, srcWriter, dataHash);
        assertEq(actual, expected);
    }

    function test_entryKey() public {
        bytes32 expected = keccak256(abi.encode(1, 2));
        assertEq(libHarness.entryKey(mockEntry), expected);
    }

    function test_entryKey(InterchainEntry memory entry) public {
        bytes32 expected = keccak256(abi.encode(entry.srcChainId, entry.dbNonce));
        assertEq(libHarness.entryKey(entry), expected);
    }

    function test_entryValue() public {
        bytes32 expected = keccak256(abi.encode(3, 4));
        assertEq(libHarness.entryValue(mockEntry), expected);
    }

    function test_entryValue(InterchainEntry memory entry) public {
        bytes32 expected = keccak256(abi.encode(entry.srcWriter, entry.dataHash));
        assertEq(libHarness.entryValue(entry), expected);
    }
}
