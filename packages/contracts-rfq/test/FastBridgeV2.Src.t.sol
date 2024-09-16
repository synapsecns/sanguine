// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {MsgValueIncorrect} from "../contracts/libs/Errors.sol";

import {FastBridgeV2, FastBridgeV2Test, IFastBridge} from "./FastBridgeV2.t.sol";

// solhint-disable func-name-mixedcase, ordering
contract FastBridgeV2SrcTest is FastBridgeV2Test {
    event BridgeRequested(
        bytes32 indexed transactionId,
        address indexed sender,
        bytes request,
        uint32 destChainId,
        address originToken,
        address destToken,
        uint256 originAmount,
        uint256 destAmount,
        bool sendChainGas
    );

    uint256 public constant LEFTOVER_BALANCE = 1 ether;
    uint256 public constant INITIAL_PROTOCOL_FEES_TOKEN = 456_789;
    uint256 public constant INITIAL_PROTOCOL_FEES_ETH = 0.123 ether;

    function setUp() public override {
        vm.chainId(SRC_CHAIN_ID);
        super.setUp();
    }

    function deployFastBridge() public override returns (FastBridgeV2) {
        return new FastBridgeV2(address(this));
    }

    function mintTokens() public override {
        // Prior Protocol fees
        srcToken.mint(address(fastBridge), INITIAL_PROTOCOL_FEES_TOKEN);
        deal(address(fastBridge), INITIAL_PROTOCOL_FEES_ETH);
        cheatCollectedProtocolFees(address(srcToken), INITIAL_PROTOCOL_FEES_TOKEN);
        cheatCollectedProtocolFees(ETH_ADDRESS, INITIAL_PROTOCOL_FEES_ETH);
        // Users
        srcToken.mint(userA, LEFTOVER_BALANCE + tokenParams.originAmount);
        srcToken.mint(userB, LEFTOVER_BALANCE + tokenParams.originAmount);
        deal(userA, LEFTOVER_BALANCE + ethParams.originAmount);
        deal(userB, LEFTOVER_BALANCE + ethParams.originAmount);
        vm.prank(userA);
        srcToken.approve(address(fastBridge), type(uint256).max);
        vm.prank(userB);
        srcToken.approve(address(fastBridge), type(uint256).max);
    }

    function bridge(address caller, uint256 msgValue, IFastBridge.BridgeParams memory params) public {
        vm.prank(caller);
        fastBridge.bridge{value: msgValue}(params);
    }

    function expectBridgeRequested(IFastBridge.BridgeTransaction memory bridgeTx, bytes32 txId) public {
        vm.expectEmit(address(fastBridge));
        emit BridgeRequested({
            transactionId: txId,
            sender: bridgeTx.originSender,
            request: abi.encode(bridgeTx),
            destChainId: bridgeTx.destChainId,
            originToken: bridgeTx.originToken,
            destToken: bridgeTx.destToken,
            originAmount: bridgeTx.originAmount,
            destAmount: bridgeTx.destAmount,
            sendChainGas: bridgeTx.sendChainGas
        });
    }

    function assertEq(FastBridgeV2.BridgeStatus a, FastBridgeV2.BridgeStatus b) public pure {
        assertEq(uint8(a), uint8(b));
    }

    // ══════════════════════════════════════════════════ BRIDGE ═══════════════════════════════════════════════════════

    function checkTokenBalancesAfterBridge(address caller) public {
        assertEq(fastBridge.protocolFees(address(srcToken)), INITIAL_PROTOCOL_FEES_TOKEN);
        assertEq(srcToken.balanceOf(caller), LEFTOVER_BALANCE);
        assertEq(srcToken.balanceOf(address(fastBridge)), INITIAL_PROTOCOL_FEES_TOKEN + tokenParams.originAmount);
    }

    function test_bridge_token() public {
        bytes32 txId = getTxId(tokenTx);
        expectBridgeRequested(tokenTx, txId);
        bridge({caller: userA, msgValue: 0, params: tokenParams});
        assertEq(fastBridge.bridgeStatuses(txId), FastBridgeV2.BridgeStatus.REQUESTED);
        checkTokenBalancesAfterBridge(userA);
    }

    function test_bridge_token_diffSender() public {
        bytes32 txId = getTxId(tokenTx);
        expectBridgeRequested(tokenTx, txId);
        bridge({caller: userB, msgValue: 0, params: tokenParams});
        assertEq(fastBridge.bridgeStatuses(txId), FastBridgeV2.BridgeStatus.REQUESTED);
        assertEq(srcToken.balanceOf(userA), LEFTOVER_BALANCE + tokenParams.originAmount);
        checkTokenBalancesAfterBridge(userB);
    }

    function checkEthBalancesAfterBridge(address caller) public {
        assertEq(fastBridge.protocolFees(ETH_ADDRESS), INITIAL_PROTOCOL_FEES_ETH);
        assertEq(address(caller).balance, LEFTOVER_BALANCE);
        assertEq(address(fastBridge).balance, INITIAL_PROTOCOL_FEES_ETH + ethParams.originAmount);
    }

    function test_bridge_eth() public {
        // bridge token first to match the nonce
        bridge({caller: userA, msgValue: 0, params: tokenParams});
        bytes32 txId = getTxId(ethTx);
        expectBridgeRequested(ethTx, txId);
        bridge({caller: userA, msgValue: ethParams.originAmount, params: ethParams});
        assertEq(fastBridge.bridgeStatuses(txId), FastBridgeV2.BridgeStatus.REQUESTED);
        checkEthBalancesAfterBridge(userA);
    }

    function test_bridge_eth_diffSender() public {
        // bridge token first to match the nonce
        bridge({caller: userA, msgValue: 0, params: tokenParams});
        bytes32 txId = getTxId(ethTx);
        expectBridgeRequested(ethTx, txId);
        bridge({caller: userB, msgValue: ethParams.originAmount, params: ethParams});
        assertEq(fastBridge.bridgeStatuses(txId), FastBridgeV2.BridgeStatus.REQUESTED);
        assertEq(userA.balance, LEFTOVER_BALANCE + ethParams.originAmount);
        checkEthBalancesAfterBridge(userB);
    }

    function test_bridge_userSpecificNonce() public {
        vm.skip(true); // TODO: unskip when implemented
        bridge({caller: userA, msgValue: 0, params: tokenParams});
        // UserB nonce is 0
        ethTx.nonce = 0;
        ethParams.sender = userB;
        ethTx.originSender = userB;
        bytes32 txId = getTxId(ethTx);
        expectBridgeRequested(ethTx, txId);
        bridge({caller: userB, msgValue: ethParams.originAmount, params: ethParams});
        assertEq(fastBridge.bridgeStatuses(txId), FastBridgeV2.BridgeStatus.REQUESTED);
        checkEthBalancesAfterBridge(userB);
    }

    function test_bridge_eth_revert_lowerMsgValue() public {
        vm.expectRevert(MsgValueIncorrect.selector);
        bridge({caller: userA, msgValue: ethParams.originAmount - 1, params: ethParams});
    }

    function test_bridge_eth_revert_higherMsgValue() public {
        vm.expectRevert(MsgValueIncorrect.selector);
        bridge({caller: userA, msgValue: ethParams.originAmount + 1, params: ethParams});
    }

    function test_bridge_eth_revert_zeroMsgValue() public {
        vm.expectRevert(MsgValueIncorrect.selector);
        bridge({caller: userA, msgValue: 0, params: ethParams});
    }
}
