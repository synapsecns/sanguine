// SPDX-License-Identifier: MIT

pragma solidity 0.8.13;

import { SystemMessage } from "../../../contracts/libs/SystemMessage.sol";

/**
 * @notice Exposes Attestation methods for testing against golang.
 */
contract SystemMessageHarness {
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
