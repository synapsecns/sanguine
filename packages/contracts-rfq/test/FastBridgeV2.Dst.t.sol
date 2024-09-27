// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {FastBridgeV2DstBaseTest, IFastBridgeV2} from "./FastBridgeV2.Dst.Base.t.sol";

// solhint-disable func-name-mixedcase, ordering
contract FastBridgeV2DstTest is FastBridgeV2DstBaseTest {
    event BridgeRelayed(
        bytes32 indexed transactionId,
        address indexed relayer,
        address indexed to,
        uint32 originChainId,
        address originToken,
        address destToken,
        uint256 originAmount,
        uint256 destAmount,
        uint256 chainGasAmount
    );

    function expectBridgeRelayed(
        IFastBridgeV2.BridgeTransactionV2 memory bridgeTx,
        bytes32 txId,
        address relayer
    )
        public
    {
        vm.expectEmit(address(fastBridge));
        emit BridgeRelayed({
            transactionId: txId,
            relayer: relayer,
            to: bridgeTx.txV1.destRecipient,
            originChainId: bridgeTx.txV1.originChainId,
            originToken: bridgeTx.txV1.originToken,
            destToken: bridgeTx.txV1.destToken,
            originAmount: bridgeTx.txV1.originAmount,
            destAmount: bridgeTx.txV1.destAmount,
            chainGasAmount: 0
        });
    }

    function checkRelayedViews(bytes32 txId, address expectedRelayer) public view {
        assertTrue(fastBridge.bridgeRelays(txId));
        (uint48 blockNumber, uint48 blockTimestamp, address relayer) = fastBridge.bridgeRelayDetails(txId);
        assertEq(blockNumber, block.number);
        assertEq(blockTimestamp, block.timestamp);
        assertEq(relayer, expectedRelayer);
    }

    /// @notice RelayerA completes the ERC20 bridge request
    function test_relay_token() public {
        bytes32 txId = getTxId(tokenTx);
        expectBridgeRelayed(tokenTx, txId, relayerA);
        relay({caller: relayerA, msgValue: 0, bridgeTx: tokenTx});
        checkRelayedViews({txId: txId, expectedRelayer: relayerA});
        assertEq(dstToken.balanceOf(userB), tokenParams.destAmount);
        assertEq(dstToken.balanceOf(relayerA), LEFTOVER_BALANCE);
        assertEq(dstToken.balanceOf(address(fastBridge)), 0);
    }

    /// @notice RelayerB completes the ERC20 bridge request, using relayerA's address
    function test_relay_token_withRelayerAddress() public {
        bytes32 txId = getTxId(tokenTx);
        expectBridgeRelayed(tokenTx, txId, relayerA);
        relayWithAddress({caller: relayerB, relayer: relayerA, msgValue: 0, bridgeTx: tokenTx});
        checkRelayedViews({txId: txId, expectedRelayer: relayerA});
        assertEq(dstToken.balanceOf(userB), tokenParams.destAmount);
        assertEq(dstToken.balanceOf(relayerB), LEFTOVER_BALANCE);
        assertEq(dstToken.balanceOf(address(fastBridge)), 0);
    }

    /// @notice RelayerB completes the ETH bridge request
    function test_relay_eth() public {
        bytes32 txId = getTxId(ethTx);
        expectBridgeRelayed(ethTx, txId, relayerB);
        relay({caller: relayerB, msgValue: ethParams.destAmount, bridgeTx: ethTx});
        checkRelayedViews({txId: txId, expectedRelayer: relayerB});
        assertEq(userB.balance, ethParams.destAmount);
        assertEq(relayerB.balance, LEFTOVER_BALANCE);
        assertEq(address(fastBridge).balance, 0);
    }

    /// @notice RelayerA completes the ETH bridge request, using relayerB's address
    function test_relay_eth_withRelayerAddress() public {
        bytes32 txId = getTxId(ethTx);
        expectBridgeRelayed(ethTx, txId, relayerB);
        relayWithAddress({caller: relayerA, relayer: relayerB, msgValue: ethParams.destAmount, bridgeTx: ethTx});
        checkRelayedViews({txId: txId, expectedRelayer: relayerB});
        assertEq(userB.balance, ethParams.destAmount);
        assertEq(relayerA.balance, LEFTOVER_BALANCE);
        assertEq(address(fastBridge).balance, 0);
    }

    /// @notice RelayerA completes the ETH bridge request, using relayerB's address
    function test_relay_eth_withRelayerAddress_checkBlockData() public {
        vm.roll(987_654_321);
        vm.warp(123_456_789);
        bytes32 txId = getTxId(ethTx);
        expectBridgeRelayed(ethTx, txId, relayerB);
        relayWithAddress({caller: relayerA, relayer: relayerB, msgValue: ethParams.destAmount, bridgeTx: ethTx});
        assertTrue(fastBridge.bridgeRelays(txId));
        (uint48 recordedBlockNumber, uint48 recordedBlockTimestamp, address recordedRelayer) =
            fastBridge.bridgeRelayDetails(txId);
        assertEq(recordedBlockNumber, 987_654_321);
        assertEq(recordedBlockTimestamp, 123_456_789);
        assertEq(recordedRelayer, relayerB);
        assertEq(userB.balance, ethParams.destAmount);
        assertEq(relayerA.balance, LEFTOVER_BALANCE);
        assertEq(address(fastBridge).balance, 0);
    }
    // ══════════════════════════════════════════════════ REVERTS ══════════════════════════════════════════════════════

    function test_relay_revert_chainIncorrect() public {
        vm.chainId(SRC_CHAIN_ID);
        vm.expectRevert(ChainIncorrect.selector);
        relay({caller: relayerA, msgValue: 0, bridgeTx: tokenTx});
    }

    function test_relay_revert_transactionRelayed() public {
        relay({caller: relayerA, msgValue: 0, bridgeTx: tokenTx});
        vm.expectRevert(TransactionRelayed.selector);
        relay({caller: relayerA, msgValue: 0, bridgeTx: tokenTx});
    }

    function test_relay_revert_deadlineExceeded() public {
        skip(DEADLINE + 1);
        vm.expectRevert(DeadlineExceeded.selector);
        relay({caller: relayerA, msgValue: 0, bridgeTx: tokenTx});
    }

    function test_relay_withRelayerAddress_revert_chainIncorrect() public {
        vm.chainId(SRC_CHAIN_ID);
        vm.expectRevert(ChainIncorrect.selector);
        relayWithAddress({caller: relayerA, relayer: relayerB, msgValue: 0, bridgeTx: tokenTx});
    }

    function test_relay_withRelayerAddress_revert_transactionRelayed() public {
        relay({caller: relayerA, msgValue: 0, bridgeTx: tokenTx});
        vm.expectRevert(TransactionRelayed.selector);
        relayWithAddress({caller: relayerA, relayer: relayerB, msgValue: 0, bridgeTx: tokenTx});
    }

    function test_relay_withRelayerAddress_revert_deadlineExceeded() public {
        skip(DEADLINE + 1);
        vm.expectRevert(DeadlineExceeded.selector);
        relayWithAddress({caller: relayerA, relayer: relayerB, msgValue: 0, bridgeTx: tokenTx});
    }

    function test_relay_withRelayerAddress_revert_zeroAddr() public {
        vm.expectRevert(ZeroAddress.selector);
        relayWithAddress({caller: relayerA, relayer: address(0), msgValue: 0, bridgeTx: tokenTx});
    }
}
