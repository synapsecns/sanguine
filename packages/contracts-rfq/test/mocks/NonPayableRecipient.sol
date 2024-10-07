// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

/// @notice Incorrectly implemented recipient mock for testing purposes. DO NOT USE IN PRODUCTION.
contract NonPayableRecipient {
    /// @notice Incorrectly implemented - method is not payable.
    function fastBridgeTransferReceived(address, uint256, bytes memory) external pure returns (bytes4) {
        return NonPayableRecipient.fastBridgeTransferReceived.selector;
    }
}
