// SPDX-License-Identifier: MIT
pragma solidity ^0.8.4;

import {IFastBridge} from "./IFastBridge.sol";

interface IFastBridgeInterceptor {
    /// @notice Parameters for adjusting a quote on the fly.
    /// @param fastBridge           The address of the FastBridge contract
    /// @param quoteOriginAmount    The quote's initial origin amount
    struct InterceptorParams {
        address fastBridge;
        uint256 quoteOriginAmount;
    }

    error FBI__OriginAmountOutOfRange(uint256 originAmount, uint256 quoteOriginAmount);
    error FBI__TokenNotContract(address token);

    /// @notice Initiates bridge on origin chain via FastBridge contract on behalf of the user.
    /// Intercepts the bridge process to adjust the destination amount in `params` based on `params.originAmount`
    /// and the price rate of the initial quote.
    /// @dev Will revert if `params.originAmount` differs from initial quote's origin amount by more than 1%.
    /// @param params               The bridge parameters with amount possibly different from the initial quote.
    /// @param interceptorParams   The interception parameters with the initial quote.
    function bridgeWithInterception(
        IFastBridge.BridgeParams memory params,
        InterceptorParams memory interceptorParams
    )
        external
        payable;
}
