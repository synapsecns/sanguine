// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {FastBridgeV2DstBaseTest, IFastBridgeV2} from "./FastBridgeV2.Dst.Base.t.sol";

import {ExcessiveReturnValueRecipient} from "./mocks/ExcessiveReturnValueRecipient.sol";
import {IncorrectReturnValueRecipient} from "./mocks/IncorrectReturnValueRecipient.sol";
import {NoOpContract} from "./mocks/NoOpContract.sol";
import {NoReturnValueRecipient} from "./mocks/NoReturnValueRecipient.sol";
import {NonPayableRecipient} from "./mocks/NonPayableRecipient.sol";

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

    uint256 public constant ZAP_NATIVE = 1_337_420;

    address public excessiveReturnValueRecipient;
    address public incorrectReturnValueRecipient;
    address public noOpRecipient;
    address public noReturnValueRecipient;
    address public nonPayableRecipient;

    function setUp() public virtual override {
        super.setUp();
        excessiveReturnValueRecipient = address(new ExcessiveReturnValueRecipient());
        vm.label(excessiveReturnValueRecipient, "ExcessiveReturnValueRecipient");
        incorrectReturnValueRecipient = address(new IncorrectReturnValueRecipient());
        vm.label(incorrectReturnValueRecipient, "IncorrectReturnValueRecipient");
        noOpRecipient = address(new NoOpContract());
        vm.label(noOpRecipient, "NoOpRecipient");
        noReturnValueRecipient = address(new NoReturnValueRecipient());
        vm.label(noReturnValueRecipient, "NoReturnValueRecipient");
        nonPayableRecipient = address(new NonPayableRecipient());
        vm.label(nonPayableRecipient, "NonPayableRecipient");
    }

    function setTokenTestRecipient(address recipient) public {
        userB = recipient;
        tokenParams.to = recipient;
        tokenTx.destRecipient = recipient;
    }

    function setEthTestRecipient(address recipient) public {
        userB = recipient;
        ethParams.to = recipient;
        ethTx.destRecipient = recipient;
    }

    function assertEmptyZapData(bytes memory zapData) public pure {
        assertEq(zapData.length, 0, "Invalid setup: zapData are not empty");
    }

    function expectBridgeRelayed(
        IFastBridgeV2.BridgeTransactionV2 memory bridgeTx,
        bytes32 txId,
        address relayer
    )
        public
        virtual
    {
        uint256 chainGasAmount = bridgeTx.destToken == ETH_ADDRESS ? 0 : bridgeTx.zapNative;
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
            chainGasAmount: chainGasAmount
        });
    }

    function checkRelayedViews(bytes32 txId, address expectedRelayer) public view {
        assertTrue(fastBridge.bridgeRelays(txId));
        (uint48 blockNumber, uint48 blockTimestamp, address relayer) = fastBridge.bridgeRelayDetails(txId);
        assertEq(blockNumber, block.number);
        assertEq(blockTimestamp, block.timestamp);
        assertEq(relayer, expectedRelayer);
    }

    function checkTokenBalances(address recipient, address relayCaller) public view {
        assertEq(dstToken.balanceOf(recipient), tokenParams.destAmount);
        assertEq(dstToken.balanceOf(relayCaller), LEFTOVER_BALANCE);
        assertEq(dstToken.balanceOf(address(fastBridge)), 0);
    }

    function checkEthBalances(address recipient, address relayCaller) public view {
        assertEq(recipient.balance, ethParams.destAmount);
        assertEq(relayCaller.balance, LEFTOVER_BALANCE);
        assertEq(address(fastBridge).balance, 0);
    }

    /// @notice RelayerA completes the ERC20 bridge request
    function test_relay_token() public {
        bytes32 txId = getTxId(tokenTx);
        expectBridgeRelayed(tokenTx, txId, relayerA);
        relay({caller: relayerA, msgValue: 0, bridgeTx: tokenTx});
        checkRelayedViews({txId: txId, expectedRelayer: relayerA});
        checkTokenBalances({recipient: userB, relayCaller: relayerA});
    }

    /// @notice RelayerB completes the ERC20 bridge request, using relayerA's address
    function test_relay_token_withRelayerAddress() public {
        bytes32 txId = getTxId(tokenTx);
        expectBridgeRelayed(tokenTx, txId, relayerA);
        relayWithAddress({caller: relayerB, relayer: relayerA, msgValue: 0, bridgeTx: tokenTx});
        checkRelayedViews({txId: txId, expectedRelayer: relayerA});
        checkTokenBalances({recipient: userB, relayCaller: relayerB});
    }

    /// @notice RelayerB completes the ETH bridge request
    function test_relay_eth() public {
        bytes32 txId = getTxId(ethTx);
        expectBridgeRelayed(ethTx, txId, relayerB);
        relay({caller: relayerB, msgValue: ethParams.destAmount, bridgeTx: ethTx});
        checkRelayedViews({txId: txId, expectedRelayer: relayerB});
        checkEthBalances({recipient: userB, relayCaller: relayerB});
    }

    /// @notice RelayerA completes the ETH bridge request, using relayerB's address
    function test_relay_eth_withRelayerAddress() public {
        bytes32 txId = getTxId(ethTx);
        expectBridgeRelayed(ethTx, txId, relayerB);
        relayWithAddress({caller: relayerA, relayer: relayerB, msgValue: ethParams.destAmount, bridgeTx: ethTx});
        checkRelayedViews({txId: txId, expectedRelayer: relayerB});
        checkEthBalances({recipient: userB, relayCaller: relayerA});
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
        checkEthBalances({recipient: userB, relayCaller: relayerA});
    }

    // ══════════════════════════════════════════ RELAYS WITH CALL VALUE ═══════════════════════════════════════════════

    function test_relay_token_withZapNative() public {
        setTokenTestZapNative(ZAP_NATIVE);
        bytes32 txId = getTxId(tokenTx);
        expectBridgeRelayed(tokenTx, txId, relayerA);
        relay({caller: relayerA, msgValue: ZAP_NATIVE, bridgeTx: tokenTx});
        checkRelayedViews({txId: txId, expectedRelayer: relayerA});
        checkTokenBalances({recipient: userB, relayCaller: relayerA});
        assertEq(userB.balance, ZAP_NATIVE);
    }

    function test_relay_token_withRelayerAddressZapNative() public {
        setTokenTestZapNative(ZAP_NATIVE);
        bytes32 txId = getTxId(tokenTx);
        expectBridgeRelayed(tokenTx, txId, relayerA);
        relayWithAddress({caller: relayerB, relayer: relayerA, msgValue: ZAP_NATIVE, bridgeTx: tokenTx});
        checkRelayedViews({txId: txId, expectedRelayer: relayerA});
        checkTokenBalances({recipient: userB, relayCaller: relayerB});
        assertEq(userB.balance, ZAP_NATIVE);
    }

    // ═════════════════════════════════════ EXCESSIVE RETURN VALUE RECIPIENT ══════════════════════════════════════════

    // Note: in this test, the zapData are not present, and the below test functions succeed.
    // Override them in the derived tests where zapData are present to check for a revert.

    function test_relay_token_excessiveReturnValueRecipient_revertWhenZapDataPresent() public virtual {
        assertEmptyZapData(tokenTx.zapData);
        setTokenTestRecipient(excessiveReturnValueRecipient);
        test_relay_token();
    }

    function test_relay_token_withRelayerAddress_excessiveReturnValueRecipient_revertWhenZapDataPresent()
        public
        virtual
    {
        assertEmptyZapData(tokenTx.zapData);
        setTokenTestRecipient(excessiveReturnValueRecipient);
        test_relay_token_withRelayerAddress();
    }

    function test_relay_token_withZapNative_excessiveReturnValueRecipient_revertWhenZapDataPresent() public virtual {
        assertEmptyZapData(tokenTx.zapData);
        setTokenTestZapNative(ZAP_NATIVE);
        setTokenTestRecipient(excessiveReturnValueRecipient);
        test_relay_token_withZapNative();
    }

    function test_relay_token_withRelayerAddressZapNative_excessiveReturnValueRecipient_revertWhenZapDataPresent()
        public
        virtual
    {
        assertEmptyZapData(tokenTx.zapData);
        setTokenTestZapNative(ZAP_NATIVE);
        setTokenTestRecipient(excessiveReturnValueRecipient);
        test_relay_token_withRelayerAddressZapNative();
    }

    function test_relay_eth_excessiveReturnValueRecipient_revertWhenZapDataPresent() public virtual {
        assertEmptyZapData(ethTx.zapData);
        setEthTestRecipient(excessiveReturnValueRecipient);
        test_relay_eth();
    }

    function test_relay_eth_withRelayerAddress_excessiveReturnValueRecipient_revertWhenZapDataPresent()
        public
        virtual
    {
        assertEmptyZapData(ethTx.zapData);
        setEthTestRecipient(excessiveReturnValueRecipient);
        test_relay_eth_withRelayerAddress();
    }

    // ═════════════════════════════════════ INCORRECT RETURN VALUE RECIPIENT ══════════════════════════════════════════

    // Note: in this test, the zapData are not present, and the below test functions succeed.
    // Override them in the derived tests where zapData are present to check for a revert.

    function test_relay_token_incorrectReturnValueRecipient_revertWhenZapDataPresent() public virtual {
        assertEmptyZapData(tokenTx.zapData);
        setTokenTestRecipient(incorrectReturnValueRecipient);
        test_relay_token();
    }

    function test_relay_token_withRelayerAddress_incorrectReturnValueRecipient_revertWhenZapDataPresent()
        public
        virtual
    {
        assertEmptyZapData(tokenTx.zapData);
        setTokenTestRecipient(incorrectReturnValueRecipient);
        test_relay_token_withRelayerAddress();
    }

    function test_relay_token_withZapNative_incorrectReturnValueRecipient_revertWhenZapDataPresent() public virtual {
        assertEmptyZapData(tokenTx.zapData);
        setTokenTestZapNative(ZAP_NATIVE);
        setTokenTestRecipient(incorrectReturnValueRecipient);
        test_relay_token_withZapNative();
    }

    function test_relay_token_withRelayerAddressZapNative_incorrectReturnValueRecipient_revertWhenZapDataPresent()
        public
        virtual
    {
        assertEmptyZapData(tokenTx.zapData);
        setTokenTestZapNative(ZAP_NATIVE);
        setTokenTestRecipient(incorrectReturnValueRecipient);
        test_relay_token_withRelayerAddressZapNative();
    }

    function test_relay_eth_incorrectReturnValueRecipient_revertWhenZapDataPresent() public virtual {
        assertEmptyZapData(ethTx.zapData);
        setEthTestRecipient(incorrectReturnValueRecipient);
        test_relay_eth();
    }

    function test_relay_eth_withRelayerAddress_incorrectReturnValueRecipient_revertWhenZapDataPresent()
        public
        virtual
    {
        assertEmptyZapData(ethTx.zapData);
        setEthTestRecipient(incorrectReturnValueRecipient);
        test_relay_eth_withRelayerAddress();
    }

    // ═══════════════════════════════════════════ NON PAYABLE RECIPIENT ═══════════════════════════════════════════════

    /// @notice Should not affect the ERC20 transfer
    function test_relay_token_nonPayableRecipient() public {
        setTokenTestRecipient(nonPayableRecipient);
        test_relay_token();
    }

    function test_relay_token_withRelayerAddress_nonPayableRecipient() public {
        setTokenTestRecipient(nonPayableRecipient);
        test_relay_token_withRelayerAddress();
    }

    function test_relay_token_withZapNative_nonPayableRecipient_revert() public {
        setTokenTestZapNative(ZAP_NATIVE);
        setTokenTestRecipient(nonPayableRecipient);
        vm.expectRevert();
        relay({caller: relayerA, msgValue: ZAP_NATIVE, bridgeTx: tokenTx});
    }

    function test_relay_token_withRelayerAddressZapNative_nonPayableRecipient_revert() public {
        setTokenTestZapNative(ZAP_NATIVE);
        setTokenTestRecipient(nonPayableRecipient);
        vm.expectRevert();
        relayWithAddress({caller: relayerB, relayer: relayerA, msgValue: ZAP_NATIVE, bridgeTx: tokenTx});
    }

    function test_relay_eth_revert_nonPayableRecipient() public {
        setEthTestRecipient(nonPayableRecipient);
        vm.expectRevert();
        relay({caller: relayerB, msgValue: ethParams.destAmount, bridgeTx: ethTx});
    }

    function test_relay_eth_withRelayerAddress_revert_nonPayableRecipient() public {
        setEthTestRecipient(nonPayableRecipient);
        vm.expectRevert();
        relayWithAddress({caller: relayerA, relayer: relayerB, msgValue: ethParams.destAmount, bridgeTx: ethTx});
    }

    // ══════════════════════════════════════════════ NO-OP RECIPIENT ══════════════════════════════════════════════════

    // Note: in this test, the zapData are not present, and the below test functions succeed.
    // Override them in the derived tests where zapData are present to check for a revert.

    function test_relay_token_noOpRecipient_revertWhenZapDataPresent() public virtual {
        assertEmptyZapData(tokenTx.zapData);
        setTokenTestRecipient(noOpRecipient);
        test_relay_token();
    }

    function test_relay_token_withRelayerAddress_noOpRecipient_revertWhenZapDataPresent() public virtual {
        assertEmptyZapData(tokenTx.zapData);
        setTokenTestRecipient(noOpRecipient);
        test_relay_token_withRelayerAddress();
    }

    function test_relay_token_withZapNative_noOpRecipient_revertWhenZapDataPresent() public virtual {
        assertEmptyZapData(tokenTx.zapData);
        setTokenTestZapNative(ZAP_NATIVE);
        setTokenTestRecipient(noOpRecipient);
        test_relay_token_withZapNative();
    }

    function test_relay_token_withRelayerAddressZapNative_noOpRecipient_revertWhenZapDataPresent() public virtual {
        assertEmptyZapData(tokenTx.zapData);
        setTokenTestZapNative(ZAP_NATIVE);
        setTokenTestRecipient(noOpRecipient);
        test_relay_token_withRelayerAddressZapNative();
    }

    function test_relay_eth_noOpRecipient_revertWhenZapDataPresent() public virtual {
        assertEmptyZapData(ethTx.zapData);
        setEthTestRecipient(noOpRecipient);
        test_relay_eth();
    }

    function test_relay_eth_withRelayerAddress_noOpRecipient_revertWhenZapDataPresent() public virtual {
        assertEmptyZapData(ethTx.zapData);
        setEthTestRecipient(noOpRecipient);
        test_relay_eth_withRelayerAddress();
    }

    // ═════════════════════════════════════════ NO RETURN VALUE RECIPIENT ═════════════════════════════════════════════

    // Note: in this test, the zapData are not present, and the below test functions succeed.
    // Override them in the derived tests where zapData are present to check for a revert.

    function test_relay_token_noReturnValueRecipient_revertWhenZapDataPresent() public virtual {
        assertEmptyZapData(tokenTx.zapData);
        setTokenTestRecipient(noReturnValueRecipient);
        test_relay_token();
    }

    function test_relay_token_withRelayerAddress_noReturnValueRecipient_revertWhenZapDataPresent() public virtual {
        assertEmptyZapData(tokenTx.zapData);
        setTokenTestRecipient(noReturnValueRecipient);
        test_relay_token_withRelayerAddress();
    }

    function test_relay_token_withZapNative_noReturnValueRecipient_revertWhenZapDataPresent() public virtual {
        assertEmptyZapData(tokenTx.zapData);
        setTokenTestZapNative(ZAP_NATIVE);
        setTokenTestRecipient(noReturnValueRecipient);
        test_relay_token_withZapNative();
    }

    function test_relay_token_withRelayerAddressZapNative_noReturnValueRecipient_revertWhenZapDataPresent()
        public
        virtual
    {
        assertEmptyZapData(tokenTx.zapData);
        setTokenTestZapNative(ZAP_NATIVE);
        setTokenTestRecipient(noReturnValueRecipient);
        test_relay_token_withRelayerAddressZapNative();
    }

    function test_relay_eth_noReturnValueRecipient_revertWhenZapDataPresent() public virtual {
        assertEmptyZapData(ethTx.zapData);
        setEthTestRecipient(noReturnValueRecipient);
        test_relay_eth();
    }

    function test_relay_eth_withRelayerAddress_noReturnValueRecipient_revertWhenZapDataPresent() public virtual {
        assertEmptyZapData(ethTx.zapData);
        setEthTestRecipient(noReturnValueRecipient);
        test_relay_eth_withRelayerAddress();
    }

    // ══════════════════════════════════════════════════ REVERTS ══════════════════════════════════════════════════════

    function test_relay_revert_requestV1() public {
        // V1 doesn't have any version field
        expectRevertUnsupportedVersion(0);
        vm.prank({msgSender: relayerA, txOrigin: relayerA});
        fastBridge.relay(mockRequestV1);
    }

    function test_relay_revert_invalidRequestV2() public {
        expectRevertInvalidEncodedTx();
        vm.prank({msgSender: relayerA, txOrigin: relayerA});
        fastBridge.relay(invalidRequestV2);
    }

    function test_relay_revert_requestV3() public {
        expectRevertUnsupportedVersion(3);
        vm.prank({msgSender: relayerA, txOrigin: relayerA});
        fastBridge.relay(mockRequestV3);
    }

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

    function test_relay_withRelayerAddress_revert_requestV1() public {
        // V1 doesn't have any version field
        expectRevertUnsupportedVersion(0);
        vm.prank({msgSender: relayerA, txOrigin: relayerA});
        fastBridge.relay(mockRequestV1, relayerB);
    }

    function test_relay_withRelayerAddress_revert_invalidRequestV2() public {
        expectRevertInvalidEncodedTx();
        vm.prank({msgSender: relayerA, txOrigin: relayerA});
        fastBridge.relay(invalidRequestV2, relayerB);
    }

    function test_relay_withRelayerAddress_revert_requestV3() public {
        expectRevertUnsupportedVersion(3);
        vm.prank({msgSender: relayerA, txOrigin: relayerA});
        fastBridge.relay(mockRequestV3, relayerB);
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

    function test_relay_token_withZapNative_revert_zeroZapNative() public {
        setTokenTestZapNative(ZAP_NATIVE);
        vm.expectRevert(MsgValueIncorrect.selector);
        relay({caller: relayerA, msgValue: 0, bridgeTx: tokenTx});
    }

    function test_relay_token_withZapNative_revert_lowerZapNative() public {
        setTokenTestZapNative(ZAP_NATIVE);
        vm.expectRevert(MsgValueIncorrect.selector);
        relay({caller: relayerA, msgValue: ZAP_NATIVE - 1, bridgeTx: tokenTx});
    }

    function test_relay_token_withZapNative_revert_higherZapNative() public {
        setTokenTestZapNative(ZAP_NATIVE);
        vm.expectRevert(MsgValueIncorrect.selector);
        relay({caller: relayerA, msgValue: ZAP_NATIVE + 1, bridgeTx: tokenTx});
    }

    function test_relay_token_withRelayerAddressZapNative_revert_zeroZapNative() public {
        setTokenTestZapNative(ZAP_NATIVE);
        vm.expectRevert(MsgValueIncorrect.selector);
        relayWithAddress({caller: relayerB, relayer: relayerA, msgValue: 0, bridgeTx: tokenTx});
    }

    function test_relay_token_withRelayerAddressZapNative_revert_lowerZapNative() public {
        setTokenTestZapNative(ZAP_NATIVE);
        vm.expectRevert(MsgValueIncorrect.selector);
        relayWithAddress({caller: relayerB, relayer: relayerA, msgValue: ZAP_NATIVE - 1, bridgeTx: tokenTx});
    }

    function test_relay_token_withRelayerAddressZapNative_revert_higherZapNative() public {
        setTokenTestZapNative(ZAP_NATIVE);
        vm.expectRevert(MsgValueIncorrect.selector);
        relayWithAddress({caller: relayerB, relayer: relayerA, msgValue: ZAP_NATIVE + 1, bridgeTx: tokenTx});
    }

    function test_relay_eth_withZapNative_revert_notSupported() public {
        setEthTestZapNative(ZAP_NATIVE);
        // Neither destAmount, ZAP_NATIVE, nor destAmount + ZAP_NATIVE should work here
        vm.expectRevert(ZapNativeNotSupported.selector);
        relay({caller: relayerB, msgValue: ZAP_NATIVE, bridgeTx: ethTx});
        vm.expectRevert(ZapNativeNotSupported.selector);
        relay({caller: relayerB, msgValue: ethParams.destAmount, bridgeTx: ethTx});
        vm.expectRevert(ZapNativeNotSupported.selector);
        relay({caller: relayerB, msgValue: ethParams.destAmount + ZAP_NATIVE, bridgeTx: ethTx});
    }

    function test_relay_eth_withRelayerAddressZapNative_revert_notSupported() public {
        setEthTestZapNative(ZAP_NATIVE);
        // Neither destAmount, ZAP_NATIVE, nor destAmount + ZAP_NATIVE should work here
        vm.expectRevert(ZapNativeNotSupported.selector);
        relayWithAddress({caller: relayerA, relayer: relayerB, msgValue: ZAP_NATIVE, bridgeTx: ethTx});
        vm.expectRevert(ZapNativeNotSupported.selector);
        relayWithAddress({caller: relayerA, relayer: relayerB, msgValue: ethParams.destAmount, bridgeTx: ethTx});
        vm.expectRevert(ZapNativeNotSupported.selector);
        relayWithAddress({
            caller: relayerA,
            relayer: relayerB,
            msgValue: ethParams.destAmount + ZAP_NATIVE,
            bridgeTx: ethTx
        });
    }

    function test_relay_token_revert_approvedZero() public {
        vm.prank(relayerA);
        dstToken.approve(address(fastBridge), 0);
        vm.expectRevert();
        relay({caller: relayerA, msgValue: 0, bridgeTx: tokenTx});
    }

    function test_relay_token_revert_approvedNotEnough() public {
        vm.prank(relayerA);
        dstToken.approve(address(fastBridge), tokenParams.destAmount - 1);
        vm.expectRevert();
        relay({caller: relayerA, msgValue: 0, bridgeTx: tokenTx});
    }

    function test_relay_token_revert_nonZeroMsgValue() public {
        vm.expectRevert();
        relay({caller: relayerA, msgValue: tokenParams.destAmount, bridgeTx: tokenTx});
    }

    function test_relay_token_withRelayerAddress_revert_approvedZero() public {
        vm.prank(relayerB);
        dstToken.approve(address(fastBridge), 0);
        vm.expectRevert();
        relayWithAddress({caller: relayerB, relayer: relayerA, msgValue: 0, bridgeTx: tokenTx});
    }

    function test_relay_token_withRelayerAddress_revert_approvedNotEnough() public {
        vm.prank(relayerB);
        dstToken.approve(address(fastBridge), tokenParams.destAmount - 1);
        vm.expectRevert();
        relayWithAddress({caller: relayerB, relayer: relayerA, msgValue: 0, bridgeTx: tokenTx});
    }

    function test_relay_token_withRelayerAddress_revert_nonZeroMsgValue() public {
        vm.expectRevert();
        relayWithAddress({caller: relayerB, relayer: relayerA, msgValue: tokenParams.destAmount, bridgeTx: tokenTx});
    }

    function test_relay_eth_revert_lowerMsgValue() public {
        vm.expectRevert();
        relay({caller: relayerB, msgValue: ethParams.destAmount - 1, bridgeTx: ethTx});
    }

    function test_relay_eth_revert_higherMsgValue() public {
        vm.expectRevert();
        relay({caller: relayerB, msgValue: ethParams.destAmount + 1, bridgeTx: ethTx});
    }

    function test_relay_eth_revert_zeroMsgValue() public {
        vm.expectRevert();
        relay({caller: relayerB, msgValue: 0, bridgeTx: ethTx});
    }

    function test_relay_eth_withRelayerAddress_revert_lowerMsgValue() public {
        vm.expectRevert();
        relayWithAddress({caller: relayerA, relayer: relayerB, msgValue: ethParams.destAmount - 1, bridgeTx: ethTx});
    }

    function test_relay_eth_withRelayerAddress_revert_higherMsgValue() public {
        vm.expectRevert();
        relayWithAddress({caller: relayerA, relayer: relayerB, msgValue: ethParams.destAmount + 1, bridgeTx: ethTx});
    }

    function test_relay_eth_withRelayerAddress_revert_zeroMsgValue() public {
        vm.expectRevert();
        relayWithAddress({caller: relayerA, relayer: relayerB, msgValue: 0, bridgeTx: ethTx});
    }
}
