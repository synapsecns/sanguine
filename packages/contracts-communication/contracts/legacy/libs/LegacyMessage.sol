// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

library LegacyMessageLib {
    /// @notice Encodes a message to be sent through the Legacy MessageBus.
    /// @param srcSender    The address of the sender on the source chain
    /// @param dstReceiver  The address of the receiver on the destination chain
    /// @param srcNonce     The nonce of the message on the source chain
    /// @param message      The arbitrary payload to pass to the destination chain receiver
    /// @return legacyMsg   The encoded legacy message
    function encodeLegacyMessage(
        address srcSender,
        address dstReceiver,
        uint64 srcNonce,
        bytes memory message
    )
        internal
        pure
        returns (bytes memory legacyMsg)
    {
        // TODO: implement
    }

    /// @notice Decodes a message received by the Legacy MessageBus.
    /// @param legacyMsg    The encoded legacy message
    /// @return srcSender   The address of the sender on the source chain
    /// @return dstReceiver The address of the receiver on the destination chain
    /// @return srcNonce    The nonce of the message on the source chain
    /// @return message     The arbitrary payload to pass to the destination chain receiver
    function decodeLegacyMessage(bytes calldata legacyMsg)
        internal
        pure
        returns (address srcSender, address dstReceiver, uint64 srcNonce, bytes memory message)
    {
        // TODO: implement
    }
}
