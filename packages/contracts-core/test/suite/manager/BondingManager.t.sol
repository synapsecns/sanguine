// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import { AgentInfo, SystemEntity } from "../../../contracts/libs/Structures.sol";
import { AgentManagerTest } from "./AgentManager.t.sol";

import { ISystemContract, Summit, SynapseTest } from "../../utils/SynapseTest.t.sol";

// solhint-disable func-name-mixedcase
contract BondingManagerTest is AgentManagerTest {
    // Deploy Production version of Summit and mocks for everything else
    constructor() SynapseTest(DEPLOY_PROD_SUMMIT) {}

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                TESTS: UNAUTHORIZED ACCESS (NOT OWNER)                ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function test_addAgent_revert_notOwner(address caller) public {
        vm.assume(caller != address(this));
        expectRevertNotOwner();
        vm.prank(caller);
        bondingManager.addAgent(1, address(1));
    }

    function test_removeAgent_revert_notOwner(address caller) public {
        vm.assume(caller != address(this));
        expectRevertNotOwner();
        vm.prank(caller);
        bondingManager.removeAgent(1, address(1));
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║            TESTS: UNAUTHORIZED ACCESS (NOT SYSTEM ROUTER)            ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function test_slashAgent_revert_notSystemRouter(
        address caller,
        uint256 submittedAt,
        uint32 callOrigin
    ) public {
        vm.assume(caller != address(systemRouterSynapse));
        AgentInfo memory info;
        for (uint256 c = 0; c < uint8(type(SystemEntity).max); ++c) {
            SystemEntity systemCaller = SystemEntity(c);
            vm.expectRevert("!systemRouter");
            vm.prank(caller);
            bondingManager.slashAgent(submittedAt, callOrigin, systemCaller, info);
        }
    }

    function test_syncAgent_revert_notSystemRouter(
        address caller,
        uint256 submittedAt,
        uint32 callOrigin
    ) public {
        vm.assume(caller != address(systemRouterSynapse));
        AgentInfo memory info;
        for (uint256 c = 0; c < uint8(type(SystemEntity).max); ++c) {
            SystemEntity systemCaller = SystemEntity(c);
            vm.expectRevert("!systemRouter");
            vm.prank(caller);
            bondingManager.syncAgent(submittedAt, callOrigin, systemCaller, info);
        }
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                       TESTS: ADD/REMOVE AGENTS                       ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function test_addAgent(uint32 domain, address agent) public {
        (bool isActive, ) = bondingManager.isActiveAgent(agent);
        // Should not be an already added agent
        vm.assume(!isActive);
        AgentInfo memory info = AgentInfo({ domain: domain, account: agent, bonded: true });
        bytes memory expectedCall = _expectedCall(ISystemContract.syncAgent.selector, info);
        // All system registries should be system called
        vm.expectCall(originSynapse, expectedCall);
        vm.expectCall(summit, expectedCall);
        bondingManager.addAgent(domain, agent);
    }

    function test_removeAgent(uint32 domain, address agent) public {
        test_addAgent(domain, agent);
        AgentInfo memory info = AgentInfo({ domain: domain, account: agent, bonded: false });
        bytes memory expectedCall = _expectedCall(ISystemContract.syncAgent.selector, info);
        // All system registries should be system called
        vm.expectCall(originSynapse, expectedCall);
        vm.expectCall(summit, expectedCall);
        bondingManager.removeAgent(domain, agent);
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                      TESTS: SLASH AGENT REVERTS                      ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function test_slashAgent_revert_localDomain_notOrigin() public {
        AgentInfo memory info;
        bytes memory data = _dataSlashAgentCall(info);
        // Only Origin on local domain is allowed to call slashAgent()
        for (uint256 c = 0; c < uint8(type(SystemEntity).max); ++c) {
            SystemEntity caller = SystemEntity(c);
            // Should reject system calls from local domain, if caller is not Origin
            if (caller == SystemEntity.Origin) continue;
            vm.expectRevert("!allowedCaller");
            _systemPrank(systemRouterSynapse, _localDomain(), caller, data);
        }
    }

    function test_slashAgent_revert_remoteDomain_notAgentManager(uint32 callOrigin) public {
        AgentInfo memory info;
        bytes memory data = _dataSlashAgentCall(info);
        vm.assume(callOrigin != DOMAIN_SYNAPSE);
        _skipBondingOptimisticPeriod();
        for (uint256 c = 0; c < uint8(type(SystemEntity).max); ++c) {
            // Should reject system calls from a remote domain, if caller is not AgentManager
            SystemEntity caller = SystemEntity(c);
            if (caller == SystemEntity.AgentManager) continue;
            vm.expectRevert("!allowedCaller");
            _systemPrank(systemRouterSynapse, callOrigin, caller, data);
        }
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                      TESTS: SYNC AGENTS REVERTS                      ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function test_syncAgent_revert(uint32 callOrigin) public {
        AgentInfo memory info;
        bytes memory data = _dataSyncAgentCall(info);
        // Should reject all syncAgent calls
        for (uint256 c = 0; c < uint8(type(SystemEntity).max); ++c) {
            SystemEntity caller = SystemEntity(c);
            vm.expectRevert("Disabled for BondingManager");
            _systemPrank(systemRouterSynapse, callOrigin, caller, data);
        }
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║              TESTS: RECEIVE SYSTEM CALLS (LOCAL DOMAIN)              ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function test_slashAgent_localDomain_origin(uint32 domain, address account) public {
        AgentInfo memory info = AgentInfo({ domain: domain, account: account, bonded: false });
        bytes memory data = _dataSlashAgentCall(info);
        bytes memory expectedCall = _expectedCall(ISystemContract.slashAgent.selector, info);
        // All system registries should be system called
        vm.expectCall(originSynapse, expectedCall);
        vm.expectCall(summit, expectedCall);
        // callOrigin is Synapse Chain
        _expectForwardCalls(DOMAIN_SYNAPSE, data);
        // Prank a local system call: [Local Origin] -> [Local AgentManager].slashAgent
        _systemPrank(systemRouterSynapse, _localDomain(), SystemEntity.Origin, data);
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║             TESTS: RECEIVE SYSTEM CALLS (REMOTE DOMAIN)              ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function test_slashAgent_remoteDomain_agentManager(
        uint32 callOrigin,
        uint32 domain,
        address account
    ) public {
        // TODO: restrict callOrigin to existing domains
        vm.assume(callOrigin != 0 && callOrigin != DOMAIN_SYNAPSE);
        _skipBondingOptimisticPeriod();
        AgentInfo memory info = AgentInfo({ domain: domain, account: account, bonded: false });
        bytes memory data = _dataSlashAgentCall(info);
        bytes memory expectedCall = _expectedCall(ISystemContract.slashAgent.selector, info);
        // All system registries should be system called
        vm.expectCall(originSynapse, expectedCall);
        vm.expectCall(summit, expectedCall);
        // data should be forwarded to remote chains
        _expectForwardCalls(callOrigin, _dataSlashAgentCall(info));
        // Prank a local system call: [Remote AgentManager] -> [Local AgentManager].slashAgent
        _systemPrank(systemRouterSynapse, callOrigin, SystemEntity.AgentManager, data);
    }

    function _expectForwardCalls(uint32 callOrigin, bytes memory data) internal {
        if (callOrigin != DOMAIN_LOCAL) {
            vm.expectCall(
                address(systemRouterSynapse),
                abi.encodeWithSelector(
                    systemRouterSynapse.systemCall.selector,
                    DOMAIN_LOCAL, // destination
                    BONDING_OPTIMISTIC_PERIOD, // optimisticSeconds
                    SystemEntity.AgentManager, //recipient
                    data
                )
            );
        }
        if (callOrigin != DOMAIN_REMOTE) {
            vm.expectCall(
                address(systemRouterSynapse),
                abi.encodeWithSelector(
                    systemRouterSynapse.systemCall.selector,
                    DOMAIN_REMOTE, // destination
                    BONDING_OPTIMISTIC_PERIOD, // optimisticSeconds
                    SystemEntity.AgentManager, //recipient
                    data
                )
            );
        }
    }

    function _localDomain() internal pure override returns (uint32) {
        return DOMAIN_SYNAPSE;
    }
}
