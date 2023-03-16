// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import { IDisputeHub } from "../../../contracts/interfaces/IDisputeHub.sol";

contract DisputeHubMock is IDisputeHub {
    /// @notice Prevents this contract from being included in the coverage report
    function testDisputeHubMock() external {}

    function submitStateReport(
        uint256 _stateIndex,
        bytes memory _srPayload,
        bytes memory _srSignature,
        bytes memory _snapPayload,
        bytes memory _snapSignature
    ) external returns (bool wasAccepted) {}

    function submitStateReportWithProof(
        uint256 _stateIndex,
        bytes memory _srPayload,
        bytes memory _srSignature,
        bytes32[] memory _snapProof,
        bytes memory _attPayload,
        bytes memory _attSignature
    ) external returns (bool wasAccepted) {}
}
