import { Zero } from '@ethersproject/constants'
import { BigNumber } from 'ethers'

import { Prettify } from '../../utils'
import { EngineID, StepParams } from '../core'

export type SwapEngineQuote = {
  engineID: EngineID
  engineName: string
  chainId: number
  fromToken: string
  fromAmount: BigNumber
  toToken: string
  expectedToAmount: BigNumber
  steps?: StepParams[]
}

export type SwapEngineRoute = Prettify<
  Required<SwapEngineQuote> & {
    minToAmount?: BigNumber
  }
>

export const getEmptyQuote = (engineID: EngineID): SwapEngineQuote => {
  return {
    engineID,
    engineName: EngineID[engineID],
    chainId: 0,
    fromToken: '',
    fromAmount: Zero,
    toToken: '',
    expectedToAmount: Zero,
  }
}

export const getEmptyRoute = (engineID: EngineID): SwapEngineRoute => {
  return {
    ...getEmptyQuote(engineID),
    steps: [],
  }
}

export const sanitizeMultiStepQuote = (
  quote: SwapEngineQuote
): SwapEngineQuote => {
  if (!quote.steps || quote.steps.length <= 1) {
    return quote
  }
  return getEmptyQuote(quote.engineID)
}

export const sanitizeMultiStepRoute = (
  route: SwapEngineRoute
): SwapEngineRoute => {
  if (route.steps.length <= 1) {
    return route
  }
  return getEmptyRoute(route.engineID)
}
