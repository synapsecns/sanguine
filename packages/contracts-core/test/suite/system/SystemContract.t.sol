// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import "../../tools/system/SystemRouterTools.t.sol";
import { SystemContractMock } from "../../mocks/system/SystemContractMock.t.sol";

// solhint-disable func-name-mixedcase
contract SystemContractTest is SystemRouterTools {
    SystemContractMock internal systemContract;

    function setUp() public override {
        super.setUp();
        systemContract = new SystemContractMock(DOMAIN_LOCAL);
        systemContract.initialize();
        systemContract.transferOwnership(owner);
    }

    function test_setup() public {
        assertEq(systemContract.owner(), owner, "!owner");
        assertEq(systemContract.localDomain(), DOMAIN_LOCAL, "!localDomain");
        assertEq(systemContract.SYNAPSE_DOMAIN(), DOMAIN_SYNAPSE, "!synapseDomain");
        assertEq(
            systemContract.ORIGIN_MASK(),
            1 << uint256(ISystemRouter.SystemEntity.Origin),
            "!originMask"
        );
        assertEq(
            systemContract.DESTINATION_MASK(),
            1 << uint256(ISystemRouter.SystemEntity.Destination),
            "!destinationMask"
        );
        assertEq(
            systemContract.BONDING_MANAGER_MASK(),
            1 << uint256(ISystemRouter.SystemEntity.BondingManager),
            "!bondingManagerMask"
        );
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                  TESTS: RESTRICTED ACCESS (REVERTS)                  ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function test_setSystemRouter_revert_notOwner(address caller) public {
        vm.assume(caller != owner);
        expectRevertNotOwner();
        vm.prank(caller);
        systemContract.setSystemRouter(suiteSystemRouter(DOMAIN_LOCAL));
    }

    function test_renounceOwnership_revert_notOwner(address caller) public {
        vm.assume(caller != owner);
        expectRevertNotOwner();
        vm.prank(caller);
        systemContract.renounceOwnership();
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                       TESTS: RESTRICTED ACCESS                       ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function test_setSystemRouter() public {
        vm.prank(owner);
        systemContract.setSystemRouter(suiteSystemRouter(DOMAIN_LOCAL));
        assertEq(
            address(systemContract.systemRouter()),
            address(suiteSystemRouter(DOMAIN_LOCAL)),
            "Failed to set system router"
        );
    }

    function test_renounceOwnership_doesNothing() public {
        vm.prank(owner);
        systemContract.renounceOwnership();
        assertEq(systemContract.owner(), owner, "!owner");
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                           TESTS: MODIFIERS                           ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function test_onlySystemRouter() public {
        test_setSystemRouter();
        vm.prank(address(suiteSystemRouter(DOMAIN_LOCAL)));
        systemContract.mockOnlySystemRouter();
    }

    function test_onlySystemRouter_revert_notSystemRouter(address caller) public {
        vm.assume(caller != address(suiteSystemRouter(DOMAIN_LOCAL)));
        test_setSystemRouter();
        vm.expectRevert("!systemRouter");
        vm.prank(caller);
        systemContract.mockOnlySystemRouter();
    }

    function test_onlySynapseChain() public {
        systemContract.mockOnlySynapseChain(DOMAIN_SYNAPSE);
    }

    function test_onlySynapseChain(uint32 domain) public {
        vm.assume(domain != DOMAIN_SYNAPSE);
        vm.expectRevert("!synapseDomain");
        systemContract.mockOnlySynapseChain(domain);
    }

    function test_onlyCallers_onlyOthers() public {
        // 111...11100 excludes Origin, Destination and nothing else
        uint256 mask = type(uint256).max ^ 0x3;
        vm.expectRevert("!allowedCaller");
        systemContract.mockOnlyCallers(mask, ISystemRouter.SystemEntity.Origin);
        vm.expectRevert("!allowedCaller");
        systemContract.mockOnlyCallers(mask, ISystemRouter.SystemEntity.Destination);
    }

    function test_onlyCallers_onlyOrigin() public {
        uint256 mask = systemContract.ORIGIN_MASK();
        systemContract.mockOnlyCallers(mask, ISystemRouter.SystemEntity.Origin);
        vm.expectRevert("!allowedCaller");
        systemContract.mockOnlyCallers(mask, ISystemRouter.SystemEntity.Destination);
    }

    function test_onlyCallers_onlyDestination() public {
        uint256 mask = systemContract.DESTINATION_MASK();
        vm.expectRevert("!allowedCaller");
        systemContract.mockOnlyCallers(mask, ISystemRouter.SystemEntity.Origin);
        systemContract.mockOnlyCallers(mask, ISystemRouter.SystemEntity.Destination);
    }

    function test_onlyCallers_onlyOriginDestination() public {
        uint256 mask = systemContract.ORIGIN_MASK() | systemContract.DESTINATION_MASK();
        systemContract.mockOnlyCallers(mask, ISystemRouter.SystemEntity.Origin);
        systemContract.mockOnlyCallers(mask, ISystemRouter.SystemEntity.Destination);
    }

    function test_onlyOptimisticPeriodOver() public {
        rootSubmittedAt = block.timestamp;
        uint256 period = 1 days;
        skip(period);
        systemContract.mockOnlyOptimisticPeriodOver(rootSubmittedAt, period);
        skip(1);
        systemContract.mockOnlyOptimisticPeriodOver(rootSubmittedAt, period);
    }

    function test_onlyOptimisticPeriodOver_revert_tooEarly() public {
        rootSubmittedAt = block.timestamp;
        uint256 period = 1 days;

        // 0 seconds passed => revert
        vm.expectRevert("!optimisticPeriod");
        systemContract.mockOnlyOptimisticPeriodOver(rootSubmittedAt, period);

        // Half period passed => revert
        skip(period / 2);
        vm.expectRevert("!optimisticPeriod");
        systemContract.mockOnlyOptimisticPeriodOver(rootSubmittedAt, period);

        // Period minus 1 second passed => revert
        skip(period / 2 - 1);
        vm.expectRevert("!optimisticPeriod");
        systemContract.mockOnlyOptimisticPeriodOver(rootSubmittedAt, period);

        // Period passed
        skip(1);
        systemContract.mockOnlyOptimisticPeriodOver(rootSubmittedAt, period);
    }

    function test_onlyOptimisticPeriodOver_revert_futureRoot() public {
        rootSubmittedAt = block.timestamp + 1 days;
        uint256 period = 1;
        // Root submission is 1 day into the future
        vm.expectRevert("!optimisticPeriod");
        systemContract.mockOnlyOptimisticPeriodOver(rootSubmittedAt, period);

        // Root submission is "now", period not passed
        skip(1 days);
        vm.expectRevert("!optimisticPeriod");
        systemContract.mockOnlyOptimisticPeriodOver(rootSubmittedAt, period);

        // Period passed
        skip(1);
        // Should now pass
        systemContract.mockOnlyOptimisticPeriodOver(rootSubmittedAt, period);
    }
}
