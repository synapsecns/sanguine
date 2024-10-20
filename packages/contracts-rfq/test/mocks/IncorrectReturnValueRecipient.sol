// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {IFastBridgeRecipient} from "../../contracts/interfaces/IFastBridgeRecipient.sol";

// solhint-disable no-empty-blocks
/// @notice Incorrectly implemented recipient mock for testing purposes. DO NOT USE IN PRODUCTION.
contract IncorrectReturnValueRecipient is IFastBridgeRecipient {
    /// @notice Mock needs to accept ETH
    receive() external payable {}

    /// @notice We include an empty "test" function so that this contract does not appear in the coverage report.
    function testIncorrectReturnValueRecipient() external {}

    /// @notice Incorrectly implemented - method returns incorrect value.
    function fastBridgeTransferReceived(address, uint256, bytes memory) external payable returns (bytes4) {
        // Flip the last bit
        return IFastBridgeRecipient.fastBridgeTransferReceived.selector ^ 0x00000001;
    }
}
