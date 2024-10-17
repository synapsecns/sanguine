// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {FastBridgeV2, FastBridgeV2Test, IFastBridge, IFastBridgeV2} from "./FastBridgeV2.t.sol";

// solhint-disable func-name-mixedcase, ordering
contract FastBridgeV2EncodingTest is FastBridgeV2Test {
    function setUp() public virtual override {
        vm.chainId(SRC_CHAIN_ID);
        super.setUp();
    }

    function deployFastBridge() public virtual override returns (FastBridgeV2) {
        return new FastBridgeV2(address(this));
    }

    function assertEq(IFastBridge.BridgeTransaction memory a, IFastBridge.BridgeTransaction memory b) public pure {
        assertEq(a.originChainId, b.originChainId);
        assertEq(a.destChainId, b.destChainId);
        assertEq(a.originSender, b.originSender);
        assertEq(a.destRecipient, b.destRecipient);
        assertEq(a.originToken, b.originToken);
        assertEq(a.destToken, b.destToken);
        assertEq(a.originAmount, b.originAmount);
        assertEq(a.destAmount, b.destAmount);
        assertEq(a.originFeeAmount, b.originFeeAmount);
        assertEq(a.sendChainGas, b.sendChainGas);
        assertEq(a.deadline, b.deadline);
        assertEq(a.nonce, b.nonce);
    }

    function assertEq(
        IFastBridgeV2.BridgeTransactionV2 memory a,
        IFastBridgeV2.BridgeTransactionV2 memory b
    )
        public
        pure
    {
        assertEq(extractV1(a), extractV1(b));
        assertEq(a.exclusivityRelayer, b.exclusivityRelayer);
        assertEq(a.exclusivityEndTime, b.exclusivityEndTime);
    }

    function test_getBridgeTransaction(IFastBridge.BridgeTransaction memory bridgeTx) public view {
        bytes memory request = abi.encode(bridgeTx);
        IFastBridge.BridgeTransaction memory decodedTx = fastBridge.getBridgeTransaction(request);
        assertEq(decodedTx, bridgeTx);
    }

    /// @notice We expect all the V1 fields except for `sendChainGas` to match.
    /// `sendChainGas` is replaced with `callValue` in V2, therefore we expect non-zero `callValue`
    /// to match `sendChainGas = true` in V1
    function test_getBridgeTransaction_supportsV2(IFastBridgeV2.BridgeTransactionV2 memory bridgeTxV2) public view {
        bytes memory request = abi.encode(bridgeTxV2);
        IFastBridge.BridgeTransaction memory decodedTx = fastBridge.getBridgeTransaction(request);
        IFastBridge.BridgeTransaction memory expectedTx = extractV1(bridgeTxV2);
        expectedTx.sendChainGas = bridgeTxV2.callValue > 0;
        assertEq(decodedTx, expectedTx);
    }

    function test_getBridgeTransactionV2(IFastBridgeV2.BridgeTransactionV2 memory bridgeTxV2) public view {
        bytes memory request = abi.encode(bridgeTxV2);
        IFastBridgeV2.BridgeTransactionV2 memory decodedTxV2 = fastBridge.getBridgeTransactionV2(request);
        assertEq(decodedTxV2, bridgeTxV2);
    }

    function test_getBridgeTransactionV2_revert_usedRequestV1(IFastBridge.BridgeTransaction memory bridgeTx) public {
        bytes memory request = abi.encode(bridgeTx);
        vm.expectRevert();
        fastBridge.getBridgeTransactionV2(request);
    }
}
