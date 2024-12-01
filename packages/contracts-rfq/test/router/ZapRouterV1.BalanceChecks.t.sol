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

    function test_depositERC20_revert_unspentERC20() public {
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

    function test_depositNative_revert_unspentNative() public {
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
}
