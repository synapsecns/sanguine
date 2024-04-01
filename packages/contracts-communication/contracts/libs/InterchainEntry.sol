// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

import {TypeCasts} from "./TypeCasts.sol";

import {SafeCast} from "@openzeppelin/contracts/utils/math/SafeCast.sol";

/// @notice Struct representing an entry in the Interchain DataBase.
/// Entry has a globally unique identifier (key) and a value.
/// - key: srcChainId + dbNonce + entryIndex
/// - value: srcWriter + dataHash
/// @param srcChainId   The chain id of the source chain
/// @param dbNonce      The database nonce of the batch containing the entry
/// @param entryIndex   The index of the entry in the batch
/// @param srcWriter    The address of the writer on the source chain
/// @param dataHash     The hash of the data written on the source chain
struct InterchainEntry {
    // TODO: can we use uint64 for chain id?
    uint64 srcChainId;
    uint64 dbNonce;
    uint64 entryIndex;
    bytes32 srcWriter;
    bytes32 dataHash;
}

using InterchainEntryLib for InterchainEntry global;

library InterchainEntryLib {
    /// @notice Constructs an InterchainEntry struct to be written on the local chain
    /// @param dbNonce      The database nonce of the entry on the source chain
    /// @param writer       The address of the writer on the local chain
    /// @param dataHash     The hash of the data written on the local chain
    /// @return entry       The constructed InterchainEntry struct
    function constructLocalEntry(
        uint64 dbNonce,
        uint64 entryIndex,
        address writer,
        bytes32 dataHash
    )
        internal
        view
        returns (InterchainEntry memory entry)
    {
        return InterchainEntry({
            srcChainId: SafeCast.toUint64(block.chainid),
            dbNonce: dbNonce,
            entryIndex: entryIndex,
            srcWriter: TypeCasts.addressToBytes32(writer),
            dataHash: dataHash
        });
    }

    /// @notice Returns the globally unique identifier of the entry
    function entryKey(InterchainEntry memory entry) internal pure returns (bytes32) {
        return keccak256(abi.encode(entry.srcChainId, entry.dbNonce, entry.entryIndex));
    }

    /// @notice Returns the value of the entry: writer + dataHash hashed together
    function entryValue(InterchainEntry memory entry) internal pure returns (bytes32) {
        return keccak256(abi.encode(entry.srcWriter, entry.dataHash));
    }

    /// @notice Returns the globally unique identifier of the batch containing the entry
    function batchKey(InterchainEntry memory entry) internal pure returns (bytes32) {
        return keccak256(abi.encode(entry.srcChainId, entry.dbNonce));
    }
}
