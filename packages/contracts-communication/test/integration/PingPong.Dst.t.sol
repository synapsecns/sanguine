// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {IInterchainClientV1} from "../../contracts/interfaces/IInterchainClientV1.sol";
import {ModuleBatchLib} from "../../contracts/libs/ModuleBatch.sol";

import {
    ICIntegrationTest,
    InterchainBatch,
    InterchainEntry,
    InterchainTransaction,
    InterchainTxDescriptor,
    OptionsV1
} from "./ICIntegration.t.sol";

// solhint-disable func-name-mixedcase
// solhint-disable ordering
contract PingPongDstIntegrationTest is ICIntegrationTest {
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
        encodedSrcTx = abi.encode(srcTx);

        moduleBatch = getModuleBatch(srcBatch);
        moduleSignatures = getModuleSignatures(srcBatch);

        dstTx = getDstTransaction();
        dstEntry = getDstInterchainEntry();
        dstDesc = getInterchainTxDescriptor(dstEntry);
        dstBatch = getInterchainBatch(dstEntry);

        dstPingFee = pingPongApp.getPingFee(SRC_CHAIN_ID);
        dstVerificationFee = icDB.getInterchainFee(SRC_CHAIN_ID, toArray(address(module)));
        dstExecutionFee = executionService.getExecutionFee({
            dstChainId: SRC_CHAIN_ID,
            txPayloadSize: abi.encode(dstTx).length,
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
        assertEq(icDB.checkVerification(address(module), srcEntry, new bytes32[](0)), INITIAL_TS);
    }

    function test_interchainExecute_events() public {
        module.verifyRemoteBatch(moduleBatch, moduleSignatures);
        skip(APP_OPTIMISTIC_PERIOD + 1);
        expectPingPongEventPingReceived(COUNTER, srcDesc);
        expectEventsPingSent(COUNTER - 1, dstTx, dstEntry, dstVerificationFee, dstExecutionFee);
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
        checkDatabaseStatePingSent(dstEntry, DST_INITIAL_DB_NONCE);
    }

    function test_interchainExecute_state_execFees() public {
        module.verifyRemoteBatch(moduleBatch, moduleSignatures);
        skip(APP_OPTIMISTIC_PERIOD + 1);
        executeTx(ppOptions);
        assertEq(address(executionFees).balance, dstExecutionFee);
        assertEq(executionFees.executionFee(SRC_CHAIN_ID, dstDesc.transactionId), dstExecutionFee);
    }

    function test_interchainExecute_state_pingPongApp() public {
        module.verifyRemoteBatch(moduleBatch, moduleSignatures);
        skip(APP_OPTIMISTIC_PERIOD + 1);
        executeTx(ppOptions);
        assertEq(address(pingPongApp).balance, PING_PONG_BALANCE - dstPingFee);
    }

    function test_interchainExecute_state_synapseModule() public {
        module.verifyRemoteBatch(moduleBatch, moduleSignatures);
        skip(APP_OPTIMISTIC_PERIOD + 1);
        executeTx(ppOptions);
        assertEq(address(module).balance, dstVerificationFee);
    }

    function localChainId() internal pure override returns (uint256) {
        return DST_CHAIN_ID;
    }

    function remoteChainId() internal pure override returns (uint256) {
        return SRC_CHAIN_ID;
    }
}
