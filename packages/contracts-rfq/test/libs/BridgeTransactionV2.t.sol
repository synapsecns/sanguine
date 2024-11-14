// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import {BridgeTransactionV2Harness, IFastBridgeV2} from "../harnesses/BridgeTransactionV2Harness.sol";

import {Test} from "forge-std/Test.sol";

// solhint-disable func-name-mixedcase, ordering
contract BridgeTransactionV2Test is Test {
    BridgeTransactionV2Harness internal harness;

    function setUp() public {
        harness = new BridgeTransactionV2Harness();
    }

    function assertEq(
        IFastBridgeV2.BridgeTransactionV2 memory a,
        IFastBridgeV2.BridgeTransactionV2 memory b
    )
        public
        pure
    {
        assertEq(a.originChainId, b.originChainId);
        assertEq(a.destChainId, b.destChainId);
        assertEq(a.originSender, b.originSender);
        assertEq(a.destRecipient, b.destRecipient);
        assertEq(a.originToken, b.originToken);
        assertEq(a.destToken, b.destToken);
        assertEq(a.originAmount, b.originAmount);
        assertEq(a.destAmount, b.destAmount);
        assertEq(a.originFeeAmount, b.originFeeAmount);
        assertEq(a.zapNative, b.zapNative);
        assertEq(a.deadline, b.deadline);
        assertEq(a.nonce, b.nonce);
        assertEq(a.exclusivityRelayer, b.exclusivityRelayer);
        assertEq(a.exclusivityEndTime, b.exclusivityEndTime);
        assertEq(a.zapData, b.zapData);
    }

    function test_roundtrip(IFastBridgeV2.BridgeTransactionV2 memory bridgeTx) public view {
        bytes memory encodedTx = harness.encodeV2(bridgeTx);
        assertEq(harness.version(encodedTx), 2);
        assertEq(harness.originChainId(encodedTx), bridgeTx.originChainId);
        assertEq(harness.destChainId(encodedTx), bridgeTx.destChainId);
        assertEq(harness.originSender(encodedTx), bridgeTx.originSender);
        assertEq(harness.destRecipient(encodedTx), bridgeTx.destRecipient);
        assertEq(harness.originToken(encodedTx), bridgeTx.originToken);
        assertEq(harness.destToken(encodedTx), bridgeTx.destToken);
        assertEq(harness.originAmount(encodedTx), bridgeTx.originAmount);
        assertEq(harness.destAmount(encodedTx), bridgeTx.destAmount);
        assertEq(harness.originFeeAmount(encodedTx), bridgeTx.originFeeAmount);
        assertEq(harness.zapNative(encodedTx), bridgeTx.zapNative);
        assertEq(harness.deadline(encodedTx), bridgeTx.deadline);
        assertEq(harness.nonce(encodedTx), bridgeTx.nonce);
        assertEq(harness.exclusivityRelayer(encodedTx), bridgeTx.exclusivityRelayer);
        assertEq(harness.exclusivityEndTime(encodedTx), bridgeTx.exclusivityEndTime);
        assertEq(harness.zapData(encodedTx), bridgeTx.zapData);
    }

    function test_roundtrip_decodeV2(IFastBridgeV2.BridgeTransactionV2 memory bridgeTx) public view {
        bytes memory encodedTx = harness.encodeV2(bridgeTx);
        IFastBridgeV2.BridgeTransactionV2 memory decodedTx = harness.decodeV2(encodedTx);
        assertEq(decodedTx, bridgeTx);
    }
}
