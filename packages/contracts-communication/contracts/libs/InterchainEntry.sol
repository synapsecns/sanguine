// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {TypeCasts} from "./TypeCasts.sol";

/// @notice Struct representing an entry in the Interchain DataBase.
/// Entry has a globally unique identifier (key) and a value.
/// - key: srcChainId + dbNonce
/// - value: srcWriter + dataHash
/// @param srcChainId   The chain id of the source chain
/// @param dbNonce      The database nonce of the entry on the source chain
/// @param srcWriter    The address of the writer on the source chain
/// @param dataHash     The hash of the data written on the source chain
struct InterchainEntry {
    uint256 srcChainId;
    uint256 dbNonce;
    bytes32 srcWriter;
    bytes32 dataHash;
}

library InterchainEntryLib {
    /// @notice Constructs an InterchainEntry struct to be written on the local chain
    /// @param dbNonce      The database nonce of the entry on the source chain
    /// @param srcWriter    The address of the writer on the local chain
    /// @param dataHash     The hash of the data written on the local chain
    /// @return entry       The constructed InterchainEntry struct
    function constructLocalEntry(
        uint256 dbNonce,
        address srcWriter,
        bytes32 dataHash
    )
        internal
        view
        returns (InterchainEntry memory entry)
    {
        return InterchainEntry({
            srcChainId: block.chainid,
            dbNonce: dbNonce,
            srcWriter: TypeCasts.addressToBytes32(srcWriter),
            dataHash: dataHash
        });
    }

    /// @notice Returns the globally unique identifier of the entry
    function entryKey(InterchainEntry memory entry) internal pure returns (bytes32) {
        return keccak256(abi.encode(entry.srcChainId, entry.dbNonce));
    }

    /// @notice Returns the value of the entry: writer + dataHash hashed together
    function entryValue(InterchainEntry memory entry) internal pure returns (bytes32) {
        return keccak256(abi.encode(entry.srcWriter, entry.dataHash));
    }
}
