// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {IInterchainClientV1} from "../../../contracts/interfaces/IInterchainClientV1.sol";
import {InterchainBatch} from "../../../contracts/libs/InterchainBatch.sol";
import {InterchainEntry} from "../../../contracts/libs/InterchainEntry.sol";
import {InterchainTransaction, InterchainTxDescriptor} from "../../../contracts/libs/InterchainTransaction.sol";
import {OptionsV1} from "../../../contracts/libs/Options.sol";

import {LegacyPingPongIntegrationTest} from "./LegacyPingPong.t.sol";

// solhint-disable func-name-mixedcase
// solhint-disable ordering
contract LegacyPingPongDstIntegrationTest is LegacyPingPongIntegrationTest {
    uint256 public constant LONG_PERIOD = 1 weeks;

    InterchainBatch public srcBatch;
    InterchainEntry public srcEntry;
    InterchainTransaction public srcTx;
    InterchainTxDescriptor public srcDesc;
    bytes public encodedSrcTx;

    InterchainBatch public dstBatch;
    InterchainEntry public dstEntry;
    InterchainTransaction public dstTx;
    InterchainTxDescriptor public dstDesc;

    uint256 public dstPingFee;
    uint256 public dstVerificationFee;
    uint256 public dstExecutionFee;

    bytes public moduleBatch;
    bytes public moduleSignatures;

    function setUp() public override {
        super.setUp();
        srcTx = getSrcTransaction();
        srcEntry = getSrcInterchainEntry();
        srcDesc = getInterchainTxDescriptor(srcEntry);
        srcBatch = getInterchainBatch(srcEntry);
        encodedSrcTx = getEncodedTx(srcTx);

        moduleBatch = getModuleBatch(srcBatch);
        moduleSignatures = getModuleSignatures(srcBatch);

        dstTx = getDstTransaction();
        dstEntry = getDstInterchainEntry();
        dstDesc = getInterchainTxDescriptor(dstEntry);
        dstBatch = getInterchainBatch(dstEntry);

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
        icClient.interchainExecute{value: options.gasAirdrop}({
            gasLimit: options.gasLimit,
            transaction: encodedSrcTx,
            proof: new bytes32[](0)
        });
    }

    function test_verifyRemoteBatch_events() public {
        expectDatabaseEventInterchainBatchVerified(srcBatch);
        expectModuleEventBatchVerified(srcBatch);
        module.verifyRemoteBatch(moduleBatch, moduleSignatures);
    }

    function test_verifyRemoteBatch_state_client() public {
        module.verifyRemoteBatch(moduleBatch, moduleSignatures);
        skip(APP_OPTIMISTIC_PERIOD + 1);
        assertEq(icClient.getExecutor(encodedSrcTx), address(0));
        assertEq(icClient.getExecutorById(srcDesc.transactionId), address(0));
        assertTrue(icClient.isExecutable(encodedSrcTx, new bytes32[](0)));
    }

    function test_verifyRemoteBatch_state_db() public {
        module.verifyRemoteBatch(moduleBatch, moduleSignatures);
        skip(LONG_PERIOD);
        assertEq(icDB.checkBatchVerification(address(module), srcBatch), INITIAL_TS);
    }

    function test_interchainExecute_callMessageBusAndLegacyPP() public {
        module.verifyRemoteBatch(moduleBatch, moduleSignatures);
        skip(APP_OPTIMISTIC_PERIOD + 1);
        expectAppCall(srcTx, icOptions);
        expectPingPongCall();
        executeTx(icOptions);
    }

    function test_interchainExecute_callMessageBusAndLegacyPP_lowerGas() public {
        OptionsV1 memory options = OptionsV1({gasLimit: icOptions.gasLimit / 2, gasAirdrop: 0});
        module.verifyRemoteBatch(moduleBatch, moduleSignatures);
        skip(APP_OPTIMISTIC_PERIOD + 1);
        // Should use the requested gas limit
        expectAppCall(srcTx, icOptions);
        expectPingPongCall();
        executeTx(options);
    }

    function test_interchainExecute_callMessageBusAndLegacyPP_higherGas() public {
        OptionsV1 memory options = OptionsV1({gasLimit: 2 * icOptions.gasLimit, gasAirdrop: 0});
        module.verifyRemoteBatch(moduleBatch, moduleSignatures);
        skip(APP_OPTIMISTIC_PERIOD + 1);
        // Should allow to use higher gas limit
        expectAppCall(srcTx, options);
        expectPingPongCall();
        executeTx(options);
    }

    function test_interchainExecute_events() public {
        module.verifyRemoteBatch(moduleBatch, moduleSignatures);
        skip(APP_OPTIMISTIC_PERIOD + 1);
        expectPingPongEventPingReceived(COUNTER);
        expectEventsPingSent(COUNTER - 1, dstTx, dstEntry, dstVerificationFee, dstExecutionFee);
        expectClientEventInterchainTransactionReceived(srcTx);
        executeTx(icOptions);
    }

    function test_interchainExecute_state_client() public {
        module.verifyRemoteBatch(moduleBatch, moduleSignatures);
        skip(APP_OPTIMISTIC_PERIOD + 1);
        executeTx(icOptions);
        assertEq(icClient.getExecutor(encodedSrcTx), executor);
        assertEq(icClient.getExecutorById(srcDesc.transactionId), executor);
        vm.expectRevert(
            abi.encodeWithSelector(
                IInterchainClientV1.InterchainClientV1__TxAlreadyExecuted.selector, srcDesc.transactionId
            )
        );
        icClient.isExecutable(encodedSrcTx, new bytes32[](0));
    }

    function test_interchainExecute_state_db() public {
        module.verifyRemoteBatch(moduleBatch, moduleSignatures);
        skip(APP_OPTIMISTIC_PERIOD + 1);
        executeTx(icOptions);
        checkDatabaseStateMsgSent(dstEntry, DST_INITIAL_DB_NONCE);
    }

    function test_interchainExecute_state_execService() public {
        module.verifyRemoteBatch(moduleBatch, moduleSignatures);
        skip(APP_OPTIMISTIC_PERIOD + 1);
        executeTx(icOptions);
        assertEq(address(executionService).balance, dstExecutionFee);
    }

    function test_interchainExecute_state_legacyPingPong() public {
        module.verifyRemoteBatch(moduleBatch, moduleSignatures);
        skip(APP_OPTIMISTIC_PERIOD + 1);
        executeTx(icOptions);
        assertEq(address(dstLegacyPingPong()).balance, PING_PONG_BALANCE - dstPingFee);
    }

    function test_interchainExecute_state_synapseModule() public {
        module.verifyRemoteBatch(moduleBatch, moduleSignatures);
        skip(APP_OPTIMISTIC_PERIOD + 1);
        executeTx(icOptions);
        assertEq(address(module).balance, dstVerificationFee);
    }

    function test_interchainExecute_revert_notConfirmed() public {
        // No module signatures
        expectClientRevertNotEnoughResponses({actual: 0, required: 1});
        executeTx(icOptions);
    }

    function test_interchainExecute_revert_notConfirmed_guardMarked() public {
        markInvalidByGuard(srcBatch);
        expectClientRevertBatchConflict(guard);
        executeTx(icOptions);
    }

    function test_interchainExecute_revert_confirmed_sameBlock() public {
        module.verifyRemoteBatch(moduleBatch, moduleSignatures);
        expectClientRevertNotEnoughResponses({actual: 0, required: 1});
        executeTx(icOptions);
    }

    function test_interchainExecute_revert_confirmed_sameBlock_guardMarked() public {
        module.verifyRemoteBatch(moduleBatch, moduleSignatures);
        markInvalidByGuard(srcBatch);
        expectClientRevertBatchConflict(guard);
        executeTx(icOptions);
    }

    function test_interchainExecute_revert_confirmed_periodMinusOneSecond() public {
        module.verifyRemoteBatch(moduleBatch, moduleSignatures);
        skip(APP_OPTIMISTIC_PERIOD);
        expectClientRevertNotEnoughResponses({actual: 0, required: 1});
        executeTx(icOptions);
    }

    function test_interchainExecute_revert_confirmed_periodMinusOneSecond_guardMarked() public {
        module.verifyRemoteBatch(moduleBatch, moduleSignatures);
        skip(APP_OPTIMISTIC_PERIOD);
        markInvalidByGuard(srcBatch);
        expectClientRevertBatchConflict(guard);
        executeTx(icOptions);
    }

    function test_interchainExecute_revert_alreadyExecuted() public {
        module.verifyRemoteBatch(moduleBatch, moduleSignatures);
        skip(APP_OPTIMISTIC_PERIOD + 1);
        executeTx(icOptions);
        expectClientRevertTxAlreadyExecuted(srcDesc);
        executeTx(icOptions);
    }

    function test_isExecutable() public {
        module.verifyRemoteBatch(moduleBatch, moduleSignatures);
        skip(APP_OPTIMISTIC_PERIOD + 1);
        assertTrue(icClient.isExecutable(encodedSrcTx, new bytes32[](0)));
    }

    function test_isExecutable_revert_notConfirmed() public {
        expectClientRevertNotEnoughResponses({actual: 0, required: 1});
        icClient.isExecutable(encodedSrcTx, new bytes32[](0));
    }

    function test_isExecutable_revert_notConfirmed_guardMarked() public {
        markInvalidByGuard(srcBatch);
        expectClientRevertBatchConflict(guard);
        icClient.isExecutable(encodedSrcTx, new bytes32[](0));
    }

    function test_isExecutable_revert_confirmed_sameBlock() public {
        module.verifyRemoteBatch(moduleBatch, moduleSignatures);
        expectClientRevertNotEnoughResponses({actual: 0, required: 1});
        icClient.isExecutable(encodedSrcTx, new bytes32[](0));
    }

    function test_isExecutable_revert_confirmed_sameBlock_guardMarked() public {
        module.verifyRemoteBatch(moduleBatch, moduleSignatures);
        markInvalidByGuard(srcBatch);
        expectClientRevertBatchConflict(guard);
        icClient.isExecutable(encodedSrcTx, new bytes32[](0));
    }

    function test_isExecutable_revert_confirmed_periodMinusOneSecond() public {
        module.verifyRemoteBatch(moduleBatch, moduleSignatures);
        skip(APP_OPTIMISTIC_PERIOD);
        expectClientRevertNotEnoughResponses({actual: 0, required: 1});
        icClient.isExecutable(encodedSrcTx, new bytes32[](0));
    }

    function test_isExecutable_revert_confirmed_periodMinusOneSecond_guardMarked() public {
        module.verifyRemoteBatch(moduleBatch, moduleSignatures);
        skip(APP_OPTIMISTIC_PERIOD);
        markInvalidByGuard(srcBatch);
        expectClientRevertBatchConflict(guard);
        icClient.isExecutable(encodedSrcTx, new bytes32[](0));
    }

    function test_isExecutable_revert_alreadyExecuted() public {
        module.verifyRemoteBatch(moduleBatch, moduleSignatures);
        skip(APP_OPTIMISTIC_PERIOD + 1);
        executeTx(icOptions);
        expectClientRevertTxAlreadyExecuted(srcDesc);
        icClient.isExecutable(encodedSrcTx, new bytes32[](0));
    }

    function localChainId() internal pure override returns (uint64) {
        return DST_CHAIN_ID;
    }

    function remoteChainId() internal pure override returns (uint64) {
        return SRC_CHAIN_ID;
    }
}
