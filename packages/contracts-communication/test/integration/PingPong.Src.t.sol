// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {
    ICIntegrationTest,
    InterchainBatch,
    InterchainEntry,
    InterchainTransaction,
    InterchainTxDescriptor
} from "./ICIntegration.t.sol";

// solhint-disable func-name-mixedcase
contract PingPongSrcIntegrationTest is ICIntegrationTest {
    InterchainBatch public batch;
    InterchainEntry public entry;
    InterchainTransaction public icTx;
    InterchainTxDescriptor public desc;

    uint256 public pingFee;
    uint256 public verificationFee;
    uint256 public executionFee;

    function setUp() public override {
        super.setUp();
        icTx = getSrcTransaction();
        entry = getSrcInterchainEntry();
        desc = getInterchainTxDescriptor(entry);
        batch = getInterchainBatch(entry);
        pingFee = pingPongApp.getPingFee(DST_CHAIN_ID);
        verificationFee = icDB.getInterchainFee(DST_CHAIN_ID, toArray(address(module)));
        executionFee = executionService.getExecutionFee({
            dstChainId: DST_CHAIN_ID,
            txPayloadSize: abi.encode(icTx).length,
            options: ppOptions.encodeOptionsV1()
        });
    }

    function test_startPingPong_events() public {
        expectDatabaseEventInterchainEntryWritten(entry);
        expectModuleEventBatchVerificationRequested(batch);
        expectDatabaseEventInterchainBatchVerificationRequested(batch);
        expectFeesEventExecutionFeeAdded(desc.transactionId, executionFee);
        expectServiceEventExecutionRequested(desc.transactionId);
        expectClientEventInterchainTransactionSent(icTx, verificationFee, executionFee);
        expectPingPongEventPingSent(COUNTER, desc);
        pingPongApp.startPingPong(DST_CHAIN_ID, COUNTER);
    }

    function localChainId() internal pure override returns (uint256) {
        return SRC_CHAIN_ID;
    }

    function remoteChainId() internal pure override returns (uint256) {
        return DST_CHAIN_ID;
    }
}
