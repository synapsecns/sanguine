// SPDX-License-Identifier: MIT

pragma solidity 0.8.13;

import { Message } from "../../contracts/libs/Message.sol";
import { Header } from "../../contracts/libs/Header.sol";
import { Tips } from "../../contracts/libs/Tips.sol";
import { TypedMemView } from "../../contracts/libs/TypedMemView.sol";

contract MessageHarness {
    using Message for bytes;
    using Message for bytes29;
    using Header for bytes29;
    using TypedMemView for bytes29;

    function formatMessage(
        uint32 _originDomain,
        bytes32 _sender,
        uint32 _nonce,
        uint32 _destinationDomain,
        bytes32 _recipient,
        uint32 _optimisticSeconds,
        // tips params
        uint96 _updaterTip,
        uint96 _relayerTip,
        uint96 _proverTip,
        uint96 _processorTip,
        bytes memory _messageBody
    ) public pure returns (bytes memory) {
        bytes memory _tips = Tips.formatTips(_updaterTip, _relayerTip, _proverTip, _processorTip);

        bytes memory _header = Header.formatHeader(
            _originDomain,
            _sender,
            _nonce,
            _destinationDomain,
            _recipient,
            _optimisticSeconds
        );
        return Message.formatMessage(_header, _tips, _messageBody);
    }

    /**
     * @notice Returns leaf of formatted message with provided fields.
     * @param _origin Domain of home chain
     * @param _sender Address of sender as bytes32
     * @param _nonce Destination-specific nonce number
     * @param _destination Domain of destination chain
     * @param _recipient Address of recipient on destination chain as bytes32
     * @param _body Raw bytes of message body
     * @return Leaf (hash) of formatted message
     **/
    function messageHash(
        uint32 _origin,
        bytes32 _sender,
        uint32 _nonce,
        uint32 _destination,
        bytes32 _recipient,
        uint32 _optimisticSeconds,
        bytes memory _tips,
        bytes memory _body
    ) public pure returns (bytes32) {
        bytes memory _header = Header.formatHeader(
            _origin,
            _sender,
            _nonce,
            _destination,
            _recipient,
            _optimisticSeconds
        );
        return Message.messageHash(_header, _tips, _body);
    }

    function tips(bytes memory _message) external view returns (bytes memory) {
        return _message.messageView().tips().clone();
    }

    function body(bytes memory _message) external view returns (bytes memory) {
        return _message.messageView().body().clone();
    }

    function origin(bytes memory _message) external pure returns (uint32) {
        return _message.messageView().header().origin();
    }

    function sender(bytes memory _message) external pure returns (bytes32) {
        return _message.messageView().header().sender();
    }

    function nonce(bytes memory _message) external pure returns (uint32) {
        return _message.messageView().header().nonce();
    }

    function destination(bytes memory _message) external pure returns (uint32) {
        return _message.messageView().header().destination();
    }

    function recipient(bytes memory _message) external pure returns (bytes32) {
        return _message.messageView().header().recipient();
    }

    function recipientAddress(bytes memory _message) external pure returns (address) {
        return _message.messageView().header().recipientAddress();
    }

    function optimisticSeconds(bytes memory _message) external pure returns (uint32) {
        return _message.messageView().header().optimisticSeconds();
    }

    function leaf(bytes memory _message) external pure returns (bytes32) {
        return _message.messageView().leaf();
    }

    function messageVersion() public pure returns (uint16) {
        return Message.MESSAGE_VERSION;
    }

    function headerOffset() public pure returns (uint16) {
        return Message.HEADER_OFFSET;
    }
}
