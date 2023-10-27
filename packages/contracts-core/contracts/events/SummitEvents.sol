// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

/// @notice A collection of events emitted by the Summit contract
abstract contract SummitEvents {
    /**
     * @notice Emitted when a tip is awarded to the actor, whether they are bonded or unbonded actor.
     * @param actor     Actor address
     * @param origin    Domain where tips were originally paid
     * @param tip       Tip value, denominated in domain's wei
     */
    event TipAwarded(address actor, uint32 origin, uint256 tip);

    /**
     * @notice Emitted when a tip withdrawal is initiated by the actor.
     * @param actor     Actor address
     * @param origin    Domain where tips were originally paid
     * @param tip       Tip value, denominated in domain's wei
     */
    event TipWithdrawalInitiated(address actor, uint32 origin, uint256 tip);
}
