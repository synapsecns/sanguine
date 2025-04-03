// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import {IFastBridge, IFastBridgeInterceptor} from "./interfaces/IFastBridgeInterceptor.sol";

contract FastBridgeInterceptor is IFastBridgeInterceptor {
    /// @inheritdoc IFastBridgeInterceptor
    function bridgeWithInterception(
        IFastBridge.BridgeParams memory params,
        InterceptorParams memory interceptorParams
    )
        external
        payable
    {
        // TODO: implement
    }
}
