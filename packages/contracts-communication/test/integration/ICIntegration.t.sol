// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {ExecutionFeesEvents} from "../../contracts/events/ExecutionFeesEvents.sol";
import {ExecutionServiceEvents} from "../../contracts/events/ExecutionServiceEvents.sol";
import {InterchainClientV1Events} from "../../contracts/events/InterchainClientV1Events.sol";
import {InterchainDBEvents} from "../../contracts/events/InterchainDBEvents.sol";
import {InterchainModuleEvents} from "../../contracts/events/InterchainModuleEvents.sol";

import {InterchainBatch} from "../../contracts/libs/InterchainBatch.sol";
import {InterchainEntry} from "../../contracts/libs/InterchainEntry.sol";
import {InterchainTransaction} from "../../contracts/libs/InterchainTransaction.sol";
import {ModuleBatchLib} from "../../contracts/libs/ModuleBatch.sol";
import {OptionsV1} from "../../contracts/libs/Options.sol";

import {ICSetup, TypeCasts} from "./ICSetup.t.sol";

// solhint-disable custom-errors
// solhint-disable ordering
abstract contract ICIntegrationTest is
    ICSetup,
    ExecutionFeesEvents,
    ExecutionServiceEvents,
    InterchainClientV1Events,
    InterchainDBEvents,
    InterchainModuleEvents
{
    using TypeCasts for address;

    uint256 public constant COUNTER = 42;

    OptionsV1 public ppOptions = OptionsV1({gasLimit: 500_000, gasAirdrop: 0});

    function expectFeesEventExecutionFeeAdded(bytes32 transactionId, uint256 totalFee) internal {
        vm.expectEmit(address(executionFees));
        emit ExecutionFeeAdded({dstChainId: remoteChainId(), transactionId: transactionId, totalFee: totalFee});
    }

    function expectServiceEventExecutionRequested(bytes32 transactionId) internal {
        vm.expectEmit(address(executionService));
        emit ExecutionRequested({transactionId: transactionId, client: address(icClient)});
    }

    function expectClientEventInterchainTransactionSent(
        InterchainTransaction memory icTx,
        uint256 verificationFee,
        uint256 executionFee
    )
        internal
    {
        vm.expectEmit(address(icClient));
        emit InterchainTransactionSent({
            transactionId: icTx.transactionId(),
            dbNonce: icTx.dbNonce,
            entryIndex: icTx.entryIndex,
            dstChainId: icTx.dstChainId,
            srcSender: icTx.srcSender,
            dstReceiver: icTx.dstReceiver,
            verificationFee: verificationFee,
            executionFee: executionFee,
            options: icTx.options,
            message: icTx.message
        });
    }

    function expectClientEventInterchainTransactionReceived(InterchainTransaction memory icTx) internal {
        vm.expectEmit(address(icClient));
        emit InterchainTransactionReceived({
            transactionId: icTx.transactionId(),
            dbNonce: icTx.dbNonce,
            entryIndex: icTx.entryIndex,
            srcChainId: icTx.srcChainId,
            srcSender: icTx.srcSender,
            dstReceiver: icTx.dstReceiver
        });
    }

    function expectDatabaseEventInterchainEntryWritten(InterchainEntry memory entry) internal {
        vm.expectEmit(address(icDB));
        emit InterchainEntryWritten({
            srcChainId: entry.srcChainId,
            dbNonce: entry.dbNonce,
            srcWriter: entry.srcWriter,
            dataHash: entry.dataHash
        });
    }

    function expectDatabaseEventInterchainBatchVerified(InterchainBatch memory batch) internal {
        vm.expectEmit(address(icDB));
        emit InterchainBatchVerified({
            module: address(module),
            srcChainId: batch.srcChainId,
            dbNonce: batch.dbNonce,
            batchRoot: batch.batchRoot
        });
    }

    function expectDatabaseEventInterchainBatchVerificationRequested(InterchainBatch memory batch) internal {
        vm.expectEmit(address(icDB));
        emit InterchainBatchVerificationRequested({
            dstChainId: remoteChainId(),
            dbNonce: batch.dbNonce,
            batchRoot: batch.batchRoot,
            srcModules: toArray(address(module))
        });
    }

    function expectModuleEventBatchVerificationRequested(InterchainBatch memory batch) internal {
        bytes memory encodedBatch = getModuleBatch(batch);
        bytes32 digest = keccak256(abi.encodePacked("\x19Ethereum Signed Message:\n32", keccak256(encodedBatch)));
        vm.expectEmit(address(module));
        emit BatchVerificationRequested({dstChainId: remoteChainId(), batch: encodedBatch, ethSignedBatchHash: digest});
    }

    function expectModuleEventBatchVerified(InterchainBatch memory batch) internal {
        bytes memory encodedBatch = getModuleBatch(batch);
        bytes32 digest = keccak256(abi.encodePacked("\x19Ethereum Signed Message:\n32", keccak256(encodedBatch)));
        vm.expectEmit(address(module));
        emit BatchVerified({srcChainId: batch.srcChainId, batch: encodedBatch, ethSignedBatchHash: digest});
    }

    // ═══════════════════════════════════════════════ DATA HELPERS ════════════════════════════════════════════════════

    function getModuleBatch(InterchainBatch memory batch) internal pure returns (bytes memory) {
        return ModuleBatchLib.encodeModuleBatch(batch, new bytes(0));
    }

    function getInterchainBatch(InterchainEntry memory entry) internal pure returns (InterchainBatch memory) {
        return InterchainBatch({
            srcChainId: entry.srcChainId,
            dbNonce: entry.dbNonce,
            batchRoot: keccak256(abi.encode(entry.srcWriter, entry.dataHash))
        });
    }

    function getSrcInterchainEntry() internal view returns (InterchainEntry memory) {
        return InterchainEntry({
            srcChainId: SRC_CHAIN_ID,
            dbNonce: SRC_INITIAL_DB_NONCE,
            entryIndex: 0,
            srcWriter: address(icClient).addressToBytes32(),
            dataHash: getSrcTransaction().transactionId()
        });
    }

    function getDstInterchainEntry() internal view returns (InterchainEntry memory) {
        return InterchainEntry({
            srcChainId: DST_CHAIN_ID,
            dbNonce: DST_INITIAL_DB_NONCE,
            entryIndex: 0,
            srcWriter: address(icClient).addressToBytes32(),
            dataHash: getDstTransaction().transactionId()
        });
    }

    function getSrcTransaction() internal view returns (InterchainTransaction memory) {
        return InterchainTransaction({
            srcChainId: SRC_CHAIN_ID,
            srcSender: address(pingPongApp).addressToBytes32(),
            dstChainId: DST_CHAIN_ID,
            dstReceiver: address(pingPongApp).addressToBytes32(),
            dbNonce: SRC_INITIAL_DB_NONCE,
            entryIndex: 0,
            options: ppOptions.encodeOptionsV1(),
            message: getPingPongSrcMessage()
        });
    }

    function getDstTransaction() internal view returns (InterchainTransaction memory) {
        return InterchainTransaction({
            srcChainId: DST_CHAIN_ID,
            srcSender: address(pingPongApp).addressToBytes32(),
            dstChainId: SRC_CHAIN_ID,
            dstReceiver: address(pingPongApp).addressToBytes32(),
            dbNonce: DST_INITIAL_DB_NONCE,
            entryIndex: 0,
            options: ppOptions.encodeOptionsV1(),
            message: getPingPongDstMessage()
        });
    }

    /// @notice Message that source chain PingPongApp sends to destination chain.
    function getPingPongSrcMessage() internal pure returns (bytes memory) {
        return abi.encode(COUNTER);
    }

    /// @notice Message that destination chain PingPongApp sends back to source chain.
    function getPingPongDstMessage() internal pure returns (bytes memory) {
        return abi.encode(COUNTER - 1);
    }

    function toArray(address addr) internal pure returns (address[] memory arr) {
        arr = new address[](1);
        arr[0] = addr;
    }
}
