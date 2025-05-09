// SPDX-License-Identifier: MIT
pragma solidity ^0.8.4;

import {BridgeMessage} from "../../src/libs/BridgeMessage.sol";

// solhint-disable no-empty-blocks
contract BridgeMessageHarness {
    /// @notice We include an empty "test" function so that this contract does not appear in the coverage report.
    function testBridgeMessageHarness() external {}

    function encodeBridgeMessage(
        address recipient,
        bytes31 symbol,
        uint256 amount
    )
        public
        pure
        returns (bytes memory)
    {
        return BridgeMessage.encodeBridgeMessage(recipient, symbol, amount);
    }

    function decodeBridgeMessage(bytes calldata payload)
        public
        pure
        returns (address recipient, bytes31 symbol, uint256 amount)
    {
        return BridgeMessage.decodeBridgeMessage(payload);
    }
}
