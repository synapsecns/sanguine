// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import "./BondingManager.t.sol";

// solhint-disable func-name-mixedcase
contract BondingSecondaryTest is BondingManagerTest {
    // Deploy mocks for every messaging contract
    constructor() SynapseTest(0) {}

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                             TESTS: SETUP                             ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function test_constructor_revert_onSynapseChain() public {
        // Should not be able to deploy on Synapse Chain
        vm.expectRevert("Can't be deployed on SynChain");
        new BondingSecondary(DOMAIN_SYNAPSE);
    }

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
        vm.assume(caller != address(systemRouter));
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
        vm.assume(caller != address(systemRouter));
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
        bytes memory data = _dataSyncAgentCall(info);
        bytes memory expectedCall = _expectedCall(ISystemContract.syncAgent.selector, info);
        // All system registries should be system called
        vm.expectCall(origin, expectedCall);
        vm.expectCall(destination, expectedCall);
        bondingManager.addAgent(domain, agent);
    }

    function test_removeAgent(uint32 domain, address agent) public {
        test_addAgent(domain, agent);
        AgentInfo memory info = AgentInfo({ domain: domain, account: agent, bonded: false });
        bytes memory data = _dataSyncAgentCall(info);
        bytes memory expectedCall = _expectedCall(ISystemContract.syncAgent.selector, info);
        // All system registries should be system called
        vm.expectCall(origin, expectedCall);
        vm.expectCall(destination, expectedCall);
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
            _systemPrank(systemRouter, _localDomain(), caller, data);
        }
    }

    function test_slashAgent_revert_synapseDomain_notBondingManager() public {
        AgentInfo memory info;
        bytes memory data = _dataSlashAgentCall(info);
        _skipBondingOptimisticPeriod();
        for (uint256 c = 0; c < uint8(type(SystemEntity).max); ++c) {
            // Should reject system calls from Synapse domain, if caller is not BondingManager
            SystemEntity caller = SystemEntity(c);
            if (caller == SystemEntity.BondingManager) continue;
            vm.expectRevert("!allowedCaller");
            _systemPrank(systemRouter, DOMAIN_SYNAPSE, caller, data);
        }
    }

    function test_slashAgent_revert_remoteNotSynapseDomain(uint32 callOrigin) public {
        // Exclude local calls and calls from Synapse Chain
        vm.assume(callOrigin != _localDomain() && callOrigin != DOMAIN_SYNAPSE);
        AgentInfo memory info;
        bytes memory data = _dataSlashAgentCall(info);
        _skipBondingOptimisticPeriod();
        for (uint256 c = 0; c < uint8(type(SystemEntity).max); ++c) {
            // Should reject cross-chain system calls from domains other than Synapse domain
            SystemEntity caller = SystemEntity(c);
            vm.expectRevert("!synapseDomain");
            // Use mocked agent info
            _systemPrank(systemRouter, callOrigin, caller, data);
        }
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                      TESTS: SYNC AGENTS REVERTS                      ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function test_syncAgent_revert_localDomain() public {
        AgentInfo memory info;
        bytes memory data = _dataSyncAgentCall(info);
        for (uint256 c = 0; c < uint8(type(SystemEntity).max); ++c) {
            // Should reject all system calls from local domain
            SystemEntity caller = SystemEntity(c);
            // Calls from local domain never pass the optimistic period check
            vm.expectRevert("!optimisticPeriod");
            _systemPrank(systemRouter, _localDomain(), caller, data);
        }
    }

    function test_syncAgent_revert_remoteNotSynapseDomain(uint32 callOrigin) public {
        AgentInfo memory info;
        bytes memory data = _dataSyncAgentCall(info);
        // Exclude local calls and calls from Synapse Chain
        vm.assume(callOrigin != _localDomain() && callOrigin != DOMAIN_SYNAPSE);
        _skipBondingOptimisticPeriod();
        for (uint256 c = 0; c < uint8(type(SystemEntity).max); ++c) {
            // Should reject all system calls from remote domains other than Synapse domain
            SystemEntity caller = SystemEntity(c);
            vm.expectRevert("!synapseDomain");
            _systemPrank(systemRouter, callOrigin, caller, data);
        }
    }

    function test_syncAgent_revert_synapseDomain_notBondingManager() public {
        AgentInfo memory info;
        bytes memory data = _dataSyncAgentCall(info);
        _skipBondingOptimisticPeriod();
        for (uint256 c = 0; c < uint8(type(SystemEntity).max); ++c) {
            SystemEntity caller = SystemEntity(c);
            // Should reject system calls from Synapse domain, if caller is not BondingManager
            if (caller == SystemEntity.BondingManager) continue;
            vm.expectRevert("!allowedCaller");
            _systemPrank(systemRouter, DOMAIN_SYNAPSE, caller, data);
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
        vm.expectCall(origin, expectedCall);
        vm.expectCall(destination, expectedCall);
        // data should be forwarded to Synapse Chain
        vm.expectCall(
            address(systemRouter),
            abi.encodeWithSelector(
                SystemRouter.systemCall.selector,
                DOMAIN_SYNAPSE, // destination
                BONDING_OPTIMISTIC_PERIOD, // optimisticSeconds
                SystemEntity.BondingManager, //recipient
                data
            )
        );
        // Prank a local system call: [Local Origin] -> [Local BondingManager].slashAgent
        _systemPrank(systemRouter, _localDomain(), SystemEntity.Origin, data);
    }

    function _localDomain() internal pure override returns (uint32) {
        return DOMAIN_LOCAL;
    }
}
