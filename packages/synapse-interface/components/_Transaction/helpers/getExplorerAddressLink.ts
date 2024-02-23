import type { Chain } from '@/utils/types'
import { CHAINS_BY_ID } from '@/constants/chains'

export const getExplorerAddressLink = (chainId: number, address: string) => {
  if (!chainId || !address) {
    return [null, null]
  }

  const chain: Chain = CHAINS_BY_ID[chainId]

  const addressExplorerUrl = `${chain.explorerUrl}/address/${address}`
  const explorerName = chain.explorerName

  return [addressExplorerUrl, explorerName]
}
