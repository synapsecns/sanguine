// SPDX-License-Identifier: MIT

pragma solidity 0.8.13;

import { Message } from "../../contracts/libs/Message.sol";

contract MessageHarness {
    using Message for bytes29;

    function formatMessage(
        uint32 _originDomain,
        bytes32 _sender,
        uint32 _nonce,
        uint32 _destinationDomain,
        bytes32 _recipient,
        bytes memory _messageBody
    ) public pure returns (bytes memory) {
        return
            Message.formatMessage(
                _originDomain,
                _sender,
                _nonce,
                _destinationDomain,
                _recipient,
                _messageBody
            );
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
        bytes memory _body
    ) public pure returns (bytes32) {
        return Message.messageHash(_origin, _sender, _nonce, _destination, _recipient, _body);
    }

    /// @notice Returns message's origin field
    function origin(bytes29 _message) public pure returns (uint32) {
        return _message.origin();
    }

    /// @notice Returns message's sender field
    function sender(bytes29 _message) public pure returns (bytes32) {
        return _message.sender();
    }

    /// @notice Returns message's nonce field
    function nonce(bytes29 _message) public pure returns (uint32) {
        return _message.nonce();
    }

    /// @notice Returns message's destination field
    function destination(bytes29 _message) public pure returns (uint32) {
        return _message.destination();
    }

    /// @notice Returns message's recipient field as bytes32
    function recipient(bytes29 _message) public pure returns (bytes32) {
        return _message.recipient();
    }

    /// @notice Returns message's recipient field as an address
    function recipientAddress(bytes29 _message) public pure returns (address) {
        return _message.recipientAddress();
    }

    /// @notice Returns message's body field as bytes29 (refer to TypedMemView library for details on bytes29 type)
    function body(bytes29 _message) public pure returns (bytes29) {
        return _message.body();
    }

    function leaf(bytes29 _message) public view returns (bytes32) {
        return _message.leaf();
    }
}
