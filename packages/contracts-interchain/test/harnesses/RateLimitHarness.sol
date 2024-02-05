// SPDX-License-Identifier: MIT
pragma solidity 0.8.23;

import {RateLimit, RateLimiting} from "../../src/libs/RateLimit.sol";

contract RateLimitHarness {
    RateLimit public rateLimit;

    function setRateLimit(RateLimit memory rateLimit_) public {
        rateLimit = rateLimit_;
    }

    function spendLimit(uint256 amount) public {
        RateLimiting.spendLimit(rateLimit, amount);
    }

    function getCurrentLimit() public view returns (uint256) {
        return RateLimiting.getCurrentLimit(rateLimit);
    }

    function getTotalLimit() public view returns (uint256) {
        return RateLimiting.getTotalLimit(rateLimit);
    }
}
