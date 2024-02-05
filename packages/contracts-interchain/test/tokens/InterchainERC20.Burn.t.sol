// SPDX-License-Identifier: MIT
pragma solidity 0.8.23;

import {InterchainERC20Test} from "./InterchainERC20.t.sol";
import {RateLimiting} from "../../src/libs/RateLimit.sol";

import {Pausable} from "@openzeppelin/contracts/utils/Pausable.sol";

// solhint-disable func-name-mixedcase
// solhint-disable ordering
contract InterchainERC20BurnTest is InterchainERC20Test {
    address public bridge;

    uint256 public constant INITIAL_TOTAL_LIMIT = 1000;
    uint256 public constant INITIAL_BURNED = 400;
    uint256 public constant INITIAL_CURRENT_LIMIT = INITIAL_TOTAL_LIMIT - INITIAL_BURNED;

    uint256 public constant SMALL_PERIOD = 1 days / 10;
    uint256 public constant SMALL_PERIOD_CURRENT_LIMIT = INITIAL_CURRENT_LIMIT + 100;
    uint256 public constant LARGE_PERIOD = 1 days * 2;

    uint256 public constant INITIAL_BRIDGE_BALANCE = 1000;
    uint256 public constant INITIAL_PROCESSOR_BALANCE = 10_000;
    uint256 public constant INITIAL_TOTAL_SUPPLY = INITIAL_BRIDGE_BALANCE + INITIAL_PROCESSOR_BALANCE;

    function setUp() public override {
        super.setUp();
        bridge = makeAddr("Bridge");
        // Set the total burn limit for the bridge and spend some of it
        authSetTotalBurnLimit(bridge, INITIAL_TOTAL_LIMIT);
        processorMintToken(bridge, INITIAL_BRIDGE_BALANCE + INITIAL_BURNED);
        processorMintToken(processor, INITIAL_PROCESSOR_BALANCE);
        authBurnToken(INITIAL_BURNED);
    }

    function processorMintToken(address account, uint256 amount) public {
        vm.prank(processor);
        token.mint(account, amount);
    }

    function authBurnToken(uint256 amount) public {
        vm.prank(bridge);
        token.burn(amount);
    }

    function test_setUp() public {
        assertEq(token.getCurrentBurnLimit(bridge), INITIAL_CURRENT_LIMIT);
        assertEq(token.getTotalBurnLimit(bridge), INITIAL_TOTAL_LIMIT);
        assertEq(token.balanceOf(bridge), INITIAL_BRIDGE_BALANCE);
        assertEq(token.balanceOf(processor), INITIAL_PROCESSOR_BALANCE);
        assertEq(token.totalSupply(), INITIAL_TOTAL_SUPPLY);
    }

    function test_burn_revert_unauthorized() public {
        address user = makeAddr("User");
        processorMintToken(user, 100);
        vm.expectRevert(abi.encodeWithSelector(RateLimiting.RateLimiting__LimitExceeded.selector, 100, 0));
        vm.prank(user);
        token.burn(100);
    }

    function test_burn_zeroTimePassed_burnUnderLimit() public {
        uint256 amount = INITIAL_CURRENT_LIMIT / 10;
        authBurnToken(amount);
        assertEq(token.balanceOf(bridge), INITIAL_BRIDGE_BALANCE - amount);
        assertEq(token.totalSupply(), INITIAL_TOTAL_SUPPLY - amount);
        assertEq(token.getCurrentBurnLimit(bridge), INITIAL_CURRENT_LIMIT - amount);
        assertEq(token.getTotalBurnLimit(bridge), INITIAL_TOTAL_LIMIT);
    }

    function test_burn_zeroTimePassed_burnExactlyLimit() public {
        authBurnToken(INITIAL_CURRENT_LIMIT);
        assertEq(token.balanceOf(bridge), INITIAL_BRIDGE_BALANCE - INITIAL_CURRENT_LIMIT);
        assertEq(token.totalSupply(), INITIAL_TOTAL_SUPPLY - INITIAL_CURRENT_LIMIT);
        assertEq(token.getCurrentBurnLimit(bridge), 0);
        assertEq(token.getTotalBurnLimit(bridge), INITIAL_TOTAL_LIMIT);
    }

    function test_burn_zeroTimePassed_revert_burnOverLimit() public {
        vm.expectRevert(
            abi.encodeWithSelector(
                RateLimiting.RateLimiting__LimitExceeded.selector, INITIAL_CURRENT_LIMIT + 1, INITIAL_CURRENT_LIMIT
            )
        );
        authBurnToken(INITIAL_CURRENT_LIMIT + 1);
    }

    function test_burn_timePassed_replenishUnderTotalLimit_burnUnderLimit() public {
        skip(SMALL_PERIOD);
        uint256 amount = SMALL_PERIOD_CURRENT_LIMIT / 10;
        authBurnToken(amount);
        assertEq(token.balanceOf(bridge), INITIAL_BRIDGE_BALANCE - amount);
        assertEq(token.totalSupply(), INITIAL_TOTAL_SUPPLY - amount);
        assertEq(token.getCurrentBurnLimit(bridge), SMALL_PERIOD_CURRENT_LIMIT - amount);
        assertEq(token.getTotalBurnLimit(bridge), INITIAL_TOTAL_LIMIT);
    }

    function test_burn_timePassed_replenishUnderTotalLimit_burnExactlyLimit() public {
        skip(SMALL_PERIOD);
        authBurnToken(SMALL_PERIOD_CURRENT_LIMIT);
        assertEq(token.balanceOf(bridge), INITIAL_BRIDGE_BALANCE - SMALL_PERIOD_CURRENT_LIMIT);
        assertEq(token.totalSupply(), INITIAL_TOTAL_SUPPLY - SMALL_PERIOD_CURRENT_LIMIT);
        assertEq(token.getCurrentBurnLimit(bridge), 0);
        assertEq(token.getTotalBurnLimit(bridge), INITIAL_TOTAL_LIMIT);
    }

    function test_burn_timePassed_replenishUnderTotalLimit_revert_burnOverLimit() public {
        skip(SMALL_PERIOD);
        vm.expectRevert(
            abi.encodeWithSelector(
                RateLimiting.RateLimiting__LimitExceeded.selector,
                SMALL_PERIOD_CURRENT_LIMIT + 1,
                SMALL_PERIOD_CURRENT_LIMIT
            )
        );
        authBurnToken(SMALL_PERIOD_CURRENT_LIMIT + 1);
    }

    function test_burn_timePassed_replenishOverTotalLimit_burnUnderLimit() public {
        skip(LARGE_PERIOD);
        uint256 amount = INITIAL_TOTAL_LIMIT / 10;
        authBurnToken(amount);
        assertEq(token.balanceOf(bridge), INITIAL_BRIDGE_BALANCE - amount);
        assertEq(token.totalSupply(), INITIAL_TOTAL_SUPPLY - amount);
        assertEq(token.getCurrentBurnLimit(bridge), INITIAL_TOTAL_LIMIT - amount);
        assertEq(token.getTotalBurnLimit(bridge), INITIAL_TOTAL_LIMIT);
    }

    function test_burn_timePassed_replenishOverTotalLimit_burnExactlyLimit() public {
        skip(LARGE_PERIOD);
        authBurnToken(INITIAL_TOTAL_LIMIT);
        assertEq(token.balanceOf(bridge), INITIAL_BRIDGE_BALANCE - INITIAL_TOTAL_LIMIT);
        assertEq(token.totalSupply(), INITIAL_TOTAL_SUPPLY - INITIAL_TOTAL_LIMIT);
        assertEq(token.getCurrentBurnLimit(bridge), 0);
        assertEq(token.getTotalBurnLimit(bridge), INITIAL_TOTAL_LIMIT);
    }

    function test_burn_timePassed_replenishOverTotalLimit_revert_burnOverLimit() public {
        skip(LARGE_PERIOD);
        vm.expectRevert(
            abi.encodeWithSelector(
                RateLimiting.RateLimiting__LimitExceeded.selector, INITIAL_TOTAL_LIMIT + 1, INITIAL_TOTAL_LIMIT
            )
        );
        authBurnToken(INITIAL_TOTAL_LIMIT + 1);
    }

    // ═════════════════════════════════════════ TESTS: BURN BY PROCESSOR ══════════════════════════════════════════════

    function test_burn_byProcessor() public {
        vm.prank(processor);
        token.burn(100);
        assertEq(token.balanceOf(processor), INITIAL_PROCESSOR_BALANCE - 100);
        assertEq(token.totalSupply(), INITIAL_TOTAL_SUPPLY - 100);
        // Should not affect the bridge's burn limit
        assertEq(token.getCurrentBurnLimit(bridge), INITIAL_CURRENT_LIMIT);
        assertEq(token.getTotalBurnLimit(bridge), INITIAL_TOTAL_LIMIT);
        // Should not affect the processor's burn limit
        assertEq(token.getCurrentBurnLimit(processor), type(uint256).max);
        assertEq(token.getTotalBurnLimit(processor), type(uint256).max);
    }

    function test_burn_byProcessor_bigAmount() public {
        uint256 amount = INITIAL_TOTAL_LIMIT * 10;
        vm.prank(processor);
        token.burn(amount);
        assertEq(token.balanceOf(processor), INITIAL_PROCESSOR_BALANCE - amount);
        assertEq(token.totalSupply(), INITIAL_TOTAL_SUPPLY - amount);
        // Should not affect the bridge's burn limit
        assertEq(token.getCurrentBurnLimit(bridge), INITIAL_CURRENT_LIMIT);
        assertEq(token.getTotalBurnLimit(bridge), INITIAL_TOTAL_LIMIT);
        // Should not affect the processor's burn limit
        assertEq(token.getCurrentBurnLimit(processor), type(uint256).max);
        assertEq(token.getTotalBurnLimit(processor), type(uint256).max);
    }

    // ════════════════════════════════════════════ TESTS: BURN + PAUSE ════════════════════════════════════════════════

    function test_burn_revert_whenPaused() public {
        authPause();
        vm.expectRevert(Pausable.EnforcedPause.selector);
        authBurnToken(100);
    }

    function test_burn_works_whenPausedAndUnpaused() public {
        authPause();
        authUnpause();
        test_burn_zeroTimePassed_burnUnderLimit();
    }

    function test_burn_byProcessor_revert_whenPaused() public {
        authPause();
        vm.expectRevert(Pausable.EnforcedPause.selector);
        vm.prank(processor);
        token.burn(100);
    }

    function test_burn_byProcessor_works_whenPausedAndUnpaused() public {
        authPause();
        authUnpause();
        test_burn_byProcessor();
    }
}
