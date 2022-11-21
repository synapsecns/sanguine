// SPDX-License-Identifier: MIT

pragma solidity 0.8.17;

import { Message } from "../../../contracts/libs/Message.sol";
import { Header } from "../../../contracts/libs/Header.sol";
import { Tips } from "../../../contracts/libs/Tips.sol";
import { TypedMemView } from "../../../contracts/libs/TypedMemView.sol";

/**
 * @notice Exposes Message methods for testing against golang.
 */
contract MessageHarness {
    using Message for bytes;
    using Message for bytes29;
    using TypedMemView for bytes;
    using TypedMemView for bytes29;

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                               GETTERS                                ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function castToMessage(uint40, bytes memory _payload)
        public
        view
        returns (uint40, bytes memory)
    {
        // Walkaround to get the forge coverage working on libraries, see
        // https://github.com/foundry-rs/foundry/pull/3128#issuecomment-1241245086
        bytes29 _view = Message.castToMessage(_payload);
        return (_view.typeOf(), _view.clone());
    }

    /// @notice Returns message's header field as bytes29 pointer.
    function header(uint40 _type, bytes memory _payload)
        public
        view
        returns (uint40, bytes memory)
    {
        bytes29 _view = _payload.ref(_type).header();
        return (_view.typeOf(), _view.clone());
    }

    /// @notice Returns message's tips field as bytes29 pointer.
    function tips(uint40 _type, bytes memory _payload) public view returns (uint40, bytes memory) {
        bytes29 _view = _payload.ref(_type).tips();
        return (_view.typeOf(), _view.clone());
    }

    /// @notice Returns message's body field as bytes29 pointer.
    function body(uint40 _type, bytes memory _payload) public view returns (uint40, bytes memory) {
        bytes29 _view = _payload.ref(_type).body();
        return (_view.typeOf(), _view.clone());
    }

    /// @notice Returns message's version field.
    function messageVersion(uint40 _type, bytes memory _payload) public pure returns (uint16) {
        return _payload.ref(_type).messageVersion();
    }

    // TODO: Do we need this function in the library? Literally never used.
    function messageHash(
        bytes memory _header,
        bytes memory _tips,
        bytes memory _messageBody
    ) public pure returns (bytes32) {
        return Message.messageHash(_header, _tips, _messageBody);
    }

    function isMessage(bytes memory _payload) public pure returns (bool) {
        return _payload.castToMessage().isMessage();
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
        bytes memory _tips = Tips.formatTips(_notaryTip, _broadcasterTip, _proverTip, _executorTip);

        bytes memory _header = Header.formatHeader(
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
            Message.formatMessage(
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
        return Message.formatMessage(_header, _tips, _messageBody);
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                           CONSTANT GETTERS                           ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function messageVersion() public pure returns (uint16) {
        return Message.MESSAGE_VERSION;
    }

    function offsetVersion() public pure returns (uint256) {
        return Message.OFFSET_VERSION;
    }

    function offsetHeader() public pure returns (uint256) {
        return Message.OFFSET_HEADER;
    }
}
