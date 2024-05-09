// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {InterchainTransaction, InterchainTxDescriptor} from "../../contracts/libs/InterchainTransaction.sol";

import {PingPongIntegrationTest} from "./PingPong.t.sol";

// solhint-disable func-name-mixedcase
// solhint-disable ordering
contract PingPongSrcIntegrationTest is PingPongIntegrationTest {
    FullEntry public fullEntry;
    InterchainTransaction public icTx;
    InterchainTxDescriptor public desc;

    uint256 public pingFee;
    uint256 public verificationFee;
    uint256 public executionFee;

    function setUp() public override {
        super.setUp();
        icTx = getSrcTransaction();
        fullEntry = getSrcFullEntry();
        desc = getInterchainTxDescriptor(fullEntry);
        pingFee = srcPingPongApp().getPingFee(DST_CHAIN_ID);
        verificationFee = icDB.getInterchainFee(DST_CHAIN_ID, toArray(address(module)));
        executionFee = executionService.getExecutionFee({
            dstChainId: DST_CHAIN_ID,
            txPayloadSize: getEncodedTx(icTx).length,
            options: ppOptions.encodeOptionsV1()
        });
    }

    function test_startPingPong_events() public {
        expectEventsPingSent(COUNTER, icTx, fullEntry, verificationFee, executionFee);
        srcPingPongApp().startPingPong(DST_CHAIN_ID, COUNTER);
    }

    function test_startPingPong_state_db() public {
        srcPingPongApp().startPingPong(DST_CHAIN_ID, COUNTER);
        checkDatabaseStateMsgSent(fullEntry, SRC_INITIAL_DB_NONCE);
    }

    function test_startPingPong_state_execService() public {
        srcPingPongApp().startPingPong(DST_CHAIN_ID, COUNTER);
        assertEq(address(executionService).balance, executionFee);
    }

    function test_startPingPong_state_pingPongApp() public {
        srcPingPongApp().startPingPong(DST_CHAIN_ID, COUNTER);
        assertEq(srcApp.balance, PING_PONG_BALANCE - pingFee);
    }

    function test_startPingPong_state_synapseModule() public {
        srcPingPongApp().startPingPong(DST_CHAIN_ID, COUNTER);
        assertEq(address(module).balance, verificationFee);
    }

    function localChainId() internal pure override returns (uint64) {
        return SRC_CHAIN_ID;
    }

    function remoteChainId() internal pure override returns (uint64) {
        return DST_CHAIN_ID;
    }
}
