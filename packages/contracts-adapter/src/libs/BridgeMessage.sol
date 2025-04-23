// SPDX-License-Identifier: MIT
pragma solidity ^0.8.4;

library BridgeMessage {
    uint256 private constant BRIDGE_MESSAGE_LENGTH = 96;

    error BridgeMessage__InvalidPayload();

    function encodeBridgeMessage(
        address recipient,
        bytes31 symbol,
        uint256 amount
    )
        internal
        pure
        returns (bytes memory)
    {
        return abi.encode(recipient, symbol, amount);
    }

    function decodeBridgeMessage(bytes calldata payload)
        internal
        pure
        returns (address recipient, bytes31 symbol, uint256 amount)
    {
        if (payload.length != BRIDGE_MESSAGE_LENGTH) revert BridgeMessage__InvalidPayload();
        return abi.decode(payload, (address, bytes31, uint256));
    }
}
