// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

abstract contract InterchainModuleEvents {
    /// @notice Emitted when a batch verification on a remote chain is requested.
    /// @param dstChainId           The chain ID of the destination chain.
    /// @param batch                The encoded batch to be verified.
    /// @param ethSignedBatchHash   The digest of the batch (EIP-191 personal signed).
    event BatchVerificationRequested(uint64 indexed dstChainId, bytes batch, bytes32 ethSignedBatchHash);

    /// @notice Emitted when a batch from the remote chain is verified.
    /// @param srcChainId           The chain ID of the source chain.
    /// @param batch                The encoded batch that was verified.
    /// @param ethSignedBatchHash   The digest of the batch (EIP-191 personal signed).
    event BatchVerified(uint64 indexed srcChainId, bytes batch, bytes32 ethSignedBatchHash);

    /// @notice Emitted when an entry verification on a remote chain is requested.
    /// @param dstChainId           The chain ID of the destination chain.
    /// @param entry                The encoded entry to be verified.
    /// @param ethSignedEntryHash   The digest of the entry (EIP-191 personal signed).
    event EntryVerificationRequested(uint64 indexed dstChainId, bytes entry, bytes32 ethSignedEntryHash);

    /// @notice Emitted when an entry from the remote chain is verified.
    /// @param srcChainId           The chain ID of the source chain.
    /// @param entry                The encoded entry that was verified.
    /// @param ethSignedEntryHash   The digest of the entry (EIP-191 personal signed).
    event EntryVerified(uint64 indexed srcChainId, bytes entry, bytes32 ethSignedEntryHash);
}
