// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import { IExecutionHub } from "../../../contracts/interfaces/IExecutionHub.sol";

// solhint-disable no-empty-blocks
contract ExecutionHubMock is IExecutionHub {
    /// @notice Prevents this contract from being included in the coverage report
    function testExecutionHubMock() external {}

    function execute(
        bytes memory msgPayload,
        bytes32[] calldata originProof,
        bytes32[] calldata snapProof,
        uint256 stateIndex
    ) external {}
}
