import { ExplorerLinks, ExplorerNames } from '../constants/explorer'

export const getExplorerAddressUrl = (chainId: number, address: string) => {
  const blockExplorer = ExplorerLinks[chainId]

  if (blockExplorer && address) {
    const explorerUrl = `${blockExplorer}/address/${address}`
    const explorerName = ExplorerNames[chainId]

    return [explorerUrl, explorerName]
  }

  console.error(
    '[Synapse Widget] Error retrieving Native Explorer Address URL: ChainId or Address missing'
  )
  return [null, null]
}
