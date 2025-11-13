/**
 * Slippage warning thresholds for USD-based slippage calculations
 *
 * These thresholds determine when to show amber vs red warning colors to users.
 * A warning (amber) is shown if EITHER threshold is exceeded.
 * A critical warning (red) is shown if BOTH thresholds are exceeded.
 */

/**
 * Percentage-based slippage threshold
 * Triggers amber warning if slippage is worse than this percentage
 *
 * @example -2.5 means losses greater than 2.5% trigger a warning
 */
export const SLIPPAGE_WARNING_THRESHOLD = -2.5

/**
 * USD-based slippage threshold
 * Triggers amber warning if USD loss is greater than this amount
 *
 * @example -1 means losses greater than $1 trigger a warning
 */
export const USD_SLIPPAGE_WARNING_THRESHOLD = -1
