// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

/// @notice A collection of events emitted by the Summit contract
abstract contract SummitEvents {
    /**
     * @notice Emitted when a tip is awarded to the actor, whether they are bonded or unbonded actor.
     * @param actor     Actor address
     * @param origin    Domain where tips were originally paid
     * @param tip       Tip value, scaled down by TIPS_MULTIPLIER
     */
    event TipAwarded(address actor, uint32 origin, uint256 tip);
}
