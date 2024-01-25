import { Chain } from '@/utils/types'
import { CHAINS_BY_ID } from '@/constants/chains'

export const getExplorerAddressLink = (chainId: number, address: string) => {
  const chain: Chain = CHAINS_BY_ID[chainId]

  if (chain && address) {
    const addressExplorerUrl = `${chain.explorerUrl}/address/${address}`
    const explorerName = chain.explorerName

    return [addressExplorerUrl, explorerName]
  }

  console.error('getExplorerAddressLink: Chain or Address missing')

  return [null, null]
}
