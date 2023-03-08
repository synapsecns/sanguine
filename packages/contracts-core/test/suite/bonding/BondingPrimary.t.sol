// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import "./BondingManager.t.sol";

// solhint-disable func-name-mixedcase
contract BondingPrimaryTest is BondingManagerTest {
    Summit internal bondingPrimary;

    // Deploy Production version of Summit and mocks for everything else
    constructor() SynapseTest(DEPLOY_PROD_SUMMIT) {}

    function setUp() public virtual override {
        super.setUp();
        bondingPrimary = Summit(summit);
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                TESTS: UNAUTHORIZED ACCESS (NOT OWNER)                ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function test_addAgent_revert_notOwner(address caller) public {
        vm.assume(caller != address(this));
        expectRevertNotOwner();
        vm.prank(caller);
        bondingPrimary.addAgent(1, address(1));
    }

    function test_removeAgent_revert_notOwner(address caller) public {
        vm.assume(caller != address(this));
        expectRevertNotOwner();
        vm.prank(caller);
        bondingPrimary.removeAgent(1, address(1));
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
            bondingPrimary.slashAgent(submittedAt, callOrigin, systemCaller, info);
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
            bondingPrimary.syncAgent(submittedAt, callOrigin, systemCaller, info);
        }
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                       TESTS: ADD/REMOVE AGENTS                       ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function test_addAgent(uint32 domain, address agent) public {
        (bool isActive, ) = bondingPrimary.isActiveAgent(agent);
        // Should not be an already added agent
        vm.assume(!isActive);
        AgentInfo memory info = AgentInfo({ domain: domain, account: agent, bonded: true });
        bytes memory data = _dataSyncAgentCall(info);
        bytes memory expectedCall = _expectedCall(ISystemContract.syncAgent.selector, info);
        // All system registries should be system called
        vm.expectCall(originSynapse, expectedCall);
        vm.expectCall(destinationSynapse, expectedCall);
        bondingPrimary.addAgent(domain, agent);
    }

    function test_removeAgent(uint32 domain, address agent) public {
        test_addAgent(domain, agent);
        AgentInfo memory info = AgentInfo({ domain: domain, account: agent, bonded: false });
        bytes memory data = _dataSyncAgentCall(info);
        bytes memory expectedCall = _expectedCall(ISystemContract.syncAgent.selector, info);
        // All system registries should be system called
        vm.expectCall(originSynapse, expectedCall);
        vm.expectCall(destinationSynapse, expectedCall);
        bondingPrimary.removeAgent(domain, agent);
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

    function test_slashAgent_revert_remoteDomain_notBondingManager(uint32 callOrigin) public {
        AgentInfo memory info;
        bytes memory data = _dataSlashAgentCall(info);
        vm.assume(callOrigin != DOMAIN_SYNAPSE);
        _skipBondingOptimisticPeriod();
        for (uint256 c = 0; c < uint8(type(SystemEntity).max); ++c) {
            // Should reject system calls from a remote domain, if caller is not BondingManager
            SystemEntity caller = SystemEntity(c);
            if (caller == SystemEntity.BondingManager) continue;
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
            vm.expectRevert("Disabled for BondingPrimary");
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
        vm.expectCall(destinationSynapse, expectedCall);
        // callOrigin is Synapse Chain
        _expectForwardCalls(DOMAIN_SYNAPSE, data);
        // Prank a local system call: [Local Origin] -> [Local BondingManager].slashAgent
        _systemPrank(systemRouterSynapse, _localDomain(), SystemEntity.Origin, data);
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║             TESTS: RECEIVE SYSTEM CALLS (REMOTE DOMAIN)              ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function test_slashAgent_remoteDomain_bondingManager(
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
        vm.expectCall(destinationSynapse, expectedCall);
        // data should be forwarded to remote chains
        _expectForwardCalls(callOrigin, _dataSlashAgentCall(info));
        // Prank a local system call: [Remote BondingManager] -> [Local BondingManager].slashAgent
        _systemPrank(systemRouterSynapse, callOrigin, SystemEntity.BondingManager, data);
    }

    function _expectForwardCalls(uint32 callOrigin, bytes memory data) internal {
        if (callOrigin != DOMAIN_LOCAL) {
            vm.expectCall(
                address(systemRouterSynapse),
                abi.encodeWithSelector(
                    SystemRouter.systemCall.selector,
                    DOMAIN_LOCAL, // destination
                    BONDING_OPTIMISTIC_PERIOD, // optimisticSeconds
                    SystemEntity.BondingManager, //recipient
                    data
                )
            );
        }
        if (callOrigin != DOMAIN_REMOTE) {
            vm.expectCall(
                address(systemRouterSynapse),
                abi.encodeWithSelector(
                    SystemRouter.systemCall.selector,
                    DOMAIN_REMOTE, // destination
                    BONDING_OPTIMISTIC_PERIOD, // optimisticSeconds
                    SystemEntity.BondingManager, //recipient
                    data
                )
            );
        }
    }

    function _localDomain() internal pure override returns (uint32) {
        return DOMAIN_SYNAPSE;
    }
}
