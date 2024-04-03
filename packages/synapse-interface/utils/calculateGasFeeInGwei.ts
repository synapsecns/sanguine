import { formatGwei } from 'viem'

/**
 * Calculates the estimated gas fee in Gwei.
 * TODO: Hardcoding gas limit to 200k for now, update dynamically
 *
 * @param {string} gasPrice - The current gas price in Gwei as a string.
 * @param {number} gasLimit - Function to format a value as Gwei.
 * @returns {number|null} The formatted estimated gas cost, or null if the calculation is not possible.
 */

export const calculateGasFeeInGwei = (gasPrice?: string, gasLimit = 200000) => {
  if (!gasPrice) return null

  const estimatedGasCostInGwei = gasLimit * parseFloat(gasPrice)

  const oneGwei = parseFloat(formatGwei(1n))

  const formattedEstimatedGasCost = estimatedGasCostInGwei
    ? estimatedGasCostInGwei * oneGwei
    : null

  return formattedEstimatedGasCost
}
