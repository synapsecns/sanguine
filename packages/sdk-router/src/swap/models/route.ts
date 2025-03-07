import { BigNumber } from 'ethers'
import { Zero } from '@ethersproject/constants'

import { EngineID, StepParams } from '../core'

export type SwapEngineQuote = {
  engineID: EngineID
  engineName: string
  chainId: number
  expectedAmountOut: BigNumber
  steps?: StepParams[]
}

export type SwapEngineRoute = Required<SwapEngineQuote>

export const getEmptyQuote = (engineID: EngineID): SwapEngineQuote => {
  return {
    engineID,
    engineName: EngineID[engineID],
    chainId: 0,
    expectedAmountOut: Zero,
  }
}

export const getEmptyRoute = (engineID: EngineID): SwapEngineRoute => {
  return {
    engineID,
    engineName: EngineID[engineID],
    chainId: 0,
    expectedAmountOut: Zero,
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
