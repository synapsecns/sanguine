// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {ByteString, CallData} from "./ByteString.sol";
import {SystemEntity} from "./Structures.sol";
import {TypedMemView} from "./TypedMemView.sol";

/// @dev SystemMessage is a memory view over the message with instructions for a system call.
type SystemMessage is bytes29;

/// @dev Attach library functions to SystemMessage
using SystemMessageLib for SystemMessage global;

library SystemMessageLib {
    using ByteString for bytes;
    using ByteString for bytes29;
    using TypedMemView for bytes29;

    /**
     * @dev SystemMessage memory layout
     * Note: calldata does not include the security arguments, these are added by SystemRouter on destination chain.
     * [000 .. 001): sender         uint8   1 byte      SystemEntity that sent the message on origin chain
     * [001 .. 002): recipient      uint8   1 byte      SystemEntity to receive the message on destination chain
     * [002 .. END]: calldata       bytes   ? bytes     Raw bytes of payload to call system recipient
     *
     * The variables below are not supposed to be used outside of the library directly.
     */

    uint256 private constant OFFSET_SENDER = 0;
    uint256 private constant OFFSET_RECIPIENT = 1;
    uint256 private constant OFFSET_CALLDATA = 2;

    // ══════════════════════════════════════════════ SYSTEM MESSAGE ═══════════════════════════════════════════════════

    /**
     * @notice Returns a formatted SystemMessage payload with provided fields.
     * @param sender_           System Contract that sent receive message
     * @param recipient_        System Contract to receive message
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
        return castToSystemMessage(payload.castToRawBytes());
    }

    /**
     * @notice Casts a memory view to a SystemMessage view.
     * @dev Will revert if the payload is not a system message.
     */
    function castToSystemMessage(bytes29 view_) internal pure returns (SystemMessage) {
        require(isSystemMessage(view_), "Not a system message");
        return SystemMessage.wrap(view_);
    }

    /**
     * @notice Checks that a payload is a formatted System Message.
     */
    function isSystemMessage(bytes29 view_) internal pure returns (bool) {
        // Check if sender and recipient exist in the payload
        if (view_.len() < OFFSET_CALLDATA) return false;
        // Check if sender fits into SystemEntity enum
        if (_sender(view_) > uint8(type(SystemEntity).max)) return false;
        // Check if recipient fits into SystemEntity enum
        if (_recipient(view_) > uint8(type(SystemEntity).max)) return false;
        bytes29 callDataView = _callData(view_);
        // Check that "calldata" field is a proper calldata
        return callDataView.isCallData();
    }

    /// @notice Convenience shortcut for unwrapping a view.
    function unwrap(SystemMessage systemMessage) internal pure returns (bytes29) {
        return SystemMessage.unwrap(systemMessage);
    }

    // ══════════════════════════════════════════ SYSTEM MESSAGE SLICING ═══════════════════════════════════════════════

    /// @notice Returns system message's recipient.
    function sender(SystemMessage systemMessage) internal pure returns (SystemEntity) {
        // Get the underlying memory view
        bytes29 view_ = systemMessage.unwrap();
        // We check that sender fits into enum, when payload is wrapped
        // into SystemMessage, so this never reverts
        return SystemEntity(_sender(view_));
    }

    /// @notice Returns system message's recipient.
    function recipient(SystemMessage systemMessage) internal pure returns (SystemEntity) {
        // Get the underlying memory view
        bytes29 view_ = systemMessage.unwrap();
        // We check that recipient fits into enum, when payload is wrapped
        // into SystemMessage, so this never reverts
        return SystemEntity(_recipient(view_));
    }

    /// @notice Returns typed memory view over the calldata used in the system message.
    function callData(SystemMessage systemMessage) internal pure returns (CallData) {
        // Get the underlying memory view
        bytes29 view_ = systemMessage.unwrap();
        return _callData(view_).castToCallData();
    }

    // ══════════════════════════════════════════════ PRIVATE HELPERS ══════════════════════════════════════════════════

    /// @dev Returns message's sender without checking that it fits into SystemEntity enum.
    function _sender(bytes29 view_) private pure returns (uint8) {
        return uint8(view_.indexUint({index_: OFFSET_SENDER, bytes_: 1}));
    }

    /// @dev Returns message's recipient without checking that it fits into SystemEntity enum.
    function _recipient(bytes29 view_) private pure returns (uint8) {
        return uint8(view_.indexUint({index_: OFFSET_RECIPIENT, bytes_: 1}));
    }

    /// @dev Returns an untyped memory view over the calldata field without checking
    /// if the whole payload or the calldata are properly formatted.
    function _callData(bytes29 view_) private pure returns (bytes29) {
        return view_.sliceFrom({index_: OFFSET_CALLDATA, newType: 0});
    }
}
