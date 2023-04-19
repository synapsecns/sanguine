// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {ByteString, CallData} from "./ByteString.sol";
import {SystemEntity} from "./Structures.sol";
import {MemView, MemViewLib} from "./MemView.sol";

/// SystemMessage is a memory view over the message with instructions for a system call.
type SystemMessage is uint256;

using SystemMessageLib for SystemMessage global;

/// SystemMessage structure represents a message sent by one of the messaging contracts
/// with instructions for a system call.
/// > See SystemRouter.sol for clarifications about the remote system calls.
/// Note: calldata does not include the security arguments, these are added by SystemRouter on destination chain.
///
/// # SystemMessage memory layout
///
/// | Position   | Field     | Type  | Bytes | Description                                              |
/// | ---------- | --------- | ----- | ----- | -------------------------------------------------------- |
/// | [000..001) | sender    | uint8 | 1     | SystemEntity that sent the message on origin chain       |
/// | [001..002) | recipient | uint8 | 1     | SystemEntity to receive the message on destination chain |
/// | [002..END) | calldata  | bytes | ??    | Raw bytes of payload to call system recipient            |
library SystemMessageLib {
    using MemViewLib for bytes;
    using ByteString for MemView;

    /// @dev The variables below are not supposed to be used outside of the library directly.
    uint256 private constant OFFSET_SENDER = 0;
    uint256 private constant OFFSET_RECIPIENT = 1;
    uint256 private constant OFFSET_CALLDATA = 2;

    // ══════════════════════════════════════════════ SYSTEM MESSAGE ═══════════════════════════════════════════════════

    /**
     * @notice Returns a formatted SystemMessage payload with provided fields.
     * @param sender_           System Contract that sent a system message
     * @param recipient_        System Contract to receive a system message
     * @param callData_         Raw bytes with calldata payload
     * @return Formatted SystemMessage payload.
     */
    function formatSystemMessage(SystemEntity sender_, SystemEntity recipient_, bytes memory callData_)
        internal
        pure
        returns (bytes memory)
    {
        return abi.encodePacked(sender_, recipient_, callData_);
    }

    /**
     * @notice Returns a SystemMessage view over for the given payload.
     * @dev Will revert if the payload is not a system message.
     */
    function castToSystemMessage(bytes memory payload) internal pure returns (SystemMessage) {
        return castToSystemMessage(payload.ref());
    }

    /**
     * @notice Casts a memory view to a SystemMessage view.
     * @dev Will revert if the payload is not a system message.
     */
    function castToSystemMessage(MemView memView) internal pure returns (SystemMessage) {
        require(isSystemMessage(memView), "Not a system message");
        return SystemMessage.wrap(MemView.unwrap(memView));
    }

    /**
     * @notice Checks that a payload is a formatted System Message.
     */
    function isSystemMessage(MemView memView) internal pure returns (bool) {
        // Check if sender and recipient exist in the payload
        if (memView.len() < OFFSET_CALLDATA) return false;
        // Check if sender fits into SystemEntity enum
        if (_sender(memView) > uint8(type(SystemEntity).max)) return false;
        // Check if recipient fits into SystemEntity enum
        if (_recipient(memView) > uint8(type(SystemEntity).max)) return false;
        // Check that "calldata" field is a proper calldata
        return _callData(memView).isCallData();
    }

    /// @notice Convenience shortcut for unwrapping a view.
    function unwrap(SystemMessage systemMessage) internal pure returns (MemView) {
        return MemView.wrap(SystemMessage.unwrap(systemMessage));
    }

    // ══════════════════════════════════════════ SYSTEM MESSAGE SLICING ═══════════════════════════════════════════════

    /// @notice Returns system message's recipient.
    function sender(SystemMessage systemMessage) internal pure returns (SystemEntity) {
        // We check that sender fits into enum, when payload is wrapped
        // into SystemMessage, so this never reverts
        return SystemEntity(_sender(systemMessage.unwrap()));
    }

    /// @notice Returns system message's recipient.
    function recipient(SystemMessage systemMessage) internal pure returns (SystemEntity) {
        // We check that recipient fits into enum, when payload is wrapped
        // into SystemMessage, so this never reverts
        return SystemEntity(_recipient(systemMessage.unwrap()));
    }

    /// @notice Returns typed memory view over the calldata used in the system message.
    function callData(SystemMessage systemMessage) internal pure returns (CallData) {
        return _callData(systemMessage.unwrap()).castToCallData();
    }

    // ══════════════════════════════════════════════ PRIVATE HELPERS ══════════════════════════════════════════════════

    /// @dev Returns message's sender without checking that it fits into SystemEntity enum.
    function _sender(MemView memView) private pure returns (uint8) {
        return uint8(memView.indexUint({index_: OFFSET_SENDER, bytes_: 1}));
    }

    /// @dev Returns message's recipient without checking that it fits into SystemEntity enum.
    function _recipient(MemView memView) private pure returns (uint8) {
        return uint8(memView.indexUint({index_: OFFSET_RECIPIENT, bytes_: 1}));
    }

    /// @dev Returns an untyped memory view over the calldata field without checking
    /// if the whole payload or the calldata are properly formatted.
    function _callData(MemView memView) private pure returns (MemView) {
        return memView.sliceFrom({index_: OFFSET_CALLDATA});
    }
}
