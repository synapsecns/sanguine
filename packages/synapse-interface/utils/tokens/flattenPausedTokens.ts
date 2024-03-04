import _ from 'lodash'
import { PAUSED_TOKENS_BY_CHAIN } from '@/constants/tokens'

const FLATTENED_PAUSED_TOKENS = _(PAUSED_TOKENS_BY_CHAIN)
  .entries()
  .map(([chainId, tokens]) => tokens.map((token) => `${token}-${chainId}`))
  .flatten()
  .value()

export const flattenPausedTokens = () => {
  return FLATTENED_PAUSED_TOKENS
}

