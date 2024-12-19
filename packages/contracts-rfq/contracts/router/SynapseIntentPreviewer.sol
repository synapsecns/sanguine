// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

// ════════════════════════════════════════════════ INTERFACES ═════════════════════════════════════════════════════

import {ISynapseIntentPreviewer} from "../interfaces/ISynapseIntentPreviewer.sol";
import {ISynapseIntentRouter} from "../interfaces/ISynapseIntentRouter.sol";
import {ISwapQuoter} from "../legacy/rfq/interfaces/ISwapQuoter.sol";
import {IDefaultExtendedPool, IDefaultPool} from "../legacy/router/interfaces/IDefaultExtendedPool.sol";
import {IWETH9} from "../legacy/router/interfaces/IWETH9.sol";

// ═════════════════════════════════════════════ INTERNAL IMPORTS ══════════════════════════════════════════════════

import {Action, DefaultParams, LimitedToken, SwapQuery} from "../legacy/router/libs/Structs.sol";
import {ZapDataV1} from "../libs/ZapDataV1.sol";

contract SynapseIntentPreviewer is ISynapseIntentPreviewer {
    /// @notice The address reserved for the native gas token (ETH on Ethereum and most L2s, AVAX on Avalanche, etc.).
    address public constant NATIVE_GAS_TOKEN = 0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE;

    /// @dev Amount value that signals that the Zap step should be performed using the full ZapRecipient balance.
    uint256 internal constant FULL_BALANCE = type(uint256).max;

    /// @dev Maximum allowed slippage for the intent preview (100%). Use extreme caution when using this value.
    uint256 internal constant MAX_SLIPPAGE = 10 ** 18;

    error SIP__NoOpForwardNotSupported();
    error SIP__PoolTokenMismatch();
    error SIP__PoolZeroAddress();
    error SIP__RawParamsEmpty();
    error SIP__TokenNotNative();

    /// @inheritdoc ISynapseIntentPreviewer
    // solhint-disable-next-line code-complexity
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
        returns (uint256 amountOut, ISynapseIntentRouter.StepParams[] memory steps)
    {
        // First, check if the intent is a no-op.
        if (tokenIn == tokenOut) {
            if (forwardTo != address(0)) revert SIP__NoOpForwardNotSupported();
            return (amountIn, new ISynapseIntentRouter.StepParams[](0));
        }

        // Obtain the swap quote, don't put any restrictions on the actions allowed to complete the intent.
        SwapQuery memory query = ISwapQuoter(swapQuoter).getAmountOut(
            LimitedToken({token: tokenIn, actionMask: type(uint256).max}), tokenOut, amountIn
        );

        // Check if a quote was returned.
        amountOut = query.minAmountOut;
        if (amountOut == 0) {
            return (0, new ISynapseIntentRouter.StepParams[](0));
        }
        uint256 lastStepMinAmountOut = 0;
        if (slippageWei < MAX_SLIPPAGE) {
            lastStepMinAmountOut = amountOut - (amountOut * slippageWei) / MAX_SLIPPAGE;
        }

        // At this point we have a quote for a non-trivial action, therefore `query.rawParams` is not empty.
        if (query.rawParams.length == 0) revert SIP__RawParamsEmpty();
        DefaultParams memory params = abi.decode(query.rawParams, (DefaultParams));

        // Create the steps for the intent based on the action type.
        if (params.action == Action.Swap) {
            steps = _createSwapSteps(tokenIn, tokenOut, amountIn, params, forwardTo, lastStepMinAmountOut);
        } else if (params.action == Action.AddLiquidity) {
            steps = _createAddLiquiditySteps(tokenIn, tokenOut, params, forwardTo, lastStepMinAmountOut);
        } else if (params.action == Action.RemoveLiquidity) {
            steps = _createRemoveLiquiditySteps(tokenIn, tokenOut, params, forwardTo, lastStepMinAmountOut);
        } else {
            steps = _createHandleHativeSteps(tokenIn, tokenOut, amountIn, forwardTo);
        }
    }

    /// @notice Helper function to create steps for a swap.
    function _createSwapSteps(
        address tokenIn,
        address tokenOut,
        uint256 amountIn,
        DefaultParams memory params,
        address forwardTo,
        uint256 lastStepMinAmountOut
    )
        internal
        view
        returns (ISynapseIntentRouter.StepParams[] memory steps)
    {
        address pool = params.pool;
        if (pool == address(0)) revert SIP__PoolZeroAddress();
        // Default Pools can only host wrapped native tokens.
        // Check if we start from the native gas token.
        if (tokenIn == NATIVE_GAS_TOKEN) {
            // Get the address of the wrapped native token.
            address wrappedNative = IDefaultPool(pool).getToken(params.tokenIndexFrom);
            // Sanity check tokenOut vs tokenIndexTo.
            if (IDefaultPool(pool).getToken(params.tokenIndexTo) != tokenOut) revert SIP__PoolTokenMismatch();
            // Native => WrappedNative + WrappedNative => TokenOut. Forwarding is done in the second step.
            return _toStepsArray(
                _createWrapNativeStep({wrappedNative: wrappedNative, msgValue: amountIn, forwardTo: address(0)}),
                _createSwapStep({
                    tokenIn: wrappedNative,
                    tokenOut: tokenOut,
                    params: params,
                    forwardTo: forwardTo,
                    minAmountOut: lastStepMinAmountOut
                })
            );
        }

        // Sanity check tokenIn vs tokenIndexFrom.
        if (IDefaultPool(pool).getToken(params.tokenIndexFrom) != tokenIn) revert SIP__PoolTokenMismatch();

        // Check if we end with the native gas token.
        if (tokenOut == NATIVE_GAS_TOKEN) {
            // Get the address of the wrapped native token.
            address wrappedNative = IDefaultPool(pool).getToken(params.tokenIndexTo);
            // TokenIn => WrappedNative + WrappedNative => Native. Forwarding/minAmountOut is done in the second step.
            return _toStepsArray(
                _createSwapStep({
                    tokenIn: tokenIn,
                    tokenOut: wrappedNative,
                    params: params,
                    forwardTo: address(0),
                    minAmountOut: 0
                }),
                _createUnwrapNativeStep({wrappedNative: wrappedNative, forwardTo: forwardTo})
            );
        }

        // Sanity check tokenOut vs tokenIndexTo.
        if (IDefaultPool(pool).getToken(params.tokenIndexTo) != tokenOut) revert SIP__PoolTokenMismatch();

        // TokenIn => TokenOut.
        ISynapseIntentRouter.StepParams memory step = _createSwapStep({
            tokenIn: tokenIn,
            tokenOut: tokenOut,
            params: params,
            forwardTo: forwardTo,
            minAmountOut: lastStepMinAmountOut
        });
        return _toStepsArray(step);
    }

    /// @notice Helper function to create steps for adding liquidity.
    function _createAddLiquiditySteps(
        address tokenIn,
        address tokenOut,
        DefaultParams memory params,
        address forwardTo,
        uint256 lastStepMinAmountOut
    )
        internal
        view
        returns (ISynapseIntentRouter.StepParams[] memory steps)
    {
        address pool = params.pool;
        if (pool == address(0)) revert SIP__PoolZeroAddress();
        // Sanity check tokenIn vs tokenIndexFrom.
        if (IDefaultPool(pool).getToken(params.tokenIndexFrom) != tokenIn) revert SIP__PoolTokenMismatch();
        // Sanity check tokenOut vs pool's LP token.
        _verifyLpToken(pool, tokenOut);
        // Figure out how many tokens does the pool support.
        uint256[] memory amounts;
        for (uint8 i = 0;; i++) {
            // solhint-disable-next-line no-empty-blocks
            try IDefaultExtendedPool(pool).getToken(i) returns (address) {
                // Token exists, continue.
            } catch {
                // No more tokens, allocate the array using the correct size.
                amounts = new uint256[](i);
                break;
            }
        }
        return _toStepsArray(
            ISynapseIntentRouter.StepParams({
                token: tokenIn,
                amount: FULL_BALANCE,
                msgValue: 0,
                zapData: ZapDataV1.encodeV1({
                    target_: pool,
                    finalToken_: tokenOut,
                    forwardTo_: forwardTo,
                    // addLiquidity(amounts, minToMint, deadline)
                    payload_: abi.encodeCall(
                        IDefaultExtendedPool.addLiquidity, (amounts, lastStepMinAmountOut, type(uint256).max)
                    ),
                    // amountIn is encoded within `amounts` at `TOKEN_IN_INDEX`, `amounts` is encoded after
                    // (amounts.offset, minToMint, deadline, amounts.length).
                    amountPosition_: 4 + 32 * 4 + 32 * uint16(params.tokenIndexFrom)
                })
            })
        );
    }

    /// @notice Helper function to create steps for removing liquidity.
    function _createRemoveLiquiditySteps(
        address tokenIn,
        address tokenOut,
        DefaultParams memory params,
        address forwardTo,
        uint256 lastStepMinAmountOut
    )
        internal
        view
        returns (ISynapseIntentRouter.StepParams[] memory steps)
    {
        address pool = params.pool;
        if (pool == address(0)) revert SIP__PoolZeroAddress();
        // Sanity check tokenIn vs pool's LP token.
        _verifyLpToken(pool, tokenIn);
        // Sanity check tokenOut vs tokenIndexTo.
        if (IDefaultPool(pool).getToken(params.tokenIndexTo) != tokenOut) revert SIP__PoolTokenMismatch();
        return _toStepsArray(
            ISynapseIntentRouter.StepParams({
                token: tokenIn,
                amount: FULL_BALANCE,
                msgValue: 0,
                zapData: ZapDataV1.encodeV1({
                    target_: pool,
                    finalToken_: tokenOut,
                    forwardTo_: forwardTo,
                    // removeLiquidityOneToken(tokenAmount, tokenIndex, minAmount, deadline)
                    payload_: abi.encodeCall(
                        IDefaultExtendedPool.removeLiquidityOneToken,
                        (0, params.tokenIndexTo, lastStepMinAmountOut, type(uint256).max)
                    ),
                    // amountIn is encoded as the first parameter: tokenAmount
                    amountPosition_: 4
                })
            })
        );
    }

    function _verifyLpToken(address pool, address token) internal view {
        (,,,,,, address lpToken) = IDefaultExtendedPool(pool).swapStorage();
        if (lpToken != token) revert SIP__PoolTokenMismatch();
    }

    /// @notice Helper function to create steps for wrapping or unwrapping native gas tokens.
    function _createHandleHativeSteps(
        address tokenIn,
        address tokenOut,
        uint256 amountIn,
        address forwardTo
    )
        internal
        pure
        returns (ISynapseIntentRouter.StepParams[] memory steps)
    {
        if (tokenIn == NATIVE_GAS_TOKEN) {
            // tokenOut is Wrapped Native
            return _toStepsArray(
                _createWrapNativeStep({wrappedNative: tokenOut, msgValue: amountIn, forwardTo: forwardTo})
            );
        }
        // Sanity check tokenOut
        if (tokenOut != NATIVE_GAS_TOKEN) revert SIP__TokenNotNative();
        // tokenIn is Wrapped Native
        return _toStepsArray(_createUnwrapNativeStep({wrappedNative: tokenIn, forwardTo: forwardTo}));
    }

    /// @notice Helper function to create a single step for a swap.
    function _createSwapStep(
        address tokenIn,
        address tokenOut,
        DefaultParams memory params,
        address forwardTo,
        uint256 minAmountOut
    )
        internal
        pure
        returns (ISynapseIntentRouter.StepParams memory)
    {
        return ISynapseIntentRouter.StepParams({
            token: tokenIn,
            amount: FULL_BALANCE,
            msgValue: 0,
            zapData: ZapDataV1.encodeV1({
                target_: params.pool,
                finalToken_: tokenOut,
                forwardTo_: forwardTo,
                // swap(tokenIndexFrom, tokenIndexTo, dx, minDy, deadline)
                payload_: abi.encodeCall(
                    IDefaultPool.swap, (params.tokenIndexFrom, params.tokenIndexTo, 0, minAmountOut, type(uint256).max)
                ),
                // amountIn is encoded as the third parameter: `dx`
                amountPosition_: 4 + 32 * 2
            })
        });
    }

    /// @notice Helper function to create a single step for wrapping native gas tokens.
    function _createWrapNativeStep(
        address wrappedNative,
        uint256 msgValue,
        address forwardTo
    )
        internal
        pure
        returns (ISynapseIntentRouter.StepParams memory)
    {
        return ISynapseIntentRouter.StepParams({
            token: NATIVE_GAS_TOKEN,
            amount: FULL_BALANCE,
            msgValue: msgValue,
            zapData: ZapDataV1.encodeV1({
                target_: wrappedNative,
                finalToken_: wrappedNative,
                forwardTo_: forwardTo,
                // deposit()
                payload_: abi.encodeCall(IWETH9.deposit, ()),
                // amountIn is not encoded
                amountPosition_: ZapDataV1.AMOUNT_NOT_PRESENT
            })
        });
    }

    /// @notice Helper function to create a single step for unwrapping native gas tokens.
    function _createUnwrapNativeStep(
        address wrappedNative,
        address forwardTo
    )
        internal
        pure
        returns (ISynapseIntentRouter.StepParams memory)
    {
        return ISynapseIntentRouter.StepParams({
            token: wrappedNative,
            amount: FULL_BALANCE,
            msgValue: 0,
            zapData: ZapDataV1.encodeV1({
                target_: wrappedNative,
                finalToken_: NATIVE_GAS_TOKEN,
                forwardTo_: forwardTo,
                // withdraw(amount)
                payload_: abi.encodeCall(IWETH9.withdraw, (0)),
                // amountIn encoded as the first parameter
                amountPosition_: 4
            })
        });
    }

    /// @notice Helper function to construct an array of steps having a single step.
    function _toStepsArray(ISynapseIntentRouter.StepParams memory step0)
        internal
        pure
        returns (ISynapseIntentRouter.StepParams[] memory)
    {
        ISynapseIntentRouter.StepParams[] memory steps = new ISynapseIntentRouter.StepParams[](1);
        steps[0] = step0;
        return steps;
    }

    /// @notice Helper function to construct an array of steps having two steps.
    function _toStepsArray(
        ISynapseIntentRouter.StepParams memory step0,
        ISynapseIntentRouter.StepParams memory step1
    )
        internal
        pure
        returns (ISynapseIntentRouter.StepParams[] memory)
    {
        ISynapseIntentRouter.StepParams[] memory steps = new ISynapseIntentRouter.StepParams[](2);
        steps[0] = step0;
        steps[1] = step1;
        return steps;
    }
}
