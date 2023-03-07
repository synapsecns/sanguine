// SPDX-License-Identifier: MIT

pragma solidity 0.8.17;

import "../../../contracts/libs/Constants.sol";
import "../../../contracts/libs/SystemMessage.sol";

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

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                              FORMATTERS                              ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function formatSystemMessage(
        uint8 _systemRecipient,
        bytes memory _callData,
        bytes memory _prefix
    ) public view returns (bytes memory) {
        return
            SystemMessageLib.formatSystemMessage(
                _systemRecipient,
                _callData.castToCallData(),
                _prefix.castToRawBytes()
            );
    }

    function formatAdjustedCallData(bytes memory _callData, bytes memory _prefix)
        public
        view
        returns (bytes memory)
    {
        return
            SystemMessageLib.formatAdjustedCallData(
                _callData.castToCallData(),
                _prefix.castToRawBytes()
            );
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                               GETTERS                                ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function castToSystemMessage(bytes memory _payload) public view returns (bytes memory) {
        // Walkaround to get the forge coverage working on libraries, see
        // https://github.com/foundry-rs/foundry/pull/3128#issuecomment-1241245086
        SystemMessage sm = SystemMessageLib.castToSystemMessage(_payload);
        return sm.unwrap().clone();
    }

    function callData(bytes memory _payload) public view returns (bytes memory) {
        return _payload.castToSystemMessage().callData().unwrap().clone();
    }

    function callRecipient(bytes memory _payload) public pure returns (uint8) {
        return _payload.castToSystemMessage().callRecipient();
    }

    function isSystemMessage(bytes memory _payload) public pure returns (bool) {
        return _payload.ref(0).isSystemMessage();
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
