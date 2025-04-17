import { BigNumber } from '@ethersproject/bignumber'

import { PopulatedTx, Query } from './misc'

/**
 * Quote for a swap transaction for SynapseRouter (V1).
 * Returned by SDK to the consumer.
 *
 * @property {string} routerAddress - Address of the router that performs the operation
 * @property {BigNumber} maxAmountOut - Maximum amount of tokens that will be received (in token decimals)
 * @property {Query} query - Query details for the swap
 */
export type SwapQuote = {
  routerAddress: string
  maxAmountOut: BigNumber
  query: Query
}

/**
 * Parameters required to create a swap transaction (v2).
 *
 * @property {number} chainId - ID of the chain where the swap will be performed
 * @property {string} fromToken - Address of the token to be swapped
 * @property {string} fromAmount - Amount of tokens to be swapped (in token decimals)
 * @property {string} toToken - Address of the token to be received
 * @property {string} [toRecipient] - Address of the recipient that will receive the swapped tokens (default: sender)
 * @property {number} [slippagePercentage] - Slippage tolerance percentage
 * @property {number} [deadline] - Timestamp after which the transaction will be rejected
 * @property {boolean} [restrictComplexity] - If true, restricts complex routes to improve transaction reliability
 */
export type SwapV2Parameters = {
  chainId: number
  fromToken: string
  fromAmount: string
  toToken: string
  toRecipient?: string
  slippagePercentage?: number
  deadline?: number
  restrictComplexity?: boolean
}

/**
 * Quote information for a swap operation (v2).
 *
 * @property {string} id - Unique identifier for the quote
 * @property {number} chainId - ID of the chain where the swap will be performed
 * @property {string} fromToken - Address of the token to be swapped
 * @property {string} fromAmount - Amount of tokens to be swapped (in token decimals)
 * @property {string} toToken - Address of the token to be received
 * @property {string} expectedToAmount - Expected amount of tokens to be received (in token decimals)
 * @property {string} minToAmount - Minimum amount of tokens to be received (slippage protected) (in token decimals)
 * @property {string} routerAddress - Address of the router that performs the operation
 * @property {number} estimatedTime - Estimated time for the swap to complete in seconds
 * @property {string[]} moduleNames - Names of the modules used for the operation
 * @property {PopulatedTx} [tx] - Optional populated transaction for the operation (returned only if sender is provided)
 */
export type SwapQuoteV2 = {
  id: string
  chainId: number
  fromToken: string
  fromAmount: string
  toToken: string
  expectedToAmount: string
  minToAmount: string
  routerAddress: string
  estimatedTime: number
  moduleNames: string[]
  tx?: PopulatedTx
}
