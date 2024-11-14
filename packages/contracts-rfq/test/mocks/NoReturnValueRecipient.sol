// SPDX-License-Identifier: MIT
pragma solidity ^0.8.4;

// solhint-disable no-empty-blocks
/// @notice Incorrectly implemented recipient mock for testing purposes. DO NOT USE IN PRODUCTION.
contract NoReturnValueRecipient {
    /// @notice Mock needs to accept ETH
    receive() external payable {}

    /// @notice We include an empty "test" function so that this contract does not appear in the coverage report.
    function testNoReturnValueRecipient() external {}

    /// @notice Incorrectly implemented - method does not return anything.
    function zap(address, uint256, bytes memory) external payable {}
}
