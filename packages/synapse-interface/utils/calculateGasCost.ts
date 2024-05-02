import { formatGwei } from 'viem'

import { OPTIMISM } from '@/constants/chains/master'

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
  rawGasCost: string
  parsedGasCost: string
} => {
  if (!gasPrice || !gasLimit) {
    return {
      rawGasCost: null,
      parsedGasCost: null,
    }
  }

  let upperLimitBuffer = 1.75

  if (chainId === OPTIMISM.id) {
    upperLimitBuffer = 3
  }

  const gasLimitFloat = parseFloat(gasLimit) ? parseFloat(gasLimit) : 1
  const gasPriceFloat = parseFloat(gasPrice) ? parseFloat(gasPrice) : 1

  const estimatedGasCostInGwei =
    gasLimitFloat * gasPriceFloat * upperLimitBuffer

  const oneGwei = parseFloat(formatGwei(1n))

  const formattedEstimatedGasCost = estimatedGasCostInGwei
    ? estimatedGasCostInGwei * oneGwei
    : null

  // console.log('estimatedGasCostInGwei: ', estimatedGasCostInGwei)
  // console.log('gasLimit: ', gasLimit)
  // console.log('gasPrice: ', gasPrice)
  // console.log('gasLimitFloat: ', gasLimitFloat)
  // console.log('gasPriceFloat: ', gasPriceFloat)
  // console.log('formattedEstimatedGasCost: ', formattedEstimatedGasCost)

  return {
    rawGasCost: estimatedGasCostInGwei?.toString(),
    parsedGasCost: formattedEstimatedGasCost?.toString(),
  }
}
