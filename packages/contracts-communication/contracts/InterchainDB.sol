// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {InterchainDBEvents} from "./events/InterchainDBEvents.sol";
import {IInterchainDB} from "./interfaces/IInterchainDB.sol";
import {IInterchainModule} from "./interfaces/IInterchainModule.sol";

import {BatchingV1Lib} from "./libs/BatchingV1.sol";
import {InterchainBatch, InterchainBatchLib, BatchKey} from "./libs/InterchainBatch.sol";
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
            revert InterchainDB__SameChainId(chainId);
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
        uint16 dbVersion = versionedBatch.getVersion();
        if (dbVersion != DB_VERSION) {
            revert InterchainDB__InvalidBatchVersion(dbVersion);
        }
        InterchainBatch memory batch = InterchainBatchLib.decodeBatch(versionedBatch.getPayload());
        if (batch.srcChainId == block.chainid) {
            revert InterchainDB__SameChainId(batch.srcChainId);
        }
        BatchKey batchKey = InterchainBatchLib.encodeBatchKey({srcChainId: batch.srcChainId, dbNonce: batch.dbNonce});
        RemoteBatch memory existingBatch = _remoteBatches[msg.sender][batchKey];
        // Check if that's the first time module verifies the batch
        if (existingBatch.verifiedAt == 0) {
            _saveVerifiedBatch(msg.sender, batchKey, batch);
        } else {
            // If the module has already verified the batch, check that the batch root is the same
            if (existingBatch.batchRoot != batch.batchRoot) {
                revert InterchainDB__ConflictingBatches(msg.sender, existingBatch.batchRoot, batch);
            }
            // No-op if the batch root is the same
        }
    }

    // ═══════════════════════════════════════════════════ VIEWS ═══════════════════════════════════════════════════════

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
        if (start != 0 || end != 1) {
            revert InterchainDB__InvalidEntryRange(dbNonce, start, end);
        }
        return getBatchLeafs(dbNonce);
    }

    /// @inheritdoc IInterchainDB
    function getEntryProof(uint64 dbNonce, uint64 entryIndex) external view returns (bytes32[] memory proof) {
        // In "no batching" mode: the batch root is the same as the entry value, hence the proof is empty
        _assertBatchFinalized(dbNonce);
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
            return InterchainBatchLib.UNVERIFIED;
        }
        // Check if the batch root matches the one verified by the module
        return remoteBatch.batchRoot == batch.batchRoot ? remoteBatch.verifiedAt : InterchainBatchLib.CONFLICT;
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
    function getBatchLeafs(uint64 dbNonce) public view returns (bytes32[] memory leafs) {
        // In "no batching" mode: the finalized batch size is 1
        _assertBatchFinalized(dbNonce);
        leafs = new bytes32[](1);
        leafs[0] = getEntryValue(dbNonce, 0);
    }

    /// @inheritdoc IInterchainDB
    function getBatchSize(uint64 dbNonce) public view returns (uint64) {
        // In "no batching" mode: the finalized batch size is 1, the pending batch size is 0
        uint64 pendingNonce = _assertBatchExists(dbNonce);
        return dbNonce < pendingNonce ? 1 : 0;
    }

    /// @inheritdoc IInterchainDB
    function getBatch(uint64 dbNonce) public view returns (InterchainBatch memory) {
        _assertBatchFinalized(dbNonce);
        // In "no batching" mode: the batch root is the same as the entry hash
        return InterchainBatchLib.constructLocalBatch(dbNonce, getEntryValue(dbNonce, 0));
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
            revert InterchainDB__IncorrectFeeAmount(msg.value, totalFee);
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
            IInterchainModule(srcModules[i]).requestBatchVerification{value: fees[i]}(dstChainId, versionedBatch);
        }
        emit InterchainBatchVerificationRequested(dstChainId, batch.dbNonce, batch.batchRoot, srcModules);
    }

    /// @dev Save the verified batch to the database and emit the event.
    function _saveVerifiedBatch(address module, BatchKey batchKey, InterchainBatch memory batch) internal {
        _remoteBatches[module][batchKey] = RemoteBatch({verifiedAt: block.timestamp, batchRoot: batch.batchRoot});
        emit InterchainBatchVerified(module, batch.srcChainId, batch.dbNonce, batch.batchRoot);
    }

    // ══════════════════════════════════════════════ INTERNAL VIEWS ═══════════════════════════════════════════════════

    /// @dev Check that the batch with the given nonce exists and return the pending nonce.
    function _assertBatchExists(uint64 dbNonce) internal view returns (uint64 pendingNonce) {
        pendingNonce = getDBNonce();
        if (dbNonce > pendingNonce) {
            revert InterchainDB__BatchDoesNotExist(dbNonce);
        }
    }

    /// @dev Check that the batch with the given nonce is finalized and return the pending nonce.
    function _assertBatchFinalized(uint64 dbNonce) internal view returns (uint64 pendingNonce) {
        pendingNonce = getDBNonce();
        if (dbNonce >= pendingNonce) {
            revert InterchainDB__BatchNotFinalized(dbNonce);
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
            revert InterchainDB__NoModulesSpecified();
        }
        fees = new uint256[](len);
        for (uint256 i = 0; i < len; ++i) {
            fees[i] = IInterchainModule(srcModules[i]).getModuleFee(dstChainId, dbNonce);
            totalFee += fees[i];
        }
    }
}
