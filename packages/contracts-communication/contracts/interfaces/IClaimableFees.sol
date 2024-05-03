// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

interface IClaimableFees {
    error ClaimableFees__ClaimerFractionAboveMax(uint256 claimerFraction, uint256 maxAllowed);
    error ClaimableFees__FeeRecipientNotSet();
    error ClaimableFees__ZeroAmount();

    function claimFees() external;

    function getClaimableAmount() external view returns (uint256);
    function getClaimerFraction() external view returns (uint256);
    function getClaimerReward() external view returns (uint256);
    function getFeeRecipient() external view returns (address);
}
