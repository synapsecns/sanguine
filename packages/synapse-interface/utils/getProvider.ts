import { CHAINS_BY_ID } from '@/constants/chains'
import { JsonRpcProvider } from '@ethersproject/providers'
export const getProvider = (chainId: number): JsonRpcProvider => {
  const chain = CHAINS_BY_ID[chainId]
  return new JsonRpcProvider(chain.rpc)
}
