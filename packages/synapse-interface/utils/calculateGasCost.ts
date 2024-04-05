import { formatGwei } from 'viem'

/**
 * Calculates the estimated gas cost for a transaction.
 *
 * TODO: Hardcoding gas limit to 200k for now, update dynamically
 *
 * @param {string} gasPrice - The current network gas price in Gwei
 * @param {number} gasLimit - Selected gasLimit to execute Transaction
 */

export const calculateGasCost = (
  gasPrice?: string,
  gasLimit = 200000
): {
  rawGasCost: number
  parsedGasCost: number
} => {
  if (!gasPrice) {
    return {
      rawGasCost: null,
      parsedGasCost: null,
    }
  }

  const estimatedGasCostInGwei = gasLimit * parseFloat(gasPrice)

  const oneGwei = parseFloat(formatGwei(1n))

  const formattedEstimatedGasCost = estimatedGasCostInGwei
    ? estimatedGasCostInGwei * oneGwei
    : null

  return {
    rawGasCost: estimatedGasCostInGwei,
    parsedGasCost: formattedEstimatedGasCost,
  }
}
