// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import {IZapRouterV1, ZapRouterV1Test} from "./ZapRouterV1.t.sol";

// solhint-disable func-name-mixedcase, ordering
contract ZapRouterV1BalanceChecksTest is ZapRouterV1Test {
    function userPerformZaps(
        uint256 msgValue,
        uint256 amountIn,
        uint256 minLastZapAmountIn,
        uint256 deadline,
        IZapRouterV1.ZapParams[] memory zapParams
    )
        public
        virtual
        override
    {
        vm.prank(user);
        router.performZapsWithBalanceChecks{value: msgValue}({
            zapRecipient: address(tokenZap),
            amountIn: amountIn,
            minLastZapAmountIn: minLastZapAmountIn,
            deadline: deadline,
            zapParams: zapParams
        });
    }

    // ═════════════════════════════════════════ SINGLE ZAP UNSPENT FUNDS ══════════════════════════════════════════════

    function test_depositERC20_exactAmount_revert_unspentERC20() public {
        IZapRouterV1.ZapParams[] memory zapParams = getDepositERC20ZapParams(AMOUNT);
        vm.expectRevert(ZapRouterV1__ZapUnspentFunds.selector);
        userPerformZaps({
            msgValue: 0,
            amountIn: AMOUNT + 1,
            minLastZapAmountIn: AMOUNT,
            deadline: block.timestamp,
            zapParams: zapParams
        });
    }

    function test_depositERC20_exactAmount_extraERC20_revert_unspentERC20() public {
        erc20.mint(address(tokenZap), EXTRA_FUNDS);
        test_depositERC20_exactAmount_revert_unspentERC20();
    }

    function test_depositNative_exactAmount_revert_unspentNative() public {
        IZapRouterV1.ZapParams[] memory zapParams = getDepositNativeZapParams(AMOUNT);
        zapParams[0].msgValue = AMOUNT + 1;
        vm.expectRevert(ZapRouterV1__ZapUnspentFunds.selector);
        userPerformZaps({
            msgValue: AMOUNT + 1,
            amountIn: AMOUNT + 1,
            minLastZapAmountIn: AMOUNT,
            deadline: block.timestamp,
            zapParams: zapParams
        });
    }

    function test_depositNative_exactAmount_extraNative_revert_unspentNative() public {
        deal(address(tokenZap), EXTRA_FUNDS);
        test_depositNative_exactAmount_revert_unspentNative();
    }

    // ═════════════════════════════════════════ DOUBLE ZAP UNSPENT FUNDS ══════════════════════════════════════════════

    function test_swapDepositERC20_exactAmounts_revert_unspentERC20() public {
        uint256 amountDeposit = AMOUNT * TOKEN_PRICE;
        IZapRouterV1.ZapParams[] memory zapParams = getSwapDepositERC20ZapParams(USE_FULL_BALANCE, amountDeposit);
        vm.expectRevert(ZapRouterV1__ZapUnspentFunds.selector);
        userPerformZaps({
            msgValue: 0,
            amountIn: AMOUNT + 1,
            minLastZapAmountIn: amountDeposit,
            deadline: block.timestamp,
            zapParams: zapParams
        });
    }

    function test_swapDepositERC20_exactAmounts_revert_unspentWETH() public {
        uint256 amountReduced = AMOUNT * TOKEN_PRICE - 1;
        IZapRouterV1.ZapParams[] memory zapParams = getSwapDepositERC20ZapParams(USE_FULL_BALANCE, amountReduced);
        vm.expectRevert(ZapRouterV1__ZapUnspentFunds.selector);
        userPerformZaps({
            msgValue: 0,
            amountIn: AMOUNT,
            minLastZapAmountIn: amountReduced,
            deadline: block.timestamp,
            zapParams: zapParams
        });
    }

    function test_swapDepositERC20_exactAmount1_extraFunds_revertWithBalanceChecks() public override {
        erc20.mint(address(tokenZap), EXTRA_FUNDS);
        weth.mint(address(tokenZap), EXTRA_FUNDS);
        uint256 amountDeposit = AMOUNT * TOKEN_PRICE;
        IZapRouterV1.ZapParams[] memory zapParams = getSwapDepositERC20ZapParams(USE_FULL_BALANCE, amountDeposit);
        vm.expectRevert(ZapRouterV1__ZapUnspentFunds.selector);
        userPerformZaps({
            msgValue: 0,
            amountIn: AMOUNT,
            minLastZapAmountIn: amountDeposit,
            deadline: block.timestamp,
            zapParams: zapParams
        });
    }

    function test_wrapDepositWETH_exactAmounts_revert_unspentNative() public {
        IZapRouterV1.ZapParams[] memory zapParams = getWrapDepositWETHZapParams(AMOUNT, AMOUNT);
        zapParams[0].msgValue = AMOUNT + 1;
        vm.expectRevert(ZapRouterV1__ZapUnspentFunds.selector);
        userPerformZaps({
            msgValue: AMOUNT + 1,
            amountIn: AMOUNT + 1,
            minLastZapAmountIn: AMOUNT,
            deadline: block.timestamp,
            zapParams: zapParams
        });
    }

    function test_wrapDepositWETH_exactAmounts_revert_unspentWETH() public {
        uint256 amountReduced = AMOUNT - 1;
        IZapRouterV1.ZapParams[] memory zapParams = getWrapDepositWETHZapParams(AMOUNT, amountReduced);
        vm.expectRevert(ZapRouterV1__ZapUnspentFunds.selector);
        userPerformZaps({
            msgValue: AMOUNT,
            amountIn: AMOUNT,
            minLastZapAmountIn: amountReduced,
            deadline: block.timestamp,
            zapParams: zapParams
        });
    }

    function test_wrapDepositWETH_exactAmount1_extraFunds_revertWithBalanceChecks() public override {
        deal(address(tokenZap), EXTRA_FUNDS);
        weth.mint(address(tokenZap), EXTRA_FUNDS);
        IZapRouterV1.ZapParams[] memory zapParams = getWrapDepositWETHZapParams(USE_FULL_BALANCE, AMOUNT);
        vm.expectRevert(ZapRouterV1__ZapUnspentFunds.selector);
        userPerformZaps({
            msgValue: AMOUNT,
            amountIn: AMOUNT,
            minLastZapAmountIn: AMOUNT,
            deadline: block.timestamp,
            zapParams: zapParams
        });
    }

    function test_unwrapDepositNative_exactAmounts_revert_unspentWETH() public {
        IZapRouterV1.ZapParams[] memory zapParams = getUnwrapDepositNativeZapParams(AMOUNT, AMOUNT);
        vm.expectRevert(ZapRouterV1__ZapUnspentFunds.selector);
        userPerformZaps({
            msgValue: 0,
            amountIn: AMOUNT + 1,
            minLastZapAmountIn: AMOUNT,
            deadline: block.timestamp,
            zapParams: zapParams
        });
    }

    function test_unwrapDepositNative_exactAmounts_revert_unspentNative() public {
        uint256 amountReduced = AMOUNT - 1;
        IZapRouterV1.ZapParams[] memory zapParams = getUnwrapDepositNativeZapParams(AMOUNT, amountReduced);
        vm.expectRevert(ZapRouterV1__ZapUnspentFunds.selector);
        userPerformZaps({
            msgValue: 0,
            amountIn: AMOUNT,
            minLastZapAmountIn: amountReduced,
            deadline: block.timestamp,
            zapParams: zapParams
        });
    }

    function test_unwrapDepositNative_exactAmount1_extraFunds_revertWithBalanceChecks() public override {
        deal(address(tokenZap), EXTRA_FUNDS);
        weth.mint(address(tokenZap), EXTRA_FUNDS);
        IZapRouterV1.ZapParams[] memory zapParams = getUnwrapDepositNativeZapParams(USE_FULL_BALANCE, AMOUNT);
        vm.expectRevert(ZapRouterV1__ZapUnspentFunds.selector);
        userPerformZaps({
            msgValue: 0,
            amountIn: AMOUNT,
            minLastZapAmountIn: AMOUNT,
            deadline: block.timestamp,
            zapParams: zapParams
        });
    }
}
