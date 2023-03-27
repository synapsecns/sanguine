// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import { IAgentManager } from "../../../contracts/interfaces/IAgentManager.sol";

import {
    SystemEntity,
    SystemRouterHarness
} from "../../harnesses/system/SystemRouterHarness.t.sol";

import { ISystemContract, SynapseTest } from "../../utils/SynapseTest.t.sol";

abstract contract AgentManagerTest is SynapseTest {
    uint256 internal rootSubmittedAt;

    /// @notice Prevents this contract from being included in the coverage report
    function testAgentManagerTest() external {}

    function checkActive(
        IAgentManager manager,
        uint32 domain,
        address agent
    ) public {
        assertTrue(manager.isActiveAgent(domain, agent), "!isActive(domain, agent)");
        (bool isActive, uint32 _domain) = manager.isActiveAgent(agent);
        assertTrue(isActive, "!isActive(agent)");
        assertEq(_domain, domain, "!isActive(agent): domain");
    }

    function checkInactive(
        IAgentManager manager,
        uint32 domain,
        address agent
    ) public {
        assertFalse(manager.isActiveAgent(domain, agent), "!isActive(domain, agent)");
        (bool isActive, ) = manager.isActiveAgent(agent);
        assertFalse(isActive, "!isActive(agent)");
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                           INTERNAL HELPERS                           ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function _localDomain() internal view virtual returns (uint32);

    function _skipBondingOptimisticPeriod() internal {
        _skipPeriod(BONDING_OPTIMISTIC_PERIOD);
    }

    function _skipPeriod(uint256 period) internal {
        rootSubmittedAt = block.timestamp;
        skip(period);
    }

    function _systemPrank(
        SystemRouterHarness router,
        uint32 callOrigin,
        SystemEntity systemCaller,
        bytes memory data
    ) internal {
        router.systemPrank({
            _recipient: SystemEntity.AgentManager,
            _rootSubmittedAt: callOrigin == _localDomain() ? block.timestamp : rootSubmittedAt,
            _callOrigin: callOrigin,
            _systemCaller: systemCaller,
            _data: data
        });
    }

    function _remoteSlashData(
        uint32 domain,
        address agent,
        address reporter
    ) internal view returns (bytes memory) {
        // (_rootSubmittedAt, _callOrigin, _systemCaller, _domain, _agent, _reporter)
        return
            abi.encodeWithSelector(
                bondingManager.remoteRegistrySlash.selector,
                0,
                0,
                0,
                domain,
                agent,
                reporter
            );
    }
}
