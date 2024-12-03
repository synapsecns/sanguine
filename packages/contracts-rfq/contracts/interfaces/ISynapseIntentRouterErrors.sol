// SPDX-License-Identifier: MIT
pragma solidity ^0.8.4;

interface ISynapseIntentRouterErrors {
    error SIR__AmountInsufficient();
    error SIR__DeadlineExceeded();
    error SIR__MsgValueIncorrect();
    error SIR__StepsNotProvided();
}
