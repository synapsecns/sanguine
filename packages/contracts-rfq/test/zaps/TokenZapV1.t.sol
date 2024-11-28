// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import {ZapDataV1} from "../../contracts/libs/ZapDataV1.sol";
import {TokenZapV1} from "../../contracts/zaps/TokenZapV1.sol";

import {MockERC20} from "../MockERC20.sol";
import {NonPayableRecipient} from "../mocks/NonPayableRecipient.sol";
import {RecipientMock} from "../mocks/RecipientMock.sol";
import {VaultManyArguments} from "../mocks/VaultManyArguments.sol";
import {WETHMock} from "../mocks/WETHMock.sol";

import {Address} from "@openzeppelin/contracts/utils/Address.sol";
import {Test} from "forge-std/Test.sol";

// solhint-disable func-name-mixedcase, ordering
contract TokenZapV1Test is Test {
    uint256 internal constant AMOUNT = 0.987 ether;

    TokenZapV1 internal tokenZap;
    VaultManyArguments internal vault;
    MockERC20 internal erc20;
    WETHMock internal weth;
    address internal payableMock;
    address internal nonPayableMock;

    address internal user;
    address internal nativeGasToken = 0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE;

    function setUp() public {
        tokenZap = new TokenZapV1();
        vault = new VaultManyArguments();
        erc20 = new MockERC20("TKN", 18);
        weth = new WETHMock();
        payableMock = address(new RecipientMock());
        nonPayableMock = address(new NonPayableRecipient());

        user = makeAddr("user");

        erc20.mint(address(this), 100 * AMOUNT);
        deal(address(this), 200 * AMOUNT);
        weth.deposit{value: 100 * AMOUNT}();
    }

    function getVaultPayload(address token, uint256 amount) public view returns (bytes memory) {
        return abi.encodeCall(vault.deposit, (token, abi.encode(token), amount, user, abi.encode(user)));
    }

    function getVaultPayloadNoAmount() public view returns (bytes memory) {
        return abi.encodeCall(vault.depositNoAmount, (user));
    }

    function getVaultPayloadWithRevert() public view returns (bytes memory) {
        return abi.encodeCall(vault.depositWithRevert, ());
    }

    function getZapData(bytes memory originalPayload) public view returns (bytes memory) {
        // Amount is the third argument of the deposit function
        return tokenZap.encodeZapData(address(vault), originalPayload, 4 + 32 * 2);
    }

    function getZapDataNoAmount(bytes memory originalPayload) public view returns (bytes memory) {
        return tokenZap.encodeZapData(address(vault), originalPayload, originalPayload.length);
    }

    function checkERC20HappyPath(bytes memory zapData, uint256 msgValue) public {
        // Transfer tokens to the zap contract first
        erc20.transfer(address(tokenZap), AMOUNT);
        bytes4 returnValue = tokenZap.zap{value: msgValue}(address(erc20), AMOUNT, zapData);
        assertEq(returnValue, tokenZap.zap.selector);
        // Check that the vault registered the deposit
        assertEq(vault.balanceOf(user, address(erc20)), AMOUNT);
    }

    function test_zap_erc20_placeholderZero() public {
        bytes memory zapData = getZapData(getVaultPayload(address(erc20), 0));
        checkERC20HappyPath(zapData, 0);
    }

    function test_zap_erc20_placeholderNonZero() public {
        // Use the approximate amount of tokens as placeholder
        bytes memory zapData = getZapData(getVaultPayload(address(erc20), 1 ether));
        checkERC20HappyPath(zapData, 0);
    }

    function test_zap_erc20_placeholderZero_withMsgValue() public {
        bytes memory zapData = getZapData(getVaultPayload(address(erc20), 0));
        checkERC20HappyPath(zapData, 123_456);
        // Should forward the msg.value to the vault
        assertEq(address(vault).balance, 123_456);
    }

    function test_zap_erc20_placeholderNonZero_withMsgValue() public {
        bytes memory zapData = getZapData(getVaultPayload(address(erc20), 1 ether));
        checkERC20HappyPath(zapData, 123_456);
        // Should forward the msg.value to the vault
        assertEq(address(vault).balance, 123_456);
    }

    function test_zap_erc20_placeholderZero_extraTokens() public {
        // Mint some extra tokens to the zap contract
        erc20.mint(address(tokenZap), AMOUNT);
        // Should not affect the zap
        test_zap_erc20_placeholderZero();
    }

    function test_zap_erc20_placeholderNonZero_extraTokens() public {
        // Mint some extra tokens to the zap contract
        erc20.mint(address(tokenZap), AMOUNT);
        // Should not affect the zap
        test_zap_erc20_placeholderNonZero();
    }

    function checkNativeHappyPath(bytes memory zapData) public {
        bytes4 returnValue = tokenZap.zap{value: AMOUNT}(nativeGasToken, AMOUNT, zapData);
        assertEq(returnValue, tokenZap.zap.selector);
        // Check that the vault registered the deposit
        assertEq(vault.balanceOf(user, nativeGasToken), AMOUNT);
    }

    function test_zap_native_placeholderZero() public {
        bytes memory zapData = getZapData(getVaultPayload(nativeGasToken, 0));
        checkNativeHappyPath(zapData);
    }

    function test_zap_native_placeholderNonZero() public {
        // Use the approximate amount of tokens as placeholder
        bytes memory zapData = getZapData(getVaultPayload(nativeGasToken, 1 ether));
        checkNativeHappyPath(zapData);
    }

    function test_zap_native_noAmount() public {
        bytes memory zapData = getZapDataNoAmount(getVaultPayloadNoAmount());
        checkNativeHappyPath(zapData);
    }

    function test_zap_native_placeholderZero_extraNative() public {
        // Mint some extra native tokens to the zap contract
        deal(address(tokenZap), AMOUNT);
        // Should not affect the zap
        test_zap_native_placeholderZero();
    }

    function test_zap_native_placeholderNonZero_extraNative() public {
        // Mint some extra native tokens to the zap contract
        deal(address(tokenZap), AMOUNT);
        // Should not affect the zap
        test_zap_native_placeholderNonZero();
    }

    function test_zap_native_noAmount_extraNative() public {
        // Mint some extra native tokens to the zap contract
        deal(address(tokenZap), AMOUNT);
        // Should not affect the zap
        test_zap_native_noAmount();
    }

    /// @notice Should be able to use amount lower than msg.value.
    function test_zap_native_msgValueHigherThanAmount() public {
        bytes memory zapData = getZapData(getVaultPayload(nativeGasToken, 1 ether));
        bytes4 returnValue = tokenZap.zap{value: AMOUNT + 1 wei}(nativeGasToken, AMOUNT, zapData);
        assertEq(returnValue, tokenZap.zap.selector);
        // Check that the vault registered the deposit
        assertEq(vault.balanceOf(user, nativeGasToken), AMOUNT);
        // Note: the extra funds are left in the contract
        assertEq(address(tokenZap).balance, 1 wei);
    }

    /// @notice Should be able to utilize both msg.value and existing native balance.
    function test_zap_native_msgValueLowerThanAmount_extraNative() public {
        deal(address(tokenZap), 1337);
        bytes memory zapData = getZapData(getVaultPayload(nativeGasToken, 1 ether));
        bytes4 returnValue = tokenZap.zap{value: AMOUNT - 1337}(nativeGasToken, AMOUNT, zapData);
        assertEq(returnValue, tokenZap.zap.selector);
        // Check that the vault registered the deposit
        assertEq(vault.balanceOf(user, nativeGasToken), AMOUNT);
    }

    // ═════════════════════════════════════════════════ MULTIHOPS ═════════════════════════════════════════════════════

    function getZapDataWithdraw(uint256 amount) public view returns (bytes memory) {
        return tokenZap.encodeZapData(address(weth), abi.encodeCall(WETHMock.withdraw, (amount)), 4);
    }

    function test_zap_withdraw_depositNative_placeholderZero() public {
        bytes memory zapDataWithdraw = getZapDataWithdraw(0);
        bytes memory zapDataDeposit = getZapDataNoAmount(getVaultPayloadNoAmount());
        weth.transfer(address(tokenZap), AMOUNT);
        // Do two Zaps in a row
        bytes4 returnValue = tokenZap.zap(address(weth), AMOUNT, zapDataWithdraw);
        assertEq(returnValue, tokenZap.zap.selector);
        returnValue = tokenZap.zap(nativeGasToken, AMOUNT, zapDataDeposit);
        assertEq(returnValue, tokenZap.zap.selector);
        // Check that the vault registered the deposit
        assertEq(vault.balanceOf(user, nativeGasToken), AMOUNT);
    }

    function test_zap_withdraw_depositNative_placeholderNonZero() public {
        // Use the approximate amount of tokens as placeholder
        bytes memory zapDataWithdraw = getZapDataWithdraw(1 ether);
        bytes memory zapDataDeposit = getZapDataNoAmount(getVaultPayloadNoAmount());
        weth.transfer(address(tokenZap), AMOUNT);
        // Do two Zaps in a row
        bytes4 returnValue = tokenZap.zap(address(weth), AMOUNT, zapDataWithdraw);
        assertEq(returnValue, tokenZap.zap.selector);
        returnValue = tokenZap.zap(nativeGasToken, AMOUNT, zapDataDeposit);
        assertEq(returnValue, tokenZap.zap.selector);
        // Check that the vault registered the deposit
        assertEq(vault.balanceOf(user, nativeGasToken), AMOUNT);
    }

    function test_zap_withdraw_depositNative_placeholderZero_extraTokens() public {
        // Transfer some extra tokens to the zap contract
        weth.transfer(address(tokenZap), AMOUNT);
        // Should not affect the zap
        test_zap_withdraw_depositNative_placeholderZero();
    }

    function test_zap_withdraw_depositNative_placeholderZero_extraNative() public {
        // Transfer some extra native tokens to the zap contract
        deal(address(tokenZap), AMOUNT);
        // Should not affect the zap
        test_zap_withdraw_depositNative_placeholderZero();
    }

    function test_zap_withdraw_depositNative_placeholderNonZero_extraTokens() public {
        // Transfer some extra tokens to the zap contract
        weth.transfer(address(tokenZap), AMOUNT);
        // Should not affect the zap
        test_zap_withdraw_depositNative_placeholderNonZero();
    }

    function test_zap_withdraw_depositNative_placeholderNonZero_extraNative() public {
        // Transfer some extra native tokens to the zap contract
        deal(address(tokenZap), AMOUNT);
        // Should not affect the zap
        test_zap_withdraw_depositNative_placeholderNonZero();
    }

    function test_zap_withdraw_transferNativeEOA_placeholderZero() public {
        bytes memory zapDataWithdraw = getZapDataWithdraw(0);
        bytes memory zapDataTransfer = tokenZap.encodeZapData({target: user, payload: "", amountPosition: 0});
        weth.transfer(address(tokenZap), AMOUNT);
        // Do two Zaps in a row
        bytes4 returnValue = tokenZap.zap(address(weth), AMOUNT, zapDataWithdraw);
        assertEq(returnValue, tokenZap.zap.selector);
        returnValue = tokenZap.zap(nativeGasToken, AMOUNT, zapDataTransfer);
        assertEq(returnValue, tokenZap.zap.selector);
        // Check that the user received the native tokens
        assertEq(user.balance, AMOUNT);
    }

    function test_zap_withdraw_transferNativeEOA_placeholderNonZero() public {
        // Use the approximate amount of tokens as placeholder
        bytes memory zapDataWithdraw = getZapDataWithdraw(1 ether);
        bytes memory zapDataTransfer = tokenZap.encodeZapData({target: user, payload: "", amountPosition: 0});
        weth.transfer(address(tokenZap), AMOUNT);
        // Do two Zaps in a row
        bytes4 returnValue = tokenZap.zap(address(weth), AMOUNT, zapDataWithdraw);
        assertEq(returnValue, tokenZap.zap.selector);
        returnValue = tokenZap.zap(nativeGasToken, AMOUNT, zapDataTransfer);
        assertEq(returnValue, tokenZap.zap.selector);
        // Check that the user received the native tokens
        assertEq(user.balance, AMOUNT);
    }

    function test_zap_withdraw_transferNativeEOA_placeholderZero_extraTokens() public {
        // Transfer some extra tokens to the zap contract
        weth.transfer(address(tokenZap), AMOUNT);
        // Should not affect the zap
        test_zap_withdraw_transferNativeEOA_placeholderZero();
    }

    function test_zap_withdraw_transferNativeEOA_placeholderZero_extraNative() public {
        // Transfer some extra native tokens to the zap contract
        deal(address(tokenZap), AMOUNT);
        // Should not affect the zap
        test_zap_withdraw_transferNativeEOA_placeholderZero();
    }

    function test_zap_withdraw_transferNativeEOA_placeholderNonZero_extraTokens() public {
        // Transfer some extra tokens to the zap contract
        weth.transfer(address(tokenZap), AMOUNT);
        // Should not affect the zap
        test_zap_withdraw_transferNativeEOA_placeholderNonZero();
    }

    function test_zap_withdraw_transferNativeEOA_placeholderNonZero_extraNative() public {
        // Transfer some extra native tokens to the zap contract
        deal(address(tokenZap), AMOUNT);
        // Should not affect the zap
        test_zap_withdraw_transferNativeEOA_placeholderNonZero();
    }

    function test_zap_withdraw_transferNativeContract_placeholderZero() public {
        bytes memory zapDataWithdraw = getZapDataWithdraw(0);
        bytes memory zapDataTransfer = tokenZap.encodeZapData({target: payableMock, payload: "", amountPosition: 0});
        weth.transfer(address(tokenZap), AMOUNT);
        // Do two Zaps in a row
        bytes4 returnValue = tokenZap.zap(address(weth), AMOUNT, zapDataWithdraw);
        assertEq(returnValue, tokenZap.zap.selector);
        returnValue = tokenZap.zap(nativeGasToken, AMOUNT, zapDataTransfer);
        assertEq(returnValue, tokenZap.zap.selector);
        // Check that the contract received the native tokens
        assertEq(payableMock.balance, AMOUNT);
    }

    function test_zap_withdraw_transferNativeContract_placeholderNonZero() public {
        // Use the approximate amount of tokens as placeholder
        bytes memory zapDataWithdraw = getZapDataWithdraw(1 ether);
        bytes memory zapDataTransfer = tokenZap.encodeZapData({target: payableMock, payload: "", amountPosition: 0});
        weth.transfer(address(tokenZap), AMOUNT);
        // Do two Zaps in a row
        bytes4 returnValue = tokenZap.zap(address(weth), AMOUNT, zapDataWithdraw);
        assertEq(returnValue, tokenZap.zap.selector);
        returnValue = tokenZap.zap(nativeGasToken, AMOUNT, zapDataTransfer);
        assertEq(returnValue, tokenZap.zap.selector);
        // Check that the contract received the native tokens
        assertEq(payableMock.balance, AMOUNT);
    }

    function test_zap_withdraw_transferNativeContract_placeholderZero_extraTokens() public {
        // Transfer some extra tokens to the zap contract
        weth.transfer(address(tokenZap), AMOUNT);
        // Should not affect the zap
        test_zap_withdraw_transferNativeContract_placeholderZero();
    }

    function test_zap_withdraw_transferNativeContract_placeholderZero_extraNative() public {
        // Transfer some extra native tokens to the zap contract
        deal(address(tokenZap), AMOUNT);
        // Should not affect the zap
        test_zap_withdraw_transferNativeContract_placeholderZero();
    }

    function test_zap_withdraw_transferNativeContract_placeholderNonZero_extraTokens() public {
        // Transfer some extra tokens to the zap contract
        weth.transfer(address(tokenZap), AMOUNT);
        // Should not affect the zap
        test_zap_withdraw_transferNativeContract_placeholderNonZero();
    }

    function test_zap_withdraw_transferNativeContract_placeholderNonZero_extraNative() public {
        // Transfer some extra native tokens to the zap contract
        deal(address(tokenZap), AMOUNT);
        // Should not affect the zap
        test_zap_withdraw_transferNativeContract_placeholderNonZero();
    }

    // ═════════════════════════════════════════════════ ENCODING ══════════════════════════════════════════════════════

    function test_encodeZapData_roundtrip(address token, uint256 placeholderAmount, uint256 amount) public view {
        bytes memory originalPayload = getVaultPayload(token, placeholderAmount);
        bytes memory expectedPayload = getVaultPayload(token, amount);

        bytes memory zapData = getZapData(originalPayload);
        (address target, bytes memory payload) = tokenZap.decodeZapData(zapData, amount);

        assertEq(target, address(vault));
        assertEq(payload, expectedPayload);
    }

    function test_encodeZapData_roundtripNoAmount(uint256 amountPosition) public view {
        bytes memory payload = getVaultPayloadNoAmount();
        // Any value >= payload.length could be used to signal that the amount is not an argument of the target function
        amountPosition = bound(amountPosition, payload.length, type(uint256).max);

        bytes memory zapData = tokenZap.encodeZapData(address(vault), payload, amountPosition);
        (address target, bytes memory decodedPayload) = tokenZap.decodeZapData(zapData, 0);
        assertEq(target, address(vault));
        assertEq(decodedPayload, payload);
    }

    // ══════════════════════════════════════════════════ REVERTS ══════════════════════════════════════════════════════

    function getZeroTargetZapData(bytes memory payload, uint16 amountPosition) public pure returns (bytes memory) {
        // Encode manually as the library checks for zero address
        return abi.encodePacked(ZapDataV1.VERSION, amountPosition, address(0), payload);
    }

    function test_zap_erc20_revert_notEnoughTokens() public {
        bytes memory zapData = getZapData(getVaultPayload(address(erc20), 0));
        // Transfer tokens to the zap contract first, but not enough
        erc20.transfer(address(tokenZap), AMOUNT - 1);
        vm.expectRevert();
        tokenZap.zap(address(erc20), AMOUNT, zapData);
    }

    function test_zap_erc20_revert_targetReverted() public {
        bytes memory zapData = getZapData(getVaultPayloadWithRevert());
        // Transfer tokens to the zap contract first
        erc20.transfer(address(tokenZap), AMOUNT);
        vm.expectRevert(VaultManyArguments.VaultManyArguments__SomeError.selector);
        tokenZap.zap(address(erc20), AMOUNT, zapData);
    }

    function test_zap_erc20_revert_targetZeroAddress_emptyPayload() public {
        bytes memory zapData = getZeroTargetZapData({payload: "", amountPosition: 0});
        // Transfer tokens to the zap contract first
        erc20.transfer(address(tokenZap), AMOUNT);
        vm.expectRevert(TokenZapV1.TokenZapV1__TargetZeroAddress.selector);
        tokenZap.zap(address(erc20), AMOUNT, zapData);
    }

    function test_zap_erc20_revert_targetZeroAddress_nonEmptyPayload() public {
        bytes memory zapData =
            getZeroTargetZapData({payload: getVaultPayload(address(erc20), 0), amountPosition: 4 + 32 * 2});
        // Transfer tokens to the zap contract first
        erc20.transfer(address(tokenZap), AMOUNT);
        vm.expectRevert(TokenZapV1.TokenZapV1__TargetZeroAddress.selector);
        tokenZap.zap(address(erc20), AMOUNT, zapData);
    }

    function test_zap_erc20_revert_targetEOA_nonEmptyPayload() public {
        bytes memory zapData = tokenZap.encodeZapData({
            target: user,
            payload: getVaultPayload(address(erc20), 0),
            amountPosition: 4 + 32 * 2
        });
        // Transfer tokens to the zap contract first
        erc20.transfer(address(tokenZap), AMOUNT);
        vm.expectRevert(abi.encodeWithSelector(Address.AddressEmptyCode.selector, user));
        tokenZap.zap(address(erc20), AMOUNT, zapData);
    }

    function test_zap_erc20_revert_targetEOA_emptyPayload() public {
        bytes memory zapData = tokenZap.encodeZapData({target: user, payload: "", amountPosition: 0});
        // Transfer tokens to the zap contract first
        erc20.transfer(address(tokenZap), AMOUNT);
        vm.expectRevert(abi.encodeWithSelector(Address.AddressEmptyCode.selector, user));
        tokenZap.zap(address(erc20), AMOUNT, zapData);
    }

    function test_zap_native_revert_targetReverted() public {
        bytes memory zapData = getZapData(getVaultPayloadWithRevert());
        vm.expectRevert(VaultManyArguments.VaultManyArguments__SomeError.selector);
        tokenZap.zap{value: AMOUNT}(nativeGasToken, AMOUNT, zapData);
    }

    function test_zap_native_revert_msgValueLowerThanExpected() public {
        bytes memory originalPayload = getVaultPayload(nativeGasToken, 0);
        bytes memory zapData = getZapData(originalPayload);

        vm.expectRevert(abi.encodeWithSelector(Address.AddressInsufficientBalance.selector, tokenZap));
        tokenZap.zap{value: 1 ether - 1 wei}(nativeGasToken, 1 ether, zapData);
    }

    function test_zap_withdraw_transferNative_revert_targetReverted() public {
        bytes memory zapDataWithdraw = getZapDataWithdraw(0);
        bytes memory zapDataTransfer = tokenZap.encodeZapData({target: nonPayableMock, payload: "", amountPosition: 0});
        weth.transfer(address(tokenZap), AMOUNT);
        tokenZap.zap(address(weth), AMOUNT, zapDataWithdraw);
        vm.expectRevert(Address.FailedInnerCall.selector);
        tokenZap.zap(address(weth), AMOUNT, zapDataTransfer);
    }

    function test_zap_withdraw_transferNative_revert_targetZeroAddress_emptyPayload() public {
        bytes memory zapDataWithdraw = getZapDataWithdraw(0);
        bytes memory zapDataTransfer = getZeroTargetZapData({payload: "", amountPosition: 0});
        weth.transfer(address(tokenZap), AMOUNT);
        tokenZap.zap(address(weth), AMOUNT, zapDataWithdraw);
        vm.expectRevert(TokenZapV1.TokenZapV1__TargetZeroAddress.selector);
        tokenZap.zap(address(weth), AMOUNT, zapDataTransfer);
    }

    function test_zap_withdraw_transferNative_revert_targetZeroAddress_nonEmptyPayload() public {
        bytes memory zapDataWithdraw = getZapDataWithdraw(0);
        bytes memory payload = getVaultPayloadNoAmount();
        bytes memory zapDataTransfer = getZeroTargetZapData({payload: payload, amountPosition: uint16(payload.length)});
        weth.transfer(address(tokenZap), AMOUNT);
        tokenZap.zap(address(weth), AMOUNT, zapDataWithdraw);
        vm.expectRevert(TokenZapV1.TokenZapV1__TargetZeroAddress.selector);
        tokenZap.zap(address(weth), AMOUNT, zapDataTransfer);
    }

    function test_zap_withdraw_transferNative_revert_targetEOA_nonEmptyPayload() public {
        bytes memory zapDataWithdraw = getZapDataWithdraw(0);
        bytes memory zapDataTransfer =
            tokenZap.encodeZapData({target: user, payload: getVaultPayloadNoAmount(), amountPosition: 0});
        weth.transfer(address(tokenZap), AMOUNT);
        tokenZap.zap(address(weth), AMOUNT, zapDataWithdraw);
        vm.expectRevert(abi.encodeWithSelector(Address.AddressEmptyCode.selector, user));
        tokenZap.zap(address(weth), AMOUNT, zapDataTransfer);
    }

    function test_encodeZapData_revert_payloadLengthAboveMax() public {
        bytes memory tooLongPayload = new bytes(2 ** 16);
        vm.expectRevert(TokenZapV1.TokenZapV1__PayloadLengthAboveMax.selector);
        tokenZap.encodeZapData(address(vault), tooLongPayload, 0);
    }

    function test_encodeZapData_revert_targetZeroAddress() public {
        bytes memory payload = getVaultPayloadNoAmount();

        vm.expectRevert(ZapDataV1.ZapDataV1__TargetZeroAddress.selector);
        tokenZap.encodeZapData(address(0), payload, payload.length);
    }
}
