// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {InterchainDBEvents} from "./events/InterchainDBEvents.sol";
import {IInterchainDB} from "./interfaces/IInterchainDB.sol";
import {IInterchainModule} from "./interfaces/IInterchainModule.sol";

import {
    InterchainEntry, InterchainEntryLib, EntryKey, ENTRY_UNVERIFIED, ENTRY_CONFLICT
} from "./libs/InterchainEntry.sol";
import {VersionedPayloadLib} from "./libs/VersionedPayload.sol";
import {TypeCasts} from "./libs/TypeCasts.sol";

contract InterchainDB is InterchainDBEvents, IInterchainDB {
    using VersionedPayloadLib for bytes;

    /// @notice Struct representing an entry from the remote Interchain DataBase,
    /// verified by the Interchain Module.
    /// @param verifiedAt   The block timestamp at which the entry was verified by the module
    /// @param entryValue   The value of the entry: srcWriter + digest hashed together
    struct RemoteEntry {
        uint256 verifiedAt;
        bytes32 entryValue;
    }

    /// @notice The version of the Interchain DataBase. Must match the version of the entries.
    uint16 public constant DB_VERSION = 1;

    /// @dev The entry values written on the local chain.
    bytes32[] internal _entryValues;
    /// @dev The remote entries verified by the modules.
    mapping(address module => mapping(EntryKey entryKey => RemoteEntry entry)) internal _remoteEntries;

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
    function writeEntry(bytes32 digest) external returns (uint64 dbNonce) {
        InterchainEntry memory entry = _writeEntry(digest);
        dbNonce = entry.dbNonce;
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

    /// @notice Allows the Interchain Module to verify the entry coming from the remote chain.
    /// The module SHOULD verify the exact finalized entry from the remote chain. If the entry with a given nonce
    /// does not exist, module CAN verify it with an empty entry value.
    /// Once the entry exists, the module SHOULD re-verify the entry with the correct entry value.
    /// Note: The DB will only accept the entry of the same version as the DB itself.
    /// @dev Will revert if the entry with the same nonce but a different non-empty entry value is already verified.
    /// @param encodedEntry   The versioned Interchain Entry to verify
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

    /// @notice Get the fee for writing data to the Interchain DataBase, and verifying it on the destination chain
    /// using the provided Interchain Modules.
    /// @dev Will revert if the empty array of modules is provided.
    /// @param dstChainId   The chain id of the destination chain
    /// @param srcModules   The source chain addresses of the Interchain Modules to use for verification
    function getInterchainFee(uint64 dstChainId, address[] calldata srcModules) external view returns (uint256 fee) {
        (, fee) = _getModuleFees(dstChainId, srcModules);
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

    /// @notice Get the nonce of the database, which is incremented every time a new entry is written.
    /// This is the nonce of the next entry to be written.
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
        (uint256[] memory fees, uint256 totalFee) = _getModuleFees(dstChainId, srcModules);
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
            IInterchainModule(srcModules[i]).requestEntryVerification{value: fees[i]}(dstChainId, versionedEntry);
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
            fees[i] = IInterchainModule(srcModules[i]).getModuleFee(dstChainId);
            totalFee += fees[i];
        }
    }
}
