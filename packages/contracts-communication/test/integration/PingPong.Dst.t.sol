// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {IInterchainClientV1} from "../../contracts/interfaces/IInterchainClientV1.sol";
import {ModuleBatchLib} from "../../contracts/libs/ModuleBatch.sol";
import {OptionsV1} from "../../contracts/libs/Options.sol";

import {
    PingPongIntegrationTest,
    InterchainBatch,
    InterchainEntry,
    InterchainTransaction,
    InterchainTxDescriptor,
    PingPongApp
} from "./PingPong.t.sol";

// solhint-disable func-name-mixedcase
// solhint-disable ordering
contract PingPongDstIntegrationTest is PingPongIntegrationTest {
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

        dstPingFee = dstPingPongApp().getPingFee(SRC_CHAIN_ID);
        dstVerificationFee = icDB.getInterchainFee(SRC_CHAIN_ID, toArray(address(module)));
        dstExecutionFee = executionService.getExecutionFee({
            dstChainId: SRC_CHAIN_ID,
            txPayloadSize: getEncodedTx(dstTx).length,
            options: ppOptions.encodeOptionsV1()
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

    function test_interchainExecute_callPingPongApp() public {
        module.verifyRemoteBatch(moduleBatch, moduleSignatures);
        skip(APP_OPTIMISTIC_PERIOD + 1);
        expectAppCall(srcTx, ppOptions);
        executeTx(ppOptions);
    }

    function test_interchainExecute_callPingPongApp_lowerGas() public {
        OptionsV1 memory options = OptionsV1({gasLimit: ppOptions.gasLimit / 2, gasAirdrop: ppOptions.gasAirdrop});
        module.verifyRemoteBatch(moduleBatch, moduleSignatures);
        skip(APP_OPTIMISTIC_PERIOD + 1);
        // Should use the requested gas limit
        expectAppCall(srcTx, ppOptions);
        executeTx(options);
    }

    function test_interchainExecute_callPingPongApp_higherGas() public {
        OptionsV1 memory options = OptionsV1({gasLimit: 2 * ppOptions.gasLimit, gasAirdrop: ppOptions.gasAirdrop});
        module.verifyRemoteBatch(moduleBatch, moduleSignatures);
        skip(APP_OPTIMISTIC_PERIOD + 1);
        // Should allow to use higher gas limit
        expectAppCall(srcTx, options);
        executeTx(options);
    }

    function test_interchainExecute_events() public {
        module.verifyRemoteBatch(moduleBatch, moduleSignatures);
        skip(APP_OPTIMISTIC_PERIOD + 1);
        expectPingPongEventPingReceived(COUNTER, srcEntry);
        expectEventsPingSent(COUNTER - 1, dstTx, dstEntry, dstVerificationFee, dstExecutionFee);
        expectClientEventInterchainTransactionReceived(srcTx);
        executeTx(ppOptions);
    }

    function test_interchainExecute_state_client() public {
        module.verifyRemoteBatch(moduleBatch, moduleSignatures);
        skip(APP_OPTIMISTIC_PERIOD + 1);
        executeTx(ppOptions);
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
        executeTx(ppOptions);
        checkDatabaseStateMsgSent(dstEntry, DST_INITIAL_DB_NONCE);
    }

    function test_interchainExecute_state_execService() public {
        module.verifyRemoteBatch(moduleBatch, moduleSignatures);
        skip(APP_OPTIMISTIC_PERIOD + 1);
        executeTx(ppOptions);
        assertEq(address(executionService).balance, dstExecutionFee);
    }

    function test_interchainExecute_state_pingPongApp() public {
        module.verifyRemoteBatch(moduleBatch, moduleSignatures);
        skip(APP_OPTIMISTIC_PERIOD + 1);
        executeTx(ppOptions);
        assertEq(dstApp.balance, PING_PONG_BALANCE - dstPingFee);
    }

    function test_interchainExecute_state_synapseModule() public {
        module.verifyRemoteBatch(moduleBatch, moduleSignatures);
        skip(APP_OPTIMISTIC_PERIOD + 1);
        executeTx(ppOptions);
        assertEq(address(module).balance, dstVerificationFee);
    }

    function test_interchainExecute_revert_notConfirmed() public {
        // No module signatures
        expectClientRevertNotEnoughResponses({actual: 0, required: 1});
        executeTx(ppOptions);
    }

    function test_interchainExecute_revert_notConfirmed_guardMarked() public {
        markInvalidByGuard(srcBatch);
        expectClientRevertBatchConflict(guard);
        executeTx(ppOptions);
    }

    function test_interchainExecute_revert_confirmed_sameBlock() public {
        module.verifyRemoteBatch(moduleBatch, moduleSignatures);
        expectClientRevertNotEnoughResponses({actual: 0, required: 1});
        executeTx(ppOptions);
    }

    function test_interchainExecute_revert_confirmed_sameBlock_guardMarked() public {
        module.verifyRemoteBatch(moduleBatch, moduleSignatures);
        markInvalidByGuard(srcBatch);
        expectClientRevertBatchConflict(guard);
        executeTx(ppOptions);
    }

    function test_interchainExecute_revert_confirmed_periodMinusOneSecond() public {
        module.verifyRemoteBatch(moduleBatch, moduleSignatures);
        skip(APP_OPTIMISTIC_PERIOD);
        expectClientRevertNotEnoughResponses({actual: 0, required: 1});
        executeTx(ppOptions);
    }

    function test_interchainExecute_revert_confirmed_periodMinusOneSecond_guardMarked() public {
        module.verifyRemoteBatch(moduleBatch, moduleSignatures);
        markInvalidByGuard(srcBatch);
        skip(APP_OPTIMISTIC_PERIOD);
        expectClientRevertBatchConflict(guard);
        executeTx(ppOptions);
    }

    function test_interchainExecute_revert_alreadyExecuted() public {
        module.verifyRemoteBatch(moduleBatch, moduleSignatures);
        skip(APP_OPTIMISTIC_PERIOD + 1);
        executeTx(ppOptions);
        expectClientRevertTxAlreadyExecuted(srcDesc);
        executeTx(ppOptions);
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
        markInvalidByGuard(srcBatch);
        skip(APP_OPTIMISTIC_PERIOD);
        expectClientRevertBatchConflict(guard);
        icClient.isExecutable(encodedSrcTx, new bytes32[](0));
    }

    function test_isExecutable_revert_alreadyExecuted() public {
        module.verifyRemoteBatch(moduleBatch, moduleSignatures);
        skip(APP_OPTIMISTIC_PERIOD + 1);
        executeTx(ppOptions);
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
