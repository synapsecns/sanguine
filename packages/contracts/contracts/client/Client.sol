// SPDX-License-Identifier: MIT
pragma solidity 0.8.13;

// ============ Internal Imports ============
import { IMessageRecipient } from "../interfaces/IMessageRecipient.sol";
import { Home } from "../Home.sol";

/// @dev Stateless contract, that can be potentially used as a parent
/// for the upgradeable contract.
abstract contract Client is IMessageRecipient {
    // ============ Immutable Variables ============

    // local chain Home: used for sending messages
    address public immutable home;

    // local chain ReplicaManager: used for receiving messages
    address public immutable replicaManager;

    // ============ Constructor ============

    constructor(address _home, address _replicaManager) {
        home = _home;
        replicaManager = _replicaManager;
    }

    /**
     * @notice          Handles an incoming message.
     * @dev             Can only be called by chain's ReplicaManager.
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
        require(msg.sender == replicaManager, "Client: !replica");
        require(
            _sender == trustedSender(_origin) && _sender != bytes32(0),
            "Client: !trustedSender"
        );
        // solhint-disable-next-line do-not-rely-on-time
        require(
            block.timestamp >= _rootTimestamp + optimisticSeconds(),
            "Client: !optimisticSeconds"
        );
        _handle(_origin, _nonce, _sender, _message);
    }

    // ============ Virtual Functions  ============

    /// @dev Internal logic for handling the message, assuming all security checks are passed
    function _handle(
        uint32 _origin,
        uint32 _nonce,
        bytes32 _sender,
        bytes memory _message
    ) internal virtual;

    /**
     * @dev                 Sends a message to given destination chain.
     * @param _destination  Domain of the destination chain
     * @param _message      The message
     */
    function _send(uint32 _destination, bytes memory _message) internal {
        bytes32 recipient = trustedSender(_destination);
        require(recipient != bytes32(0), "Client: !recipient");
        Home(home).dispatch(_destination, recipient, optimisticSeconds(), _message);
    }

    /// @dev Period of time since the root was submitted to Replica. Once this period is over,
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
