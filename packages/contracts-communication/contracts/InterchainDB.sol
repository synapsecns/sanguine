// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {IInterchainDB} from './interfaces/IInterchainDB.sol';
import {IInterchainDBEvents} from './interfaces/IInterchainDBEvents.sol';
import {IInterchainModule} from './interfaces/IInterchainModule.sol';

import {InterchainEntry, InterchainEntryLib} from './libs/InterchainEntry.sol';
import {TypeCasts} from './libs/TypeCasts.sol';

contract InterchainDB is IInterchainDB, IInterchainDBEvents {
  mapping(address writer => bytes32[] dataHashes) internal _entries;
  mapping(bytes32 entryId => mapping(address module => RemoteEntry entry))
    internal _remoteEntries;

  modifier onlyRemoteChainId(uint256 chainId) {
    if (chainId == block.chainid) {
      revert InterchainDB__SameChainId();
    }
    _;
  }

  // ═══════════════════════════════════════════════ WRITER-FACING ═══════════════════════════════════════════════════

  /// @inheritdoc IInterchainDB
  function writeEntry(bytes32 dataHash) external returns (uint256 writerNonce) {
    return _writeEntry(dataHash);
  }

  /// @inheritdoc IInterchainDB
  function requestVerification(
    uint256 destChainId,
    address writer,
    uint256 writerNonce,
    address[] calldata srcModules
  ) external payable onlyRemoteChainId(destChainId) {
    InterchainEntry memory entry = getEntry(writer, writerNonce);
    _requestVerification(destChainId, entry, srcModules);
  }

  /// @inheritdoc IInterchainDB
  function writeEntryWithVerification(
    uint256 destChainId,
    bytes32 dataHash,
    address[] calldata srcModules
  )
    external
    payable
    onlyRemoteChainId(destChainId)
    returns (uint256 writerNonce)
  {
    writerNonce = _writeEntry(dataHash);
    InterchainEntry memory entry = InterchainEntryLib.constructLocalEntry(
      msg.sender,
      writerNonce,
      dataHash
    );
    _requestVerification(destChainId, entry, srcModules);
  }

  // ═══════════════════════════════════════════════ MODULE-FACING ═══════════════════════════════════════════════════

  /// @inheritdoc IInterchainDB
  function verifyEntry(
    InterchainEntry memory entry
  ) external onlyRemoteChainId(entry.srcChainId) {
    bytes32 entryId = InterchainEntryLib.entryId(entry);
    RemoteEntry memory existingEntry = _remoteEntries[entryId][msg.sender];
    // Check if that's the first time module verifies the entry
    if (existingEntry.verifiedAt == 0) {
      _remoteEntries[entryId][msg.sender] = RemoteEntry({
        verifiedAt: block.timestamp,
        dataHash: entry.dataHash
      });
      emit InterchainEntryVerified(
        msg.sender,
        entry.srcChainId,
        entry.srcWriter,
        entry.writerNonce,
        entry.dataHash
      );
    } else {
      // If the module has already verified the entry, check that the data hash is the same
      if (existingEntry.dataHash != entry.dataHash) {
        revert InterchainDB__ConflictingEntries(existingEntry.dataHash, entry);
      }
      // No-op if the data hash is the same
    }
  }

  // ═══════════════════════════════════════════════════ VIEWS ═══════════════════════════════════════════════════════

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
    RemoteEntry memory remoteEntry = _remoteEntries[
      InterchainEntryLib.entryId(entry)
    ][dstModule];
    // Check data against the one verified by the module
    return remoteEntry.dataHash == entry.dataHash ? remoteEntry.verifiedAt : 0;
  }

  /// @inheritdoc IInterchainDB
  function getInterchainFee(
    uint256 destChainId,
    address[] calldata srcModules
  ) external view returns (uint256 fee) {
    (, fee) = _getModuleFees(destChainId, srcModules);
  }

  /// @inheritdoc IInterchainDB
  function getEntry(
    address writer,
    uint256 writerNonce
  ) public view returns (InterchainEntry memory) {
    if (getWriterNonce(writer) <= writerNonce) {
      revert InterchainDB__EntryDoesNotExist(writer, writerNonce);
    }
    return
      InterchainEntryLib.constructLocalEntry(
        writer,
        writerNonce,
        _entries[writer][writerNonce]
      );
  }

  /// @inheritdoc IInterchainDB
  function getWriterNonce(address writer) public view returns (uint256) {
    return _entries[writer].length;
  }

  // ══════════════════════════════════════════════ INTERNAL LOGIC ═══════════════════════════════════════════════════

  /// @dev Write the entry to the database and emit the event.
  function _writeEntry(
    bytes32 dataHash
  ) internal returns (uint256 writerNonce) {
    writerNonce = _entries[msg.sender].length;
    _entries[msg.sender].push(dataHash);
    emit InterchainEntryWritten(
      block.chainid,
      TypeCasts.addressToBytes32(msg.sender),
      writerNonce,
      dataHash
    );
  }

  /// @dev Request the verification of the entry by the modules, and emit the event.
  /// Note: the validity of the passed entry and chain id being remote is enforced in the calling function.
  function _requestVerification(
    uint256 destChainId,
    InterchainEntry memory entry,
    address[] calldata srcModules
  ) internal {
    (uint256[] memory fees, uint256 totalFee) = _getModuleFees(
      destChainId,
      srcModules
    );
    if (msg.value != totalFee) {
      revert InterchainDB__IncorrectFeeAmount(msg.value, totalFee);
    }
    uint256 len = srcModules.length;
    for (uint256 i = 0; i < len; ++i) {
      IInterchainModule(srcModules[i]).requestVerification{value: fees[i]}(
        destChainId,
        entry
      );
    }
    emit InterchainVerificationRequested(
      destChainId,
      entry.srcWriter,
      entry.writerNonce,
      srcModules
    );
  }

  // ══════════════════════════════════════════════ INTERNAL VIEWS ═══════════════════════════════════════════════════

  /// @dev Get the verification fees for the modules
  function _getModuleFees(
    uint256 destChainId,
    address[] calldata srcModules
  ) internal view returns (uint256[] memory fees, uint256 totalFee) {
    uint256 len = srcModules.length;
    if (len == 0) {
      revert InterchainDB__NoModulesSpecified();
    }
    fees = new uint256[](len);
    for (uint256 i = 0; i < len; ++i) {
      fees[i] = IInterchainModule(srcModules[i]).getModuleFee(destChainId);
      totalFee += fees[i];
    }
  }
}
