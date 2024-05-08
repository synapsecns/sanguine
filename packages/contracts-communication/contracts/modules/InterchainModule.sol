// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {InterchainModuleEvents} from "../events/InterchainModuleEvents.sol";
import {IInterchainDB} from "../interfaces/IInterchainDB.sol";
import {IInterchainModule} from "../interfaces/IInterchainModule.sol";

import {InterchainBatch, InterchainBatchLib} from "../libs/InterchainBatch.sol";
import {ModuleBatchLib} from "../libs/ModuleBatch.sol";
import {VersionedPayloadLib} from "../libs/VersionedPayload.sol";

import {MessageHashUtils} from "@openzeppelin/contracts/utils/cryptography/MessageHashUtils.sol";

/// @notice Common logic for all Interchain Modules.
abstract contract InterchainModule is InterchainModuleEvents, IInterchainModule {
    using VersionedPayloadLib for bytes;

    /// @notice The address of the Interchain DataBase contract: used for verifying the batches.
    address public immutable INTERCHAIN_DB;

    constructor(address interchainDB) {
        INTERCHAIN_DB = interchainDB;
    }

    /// @notice Request the verification of a batch from the Interchain DataBase by the module.
    /// If the batch is not yet finalized, the verification on destination chain will be delayed until
    /// the finalization is done and batch root is saved on the source chain.
    /// Note: a fee is paid to the module for verification, and could be retrieved by using `getModuleFee`.
    /// Note: this will eventually trigger `InterchainDB.verifyRemoteBatch(batch)` function on destination chain,
    /// with no guarantee of ordering.
    /// @dev Could be only called by the Interchain DataBase contract.
    /// @param dstChainId       The chain id of the destination chain
    /// @param batchNonce       The nonce of the batch on the source chain
    /// @param versionedBatch   The versioned batch to verify
    function requestBatchVerification(
        uint64 dstChainId,
        uint64 batchNonce,
        bytes calldata versionedBatch
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
        uint256 requiredFee = _getModuleFee(dstChainId, batchNonce);
        if (msg.value < requiredFee) {
            revert InterchainModule__FeeAmountBelowMin({feeAmount: msg.value, minRequired: requiredFee});
        }
        bytes memory moduleData = _fillModuleData(dstChainId, batchNonce);
        bytes memory encodedBatch = ModuleBatchLib.encodeVersionedModuleBatch(versionedBatch, moduleData);
        bytes32 ethSignedBatchHash = MessageHashUtils.toEthSignedMessageHash(keccak256(encodedBatch));
        _requestVerification(dstChainId, encodedBatch);
        emit BatchVerificationRequested(dstChainId, encodedBatch, ethSignedBatchHash);
    }

    function requestEntryVerification(
        uint64 dstChainId,
        uint64 dbNonce,
        bytes memory versionedEntry
    )
        external
        payable
    {
        // TODO: implement
    }

    /// @notice Get the Module fee for verifying a batch on the specified destination chain.
    /// @param dstChainId   The chain id of the destination chain
    /// @param dbNonce      The database nonce of the batch on the source chain
    function getModuleFee(uint64 dstChainId, uint64 dbNonce) external view returns (uint256) {
        return _getModuleFee(dstChainId, dbNonce);
    }

    /// @dev Should be called once the Module has verified the batch and needs to signal this
    /// to the InterchainDB.
    function _verifyBatch(bytes memory encodedModuleBatch) internal {
        (bytes memory versionedBatch, bytes memory moduleData) =
            ModuleBatchLib.decodeVersionedModuleBatch(encodedModuleBatch);
        InterchainBatch memory batch = InterchainBatchLib.decodeBatchFromMemory(versionedBatch.getPayloadFromMemory());
        if (batch.srcChainId == block.chainid) {
            revert InterchainModule__ChainIdNotRemote(batch.srcChainId);
        }
        IInterchainDB(INTERCHAIN_DB).verifyRemoteBatch(versionedBatch);
        _receiveModuleData(batch.srcChainId, batch.dbNonce, moduleData);
        emit BatchVerified(
            batch.srcChainId, encodedModuleBatch, MessageHashUtils.toEthSignedMessageHash(keccak256(encodedModuleBatch))
        );
    }

    // solhint-disable no-empty-blocks
    /// @dev Internal logic to request the verification of an batch on the destination chain.
    function _requestVerification(uint64 dstChainId, bytes memory encodedBatch) internal virtual {}

    /// @dev Internal logic to fill the module data for the specified destination chain.
    function _fillModuleData(uint64 dstChainId, uint64 dbNonce) internal virtual returns (bytes memory) {}

    /// @dev Internal logic to handle the auxiliary module data relayed from the remote chain.
    function _receiveModuleData(uint64 srcChainId, uint64 dbNonce, bytes memory moduleData) internal virtual {}

    /// @dev Internal logic to get the module fee for verifying an batch on the specified destination chain.
    function _getModuleFee(uint64 dstChainId, uint64 dbNonce) internal view virtual returns (uint256);
}
