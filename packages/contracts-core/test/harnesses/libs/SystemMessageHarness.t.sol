// SPDX-License-Identifier: MIT

pragma solidity 0.8.17;

import {SYSTEM_ROUTER} from "../../../contracts/libs/Constants.sol";
import {
    ByteString,
    CallData,
    SystemMessage,
    SystemMessageLib,
    TypedMemView
} from "../../../contracts/libs/SystemMessage.sol";

/**
 * @notice Exposes SystemMessageLib methods for testing against golang.
 */
contract SystemMessageHarness {
    using ByteString for bytes;
    using ByteString for CallData;
    using SystemMessageLib for bytes;
    using SystemMessageLib for bytes29;
    using SystemMessageLib for SystemMessage;
    using TypedMemView for bytes;
    using TypedMemView for bytes29;

    // Note: we don't add an empty test() function here, as it currently leads
    // to zero coverage on the corresponding library.

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                              FORMATTERS                              ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function formatSystemMessage(uint8 systemRecipient, bytes memory callData_, bytes memory prefix)
        public
        view
        returns (bytes memory)
    {
        return
            SystemMessageLib.formatSystemMessage(systemRecipient, callData_.castToCallData(), prefix.castToRawBytes());
    }

    function formatAdjustedCallData(bytes memory callData_, bytes memory prefix) public view returns (bytes memory) {
        return SystemMessageLib.formatAdjustedCallData(callData_.castToCallData(), prefix.castToRawBytes());
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                               GETTERS                                ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function castToSystemMessage(bytes memory payload) public view returns (bytes memory) {
        // Walkaround to get the forge coverage working on libraries, see
        // https://github.com/foundry-rs/foundry/pull/3128#issuecomment-1241245086
        SystemMessage sm = SystemMessageLib.castToSystemMessage(payload);
        return sm.unwrap().clone();
    }

    function callData(bytes memory payload) public view returns (bytes memory) {
        return payload.castToSystemMessage().callData().unwrap().clone();
    }

    function callRecipient(bytes memory payload) public pure returns (uint8) {
        return payload.castToSystemMessage().callRecipient();
    }

    function isSystemMessage(bytes memory payload) public pure returns (bool) {
        return payload.ref(0).isSystemMessage();
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                           CONSTANT GETTERS                           ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function systemRouter() public pure returns (bytes32) {
        return SYSTEM_ROUTER;
    }

    function offsetRecipient() public pure returns (uint256) {
        return SystemMessageLib.OFFSET_RECIPIENT;
    }

    function offsetCallData() public pure returns (uint256) {
        return SystemMessageLib.OFFSET_CALLDATA;
    }
}
