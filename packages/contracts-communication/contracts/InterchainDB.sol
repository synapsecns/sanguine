// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {InterchainDBEvents} from "./events/InterchainDBEvents.sol";
import {IInterchainDB} from "./interfaces/IInterchainDB.sol";
import {IInterchainModule} from "./interfaces/IInterchainModule.sol";

import {
    InterchainBatch, InterchainBatchLib, BatchKey, BATCH_UNVERIFIED, BATCH_CONFLICT
} from "./libs/InterchainBatch.sol";
import {
    InterchainEntry, InterchainEntryLib, EntryKey, ENTRY_UNVERIFIED, ENTRY_CONFLICT
} from "./libs/InterchainEntry.sol";
import {VersionedPayloadLib} from "./libs/VersionedPayload.sol";
import {TypeCasts} from "./libs/TypeCasts.sol";

contract InterchainDB is InterchainDBEvents, IInterchainDB {
    using VersionedPayloadLib for bytes;

    /// @notice Struct representing a batch of entries from the remote Interchain DataBase,
    /// verified by the Interchain Module.
    /// @param verifiedAt   The block timestamp at which the entry was verified by the module
    /// @param entryValue   The value of the entry: srcWriter + digest hashed together
    struct RemoteEntry {
        uint256 verifiedAt;
        bytes32 entryValue;
    }

    /// @notice The version of the Interchain DataBase. Must match the version of the batches.
    uint16 public constant DB_VERSION = 1;

    /// @dev The entry values written on the local chain.
    bytes32[] internal _entryValues;
    /// @dev The remote entries verified by the modules.
    mapping(address module => mapping(EntryKey entryKey => RemoteEntry batch)) internal _remoteEntries;

    /// @dev Checks if the chain id is not the same as the local chain id.
    modifier onlyRemoteChainId(uint64 chainId) {
        if (chainId == block.chainid) {
            revert InterchainDB__ChainIdNotRemote(chainId);
        }
        _;
    }

    // ═══════════════════════════════════════════════ WRITER-FACING ═══════════════════════════════════════════════════

    /// @notice Write a data digest to the Interchain DataBase as a new entry.
    /// Note: there are no guarantees that this entry will be available for reading on any of the remote chains.
    /// Use `requestEntryVerification` to ensure that the entry is available for reading on the destination chain.
    /// @param digest       The digest of the data to be written to the Interchain DataBase
    /// @return dbNonce     The database nonce of the entry
    function writeEntry(bytes32 digest) external returns (uint64 dbNonce, uint64 entryIndex) {
        InterchainEntry memory entry = _writeEntry(digest);
        dbNonce = entry.dbNonce;
        // TODO: remove entryIndex
    }

    function requestBatchVerification(
        uint64 dstChainId,
        uint64 dbNonce,
        address[] calldata srcModules
    )
        external
        payable
        onlyRemoteChainId(dstChainId)
    {
        // TODO: remove
    }

    /// @notice Request the given Interchain Modules to verify an existing entry.
    /// Note: every module has a separate fee paid in the native gas token of the source chain,
    /// and `msg.value` must be equal to the sum of all fees.
    /// Note: this method is permissionless, and anyone can request verification for any entry.
    /// Note: will request the verification of an empty entry, if the entry with the given nonce does not exist.
    /// This could be useful for providing the proof of non-existence of the entry.
    /// @param dstChainId    The chain id of the destination chain
    /// @param dbNonce       The database nonce of the existing entry
    /// @param srcModules    The source chain addresses of the Interchain Modules to use for verification
    function requestEntryVerification(
        uint64 dstChainId,
        uint64 dbNonce,
        address[] calldata srcModules
    )
        external
        payable
        onlyRemoteChainId(dstChainId)
    {
        InterchainEntry memory entry = getEntry(dbNonce);
        _requestVerification(dstChainId, entry, srcModules);
    }

    function writeEntryWithVerification(
        uint64 dstChainId,
        bytes32 digest,
        address[] calldata srcModules
    )
        external
        payable
        onlyRemoteChainId(dstChainId)
        returns (uint64 dbNonce, uint64 entryIndex)
    {
        // TODO: remove
    }

    /// @notice Write a data digest to the Interchain DataBase as a new entry.
    /// Then request the Interchain Modules to verify the entry on the destination chain.
    /// See `writeEntry` and `requestEntryVerification` for more details.
    /// @dev Will revert if the empty array of modules is provided.
    /// @param dstChainId   The chain id of the destination chain
    /// @param digest       The digest of the data to be written to the Interchain DataBase
    /// @param srcModules   The source chain addresses of the Interchain Modules to use for verification
    /// @return dbNonce     The database nonce of the entry
    function writeEntryRequestVerification(
        uint64 dstChainId,
        bytes32 digest,
        address[] calldata srcModules
    )
        external
        payable
        onlyRemoteChainId(dstChainId)
        returns (uint64 dbNonce)
    {
        InterchainEntry memory entry = _writeEntry(digest);
        _requestVerification(dstChainId, entry, srcModules);
        return entry.dbNonce;
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
        // TODO: remove
    }

    function verifyRemoteEntry(bytes calldata encodedEntry) external {
        InterchainEntry memory entry = _assertCorrectEntry(encodedEntry);
        EntryKey entryKey = InterchainEntryLib.encodeEntryKey({srcChainId: entry.srcChainId, dbNonce: entry.dbNonce});
        RemoteEntry memory existingEntry = _remoteEntries[msg.sender][entryKey];
        // Check if that's the first time module verifies the entry
        if (existingEntry.verifiedAt == 0) {
            _saveVerifiedEntry(msg.sender, entryKey, entry);
            return;
        }
        // No-op if the entry value is the same
        if (existingEntry.entryValue == entry.entryValue) {
            return;
        }
        // Overwriting an empty (non-existent) entry with a different one is allowed
        if (existingEntry.entryValue == 0) {
            _saveVerifiedEntry(msg.sender, entryKey, entry);
            return;
        }
        // Overwriting an existing entry with a different one is not allowed
        revert InterchainDB__EntryConflict(msg.sender, entry);
    }

    // ═══════════════════════════════════════════════════ VIEWS ═══════════════════════════════════════════════════════

    /// @notice Returns the list of leafs of the finalized batch with the given nonce.
    /// Note: the leafs are ordered by the index of the written entry in the current batch,
    /// and the leafs value match the value of the written entry (srcWriter + dataHash hashed together).
    /// @dev Will revert if the batch with the given nonce does not exist, or is not finalized.
    /// @param dbNonce      The database nonce of the finalized batch
    function getBatchLeafs(uint64 dbNonce) external view returns (bytes32[] memory leafs) {
        // TODO: remove
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
        // TODO: remove
    }

    /// @notice Get the Merkle proof of inclusion for the entry with the given index
    /// in the finalized batch with the given nonce.
    /// @dev Will revert if the batch with the given nonce does not exist, or is not finalized.
    /// Will revert if the entry with the given index does not exist within the batch.
    /// @param dbNonce      The database nonce of the finalized batch
    /// @param entryIndex   The index of the written entry within the batch
    /// @return proof       The Merkle proof of inclusion for the entry
    function getEntryProof(uint64 dbNonce, uint64 entryIndex) external view returns (bytes32[] memory proof) {
        // TODO: remove
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
        // TODO: remove
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
        // TODO: remove
    }

    /// @notice Check if the entry is verified by the Interchain Module on the destination chain.
    /// - returned value `ENTRY_UNVERIFIED` indicates that the module has not verified the entry.
    /// - returned value `ENTRY_CONFLICT` indicates that the module has verified a different entry value
    /// for the same entry key.
    /// @param dstModule    The destination chain addresses of the Interchain Modules to use for verification
    /// @param entry        The Interchain Entry to check
    /// @return moduleVerifiedAt    The block timestamp at which the entry was verified by the module,
    /// ENTRY_UNVERIFIED if the module has not verified the entry,
    /// ENTRY_CONFLICT if the module has verified a different entry value for the same entry key.
    function checkEntryVerification(
        address dstModule,
        InterchainEntry memory entry
    )
        external
        view
        onlyRemoteChainId(entry.srcChainId)
        returns (uint256 moduleVerifiedAt)
    {
        EntryKey entryKey = InterchainEntryLib.encodeEntryKey({srcChainId: entry.srcChainId, dbNonce: entry.dbNonce});
        RemoteEntry memory remoteEntry = _remoteEntries[dstModule][entryKey];
        // Check if module verified anything for this entry key first
        if (remoteEntry.verifiedAt == 0) {
            return ENTRY_UNVERIFIED;
        }
        // Check if the entry value matches the one verified by the module
        return remoteEntry.entryValue == entry.entryValue ? remoteEntry.verifiedAt : ENTRY_CONFLICT;
    }

    /// @notice Get the versioned Interchain Batch with the given nonce.
    /// Note: will return a batch with an empty root if the batch does not exist, or is not finalized.
    /// @param dbNonce      The database nonce of the batch
    function getVersionedBatch(uint64 dbNonce) external view returns (bytes memory versionedBatch) {
        // TODO: remove
    }

    /// @notice Get the versioned Interchain Entry with the given nonce.
    /// Note: will return an entry with an empty value if the entry does not exist.
    /// @param dbNonce      The database nonce of the entry
    function getEncodedEntry(uint64 dbNonce) external view returns (bytes memory) {
        InterchainEntry memory entry = getEntry(dbNonce);
        return VersionedPayloadLib.encodeVersionedPayload({
            version: DB_VERSION,
            payload: InterchainEntryLib.encodeEntry(entry)
        });
    }

    /// @notice Get the batch root containing the Interchain Entry with the given index.
    /// @param entry         The Interchain Entry to get the batch root for
    /// @param proof         The Merkle proof of inclusion for the entry in the batch
    function getBatchRoot(InterchainEntry memory entry, bytes32[] calldata proof) external pure returns (bytes32) {
        // TODO: remove
    }

    /// @notice Returns the size of the finalized batch with the given nonce.
    /// @dev Will return 0 for non-existent or non-finalized batches.
    /// @param dbNonce      The database nonce of the finalized batch
    function getBatchSize(uint64 dbNonce) public view returns (uint64) {
        // TODO: remove
    }

    /// @notice Get the finalized Interchain Batch with the given nonce.
    /// @dev Will return a batch with an empty root if the batch does not exist, or is not finalized.
    /// @param dbNonce      The database nonce of the finalized batch
    function getBatch(uint64 dbNonce) public view returns (InterchainBatch memory) {
        // TODO: remove
    }

    /// @notice Get the Interchain Entry with the given nonce written on the local chain.
    /// @dev Will return an entry with an empty value if the entry does not exist.
    /// @param dbNonce      The database nonce of the entry
    function getEntry(uint64 dbNonce) public view returns (InterchainEntry memory) {
        return InterchainEntryLib.constructLocalEntry({dbNonce: dbNonce, entryValue: getEntryValue(dbNonce)});
    }

    /// @notice Get the Interchain Entry's value written on the local chain with the given nonce.
    /// Entry value is calculated as the hash of the writer address and the written data hash.
    /// @dev Will return an empty value if the entry does not exist.
    /// @param dbNonce      The database nonce of the entry
    function getEntryValue(uint64 dbNonce) public view returns (bytes32) {
        // For non-existent entries, the value is zero
        return dbNonce < getDBNonce() ? _entryValues[dbNonce] : bytes32(0);
    }

    /// @notice Get the Interchain Entry's value written on the local chain with the given batch nonce and entry index.
    /// Entry value is calculated as the hash of the writer address and the written data hash.
    /// Note: the batch does not have to be finalized to fetch the entry value.
    /// @dev Will revert if the batch with the given nonce does not exist,
    /// or the entry with the given index does not exist within the batch.
    /// @param dbNonce      The database nonce of the existing batch
    /// @param entryIndex   The index of the written entry within the batch
    function getEntryValue(uint64 dbNonce, uint64 entryIndex) public view returns (bytes32) {
        // TODO: remove
    }

    /// @notice Get the nonce of the database, which is incremented every time a new batch is finalized.
    /// This is the nonce of the current non-finalized batch.
    function getDBNonce() public view returns (uint64) {
        // We can do the unsafe cast here as writing more than 2^64 entries is practically impossible
        return uint64(_entryValues.length);
    }

    // ══════════════════════════════════════════════ INTERNAL LOGIC ═══════════════════════════════════════════════════

    /// @dev Write the entry to the database and emit the event.
    function _writeEntry(bytes32 digest) internal returns (InterchainEntry memory entry) {
        uint64 dbNonce = getDBNonce();
        bytes32 srcWriter = TypeCasts.addressToBytes32(msg.sender);
        bytes32 entryValue = InterchainEntryLib.getEntryValue({srcWriter: msg.sender, digest: digest});
        _entryValues.push(entryValue);
        entry = InterchainEntryLib.constructLocalEntry({dbNonce: dbNonce, entryValue: entryValue});
        emit InterchainEntryWritten({dbNonce: dbNonce, srcWriter: srcWriter, digest: digest, entryValue: entryValue});
    }

    /// @dev Request the verification of the entry by the modules, and emit the event.
    /// Note: the validity of the passed entry and chain id being remote is enforced in the calling function.
    function _requestVerification(
        uint64 dstChainId,
        InterchainEntry memory entry,
        address[] calldata srcModules
    )
        internal
    {
        (uint256[] memory fees, uint256 totalFee) = _getModuleFees(dstChainId, entry.dbNonce, srcModules);
        if (msg.value < totalFee) {
            revert InterchainDB__FeeAmountBelowMin(msg.value, totalFee);
        } else if (msg.value > totalFee) {
            // The exceeding amount goes to the first module
            fees[0] += msg.value - totalFee;
        }
        uint256 len = srcModules.length;
        bytes memory versionedEntry = VersionedPayloadLib.encodeVersionedPayload({
            version: DB_VERSION,
            payload: InterchainEntryLib.encodeEntry(entry)
        });
        for (uint256 i = 0; i < len; ++i) {
            IInterchainModule(srcModules[i]).requestEntryVerification{value: fees[i]}(
                dstChainId, entry.dbNonce, versionedEntry
            );
        }
        emit InterchainEntryVerificationRequested(dstChainId, entry.dbNonce, srcModules);
    }

    /// @dev Save the verified entry to the database and emit the event.
    function _saveVerifiedEntry(address module, EntryKey entryKey, InterchainEntry memory entry) internal {
        _remoteEntries[module][entryKey] = RemoteEntry({verifiedAt: block.timestamp, entryValue: entry.entryValue});
        emit InterchainEntryVerified(module, entry.srcChainId, entry.dbNonce, entry.entryValue);
    }

    // ══════════════════════════════════════════════ INTERNAL VIEWS ═══════════════════════════════════════════════════

    /// @dev Asserts that the entry version is correct and that entry originates from a remote chain.
    /// Note: returns the decoded entry for chaining purposes.
    function _assertCorrectEntry(bytes calldata versionedEntry) internal view returns (InterchainEntry memory entry) {
        uint16 dbVersion = versionedEntry.getVersion();
        if (dbVersion != DB_VERSION) {
            revert InterchainDB__EntryVersionMismatch(dbVersion, DB_VERSION);
        }
        entry = InterchainEntryLib.decodeEntry(versionedEntry.getPayload());
        if (entry.srcChainId == block.chainid) {
            revert InterchainDB__ChainIdNotRemote(entry.srcChainId);
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
