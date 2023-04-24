// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {AgentStatus, InterfaceDestination} from "../../contracts/interfaces/InterfaceDestination.sol";
import {ExecutionHubMock} from "./hubs/ExecutionHubMock.t.sol";
import {DisputeHubMock} from "./hubs/DisputeHubMock.t.sol";
import {SystemRegistryMock} from "./system/SystemRegistryMock.t.sol";

// solhint-disable no-empty-blocks
contract DestinationMock is ExecutionHubMock, DisputeHubMock, SystemRegistryMock, InterfaceDestination {
    /// @notice Prevents this contract from being included in the coverage report
    function testDestinationMock() external {}

    function passAgentRoot() external returns (bool rootPassed, bool rootPending) {}

    function acceptAttestation(
        address notary,
        AgentStatus memory status,
        bytes memory attPayload,
        bytes memory attSignature
    ) external returns (bool wasAccepted) {}

    function attestationsAmount() external view returns (uint256) {}

    function getSignedAttestation(uint256 index)
        external
        view
        returns (bytes memory attPayload, bytes memory attSignature)
    {}

    function destStatus() external view returns (uint48 snapRootTime, uint48 agentRootTime, address notary) {}

    function nextAgentRoot() external view returns (bytes32) {}
}
