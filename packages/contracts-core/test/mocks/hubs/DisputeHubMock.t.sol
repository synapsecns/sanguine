// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {DisputeStatus, IDisputeHub} from "../../../contracts/interfaces/IDisputeHub.sol";

// solhint-disable no-empty-blocks
contract DisputeHubMock is IDisputeHub {
    /// @notice Prevents this contract from being included in the coverage report
    function testDisputeHubMock() external {}

    function submitStateReport(
        uint256 stateIndex,
        bytes memory srPayload,
        bytes memory srSignature,
        bytes memory snapPayload,
        bytes memory snapSignature
    ) external returns (bool wasAccepted) {}

    function submitStateReportWithProof(
        uint256 stateIndex,
        bytes memory srPayload,
        bytes memory srSignature,
        bytes32[] memory snapProof,
        bytes memory attPayload,
        bytes memory attSignature
    ) external returns (bool wasAccepted) {}

    function disputeStatus(address agent) external view returns (DisputeStatus memory status) {}
}
