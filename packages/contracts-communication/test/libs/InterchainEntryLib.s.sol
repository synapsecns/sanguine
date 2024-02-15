// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {InterchainEntry, InterchainEntryLibHarness} from "../harnesses/InterchainEntryLibHarness.sol";

import {Test} from "forge-std/Test.sol";

contract InterchainEntryLibTest is Test {
    InterchainEntryLibHarness public libHarness;

    InterchainEntry public mockEntry =
        InterchainEntry({srcChainId: 1, srcWriter: bytes32(uint256(2)), writerNonce: 3, dataHash: bytes32(uint256(4))});

    function setUp() public {
        libHarness = new InterchainEntryLibHarness();
    }

    function assertEq(InterchainEntry memory actual, InterchainEntry memory expected) public {
        assertEq(actual.srcChainId, expected.srcChainId, "!srcChainId");
        assertEq(actual.srcWriter, expected.srcWriter, "!srcWriter");
        assertEq(actual.writerNonce, expected.writerNonce, "!writerNonce");
        assertEq(actual.dataHash, expected.dataHash, "!dataHash");
    }

    function test_constructLocalEntry() public {
        vm.chainId(1);
        address srcWriter = address(2);
        uint256 writerNonce = 3;
        bytes32 dataHash = bytes32(uint256(4));
        InterchainEntry memory actual = libHarness.constructLocalEntry(srcWriter, writerNonce, dataHash);
        assertEq(actual, mockEntry);
    }

    function test_constructLocalEntry(
        uint64 chainId,
        address srcWriter,
        uint256 writerNonce,
        bytes32 dataHash
    )
        public
    {
        vm.chainId(chainId);
        InterchainEntry memory expected = InterchainEntry({
            srcChainId: chainId,
            srcWriter: bytes32(uint256(uint160(srcWriter))),
            writerNonce: writerNonce,
            dataHash: dataHash
        });
        InterchainEntry memory actual = libHarness.constructLocalEntry(srcWriter, writerNonce, dataHash);
        assertEq(actual, expected);
    }

    function test_entryId() public {
        bytes32 expected = keccak256(abi.encode(1, 2, 3));
        assertEq(libHarness.entryId(mockEntry), expected);
    }

    function test_entryId(InterchainEntry memory entry) public {
        bytes32 expected = keccak256(abi.encode(entry.srcChainId, entry.srcWriter, entry.writerNonce));
        assertEq(libHarness.entryId(entry), expected);
    }
}
