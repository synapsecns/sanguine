// SPDX-License-Identifier: MIT
pragma solidity ^0.8.4;

interface IZapRouterV1Errors {
    error ZapRouterV1__AmountInsufficient();
    error ZapRouterV1__DeadlineExceeded();
    error ZapRouterV1__MsgValueIncorrect();
    error ZapRouterV1__NoZapsProvided();
}
