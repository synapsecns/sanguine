// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

// ══════════════════════════════ LIBRARY IMPORTS ══════════════════════════════
import {BaseClientOptimisticPeriod, IncorrectSender} from "../libs/Errors.sol";
// ═════════════════════════════ INTERNAL IMPORTS ══════════════════════════════
import {MessageRecipient} from "./MessageRecipient.sol";

/**
 * @dev Implementation of IMessageRecipient interface, to be used as the recipient of
 * base messages passed by the Destination contract.
 * BaseClient could be used as a backbone for cross-chain apps:
 * - A single BaseClient contract per chain (aka trusted sender)
 * - Only BaseClient instances from other chains are able to send messages to this contract
 * - BaseClient enforces a common optimistic period for all types of messages
 * Note: BaseClient is forever stateless, meaning it can be potentially used as a parent
 * for the upgradeable contract without worrying about storage collision.
 */
abstract contract BaseClient is MessageRecipient {
    // solhint-disable-next-line no-empty-blocks
    constructor(address origin_, address destination_) MessageRecipient(origin_, destination_) {}

    // ═══════════════════════════════════════════════════ VIEWS ═══════════════════════════════════════════════════════

    /**
     * @notice Period of time since the root was submitted to Destination. Once this period is over,
     * root can be used for proving and executing a message though this Client.
     */
    function optimisticPeriod() public view virtual returns (uint32);

    /**
     * @notice Address of the trusted sender on the destination chain. The trusted sender will be able to:
     *  - Send messages to this contract from other chains.
     *  - Receive messages from this contract on other chains.
     */
    function trustedSender(uint32 destination) public view virtual returns (bytes32);

    // ══════════════════════════════════════════════ INTERNAL LOGIC ═══════════════════════════════════════════════════

    /**
     * @dev Child contracts should implement the logic for receiving a Base Message in an "unsafe way".
     * Following checks HAVE been performed:
     *  - receiveBaseMessage() was called by Destination (i.e. this is a legit base message).
     *  - Nonce is not zero.
     *  - Message sender on origin chain is not a zero address.
     *  - Proof maturity is not zero.
     * Following checks HAVE NOT been performed (thus "unsafe"):
     *  - Message sender on origin chain could be anything non-zero at this point.
     *  - Proof maturity could be anything non-zero at this point.
     */
    function _receiveBaseMessageUnsafe(
        uint32 origin_,
        uint32 nonce,
        bytes32 sender,
        uint256 proofMaturity,
        uint32 version,
        bytes memory content
    ) internal override {
        if (sender != trustedSender(origin_)) revert IncorrectSender();
        if (proofMaturity < optimisticPeriod()) revert BaseClientOptimisticPeriod();
        _receiveBaseMessage(origin_, nonce, version, content);
    }

    /**
     * @dev Child contracts should implement the logic for receiving a Base Message.
     * Following checks HAVE been performed:
     *  - receiveBaseMessage() was called by Destination (i.e. this is a legit base message).
     *  - Nonce is not zero.
     *  - Message sender on origin chain is a trusted sender (and is not zero).
     *  - Optimistic period for the message have passed (and is not zero).
     */
    function _receiveBaseMessage(uint32 origin_, uint32 nonce, uint32 version, bytes memory content) internal virtual;

    /**
     * @dev Sends a message to given destination chain. Full `msg.value` is used to pay for the message tips.
     * `_getMinimumTipsValue()` could be used to calculate the minimum required tips value, and should be also
     * exposed as a public view function to estimate the tips value before sending a message off-chain.
     * This function is not exposed in BaseClient, as the message encoding is implemented by the child contract.
     * > Will revert if the trusted sender is not set for the destination domain.
     * @param destination_          Domain of the destination chain
     * @param tipsValue             Tips to be paid for sending the message
     * @param request               Encoded message execution request on destination chain
     * @param content               The message content
     */
    function _sendBaseMessage(
        uint32 destination_,
        uint256 tipsValue,
        MessageRequest memory request,
        bytes memory content
    ) internal returns (uint32 messageNonce, bytes32 messageHash) {
        // Send message to the trusted sender on destination chain with the defined optimistic period.
        // Note: this will revert if the trusted sender is not set for the destination domain.
        return _sendBaseMessage({
            destination_: destination_,
            recipient: trustedSender(destination_),
            optimisticPeriod: optimisticPeriod(),
            tipsValue: tipsValue,
            request: request,
            content: content
        });
    }
}
