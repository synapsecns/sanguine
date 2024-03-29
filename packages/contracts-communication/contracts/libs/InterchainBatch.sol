// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {VersionedPayloadLib} from "./VersionedPayload.sol";

/// @notice Struct representing a batch of entries in the Interchain DataBase.
/// Batched entries are put together in a Merkle tree, which root is saved.
/// Batch has a globally unique identifier (key) and a value.
/// - key: srcChainId + dbNonce
/// - value: batchRoot
/// @param srcChainId   The chain id of the source chain
/// @param dbNonce      The database nonce of the batch
/// @param batchRoot    The root of the Merkle tree containing the batched entries
struct InterchainBatch {
    // TODO: can we use uint64 for chain id?
    uint256 srcChainId;
    uint256 dbNonce;
    bytes32 batchRoot;
}

library InterchainBatchLib {
    using VersionedPayloadLib for bytes;

    /// @notice Constructs an InterchainBatch struct to be saved on the local chain.
    /// @param dbNonce      The database nonce of the batch
    /// @param batchRoot    The root of the Merkle tree containing the batched entries
    /// @return batch       The constructed InterchainBatch struct
    function constructLocalBatch(
        uint256 dbNonce,
        bytes32 batchRoot
    )
        internal
        view
        returns (InterchainBatch memory batch)
    {
        return InterchainBatch({srcChainId: block.chainid, dbNonce: dbNonce, batchRoot: batchRoot});
    }

    /// @notice Decodes the versioned batch payload from memory into version and InterchainBatch struct.
    /// @dev See `VersionedPayloadLib` for more details about calldata/memory locations.
    /// @param versionedBatch   The versioned batch payload
    /// @return dbVersion       The version of the InterchainDB contract that created the batch
    /// @return batch           The InterchainBatch struct
    function decodeVersionedBatchFromMemory(bytes memory versionedBatch)
        internal
        view
        returns (uint16 dbVersion, InterchainBatch memory batch)
    {
        dbVersion = versionedBatch.getVersionFromMemory();
        batch = abi.decode(versionedBatch.getPayloadFromMemory(), (InterchainBatch));
    }

    /// @notice Decodes the versioned batch payload into version and InterchainBatch struct.
    /// @param versionedBatch   The versioned batch payload
    /// @return dbVersion       The version of the InterchainDB contract that created the batch
    /// @return batch           The InterchainBatch struct
    function decodeVersionedBatch(bytes calldata versionedBatch)
        internal
        pure
        returns (uint16 dbVersion, InterchainBatch memory batch)
    {
        dbVersion = versionedBatch.getVersion();
        batch = abi.decode(versionedBatch.getPayload(), (InterchainBatch));
    }

    /// @notice Encodes the InterchainBatch struct into a versioned batch payload.
    /// @param dbVersion        The version of the InterchainDB contract that created the batch
    /// @param batch            The InterchainBatch struct
    /// @return versionedBatch  The versioned batch payload
    function encodeVersionedBatch(
        uint16 dbVersion,
        InterchainBatch memory batch
    )
        internal
        pure
        returns (bytes memory versionedBatch)
    {
        versionedBatch = VersionedPayloadLib.encodeVersionedPayload(dbVersion, abi.encode(batch));
    }

    /// @notice Returns the globally unique identifier of the batch
    function batchKey(InterchainBatch memory batch) internal pure returns (bytes32) {
        return keccak256(abi.encode(batch.srcChainId, batch.dbNonce));
    }
}
