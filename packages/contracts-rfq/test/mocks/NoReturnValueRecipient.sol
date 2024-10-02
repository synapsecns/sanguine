// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

// solhint-disable

/// @notice Incorrectly implemented recipient mock for testing purposes. DO NOT USE IN PRODUCTION.
contract NoReturnValueRecipient {
    /// @notice Mock needs to accept ETH
    receive() external payable {}

    /// @notice Incorrectly implemented - method does not return anything.
    function fastBridgeTransferReceived(address, uint256, bytes memory) external payable {}
}
