import { BigNumberish } from 'ethers'

import { PopulatedTx } from './misc'

/**
 * Parameters required to create an intent.
 *
 * @property {number} fromChainId - ID of the chain where funds will be sent from
 * @property {string} fromToken - Address of the token to be sent
 * @property {BigNumberish} fromAmount - Amount of tokens to be sent (in token decimals)
 * @property {string} [fromSender] - Address of the account sending the funds
 * @property {number} toChainId - ID of the chain where funds will be received
 * @property {string} toToken - Address of the token to be received
 * @property {string} [toRecipient] - Address of the recipient that will receive the funds (default: `fromSender`)
 * @property {number} [slippagePercentage] - Slippage tolerance percentage
 * @property {number} [deadline] - Timestamp after which the transaction will be rejected
 * @property {boolean} [allowMultipleTxs] - Whether multiple transactions are allowed
 */
export type IntentParameters = {
  fromChainId: number
  fromToken: string
  fromAmount: BigNumberish
  fromSender?: string
  toChainId: number
  toToken: string
  toRecipient?: string
  slippagePercentage?: number
  deadline?: number
  allowMultipleTxs?: boolean
}

/**
 * Quote information for an intent, including pricing and routing details.
 *
 * @property {string} id - Unique identifier for the quote
 * @property {number} fromChainId - ID of the chain where funds will be sent from
 * @property {string} fromToken - Address of the token to be sent
 * @property {BigNumber} fromAmount - Amount of tokens to be sent (in token decimals)
 * @property {number} toChainId - ID of the chain where funds will be received
 * @property {string} toToken - Address of the token to be received
 * @property {BigNumber} expectedToAmount - Expected amount of tokens to be received (in token decimals)
 * @property {BigNumber} minToAmount - Minimum amount of tokens to be received (slippage protected) (in token decimals)
 * @property {number} estimatedTime - Estimated time for the intent to complete in seconds
 * @property {IntentStep[]} steps - Ordered array of steps that make up this intent
 */
export type IntentQuote = {
  id: string
  fromChainId: number
  fromToken: string
  fromAmount: string
  toChainId: number
  toToken: string
  expectedToAmount: string
  minToAmount: string
  estimatedTime: number
  steps: IntentStep[]
}

/**
 * Atomic step of an intent.
 *
 * @property {number} fromChainId - ID of the chain, where funds will be sent from.
 * @property {string} fromToken - Address of the token to be sent.
 * @property {BigNumber} fromAmount - Amount of tokens to be sent (in token decimals).
 * @property {number} toChainId - ID of the chain, where funds will be received (can be the same as `fromChainId`).
 * @property {string} toToken - Address of the token to be received.
 * @property {BigNumber} expectedToAmount - Expected amount of tokens to be received (in token decimals).
 * @property {BigNumber} minToAmount - Minimum amount of tokens to be received (slippage protected) (in token decimals).
 * @property {string} routerAddress - Address of the router that performs the operation.
 * @property {number} estimatedTime - Estimated time for the operation to complete in seconds.
 * @property {string[]} moduleNames - Names of the modules used for the operation.
 * @property {BigNumber} gasDropAmount - Amount of gas to be dropped after the operation alongside `toToken`.
 * @property {PopulatedTransaction} [tx] - Optional populated transaction for the operation (returned only if `fromSender` is provided).
 */
export type IntentStep = {
  fromChainId: number
  fromToken: string
  fromAmount: string
  toChainId: number
  toToken: string
  expectedToAmount: string
  minToAmount: string
  routerAddress: string
  estimatedTime: number
  moduleNames: string[]
  gasDropAmount: string
  tx?: PopulatedTx
}
