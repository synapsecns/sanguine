// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import {ISynapseIntentRouter, SynapseIntentRouterTest} from "./SynapseIntentRouter.t.sol";

// solhint-disable func-name-mixedcase, ordering
contract SynapseIntentRouterBalanceChecksTest is SynapseIntentRouterTest {
    function completeUserIntent(
        uint256 msgValue,
        uint256 amountIn,
        uint256 minLastStepAmountIn,
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
            minLastStepAmountIn: minLastStepAmountIn,
            deadline: deadline,
            steps: steps
        });
    }

    // ═════════════════════════════════════════ SINGLE ZAP UNSPENT FUNDS ══════════════════════════════════════════════

    function test_depositERC20_revert_unspentERC20() public {
        ISynapseIntentRouter.StepParams[] memory steps = getDepositERC20StepParams(AMOUNT);
        vm.expectRevert(SIR__UnspentFunds.selector);
        completeUserIntent({
            msgValue: 0,
            amountIn: AMOUNT + 1,
            minLastStepAmountIn: AMOUNT,
            deadline: block.timestamp,
            steps: steps
        });
    }

    function test_depositNative_revert_unspentNative() public {
        ISynapseIntentRouter.StepParams[] memory steps = getDepositNativeStepParams(AMOUNT);
        steps[0].msgValue = AMOUNT + 1;
        vm.expectRevert(SIR__UnspentFunds.selector);
        completeUserIntent({
            msgValue: AMOUNT + 1,
            amountIn: AMOUNT + 1,
            minLastStepAmountIn: AMOUNT,
            deadline: block.timestamp,
            steps: steps
        });
    }
}
