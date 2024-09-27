// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {
    AmountIncorrect,
    ChainIncorrect,
    DisputePeriodNotPassed,
    DisputePeriodPassed,
    DeadlineNotExceeded,
    DeadlineTooShort,
    MsgValueIncorrect,
    SenderIncorrect,
    StatusIncorrect,
    ZeroAddress
} from "../contracts/libs/Errors.sol";

import {FastBridgeV2SrcBaseTest} from "./FastBridgeV2.Src.Base.t.sol";
import {IFastBridge, IFastBridgeV2} from "./FastBridgeV2.t.sol";

// solhint-disable func-name-mixedcase, ordering
contract FastBridgeV2SrcTest is FastBridgeV2SrcBaseTest {
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

    event BridgeDepositClaimed(
        bytes32 indexed transactionId, address indexed relayer, address indexed to, address token, uint256 amount
    );

    event BridgeProofDisputed(bytes32 indexed transactionId, address indexed relayer);

    event BridgeDepositRefunded(bytes32 indexed transactionId, address indexed to, address token, uint256 amount);

    address public claimTo = makeAddr("Claim To");

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

    function expectBridgeDepositClaimed(
        IFastBridge.BridgeTransaction memory bridgeTx,
        bytes32 txId,
        address relayer,
        address to
    )
        public
    {
        vm.expectEmit(address(fastBridge));
        emit BridgeDepositClaimed({
            transactionId: txId,
            relayer: relayer,
            to: to,
            token: bridgeTx.originToken,
            amount: bridgeTx.originAmount
        });
    }

    function expectBridgeProofDisputed(bytes32 txId, address guard) public {
        vm.expectEmit(address(fastBridge));
        // Note: BridgeProofDisputed event has a mislabeled address parameter, this is actually the guard
        emit BridgeProofDisputed({transactionId: txId, relayer: guard});
    }

    function expectBridgeDepositRefunded(IFastBridge.BridgeParams memory bridgeParams, bytes32 txId) public {
        vm.expectEmit(address(fastBridge));
        emit BridgeDepositRefunded({
            transactionId: txId,
            to: bridgeParams.sender,
            token: bridgeParams.originToken,
            amount: bridgeParams.originAmount
        });
    }

    // ══════════════════════════════════════════════════ BRIDGE ═══════════════════════════════════════════════════════

    function checkTokenBalancesAfterBridge(address caller) public view {
        assertEq(fastBridge.protocolFees(address(srcToken)), INITIAL_PROTOCOL_FEES_TOKEN);
        assertEq(srcToken.balanceOf(caller), LEFTOVER_BALANCE);
        assertEq(srcToken.balanceOf(address(fastBridge)), INITIAL_PROTOCOL_FEES_TOKEN + tokenParams.originAmount);
    }

    function test_bridge_token() public {
        bytes32 txId = getTxId(tokenTx);
        expectBridgeRequested(tokenTx, txId);
        bridge({caller: userA, msgValue: 0, params: tokenParams});
        assertEq(fastBridge.bridgeStatuses(txId), IFastBridgeV2.BridgeStatus.REQUESTED);
        checkTokenBalancesAfterBridge(userA);
    }

    function test_bridge_token_diffSender() public {
        bytes32 txId = getTxId(tokenTx);
        expectBridgeRequested(tokenTx, txId);
        bridge({caller: userB, msgValue: 0, params: tokenParams});
        assertEq(fastBridge.bridgeStatuses(txId), IFastBridgeV2.BridgeStatus.REQUESTED);
        assertEq(srcToken.balanceOf(userA), LEFTOVER_BALANCE + tokenParams.originAmount);
        checkTokenBalancesAfterBridge(userB);
    }

    function checkEthBalancesAfterBridge(address caller) public view {
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
        assertEq(fastBridge.bridgeStatuses(txId), IFastBridgeV2.BridgeStatus.REQUESTED);
        checkEthBalancesAfterBridge(userA);
    }

    function test_bridge_eth_diffSender() public {
        // bridge token first to match the nonce
        bridge({caller: userA, msgValue: 0, params: tokenParams});
        bytes32 txId = getTxId(ethTx);
        expectBridgeRequested(ethTx, txId);
        bridge({caller: userB, msgValue: ethParams.originAmount, params: ethParams});
        assertEq(fastBridge.bridgeStatuses(txId), IFastBridgeV2.BridgeStatus.REQUESTED);
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
        assertEq(fastBridge.bridgeStatuses(txId), IFastBridgeV2.BridgeStatus.REQUESTED);
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

    function test_bridge_revert_sameDestinationChain() public {
        tokenParams.dstChainId = SRC_CHAIN_ID;
        vm.expectRevert(ChainIncorrect.selector);
        bridge({caller: userA, msgValue: 0, params: tokenParams});
    }

    function test_bridge_revert_zeroOriginAmount() public {
        tokenParams.originAmount = 0;
        vm.expectRevert(AmountIncorrect.selector);
        bridge({caller: userA, msgValue: 0, params: tokenParams});
    }

    function test_bridge_revert_zeroDestAmount() public {
        tokenParams.destAmount = 0;
        vm.expectRevert(AmountIncorrect.selector);
        bridge({caller: userA, msgValue: 0, params: tokenParams});
    }

    function test_bridge_revert_zeroOriginToken() public {
        tokenParams.originToken = address(0);
        vm.expectRevert(ZeroAddress.selector);
        bridge({caller: userA, msgValue: 0, params: tokenParams});
    }

    function test_bridge_revert_zeroDestToken() public {
        tokenParams.destToken = address(0);
        vm.expectRevert(ZeroAddress.selector);
        bridge({caller: userA, msgValue: 0, params: tokenParams});
    }

    function test_bridge_revert_zeroSender() public {
        tokenParams.sender = address(0);
        vm.expectRevert(ZeroAddress.selector);
        bridge({caller: userA, msgValue: 0, params: tokenParams});
    }

    function test_bridge_revert_zeroRecipient() public {
        tokenParams.to = address(0);
        vm.expectRevert(ZeroAddress.selector);
        bridge({caller: userA, msgValue: 0, params: tokenParams});
    }

    function test_bridge_revert_deadlineTooClose() public {
        tokenParams.deadline = block.timestamp + MIN_DEADLINE - 1;
        vm.expectRevert(DeadlineTooShort.selector);
        bridge({caller: userA, msgValue: 0, params: tokenParams});
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

    // ════════════════════════════════════════ PROVE OTHER RELAYER ════════════════════════════════════════════

    function test_proveOther_token() public {
        bytes32 txId = getTxId(tokenTx);
        bridge({caller: userA, msgValue: 0, params: tokenParams});
        expectBridgeProofProvided({txId: txId, relayer: relayerA, destTxHash: hex"01"});
        prove({caller: relayerB, transactionId: txId, destTxHash: hex"01", relayer: relayerA});
        (uint96 timestamp, address relayer) = fastBridge.bridgeProofs(txId);
        assertEq(timestamp, block.timestamp);
        assertEq(relayer, relayerA);
        assertEq(srcToken.balanceOf(address(fastBridge)), INITIAL_PROTOCOL_FEES_TOKEN + tokenParams.originAmount);
        assertEq(fastBridge.protocolFees(address(srcToken)), INITIAL_PROTOCOL_FEES_TOKEN);
    }

    function test_proveOther_eth() public {
        // bridge token first to match the nonce
        bridge({caller: userA, msgValue: 0, params: tokenParams});
        bytes32 txId = getTxId(ethTx);
        bridge({caller: userA, msgValue: ethParams.originAmount, params: ethParams});
        expectBridgeProofProvided({txId: txId, relayer: relayerA, destTxHash: hex"01"});
        prove({caller: relayerB, transactionId: txId, destTxHash: hex"01", relayer: relayerA});
        (uint96 timestamp, address relayer) = fastBridge.bridgeProofs(txId);
        assertEq(timestamp, block.timestamp);
        assertEq(relayer, relayerA);
        assertEq(address(fastBridge).balance, INITIAL_PROTOCOL_FEES_ETH + ethParams.originAmount);
        assertEq(fastBridge.protocolFees(ETH_ADDRESS), INITIAL_PROTOCOL_FEES_ETH);
    }

    // relayer self-proving using tx id, which is capable of proving for another & most tests focus on that angle.
    function test_proveOther_self() public {
        bytes32 txId = getTxId(tokenTx);
        bridge({caller: userA, msgValue: 0, params: tokenParams});
        expectBridgeProofProvided({txId: txId, relayer: relayerA, destTxHash: hex"01"});
        prove({caller: relayerA, transactionId: txId, destTxHash: hex"01", relayer: relayerA});
        (uint96 timestamp, address relayer) = fastBridge.bridgeProofs(txId);
        assertEq(timestamp, block.timestamp);
        assertEq(relayer, relayerA);
        assertEq(srcToken.balanceOf(address(fastBridge)), INITIAL_PROTOCOL_FEES_TOKEN + tokenParams.originAmount);
        assertEq(fastBridge.protocolFees(address(srcToken)), INITIAL_PROTOCOL_FEES_TOKEN);
    }

    // arbitrary non-privileged address can be asserted as the relayer
    function test_proveOther_permless() public {
        bytes32 txId = getTxId(tokenTx);

        bridge({caller: userA, msgValue: 0, params: tokenParams});
        expectBridgeProofProvided({txId: txId, relayer: address(0x1234), destTxHash: hex"01"});
        prove({caller: relayerA, transactionId: txId, destTxHash: hex"01", relayer: address(0x1234)});
        (uint96 timestamp, address relayer) = fastBridge.bridgeProofs(txId);
        assertEq(timestamp, block.timestamp);
        assertEq(relayer, address(0x1234));
        assertEq(srcToken.balanceOf(address(fastBridge)), INITIAL_PROTOCOL_FEES_TOKEN + tokenParams.originAmount);
        assertEq(fastBridge.protocolFees(address(srcToken)), INITIAL_PROTOCOL_FEES_TOKEN);
    }

    function test_proveOther_reProveAfterDispute() public {
        bridge({caller: userA, msgValue: 0, params: tokenParams});
        bytes32 txId = getTxId(ethTx);
        bridge({caller: userA, msgValue: ethParams.originAmount, params: ethParams});
        expectBridgeProofProvided({txId: txId, relayer: relayerA, destTxHash: hex"01"});
        prove({caller: relayerB, transactionId: txId, destTxHash: hex"01", relayer: relayerA});
        expectBridgeProofDisputed(txId, guard);
        dispute(guard, txId);
        expectBridgeProofProvided({txId: txId, relayer: relayerA, destTxHash: hex"02"});
        prove({caller: relayerB, transactionId: txId, destTxHash: hex"02", relayer: relayerA});
        expectBridgeProofDisputed(txId, guard);
        dispute(guard, txId);
        expectBridgeProofProvided({txId: txId, relayer: relayerA, destTxHash: hex"03"});
        prove({caller: relayerB, transactionId: txId, destTxHash: hex"03", relayer: relayerA});
        (uint96 timestamp, address relayer) = fastBridge.bridgeProofs(txId);
        assertEq(timestamp, block.timestamp);
        assertEq(relayer, relayerA);
        assertEq(srcToken.balanceOf(address(fastBridge)), INITIAL_PROTOCOL_FEES_TOKEN + tokenParams.originAmount);
        assertEq(fastBridge.protocolFees(address(srcToken)), INITIAL_PROTOCOL_FEES_TOKEN);
    }

    // can prove long after relaying as long as status is still good
    function test_proveOther_longDelay() public {
        bytes32 txId = getTxId(tokenTx);
        bridge({caller: userA, msgValue: 0, params: tokenParams});
        skip(10 days);
        expectBridgeProofProvided({txId: txId, relayer: relayerA, destTxHash: hex"01"});
        prove({caller: relayerA, transactionId: txId, destTxHash: hex"01", relayer: relayerA});
        (uint96 timestamp, address relayer) = fastBridge.bridgeProofs(txId);
        assertEq(timestamp, block.timestamp);
        assertEq(relayer, relayerA);
        assertEq(srcToken.balanceOf(address(fastBridge)), INITIAL_PROTOCOL_FEES_TOKEN + tokenParams.originAmount);
        assertEq(fastBridge.protocolFees(address(srcToken)), INITIAL_PROTOCOL_FEES_TOKEN);
    }

    function test_proveOther_revert_statusProved() public {
        bytes32 txId = getTxId(tokenTx);
        bridge({caller: userA, msgValue: 0, params: tokenParams});
        prove({caller: relayerB, transactionId: txId, destTxHash: hex"01", relayer: relayerA});
        vm.expectRevert(StatusIncorrect.selector);
        prove({caller: relayerB, transactionId: txId, destTxHash: hex"02", relayer: relayerA});
    }

    function test_proveOther_revert_statusClaimed() public {
        bytes32 txId = getTxId(tokenTx);
        bridge({caller: userA, msgValue: 0, params: tokenParams});
        prove({caller: relayerB, transactionId: txId, destTxHash: hex"01", relayer: relayerA});
        skip(CLAIM_DELAY + 1);
        claim({caller: relayerA, bridgeTx: tokenTx, to: relayerA});
        vm.expectRevert(StatusIncorrect.selector);
        prove({caller: relayerB, transactionId: txId, destTxHash: hex"02", relayer: relayerA});
    }

    function test_proveOther_revert_statusRefunded() public {
        bytes32 txId = getTxId(tokenTx);
        bridge({caller: userA, msgValue: 0, params: tokenParams});
        skip(DEADLINE + 1);
        refund({caller: refunder, bridgeTx: tokenTx});
        vm.expectRevert(StatusIncorrect.selector);
        prove({caller: relayerB, transactionId: txId, destTxHash: hex"01", relayer: relayerA});
    }

    function test_proveOther_revert_callerNotAuthed(address caller) public {
        bytes32 txId = getTxId(tokenTx);
        vm.assume(caller != relayerA && caller != relayerB);
        bridge({caller: userA, msgValue: 0, params: tokenParams});
        expectUnauthorized(caller, fastBridge.RELAYER_ROLE());
        prove({caller: caller, transactionId: txId, destTxHash: hex"01", relayer: relayerA});
    }

    // ═══════════════════════════════════════════════════ CLAIM ═══════════════════════════════════════════════════════

    function checkTokenBalancesAfterClaim(address relayer) public view {
        assertEq(fastBridge.protocolFees(address(srcToken)), INITIAL_PROTOCOL_FEES_TOKEN + tokenTx.originFeeAmount);
        assertEq(srcToken.balanceOf(relayer), tokenTx.originAmount);
        assertEq(srcToken.balanceOf(address(fastBridge)), INITIAL_PROTOCOL_FEES_TOKEN + tokenTx.originFeeAmount);
    }

    function test_claim_token() public {
        bytes32 txId = getTxId(tokenTx);
        bridge({caller: userA, msgValue: 0, params: tokenParams});
        prove({caller: relayerA, bridgeTx: tokenTx, destTxHash: hex"01"});
        skip(CLAIM_DELAY + 1);
        assertTrue(fastBridge.canClaim(txId, relayerA));
        expectBridgeDepositClaimed({bridgeTx: tokenTx, txId: txId, relayer: relayerA, to: relayerA});
        claim({caller: relayerA, bridgeTx: tokenTx, to: relayerA});
        assertEq(fastBridge.bridgeStatuses(txId), IFastBridgeV2.BridgeStatus.RELAYER_CLAIMED);
        checkTokenBalancesAfterClaim(relayerA);
    }

    function test_claim_token_permissionless(address caller) public {
        vm.assume(caller != relayerA);
        bytes32 txId = getTxId(tokenTx);
        bridge({caller: userA, msgValue: 0, params: tokenParams});
        prove({caller: relayerA, bridgeTx: tokenTx, destTxHash: hex"01"});
        skip(CLAIM_DELAY + 1);
        expectBridgeDepositClaimed({bridgeTx: tokenTx, txId: txId, relayer: relayerA, to: relayerA});
        claim({caller: caller, bridgeTx: tokenTx});
        assertEq(fastBridge.bridgeStatuses(txId), IFastBridgeV2.BridgeStatus.RELAYER_CLAIMED);
        checkTokenBalancesAfterClaim(relayerA);
    }

    function test_claim_token_permissionless_toZeroAddress(address caller) public {
        vm.assume(caller != relayerA);
        bytes32 txId = getTxId(tokenTx);
        bridge({caller: userA, msgValue: 0, params: tokenParams});
        prove({caller: relayerA, bridgeTx: tokenTx, destTxHash: hex"01"});
        skip(CLAIM_DELAY + 1);
        expectBridgeDepositClaimed({bridgeTx: tokenTx, txId: txId, relayer: relayerA, to: relayerA});
        claim({caller: caller, bridgeTx: tokenTx, to: address(0)});
        assertEq(fastBridge.bridgeStatuses(txId), IFastBridgeV2.BridgeStatus.RELAYER_CLAIMED);
        checkTokenBalancesAfterClaim(relayerA);
    }

    function test_claim_token_toDiffAddress() public {
        bytes32 txId = getTxId(tokenTx);
        bridge({caller: userA, msgValue: 0, params: tokenParams});
        prove({caller: relayerA, bridgeTx: tokenTx, destTxHash: hex"01"});
        skip(CLAIM_DELAY + 1);
        expectBridgeDepositClaimed({bridgeTx: tokenTx, txId: txId, relayer: relayerA, to: claimTo});
        claim({caller: relayerA, bridgeTx: tokenTx, to: claimTo});
        assertEq(fastBridge.bridgeStatuses(txId), IFastBridgeV2.BridgeStatus.RELAYER_CLAIMED);
        assertEq(srcToken.balanceOf(relayerA), 0);
        checkTokenBalancesAfterClaim(claimTo);
    }

    function test_claim_token_longDelay() public {
        bytes32 txId = getTxId(tokenTx);
        bridge({caller: userA, msgValue: 0, params: tokenParams});
        prove({caller: relayerA, bridgeTx: tokenTx, destTxHash: hex"01"});
        skip(CLAIM_DELAY + 30 days);
        expectBridgeDepositClaimed({bridgeTx: tokenTx, txId: txId, relayer: relayerA, to: relayerA});
        claim({caller: relayerA, bridgeTx: tokenTx, to: relayerA});
        assertEq(fastBridge.bridgeStatuses(txId), IFastBridgeV2.BridgeStatus.RELAYER_CLAIMED);
        checkTokenBalancesAfterClaim(relayerA);
    }

    function checkEthBalancesAfterClaim(address relayer) public view {
        assertEq(fastBridge.protocolFees(ETH_ADDRESS), INITIAL_PROTOCOL_FEES_ETH + ethTx.originFeeAmount);
        assertEq(address(relayer).balance, ethTx.originAmount);
        assertEq(address(fastBridge).balance, INITIAL_PROTOCOL_FEES_ETH + ethTx.originFeeAmount);
    }

    function test_claim_eth() public {
        bridge({caller: userA, msgValue: 0, params: tokenParams});
        bytes32 txId = getTxId(ethTx);
        bridge({caller: userA, msgValue: ethParams.originAmount, params: ethParams});
        prove({caller: relayerA, bridgeTx: ethTx, destTxHash: hex"01"});
        skip(CLAIM_DELAY + 1);
        assertTrue(fastBridge.canClaim(txId, relayerA));
        expectBridgeDepositClaimed({bridgeTx: ethTx, txId: txId, relayer: relayerA, to: relayerA});
        claim({caller: relayerA, bridgeTx: ethTx, to: relayerA});
        assertEq(fastBridge.bridgeStatuses(txId), IFastBridgeV2.BridgeStatus.RELAYER_CLAIMED);
        checkEthBalancesAfterClaim(relayerA);
    }

    function test_claim_eth_permissionless(address caller) public {
        vm.assume(caller != relayerA);
        bridge({caller: userA, msgValue: 0, params: tokenParams});
        bytes32 txId = getTxId(ethTx);
        bridge({caller: userA, msgValue: ethParams.originAmount, params: ethParams});
        prove({caller: relayerA, bridgeTx: ethTx, destTxHash: hex"01"});
        skip(CLAIM_DELAY + 1);
        expectBridgeDepositClaimed({bridgeTx: ethTx, txId: txId, relayer: relayerA, to: relayerA});
        claim({caller: caller, bridgeTx: ethTx});
        assertEq(fastBridge.bridgeStatuses(txId), IFastBridgeV2.BridgeStatus.RELAYER_CLAIMED);
        checkEthBalancesAfterClaim(relayerA);
    }

    function test_claim_eth_permissionless_toZeroAddress(address caller) public {
        vm.assume(caller != relayerA);
        bridge({caller: userA, msgValue: 0, params: tokenParams});
        bytes32 txId = getTxId(ethTx);
        bridge({caller: userA, msgValue: ethParams.originAmount, params: ethParams});
        prove({caller: relayerA, bridgeTx: ethTx, destTxHash: hex"01"});
        skip(CLAIM_DELAY + 1);
        expectBridgeDepositClaimed({bridgeTx: ethTx, txId: txId, relayer: relayerA, to: relayerA});
        claim({caller: caller, bridgeTx: ethTx, to: address(0)});
        assertEq(fastBridge.bridgeStatuses(txId), IFastBridgeV2.BridgeStatus.RELAYER_CLAIMED);
        checkEthBalancesAfterClaim(relayerA);
    }

    function test_claim_eth_toDiffAddress() public {
        bridge({caller: userA, msgValue: 0, params: tokenParams});
        bytes32 txId = getTxId(ethTx);
        bridge({caller: userA, msgValue: ethParams.originAmount, params: ethParams});
        prove({caller: relayerA, bridgeTx: ethTx, destTxHash: hex"01"});
        skip(CLAIM_DELAY + 1);
        expectBridgeDepositClaimed({bridgeTx: ethTx, txId: txId, relayer: relayerA, to: claimTo});
        claim({caller: relayerA, bridgeTx: ethTx, to: claimTo});
        assertEq(fastBridge.bridgeStatuses(txId), IFastBridgeV2.BridgeStatus.RELAYER_CLAIMED);
        checkEthBalancesAfterClaim(claimTo);
    }

    function test_claim_eth_longDelay() public {
        bridge({caller: userA, msgValue: 0, params: tokenParams});
        bytes32 txId = getTxId(ethTx);
        bridge({caller: userA, msgValue: ethParams.originAmount, params: ethParams});
        prove({caller: relayerA, bridgeTx: ethTx, destTxHash: hex"01"});
        skip(CLAIM_DELAY + 30 days);
        expectBridgeDepositClaimed({bridgeTx: ethTx, txId: txId, relayer: relayerA, to: relayerA});
        claim({caller: relayerA, bridgeTx: ethTx, to: relayerA});
        assertEq(fastBridge.bridgeStatuses(txId), IFastBridgeV2.BridgeStatus.RELAYER_CLAIMED);
        checkEthBalancesAfterClaim(relayerA);
    }

    function test_claim_revert_zeroDelay() public {
        bridge({caller: userA, msgValue: 0, params: tokenParams});
        prove({caller: relayerA, bridgeTx: tokenTx, destTxHash: hex"01"});
        assertFalse(fastBridge.canClaim(getTxId(tokenTx), relayerA));
        vm.expectRevert(DisputePeriodNotPassed.selector);
        claim({caller: relayerA, bridgeTx: tokenTx, to: relayerA});
    }

    function test_claim_revert_smallerDelay() public {
        bridge({caller: userA, msgValue: 0, params: tokenParams});
        prove({caller: relayerA, bridgeTx: tokenTx, destTxHash: hex"01"});
        skip(CLAIM_DELAY);
        assertFalse(fastBridge.canClaim(getTxId(tokenTx), relayerA));
        vm.expectRevert(DisputePeriodNotPassed.selector);
        claim({caller: relayerA, bridgeTx: tokenTx, to: relayerA});
    }

    function test_claim_revert_callerNotProven(address caller, address to) public {
        vm.assume(caller != relayerA && to != address(0));
        bridge({caller: userA, msgValue: 0, params: tokenParams});
        prove({caller: relayerA, bridgeTx: tokenTx, destTxHash: hex"01"});
        skip(CLAIM_DELAY + 1);
        vm.expectRevert(SenderIncorrect.selector);
        fastBridge.canClaim(getTxId(tokenTx), caller);
        vm.expectRevert(SenderIncorrect.selector);
        claim({caller: caller, bridgeTx: tokenTx, to: to});
    }

    function test_claim_revert_statusNull() public {
        bytes32 txId = getTxId(tokenTx);
        vm.expectRevert(StatusIncorrect.selector);
        fastBridge.canClaim(txId, relayerA);
        vm.expectRevert(StatusIncorrect.selector);
        claim({caller: relayerA, bridgeTx: tokenTx, to: relayerA});
    }

    function test_claim_revert_statusRequested() public {
        bytes32 txId = getTxId(tokenTx);
        bridge({caller: userA, msgValue: 0, params: tokenParams});
        vm.expectRevert(StatusIncorrect.selector);
        fastBridge.canClaim(txId, relayerA);
        vm.expectRevert(StatusIncorrect.selector);
        claim({caller: relayerA, bridgeTx: tokenTx, to: relayerA});
    }

    function test_claim_revert_statusClaimed() public {
        bytes32 txId = getTxId(tokenTx);
        bridge({caller: userA, msgValue: 0, params: tokenParams});
        prove({caller: relayerA, bridgeTx: tokenTx, destTxHash: hex"01"});
        skip(CLAIM_DELAY + 1);
        claim({caller: relayerA, bridgeTx: tokenTx, to: relayerA});
        vm.expectRevert(StatusIncorrect.selector);
        fastBridge.canClaim(txId, relayerA);
        vm.expectRevert(StatusIncorrect.selector);
        claim({caller: relayerA, bridgeTx: tokenTx, to: relayerA});
    }

    function test_claim_revert_statusRefunded() public {
        bytes32 txId = getTxId(tokenTx);
        bridge({caller: userA, msgValue: 0, params: tokenParams});
        skip(DEADLINE + 1);
        refund({caller: refunder, bridgeTx: tokenTx});
        vm.expectRevert(StatusIncorrect.selector);
        fastBridge.canClaim(txId, relayerA);
        vm.expectRevert(StatusIncorrect.selector);
        claim({caller: relayerA, bridgeTx: tokenTx, to: relayerA});
    }

    // ══════════════════════════════════════════════════ DISPUTE ══════════════════════════════════════════════════════

    function test_dispute_token() public {
        bytes32 txId = getTxId(tokenTx);
        bridge({caller: userA, msgValue: 0, params: tokenParams});
        prove({caller: relayerA, bridgeTx: tokenTx, destTxHash: hex"01"});
        expectBridgeProofDisputed({txId: txId, guard: guard});
        dispute({caller: guard, txId: txId});
        assertEq(fastBridge.bridgeStatuses(txId), IFastBridgeV2.BridgeStatus.REQUESTED);
        assertEq(fastBridge.protocolFees(address(srcToken)), INITIAL_PROTOCOL_FEES_TOKEN);
        assertEq(srcToken.balanceOf(address(fastBridge)), INITIAL_PROTOCOL_FEES_TOKEN + tokenParams.originAmount);
    }

    function test_dispute_token_justBeforeDeadline() public {
        bytes32 txId = getTxId(tokenTx);
        bridge({caller: userA, msgValue: 0, params: tokenParams});
        prove({caller: relayerA, bridgeTx: tokenTx, destTxHash: hex"01"});
        skip(CLAIM_DELAY);
        expectBridgeProofDisputed({txId: txId, guard: guard});
        dispute({caller: guard, txId: txId});
        assertEq(fastBridge.bridgeStatuses(txId), IFastBridgeV2.BridgeStatus.REQUESTED);
        assertEq(fastBridge.protocolFees(address(srcToken)), INITIAL_PROTOCOL_FEES_TOKEN);
        assertEq(srcToken.balanceOf(address(fastBridge)), INITIAL_PROTOCOL_FEES_TOKEN + tokenParams.originAmount);
    }

    function test_dispute_eth() public {
        bridge({caller: userA, msgValue: 0, params: tokenParams});
        bytes32 txId = getTxId(ethTx);
        bridge({caller: userA, msgValue: ethParams.originAmount, params: ethParams});
        prove({caller: relayerA, bridgeTx: ethTx, destTxHash: hex"01"});
        expectBridgeProofDisputed({txId: txId, guard: guard});
        dispute({caller: guard, txId: txId});
        assertEq(fastBridge.bridgeStatuses(txId), IFastBridgeV2.BridgeStatus.REQUESTED);
        assertEq(fastBridge.protocolFees(ETH_ADDRESS), INITIAL_PROTOCOL_FEES_ETH);
        assertEq(address(fastBridge).balance, INITIAL_PROTOCOL_FEES_ETH + ethParams.originAmount);
    }

    function test_dispute_eth_justBeforeDeadline() public {
        bridge({caller: userA, msgValue: 0, params: tokenParams});
        bytes32 txId = getTxId(ethTx);
        bridge({caller: userA, msgValue: ethParams.originAmount, params: ethParams});
        prove({caller: relayerA, bridgeTx: ethTx, destTxHash: hex"01"});
        skip(CLAIM_DELAY);
        expectBridgeProofDisputed({txId: txId, guard: guard});
        dispute({caller: guard, txId: txId});
        assertEq(fastBridge.bridgeStatuses(txId), IFastBridgeV2.BridgeStatus.REQUESTED);
        assertEq(fastBridge.protocolFees(ETH_ADDRESS), INITIAL_PROTOCOL_FEES_ETH);
        assertEq(address(fastBridge).balance, INITIAL_PROTOCOL_FEES_ETH + ethParams.originAmount);
    }

    function test_dispute_revert_afterDeadline() public {
        bytes32 txId = getTxId(tokenTx);
        bridge({caller: userA, msgValue: 0, params: tokenParams});
        prove({caller: relayerA, bridgeTx: tokenTx, destTxHash: hex"01"});
        skip(CLAIM_DELAY + 1);
        vm.expectRevert(DisputePeriodPassed.selector);
        dispute({caller: guard, txId: txId});
    }

    function test_dispute_revert_callerNotGuard(address caller) public {
        vm.assume(caller != guard);
        bytes32 txId = getTxId(tokenTx);
        bridge({caller: userA, msgValue: 0, params: tokenParams});
        prove({caller: relayerA, bridgeTx: tokenTx, destTxHash: hex"01"});
        expectUnauthorized(caller, fastBridge.GUARD_ROLE());
        dispute({caller: caller, txId: txId});
    }

    function test_dispute_revert_statusNull() public {
        bytes32 txId = getTxId(tokenTx);
        vm.expectRevert(StatusIncorrect.selector);
        dispute({caller: guard, txId: txId});
    }

    function test_dispute_revert_statusRequested() public {
        bytes32 txId = getTxId(tokenTx);
        bridge({caller: userA, msgValue: 0, params: tokenParams});
        vm.expectRevert(StatusIncorrect.selector);
        dispute({caller: guard, txId: txId});
    }

    function test_dispute_revert_statusClaimed() public {
        bytes32 txId = getTxId(tokenTx);
        bridge({caller: userA, msgValue: 0, params: tokenParams});
        prove({caller: relayerA, bridgeTx: tokenTx, destTxHash: hex"01"});
        skip(CLAIM_DELAY + 1);
        claim({caller: relayerA, bridgeTx: tokenTx, to: relayerA});
        vm.expectRevert(StatusIncorrect.selector);
        dispute({caller: guard, txId: txId});
    }

    function test_dispute_revert_statusRefunded() public {
        bytes32 txId = getTxId(tokenTx);
        bridge({caller: userA, msgValue: 0, params: tokenParams});
        skip(DEADLINE + 1);
        refund({caller: refunder, bridgeTx: tokenTx});
        vm.expectRevert(StatusIncorrect.selector);
        dispute({caller: guard, txId: txId});
    }

    // ══════════════════════════════════════════════════ REFUND ═══════════════════════════════════════════════════════

    function test_refund_token() public {
        bytes32 txId = getTxId(tokenTx);
        bridge({caller: userA, msgValue: 0, params: tokenParams});
        skip(DEADLINE + 1);
        expectBridgeDepositRefunded({bridgeParams: tokenParams, txId: txId});
        refund({caller: refunder, bridgeTx: tokenTx});
        assertEq(fastBridge.bridgeStatuses(txId), IFastBridgeV2.BridgeStatus.REFUNDED);
        assertEq(fastBridge.protocolFees(address(srcToken)), INITIAL_PROTOCOL_FEES_TOKEN);
        assertEq(srcToken.balanceOf(userA), LEFTOVER_BALANCE + tokenParams.originAmount);
        assertEq(srcToken.balanceOf(address(fastBridge)), INITIAL_PROTOCOL_FEES_TOKEN);
    }

    /// @notice Deposit should be refunded to the BridgeParams.sender, regardless of the actual caller
    function test_refund_token_diffSender() public {
        bytes32 txId = getTxId(tokenTx);
        bridge({caller: userB, msgValue: 0, params: tokenParams});
        skip(DEADLINE + 1);
        expectBridgeDepositRefunded({bridgeParams: tokenParams, txId: txId});
        refund({caller: refunder, bridgeTx: tokenTx});
        assertEq(fastBridge.bridgeStatuses(txId), IFastBridgeV2.BridgeStatus.REFUNDED);
        assertEq(fastBridge.protocolFees(address(srcToken)), INITIAL_PROTOCOL_FEES_TOKEN);
        assertEq(srcToken.balanceOf(userA), LEFTOVER_BALANCE + 2 * tokenParams.originAmount);
        assertEq(srcToken.balanceOf(userB), LEFTOVER_BALANCE);
        assertEq(srcToken.balanceOf(address(fastBridge)), INITIAL_PROTOCOL_FEES_TOKEN);
    }

    function test_refund_token_longDelay() public {
        bytes32 txId = getTxId(tokenTx);
        bridge({caller: userA, msgValue: 0, params: tokenParams});
        skip(DEADLINE + 30 days);
        expectBridgeDepositRefunded({bridgeParams: tokenParams, txId: txId});
        refund({caller: refunder, bridgeTx: tokenTx});
        assertEq(fastBridge.bridgeStatuses(txId), IFastBridgeV2.BridgeStatus.REFUNDED);
        assertEq(fastBridge.protocolFees(address(srcToken)), INITIAL_PROTOCOL_FEES_TOKEN);
        assertEq(srcToken.balanceOf(userA), LEFTOVER_BALANCE + tokenParams.originAmount);
        assertEq(srcToken.balanceOf(address(fastBridge)), INITIAL_PROTOCOL_FEES_TOKEN);
    }

    function test_refund_token_permisionless(address caller) public {
        vm.assume(caller != refunder);
        bytes32 txId = getTxId(tokenTx);
        bridge({caller: userA, msgValue: 0, params: tokenParams});
        skip(DEADLINE + PERMISSIONLESS_REFUND_DELAY + 1);
        expectBridgeDepositRefunded({bridgeParams: tokenParams, txId: txId});
        refund({caller: caller, bridgeTx: tokenTx});
        assertEq(fastBridge.bridgeStatuses(txId), IFastBridgeV2.BridgeStatus.REFUNDED);
        assertEq(fastBridge.protocolFees(address(srcToken)), INITIAL_PROTOCOL_FEES_TOKEN);
        assertEq(srcToken.balanceOf(userA), LEFTOVER_BALANCE + tokenParams.originAmount);
        assertEq(srcToken.balanceOf(address(fastBridge)), INITIAL_PROTOCOL_FEES_TOKEN);
    }

    function test_refund_eth() public {
        bridge({caller: userA, msgValue: 0, params: tokenParams});
        bytes32 txId = getTxId(ethTx);
        bridge({caller: userA, msgValue: ethParams.originAmount, params: ethParams});
        skip(DEADLINE + 1);
        expectBridgeDepositRefunded({bridgeParams: ethParams, txId: txId});
        refund({caller: refunder, bridgeTx: ethTx});
        assertEq(fastBridge.bridgeStatuses(txId), IFastBridgeV2.BridgeStatus.REFUNDED);
        assertEq(fastBridge.protocolFees(ETH_ADDRESS), INITIAL_PROTOCOL_FEES_ETH);
        assertEq(address(userA).balance, LEFTOVER_BALANCE + ethParams.originAmount);
        assertEq(address(fastBridge).balance, INITIAL_PROTOCOL_FEES_ETH);
    }

    /// @notice Deposit should be refunded to the BridgeParams.sender, regardless of the actual caller
    function test_refund_eth_diffSender() public {
        bridge({caller: userA, msgValue: 0, params: tokenParams});
        bytes32 txId = getTxId(ethTx);
        bridge({caller: userB, msgValue: ethParams.originAmount, params: ethParams});
        skip(DEADLINE + 1);
        expectBridgeDepositRefunded({bridgeParams: ethParams, txId: txId});
        refund({caller: refunder, bridgeTx: ethTx});
        assertEq(fastBridge.bridgeStatuses(txId), IFastBridgeV2.BridgeStatus.REFUNDED);
        assertEq(fastBridge.protocolFees(ETH_ADDRESS), INITIAL_PROTOCOL_FEES_ETH);
        assertEq(address(userA).balance, LEFTOVER_BALANCE + 2 * ethParams.originAmount);
        assertEq(address(userB).balance, LEFTOVER_BALANCE);
        assertEq(address(fastBridge).balance, INITIAL_PROTOCOL_FEES_ETH);
    }

    function test_refund_eth_longDelay() public {
        bridge({caller: userA, msgValue: 0, params: tokenParams});
        bytes32 txId = getTxId(ethTx);
        bridge({caller: userA, msgValue: ethParams.originAmount, params: ethParams});
        skip(DEADLINE + 30 days);
        expectBridgeDepositRefunded({bridgeParams: ethParams, txId: txId});
        refund({caller: refunder, bridgeTx: ethTx});
        assertEq(fastBridge.bridgeStatuses(txId), IFastBridgeV2.BridgeStatus.REFUNDED);
        assertEq(fastBridge.protocolFees(ETH_ADDRESS), INITIAL_PROTOCOL_FEES_ETH);
        assertEq(address(userA).balance, LEFTOVER_BALANCE + ethParams.originAmount);
        assertEq(address(fastBridge).balance, INITIAL_PROTOCOL_FEES_ETH);
    }

    function test_refund_eth_permisionless(address caller) public {
        vm.assume(caller != refunder);
        bytes32 txId = getTxId(ethTx);
        bridge({caller: userA, msgValue: 0, params: tokenParams});
        bridge({caller: userA, msgValue: ethParams.originAmount, params: ethParams});
        skip(DEADLINE + PERMISSIONLESS_REFUND_DELAY + 1);
        expectBridgeDepositRefunded({bridgeParams: ethParams, txId: txId});
        refund({caller: caller, bridgeTx: ethTx});
        assertEq(fastBridge.bridgeStatuses(txId), IFastBridgeV2.BridgeStatus.REFUNDED);
        assertEq(fastBridge.protocolFees(ETH_ADDRESS), INITIAL_PROTOCOL_FEES_ETH);
        assertEq(address(userA).balance, LEFTOVER_BALANCE + ethParams.originAmount);
        assertEq(address(fastBridge).balance, INITIAL_PROTOCOL_FEES_ETH);
    }

    function test_refund_revert_zeroDelay() public {
        bridge({caller: userA, msgValue: 0, params: tokenParams});
        vm.expectRevert(DeadlineNotExceeded.selector);
        refund({caller: refunder, bridgeTx: tokenTx});
    }

    function test_refund_revert_justBeforeDeadline() public {
        bridge({caller: userA, msgValue: 0, params: tokenParams});
        skip(DEADLINE);
        vm.expectRevert(DeadlineNotExceeded.selector);
        refund({caller: refunder, bridgeTx: tokenTx});
    }

    function test_refund_revert_justBeforeDeadline_permisionless(address caller) public {
        vm.assume(caller != refunder);
        bridge({caller: userA, msgValue: 0, params: tokenParams});
        skip(DEADLINE + PERMISSIONLESS_REFUND_DELAY);
        vm.expectRevert(DeadlineNotExceeded.selector);
        refund({caller: caller, bridgeTx: tokenTx});
    }

    function test_refund_revert_statusNull() public {
        vm.expectRevert(StatusIncorrect.selector);
        refund({caller: refunder, bridgeTx: ethTx});
    }

    function test_refund_revert_statusProven() public {
        bridge({caller: userA, msgValue: 0, params: tokenParams});
        prove({caller: relayerA, bridgeTx: tokenTx, destTxHash: hex"01"});
        vm.expectRevert(StatusIncorrect.selector);
        refund({caller: refunder, bridgeTx: tokenTx});
    }

    function test_refund_revert_statusClaimed() public {
        bridge({caller: userA, msgValue: 0, params: tokenParams});
        prove({caller: relayerA, bridgeTx: tokenTx, destTxHash: hex"01"});
        skip(CLAIM_DELAY + 1);
        claim({caller: relayerA, bridgeTx: tokenTx, to: relayerA});
        vm.expectRevert(StatusIncorrect.selector);
        refund({caller: refunder, bridgeTx: tokenTx});
    }

    function test_refund_revert_statusRefunded() public {
        bridge({caller: userA, msgValue: 0, params: tokenParams});
        skip(DEADLINE + 1);
        refund({caller: refunder, bridgeTx: tokenTx});
        vm.expectRevert(StatusIncorrect.selector);
        refund({caller: refunder, bridgeTx: tokenTx});
    }
}
