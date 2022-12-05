// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import "../../utils/SynapseTestSuite.t.sol";
import "../../tools/system/SystemContractTools.t.sol";
import { SystemRegistryMock } from "../../mocks/system/SystemRegistryMock.t.sol";

// solhint-disable func-name-mixedcase
contract SystemRegistryTest is SystemContractTools, SynapseTestSuite {
    SystemRegistryMock internal systemRegistry;
    SystemRouterHarness internal systemRouter;

    function setUp() public override {
        super.setUp();
        systemRegistry = new SystemRegistryMock(DOMAIN_LOCAL);
        systemRegistry.initialize();
        // Deploy a mock for System Router, with deployed registry as "origin"
        systemRouter = new SystemRouterHarness({
            _domain: DOMAIN_LOCAL,
            _origin: address(systemRegistry),
            _destination: address(0),
            _bondingManager: address(0)
        });
        systemRegistry.setSystemRouter(systemRouter);
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                   TESTS: CALLED BY BONDING MANAGER                   ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function test_bondNotary() public {
        for (uint256 d = 0; d < DOMAINS; ++d) {
            uint32 domain = domains[d];
            address notary = suiteNotary(domain);
            SystemContract.AgentInfo[] memory infos = infoToArray(
                agentInfo({ domain: domain, account: notary, bonded: true })
            );
            assertFalse(systemRegistry.isNotary(domain, notary), "!bondNotary: already added");
            _mockValidSyncAgentsCall(infos);
            assertTrue(systemRegistry.isNotary(domain, notary), "!bondNotary: not added");
        }
    }

    function test_unbondNotary() public {
        test_bondNotary();
        for (uint256 d = 0; d < DOMAINS; ++d) {
            uint32 domain = domains[d];
            address notary = suiteNotary(domain);
            SystemContract.AgentInfo[] memory infos = infoToArray(
                agentInfo({ domain: domain, account: notary, bonded: false })
            );
            assertTrue(systemRegistry.isNotary(domain, notary), "!unbondNotary: not active");
            _mockValidSyncAgentsCall(infos);
            assertFalse(systemRegistry.isNotary(domain, notary), "!unbondNotary: not removed");
        }
    }

    function test_slashNotary() public {
        test_bondNotary();
        for (uint256 d = 0; d < DOMAINS; ++d) {
            uint32 domain = domains[d];
            address notary = suiteNotary(domain);
            SystemContract.AgentInfo memory info = agentInfo({
                domain: domain,
                account: notary,
                bonded: false
            });
            assertTrue(systemRegistry.isNotary(domain, notary), "!slashNotary: not active");
            vm.expectEmit(true, true, true, true, address(systemRegistry));
            emit SlashAgentCall(info);
            _mockValidSlashAgentCall(info);
            assertFalse(systemRegistry.isNotary(domain, notary), "!slashNotary: not removed");
        }
    }

    function test_bondGuard() public {
        address guard = suiteGuard();
        SystemContract.AgentInfo[] memory infos = infoToArray(
            guardInfo({ guard: guard, bonded: true })
        );
        assertFalse(systemRegistry.isGuard(guard), "!bondGuard: already added");
        _mockValidSyncAgentsCall(infos);
        assertTrue(systemRegistry.isGuard(guard), "!bondGuard: not added");
    }

    function test_unbondGuard() public {
        test_bondGuard();
        address guard = suiteGuard();
        SystemContract.AgentInfo[] memory infos = infoToArray(
            guardInfo({ guard: guard, bonded: false })
        );
        assertTrue(systemRegistry.isGuard(guard), "!unbondGuard: not added");
        _mockValidSyncAgentsCall(infos);
        assertFalse(systemRegistry.isGuard(guard), "!unbondGuard: not removed");
    }

    function test_slashGuard() public {
        test_bondGuard();
        address guard = suiteGuard();
        SystemContract.AgentInfo memory info = guardInfo({ guard: guard, bonded: false });
        assertTrue(systemRegistry.isGuard(guard), "!slashGuard: not added");
        vm.expectEmit(true, true, true, true, address(systemRegistry));
        emit SlashAgentCall(info);
        _mockValidSlashAgentCall(info);
        assertFalse(systemRegistry.isGuard(guard), "!slashGuard: not removed");
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║            TESTS: UNAUTHORIZED ACCESS (NOT SYSTEM ROUTER)            ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function test_syncAgents_revert_notSystemRouter(address caller) public {
        vm.assume(caller != address(systemRouter));
        vm.expectRevert("!systemRouter");
        vm.prank(caller);
        systemRegistry.syncAgents(
            block.timestamp,
            DOMAIN_LOCAL,
            ISystemRouter.SystemEntity.BondingManager,
            0,
            false,
            new SystemContract.AgentInfo[](0)
        );
    }

    function test_slashAgent_revert_notSystemRouter(address caller) public {
        vm.assume(caller != address(systemRouter));
        SystemContract.AgentInfo memory info = guardInfo({ guard: address(0), bonded: false });
        vm.expectRevert("!systemRouter");
        vm.prank(caller);
        systemRegistry.slashAgent(
            block.timestamp,
            DOMAIN_LOCAL,
            ISystemRouter.SystemEntity.BondingManager,
            info
        );
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║            TESTS: UNAUTHORIZED ACCESS (NOT LOCAL DOMAIN)             ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function test_syncAgents_revert_notLocalDomain(uint32 domain) public {
        vm.assume(domain != DOMAIN_LOCAL);
        // Should reject system call from other domains (regardless of the caller)
        for (uint256 c = 0; c < uint8(type(ISystemRouter.SystemEntity).max); ++c) {
            ISystemRouter.SystemEntity caller = ISystemRouter.SystemEntity(c);
            vm.expectRevert("!localDomain");
            _mockRevertedSyncAgentsCall({ callOrigin: domain, systemCaller: caller });
        }
    }

    function test_slashAgent_revert_notLocalDomain(uint32 domain) public {
        vm.assume(domain != DOMAIN_LOCAL);
        // Should reject system call from other domains (regardless of the caller)
        for (uint256 c = 0; c < uint8(type(ISystemRouter.SystemEntity).max); ++c) {
            ISystemRouter.SystemEntity caller = ISystemRouter.SystemEntity(c);
            vm.expectRevert("!localDomain");
            _mockRevertedSlashAgentCall({ callOrigin: domain, systemCaller: caller });
        }
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║           TESTS: UNAUTHORIZED ACCESS (NOT BONDING MANAGER)           ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function test_syncAgents_revert_localDomain_notBondingManager() public {
        for (uint256 c = 0; c < uint8(type(ISystemRouter.SystemEntity).max); ++c) {
            // Should reject system calls from local domain, if caller is not BondingManager
            if (c == uint8(ISystemRouter.SystemEntity.BondingManager)) continue;
            ISystemRouter.SystemEntity caller = ISystemRouter.SystemEntity(c);
            vm.expectRevert("!allowedCaller");
            _mockRevertedSyncAgentsCall({ callOrigin: DOMAIN_LOCAL, systemCaller: caller });
        }
    }

    function test_slashAgent_revert_localDomain_notBondingManager() public {
        for (uint256 c = 0; c < uint8(type(ISystemRouter.SystemEntity).max); ++c) {
            // Should reject system calls from local domain, if caller is not BondingManager
            if (c == uint8(ISystemRouter.SystemEntity.BondingManager)) continue;
            ISystemRouter.SystemEntity caller = ISystemRouter.SystemEntity(c);
            vm.expectRevert("!allowedCaller");
            _mockRevertedSlashAgentCall({ callOrigin: DOMAIN_LOCAL, systemCaller: caller });
        }
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                           INTERNAL HELPERS                           ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function _mockValidSyncAgentsCall(SystemContract.AgentInfo[] memory infos) internal {
        // Mock system call from local Bonding Manager by default
        _mockSyncAgentsCall({
            callOrigin: DOMAIN_LOCAL,
            systemCaller: ISystemRouter.SystemEntity.BondingManager,
            infos: infos
        });
    }

    function _mockRevertedSyncAgentsCall(uint32 callOrigin, ISystemRouter.SystemEntity systemCaller)
        internal
    {
        // Mock Agent data for the "revert tests"
        _mockSyncAgentsCall({
            callOrigin: callOrigin,
            systemCaller: systemCaller,
            infos: infoToArray(guardInfo({ guard: address(0), bonded: false }))
        });
    }

    function _mockSyncAgentsCall(
        uint32 callOrigin,
        ISystemRouter.SystemEntity systemCaller,
        SystemContract.AgentInfo[] memory infos
    ) internal {
        // TODO: add coverage when these params are no longer ignored in SystemRegistry
        uint256 requestID = 0;
        bool removeExisting = false;
        systemRouter.mockSystemCall({
            _recipient: ISystemRouter.SystemEntity.Origin,
            _rootSubmittedAt: block.timestamp,
            _callOrigin: callOrigin,
            _systemCaller: systemCaller,
            _data: abi.encodeWithSelector(
                SystemContract.syncAgents.selector,
                0, // rootSubmittedAt
                0, // callOrigin
                0, // systemCaller
                requestID,
                removeExisting,
                infos
            )
        });
    }

    function _mockValidSlashAgentCall(SystemContract.AgentInfo memory info) internal {
        // Mock system call from local Bonding Manager by default
        _mockSlashAgentCall({
            callOrigin: DOMAIN_LOCAL,
            systemCaller: ISystemRouter.SystemEntity.BondingManager,
            info: info
        });
    }

    function _mockRevertedSlashAgentCall(uint32 callOrigin, ISystemRouter.SystemEntity systemCaller)
        internal
    {
        // Mock Agent data for the "revert tests"
        _mockSlashAgentCall({
            callOrigin: callOrigin,
            systemCaller: systemCaller,
            info: guardInfo({ guard: address(0), bonded: false })
        });
    }

    function _mockSlashAgentCall(
        uint32 callOrigin,
        ISystemRouter.SystemEntity systemCaller,
        SystemContract.AgentInfo memory info
    ) internal {
        // Mock system call: slashAgent(rootSubmittedAt, callOrigin, systemCaller, info)
        systemRouter.mockSystemCall({
            _recipient: ISystemRouter.SystemEntity.Origin,
            _rootSubmittedAt: block.timestamp,
            _callOrigin: callOrigin,
            _systemCaller: systemCaller,
            _data: abi.encodeWithSelector(SystemContract.slashAgent.selector, 0, 0, 0, info)
        });
    }
}
