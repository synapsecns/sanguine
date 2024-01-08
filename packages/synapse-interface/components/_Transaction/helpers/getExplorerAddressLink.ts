import { ExplorerLinks, ExplorerNames } from './constants'

export const getExplorerAddressLink = (chainId: number, address: string) => {
  const blockExplorer = ExplorerLinks[chainId]

  if (blockExplorer && address) {
    const explorerUrl = `${blockExplorer}/address/${address}`
    const explorerName = ExplorerNames[chainId]

    return [explorerUrl, explorerName]
  }

  console.error('getExplorerAddressLink: ChainId or Address missing')
  return [null, null]
}
