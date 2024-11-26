// SPDX-License-Identifier: MIT
pragma solidity ^0.8.4;

import {IZapRecipient} from "../../contracts/interfaces/IZapRecipient.sol";

// solhint-disable no-empty-blocks
/// @notice Recipient mock for testing purposes. DO NOT USE IN PRODUCTION.
contract RecipientMock is IZapRecipient {
    /// @notice Mock needs to accept ETH
    receive() external payable {}

    /// @notice We include an empty "test" function so that this contract does not appear in the coverage report.
    function testRecipientMock() external {}

    /// @notice Minimal viable implementation of the zap hook.
    function zap(address, uint256, bytes memory) external payable returns (bytes4) {
        return IZapRecipient.zap.selector;
    }
}
