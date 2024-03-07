// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {InterchainDBEvents} from "./events/InterchainDBEvents.sol";
import {IInterchainDB} from "./interfaces/IInterchainDB.sol";
import {IInterchainModule} from "./interfaces/IInterchainModule.sol";

import {InterchainBatch, InterchainBatchLib} from "./libs/InterchainBatch.sol";
import {InterchainEntry, InterchainEntryLib} from "./libs/InterchainEntry.sol";
import {TypeCasts} from "./libs/TypeCasts.sol";

contract InterchainDB is InterchainDBEvents, IInterchainDB {
    LocalEntry[] internal _entries;
    mapping(address module => mapping(bytes32 entryKey => RemoteEntry entry)) internal _remoteEntries;

    modifier onlyRemoteChainId(uint256 chainId) {
        if (chainId == block.chainid) {
            revert InterchainDB__SameChainId(block.chainid);
        }
        _;
    }

    // ═══════════════════════════════════════════════ WRITER-FACING ═══════════════════════════════════════════════════

    /// @inheritdoc IInterchainDB
    function writeEntry(bytes32 dataHash) external returns (uint256 dbNonce, uint64 entryIndex) {
        return _writeEntry(dataHash);
    }

    /// @inheritdoc IInterchainDB
    function requestBatchVerification(
        uint256 dstChainId,
        uint256 dbNonce,
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
        uint256 dstChainId,
        bytes32 dataHash,
        address[] calldata srcModules
    )
        external
        payable
        onlyRemoteChainId(dstChainId)
        returns (uint256 dbNonce, uint64 entryIndex)
    {
        (dbNonce, entryIndex) = _writeEntry(dataHash);
        // In "no batching" mode: the batch root is the same as the entry hash
        InterchainBatch memory batch = InterchainBatchLib.constructLocalBatch(dbNonce, dataHash);
        _requestVerification(dstChainId, batch, srcModules);
    }

    // ═══════════════════════════════════════════════ MODULE-FACING ═══════════════════════════════════════════════════

    function verifyEntry(InterchainEntry memory entry) external onlyRemoteChainId(entry.srcChainId) {
        // TODO: deprecated
        bytes32 entryKey = InterchainEntryLib.entryKey(entry);
        bytes32 entryValue = InterchainEntryLib.entryValue(entry);
        RemoteEntry memory existingEntry = _remoteEntries[msg.sender][entryKey];
        // Check if that's the first time module verifies the entry
        if (existingEntry.verifiedAt == 0) {
            _remoteEntries[msg.sender][entryKey] = RemoteEntry({verifiedAt: block.timestamp, entryValue: entryValue});
            emit InterchainEntryVerified(msg.sender, entry.srcChainId, entry.dbNonce, entry.srcWriter, entry.dataHash);
        } else {
            // If the module has already verified the entry, check that the entry value is the same
            if (existingEntry.entryValue != entryValue) {
                revert InterchainDB__ConflictingEntries(existingEntry.entryValue, entry);
            }
            // No-op if the entry value is the same
        }
    }

    /// @inheritdoc IInterchainDB
    function verifyRemoteBatch(InterchainBatch memory batch) external {
        // TODO: implement
    }

    // ═══════════════════════════════════════════════════ VIEWS ═══════════════════════════════════════════════════════

    /// @inheritdoc IInterchainDB
    function getBatchLeafs(uint256 dbNonce) external view returns (bytes32[] memory leafs) {
        // In "no batching" mode: the finalized batch size is 1
        _assertBatchFinalized(dbNonce);
        leafs = new bytes32[](1);
        leafs[0] = _entries[dbNonce].dataHash;
    }

    /// @inheritdoc IInterchainDB
    function getBatchLeafsPaginated(
        uint256 dbNonce,
        uint64 start,
        uint64 end
    )
        external
        view
        returns (bytes32[] memory leafs)
    {
        // In "no batching" mode: the finalized batch size is 1
        _assertBatchFinalized(dbNonce);
        if (start != 0 || end != 1) {
            revert InterchainDB__InvalidEntryRange(dbNonce, start, end);
        }
        leafs = new bytes32[](1);
        leafs[0] = _entries[dbNonce].dataHash;
    }

    /// @inheritdoc IInterchainDB
    function getEntryProof(uint256 dbNonce, uint64 entryIndex) external view returns (bytes32[] memory proof) {
        // In "no batching" mode: the batch root is the same as the entry hash, hence the proof is empty
        _assertBatchFinalized(dbNonce);
        _assertEntryExists(dbNonce, entryIndex);
        return new bytes32[](0);
    }

    function readEntry(
        address dstModule,
        InterchainEntry memory entry
    )
        external
        view
        onlyRemoteChainId(entry.srcChainId)
        returns (uint256 moduleVerifiedAt)
    {
        // TODO: deprecated
        RemoteEntry memory remoteEntry = _remoteEntries[dstModule][InterchainEntryLib.entryKey(entry)];
        bytes32 entryValue = InterchainEntryLib.entryValue(entry);
        // Check entry value against the one verified by the module
        return remoteEntry.entryValue == entryValue ? remoteEntry.verifiedAt : 0;
    }

    /// @inheritdoc IInterchainDB
    function getInterchainFee(uint256 dstChainId, address[] calldata srcModules) external view returns (uint256 fee) {
        (, fee) = _getModuleFees(dstChainId, getDBNonce(), srcModules);
    }

    /// @inheritdoc IInterchainDB
    function getNextEntryIndex() external view returns (uint256 dbNonce, uint64 entryIndex) {
        // In "no batching" mode: entry index is 0, batch size is 1
        dbNonce = getDBNonce();
        entryIndex = 0;
    }

    /// @inheritdoc IInterchainDB
    function checkVerification(
        address dstModule,
        InterchainEntry memory entry,
        bytes32[] memory proof
    )
        external
        view
        returns (uint256 moduleVerifiedAt)
    {
        // TODO: implement
    }

    /// @inheritdoc IInterchainDB
    function getBatchSize(uint256 dbNonce) public view returns (uint64) {
        // In "no batching" mode: the finalized batch size is 1, the pending batch size is 0
        uint256 pendingNonce = _assertBatchExists(dbNonce);
        return dbNonce < pendingNonce ? 1 : 0;
    }

    /// @inheritdoc IInterchainDB
    function getBatch(uint256 dbNonce) public view returns (InterchainBatch memory) {
        _assertBatchFinalized(dbNonce);
        // In "no batching" mode: the batch root is the same as the entry hash
        return InterchainBatchLib.constructLocalBatch(dbNonce, _entries[dbNonce].dataHash);
    }

    /// @inheritdoc IInterchainDB
    function getEntry(uint256 dbNonce, uint64 entryIndex) public view returns (InterchainEntry memory) {
        _assertEntryExists(dbNonce, entryIndex);
        return InterchainEntryLib.constructLocalEntry(
            dbNonce, entryIndex, _entries[dbNonce].writer, _entries[dbNonce].dataHash
        );
    }

    /// @inheritdoc IInterchainDB
    function getDBNonce() public view returns (uint256) {
        return _entries.length;
    }

    // ══════════════════════════════════════════════ INTERNAL LOGIC ═══════════════════════════════════════════════════

    /// @dev Write the entry to the database and emit the event.
    function _writeEntry(bytes32 dataHash) internal returns (uint256 dbNonce, uint64 entryIndex) {
        dbNonce = _entries.length;
        entryIndex = 0;
        _entries.push(LocalEntry(msg.sender, dataHash));
        emit InterchainEntryWritten(block.chainid, dbNonce, TypeCasts.addressToBytes32(msg.sender), dataHash);
    }

    /// @dev Request the verification of the entry by the modules, and emit the event.
    /// Note: the validity of the passed entry and chain id being remote is enforced in the calling function.
    function _requestVerification(
        uint256 dstChainId,
        InterchainBatch memory batch,
        address[] calldata srcModules
    )
        internal
    {
        (uint256[] memory fees, uint256 totalFee) = _getModuleFees(dstChainId, batch.dbNonce, srcModules);
        // TODO: handle the case where fees are overpaid
        if (msg.value != totalFee) {
            revert InterchainDB__IncorrectFeeAmount(msg.value, totalFee);
        }
        uint256 len = srcModules.length;
        for (uint256 i = 0; i < len; ++i) {
            IInterchainModule(srcModules[i]).requestBatchVerification{value: fees[i]}(dstChainId, batch);
        }
        emit InterchainBatchVerificationRequested(dstChainId, batch.dbNonce, batch.batchRoot, srcModules);
    }

    // ══════════════════════════════════════════════ INTERNAL VIEWS ═══════════════════════════════════════════════════

    /// @dev Check that the batch with the given nonce exists and return the pending nonce.
    function _assertBatchExists(uint256 dbNonce) internal view returns (uint256 pendingNonce) {
        pendingNonce = getDBNonce();
        if (dbNonce > pendingNonce) {
            revert InterchainDB__BatchDoesNotExist(dbNonce);
        }
    }

    /// @dev Check that the batch with the given nonce is finalized and return the pending nonce.
    function _assertBatchFinalized(uint256 dbNonce) internal view returns (uint256 pendingNonce) {
        pendingNonce = getDBNonce();
        if (dbNonce >= pendingNonce) {
            revert InterchainDB__BatchNotFinalized(dbNonce);
        }
    }

    /// @dev Check that the entry index is within the batch size. Also checks that the batch exists.
    function _assertEntryExists(uint256 dbNonce, uint64 entryIndex) internal view {
        // This will revert if the batch does not exist
        uint64 batchSize = getBatchSize(dbNonce);
        if (entryIndex >= batchSize) {
            revert InterchainDB__EntryIndexOutOfRange(dbNonce, entryIndex, batchSize);
        }
    }

    /// @dev Get the verification fees for the modules
    function _getModuleFees(
        uint256 dstChainId,
        uint256 dbNonce,
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
