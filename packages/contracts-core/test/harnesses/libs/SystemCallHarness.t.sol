// SPDX-License-Identifier: MIT

pragma solidity 0.8.17;

import { SystemCall } from "../../../contracts/libs/SystemCall.sol";
import { TypedMemView } from "../../../contracts/libs/TypedMemView.sol";

/**
 * @notice Exposes SystemCall methods for testing against golang.
 */
contract SystemCallHarness {
    using SystemCall for bytes;
    using SystemCall for bytes29;
    using TypedMemView for bytes;
    using TypedMemView for bytes29;

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

    function callPayload(uint40 _type, bytes memory _payload)
        public
        view
        returns (uint40, bytes memory)
    {
        bytes29 _view = _payload.ref(_type).callPayload();
        return (_view.typeOf(), _view.clone());
    }

    function callRecipient(uint40 _type, bytes memory _payload) public pure returns (uint8) {
        return _payload.ref(_type).callRecipient();
    }

    function isSystemCall(bytes memory _payload) public pure returns (bool) {
        return _payload.castToSystemCall().isSystemCall();
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                              FORMATTERS                              ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function formatSystemCall(uint8 _systemRecipient, bytes memory _payload)
        public
        pure
        returns (bytes memory)
    {
        return SystemCall.formatSystemCall(_systemRecipient, _payload);
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                           CONSTANT GETTERS                           ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function systemRouter() public pure returns (bytes32) {
        return SystemCall.SYSTEM_ROUTER;
    }

    function offsetCallRecipient() public pure returns (uint256) {
        return SystemCall.OFFSET_CALL_RECIPIENT;
    }

    function offsetCallPayload() public pure returns (uint256) {
        return SystemCall.OFFSET_CALL_PAYLOAD;
    }
}
