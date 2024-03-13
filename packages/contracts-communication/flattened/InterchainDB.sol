// SPDX-License-Identifier: MIT
pragma solidity =0.8.20 ^0.8.0;

// contracts/events/InterchainDBEvents.sol

abstract contract InterchainDBEvents {
    // TODO: figure out indexing
    event InterchainEntryWritten(uint256 srcChainId, uint256 dbNonce, bytes32 srcWriter, bytes32 dataHash);

    event InterchainBatchVerified(address module, uint256 srcChainId, uint256 dbNonce, bytes32 batchRoot);

    event InterchainBatchVerificationRequested(
        uint256 dstChainId, uint256 dbNonce, bytes32 batchRoot, address[] srcModules
    );
}

// contracts/libs/InterchainBatch.sol

/// @notice Struct representing a batch of entries in the Interchain DataBase.
/// Batched entries are put together in a Merkle tree, which root is saved.
/// Batch has a globally unique identifier (key) and a value.
/// - key: srcChainId + dbNonce
/// - value: batchRoot
/// @param srcChainId   The chain id of the source chain
/// @param dbNonce      The database nonce of the batch
/// @param batchRoot    The root of the Merkle tree containing the batched entries
struct InterchainBatch {
    // TODO: can we use uint64 for chain id?
    uint256 srcChainId;
    uint256 dbNonce;
    bytes32 batchRoot;
}

library InterchainBatchLib {
    /// @notice Constructs an InterchainBatch struct to be saved on the local chain.
    /// @param dbNonce      The database nonce of the batch
    /// @param batchRoot    The root of the Merkle tree containing the batched entries
    /// @return batch       The constructed InterchainBatch struct
    function constructLocalBatch(
        uint256 dbNonce,
        bytes32 batchRoot
    )
        internal
        view
        returns (InterchainBatch memory batch)
    {
        return InterchainBatch({srcChainId: block.chainid, dbNonce: dbNonce, batchRoot: batchRoot});
    }

    /// @notice Returns the globally unique identifier of the batch
    function batchKey(InterchainBatch memory batch) internal pure returns (bytes32) {
        return keccak256(abi.encode(batch.srcChainId, batch.dbNonce));
    }
}

// contracts/libs/TypeCasts.sol

library TypeCasts {
    function addressToBytes32(address addr) internal pure returns (bytes32) {
        return bytes32(uint256(uint160(addr)));
    }

    function bytes32ToAddress(bytes32 b) internal pure returns (address) {
        return address(uint160(uint256(b)));
    }
}

// contracts/libs/InterchainEntry.sol

/// @notice Struct representing an entry in the Interchain DataBase.
/// Entry has a globally unique identifier (key) and a value.
/// - key: srcChainId + dbNonce + entryIndex
/// - value: srcWriter + dataHash
/// @param srcChainId   The chain id of the source chain
/// @param dbNonce      The database nonce of the batch containing the entry
/// @param entryIndex   The index of the entry in the batch
/// @param srcWriter    The address of the writer on the source chain
/// @param dataHash     The hash of the data written on the source chain
struct InterchainEntry {
    // TODO: can we use uint64 for chain id?
    uint256 srcChainId;
    uint256 dbNonce;
    uint64 entryIndex;
    bytes32 srcWriter;
    bytes32 dataHash;
}

library InterchainEntryLib {
    /// @notice Constructs an InterchainEntry struct to be written on the local chain
    /// @param dbNonce      The database nonce of the entry on the source chain
    /// @param writer       The address of the writer on the local chain
    /// @param dataHash     The hash of the data written on the local chain
    /// @return entry       The constructed InterchainEntry struct
    function constructLocalEntry(
        uint256 dbNonce,
        uint64 entryIndex,
        address writer,
        bytes32 dataHash
    )
        internal
        view
        returns (InterchainEntry memory entry)
    {
        return InterchainEntry({
            srcChainId: block.chainid,
            dbNonce: dbNonce,
            entryIndex: entryIndex,
            srcWriter: TypeCasts.addressToBytes32(writer),
            dataHash: dataHash
        });
    }

    /// @notice Returns the globally unique identifier of the entry
    function entryKey(InterchainEntry memory entry) internal pure returns (bytes32) {
        return keccak256(abi.encode(entry.srcChainId, entry.dbNonce, entry.entryIndex));
    }

    /// @notice Returns the value of the entry: writer + dataHash hashed together
    function entryValue(InterchainEntry memory entry) internal pure returns (bytes32) {
        return keccak256(abi.encode(entry.srcWriter, entry.dataHash));
    }

    /// @notice Returns the globally unique identifier of the batch containing the entry
    function batchKey(InterchainEntry memory entry) internal pure returns (bytes32) {
        return keccak256(abi.encode(entry.srcChainId, entry.dbNonce));
    }
}

// contracts/interfaces/IInterchainDB.sol

interface IInterchainDB {
    /// @notice Struct representing an entry from the local Interchain DataBase.
    /// @param writer       The address of the writer on the local chain
    /// @param dataHash     The hash of the data written on the local chain
    struct LocalEntry {
        address writer;
        bytes32 dataHash;
    }

    /// @notice Struct representing a batch of entries from the remote Interchain DataBase,
    /// verified by the Interchain Module.
    /// @param verifiedAt   The block timestamp at which the entry was verified by the module
    /// @param batchRoot    The Merkle root of the batch
    struct RemoteBatch {
        uint256 verifiedAt;
        bytes32 batchRoot;
    }

    error InterchainDB__BatchDoesNotExist(uint256 dbNonce);
    error InterchainDB__BatchNotFinalized(uint256 dbNonce);
    error InterchainDB__ConflictingBatches(address module, bytes32 existingBatchRoot, InterchainBatch newBatch);
    error InterchainDB__EntryIndexOutOfRange(uint256 dbNonce, uint64 entryIndex, uint64 batchSize);
    error InterchainDB__IncorrectFeeAmount(uint256 actualFee, uint256 expectedFee);
    error InterchainDB__InvalidEntryRange(uint256 dbNonce, uint64 start, uint64 end);
    error InterchainDB__NoModulesSpecified();
    error InterchainDB__SameChainId(uint256 chainId);

    /// @notice Write data to the Interchain DataBase as a new entry in the current batch.
    /// Note: there are no guarantees that this entry will be available for reading on any of the remote chains.
    /// Use `requestBatchVerification` to ensure that the entry is available for reading on the destination chain.
    /// @param dataHash     The hash of the data to be written to the Interchain DataBase as a new entry
    /// @return dbNonce     The database nonce of the batch containing the written entry
    /// @return entryIndex  The index of the written entry within the batch
    function writeEntry(bytes32 dataHash) external returns (uint256 dbNonce, uint64 entryIndex);

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
        uint256 dstChainId,
        uint256 dbNonce,
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
        uint256 dstChainId,
        bytes32 dataHash,
        address[] memory srcModules
    )
        external
        payable
        returns (uint256 dbNonce, uint64 entryIndex);

    /// @notice Allows the Interchain Module to verify the batch coming from the remote chain.
    /// @param batch        The Interchain Batch to confirm
    function verifyRemoteBatch(InterchainBatch memory batch) external;

    // ═══════════════════════════════════════════════════ VIEWS ═══════════════════════════════════════════════════════

    /// @notice Get the fee for writing data to the Interchain DataBase, and verifying it on the destination chain
    /// using the provided Interchain Modules.
    /// @dev Will revert if the empty array of modules is provided.
    /// @param dstChainId   The chain id of the destination chain
    /// @param srcModules   The source chain addresses of the Interchain Modules to use for verification
    function getInterchainFee(uint256 dstChainId, address[] memory srcModules) external view returns (uint256);

    /// @notice Returns the list of leafs of the finalized batch with the given nonce.
    /// Note: the leafs are ordered by the index of the written entry in the current batch,
    /// and the leafs value match the value of the written entry (srcWriter + dataHash hashed together).
    /// @dev Will revert if the batch with the given nonce does not exist, or is not finalized.
    /// @param dbNonce      The database nonce of the finalized batch
    function getBatchLeafs(uint256 dbNonce) external view returns (bytes32[] memory);

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
        uint256 dbNonce,
        uint64 start,
        uint64 end
    )
        external
        view
        returns (bytes32[] memory);

    /// @notice Returns the size of the finalized batch with the given nonce.
    /// @dev Will revert if the batch with the given nonce does not exist, or is not finalized.
    /// @param dbNonce      The database nonce of the finalized batch
    function getBatchSize(uint256 dbNonce) external view returns (uint64);

    /// @notice Get the finalized Interchain Batch with the given nonce.
    /// @dev Will revert if the batch with the given nonce does not exist, or is not finalized.
    /// @param dbNonce      The database nonce of the finalized batch
    function getBatch(uint256 dbNonce) external view returns (InterchainBatch memory);

    /// @notice Get the Interchain Entry written on the local chain with the given batch nonce and entry index.
    /// Note: the batch does not have to be finalized to fetch the local entry.
    /// @dev Will revert if the batch with the given nonce does not exist,
    /// or the entry with the given index does not exist within the batch.
    /// @param dbNonce      The database nonce of the existing batch
    /// @param entryIndex   The index of the written entry within the batch
    function getEntry(uint256 dbNonce, uint64 entryIndex) external view returns (InterchainEntry memory);

    /// @notice Get the Merkle proof of inclusion for the entry with the given index
    /// in the finalized batch with the given nonce.
    /// @dev Will revert if the batch with the given nonce does not exist, or is not finalized.
    /// Will revert if the entry with the given index does not exist within the batch.
    /// @param dbNonce      The database nonce of the finalized batch
    /// @param entryIndex   The index of the written entry within the batch
    /// @return proof       The Merkle proof of inclusion for the entry
    function getEntryProof(uint256 dbNonce, uint64 entryIndex) external view returns (bytes32[] memory proof);

    /// @notice Get the nonce of the database, which is incremented every time a new batch is finalized.
    /// This is the nonce of the current non-finalized batch.
    function getDBNonce() external view returns (uint256);

    /// @notice Get the index of the next entry to be written to the database.
    /// @return dbNonce      The database nonce of the batch including the next entry
    /// @return entryIndex   The index of the next entry within that batch
    function getNextEntryIndex() external view returns (uint256 dbNonce, uint64 entryIndex);

    /// @notice Read the data written on specific source chain by a specific writer,
    /// and verify it on the destination chain using the provided Interchain Module.
    /// Note: returned zero value indicates that the module has not verified the entry.
    /// @param entry        The Interchain Entry to read
    /// @param dstModule    The destination chain addresses of the Interchain Modules to use for verification
    /// @return moduleVerifiedAt   The block timestamp at which the entry was verified by the module,
    ///                             or ZERO if the module has not verified the entry.
    function checkVerification(
        address dstModule,
        InterchainEntry memory entry,
        bytes32[] memory proof
    )
        external
        view
        returns (uint256 moduleVerifiedAt);
}

// contracts/interfaces/IInterchainModule.sol

/// @notice Every Module may opt a different method to confirm the verified entries on destination chain,
/// therefore this is not a part of a common interface.
interface IInterchainModule {
    error InterchainModule__NotInterchainDB(address caller);
    error InterchainModule__IncorrectSourceChainId(uint256 chainId);
    error InterchainModule__InsufficientFee(uint256 actual, uint256 required);
    error InterchainModule__SameChainId(uint256 chainId);

    /// @notice Request the verification of a batch from the Interchain DataBase by the module.
    /// If the batch is not yet finalized, the verification on destination chain will be delayed until
    /// the finalization is done and batch root is saved on the source chain.
    /// Note: a fee is paid to the module for verification, and could be retrieved by using `getModuleFee`.
    /// Note: this will eventually trigger `InterchainDB.verifyRemoteBatch(batch)` function on destination chain,
    /// with no guarantee of ordering.
    /// @dev Could be only called by the Interchain DataBase contract.
    /// @param dstChainId   The chain id of the destination chain
    /// @param batch        The batch to verify
    function requestBatchVerification(uint256 dstChainId, InterchainBatch memory batch) external payable;

    /// @notice Get the Module fee for verifying a batch on the specified destination chain.
    /// @param dstChainId   The chain id of the destination chain
    /// @param dbNonce      The database nonce of the batch on the source chain
    function getModuleFee(uint256 dstChainId, uint256 dbNonce) external view returns (uint256);
}

// contracts/InterchainDB.sol

contract InterchainDB is InterchainDBEvents, IInterchainDB {
    LocalEntry[] internal _entries;
    mapping(address module => mapping(bytes32 batchKey => RemoteBatch batch)) internal _remoteBatches;

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

    /// @inheritdoc IInterchainDB
    function verifyRemoteBatch(InterchainBatch memory batch) external onlyRemoteChainId(batch.srcChainId) {
        bytes32 batchKey = InterchainBatchLib.batchKey(batch);
        RemoteBatch memory existingBatch = _remoteBatches[msg.sender][batchKey];
        // Check if that's the first time module verifies the batch
        if (existingBatch.verifiedAt == 0) {
            _remoteBatches[msg.sender][batchKey] =
                RemoteBatch({verifiedAt: block.timestamp, batchRoot: batch.batchRoot});
            emit InterchainBatchVerified(msg.sender, batch.srcChainId, batch.dbNonce, batch.batchRoot);
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
        bytes32[] calldata proof
    )
        external
        view
        onlyRemoteChainId(entry.srcChainId)
        returns (uint256 moduleVerifiedAt)
    {
        // In "no batching" mode: the batch root is the same as the entry hash, hence the proof is empty
        if (proof.length != 0) {
            // If proof is not empty, the batch root is not verified
            return 0;
        }
        // In "no batching" mode: entry index is 0, batch size is 1
        if (entry.entryIndex != 0) {
            // If entry index is not 0, it does not belong to the batch
            return 0;
        }
        RemoteBatch memory remoteBatch = _remoteBatches[dstModule][InterchainEntryLib.batchKey(entry)];
        bytes32 entryValue = InterchainEntryLib.entryValue(entry);
        // Check entry value against the batch root verified by the module
        return remoteBatch.batchRoot == entryValue ? remoteBatch.verifiedAt : 0;
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
