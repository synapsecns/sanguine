// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {AgentStatus, DisputeStatus, IAgentSecured} from "../../../contracts/interfaces/IAgentSecured.sol";

contract AgentSecuredMock is IAgentSecured {
    /// @notice Prevents this contract from being included in the coverage report
    function testAgentSecuredMock() external {}

    function openDispute(uint32 guardIndex, uint32 notaryIndex) external {}

    function resolveDispute(uint32 slashedIndex, uint32 rivalIndex) external {}

    function agentManager() external view returns (address) {}

    function inbox() external view returns (address) {}

    function agentStatus(address agent) external view returns (AgentStatus memory) {}

    function getAgent(uint256 index) external view returns (address agent, AgentStatus memory status) {}

    function latestDisputeStatus(uint32 agentIndex) external view returns (DisputeStatus memory) {}
}
