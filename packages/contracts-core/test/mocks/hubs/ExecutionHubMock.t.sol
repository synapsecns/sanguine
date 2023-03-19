// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import { IExecutionHub, TREE_DEPTH } from "../../../contracts/interfaces/IExecutionHub.sol";

contract ExecutionHubMock is IExecutionHub {
    /// @notice Prevents this contract from being included in the coverage report
    function testExecutionHubMock() external {}

    function execute(
        bytes memory _message,
        bytes32[TREE_DEPTH] calldata _originProof,
        bytes32[] calldata _snapProof,
        uint256 _stateIndex
    ) external {}
}
