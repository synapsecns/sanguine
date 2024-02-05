// SPDX-License-Identifier: MIT
pragma solidity 0.8.23;

import {InterchainERC20Test} from "./InterchainERC20.t.sol";
import {RateLimiting} from "../../src/libs/RateLimit.sol";

import {IERC20Errors} from "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import {Pausable} from "@openzeppelin/contracts/utils/Pausable.sol";

// solhint-disable func-name-mixedcase
// solhint-disable ordering
contract InterchainERC20BurnTest is InterchainERC20Test {
    address public bridge;
    address public user;

    uint256 public constant INITIAL_TOTAL_LIMIT = 1000;
    uint256 public constant INITIAL_BURNED = 400;
    uint256 public constant INITIAL_CURRENT_LIMIT = INITIAL_TOTAL_LIMIT - INITIAL_BURNED;

    uint256 public constant SMALL_PERIOD = 1 days / 10;
    uint256 public constant SMALL_PERIOD_CURRENT_LIMIT = INITIAL_CURRENT_LIMIT + 100;
    uint256 public constant LARGE_PERIOD = 1 days * 2;

    uint256 public constant INITIAL_USER_BALANCE = 1000;
    uint256 public constant INITIAL_APPROVE_AMOUNT = 2 * INITIAL_USER_BALANCE;
    uint256 public constant INITIAL_PROCESSOR_BALANCE = 10_000;
    uint256 public constant INITIAL_TOTAL_SUPPLY = INITIAL_USER_BALANCE + INITIAL_PROCESSOR_BALANCE;

    function setUp() public override {
        super.setUp();
        bridge = makeAddr("Bridge");
        user = makeAddr("User");
        // Set the total burn limit for the bridge and spend some of it
        authSetTotalBurnLimit(bridge, INITIAL_TOTAL_LIMIT);
        processorMintToken(user, INITIAL_USER_BALANCE);
        processorMintToken(processor, INITIAL_PROCESSOR_BALANCE);
        processorMintToken(bridge, INITIAL_BURNED);
        vm.prank(bridge);
        token.burn(INITIAL_BURNED);
        vm.prank(user);
        token.approve(bridge, INITIAL_APPROVE_AMOUNT);
    }

    function processorMintToken(address account, uint256 amount) public {
        vm.prank(processor);
        token.mint(account, amount);
    }

    function authBurnFrom(address account, uint256 amount) public {
        vm.prank(bridge);
        token.burnFrom(account, amount);
    }

    function test_setUp() public {
        assertEq(token.getCurrentBurnLimit(bridge), INITIAL_CURRENT_LIMIT);
        assertEq(token.getTotalBurnLimit(bridge), INITIAL_TOTAL_LIMIT);
        assertEq(token.allowance(user, bridge), INITIAL_APPROVE_AMOUNT);
        assertEq(token.allowance(user, processor), 0);
        assertEq(token.balanceOf(user), INITIAL_USER_BALANCE);
        assertEq(token.balanceOf(processor), INITIAL_PROCESSOR_BALANCE);
        assertEq(token.totalSupply(), INITIAL_TOTAL_SUPPLY);
    }

    function test_burnFrom_revert_unauthorized() public {
        vm.startPrank(user);
        token.approve(user, 100);
        vm.expectRevert(abi.encodeWithSelector(RateLimiting.RateLimiting__LimitExceeded.selector, 100, 0));
        token.burnFrom(user, 100);
        vm.stopPrank();
    }

    function test_burnFrom_zeroTimePassed_burnUnderLimit() public {
        uint256 amount = INITIAL_CURRENT_LIMIT / 10;
        authBurnFrom(user, amount);
        assertEq(token.allowance(user, bridge), INITIAL_APPROVE_AMOUNT - amount);
        assertEq(token.balanceOf(user), INITIAL_USER_BALANCE - amount);
        assertEq(token.totalSupply(), INITIAL_TOTAL_SUPPLY - amount);
        assertEq(token.getCurrentBurnLimit(bridge), INITIAL_CURRENT_LIMIT - amount);
        assertEq(token.getTotalBurnLimit(bridge), INITIAL_TOTAL_LIMIT);
    }

    function test_burnFrom_zeroTimePassed_burnExactlyLimit() public {
        authBurnFrom(user, INITIAL_CURRENT_LIMIT);
        assertEq(token.allowance(user, bridge), INITIAL_APPROVE_AMOUNT - INITIAL_CURRENT_LIMIT);
        assertEq(token.balanceOf(user), INITIAL_USER_BALANCE - INITIAL_CURRENT_LIMIT);
        assertEq(token.totalSupply(), INITIAL_TOTAL_SUPPLY - INITIAL_CURRENT_LIMIT);
        assertEq(token.getCurrentBurnLimit(bridge), 0);
        assertEq(token.getTotalBurnLimit(bridge), INITIAL_TOTAL_LIMIT);
    }

    function test_burnFrom_zeroTimePassed_revert_burnOverLimit() public {
        vm.expectRevert(
            abi.encodeWithSelector(
                RateLimiting.RateLimiting__LimitExceeded.selector, INITIAL_CURRENT_LIMIT + 1, INITIAL_CURRENT_LIMIT
            )
        );
        authBurnFrom(user, INITIAL_CURRENT_LIMIT + 1);
    }

    function test_burnFrom_timePassed_replenishUnderTotalLimit_burnUnderLimit() public {
        skip(SMALL_PERIOD);
        uint256 amount = SMALL_PERIOD_CURRENT_LIMIT / 10;
        authBurnFrom(user, amount);
        assertEq(token.allowance(user, bridge), INITIAL_APPROVE_AMOUNT - amount);
        assertEq(token.balanceOf(user), INITIAL_USER_BALANCE - amount);
        assertEq(token.totalSupply(), INITIAL_TOTAL_SUPPLY - amount);
        assertEq(token.getCurrentBurnLimit(bridge), SMALL_PERIOD_CURRENT_LIMIT - amount);
        assertEq(token.getTotalBurnLimit(bridge), INITIAL_TOTAL_LIMIT);
    }

    function test_burnFrom_timePassed_replenishUnderTotalLimit_burnExactlyLimit() public {
        skip(SMALL_PERIOD);
        authBurnFrom(user, SMALL_PERIOD_CURRENT_LIMIT);
        assertEq(token.allowance(user, bridge), INITIAL_APPROVE_AMOUNT - SMALL_PERIOD_CURRENT_LIMIT);
        assertEq(token.balanceOf(user), INITIAL_USER_BALANCE - SMALL_PERIOD_CURRENT_LIMIT);
        assertEq(token.totalSupply(), INITIAL_TOTAL_SUPPLY - SMALL_PERIOD_CURRENT_LIMIT);
        assertEq(token.getCurrentBurnLimit(bridge), 0);
        assertEq(token.getTotalBurnLimit(bridge), INITIAL_TOTAL_LIMIT);
    }

    function test_burnFrom_timePassed_replenishUnderTotalLimit_revert_burnOverLimit() public {
        skip(SMALL_PERIOD);
        vm.expectRevert(
            abi.encodeWithSelector(
                RateLimiting.RateLimiting__LimitExceeded.selector,
                SMALL_PERIOD_CURRENT_LIMIT + 1,
                SMALL_PERIOD_CURRENT_LIMIT
            )
        );
        authBurnFrom(user, SMALL_PERIOD_CURRENT_LIMIT + 1);
    }

    function test_burnFrom_timePassed_replenishOverTotalLimit_burnUnderLimit() public {
        skip(LARGE_PERIOD);
        uint256 amount = INITIAL_TOTAL_LIMIT / 10;
        authBurnFrom(user, amount);
        assertEq(token.allowance(user, bridge), INITIAL_APPROVE_AMOUNT - amount);
        assertEq(token.balanceOf(user), INITIAL_USER_BALANCE - amount);
        assertEq(token.totalSupply(), INITIAL_TOTAL_SUPPLY - amount);
        assertEq(token.getCurrentBurnLimit(bridge), INITIAL_TOTAL_LIMIT - amount);
        assertEq(token.getTotalBurnLimit(bridge), INITIAL_TOTAL_LIMIT);
    }

    function test_burnFrom_timePassed_replenishOverTotalLimit_burnExactlyLimit() public {
        skip(LARGE_PERIOD);
        authBurnFrom(user, INITIAL_TOTAL_LIMIT);
        assertEq(token.allowance(user, bridge), INITIAL_APPROVE_AMOUNT - INITIAL_TOTAL_LIMIT);
        assertEq(token.balanceOf(user), INITIAL_USER_BALANCE - INITIAL_TOTAL_LIMIT);
        assertEq(token.totalSupply(), INITIAL_TOTAL_SUPPLY - INITIAL_TOTAL_LIMIT);
        assertEq(token.getCurrentBurnLimit(bridge), 0);
        assertEq(token.getTotalBurnLimit(bridge), INITIAL_TOTAL_LIMIT);
    }

    function test_burnFrom_timePassed_replenishOverTotalLimit_revert_burnOverLimit() public {
        skip(LARGE_PERIOD);
        vm.expectRevert(
            abi.encodeWithSelector(
                RateLimiting.RateLimiting__LimitExceeded.selector, INITIAL_TOTAL_LIMIT + 1, INITIAL_TOTAL_LIMIT
            )
        );
        authBurnFrom(user, INITIAL_TOTAL_LIMIT + 1);
    }

    // ═════════════════════════════════════════ TESTS: BURN BY PROCESSOR ══════════════════════════════════════════════

    function test_burnFrom_byProcessor() public {
        vm.prank(user);
        token.approve(processor, 100);
        vm.prank(processor);
        token.burnFrom(user, 100);
        assertEq(token.balanceOf(user), INITIAL_USER_BALANCE - 100);
        assertEq(token.totalSupply(), INITIAL_TOTAL_SUPPLY - 100);
        // Should not affect the bridge's burn limit
        assertEq(token.getCurrentBurnLimit(bridge), INITIAL_CURRENT_LIMIT);
        assertEq(token.getTotalBurnLimit(bridge), INITIAL_TOTAL_LIMIT);
        // Should not affect the processor's burn limit
        assertEq(token.getCurrentBurnLimit(processor), type(uint256).max);
        assertEq(token.getTotalBurnLimit(processor), type(uint256).max);
    }

    function test_burnFrom_byProcessor_bigAmount() public {
        uint256 amount = INITIAL_PROCESSOR_BALANCE;
        vm.prank(processor);
        token.transfer(user, amount);
        vm.prank(user);
        token.approve(processor, amount);
        vm.prank(processor);
        token.burnFrom(user, amount);
        assertEq(token.balanceOf(user), INITIAL_USER_BALANCE);
        assertEq(token.totalSupply(), INITIAL_TOTAL_SUPPLY - INITIAL_PROCESSOR_BALANCE);
        // Should not affect the bridge's burn limit
        assertEq(token.getCurrentBurnLimit(bridge), INITIAL_CURRENT_LIMIT);
        assertEq(token.getTotalBurnLimit(bridge), INITIAL_TOTAL_LIMIT);
        // Should not affect the processor's burn limit
        assertEq(token.getCurrentBurnLimit(processor), type(uint256).max);
        assertEq(token.getTotalBurnLimit(processor), type(uint256).max);
    }

    // ════════════════════════════════════════════ TESTS: BURN + PAUSE ════════════════════════════════════════════════

    function test_burnFrom_revert_whenPaused() public {
        authPause();
        vm.expectRevert(Pausable.EnforcedPause.selector);
        authBurnFrom(user, 100);
    }

    function test_burnFrom_works_whenPausedAndUnpaused() public {
        authPause();
        authUnpause();
        test_burnFrom_zeroTimePassed_burnUnderLimit();
    }

    function test_burnFrom_byProcessor_revert_whenPaused() public {
        vm.prank(user);
        token.approve(processor, 100);
        authPause();
        vm.expectRevert(Pausable.EnforcedPause.selector);
        vm.prank(processor);
        token.burnFrom(user, 100);
    }

    function test_burnFrom_byProcessor_works_whenPausedAndUnpaused() public {
        authPause();
        authUnpause();
        test_burnFrom_byProcessor();
    }

    // ═══════════════════════════════════════════ TESTS: BURN ALLOWANCE ═══════════════════════════════════════════════

    function test_burnFrom_revert_allowanceExceeded() public {
        vm.prank(user);
        token.approve(bridge, 100);
        vm.expectRevert(abi.encodeWithSelector(IERC20Errors.ERC20InsufficientAllowance.selector, bridge, 100, 101));
        authBurnFrom(user, 101);
    }
}
