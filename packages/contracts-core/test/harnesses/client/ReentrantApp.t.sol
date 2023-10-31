// SPDX-License-Identifier: MIT

pragma solidity 0.8.17;

import {ORIGIN_TREE_HEIGHT} from "../../../contracts/libs/Constants.sol";
import {IMessageRecipient} from "../../../contracts/interfaces/IMessageRecipient.sol";
import {IExecutionHub} from "../../../contracts/interfaces/IExecutionHub.sol";

// solhint-disable no-empty-blocks
contract ReentrantApp is IMessageRecipient {
    bytes internal msgPayload;
    bytes32[] internal originProof;
    bytes32[] internal snapProof;
    uint8 internal stateIndex;

    /// @notice Prevents this contract from being included in the coverage report
    function testReentrantApp() external {}

    function prepare(
        bytes memory msgPayload_,
        bytes32[] memory originProof_,
        bytes32[] memory snapProof_,
        uint8 stateIndex_
    ) external {
        msgPayload = msgPayload_;
        originProof = originProof_;
        snapProof = snapProof_;
        stateIndex = stateIndex_;
    }

    function receiveBaseMessage(uint32, uint32, bytes32, uint256, uint32, bytes memory) external payable {
        IExecutionHub(msg.sender).execute(msgPayload, originProof, snapProof, stateIndex, 0);
    }
}
