// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

/// @notice A collection of events emitted by the ExecutionHub contract
abstract contract ExecutionHubEvents {
    /**
     * @notice Emitted when message is executed.
     * @param remoteDomain  Remote domain where message originated
     * @param messageHash   The keccak256 hash of the message that was executed
     */
    event Executed(uint32 indexed remoteDomain, bytes32 indexed messageHash);

    /**
     * @notice Emitted when a proof of invalid receipt statement is submitted.
     * @param rcptPayload   Raw payload with the receipt statement
     * @param rcptSignature Notary signature for the receipt statement
     */
    event InvalidReceipt(bytes rcptPayload, bytes rcptSignature);

    /**
     * @notice Emitted when message tips are recorded.
     * @param messageHash   The keccak256 hash of the message that was executed
     * @param tipsPayload   Raw payload with tips paid for the off-chain agents
     */
    event TipsRecorded(bytes32 messageHash, bytes tipsPayload);
}
