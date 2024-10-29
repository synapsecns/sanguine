// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import {TokenZap} from "../../contracts/zaps/TokenZap.sol";
import {VaultManyArguments} from "../mocks/VaultManyArguments.sol";
import {MockERC20} from "../MockERC20.sol";

import {Test} from "forge-std/Test.sol";

// solhint-disable func-name-mixedcase, ordering
contract TokenZapTest is Test {
    uint256 internal constant AMOUNT = 0.987 ether;

    TokenZap internal tokenZap;
    VaultManyArguments internal vault;
    MockERC20 internal erc20;

    address internal user;
    address internal nativeGasToken = 0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE;

    function setUp() public {
        tokenZap = new TokenZap();
        vault = new VaultManyArguments();
        erc20 = new MockERC20("TKN", 18);

        user = makeAddr("user");

        erc20.mint(address(this), 100 * AMOUNT);
        deal(address(this), 100 * AMOUNT);
    }

    function getVaultPayload(address token, uint256 amount) public view returns (bytes memory) {
        return abi.encodeCall(vault.deposit, (token, abi.encode(token), amount, user, abi.encode(user)));
    }

    function getVaultPayloadNoAmount() public view returns (bytes memory) {
        return abi.encodeCall(vault.depositNative, (user));
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

    function checkERC20HappyPath(bytes memory zapData) public {
        // Transfer tokens to the zap contract first
        erc20.transfer(address(tokenZap), AMOUNT);
        bytes4 returnValue = tokenZap.zap(address(erc20), AMOUNT, zapData);
        assertEq(returnValue, tokenZap.zap.selector);
        // Check that the vault registered the deposit
        assertEq(vault.balanceOf(user, address(erc20)), AMOUNT);
    }

    function test_zap_erc20_placeholderZero() public {
        bytes memory zapData = getZapData(getVaultPayload(address(erc20), 0));
        checkERC20HappyPath(zapData);
    }

    function test_zap_erc20_placeholderNonZero() public {
        // Use the approximate amount of tokens as placeholder
        bytes memory zapData = getZapData(getVaultPayload(address(erc20), 1 ether));
        checkERC20HappyPath(zapData);
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

    function test_zap_erc20_revert_targetReverted() public {
        bytes memory zapData = getZapData(getVaultPayloadWithRevert());
        // Transfer tokens to the zap contract first
        erc20.transfer(address(tokenZap), AMOUNT);
        vm.expectRevert(VaultManyArguments.VaultManyArguments__SomeError.selector);
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

        vm.expectRevert(TokenZap.TokenZap__AmountIncorrect.selector);
        tokenZap.zap{value: 1 ether - 1 wei}(nativeGasToken, 1 ether, zapData);
    }

    function test_zap_native_revert_msgValueHigherThanExpected() public {
        bytes memory originalPayload = getVaultPayload(nativeGasToken, 0);
        bytes memory zapData = getZapData(originalPayload);

        vm.expectRevert(TokenZap.TokenZap__AmountIncorrect.selector);
        tokenZap.zap{value: 1 ether + 1 wei}(nativeGasToken, 1 ether, zapData);
    }

    function test_encodeZapData_revert_payloadLengthAboveMax() public {
        bytes memory tooLongPayload = new bytes(2 ** 16);
        vm.expectRevert(TokenZap.TokenZap__PayloadLengthAboveMax.selector);
        tokenZap.encodeZapData(address(vault), tooLongPayload, 0);
    }
}
