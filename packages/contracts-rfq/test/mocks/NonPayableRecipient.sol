// SPDX-License-Identifier: MIT
pragma solidity ^0.8.4;

// solhint-disable no-empty-blocks
/// @notice Incorrectly implemented recipient mock for testing purposes. DO NOT USE IN PRODUCTION.
contract NonPayableRecipient {
    /// @notice We include an empty "test" function so that this contract does not appear in the coverage report.
    function testNonPayableRecipient() external {}

    /// @notice Incorrectly implemented - method is not payable.
    function zap(address, uint256, bytes memory) external pure returns (bytes4) {
        return NonPayableRecipient.zap.selector;
    }
}
