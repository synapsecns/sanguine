// SPDX-License-Identifier: MIT
pragma solidity 0.8.13;

// ============ Internal Imports ============
import { IMessageRecipient } from "../interfaces/IMessageRecipient.sol";
import { Origin } from "../Origin.sol";

/// @dev Stateless contract, that can be potentially used as a parent
/// for the upgradeable contract.
abstract contract Client is IMessageRecipient {
    // ============ Immutable Variables ============

    // local chain Origin: used for sending messages
    address public immutable origin;

    // local chain Destination: used for receiving messages
    address public immutable destination;

    // ============ Constructor ============

    constructor(address _origin, address _destination) {
        origin = _origin;
        destination = _destination;
    }

    /**
     * @notice          Handles an incoming message.
     * @dev             Can only be called by chain's Destination.
     *                  Can only be sent from a trusted sender on the remote chain.
     * @param _origin   Domain of the remote chain, where message originated
     * @param _nonce    Unique identifier for the message from origin to destination chain
     * @param _sender   Sender of the message on the origin chain
     * @param _message  The message
     */
    function handle(
        uint32 _origin,
        uint32 _nonce,
        bytes32 _sender,
        uint256 _rootTimestamp,
        bytes memory _message
    ) external {
        require(msg.sender == destination, "Client: !destination");
        require(
            _sender == trustedSender(_origin) && _sender != bytes32(0),
            "Client: !trustedSender"
        );
        // solhint-disable-next-line do-not-rely-on-time
        require(
            block.timestamp >= _rootTimestamp + optimisticSeconds(),
            "Client: !optimisticSeconds"
        );
        // No need to pass both _origin and _sender: _sender == trustedSender(_origin)
        _handle(_origin, _nonce, _message);
    }

    // ============ Virtual Functions  ============

    /// @dev Internal logic for handling the message, assuming all security checks are passed
    function _handle(
        uint32 _origin,
        uint32 _nonce,
        bytes memory _message
    ) internal virtual;

    /**
     * @dev                 Sends a message to given destination chain.
     * @param _destination  Domain of the destination chain
     * @param _message      The message
     */
    function _send(
        uint32 _destination,
        bytes memory _tips,
        bytes memory _message
    ) internal {
        bytes32 recipient = trustedSender(_destination);
        require(recipient != bytes32(0), "Client: !recipient");
        Origin(origin).dispatch{ value: msg.value }(
            _destination,
            recipient,
            optimisticSeconds(),
            _tips,
            _message
        );
    }

    /// @dev Period of time since the root was submitted to Mirror. Once this period is over,
    /// root can be used for proving and executing a message though this Client.
    function optimisticSeconds() public view virtual returns (uint32);

    /**
     * @dev Address of the trusted sender on the destination chain.
     *      The trusted sender will be able to:
     *          (1) send messages to this contract
     *          (2) receive messages from this contract
     */
    function trustedSender(uint32 _destination) public view virtual returns (bytes32);
}
