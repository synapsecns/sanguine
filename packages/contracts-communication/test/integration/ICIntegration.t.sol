// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {SynapseExecutionServiceEvents} from "../../contracts/events/SynapseExecutionServiceEvents.sol";
import {InterchainClientV1Events} from "../../contracts/events/InterchainClientV1Events.sol";
import {InterchainDBEvents} from "../../contracts/events/InterchainDBEvents.sol";
import {InterchainModuleEvents} from "../../contracts/events/InterchainModuleEvents.sol";

import {IInterchainApp} from "../../contracts/interfaces/IInterchainApp.sol";
import {IInterchainClientV1} from "../../contracts/interfaces/IInterchainClientV1.sol";
import {InterchainEntry} from "../../contracts/libs/InterchainEntry.sol";
import {InterchainTransaction, InterchainTxDescriptor} from "../../contracts/libs/InterchainTransaction.sol";
import {ModuleEntryLib} from "../../contracts/libs/ModuleEntry.sol";
import {OptionsV1} from "../../contracts/libs/Options.sol";

import {ICSetup, TypeCasts} from "./ICSetup.t.sol";

// solhint-disable custom-errors
// solhint-disable ordering
abstract contract ICIntegrationTest is
    ICSetup,
    SynapseExecutionServiceEvents,
    InterchainClientV1Events,
    InterchainDBEvents,
    InterchainModuleEvents
{
    struct FullEntry {
        uint64 srcChainId;
        uint64 dbNonce;
        bytes32 entryValue;
        bytes32 srcWriter;
        bytes32 digest;
    }

    using TypeCasts for address;

    function assertEq(InterchainEntry memory entry, InterchainEntry memory expected) internal pure {
        assertEq(entry.srcChainId, expected.srcChainId);
        assertEq(entry.dbNonce, expected.dbNonce);
        assertEq(entry.entryValue, expected.entryValue);
    }

    function expectServiceEventExecutionRequested(bytes32 transactionId, uint256 executionFee) internal {
        vm.expectEmit(address(executionService));
        emit ExecutionRequested({transactionId: transactionId, client: address(icClient), executionFee: executionFee});
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
            transactionId: getTxId(icTx),
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
            transactionId: getTxId(icTx),
            dbNonce: icTx.dbNonce,
            entryIndex: icTx.entryIndex,
            srcChainId: icTx.srcChainId,
            srcSender: icTx.srcSender,
            dstReceiver: icTx.dstReceiver
        });
    }

    function expectDatabaseEventInterchainEntryWritten(FullEntry memory fullEntry) internal {
        vm.expectEmit(address(icDB));
        emit InterchainEntryWritten({
            dbNonce: fullEntry.dbNonce,
            srcWriter: fullEntry.srcWriter,
            digest: fullEntry.digest,
            entryValue: fullEntry.entryValue
        });
    }

    function expectDatabaseEventInterchainEntryVerified(FullEntry memory fullEntry) internal {
        vm.expectEmit(address(icDB));
        emit InterchainEntryVerified({
            module: address(module),
            srcChainId: fullEntry.srcChainId,
            dbNonce: fullEntry.dbNonce,
            entryValue: fullEntry.entryValue
        });
    }

    function expectDatabaseEventInterchainEntryVerificationRequested(FullEntry memory fullEntry) internal {
        vm.expectEmit(address(icDB));
        emit InterchainEntryVerificationRequested({
            dstChainId: remoteChainId(),
            dbNonce: fullEntry.dbNonce,
            srcModules: toArray(address(module))
        });
    }

    function expectModuleEventEntryVerificationRequested(FullEntry memory fullEntry) internal {
        bytes memory encodedEntry = getModuleEntry(fullEntry);
        bytes32 digest = getEthSignedEntryHash(fullEntry);
        vm.expectEmit(address(module));
        emit EntryVerificationRequested({dstChainId: remoteChainId(), entry: encodedEntry, ethSignedEntryHash: digest});
    }

    function expectModuleEventEntryVerified(FullEntry memory fullEntry) internal {
        bytes memory encodedEntry = getModuleEntry(fullEntry);
        bytes32 digest = getEthSignedEntryHash(fullEntry);
        vm.expectEmit(address(module));
        emit EntryVerified({srcChainId: fullEntry.srcChainId, entry: encodedEntry, ethSignedEntryHash: digest});
    }

    function expectAppCall(InterchainTransaction memory icTx, OptionsV1 memory options) internal {
        bytes memory expectedCalldata = abi.encodeCall(
            IInterchainApp.appReceive, (icTx.srcChainId, icTx.srcSender, icTx.dbNonce, icTx.entryIndex, icTx.message)
        );
        vm.expectCall({
            callee: localApp(),
            msgValue: options.gasAirdrop,
            gas: uint64(options.gasLimit),
            data: expectedCalldata,
            count: 1
        });
    }

    // ═══════════════════════════════════════════ COMPLEX SERIES CHECKS ═══════════════════════════════════════════════

    function expectEventsMessageSent(
        InterchainTransaction memory icTx,
        FullEntry memory fullEntry,
        uint256 verificationFee,
        uint256 executionFee
    )
        internal
    {
        InterchainTxDescriptor memory desc = getInterchainTxDescriptor(fullEntry);
        expectDatabaseEventInterchainEntryWritten(fullEntry);
        expectModuleEventEntryVerificationRequested(fullEntry);
        expectDatabaseEventInterchainEntryVerificationRequested(fullEntry);
        expectServiceEventExecutionRequested(desc.transactionId, executionFee);
        expectClientEventInterchainTransactionSent(icTx, verificationFee, executionFee);
    }

    // ══════════════════════════════════════════════ EXPECT REVERTS ═══════════════════════════════════════════════════

    function expectClientRevertEntryConflict(address module) internal {
        vm.expectRevert(abi.encodeWithSelector(IInterchainClientV1.InterchainClientV1__EntryConflict.selector, module));
    }

    function expectClientRevertResponsesAmountBelowMin(uint256 actual, uint256 required) internal {
        vm.expectRevert(
            abi.encodeWithSelector(
                IInterchainClientV1.InterchainClientV1__ResponsesAmountBelowMin.selector, actual, required
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

    function checkDatabaseStateMsgSent(FullEntry memory fullEntry, uint64 initialDBNonce) internal view {
        InterchainEntry memory entry = getInterchainEntry(fullEntry);
        InterchainTxDescriptor memory desc = getInterchainTxDescriptor(fullEntry);
        assertEq(desc.dbNonce, initialDBNonce);
        // Check getters related to the txs' dbNonce
        assertEq(icDB.getEntry(desc.dbNonce), entry);
        assertEq(icDB.getEntryValue(desc.dbNonce), entry.entryValue);
        // Check getters related to the next dbNonce
        assertEq(icDB.getDBNonce(), desc.dbNonce + 1);
    }

    function markInvalidByGuard(FullEntry memory fullEntry) internal {
        InterchainEntry memory conflictingEntry = InterchainEntry({
            srcChainId: fullEntry.srcChainId,
            dbNonce: fullEntry.dbNonce,
            entryValue: keccak256("Some other data")
        });
        bytes memory encodedEntry =
            payloadLibHarness.encodeVersionedPayload(DB_VERSION, entryLibHarness.encodeEntry(conflictingEntry));
        vm.prank(guard);
        icDB.verifyRemoteEntry(encodedEntry);
    }

    // ═══════════════════════════════════════════════ DATA HELPERS ════════════════════════════════════════════════════

    function getModuleSignatures(FullEntry memory fullEntry) internal view returns (bytes memory signatures) {
        bytes32 digest = getEthSignedEntryHash(fullEntry);
        signatures = "";
        for (uint256 i = 0; i < signerPKs.length; i++) {
            (uint8 v, bytes32 r, bytes32 s) = vm.sign(signerPKs[i], digest);
            signatures = bytes.concat(signatures, abi.encodePacked(r, s, v));
        }
    }

    function getEthSignedEntryHash(FullEntry memory fullEntry) internal view returns (bytes32) {
        bytes memory moduleEntry = getModuleEntry(fullEntry);
        return keccak256(abi.encodePacked("\x19Ethereum Signed Message:\n32", keccak256(moduleEntry)));
    }

    function getModuleEntry(FullEntry memory fullEntry) internal view returns (bytes memory) {
        bytes memory versionedEntry = payloadLibHarness.encodeVersionedPayload(
            DB_VERSION, entryLibHarness.encodeEntry(getInterchainEntry(fullEntry))
        );
        return ModuleEntryLib.encodeVersionedModuleEntry(versionedEntry, new bytes(0));
    }

    function getInterchainEntry(FullEntry memory fullEntry) internal pure returns (InterchainEntry memory) {
        return InterchainEntry({
            srcChainId: fullEntry.srcChainId,
            dbNonce: fullEntry.dbNonce,
            entryValue: fullEntry.entryValue
        });
    }

    function getInterchainTxDescriptor(FullEntry memory fullEntry)
        internal
        pure
        returns (InterchainTxDescriptor memory)
    {
        return InterchainTxDescriptor({dbNonce: fullEntry.dbNonce, transactionId: fullEntry.digest});
    }

    function getSrcFullEntry() internal view returns (FullEntry memory) {
        bytes32 srcWriter = address(icClient).addressToBytes32();
        bytes32 digest = getTxId(getSrcTransaction());
        return FullEntry({
            srcChainId: SRC_CHAIN_ID,
            dbNonce: SRC_INITIAL_DB_NONCE,
            entryValue: entryLibHarness.getEntryValue(srcWriter, digest),
            srcWriter: srcWriter,
            digest: digest
        });
    }

    function getDstFullEntry() internal view returns (FullEntry memory) {
        bytes32 srcWriter = address(icClient).addressToBytes32();
        bytes32 digest = getTxId(getDstTransaction());
        return FullEntry({
            srcChainId: DST_CHAIN_ID,
            dbNonce: DST_INITIAL_DB_NONCE,
            entryValue: entryLibHarness.getEntryValue(srcWriter, digest),
            srcWriter: srcWriter,
            digest: digest
        });
    }

    function getTxId(InterchainTransaction memory icTx) internal view returns (bytes32) {
        return keccak256(getEncodedTx(icTx));
    }

    function getEncodedTx(InterchainTransaction memory icTx) internal view returns (bytes memory) {
        return payloadLibHarness.encodeVersionedPayload(CLIENT_VERSION, txLibHarness.encodeTransaction(icTx));
    }

    function getSrcTransaction() internal view returns (InterchainTransaction memory) {
        return InterchainTransaction({
            srcChainId: SRC_CHAIN_ID,
            srcSender: address(srcApp).addressToBytes32(),
            dstChainId: DST_CHAIN_ID,
            dstReceiver: address(dstApp).addressToBytes32(),
            dbNonce: SRC_INITIAL_DB_NONCE,
            options: getSrcOptions().encodeOptionsV1(),
            message: getSrcMessage()
        });
    }

    function getDstTransaction() internal view returns (InterchainTransaction memory) {
        return InterchainTransaction({
            srcChainId: DST_CHAIN_ID,
            srcSender: address(dstApp).addressToBytes32(),
            dstChainId: SRC_CHAIN_ID,
            dstReceiver: address(srcApp).addressToBytes32(),
            dbNonce: DST_INITIAL_DB_NONCE,
            options: getDstOptions().encodeOptionsV1(),
            message: getDstMessage()
        });
    }

    function getSrcOptions() internal view virtual returns (OptionsV1 memory);
    function getSrcMessage() internal view virtual returns (bytes memory);

    function getDstOptions() internal view virtual returns (OptionsV1 memory);
    function getDstMessage() internal view virtual returns (bytes memory);

    function toArray(address addr) internal pure returns (address[] memory arr) {
        arr = new address[](1);
        arr[0] = addr;
    }
}
