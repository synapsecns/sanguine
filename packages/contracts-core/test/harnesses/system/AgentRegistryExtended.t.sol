// SPDX-License-Identifier: MIT

pragma solidity 0.8.17;

import { AgentRegistry } from "../../../contracts/system/AgentRegistry.sol";

contract AgentRegistryExtended is AgentRegistry {
    /// @notice Prevents this contract from being included in the coverage report
    function testAgentRegistryExtended() external {}

    function addAgent(uint32 _domain, address _account) external returns (bool) {
        return _addAgent(_domain, _account);
    }

    function removeAgent(uint32 _domain, address _account) external returns (bool) {
        return _removeAgent(_domain, _account);
    }

    function removeAllAgents() external {
        _resetAgents();
    }

    function _isIgnoredAgent(uint32, address) internal view virtual override returns (bool) {
        return false;
    }
}
