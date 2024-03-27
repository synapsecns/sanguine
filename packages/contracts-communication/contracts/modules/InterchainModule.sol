// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {InterchainModuleEvents} from "../events/InterchainModuleEvents.sol";
import {IInterchainDB} from "../interfaces/IInterchainDB.sol";
import {IInterchainModule} from "../interfaces/IInterchainModule.sol";

import {InterchainBatch, ModuleBatchLib} from "../libs/ModuleBatch.sol";

import {MessageHashUtils} from "@openzeppelin/contracts/utils/cryptography/MessageHashUtils.sol";

/// @notice Common logic for all Interchain Modules.
abstract contract InterchainModule is InterchainModuleEvents, IInterchainModule {
    address public immutable INTERCHAIN_DB;

    constructor(address interchainDB) {
        INTERCHAIN_DB = interchainDB;
    }

    /// @inheritdoc IInterchainModule
    function requestBatchVerification(uint256 dstChainId, InterchainBatch memory batch) external payable {
        if (msg.sender != INTERCHAIN_DB) {
            revert InterchainModule__NotInterchainDB(msg.sender);
        }
        if (dstChainId == block.chainid) {
            revert InterchainModule__SameChainId(block.chainid);
        }
        if (batch.srcChainId != block.chainid) {
            revert InterchainModule__IncorrectSourceChainId({chainId: batch.srcChainId});
        }
        uint256 requiredFee = _getModuleFee(dstChainId, batch.dbNonce);
        if (msg.value < requiredFee) {
            revert InterchainModule__InsufficientFee({actual: msg.value, required: requiredFee});
        }
        bytes memory moduleData = _fillModuleData(dstChainId, batch.dbNonce);
        bytes memory encodedBatch = ModuleBatchLib.encodeModuleBatch(batch, moduleData);
        bytes32 ethSignedBatchHash = MessageHashUtils.toEthSignedMessageHash(keccak256(encodedBatch));
        _requestVerification(dstChainId, encodedBatch);
        emit BatchVerificationRequested(dstChainId, encodedBatch, ethSignedBatchHash);
    }

    /// @inheritdoc IInterchainModule
    function getModuleFee(uint256 dstChainId, uint256 dbNonce) external view returns (uint256) {
        return _getModuleFee(dstChainId, dbNonce);
    }

    /// @dev Should be called once the Module has verified the batch and needs to signal this
    /// to the InterchainDB.
    function _verifyBatch(bytes memory encodedBatch) internal {
        (InterchainBatch memory batch, bytes memory moduleData) = ModuleBatchLib.decodeModuleBatch(encodedBatch);
        if (batch.srcChainId == block.chainid) {
            revert InterchainModule__SameChainId(block.chainid);
        }
        IInterchainDB(INTERCHAIN_DB).verifyRemoteBatch(batch);
        _receiveModuleData(batch.srcChainId, batch.dbNonce, moduleData);
        emit BatchVerified(
            batch.srcChainId, encodedBatch, MessageHashUtils.toEthSignedMessageHash(keccak256(encodedBatch))
        );
    }

    // solhint-disable no-empty-blocks
    /// @dev Internal logic to request the verification of an batch on the destination chain.
    function _requestVerification(uint256 dstChainId, bytes memory encodedBatch) internal virtual {}

    /// @dev Internal logic to fill the module data for the specified destination chain.
    function _fillModuleData(uint256 dstChainId, uint256 dbNonce) internal virtual returns (bytes memory) {}

    /// @dev Internal logic to handle the auxiliary module data relayed from the remote chain.
    function _receiveModuleData(uint256 srcChainId, uint256 dbNonce, bytes memory moduleData) internal virtual {}

    /// @dev Internal logic to get the module fee for verifying an batch on the specified destination chain.
    function _getModuleFee(uint256 dstChainId, uint256 dbNonce) internal view virtual returns (uint256);
}
