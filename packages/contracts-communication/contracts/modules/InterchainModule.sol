// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {InterchainModuleEvents} from "../events/InterchainModuleEvents.sol";
import {IInterchainDB} from "../interfaces/IInterchainDB.sol";
import {IInterchainModule} from "../interfaces/IInterchainModule.sol";

import {InterchainEntry} from "../libs/InterchainEntry.sol";

import {MessageHashUtils} from "@openzeppelin/contracts/utils/cryptography/MessageHashUtils.sol";

/// @notice Common logic for all Interchain Modules.
abstract contract InterchainModule is InterchainModuleEvents, IInterchainModule {
    address public immutable INTERCHAIN_DB;

    constructor(address interchainDB) {
        INTERCHAIN_DB = interchainDB;
    }

    /// @inheritdoc IInterchainModule
    function requestVerification(uint256 destChainId, InterchainEntry memory entry) external payable {
        if (msg.sender != INTERCHAIN_DB) {
            revert InterchainModule__NotInterchainDB();
        }
        if (destChainId == block.chainid) {
            revert InterchainModule__SameChainId();
        }
        if (entry.srcChainId != block.chainid) {
            revert InterchainModule__IncorrectSourceChainId({chainId: entry.srcChainId});
        }
        uint256 requiredFee = _getModuleFee(destChainId);
        if (msg.value < requiredFee) {
            revert InterchainModule__InsufficientFee({actual: msg.value, required: requiredFee});
        }
        bytes memory encodedEntry = abi.encode(entry);
        bytes32 ethSignedEntryHash = MessageHashUtils.toEthSignedMessageHash(keccak256(encodedEntry));
        _requestVerification(destChainId, encodedEntry);
        emit VerificationRequested(destChainId, encodedEntry, ethSignedEntryHash);
    }

    /// @inheritdoc IInterchainModule
    function getModuleFee(uint256 destChainId) external view returns (uint256) {
        return _getModuleFee(destChainId);
    }

    /// @dev Should be called once the Module has verified the entry and needs to signal this
    /// to the InterchainDB.
    function _verifyEntry(bytes memory encodedEntry) internal {
        InterchainEntry memory entry = abi.decode(encodedEntry, (InterchainEntry));
        if (entry.srcChainId == block.chainid) {
            revert InterchainModule__SameChainId();
        }
        IInterchainDB(INTERCHAIN_DB).verifyEntry(entry);
        emit EntryVerified(entry);
    }

    /// @dev Internal logic to request the verification of an entry on the destination chain.
    function _requestVerification(uint256 destChainId, bytes memory encodedEntry) internal virtual;

    /// @dev Internal logic to get the module fee for verifying an entry on the specified destination chain.
    function _getModuleFee(uint256 destChainId) internal view virtual returns (uint256);
}
