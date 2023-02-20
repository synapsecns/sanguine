// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import { DomainContext } from "../../../contracts/context/DomainContext.sol";
import { SystemRegistry } from "../../../contracts/system/SystemRegistry.sol";
import "../events/SystemContractMockEvents.sol";

// solhint-disable no-empty-blocks
contract SystemRegistryMock is SystemContractMockEvents, SystemRegistry {
    constructor(uint32 _domain) DomainContext(_domain) {}

    function initialize() external initializer {
        __SystemContract_initialize();
    }

    function isNotary(uint32 _domain, address _notary) public view returns (bool) {
        return _isActiveAgent(_domain, _notary);
    }

    function isGuard(address _guard) public view returns (bool) {
        return _isActiveAgent(0, _guard);
    }

    /**
     * @notice Hook that is called before the specified agent was slashed via a system call.
     */
    function _beforeAgentSlashed(AgentInfo memory _info) internal override {
        emit SlashAgentCall(_info);
    }

    function _isIgnoredAgent(uint32, address) internal pure override returns (bool) {
        return false;
    }
}
