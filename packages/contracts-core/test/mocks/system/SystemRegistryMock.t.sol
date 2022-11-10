// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import { LocalDomainContext } from "../../../contracts/context/LocalDomainContext.sol";
import { SystemRegistry } from "../../../contracts/system/SystemRegistry.sol";
import "../events/SystemContractMockEvents.sol";
import "../../harnesses/registry/GlobalNotaryRegistryHarness.t.sol";
import "../../harnesses/registry/GuardRegistryHarness.t.sol";

// solhint-disable no-empty-blocks
contract SystemRegistryMock is
    SystemContractMockEvents,
    LocalDomainContext,
    SystemRegistry,
    GlobalNotaryRegistry,
    GuardRegistry
{
    constructor(uint32 _domain) LocalDomainContext(_domain) {}

    function initialize() external initializer {
        __SystemContract_initialize();
    }

    function isNotary(uint32 _domain, address _notary) public view returns (bool) {
        return _isNotary(_domain, _notary);
    }

    function isGuard(address _guard) public view returns (bool) {
        return _isGuard(_guard);
    }

    /**
     * @notice Hook that is called before the specified agent was slashed via a system call.
     */
    function _beforeAgentSlashed(AgentInfo memory _info) internal override {
        emit SlashAgentCall(_info);
    }
}
