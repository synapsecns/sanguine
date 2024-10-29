// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import {IZapRecipient} from "../interfaces/IZapRecipient.sol";
import {ZapDataV1} from "../libs/ZapDataV1.sol";

import {Address} from "@openzeppelin/contracts/utils/Address.sol";
import {SafeERC20, IERC20} from "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";

contract TokenZap is IZapRecipient {
    using SafeERC20 for IERC20;
    using ZapDataV1 for bytes;

    address public constant NATIVE_GAS_TOKEN = 0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE;

    error TokenZap__AmountIncorrect();
    error TokenZap__PayloadLengthAboveMax();

    /// @notice Performs a Zap action using the specified token and amount. This amount must be previously
    /// transferred to this contract (or supplied as msg.value if the token is native gas token).
    /// @dev The provided ZapData contains the target address and calldata for the Zap action, and must be
    /// encoded using the encodeZapData function.
    /// @param token        Address of the token to be used for the Zap action.
    /// @param amount       Amount of the token to be used for the Zap action.
    ///                     Must match msg.value if the token is native gas token.
    /// @param zapData      Encoded Zap Data containing the target address and calldata for the Zap action.
    /// @return selector    Selector of this function to signal the caller about the success of the Zap action.
    function zap(address token, uint256 amount, bytes calldata zapData) external payable returns (bytes4) {
        // Check that the ZapData is valid before decoding it
        zapData.validateV1();
        address target = zapData.target();
        // Approve the target contract to spend the token. TokenZap does not custody any tokens outside of the
        // zap action, so we can approve the arbitrary target contract.
        if (token == NATIVE_GAS_TOKEN) {
            // No approvals are needed for the native gas token, just check that the amount is correct
            if (msg.value != amount) revert TokenZap__AmountIncorrect();
        } else {
            // Issue the approval only if the current allowance is less than the required amount
            if (IERC20(token).allowance(address(this), target) < amount) {
                IERC20(token).forceApprove(target, type(uint256).max);
            }
        }
        // Perform the Zap action, forwarding full msg.value to the target contract
        // Note: this will bubble up any revert from the target contract
        bytes memory payload = zapData.payload(amount);
        Address.functionCallWithValue({target: target, data: payload, value: msg.value});
        return this.zap.selector;
    }

    /// @notice Encodes the ZapData for a Zap action.
    /// Note: at the time of encoding we don't know the exact amount of tokens that will be used for the Zap,
    /// as we don't have a quote for performing a Zap. Therefore a placeholder value for amount must be used
    /// when abi-encoding the payload. A reference index where the actual amount is encoded within the payload
    /// must be provided in order to replace the placeholder with the actual amount when the Zap is performed.
    /// @param target           Address of the target contract.
    /// @param payload          ABI-encoded calldata to be used for the `target` contract call.
    ///                         If the target function has the token amount as an argument, any placeholder amount value
    ///                         can be used for the original ABI encoding of `payload`. The placeholder amount will
    ///                         be replaced with the actual amount, when the Zap Data is decoded.
    /// @param amountPosition   Position (start index) where the token amount is encoded within `payload`.
    ///                         This will usually be `4 + 32 * n`, where `n` is the position of the token amount in
    ///                         the list of parameters of the target function (starting from 0).
    ///                         Any value greater or equal to `payload.length` can be used if the token amount is
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
            revert TokenZap__PayloadLengthAboveMax();
        }
        if (amountPosition >= payload.length) {
            amountPosition = ZapDataV1.AMOUNT_NOT_PRESENT;
        }
        // At this point we checked that both amountPosition and payload.length fit in uint16
        return ZapDataV1.encodeV1(uint16(amountPosition), target, payload);
    }

    /// @notice Decodes the ZapData for a Zap action. Replaces the placeholder amount with the actual amount,
    /// if it was present in the original `payload`. Otherwise returns the original `payload` as is.
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
