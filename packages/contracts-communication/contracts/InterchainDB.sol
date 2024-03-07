// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {InterchainDBEvents} from "./events/InterchainDBEvents.sol";
import {IInterchainDB} from "./interfaces/IInterchainDB.sol";
import {IInterchainModule} from "./interfaces/IInterchainModule.sol";

import {InterchainBatch} from "./libs/InterchainBatch.sol";
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
        InterchainEntry memory entry = getEntry(dbNonce);
        _requestVerification(dstChainId, entry, srcModules);
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
        dbNonce = _writeEntry(dataHash);
        // TODO: entryIndex
        InterchainEntry memory entry = InterchainEntryLib.constructLocalEntry(dbNonce, 0, msg.sender, dataHash);
        _requestVerification(dstChainId, entry, srcModules);
    }

    // ═══════════════════════════════════════════════ MODULE-FACING ═══════════════════════════════════════════════════

    /// @inheritdoc IInterchainDB
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
    function getBatchLeafs(uint256 dbNonce) external view returns (bytes32[] memory) {}

    /// @inheritdoc IInterchainDB
    function getBatchLeafsPaginated(
        uint256 dbNonce,
        uint64 start,
        uint64 end
    )
        external
        view
        returns (bytes32[] memory)
    {
        // TODO: implement
    }

    /// @inheritdoc IInterchainDB
    function getBatchSize(uint256 dbNonce) external view returns (uint64) {
        // TODO: implement
    }

    /// @inheritdoc IInterchainDB
    function getEntryProof(uint256 dbNonce, uint64 entryIndex) external view returns (bytes32[] memory proof) {
        // TODO: implement
    }

    /// @inheritdoc IInterchainDB
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
        // TODO: implement
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
    function getBatch(uint256 dbNonce) public view returns (InterchainBatch memory) {
        // TODO: implement
    }

    /// @inheritdoc IInterchainDB
    function getEntry(uint256 dbNonce, uint64 entryIndex) public view returns (InterchainEntry memory) {
        if (getDBNonce() <= dbNonce) {
            revert InterchainDB__EntryDoesNotExist(dbNonce);
        }
        // TODO: entryIndex
        return InterchainEntryLib.constructLocalEntry(dbNonce, 0, _entries[dbNonce].writer, _entries[dbNonce].dataHash);
    }

    /// @inheritdoc IInterchainDB
    function getDBNonce() public view returns (uint256) {
        return _entries.length;
    }

    // ══════════════════════════════════════════════ INTERNAL LOGIC ═══════════════════════════════════════════════════

    /// @dev Write the entry to the database and emit the event.
    function _writeEntry(bytes32 dataHash) internal returns (uint256 dbNonce, uint64 entryIndex) {
        dbNonce = _entries.length;
        // TODO: entryIndex
        _entries.push(LocalEntry(msg.sender, dataHash));
        emit InterchainEntryWritten(block.chainid, dbNonce, TypeCasts.addressToBytes32(msg.sender), dataHash);
    }

    /// @dev Request the verification of the entry by the modules, and emit the event.
    /// Note: the validity of the passed entry and chain id being remote is enforced in the calling function.
    function _requestVerification(
        uint256 dstChainId,
        InterchainEntry memory entry,
        address[] calldata srcModules
    )
        internal
    {
        (uint256[] memory fees, uint256 totalFee) = _getModuleFees(dstChainId, entry.dbNonce, srcModules);
        if (msg.value != totalFee) {
            revert InterchainDB__IncorrectFeeAmount(msg.value, totalFee);
        }
        uint256 len = srcModules.length;
        for (uint256 i = 0; i < len; ++i) {
            // TODO: proper requestBatchVerification
            IInterchainModule(srcModules[i]).requestBatchVerification{value: fees[i]}(
                dstChainId, InterchainBatch({srcChainId: block.chainid, dbNonce: entry.dbNonce, batchRoot: 0})
            );
        }
        emit InterchainVerificationRequested(dstChainId, entry.dbNonce, srcModules);
    }

    // ══════════════════════════════════════════════ INTERNAL VIEWS ═══════════════════════════════════════════════════

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
