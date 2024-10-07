// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

interface IFastBridgeV2Errors {
    error AmountIncorrect();
    error CallParamsLengthAboveMax();
    error ChainIncorrect();
    error ExclusivityParamsIncorrect();
    error MsgValueIncorrect();
    error SenderIncorrect();
    error StatusIncorrect();
    error ZeroAddress();

    error RecipientIncorrectReturnValue();
    error RecipientNoReturnValue();

    error DeadlineExceeded();
    error DeadlineNotExceeded();
    error DeadlineTooShort();
    error DisputePeriodNotPassed();
    error DisputePeriodPassed();
    error ExclusivityPeriodNotPassed();

    error TransactionRelayed();
}
