// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import { IExecutionHub } from "../../../contracts/interfaces/IExecutionHub.sol";

contract ExecutionHubMock is IExecutionHub {
    /// @notice Prevents this contract from being included in the coverage report
    function testExecutionHubMock() external {}

    function execute(
        bytes memory _message,
        bytes32[] calldata _originProof,
        bytes32[] calldata _snapProof,
        uint256 _stateIndex
    ) external {}
}
