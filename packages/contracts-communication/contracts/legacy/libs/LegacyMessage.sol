// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {MathLib} from "../../libs/Math.sol";

library LegacyMessageLib {
    using MathLib for uint256;

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
        return abi.encode(srcSender, dstReceiver, srcNonce, message);
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
        return abi.decode(legacyMsg, (address, address, uint64, bytes));
    }

    /// @notice Returns the length of encoded legacy message, that has message field of given length.
    /// @param messageLen   The length of the message field
    function payloadSize(uint256 messageLen) internal pure returns (uint256) {
        // 4 fields * 32 bytes (3 values for static, 1 offset for dynamic) + 1 * 32 bytes (message length)
        // message is dynamic field, which is padded up to 32 bytes
        return 160 + messageLen.roundUpToWord();
    }
}
