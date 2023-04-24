// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {AgentStatus, IAgentSecured} from "../../../contracts/interfaces/IAgentSecured.sol";

contract AgentSecuredMock is IAgentSecured {
    /// @notice Prevents this contract from being included in the coverage report
    function testAgentSecuredMock() external {}

    function managerSlash(uint32 domain, address agent, address prover) external {}

    function agentManager() external view returns (address) {}

    function agentStatus(address agent) external view returns (AgentStatus memory) {}

    function getAgent(uint256 index) external view returns (address agent, AgentStatus memory status) {}
}
