// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {IFastBridgeRecipient} from "../../contracts/interfaces/IFastBridgeRecipient.sol";

/// @notice Recipient mock for testing purposes. DO NOT USE IN PRODUCTION.
contract FastBridgeRecipientMock is IFastBridgeRecipient {
    function fastBridgeTransferReceived(address, uint256, bytes memory) external payable returns (bytes4) {
        return IFastBridgeRecipient.fastBridgeTransferReceived.selector;
    }
}
