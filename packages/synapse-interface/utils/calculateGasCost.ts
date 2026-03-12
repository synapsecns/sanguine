import { formatGwei } from 'viem'

/**
 * Calculates the estimated gas cost for a transaction.
 *
 * @param {string} gasPrice - The current network gas price in Gwei
 * @param {number} gasLimit - Selected gasLimit to execute Transaction
 */

export const calculateGasCost = (
  gasPrice: string | undefined,
  gasLimit: string,
  chainId: number | null
): {
  rawGasCost: string | null | undefined
  parsedGasCost: string | null | undefined
} => {
  if (!gasPrice || !gasLimit || !chainId) {
    return {
      rawGasCost: null,
      parsedGasCost: null,
    }
  }

  const upperLimitBuffer = 3
  const gasLimitFloat = parseFloat(gasLimit) ? parseFloat(gasLimit) : 1
  const gasPriceFloat = parseFloat(gasPrice) ? parseFloat(gasPrice) : 1

  const estimatedGasCostInGwei =
    gasLimitFloat * gasPriceFloat * upperLimitBuffer

  const oneGwei = parseFloat(formatGwei(1n))

  const formattedEstimatedGasCost = estimatedGasCostInGwei
    ? estimatedGasCostInGwei * oneGwei
    : null

  return {
    rawGasCost: estimatedGasCostInGwei?.toString(),
    parsedGasCost: formattedEstimatedGasCost?.toString(),
  }
}
