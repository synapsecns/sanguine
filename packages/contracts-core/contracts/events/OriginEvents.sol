// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

/// @notice A collection of events emitted by the Origin contract
abstract contract OriginEvents {
    // Old Event to ensure that go generation works with the existing Agents
    // TODO: remove once agents are updated to handle the new "Sent" event
    event Dispatched(bytes32 indexed messageHash, uint32 indexed nonce, uint32 indexed destination, bytes message);

    /**
     * @notice Emitted when a new message is sent.
     * @param messageHash   Hash of message; the leaf inserted to the Merkle tree for the message
     * @param nonce         Nonce of sent message (starts from 1)
     * @param destination   Destination domain
     * @param message       Raw bytes of message
     */
    event Sent(bytes32 indexed messageHash, uint32 indexed nonce, uint32 indexed destination, bytes message);
}
