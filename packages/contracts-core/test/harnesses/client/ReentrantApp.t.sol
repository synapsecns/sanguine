// SPDX-License-Identifier: MIT

pragma solidity 0.8.17;

import { ORIGIN_TREE_HEIGHT } from "../../../contracts/libs/Constants.sol";
import { IMessageRecipient } from "../../../contracts/interfaces/IMessageRecipient.sol";
import { IExecutionHub } from "../../../contracts/interfaces/IExecutionHub.sol";

contract ReentrantApp is IMessageRecipient {
    bytes internal message;
    bytes32[] internal originProof;
    bytes32[] internal snapProof;
    uint256 internal stateIndex;

    /// @notice Prevents this contract from being included in the coverage report
    function testReentrantApp() external {}

    function prepare(
        bytes memory message,
        bytes32[] memory originProof,
        bytes32[] memory snapProof,
        uint256 stateIndex
    ) external {
        message = message;
        originProof = originProof;
        snapProof = snapProof;
        stateIndex = stateIndex;
    }

    function handle(
        uint32,
        uint32,
        bytes32,
        uint256,
        bytes memory
    ) external {
        IExecutionHub(msg.sender).execute(message, originProof, snapProof, stateIndex);
    }
}
