// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

// ══════════════════════════════ LIBRARY IMPORTS ══════════════════════════════
import {
    BaseClientOptimisticPeriod, CallerNotDestination, IncorrectSender, IncorrectRecipient
} from "../libs/Errors.sol";
import {Request} from "../libs/Request.sol";
// ═════════════════════════════ INTERNAL IMPORTS ══════════════════════════════
import {IMessageRecipient} from "../interfaces/IMessageRecipient.sol";
import {InterfaceOrigin} from "../interfaces/InterfaceOrigin.sol";

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
abstract contract BaseClient is IMessageRecipient {
    /// @notice Local chain Origin: used for sending messages
    address public immutable origin;

    /// @notice Local chain Destination: used for receiving messages
    address public immutable destination;

    constructor(address origin_, address destination_) {
        origin = origin_;
        destination = destination_;
    }

    // ═══════════════════════════════════════════ RECEIVE BASE MESSAGES ═══════════════════════════════════════════════

    /// @inheritdoc IMessageRecipient
    function receiveBaseMessage(
        uint32 origin_,
        uint32 nonce,
        bytes32 sender,
        uint256 proofMaturity,
        uint32 version,
        bytes memory content
    ) external payable {
        if (msg.sender != destination) revert CallerNotDestination();
        if (sender != trustedSender(origin_) || sender == 0) revert IncorrectSender();
        if (proofMaturity < optimisticPeriod()) revert BaseClientOptimisticPeriod();
        // All security checks are passed, handle the message content
        _receiveBaseMessage(origin_, nonce, version, content);
    }

    // ═══════════════════════════════════════════════════ VIEWS ═══════════════════════════════════════════════════════

    /**
     * @dev Period of time since the root was submitted to Destination. Once this period is over,
     * root can be used for proving and executing a message though this Client.
     */
    function optimisticPeriod() public view virtual returns (uint32);

    /**
     * @dev Address of the trusted sender on the destination chain.
     *      The trusted sender will be able to:
     *          (1) send messages to this contract
     *          (2) receive messages from this contract
     */
    function trustedSender(uint32 destination) public view virtual returns (bytes32);

    // ══════════════════════════════════════════════ INTERNAL LOGIC ═══════════════════════════════════════════════════

    /**
     * @dev Child contracts should implement the logic for receiving a Base Message.
     * At this point it has been confirmed:
     *  - receiveBaseMessage() was called by Destination (i.e. this is a legit base message)
     *  - Message sender on origin chain is a trusted sender
     *  - Optimistic period for the message have passed
     */
    function _receiveBaseMessage(uint32 origin_, uint32 nonce, uint32 version, bytes memory content) internal virtual;

    /**
     * @dev Sends a message to given destination chain.
     * @param destination_          Domain of the destination chain
     * @param request               Encoded message execution request on destination chain
     * @param content               The message content
     */
    function _sendBaseMessage(uint32 destination_, Request request, bytes memory content) internal {
        bytes32 recipient = trustedSender(destination_);
        if (recipient == 0) revert IncorrectRecipient();
        InterfaceOrigin(origin).sendBaseMessage{value: msg.value}(
            destination_, recipient, optimisticPeriod(), Request.unwrap(request), content
        );
    }
}
