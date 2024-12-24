// SPDX-License-Identifier: MIT
pragma solidity ^0.8.17;

// solhint-disable no-empty-blocks
contract MockRevertingRecipient {
    receive() external payable {
        revert("GM");
    }

    /// @notice We include an empty "test" function so that this contract does not appear in the coverage report.
    function testMockRevertingRecipient() external {}
}
