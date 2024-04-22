// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {InterchainBatch} from "../libs/InterchainBatch.sol";
import {InterchainEntry} from "../libs/InterchainEntry.sol";

interface IInterchainDB {
    error InterchainDB__BatchDoesNotExist(uint64 dbNonce);
    error InterchainDB__BatchNotFinalized(uint64 dbNonce);
    error InterchainDB__ConflictingBatches(address module, bytes32 existingBatchRoot, InterchainBatch newBatch);
    error InterchainDB__EntryIndexOutOfRange(uint64 dbNonce, uint64 entryIndex, uint64 batchSize);
    error InterchainDB__IncorrectFeeAmount(uint256 actualFee, uint256 expectedFee);
    error InterchainDB__InvalidBatchVersion(uint16 version);
    error InterchainDB__InvalidEntryRange(uint64 dbNonce, uint64 start, uint64 end);
    error InterchainDB__NoModulesSpecified();
    error InterchainDB__SameChainId(uint64 chainId);

    /// @notice Write data to the Interchain DataBase as a new entry in the current batch.
    /// Note: there are no guarantees that this entry will be available for reading on any of the remote chains.
    /// Use `requestBatchVerification` to ensure that the entry is available for reading on the destination chain.
    /// @param dataHash     The hash of the data to be written to the Interchain DataBase as a new entry
    /// @return dbNonce     The database nonce of the batch containing the written entry
    /// @return entryIndex  The index of the written entry within the batch
    function writeEntry(bytes32 dataHash) external returns (uint64 dbNonce, uint64 entryIndex);

    /// @notice Request the given Interchain Modules to verify an existing batch.
    /// If the batch is not finalized, the module will verify it after finalization.
    /// For the finalized batch the batch root is already available, and the module can verify it immediately.
    /// Note: every module has a separate fee paid in the native gas token of the source chain,
    /// and `msg.value` must be equal to the sum of all fees.
    /// Note: this method is permissionless, and anyone can request verification for any batch.
    /// @dev Will revert if the batch with the given nonce does not exist.
    /// @param dstChainId    The chain id of the destination chain
    /// @param dbNonce       The database nonce of the existing batch
    /// @param srcModules    The source chain addresses of the Interchain Modules to use for verification
    function requestBatchVerification(
        uint64 dstChainId,
        uint64 dbNonce,
        address[] memory srcModules
    )
        external
        payable;

    /// @notice Write data to the Interchain DataBase as a new entry in the current batch.
    /// Then request the Interchain Modules to verify the batch containing the written entry on the destination chain.
    /// See `writeEntry` and `requestBatchVerification` for more details.
    /// @dev Will revert if the empty array of modules is provided.
    /// @param dstChainId   The chain id of the destination chain
    /// @param dataHash     The hash of the data to be written to the Interchain DataBase as a new entry
    /// @param srcModules   The source chain addresses of the Interchain Modules to use for verification
    /// @return dbNonce     The database nonce of the batch containing the written entry
    /// @return entryIndex  The index of the written entry within the batch
    function writeEntryWithVerification(
        uint64 dstChainId,
        bytes32 dataHash,
        address[] memory srcModules
    )
        external
        payable
        returns (uint64 dbNonce, uint64 entryIndex);

    /// @notice Allows the Interchain Module to verify the batch coming from the remote chain.
    /// Note: The DB will only accept the batch of the same version as the DB itself.
    /// @param versionedBatch   The versioned Interchain Batch to verify
    function verifyRemoteBatch(bytes memory versionedBatch) external;

    // ═══════════════════════════════════════════════════ VIEWS ═══════════════════════════════════════════════════════

    /// @notice Get the fee for writing data to the Interchain DataBase, and verifying it on the destination chain
    /// using the provided Interchain Modules.
    /// @dev Will revert if the empty array of modules is provided.
    /// @param dstChainId   The chain id of the destination chain
    /// @param srcModules   The source chain addresses of the Interchain Modules to use for verification
    function getInterchainFee(uint64 dstChainId, address[] memory srcModules) external view returns (uint256);

    /// @notice Returns the list of leafs of the finalized batch with the given nonce.
    /// Note: the leafs are ordered by the index of the written entry in the current batch,
    /// and the leafs value match the value of the written entry (srcWriter + dataHash hashed together).
    /// @dev Will revert if the batch with the given nonce does not exist, or is not finalized.
    /// @param dbNonce      The database nonce of the finalized batch
    function getBatchLeafs(uint64 dbNonce) external view returns (bytes32[] memory);

    /// @notice Returns the list of leafs of the finalized batch with the given nonce,
    /// paginated by the given start and end indexes. The end index is exclusive.
    /// Note: this is useful when the batch contains a large number of leafs, and calling `getBatchLeafs`
    /// would result in a gas limit exceeded error.
    /// @dev Will revert if the batch with the given nonce does not exist, or is not finalized.
    /// Will revert if the provided range is invalid.
    /// @param dbNonce      The database nonce of the finalized batch
    /// @param start        The start index of the paginated leafs, inclusive
    /// @param end          The end index of the paginated leafs, exclusive
    function getBatchLeafsPaginated(
        uint64 dbNonce,
        uint64 start,
        uint64 end
    )
        external
        view
        returns (bytes32[] memory);

    /// @notice Returns the size of the finalized batch with the given nonce.
    /// @dev Will revert if the batch with the given nonce does not exist, or is not finalized.
    /// @param dbNonce      The database nonce of the finalized batch
    function getBatchSize(uint64 dbNonce) external view returns (uint64);

    /// @notice Get the finalized Interchain Batch with the given nonce.
    /// @dev Will revert if the batch with the given nonce does not exist, or is not finalized.
    /// @param dbNonce      The database nonce of the finalized batch
    function getBatch(uint64 dbNonce) external view returns (InterchainBatch memory);

    /// @notice Get the Interchain Entry's value written on the local chain with the given batch nonce and entry index.
    /// Entry value is calculated as the hash of the writer address and the written data hash.
    /// Note: the batch does not have to be finalized to fetch the entry value.
    /// @dev Will revert if the batch with the given nonce does not exist,
    /// or the entry with the given index does not exist within the batch.
    /// @param dbNonce      The database nonce of the existing batch
    /// @param entryIndex   The index of the written entry within the batch
    function getEntryValue(uint64 dbNonce, uint64 entryIndex) external view returns (bytes32);

    /// @notice Get the Merkle proof of inclusion for the entry with the given index
    /// in the finalized batch with the given nonce.
    /// @dev Will revert if the batch with the given nonce does not exist, or is not finalized.
    /// Will revert if the entry with the given index does not exist within the batch.
    /// @param dbNonce      The database nonce of the finalized batch
    /// @param entryIndex   The index of the written entry within the batch
    /// @return proof       The Merkle proof of inclusion for the entry
    function getEntryProof(uint64 dbNonce, uint64 entryIndex) external view returns (bytes32[] memory proof);

    /// @notice Get the nonce of the database, which is incremented every time a new batch is finalized.
    /// This is the nonce of the current non-finalized batch.
    function getDBNonce() external view returns (uint64);

    /// @notice Get the index of the next entry to be written to the database.
    /// @return dbNonce      The database nonce of the batch including the next entry
    /// @return entryIndex   The index of the next entry within that batch
    function getNextEntryIndex() external view returns (uint64 dbNonce, uint64 entryIndex);

    /// @notice Check if the batch is verified by the Interchain Module on the destination chain.
    /// Note: returned zero value indicates that the module has not verified the batch.
    /// @param dstModule    The destination chain addresses of the Interchain Modules to use for verification
    /// @param batch        The Interchain Batch to check
    /// @return moduleVerifiedAt    The block timestamp at which the batch was verified by the module,
    ///                             or ZERO if the module has not verified the batch.
    function checkBatchVerification(
        address dstModule,
        InterchainBatch memory batch
    )
        external
        view
        returns (uint256 moduleVerifiedAt);

    /// @notice Get the batch root containing the Interchain Entry with the given index.
    /// @param entry         The Interchain Entry to get the batch root for
    /// @param proof         The Merkle proof of inclusion for the entry in the batch
    function getBatchRoot(InterchainEntry memory entry, bytes32[] memory proof) external pure returns (bytes32);

    /// @notice Get the version of the Interchain DataBase.
    // solhint-disable-next-line func-name-mixedcase
    function DB_VERSION() external pure returns (uint16);
}
