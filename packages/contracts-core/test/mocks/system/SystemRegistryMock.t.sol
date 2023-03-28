// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import { AgentStatus, ISystemRegistry } from "../../../contracts/interfaces/ISystemRegistry.sol";

contract SystemRegistryMock is ISystemRegistry {
    /// @notice Prevents this contract from being included in the coverage report
    function testSystemRegistryMock() external {}

    function managerSlash(
        uint32 _domain,
        address _agent,
        address _prover
    ) external {}

    function agentStatus(address _agent) external view returns (AgentStatus memory) {}
}
