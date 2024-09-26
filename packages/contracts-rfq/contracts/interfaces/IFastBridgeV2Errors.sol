// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

interface IFastBridgeV2Errors {
    error AmountIncorrect();
    error ChainIncorrect();
    error MsgValueIncorrect();
    error SenderIncorrect();
    error StatusIncorrect();
    error ZeroAddress();

    error DeadlineExceeded();
    error DeadlineNotExceeded();
    error DeadlineTooShort();
    error DisputePeriodNotPassed();
    error DisputePeriodPassed();

    error TransactionRelayed();
}
