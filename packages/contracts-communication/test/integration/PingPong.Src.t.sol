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
// solhint-disable ordering
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
            txPayloadSize: getEncodedTx(icTx).length,
            options: ppOptions.encodeOptionsV1()
        });
    }

    function test_startPingPong_events() public {
        expectEventsPingSent(COUNTER, icTx, entry, verificationFee, executionFee);
        pingPongApp.startPingPong(DST_CHAIN_ID, COUNTER);
    }

    function test_startPingPong_state_db() public {
        pingPongApp.startPingPong(DST_CHAIN_ID, COUNTER);
        checkDatabaseStatePingSent(entry, SRC_INITIAL_DB_NONCE);
    }

    function test_startPingPong_state_execFees() public {
        pingPongApp.startPingPong(DST_CHAIN_ID, COUNTER);
        assertEq(address(executionFees).balance, executionFee);
        assertEq(executionFees.executionFee(DST_CHAIN_ID, desc.transactionId), executionFee);
    }

    function test_startPingPong_state_pingPongApp() public {
        pingPongApp.startPingPong(DST_CHAIN_ID, COUNTER);
        assertEq(address(pingPongApp).balance, PING_PONG_BALANCE - pingFee);
    }

    function test_startPingPong_state_synapseModule() public {
        pingPongApp.startPingPong(DST_CHAIN_ID, COUNTER);
        assertEq(address(module).balance, verificationFee);
    }

    function localChainId() internal pure override returns (uint64) {
        return SRC_CHAIN_ID;
    }

    function remoteChainId() internal pure override returns (uint64) {
        return DST_CHAIN_ID;
    }
}
