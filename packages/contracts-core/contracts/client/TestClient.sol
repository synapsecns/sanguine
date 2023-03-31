// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;
// ══════════════════════════════ LIBRARY IMPORTS ══════════════════════════════
import { TipsLib } from "../libs/Tips.sol";
import { TypeCasts } from "../libs/TypeCasts.sol";

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

    constructor(address origin, address destination) {
        origin = origin;
        destination = destination;
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                          RECEIVING MESSAGES                          ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function handle(
        uint32 origin,
        uint32 nonce,
        bytes32 sender,
        uint256 rootSubmittedAt,
        bytes memory message
    ) external {
        require(msg.sender == destination, "TestClient: !destination");
        emit MessageReceived(origin, nonce, sender, rootSubmittedAt, message);
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                           SENDING MESSAGES                           ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function sendMessage(
        uint32 destination,
        address recipient,
        uint32 optimisticSeconds,
        bytes memory message
    ) external {
        bytes32 recipient = TypeCasts.addressToBytes32(recipient);
        bytes memory tips = TipsLib.emptyTips();
        (uint32 nonce, ) = InterfaceOrigin(origin).dispatch(
            destination,
            recipient,
            optimisticSeconds,
            tips,
            message
        );
        emit MessageSent(
            destination,
            nonce,
            TypeCasts.addressToBytes32(address(this)),
            recipient,
            message
        );
    }
}
