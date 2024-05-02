// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {InterchainBatch} from "../../../contracts/libs/InterchainBatch.sol";
import {InterchainEntry} from "../../../contracts/libs/InterchainEntry.sol";
import {InterchainTransaction, InterchainTxDescriptor} from "../../../contracts/libs/InterchainTransaction.sol";

import {LegacyPingPongIntegrationTest} from "./LegacyPingPong.t.sol";

contract LegacyPingPongSrcIntegrationTest is LegacyPingPongIntegrationTest {
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
        pingFee = srcLegacyPingPong().getPingFee(DST_CHAIN_ID);
        verificationFee = icDB.getInterchainFee(DST_CHAIN_ID, toArray(address(module)));
        executionFee = executionService.getExecutionFee({
            dstChainId: DST_CHAIN_ID,
            txPayloadSize: getEncodedTx(icTx).length,
            options: icOptions.encodeOptionsV1()
        });
    }

    function test_startPingPong_events() public {
        expectEventsPingSent(COUNTER, icTx, entry, verificationFee, executionFee);
        srcLegacyPingPong().startPingPong(DST_CHAIN_ID, COUNTER);
    }

    function test_startPingPong_state_db() public {
        srcLegacyPingPong().startPingPong(DST_CHAIN_ID, COUNTER);
        checkDatabaseStateMsgSent(entry, SRC_INITIAL_DB_NONCE);
    }

    function test_startPingPong_state_execService() public {
        srcLegacyPingPong().startPingPong(DST_CHAIN_ID, COUNTER);
        assertEq(address(executionService).balance, executionFee);
    }

    function test_startPingPong_state_legacyPingPong() public {
        srcLegacyPingPong().startPingPong(DST_CHAIN_ID, COUNTER);
        assertEq(srcPingPong.balance, PING_PONG_BALANCE - pingFee);
    }

    function test_startPingPong_state_synapseModule() public {
        srcLegacyPingPong().startPingPong(DST_CHAIN_ID, COUNTER);
        assertEq(address(module).balance, verificationFee);
    }

    function localChainId() internal pure override returns (uint64) {
        return SRC_CHAIN_ID;
    }

    function remoteChainId() internal pure override returns (uint64) {
        return DST_CHAIN_ID;
    }
}
