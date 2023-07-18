import { PAUSED_TOKENS_BY_CHAIN } from '@/constants/tokens'

export const flattenPausedTokens = () => {
  let flatList = []

  for (let chainId in PAUSED_TOKENS_BY_CHAIN) {
    PAUSED_TOKENS_BY_CHAIN[chainId].forEach((token) => {
      flatList.push(`${token}-${chainId}`)
    })
  }

  return flatList
}
