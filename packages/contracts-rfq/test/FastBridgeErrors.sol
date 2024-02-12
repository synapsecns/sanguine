// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

import {
    AmountIncorrect,
    ChainIncorrect,
    DeadlineExceeded,
    DeadlineNotExceeded,
    DeadlineTooShort,
    DisputePeriodNotPassed,
    DisputePeriodPassed,
    MsgValueIncorrect,
    SenderIncorrect,
    StatusIncorrect,
    TransactionRelayed,
    ZeroAddress
} from "../contracts/libs/Errors.sol";

contract FastBridgeErrors {
    function amountIncorrectSelector() internal pure virtual returns (bytes4) {
        return AmountIncorrect.selector;
    }

    function chainIncorrectSelector() internal pure virtual returns (bytes4) {
        return ChainIncorrect.selector;
    }

    function deadlineExceededSelector() internal pure virtual returns (bytes4) {
        return DeadlineExceeded.selector;
    }

    function deadlineNotExceededSelector() internal pure virtual returns (bytes4) {
        return DeadlineNotExceeded.selector;
    }

    function deadlineTooShortSelector() internal pure virtual returns (bytes4) {
        return DeadlineTooShort.selector;
    }

    function disputePeriodNotPassedSelector() internal pure virtual returns (bytes4) {
        return DisputePeriodNotPassed.selector;
    }

    function disputePeriodPassedSelector() internal pure virtual returns (bytes4) {
        return DisputePeriodPassed.selector;
    }

    function msgValueIncorrectSelector() internal pure virtual returns (bytes4) {
        return MsgValueIncorrect.selector;
    }

    function senderIncorrectSelector() internal pure virtual returns (bytes4) {
        return SenderIncorrect.selector;
    }

    function statusIncorrectSelector() internal pure virtual returns (bytes4) {
        return StatusIncorrect.selector;
    }

    function transactionRelayedSelector() internal pure virtual returns (bytes4) {
        return TransactionRelayed.selector;
    }

    function zeroAddressSelector() internal pure virtual returns (bytes4) {
        return ZeroAddress.selector;
    }
}
