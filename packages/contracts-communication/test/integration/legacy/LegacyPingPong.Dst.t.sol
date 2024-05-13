// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {IInterchainClientV1} from "../../../contracts/interfaces/IInterchainClientV1.sol";
import {InterchainEntry} from "../../../contracts/libs/InterchainEntry.sol";
import {InterchainTransaction, InterchainTxDescriptor} from "../../../contracts/libs/InterchainTransaction.sol";
import {OptionsV1} from "../../../contracts/libs/Options.sol";

import {LegacyPingPongIntegrationTest} from "./LegacyPingPong.t.sol";

// solhint-disable func-name-mixedcase
// solhint-disable ordering
contract LegacyPingPongDstIntegrationTest is LegacyPingPongIntegrationTest {
    uint256 public constant LONG_PERIOD = 1 weeks;

    FullEntry public srcFullEntry;
    InterchainTransaction public srcTx;
    InterchainTxDescriptor public srcDesc;
    bytes public encodedSrcTx;

    FullEntry public dstFullEntry;
    InterchainTransaction public dstTx;
    InterchainTxDescriptor public dstDesc;

    uint256 public dstPingFee;
    uint256 public dstVerificationFee;
    uint256 public dstExecutionFee;

    bytes public moduleEntry;
    bytes public moduleSignatures;

    function setUp() public override {
        super.setUp();
        srcTx = getSrcTransaction();
        srcFullEntry = getSrcFullEntry();
        srcDesc = getInterchainTxDescriptor(srcFullEntry);
        encodedSrcTx = getEncodedTx(srcTx);

        moduleEntry = getModuleEntry(srcFullEntry);
        moduleSignatures = getModuleSignatures(srcFullEntry);

        dstTx = getDstTransaction();
        dstFullEntry = getDstFullEntry();
        dstDesc = getInterchainTxDescriptor(dstFullEntry);

        dstPingFee = dstLegacyPingPong().getPingFee(SRC_CHAIN_ID);
        dstVerificationFee = icDB.getInterchainFee(SRC_CHAIN_ID, toArray(address(module)));
        dstExecutionFee = executionService.getExecutionFee({
            dstChainId: SRC_CHAIN_ID,
            txPayloadSize: getEncodedTx(dstTx).length,
            options: icOptions.encodeOptionsV1()
        });
    }

    function executeTx(OptionsV1 memory options) internal {
        vm.prank(executor);
        icClient.interchainExecute{value: options.gasAirdrop}({gasLimit: options.gasLimit, transaction: encodedSrcTx});
    }

    function test_verifyRemoteEntry_events() public {
        expectDatabaseEventInterchainEntryVerified(srcFullEntry);
        expectModuleEventEntryVerified(srcFullEntry);
        module.verifyRemoteEntry(moduleEntry, moduleSignatures);
    }

    function test_verifyRemoteEntry_state_client() public {
        module.verifyRemoteEntry(moduleEntry, moduleSignatures);
        skip(APP_OPTIMISTIC_PERIOD + 1);
        assertEq(icClient.getExecutor(encodedSrcTx), address(0));
        assertEq(icClient.getExecutorById(srcDesc.transactionId), address(0));
        assertTrue(icClient.isExecutable(encodedSrcTx));
    }

    function test_verifyRemoteEntry_state_db() public {
        InterchainEntry memory entry = getInterchainEntry(srcFullEntry);
        module.verifyRemoteEntry(moduleEntry, moduleSignatures);
        skip(LONG_PERIOD);
        assertEq(icDB.checkEntryVerification(address(module), entry), INITIAL_TS);
    }

    function test_interchainExecute_callMessageBusAndLegacyPP() public {
        module.verifyRemoteEntry(moduleEntry, moduleSignatures);
        skip(APP_OPTIMISTIC_PERIOD + 1);
        expectAppCall(srcTx, icOptions);
        expectPingPongCall();
        executeTx(icOptions);
    }

    function test_interchainExecute_callMessageBusAndLegacyPP_lowerGas() public {
        OptionsV1 memory options = OptionsV1({gasLimit: icOptions.gasLimit / 2, gasAirdrop: 0});
        module.verifyRemoteEntry(moduleEntry, moduleSignatures);
        skip(APP_OPTIMISTIC_PERIOD + 1);
        // Should use the requested gas limit
        expectAppCall(srcTx, icOptions);
        expectPingPongCall();
        executeTx(options);
    }

    function test_interchainExecute_callMessageBusAndLegacyPP_higherGas() public {
        OptionsV1 memory options = OptionsV1({gasLimit: 2 * icOptions.gasLimit, gasAirdrop: 0});
        module.verifyRemoteEntry(moduleEntry, moduleSignatures);
        skip(APP_OPTIMISTIC_PERIOD + 1);
        // Should allow to use higher gas limit
        expectAppCall(srcTx, options);
        expectPingPongCall();
        executeTx(options);
    }

    function test_interchainExecute_events() public {
        module.verifyRemoteEntry(moduleEntry, moduleSignatures);
        skip(APP_OPTIMISTIC_PERIOD + 1);
        expectPingPongEventPingReceived(COUNTER);
        expectEventsPingSent(COUNTER - 1, dstTx, dstFullEntry, dstVerificationFee, dstExecutionFee);
        expectClientEventInterchainTransactionReceived(srcTx);
        executeTx(icOptions);
    }

    function test_interchainExecute_state_client() public {
        module.verifyRemoteEntry(moduleEntry, moduleSignatures);
        skip(APP_OPTIMISTIC_PERIOD + 1);
        executeTx(icOptions);
        assertEq(icClient.getExecutor(encodedSrcTx), executor);
        assertEq(icClient.getExecutorById(srcDesc.transactionId), executor);
        vm.expectRevert(
            abi.encodeWithSelector(
                IInterchainClientV1.InterchainClientV1__TxAlreadyExecuted.selector, srcDesc.transactionId
            )
        );
        icClient.isExecutable(encodedSrcTx);
    }

    function test_interchainExecute_state_db() public {
        module.verifyRemoteEntry(moduleEntry, moduleSignatures);
        skip(APP_OPTIMISTIC_PERIOD + 1);
        executeTx(icOptions);
        checkDatabaseStateMsgSent(dstFullEntry, DST_INITIAL_DB_NONCE);
    }

    function test_interchainExecute_state_execService() public {
        module.verifyRemoteEntry(moduleEntry, moduleSignatures);
        skip(APP_OPTIMISTIC_PERIOD + 1);
        executeTx(icOptions);
        assertEq(address(executionService).balance, dstExecutionFee);
    }

    function test_interchainExecute_state_legacyPingPong() public {
        module.verifyRemoteEntry(moduleEntry, moduleSignatures);
        skip(APP_OPTIMISTIC_PERIOD + 1);
        executeTx(icOptions);
        assertEq(address(dstLegacyPingPong()).balance, PING_PONG_BALANCE - dstPingFee);
    }

    function test_interchainExecute_state_synapseModule() public {
        module.verifyRemoteEntry(moduleEntry, moduleSignatures);
        skip(APP_OPTIMISTIC_PERIOD + 1);
        executeTx(icOptions);
        assertEq(address(module).balance, dstVerificationFee);
    }

    function test_interchainExecute_revert_notConfirmed() public {
        // No module signatures
        expectClientRevertResponsesAmountBelowMin({actual: 0, required: 1});
        executeTx(icOptions);
    }

    function test_interchainExecute_revert_confirmed_sameBlock() public {
        module.verifyRemoteEntry(moduleEntry, moduleSignatures);
        expectClientRevertResponsesAmountBelowMin({actual: 0, required: 1});
        executeTx(icOptions);
    }

    function test_interchainExecute_revert_confirmed_periodMinusOneSecond() public {
        module.verifyRemoteEntry(moduleEntry, moduleSignatures);
        skip(APP_OPTIMISTIC_PERIOD);
        expectClientRevertResponsesAmountBelowMin({actual: 0, required: 1});
        executeTx(icOptions);
    }

    /// @notice MessageBus doesn't opt in for the guard, so the guard conflict is not checked
    function test_interchainExecute_state_db_guardMarked() public {
        markInvalidByGuard(srcFullEntry);
        test_interchainExecute_state_db();
    }

    function test_interchainExecute_state_execService_guardMarked() public {
        markInvalidByGuard(srcFullEntry);
        test_interchainExecute_state_execService();
    }

    function test_interchainExecute_state_legacyPingPong_guardMarked() public {
        markInvalidByGuard(srcFullEntry);
        test_interchainExecute_state_legacyPingPong();
    }

    function test_interchainExecute_state_synapseModule_guardMarked() public {
        markInvalidByGuard(srcFullEntry);
        test_interchainExecute_state_synapseModule();
    }

    function test_interchainExecute_revert_notConfirmed_guardMarked() public {
        markInvalidByGuard(srcFullEntry);
        test_interchainExecute_revert_notConfirmed();
    }

    function test_interchainExecute_revert_confirmed_sameBlock_guardMarked() public {
        markInvalidByGuard(srcFullEntry);
        test_interchainExecute_revert_confirmed_sameBlock();
    }

    function test_interchainExecute_revert_confirmed_periodMinusOneSecond_guardMarked() public {
        markInvalidByGuard(srcFullEntry);
        test_interchainExecute_revert_confirmed_periodMinusOneSecond();
    }

    function test_interchainExecute_revert_alreadyExecuted() public {
        module.verifyRemoteEntry(moduleEntry, moduleSignatures);
        skip(APP_OPTIMISTIC_PERIOD + 1);
        executeTx(icOptions);
        expectClientRevertTxAlreadyExecuted(srcDesc);
        executeTx(icOptions);
    }

    function test_isExecutable() public {
        module.verifyRemoteEntry(moduleEntry, moduleSignatures);
        skip(APP_OPTIMISTIC_PERIOD + 1);
        assertTrue(icClient.isExecutable(encodedSrcTx));
    }

    function test_isExecutable_revert_notConfirmed() public {
        expectClientRevertResponsesAmountBelowMin({actual: 0, required: 1});
        icClient.isExecutable(encodedSrcTx);
    }

    function test_isExecutable_revert_notConfirmed_guardMarked() public {
        markInvalidByGuard(srcFullEntry);
        test_isExecutable_revert_notConfirmed();
    }

    function test_isExecutable_revert_confirmed_sameBlock() public {
        module.verifyRemoteEntry(moduleEntry, moduleSignatures);
        expectClientRevertResponsesAmountBelowMin({actual: 0, required: 1});
        icClient.isExecutable(encodedSrcTx);
    }

    function test_isExecutable_revert_confirmed_sameBlock_guardMarked() public {
        markInvalidByGuard(srcFullEntry);
        test_isExecutable_revert_confirmed_sameBlock();
    }

    function test_isExecutable_revert_confirmed_periodMinusOneSecond() public {
        module.verifyRemoteEntry(moduleEntry, moduleSignatures);
        skip(APP_OPTIMISTIC_PERIOD);
        expectClientRevertResponsesAmountBelowMin({actual: 0, required: 1});
        icClient.isExecutable(encodedSrcTx);
    }

    function test_isExecutable_revert_confirmed_periodMinusOneSecond_guardMarked() public {
        markInvalidByGuard(srcFullEntry);
        test_isExecutable_revert_confirmed_periodMinusOneSecond();
    }

    function test_isExecutable_revert_alreadyExecuted() public {
        module.verifyRemoteEntry(moduleEntry, moduleSignatures);
        skip(APP_OPTIMISTIC_PERIOD + 1);
        executeTx(icOptions);
        expectClientRevertTxAlreadyExecuted(srcDesc);
        icClient.isExecutable(encodedSrcTx);
    }

    function localChainId() internal pure override returns (uint64) {
        return DST_CHAIN_ID;
    }

    function remoteChainId() internal pure override returns (uint64) {
        return SRC_CHAIN_ID;
    }
}
