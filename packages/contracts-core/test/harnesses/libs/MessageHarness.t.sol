// SPDX-License-Identifier: MIT

pragma solidity 0.8.17;

import {
    HeaderLib,
    Message,
    MessageLib,
    TipsLib,
    TypedMemView
} from "../../../contracts/libs/Message.sol";

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

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                               GETTERS                                ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function castToMessage(bytes memory payload) public view returns (bytes memory) {
        // Walkaround to get the forge coverage working on libraries, see
        // https://github.com/foundry-rs/foundry/pull/3128#issuecomment-1241245086
        Message message = MessageLib.castToMessage(payload);
        return message.unwrap().clone();
    }

    function header(bytes memory payload) public view returns (bytes memory) {
        return payload.castToMessage().header().unwrap().clone();
    }

    function tips(bytes memory payload) public view returns (bytes memory) {
        return payload.castToMessage().tips().unwrap().clone();
    }

    function body(bytes memory payload) public view returns (bytes memory) {
        return payload.castToMessage().body().clone();
    }

    function version(bytes memory payload) public pure returns (uint16) {
        return payload.castToMessage().version();
    }

    function leaf(bytes memory payload) public pure returns (bytes32) {
        return payload.castToMessage().leaf();
    }

    function isMessage(bytes memory payload) public pure returns (bool) {
        return payload.ref(0).isMessage();
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                              FORMATTERS                              ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function formatMessage(
        uint32 origin,
        bytes32 sender,
        uint32 nonce,
        uint32 destination,
        bytes32 recipient,
        uint32 optimisticSeconds,
        // tips params
        uint96 notaryTip,
        uint96 broadcasterTip,
        uint96 proverTip,
        uint96 executorTip,
        bytes memory messageBody
    ) public pure returns (bytes memory) {
        bytes memory tips_ = TipsLib.formatTips(notaryTip, broadcasterTip, proverTip, executorTip);

        bytes memory header_ = HeaderLib.formatHeader(
            origin,
            sender,
            nonce,
            destination,
            recipient,
            optimisticSeconds
        );
        return formatMessage(header_, tips_, messageBody);
    }

    function formatMessage(
        uint32 origin,
        bytes32 sender,
        uint32 nonce,
        uint32 destination,
        bytes32 recipient,
        uint32 optimisticSeconds,
        bytes memory tips_,
        bytes memory messageBody
    ) public pure returns (bytes memory) {
        return
            MessageLib.formatMessage(
                origin,
                sender,
                nonce,
                destination,
                recipient,
                optimisticSeconds,
                tips_,
                messageBody
            );
    }

    function formatMessage(
        bytes memory header_,
        bytes memory tips_,
        bytes memory messageBody
    ) public pure returns (bytes memory) {
        return MessageLib.formatMessage(header_, tips_, messageBody);
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                           CONSTANT GETTERS                           ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function messageVersion() public pure returns (uint16) {
        return MessageLib.MESSAGE_VERSION;
    }

    function offsetVersion() public pure returns (uint256) {
        return MessageLib.OFFSET_VERSION;
    }

    function offsetHeader() public pure returns (uint256) {
        return MessageLib.OFFSET_HEADER;
    }
}
