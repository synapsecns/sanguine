// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

interface IFastBridgeV2Errors {
    error AmountIncorrect();
    error ChainIncorrect();
    error ExclusivityParamsIncorrect();
    error MsgValueIncorrect();
    error SenderIncorrect();
    error StatusIncorrect();
    error TokenNotContract();
    error ZapDataLengthAboveMax();
    error ZapNativeNotSupported();
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
