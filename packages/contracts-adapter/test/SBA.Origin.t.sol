// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import {ISynapseBridgeAdapter, SynapseBridgeAdapter, SynapseBridgeAdapterTest} from "./SBA.t.sol";

import {SynapseBridgeMock} from "./mocks/SynapseBridgeMock.sol";
import {ERC20, TestToken} from "./mocks/TestToken.sol";

import {
    ILayerZeroEndpointV2,
    MessagingFee,
    MessagingParams,
    MessagingReceipt
} from "@layerzerolabs/lz-evm-protocol-v2/contracts/interfaces/ILayerZeroEndpointV2.sol";

// solhint-disable check-send-result, func-name-mixedcase, ordering
contract SynapseBridgeAdapterOriginTest is SynapseBridgeAdapterTest {
    uint64 internal constant MIN_GAS_LIMIT = 100_000;

    address internal bridge;
    TestToken internal token;

    address internal user = makeAddr("User");
    address internal recipient = makeAddr("Recipient");
    bytes31 internal symbol = "SYMBOL";
    uint256 internal initialBalance = 1 ether;
    uint256 internal amount = 0.123456 ether;
    uint64 internal gasLimit = 123_321;
    uint256 internal nativeFee = 123_456_789 wei;

    // Based on https://docs.layerzero.network/v2/developers/evm/oapp/overview#message-execution-options
    // Replacing 60_000 with 123_321
    bytes internal expectedOptions = hex"0003010011010000000000000000000000000001e1b9";

    bytes internal expectedBridgeMessage;

    modifier withBridgeSet() {
        adapter.setBridge(bridge);
        _;
    }

    modifier withMintTokenAdded() {
        adapter.addToken(address(token), ISynapseBridgeAdapter.TokenType.MintBurn, symbol);
        _;
    }

    modifier withWithdrawTokenAdded() {
        adapter.addToken(address(token), ISynapseBridgeAdapter.TokenType.WithdrawDeposit, symbol);
        _;
    }

    function afterAdapterDeployed() internal virtual override {
        adapter.setPeer(DST_EID, DEST_ADAPTER);

        bridge = address(new SynapseBridgeMock());
        token = new TestToken();
        token.mintTestTokens(address(user), initialBalance);
        deal(user, 100 ether);
        vm.prank(user);
        token.approve(address(adapter), type(uint256).max);

        expectedBridgeMessage = bridgeMessageLib.encodeBridgeMessage(recipient, symbol, amount);

        mockSendReceipt();
    }

    function deployAdapter() internal virtual override returns (SynapseBridgeAdapter) {
        return new SynapseBridgeAdapter(endpoint, address(this));
    }

    function userBridgesToken() internal {
        vm.prank({msgSender: user, txOrigin: user});
        adapter.bridgeERC20{value: nativeFee}(DST_EID, recipient, address(token), amount, gasLimit);
    }

    function mockSendReceipt() internal {
        vm.mockCall({
            callee: endpoint,
            data: abi.encodeWithSelector(ILayerZeroEndpointV2.send.selector),
            returnData: abi.encode(
                MessagingReceipt({guid: MOCK_GUID, nonce: 1, fee: MessagingFee({nativeFee: nativeFee, lzTokenFee: 0})})
            )
        });
    }

    // ═══════════════════════════════════════════ TEST: MINT-BURN TOKEN ═══════════════════════════════════════════════

    function test_bridge_mintBurn() public withBridgeSet withMintTokenAdded {
        // Expected token action: burn from user
        vm.expectCall({callee: address(token), data: abi.encodeCall(TestToken.burnFrom, (user, amount))});
        // Expected bridge message
        vm.expectCall({
            callee: endpoint,
            msgValue: nativeFee,
            data: abi.encodeCall(
                ILayerZeroEndpointV2.send,
                (
                    MessagingParams({
                        dstEid: DST_EID,
                        receiver: DEST_ADAPTER,
                        message: expectedBridgeMessage,
                        options: expectedOptions,
                        payInLzToken: false
                    }),
                    // refund to msg.sender
                    user
                )
            )
        });
        // Expected event
        expectEventTokenSent(DST_EID, recipient, address(token), amount, MOCK_GUID);
        userBridgesToken();
        // Check token balances
        assertEq(token.balanceOf(address(user)), initialBalance - amount);
        assertEq(token.balanceOf(address(adapter)), 0);
        assertEq(token.balanceOf(address(bridge)), 0);
        assertEq(token.totalSupply(), initialBalance - amount);
    }

    function test_bridge_mintBurn_tokenNotApproved() public withBridgeSet withMintTokenAdded {
        vm.prank(user);
        token.approve(address(adapter), 0);
        vm.expectRevert();
        userBridgesToken();
    }

    function test_bridge_mintBurn_revert_bridgeNotSet() public withMintTokenAdded {
        expectRevertBridgeNotSet();
        userBridgesToken();
    }

    function test_bridge_mintBurn_revert_tokenNotAdded() public withBridgeSet {
        expectRevertTokenUnknown(address(token));
        userBridgesToken();
    }

    function test_bridge_mintBurn_revert_bridgeNotSetTokenNotAdded() public {
        expectRevertBridgeNotSet();
        userBridgesToken();
    }

    function test_bridge_mintBurn_revert_eidUnknown() public withBridgeSet withMintTokenAdded {
        vm.expectRevert();
        vm.prank({msgSender: user, txOrigin: user});
        adapter.bridgeERC20{value: nativeFee}(UNKNOWN_EID, recipient, address(token), amount, gasLimit);
    }

    function test_bridge_mintBurn_revert_recipientZero() public withBridgeSet withMintTokenAdded {
        expectRevertZeroAddress();
        vm.prank({msgSender: user, txOrigin: user});
        adapter.bridgeERC20{value: nativeFee}(DST_EID, address(0), address(token), amount, gasLimit);
    }

    function test_bridge_mintBurn_revert_tokenZero() public withBridgeSet withMintTokenAdded {
        expectRevertTokenUnknown(address(0));
        vm.prank({msgSender: user, txOrigin: user});
        adapter.bridgeERC20{value: nativeFee}(DST_EID, recipient, address(0), amount, gasLimit);
    }

    function test_bridge_mintBurn_revert_amountZero() public withBridgeSet withMintTokenAdded {
        expectRevertZeroAmount();
        vm.prank({msgSender: user, txOrigin: user});
        adapter.bridgeERC20{value: nativeFee}(DST_EID, recipient, address(token), 0, gasLimit);
    }

    function test_bridge_mintBurn_revert_gasLimitBelowMinimum() public withBridgeSet withMintTokenAdded {
        expectRevertGasLimitBelowMinimum();
        vm.prank({msgSender: user, txOrigin: user});
        adapter.bridgeERC20{value: nativeFee}(DST_EID, recipient, address(token), amount, MIN_GAS_LIMIT - 1);
    }

    // ═══════════════════════════════════════ TEST: WITHDRAW-DEPOSIT TOKEN ════════════════════════════════════════════

    function test_bridge_withdrawDeposit() public withBridgeSet withWithdrawTokenAdded {
        // Expected token action: transfer from user to bridge
        vm.expectCall({
            callee: address(token),
            data: abi.encodeCall(ERC20.transferFrom, (user, address(bridge), amount))
        });
        // Expected bridge message
        vm.expectCall({
            callee: endpoint,
            msgValue: nativeFee,
            data: abi.encodeCall(
                ILayerZeroEndpointV2.send,
                (
                    MessagingParams({
                        dstEid: DST_EID,
                        receiver: DEST_ADAPTER,
                        message: expectedBridgeMessage,
                        options: expectedOptions,
                        payInLzToken: false
                    }),
                    // refund to msg.sender
                    user
                )
            )
        });
        // Expected event
        expectEventTokenSent(DST_EID, recipient, address(token), amount, MOCK_GUID);
        userBridgesToken();
        // Check token balances
        assertEq(token.balanceOf(address(user)), initialBalance - amount);
        assertEq(token.balanceOf(address(adapter)), 0);
        assertEq(token.balanceOf(address(bridge)), amount);
        assertEq(token.totalSupply(), initialBalance);
    }

    function test_bridge_withdrawDeposit_tokenNotApproved() public withBridgeSet withWithdrawTokenAdded {
        vm.prank(user);
        token.approve(address(adapter), 0);
        vm.expectRevert();
        userBridgesToken();
    }

    function test_bridge_withdrawDeposit_revert_bridgeNotSet() public withWithdrawTokenAdded {
        expectRevertBridgeNotSet();
        userBridgesToken();
    }

    function test_bridge_withdrawDeposit_revert_tokenNotAdded() public withBridgeSet {
        expectRevertTokenUnknown(address(token));
        userBridgesToken();
    }

    function test_bridge_withdrawDeposit_revert_bridgeNotSetTokenNotAdded() public {
        expectRevertBridgeNotSet();
        userBridgesToken();
    }

    function test_bridge_withdrawDeposit_revert_eidUnknown() public withBridgeSet withWithdrawTokenAdded {
        vm.expectRevert();
        vm.prank({msgSender: user, txOrigin: user});
        adapter.bridgeERC20{value: nativeFee}(UNKNOWN_EID, recipient, address(token), amount, gasLimit);
    }

    function test_bridge_withdrawDeposit_revert_recipientZero() public withBridgeSet withWithdrawTokenAdded {
        expectRevertZeroAddress();
        vm.prank({msgSender: user, txOrigin: user});
        adapter.bridgeERC20{value: nativeFee}(DST_EID, address(0), address(token), amount, gasLimit);
    }

    function test_bridge_withdrawDeposit_revert_tokenZero() public withBridgeSet withWithdrawTokenAdded {
        expectRevertTokenUnknown(address(0));
        vm.prank({msgSender: user, txOrigin: user});
        adapter.bridgeERC20{value: nativeFee}(DST_EID, recipient, address(0), amount, gasLimit);
    }

    function test_bridge_withdrawDeposit_revert_amountZero() public withBridgeSet withWithdrawTokenAdded {
        expectRevertZeroAmount();
        vm.prank({msgSender: user, txOrigin: user});
        adapter.bridgeERC20{value: nativeFee}(DST_EID, recipient, address(token), 0, gasLimit);
    }

    function test_bridge_withdrawDeposit_revert_gasLimitBelowMinimum() public withBridgeSet withWithdrawTokenAdded {
        expectRevertGasLimitBelowMinimum();
        vm.prank({msgSender: user, txOrigin: user});
        adapter.bridgeERC20{value: nativeFee}(DST_EID, recipient, address(token), amount, MIN_GAS_LIMIT - 1);
    }

    // ══════════════════════════════════════════════ GET NATIVE FEE ═══════════════════════════════════════════════════

    function test_getNativeFee() public {
        bytes memory mockMessage = bridgeMessageLib.encodeBridgeMessage(address(0), 0, 0);
        vm.mockCall({
            callee: endpoint,
            data: abi.encodeCall(
                ILayerZeroEndpointV2.quote,
                (
                    MessagingParams({
                        dstEid: DST_EID,
                        receiver: DEST_ADAPTER,
                        message: mockMessage,
                        options: expectedOptions,
                        payInLzToken: false
                    }),
                    address(adapter)
                )
            ),
            returnData: abi.encode(MessagingFee({nativeFee: nativeFee, lzTokenFee: 0}))
        });
        assertEq(adapter.getNativeFee(DST_EID, gasLimit), nativeFee);
    }

    function test_getNativeFee_revert_eidUnknown() public {
        vm.expectRevert();
        adapter.getNativeFee(UNKNOWN_EID, gasLimit);
    }

    function test_getNativeFee_revert_gasLimitBelowMinimum() public {
        vm.expectRevert();
        adapter.getNativeFee(DST_EID, MIN_GAS_LIMIT - 1);
    }
}
