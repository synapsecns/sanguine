import { Zero } from '@ethersproject/constants'

import { SupportedChainId } from '../../constants'
import { EngineID, SwapEngineQuote } from './swapEngine'

const PRIORITY_NULL = 0
const PRIORITY_INEFFICIENT_QUOTES = 1
const PRIORITY_NORMAL = 2

type Priority = {
  value: number
  overrides: Partial<Record<SupportedChainId, number>>
}

const ENGINE_PRIORITY: Record<EngineID, Priority> = {
  [EngineID.Null]: { value: PRIORITY_NULL, overrides: {} },
  [EngineID.NoOp]: { value: PRIORITY_NORMAL, overrides: {} },
  [EngineID.Default]: { value: PRIORITY_NORMAL, overrides: {} },
  [EngineID.KyberSwap]: {
    value: PRIORITY_NORMAL,
    overrides: {
      [SupportedChainId.ARBITRUM]: PRIORITY_INEFFICIENT_QUOTES,
      [SupportedChainId.OPTIMISM]: PRIORITY_INEFFICIENT_QUOTES,
    },
  },
  [EngineID.ParaSwap]: { value: PRIORITY_NORMAL, overrides: {} },
  [EngineID.LiFi]: { value: PRIORITY_NORMAL, overrides: {} },
}

const getEnginePriority = (
  engineID: EngineID,
  chainId: SupportedChainId
): number => {
  return (
    ENGINE_PRIORITY[engineID].overrides[chainId] ??
    ENGINE_PRIORITY[engineID].value
  )
}

export const compareQuotesWithPriority = (
  quoteA: SwapEngineQuote,
  quoteB: SwapEngineQuote
): SwapEngineQuote => {
  const priorityA = getEnginePriority(quoteA.engineID, quoteA.chainId)
  const priorityB = getEnginePriority(quoteB.engineID, quoteB.chainId)
  // Compare priorities only if both quotes have a non-zero amountOut.
  if (
    quoteA.expectedAmountOut.gt(Zero) &&
    quoteB.expectedAmountOut.gt(Zero) &&
    priorityA !== priorityB
  ) {
    return priorityA > priorityB ? quoteA : quoteB
  }
  // Otherwise (same priority or at least one zero quote), compare amountOut.
  return quoteA.expectedAmountOut.gte(quoteB.expectedAmountOut)
    ? quoteA
    : quoteB
}
