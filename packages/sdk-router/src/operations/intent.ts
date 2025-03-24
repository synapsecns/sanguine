import { Zero } from '@ethersproject/constants'
import { BigNumber } from 'ethers'
import { uuidv7 } from 'uuidv7'

import { SynapseSDK } from '../sdk'
import { _bridgeV2Internal } from './bridge'
import { swapV2 } from './swap'
import { IntentParameters, IntentQuote, IntentStep } from '../types'
import { isSameAddress } from '../utils'

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
    estimatedTime: swapQuote.estimatedTime,
  }
  const swapStep: IntentStep = {
    ...intentCommon,
    routerAddress: swapQuote.routerAddress,
    moduleNames: swapQuote.moduleNames,
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
  // First, collect V2 quotes into either the requested token (can fallback to the bridge token if not available).
  const toRecipient = params.toRecipient || params.fromSender
  const bridgeQuotes = await _bridgeV2Internal.call(this, {
    fromChainId: params.fromChainId,
    toChainId: params.toChainId,
    fromToken: params.fromToken,
    toToken: params.toToken,
    fromAmount: params.fromAmount,
    fromSender: params.fromSender,
    toRecipient,
    slippage: params.slippage,
    deadline: params.deadline,
    allowMultipleTxs: params.allowMultipleTxs,
  })

  // Then, iterate over the quotes and add the swap step, if needed.
  const intentQuotes: IntentQuote[][] = await Promise.all(
    bridgeQuotes.map(async (bridgeQuote) => {
      const id = uuidv7()
      const intentCommon = {
        fromChainId: params.fromChainId,
        fromToken: params.fromToken,
        fromAmount: BigNumber.from(params.fromAmount),
        toChainId: params.toChainId,
      }
      const bridgeStep: IntentStep = {
        ...intentCommon,
        toToken: bridgeQuote.toToken,
        expectedToAmount: bridgeQuote.expectedToAmount,
        minToAmount: bridgeQuote.minToAmount,
        routerAddress: bridgeQuote.routerAddress,
        estimatedTime: bridgeQuote.estimatedTime,
        moduleNames: bridgeQuote.moduleNames,
        gasDropAmount: bridgeQuote.gasDropAmount,
        tx: bridgeQuote.tx,
      }
      // Check if the token out is the same as the requested token out.
      if (isSameAddress(params.toToken, bridgeQuote.toToken)) {
        const intentQuote: IntentQuote = {
          id,
          ...intentCommon,
          toToken: params.toToken,
          expectedToAmount: bridgeQuote.expectedToAmount,
          minToAmount: bridgeQuote.minToAmount,
          estimatedTime: bridgeQuote.estimatedTime,
          steps: [bridgeStep],
        }
        return [intentQuote]
      }
      // Otherwise we need to find a swap quote between the bridge token and the requested token out on the destination chain.
      const swapQuotes = await _getSameChainIntentQuotes.call(this, {
        fromChainId: params.toChainId,
        fromToken: bridgeQuote.toToken,
        fromAmount: bridgeQuote.expectedToAmount,
        fromSender: toRecipient,
        toChainId: params.toChainId,
        toToken: params.toToken,
        toRecipient,
        slippage: params.slippage,
      })
      return swapQuotes.map((swapQuote) => {
        const intentQuote: IntentQuote = {
          id,
          ...intentCommon,
          toToken: params.toToken,
          expectedToAmount: swapQuote.expectedToAmount,
          minToAmount: swapQuote.minToAmount,
          estimatedTime: bridgeQuote.estimatedTime + swapQuote.estimatedTime,
          steps: [bridgeStep, ...swapQuote.steps],
        }
        return intentQuote
      })
    })
  )
  // Flatten the quotes and sort them by expectedToAmount in descending order
  return intentQuotes
    .flat()
    .sort((a, b) => (a.expectedToAmount.gte(b.expectedToAmount) ? -1 : 1))
}
