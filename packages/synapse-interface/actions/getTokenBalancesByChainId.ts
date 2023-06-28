import { getAccount } from '@wagmi/core'
import { BRIDGABLE_TOKENS } from '@/constants/tokens'
import { sortByTokenBalance } from '@/utils/sortTokens'

export const getTokensByChainId = async (chainId: number) => {
  const { address } = getAccount()

  const tokens = BRIDGABLE_TOKENS[chainId]

  return await sortByTokenBalance(tokens, chainId, address)
}
