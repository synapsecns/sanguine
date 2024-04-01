// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {VersionedPayloadLib} from "./VersionedPayload.sol";

import {SafeCast} from "@openzeppelin/contracts/utils/math/SafeCast.sol";

type BatchKey is uint128;

/// @notice Struct representing a batch of entries in the Interchain DataBase.
/// Batched entries are put together in a Merkle tree, which root is saved.
/// Batch has a globally unique identifier (key) and a value.
/// - key: srcChainId + dbNonce
/// - value: batchRoot
/// @param srcChainId   The chain id of the source chain
/// @param dbNonce      The database nonce of the batch
/// @param batchRoot    The root of the Merkle tree containing the batched entries
struct InterchainBatch {
    uint64 srcChainId;
    uint64 dbNonce;
    bytes32 batchRoot;
}

library InterchainBatchLib {
    using VersionedPayloadLib for bytes;

    /// @notice Constructs an InterchainBatch struct to be saved on the local chain.
    /// @param dbNonce      The database nonce of the batch
    /// @param batchRoot    The root of the Merkle tree containing the batched entries
    /// @return batch       The constructed InterchainBatch struct
    function constructLocalBatch(
        uint64 dbNonce,
        bytes32 batchRoot
    )
        internal
        view
        returns (InterchainBatch memory batch)
    {
        return InterchainBatch({srcChainId: SafeCast.toUint64(block.chainid), dbNonce: dbNonce, batchRoot: batchRoot});
    }

    /// @notice Encodes the InterchainBatch struct into a non-versioned batch payload.
    function encodeBatch(InterchainBatch memory batch) internal pure returns (bytes memory) {
        return abi.encode(encodeBatchKey(batch.srcChainId, batch.dbNonce), batch.batchRoot);
    }

    /// @notice Decodes the InterchainBatch struct from a non-versioned batch payload in calldata.
    function decodeBatch(bytes calldata data) internal pure returns (InterchainBatch memory batch) {
        BatchKey key;
        (key, batch.batchRoot) = abi.decode(data, (BatchKey, bytes32));
        (batch.srcChainId, batch.dbNonce) = decodeBatchKey(key);
    }

    /// @notice Decodes the InterchainBatch struct from a non-versioned batch payload in memory.
    function decodeBatchFromMemory(bytes memory data) internal pure returns (InterchainBatch memory batch) {
        BatchKey key;
        (key, batch.batchRoot) = abi.decode(data, (BatchKey, bytes32));
        (batch.srcChainId, batch.dbNonce) = decodeBatchKey(key);
    }

    /// @notice Encodes the uint128 key of the batch from uint64 srcChainId and uint64 dbNonce.
    function encodeBatchKey(uint64 srcChainId, uint64 dbNonce) internal pure returns (BatchKey) {
        return BatchKey.wrap((uint128(srcChainId) << 64) | dbNonce);
    }

    /// @notice Decodes the uint128 key of the batch into uint64 srcChainId and uint64 dbNonce.
    function decodeBatchKey(BatchKey key) internal pure returns (uint64 srcChainId, uint64 dbNonce) {
        srcChainId = uint64(BatchKey.unwrap(key) >> 64);
        dbNonce = uint64(BatchKey.unwrap(key));
    }
}
