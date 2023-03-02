// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;
// ═════════════════════════════ INTERNAL IMPORTS ══════════════════════════════
import { IMessageRecipient } from "../interfaces/IMessageRecipient.sol";
import { InterfaceOrigin } from "../interfaces/InterfaceOrigin.sol";

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

    constructor(address _origin, address _destination) {
        origin = _origin;
        destination = _destination;
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                          EXTERNAL FUNCTIONS                          ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @notice  Handles an incoming message.
     * @dev     Can only be called by chain's Destination.
     *          Message can only be sent from a trusted sender on the origin chain.
     * @param _origin           Domain of the remote chain, where message originated
     * @param _nonce            Unique identifier for the message from origin to destination chain
     * @param _sender           Sender of the message on the origin chain
     * @param _rootSubmittedAt  Time when merkle root (used for proving this message) was submitted
     * @param _message          The message
     */
    function handle(
        uint32 _origin,
        uint32 _nonce,
        bytes32 _sender,
        uint256 _rootSubmittedAt,
        bytes memory _message
    ) external {
        require(msg.sender == destination, "BasicClient: !destination");
        require(
            _sender == trustedSender(_origin) && _sender != bytes32(0),
            "BasicClient: !trustedSender"
        );
        /// @dev root timestamp wasn't checked => potentially unsafe
        /// No need to pass both _origin and _sender: _sender == trustedSender(_origin)
        _handleUnsafe(_origin, _nonce, _rootSubmittedAt, _message);
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
    function trustedSender(uint32 _destination) public view virtual returns (bytes32);

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
    function _handleUnsafe(
        uint32 _origin,
        uint32 _nonce,
        uint256 _rootSubmittedAt,
        bytes memory _message
    ) internal virtual;

    /**
     * @dev Sends a message to given destination chain.
     * @param _destination          Domain of the destination chain
     * @param _optimisticSeconds    Optimistic period for message execution on destination chain
     * @param _tips                 Payload with information about paid tips
     * @param _message              The message
     */
    function _send(
        uint32 _destination,
        uint32 _optimisticSeconds,
        bytes memory _tips,
        bytes memory _message
    ) internal {
        bytes32 recipient = trustedSender(_destination);
        require(recipient != bytes32(0), "BasicClient: !recipient");
        InterfaceOrigin(origin).dispatch{ value: msg.value }(
            _destination,
            recipient,
            _optimisticSeconds,
            _tips,
            _message
        );
    }
}
