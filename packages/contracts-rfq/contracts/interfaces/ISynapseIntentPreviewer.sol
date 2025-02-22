// SPDX-License-Identifier: MIT
pragma solidity ^0.8.4;

import {ISynapseIntentRouter} from "./ISynapseIntentRouter.sol";

interface ISynapseIntentPreviewer {
    /// @notice Preview the completion of a user intent.
    /// @dev Will not revert if the intent cannot be completed, returns empty values instead.
    /// @dev Returns (amountIn, []) if the intent is a no-op (tokenIn == tokenOut).
    /// @param swapQuoter   Peripheral contract to use for swap quoting
    /// @param forwardTo    The address to which the proceeds of the intent should be forwarded to.
    ///                     Note: if no forwarding is required (or done within the intent), use address(0).
    /// @param slippageWei  Slippage settings to apply when encoding the intent steps. Capped at 100%.
    /// @param tokenIn      Initial token for the intent
    /// @param tokenOut     Final token for the intent
    /// @param amountIn     Initial amount of tokens to use for the intent
    /// @return amountOut   Final amount of tokens to receive. Zero if the intent cannot be completed.
    /// @return steps       Steps to use in SynapseIntentRouter in order to complete the intent.
    ///                     Empty if the intent cannot be completed, or if intent is a no-op (tokenIn == tokenOut).
    function previewIntent(
        address swapQuoter,
        address forwardTo,
        uint256 slippageWei,
        address tokenIn,
        address tokenOut,
        uint256 amountIn
    )
        external
        view
        returns (uint256 amountOut, ISynapseIntentRouter.StepParams[] memory steps);
}
