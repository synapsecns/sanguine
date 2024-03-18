// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {ModuleBatchLib} from "../../contracts/libs/ModuleBatch.sol";

import {
    ICIntegrationTest,
    InterchainBatch,
    InterchainEntry,
    InterchainTransaction,
    InterchainTxDescriptor
} from "./ICIntegration.t.sol";

// solhint-disable func-name-mixedcase
// solhint-disable ordering
contract PingPongDstIntegrationTest is ICIntegrationTest {
    uint256 public constant LONG_PERIOD = 1 weeks;

    InterchainBatch public srcBatch;
    InterchainEntry public srcEntry;
    InterchainTransaction public srcTx;
    InterchainTxDescriptor public srcDesc;

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

    function localChainId() internal pure override returns (uint256) {
        return DST_CHAIN_ID;
    }

    function remoteChainId() internal pure override returns (uint256) {
        return SRC_CHAIN_ID;
    }
}
