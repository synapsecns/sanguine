// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

abstract contract ClaimableFeesEvents {
    /// @notice Emitted when the claim fee fraction is set. This fraction of the fees will be paid
    /// to the caller of the `claimFees` function.
    /// This encourages rational actors to call the function as soon as claim fee is higher than the gas cost.
    /// @param claimerFraction  The fraction of the fees to be paid to the claimer (100% = 1e18)
    event ClaimerFractionSet(uint256 claimerFraction);

    /// @notice Emitted when fees are claimed to the fee recipient address.
    /// @param feeRecipient     The address that receives the claimed fees.
    /// @param claimedFees      The amount of fees claimed, after the claimer reward is deducted.
    /// @param claimer          The address of the claimer (who called `claimFees`)
    /// @param claimerReward    The reward paid to the claimer for calling the `claimFees` function.
    event FeesClaimed(address feeRecipient, uint256 claimedFees, address claimer, uint256 claimerReward);
}
