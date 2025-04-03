// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import {IFastBridge, IFastBridgeInterceptor} from "./interfaces/IFastBridgeInterceptor.sol";

import {IERC20, SafeERC20} from "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";

/// @title FastBridgeInterceptor
/// @notice Contract that enables flexible bridging with automatic price adjustment.
/// @dev Enforces a maximum 1% deviation limit between quoted and actual amounts.
contract FastBridgeInterceptor is IFastBridgeInterceptor {
    using SafeERC20 for IERC20;

    /// @notice Maximum allowed difference (1%) from the quoted origin amount.
    uint256 public constant MAX_ORIGIN_AMOUNT_DIFF = 0.01e18;

    /// @notice Special address representing the native gas token (ETH).
    address public constant NATIVE_GAS_TOKEN = 0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE;

    /// @dev Base unit for percentage calculations (100% = 1e18).
    uint256 internal constant WEI = 1e18;

    /// @dev Minimum allowed percentage (99%) for origin amount comparisons.
    uint256 internal constant MIN_ORIGIN_AMOUNT = WEI - MAX_ORIGIN_AMOUNT_DIFF;

    /// @dev Maximum allowed percentage (101%) for origin amount comparisons.
    uint256 internal constant MAX_ORIGIN_AMOUNT = WEI + MAX_ORIGIN_AMOUNT_DIFF;

    /// @inheritdoc IFastBridgeInterceptor
    function bridgeWithInterception(
        IFastBridge.BridgeParams memory params,
        InterceptorParams memory interceptorParams
    )
        external
        payable
    {
        // Cache amounts from memory.
        uint256 originAmount = params.originAmount;
        uint256 quoteOriginAmount = interceptorParams.quoteOriginAmount;
        // Check if origin amount is within 1% of the initial quote origin amount.
        uint256 minOriginAmount = quoteOriginAmount * MIN_ORIGIN_AMOUNT / WEI;
        uint256 maxOriginAmount = quoteOriginAmount * MAX_ORIGIN_AMOUNT / WEI;
        if (originAmount < minOriginAmount || originAmount > maxOriginAmount) {
            revert FBI__OriginAmountOutOfRange(originAmount, quoteOriginAmount);
        }
        // Adjust the destination amount using the initial quote rate.
        if (originAmount != quoteOriginAmount) {
            uint256 quoteDestAmount = params.destAmount;
            params.destAmount = (originAmount * quoteDestAmount) / quoteOriginAmount;
        }
        // Collect the token from msg.sender and approve the FastBridge contract, if needed.
        address token = params.originToken;
        address fastBridge = interceptorParams.fastBridge;
        if (token != NATIVE_GAS_TOKEN) {
            if (token.code.length == 0) revert FBI__TokenNotContract(token);
            IERC20(token).safeTransferFrom(msg.sender, address(this), originAmount);
            // Infinite approval as this contract is stateless and doesn't custody tokens.
            if (IERC20(token).allowance(address(this), fastBridge) < originAmount) {
                IERC20(token).forceApprove(fastBridge, type(uint256).max);
            }
        }
        // Bridge the token, forwarding the original msg.value.
        IFastBridge(fastBridge).bridge{value: msg.value}(params);
    }
}
