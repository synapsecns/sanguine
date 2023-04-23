// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {IExecutionHub, MessageStatus} from "../../../contracts/interfaces/IExecutionHub.sol";

// solhint-disable no-empty-blocks
contract ExecutionHubMock is IExecutionHub {
    /// @notice Prevents this contract from being included in the coverage report
    function testExecutionHubMock() external {}

    function execute(
        bytes memory msgPayload,
        bytes32[] calldata originProof,
        bytes32[] calldata snapProof,
        uint256 stateIndex,
        uint64 gasLimit
    ) external {}

    function isValidReceipt(bytes memory rcptPayload) external view returns (bool isValid) {}

    function messageStatus(bytes32 messageHash) external view returns (MessageStatus status) {}

    function receiptBody(bytes32 messageHash) external view returns (bytes memory data) {}
}
