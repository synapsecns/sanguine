// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {InterchainDBEvents} from "./events/InterchainDBEvents.sol";
import {IInterchainDB} from "./interfaces/IInterchainDB.sol";
import {IInterchainModule} from "./interfaces/IInterchainModule.sol";

import {BatchingV1Lib} from "./libs/BatchingV1.sol";
import {
    InterchainBatch, InterchainBatchLib, BatchKey, BATCH_UNVERIFIED, BATCH_CONFLICT
} from "./libs/InterchainBatch.sol";
import {InterchainEntry, InterchainEntryLib} from "./libs/InterchainEntry.sol";
import {VersionedPayloadLib} from "./libs/VersionedPayload.sol";
import {TypeCasts} from "./libs/TypeCasts.sol";

contract InterchainDB is InterchainDBEvents, IInterchainDB {
    using VersionedPayloadLib for bytes;

    /// @notice Struct representing a batch of entries from the remote Interchain DataBase,
    /// verified by the Interchain Module.
    /// @param verifiedAt   The block timestamp at which the entry was verified by the module
    /// @param batchRoot    The Merkle root of the batch
    struct RemoteBatch {
        uint256 verifiedAt;
        bytes32 batchRoot;
    }

    /// @notice The version of the Interchain DataBase. Must match the version of the batches.
    uint16 public constant DB_VERSION = 1;

    /// @dev The entry values written on the local chain.
    bytes32[] internal _entryValues;
    /// @dev The remote batches verified by the modules.
    mapping(address module => mapping(BatchKey batchKey => RemoteBatch batch)) internal _remoteBatches;

    /// @dev Checks if the chain id is not the same as the local chain id.
    modifier onlyRemoteChainId(uint64 chainId) {
        if (chainId == block.chainid) {
            revert InterchainDB__ChainIdNotRemote(chainId);
        }
        _;
    }

    // ═══════════════════════════════════════════════ WRITER-FACING ═══════════════════════════════════════════════════

    /// @notice Write data to the Interchain DataBase as a new entry in the current batch.
    /// Note: there are no guarantees that this entry will be available for reading on any of the remote chains.
    /// Use `requestBatchVerification` to ensure that the entry is available for reading on the destination chain.
    /// @param dataHash     The hash of the data to be written to the Interchain DataBase as a new entry
    /// @return dbNonce     The database nonce of the batch containing the written entry
    /// @return entryIndex  The index of the written entry within the batch
    function writeEntry(bytes32 dataHash) external returns (uint64 dbNonce, uint64 entryIndex) {
        InterchainEntry memory entry = _writeEntry(dataHash);
        (dbNonce, entryIndex) = (entry.dbNonce, entry.entryIndex);
    }

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
        address[] calldata srcModules
    )
        external
        payable
        onlyRemoteChainId(dstChainId)
    {
        InterchainBatch memory batch = getBatch(dbNonce);
        _requestVerification(dstChainId, batch, srcModules);
    }

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
        address[] calldata srcModules
    )
        external
        payable
        onlyRemoteChainId(dstChainId)
        returns (uint64 dbNonce, uint64 entryIndex)
    {
        InterchainEntry memory entry = _writeEntry(dataHash);
        (dbNonce, entryIndex) = (entry.dbNonce, entry.entryIndex);
        // In "no batching" mode: the batch root is the same as the entry value
        InterchainBatch memory batch = InterchainBatchLib.constructLocalBatch(dbNonce, entry.entryValue());
        _requestVerification(dstChainId, batch, srcModules);
    }

    // ═══════════════════════════════════════════════ MODULE-FACING ═══════════════════════════════════════════════════

    /// @notice Allows the Interchain Module to verify the batch coming from the remote chain.
    /// The module SHOULD verify the exact finalized batch from the remote chain. If the batch with a given nonce
    /// is not finalized or does not exist, module CAN verify it with an empty root value. Once the batch is
    /// finalized, the module SHOULD re-verify the batch with the correct root value.
    /// Note: The DB will only accept the batch of the same version as the DB itself.
    /// @dev Will revert if the batch with the same nonce but a different non-empty root is already verified.
    /// @param versionedBatch   The versioned Interchain Batch to verify
    function verifyRemoteBatch(bytes calldata versionedBatch) external {
        InterchainBatch memory batch = _assertCorrectBatch(versionedBatch);
        BatchKey batchKey = InterchainBatchLib.encodeBatchKey({srcChainId: batch.srcChainId, dbNonce: batch.dbNonce});
        RemoteBatch memory existingBatch = _remoteBatches[msg.sender][batchKey];
        // Check if that's the first time module verifies the batch
        if (existingBatch.verifiedAt == 0) {
            _saveVerifiedBatch(msg.sender, batchKey, batch);
            return;
        }
        // No-op if the batch root is the same
        if (existingBatch.batchRoot == batch.batchRoot) {
            return;
        }
        // Overwriting an empty (non-existent) batch with a different one is allowed
        if (existingBatch.batchRoot == 0) {
            _saveVerifiedBatch(msg.sender, batchKey, batch);
            return;
        }
        // Overwriting an existing batch with a different one is not allowed
        revert InterchainDB__BatchConflict(msg.sender, existingBatch.batchRoot, batch);
    }

    // ═══════════════════════════════════════════════════ VIEWS ═══════════════════════════════════════════════════════

    /// @notice Returns the list of leafs of the finalized batch with the given nonce.
    /// Note: the leafs are ordered by the index of the written entry in the current batch,
    /// and the leafs value match the value of the written entry (srcWriter + dataHash hashed together).
    /// @dev Will revert if the batch with the given nonce does not exist, or is not finalized.
    /// @param dbNonce      The database nonce of the finalized batch
    function getBatchLeafs(uint64 dbNonce) external view returns (bytes32[] memory leafs) {
        uint256 batchSize = getBatchSize(dbNonce);
        leafs = new bytes32[](batchSize);
        for (uint64 i = 0; i < batchSize; ++i) {
            leafs[i] = getEntryValue(dbNonce, i);
        }
    }

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
        returns (bytes32[] memory leafs)
    {
        uint256 size = getBatchSize(dbNonce);
        if (start > end || end > size) {
            revert InterchainDB__EntryRangeInvalid(dbNonce, start, end);
        }
        leafs = new bytes32[](end - start);
        for (uint64 i = start; i < end; ++i) {
            leafs[i - start] = getEntryValue(dbNonce, i);
        }
    }

    /// @notice Get the Merkle proof of inclusion for the entry with the given index
    /// in the finalized batch with the given nonce.
    /// @dev Will revert if the batch with the given nonce does not exist, or is not finalized.
    /// Will revert if the entry with the given index does not exist within the batch.
    /// @param dbNonce      The database nonce of the finalized batch
    /// @param entryIndex   The index of the written entry within the batch
    /// @return proof       The Merkle proof of inclusion for the entry
    function getEntryProof(uint64 dbNonce, uint64 entryIndex) external view returns (bytes32[] memory proof) {
        // In "no batching" mode: the batch root is the same as the entry value, hence the proof is empty
        _assertEntryExists(dbNonce, entryIndex);
        return new bytes32[](0);
    }

    /// @notice Get the fee for writing data to the Interchain DataBase, and verifying it on the destination chain
    /// using the provided Interchain Modules.
    /// @dev Will revert if the empty array of modules is provided.
    /// @param dstChainId   The chain id of the destination chain
    /// @param srcModules   The source chain addresses of the Interchain Modules to use for verification
    function getInterchainFee(uint64 dstChainId, address[] calldata srcModules) external view returns (uint256 fee) {
        (, fee) = _getModuleFees(dstChainId, getDBNonce(), srcModules);
    }

    /// @notice Get the index of the next entry to be written to the database.
    /// @return dbNonce      The database nonce of the batch including the next entry
    /// @return entryIndex   The index of the next entry within that batch
    function getNextEntryIndex() external view returns (uint64 dbNonce, uint64 entryIndex) {
        // In "no batching" mode: entry index is 0, batch size is 1
        dbNonce = getDBNonce();
        entryIndex = 0;
    }

    /// @notice Check if the batch is verified by the Interchain Module on the destination chain.
    /// - returned value `BATCH_UNVERIFIED` indicates that the module has not verified the batch.
    /// - returned value `BATCH_CONFLICT` indicates that the module has verified a different batch root
    /// for the same batch key.
    /// @param dstModule    The destination chain addresses of the Interchain Modules to use for verification
    /// @param batch        The Interchain Batch to check
    /// @return moduleVerifiedAt    The block timestamp at which the batch was verified by the module,
    /// BATCH_UNVERIFIED if the module has not verified the batch,
    /// BATCH_CONFLICT if the module has verified a different batch root for the same batch key.
    function checkBatchVerification(
        address dstModule,
        InterchainBatch memory batch
    )
        external
        view
        onlyRemoteChainId(batch.srcChainId)
        returns (uint256 moduleVerifiedAt)
    {
        BatchKey batchKey = InterchainBatchLib.encodeBatchKey({srcChainId: batch.srcChainId, dbNonce: batch.dbNonce});
        RemoteBatch memory remoteBatch = _remoteBatches[dstModule][batchKey];
        // Check if module verified anything for this batch key first
        if (remoteBatch.verifiedAt == 0) {
            return BATCH_UNVERIFIED;
        }
        // Check if the batch root matches the one verified by the module
        return remoteBatch.batchRoot == batch.batchRoot ? remoteBatch.verifiedAt : BATCH_CONFLICT;
    }

    /// @notice Get the versioned Interchain Batch with the given nonce.
    /// Note: will return a batch with an empty root if the batch does not exist, or is not finalized.
    /// @param dbNonce      The database nonce of the batch
    function getVersionedBatch(uint64 dbNonce) external view returns (bytes memory versionedBatch) {
        InterchainBatch memory batch = getBatch(dbNonce);
        return VersionedPayloadLib.encodeVersionedPayload({
            version: DB_VERSION,
            payload: InterchainBatchLib.encodeBatch(batch)
        });
    }

    /// @notice Get the batch root containing the Interchain Entry with the given index.
    /// @param entry         The Interchain Entry to get the batch root for
    /// @param proof         The Merkle proof of inclusion for the entry in the batch
    function getBatchRoot(InterchainEntry memory entry, bytes32[] calldata proof) external pure returns (bytes32) {
        return BatchingV1Lib.getBatchRoot({
            srcWriter: entry.srcWriter,
            dataHash: entry.dataHash,
            entryIndex: entry.entryIndex,
            proof: proof
        });
    }

    /// @notice Returns the size of the finalized batch with the given nonce.
    /// @dev Will return 0 for non-existent or non-finalized batches.
    /// @param dbNonce      The database nonce of the finalized batch
    function getBatchSize(uint64 dbNonce) public view returns (uint64) {
        // In "no batching" mode: the finalized batch size is 1, the pending batch size is 0
        // We also return 0 for non-existent batches
        return dbNonce < getDBNonce() ? 1 : 0;
    }

    /// @notice Get the finalized Interchain Batch with the given nonce.
    /// @dev Will return a batch with an empty root if the batch does not exist, or is not finalized.
    /// @param dbNonce      The database nonce of the finalized batch
    function getBatch(uint64 dbNonce) public view returns (InterchainBatch memory) {
        // In "no batching" mode: the batch root is the same as the entry hash.
        // For non-finalized or non-existent batches, the batch root is 0.
        bytes32 batchRoot = dbNonce < getDBNonce() ? getEntryValue(dbNonce, 0) : bytes32(0);
        return InterchainBatchLib.constructLocalBatch(dbNonce, batchRoot);
    }

    /// @notice Get the Interchain Entry's value written on the local chain with the given batch nonce and entry index.
    /// Entry value is calculated as the hash of the writer address and the written data hash.
    /// Note: the batch does not have to be finalized to fetch the entry value.
    /// @dev Will revert if the batch with the given nonce does not exist,
    /// or the entry with the given index does not exist within the batch.
    /// @param dbNonce      The database nonce of the existing batch
    /// @param entryIndex   The index of the written entry within the batch
    function getEntryValue(uint64 dbNonce, uint64 entryIndex) public view returns (bytes32) {
        _assertEntryExists(dbNonce, entryIndex);
        return _entryValues[dbNonce];
    }

    /// @notice Get the nonce of the database, which is incremented every time a new batch is finalized.
    /// This is the nonce of the current non-finalized batch.
    function getDBNonce() public view returns (uint64) {
        // We can do the unsafe cast here as writing more than 2^64 entries is practically impossible
        return uint64(_entryValues.length);
    }

    // ══════════════════════════════════════════════ INTERNAL LOGIC ═══════════════════════════════════════════════════

    /// @dev Write the entry to the database and emit the event.
    function _writeEntry(bytes32 dataHash) internal returns (InterchainEntry memory entry) {
        uint64 dbNonce = getDBNonce();
        entry = InterchainEntryLib.constructLocalEntry({
            dbNonce: dbNonce,
            entryIndex: 0,
            writer: msg.sender,
            dataHash: dataHash
        });
        bytes32 entryValue = entry.entryValue();
        _entryValues.push(entryValue);
        emit InterchainEntryWritten({
            dbNonce: dbNonce,
            entryIndex: 0,
            srcWriter: TypeCasts.addressToBytes32(msg.sender),
            dataHash: dataHash
        });
        // In the InterchainDB V1 the batch is finalized immediately after the entry is written
        emit InterchainBatchFinalized({dbNonce: dbNonce, batchRoot: entryValue});
    }

    /// @dev Request the verification of the entry by the modules, and emit the event.
    /// Note: the validity of the passed entry and chain id being remote is enforced in the calling function.
    function _requestVerification(
        uint64 dstChainId,
        InterchainBatch memory batch,
        address[] calldata srcModules
    )
        internal
    {
        (uint256[] memory fees, uint256 totalFee) = _getModuleFees(dstChainId, batch.dbNonce, srcModules);
        if (msg.value < totalFee) {
            revert InterchainDB__FeeAmountBelowMin(msg.value, totalFee);
        } else if (msg.value > totalFee) {
            // The exceeding amount goes to the first module
            fees[0] += msg.value - totalFee;
        }
        uint256 len = srcModules.length;
        bytes memory versionedBatch = VersionedPayloadLib.encodeVersionedPayload({
            version: DB_VERSION,
            payload: InterchainBatchLib.encodeBatch(batch)
        });
        for (uint256 i = 0; i < len; ++i) {
            IInterchainModule(srcModules[i]).requestBatchVerification{value: fees[i]}(
                dstChainId, batch.dbNonce, versionedBatch
            );
        }
        emit InterchainBatchVerificationRequested(dstChainId, batch.dbNonce, batch.batchRoot, srcModules);
    }

    /// @dev Save the verified batch to the database and emit the event.
    function _saveVerifiedBatch(address module, BatchKey batchKey, InterchainBatch memory batch) internal {
        _remoteBatches[module][batchKey] = RemoteBatch({verifiedAt: block.timestamp, batchRoot: batch.batchRoot});
        emit InterchainBatchVerified(module, batch.srcChainId, batch.dbNonce, batch.batchRoot);
    }

    // ══════════════════════════════════════════════ INTERNAL VIEWS ═══════════════════════════════════════════════════

    /// @dev Asserts that the batch version is correct and that batch originates from a remote chain.
    /// Note: returns the decoded batch for chaining purposes.
    function _assertCorrectBatch(bytes calldata versionedBatch) internal view returns (InterchainBatch memory batch) {
        uint16 dbVersion = versionedBatch.getVersion();
        if (dbVersion != DB_VERSION) {
            revert InterchainDB__BatchVersionMismatch(dbVersion, DB_VERSION);
        }
        batch = InterchainBatchLib.decodeBatch(versionedBatch.getPayload());
        if (batch.srcChainId == block.chainid) {
            revert InterchainDB__ChainIdNotRemote(batch.srcChainId);
        }
    }

    /// @dev Check that the entry index is within the batch size. Also checks that the batch exists.
    function _assertEntryExists(uint64 dbNonce, uint64 entryIndex) internal view {
        // This will revert if the batch does not exist
        uint64 batchSize = getBatchSize(dbNonce);
        if (entryIndex >= batchSize) {
            revert InterchainDB__EntryIndexOutOfRange(dbNonce, entryIndex, batchSize);
        }
    }

    /// @dev Get the verification fees for the modules
    function _getModuleFees(
        uint64 dstChainId,
        uint64 dbNonce,
        address[] calldata srcModules
    )
        internal
        view
        returns (uint256[] memory fees, uint256 totalFee)
    {
        uint256 len = srcModules.length;
        if (len == 0) {
            revert InterchainDB__ModulesNotProvided();
        }
        fees = new uint256[](len);
        for (uint256 i = 0; i < len; ++i) {
            fees[i] = IInterchainModule(srcModules[i]).getModuleFee(dstChainId, dbNonce);
            totalFee += fees[i];
        }
    }
}
