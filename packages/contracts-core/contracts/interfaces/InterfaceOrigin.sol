// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

interface InterfaceOrigin {
    // ═══════════════════════════════════════════════ SEND MESSAGES ═══════════════════════════════════════════════════

    /**
     * @notice Send a message to the recipient located on destination domain.
     * @dev Recipient has to conform to IMessageRecipient interface, otherwise message won't be delivered.
     * Will revert if any of these is true:
     * - `destination` is equal to contract's local domain
     * - `content` length is greater than `MAX_CONTENT_BYTES`
     * - `msg.value` is lower than value of minimum tips for the given message
     * @param destination           Domain of destination chain
     * @param recipient             Address of recipient on destination chain as bytes32
     * @param optimisticPeriod      Optimistic period for message execution on destination chain
     * @param paddedRequest         Padded encoded message execution request on destination chain
     * @param content               Raw bytes content of message
     * @return messageNonce         Nonce of the sent message
     * @return messageHash          Hash of the sent message
     */
    function sendBaseMessage(
        uint32 destination,
        bytes32 recipient,
        uint32 optimisticPeriod,
        uint256 paddedRequest,
        bytes memory content
    ) external payable returns (uint32 messageNonce, bytes32 messageHash);

    /**
     * @notice Send a manager message to the destination domain.
     * @dev This could only be called by AgentManager, which takes care of encoding the calldata payload.
     * Note: (msgOrigin, proofMaturity) security args will be added to payload on the destination chain
     * so that the AgentManager could verify where the Manager Message came from and how mature is the proof.
     * Note: function is not payable, as no tips are required for sending a manager message.
     * Will revert if `destination` is equal to contract's local domain.
     * @param destination           Domain of destination chain
     * @param optimisticPeriod      Optimistic period for message execution on destination chain
     * @param payload               Payload for calling AgentManager on destination chain (with extra security args)
     */
    function sendManagerMessage(uint32 destination, uint32 optimisticPeriod, bytes memory payload)
        external
        returns (uint32 messageNonce, bytes32 messageHash);

    // ════════════════════════════════════════════════ TIPS LOGIC ═════════════════════════════════════════════════════

    /**
     * @notice Withdraws locked base message tips to the recipient.
     * @dev Could only be called by a local AgentManager.
     * @param recipient     Address to withdraw tips to
     * @param amount        Tips value to withdraw
     */
    function withdrawTips(address recipient, uint256 amount) external;

    // ═══════════════════════════════════════════════════ VIEWS ═══════════════════════════════════════════════════════

    /**
     * @notice Returns the minimum tips value for sending a message to a given destination.
     * @dev Using at least `tipsValue` as `msg.value` for `sendBaseMessage()`
     * will guarantee that the message will be accepted.
     * @param destination       Domain of destination chain
     * @param paddedRequest     Padded encoded message execution request on destination chain
     * @param contentLength     The length of the message content
     * @return tipsValue        Minimum tips value for a message to be accepted
     */
    function getMinimumTipsValue(uint32 destination, uint256 paddedRequest, uint256 contentLength)
        external
        view
        returns (uint256 tipsValue);
}
