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

    uint16 public constant DB_VERSION = 1;

    bytes32[] internal _entryValues;
    mapping(address module => mapping(BatchKey batchKey => RemoteBatch batch)) internal _remoteBatches;

    modifier onlyRemoteChainId(uint64 chainId) {
        if (chainId == block.chainid) {
            revert InterchainDB__ChainIdNotRemote(chainId);
        }
        _;
    }

    // ═══════════════════════════════════════════════ WRITER-FACING ═══════════════════════════════════════════════════

    /// @inheritdoc IInterchainDB
    function writeEntry(bytes32 dataHash) external returns (uint64 dbNonce, uint64 entryIndex) {
        InterchainEntry memory entry = _writeEntry(dataHash);
        (dbNonce, entryIndex) = (entry.dbNonce, entry.entryIndex);
    }

    /// @inheritdoc IInterchainDB
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

    /// @inheritdoc IInterchainDB
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

    /// @inheritdoc IInterchainDB
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

    /// @inheritdoc IInterchainDB
    function getBatchLeafs(uint64 dbNonce) external view returns (bytes32[] memory leafs) {
        uint256 batchSize = getBatchSize(dbNonce);
        leafs = new bytes32[](batchSize);
        for (uint64 i = 0; i < batchSize; ++i) {
            leafs[i] = getEntryValue(dbNonce, i);
        }
    }

    /// @inheritdoc IInterchainDB
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

    /// @inheritdoc IInterchainDB
    function getEntryProof(uint64 dbNonce, uint64 entryIndex) external view returns (bytes32[] memory proof) {
        // In "no batching" mode: the batch root is the same as the entry value, hence the proof is empty
        _assertEntryExists(dbNonce, entryIndex);
        return new bytes32[](0);
    }

    /// @inheritdoc IInterchainDB
    function getInterchainFee(uint64 dstChainId, address[] calldata srcModules) external view returns (uint256 fee) {
        (, fee) = _getModuleFees(dstChainId, getDBNonce(), srcModules);
    }

    /// @inheritdoc IInterchainDB
    function getNextEntryIndex() external view returns (uint64 dbNonce, uint64 entryIndex) {
        // In "no batching" mode: entry index is 0, batch size is 1
        dbNonce = getDBNonce();
        entryIndex = 0;
    }

    /// @inheritdoc IInterchainDB
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

    /// @inheritdoc IInterchainDB
    function getVersionedBatch(uint64 dbNonce) external view returns (bytes memory versionedBatch) {
        InterchainBatch memory batch = getBatch(dbNonce);
        return VersionedPayloadLib.encodeVersionedPayload({
            version: DB_VERSION,
            payload: InterchainBatchLib.encodeBatch(batch)
        });
    }

    /// @inheritdoc IInterchainDB
    function getBatchRoot(InterchainEntry memory entry, bytes32[] calldata proof) external pure returns (bytes32) {
        return BatchingV1Lib.getBatchRoot({
            srcWriter: entry.srcWriter,
            dataHash: entry.dataHash,
            entryIndex: entry.entryIndex,
            proof: proof
        });
    }

    /// @inheritdoc IInterchainDB
    function getBatchSize(uint64 dbNonce) public view returns (uint64) {
        // In "no batching" mode: the finalized batch size is 1, the pending batch size is 0
        // We also return 0 for non-existent batches
        return dbNonce < getDBNonce() ? 1 : 0;
    }

    /// @inheritdoc IInterchainDB
    function getBatch(uint64 dbNonce) public view returns (InterchainBatch memory) {
        // In "no batching" mode: the batch root is the same as the entry hash.
        // For non-finalized or non-existent batches, the batch root is 0.
        bytes32 batchRoot = dbNonce < getDBNonce() ? getEntryValue(dbNonce, 0) : bytes32(0);
        return InterchainBatchLib.constructLocalBatch(dbNonce, batchRoot);
    }

    /// @inheritdoc IInterchainDB
    function getEntryValue(uint64 dbNonce, uint64 entryIndex) public view returns (bytes32) {
        _assertEntryExists(dbNonce, entryIndex);
        return _entryValues[dbNonce];
    }

    /// @inheritdoc IInterchainDB
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
