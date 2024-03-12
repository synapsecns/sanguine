// SPDX-License-Identifier: MIT
pragma solidity 0.8.23;

import {InterchainERC20Test} from "./InterchainERC20.t.sol";
import {RateLimit} from "../../src/libs/RateLimit.sol";

import {IAccessControl} from "@openzeppelin/contracts/access/IAccessControl.sol";

// solhint-disable func-name-mixedcase
// solhint-disable ordering
contract InterchainERC20AdminTest is InterchainERC20Test {
    address public bridge;

    event BurnLimitSet(address indexed bridge, uint256 limit);
    event MintLimitSet(address indexed bridge, uint256 limit);

    function setUp() public override {
        super.setUp();
        bridge = makeAddr("Bridge");
    }

    // ═════════════════════════════════════════ TESTS: SETTING BURN LIMIT ═════════════════════════════════════════════

    function test_setTotalBurnLimit_limitNotSet() public {
        vm.expectEmit(address(token));
        emit BurnLimitSet(bridge, 1000);
        authSetTotalBurnLimit(bridge, 1000);
        assertEq(token.getCurrentBurnLimit(bridge), 1000);
        assertEq(token.getTotalBurnLimit(bridge), 1000);
        assertEq(token.getCurrentMintLimit(bridge), 0);
        assertEq(token.getTotalMintLimit(bridge), 0);
    }

    function test_setTotalBurnLimit_limitSet_fullySpent_increase() public {
        RateLimit memory initialLimit = RateLimit({lastUpdatedAt: block.timestamp, lastRemaining: 0, totalLimit: 1000});
        token.exposed__setBurnRateLimit(bridge, initialLimit);
        vm.expectEmit(address(token));
        emit BurnLimitSet(bridge, 2000);
        authSetTotalBurnLimit(bridge, 2000);
        // 2000 - 1000 = 1000
        assertEq(token.getCurrentBurnLimit(bridge), 1000);
        assertEq(token.getTotalBurnLimit(bridge), 2000);
        assertEq(token.getCurrentMintLimit(bridge), 0);
        assertEq(token.getTotalMintLimit(bridge), 0);
    }

    function test_setTotalBurnLimit_limitSet_fullySpent_decrease() public {
        RateLimit memory initialLimit = RateLimit({lastUpdatedAt: block.timestamp, lastRemaining: 0, totalLimit: 1000});
        token.exposed__setBurnRateLimit(bridge, initialLimit);
        vm.expectEmit(address(token));
        emit BurnLimitSet(bridge, 500);
        authSetTotalBurnLimit(bridge, 500);
        // 500 (total) < 1000 (already spent) => current limit is zero
        assertEq(token.getCurrentBurnLimit(bridge), 0);
        assertEq(token.getTotalBurnLimit(bridge), 500);
        assertEq(token.getCurrentMintLimit(bridge), 0);
        assertEq(token.getTotalMintLimit(bridge), 0);
    }

    function test_setTotalBurnLimit_limitSet_fullySpent_decrease_toZero() public {
        RateLimit memory initialLimit = RateLimit({lastUpdatedAt: block.timestamp, lastRemaining: 0, totalLimit: 1000});
        token.exposed__setBurnRateLimit(bridge, initialLimit);
        vm.expectEmit(address(token));
        emit BurnLimitSet(bridge, 0);
        authSetTotalBurnLimit(bridge, 0);
        assertEq(token.getCurrentBurnLimit(bridge), 0);
        assertEq(token.getTotalBurnLimit(bridge), 0);
        assertEq(token.getCurrentMintLimit(bridge), 0);
        assertEq(token.getTotalMintLimit(bridge), 0);
    }

    function test_setTotalBurnLimit_limitSet_notSpent_increase() public {
        RateLimit memory initialLimit =
            RateLimit({lastUpdatedAt: block.timestamp, lastRemaining: 1000, totalLimit: 1000});
        token.exposed__setBurnRateLimit(bridge, initialLimit);
        vm.expectEmit(address(token));
        emit BurnLimitSet(bridge, 2000);
        authSetTotalBurnLimit(bridge, 2000);
        assertEq(token.getCurrentBurnLimit(bridge), 2000);
        assertEq(token.getTotalBurnLimit(bridge), 2000);
        assertEq(token.getCurrentMintLimit(bridge), 0);
        assertEq(token.getTotalMintLimit(bridge), 0);
    }

    function test_setTotalBurnLimit_limitSet_notSpent_decrease() public {
        RateLimit memory initialLimit =
            RateLimit({lastUpdatedAt: block.timestamp, lastRemaining: 1000, totalLimit: 1000});
        token.exposed__setBurnRateLimit(bridge, initialLimit);
        vm.expectEmit(address(token));
        emit BurnLimitSet(bridge, 500);
        authSetTotalBurnLimit(bridge, 500);
        assertEq(token.getCurrentBurnLimit(bridge), 500);
        assertEq(token.getTotalBurnLimit(bridge), 500);
        assertEq(token.getCurrentMintLimit(bridge), 0);
        assertEq(token.getTotalMintLimit(bridge), 0);
    }

    function test_setTotalBurnLimit_limitSet_notSpent_decrease_toZero() public {
        RateLimit memory initialLimit =
            RateLimit({lastUpdatedAt: block.timestamp, lastRemaining: 1000, totalLimit: 1000});
        token.exposed__setBurnRateLimit(bridge, initialLimit);
        vm.expectEmit(address(token));
        emit BurnLimitSet(bridge, 0);
        authSetTotalBurnLimit(bridge, 0);
        assertEq(token.getCurrentBurnLimit(bridge), 0);
        assertEq(token.getTotalBurnLimit(bridge), 0);
        assertEq(token.getCurrentMintLimit(bridge), 0);
        assertEq(token.getTotalMintLimit(bridge), 0);
    }

    function test_setTotalBurnLimit_limitSet_partiallySpent_increase() public {
        RateLimit memory initialLimit =
            RateLimit({lastUpdatedAt: block.timestamp, lastRemaining: 100, totalLimit: 1000});
        token.exposed__setBurnRateLimit(bridge, initialLimit);
        // Skip 12 hours => 500 replenished, a total of 600 available
        // e.g. 400 token were spent from the limit
        skip(12 hours);
        vm.expectEmit(address(token));
        emit BurnLimitSet(bridge, 2000);
        authSetTotalBurnLimit(bridge, 2000);
        // 2000 - 400 = 1600
        assertEq(token.getCurrentBurnLimit(bridge), 1600);
        assertEq(token.getTotalBurnLimit(bridge), 2000);
        assertEq(token.getCurrentMintLimit(bridge), 0);
        assertEq(token.getTotalMintLimit(bridge), 0);
    }

    function test_setTotalBurnLimit_limitSet_partiallySpent_decrease() public {
        RateLimit memory initialLimit =
            RateLimit({lastUpdatedAt: block.timestamp, lastRemaining: 100, totalLimit: 1000});
        token.exposed__setBurnRateLimit(bridge, initialLimit);
        // Skip 12 hours => 500 replenished, a total of 600 available
        // e.g. 400 token were spent from the limit
        skip(12 hours);
        vm.expectEmit(address(token));
        emit BurnLimitSet(bridge, 500);
        authSetTotalBurnLimit(bridge, 500);
        // 500 - 400 = 100
        assertEq(token.getCurrentBurnLimit(bridge), 100);
        assertEq(token.getTotalBurnLimit(bridge), 500);
        assertEq(token.getCurrentMintLimit(bridge), 0);
        assertEq(token.getTotalMintLimit(bridge), 0);
    }

    function test_setTotalBurnLimit_limitSet_partiallySpent_decrease_underSpentAmount() public {
        RateLimit memory initialLimit =
            RateLimit({lastUpdatedAt: block.timestamp, lastRemaining: 100, totalLimit: 1000});
        token.exposed__setBurnRateLimit(bridge, initialLimit);
        // Skip 12 hours => 500 replenished, a total of 600 available
        // e.g. 400 token were spent from the limit
        skip(12 hours);
        vm.expectEmit(address(token));
        emit BurnLimitSet(bridge, 300);
        authSetTotalBurnLimit(bridge, 300);
        // 300 (total) < 400 (already spent) => current limit is zero
        assertEq(token.getCurrentBurnLimit(bridge), 0);
        assertEq(token.getTotalBurnLimit(bridge), 300);
        assertEq(token.getCurrentMintLimit(bridge), 0);
        assertEq(token.getTotalMintLimit(bridge), 0);
    }

    function test_setTotalBurnLimit_limitSet_partiallySpent_decrease_toZero() public {
        RateLimit memory initialLimit =
            RateLimit({lastUpdatedAt: block.timestamp, lastRemaining: 100, totalLimit: 1000});
        token.exposed__setBurnRateLimit(bridge, initialLimit);
        skip(12 hours);
        vm.expectEmit(address(token));
        emit BurnLimitSet(bridge, 0);
        authSetTotalBurnLimit(bridge, 0);
        assertEq(token.getCurrentBurnLimit(bridge), 0);
        assertEq(token.getTotalBurnLimit(bridge), 0);
        assertEq(token.getCurrentMintLimit(bridge), 0);
        assertEq(token.getTotalMintLimit(bridge), 0);
    }

    // ═════════════════════════════════════════ TESTS: SETTING MINT LIMIT ═════════════════════════════════════════════

    function test_setTotalMintLimit_limitNotSet() public {
        vm.expectEmit(address(token));
        emit MintLimitSet(bridge, 1000);
        authSetTotalMintLimit(bridge, 1000);
        assertEq(token.getCurrentMintLimit(bridge), 1000);
        assertEq(token.getTotalMintLimit(bridge), 1000);
        assertEq(token.getCurrentBurnLimit(bridge), 0);
        assertEq(token.getTotalBurnLimit(bridge), 0);
    }

    function test_setTotalMintLimit_limitSet_fullySpent_increase() public {
        RateLimit memory initialLimit = RateLimit({lastUpdatedAt: block.timestamp, lastRemaining: 0, totalLimit: 1000});
        token.exposed__setMintRateLimit(bridge, initialLimit);
        vm.expectEmit(address(token));
        emit MintLimitSet(bridge, 2000);
        authSetTotalMintLimit(bridge, 2000);
        // 2000 - 1000 = 1000
        assertEq(token.getCurrentMintLimit(bridge), 1000);
        assertEq(token.getTotalMintLimit(bridge), 2000);
        assertEq(token.getCurrentBurnLimit(bridge), 0);
        assertEq(token.getTotalBurnLimit(bridge), 0);
    }

    function test_setTotalMintLimit_limitSet_fullySpent_decrease() public {
        RateLimit memory initialLimit = RateLimit({lastUpdatedAt: block.timestamp, lastRemaining: 0, totalLimit: 1000});
        token.exposed__setMintRateLimit(bridge, initialLimit);
        vm.expectEmit(address(token));
        emit MintLimitSet(bridge, 500);
        authSetTotalMintLimit(bridge, 500);
        // 500 (total) < 1000 (already spent) => current limit is zero
        assertEq(token.getCurrentMintLimit(bridge), 0);
        assertEq(token.getTotalMintLimit(bridge), 500);
        assertEq(token.getCurrentBurnLimit(bridge), 0);
        assertEq(token.getTotalBurnLimit(bridge), 0);
    }

    function test_setTotalMintLimit_limitSet_fullySpent_decrease_toZero() public {
        RateLimit memory initialLimit = RateLimit({lastUpdatedAt: block.timestamp, lastRemaining: 0, totalLimit: 1000});
        token.exposed__setMintRateLimit(bridge, initialLimit);
        vm.expectEmit(address(token));
        emit MintLimitSet(bridge, 0);
        authSetTotalMintLimit(bridge, 0);
        assertEq(token.getCurrentMintLimit(bridge), 0);
        assertEq(token.getTotalMintLimit(bridge), 0);
        assertEq(token.getCurrentBurnLimit(bridge), 0);
        assertEq(token.getTotalBurnLimit(bridge), 0);
    }

    function test_setTotalMintLimit_limitSet_notSpent_increase() public {
        RateLimit memory initialLimit =
            RateLimit({lastUpdatedAt: block.timestamp, lastRemaining: 1000, totalLimit: 1000});
        token.exposed__setMintRateLimit(bridge, initialLimit);
        vm.expectEmit(address(token));
        emit MintLimitSet(bridge, 2000);
        authSetTotalMintLimit(bridge, 2000);
        assertEq(token.getCurrentMintLimit(bridge), 2000);
        assertEq(token.getTotalMintLimit(bridge), 2000);
        assertEq(token.getCurrentBurnLimit(bridge), 0);
        assertEq(token.getTotalBurnLimit(bridge), 0);
    }

    function test_setTotalMintLimit_limitSet_notSpent_decrease() public {
        RateLimit memory initialLimit =
            RateLimit({lastUpdatedAt: block.timestamp, lastRemaining: 1000, totalLimit: 1000});
        token.exposed__setMintRateLimit(bridge, initialLimit);
        vm.expectEmit(address(token));
        emit MintLimitSet(bridge, 500);
        authSetTotalMintLimit(bridge, 500);
        assertEq(token.getCurrentMintLimit(bridge), 500);
        assertEq(token.getTotalMintLimit(bridge), 500);
        assertEq(token.getCurrentBurnLimit(bridge), 0);
        assertEq(token.getTotalBurnLimit(bridge), 0);
    }

    function test_setTotalMintLimit_limitSet_notSpent_decrease_toZero() public {
        RateLimit memory initialLimit =
            RateLimit({lastUpdatedAt: block.timestamp, lastRemaining: 1000, totalLimit: 1000});
        token.exposed__setMintRateLimit(bridge, initialLimit);
        vm.expectEmit(address(token));
        emit MintLimitSet(bridge, 0);
        authSetTotalMintLimit(bridge, 0);
        assertEq(token.getCurrentMintLimit(bridge), 0);
        assertEq(token.getTotalMintLimit(bridge), 0);
        assertEq(token.getCurrentBurnLimit(bridge), 0);
        assertEq(token.getTotalBurnLimit(bridge), 0);
    }

    function test_setTotalMintLimit_limitSet_partiallySpent_increase() public {
        RateLimit memory initialLimit =
            RateLimit({lastUpdatedAt: block.timestamp, lastRemaining: 100, totalLimit: 1000});
        token.exposed__setMintRateLimit(bridge, initialLimit);
        // Skip 12 hours => 500 replenished, a total of 600 available
        // e.g. 400 token were spent from the limit
        skip(12 hours);
        vm.expectEmit(address(token));
        emit MintLimitSet(bridge, 2000);
        authSetTotalMintLimit(bridge, 2000);
        // 2000 - 400 = 1600
        assertEq(token.getCurrentMintLimit(bridge), 1600);
        assertEq(token.getTotalMintLimit(bridge), 2000);
        assertEq(token.getCurrentBurnLimit(bridge), 0);
        assertEq(token.getTotalBurnLimit(bridge), 0);
    }

    function test_setTotalMintLimit_limitSet_partiallySpent_decrease() public {
        RateLimit memory initialLimit =
            RateLimit({lastUpdatedAt: block.timestamp, lastRemaining: 100, totalLimit: 1000});
        token.exposed__setMintRateLimit(bridge, initialLimit);
        // Skip 12 hours => 500 replenished, a total of 600 available
        // e.g. 400 token were spent from the limit
        skip(12 hours);
        vm.expectEmit(address(token));
        emit MintLimitSet(bridge, 500);
        authSetTotalMintLimit(bridge, 500);
        // 500 - 400 = 100
        assertEq(token.getCurrentMintLimit(bridge), 100);
        assertEq(token.getTotalMintLimit(bridge), 500);
        assertEq(token.getCurrentBurnLimit(bridge), 0);
        assertEq(token.getTotalBurnLimit(bridge), 0);
    }

    function test_setTotalMintLimit_limitSet_partiallySpent_decrease_underSpentAmount() public {
        RateLimit memory initialLimit =
            RateLimit({lastUpdatedAt: block.timestamp, lastRemaining: 100, totalLimit: 1000});
        token.exposed__setMintRateLimit(bridge, initialLimit);
        // Skip 12 hours => 500 replenished, a total of 600 available
        // e.g. 400 token were spent from the limit
        skip(12 hours);
        vm.expectEmit(address(token));
        emit MintLimitSet(bridge, 300);
        authSetTotalMintLimit(bridge, 300);
        // 300 (total) < 400 (already spent) => current limit is zero
        assertEq(token.getCurrentMintLimit(bridge), 0);
        assertEq(token.getTotalMintLimit(bridge), 300);
        assertEq(token.getCurrentBurnLimit(bridge), 0);
        assertEq(token.getTotalBurnLimit(bridge), 0);
    }

    function test_setTotalMintLimit_limitSet_partiallySpent_decrease_toZero() public {
        RateLimit memory initialLimit =
            RateLimit({lastUpdatedAt: block.timestamp, lastRemaining: 100, totalLimit: 1000});
        token.exposed__setMintRateLimit(bridge, initialLimit);
        skip(12 hours);
        vm.expectEmit(address(token));
        emit MintLimitSet(bridge, 0);
        authSetTotalMintLimit(bridge, 0);
        assertEq(token.getCurrentMintLimit(bridge), 0);
        assertEq(token.getTotalMintLimit(bridge), 0);
        assertEq(token.getCurrentBurnLimit(bridge), 0);
        assertEq(token.getTotalBurnLimit(bridge), 0);
    }

    // ════════════════════════════════════════════ TESTS: UNAUTHORIZED ════════════════════════════════════════════════

    function test_grantRole_revert_notAdmin(address unauthorizedActor) public {
        vm.assume(unauthorizedActor != admin);
        bytes32 role = token.EMERGENCY_PAUSER_ROLE();
        vm.prank(unauthorizedActor);
        vm.expectRevert(
            abi.encodeWithSelector(IAccessControl.AccessControlUnauthorizedAccount.selector, unauthorizedActor, 0)
        );
        token.grantRole(role, address(1337));
    }

    function test_setTotalBurnLimit_revert_notGovernor(address unauthorizedActor) public {
        vm.assume(unauthorizedActor != governor);
        bytes32 requiredRole = token.GOVERNOR_ROLE();
        vm.prank(unauthorizedActor);
        vm.expectRevert(
            abi.encodeWithSelector(
                IAccessControl.AccessControlUnauthorizedAccount.selector, unauthorizedActor, requiredRole
            )
        );
        token.setTotalBurnLimit(address(1337), 1);
    }

    function test_setTotalMintLimit_revert_notGovernor(address unauthorizedActor) public {
        vm.assume(unauthorizedActor != governor);
        bytes32 requiredRole = token.GOVERNOR_ROLE();
        vm.prank(unauthorizedActor);
        vm.expectRevert(
            abi.encodeWithSelector(
                IAccessControl.AccessControlUnauthorizedAccount.selector, unauthorizedActor, requiredRole
            )
        );
        token.setTotalMintLimit(address(1337), 1);
    }

    function test_pause_revert_notEmergencyPauser(address unauthorizedActor) public {
        vm.assume(unauthorizedActor != emergencyPauser);
        bytes32 requiredRole = token.EMERGENCY_PAUSER_ROLE();
        vm.prank(unauthorizedActor);
        vm.expectRevert(
            abi.encodeWithSelector(
                IAccessControl.AccessControlUnauthorizedAccount.selector, unauthorizedActor, requiredRole
            )
        );
        token.pause();
    }

    function test_unpause_revert_notEmergencyPauser(address unauthorizedActor) public {
        vm.assume(unauthorizedActor != emergencyPauser);
        authPause();
        bytes32 requiredRole = token.EMERGENCY_PAUSER_ROLE();
        vm.prank(unauthorizedActor);
        vm.expectRevert(
            abi.encodeWithSelector(
                IAccessControl.AccessControlUnauthorizedAccount.selector, unauthorizedActor, requiredRole
            )
        );
        token.unpause();
    }
}
