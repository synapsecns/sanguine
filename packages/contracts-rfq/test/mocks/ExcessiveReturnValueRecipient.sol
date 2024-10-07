// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {IFastBridgeRecipient} from "../../contracts/interfaces/IFastBridgeRecipient.sol";

/// @notice Incorrectly implemented recipient mock for testing purposes. DO NOT USE IN PRODUCTION.
contract ExcessiveReturnValueRecipient {
    /// @notice Mock needs to accept ETH
    receive() external payable {}

    /// @notice Incorrectly implemented - method returns excessive bytes.
    function fastBridgeTransferReceived(address, uint256, bytes memory) external payable returns (bytes4, uint256) {
        return (IFastBridgeRecipient.fastBridgeTransferReceived.selector, 1337);
    }
}
