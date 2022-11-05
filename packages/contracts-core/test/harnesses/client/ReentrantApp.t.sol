// SPDX-License-Identifier: MIT

pragma solidity 0.8.17;

import { IMessageRecipient } from "../../../contracts/interfaces/IMessageRecipient.sol";
import { Destination } from "../../../contracts/Destination.sol";

contract ReentrantApp is IMessageRecipient {
    bytes internal message;
    bytes32[32] internal proof;
    uint256 internal index;

    function prepare(
        bytes memory _message,
        bytes32[32] memory _proof,
        uint256 _index
    ) external {
        message = _message;
        proof = _proof;
        index = _index;
    }

    function handle(
        uint32,
        uint32,
        bytes32,
        uint256,
        bytes memory
    ) external {
        Destination(msg.sender).execute(message, proof, index);
    }
}
