// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {ChainIncorrect, DeadlineExceeded, TransactionRelayed, ZeroAddress} from "../contracts/libs/Errors.sol";

import {FastBridgeV2, FastBridgeV2Test, IFastBridge} from "./FastBridgeV2.t.sol";

// solhint-disable func-name-mixedcase, ordering
contract FastBridgeV2DstTest is FastBridgeV2Test {
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

    uint256 public constant LEFTOVER_BALANCE = 1 ether;

    function setUp() public override {
        vm.chainId(DST_CHAIN_ID);
        super.setUp();
    }

    function deployFastBridge() public override returns (FastBridgeV2) {
        return new FastBridgeV2(address(this));
    }

    function mintTokens() public override {
        dstToken.mint(address(relayerA), LEFTOVER_BALANCE + tokenParams.destAmount);
        deal(relayerB, LEFTOVER_BALANCE + ethParams.destAmount);
        vm.prank(relayerA);
        dstToken.approve(address(fastBridge), type(uint256).max);
    }

    function expectBridgeRelayed(IFastBridge.BridgeTransaction memory bridgeTx, bytes32 txId, address relayer) public {
        vm.expectEmit(address(fastBridge));
        emit BridgeRelayed({
            transactionId: txId,
            relayer: relayer,
            to: bridgeTx.destRecipient,
            originChainId: bridgeTx.originChainId,
            originToken: bridgeTx.originToken,
            destToken: bridgeTx.destToken,
            originAmount: bridgeTx.originAmount,
            destAmount: bridgeTx.destAmount,
            chainGasAmount: 0
        });
    }

    function relay(address caller, uint256 msgValue, IFastBridge.BridgeTransaction memory bridgeTx) public {
        bytes memory request = abi.encode(bridgeTx);
        vm.prank(caller);
        fastBridge.relay{value: msgValue}(request);
    }

    function relayWithAddress(
        address caller,
        address relayer,
        uint256 msgValue,
        IFastBridge.BridgeTransaction memory bridgeTx
    )
        public
    {
        bytes memory request = abi.encode(bridgeTx);
        vm.prank(caller);
        fastBridge.relay{value: msgValue}(request, relayer);
    }

    /// @notice RelayerA completes the ERC20 bridge request
    function test_relay_token() public {
        bytes32 txId = getTxId(tokenTx);
        expectBridgeRelayed(tokenTx, txId, address(relayerA));
        relay({caller: relayerA, msgValue: 0, bridgeTx: tokenTx});
        assertTrue(fastBridge.bridgeRelays(txId));
        assertEq(dstToken.balanceOf(address(userB)), tokenParams.destAmount);
        assertEq(dstToken.balanceOf(address(relayerA)), LEFTOVER_BALANCE);
        assertEq(dstToken.balanceOf(address(fastBridge)), 0);
    }

    /// @notice RelayerA completes the ERC20 bridge request, using relayerB's address
    function test_relay_token_withRelayerAddress() public {
        bytes32 txId = getTxId(tokenTx);
        expectBridgeRelayed(tokenTx, txId, address(relayerB));
        relayWithAddress({caller: relayerA, relayer: relayerB, msgValue: 0, bridgeTx: tokenTx});
        assertTrue(fastBridge.bridgeRelays(txId));
        assertEq(dstToken.balanceOf(address(userB)), tokenParams.destAmount);
        assertEq(dstToken.balanceOf(address(relayerA)), LEFTOVER_BALANCE);
        assertEq(dstToken.balanceOf(address(fastBridge)), 0);
    }

    /// @notice RelayerB completes the ETH bridge request
    function test_relay_eth() public {
        bytes32 txId = getTxId(ethTx);
        expectBridgeRelayed(ethTx, txId, address(relayerB));
        relay({caller: relayerB, msgValue: ethParams.destAmount, bridgeTx: ethTx});
        assertTrue(fastBridge.bridgeRelays(txId));
        assertEq(address(userB).balance, ethParams.destAmount);
        assertEq(address(relayerB).balance, LEFTOVER_BALANCE);
        assertEq(address(fastBridge).balance, 0);
    }

    /// @notice RelayerB completes the ETH bridge request, using relayerA's address
    function test_relay_eth_withRelayerAddress() public {
        bytes32 txId = getTxId(ethTx);
        expectBridgeRelayed(ethTx, txId, address(relayerA));
        relayWithAddress({caller: relayerB, relayer: relayerA, msgValue: ethParams.destAmount, bridgeTx: ethTx});
        assertTrue(fastBridge.bridgeRelays(txId));
        assertEq(address(userB).balance, ethParams.destAmount);
        assertEq(address(relayerB).balance, LEFTOVER_BALANCE);
        assertEq(address(fastBridge).balance, 0);
    }

    /// @notice RelayerB completes the ETH bridge request, using relayerA's address
    function test_relay_eth_withRelayerAddress_checkBlockData() public {
        vm.roll(987_654_321);
        vm.warp(123_456_789);
        bytes32 txId = getTxId(ethTx);
        expectBridgeRelayed(ethTx, txId, address(relayerA));
        relayWithAddress({caller: relayerB, relayer: relayerA, msgValue: ethParams.destAmount, bridgeTx: ethTx});
        assertTrue(fastBridge.bridgeRelays(txId));
        (uint48 recordedBlockNumber, uint48 recordedblockTimestamp,) = fastBridge.bridgeRelayDetails(txId);
        assertEq(recordedBlockNumber, 987_654_321);
        assertEq(recordedblockTimestamp, 123_456_789);
        assertEq(address(userB).balance, ethParams.destAmount);
        assertEq(address(relayerB).balance, LEFTOVER_BALANCE);
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
