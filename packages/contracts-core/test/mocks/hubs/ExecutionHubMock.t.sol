// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {IExecutionHub, MessageStatus} from "../../../contracts/interfaces/IExecutionHub.sol";
import {BaseMock} from "../base/BaseMock.t.sol";

// solhint-disable no-empty-blocks
contract ExecutionHubMock is BaseMock, IExecutionHub {
    /// @notice Prevents this contract from being included in the coverage report
    function testExecutionHubMock() external {}

    function execute(
        bytes memory msgPayload,
        bytes32[] calldata originProof,
        bytes32[] calldata snapProof,
        uint8 stateIndex,
        uint64 gasLimit
    ) external {}

    function getAttestationNonce(bytes32) external view returns (uint32 attNonce) {
        return getReturnValueUint32();
    }

    function isValidReceipt(bytes memory) external view returns (bool isValid) {
        return getReturnValueBool();
    }

    function messageStatus(bytes32 messageHash) external view returns (MessageStatus status) {}

    function messageReceipt(bytes32 messageHash) external view returns (bytes memory data) {}
}
