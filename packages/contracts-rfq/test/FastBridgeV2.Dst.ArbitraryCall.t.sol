// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {IFastBridgeV2} from "../contracts/interfaces/IFastBridgeV2.sol";
import {FastBridgeV2DstExclusivityTest} from "./FastBridgeV2.Dst.Exclusivity.t.sol";
import {RecipientMock} from "./mocks/RecipientMock.sol";

import {Address} from "@openzeppelin/contracts/utils/Address.sol";

// solhint-disable func-name-mixedcase, ordering
contract FastBridgeV2DstArbitraryCallTest is FastBridgeV2DstExclusivityTest {
    bytes public constant ZAP_DATA = abi.encode("Hello, world!");
    bytes public constant REVERT_MSG = "GM, this is a revert";

    function createFixtures() public virtual override {
        // In the inherited tests userB is always used as the recipient of the tokens.
        userB = address(new RecipientMock());
        vm.label(userB, "ContractRecipient");
        super.createFixtures();
    }

    function createFixturesV2() public virtual override {
        super.createFixturesV2();
        setTokenTestZapData(ZAP_DATA);
        setEthTestZapData(ZAP_DATA);
    }

    /// @notice We override the "expect event" function to also check for the Zap
    /// made to the token recipient.
    function expectBridgeRelayed(
        IFastBridgeV2.BridgeTransactionV2 memory bridgeTx,
        bytes32 txId,
        address relayer
    )
        public
        virtual
        override
    {
        vm.expectCall({callee: userB, data: getExpectedCalldata(bridgeTx), count: 1});
        super.expectBridgeRelayed(bridgeTx, txId, relayer);
    }

    function mockRecipientRevert(IFastBridgeV2.BridgeTransactionV2 memory bridgeTx) public {
        vm.mockCallRevert({callee: userB, data: getExpectedCalldata(bridgeTx), revertData: bytes(REVERT_MSG)});
    }

    function getExpectedCalldata(IFastBridgeV2.BridgeTransactionV2 memory bridgeTx)
        public
        pure
        returns (bytes memory)
    {
        // fastBridgeTransferReceived(token, amount, zapData)
        return abi.encodeCall(
            RecipientMock.fastBridgeTransferReceived, (bridgeTx.destToken, bridgeTx.destAmount, ZAP_DATA)
        );
    }

    // ═══════════════════════════════════════════════ RECIPIENT EOA ═══════════════════════════════════════════════════

    function test_relay_token_revert_recipientNotContract() public {
        setTokenTestRecipient(userA);
        vm.expectRevert(abi.encodeWithSelector(Address.AddressEmptyCode.selector, userA));
        relay({caller: relayerA, msgValue: 0, bridgeTx: tokenTx});
    }

    function test_relay_token_withRelayerAddress_revert_recipientNotContract() public {
        setTokenTestRecipient(userA);
        vm.expectRevert(abi.encodeWithSelector(Address.AddressEmptyCode.selector, userA));
        relayWithAddress({caller: relayerB, relayer: relayerA, msgValue: 0, bridgeTx: tokenTx});
    }

    function test_relay_eth_revert_recipientNotContract() public {
        setEthTestRecipient(userA);
        vm.expectRevert(abi.encodeWithSelector(Address.AddressEmptyCode.selector, userA));
        relay({caller: relayerB, msgValue: ethParams.destAmount, bridgeTx: ethTx});
    }

    function test_relay_eth_withRelayerAddress_revert_recipientNotContract() public {
        setEthTestRecipient(userA);
        vm.expectRevert(abi.encodeWithSelector(Address.AddressEmptyCode.selector, userA));
        relayWithAddress({caller: relayerA, relayer: relayerB, msgValue: ethParams.destAmount, bridgeTx: ethTx});
    }

    // ═════════════════════════════════════ EXCESSIVE RETURN VALUE RECIPIENT ══════════════════════════════════════════

    function test_relay_token_excessiveReturnValueRecipient_revertWhenZapDataPresent() public virtual override {
        setTokenTestRecipient(excessiveReturnValueRecipient);
        vm.expectRevert(RecipientIncorrectReturnValue.selector);
        relay({caller: relayerA, msgValue: 0, bridgeTx: tokenTx});
    }

    function test_relay_token_withRelayerAddress_excessiveReturnValueRecipient_revertWhenZapDataPresent()
        public
        virtual
        override
    {
        setTokenTestRecipient(excessiveReturnValueRecipient);
        vm.expectRevert(RecipientIncorrectReturnValue.selector);
        relayWithAddress({caller: relayerB, relayer: relayerA, msgValue: 0, bridgeTx: tokenTx});
    }

    function test_relay_token_withZapNative_excessiveReturnValueRecipient_revertWhenZapDataPresent()
        public
        virtual
        override
    {
        setTokenTestZapNative(ZAP_NATIVE);
        setTokenTestRecipient(excessiveReturnValueRecipient);
        vm.expectRevert(RecipientIncorrectReturnValue.selector);
        relay({caller: relayerA, msgValue: ZAP_NATIVE, bridgeTx: tokenTx});
    }

    function test_relay_token_withRelayerAddressZapNative_excessiveReturnValueRecipient_revertWhenZapDataPresent()
        public
        virtual
        override
    {
        setTokenTestZapNative(ZAP_NATIVE);
        setTokenTestRecipient(excessiveReturnValueRecipient);
        vm.expectRevert(RecipientIncorrectReturnValue.selector);
        relayWithAddress({caller: relayerB, relayer: relayerA, msgValue: ZAP_NATIVE, bridgeTx: tokenTx});
    }

    function test_relay_eth_excessiveReturnValueRecipient_revertWhenZapDataPresent() public virtual override {
        setEthTestRecipient(excessiveReturnValueRecipient);
        vm.expectRevert(RecipientIncorrectReturnValue.selector);
        relay({caller: relayerB, msgValue: ethParams.destAmount, bridgeTx: ethTx});
    }

    function test_relay_eth_withRelayerAddress_excessiveReturnValueRecipient_revertWhenZapDataPresent()
        public
        virtual
        override
    {
        setEthTestRecipient(excessiveReturnValueRecipient);
        vm.expectRevert(RecipientIncorrectReturnValue.selector);
        relayWithAddress({caller: relayerA, relayer: relayerB, msgValue: ethParams.destAmount, bridgeTx: ethTx});
    }

    // ═════════════════════════════════════ INCORRECT RETURN VALUE RECIPIENT ══════════════════════════════════════════

    function test_relay_token_incorrectReturnValueRecipient_revertWhenZapDataPresent() public virtual override {
        setTokenTestRecipient(incorrectReturnValueRecipient);
        vm.expectRevert(RecipientIncorrectReturnValue.selector);
        relay({caller: relayerA, msgValue: 0, bridgeTx: tokenTx});
    }

    function test_relay_token_withRelayerAddress_incorrectReturnValueRecipient_revertWhenZapDataPresent()
        public
        virtual
        override
    {
        setTokenTestRecipient(incorrectReturnValueRecipient);
        vm.expectRevert(RecipientIncorrectReturnValue.selector);
        relayWithAddress({caller: relayerB, relayer: relayerA, msgValue: 0, bridgeTx: tokenTx});
    }

    function test_relay_token_withZapNative_incorrectReturnValueRecipient_revertWhenZapDataPresent()
        public
        virtual
        override
    {
        setTokenTestZapNative(ZAP_NATIVE);
        setTokenTestRecipient(incorrectReturnValueRecipient);
        vm.expectRevert(RecipientIncorrectReturnValue.selector);
        relay({caller: relayerA, msgValue: ZAP_NATIVE, bridgeTx: tokenTx});
    }

    function test_relay_token_withRelayerAddressZapNative_incorrectReturnValueRecipient_revertWhenZapDataPresent()
        public
        virtual
        override
    {
        setTokenTestZapNative(ZAP_NATIVE);
        setTokenTestRecipient(incorrectReturnValueRecipient);
        vm.expectRevert(RecipientIncorrectReturnValue.selector);
        relayWithAddress({caller: relayerB, relayer: relayerA, msgValue: ZAP_NATIVE, bridgeTx: tokenTx});
    }

    function test_relay_eth_incorrectReturnValueRecipient_revertWhenZapDataPresent() public virtual override {
        setEthTestRecipient(incorrectReturnValueRecipient);
        vm.expectRevert(RecipientIncorrectReturnValue.selector);
        relay({caller: relayerB, msgValue: ethParams.destAmount, bridgeTx: ethTx});
    }

    function test_relay_eth_withRelayerAddress_incorrectReturnValueRecipient_revertWhenZapDataPresent()
        public
        virtual
        override
    {
        setEthTestRecipient(incorrectReturnValueRecipient);
        vm.expectRevert(RecipientIncorrectReturnValue.selector);
        relayWithAddress({caller: relayerA, relayer: relayerB, msgValue: ethParams.destAmount, bridgeTx: ethTx});
    }

    // ══════════════════════════════════════════════ NO-OP RECIPIENT ══════════════════════════════════════════════════

    // Note: in these tests NoOpRecipient doesn't implement hook function, so we expect a generic OZ library revert.

    function test_relay_token_noOpRecipient_revertWhenZapDataPresent() public virtual override {
        setTokenTestRecipient(noOpRecipient);
        vm.expectRevert(Address.FailedInnerCall.selector);
        relay({caller: relayerA, msgValue: 0, bridgeTx: tokenTx});
    }

    function test_relay_token_withRelayerAddress_noOpRecipient_revertWhenZapDataPresent() public virtual override {
        setTokenTestRecipient(noOpRecipient);
        vm.expectRevert(Address.FailedInnerCall.selector);
        relayWithAddress({caller: relayerB, relayer: relayerA, msgValue: 0, bridgeTx: tokenTx});
    }

    function test_relay_token_withZapNative_noOpRecipient_revertWhenZapDataPresent() public virtual override {
        setTokenTestZapNative(ZAP_NATIVE);
        setTokenTestRecipient(noOpRecipient);
        vm.expectRevert(Address.FailedInnerCall.selector);
        relay({caller: relayerA, msgValue: ZAP_NATIVE, bridgeTx: tokenTx});
    }

    function test_relay_token_withRelayerAddressZapNative_noOpRecipient_revertWhenZapDataPresent()
        public
        virtual
        override
    {
        setTokenTestZapNative(ZAP_NATIVE);
        setTokenTestRecipient(noOpRecipient);
        vm.expectRevert(Address.FailedInnerCall.selector);
        relayWithAddress({caller: relayerB, relayer: relayerA, msgValue: ZAP_NATIVE, bridgeTx: tokenTx});
    }

    function test_relay_eth_noOpRecipient_revertWhenZapDataPresent() public virtual override {
        setEthTestRecipient(noOpRecipient);
        vm.expectRevert(Address.FailedInnerCall.selector);
        relay({caller: relayerB, msgValue: ethParams.destAmount, bridgeTx: ethTx});
    }

    function test_relay_eth_withRelayerAddress_noOpRecipient_revertWhenZapDataPresent() public virtual override {
        setEthTestRecipient(noOpRecipient);
        vm.expectRevert(Address.FailedInnerCall.selector);
        relayWithAddress({caller: relayerA, relayer: relayerB, msgValue: ethParams.destAmount, bridgeTx: ethTx});
    }

    // ═════════════════════════════════════════ NO RETURN VALUE RECIPIENT ═════════════════════════════════════════════

    function test_relay_token_noReturnValueRecipient_revertWhenZapDataPresent() public virtual override {
        setTokenTestRecipient(noReturnValueRecipient);
        vm.expectRevert(RecipientNoReturnValue.selector);
        relay({caller: relayerA, msgValue: 0, bridgeTx: tokenTx});
    }

    function test_relay_token_withRelayerAddress_noReturnValueRecipient_revertWhenZapDataPresent()
        public
        virtual
        override
    {
        setTokenTestRecipient(noReturnValueRecipient);
        vm.expectRevert(RecipientNoReturnValue.selector);
        relayWithAddress({caller: relayerB, relayer: relayerA, msgValue: 0, bridgeTx: tokenTx});
    }

    function test_relay_token_withZapNative_noReturnValueRecipient_revertWhenZapDataPresent() public virtual override {
        setTokenTestZapNative(ZAP_NATIVE);
        setTokenTestRecipient(noReturnValueRecipient);
        vm.expectRevert(RecipientNoReturnValue.selector);
        relay({caller: relayerA, msgValue: ZAP_NATIVE, bridgeTx: tokenTx});
    }

    function test_relay_token_withRelayerAddressZapNative_noReturnValueRecipient_revertWhenZapDataPresent()
        public
        virtual
        override
    {
        setTokenTestZapNative(ZAP_NATIVE);
        setTokenTestRecipient(noReturnValueRecipient);
        vm.expectRevert(RecipientNoReturnValue.selector);
        relayWithAddress({caller: relayerB, relayer: relayerA, msgValue: ZAP_NATIVE, bridgeTx: tokenTx});
    }

    function test_relay_eth_noReturnValueRecipient_revertWhenZapDataPresent() public virtual override {
        setEthTestRecipient(noReturnValueRecipient);
        vm.expectRevert(RecipientNoReturnValue.selector);
        relay({caller: relayerB, msgValue: ethParams.destAmount, bridgeTx: ethTx});
    }

    function test_relay_eth_withRelayerAddress_noReturnValueRecipient_revertWhenZapDataPresent()
        public
        virtual
        override
    {
        setEthTestRecipient(noReturnValueRecipient);
        vm.expectRevert(RecipientNoReturnValue.selector);
        relayWithAddress({caller: relayerA, relayer: relayerB, msgValue: ethParams.destAmount, bridgeTx: ethTx});
    }

    // ═════════════════════════════════════════════ RECIPIENT REVERTS ═════════════════════════════════════════════════

    function test_relay_token_revert_recipientReverts() public {
        mockRecipientRevert(tokenTx);
        vm.expectRevert(REVERT_MSG);
        relay({caller: relayerA, msgValue: 0, bridgeTx: tokenTx});
    }

    function test_relay_token_withRelayerAddress_revert_recipientReverts() public {
        mockRecipientRevert(tokenTx);
        vm.expectRevert(REVERT_MSG);
        relayWithAddress({caller: relayerB, relayer: relayerA, msgValue: 0, bridgeTx: tokenTx});
    }

    function test_relay_token_withZapNative_revert_recipientReverts() public {
        setTokenTestZapNative(ZAP_NATIVE);
        mockRecipientRevert(tokenTx);
        vm.expectRevert(REVERT_MSG);
        relay({caller: relayerA, msgValue: ZAP_NATIVE, bridgeTx: tokenTx});
    }

    function test_relay_token_withRelayerAddressZapNative_revert_recipientReverts() public {
        setTokenTestZapNative(ZAP_NATIVE);
        mockRecipientRevert(tokenTx);
        vm.expectRevert(REVERT_MSG);
        relayWithAddress({caller: relayerB, relayer: relayerA, msgValue: ZAP_NATIVE, bridgeTx: tokenTx});
    }

    function test_relay_eth_revert_recipientReverts() public {
        mockRecipientRevert(ethTx);
        vm.expectRevert(REVERT_MSG);
        relay({caller: relayerB, msgValue: ethParams.destAmount, bridgeTx: ethTx});
    }

    function test_relay_eth_withRelayerAddress_revert_recipientReverts() public {
        mockRecipientRevert(ethTx);
        vm.expectRevert(REVERT_MSG);
        relayWithAddress({caller: relayerA, relayer: relayerB, msgValue: ethParams.destAmount, bridgeTx: ethTx});
    }

    function test_relay_eth_noZapData_revert_recipientReverts() public {
        setEthTestZapData("");
        vm.mockCallRevert({callee: userB, data: "", revertData: bytes(REVERT_MSG)});
        // Note: OZ library doesn't bubble the revert message for just sending ETH
        // (as opposed to doing an external hook call). Therefore we expect a generic library revert.
        vm.expectRevert(Address.FailedInnerCall.selector);
        relay({caller: relayerB, msgValue: ethParams.destAmount, bridgeTx: ethTx});
    }

    function test_relay_eth_withRelayerAddress_noZapData_revert_recipientReverts() public {
        setEthTestZapData("");
        vm.mockCallRevert({callee: userB, data: "", revertData: bytes(REVERT_MSG)});
        // Note: OZ library doesn't bubble the revert message for just sending ETH
        // (as opposed to doing an external hook call). Therefore we expect a generic library revert.
        vm.expectRevert(Address.FailedInnerCall.selector);
        relayWithAddress({caller: relayerA, relayer: relayerB, msgValue: ethParams.destAmount, bridgeTx: ethTx});
    }
}
