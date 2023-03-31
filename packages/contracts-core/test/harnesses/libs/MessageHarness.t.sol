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

    function castToMessage(bytes memory _payload) public view returns (bytes memory) {
        // Walkaround to get the forge coverage working on libraries, see
        // https://github.com/foundry-rs/foundry/pull/3128#issuecomment-1241245086
        Message _msg = MessageLib.castToMessage(_payload);
        return _msg.unwrap().clone();
    }

    function header(bytes memory _payload) public view returns (bytes memory) {
        return _payload.castToMessage().header().unwrap().clone();
    }

    function tips(bytes memory _payload) public view returns (bytes memory) {
        return _payload.castToMessage().tips().unwrap().clone();
    }

    function body(bytes memory _payload) public view returns (bytes memory) {
        return _payload.castToMessage().body().clone();
    }

    function version(bytes memory _payload) public pure returns (uint16) {
        return _payload.castToMessage().version();
    }

    function leaf(bytes memory _payload) public pure returns (bytes32) {
        return _payload.castToMessage().leaf();
    }

    function isMessage(bytes memory _payload) public pure returns (bool) {
        return _payload.ref(0).isMessage();
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
        uint96 _notaryTip,
        uint96 _broadcasterTip,
        uint96 _proverTip,
        uint96 _executorTip,
        bytes memory messageBody
    ) public pure returns (bytes memory) {
        bytes memory tips = TipsLib.formatTips(
            _notaryTip,
            _broadcasterTip,
            _proverTip,
            _executorTip
        );

        bytes memory _header = HeaderLib.formatHeader(
            origin,
            sender,
            nonce,
            destination,
            recipient,
            optimisticSeconds
        );
        return formatMessage(_header, tips, messageBody);
    }

    function formatMessage(
        uint32 origin,
        bytes32 sender,
        uint32 nonce,
        uint32 destination,
        bytes32 recipient,
        uint32 optimisticSeconds,
        bytes memory tips,
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
                tips,
                messageBody
            );
    }

    function formatMessage(
        bytes memory _header,
        bytes memory tips,
        bytes memory messageBody
    ) public pure returns (bytes memory) {
        return MessageLib.formatMessage(_header, tips, messageBody);
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
