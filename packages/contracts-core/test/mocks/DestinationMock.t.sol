// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {AgentStatus, InterfaceDestination} from "../../contracts/interfaces/InterfaceDestination.sol";
import {ExecutionHubMock} from "./hubs/ExecutionHubMock.t.sol";
import {AgentSecuredMock} from "./base/AgentSecuredMock.t.sol";

// solhint-disable no-empty-blocks
contract DestinationMock is ExecutionHubMock, AgentSecuredMock, InterfaceDestination {
    /// @notice Prevents this contract from being included in the coverage report
    function testDestinationMock() external {}

    function passAgentRoot() external returns (bool rootPassed, bool rootPending) {}

    function acceptAttestation(uint32 notaryIndex, uint256 sigIndex, bytes memory attPayload)
        external
        returns (bool wasAccepted)
    {}

    function attestationsAmount() external view returns (uint256) {}

    function getAttestation(uint256 index) external view returns (bytes memory attPayload, bytes memory attSignature) {}

    function destStatus() external view returns (uint40 snapRootTime, uint40 agentRootTime, uint32 notaryIndex) {}

    function nextAgentRoot() external view returns (bytes32) {}
}
