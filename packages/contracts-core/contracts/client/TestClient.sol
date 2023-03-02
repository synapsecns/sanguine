// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;
// ══════════════════════════════ LIBRARY IMPORTS ══════════════════════════════
import "../libs/Tips.sol";
import "../libs/TypeCasts.sol";
// ═════════════════════════════ INTERNAL IMPORTS ══════════════════════════════
import { IMessageRecipient } from "../interfaces/IMessageRecipient.sol";
import { InterfaceOrigin } from "../interfaces/InterfaceOrigin.sol";

contract TestClient is IMessageRecipient {
    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                              IMMUTABLES                              ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    // local chain Origin: used for sending messages
    address public immutable origin;

    // local chain Destination: used for receiving messages
    address public immutable destination;

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                                EVENTS                                ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    event MessageReceived(
        uint32 origin,
        uint32 nonce,
        bytes32 sender,
        uint256 rootSubmittedAt,
        bytes message
    );

    event MessageSent(
        uint32 destination,
        uint32 nonce,
        bytes32 sender,
        bytes32 recipient,
        bytes message
    );

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                             CONSTRUCTOR                              ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    constructor(address _origin, address _destination) {
        origin = _origin;
        destination = _destination;
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                          RECEIVING MESSAGES                          ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function handle(
        uint32 _origin,
        uint32 _nonce,
        bytes32 _sender,
        uint256 _rootSubmittedAt,
        bytes memory _message
    ) external {
        require(msg.sender == destination, "TestClient: !destination");
        emit MessageReceived(_origin, _nonce, _sender, _rootSubmittedAt, _message);
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                           SENDING MESSAGES                           ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function sendMessage(
        uint32 _destination,
        address _recipient,
        uint32 _optimisticSeconds,
        bytes memory _message
    ) external {
        bytes32 recipient = TypeCasts.addressToBytes32(_recipient);
        bytes memory tips = TipsLib.emptyTips();
        (uint32 nonce, ) = InterfaceOrigin(origin).dispatch(
            _destination,
            recipient,
            _optimisticSeconds,
            tips,
            _message
        );
        emit MessageSent(
            _destination,
            nonce,
            TypeCasts.addressToBytes32(address(this)),
            recipient,
            _message
        );
    }
}
