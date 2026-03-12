/**
 * Slippage warning thresholds for USD-based slippage calculations
 *
 * These thresholds determine when to show amber vs red warning colors to users.
 * The color logic works as follows:
 * - Green: Positive slippage and difference (both higher than neutral thresholds)
 * - White: Either slippage or difference is below neutral thresholds (positive or negative)
 * - Amber: Slippage or difference is below warning thresholds (negative)
 * - Red: Negative slippage and difference (both higher than warning thresholds)
 */

export enum PercentageThreshold {
  NEUTRAL = 0.1,
  WARNING = 1,
}

export enum AbsoluteThreshold {
  NEUTRAL = 1,
  WARNING = 10,
}
