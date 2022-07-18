// SPDX-License-Identifier: MIT
pragma solidity 0.8.13;

import { TypedMemView } from "../libs/TypedMemView.sol";

library SystemMessage {
    using TypedMemView for bytes29;

    enum SystemMessageType {
        None,
        Call,
        Adjust
    }

    /**
     * @dev Custom address, used for receiving and sending system messages.
     *      Home is supposed to dispatch messages from SystemMessenger as if they were sent by this address.
     *      ReplicaManager is supposed to reroute messages for this address to SystemMessenger.
     *      Note: all bits except for lower 20 bytes are set to 1.
     *      Note: TypeCasts.bytes32ToAddress(SYSTEM_SENDER) = address(0)
     */
    bytes32 internal constant SYSTEM_SENDER = bytes32(type(uint256).max << 160);

    /**
     * @dev SystemMessage memory layout
     * [000 .. 001): messageType    uint8   1 bytes
     * [001 .. END]: messageBody    bytes   ? bytes
     */

    uint256 internal constant OFFSET_BODY = 1;

    /**
     * @dev SystemMessageType.Call memory layout
     * [000 .. 001): recipient      uint8   1 bytes
     * [001 .. END]: payload        bytes   ? bytes
     */

    uint256 internal constant OFFSET_CALL_PAYLOAD = 1;

    // TODO: memory layout + setter/getters for SystemMessageType.Adjust

    modifier onlyType(SystemMessageType _type, bytes29 _view) {
        _view.assertType(uint40(_type));
        _;
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                              FORMATTERS                              ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function formatSystemMessage(SystemMessageType _messageType, bytes memory _messageBody)
        internal
        pure
        returns (bytes memory)
    {
        return abi.encodePacked(uint8(_messageType), _messageBody);
    }

    function formatCall(uint8 _recipientType, bytes memory _payload)
        internal
        pure
        returns (bytes memory)
    {
        return
            formatSystemMessage(SystemMessageType.Call, abi.encodePacked(_recipientType, _payload));
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                        SYSTEM MESSAGE GETTERS                        ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function systemMessage(bytes29 _message)
        internal
        pure
        returns (SystemMessageType _messageType, bytes29 _messageView)
    {
        _messageType = SystemMessageType(_message.indexUint(0, 1));
        _messageView = _message.slice(
            OFFSET_BODY,
            _message.len() - OFFSET_BODY,
            uint40(_messageType)
        );
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                   SYSTEM_MESSAGE_TYPE.CALL GETTERS                   ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function callRecipient(bytes29 _call)
        internal
        pure
        onlyType(SystemMessageType.Call, _call)
        returns (uint8)
    {
        return uint8(_call.indexUint(0, 1));
    }

    function callPayload(bytes29 _call)
        internal
        pure
        onlyType(SystemMessageType.Call, _call)
        returns (bytes29)
    {
        return _call.slice(OFFSET_CALL_PAYLOAD, _call.len() - OFFSET_CALL_PAYLOAD, 0);
    }
}
