import { Chain } from '@/utils/types'
import { CHAINS_BY_ID } from '@/constants/chains'

export const getTxBlockExplorerLink = (chainId: number, txHash: string) => {
  const chain: Chain = CHAINS_BY_ID[chainId]

  if (chain && txHash) {
    const txExplorerUrl = `${chain.explorerUrl}/tx/${txHash}`
    const explorerName = chain.explorerName

    return [txExplorerUrl, explorerName]
  }

  console.error('getTxBlockExplorerLink: Chain or Transaction Hash missing')
  return [null, null]
}
