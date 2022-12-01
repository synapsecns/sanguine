// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import "./BondingManager.t.sol";

// solhint-disable func-name-mixedcase
contract BondingPrimaryTest is BondingManagerTest {
    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                             TESTS: SETUP                             ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function test_constructor_revert_onlySynapseChain(uint32 domain) public {
        vm.assume(domain != DOMAIN_SYNAPSE);
        // Should be able to deploy on Synapse Chain only
        vm.expectRevert("Only deployed on SynChain");
        new BondingPrimary(domain);
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                  TESTS: ADD/REMOVE AGENTS (REVERTS)                  ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function test_addNotary_revert_notOwner(address caller) public {
        vm.assume(caller != owner);
        expectRevertNotOwner();
        vm.prank(caller);
        _castToPrimary().addNotary(1, address(1));
    }

    function test_removeNotary_revert_notOwner(address caller) public {
        vm.assume(caller != owner);
        expectRevertNotOwner();
        vm.prank(caller);
        _castToPrimary().removeNotary(1, address(1));
    }

    function test_addGuard_revert_notOwner(address caller) public {
        vm.assume(caller != owner);
        expectRevertNotOwner();
        vm.prank(caller);
        _castToPrimary().addGuard(address(1));
    }

    function test_removeGuard_revert_notOwner(address caller) public {
        vm.assume(caller != owner);
        expectRevertNotOwner();
        vm.prank(caller);
        _castToPrimary().removeGuard(address(1));
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                       TESTS: ADD/REMOVE AGENTS                       ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function test_addNotary(uint32 domain, address notary) public {
        vm.assume(domain != 0);
        SystemContract.AgentInfo[] memory infos = infoToArray(
            agentInfo({ domain: domain, account: notary, bonded: true })
        );
        // All system registries should be system called
        for (uint256 r = 0; r < systemRegistries.length; ++r) {
            vm.expectEmit(true, true, true, true, systemRegistries[r]);
            // This is the first BondingPrimary request
            emit SyncAgentsCall({ requestID: 1, removeExisting: false, infos: infos });
        }
        vm.prank(owner);
        _castToPrimary().addNotary(domain, notary);
    }

    function test_removeNotary(uint32 domain, address notary) public {
        test_addNotary(domain, notary);
        SystemContract.AgentInfo[] memory infos = infoToArray(
            agentInfo({ domain: domain, account: notary, bonded: false })
        );
        // All system registries should be system called
        for (uint256 r = 0; r < systemRegistries.length; ++r) {
            vm.expectEmit(true, true, true, true, systemRegistries[r]);
            // This is the second BondingPrimary request
            emit SyncAgentsCall({ requestID: 2, removeExisting: false, infos: infos });
        }
        vm.prank(owner);
        _castToPrimary().removeNotary(domain, notary);
    }

    function test_addGuard(address guard) public {
        SystemContract.AgentInfo[] memory infos = infoToArray(
            guardInfo({ guard: guard, bonded: true })
        );
        // All system registries should be system called
        for (uint256 r = 0; r < systemRegistries.length; ++r) {
            vm.expectEmit(true, true, true, true, systemRegistries[r]);
            // This is the first BondingPrimary request
            emit SyncAgentsCall({ requestID: 1, removeExisting: false, infos: infos });
        }
        vm.prank(owner);
        _castToPrimary().addGuard(guard);
    }

    function test_removeGuard(address guard) public {
        test_addGuard(guard);
        SystemContract.AgentInfo[] memory infos = infoToArray(
            guardInfo({ guard: guard, bonded: false })
        );
        // All system registries should be system called
        for (uint256 r = 0; r < systemRegistries.length; ++r) {
            vm.expectEmit(true, true, true, true, systemRegistries[r]);
            // This is the second BondingPrimary request
            emit SyncAgentsCall({ requestID: 2, removeExisting: false, infos: infos });
        }
        vm.prank(owner);
        _castToPrimary().removeGuard(guard);
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                      TESTS: SLASH AGENT REVERTS                      ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    // slashAgent(): localDomain_notOrigin is tested in BondingManager.t.sol

    function test_slashAgent_revert_remoteDomain_notBondingManager(uint32 callOrigin) public {
        vm.assume(callOrigin != localDomain);
        _skipBondingOptimisticPeriod();
        for (uint256 c = 0; c < uint8(type(ISystemRouter.SystemEntity).max); ++c) {
            // Should reject system calls from a remote domain, if caller is not BondingManager
            ISystemRouter.SystemEntity caller = ISystemRouter.SystemEntity(c);
            if (caller == ISystemRouter.SystemEntity.BondingManager) continue;
            vm.expectRevert("!allowedCaller");
            // Use mocked agent info
            _mockSlashAgentCall({
                callOrigin: callOrigin,
                systemCaller: caller,
                info: guardInfo({ guard: address(0), bonded: false })
            });
        }
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                      TESTS: SYNC AGENTS REVERTS                      ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function test_syncAgents_revert_localDomain() public {
        for (uint256 c = 0; c < uint8(type(ISystemRouter.SystemEntity).max); ++c) {
            // Should reject all system calls from local domain
            ISystemRouter.SystemEntity caller = ISystemRouter.SystemEntity(c);
            // Calls from local domain never pass the optimistic period check
            vm.expectRevert("!optimisticPeriod");
            // Use mocked list of agents
            _mockSyncAgentsCall({
                callOrigin: localDomain,
                systemCaller: caller,
                requestID: 0,
                removeExisting: false,
                infos: new SystemContract.AgentInfo[](0)
            });
        }
    }

    function test_syncAgents_revert_remoteDomain_notBondingManager(uint32 callOrigin) public {
        vm.assume(callOrigin != localDomain);
        _skipBondingOptimisticPeriod();
        for (uint256 c = 0; c < uint8(type(ISystemRouter.SystemEntity).max); ++c) {
            ISystemRouter.SystemEntity caller = ISystemRouter.SystemEntity(c);
            // Should reject system calls from a remote domain, if caller is not BondingManager
            if (caller == ISystemRouter.SystemEntity.BondingManager) continue;
            vm.expectRevert("!allowedCaller");
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
    ▏*║              TESTS: RECEIVE SYSTEM CALLS (LOCAL DOMAIN)              ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function test_slashAgent_localDomain_origin(uint32 domain, address account) public {
        SystemContract.AgentInfo memory info = agentInfo({
            domain: domain,
            account: account,
            bonded: false
        });
        bytes memory data = abi.encodeWithSelector(
            SystemContract.slashAgent.selector,
            0, // rootSubmittedAt
            0, // callOrigin
            0, // systemCaller
            info
        );
        // All system registries should be system called
        for (uint256 r = 0; r < systemRegistries.length; ++r) {
            vm.expectEmit(true, true, true, true, systemRegistries[r]);
            emit SlashAgentCall(info);
        }
        // TODO: add test for forwarding the data once implemented
        data;
        // Mock a local system call: [Local Origin] -> [Local BondingManager].slashAgent
        _mockSlashAgentCall({
            callOrigin: localDomain,
            systemCaller: ISystemRouter.SystemEntity.Origin,
            info: info
        });
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
        SystemContract.AgentInfo memory info = agentInfo({
            domain: domain,
            account: account,
            bonded: false
        });
        bytes memory data = abi.encodeWithSelector(
            SystemContract.slashAgent.selector,
            0, // rootSubmittedAt
            0, // callOrigin
            0, // systemCaller
            info
        );
        // All system registries should be system called
        for (uint256 r = 0; r < systemRegistries.length; ++r) {
            vm.expectEmit(true, true, true, true, systemRegistries[r]);
            emit SlashAgentCall(info);
        }
        // TODO: add test for forwarding the data once implemented
        data;
        // Mock a local system call: [Remote BondingManager] -> [Local BondingManager].slashAgent
        _mockSlashAgentCall({
            callOrigin: callOrigin,
            systemCaller: ISystemRouter.SystemEntity.BondingManager,
            info: info
        });
    }

    // TODO: test for handling PONGs in syncAgents() once implemented

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                            INTERNAL VIEWS                            ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function _deployBondingManager(uint32 domain) internal override returns (BondingManager) {
        return new BondingPrimary(domain);
    }

    function _castToPrimary() internal view returns (BondingPrimary) {
        return BondingPrimary(address(bondingManager));
    }

    function _getTestLocalDomain() internal pure override returns (uint32) {
        return DOMAIN_SYNAPSE;
    }
}
