// SPDX-License-Identifier: MIT

pragma solidity 0.8.17;

import {
    Header,
    Message,
    MessageFlag,
    MessageLib,
    MemView,
    MemViewLib
} from "../../../../contracts/libs/memory/Message.sol";

// solhint-disable ordering
/**
 * @notice Exposes Message methods for testing against golang.
 */
contract MessageHarness {
    using MessageLib for bytes;
    using MessageLib for MemView;
    using MemViewLib for bytes;

    // Note: we don't add an empty test() function here, as it currently leads
    // to zero coverage on the corresponding library.

    // ══════════════════════════════════════════════════ GETTERS ══════════════════════════════════════════════════════

    function castToMessage(bytes memory payload) public view returns (bytes memory) {
        // Walkaround to get the forge coverage working on libraries, see
        // https://github.com/foundry-rs/foundry/pull/3128#issuecomment-1241245086
        Message message = MessageLib.castToMessage(payload);
        return message.unwrap().clone();
    }

    function header(bytes memory payload) public pure returns (Header) {
        return payload.castToMessage().header();
    }

    function body(bytes memory payload) public view returns (bytes memory) {
        return payload.castToMessage().body().clone();
    }

    function leaf(bytes memory payload) public pure returns (bytes32) {
        return payload.castToMessage().leaf();
    }

    function isMessage(bytes memory payload) public pure returns (bool) {
        return payload.ref().isMessage();
    }

    // ════════════════════════════════════════════════ FORMATTERS ═════════════════════════════════════════════════════

    function formatMessage(Header header_, bytes memory body_) public pure returns (bytes memory) {
        return MessageLib.formatMessage(header_, body_);
    }
}
