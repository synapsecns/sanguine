// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import "./BondingManager.t.sol";
import { BondingMVP } from "../../../contracts/bonding/BondingMVP.sol";

// solhint-disable func-name-mixedcase
contract BondingMVPTest is BondingManagerTest {
    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                        ADD/REMOVE AGENT TESTS                        ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function test_addAgent_revert_onlyOwner(address caller) public {
        vm.assume(caller != owner);
        expectRevertNotOwner();
        vm.prank(caller);
        _castToMVP().addAgent(1, address(1));
    }

    function test_removeAgent_revert_onlyOwner(address caller) public {
        vm.assume(caller != owner);
        expectRevertNotOwner();
        vm.prank(caller);
        _castToMVP().removeAgent(1, address(1));
    }

    function test_addAgent(uint32 domain, address agent) public {
        SystemContract.AgentInfo[] memory infos = infoToArray(agentInfo(domain, agent, true));
        // All system registries should be system called
        for (uint256 r = 0; r < systemRegistries.length; ++r) {
            vm.expectEmit(true, true, true, true, systemRegistries[r]);
            // Default values are used in MVP implementation
            emit SyncAgentsCall({ requestID: 0, removeExisting: false, infos: infos });
        }
        vm.prank(owner);
        _castToMVP().addAgent(domain, agent);
    }

    function test_removeAgent(uint32 domain, address agent) public {
        SystemContract.AgentInfo[] memory infos = infoToArray(agentInfo(domain, agent, false));
        // All system registries should be system called
        for (uint256 r = 0; r < systemRegistries.length; ++r) {
            vm.expectEmit(true, true, true, true, systemRegistries[r]);
            // Default values are used in MVP implementation
            emit SyncAgentsCall({ requestID: 0, removeExisting: false, infos: infos });
        }
        vm.prank(owner);
        _castToMVP().removeAgent(domain, agent);
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                      CROSS-CHAIN CALLS: REVERTS                      ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function test_slashAgent_revert_remoteDomain(uint32 callOrigin) public {
        // Exclude local calls and calls from Synapse Chain
        vm.assume(callOrigin != DOMAIN_LOCAL);
        _skipBondingOptimisticPeriod();
        for (uint256 c = 0; c < uint8(type(InterfaceSystemRouter.SystemEntity).max); ++c) {
            // Should reject all system calls from remote domains
            InterfaceSystemRouter.SystemEntity caller = InterfaceSystemRouter.SystemEntity(c);
            vm.expectRevert("Cross-chain disabled");
            // Use mocked agent info
            _mockSlashAgentCall({
                callOrigin: callOrigin,
                systemCaller: caller,
                info: guardInfo({ guard: address(0), bonded: false })
            });
        }
    }

    function test_syncAgents_revert_remoteDomain(uint32 callOrigin) public {
        // Exclude local calls
        vm.assume(callOrigin != localDomain);
        _skipBondingOptimisticPeriod();
        for (uint256 c = 0; c < uint8(type(InterfaceSystemRouter.SystemEntity).max); ++c) {
            // Should reject all system calls from remote domains
            InterfaceSystemRouter.SystemEntity caller = InterfaceSystemRouter.SystemEntity(c);
            vm.expectRevert("Cross-chain disabled");
            // Use mocked list of agents
            _mockSyncAgentsCall({
                callOrigin: callOrigin,
                systemCaller: caller,
                requestID: 0,
                removeExisting: false,
                infos: new SystemContract.AgentInfo[](0)
            });
        }
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                           INTERNAL HELPERS                           ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function _deployBondingManager(uint32 domain) internal override returns (BondingManager) {
        return new BondingMVP(domain);
    }

    function _castToMVP() internal view returns (BondingMVP) {
        return BondingMVP(address(bondingManager));
    }

    function _getTestLocalDomain() internal pure override returns (uint32) {
        return DOMAIN_LOCAL;
    }
}
