// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

// ════════════════════════════════════════════════ INTERFACES ═════════════════════════════════════════════════════

import {IZapRecipient} from "../interfaces/IZapRecipient.sol";

// ═════════════════════════════════════════════ INTERNAL IMPORTS ══════════════════════════════════════════════════

import {ZapDataV1} from "../libs/ZapDataV1.sol";

// ═════════════════════════════════════════════ EXTERNAL IMPORTS ══════════════════════════════════════════════════

import {IERC20, SafeERC20} from "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import {Address} from "@openzeppelin/contracts/utils/Address.sol";

/// @title TokenZapV1
/// @notice Facilitates atomic token operations known as "Zaps," allowing to execute predefined actions
/// on behalf of users like deposits or swaps. Supports ERC20 tokens and native gas tokens (e.g., ETH).
/// @dev Tokens must be pre-transferred to the contract for execution, with native tokens sent as msg.value.
/// This contract is stateless and does not hold assets between Zaps; leftover tokens can be claimed by anyone.
/// Ensure Zaps fully utilize tokens or revert to prevent fund loss.
contract TokenZapV1 is IZapRecipient {
    using SafeERC20 for IERC20;
    using ZapDataV1 for bytes;

    address public constant NATIVE_GAS_TOKEN = 0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE;

    error TokenZapV1__PayloadLengthAboveMax();

    /// @notice Allows the contract to receive ETH.
    /// @dev Leftover ETH can be claimed by anyone, make sure to spend the full balance during the Zaps.
    receive() external payable {}

    /// @notice Performs a Zap action using the specified token and amount. This amount must be previously
    /// transferred to this contract (could also be supplied as msg.value if the token is native gas token).
    /// Note: all funds remaining after the Zap action is performed can be claimed by anyone.
    /// Make sure to spend the full balance during the Zaps and avoid sending extra funds if a single Zap is performed.
    /// @dev The provided ZapData contains the target address and calldata for the Zap action, and must be
    /// encoded using the encodeZapData function.
    /// @param token        Address of the token to be used for the Zap action.
    /// @param amount       Amount of the token to be used for the Zap action.
    /// @param zapData      Encoded Zap Data containing the target address and calldata for the Zap action.
    /// @return selector    Selector of this function to signal the caller about the success of the Zap action.
    function zap(address token, uint256 amount, bytes calldata zapData) external payable returns (bytes4) {
        // Validate the ZapData format and extract the target address.
        zapData.validateV1();
        address target = zapData.target();
        uint256 msgValue = msg.value;
        // Note: we don't check the amount that was transferred to TokenZapV1 (or msg.value for native gas tokens),
        // so transfering more than `amount` will lead to remaining funds in TokenZapV1, which can be claimed by anyone.
        // Make sure to send the exact amount for a single Zap or spend the full balance for multiple `zap()` calls.
        if (token == NATIVE_GAS_TOKEN) {
            // For native gas tokens we forward the requested amount to the target contract during the Zap action.
            // Silimar to ERC20s, we allow to use pre-transferred native tokens for the Zap.
            msgValue = amount;
            // No approval needed since native token doesn't use allowances.
            // Note: balance check is performed within `Address.sendValue` or `Address.functionCallWithValue` below.
        } else {
            // For ERC20 tokens, grant unlimited approval to the target if the current allowance is insufficient.
            // This is safe since the contract doesn't custody tokens between zaps.
            if (IERC20(token).allowance(address(this), target) < amount) {
                IERC20(token).forceApprove(target, type(uint256).max);
            }
            // Note: balance check is omitted as the target contract will revert if there are insufficient funds.
        }
        // Construct the payload for the target contract call with the Zap action.
        // The payload is modified to replace the placeholder amount with the actual amount.
        bytes memory payload = zapData.payload(amount);
        if (payload.length == 0) {
            // No payload provided, perform the native gas token transfer to the target.
            // Note: we avoid using `functionCallWithValue` because target might be an EOA. This will
            // revert with a generic custom error should the target contract revert on incoming transfer.
            Address.sendValue({recipient: payable(target), amount: msgValue});
        } else {
            // Perform the Zap action, forwarding full msg.value to the target contract.
            // Note: this will bubble up any revert from the target contract.
            Address.functionCallWithValue({target: target, data: payload, value: msgValue});
        }
        // Return function selector to indicate successful execution
        return this.zap.selector;
    }

    /// @notice Encodes the ZapData for a Zap action.
    /// @dev At the time of encoding, we don't know the exact amount of tokens that will be used for the Zap,
    /// as we don't have a quote for performing a Zap. Therefore, a placeholder value for the amount must be used
    /// when ABI-encoding the payload. A reference index where the actual amount is encoded within the payload
    /// must be provided in order to replace the placeholder with the actual amount when the Zap is performed.
    /// @param target           Address of the target contract.
    /// @param payload          ABI-encoded calldata to be used for the `target` contract call.
    ///                         If the target function has the token amount as an argument, any placeholder amount value
    ///                         can be used for the original ABI encoding of `payload`. The placeholder amount will
    ///                         be replaced with the actual amount when the Zap Data is decoded.
    /// @param amountPosition   Position (start index) where the token amount is encoded within `payload`.
    ///                         This will usually be `4 + 32 * n`, where `n` is the position of the token amount in
    ///                         the list of parameters of the target function (starting from 0).
    ///                         Any value greater than or equal to `payload.length` can be used if the token amount is
    ///                         not an argument of the target function.
    function encodeZapData(
        address target,
        bytes memory payload,
        uint256 amountPosition
    )
        external
        pure
        returns (bytes memory)
    {
        if (payload.length > ZapDataV1.AMOUNT_NOT_PRESENT) {
            revert TokenZapV1__PayloadLengthAboveMax();
        }
        // External integrations do not need to understand the specific `AMOUNT_NOT_PRESENT` semantics.
        // Therefore, they can specify any value greater than or equal to `payload.length` to indicate
        // that the amount is not present in the payload.
        if (amountPosition >= payload.length) {
            amountPosition = ZapDataV1.AMOUNT_NOT_PRESENT;
        }
        // At this point, we have checked that both `amountPosition` and `payload.length` fit in uint16.
        return ZapDataV1.encodeV1(uint16(amountPosition), target, payload);
    }

    /// @notice Decodes the ZapData for a Zap action. Replaces the placeholder amount with the actual amount,
    /// if it was present in the original `payload`. Otherwise, returns the original `payload` as is.
    /// @param zapData          Encoded Zap Data containing the target address and calldata for the Zap action.
    /// @param amount           Actual amount of the token to be used for the Zap action.
    function decodeZapData(
        bytes calldata zapData,
        uint256 amount
    )
        public
        pure
        returns (address target, bytes memory payload)
    {
        zapData.validateV1();
        target = zapData.target();
        payload = zapData.payload(amount);
    }
}
