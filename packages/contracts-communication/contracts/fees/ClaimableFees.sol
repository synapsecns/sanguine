// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {ClaimableFeesEvents} from "../events/ClaimableFeesEvents.sol";
import {IClaimableFees} from "../interfaces/IClaimableFees.sol";

import {Address} from "@openzeppelin/contracts/utils/Address.sol";

/// @notice A simple abstraction for a contract that is collecting fees in native chain token.
/// The claim process could be performed by anyone, but the fees will be sent to
/// the predefined address. The claimer will receive a fraction of the fees to offset
/// the gas costs.
/// @dev The contract is implemented in a stateless way to allow the inheriting
/// contract to be immutable or upgradeable.
abstract contract ClaimableFees is ClaimableFeesEvents, IClaimableFees {
    uint256 private constant FEE_PRECISION = 1e18;
    /// @dev The maximum fraction that the claimer can receive is 1%.
    uint256 internal constant MAX_CLAIMER_FRACTION = 1e16;

    /// @notice Transfers the accumulated fees to the fee recipient.
    /// Message caller receives a fraction of the fees as a reward to offset the gas costs.
    /// The reward amount could be obtained by calling the `getClaimerReward` function beforehand.
    /// @dev Will revert if the claimable amount is zero or the fee recipient is not set.
    function claimFees() external {
        uint256 amount = getClaimableAmount();
        if (amount == 0) {
            revert ClaimableFees__ZeroAmount();
        }
        address recipient = getFeeRecipient();
        if (recipient == address(0)) {
            revert ClaimableFees__FeeRecipientNotSet();
        }
        // Subtract the claimer reward from the total amount
        uint256 reward = getClaimerReward();
        _beforeFeesClaimed(amount, reward);
        // We can do unchecked subtraction because `getClaimerReward` ensures that `reward <= amount * 0.01`
        unchecked {
            amount -= reward;
        }
        // Emit the event before transferring the fees
        emit FeesClaimed(recipient, amount, msg.sender, reward);
        Address.sendValue(payable(recipient), amount);
        Address.sendValue(payable(msg.sender), reward);
    }

    /// @notice Returns the amount of native chain token that the claimer will receive
    /// after calling the `claimFees` function.
    function getClaimerReward() public view returns (uint256) {
        uint256 fraction = getClaimerFraction();
        if (fraction > MAX_CLAIMER_FRACTION) {
            revert ClaimableFees__ClaimerFractionExceedsMax(fraction);
        }
        // The returned value is in the range [0, _getClaimableAmount() * 0.01]
        return (getClaimableAmount() * fraction) / FEE_PRECISION;
    }

    /// @notice Returns the amount of fees that can be claimed.
    function getClaimableAmount() public view virtual returns (uint256);

    /// @notice Returns the fraction of the fees that the claimer will receive.
    /// The result is in the range [0, 1e18], where 1e18 is 100%.
    function getClaimerFraction() public view virtual returns (uint256);

    /// @notice Returns the address that will receive the claimed fees.
    function getFeeRecipient() public view virtual returns (address);

    /// @dev Hook that is called before the fees are claimed.
    /// Useful if the inheriting contract needs to manage the state when the fees are claimed.
    function _beforeFeesClaimed(uint256 fullAmount, uint256 reward) internal virtual;
}
