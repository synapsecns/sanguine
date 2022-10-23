// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

// ============ Internal Imports ============
import { BasicClient } from "./BasicClient.sol";

/**
 * @dev Implementation of IMessageRecipient interface, to be used as recipient of
 * messages passed by Destination contract.
 * Client could be used as a backbone for cross-chain apps, assuming:
 * - A single app contract per chain (aka trusted sender)
 * - Only app contracts from other chains are able to send messages to app (enforced in BasicClient)
 * - App has the same optimistic period on all chains (enforced in Client)
 *
 * Note: Client is forever stateless, meaning it can be potentially used as a parent
 * for the upgradeable contract without worrying about storage collision.
 */
abstract contract Client is BasicClient {
    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                             CONSTRUCTOR                              ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/
    // solhint-disable-next-line no-empty-blocks
    constructor(address _origin, address _destination) BasicClient(_origin, _destination) {}

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                                VIEWS                                 ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @dev Period of time since the root was submitted to Destination. Once this period is over,
     * root can be used for proving and executing a message though this Client.
     */
    function optimisticSeconds() public view virtual returns (uint32);

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                          INTERNAL FUNCTIONS                          ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @notice The handling logic.
     * At this point it has been confirmed:
     * - Destination called this.handle()
     * - Sender on origin chain is a trusted sender
     * Note: no checks have been done for root timestamp, make sure to enforce optimistic period
     * to protect against executed fake messages on Destination.
     * @param _origin           Domain of the remote chain, where message originated
     * @param _nonce            Unique identifier for the message from origin to destination chain
     * @param _rootSubmittedAt  Time when merkle root (sed for proving this message) was submitted
     * @param _message          The message
     */
    function _handleUnsafe(
        uint32 _origin,
        uint32 _nonce,
        uint256 _rootSubmittedAt,
        bytes memory _message
    ) internal override {
        // solhint-disable-next-line do-not-rely-on-time
        require(
            block.timestamp >= _rootSubmittedAt + optimisticSeconds(),
            "Client: !optimisticSeconds"
        );
        _handle(_origin, _nonce, _message);
    }

    /**
     * @dev Child contracts should implement the handling logic.
     * At this point it has been confirmed:
     * - Destination called this.handle()
     * - Sender on origin chain is a trusted sender
     * - Optimistic period has passed since merkle root submission
     * Note: this usually means that all security checks have passed
     * and message could be safely executed.
     */
    function _handle(
        uint32 _origin,
        uint32 _nonce,
        bytes memory _message
    ) internal virtual;

    /**
     * @dev Sends a message to given destination chain.
     * @param _destination  Domain of the destination chain
     * @param _message      The message
     */
    function _send(
        uint32 _destination,
        bytes memory _tips,
        bytes memory _message
    ) internal {
        _send(_destination, optimisticSeconds(), _tips, _message);
    }
}
