// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {TypeCasts} from "./TypeCasts.sol";

import {SafeCast} from "@openzeppelin/contracts/utils/math/SafeCast.sol";

/// @notice Struct representing an entry in the Interchain DataBase.
/// Entry has a globally unique identifier (key) and a value.
/// Assuming `srcWriter` written data `digest` on the source chain:
/// - key: (srcChainId, dbNonce)
/// - entryValue = keccak256(srcWriter, digest)
/// @param srcChainId   The chain id of the source chain
/// @param dbNonce      The database nonce of the entry
/// @param entryValue   The entry value
struct InterchainEntry {
    uint64 srcChainId;
    uint64 dbNonce;
    bytes32 entryValue;
}

type EntryKey is uint128;

/// @dev Signals that the module has not verified any entry with the given key.
uint256 constant ENTRY_UNVERIFIED = 0;
/// @dev Signals that the module has verified a conflicting entry with the given key.
uint256 constant ENTRY_CONFLICT = type(uint256).max;

library InterchainEntryLib {
    /// @notice Constructs an InterchainEntry struct to be written on the local chain
    /// @param dbNonce      The database nonce of the entry on the source chain
    /// @param entryValue   The value of the entry
    /// @return entry       The constructed InterchainEntry struct
    function constructLocalEntry(
        uint64 dbNonce,
        bytes32 entryValue
    )
        internal
        view
        returns (InterchainEntry memory entry)
    {
        return InterchainEntry({srcChainId: SafeCast.toUint64(block.chainid), dbNonce: dbNonce, entryValue: entryValue});
    }

    /// @notice Returns the value of the entry: writer + digest hashed together
    function getEntryValue(bytes32 srcWriter, bytes32 digest) internal pure returns (bytes32) {
        return keccak256(abi.encode(srcWriter, digest));
    }

    /// @notice Returns the value of the entry: writer + digest hashed together.
    /// Note: this is exposed for convenience to avoid typecasts prior to abi-encoding.
    function getEntryValue(address srcWriter, bytes32 digest) internal pure returns (bytes32) {
        return keccak256(abi.encode(srcWriter, digest));
    }

    /// @notice Encodes the InterchainEntry struct into a non-versioned entry payload.
    function encodeEntry(InterchainEntry memory entry) internal pure returns (bytes memory) {
        return abi.encode(encodeEntryKey(entry.srcChainId, entry.dbNonce), entry.entryValue);
    }

    /// @notice Decodes the InterchainEntry struct from a non-versioned entry payload in calldata.
    function decodeEntry(bytes calldata data) internal pure returns (InterchainEntry memory entry) {
        EntryKey key;
        (key, entry.entryValue) = abi.decode(data, (EntryKey, bytes32));
        (entry.srcChainId, entry.dbNonce) = decodeEntryKey(key);
    }

    /// @notice Decodes the InterchainEntry struct from a non-versioned entry payload in memory.
    function decodeEntryFromMemory(bytes memory data) internal pure returns (InterchainEntry memory entry) {
        EntryKey key;
        (key, entry.entryValue) = abi.decode(data, (EntryKey, bytes32));
        (entry.srcChainId, entry.dbNonce) = decodeEntryKey(key);
    }

    /// @notice Encodes the uint128 key of the entry from uint64 srcChainId and uint64 dbNonce.
    function encodeEntryKey(uint64 srcChainId, uint64 dbNonce) internal pure returns (EntryKey) {
        return EntryKey.wrap((uint128(srcChainId) << 64) | dbNonce);
    }

    /// @notice Decodes the uint128 key of the entry into uint64 srcChainId and uint64 dbNonce.
    function decodeEntryKey(EntryKey key) internal pure returns (uint64 srcChainId, uint64 dbNonce) {
        srcChainId = uint64(EntryKey.unwrap(key) >> 64);
        dbNonce = uint64(EntryKey.unwrap(key));
    }
}
