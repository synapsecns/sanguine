// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import { SynapseTypes } from "./SynapseTypes.sol";
import { TypedMemView } from "./TypedMemView.sol";

library SystemMessage {
    using TypedMemView for bytes;
    using TypedMemView for bytes29;

    // TODO: Looks as if SystemCall covers all the possible use cases
    // for a system message. Determine if that's the case and adjust the code accordingly.

    /**
     * @dev More flag values could be added in the future,
     * e.g. flag indicating another type of system message.
     *
     * - MessageFlag.Call indicates a system contract needs
     * to be called on destination chain.
     */
    enum MessageFlag {
        Call
    }

    /**
     * @dev Custom address, used for sending and receiving system messages.
     *      Origin is supposed to dispatch messages from SystemRouter
     *      as if they were sent by this address.
     *      Destination is supposed to reroute messages for this address to SystemRouter.
     *
     *      Note: all bits except for lower 20 bytes are set to 1.
     *      Note: TypeCasts.bytes32ToAddress(SYSTEM_ROUTER) == address(0)
     */
    bytes32 internal constant SYSTEM_ROUTER = bytes32(type(uint256).max << 160);

    /**
     * @dev SystemMessage memory layout
     * [000 .. 001): messageFlag    uint8   1 bytes
     * [001 .. END]: messageBody    bytes   ? bytes
     */

    uint256 internal constant OFFSET_FLAG = 0;
    uint256 internal constant OFFSET_BODY = 1;

    /**
     * @dev MessageFlag.Call memory layout
     * [000 .. 001): recipient      uint8   1 bytes
     * [001 .. END]: payload        bytes   ? bytes
     */

    uint256 internal constant OFFSET_CALL_RECIPIENT = 0;
    uint256 internal constant OFFSET_CALL_PAYLOAD = 1;

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                              MODIFIERS                               ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    modifier onlyType(bytes29 _view, uint40 _type) {
        _view.assertType(_type);
        _;
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                              FORMATTERS                              ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @notice Returns a properly typed bytes29 pointer for a system message payload.
     */
    function castToSystemMessage(bytes memory _payload) internal pure returns (bytes29) {
        return _payload.ref(SynapseTypes.SYSTEM_MESSAGE);
    }

    /**
     * @notice Returns a formatted System Message payload with provided fields
     * @dev `_messageBody` should be formatted in accordance to `_messageFlag`
     * @param _messageFlag  Flag specifying system message type (see enum MessageFlag)
     * @param _messageBody  Raw bytes of system message body
     * @return Formatted System Message
     **/
    function formatSystemMessage(MessageFlag _messageFlag, bytes memory _messageBody)
        internal
        pure
        returns (bytes memory)
    {
        return abi.encodePacked(uint8(_messageFlag), _messageBody);
    }

    /**
     * @notice Returns a formatted System Call payload with provided fields
     * @param _systemRecipient  System Contract to receive message
     *                          (see ISystemRouter.SystemEntity)
     * @param _payload          Payload for call on destination chain
     * @return Formatted System Call
     **/
    function formatSystemCall(uint8 _systemRecipient, bytes memory _payload)
        internal
        pure
        returns (bytes memory)
    {
        return formatSystemMessage(MessageFlag.Call, abi.encodePacked(_systemRecipient, _payload));
    }

    /**
     * @notice Checks that a payload is a formatted System Message.
     */
    function isSystemMessage(bytes29 _view) internal pure returns (bool) {
        // Message body needs to exist
        if (_view.len() <= OFFSET_BODY) return false;
        // Message flag needs to match an existing enum value
        if (_messageFlagIntValue(_view) > uint8(type(MessageFlag).max)) return false;
        (MessageFlag messageFlag, bytes29 bodyView) = unpackMessage(_view);
        if (messageFlag == MessageFlag.Call) {
            return isSystemCall(bodyView);
        }
        // Unknown messageFlag: should be unreachable
        return false;
    }

    /**
     * @notice Checks that a payload is a formatted System Call.
     */
    function isSystemCall(bytes29 _view) internal pure returns (bool) {
        // Payload needs to exist (system calls are never done via fallback function)
        return _view.len() > OFFSET_CALL_PAYLOAD;
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                        SYSTEM MESSAGE SLICING                        ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @notice Unpacks System Message into flag and body.
     * @dev `bodyView` will have a type reflecting `messageFlag`
     * @param _view         Memory view over System Message
     * @return messageFlag  Flag specifying system message type (see enum MessageFlag)
     * @return bodyView     Memory view over message body
     */
    function unpackMessage(bytes29 _view)
        internal
        pure
        onlyType(_view, SynapseTypes.SYSTEM_MESSAGE)
        returns (MessageFlag messageFlag, bytes29 bodyView)
    {
        messageFlag = MessageFlag(_messageFlagIntValue(_view));
        bodyView = _view.sliceFrom(OFFSET_BODY, _getFlagType(messageFlag));
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                     SYSTEM MESSAGE CALL SLICING                      ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @notice Returns System Call's recipient (see ISystemRouter.SystemEntity).
     */
    function callRecipient(bytes29 _view)
        internal
        pure
        onlyType(_view, SynapseTypes.SYSTEM_MESSAGE_CALL)
        returns (uint8)
    {
        return uint8(_view.indexUint(OFFSET_CALL_RECIPIENT, 1));
    }

    /**
     * @notice Returns System Call's payload.
     */
    function callPayload(bytes29 _view)
        internal
        pure
        onlyType(_view, SynapseTypes.SYSTEM_MESSAGE_CALL)
        returns (bytes29)
    {
        return _view.slice(OFFSET_CALL_PAYLOAD, _view.len() - OFFSET_CALL_PAYLOAD, 0);
    }

    // TODO: system message Adjust slicing when Adjust structure is finalized.

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                          PRIVATE FUNCTIONS                           ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @dev Returns respective uint40 type for message flag. Needed for bytes29 strict typing.
     */
    function _getFlagType(MessageFlag _messageFlag) private pure returns (uint40 _type) {
        if (_messageFlag == MessageFlag.Call) {
            _type = SynapseTypes.SYSTEM_MESSAGE_CALL;
        }
    }

    /**
     * @dev Returns int value of System Message messageFlag.
     *      Needed to prevent overflow when casting to MessageFlag.
     */
    function _messageFlagIntValue(bytes29 _view) private pure returns (uint8 flagIntValue) {
        flagIntValue = uint8(_view.indexUint(OFFSET_FLAG, 1));
    }
}
