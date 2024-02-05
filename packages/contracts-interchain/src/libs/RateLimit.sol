// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

/// @notice Struct defining the self-replenishing rate limit for an operation with a token.
/// - Limits are defined in the token's decimals.
/// - Every operation spends the limit by the amount of tokens processed.
/// - Limits are replenished at a constant rate.
/// - Limit can never surpass the total limit.
/// - It takes 24 hours to fully replenish the limit (assuming no operations are performed during that time).
/// @param lastUpdatedAt    The timestamp of the last operation
/// @param lastRemaining    The remaining limit after the last operation
/// @param totalLimit       The maximum value the remaining limit can reach
struct RateLimit {
    uint256 lastUpdatedAt;
    uint256 lastRemaining;
    uint256 totalLimit;
}

using RateLimiting for RateLimit global;

library RateLimiting {
    /// @notice Period of time it takes to fully replenish the exhausted limit.
    uint256 internal constant REPLENISH_PERIOD = 1 days;

    /// @notice Error that is thrown when the rate limit is exceeded.
    error RateLimiting__LimitExceeded(uint256 amount, uint256 limit);

    /// @notice Spend the limit associated with the operation by the given amount.
    /// @dev Will revert if the amount exceeds the current limit.
    /// Note: this function will update the RateLimit struct in storage.
    /// @param self     The RateLimit struct
    /// @param amount   The amount to spend
    function spendLimit(RateLimit storage self, uint256 amount) internal {
        // First, check if the amount exceeds the limit
        uint256 limit = getCurrentLimit(self);
        if (amount > limit) {
            revert RateLimiting__LimitExceeded(amount, limit);
        }
        // Now that we know the amount is within the limit, update the RateLimit struct
        unchecked {
            self.lastRemaining = limit - amount;
        }
        self.lastUpdatedAt = block.timestamp;
    }

    /// @notice Gets the current limit for the operation (at the time of the call).
    /// @dev This is a view function and does not modify the RateLimit struct.
    /// @param self     The RateLimit struct
    /// @return limit   The remaining limit at the time of the call
    function getCurrentLimit(RateLimit storage self) internal view returns (uint256 limit) {
        uint256 timePassed = block.timestamp - self.lastUpdatedAt;
        // Check if the full replenish period has passed since the last update
        if (timePassed >= REPLENISH_PERIOD) {
            return self.totalLimit;
        }
        // Otherwise, figure out how much the limit has replenished since the last update
        uint256 replenished = (timePassed * self.totalLimit) / REPLENISH_PERIOD;
        limit = self.lastRemaining + replenished;
        // Limit cannot surpass the total limit
        if (limit > self.totalLimit) {
            limit = self.totalLimit;
        }
    }

    /// @notice Gets the total limit for the operation.
    /// @dev This is a view function and does not modify the RateLimit struct.
    /// @param self     The RateLimit struct
    /// @return limit   The total limit
    function getTotalLimit(RateLimit storage self) internal view returns (uint256 limit) {
        return self.totalLimit;
    }
}
