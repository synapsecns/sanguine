// SPDX-License-Identifier: MIT

pragma solidity 0.8.17;

import { SystemMessage } from "../../../contracts/libs/SystemMessage.sol";
import { TypedMemView } from "../../../contracts/libs/TypedMemView.sol";

/**
 * @notice Exposes Attestation methods for testing against golang.
 */
contract SystemMessageHarness {
    using SystemMessage for bytes;
    using SystemMessage for bytes29;
    using TypedMemView for bytes;
    using TypedMemView for bytes29;

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                               GETTERS                                ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function castToSystemMessage(uint40, bytes memory _payload)
        public
        view
        returns (uint40, bytes memory)
    {
        // Walkaround to get the forge coverage working on libraries, see
        // https://github.com/foundry-rs/foundry/pull/3128#issuecomment-1241245086
        bytes29 _view = SystemMessage.castToSystemMessage(_payload);
        return (_view.typeOf(), _view.clone());
    }

    function systemMessageBody(uint40 _type, bytes memory _payload)
        public
        view
        returns (uint40, bytes memory)
    {
        (, bytes29 _view) = _payload.ref(_type).unpackMessage();
        return (_view.typeOf(), _view.clone());
    }

    function callPayload(uint40 _type, bytes memory _payload)
        public
        view
        returns (uint40, bytes memory)
    {
        bytes29 _view = _payload.ref(_type).callPayload();
        return (_view.typeOf(), _view.clone());
    }

    function systemMessageFlag(uint40 _type, bytes memory _payload)
        public
        pure
        returns (SystemMessage.MessageFlag flag)
    {
        (flag, ) = _payload.ref(_type).unpackMessage();
    }

    function callRecipient(uint40 _type, bytes memory _payload) public pure returns (uint8) {
        return _payload.ref(_type).callRecipient();
    }

    function isSystemMessage(bytes memory _payload) public pure returns (bool) {
        return _payload.castToSystemMessage().isSystemMessage();
    }

    function isSystemCall(bytes memory _payload) public pure returns (bool) {
        return _payload.ref(0).isSystemCall();
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                              FORMATTERS                              ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function formatSystemMessage(SystemMessage.MessageFlag _messageFlag, bytes memory _messageBody)
        public
        pure
        returns (bytes memory)
    {
        return SystemMessage.formatSystemMessage(_messageFlag, _messageBody);
    }

    function formatSystemCall(uint8 _systemRecipient, bytes memory _payload)
        public
        pure
        returns (bytes memory)
    {
        return SystemMessage.formatSystemCall(_systemRecipient, _payload);
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                           CONSTANT GETTERS                           ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function systemRouter() public pure returns (bytes32) {
        return SystemMessage.SYSTEM_ROUTER;
    }

    function offsetFlag() public pure returns (uint256) {
        return SystemMessage.OFFSET_FLAG;
    }

    function offsetBody() public pure returns (uint256) {
        return SystemMessage.OFFSET_BODY;
    }

    function offsetCallRecipient() public pure returns (uint256) {
        return SystemMessage.OFFSET_CALL_RECIPIENT;
    }

    function offsetCallPayload() public pure returns (uint256) {
        return SystemMessage.OFFSET_CALL_PAYLOAD;
    }
}
