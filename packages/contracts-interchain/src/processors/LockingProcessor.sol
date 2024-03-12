// SPDX-License-Identifier: MIT
pragma solidity 0.8.23;

import {AbstractProcessor} from "./AbstractProcessor.sol";

import {ICERC20} from "../interfaces/ICERC20.sol";

import {IERC20, SafeERC20} from "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";

/// @notice LockingProcessor is a contract that enables the conversion between
/// the ERC20 token (underlying) and its ICERC20 counterpart (interchain) by using the lock-unlock
/// mechanism.
/// - Interchain token is minted when the ERC20 token is locked.
/// - ERC20 token is unlocked when the Interchain token is burned.
/// See AbstractProcessor.sol for more details.
contract LockingProcessor is AbstractProcessor {
    using SafeERC20 for IERC20;

    /// @inheritdoc AbstractProcessor
    function calculateSwap(
        uint8 tokenIndexFrom,
        uint8 tokenIndexTo,
        uint256 dx
    )
        external
        view
        override
        returns (uint256 amountOut)
    {
        // InterchainToken (0) -> UnderlyingToken (1)
        if (tokenIndexFrom == 0 && tokenIndexTo == 1) {
            // Check that amount does not exceed the processor balance of the UnderlyingToken
            if (dx > IERC20(UNDERLYING_TOKEN).balanceOf(address(this))) {
                return 0;
            }
            return dx;
        }
        // UnderlyingToken (1) -> InterchainToken (0)
        if (tokenIndexFrom == 1 && tokenIndexTo == 0) {
            return dx;
        }
        // Return 0 for unsupported operations
        return 0;
    }

    /// @dev Burns the ICERC20 token taken from `msg.sender`, then
    /// unlocks the same amount of the underlying token to `msg.sender`.
    function _burnInterchainToken(uint256 amount) internal override {
        ICERC20(INTERCHAIN_TOKEN).burn(amount);
        IERC20(UNDERLYING_TOKEN).safeTransfer(msg.sender, amount);
    }

    /// @dev Locks the underlying token taken from `msg.sender`, then
    /// mints the same amount of the ICERC20 token to `msg.sender`.
    function _mintInterchainToken(uint256 amount) internal override {
        // UNDERLYING_TOKEN is already transferred to this contract, no extra steps for locking are needed
        ICERC20(INTERCHAIN_TOKEN).mint(msg.sender, amount);
    }
}
