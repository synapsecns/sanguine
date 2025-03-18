import { Zero } from '@ethersproject/constants'
import { BigNumber, BigNumberish } from 'ethers'
import { uuidv7 } from 'uuidv7'

import { IntentStep } from '../module'
import { SynapseSDK } from '../sdk'
import { Slippage } from '../swap'
import { swapV2 } from './swap'

export type IntentParameters = {
  fromChainId: number
  fromToken: string
  fromAmount: BigNumberish
  fromSender?: string
  toChainId: number
  toToken: string
  toRecipient?: string
  slippage?: Slippage
  deadline?: number
  allowMultipleTxs?: boolean
}

export type IntentQuote = {
  id: string
  fromChainId: number
  fromToken: string
  fromAmount: BigNumber
  toChainId: number
  toToken: string
  expectedToAmount: BigNumber
  minToAmount: BigNumber
  estimatedTime: number
  steps: IntentStep[]
}

export async function intent(
  this: SynapseSDK,
  params: IntentParameters
): Promise<IntentQuote[]> {
  return params.fromChainId === params.toChainId
    ? _getSameChainIntentQuotes.call(this, params)
    : _getCrossChainIntentQuotes.call(this, params)
}

async function _getSameChainIntentQuotes(
  this: SynapseSDK,
  params: IntentParameters
): Promise<IntentQuote[]> {
  const swapQuote = await swapV2.call(this, {
    chainId: params.fromChainId,
    fromToken: params.fromToken,
    toToken: params.toToken,
    fromAmount: params.fromAmount,
    toRecipient: params.toRecipient || params.fromSender,
    slippage: params.slippage,
    deadline: params.deadline,
  })
  if (!swapQuote || swapQuote.expectedToAmount.isZero()) {
    return []
  }
  const intentCommon = {
    fromChainId: params.fromChainId,
    fromToken: params.fromToken,
    fromAmount: BigNumber.from(params.fromAmount),
    toChainId: params.toChainId,
    toToken: params.toToken,
    expectedToAmount: swapQuote.expectedToAmount,
    minToAmount: swapQuote.minToAmount,
    // TODO: chain block time
    estimatedTime: 0,
  }
  const swapStep: IntentStep = {
    ...intentCommon,
    routerAddress: swapQuote.routerAddress,
    moduleName: swapQuote.moduleName,
    gasDropAmount: Zero,
    tx: swapQuote.tx,
  }
  const intentQuote: IntentQuote = {
    id: uuidv7(),
    ...intentCommon,
    steps: [swapStep],
  }
  // TODO: do we need to return multiple quotes?
  return [intentQuote]
}

async function _getCrossChainIntentQuotes(
  this: SynapseSDK,
  params: IntentParameters
): Promise<IntentQuote[]> {
  // TODO: implement
}
