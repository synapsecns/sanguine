// SPDX-License-Identifier: MIT

pragma solidity 0.8.17;

import "../../../contracts/libs/ByteString.sol";
import { SynapseTypes } from "../../../contracts/libs/SynapseTypes.sol";
import { SystemCall } from "../../../contracts/libs/SystemCall.sol";
import { TypedMemView } from "../../../contracts/libs/TypedMemView.sol";

/**
 * @notice Exposes SystemCall methods for testing against golang.
 */
contract SystemCallHarness {
    using ByteString for bytes;
    using SystemCall for bytes;
    using SystemCall for bytes29;
    using TypedMemView for bytes;
    using TypedMemView for bytes29;

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                              FORMATTERS                              ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function formatSystemCall(
        uint8 _systemRecipient,
        bytes memory _callData,
        bytes memory _prefix
    ) public view returns (bytes memory) {
        return
            SystemCall.formatSystemCall(
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
            SystemCall.formatAdjustedCallData(_callData.castToCallData(), _prefix.castToRawBytes());
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                               GETTERS                                ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function castToSystemCall(uint40, bytes memory _payload)
        public
        view
        returns (uint40, bytes memory)
    {
        // Walkaround to get the forge coverage working on libraries, see
        // https://github.com/foundry-rs/foundry/pull/3128#issuecomment-1241245086
        bytes29 _view = SystemCall.castToSystemCall(_payload);
        return (_view.typeOf(), _view.clone());
    }

    function callData(uint40 _type, bytes memory _payload) public view returns (bytes memory) {
        return CallData.unwrap(_payload.ref(_type).callData()).clone();
    }

    function callRecipient(uint40 _type, bytes memory _payload) public pure returns (uint8) {
        return _payload.ref(_type).callRecipient();
    }

    function isSystemCall(bytes memory _payload) public pure returns (bool) {
        return _payload.castToSystemCall().isSystemCall();
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                           CONSTANT GETTERS                           ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function systemRouter() public pure returns (bytes32) {
        return SystemCall.SYSTEM_ROUTER;
    }

    function offsetRecipient() public pure returns (uint256) {
        return SystemCall.OFFSET_RECIPIENT;
    }

    function offsetCallData() public pure returns (uint256) {
        return SystemCall.OFFSET_CALLDATA;
    }
}
