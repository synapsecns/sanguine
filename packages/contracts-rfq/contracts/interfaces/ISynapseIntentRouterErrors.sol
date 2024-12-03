// SPDX-License-Identifier: MIT
pragma solidity ^0.8.4;

interface ISynapseIntentRouterErrors {
    error SIR__AmountInsufficient();
    error SIR__DeadlineExceeded();
    error SIR__MsgValueIncorrect();
    error SIR__StepsNotProvided();
    error SIR__TokenNotContract();
    error SIR__UnspentFunds();
    error SIR__ZapIncorrectReturnValue();
    error SIR__ZapNoReturnValue();
}
