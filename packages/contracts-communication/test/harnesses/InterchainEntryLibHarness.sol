// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {InterchainEntry, InterchainEntryLib, EntryKey} from "../../contracts/libs/InterchainEntry.sol";

contract InterchainEntryLibHarness {
    function constructLocalEntry(uint64 dbNonce, bytes32 entryValue) external view returns (InterchainEntry memory) {
        return InterchainEntryLib.constructLocalEntry(dbNonce, entryValue);
    }

    function getEntryValue(bytes32 srcWriter, bytes32 digest) external pure returns (bytes32) {
        return InterchainEntryLib.getEntryValue(srcWriter, digest);
    }

    function getEntryValue(address srcWriter, bytes32 digest) external pure returns (bytes32) {
        return InterchainEntryLib.getEntryValue(srcWriter, digest);
    }

    function encodeEntry(InterchainEntry memory entry) external pure returns (bytes memory) {
        return InterchainEntryLib.encodeEntry(entry);
    }

    function decodeEntry(bytes calldata entry) external pure returns (InterchainEntry memory) {
        return InterchainEntryLib.decodeEntry(entry);
    }

    function decodeEntryFromMemory(bytes memory entry) external pure returns (InterchainEntry memory) {
        return InterchainEntryLib.decodeEntryFromMemory(entry);
    }

    function encodeEntryKey(uint64 srcChainId, uint64 dbNonce) external pure returns (EntryKey) {
        return InterchainEntryLib.encodeEntryKey(srcChainId, dbNonce);
    }

    function decodeEntryKey(EntryKey key) external pure returns (uint64 srcChainId, uint64 dbNonce) {
        return InterchainEntryLib.decodeEntryKey(key);
    }
}
