// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {IAgentManager} from "../../../contracts/interfaces/IAgentManager.sol";

import {SystemEntity, SystemRouterHarness} from "../../harnesses/system/SystemRouterHarness.t.sol";

import {ISystemContract, SynapseTest} from "../../utils/SynapseTest.t.sol";

// solhint-disable no-empty-blocks
// solhint-disable ordering
abstract contract AgentManagerTest is SynapseTest {
    uint256 internal rootSubmittedAt;

    /// @notice Prevents this contract from being included in the coverage report
    function testAgentManagerTest() external {}

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
        bytes memory payload
    ) internal {
        router.systemPrank({
            recipient: SystemEntity.AgentManager,
            proofMaturity: block.timestamp - rootSubmittedAt,
            callOrigin: callOrigin,
            systemCaller: systemCaller,
            payload: payload
        });
    }

    function _remoteSlashPayload(uint32 domain, address agent, address prover) internal view returns (bytes memory) {
        // (proofMaturity, callOrigin, systemCaller) are omitted; (domain, agent, prover)
        return abi.encodeWithSelector(bondingManager.remoteRegistrySlash.selector, domain, agent, prover);
    }
}
