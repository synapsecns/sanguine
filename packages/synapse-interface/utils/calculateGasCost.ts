import { formatGwei } from 'viem'

/**
 * Calculates the estimated gas cost for a transaction.
 *
 * @param {string} gasPrice - The current network gas price in Gwei
 * @param {number} gasLimit - Selected gasLimit to execute Transaction
 */

export const calculateGasCost = (
  gasPrice: string | undefined,
  gasLimit: string
): {
  rawGasCost: string
  parsedGasCost: string
} => {
  if (!gasPrice || !gasLimit) {
    return {
      rawGasCost: null,
      parsedGasCost: null,
    }
  }

  const upperLimitBuffer = 1.5

  const estimatedGasCostInGwei =
    parseFloat(gasLimit) * parseFloat(gasPrice) * upperLimitBuffer

  const oneGwei = parseFloat(formatGwei(1n))

  const formattedEstimatedGasCost = estimatedGasCostInGwei
    ? estimatedGasCostInGwei * oneGwei
    : null

  return {
    rawGasCost: estimatedGasCostInGwei?.toString(),
    parsedGasCost: formattedEstimatedGasCost?.toString(),
  }
}
