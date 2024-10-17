// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {IFastBridgeRecipient} from "../../contracts/interfaces/IFastBridgeRecipient.sol";

// solhint-disable no-empty-blocks
/// @notice Recipient mock for testing purposes. DO NOT USE IN PRODUCTION.
contract RecipientMock is IFastBridgeRecipient {
    /// @notice Mock needs to accept ETH
    receive() external payable {}

    /// @notice We include an empty "test" function so that this contract does not appear in the coverage report.
    function testRecipientMock() external {}

    /// @notice Minimal viable implementation of the fastBridgeTransferReceived hook.
    function fastBridgeTransferReceived(address, uint256, bytes memory) external payable returns (bytes4) {
        return IFastBridgeRecipient.fastBridgeTransferReceived.selector;
    }
}
