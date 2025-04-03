import { BigNumber } from 'ethers'

import { FeeConfig, PopulatedTx, Query } from './misc'

/**
 * Quote information for a bridge operation.
 *
 * @property {string} id - Unique identifier for the quote
 * @property {BigNumber} feeAmount - Amount of fee to be paid (in token decimals)
 * @property {FeeConfig} feeConfig - Fee configuration for the bridge
 * @property {string} routerAddress - Address of the router that performs the operation
 * @property {BigNumber} maxAmountOut - Maximum amount of tokens that will be received (in token decimals)
 * @property {Query} originQuery - Query details for the source chain
 * @property {Query} destQuery - Query details for the destination chain
 * @property {number} estimatedTime - Estimated time for the bridge to complete in seconds
 * @property {string} bridgeModuleName - Name of the bridge module used for the operation
 * @property {BigNumber} gasDropAmount - Amount of gas to be dropped after the operation
 * @property {number} originChainId - ID of the chain where funds will be sent from
 * @property {number} destChainId - ID of the chain where funds will be received
 */
export type BridgeQuote = {
  id: string
  feeAmount: BigNumber
  feeConfig: FeeConfig
  routerAddress: string
  maxAmountOut: BigNumber
  originQuery: Query
  destQuery: Query
  estimatedTime: number
  bridgeModuleName: string
  gasDropAmount: BigNumber
  originChainId: number
  destChainId: number
}

/**
 * Parameters required to create a bridge transaction (v2).
 *
 * @property {number} fromChainId - ID of the chain where funds will be sent from
 * @property {string} fromToken - Address of the token to be sent
 * @property {string} fromAmount - Amount of tokens to be sent (in token decimals)
 * @property {string} [fromSender] - Address of the account sending the funds
 * @property {number} toChainId - ID of the chain where funds will be received
 * @property {string} toToken - Address of the token to be received
 * @property {string} [toRecipient] - Address of the recipient that will receive the funds (default: `fromSender`)
 * @property {number} [slippagePercentage] - Slippage tolerance percentage
 * @property {number} [deadline] - Timestamp after which the transaction will be rejected
 */
export type BridgeV2Parameters = {
  fromChainId: number
  fromToken: string
  fromAmount: string
  fromSender?: string
  toChainId: number
  toToken: string
  toRecipient?: string
  slippagePercentage?: number
  deadline?: number
}

/**
 * Quote information for a bridge operation (v2).
 *
 * @property {string} id - Unique identifier for the quote
 * @property {number} fromChainId - ID of the chain where funds will be sent from
 * @property {string} fromToken - Address of the token to be sent
 * @property {string} fromAmount - Amount of tokens to be sent (in token decimals)
 * @property {number} toChainId - ID of the chain where funds will be received
 * @property {string} toToken - Address of the token to be received
 * @property {string} expectedToAmount - Expected amount of tokens to be received (in token decimals)
 * @property {string} minToAmount - Minimum amount of tokens to be received (slippage protected) (in token decimals)
 * @property {string} routerAddress - Address of the router that performs the operation
 * @property {number} estimatedTime - Estimated time for the bridge to complete in seconds
 * @property {string[]} moduleNames - Names of the modules used for the operation
 * @property {string} gasDropAmount - Amount of gas to be dropped after the operation
 * @property {PopulatedTx} [tx] - Optional populated transaction for the operation (returned only if `fromSender` is provided)
 */
export type BridgeQuoteV2 = {
  id: string
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
