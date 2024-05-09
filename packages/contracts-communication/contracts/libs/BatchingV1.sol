// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {InterchainEntryLib} from "./InterchainEntry.sol";

library BatchingV1Lib {
    error BatchingV1__EntryIndexNotZero(uint64 entryIndex);
    error BatchingV1__ProofNotEmpty();

    /// @notice Get the batch root containing the Interchain Entry with the given index.
    /// @param srcWriter    The entry writer of the source chain
    /// @param dataHash     The hash of the data of the entry
    /// @param entryIndex   The index of the entry in the batch
    /// @param proof        The Merkle proof of inclusion for the entry in the batch
    /// @return batchRoot   The root of the batch containing the entry
    function getBatchRoot(
        bytes32 srcWriter,
        bytes32 dataHash,
        uint64 entryIndex,
        bytes32[] calldata proof
    )
        internal
        pure
        returns (bytes32 batchRoot)
    {
        // In "no batching" mode: entry index is 0, proof is empty
        if (entryIndex != 0) {
            revert BatchingV1__EntryIndexNotZero(entryIndex);
        }
        if (proof.length != 0) {
            revert BatchingV1__ProofNotEmpty();
        }
        // In "no batching" mode: the batch root is the same as the entry value
        return InterchainEntryLib.getEntryValue({srcWriter: srcWriter, dataHash: dataHash});
    }
}
