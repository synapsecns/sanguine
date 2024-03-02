import { PAUSED_TOKENS_BY_CHAIN } from '@/constants/tokens'

export const flattenPausedTokens = () => {
  const flatList = []

  for (const chainId in PAUSED_TOKENS_BY_CHAIN) {
    if (PAUSED_TOKENS_BY_CHAIN.hasOwnProperty(chainId)) {
      PAUSED_TOKENS_BY_CHAIN[chainId].forEach((token) => {
        flatList.push(`${token}-${chainId}`)
      })
    }
  }

  return flatList
}
