// SPDX-License-Identifier: MIT

pragma solidity 0.8.17;

import {BaseMessage, BaseMessageLib, TypedMemView} from "../../../contracts/libs/BaseMessage.sol";

// solhint-disable ordering
/**
 * @notice Exposes BaseMessage methods for testing against golang.
 */
contract BaseMessageHarness {
    using BaseMessageLib for bytes;
    using BaseMessageLib for bytes29;
    using TypedMemView for bytes;
    using TypedMemView for bytes29;

    // Note: we don't add an empty test() function here, as it currently leads
    // to zero coverage on the corresponding library.

    // ══════════════════════════════════════════════════ GETTERS ══════════════════════════════════════════════════════

    function castToBaseMessage(bytes memory payload) public view returns (bytes memory) {
        // Walkaround to get the forge coverage working on libraries, see
        // https://github.com/foundry-rs/foundry/pull/3128#issuecomment-1241245086
        BaseMessage baseMessage = BaseMessageLib.castToBaseMessage(payload);
        return baseMessage.unwrap().clone();
    }

    /// @notice Returns baseMessage's sender field
    function sender(bytes memory payload) public pure returns (bytes32) {
        return payload.castToBaseMessage().sender();
    }

    /// @notice Returns baseMessage's recipient field as bytes32
    function recipient(bytes memory payload) public pure returns (bytes32) {
        return payload.castToBaseMessage().recipient();
    }

    /// @notice Returns baseMessage's optimistic seconds field
    function tips(bytes memory payload) public view returns (bytes memory) {
        return payload.castToBaseMessage().tips().unwrap().clone();
    }

    /// @notice Returns baseMessage's recipient field as an address
    function content(bytes memory payload) public view returns (bytes memory) {
        return payload.castToBaseMessage().content().clone();
    }

    function isBaseMessage(bytes memory payload) public pure returns (bool) {
        return payload.ref(0).isBaseMessage();
    }

    // ════════════════════════════════════════════════ FORMATTERS ═════════════════════════════════════════════════════

    function formatBaseMessage(bytes32 sender_, bytes32 recipient_, bytes memory tipsPayload, bytes memory content_)
        public
        pure
        returns (bytes memory)
    {
        return BaseMessageLib.formatBaseMessage(sender_, recipient_, tipsPayload, content_);
    }
}
