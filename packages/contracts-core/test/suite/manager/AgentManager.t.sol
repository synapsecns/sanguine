// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {IAgentManager} from "../../../contracts/interfaces/IAgentManager.sol";

import {SystemEntity, SystemRouterHarness} from "../../harnesses/system/SystemRouterHarness.t.sol";

import {SystemContractTest} from "../system/SystemContract.t.sol";

// solhint-disable no-empty-blocks
// solhint-disable ordering
abstract contract AgentManagerTest is SystemContractTest {
    uint256 internal rootSubmittedAt;

    /// @notice Prevents this contract from being included in the coverage report
    function testAgentManagerTest() external {}

    function skipBondingOptimisticPeriod() public {
        skipPeriod(BONDING_OPTIMISTIC_PERIOD);
    }

    function skipPeriod(uint256 period) public {
        rootSubmittedAt = block.timestamp;
        skip(period);
    }

    function systemPrank(SystemRouterHarness router, uint32 callOrigin, SystemEntity systemCaller, bytes memory payload)
        public
    {
        router.systemPrank({
            recipient: SystemEntity.AgentManager,
            proofMaturity: block.timestamp - rootSubmittedAt,
            callOrigin: callOrigin,
            systemCaller: systemCaller,
            payload: payload
        });
    }

    // ══════════════════════════════════════════════════ HELPERS ══════════════════════════════════════════════════════

    function remoteSlashPayload(uint32 domain, address agent, address prover) public view returns (bytes memory) {
        // (proofMaturity, callOrigin, systemCaller) are omitted; (domain, agent, prover)
        return abi.encodeWithSelector(bondingManager.remoteRegistrySlash.selector, domain, agent, prover);
    }

    /// @notice Returns address of the tested system contract
    function systemContract() public view override returns (address) {
        return localAgentManager();
    }
}
