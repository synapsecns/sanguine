// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

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

    /// @notice Returns the globally unique identifier of the batch
    function batchKey(InterchainBatch memory batch) internal pure returns (bytes32) {
        return keccak256(abi.encode(batch.srcChainId, batch.dbNonce));
    }
}
