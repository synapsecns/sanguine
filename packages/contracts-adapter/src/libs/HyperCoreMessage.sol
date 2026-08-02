// SPDX-License-Identifier: MIT
pragma solidity ^0.8.4;

library HyperCoreMessage {
    uint8 internal constant VERSION = 1;
    uint256 internal constant HYPERCORE_MESSAGE_LENGTH = 128;

    error HyperCoreMessage__InvalidPayload();
    error HyperCoreMessage__UnsupportedVersion(uint8 version);

    function encodeHyperCoreMessage(
        address recipient,
        address srcToken,
        uint256 amount
    )
        internal
        pure
        returns (bytes memory)
    {
        return abi.encode(VERSION, recipient, srcToken, amount);
    }

    function decodeHyperCoreMessage(bytes calldata payload)
        internal
        pure
        returns (address recipient, address srcToken, uint256 amount)
    {
        if (payload.length != HYPERCORE_MESSAGE_LENGTH) revert HyperCoreMessage__InvalidPayload();
        uint8 version;
        (version, recipient, srcToken, amount) = abi.decode(payload, (uint8, address, address, uint256));
        if (version != VERSION) revert HyperCoreMessage__UnsupportedVersion(version);
    }
}
