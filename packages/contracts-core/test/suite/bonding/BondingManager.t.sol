// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import "../../utils/SynapseTestSuite.t.sol";
import "../../tools/system/SystemContractTools.t.sol";
import { OriginMock } from "../../mocks/OriginMock.t.sol";
import { SystemContractMock } from "../../mocks/system/SystemContractMock.t.sol";

// solhint-disable func-name-mixedcase
abstract contract BondingManagerTest is SystemContractTools, SynapseTestSuite {
    OriginMock internal origin;
    SystemContractMock internal destination;
    BondingManager internal bondingManager;

    address[] internal systemRegistries;

    SystemRouterHarness internal systemRouter;

    uint32 internal localDomain;
    uint256 internal rootSubmittedAt;

    function setUp() public override {
        super.setUp();
        localDomain = _getTestLocalDomain();
        // Deploy mocks for tests
        origin = new OriginMock();
        destination = new SystemContractMock();
        bondingManager = _deployBondingManager(localDomain);
        systemRouter = new SystemRouterHarness(
            localDomain,
            address(origin),
            address(destination),
            address(bondingManager)
        );
        // Save Registries
        systemRegistries.push(address(origin));
        systemRegistries.push(address(destination));
        // Initialize
        bondingManager.initialize();
        // Set router
        origin.setSystemRouter(systemRouter);
        destination.setSystemRouter(systemRouter);
        bondingManager.setSystemRouter(systemRouter);
        // Transfer ownership
        bondingManager.transferOwnership(owner);
        // Label everything
        vm.label(address(origin), "Origin Mock");
        vm.label(address(destination), "Destination Mock");
        vm.label(address(bondingManager), "Bonding Manager");
        vm.label(address(systemRouter), "System Router");
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                             TESTS: SETUP                             ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function test_setup() public {
        assertEq(bondingManager.localDomain(), localDomain, "!localDomain");
        assertEq(address(bondingManager.systemRouter()), address(systemRouter), "!systemRouter");
        assertEq(bondingManager.owner(), owner, "!owner");
    }

    function test_initialize() public {
        bondingManager = _deployBondingManager(localDomain);
        assertEq(bondingManager.owner(), address(0), "owner existed pre initialize");
        bondingManager.initialize();
        assertEq(bondingManager.owner(), address(this), "failed to initialize owner");
    }

    function test_initialize_revert_onlyOnce() public {
        expectRevertAlreadyInitialized();
        bondingManager.initialize();
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║            TESTS: UNAUTHORIZED ACCESS (NOT SYSTEM ROUTER)            ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function test_slashAgent_revert_notSystemRouter(address caller) public {
        vm.assume(caller != address(systemRouter));
        AgentInfo memory info = guardInfo({ guard: address(0), bonded: false });
        vm.expectRevert("!systemRouter");
        vm.prank(caller);
        bondingManager.slashAgent(block.timestamp, DOMAIN_LOCAL, SystemEntity.BondingManager, info);
    }

    function test_syncAgents_revert_notSystemRouter(address caller) public {
        vm.assume(caller != address(systemRouter));
        vm.expectRevert("!systemRouter");
        vm.prank(caller);
        bondingManager.syncAgents(
            block.timestamp,
            DOMAIN_LOCAL,
            SystemEntity.BondingManager,
            0,
            false,
            new AgentInfo[](0)
        );
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                      TESTS: SLASH AGENT REVERTS                      ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function test_slashAgent_revert_localDomain_notOrigin() public {
        // Only Origin on local domain is allowed to call slashNotary()
        for (uint256 c = 0; c < uint8(type(SystemEntity).max); ++c) {
            SystemEntity caller = SystemEntity(c);
            // Should reject system calls from local domain, if caller is not Origin
            if (caller == SystemEntity.Origin) continue;
            vm.expectRevert("!allowedCaller");
            // Use mocked agent info
            _mockSlashAgentCall({
                callOrigin: localDomain,
                systemCaller: caller,
                info: guardInfo({ guard: address(0), bonded: false })
            });
        }
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                           INTERNAL HELPERS                           ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    // @notice Deploys relevant version of BondingManager for tests
    function _deployBondingManager(uint32 domain) internal virtual returns (BondingManager);

    function _skipBondingOptimisticPeriod() internal {
        _skipPeriod(BONDING_OPTIMISTIC_PERIOD);
    }

    function _skipPeriod(uint256 period) internal {
        rootSubmittedAt = block.timestamp;
        skip(period);
    }

    function _mockSyncAgentsCall(
        uint32 callOrigin,
        SystemEntity systemCaller,
        uint256 requestID,
        bool removeExisting,
        AgentInfo[] memory infos
    ) internal {
        systemRouter.mockSystemCall({
            _recipient: SystemEntity.BondingManager,
            _rootSubmittedAt: callOrigin == localDomain ? block.timestamp : rootSubmittedAt,
            _callOrigin: callOrigin,
            _systemCaller: systemCaller,
            _data: abi.encodeWithSelector(
                ISystemContract.syncAgents.selector,
                0, // rootSubmittedAt
                0, // callOrigin
                0, // systemCaller
                requestID,
                removeExisting,
                infos
            )
        });
    }

    function _mockSlashAgentCall(
        uint32 callOrigin,
        SystemEntity systemCaller,
        AgentInfo memory info
    ) internal {
        // Mock system call: slashAgent(rootSubmittedAt, callOrigin, systemCaller, info)
        systemRouter.mockSystemCall({
            _recipient: SystemEntity.BondingManager,
            _rootSubmittedAt: callOrigin == localDomain ? block.timestamp : rootSubmittedAt,
            _callOrigin: callOrigin,
            _systemCaller: systemCaller,
            _data: abi.encodeWithSelector(ISystemContract.slashAgent.selector, 0, 0, 0, info)
        });
    }

    function _getTestLocalDomain() internal pure virtual returns (uint32);
}
