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
        uint32 _origin,
        bytes32 _sender,
        uint32 _nonce,
        uint32 _destination,
        bytes32 _recipient,
        uint32 _optimisticSeconds,
        // tips params
        uint96 _notaryTip,
        uint96 _broadcasterTip,
        uint96 _proverTip,
        uint96 _executorTip,
        bytes memory _messageBody
    ) public pure returns (bytes memory) {
        bytes memory _tips = TipsLib.formatTips(
            _notaryTip,
            _broadcasterTip,
            _proverTip,
            _executorTip
        );

        bytes memory _header = HeaderLib.formatHeader(
            _origin,
            _sender,
            _nonce,
            _destination,
            _recipient,
            _optimisticSeconds
        );
        return formatMessage(_header, _tips, _messageBody);
    }

    function formatMessage(
        uint32 _origin,
        bytes32 _sender,
        uint32 _nonce,
        uint32 _destination,
        bytes32 _recipient,
        uint32 _optimisticSeconds,
        bytes memory _tips,
        bytes memory _messageBody
    ) public pure returns (bytes memory) {
        return
            MessageLib.formatMessage(
                _origin,
                _sender,
                _nonce,
                _destination,
                _recipient,
                _optimisticSeconds,
                _tips,
                _messageBody
            );
    }

    function formatMessage(
        bytes memory _header,
        bytes memory _tips,
        bytes memory _messageBody
    ) public pure returns (bytes memory) {
        return MessageLib.formatMessage(_header, _tips, _messageBody);
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
