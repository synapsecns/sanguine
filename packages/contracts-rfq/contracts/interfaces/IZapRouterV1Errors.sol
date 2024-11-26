// SPDX-License-Identifier: MIT
pragma solidity ^0.8.4;

interface IZapRouterV1Errors {
    error ZapRouterV1__AmountInsufficient();
    error ZapRouterV1__DeadlineExceeded();
    error ZapRouterV1__MsgValueIncorrect();
    error ZapRouterV1__NoZapsProvided();
    error ZapRouterV1__TokenNotContract();
    error ZapRouterV1__ZapIncorrectReturnValue();
    error ZapRouterV1__ZapNoReturnValue();
}
