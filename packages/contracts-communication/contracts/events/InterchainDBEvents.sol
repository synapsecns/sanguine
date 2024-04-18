// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

abstract contract InterchainDBEvents {
    /// @notice Emitted when a local entry is written to the database.
    /// @param dbNonce      The nonce of the batch containing the entry.
    /// @param entryIndex   The index of the entry within the batch.
    /// @param srcWriter    The address of the writer.
    /// @param dataHash     The written data hash.
    event InterchainEntryWritten(
        uint64 indexed dbNonce, uint64 entryIndex, bytes32 indexed srcWriter, bytes32 dataHash
    );

    /// @notice Emitted when a local batch is finalized.
    /// @param dbNonce      The nonce of the finalized batch.
    /// @param batchRoot    The Merkle root hash of the finalized batch.
    event InterchainBatchFinalized(uint64 indexed dbNonce, bytes32 batchRoot);

    /// @notice Emitted when a remote batch is verified by the Interchain Module.
    /// @param module       The address of the Interchain Module that verified the batch.
    /// @param srcChainId   The ID of the source chain.
    /// @param dbNonce      The nonce of the verified batch.
    /// @param batchRoot    The Merkle root hash of the verified batch.
    event InterchainBatchVerified(
        address indexed module, uint64 indexed srcChainId, uint64 indexed dbNonce, bytes32 batchRoot
    );

    /// @notice Emitted when a local batch is requested to be verified on a remote chain
    /// using the set of Interchain Modules.
    /// @param dstChainId   The ID of the destination chain.
    /// @param dbNonce      The nonce of the batch to be verified.
    /// @param batchRoot    The Merkle root hash of the batch to be verified.
    /// @param srcModules   The addresses of the Interchain Modules that will verify the batch.
    event InterchainBatchVerificationRequested(
        uint64 indexed dstChainId, uint64 indexed dbNonce, bytes32 batchRoot, address[] srcModules
    );
}
