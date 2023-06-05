// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

/// @notice A collection of events emitted by the ExecutionHub contract
abstract contract ExecutionHubEvents {
    /**
     * @notice Emitted when message is executed.
     * @param remoteDomain  Remote domain where message originated
     * @param messageHash   The keccak256 hash of the message that was executed
     * @param success       Whether the message was executed successfully
     */
    event Executed(uint32 indexed remoteDomain, bytes32 indexed messageHash, bool success);

    /**
     * @notice Emitted when message tips are recorded.
     * @param messageHash   The keccak256 hash of the message that was executed
     * @param paddedTips    Padded encoded paid tips information
     */
    event TipsRecorded(bytes32 messageHash, uint256 paddedTips);
}
