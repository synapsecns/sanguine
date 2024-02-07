// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

/// @notice A collection of events emitted by the Origin contract
abstract contract OriginEvents {
    /**
     * @notice Emitted when a new message is sent.
     * @param messageHash   Hash of message; the leaf inserted to the Merkle tree for the message
     * @param nonce         Nonce of sent message (starts from 1)
     * @param destination   Destination domain
     * @param message       Raw bytes of message
     */
    event Sent(bytes32 indexed messageHash, uint32 indexed nonce, uint32 indexed destination, bytes message);

    /**
     * @notice Emitted when a tip withdrawal is completed.
     * @param actor     Actor address
     * @param tip       Tip value, denominated in local domain's wei
     */
    event TipWithdrawalCompleted(address actor, uint256 tip);
}
