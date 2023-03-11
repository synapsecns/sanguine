// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import { IAgentRegistry } from "../../../contracts/interfaces/IAgentRegistry.sol";

// solhint-disable no-empty-blocks
abstract contract AgentRegistryMock is IAgentRegistry {
    /// @notice Prevents this contract from being included in the coverage report
    function testAgentRegistryMock() external {}

    function amountAgents(uint32 _domain) external view returns (uint256) {}

    function amountDomains() external view returns (uint256) {}

    function getAgent(uint32 _domain, uint256 _agentIndex) external view returns (address) {}

    function getDomain(uint256 _domainIndex) external view returns (uint32) {}

    function allAgents(uint32 _domain) external view returns (address[] memory) {}

    function allDomains() external view returns (uint32[] memory domains_) {}

    function isActiveAgent(address _account) external view returns (bool isActive, uint32 domain) {}

    function isActiveAgent(uint32 _domain, address _account) external view returns (bool) {}

    function isActiveDomain(uint32 _domain) external view returns (bool) {}
}
