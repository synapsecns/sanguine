// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {TypeCasts} from "./TypeCasts.sol";

/// @notice Struct representing an entry in the Interchain DataBase
/// @param srcChainId   The chain id of the source chain
/// @param srcWriter    The address of the writer on the source chain
/// @param writerNonce  The nonce of the writer on the source chain
/// @param dataHash     The hash of the data written on the source chain
struct InterchainEntry {
    uint256 srcChainId;
    bytes32 srcWriter;
    uint256 writerNonce;
    bytes32 dataHash;
}

library InterchainEntryLib {
    /// @notice Constructs an InterchainEntry struct to be written on the local chain
    /// @param srcWriter    The address of the writer on the local chain
    /// @param writerNonce  The nonce of the writer on the local chain
    /// @param dataHash     The hash of the data written on the local chain
    /// @return entry       The constructed InterchainEntry struct
    function constructLocalEntry(
        address srcWriter,
        uint256 writerNonce,
        bytes32 dataHash
    )
        internal
        view
        returns (InterchainEntry memory entry)
    {
        return InterchainEntry({
            srcChainId: block.chainid,
            srcWriter: TypeCasts.addressToBytes32(srcWriter),
            writerNonce: writerNonce,
            dataHash: dataHash
        });
    }

    /// @notice Returns the globally unique identifier of the entry
    function entryId(InterchainEntry memory entry) internal pure returns (bytes32) {
        return keccak256(abi.encode(entry.srcChainId, entry.srcWriter, entry.writerNonce));
    }
}
