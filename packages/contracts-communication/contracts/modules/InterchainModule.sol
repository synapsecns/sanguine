// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {InterchainModuleEvents} from "../events/InterchainModuleEvents.sol";
import {IInterchainDB} from "../interfaces/IInterchainDB.sol";
import {IInterchainModule} from "../interfaces/IInterchainModule.sol";

import {InterchainEntry, ModuleEntryLib} from "../libs/ModuleEntry.sol";

import {MessageHashUtils} from "@openzeppelin/contracts/utils/cryptography/MessageHashUtils.sol";

/// @notice Common logic for all Interchain Modules.
abstract contract InterchainModule is InterchainModuleEvents, IInterchainModule {
    address public immutable INTERCHAIN_DB;

    constructor(address interchainDB) {
        INTERCHAIN_DB = interchainDB;
    }

    /// @inheritdoc IInterchainModule
    function requestVerification(uint256 dstChainId, InterchainEntry memory entry) external payable {
        if (msg.sender != INTERCHAIN_DB) {
            revert InterchainModule__NotInterchainDB(msg.sender);
        }
        if (dstChainId == block.chainid) {
            revert InterchainModule__SameChainId(block.chainid);
        }
        if (entry.srcChainId != block.chainid) {
            revert InterchainModule__IncorrectSourceChainId({chainId: entry.srcChainId});
        }
        uint256 requiredFee = _getModuleFee(dstChainId);
        if (msg.value < requiredFee) {
            revert InterchainModule__InsufficientFee({actual: msg.value, required: requiredFee});
        }
        bytes memory moduleData = _fillModuleData(dstChainId, entry.dbNonce);
        bytes memory encodedEntry = ModuleEntryLib.encodeModuleEntry(entry, moduleData);
        bytes32 ethSignedEntryHash = MessageHashUtils.toEthSignedMessageHash(keccak256(encodedEntry));
        _requestVerification(dstChainId, encodedEntry);
        emit VerificationRequested(dstChainId, encodedEntry, ethSignedEntryHash);
    }

    /// @inheritdoc IInterchainModule
    function getModuleFee(uint256 dstChainId) external view returns (uint256) {
        return _getModuleFee(dstChainId);
    }

    /// @dev Should be called once the Module has verified the entry and needs to signal this
    /// to the InterchainDB.
    function _verifyEntry(bytes memory encodedEntry) internal {
        (InterchainEntry memory entry, bytes memory moduleData) = ModuleEntryLib.decodeModuleEntry(encodedEntry);
        if (entry.srcChainId == block.chainid) {
            revert InterchainModule__SameChainId(block.chainid);
        }
        IInterchainDB(INTERCHAIN_DB).verifyEntry(entry);
        _receiveModuleData(entry.srcChainId, entry.dbNonce, moduleData);
        emit EntryVerified(
            entry.srcChainId, encodedEntry, MessageHashUtils.toEthSignedMessageHash(keccak256(encodedEntry))
        );
    }

    // solhint-disable no-empty-blocks
    /// @dev Internal logic to request the verification of an entry on the destination chain.
    function _requestVerification(uint256 dstChainId, bytes memory encodedEntry) internal virtual {}

    /// @dev Internal logic to fill the module data for the specified destination chain.
    function _fillModuleData(uint256 dstChainId, uint256 dbNonce) internal virtual returns (bytes memory) {}

    /// @dev Internal logic to handle the auxiliary module data relayed from the remote chain.
    function _receiveModuleData(uint256 srcChainId, uint256 dbNonce, bytes memory moduleData) internal virtual {}

    /// @dev Internal logic to get the module fee for verifying an entry on the specified destination chain.
    function _getModuleFee(uint256 dstChainId) internal view virtual returns (uint256);
}
