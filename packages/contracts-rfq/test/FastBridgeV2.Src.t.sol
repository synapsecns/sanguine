// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {BridgeTransactionV2Lib} from "../contracts/libs/BridgeTransactionV2.sol";
import {TokenNotContract} from "../contracts/libs/Errors.sol";

import {FastBridgeV2SrcBaseTest, IFastBridge, IFastBridgeV2} from "./FastBridgeV2.Src.Base.t.sol";

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

    event BridgeQuoteDetails(bytes32 indexed transactionId, bytes quoteId);

    event DisputePenaltyTimeApplied(address prover, uint256 inactiveUntilTimestamp);

    address public claimTo = makeAddr("Claim To");

    function expectBridgeRequested(IFastBridgeV2.BridgeTransactionV2 memory bridgeTx, bytes32 txId) public {
        vm.expectEmit(address(fastBridge));
        emit BridgeRequested({
            transactionId: txId,
            sender: bridgeTx.originSender,
            request: BridgeTransactionV2Lib.encodeV2(bridgeTx),
            destChainId: bridgeTx.destChainId,
            originToken: bridgeTx.originToken,
            destToken: bridgeTx.destToken,
            originAmount: bridgeTx.originAmount,
            destAmount: bridgeTx.destAmount,
            sendChainGas: bridgeTx.zapNative > 0
        });
    }

    function expectBridgeQuoteDetails(bytes32 txId, bytes memory quoteId) public {
        vm.expectEmit(address(fastBridge));
        emit BridgeQuoteDetails({transactionId: txId, quoteId: quoteId});
    }

    function expectBridgeProofProvided(bytes32 txId, address relayer, bytes32 destTxHash) public {
        vm.expectEmit(address(fastBridge));
        emit BridgeProofProvided({transactionId: txId, relayer: relayer, transactionHash: destTxHash});
    }

    function expectBridgeDepositClaimed(
        IFastBridgeV2.BridgeTransactionV2 memory bridgeTx,
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

    function expectBridgeProofDisputed(bytes32 txId, address relayer) public {
        vm.expectEmit(address(fastBridge));
        emit BridgeProofDisputed({transactionId: txId, relayer: relayer});
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

    function expectDisputePenaltyTimeApplied(address prover) public {
        uint256 inactiveUntilTimestamp = block.timestamp + DISPUTE_PENALTY_TIME;
        vm.expectEmit(address(fastBridge));
        emit DisputePenaltyTimeApplied(prover, inactiveUntilTimestamp);
    }

    // ══════════════════════════════════════════════════ BRIDGE ═══════════════════════════════════════════════════════

    function checkStatusAndProofAfterBridge(bytes32 txId) public view {
        assertEq(fastBridge.bridgeStatuses(txId), IFastBridgeV2.BridgeStatus.REQUESTED);
        (
            IFastBridgeV2.BridgeStatus status,
            uint32 destChainId,
            uint16 proverID,
            uint256 proofBlockTimestamp,
            address proofRelayer
        ) = fastBridge.bridgeTxDetails(txId);
        assertEq(status, IFastBridgeV2.BridgeStatus.REQUESTED);
        assertEq(destChainId, DST_CHAIN_ID);
        assertEq(proverID, 0);
        assertEq(proofBlockTimestamp, 0);
        assertEq(proofRelayer, address(0));
        (proofBlockTimestamp, proofRelayer) = fastBridge.bridgeProofs(txId);
        assertEq(proofBlockTimestamp, 0);
        assertEq(proofRelayer, address(0));
    }

    function checkTokenBalancesAfterBridge(address caller) public view {
        assertEq(fastBridge.protocolFees(address(srcToken)), INITIAL_PROTOCOL_FEES_TOKEN);
        assertEq(srcToken.balanceOf(caller), LEFTOVER_BALANCE);
        assertEq(srcToken.balanceOf(address(fastBridge)), INITIAL_PROTOCOL_FEES_TOKEN + tokenParams.originAmount);
    }

    function test_bridge_token() public {
        bytes32 txId = getTxId(tokenTx);
        expectBridgeRequested(tokenTx, txId);
        expectBridgeQuoteDetails(txId, tokenParamsV2.quoteId);
        bridge({caller: userA, msgValue: 0, params: tokenParams});
        assertEq(fastBridge.senderNonces(userA), 1);
        assertEq(fastBridge.senderNonces(userB), 0);
        checkStatusAndProofAfterBridge(txId);
        checkTokenBalancesAfterBridge(userA);
    }

    function test_bridge_token_diffSender() public {
        bytes32 txId = getTxId(tokenTx);
        expectBridgeRequested(tokenTx, txId);
        expectBridgeQuoteDetails(txId, tokenParamsV2.quoteId);
        bridge({caller: userB, msgValue: 0, params: tokenParams});
        assertEq(fastBridge.senderNonces(userA), 1);
        assertEq(fastBridge.senderNonces(userB), 0);
        checkStatusAndProofAfterBridge(txId);
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
        assertEq(fastBridge.senderNonces(userA), 1);
        assertEq(fastBridge.senderNonces(userB), 0);
        bytes32 txId = getTxId(ethTx);
        expectBridgeRequested(ethTx, txId);
        expectBridgeQuoteDetails(txId, ethParamsV2.quoteId);
        bridge({caller: userA, msgValue: ethParams.originAmount, params: ethParams});
        assertEq(fastBridge.senderNonces(userA), 2);
        assertEq(fastBridge.senderNonces(userB), 0);
        checkStatusAndProofAfterBridge(txId);
        checkEthBalancesAfterBridge(userA);
    }

    function test_bridge_eth_diffSender() public {
        // bridge token first to match the nonce
        bridge({caller: userA, msgValue: 0, params: tokenParams});
        assertEq(fastBridge.senderNonces(userA), 1);
        assertEq(fastBridge.senderNonces(userB), 0);
        bytes32 txId = getTxId(ethTx);
        expectBridgeRequested(ethTx, txId);
        expectBridgeQuoteDetails(txId, ethParamsV2.quoteId);
        // bridge for user A as sender, called by userB
        bridge({caller: userB, msgValue: ethParams.originAmount, params: ethParams});
        assertEq(fastBridge.senderNonces(userA), 2);
        assertEq(fastBridge.senderNonces(userB), 0);
        checkStatusAndProofAfterBridge(txId);
        assertEq(userA.balance, LEFTOVER_BALANCE + ethParams.originAmount);
        checkEthBalancesAfterBridge(userB);
    }

    function test_bridge_userSpecificNonce() public {
        bridge({caller: userA, msgValue: 0, params: tokenParams});
        assertEq(fastBridge.senderNonces(userA), 1);
        assertEq(fastBridge.senderNonces(userB), 0);
        // UserB nonce is 0
        ethTx.nonce = 0;
        ethParams.sender = userB;
        ethTx.originSender = userB;
        bytes32 txId = getTxId(ethTx);
        expectBridgeRequested(ethTx, txId);
        expectBridgeQuoteDetails(txId, ethParamsV2.quoteId);
        bridge({caller: userB, msgValue: ethParams.originAmount, params: ethParams});
        assertEq(fastBridge.senderNonces(userA), 1);
        assertEq(fastBridge.senderNonces(userB), 1);
        checkStatusAndProofAfterBridge(txId);
        checkEthBalancesAfterBridge(userB);
    }

    function test_bridge_token_revert_approvedZero() public virtual {
        vm.prank(userA);
        srcToken.approve(address(fastBridge), 0);
        vm.expectRevert();
        bridge({caller: userA, msgValue: 0, params: tokenParams});
    }

    function test_bridge_token_revert_approvedNotEnough() public virtual {
        vm.prank(userA);
        srcToken.approve(address(fastBridge), tokenParams.originAmount - 1);
        vm.expectRevert();
        bridge({caller: userA, msgValue: 0, params: tokenParams});
    }

    function test_bridge_token_revert_nonZeroMsgValue() public {
        vm.expectRevert();
        bridge({caller: userA, msgValue: tokenParams.originAmount, params: tokenParams});
    }

    function test_bridge_eth_revert_lowerMsgValue() public {
        vm.expectRevert(MsgValueIncorrect.selector);
        bridge({caller: userA, msgValue: ethParams.originAmount - 1, params: ethParams});
    }

    function test_bridge_eth_revert_higherMsgValue() public {
        vm.expectRevert(MsgValueIncorrect.selector);
        bridge({caller: userA, msgValue: ethParams.originAmount + 1, params: ethParams});
    }

    function test_bridge_eth_revert_zeroMsgValue() public virtual {
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

    function test_bridge_revert_zeroOriginToken() public virtual {
        tokenParams.originToken = address(0);
        vm.expectRevert(ZeroAddress.selector);
        bridge({caller: userA, msgValue: 0, params: tokenParams});
    }

    function test_bridge_revert_zeroDestToken() public {
        tokenParams.destToken = address(0);
        vm.expectRevert(ZeroAddress.selector);
        bridge({caller: userA, msgValue: 0, params: tokenParams});
    }

    function test_bridge_revert_zeroSender() public virtual {
        tokenParams.sender = address(0);
        vm.expectRevert(ZeroAddress.selector);
        bridge({caller: userA, msgValue: 0, params: tokenParams});
    }

    function test_bridge_revert_zeroRecipient() public {
        tokenParams.to = address(0);
        vm.expectRevert(ZeroAddress.selector);
        bridge({caller: userA, msgValue: 0, params: tokenParams});
    }

    function test_bridge_revert_originTokenNotContract() public {
        tokenParams.originToken = makeAddr("Random EOA");
        vm.expectRevert(TokenNotContract.selector);
        bridge({caller: userA, msgValue: 0, params: tokenParams});
    }

    function test_bridge_revert_deadlineTooClose() public {
        tokenParams.deadline = block.timestamp + MIN_DEADLINE - 1;
        vm.expectRevert(DeadlineTooShort.selector);
        bridge({caller: userA, msgValue: 0, params: tokenParams});
    }

    function test_bridge_revert_weirdToken_amountHigher() public {
        tokenParams.originToken = address(weirdToken);
        weirdToken.setTransferToFastBridgeValue(tokenParams.originAmount + 1);
        vm.expectRevert(AmountIncorrect.selector);
        bridge({caller: userA, msgValue: 0, params: tokenParams});
    }

    function test_bridge_revert_weirdToken_amountLower() public {
        tokenParams.originToken = address(weirdToken);
        weirdToken.setTransferToFastBridgeValue(tokenParams.originAmount - 1);
        vm.expectRevert(AmountIncorrect.selector);
        bridge({caller: userA, msgValue: 0, params: tokenParams});
    }

    // ═══════════════════════════════════════════════════ PROVE ═══════════════════════════════════════════════════════

    function checkStatusAndProofAfterProve(bytes32 txId, uint16 expectedProverID, address relayer) public view {
        assertEq(fastBridge.bridgeStatuses(txId), IFastBridgeV2.BridgeStatus.RELAYER_PROVED);
        (
            IFastBridgeV2.BridgeStatus status,
            uint32 destChainId,
            uint16 proverID,
            uint256 proofBlockTimestamp,
            address proofRelayer
        ) = fastBridge.bridgeTxDetails(txId);
        assertEq(status, IFastBridgeV2.BridgeStatus.RELAYER_PROVED);
        assertEq(destChainId, DST_CHAIN_ID);
        assertEq(proverID, expectedProverID);
        assertEq(proofBlockTimestamp, block.timestamp);
        assertEq(proofRelayer, relayer);
        (proofBlockTimestamp, proofRelayer) = fastBridge.bridgeProofs(txId);
        assertEq(proofBlockTimestamp, block.timestamp);
        assertEq(proofRelayer, relayer);
    }

    function test_prove_token() public {
        bytes32 txId = getTxId(tokenTx);
        bridge({caller: userA, msgValue: 0, params: tokenParams});
        expectBridgeProofProvided({txId: txId, relayer: relayerA, destTxHash: hex"01"});
        prove({caller: relayerA, bridgeTx: tokenTx, destTxHash: hex"01"});
        checkStatusAndProofAfterProve(txId, 1, relayerA);
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
        checkStatusAndProofAfterProve(txId, 1, relayerA);
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
        cancel({caller: canceler, bridgeTx: tokenTx});
        vm.expectRevert(StatusIncorrect.selector);
        prove({caller: relayerA, bridgeTx: tokenTx, destTxHash: hex"01"});
    }

    function test_prove_revert_callerNotProver(address caller) public {
        vm.assume(caller != relayerA && caller != relayerB);
        bridge({caller: userA, msgValue: 0, params: tokenParams});
        vm.expectRevert(ProverNotActive.selector);
        prove({caller: caller, bridgeTx: tokenTx, destTxHash: hex"01"});
    }

    function test_prove_revert_disputePenaltyTime() public {
        bytes32 txId = getTxId(tokenTx);
        bridge({caller: userA, msgValue: 0, params: tokenParams});
        prove({caller: relayerA, bridgeTx: tokenTx, destTxHash: hex"01"});
        dispute({caller: guard, txId: txId});
        vm.expectRevert(ProverNotActive.selector);
        prove({caller: relayerA, bridgeTx: tokenTx, destTxHash: hex"02"});
        skip(DISPUTE_PENALTY_TIME - 1);
        vm.expectRevert(ProverNotActive.selector);
        prove({caller: relayerA, bridgeTx: tokenTx, destTxHash: hex"02"});
    }

    // ════════════════════════════════════════ PROVE OTHER RELAYER ════════════════════════════════════════════

    function test_proveOther_token() public {
        bytes32 txId = getTxId(tokenTx);
        bridge({caller: userA, msgValue: 0, params: tokenParams});
        expectBridgeProofProvided({txId: txId, relayer: relayerA, destTxHash: hex"01"});
        prove({caller: relayerB, transactionId: txId, destTxHash: hex"01", relayer: relayerA});
        checkStatusAndProofAfterProve(txId, 2, relayerA);
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
        checkStatusAndProofAfterProve(txId, 2, relayerA);
        assertEq(address(fastBridge).balance, INITIAL_PROTOCOL_FEES_ETH + ethParams.originAmount);
        assertEq(fastBridge.protocolFees(ETH_ADDRESS), INITIAL_PROTOCOL_FEES_ETH);
    }

    // relayer self-proving using tx id, which is capable of proving for another & most tests focus on that angle.
    function test_proveOther_self() public {
        bytes32 txId = getTxId(tokenTx);
        bridge({caller: userA, msgValue: 0, params: tokenParams});
        expectBridgeProofProvided({txId: txId, relayer: relayerA, destTxHash: hex"01"});
        prove({caller: relayerA, transactionId: txId, destTxHash: hex"01", relayer: relayerA});
        checkStatusAndProofAfterProve(txId, 1, relayerA);
        assertEq(srcToken.balanceOf(address(fastBridge)), INITIAL_PROTOCOL_FEES_TOKEN + tokenParams.originAmount);
        assertEq(fastBridge.protocolFees(address(srcToken)), INITIAL_PROTOCOL_FEES_TOKEN);
    }

    // arbitrary non-privileged address can be asserted as the relayer
    function test_proveOther_permless() public {
        bytes32 txId = getTxId(tokenTx);

        bridge({caller: userA, msgValue: 0, params: tokenParams});
        expectBridgeProofProvided({txId: txId, relayer: address(0x1234), destTxHash: hex"01"});
        prove({caller: relayerA, transactionId: txId, destTxHash: hex"01", relayer: address(0x1234)});
        checkStatusAndProofAfterProve(txId, 1, address(0x1234));
        assertEq(srcToken.balanceOf(address(fastBridge)), INITIAL_PROTOCOL_FEES_TOKEN + tokenParams.originAmount);
        assertEq(fastBridge.protocolFees(address(srcToken)), INITIAL_PROTOCOL_FEES_TOKEN);
    }

    function test_proveOther_afterDispute() public {
        bytes32 txId = getTxId(tokenTx);
        bridge({caller: userA, msgValue: 0, params: tokenParams});
        prove({caller: relayerA, relayer: relayerB, transactionId: txId, destTxHash: hex"01"});
        dispute({caller: guard, txId: txId});
        expectBridgeProofProvided({txId: txId, relayer: relayerA, destTxHash: hex"02"});
        prove({caller: relayerB, relayer: relayerA, transactionId: txId, destTxHash: hex"02"});
        checkStatusAndProofAfterProve(txId, 2, relayerA);
        assertEq(srcToken.balanceOf(address(fastBridge)), INITIAL_PROTOCOL_FEES_TOKEN + tokenParams.originAmount);
        assertEq(fastBridge.protocolFees(address(srcToken)), INITIAL_PROTOCOL_FEES_TOKEN);
    }

    function test_proveOther_reProveAfterDispute() public {
        bridge({caller: userA, msgValue: 0, params: tokenParams});
        bytes32 txId = getTxId(ethTx);
        bridge({caller: userA, msgValue: ethParams.originAmount, params: ethParams});
        expectBridgeProofProvided({txId: txId, relayer: relayerA, destTxHash: hex"01"});
        prove({caller: relayerB, transactionId: txId, destTxHash: hex"01", relayer: relayerA});
        checkStatusAndProofAfterProve(txId, 2, relayerA);
        expectBridgeProofDisputed(txId, relayerA);
        dispute(guard, txId);
        skip(DISPUTE_PENALTY_TIME);
        expectBridgeProofProvided({txId: txId, relayer: relayerA, destTxHash: hex"02"});
        prove({caller: relayerB, transactionId: txId, destTxHash: hex"02", relayer: relayerA});
        checkStatusAndProofAfterProve(txId, 2, relayerA);
        expectBridgeProofDisputed(txId, relayerA);
        dispute(guard, txId);
        skip(DISPUTE_PENALTY_TIME);
        expectBridgeProofProvided({txId: txId, relayer: relayerA, destTxHash: hex"03"});
        prove({caller: relayerB, transactionId: txId, destTxHash: hex"03", relayer: relayerA});
        checkStatusAndProofAfterProve(txId, 2, relayerA);
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
        checkStatusAndProofAfterProve(txId, 1, relayerA);
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
        cancel({caller: canceler, bridgeTx: tokenTx});
        vm.expectRevert(StatusIncorrect.selector);
        prove({caller: relayerB, transactionId: txId, destTxHash: hex"01", relayer: relayerA});
    }

    function test_proveOther_revert_callerNotAuthed(address caller) public {
        bytes32 txId = getTxId(tokenTx);
        vm.assume(caller != relayerA && caller != relayerB);
        bridge({caller: userA, msgValue: 0, params: tokenParams});
        vm.expectRevert(ProverNotActive.selector);
        prove({caller: caller, transactionId: txId, destTxHash: hex"01", relayer: relayerA});
    }

    function test_proveOther_revert_disputePenaltyTime() public {
        bytes32 txId = getTxId(tokenTx);
        bridge({caller: userA, msgValue: 0, params: tokenParams});
        prove({caller: relayerA, relayer: relayerB, transactionId: txId, destTxHash: hex"01"});
        dispute({caller: guard, txId: txId});
        vm.expectRevert(ProverNotActive.selector);
        prove({caller: relayerA, relayer: relayerB, transactionId: txId, destTxHash: hex"02"});
        skip(DISPUTE_PENALTY_TIME - 1);
        vm.expectRevert(ProverNotActive.selector);
        prove({caller: relayerA, relayer: relayerB, transactionId: txId, destTxHash: hex"02"});
    }

    // ═══════════════════════════════════════════════════ CLAIM ═══════════════════════════════════════════════════════

    function checkStatusAndProofAfterClaim(
        bytes32 txId,
        uint16 expectedProverID,
        address relayer,
        uint256 expectedProofTS
    )
        public
        view
    {
        assertEq(fastBridge.bridgeStatuses(txId), IFastBridgeV2.BridgeStatus.RELAYER_CLAIMED);
        (
            IFastBridgeV2.BridgeStatus status,
            uint32 destChainId,
            uint16 proverID,
            uint256 proofBlockTimestamp,
            address proofRelayer
        ) = fastBridge.bridgeTxDetails(txId);
        assertEq(status, IFastBridgeV2.BridgeStatus.RELAYER_CLAIMED);
        assertEq(destChainId, DST_CHAIN_ID);
        assertEq(proverID, expectedProverID);
        assertEq(proofBlockTimestamp, expectedProofTS);
        assertEq(proofRelayer, relayer);
        (proofBlockTimestamp, proofRelayer) = fastBridge.bridgeProofs(txId);
        assertEq(proofBlockTimestamp, expectedProofTS);
        assertEq(proofRelayer, relayer);
    }

    function checkTokenBalancesAfterClaim(address relayer) public view {
        uint256 expectedProtocolFees = INITIAL_PROTOCOL_FEES_TOKEN + tokenTx.originFeeAmount;
        assertEq(fastBridge.protocolFees(address(srcToken)), expectedProtocolFees);
        assertEq(srcToken.balanceOf(relayer), tokenTx.originAmount);
        assertEq(srcToken.balanceOf(address(fastBridge)), expectedProtocolFees);
    }

    function test_claim_token() public {
        bytes32 txId = getTxId(tokenTx);
        bridge({caller: userA, msgValue: 0, params: tokenParams});
        prove({caller: relayerA, bridgeTx: tokenTx, destTxHash: hex"01"});
        uint256 expectedProofTS = block.timestamp;
        skip(CLAIM_DELAY + 1);
        assertTrue(fastBridge.canClaim(txId, relayerA));
        expectBridgeDepositClaimed({bridgeTx: tokenTx, txId: txId, relayer: relayerA, to: relayerA});
        claim({caller: relayerA, bridgeTx: tokenTx, to: relayerA});
        checkStatusAndProofAfterClaim(txId, 1, relayerA, expectedProofTS);
        checkTokenBalancesAfterClaim(relayerA);
    }

    function test_claim_token_permissionless(address caller) public {
        vm.assume(caller != relayerA);
        bytes32 txId = getTxId(tokenTx);
        bridge({caller: userA, msgValue: 0, params: tokenParams});
        prove({caller: relayerA, bridgeTx: tokenTx, destTxHash: hex"01"});
        uint256 expectedProofTS = block.timestamp;
        skip(CLAIM_DELAY + 1);
        expectBridgeDepositClaimed({bridgeTx: tokenTx, txId: txId, relayer: relayerA, to: relayerA});
        claim({caller: caller, bridgeTx: tokenTx});
        checkStatusAndProofAfterClaim(txId, 1, relayerA, expectedProofTS);
        checkTokenBalancesAfterClaim(relayerA);
    }

    function test_claim_token_permissionless_toZeroAddress(address caller) public {
        vm.assume(caller != relayerA);
        bytes32 txId = getTxId(tokenTx);
        bridge({caller: userA, msgValue: 0, params: tokenParams});
        prove({caller: relayerA, bridgeTx: tokenTx, destTxHash: hex"01"});
        uint256 expectedProofTS = block.timestamp;
        skip(CLAIM_DELAY + 1);
        expectBridgeDepositClaimed({bridgeTx: tokenTx, txId: txId, relayer: relayerA, to: relayerA});
        claim({caller: caller, bridgeTx: tokenTx, to: address(0)});
        checkStatusAndProofAfterClaim(txId, 1, relayerA, expectedProofTS);
        checkTokenBalancesAfterClaim(relayerA);
    }

    function test_claim_token_toDiffAddress() public {
        bytes32 txId = getTxId(tokenTx);
        bridge({caller: userA, msgValue: 0, params: tokenParams});
        prove({caller: relayerA, bridgeTx: tokenTx, destTxHash: hex"01"});
        uint256 expectedProofTS = block.timestamp;
        skip(CLAIM_DELAY + 1);
        expectBridgeDepositClaimed({bridgeTx: tokenTx, txId: txId, relayer: relayerA, to: claimTo});
        claim({caller: relayerA, bridgeTx: tokenTx, to: claimTo});
        checkStatusAndProofAfterClaim(txId, 1, relayerA, expectedProofTS);
        assertEq(srcToken.balanceOf(relayerA), 0);
        checkTokenBalancesAfterClaim(claimTo);
    }

    function test_claim_token_longDelay() public {
        bytes32 txId = getTxId(tokenTx);
        bridge({caller: userA, msgValue: 0, params: tokenParams});
        prove({caller: relayerA, bridgeTx: tokenTx, destTxHash: hex"01"});
        uint256 expectedProofTS = block.timestamp;
        skip(CLAIM_DELAY + 30 days);
        expectBridgeDepositClaimed({bridgeTx: tokenTx, txId: txId, relayer: relayerA, to: relayerA});
        claim({caller: relayerA, bridgeTx: tokenTx, to: relayerA});
        checkStatusAndProofAfterClaim(txId, 1, relayerA, expectedProofTS);
        checkTokenBalancesAfterClaim(relayerA);
    }

    function checkEthBalancesAfterClaim(address relayer) public view {
        uint256 expectedProtocolFees = INITIAL_PROTOCOL_FEES_ETH + ethTx.originFeeAmount;
        assertEq(fastBridge.protocolFees(ETH_ADDRESS), expectedProtocolFees);
        assertEq(address(relayer).balance, ethTx.originAmount);
        assertEq(address(fastBridge).balance, expectedProtocolFees);
    }

    function test_claim_eth() public {
        bridge({caller: userA, msgValue: 0, params: tokenParams});
        bytes32 txId = getTxId(ethTx);
        bridge({caller: userA, msgValue: ethParams.originAmount, params: ethParams});
        prove({caller: relayerA, bridgeTx: ethTx, destTxHash: hex"01"});
        uint256 expectedProofTS = block.timestamp;
        skip(CLAIM_DELAY + 1);
        assertTrue(fastBridge.canClaim(txId, relayerA));
        expectBridgeDepositClaimed({bridgeTx: ethTx, txId: txId, relayer: relayerA, to: relayerA});
        claim({caller: relayerA, bridgeTx: ethTx, to: relayerA});
        checkStatusAndProofAfterClaim(txId, 1, relayerA, expectedProofTS);
        checkEthBalancesAfterClaim(relayerA);
    }

    function test_claim_eth_permissionless(address caller) public {
        vm.assume(caller != relayerA);
        bridge({caller: userA, msgValue: 0, params: tokenParams});
        bytes32 txId = getTxId(ethTx);
        bridge({caller: userA, msgValue: ethParams.originAmount, params: ethParams});
        prove({caller: relayerA, bridgeTx: ethTx, destTxHash: hex"01"});
        uint256 expectedProofTS = block.timestamp;
        skip(CLAIM_DELAY + 1);
        expectBridgeDepositClaimed({bridgeTx: ethTx, txId: txId, relayer: relayerA, to: relayerA});
        claim({caller: caller, bridgeTx: ethTx});
        checkStatusAndProofAfterClaim(txId, 1, relayerA, expectedProofTS);
        checkEthBalancesAfterClaim(relayerA);
    }

    function test_claim_eth_permissionless_toZeroAddress(address caller) public {
        vm.assume(caller != relayerA);
        bridge({caller: userA, msgValue: 0, params: tokenParams});
        bytes32 txId = getTxId(ethTx);
        bridge({caller: userA, msgValue: ethParams.originAmount, params: ethParams});
        prove({caller: relayerA, bridgeTx: ethTx, destTxHash: hex"01"});
        uint256 expectedProofTS = block.timestamp;
        skip(CLAIM_DELAY + 1);
        expectBridgeDepositClaimed({bridgeTx: ethTx, txId: txId, relayer: relayerA, to: relayerA});
        claim({caller: caller, bridgeTx: ethTx, to: address(0)});
        checkStatusAndProofAfterClaim(txId, 1, relayerA, expectedProofTS);
        checkEthBalancesAfterClaim(relayerA);
    }

    function test_claim_eth_toDiffAddress() public {
        bridge({caller: userA, msgValue: 0, params: tokenParams});
        bytes32 txId = getTxId(ethTx);
        bridge({caller: userA, msgValue: ethParams.originAmount, params: ethParams});
        prove({caller: relayerA, bridgeTx: ethTx, destTxHash: hex"01"});
        uint256 expectedProofTS = block.timestamp;
        skip(CLAIM_DELAY + 1);
        expectBridgeDepositClaimed({bridgeTx: ethTx, txId: txId, relayer: relayerA, to: claimTo});
        claim({caller: relayerA, bridgeTx: ethTx, to: claimTo});
        checkStatusAndProofAfterClaim(txId, 1, relayerA, expectedProofTS);
        checkEthBalancesAfterClaim(claimTo);
    }

    function test_claim_eth_longDelay() public {
        bridge({caller: userA, msgValue: 0, params: tokenParams});
        bytes32 txId = getTxId(ethTx);
        bridge({caller: userA, msgValue: ethParams.originAmount, params: ethParams});
        prove({caller: relayerA, bridgeTx: ethTx, destTxHash: hex"01"});
        uint256 expectedProofTS = block.timestamp;
        skip(CLAIM_DELAY + 30 days);
        expectBridgeDepositClaimed({bridgeTx: ethTx, txId: txId, relayer: relayerA, to: relayerA});
        claim({caller: relayerA, bridgeTx: ethTx, to: relayerA});
        checkStatusAndProofAfterClaim(txId, 1, relayerA, expectedProofTS);
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
        cancel({caller: canceler, bridgeTx: tokenTx});
        vm.expectRevert(StatusIncorrect.selector);
        fastBridge.canClaim(txId, relayerA);
        vm.expectRevert(StatusIncorrect.selector);
        claim({caller: relayerA, bridgeTx: tokenTx, to: relayerA});
    }

    // ══════════════════════════════════════════════════ DISPUTE ══════════════════════════════════════════════════════

    function checkStatusAndProofAfterDispute(bytes32 txId) public view {
        // Should be identical to a requested tx that was never proven
        checkStatusAndProofAfterBridge(txId);
    }

    function checkProver(address prover, uint16 expectedProverID, uint256 expectedActiveFromTimestamp) public view {
        (uint16 proverID, uint256 activeFromTimestamp) = fastBridge.getProverInfo(prover);
        assertEq(proverID, expectedProverID);
        assertEq(activeFromTimestamp, expectedActiveFromTimestamp);
        address p;
        (p, activeFromTimestamp) = fastBridge.getProverInfoByID(expectedProverID);
        assertEq(p, prover);
        assertEq(activeFromTimestamp, expectedActiveFromTimestamp);
    }

    function test_dispute_token() public {
        bytes32 txId = getTxId(tokenTx);
        bridge({caller: userA, msgValue: 0, params: tokenParams});
        prove({caller: relayerA, bridgeTx: tokenTx, destTxHash: hex"01"});
        expectDisputePenaltyTimeApplied({prover: relayerA});
        expectBridgeProofDisputed({txId: txId, relayer: relayerA});
        dispute({caller: guard, txId: txId});
        checkStatusAndProofAfterDispute(txId);
        assertEq(fastBridge.protocolFees(address(srcToken)), INITIAL_PROTOCOL_FEES_TOKEN);
        assertEq(srcToken.balanceOf(address(fastBridge)), INITIAL_PROTOCOL_FEES_TOKEN + tokenParams.originAmount);
        // Check disputed prover
        assertEq(fastBridge.getActiveProverID(relayerA), 0);
        checkProver(relayerA, 1, block.timestamp + DISPUTE_PENALTY_TIME);
    }

    function test_dispute_token_provedOther() public {
        bytes32 txId = getTxId(tokenTx);
        bridge({caller: userA, msgValue: 0, params: tokenParams});
        prove({caller: relayerA, relayer: relayerB, transactionId: txId, destTxHash: hex"01"});
        expectDisputePenaltyTimeApplied({prover: relayerA});
        expectBridgeProofDisputed({txId: txId, relayer: relayerB});
        dispute({caller: guard, txId: txId});
        checkStatusAndProofAfterDispute(txId);
        assertEq(fastBridge.protocolFees(address(srcToken)), INITIAL_PROTOCOL_FEES_TOKEN);
        assertEq(srcToken.balanceOf(address(fastBridge)), INITIAL_PROTOCOL_FEES_TOKEN + tokenParams.originAmount);
        // Check disputed prover
        assertEq(fastBridge.getActiveProverID(relayerA), 0);
        checkProver(relayerA, 1, block.timestamp + DISPUTE_PENALTY_TIME);
    }

    function test_dispute_token_proverAlreadyRemoved() public {
        bytes32 txId = getTxId(tokenTx);
        bridge({caller: userA, msgValue: 0, params: tokenParams});
        prove({caller: relayerA, relayer: relayerB, transactionId: txId, destTxHash: hex"01"});
        fastBridge.removeProver(relayerA);
        expectBridgeProofDisputed({txId: txId, relayer: relayerB});
        dispute({caller: guard, txId: txId});
        checkStatusAndProofAfterDispute(txId);
        assertEq(fastBridge.protocolFees(address(srcToken)), INITIAL_PROTOCOL_FEES_TOKEN);
        assertEq(srcToken.balanceOf(address(fastBridge)), INITIAL_PROTOCOL_FEES_TOKEN + tokenParams.originAmount);
        // Check disputed prover
        assertEq(fastBridge.getActiveProverID(relayerA), 0);
        checkProver(relayerA, 1, 0);
    }

    function test_dispute_token_justBeforeDeadline() public {
        bytes32 txId = getTxId(tokenTx);
        bridge({caller: userA, msgValue: 0, params: tokenParams});
        prove({caller: relayerA, bridgeTx: tokenTx, destTxHash: hex"01"});
        skip(CLAIM_DELAY);
        expectDisputePenaltyTimeApplied({prover: relayerA});
        expectBridgeProofDisputed({txId: txId, relayer: relayerA});
        dispute({caller: guard, txId: txId});
        checkStatusAndProofAfterDispute(txId);
        assertEq(fastBridge.protocolFees(address(srcToken)), INITIAL_PROTOCOL_FEES_TOKEN);
        assertEq(srcToken.balanceOf(address(fastBridge)), INITIAL_PROTOCOL_FEES_TOKEN + tokenParams.originAmount);
        // Check disputed prover
        assertEq(fastBridge.getActiveProverID(relayerA), 0);
        checkProver(relayerA, 1, block.timestamp + DISPUTE_PENALTY_TIME);
    }

    function test_dispute_eth() public {
        bridge({caller: userA, msgValue: 0, params: tokenParams});
        bytes32 txId = getTxId(ethTx);
        bridge({caller: userA, msgValue: ethParams.originAmount, params: ethParams});
        prove({caller: relayerB, bridgeTx: ethTx, destTxHash: hex"01"});
        expectDisputePenaltyTimeApplied({prover: relayerB});
        expectBridgeProofDisputed({txId: txId, relayer: relayerB});
        dispute({caller: guard, txId: txId});
        checkStatusAndProofAfterDispute(txId);
        assertEq(fastBridge.protocolFees(ETH_ADDRESS), INITIAL_PROTOCOL_FEES_ETH);
        assertEq(address(fastBridge).balance, INITIAL_PROTOCOL_FEES_ETH + ethParams.originAmount);
        // Check disputed prover
        assertEq(fastBridge.getActiveProverID(relayerB), 0);
        checkProver(relayerB, 2, block.timestamp + DISPUTE_PENALTY_TIME);
    }

    function test_dispute_eth_provedOther() public {
        bridge({caller: userA, msgValue: 0, params: tokenParams});
        bytes32 txId = getTxId(ethTx);
        bridge({caller: userA, msgValue: ethParams.originAmount, params: ethParams});
        prove({caller: relayerB, relayer: relayerA, transactionId: txId, destTxHash: hex"01"});
        expectDisputePenaltyTimeApplied({prover: relayerB});
        expectBridgeProofDisputed({txId: txId, relayer: relayerA});
        dispute({caller: guard, txId: txId});
        checkStatusAndProofAfterDispute(txId);
        assertEq(fastBridge.protocolFees(ETH_ADDRESS), INITIAL_PROTOCOL_FEES_ETH);
        assertEq(address(fastBridge).balance, INITIAL_PROTOCOL_FEES_ETH + ethParams.originAmount);
        // Check disputed prover
        assertEq(fastBridge.getActiveProverID(relayerB), 0);
        checkProver(relayerB, 2, block.timestamp + DISPUTE_PENALTY_TIME);
    }

    function test_dispute_eth_proverAlreadyRemoved() public {
        bridge({caller: userA, msgValue: 0, params: tokenParams});
        bytes32 txId = getTxId(ethTx);
        bridge({caller: userA, msgValue: ethParams.originAmount, params: ethParams});
        prove({caller: relayerB, relayer: relayerA, transactionId: txId, destTxHash: hex"01"});
        fastBridge.removeProver(relayerB);
        dispute({caller: guard, txId: txId});
        checkStatusAndProofAfterDispute(txId);
        assertEq(fastBridge.protocolFees(ETH_ADDRESS), INITIAL_PROTOCOL_FEES_ETH);
        assertEq(address(fastBridge).balance, INITIAL_PROTOCOL_FEES_ETH + ethParams.originAmount);
        // Check disputed prover
        assertEq(fastBridge.getActiveProverID(relayerB), 0);
        checkProver(relayerB, 2, 0);
    }

    function test_dispute_eth_justBeforeDeadline() public {
        bridge({caller: userA, msgValue: 0, params: tokenParams});
        bytes32 txId = getTxId(ethTx);
        bridge({caller: userA, msgValue: ethParams.originAmount, params: ethParams});
        prove({caller: relayerB, bridgeTx: ethTx, destTxHash: hex"01"});
        skip(CLAIM_DELAY);
        expectDisputePenaltyTimeApplied({prover: relayerB});
        expectBridgeProofDisputed({txId: txId, relayer: relayerB});
        dispute({caller: guard, txId: txId});
        checkStatusAndProofAfterDispute(txId);
        assertEq(fastBridge.protocolFees(ETH_ADDRESS), INITIAL_PROTOCOL_FEES_ETH);
        assertEq(address(fastBridge).balance, INITIAL_PROTOCOL_FEES_ETH + ethParams.originAmount);
        // Check disputed prover
        assertEq(fastBridge.getActiveProverID(relayerB), 0);
        checkProver(relayerB, 2, block.timestamp + DISPUTE_PENALTY_TIME);
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
        cancel({caller: canceler, bridgeTx: tokenTx});
        vm.expectRevert(StatusIncorrect.selector);
        dispute({caller: guard, txId: txId});
    }

    // ══════════════════════════════════════════════════ CANCEL ═══════════════════════════════════════════════════════

    function checkStatusAndProofAfterCancel(bytes32 txId) public view {
        assertEq(fastBridge.bridgeStatuses(txId), IFastBridgeV2.BridgeStatus.REFUNDED);
        (
            IFastBridgeV2.BridgeStatus status,
            uint32 destChainId,
            uint16 proverID,
            uint256 proofBlockTimestamp,
            address proofRelayer
        ) = fastBridge.bridgeTxDetails(txId);
        assertEq(status, IFastBridgeV2.BridgeStatus.REFUNDED);
        assertEq(destChainId, DST_CHAIN_ID);
        assertEq(proverID, 0);
        assertEq(proofBlockTimestamp, 0);
        assertEq(proofRelayer, address(0));
        (proofBlockTimestamp, proofRelayer) = fastBridge.bridgeProofs(txId);
        assertEq(proofBlockTimestamp, 0);
        assertEq(proofRelayer, address(0));
    }

    function test_cancel_token() public {
        bytes32 txId = getTxId(tokenTx);
        bridge({caller: userA, msgValue: 0, params: tokenParams});
        skip(DEADLINE + 1);
        expectBridgeDepositRefunded({bridgeParams: tokenParams, txId: txId});
        cancel({caller: canceler, bridgeTx: tokenTx});
        checkStatusAndProofAfterCancel(txId);
        assertEq(fastBridge.protocolFees(address(srcToken)), INITIAL_PROTOCOL_FEES_TOKEN);
        assertEq(srcToken.balanceOf(userA), LEFTOVER_BALANCE + tokenParams.originAmount);
        assertEq(srcToken.balanceOf(address(fastBridge)), INITIAL_PROTOCOL_FEES_TOKEN);
    }

    /// @notice Deposit should be refunded to the BridgeParams.sender, regardless of the actual caller
    function test_cancel_token_diffSender() public {
        bytes32 txId = getTxId(tokenTx);
        bridge({caller: userB, msgValue: 0, params: tokenParams});
        skip(DEADLINE + 1);
        expectBridgeDepositRefunded({bridgeParams: tokenParams, txId: txId});
        cancel({caller: canceler, bridgeTx: tokenTx});
        checkStatusAndProofAfterCancel(txId);
        assertEq(fastBridge.protocolFees(address(srcToken)), INITIAL_PROTOCOL_FEES_TOKEN);
        assertEq(srcToken.balanceOf(userA), LEFTOVER_BALANCE + 2 * tokenParams.originAmount);
        assertEq(srcToken.balanceOf(userB), LEFTOVER_BALANCE);
        assertEq(srcToken.balanceOf(address(fastBridge)), INITIAL_PROTOCOL_FEES_TOKEN);
    }

    function test_cancel_token_longDelay() public {
        bytes32 txId = getTxId(tokenTx);
        bridge({caller: userA, msgValue: 0, params: tokenParams});
        skip(DEADLINE + 30 days);
        expectBridgeDepositRefunded({bridgeParams: tokenParams, txId: txId});
        cancel({caller: canceler, bridgeTx: tokenTx});
        checkStatusAndProofAfterCancel(txId);
        assertEq(fastBridge.protocolFees(address(srcToken)), INITIAL_PROTOCOL_FEES_TOKEN);
        assertEq(srcToken.balanceOf(userA), LEFTOVER_BALANCE + tokenParams.originAmount);
        assertEq(srcToken.balanceOf(address(fastBridge)), INITIAL_PROTOCOL_FEES_TOKEN);
    }

    function test_cancel_token_permisionless(address caller) public {
        vm.assume(caller != canceler);
        bytes32 txId = getTxId(tokenTx);
        bridge({caller: userA, msgValue: 0, params: tokenParams});
        skip(DEADLINE + PERMISSIONLESS_CANCEL_DELAY + 1);
        expectBridgeDepositRefunded({bridgeParams: tokenParams, txId: txId});
        cancel({caller: caller, bridgeTx: tokenTx});
        checkStatusAndProofAfterCancel(txId);
        assertEq(fastBridge.protocolFees(address(srcToken)), INITIAL_PROTOCOL_FEES_TOKEN);
        assertEq(srcToken.balanceOf(userA), LEFTOVER_BALANCE + tokenParams.originAmount);
        assertEq(srcToken.balanceOf(address(fastBridge)), INITIAL_PROTOCOL_FEES_TOKEN);
    }

    function test_cancel_eth() public {
        bridge({caller: userA, msgValue: 0, params: tokenParams});
        bytes32 txId = getTxId(ethTx);
        bridge({caller: userA, msgValue: ethParams.originAmount, params: ethParams});
        skip(DEADLINE + 1);
        expectBridgeDepositRefunded({bridgeParams: ethParams, txId: txId});
        cancel({caller: canceler, bridgeTx: ethTx});
        checkStatusAndProofAfterCancel(txId);
        assertEq(fastBridge.protocolFees(ETH_ADDRESS), INITIAL_PROTOCOL_FEES_ETH);
        assertEq(userA.balance, LEFTOVER_BALANCE + ethParams.originAmount);
        assertEq(address(fastBridge).balance, INITIAL_PROTOCOL_FEES_ETH);
    }

    /// @notice Deposit should be refunded to the BridgeParams.sender, regardless of the actual caller
    function test_cancel_eth_diffSender() public {
        bridge({caller: userA, msgValue: 0, params: tokenParams});
        bytes32 txId = getTxId(ethTx);
        bridge({caller: userB, msgValue: ethParams.originAmount, params: ethParams});
        skip(DEADLINE + 1);
        expectBridgeDepositRefunded({bridgeParams: ethParams, txId: txId});
        cancel({caller: canceler, bridgeTx: ethTx});
        checkStatusAndProofAfterCancel(txId);
        assertEq(fastBridge.protocolFees(ETH_ADDRESS), INITIAL_PROTOCOL_FEES_ETH);
        assertEq(userA.balance, LEFTOVER_BALANCE + 2 * ethParams.originAmount);
        assertEq(userB.balance, LEFTOVER_BALANCE);
        assertEq(address(fastBridge).balance, INITIAL_PROTOCOL_FEES_ETH);
    }

    function test_cancel_eth_longDelay() public {
        bridge({caller: userA, msgValue: 0, params: tokenParams});
        bytes32 txId = getTxId(ethTx);
        bridge({caller: userA, msgValue: ethParams.originAmount, params: ethParams});
        skip(DEADLINE + 30 days);
        expectBridgeDepositRefunded({bridgeParams: ethParams, txId: txId});
        cancel({caller: canceler, bridgeTx: ethTx});
        checkStatusAndProofAfterCancel(txId);
        assertEq(fastBridge.protocolFees(ETH_ADDRESS), INITIAL_PROTOCOL_FEES_ETH);
        assertEq(userA.balance, LEFTOVER_BALANCE + ethParams.originAmount);
        assertEq(address(fastBridge).balance, INITIAL_PROTOCOL_FEES_ETH);
    }

    function test_cancel_eth_permisionless(address caller) public {
        vm.assume(caller != canceler);
        bytes32 txId = getTxId(ethTx);
        bridge({caller: userA, msgValue: 0, params: tokenParams});
        bridge({caller: userA, msgValue: ethParams.originAmount, params: ethParams});
        skip(DEADLINE + PERMISSIONLESS_CANCEL_DELAY + 1);
        expectBridgeDepositRefunded({bridgeParams: ethParams, txId: txId});
        cancel({caller: caller, bridgeTx: ethTx});
        checkStatusAndProofAfterCancel(txId);
        assertEq(fastBridge.protocolFees(ETH_ADDRESS), INITIAL_PROTOCOL_FEES_ETH);
        assertEq(userA.balance, LEFTOVER_BALANCE + ethParams.originAmount);
        assertEq(address(fastBridge).balance, INITIAL_PROTOCOL_FEES_ETH);
    }

    function test_cancel_revert_zeroDelay() public {
        bridge({caller: userA, msgValue: 0, params: tokenParams});
        vm.expectRevert(DeadlineNotExceeded.selector);
        cancel({caller: canceler, bridgeTx: tokenTx});
    }

    function test_cancel_revert_justBeforeDeadline() public {
        bridge({caller: userA, msgValue: 0, params: tokenParams});
        skip(DEADLINE);
        vm.expectRevert(DeadlineNotExceeded.selector);
        cancel({caller: canceler, bridgeTx: tokenTx});
    }

    function test_cancel_revert_justBeforeDeadline_permisionless(address caller) public {
        vm.assume(caller != canceler);
        bridge({caller: userA, msgValue: 0, params: tokenParams});
        skip(DEADLINE + PERMISSIONLESS_CANCEL_DELAY);
        vm.expectRevert(DeadlineNotExceeded.selector);
        cancel({caller: caller, bridgeTx: tokenTx});
    }

    function test_cancel_revert_statusNull() public {
        vm.expectRevert(StatusIncorrect.selector);
        cancel({caller: canceler, bridgeTx: ethTx});
    }

    function test_cancel_revert_statusProven() public {
        bridge({caller: userA, msgValue: 0, params: tokenParams});
        prove({caller: relayerA, bridgeTx: tokenTx, destTxHash: hex"01"});
        vm.expectRevert(StatusIncorrect.selector);
        cancel({caller: canceler, bridgeTx: tokenTx});
    }

    function test_cancel_revert_statusClaimed() public {
        bridge({caller: userA, msgValue: 0, params: tokenParams});
        prove({caller: relayerA, bridgeTx: tokenTx, destTxHash: hex"01"});
        skip(CLAIM_DELAY + 1);
        claim({caller: relayerA, bridgeTx: tokenTx, to: relayerA});
        vm.expectRevert(StatusIncorrect.selector);
        cancel({caller: canceler, bridgeTx: tokenTx});
    }

    function test_cancel_revert_statusRefunded() public {
        bridge({caller: userA, msgValue: 0, params: tokenParams});
        skip(DEADLINE + 1);
        cancel({caller: canceler, bridgeTx: tokenTx});
        vm.expectRevert(StatusIncorrect.selector);
        cancel({caller: canceler, bridgeTx: tokenTx});
    }

    // ═════════════════════════════════════════════ INVALID PAYLOADS ══════════════════════════════════════════════════

    function test_prove_revert_requestV1() public {
        // V1 doesn't have any version field
        expectRevertUnsupportedVersion(0);
        vm.prank({msgSender: relayerA, txOrigin: relayerA});
        fastBridge.prove(mockRequestV1, hex"01");
    }

    function test_prove_revert_invalidRequestV2() public {
        expectRevertInvalidEncodedTx();
        vm.prank({msgSender: relayerA, txOrigin: relayerA});
        fastBridge.prove(invalidRequestV2, hex"01");
    }

    function test_prove_revert_requestV3() public {
        expectRevertUnsupportedVersion(3);
        vm.prank({msgSender: relayerA, txOrigin: relayerA});
        fastBridge.prove(mockRequestV3, hex"01");
    }

    function test_claim_revert_requestV1() public {
        // V1 doesn't have any version field
        expectRevertUnsupportedVersion(0);
        vm.prank({msgSender: relayerA, txOrigin: relayerA});
        fastBridge.claimV2(mockRequestV1);
    }

    function test_claim_revert_invalidRequestV2() public {
        expectRevertInvalidEncodedTx();
        vm.prank({msgSender: relayerA, txOrigin: relayerA});
        fastBridge.claimV2(invalidRequestV2);
    }

    function test_claim_revert_requestV3() public {
        expectRevertUnsupportedVersion(3);
        vm.prank({msgSender: relayerA, txOrigin: relayerA});
        fastBridge.claimV2(mockRequestV3);
    }

    function test_claim_toDiffAddress_revert_requestV1() public {
        // V1 doesn't have any version field
        expectRevertUnsupportedVersion(0);
        vm.prank({msgSender: relayerA, txOrigin: relayerA});
        fastBridge.claim(mockRequestV1, relayerB);
    }

    function test_claim_toDiffAddress_revert_invalidRequestV2() public {
        expectRevertInvalidEncodedTx();
        vm.prank({msgSender: relayerA, txOrigin: relayerA});
        fastBridge.claim(invalidRequestV2, relayerB);
    }

    function test_claim_toDiffAddress_revert_requestV3() public {
        expectRevertUnsupportedVersion(3);
        vm.prank({msgSender: relayerA, txOrigin: relayerA});
        fastBridge.claim(mockRequestV3, relayerB);
    }

    function test_cancel_revert_requestV1() public {
        // V1 doesn't have any version field
        expectRevertUnsupportedVersion(0);
        vm.prank({msgSender: relayerA, txOrigin: relayerA});
        fastBridge.cancelV2(mockRequestV1);
    }

    function test_cancel_revert_invalidRequestV2() public {
        expectRevertInvalidEncodedTx();
        vm.prank({msgSender: relayerA, txOrigin: relayerA});
        fastBridge.cancelV2(invalidRequestV2);
    }

    function test_cancel_revert_requestV3() public {
        expectRevertUnsupportedVersion(3);
        vm.prank({msgSender: relayerA, txOrigin: relayerA});
        fastBridge.cancelV2(mockRequestV3);
    }
}
