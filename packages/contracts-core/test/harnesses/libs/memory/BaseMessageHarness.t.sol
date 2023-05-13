// SPDX-License-Identifier: MIT

pragma solidity 0.8.17;

import {
    BaseMessage,
    BaseMessageLib,
    MemView,
    MemViewLib,
    Request,
    Tips
} from "../../../../contracts/libs/memory/BaseMessage.sol";

// solhint-disable ordering
/**
 * @notice Exposes BaseMessage methods for testing against golang.
 */
contract BaseMessageHarness {
    using BaseMessageLib for bytes;
    using BaseMessageLib for MemView;
    using MemViewLib for bytes;

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

    /// @notice Returns baseMessage's tips field
    function tips(bytes memory payload) public pure returns (uint256) {
        return Tips.unwrap(payload.castToBaseMessage().tips());
    }

    /// @notice Returns baseMessage's request field
    function request(bytes memory payload) public pure returns (uint192) {
        return Request.unwrap(payload.castToBaseMessage().request());
    }

    /// @notice Returns baseMessage's content field
    function content(bytes memory payload) public view returns (bytes memory) {
        return payload.castToBaseMessage().content().clone();
    }

    function isBaseMessage(bytes memory payload) public pure returns (bool) {
        return payload.ref().isBaseMessage();
    }

    // ════════════════════════════════════════════════ FORMATTERS ═════════════════════════════════════════════════════

    function formatBaseMessage(Tips tips_, bytes32 sender_, bytes32 recipient_, Request request_, bytes memory content_)
        public
        pure
        returns (bytes memory)
    {
        return BaseMessageLib.formatBaseMessage(tips_, sender_, recipient_, request_, content_);
    }

    function leaf(bytes memory payload) public pure returns (bytes32) {
        return payload.castToBaseMessage().leaf();
    }

    function bodyLeaf(bytes memory payload) public pure returns (bytes32) {
        return payload.castToBaseMessage().bodyLeaf();
    }
}
