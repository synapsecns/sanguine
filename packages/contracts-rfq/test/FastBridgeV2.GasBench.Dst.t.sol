// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {FastBridgeV2DstBaseTest} from "./FastBridgeV2.Dst.Base.t.sol";

// solhint-disable func-name-mixedcase, ordering
/// @notice This test is used to estimate the gas cost of FastBridgeV2 destination chain functions.
/// Very little state checks are performed, make sure to do full coverage in different tests.
contract FastBridgeV2DstGasBenchmarkTest is FastBridgeV2DstBaseTest {
    uint256 public constant INITIAL_USER_BALANCE = 100 ether;

    function mintTokens() public virtual override {
        super.mintTokens();
        deal(userB, INITIAL_USER_BALANCE);
        dstToken.mint(userB, INITIAL_USER_BALANCE);
    }

    // ═══════════════════════════════════════════════════ TOKEN ═══════════════════════════════════════════════════════

    function test_relay_token() public {
        bytes32 txId = getTxId(tokenTx);
        relay({caller: relayerA, msgValue: 0, bridgeTx: tokenTx});
        (uint256 blockNumber, uint256 blockTimestamp, address relayer) = fastBridge.bridgeRelayDetails(txId);
        assertEq(blockNumber, block.number);
        assertEq(blockTimestamp, block.timestamp);
        assertEq(relayer, relayerA);
        assertEq(dstToken.balanceOf(userB), INITIAL_USER_BALANCE + tokenParams.destAmount);
        assertEq(dstToken.balanceOf(relayerA), LEFTOVER_BALANCE);
    }

    function test_relay_token_withRelayerAddress() public {
        bytes32 txId = getTxId(tokenTx);
        relayWithAddress({caller: relayerB, relayer: relayerA, msgValue: 0, bridgeTx: tokenTx});
        (uint256 blockNumber, uint256 blockTimestamp, address relayer) = fastBridge.bridgeRelayDetails(txId);
        assertEq(blockNumber, block.number);
        assertEq(blockTimestamp, block.timestamp);
        assertEq(relayer, relayerA);
        assertEq(dstToken.balanceOf(userB), INITIAL_USER_BALANCE + tokenParams.destAmount);
        assertEq(dstToken.balanceOf(relayerB), LEFTOVER_BALANCE);
    }

    // ════════════════════════════════════════════════════ ETH ════════════════════════════════════════════════════════

    function test_relay_eth() public {
        bytes32 txId = getTxId(ethTx);
        relay({caller: relayerA, msgValue: ethParams.destAmount, bridgeTx: ethTx});
        (uint256 blockNumber, uint256 blockTimestamp, address relayer) = fastBridge.bridgeRelayDetails(txId);
        assertEq(blockNumber, block.number);
        assertEq(blockTimestamp, block.timestamp);
        assertEq(relayer, relayerA);
        assertEq(address(userB).balance, INITIAL_USER_BALANCE + ethParams.destAmount);
        assertEq(address(relayerA).balance, LEFTOVER_BALANCE);
    }

    function test_relay_eth_withRelayerAddress() public {
        bytes32 txId = getTxId(ethTx);
        relayWithAddress({caller: relayerB, relayer: relayerA, msgValue: ethParams.destAmount, bridgeTx: ethTx});
        (uint256 blockNumber, uint256 blockTimestamp, address relayer) = fastBridge.bridgeRelayDetails(txId);
        assertEq(blockNumber, block.number);
        assertEq(blockTimestamp, block.timestamp);
        assertEq(relayer, relayerA);
        assertEq(address(userB).balance, INITIAL_USER_BALANCE + ethParams.destAmount);
        assertEq(address(relayerB).balance, LEFTOVER_BALANCE);
    }
}
