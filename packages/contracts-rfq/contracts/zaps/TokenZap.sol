// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import {IZapRecipient} from "../interfaces/IZapRecipient.sol";

contract TokenZap is IZapRecipient {
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
        // TODO: implement
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
        view
        returns (bytes memory)
    {
        // TODO: implement
    }

    /// @notice Decodes the ZapData for a Zap action. Replaces the placeholder amount with the actual amount,
    /// if it was present in the original `payload`. Otherwise returns the original `payload` as is.
    /// @param zapData          Encoded Zap Data containing the target address and calldata for the Zap action.
    /// @param amount           Actual amount of the token to be used for the Zap action.
    function decodeZapData(
        bytes calldata zapData,
        uint256 amount
    )
        external
        view
        returns (address target, bytes memory payload)
    {
        // TODO: implement
    }
}
