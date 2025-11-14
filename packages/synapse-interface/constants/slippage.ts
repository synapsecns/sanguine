/**
 * Slippage warning thresholds for USD-based slippage calculations
 *
 * These thresholds determine when to show amber vs red warning colors to users.
 * The color logic works as follows:
 * - Green: Positive slippage (gain)
 * - Amber: Loss <= 2.5% OR loss <= $1 (either threshold provides warning)
 * - Red: Loss > 2.5% AND loss > $1 (both thresholds must be exceeded)
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
