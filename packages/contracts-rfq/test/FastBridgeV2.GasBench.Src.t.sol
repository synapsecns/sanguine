// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {FastBridgeV2, FastBridgeV2SrcBaseTest} from "./FastBridgeV2.Src.Base.t.sol";

// solhint-disable func-name-mixedcase, ordering
/// @notice This test is used to estimate the gas cost of FastBridgeV2 source chain functions.
/// Very little state checks are performed, make sure to do full coverage in different tests.
contract FastBridgeV2GasBenchmarkSrcTest is FastBridgeV2SrcBaseTest {
    uint256 public constant BLOCK_TIME = 12 seconds;

    function createFixtures() public virtual override {
        super.createFixtures();
        // Make both nonces 0 to simplify the tests
        ethTx.nonce = 0;
    }

    function skipBlocksExactly(uint256 blocks) public {
        vm.roll(block.number + blocks);
        vm.warp(block.timestamp + blocks * BLOCK_TIME);
    }

    function rewindBlocksExactly(uint256 blocks) public {
        vm.roll(block.number - blocks);
        vm.warp(block.timestamp - blocks * BLOCK_TIME);
    }

    function skipTimeAtLeast(uint256 time) public {
        uint256 blocksToSkip = time / BLOCK_TIME;
        if (blocksToSkip * BLOCK_TIME < time) blocksToSkip++;
        skipBlocksExactly(blocksToSkip);
    }

    function rewindTimeAtLeast(uint256 time) public {
        uint256 blocksToRewind = time / BLOCK_TIME;
        if (blocksToRewind * BLOCK_TIME < time) blocksToRewind++;
        rewindBlocksExactly(blocksToRewind);
    }

    // ═══════════════════════════════════════════════════ TOKEN ═══════════════════════════════════════════════════════

    function test_bridge_token() public {
        bridge({caller: userA, msgValue: 0, params: tokenParams});
        assertEq(srcToken.balanceOf(userA), LEFTOVER_BALANCE);
        assertEq(srcToken.balanceOf(address(fastBridge)), INITIAL_PROTOCOL_FEES_TOKEN + tokenParams.originAmount);
        skipBlocksExactly(1);
    }

    function test_prove_token() public {
        test_bridge_token();
        prove({caller: relayerA, bridgeTx: tokenTx, destTxHash: hex"01"});
        (uint96 timestamp, address relayer) = fastBridge.bridgeProofs(getTxId(tokenTx));
        assertEq(timestamp, block.timestamp);
        assertEq(relayer, relayerA);
        skipTimeAtLeast(CLAIM_DELAY + 1);
    }

    function test_proveWithAddress_token() public {
        test_bridge_token();
        prove({caller: relayerB, transactionId: getTxId(tokenTx), destTxHash: hex"01", relayer: relayerA});
        (uint96 timestamp, address relayer) = fastBridge.bridgeProofs(getTxId(tokenTx));
        assertEq(timestamp, block.timestamp);
        assertEq(relayer, relayerA);
        skipTimeAtLeast(CLAIM_DELAY + 1);
    }

    function test_claim_token() public {
        test_prove_token();
        claim({caller: relayerA, bridgeTx: tokenTx});
        assertEq(srcToken.balanceOf(relayerA), tokenTx.originAmount);
        assertEq(srcToken.balanceOf(address(fastBridge)), INITIAL_PROTOCOL_FEES_TOKEN + tokenTx.originFeeAmount);
    }

    function test_claimWithAddress_token() public {
        test_proveWithAddress_token();
        claim({caller: relayerA, bridgeTx: tokenTx, to: relayerB});
        assertEq(srcToken.balanceOf(relayerB), tokenTx.originAmount);
        assertEq(srcToken.balanceOf(address(fastBridge)), INITIAL_PROTOCOL_FEES_TOKEN + tokenTx.originFeeAmount);
    }

    function test_dispute_token() public {
        test_prove_token();
        rewindTimeAtLeast({time: CLAIM_DELAY / 2});
        dispute({caller: guard, txId: getTxId(tokenTx)});
        assertEq(fastBridge.bridgeStatuses(getTxId(tokenTx)), FastBridgeV2.BridgeStatus.REQUESTED);
    }

    function test_refundPermissioned_token() public {
        test_bridge_token();
        skipTimeAtLeast({time: DEADLINE});
        refund({caller: refunder, bridgeTx: tokenTx});
        assertEq(srcToken.balanceOf(userA), LEFTOVER_BALANCE + tokenParams.originAmount);
        assertEq(srcToken.balanceOf(address(fastBridge)), INITIAL_PROTOCOL_FEES_TOKEN);
    }

    function test_refundPermissionless_token() public {
        test_bridge_token();
        skipTimeAtLeast({time: DEADLINE + PERMISSIONLESS_REFUND_DELAY});
        refund({caller: userB, bridgeTx: tokenTx});
        assertEq(srcToken.balanceOf(userA), LEFTOVER_BALANCE + tokenParams.originAmount);
        assertEq(srcToken.balanceOf(address(fastBridge)), INITIAL_PROTOCOL_FEES_TOKEN);
    }

    // ════════════════════════════════════════════════════ ETH ════════════════════════════════════════════════════════

    function test_bridge_eth() public {
        bridge({caller: userA, msgValue: ethParams.originAmount, params: ethParams});
        assertEq(userA.balance, LEFTOVER_BALANCE);
        assertEq(address(fastBridge).balance, INITIAL_PROTOCOL_FEES_ETH + ethParams.originAmount);
        skipBlocksExactly(1);
    }

    function test_prove_eth() public {
        test_bridge_eth();
        prove({caller: relayerA, bridgeTx: ethTx, destTxHash: hex"01"});
        (uint96 timestamp, address relayer) = fastBridge.bridgeProofs(getTxId(ethTx));
        assertEq(timestamp, block.timestamp);
        assertEq(relayer, relayerA);
        skipTimeAtLeast(CLAIM_DELAY + 1);
    }

    function test_proveWithAddress_eth() public {
        test_bridge_eth();
        prove({caller: relayerB, transactionId: getTxId(ethTx), destTxHash: hex"01", relayer: relayerA});
        (uint96 timestamp, address relayer) = fastBridge.bridgeProofs(getTxId(ethTx));
        assertEq(timestamp, block.timestamp);
        assertEq(relayer, relayerA);
        skipTimeAtLeast(CLAIM_DELAY + 1);
    }

    function test_claim_eth() public {
        test_prove_eth();
        claim({caller: relayerA, bridgeTx: ethTx});
        assertEq(relayerA.balance, ethTx.originAmount);
        assertEq(address(fastBridge).balance, INITIAL_PROTOCOL_FEES_ETH + ethTx.originFeeAmount);
    }

    function test_claimWithAddress_eth() public {
        test_proveWithAddress_eth();
        claim({caller: relayerA, bridgeTx: ethTx, to: relayerB});
        assertEq(relayerB.balance, ethTx.originAmount);
        assertEq(address(fastBridge).balance, INITIAL_PROTOCOL_FEES_ETH + ethTx.originFeeAmount);
    }

    function test_dispute_eth() public {
        test_prove_eth();
        rewindTimeAtLeast({time: CLAIM_DELAY / 2});
        dispute({caller: guard, txId: getTxId(ethTx)});
        assertEq(fastBridge.bridgeStatuses(getTxId(ethTx)), FastBridgeV2.BridgeStatus.REQUESTED);
    }

    function test_refundPermissioned_eth() public {
        test_bridge_eth();
        skipTimeAtLeast({time: DEADLINE});
        refund({caller: refunder, bridgeTx: ethTx});
        assertEq(userA.balance, LEFTOVER_BALANCE + ethParams.originAmount);
        assertEq(address(fastBridge).balance, INITIAL_PROTOCOL_FEES_ETH);
    }

    function test_refundPermissionless_eth() public {
        test_bridge_eth();
        skipTimeAtLeast({time: DEADLINE + PERMISSIONLESS_REFUND_DELAY});
        refund({caller: userB, bridgeTx: ethTx});
        assertEq(userA.balance, LEFTOVER_BALANCE + ethParams.originAmount);
        assertEq(address(fastBridge).balance, INITIAL_PROTOCOL_FEES_ETH);
    }
}
