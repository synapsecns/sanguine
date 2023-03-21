// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import { ISystemRegistry } from "../../../contracts/interfaces/ISystemRegistry.sol";

contract SystemRegistryMock is ISystemRegistry {
    /// @notice Prevents this contract from being included in the coverage report
    function testSystemRegistryMock() external {}

    function managerSlash(uint32 _domain, address _agent) external {}

    function isActiveAgent(address _account) external view returns (bool isActive, uint32 domain) {}

    function isActiveAgent(uint32 _domain, address _account) external view returns (bool) {}
}
