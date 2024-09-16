// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {MsgValueIncorrect, StatusIncorrect} from "../contracts/libs/Errors.sol";

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

    event BridgeProofProvided(bytes32 indexed transactionId, address indexed relayer, bytes32 transactionHash);

    uint256 public constant CLAIM_DELAY = 30 minutes;

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

    function configureFastBridge() public virtual override {
        fastBridge.grantRole(fastBridge.RELAYER_ROLE(), relayerA);
        fastBridge.grantRole(fastBridge.RELAYER_ROLE(), relayerB);
        fastBridge.grantRole(fastBridge.GUARD_ROLE(), guard);
        fastBridge.grantRole(fastBridge.REFUNDER_ROLE(), refunder);
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

    function prove(address caller, IFastBridge.BridgeTransaction memory bridgeTx, bytes32 destTxHash) public {
        vm.prank(caller);
        fastBridge.prove(abi.encode(bridgeTx), destTxHash);
    }

    function claim(address caller, IFastBridge.BridgeTransaction memory bridgeTx, address to) public {
        vm.prank(caller);
        fastBridge.claim(abi.encode(bridgeTx), to);
    }

    function dispute(address caller, bytes32 txId) public {
        vm.prank(caller);
        fastBridge.dispute(txId);
    }

    function refund(address caller, IFastBridge.BridgeTransaction memory bridgeTx) public {
        vm.prank(caller);
        fastBridge.refund(abi.encode(bridgeTx));
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

    function expectBridgeProofProvided(bytes32 txId, address relayer, bytes32 destTxHash) public {
        vm.expectEmit(address(fastBridge));
        emit BridgeProofProvided({transactionId: txId, relayer: relayer, transactionHash: destTxHash});
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

    // ═══════════════════════════════════════════════════ PROVE ═══════════════════════════════════════════════════════

    function test_prove_token() public {
        bytes32 txId = getTxId(tokenTx);
        bridge({caller: userA, msgValue: 0, params: tokenParams});
        expectBridgeProofProvided({txId: txId, relayer: relayerA, destTxHash: hex"01"});
        prove({caller: relayerA, bridgeTx: tokenTx, destTxHash: hex"01"});
        (uint96 timestamp, address relayer) = fastBridge.bridgeProofs(txId);
        assertEq(timestamp, block.timestamp);
        assertEq(relayer, relayerA);
        assertEq(srcToken.balanceOf(address(fastBridge)), INITIAL_PROTOCOL_FEES_TOKEN + tokenParams.originAmount);
        assertEq(fastBridge.protocolFees(address(srcToken)), INITIAL_PROTOCOL_FEES_TOKEN);
    }

    function test_prove_eth() public {
        // bridge token first to match the nonce
        bridge({caller: userA, msgValue: 0, params: tokenParams});
        bytes32 txId = getTxId(ethTx);
        bridge({caller: userA, msgValue: ethParams.originAmount, params: ethParams});
        expectBridgeProofProvided({txId: txId, relayer: relayerA, destTxHash: hex"01"});
        prove({caller: relayerA, bridgeTx: ethTx, destTxHash: hex"01"});
        (uint96 timestamp, address relayer) = fastBridge.bridgeProofs(txId);
        assertEq(timestamp, block.timestamp);
        assertEq(relayer, relayerA);
        assertEq(address(fastBridge).balance, INITIAL_PROTOCOL_FEES_ETH + ethParams.originAmount);
        assertEq(fastBridge.protocolFees(ETH_ADDRESS), INITIAL_PROTOCOL_FEES_ETH);
    }

    function test_prove_revert_statusNull() public {
        vm.expectRevert(StatusIncorrect.selector);
        prove({caller: relayerA, bridgeTx: tokenTx, destTxHash: hex"01"});
    }

    function test_prove_revert_statusProved() public {
        bridge({caller: userA, msgValue: 0, params: tokenParams});
        prove({caller: relayerA, bridgeTx: tokenTx, destTxHash: hex"01"});
        vm.expectRevert(StatusIncorrect.selector);
        prove({caller: relayerB, bridgeTx: tokenTx, destTxHash: hex"02"});
    }

    function test_prove_revert_statusClaimed() public {
        bridge({caller: userA, msgValue: 0, params: tokenParams});
        prove({caller: relayerA, bridgeTx: tokenTx, destTxHash: hex"01"});
        skip(CLAIM_DELAY + 1);
        claim({caller: relayerA, bridgeTx: tokenTx, to: relayerA});
        vm.expectRevert(StatusIncorrect.selector);
        prove({caller: relayerB, bridgeTx: tokenTx, destTxHash: hex"02"});
    }

    function test_prove_revert_statusRefunded() public {
        bridge({caller: userA, msgValue: 0, params: tokenParams});
        skip(DEADLINE + 1);
        refund({caller: refunder, bridgeTx: tokenTx});
        vm.expectRevert(StatusIncorrect.selector);
        prove({caller: relayerA, bridgeTx: tokenTx, destTxHash: hex"01"});
    }

    function test_prove_revert_callerNotRelayer(address caller) public {
        vm.assume(caller != relayerA && caller != relayerB);
        bridge({caller: userA, msgValue: 0, params: tokenParams});
        expectUnauthorized(caller, fastBridge.RELAYER_ROLE());
        prove({caller: caller, bridgeTx: tokenTx, destTxHash: hex"01"});
    }
}
