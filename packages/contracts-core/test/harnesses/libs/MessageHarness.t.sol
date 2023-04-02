// SPDX-License-Identifier: MIT

pragma solidity 0.8.17;

import {Message, MessageFlag, MessageLib, TypedMemView} from "../../../contracts/libs/Message.sol";

// solhint-disable ordering
/**
 * @notice Exposes Message methods for testing against golang.
 */
contract MessageHarness {
    using MessageLib for bytes;
    using MessageLib for bytes29;
    using TypedMemView for bytes;
    using TypedMemView for bytes29;

    // Note: we don't add an empty test() function here, as it currently leads
    // to zero coverage on the corresponding library.

    // ══════════════════════════════════════════════════ GETTERS ══════════════════════════════════════════════════════

    function castToMessage(bytes memory payload) public view returns (bytes memory) {
        // Walkaround to get the forge coverage working on libraries, see
        // https://github.com/foundry-rs/foundry/pull/3128#issuecomment-1241245086
        Message message = MessageLib.castToMessage(payload);
        return message.unwrap().clone();
    }

    function flag(bytes memory payload) public pure returns (MessageFlag) {
        return payload.castToMessage().flag();
    }

    function header(bytes memory payload) public view returns (bytes memory) {
        return payload.castToMessage().header().unwrap().clone();
    }

    function body(bytes memory payload) public view returns (bytes memory) {
        return payload.castToMessage().body().clone();
    }

    function leaf(bytes memory payload) public pure returns (bytes32) {
        return payload.castToMessage().leaf();
    }

    function isMessage(bytes memory payload) public pure returns (bool) {
        return payload.ref(0).isMessage();
    }

    // ════════════════════════════════════════════════ FORMATTERS ═════════════════════════════════════════════════════

    function formatMessage(MessageFlag flag_, bytes memory header_, bytes memory body_)
        public
        pure
        returns (bytes memory)
    {
        return MessageLib.formatMessage(flag_, header_, body_);
    }
}
