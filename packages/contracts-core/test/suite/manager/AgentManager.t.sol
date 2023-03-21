// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import { AgentInfo } from "../../../contracts/libs/Structures.sol";

import {
    SystemEntity,
    SystemRouterHarness
} from "../../harnesses/system/SystemRouterHarness.t.sol";

import { ISystemContract, SynapseTest } from "../../utils/SynapseTest.t.sol";

abstract contract AgentManagerTest is SynapseTest {
    uint256 internal rootSubmittedAt;

    /// @notice Prevents this contract from being included in the coverage report
    function testAgentManagerTest() external {}

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
}
