// SPDX-License-Identifier: MIT
pragma solidity 0.8.23;

import {RateLimit, RateLimiting, RateLimitHarness} from "../harnesses/RateLimitHarness.sol";

import {Test} from "forge-std/Test.sol";

// solhint-disable func-name-mixedcase
// solhint-disable ordering
contract RateLimitingLibraryTest is Test {
    RateLimitHarness public harness;

    function setUp() public {
        harness = new RateLimitHarness();
    }

    // Restrict the RateLimit struct to reasonable values
    function boundRateLimit(RateLimit memory rateLimit) internal pure {
        // 2024-01-01 .. 2024-02-01
        rateLimit.lastUpdatedAt = bound(rateLimit.lastUpdatedAt, 1_704_067_200, 1_706_745_600);
        rateLimit.totalLimit = bound(rateLimit.totalLimit, 1 ether, 1000 ether);
        rateLimit.lastRemaining = bound(rateLimit.lastRemaining, 0, rateLimit.totalLimit);
    }

    function useRateLimitFixture(RateLimit memory rateLimit) public {
        boundRateLimit(rateLimit);
        harness.setRateLimit(rateLimit);
        vm.warp(rateLimit.lastUpdatedAt);
    }

    function replenishUnderTotalLimitParams(
        RateLimit memory rateLimit,
        uint256 randomValue
    )
        public
        pure
        returns (uint256 timePassed, uint256 replenished)
    {
        uint256 spent = rateLimit.totalLimit - rateLimit.lastRemaining;
        // Minimum amount replenished (for 1 second passed)
        uint256 minReplenished = rateLimit.totalLimit / RateLimiting.REPLENISH_PERIOD;
        vm.assume(spent > minReplenished);
        replenished = bound(randomValue, minReplenished, spent);
        // Take care of the rounding
        timePassed = (replenished * RateLimiting.REPLENISH_PERIOD) / rateLimit.totalLimit;
        replenished = (timePassed * rateLimit.totalLimit) / RateLimiting.REPLENISH_PERIOD;
    }

    function replenishOverTotalLimitParams(
        RateLimit memory rateLimit,
        uint256 randomValue
    )
        public
        pure
        returns (uint256 timePassed)
    {
        uint256 spent = rateLimit.totalLimit - rateLimit.lastRemaining;
        // How much time takes to replenish the spent amount
        uint256 maxTimePassed = spent * RateLimiting.REPLENISH_PERIOD / rateLimit.totalLimit;
        vm.assume(maxTimePassed > 0);
        timePassed = bound(randomValue, maxTimePassed + 1, maxTimePassed + 30 days);
    }

    function spendWithinLimitParams(
        uint256 currentRemaining,
        uint256 randomValue
    )
        public
        pure
        returns (uint256 amount)
    {
        vm.assume(currentRemaining > 0);
        amount = bound(randomValue, 1, currentRemaining);
    }

    function spendExceedsLimitParams(
        uint256 currentRemaining,
        uint256 randomValue
    )
        public
        pure
        returns (uint256 amount)
    {
        amount = bound(randomValue, currentRemaining + 1, currentRemaining + 1 ether);
    }

    // ════════════════════════════════════════ TESTS: SETTING TOTAL LIMIT ═════════════════════════════════════════════

    function changeLimitAndCheck(uint256 newLimit, uint256 expectedRemaining) public {
        harness.setTotalLimit(newLimit);
        (uint256 lastUpdatedAt, uint256 lastRemaining, uint256 totalLimit) = harness.rateLimit();
        assertEq(lastUpdatedAt, block.timestamp);
        assertEq(lastRemaining, expectedRemaining);
        assertEq(totalLimit, newLimit);
    }

    function test_setTotalLimit_fromZero(uint256 newLimit) public {
        newLimit = bound(newLimit, 1 ether, 1000 ether);
        changeLimitAndCheck({newLimit: newLimit, expectedRemaining: newLimit});
    }

    function test_setTotalLimit_zeroTimePassed_limitIncreased(RateLimit memory rateLimit, uint256 newLimit) public {
        useRateLimitFixture(rateLimit);
        newLimit = bound(newLimit, rateLimit.totalLimit + 1, rateLimit.totalLimit + 1 ether);
        uint256 expectedRemaining = rateLimit.lastRemaining + newLimit - rateLimit.totalLimit;
        changeLimitAndCheck(newLimit, expectedRemaining);
    }

    function test_setTotalLimit_zeroTimePassed_limitDecreased_overPreviouslySpent(
        RateLimit memory rateLimit,
        uint256 newLimit
    )
        public
    {
        useRateLimitFixture(rateLimit);
        uint256 previouslySpent = rateLimit.totalLimit - rateLimit.lastRemaining;
        vm.assume(previouslySpent < rateLimit.totalLimit - 1);
        newLimit = bound(newLimit, previouslySpent, rateLimit.totalLimit - 1);
        uint256 expectedRemaining = rateLimit.lastRemaining + newLimit - rateLimit.totalLimit;
        changeLimitAndCheck(newLimit, expectedRemaining);
    }

    function test_setTotalLimit_zeroTimePassed_limitDecreased_underPreviouslySpent(
        RateLimit memory rateLimit,
        uint256 newLimit
    )
        public
    {
        useRateLimitFixture(rateLimit);
        uint256 previouslySpent = rateLimit.totalLimit - rateLimit.lastRemaining;
        vm.assume(previouslySpent > 0.1 ether);
        newLimit = bound(newLimit, 0.1 ether, previouslySpent - 1);
        uint256 expectedRemaining = 0;
        changeLimitAndCheck(newLimit, expectedRemaining);
    }

    function test_setTotalLimit_zeroTimePassed_limitDecreased_toZero(RateLimit memory rateLimit) public {
        useRateLimitFixture(rateLimit);
        uint256 newLimit = 0;
        uint256 expectedRemaining = 0;
        changeLimitAndCheck(newLimit, expectedRemaining);
    }

    function test_setTotalLimit_timePassed_replenishUnderTotalLimit_limitIncreased(
        RateLimit memory rateLimit,
        uint256 randomValue,
        uint256 newLimit
    )
        public
    {
        useRateLimitFixture(rateLimit);
        (uint256 timePassed, uint256 replenished) = replenishUnderTotalLimitParams(rateLimit, randomValue);
        skip(timePassed);
        newLimit = bound(newLimit, rateLimit.totalLimit + 1, rateLimit.totalLimit + 1 ether);
        uint256 expectedRemaining = rateLimit.lastRemaining + replenished + newLimit - rateLimit.totalLimit;
        changeLimitAndCheck(newLimit, expectedRemaining);
    }

    function test_setTotalLimit_timePassed_replenishUnderTotalLimit_limitDecreased_overPreviouslySpent(
        RateLimit memory rateLimit,
        uint256 randomValue,
        uint256 newLimit
    )
        public
    {
        useRateLimitFixture(rateLimit);
        (uint256 timePassed, uint256 replenished) = replenishUnderTotalLimitParams(rateLimit, randomValue);
        skip(timePassed);
        uint256 previouslySpent = rateLimit.totalLimit - (rateLimit.lastRemaining + replenished);
        vm.assume(previouslySpent < rateLimit.totalLimit - 1);
        newLimit = bound(newLimit, previouslySpent, rateLimit.totalLimit - 1);
        uint256 expectedRemaining = rateLimit.lastRemaining + replenished + newLimit - rateLimit.totalLimit;
        changeLimitAndCheck(newLimit, expectedRemaining);
    }

    function test_setTotalLimit_timePassed_replenishUnderTotalLimit_limitDecreased_underPreviouslySpent(
        RateLimit memory rateLimit,
        uint256 randomValue,
        uint256 newLimit
    )
        public
    {
        useRateLimitFixture(rateLimit);
        (uint256 timePassed, uint256 replenished) = replenishUnderTotalLimitParams(rateLimit, randomValue);
        skip(timePassed);
        uint256 previouslySpent = rateLimit.totalLimit - (rateLimit.lastRemaining + replenished);
        vm.assume(previouslySpent > 0.1 ether);
        newLimit = bound(newLimit, 0.1 ether, previouslySpent - 1);
        uint256 expectedRemaining = 0;
        changeLimitAndCheck(newLimit, expectedRemaining);
    }

    function test_setTotalLimit_timePassed_replenishUnderTotalLimit_limitDecreased_toZero(
        RateLimit memory rateLimit,
        uint256 randomValue
    )
        public
    {
        useRateLimitFixture(rateLimit);
        (uint256 timePassed,) = replenishUnderTotalLimitParams(rateLimit, randomValue);
        skip(timePassed);
        uint256 newLimit = 0;
        uint256 expectedRemaining = 0;
        changeLimitAndCheck(newLimit, expectedRemaining);
    }

    function test_setTotalLimit_timePassed_replenishOverTotalLimit_limitIncreased(
        RateLimit memory rateLimit,
        uint256 randomValueTime,
        uint256 newLimit
    )
        public
    {
        useRateLimitFixture(rateLimit);
        uint256 timePassed = replenishOverTotalLimitParams(rateLimit, randomValueTime);
        skip(timePassed);
        newLimit = bound(newLimit, rateLimit.totalLimit + 1, rateLimit.totalLimit + 1 ether);
        uint256 expectedRemaining = newLimit;
        changeLimitAndCheck(newLimit, expectedRemaining);
    }

    function test_setTotalLimit_timePassed_replenishOverTotalLimit_limitDecreased(
        RateLimit memory rateLimit,
        uint256 randomValueTime,
        uint256 newLimit
    )
        public
    {
        useRateLimitFixture(rateLimit);
        uint256 timePassed = replenishOverTotalLimitParams(rateLimit, randomValueTime);
        skip(timePassed);
        newLimit = bound(newLimit, 0.1 ether, rateLimit.totalLimit - 1);
        uint256 expectedRemaining = newLimit;
        changeLimitAndCheck(newLimit, expectedRemaining);
    }

    function test_setTotalLimit_timePassed_replenishOverTotalLimit_limitDecreased_toZero(
        RateLimit memory rateLimit,
        uint256 randomValueTime
    )
        public
    {
        useRateLimitFixture(rateLimit);
        uint256 timePassed = replenishOverTotalLimitParams(rateLimit, randomValueTime);
        skip(timePassed);
        uint256 newLimit = 0;
        uint256 expectedRemaining = 0;
        changeLimitAndCheck(newLimit, expectedRemaining);
    }

    // ═══════════════════════════════════════════ TESTS: SPENDING LIMIT ═══════════════════════════════════════════════

    function test_spendLimit_zeroTimePassed_spendWithinLimit(RateLimit memory rateLimit, uint256 randomValue) public {
        useRateLimitFixture(rateLimit);
        uint256 amount = spendWithinLimitParams(rateLimit.lastRemaining, randomValue);
        harness.spendLimit(amount);
        (uint256 lastUpdatedAt, uint256 lastRemaining, uint256 totalLimit) = harness.rateLimit();
        assertEq(lastUpdatedAt, block.timestamp);
        assertEq(lastRemaining, rateLimit.lastRemaining - amount);
        assertEq(totalLimit, rateLimit.totalLimit);
    }

    function test_spendLimit_zeroTimePassed_spendExceedsLimit(RateLimit memory rateLimit, uint256 randomValue) public {
        useRateLimitFixture(rateLimit);
        uint256 amount = spendExceedsLimitParams(rateLimit.lastRemaining, randomValue);
        vm.expectRevert(
            abi.encodeWithSelector(RateLimiting.RateLimiting__LimitExceeded.selector, amount, rateLimit.lastRemaining)
        );
        harness.spendLimit(amount);
    }

    function test_spendLimit_timePassed_replenishUnderTotalLimit_spendWithinLimit(
        RateLimit memory rateLimit,
        uint256 randomValueTime,
        uint256 randomValueAmount
    )
        public
    {
        useRateLimitFixture(rateLimit);
        (uint256 timePassed, uint256 replenished) = replenishUnderTotalLimitParams(rateLimit, randomValueTime);
        skip(timePassed);
        uint256 currentRemaining = rateLimit.lastRemaining + replenished;
        uint256 amount = spendWithinLimitParams(currentRemaining, randomValueAmount);
        harness.spendLimit(amount);
        (uint256 lastUpdatedAt, uint256 lastRemaining, uint256 totalLimit) = harness.rateLimit();
        assertEq(lastUpdatedAt, block.timestamp);
        assertEq(lastRemaining, currentRemaining - amount);
        assertEq(totalLimit, rateLimit.totalLimit);
    }

    function test_spendLimit_timePassed_replenishUnderTotalLimit_spendExceedsLimit(
        RateLimit memory rateLimit,
        uint256 randomValueTime,
        uint256 randomValueAmount
    )
        public
    {
        useRateLimitFixture(rateLimit);
        (uint256 timePassed, uint256 replenished) = replenishUnderTotalLimitParams(rateLimit, randomValueTime);
        skip(timePassed);
        uint256 currentRemaining = rateLimit.lastRemaining + replenished;
        uint256 amount = spendExceedsLimitParams(currentRemaining, randomValueAmount);
        vm.expectRevert(
            abi.encodeWithSelector(RateLimiting.RateLimiting__LimitExceeded.selector, amount, currentRemaining)
        );
        harness.spendLimit(amount);
    }

    function test_spendLimit_timePassed_replenishOverTotalLimit_spendWithinLimit(
        RateLimit memory rateLimit,
        uint256 randomValueTime,
        uint256 randomValueAmount
    )
        public
    {
        useRateLimitFixture(rateLimit);
        uint256 timePassed = replenishOverTotalLimitParams(rateLimit, randomValueTime);
        skip(timePassed);
        uint256 amount = spendWithinLimitParams(rateLimit.totalLimit, randomValueAmount);
        harness.spendLimit(amount);
        (uint256 lastUpdatedAt, uint256 lastRemaining, uint256 totalLimit) = harness.rateLimit();
        assertEq(lastUpdatedAt, block.timestamp);
        assertEq(lastRemaining, rateLimit.totalLimit - amount);
        assertEq(totalLimit, rateLimit.totalLimit);
    }

    function test_spendLimit_timePassed_replenishOverTotalLimit_spendExceedsLimit(
        RateLimit memory rateLimit,
        uint256 randomValueTime,
        uint256 randomValueAmount
    )
        public
    {
        useRateLimitFixture(rateLimit);
        uint256 timePassed = replenishOverTotalLimitParams(rateLimit, randomValueTime);
        skip(timePassed);
        uint256 amount = spendExceedsLimitParams(rateLimit.totalLimit, randomValueAmount);
        vm.expectRevert(
            abi.encodeWithSelector(RateLimiting.RateLimiting__LimitExceeded.selector, amount, rateLimit.totalLimit)
        );
        harness.spendLimit(amount);
    }

    // ═══════════════════════════════════════════════ TESTS: VIEWS ════════════════════════════════════════════════════

    function test_REPLENISH_PERIOD_oneDay() public {
        assertEq(RateLimiting.REPLENISH_PERIOD, 1 days);
    }

    function test_getCurrentLimit_zeroTimePassed(RateLimit memory rateLimit) public {
        useRateLimitFixture(rateLimit);
        assertEq(harness.getCurrentLimit(), rateLimit.lastRemaining);
    }

    function test_getCurrentLimit_timePassed_replenishUnderTotalLimit(
        RateLimit memory rateLimit,
        uint256 randomValue
    )
        public
    {
        useRateLimitFixture(rateLimit);
        (uint256 timePassed, uint256 replenished) = replenishUnderTotalLimitParams(rateLimit, randomValue);
        skip(timePassed);
        uint256 expected = rateLimit.lastRemaining + replenished;
        assertEq(harness.getCurrentLimit(), expected);
    }

    function test_getCurrentLimit_timePassed_replenishOverTotalLimit(
        RateLimit memory rateLimit,
        uint256 randomValue
    )
        public
    {
        useRateLimitFixture(rateLimit);
        uint256 timePassed = replenishOverTotalLimitParams(rateLimit, randomValue);
        skip(timePassed);
        assertEq(harness.getCurrentLimit(), rateLimit.totalLimit);
    }

    function test_getTotalLimit_zeroTimePassed(RateLimit memory rateLimit) public {
        useRateLimitFixture(rateLimit);
        assertEq(harness.getTotalLimit(), rateLimit.totalLimit);
    }

    function test_getTotalLimit_withTimePassed(RateLimit memory rateLimit) public {
        useRateLimitFixture(rateLimit);
        skip(12 hours);
        assertEq(harness.getTotalLimit(), rateLimit.totalLimit);
    }
}
