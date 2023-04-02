// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;
// ═════════════════════════════ INTERNAL IMPORTS ══════════════════════════════

import {IMessageRecipient} from "../interfaces/IMessageRecipient.sol";
import {InterfaceOrigin} from "../interfaces/InterfaceOrigin.sol";

/**
 * @dev Basic implementation of IMessageRecipient interface, to be used as recipient of
 * messages passed by Destination contract.
 * BasicClient could be used as a backbone for cross-chain apps, assuming:
 * - A single app contract per chain (aka trusted sender)
 * - Only app contracts from other chains are able to send messages to app (enforced in BasicClient)
 * - App is responsible for enforcing optimistic period (not enforced in BasicClient)
 *
 * Note: BasicClient is forever stateless, meaning it can be potentially used as a parent
 * for the upgradeable contract without worrying about storage collision.
 */
abstract contract BasicClient is IMessageRecipient {
    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                              IMMUTABLES                              ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    // local chain Origin: used for sending messages
    address public immutable origin;

    // local chain Destination: used for receiving messages
    address public immutable destination;

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                             CONSTRUCTOR                              ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    constructor(address origin_, address destination_) {
        origin = origin_;
        destination = destination_;
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                          EXTERNAL FUNCTIONS                          ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @notice  Handles an incoming message.
     * @dev     Can only be called by chain's Destination.
     *          Message can only be sent from a trusted sender on the origin chain.
     * @param origin_           Domain of the remote chain, where message originated
     * @param nonce             Unique identifier for the message from origin to destination chain
     * @param sender            Sender of the message on the origin chain
     * @param rootSubmittedAt   Time when merkle root (used for proving this message) was submitted
     * @param content           The message content
     */
    function handle(uint32 origin_, uint32 nonce, bytes32 sender, uint256 rootSubmittedAt, bytes memory content)
        external
    {
        require(msg.sender == destination, "BasicClient: !destination");
        require(sender == trustedSender(origin_) && sender != bytes32(0), "BasicClient: !trustedSender");
        /// @dev root timestamp wasn't checked => potentially unsafe
        /// No need to pass both origin and sender: sender == trustedSender(origin)
        _handleUnsafe(origin_, nonce, rootSubmittedAt, content);
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                                VIEWS                                 ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @dev Address of the trusted sender on the destination chain.
     *      The trusted sender will be able to:
     *          (1) send messages to this contract
     *          (2) receive messages from this contract
     */
    function trustedSender(uint32 destination) public view virtual returns (bytes32);

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                          INTERNAL FUNCTIONS                          ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @dev Child contracts should implement the handling logic.
     * At this point it has been confirmed:
     * - Destination called this.handle()
     * - Sender on origin chain is a trusted sender
     * Note: no checks have been done for root timestamp, make sure to enforce optimistic period
     * to protect against executed fake messages on Destination. Hence the "Unsafe" in the name.
     */
    function _handleUnsafe(uint32 origin_, uint32 nonce, uint256 rootSubmittedAt, bytes memory content)
        internal
        virtual;

    /**
     * @dev Sends a message to given destination chain.
     * @param destination_          Domain of the destination chain
     * @param optimisticSeconds     Optimistic period for message execution on destination chain
     * @param tipsPayload           Payload with information about paid tips
     * @param content               The message content
     */
    function _send(uint32 destination_, uint32 optimisticSeconds, bytes memory tipsPayload, bytes memory content)
        internal
    {
        bytes32 recipient = trustedSender(destination_);
        require(recipient != bytes32(0), "BasicClient: !recipient");
        InterfaceOrigin(origin).sendBaseMessage{value: msg.value}(
            destination_, recipient, optimisticSeconds, tipsPayload, content
        );
    }
}
