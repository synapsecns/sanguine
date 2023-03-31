// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;
// ══════════════════════════════ LIBRARY IMPORTS ══════════════════════════════

import {TipsLib} from "../libs/Tips.sol";
import {TypeCasts} from "../libs/TypeCasts.sol";

// ═════════════════════════════ INTERNAL IMPORTS ══════════════════════════════
import {IMessageRecipient} from "../interfaces/IMessageRecipient.sol";
import {InterfaceOrigin} from "../interfaces/InterfaceOrigin.sol";

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

    event MessageReceived(uint32 origin, uint32 nonce, bytes32 sender, uint256 rootSubmittedAt, bytes content);

    event MessageSent(uint32 destination, uint32 nonce, bytes32 sender, bytes32 recipient, bytes content);

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                             CONSTRUCTOR                              ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    constructor(address origin_, address destination_) {
        origin = origin_;
        destination = destination_;
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                          RECEIVING MESSAGES                          ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function handle(uint32 origin_, uint32 nonce, bytes32 sender, uint256 rootSubmittedAt, bytes memory content)
        external
    {
        require(msg.sender == destination, "TestClient: !destination");
        emit MessageReceived(origin_, nonce, sender, rootSubmittedAt, content);
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                           SENDING MESSAGES                           ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function sendMessage(uint32 destination_, address recipientAddress, uint32 optimisticSeconds, bytes memory content)
        external
    {
        bytes32 recipient = TypeCasts.addressToBytes32(recipientAddress);
        bytes memory tipsPayload = TipsLib.emptyTips();
        (uint32 nonce,) =
            InterfaceOrigin(origin).dispatch(destination_, recipient, optimisticSeconds, tipsPayload, content);
        emit MessageSent(destination_, nonce, TypeCasts.addressToBytes32(address(this)), recipient, content);
    }
}
