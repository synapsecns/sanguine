// SPDX-License-Identifier: MIT

pragma solidity 0.8.17;

import { IMessageRecipient } from "../../../contracts/interfaces/IMessageRecipient.sol";
import "../../../contracts/interfaces/InterfaceDestination.sol";

contract ReentrantApp is IMessageRecipient {
    bytes internal message;
    bytes32[ORIGIN_TREE_DEPTH] internal originProof;
    bytes32[] internal snapProof;
    uint256 internal stateIndex;

    function prepare(
        bytes memory _message,
        bytes32[ORIGIN_TREE_DEPTH] memory _originProof,
        bytes32[] memory _snapProof,
        uint256 _stateIndex
    ) external {
        message = _message;
        originProof = _originProof;
        snapProof = _snapProof;
        stateIndex = _stateIndex;
    }

    function handle(
        uint32,
        uint32,
        bytes32,
        uint256,
        bytes memory
    ) external {
        InterfaceDestination(msg.sender).execute(message, originProof, snapProof, stateIndex);
    }
}
