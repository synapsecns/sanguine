// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {FastBridgeV2DstTest, IFastBridgeV2} from "./FastBridgeV2.Dst.t.sol";

// solhint-disable func-name-mixedcase, ordering
contract FastBridgeV2DstExclusivityTest is FastBridgeV2DstTest {
    uint256 public constant EXCLUSIVITY_PERIOD = 60 seconds;

    function createFixturesV2() public virtual override {
        tokenParamsV2 = IFastBridgeV2.BridgeParamsV2({
            quoteRelayer: relayerA,
            quoteExclusivitySeconds: int256(EXCLUSIVITY_PERIOD),
            quoteId: "",
            callParams: ""
        });
        ethParamsV2 = IFastBridgeV2.BridgeParamsV2({
            quoteRelayer: relayerB,
            quoteExclusivitySeconds: int256(EXCLUSIVITY_PERIOD),
            quoteId: "",
            callParams: ""
        });

        tokenTx.exclusivityRelayer = relayerA;
        tokenTx.exclusivityEndTime = block.timestamp + EXCLUSIVITY_PERIOD;
        ethTx.exclusivityRelayer = relayerB;
        ethTx.exclusivityEndTime = block.timestamp + EXCLUSIVITY_PERIOD;
    }

    // ═══════════════════════════════════════════════ RELAY: TOKEN ════════════════════════════════════════════════════

    // Relayer A has the exclusivity fill rights within the EXCLUSIVITY_PERIOD

    function test_relay_token_exclusivityLastSecond() public {
        skip(EXCLUSIVITY_PERIOD);
        test_relay_token();
    }

    function test_relay_token_exclusivityOver() public {
        skip(EXCLUSIVITY_PERIOD + 1);
        test_relay_token();
    }

    function test_relay_token_notQuotedRelayer_revert() public {
        vm.expectRevert(ExclusivityPeriodNotPassed.selector);
        relay({caller: relayerB, msgValue: 0, bridgeTx: tokenTx});
    }

    function test_relay_token_notQuotedRelayer_exclusivityLastSecond_revert() public {
        skip(EXCLUSIVITY_PERIOD);
        vm.expectRevert(ExclusivityPeriodNotPassed.selector);
        relay({caller: relayerB, msgValue: 0, bridgeTx: tokenTx});
    }

    function test_relay_token_notQuotedRelayer_exclusivityOver() public {
        skip(EXCLUSIVITY_PERIOD + 1);
        bytes32 txId = getTxId(tokenTx);
        expectBridgeRelayed(tokenTx, txId, address(relayerB));
        relay({caller: relayerB, msgValue: 0, bridgeTx: tokenTx});
        assertTrue(fastBridge.bridgeRelays(txId));
        assertEq(dstToken.balanceOf(address(userB)), tokenParams.destAmount);
        assertEq(dstToken.balanceOf(address(relayerB)), LEFTOVER_BALANCE);
        assertEq(dstToken.balanceOf(address(fastBridge)), 0);
    }

    function test_relay_token_withRelayerAddress_exclusivityLastSecond() public {
        skip(EXCLUSIVITY_PERIOD);
        test_relay_token_withRelayerAddress();
    }

    function test_relay_token_withRelayerAddress_exclusivityOver() public {
        skip(EXCLUSIVITY_PERIOD + 1);
        test_relay_token_withRelayerAddress();
    }

    function test_relay_token_withRelayerAddress_notQuotedRelayer_revert() public {
        vm.expectRevert(ExclusivityPeriodNotPassed.selector);
        relayWithAddress({caller: relayerA, relayer: relayerB, msgValue: 0, bridgeTx: tokenTx});
    }

    function test_relay_token_withRelayerAddress_notQuotedRelayer_exclusivityLastSecond_revert() public {
        skip(EXCLUSIVITY_PERIOD);
        vm.expectRevert(ExclusivityPeriodNotPassed.selector);
        relayWithAddress({caller: relayerA, relayer: relayerB, msgValue: 0, bridgeTx: tokenTx});
    }

    function test_relay_token_withRelayerAddress_notQuotedRelayer_exclusivityOver() public {
        skip(EXCLUSIVITY_PERIOD + 1);
        bytes32 txId = getTxId(tokenTx);
        expectBridgeRelayed(tokenTx, txId, address(relayerB));
        relayWithAddress({caller: relayerA, relayer: relayerB, msgValue: 0, bridgeTx: tokenTx});
        assertTrue(fastBridge.bridgeRelays(txId));
        assertEq(dstToken.balanceOf(address(userB)), tokenParams.destAmount);
        assertEq(dstToken.balanceOf(address(relayerA)), LEFTOVER_BALANCE);
        assertEq(dstToken.balanceOf(address(fastBridge)), 0);
    }

    // ════════════════════════════════════════════════ RELAY: ETH ═════════════════════════════════════════════════════

    // Relayer B has the exclusivity fill rights within the EXCLUSIVITY_PERIOD

    function test_relay_eth_exclusivityLastSecond() public {
        skip(EXCLUSIVITY_PERIOD);
        test_relay_eth();
    }

    function test_relay_eth_exclusivityOver() public {
        skip(EXCLUSIVITY_PERIOD + 1);
        test_relay_eth();
    }

    function test_relay_eth_notQuotedRelayer_revert() public {
        vm.expectRevert(ExclusivityPeriodNotPassed.selector);
        relay({caller: relayerA, msgValue: ethParams.destAmount, bridgeTx: ethTx});
    }

    function test_relay_eth_notQuotedRelayer_exclusivityLastSecond_revert() public {
        skip(EXCLUSIVITY_PERIOD);
        vm.expectRevert(ExclusivityPeriodNotPassed.selector);
        relay({caller: relayerA, msgValue: ethParams.destAmount, bridgeTx: ethTx});
    }

    function test_relay_eth_notQuotedRelayer_exclusivityOver() public {
        skip(EXCLUSIVITY_PERIOD + 1);
        bytes32 txId = getTxId(ethTx);
        expectBridgeRelayed(ethTx, txId, address(relayerA));
        relay({caller: relayerA, msgValue: ethParams.destAmount, bridgeTx: ethTx});
        assertTrue(fastBridge.bridgeRelays(txId));
        assertEq(address(userB).balance, ethParams.destAmount);
        assertEq(address(relayerA).balance, LEFTOVER_BALANCE);
        assertEq(address(fastBridge).balance, 0);
    }

    function test_relay_eth_withRelayerAddress_exclusivityLastSecond() public {
        skip(EXCLUSIVITY_PERIOD);
        test_relay_eth_withRelayerAddress();
    }

    function test_relay_eth_withRelayerAddress_exclusivityOver() public {
        skip(EXCLUSIVITY_PERIOD + 1);
        test_relay_eth_withRelayerAddress();
    }

    function test_relay_eth_withRelayerAddress_notQuotedRelayer_revert() public {
        vm.expectRevert(ExclusivityPeriodNotPassed.selector);
        relayWithAddress({caller: relayerB, relayer: relayerA, msgValue: ethParams.destAmount, bridgeTx: ethTx});
    }

    function test_relay_eth_withRelayerAddress_notQuotedRelayer_exclusivityLastSecond_revert() public {
        skip(EXCLUSIVITY_PERIOD);
        vm.expectRevert(ExclusivityPeriodNotPassed.selector);
        relayWithAddress({caller: relayerB, relayer: relayerA, msgValue: ethParams.destAmount, bridgeTx: ethTx});
    }

    function test_relay_eth_withRelayerAddress_notQuotedRelayer_exclusivityOver() public {
        skip(EXCLUSIVITY_PERIOD + 1);
        bytes32 txId = getTxId(ethTx);
        expectBridgeRelayed(ethTx, txId, address(relayerA));
        relayWithAddress({caller: relayerB, relayer: relayerA, msgValue: ethParams.destAmount, bridgeTx: ethTx});
        assertTrue(fastBridge.bridgeRelays(txId));
        assertEq(address(userB).balance, ethParams.destAmount);
        assertEq(address(relayerB).balance, LEFTOVER_BALANCE);
        assertEq(address(fastBridge).balance, 0);
    }
}
