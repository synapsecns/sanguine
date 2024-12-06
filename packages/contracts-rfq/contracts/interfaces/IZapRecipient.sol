// SPDX-License-Identifier: MIT
pragma solidity ^0.8.4;

/// @notice Interface for contracts that can perform Zap operations. Such contracts could be used as Recipients
/// in a FastBridge transaction that includes a Zap operation. The Zap Data should include instructions on how
/// exactly the Zap operation should be executed, which would typically include the target address and calldata
/// to use. The exact implementation of the Zap Data encoding is up to the Recipient contract.
interface IZapRecipient {
    /// @notice Performs a Zap operation with the given token and amount according to the provided Zap data.
    /// @param token    The address of the token being used for the Zap operation.
    /// @param amount   The amount of tokens to be used.
    /// @param zapData  The encoded data specifying how the Zap operation should be executed.
    /// @return         The function selector to indicate successful execution.
    function zap(address token, uint256 amount, bytes memory zapData) external payable returns (bytes4);
}
