// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {ExecutionFeesEvents} from "../../contracts/events/ExecutionFeesEvents.sol";
import {SynapseExecutionServiceEvents} from "../../contracts/events/SynapseExecutionServiceEvents.sol";
import {InterchainClientV1Events} from "../../contracts/events/InterchainClientV1Events.sol";
import {InterchainDBEvents} from "../../contracts/events/InterchainDBEvents.sol";
import {InterchainModuleEvents} from "../../contracts/events/InterchainModuleEvents.sol";

import {IInterchainApp} from "../../contracts/interfaces/IInterchainApp.sol";
import {IInterchainClientV1} from "../../contracts/interfaces/IInterchainClientV1.sol";
import {InterchainBatch} from "../../contracts/libs/InterchainBatch.sol";
import {InterchainEntry} from "../../contracts/libs/InterchainEntry.sol";
import {InterchainTransaction, InterchainTxDescriptor} from "../../contracts/libs/InterchainTransaction.sol";
import {ModuleBatchLib} from "../../contracts/libs/ModuleBatch.sol";
import {OptionsV1} from "../../contracts/libs/Options.sol";

import {ICSetup, TypeCasts} from "./ICSetup.t.sol";

// solhint-disable custom-errors
// solhint-disable ordering
abstract contract ICIntegrationTest is
    ICSetup,
    ExecutionFeesEvents,
    SynapseExecutionServiceEvents,
    InterchainClientV1Events,
    InterchainDBEvents,
    InterchainModuleEvents
{
    using TypeCasts for address;

    uint256 public constant COUNTER = 42;

    OptionsV1 public ppOptions = OptionsV1({gasLimit: 500_000, gasAirdrop: 0});

    event PingReceived(uint256 counter, uint256 dbNonce, uint64 entryIndex);
    event PingSent(uint256 counter, uint256 dbNonce, uint64 entryIndex);

    function assertEq(InterchainBatch memory batch, InterchainBatch memory expected) internal {
        assertEq(batch.srcChainId, expected.srcChainId);
        assertEq(batch.dbNonce, expected.dbNonce);
        assertEq(batch.batchRoot, expected.batchRoot);
    }

    function assertEq(InterchainEntry memory entry, InterchainEntry memory expected) internal {
        assertEq(entry.srcChainId, expected.srcChainId);
        assertEq(entry.dbNonce, expected.dbNonce);
        assertEq(entry.entryIndex, expected.entryIndex);
        assertEq(entry.srcWriter, expected.srcWriter);
        assertEq(entry.dataHash, expected.dataHash);
    }

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
        bytes32 digest = getEthSignedBatchHash(batch);
        vm.expectEmit(address(module));
        emit BatchVerificationRequested({dstChainId: remoteChainId(), batch: encodedBatch, ethSignedBatchHash: digest});
    }

    function expectModuleEventBatchVerified(InterchainBatch memory batch) internal {
        bytes memory encodedBatch = getModuleBatch(batch);
        bytes32 digest = getEthSignedBatchHash(batch);
        vm.expectEmit(address(module));
        emit BatchVerified({srcChainId: batch.srcChainId, batch: encodedBatch, ethSignedBatchHash: digest});
    }

    function expectPingPongEventPingReceived(uint256 counter, InterchainTxDescriptor memory desc) internal {
        vm.expectEmit(address(pingPongApp));
        emit PingReceived(counter, desc.dbNonce, desc.entryIndex);
    }

    function expectPingPongEventPingSent(uint256 counter, InterchainTxDescriptor memory desc) internal {
        vm.expectEmit(address(pingPongApp));
        emit PingSent(counter, desc.dbNonce, desc.entryIndex);
    }

    function expectPingPongCall(InterchainTransaction memory icTx, OptionsV1 memory options) internal {
        bytes memory expectedCalldata = abi.encodeCall(
            IInterchainApp.appReceive, (icTx.srcChainId, icTx.srcSender, icTx.dbNonce, icTx.entryIndex, icTx.message)
        );
        vm.expectCall({
            callee: address(pingPongApp),
            msgValue: options.gasAirdrop,
            gas: uint64(options.gasLimit),
            data: expectedCalldata,
            count: 1
        });
    }

    // ══════════════════════════════════════════════ EXPECT REVERTS ═══════════════════════════════════════════════════

    function expectClientRevertNotEnoughResponses(uint256 actual, uint256 required) internal {
        vm.expectRevert(
            abi.encodeWithSelector(
                IInterchainClientV1.InterchainClientV1__NotEnoughResponses.selector, actual, required
            )
        );
    }

    function expectClientRevertTxAlreadyExecuted(InterchainTxDescriptor memory desc) internal {
        vm.expectRevert(
            abi.encodeWithSelector(
                IInterchainClientV1.InterchainClientV1__TxAlreadyExecuted.selector, desc.transactionId
            )
        );
    }

    // ═══════════════════════════════════════════ COMPLEX SERIES CHECKS ═══════════════════════════════════════════════

    function expectEventsPingSent(
        uint256 counter,
        InterchainTransaction memory icTx,
        InterchainEntry memory entry,
        uint256 verificationFee,
        uint256 executionFee
    )
        internal
    {
        InterchainBatch memory batch = getInterchainBatch(entry);
        InterchainTxDescriptor memory desc = getInterchainTxDescriptor(entry);
        expectDatabaseEventInterchainEntryWritten(entry);
        expectModuleEventBatchVerificationRequested(batch);
        expectDatabaseEventInterchainBatchVerificationRequested(batch);
        expectFeesEventExecutionFeeAdded(desc.transactionId, executionFee);
        expectServiceEventExecutionRequested(desc.transactionId);
        expectClientEventInterchainTransactionSent(icTx, verificationFee, executionFee);
        expectPingPongEventPingSent(counter, desc);
    }

    function checkBatchLeafs(InterchainBatch memory batch, bytes32[] memory leafs) internal {
        assertEq(leafs.length, 1);
        assertEq(leafs[0], batch.batchRoot);
    }

    function checkDatabaseStatePingSent(InterchainEntry memory entry, uint256 initialDBNonce) internal {
        InterchainBatch memory batch = getInterchainBatch(entry);
        InterchainTxDescriptor memory desc = getInterchainTxDescriptor(entry);
        assertEq(desc.dbNonce, initialDBNonce);
        // Check getters related to the txs' dbNonce
        checkBatchLeafs(batch, icDB.getBatchLeafs(desc.dbNonce));
        checkBatchLeafs(batch, icDB.getBatchLeafsPaginated(desc.dbNonce, 0, 1));
        assertEq(icDB.getBatchSize(desc.dbNonce), 1);
        assertEq(icDB.getBatch(desc.dbNonce), batch);
        assertEq(icDB.getEntry(desc.dbNonce, 0), entry);
        assertEq(icDB.getEntryProof(desc.dbNonce, 0).length, 0);
        // Check getters related to the next dbNonce
        assertEq(icDB.getDBNonce(), desc.dbNonce + 1);
        (uint256 dbNonce, uint64 entryIndex) = icDB.getNextEntryIndex();
        assertEq(dbNonce, desc.dbNonce + 1);
        assertEq(entryIndex, 0);
    }

    // ═══════════════════════════════════════════════ DATA HELPERS ════════════════════════════════════════════════════

    function getModuleSignatures(InterchainBatch memory batch) internal view returns (bytes memory signatures) {
        bytes32 digest = getEthSignedBatchHash(batch);
        signatures = "";
        for (uint256 i = 0; i < signerPKs.length; i++) {
            (uint8 v, bytes32 r, bytes32 s) = vm.sign(signerPKs[i], digest);
            signatures = bytes.concat(signatures, abi.encodePacked(r, s, v));
        }
    }

    function getEthSignedBatchHash(InterchainBatch memory batch) internal pure returns (bytes32) {
        bytes memory moduleBatch = getModuleBatch(batch);
        return keccak256(abi.encodePacked("\x19Ethereum Signed Message:\n32", keccak256(moduleBatch)));
    }

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

    function getInterchainTxDescriptor(InterchainEntry memory entry)
        internal
        pure
        returns (InterchainTxDescriptor memory)
    {
        return InterchainTxDescriptor({
            dbNonce: entry.dbNonce,
            entryIndex: entry.entryIndex,
            transactionId: entry.dataHash
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
