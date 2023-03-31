// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import { AgentStatus, ISystemRegistry } from "../../../contracts/interfaces/ISystemRegistry.sol";

// solhint-disable no-empty-blocks
contract SystemRegistryMock is ISystemRegistry {
    /// @notice Prevents this contract from being included in the coverage report
    function testSystemRegistryMock() external {}

    function managerSlash(
        uint32 domain,
        address agent,
        address prover
    ) external {}

    function agentStatus(address agent) external view returns (AgentStatus memory) {}
}
