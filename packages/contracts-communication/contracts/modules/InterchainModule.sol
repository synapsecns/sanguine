// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {InterchainModuleEvents} from "../events/InterchainModuleEvents.sol";
import {IInterchainDB} from "../interfaces/IInterchainDB.sol";
import {IInterchainModule} from "../interfaces/IInterchainModule.sol";

import {InterchainEntry, InterchainEntryLib} from "../libs/InterchainEntry.sol";
import {ModuleEntryLib} from "../libs/ModuleEntry.sol";
import {VersionedPayloadLib} from "../libs/VersionedPayload.sol";

import {MessageHashUtils} from "@openzeppelin/contracts/utils/cryptography/MessageHashUtils.sol";

/// @notice Common logic for all Interchain Modules.
abstract contract InterchainModule is InterchainModuleEvents, IInterchainModule {
    using VersionedPayloadLib for bytes;

    /// @notice The address of the Interchain DataBase contract: used for verifying the entries.
    address public immutable INTERCHAIN_DB;

    constructor(address interchainDB) {
        INTERCHAIN_DB = interchainDB;
    }

    /// @notice Request the verification of an entry from the Interchain DataBase by the module.
    /// Note: a fee is paid to the module for verification, and could be retrieved by using `getModuleFee`.
    /// Note: this will eventually trigger `InterchainDB.verifyRemoteEntry(entry)` function on destination chain,
    /// with no guarantee of ordering.
    /// @dev Could be only called by the Interchain DataBase contract.
    /// @param dstChainId       The chain id of the destination chain
    /// @param dbNonce          The database nonce of the entry on the source chain
    /// @param versionedEntry   The versioned entry to verify
    function requestEntryVerification(
        uint64 dstChainId,
        uint64 dbNonce,
        bytes memory versionedEntry
    )
        external
        payable
    {
        if (msg.sender != INTERCHAIN_DB) {
            revert InterchainModule__CallerNotInterchainDB(msg.sender);
        }
        if (dstChainId == block.chainid) {
            revert InterchainModule__ChainIdNotRemote(dstChainId);
        }
        uint256 requiredFee = _getModuleFee(dstChainId, dbNonce);
        if (msg.value < requiredFee) {
            revert InterchainModule__FeeAmountBelowMin({feeAmount: msg.value, minRequired: requiredFee});
        }
        bytes memory moduleData = _fillModuleData(dstChainId, dbNonce);
        bytes memory encodedEntry = ModuleEntryLib.encodeVersionedModuleEntry(versionedEntry, moduleData);
        bytes32 ethSignedEntryHash = MessageHashUtils.toEthSignedMessageHash(keccak256(encodedEntry));
        _requestVerification(dstChainId, encodedEntry);
        emit EntryVerificationRequested(dstChainId, encodedEntry, ethSignedEntryHash);
    }

    /// @notice Get the Module fee for verifying an entry on the specified destination chain.
    /// @param dstChainId   The chain id of the destination chain
    /// @param dbNonce      The database nonce of the entry on the source chain
    function getModuleFee(uint64 dstChainId, uint64 dbNonce) external view returns (uint256) {
        return _getModuleFee(dstChainId, dbNonce);
    }

    /// @dev Should be called once the Module has verified the entry and needs to signal this
    /// to the InterchainDB.
    function _verifyEntry(bytes memory encodedModuleEntry) internal {
        (bytes memory versionedEntry, bytes memory moduleData) =
            ModuleEntryLib.decodeVersionedModuleEntry(encodedModuleEntry);
        InterchainEntry memory entry = InterchainEntryLib.decodeEntryFromMemory(versionedEntry.getPayloadFromMemory());
        if (entry.srcChainId == block.chainid) {
            revert InterchainModule__ChainIdNotRemote(entry.srcChainId);
        }
        IInterchainDB(INTERCHAIN_DB).verifyRemoteEntry(versionedEntry);
        _receiveModuleData(entry.srcChainId, entry.dbNonce, moduleData);
        emit EntryVerified(
            entry.srcChainId, encodedModuleEntry, MessageHashUtils.toEthSignedMessageHash(keccak256(encodedModuleEntry))
        );
    }

    // solhint-disable no-empty-blocks
    /// @dev Internal logic to request the verification of an entry on the destination chain.
    function _requestVerification(uint64 dstChainId, bytes memory encodedEntry) internal virtual {}

    /// @dev Internal logic to fill the module data for the specified destination chain.
    function _fillModuleData(uint64 dstChainId, uint64 dbNonce) internal virtual returns (bytes memory) {}

    /// @dev Internal logic to handle the auxiliary module data relayed from the remote chain.
    function _receiveModuleData(uint64 srcChainId, uint64 dbNonce, bytes memory moduleData) internal virtual {}

    /// @dev Internal logic to get the module fee for verifying an entry on the specified destination chain.
    function _getModuleFee(uint64 dstChainId, uint64 dbNonce) internal view virtual returns (uint256);
}
