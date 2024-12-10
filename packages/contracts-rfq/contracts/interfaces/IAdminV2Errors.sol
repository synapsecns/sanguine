// SPDX-License-Identifier: MIT
pragma solidity ^0.8.4;

interface IAdminV2Errors {
    error CancelDelayBelowMin();
    error FeeRateAboveMax();
    error ProverAlreadyActive();
    error ProverCapacityExceeded();
    error ProverNotActive();
    error DisputePenaltyTimeBelowMin();
}
