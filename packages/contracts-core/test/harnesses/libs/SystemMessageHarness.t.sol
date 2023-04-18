// SPDX-License-Identifier: MIT

pragma solidity 0.8.17;

import {
    SystemEntity,
    SystemMessage,
    SystemMessageLib,
    MemView,
    MemViewLib
} from "../../../contracts/libs/SystemMessage.sol";

// solhint-disable ordering
/**
 * @notice Exposes SystemMessageLib methods for testing against golang.
 */
contract SystemMessageHarness {
    using SystemMessageLib for bytes;
    using SystemMessageLib for MemView;
    using MemViewLib for bytes;

    // Note: we don't add an empty test() function here, as it currently leads
    // to zero coverage on the corresponding library.

    // ════════════════════════════════════════════════ FORMATTERS ═════════════════════════════════════════════════════

    function formatSystemMessage(SystemEntity sender_, SystemEntity recipient_, bytes memory callData_)
        public
        pure
        returns (bytes memory)
    {
        return SystemMessageLib.formatSystemMessage(sender_, recipient_, callData_);
    }

    // ══════════════════════════════════════════════════ GETTERS ══════════════════════════════════════════════════════

    function castToSystemMessage(bytes memory payload) public view returns (bytes memory) {
        // Walkaround to get the forge coverage working on libraries, see
        // https://github.com/foundry-rs/foundry/pull/3128#issuecomment-1241245086
        SystemMessage sm = SystemMessageLib.castToSystemMessage(payload);
        return sm.unwrap().clone();
    }

    function callData(bytes memory payload) public view returns (bytes memory) {
        return payload.castToSystemMessage().callData().unwrap().clone();
    }

    function sender(bytes memory payload) public pure returns (SystemEntity) {
        return payload.castToSystemMessage().sender();
    }

    function recipient(bytes memory payload) public pure returns (SystemEntity) {
        return payload.castToSystemMessage().recipient();
    }

    function isSystemMessage(bytes memory payload) public pure returns (bool) {
        return payload.ref().isSystemMessage();
    }
}
