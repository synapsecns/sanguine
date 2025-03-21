import { Zero } from '@ethersproject/constants'

import { SupportedChainId } from '../../constants'
import { EngineID } from '../core'
import { SwapEngineQuote } from './route'

enum Priority {
  Null,
  InefficientQuotes,
  Normal,
}

type PriorityConfig = {
  value: Priority
  overrides?: Partial<Record<SupportedChainId, Priority>>
}

const ENGINE_PRIORITY: Record<EngineID, PriorityConfig> = {
  [EngineID.Null]: { value: Priority.Null },
  [EngineID.NoOp]: { value: Priority.Normal },
  [EngineID.Default]: { value: Priority.InefficientQuotes },
  [EngineID.KyberSwap]: { value: Priority.Normal },
  [EngineID.ParaSwap]: { value: Priority.Normal },
  [EngineID.LiFi]: { value: Priority.InefficientQuotes },
}

const getEnginePriority = (
  engineID: EngineID,
  chainId: SupportedChainId
): Priority => {
  const config = ENGINE_PRIORITY[engineID]
  return config.overrides?.[chainId] ?? config.value
}

export const compareQuotesWithPriority = (
  quoteA: SwapEngineQuote,
  quoteB: SwapEngineQuote
): SwapEngineQuote => {
  const bothQuotesNonZero =
    quoteA.expectedToAmount.gt(Zero) && quoteB.expectedToAmount.gt(Zero)

  if (bothQuotesNonZero) {
    const priorityA = getEnginePriority(quoteA.engineID, quoteA.chainId)
    const priorityB = getEnginePriority(quoteB.engineID, quoteB.chainId)

    if (priorityA !== priorityB) {
      return priorityA > priorityB ? quoteA : quoteB
    }
  }

  // Same priority or at least one zero quote, compare expectedToAmount
  return quoteA.expectedToAmount.gte(quoteB.expectedToAmount) ? quoteA : quoteB
}
