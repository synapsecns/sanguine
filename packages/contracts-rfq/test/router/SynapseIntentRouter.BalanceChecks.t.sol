// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import {ISynapseIntentRouter, SynapseIntentRouterTest} from "./SynapseIntentRouter.t.sol";

// solhint-disable func-name-mixedcase, ordering
contract SynapseIntentRouterBalanceChecksTest is SynapseIntentRouterTest {
    function completeUserIntent(
        uint256 msgValue,
        uint256 amountIn,
        uint256 deadline,
        ISynapseIntentRouter.StepParams[] memory steps
    )
        public
        virtual
        override
    {
        vm.prank(user);
        router.completeIntentWithBalanceChecks{value: msgValue}({
            zapRecipient: address(tokenZap),
            amountIn: amountIn,
            deadline: deadline,
            steps: steps
        });
    }

    // ═════════════════════════════════════════ SINGLE ZAP UNSPENT FUNDS ══════════════════════════════════════════════

    function test_depositERC20_exactAmount_revert_unspentERC20() public {
        ISynapseIntentRouter.StepParams[] memory steps = getDepositERC20Steps(AMOUNT);
        vm.expectRevert(SIR__UnspentFunds.selector);
        completeUserIntent({msgValue: 0, amountIn: AMOUNT + 1, deadline: block.timestamp, steps: steps});
    }

    function test_depositERC20_exactAmount_extraFunds_revert_unspentERC20() public withExtraFunds {
        test_depositERC20_exactAmount_revert_unspentERC20();
    }

    function test_depositNative_exactAmount_revert_unspentNative() public {
        ISynapseIntentRouter.StepParams[] memory steps = getDepositNativeSteps(AMOUNT);
        steps[0].msgValue = AMOUNT + 1;
        vm.expectRevert(SIR__UnspentFunds.selector);
        completeUserIntent({msgValue: AMOUNT + 1, amountIn: AMOUNT + 1, deadline: block.timestamp, steps: steps});
    }

    function test_depositNative_exactAmount_extraFunds_revert_unspentNative() public withExtraFunds {
        test_depositNative_exactAmount_revert_unspentNative();
    }

    // ═════════════════════════════════════════ DOUBLE ZAP UNSPENT FUNDS ══════════════════════════════════════════════

    function test_swapDepositERC20_exactAmounts_revert_unspentERC20() public {
        uint256 amountDeposit = AMOUNT * TOKEN_PRICE;
        ISynapseIntentRouter.StepParams[] memory steps = getSwapDepositERC20Steps(FULL_BALANCE, amountDeposit);
        vm.expectRevert(SIR__UnspentFunds.selector);
        completeUserIntent({msgValue: 0, amountIn: AMOUNT + 1, deadline: block.timestamp, steps: steps});
    }

    function test_swapDepositERC20_exactAmounts_revert_unspentWETH() public {
        uint256 amountReduced = AMOUNT * TOKEN_PRICE - 1;
        ISynapseIntentRouter.StepParams[] memory steps = getSwapDepositERC20Steps(FULL_BALANCE, amountReduced);
        vm.expectRevert(SIR__UnspentFunds.selector);
        completeUserIntent({msgValue: 0, amountIn: AMOUNT, deadline: block.timestamp, steps: steps});
    }

    function test_swapDepositERC20_exactAmounts_extraFunds_revert_unspentERC20() public withExtraFunds {
        test_swapDepositERC20_exactAmounts_revert_unspentERC20();
    }

    function test_swapDepositERC20_exactAmounts_extraFunds_revert_unspentWETH() public withExtraFunds {
        test_swapDepositERC20_exactAmounts_revert_unspentWETH();
    }

    function test_swapDepositERC20_exactAmount1_extraFunds_revertWithBalanceChecks() public override withExtraFunds {
        uint256 amountDeposit = AMOUNT * TOKEN_PRICE;
        ISynapseIntentRouter.StepParams[] memory steps = getSwapDepositERC20Steps(FULL_BALANCE, amountDeposit);
        vm.expectRevert(SIR__UnspentFunds.selector);
        completeUserIntent({msgValue: 0, amountIn: AMOUNT, deadline: block.timestamp, steps: steps});
    }

    function test_wrapDepositWETH_exactAmounts_revert_unspentNative() public {
        ISynapseIntentRouter.StepParams[] memory steps = getWrapDepositWETHSteps(AMOUNT, AMOUNT);
        steps[0].msgValue = AMOUNT + 1;
        vm.expectRevert(SIR__UnspentFunds.selector);
        completeUserIntent({msgValue: AMOUNT + 1, amountIn: AMOUNT + 1, deadline: block.timestamp, steps: steps});
    }

    function test_wrapDepositWETH_exactAmounts_revert_unspentWETH() public {
        uint256 amountReduced = AMOUNT - 1;
        ISynapseIntentRouter.StepParams[] memory steps = getWrapDepositWETHSteps(AMOUNT, amountReduced);
        vm.expectRevert(SIR__UnspentFunds.selector);
        completeUserIntent({msgValue: AMOUNT, amountIn: AMOUNT, deadline: block.timestamp, steps: steps});
    }

    function test_wrapDepositWETH_exactAmounts_extraFunds_revert_unspentNative() public withExtraFunds {
        test_wrapDepositWETH_exactAmounts_revert_unspentNative();
    }

    function test_wrapDepositWETH_exactAmounts_extraFunds_revert_unspentWETH() public withExtraFunds {
        test_wrapDepositWETH_exactAmounts_revert_unspentWETH();
    }

    function test_wrapDepositWETH_exactAmount1_extraFunds_revertWithBalanceChecks() public override withExtraFunds {
        ISynapseIntentRouter.StepParams[] memory steps = getWrapDepositWETHSteps(FULL_BALANCE, AMOUNT);
        vm.expectRevert(SIR__UnspentFunds.selector);
        completeUserIntent({msgValue: AMOUNT, amountIn: AMOUNT, deadline: block.timestamp, steps: steps});
    }

    function test_unwrapDepositNative_exactAmounts_revert_unspentWETH() public {
        ISynapseIntentRouter.StepParams[] memory steps = getUnwrapDepositNativeSteps(AMOUNT, AMOUNT);
        vm.expectRevert(SIR__UnspentFunds.selector);
        completeUserIntent({msgValue: 0, amountIn: AMOUNT + 1, deadline: block.timestamp, steps: steps});
    }

    function test_unwrapDepositNative_exactAmounts_revert_unspentNative() public {
        uint256 amountReduced = AMOUNT - 1;
        ISynapseIntentRouter.StepParams[] memory steps = getUnwrapDepositNativeSteps(AMOUNT, amountReduced);
        vm.expectRevert(SIR__UnspentFunds.selector);
        completeUserIntent({msgValue: 0, amountIn: AMOUNT, deadline: block.timestamp, steps: steps});
    }

    function test_unwrapDepositNative_exactAmounts_extraFunds_revert_unspentWETH() public withExtraFunds {
        test_unwrapDepositNative_exactAmounts_revert_unspentWETH();
    }

    function test_unwrapDepositNative_exactAmounts_extraFunds_revert_unspentNative() public withExtraFunds {
        test_unwrapDepositNative_exactAmounts_revert_unspentNative();
    }

    function test_unwrapDepositNative_exactAmount1_extraFunds_revertWithBalanceChecks()
        public
        override
        withExtraFunds
    {
        ISynapseIntentRouter.StepParams[] memory steps = getUnwrapDepositNativeSteps(FULL_BALANCE, AMOUNT);
        vm.expectRevert(SIR__UnspentFunds.selector);
        completeUserIntent({msgValue: 0, amountIn: AMOUNT, deadline: block.timestamp, steps: steps});
    }

    function test_swapUnwrapForwardNative_exactAmounts_revert_unspentERC20() public {
        uint256 amountSwap = AMOUNT / TOKEN_PRICE;
        ISynapseIntentRouter.StepParams[] memory steps = getSwapUnwrapForwardNativeSteps(AMOUNT, amountSwap);
        vm.expectRevert(SIR__UnspentFunds.selector);
        completeUserIntent({msgValue: 0, amountIn: AMOUNT + 1, deadline: block.timestamp, steps: steps});
    }

    function test_swapUnwrapForwardNative_exactAmounts_revert_unspentWETH() public {
        uint256 amountReduced = AMOUNT / TOKEN_PRICE - 1;
        ISynapseIntentRouter.StepParams[] memory steps = getSwapUnwrapForwardNativeSteps({
            amountSwap: AMOUNT,
            amountUnwrap: amountReduced,
            minFinalBalance: amountReduced
        });
        vm.expectRevert(SIR__UnspentFunds.selector);
        completeUserIntent({msgValue: 0, amountIn: AMOUNT, deadline: block.timestamp, steps: steps});
    }

    function test_swapUnwrapForwardNative_exactAmounts_extraFunds_revert_unspentERC20() public withExtraFunds {
        test_swapUnwrapForwardNative_exactAmounts_revert_unspentERC20();
    }

    function test_swapUnwrapForwardNative_exactAmounts_extraFunds_revert_unspentWETH() public withExtraFunds {
        test_swapUnwrapForwardNative_exactAmounts_revert_unspentWETH();
    }

    function test_swapUnwrapForwardNative_exactAmount1_extraFunds_revertWithBalanceChecks()
        public
        override
        withExtraFunds
    {
        uint256 amountSwap = AMOUNT / TOKEN_PRICE;
        ISynapseIntentRouter.StepParams[] memory steps = getSwapUnwrapForwardNativeSteps(FULL_BALANCE, amountSwap);
        vm.expectRevert(SIR__UnspentFunds.selector);
        completeUserIntent({msgValue: 0, amountIn: AMOUNT, deadline: block.timestamp, steps: steps});
    }
}
