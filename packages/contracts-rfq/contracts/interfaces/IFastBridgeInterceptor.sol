// SPDX-License-Identifier: MIT
pragma solidity ^0.8.4;

import {IFastBridge} from "./IFastBridge.sol";

interface IFastBridgeInterceptor {
    /// @notice Parameters for adjusting a bridge quote on the fly.
    /// @param fastBridge           The FastBridge contract address to call.
    /// @param quoteOriginAmount    The original quoted origin amount (for price calculations).
    struct InterceptorParams {
        address fastBridge;
        uint256 quoteOriginAmount;
    }

    /// @notice Error thrown when the origin amount is outside the allowed range (±1% of quote).
    /// @param originAmount         The actual origin amount provided.
    /// @param quoteOriginAmount    The quote's original origin amount.
    error FBI__OriginAmountOutOfRange(uint256 originAmount, uint256 quoteOriginAmount);

    /// @notice Error thrown when the token address provided has no code (is not a contract).
    /// @param token                The address that was expected to be a token contract.
    error FBI__TokenNotContract(address token);

    /// @notice Bridges tokens with automatic destination amount adjustment based on actual sent amount.
    ///         For ERC20 tokens: Requires token approval to this contract (not FastBridge).
    ///         For ETH: Send the desired ETH amount via msg.value.
    ///         Will revert if amounts differ by more than 1% or an invalid token is provided.
    /// @dev Handles token transfers/approvals and forwards calls to FastBridge with adjusted destination amounts.
    /// @param params               Bridge parameters with the actual amount to bridge.
    /// @param interceptorParams    Original quote parameters used for price calculation.
    function bridgeWithInterception(
        IFastBridge.BridgeParams memory params,
        InterceptorParams memory interceptorParams
    )
        external
        payable;
}
